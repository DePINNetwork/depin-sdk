package genutil_test

import (
	"os"
	"path/filepath"
	"testing"

	addresscodec "github.com/depinnetwork/depin-sdk/codec/address"
	sdk "github.com/depinnetwork/depin-sdk/types"
	"github.com/depinnetwork/depin-sdk/x/genutil"
	"github.com/depinnetwork/depin-sdk/x/genutil/types"
)

// Ensures that CollectTx correctly traverses directories and won't error out on encountering
// a directory during traversal of the first level. See issue https://github.com/cosmos/cosmos-sdk/issues/6788.
func TestCollectTxsHandlesDirectories(t *testing.T) {
	testDir := t.TempDir()

	// 1. We'll insert a directory as the first element before JSON file.
	subDirPath := filepath.Join(testDir, "_adir")
	if err := os.MkdirAll(subDirPath, 0o755); err != nil {
		t.Fatal(err)
	}

	txDecoder := sdk.TxDecoder(func(txBytes []byte) (sdk.Tx, error) {
		return nil, nil
	})

	// 2. Ensure that we don't encounter any error traversing the directory.
	genesis := &types.AppGenesis{AppState: []byte("{}")}

	if _, _, err := genutil.CollectTxs(txDecoder, "foo", testDir, genesis, types.DefaultMessageValidator,
		addresscodec.NewBech32Codec("cosmosvaloper"), addresscodec.NewBech32Codec("cosmos")); err != nil {
		t.Fatal(err)
	}
}
