// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/snlansky/glibs/logging"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const cmdRoot = "api"

var (
	logger  = logging.MustGetLogger(cmdRoot)
	cfgFile = ""
	rootCmd = &cobra.Command{
		Use:   "api",
		Short: "api",
		Run:   start,
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config.yaml)")

	cobra.OnInitialize(initConfig)
	initializeLogging()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Fatal(err)
	}
}

func initializeLogging() {
	loggingSpec := os.Getenv("ACCESS_LOGGING_SPEC")
	loggingFormat := os.Getenv("ACCESS_LOGGING_FORMAT")
	logging.Init(logging.Config{
		Format:  loggingFormat,
		Writer:  os.Stderr,
		LogSpec: loggingSpec,
	})
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetEnvPrefix(cmdRoot)
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AddConfigPath("./")
	viper.SetConfigType("yaml")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else if os.Getenv("CONFIG_FILE") != "" {
		viper.SetConfigFile(os.Getenv("CONFIG_FILE"))
	} else {
		home, err := homedir.Dir()
		if err != nil {
			logger.Fatal(err)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName("config-local")
	}

	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Using config file:%s\n", viper.ConfigFileUsed())
}
