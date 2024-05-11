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
	"strings"
)

func StoreNodeSpecificationWithoutEdges(ns *definition.NodeSpecification) error {
	c := getDBConn()
	nodeStmt, err := c.Prepare(insertNodeSQL)
	if err != nil {
		log.Error().Err(err).Msg(logCannotPrepareNodeStmt)
		return err
	}

	attributeStmt, err := c.Prepare(insertNodeAttributeSQL)
	if err != nil {
		log.Error().Err(err).Msg(logCannotPrepareNodeAttributeStmt)
		return err
	}

	specClassID := strings.TrimSpace(ns.ClassID)
	for nodeID, node := range ns.Definitions {
		nodeClassID := strings.TrimSpace(node.ClassID)
		if nodeClassID == "" {
			nodeClassID = specClassID
		}
		log.Debug().Msgf(logAboutToCreateNode, nodeID, node)
		_, err := nodeStmt.Exec(nodeID, nodeClassID)
		if err != nil {
			log.Warn().Err(err).Msgf(logCannotExecuteNodeStmt, nodeID, node)
		}

		// create NodeClassAttribute records
		for attributeID, attribute := range node.Attributes {
			log.Debug().Msgf(logAboutToCreateNodeAttribute, nodeClassID, nodeID, attribute)
			// TODO needs namespace
			_, err := attributeStmt.Exec(nodeID, nodeClassID, "default", attributeID, attribute)
			if err != nil {
				log.Warn().Err(err).Msgf(logCannotExecuteNodeAttributeStmt, attributeID, nodeID, attribute)
			}
		}
	}

	return nil
}

func StoreNodeSpecificationOnlyEdges(ns *definition.NodeSpecification) error {
	c := getDBConn()
	edgeStmt, err := c.Prepare(insertNodeEdgeSQL)
	if err != nil {
		log.Error().Err(err).Msg(logCannotPrepareNodeEdgeStmt)
		return err
	}

	specClassID := strings.TrimSpace(ns.ClassID)
	for nodeID, node := range ns.Definitions {
		nodeClassID := strings.TrimSpace(node.ClassID)
		if nodeClassID == "" {
			nodeClassID = specClassID
		}

		// create NodeClassEdge records
		for _, edge := range node.Edges {
			log.Debug().Msgf(logAboutToCreateNodeEdge, nodeClassID, nodeID, edge)
			//TODO needs namespace
			_, err := edgeStmt.Exec(nodeID, nodeClassID, "default", edge.DestinationNodeID, edge.DestinationNodeClassID, "default", edge.Relationship)
			if err != nil {
				log.Warn().Err(err).Msgf(logCannotExecuteNodeEdgeStmt, nodeClassID, edge)
			}
			if edge.IsBidirectional {
				_, err := edgeStmt.Exec(edge.DestinationNodeID, edge.DestinationNodeClassID, nodeID, nodeClassID, edge.Relationship)
				if err != nil {
					log.Warn().Err(err).Msgf(logCannotExecuteNodeEdgeStmt, nodeClassID, edge)
				}
			}
		}
	}

	return nil
}
