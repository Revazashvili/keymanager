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
	"github.com/spf13/cobra"
	"os"

	"github.com/spf13/viper"
)

var cfgFile string

var logo = " _  __             __  __                                         \n| |/ / ___  _   _ |  \\/  |  __ _  _ __    __ _   __ _   ___  _ __ \n| ' / / _ \\| | | || |\\/| | / _` || '_ \\  / _` | / _` | / _ \\| '__|\n| . \\|  __/| |_| || |  | || (_| || | | || (_| || (_| ||  __/| |   \n|_|\\_\\\\___| \\__, ||_|  |_| \\__,_||_| |_| \\__,_| \\__, | \\___||_|   \n            |___/                               |___/             "

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "keymanager",
	Short: "app to manage keys",
	Long:  `Application to manage keys in CIB database. generate scripts for production environment from test database.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(logo)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.keymanager.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".keymanager" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".keymanager")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
