/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package cmd

import (
	"github.com/spf13/cobra"
)

var (
	restoreCmd = &cobra.Command{
		Use:   "restore",
		Short: "Restore the data for a resource",
		Long:  "Restore the data for a resource",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
	restoreCmdPathFlag = restoreCmd.PersistentFlags().String("path", "", "the absolute path to a backup file")
	restoreCmdPipeFlag = restoreCmd.PersistentFlags().Bool("pipe", false, "whether to retrieve the backup data from STDIN")
)

func init() {
	rootCmd.AddCommand(restoreCmd)
}
