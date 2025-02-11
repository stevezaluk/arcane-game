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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stevezaluk/arcane-game/game"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		gameClient, err := game.NewClient()
		if err != nil {
			fmt.Println("Error while creating game client", err.Error())
			return
		}

		err = gameClient.Connect(viper.GetString("client.server_ip"), viper.GetInt("client.server_port"))
		if err != nil {
			fmt.Println("Error while connecting client to game server", err.Error())
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
