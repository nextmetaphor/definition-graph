/*
 * Copyright (C) 2024 Paul Tatham <paul@nextmetaphor.io>
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

package definition

import "fmt"

const (
	graphNodeIDFormatString = "%s:%s"
)

type (
	Graph struct {
		Nodes []GraphNode `json:"nodes"`
		Links []GraphLink `json:"links"`
	}

	GraphNode struct {
		ID          string `json:"id"`
		Class       string `json:"class"`
		Namespace   string `json:"namespace"`
		Description string `json:"description"`
	}

	GraphLink struct {
		Source       string `json:"source"`
		Target       string `json:"target"`
		Relationship string `json:"relationship"`
	}
)

func GraphNodeID(nodeClassID, nodeID string) string {
	return fmt.Sprintf(graphNodeIDFormatString, nodeClassID, nodeID)
}
