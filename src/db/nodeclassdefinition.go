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

func StoreNodeClassSpecification(ncs *definition.NodeClassSpecification) error {
	c := getDBConn()
	stmt, err := c.Prepare(createNodeClassSQL)
	if err != nil {
		log.Error().Err(err).Msg(logCannotPrepareNodeClassStmt)
		return err
	}

	attributeStmt, err := c.Prepare(insertNodeClassAttributeSQL)
	if err != nil {
		log.Error().Err(err)
		return err
	}

	edgeStmt, err := c.Prepare(insertNodeClassEdgeSQL)
	if err != nil {
		log.Error().Err(err).Msg(logCannotPrepareNodeClassEdgeStmt)
		return err
	}

	for classID, classDefinition := range ncs.Definitions {
		// create NodeClass record
		_, err := stmt.Exec(classID, "default", classDefinition.Description)
		if err != nil {
			log.Warn().Err(err).Msgf(logCannotExecuteNodeClassStmt, classID, classDefinition)
		}

		// create NodeClassAttribute records
		for attributeID, attribute := range classDefinition.Attributes {
			_, err := attributeStmt.Exec(attributeID, classID, "default", attribute.Description, attribute.Type, boolToInt(attribute.IsRequired))
			if err != nil {
				log.Warn().Err(err)
			}
		}

		// create NodeClassEdge records
		for _, edge := range classDefinition.Edges {
			// TODO - needs namespaces in definition files
			_, err := edgeStmt.Exec(classID, "default", edge.DestinationNodeClassID, "default", edge.Relationship)
			if err != nil {
				log.Warn().Err(err).Msgf(logCannotExecuteNodeClassEdgeStmt, classID, edge)
			}
			if edge.IsBidirectional {
				_, err := edgeStmt.Exec(edge.DestinationNodeClassID, classID, edge.Relationship)
				if err != nil {
					log.Warn().Err(err).Msgf(logCannotExecuteNodeClassEdgeStmt, classID, edge)
				}
			}
		}
	}

	return nil
}
