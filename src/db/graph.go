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

package db

import (
	"github.com/nextmetaphor/definition-graph/definition"
	"github.com/rs/zerolog/log"
)

const (
	selectGraphNodeEdgeSQL = `SELECT SourceNodeID, SourceNodeClassID, SourceNodeClassNamespace, DestinationNodeID, DestinationNodeClassID, DestinationNodeClassNamespace, Relationship from NodeEdge;`
)

func SelectNodeClassGraph() (graph definition.Graph, err error) {
	c := getDBConn()
	nodeRows, err := c.Query(selectNodeClassSQL)
	if err != nil {
		log.Error().Err(err).Msg(logCannotQueryNodeClassSelectStmt)
		return
	}
	defer nodeRows.Close()

	for nodeRows.Next() {
		var node definition.GraphNode
		if err = nodeRows.Scan(&node.ID, &node.Namespace, &node.Description); err != nil {
			return
		}
		node.Class = node.ID
		graph.Nodes = append(graph.Nodes, node)
	}

	linkRows, err := c.Query(selectNodeClassEdgeSQL)

	if err != nil {
		log.Error().Err(err).Msg(logCannotQueryNodeClassEdgeSelectStmt)
		return
	}
	defer linkRows.Close()

	for linkRows.Next() {
		var link definition.GraphLink
		if err = linkRows.Scan(&link.Source, &link.Target, &link.Relationship); err != nil {
			return
		}
		graph.Links = append(graph.Links, link)
	}

	return
}

func SelectNodeGraph() (graph definition.Graph, err error) {
	c := getDBConn()
	nodeRows, err := c.Query(selectNodeSQL)
	if err != nil {
		log.Error().Err(err).Msg(logCannotQueryNodeSelectStmt)
		return
	}
	defer nodeRows.Close()

	for nodeRows.Next() {
		var node definition.GraphNode
		var nodeID, classID string
		if err = nodeRows.Scan(&nodeID, &classID); err != nil {
			return
		}
		node.ID = definition.GraphNodeID(classID, nodeID)
		node.Class = classID
		node.Description = node.ID
		graph.Nodes = append(graph.Nodes, node)
	}

	linkRows, err := c.Query(selectGraphNodeEdgeSQL)
	if err != nil {
		log.Error().Err(err).Msg(logCannotQueryNodeEdgeSelectStmt)
		return
	}
	defer linkRows.Close()

	for linkRows.Next() {
		var link definition.GraphLink
		var sourceNodeID, sourceNodeClassID, destinationNodeID, destinationNodeClassID string
		if err = linkRows.Scan(&sourceNodeID, &sourceNodeClassID, &destinationNodeID, &destinationNodeClassID, &link.Relationship); err != nil {
			return
		}
		link.Source = definition.GraphNodeID(sourceNodeClassID, sourceNodeID)
		link.Target = definition.GraphNodeID(destinationNodeClassID, destinationNodeID)

		graph.Links = append(graph.Links, link)
	}

	return
}
