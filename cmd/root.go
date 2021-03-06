/*
Copyright © 2021 Sascha Andres <sascha.andres@outlook.com>

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

var rootCmd = &cobra.Command{
	Use:   "dupfinder",
	Short: "a duplicate file finder",
	Long: `scan the current directory and subdirectory for duplicate files.

Method used: calculate hash for each file and compare to existing files`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := createRootLogger()
		worker := collectData(logger)
		generateReport(logger, worker)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dupfinder.yaml)")

	rootCmd.Flags().BoolP("verbose", "v", false, "turn on verbose logging")
	rootCmd.Flags().StringP("reporter", "r", "console", "select reporter")
	_ = viper.BindPFlag("logging.verbose", rootCmd.Flags().Lookup("verbose"))
	_ = viper.BindPFlag("reporter.type", rootCmd.Flags().Lookup("reporter"))
}
