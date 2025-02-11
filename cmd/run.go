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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stevezaluk/arcane-game/game"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the server according to the configuration values you have provided",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		server, err := game.NewServer(viper.GetString("server.lobby_name"), viper.GetString("server.game_mode"))
		if err != nil {
			panic(err)
		}

		server.Start()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringP("name", "n", "Default Server Name", "Describe the name of the lobby that you are starting")
	viper.BindPFlag("server.lobby_name", runCmd.Flags().Lookup("name"))

	runCmd.Flags().StringP("mode", "m", game.CommanderGameMode, "Set the game mode for the lobby that you are starting")
	viper.BindPFlag("server.game_mode", runCmd.Flags().Lookup("mode"))
}
