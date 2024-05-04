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

type (
	// Attributes TODO
	Attributes map[string]interface{}

	// NodeEdges TODO
	NodeEdges []NodeEdge

	// NodeEdge TODO
	NodeEdge struct {
		DestinationNodeID      string `yaml:"DestinationNode"`
		DestinationNodeClassID string `yaml:"DestinationNodeClass"`
		Relationship           string `yaml:"Relationship"`
		IsBidirectional        bool   `yaml:"IsBidirectional,omitempty"`
	}

	// NodeDefinition TODO
	NodeDefinition struct {
		// ClassID TODO
		ClassID string `yaml:"Class"`

		// Attributes TODO
		Attributes Attributes `yaml:"Attributes"`

		// Edges TODO
		Edges NodeEdges `yaml:"Edges"`
	}

	// NodeSpecification TODO
	NodeSpecification struct {
		// Class allows the class for all the definitions within the document to be specified.
		ClassID string `yaml:"Class,omitempty"`

		// Definitions TODO
		Definitions map[string]NodeDefinition `yaml:"Definitions,omitempty"`
	}
)
