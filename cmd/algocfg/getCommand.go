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
	"fmt"
	"os"
	"reflect"

	"github.com/spf13/cobra"

	"github.com/DePINNetwork/depin-sdk/cmd/util/datadir"
	"github.com/DePINNetwork/depin-sdk/config"
)

var (
	getParameterArg string
)

func init() {
	getCmd.Flags().StringVarP(&getParameterArg, "parameter", "p", "", "Parameter to query")
	getCmd.MarkFlagRequired("parameter")

	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve the current value for the specified parameter",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {
		anyError := false
		datadir.OnDataDirs(func(dataDir string) {
			cfg, err := config.LoadConfigFromDisk(dataDir)
			if err != nil && !os.IsNotExist(err) {
				reportWarnf("Error loading config file from '%s' - %s", dataDir, err)
				anyError = true
				return
			}

			val, err := serializeObjectProperty(cfg, getParameterArg)
			if err != nil {
				reportWarnf("Error retrieving property '%s' - %s", getParameterArg, err)
				anyError = true
				return
			}

			fmt.Print(val)
		})
		if anyError {
			os.Exit(1)
		}
	},
}

func serializeObjectProperty(object interface{}, property string) (ret string, err error) {
	v := reflect.ValueOf(object)
	val := reflect.Indirect(v)
	f := val.FieldByName(property)

	if !f.IsValid() {
		return "", fmt.Errorf("unknown property named '%s'", property)
	}

	return fmt.Sprintf("%v", f.Interface()), nil
}
