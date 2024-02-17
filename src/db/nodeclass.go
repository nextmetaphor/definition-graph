package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nextmetaphor/definition-graph/definition"
	"github.com/rs/zerolog/log"
)

const (
	insertNodeClassSQL = `INSERT INTO NodeClass (ID, Description) values (?, ?);`

	insertNodeClassAttributeSQL = `INSERT INTO NodeClassAttribute (ID, NodeClassID, Description, Type, IsRequired) values (?, ?, ?, ?, ?);`

	insertNodeClassEdgeSQL = `INSERT INTO NodeClassEdge (SourceNodeClassID, DestinationNodeClassID, Relationship, IsFromSource, IsFromDestination) values (?, ?, ?, ?, ?);`

	logCannotPrepareNodeClassStmt          = "cannot prepare NodeClass insert statement"
	logCannotPrepareNodeClassAttributeStmt = "cannot prepare NodeClassAttribute insert statement"
	logCannotPrepareNodeClassEdgeStmt      = "cannot prepare NodeClassEdge insert statement"
	logCannotExecuteNodeClassStmt          = "cannot execute NodeClass insert statement, id=[%s], [%#v]"
	logCannotExecuteNodeClassAttributeStmt = "cannot execute NodeClassAttribute insert statement, classid=[%s], id=[%s], [%#v]"
	logCannotExecuteNodeClassEdgeStmt      = "cannot execute NodeClassEdge insert statement, classid=[%s], [%#v]"
)

func StoreNodeClassSpecification(db *sql.DB, ncs *definition.NodeClassSpecification) error {
	stmt, err := db.Prepare(insertNodeClassSQL)
	if err != nil {
		log.Error().Err(err).Msg(logCannotPrepareNodeClassStmt)
		return err
	}

	attributeStmt, err := db.Prepare(insertNodeClassAttributeSQL)
	if err != nil {
		log.Error().Err(err).Msg(logCannotPrepareNodeClassAttributeStmt)
		return err
	}

	edgeStmt, err := db.Prepare(insertNodeClassEdgeSQL)
	if err != nil {
		log.Error().Err(err).Msg(logCannotPrepareNodeClassEdgeStmt)
		return err
	}

	for classID, classDefinition := range ncs.Definitions {
		// create NodeClass record
		_, err := stmt.Exec(classID, classDefinition.Description)
		if err != nil {
			log.Warn().Err(err).Msgf(logCannotExecuteNodeClassStmt, classID, classDefinition)
		}

		// create NodeClassAttribute records
		for attributeID, attribute := range classDefinition.Attributes {
			_, err := attributeStmt.Exec(attributeID, classID, attribute.Description, attribute.Type, attribute.IsRequired)
			if err != nil {
				log.Warn().Err(err).Msgf(logCannotExecuteNodeClassAttributeStmt, attributeID, classID, attribute)
			}
		}

		// create NodeClassEdge records
		for _, edge := range classDefinition.Edges {
			_, err := edgeStmt.Exec(classID, edge.DestinationNodeClassID, edge.Relationship, edge.IsToDestination, edge.IsFromDestination)
			if err != nil {
				log.Warn().Err(err).Msgf(logCannotExecuteNodeClassEdgeStmt, classID, edge)
			}
		}
	}

	return nil
}
