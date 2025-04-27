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
	"os"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Error: Client simulation not implemented")
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	/*
		Connection CLI Flags
	*/
	clientCmd.Flags().String("client.server_hostname", "127.0.0.1", "The hostname of the Arcane Game server you want to connect to")
	clientCmd.Flags().Int("client.server_port", 44444, "The port of the Arcane Game Server you want to connect to")

	/*
		Iterates through each command and binds there long values to viper values
	*/
	err := viper.BindPFlags(clientCmd.Flags())
	if err != nil {
		fmt.Println("Error binding Cobra flags to viper: ", err.Error())
		os.Exit(1)
	}

}
