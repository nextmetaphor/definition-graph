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
	// NodeClassAttributes TODO
	NodeClassAttributes map[string]NodeClassAttribute

	// NodeClassAttribute TODO
	NodeClassAttribute struct {
		Description string `yaml:"Description,omitempty"`
		Type        string `yaml:"Type"`
		IsRequired  bool   `yaml:"IsRequired"`
	}

	// NodeClassEdges TODO
	NodeClassEdges []NodeClassEdge

	// NodeClassEdge TODO
	NodeClassEdge struct {
		DestinationNodeClassID string `yaml:"DestinationNodeClass"`
		Relationship           string `yaml:"Relationship"`
		IsBidirectional        bool   `yaml:"IsBidirectional,omitempty"`
	}

	// NodeClassDefinition TODO
	NodeClassDefinition struct {
		// Description TODO
		Description string `yaml:"Description,omitempty"`

		// Attributes TODO
		Attributes NodeClassAttributes `yaml:"Attributes"`

		// Edges TODO
		Edges NodeClassEdges `yaml:"Edges"`
	}

	// NodeClassSpecification TODO
	NodeClassSpecification struct {
		// Definitions TODO
		Definitions map[string]NodeClassDefinition `yaml:"Definitions,omitempty"`
	}
)
