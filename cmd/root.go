/*
Copyright © 2025 Steven A. Zaluk <arcanegame@protonmail.com>

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
	"github.com/mitchellh/go-homedir"
	slogmulti "github.com/samber/slog-multi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stevezaluk/mtgjson-sdk-client/config"
	"log/slog"
	"os"
	"time"
)

const (
	defaultConfigPath = "/.config/arcane-game-server"
	defaultConfigName = "config.json"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "arcane-game",
	Short: "An Un-official Magic: The Gathering game server and client",
	Long: `Command line interface for interacting with both the client and server
side infrastructure of the Arcane Game Server.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(initLogger)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.arcane-game.yaml)")

	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Increase logging verbosity for client/server")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}

func initConfig() {
	if cfgFile == "" {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println("error: Failed to find home directory (", err.Error(), ")")
			os.Exit(1)
		}

		cfgFile = home + defaultConfigPath + "/" + defaultConfigName
	}

	err := config.ReadConfigFile(cfgFile)
	if err != nil {
		fmt.Println("error: Failed to read config file (", err.Error(), ")")
		os.Exit(1)
	}
}

func initLogger() {
	buildFileName := func() string {
		timestamp := time.Now().Format(time.RFC3339)
		return viper.GetString("log.path") + "/arcane-" + timestamp + ".json"
	}

	file, err := os.OpenFile(buildFileName(), os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("error: Failed to init logger (", err.Error(), ")")
		os.Exit(1)
	}

	handler := slogmulti.Fanout(
		slog.NewJSONHandler(file, nil),
		slog.NewTextHandler(os.Stdout, nil))

	slog.SetDefault(slog.New(handler))

	viper.Set("log.fileRef", file)
	viper.Set("log.file", file.Name())
}
