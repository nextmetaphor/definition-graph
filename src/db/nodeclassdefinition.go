package db

import (
	"database/sql"
	"github.com/nextmetaphor/definition-graph/definition"
	"github.com/rs/zerolog/log"
)

func StoreNodeClassSpecification(db *sql.DB, ncs *definition.NodeClassSpecification) error {
	stmt, err := db.Prepare(createNodeClassSQL)
	if err != nil {
		log.Error().Err(err).Msg(logCannotPrepareNodeClassStmt)
		return err
	}

	attributeStmt, err := db.Prepare(insertNodeClassAttributeSQL)
	if err != nil {
		log.Error().Err(err)
		return err
	}

	edgeStmt, err := db.Prepare(insertNodeClassEdgeSQL)
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
