/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(modifyCmd)
}

var modifyCmd = &cobra.Command{
	Use:   "modify",
	Short: "Modify a resource",
	Long:  "Modify a resource",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}