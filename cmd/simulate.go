/*
Copyright © 2025 Steven Zaluk <arcanegame@protonmail.com>

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
)

var simulateCmd = &cobra.Command{
	Use:   "simulate",
	Short: "Simulate a game of MTG without initializing the server or making API Calls",
	Long: `The simulate command allows the user to initialize a game without making API calls
or setting up the network infrastructure required to start a game server. You can use this
for experimenting with the Game engine itself`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(simulateCmd)
}
