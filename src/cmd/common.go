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
	"github.com/spf13/cobra"
	"os"
)

const (
	appName    = "definition-graph"
	appVersion = "0.0.1"

	commandRootUse      = appName
	commandRootUseShort = appName + ": generate graph representations from simple definition files"
	commandRootUseLong  = "Define data in YAML/JSON then generate graph representations to model relationships"

	commandVersionUse    = "version"
	commandVersionShort  = "Print the version number of " + appName
	commandVersionString = appVersion

	commandAPIUse   = "api"
	commandAPIShort = "Start the " + appName + " API server"

	commandDefineUse   = "define"
	commandDefineShort = "Load nodeClass definition files from a specified directory"

	commandPopulateUse   = "populate"
	commandPopulateShort = "Load node definition files from a specified directory"

	flagAPIAddress          = "address"
	flagAPIAddressShorthand = "a"
	flagAPIAddressDefault   = ""
	flagAPIAddressUsage     = "address for api"

	flagAPIPort          = "port"
	flagAPIPortShorthand = "p"
	flagAPIPortDefault   = 8080
	flagAPIPortUsage     = "port for api"

	flagDefinitionDirectory          = "dir"
	flagDefinitionDirectoryShorthand = "d"
	flagDefinitionDirectoryUsage     = "directory from which to load the definition files"

	flagDefinitionFormat          = "format"
	flagDefinitionFormatShorthand = "f"
	flagDefinitionFormatDefault   = "yaml"
	flagDefinitionFormatUsage     = "format of the definition files (yaml or json)"

	flagRecreateDatabase          = "recreateDB"
	flagRecreateDatabaseShorthand = "c"
	flagRecreateDatabaseUsage     = "recreate database"
	flagRecreateDatabaseDefault   = false

	exitCodeRootCmdFailed = 1
)

var (
	rootCmd = &cobra.Command{
		Use:   commandRootUse,
		Short: commandRootUseShort,
		Long:  commandRootUseLong,
	}
)

var (
	// variable for flagAPIPort parameter
	apiServerPort int

	// variable for flagAPIAddress parameter
	apiServerHost string

	defineDefinitionDirectory   string
	defineDefinitionFormat      string
	populateDefinitionDirectory string
	populateDefinitionFormat    string
	recreateDatabase            bool
)

// Execute TODO
func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	if err := rootCmd.Execute(); err != nil {
		os.Exit(exitCodeRootCmdFailed)
	}
}
