// Copyright (C) 2019-2025 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/DePINNetwork/avm-abi/apps"
	cmdutil "github.com/DePINNetwork/depin-sdk/cmd/util"
	"github.com/DePINNetwork/depin-sdk/config"
	"github.com/DePINNetwork/depin-sdk/crypto"
	"github.com/DePINNetwork/depin-sdk/data/basics"
	"github.com/DePINNetwork/depin-sdk/data/bookkeeping"
	"github.com/DePINNetwork/depin-sdk/ledger"
	"github.com/DePINNetwork/depin-sdk/ledger/ledgercore"
	"github.com/DePINNetwork/depin-sdk/ledger/store/trackerdb/sqlitedriver"
	"github.com/DePINNetwork/depin-sdk/logging"
	"github.com/DePINNetwork/depin-sdk/protocol"
	"github.com/DePINNetwork/depin-sdk/util/db"
)

var catchpointFile string
var outFileName string
var excludedFields = cmdutil.MakeCobraStringSliceValue(nil, []string{"version", "catchpoint"})
var printDigests bool

func init() {
	fileCmd.Flags().StringVarP(&catchpointFile, "tar", "t", "", "Specify the catchpoint file (either .tar or .tar.gz) to process")
	fileCmd.Flags().StringVarP(&outFileName, "output", "o", "", "Specify an outfile for the dump ( i.e. tracker.dump.txt )")
	fileCmd.Flags().BoolVarP(&loadOnly, "load", "l", false, "Load only, do not dump")
	fileCmd.Flags().BoolVarP(&printDigests, "digest", "d", false, "Print balances and spver digests")
	fileCmd.Flags().VarP(excludedFields, "exclude-fields", "e", "List of fields to exclude from the dump: ["+excludedFields.AllowedString()+"]")
}

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "Specify a file to dump",
	Long:  "Specify a file to dump",
	Args:  validateNoPosArgsFn,
	Run: func(cmd *cobra.Command, args []string) {
		if catchpointFile == "" {
			cmd.HelpFunc()(cmd, args)
			return
		}
		stats, err := os.Stat(catchpointFile)
		if err != nil {
			reportErrorf("Unable to stat '%s' : %v", catchpointFile, err)
		}

		catchpointSize := stats.Size()
		if catchpointSize == 0 {
			reportErrorf("Empty file '%s' : %v", catchpointFile, err)
		}
		// TODO: store CurrentProtocol in catchpoint file header.
		// As a temporary workaround use a current protocol version.
		genesisInitState := ledgercore.InitState{
			Block: bookkeeping.Block{BlockHeader: bookkeeping.BlockHeader{
				UpgradeState: bookkeeping.UpgradeState{
					CurrentProtocol: protocol.ConsensusCurrentVersion,
				},
			}},
		}
		cfg := config.GetDefaultLocal()
		l, err := ledger.OpenLedger(logging.Base(), "./ledger", false, genesisInitState, cfg)
		if err != nil {
			reportErrorf("Unable to open ledger : %v", err)
		}

		defer os.Remove("./ledger.block.sqlite")
		defer os.Remove("./ledger.block.sqlite-shm")
		defer os.Remove("./ledger.block.sqlite-wal")
		if !loadOnly {
			defer os.Remove("./ledger.tracker.sqlite")
			defer os.Remove("./ledger.tracker.sqlite-shm")
			defer os.Remove("./ledger.tracker.sqlite-wal")
		}
		defer l.Close()

		catchupAccessor := ledger.MakeCatchpointCatchupAccessor(l, logging.Base())
		err = catchupAccessor.ResetStagingBalances(context.Background(), true)
		if err != nil {
			reportErrorf("Unable to initialize catchup database : %v", err)
		}
		var fileHeader ledger.CatchpointFileHeader

		reader, err := os.Open(catchpointFile)
		if err != nil {
			reportErrorf("Unable to read '%s' : %v", catchpointFile, err)
		}
		defer reader.Close()

		fileHeader, err = loadCatchpointIntoDatabase(context.Background(), catchupAccessor, reader, catchpointSize)
		if err != nil {
			reportErrorf("Unable to load catchpoint file into in-memory database : %v", err)
		}

		if !loadOnly {
			outFile := os.Stdout
			if outFileName != "" {
				outFile, err = os.OpenFile(outFileName, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0755)
				if err != nil {
					reportErrorf("Unable to create file '%s' : %v", outFileName, err)
				}
				defer outFile.Close()
			}
			err = printAccountsDatabase("./ledger.tracker.sqlite", true, fileHeader, outFile, excludedFields.GetSlice())
			if err != nil {
				reportErrorf("Unable to print account database : %v", err)
			}
			err = printKeyValueStore("./ledger.tracker.sqlite", true, outFile)
			if err != nil {
				reportErrorf("Unable to print key value store : %v", err)
			}
			err = printStateProofVerificationContext("./ledger.tracker.sqlite", true, outFile)
			if err != nil {
				reportErrorf("Unable to print state proof verification database : %v", err)
			}
			err = printOnlineAccounts("./ledger.tracker.sqlite", true, outFile)
			if err != nil {
				reportErrorf("Unable to print online accounts : %v", err)
			}
			err = printOnlineRoundParams("./ledger.tracker.sqlite", true, outFile)
			if err != nil {
				reportErrorf("Unable to print online round params : %v", err)
			}
		}
	},
}

func printLoadCatchpointProgressLine(progress int, barLength int, dld int64) {
	if barLength == 0 {
		fmt.Printf(escapeCursorUp + escapeDeleteLine + "[ Done ] Loaded\n")
		return
	}

	outString := "[" + strings.Repeat(escapeSquare, progress) + strings.Repeat(escapeDot, barLength-progress) + "] Loading..."
	fmt.Printf(escapeCursorUp+escapeDeleteLine+outString+" %s\n", formatSize(dld))
}

func isGzipCompressed(catchpointReader *bufio.Reader, catchpointFileSize int64) bool {
	const gzipPrefixSize = 2
	const gzipPrefix = "\x1F\x8B"

	if catchpointFileSize < gzipPrefixSize {
		return false
	}

	prefixBytes, err := catchpointReader.Peek(gzipPrefixSize)

	if err != nil {
		return false
	}

	return prefixBytes[0] == gzipPrefix[0] && prefixBytes[1] == gzipPrefix[1]
}

func getCatchpointTarReader(catchpointReader *bufio.Reader, catchpointFileSize int64) (*tar.Reader, bool, error) {
	if isGzipCompressed(catchpointReader, catchpointFileSize) {
		gzipReader, err := gzip.NewReader(catchpointReader)
		if err != nil {
			return nil, false, err
		}
		return tar.NewReader(gzipReader), true, nil
	}

	return tar.NewReader(catchpointReader), false, nil
}

func loadCatchpointIntoDatabase(ctx context.Context, catchupAccessor ledger.CatchpointCatchupAccessor, catchpointFile io.Reader, catchpointFileSize int64) (fileHeader ledger.CatchpointFileHeader, err error) {
	fmt.Printf("\n")
	const barLength = 50
	printLoadCatchpointProgressLine(0, barLength, 0)
	lastProgressUpdate := time.Now()
	progress := uint64(0)
	defer printLoadCatchpointProgressLine(0, 0, 0)

	catchpointReader := bufio.NewReader(catchpointFile)
	tarReader, isCompressed, err := getCatchpointTarReader(catchpointReader, catchpointFileSize)
	if err != nil {
		return fileHeader, err
	}
	if isCompressed {
		// gzip'ed file is about 3-6 times smaller than tar
		// modify catchpointFileSize to make the progress bar more-less reflecting the state
		catchpointFileSize = 4 * catchpointFileSize
	}

	var downloadProgress ledger.CatchpointCatchupAccessorProgress
	for {
		header, err := tarReader.Next()
		if err != nil {
			if err == io.EOF {
				if printDigests {
					err = catchupAccessor.BuildMerkleTrie(ctx, func(uint64, uint64) {})
					if err != nil {
						return fileHeader, err
					}
					var balanceHash, spverHash, onlineAccountsHash, onlineRoundParamsHash crypto.Digest
					balanceHash, spverHash, onlineAccountsHash, onlineRoundParamsHash, _, err = catchupAccessor.GetVerifyData(ctx)
					if err != nil {
						return fileHeader, err
					}
					fmt.Printf("accounts digest=%s, spver digest=%s, onlineaccounts digest=%s onlineroundparams digest=%s\n",
						balanceHash, spverHash, onlineAccountsHash, onlineRoundParamsHash)

					fmt.Printf("Catchpoint label: %s\n", fileHeader.Catchpoint)
					// make v7 label
					v7Label := ledgercore.MakeCatchpointLabelMakerV7(fileHeader.BlocksRound, &fileHeader.BlockHeaderDigest, &balanceHash, fileHeader.Totals, &spverHash)
					fmt.Printf("catchpoint v7 label: %s\n", ledgercore.MakeLabel(v7Label))

					// make v8 label (current)
					v8Label := ledgercore.MakeCatchpointLabelMakerCurrent(fileHeader.BlocksRound, &fileHeader.BlockHeaderDigest, &balanceHash, fileHeader.Totals, &spverHash, &onlineAccountsHash, &onlineRoundParamsHash)
					fmt.Printf("catchpoint v8 label: %s\n\n", ledgercore.MakeLabel(v8Label))
				}
				return fileHeader, nil
			}
			return fileHeader, err
		}
		balancesBlockBytes := make([]byte, header.Size)
		readComplete := int64(0)

		for readComplete < header.Size {
			bytesRead, err := tarReader.Read(balancesBlockBytes[readComplete:])
			readComplete += int64(bytesRead)
			progress += uint64(bytesRead)
			if err != nil {
				if err == io.EOF {
					if readComplete == header.Size {
						break
					}
					err = fmt.Errorf("getPeerLedger received io.EOF while reading from tar file stream prior of reaching chunk size %d / %d", readComplete, header.Size)
				}
				return fileHeader, err
			}
		}
		err = catchupAccessor.ProcessStagingBalances(ctx, header.Name, balancesBlockBytes, &downloadProgress)
		if err != nil {
			return fileHeader, err
		}
		if header.Name == ledger.CatchpointContentFileName {
			// we already know it's valid, since we validated that above.
			protocol.Decode(balancesBlockBytes, &fileHeader)
		}
		if time.Since(lastProgressUpdate) > 50*time.Millisecond && catchpointFileSize > 0 {
			lastProgressUpdate = time.Now()
			progressRatio := int(float64(progress) * barLength / float64(catchpointFileSize))
			if progressRatio > barLength {
				progressRatio = barLength
			}
			printLoadCatchpointProgressLine(progressRatio, barLength, int64(progress))
		}
	}
}

func printDumpingCatchpointProgressLine(progress int, barLength int, dld int64) {
	if barLength == 0 {
		fmt.Printf(escapeCursorUp + escapeDeleteLine + "[ Done ] Dumped\n")
		return
	}

	outString := "[" + strings.Repeat(escapeSquare, progress) + strings.Repeat(escapeDot, barLength-progress) + "] Dumping..."
	if dld > 0 {
		outString = fmt.Sprintf(outString+" %d", dld)
	}
	fmt.Printf(escapeCursorUp + escapeDeleteLine + outString + "\n")
}

func printAccountsDatabase(databaseName string, stagingTables bool, fileHeader ledger.CatchpointFileHeader, outFile *os.File, excludeFields []string) error {
	lastProgressUpdate := time.Now()
	progress := uint64(0)
	defer printDumpingCatchpointProgressLine(0, 0, 0)

	fileWriter := bufio.NewWriterSize(outFile, 1024*1024)
	defer fileWriter.Flush()

	dbAccessor, err := db.MakeAccessor(databaseName, true, false)
	if err != nil || dbAccessor.Handle == nil {
		return err
	}
	if fileHeader.Version != 0 {
		var headerFields = []string{
			"Version: %d",
			"Balances Round: %d",
			"Block Round: %d",
			"Block Header Digest: %s",
			"Catchpoint: %s",
			"Total Accounts: %d",
			"Total KVs: %d",
			"Total Online Accounts: %d",
			"Total Online Round Params: %d",
			"Total Chunks: %d",
		}
		var headerValues = []interface{}{
			fileHeader.Version,
			fileHeader.BalancesRound,
			fileHeader.BlocksRound,
			fileHeader.BlockHeaderDigest.String(),
			fileHeader.Catchpoint,
			fileHeader.TotalAccounts,
			fileHeader.TotalKVs,
			fileHeader.TotalOnlineAccounts,
			fileHeader.TotalOnlineRoundParams,
			fileHeader.TotalChunks,
		}
		// safety check
		if len(headerFields) != len(headerValues) {
			return fmt.Errorf("printing failed: header formatting mismatch")
		}

		var actualFields []string
		var actualValues []interface{}
		if len(excludeFields) == 0 {
			actualFields = headerFields
			actualValues = headerValues
		} else {
			actualFields = make([]string, 0, len(headerFields)-len(excludeFields))
			actualValues = make([]interface{}, 0, len(headerFields)-len(excludeFields))
			for i, field := range headerFields {
				lower := strings.ToLower(field)
				excluded := false
				for _, filter := range excludeFields {
					if strings.HasPrefix(lower, filter) {
						excluded = true
						break
					}
				}
				if !excluded {
					actualFields = append(actualFields, field)
					actualValues = append(actualValues, headerValues[i])
				}
			}
		}

		fmt.Fprintf(fileWriter, strings.Join(actualFields, "\n")+"\n", actualValues...)

		totals := fileHeader.Totals
		fmt.Fprintf(fileWriter, "AccountTotals - Online Money: %d\nAccountTotals - Online RewardUnits : %d\nAccountTotals - Offline Money: %d\nAccountTotals - Offline RewardUnits : %d\nAccountTotals - Not Participating Money: %d\nAccountTotals - Not Participating Money RewardUnits: %d\nAccountTotals - Rewards Level: %d\n",
			totals.Online.Money.Raw, totals.Online.RewardUnits,
			totals.Offline.Money.Raw, totals.Offline.RewardUnits,
			totals.NotParticipating.Money.Raw, totals.NotParticipating.RewardUnits,
			totals.RewardsLevel)
	}
	return dbAccessor.Atomic(func(ctx context.Context, tx *sql.Tx) (err error) {
		arw := sqlitedriver.NewAccountsSQLReaderWriter(tx)

		fmt.Printf("\n")
		printDumpingCatchpointProgressLine(0, 50, 0)

		if fileHeader.Version == 0 {
			var totals ledgercore.AccountTotals
			id := ""
			if stagingTables {
				id = "catchpointStaging"
			}
			row := tx.QueryRow("SELECT online, onlinerewardunits, offline, offlinerewardunits, notparticipating, notparticipatingrewardunits, rewardslevel FROM accounttotals WHERE id=?", id)
			err = row.Scan(&totals.Online.Money.Raw, &totals.Online.RewardUnits,
				&totals.Offline.Money.Raw, &totals.Offline.RewardUnits,
				&totals.NotParticipating.Money.Raw, &totals.NotParticipating.RewardUnits,
				&totals.RewardsLevel)
			if err != nil {
				return err
			}
			fmt.Fprintf(fileWriter, "AccountTotals - Online Money: %d\nAccountTotals - Online RewardUnits : %d\nAccountTotals - Offline Money: %d\nAccountTotals - Offline RewardUnits : %d\nAccountTotals - Not Participating Money: %d\nAccountTotals - Not Participating Money RewardUnits: %d\nAccountTotals - Rewards Level: %d\n",
				totals.Online.Money.Raw, totals.Online.RewardUnits,
				totals.Offline.Money.Raw, totals.Offline.RewardUnits,
				totals.NotParticipating.Money.Raw, totals.NotParticipating.RewardUnits,
				totals.RewardsLevel)
		}

		balancesTable := "accountbase"
		resourcesTable := "resources"
		if stagingTables {
			balancesTable = "catchpointbalances"
			resourcesTable = "catchpointresources"
		}

		var rowsCount int64
		err = tx.QueryRow(fmt.Sprintf("SELECT count(*) from %s", balancesTable)).Scan(&rowsCount)
		if err != nil {
			return
		}

		printer := func(addr basics.Address, data interface{}, progress uint64) (err error) {
			jsonData, err := json.Marshal(data)
			if err != nil {
				return err
			}

			fmt.Fprintf(fileWriter, "%v : %s\n", addr, string(jsonData))

			if time.Since(lastProgressUpdate) > 50*time.Millisecond && rowsCount > 0 {
				lastProgressUpdate = time.Now()
				printDumpingCatchpointProgressLine(int(float64(progress)*50.0/float64(rowsCount)), 50, int64(progress))
			}
			return nil
		}

		if fileHeader.Version != 0 && fileHeader.Version < ledger.CatchpointFileVersionV6 {
			var rows *sql.Rows
			rows, err = tx.Query(fmt.Sprintf("SELECT address, data FROM %s order by address", balancesTable))
			if err != nil {
				return
			}
			defer rows.Close()

			for rows.Next() {
				var addrbuf []byte
				var buf []byte
				err = rows.Scan(&addrbuf, &buf)
				if err != nil {
					return
				}

				var addr basics.Address
				if len(addrbuf) != len(addr) {
					err = fmt.Errorf("account DB address length mismatch: %d != %d", len(addrbuf), len(addr))
					return
				}
				copy(addr[:], addrbuf)

				var data basics.AccountData
				err = protocol.Decode(buf, &data)
				if err != nil {
					return
				}

				err = printer(addr, data, progress)
				if err != nil {
					return
				}

				progress++
			}
			err = rows.Err()
		} else {
			acctCount := 0
			acctCb := func(addr basics.Address, data basics.AccountData) {
				err = printer(addr, data, progress)
				if err != nil {
					return
				}
				progress++
				acctCount++
			}
			_, err = arw.LoadAllFullAccounts(context.Background(), balancesTable, resourcesTable, acctCb)
			if err != nil {
				return
			}
			if acctCount != int(rowsCount) {
				return fmt.Errorf("expected %d accounts but got only %d", rowsCount, acctCount)
			}
		}
		// increase the deadline warning to disable the warning message.
		_, _ = db.ResetTransactionWarnDeadline(ctx, tx, time.Now().Add(5*time.Second))
		return err
	})
}

func printStateProofVerificationContext(databaseName string, stagingTables bool, outFile *os.File) error {
	fileWriter := bufio.NewWriterSize(outFile, 1024*1024)
	defer fileWriter.Flush()

	dbAccessor, err := db.MakeAccessor(databaseName, true, false)
	if err != nil || dbAccessor.Handle == nil {
		return err
	}
	defer dbAccessor.Close()

	var stateProofVerificationContext []ledgercore.StateProofVerificationContext
	err = dbAccessor.Atomic(func(ctx context.Context, tx *sql.Tx) (err error) {
		if stagingTables {
			stateProofVerificationContext, err = sqlitedriver.MakeStateProofVerificationReader(tx).GetAllSPContextsFromCatchpointTbl(ctx)
		} else {
			stateProofVerificationContext, err = sqlitedriver.MakeStateProofVerificationReader(tx).GetAllSPContexts(ctx)
		}
		return err
	})

	if err != nil {
		return err
	}

	fmt.Fprintf(fileWriter, "State Proof Verification Data:\n")
	for _, ctx := range stateProofVerificationContext {
		jsonData, err := json.Marshal(ctx)
		if err != nil {
			return err
		}
		fmt.Fprintf(fileWriter, "%d : %s\n", ctx.LastAttestedRound, string(jsonData))
	}

	return nil
}

func printKeyValue(writer *bufio.Writer, key, value []byte) {
	var pretty string
	ai, rest, err := apps.SplitBoxKey(string(key))
	if err == nil {
		pretty = fmt.Sprintf("box(%d, %s)", ai, base64.StdEncoding.EncodeToString([]byte(rest)))
	} else {
		pretty = base64.StdEncoding.EncodeToString(key)
	}

	fmt.Fprintf(writer, "%s : %v\n", pretty, base64.StdEncoding.EncodeToString(value))
}

func printKeyValueStore(databaseName string, stagingTables bool, outFile *os.File) error {
	printDumpingCatchpointProgressLine(0, 50, 0)
	lastProgressUpdate := time.Now()
	progress := uint64(0)
	defer printDumpingCatchpointProgressLine(0, 0, 0)

	fileWriter := bufio.NewWriterSize(outFile, 1024*1024)
	defer fileWriter.Flush()

	dbAccessor, err := db.MakeAccessor(databaseName, true, false)
	if err != nil || dbAccessor.Handle == nil {
		return err
	}

	kvTable := "kvstore"
	if stagingTables {
		kvTable = "catchpointkvstore"
	}

	return dbAccessor.Atomic(func(ctx context.Context, tx *sql.Tx) error {
		var rowsCount int64
		err := tx.QueryRow(fmt.Sprintf("SELECT count(*) from %s", kvTable)).Scan(&rowsCount)
		if err != nil {
			return err
		}

		// ordered to make dumps more "diffable"
		rows, err := tx.Query(fmt.Sprintf("SELECT key, value FROM %s order by key", kvTable))
		if err != nil {
			return err
		}
		defer rows.Close()
		for rows.Next() {
			progress++
			var key []byte
			var value []byte
			err := rows.Scan(&key, &value)
			if err != nil {
				return err
			}
			printKeyValue(fileWriter, key, value)
			if time.Since(lastProgressUpdate) > 50*time.Millisecond {
				lastProgressUpdate = time.Now()
				printDumpingCatchpointProgressLine(int(float64(progress)*50.0/float64(rowsCount)), 50, int64(progress))
			}
		}
		return nil
	})
}

func printOnlineAccounts(databaseName string, stagingTables bool, outFile *os.File) error {
	fileWriter := bufio.NewWriterSize(outFile, 1024*1024)
	defer fileWriter.Flush()

	dbAccessor, err := db.MakeAccessor(databaseName, true, false)
	if err != nil || dbAccessor.Handle == nil {
		return err
	}

	return dbAccessor.Atomic(func(ctx context.Context, tx *sql.Tx) error {
		rows, err := sqlitedriver.MakeOrderedOnlineAccountsIter(ctx, tx, stagingTables, 0)
		if err != nil {
			return err
		}
		defer rows.Close()
		for rows.Next() {
			row, err := rows.GetItem()
			if err != nil {
				return err
			}
			jsonData, err := json.Marshal(row)
			if err != nil {
				return err
			}

			fmt.Fprintf(fileWriter, "onlineaccount: %s\n", string(jsonData))
		}
		return nil
	})
}

func printOnlineRoundParams(databaseName string, stagingTables bool, outFile *os.File) error {
	fileWriter := bufio.NewWriterSize(outFile, 1024*1024)
	defer fileWriter.Flush()

	dbAccessor, err := db.MakeAccessor(databaseName, true, false)
	if err != nil || dbAccessor.Handle == nil {
		return err
	}

	return dbAccessor.Atomic(func(ctx context.Context, tx *sql.Tx) error {
		rows, err := sqlitedriver.MakeOnlineRoundParamsIter(ctx, tx, stagingTables, 0)
		if err != nil {
			return err
		}
		defer rows.Close()
		for rows.Next() {
			row, err := rows.GetItem()
			if err != nil {
				return err
			}
			jsonData, err := json.Marshal(row)
			if err != nil {
				return err
			}

			fmt.Fprintf(fileWriter, "onlineroundparams: %s\n", string(jsonData))
		}
		return nil
	})
}
