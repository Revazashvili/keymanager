/*
Copyright © 2022 levanrevazashvili@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	keymanager "github.com/Revazashvili/keymanager/core"
	"github.com/Revazashvili/keymanager/generators"
	"github.com/spf13/cobra"
	"io/fs"
	"io/ioutil"
	"os"
)

const (
	prefix            = "prefix"
	prefixShortHand   = "p"
	name              = "name"
	nameShortHand     = "n"
	location          = "location"
	locationShortHand = "l"
)

var storage = keymanager.NewOracleStorage()
var dictKeyGen = generators.NewDictionaryKeyGenerator()
var dictGen = generators.NewDictionaryGenerator()

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "command to generate queries.",
	Long:  `command to generate queries to insert test environment ready data into production database for tables CIB_DICTIONARY_KEYS and CIB_DICTIONARY.`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := cmd.Flags().GetString(prefix) //TODO: use this prefix to retrieve data from database
		keymanager.CheckErr(err)

		l := getLocation(cmd)

		dictKeyInsertModel, err := storage.GetDictionaryKeys(sw)
		keymanager.CheckErr(err)

		r, err := dictKeyGen.Generate(dictKeyInsertModel)
		keymanager.CheckErr(err)

		dictInsertModel, err := storage.GetDictionaries(sw)
		keymanager.CheckErr(err)

		r2, err := dictGen.Generate(dictInsertModel)
		keymanager.CheckErr(err)

		err = ioutil.WriteFile(l, []byte(r+"\n"+r2), fs.ModePerm)
		keymanager.CheckErr(err)

		fmt.Printf("file generated successfully. see %s.sql at location %s \n", sw, l)
	},
}

func getLocation(cmd *cobra.Command) string {
	n, err := cmd.Flags().GetString(name)
	keymanager.CheckErr(err)
	if len(n) < 1 {
		sw, _ := cmd.Flags().GetString(prefix)
		n = sw
	}
	l, err := cmd.Flags().GetString(location)
	keymanager.CheckErr(err)
	if len(l) < 1 {
		home, err := os.UserHomeDir()
		keymanager.CheckErr(err)
		l = home
	}
	return fmt.Sprintf("%s/%s.sql", l, n)
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP(prefix, prefixShortHand, "", "prefix for keys to search in database.")
	generateCmd.Flags().StringP(name, nameShortHand, "", "name for generated file. default will be prefix.")
	generateCmd.Flags().StringP(location, locationShortHand, "", "location for generated file. default will be $HOME/{name}.sql")
}
