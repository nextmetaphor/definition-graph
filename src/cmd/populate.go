/*
 * Copyright (C) 2024  Paul Tatham <paul@nextmetaphor.io>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package cmd

import (
	"github.com/nextmetaphor/definition-graph/core"
	"github.com/nextmetaphor/definition-graph/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(populateCmd)
	populateCmd.Flags().StringVarP(&populateDefinitionDirectory, flagDefinitionDirectory, flagDefinitionDirectoryShorthand, "", flagDefinitionDirectoryUsage)
	populateCmd.Flags().StringVarP(&populateDefinitionFormat, flagDefinitionFormat, flagDefinitionFormatShorthand, flagDefinitionFormatDefault, flagDefinitionFormatUsage)
}

var populateCmd = &cobra.Command{
	Use:   commandPopulateUse,
	Short: commandPopulateShort,
	Run: func(cmd *cobra.Command, args []string) {
		db.OpenDatabase()
		defer db.CloseDatabase()
		var e error
		if e = core.LoadNodeDefinitionsWithoutEdges([]string{populateDefinitionDirectory}, populateDefinitionFormat); e == nil {
			e = core.LoadNodeDefinitionsOnlyEdges([]string{populateDefinitionDirectory}, populateDefinitionFormat)
		}
	},
}
