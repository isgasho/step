/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package cmd

import (
	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Display version information",
		Long:  "Display version information",
		Run: func(cmd *cobra.Command, args []string) {
			rootCmd.SetArgs([]string{"--version"})
			rootCmd.Execute()
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
	changeHelpUsageText(versionCmd)
}
