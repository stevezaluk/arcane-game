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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stevezaluk/arcane-game/game"
	"log/slog"
	"os"
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
	PreRun: func(cmd *cobra.Command, args []string) {
		if viper.GetBool("verbose") {
			slog.SetLogLoggerLevel(slog.LevelDebug)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	/*
		General CLI Flags
	*/
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.arcane-game.yaml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Increase logging verbosity for client/server")

	/*
		Lobby CLI Flags
	*/
	rootCmd.Flags().StringP("lobby.name", "n", "Arcane Lobby", "The name of the Lobby to use for user discovery")
	rootCmd.Flags().StringP("lobby.game_mode", "g", game.CommanderGameMode, "The name of the Lobby to use for user discovery")

	/*
		Log CLI Flags
	*/
	rootCmd.Flags().StringP("log.path", "l", "/var/log/arcane-game-server", "The default path that logs will be written to")

	/*
		Server CLI Flags
	*/
	rootCmd.Flags().IntP("server.port", "p", 44444, "The port that the game server should listen for connections on")
	rootCmd.Flags().Int("server.max_connections", 4, "The maximum number of connections the server will accept before closing the lobby. Effectively acts a max player count")
	rootCmd.Flags().Bool("server.secure_connections", true, "Whether or not to use secure connections (end-to-end encryption")

	/*
		MTGJSON API Flags
	*/
	rootCmd.Flags().String("api.hostname", "127.0.0.1", "The default hostname of the MTGJSON API for user/card metadata")
	rootCmd.Flags().Int("api.port", 8080, "The default port of the MTGJSON API to use for user/card metadata")
	rootCmd.Flags().String("api.username", "", "The username to use for MTGJSON API authentication. See README for required scopes")
	rootCmd.Flags().String("api.password", "", "The password to use for MTGJSON API authentication. See README for required scopes")
	rootCmd.Flags().Bool("api.prompt_password", false, "Prompt for the MTGJSON API password to avoid saving it to command history. Ignores any other api.password values")

	/*
		Unused CLI Flags - Flags that will be implemented in the future
	*/
	rootCmd.Flags().String("crypto.key_algorithm", "rsa", "The default key exchange algorithm used for creating end-to-end encrypted connections. Clients must also use this algorithm")
	rootCmd.Flags().Int("crypto.key_size", 4096, "The default size of the key to use for server and client encryption keys")

	/*
		Iterates through each command and binds there long values to viper values
	*/
	err := viper.BindPFlags(rootCmd.Flags())
	if err != nil {
		fmt.Println("Error binding Cobra flags to viper: ", err.Error())
		os.Exit(1)
	}
}

/*
initConfig - Initialize viper with values from config files or environmental variables. Defaults
are not set here as CLI arguments are bound to viper config values. These provide defaults. Should
not be called directly, automatically called as a part of viper's initialization stack
*/
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.SetConfigType("json")
		viper.AddConfigPath(home + "/.config/arcane-game-server/")
		viper.SetConfigName("config.json")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err.Error())
		os.Exit(1)
	}
}
