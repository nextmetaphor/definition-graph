package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nextmetaphor/definition-graph/definition"
	"github.com/rs/zerolog/log"
	"strings"
)

const (
	enableForeignKeysSQL    = `PRAGMA foreign_keys = ON`
	createDatabaseSchemaSQL = `
		CREATE TABLE NodeClass
		(
			ID          TEXT not null primary key,
			Description TEXT
		);
		
		CREATE TABLE NodeClassAttribute
		(
			ID          TEXT    not null,
			NodeClassID TEXT    not null references NodeClass,
			Type        TEXT    not null,
			IsRequired  INTEGER not null,
			Description TEXT,
			primary key (ID, NodeClassID)
		);
		
		CREATE TABLE NodeClassEdge
		(
			SourceNodeClassID      TEXT    not null references NodeClass,
			DestinationNodeClassID TEXT    not null references NodeClass,
			Relationship           TEXT    not null,
			IsFromSource           INTEGER not null,
			IsFromDestination      INTEGER not null,
			primary key (SourceNodeClassID, DestinationNodeClassID, Relationship, IsFromSource, IsFromDestination)
		);
		
		CREATE TABLE Node
		(
			ID          TEXT not null,
			NodeClassID TEXT references NodeClass,
			primary key (ID, NodeClassID)
		);
		
		CREATE TABLE NodeAttribute
		(
			NodeID               TEXT not null,
			NodeClassID          TEXT not null,
			NodeClassAttributeID TEXT not null,
			Value                TEXT not null,
			primary key (NodeID, NodeClassID, NodeClassAttributeID),
			foreign key (NodeID, NodeClassID) references Node (ID, NodeClassID),
			foreign key (NodeClassAttributeID, NodeClassID) references NodeClassAttribute (ID, NodeClassID)
		);
		
		CREATE TABLE NodeEdge
		(
			SourceNodeID           TEXT    not null,
			SourceNodeClassID      TEXT    not null references NodeClass,
			DestinationNodeID      TEXT    not null,
			DestinationNodeClassID TEXT    not null references NodeClass,
			Relationship           TEXT    not null,
			IsFromSource           INTEGER not null,
			IsFromDestination      INTEGER not null,
			primary key (SourceNodeID, SourceNodeClassID, DestinationNodeID, DestinationNodeClassID, Relationship, IsFromSource,
						 IsFromDestination),
			foreign key (SourceNodeID, SourceNodeClassID) references Node (ID, NodeClassID),
			foreign key (DestinationNodeID, DestinationNodeClassID) references Node (ID, NodeClassID),
			foreign key (SourceNodeClassID, DestinationNodeClassID, Relationship, IsFromSource,
						 IsFromDestination) references NodeClassEdge (SourceNodeClassID, DestinationNodeClassID, Relationship,
																	  IsFromSource, IsFromDestination)
		)
	`

	dropDatabaseSchemeSQL = `
		drop table NodeEdge;
		drop table NodeAttribute;
		drop table Node;
		drop table NodeClassEdge;
		drop table NodeClassAttribute;
		drop TABLE NodeClass;
	`

	insertNodeClassSQL = `INSERT INTO NodeClass (ID, Description) values (?, ?);`

	insertNodeClassAttributeSQL = `INSERT INTO NodeClassAttribute (ID, NodeClassID, Description, Type, IsRequired) values (?, ?, ?, ?, ?);`

	insertNodeSQL = `INSERT INTO Node (ID, NodeClassID) values (?, ?);`

	insertNodeAttributeSQL = `INSERT INTO NodeAttribute (NodeID, NodeClassID, NodeClassAttributeID, Value) values (?, ?, ?, ?);`

	databaseName = "definition-graph.db"

	databaseDriver = "sqlite3"

	logCannotOpenDatabase                  = "cannot open database [%s]"
	logCannotDropDatabase                  = "cannot drop database schema"
	logCannotCreateDatabase                = "cannot create database schema"
	logCannotCloseDatabase                 = "cannot close database"
	logCannotEnableForeignKeys             = "cannot enable foreign keys"
	logCannotPrepareNodeClassStmt          = "cannot prepare NodeClass insert statement"
	logCannotPrepareNodeClassAttributeStmt = "cannot prepare NodeClassAttribute insert statement"
	logCannotExecuteNodeClassStmt          = "cannot execute NodeClass insert statement, id=[%s], [%#v]"
	logCannotExecuteNodeClassAttributeStmt = "cannot execute NodeClassAttribute insert statement, classid=[%s], id=[%s], [%#v]"
	logCannotPrepareNodeStmt               = "cannot prepare Node insert statement"
	logCannotPrepareNodeAttributeStmt      = "cannot prepare NodeAttribute insert statement"
	logCannotExecuteNodeStmt               = "cannot execute Node insert statement, id=[%s], [%#v]"
	logCannotExecuteNodeAttributeStmt      = "cannot execute NodeAttribute insert statement, classid=[%s], id=[%s], [%#v]"
)

func OpenDatabase() (*sql.DB, error) {
	db, err := sql.Open(databaseDriver, fmt.Sprintf("./%s", databaseName))
	if err != nil {
		log.Error().Err(err).Msgf(logCannotOpenDatabase, databaseName)
		return nil, err
	}

	_, err = db.Exec(dropDatabaseSchemeSQL)
	if err != nil {
		log.Error().Err(err).Msg(logCannotDropDatabase)
		return nil, err
	}

	_, err = db.Exec(enableForeignKeysSQL)
	if err != nil {
		log.Error().Err(err).Msg(logCannotEnableForeignKeys)
		return nil, err
	}

	_, err = db.Exec(createDatabaseSchemaSQL)
	if err != nil {
		log.Error().Err(err).Msg(logCannotCreateDatabase)
		return nil, err
	}

	return db, nil
}

func CloseDatabase(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		log.Error().Err(err).Msg(logCannotCloseDatabase)
	}
	return err
}

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

	for id, classDefinition := range ncs.Definitions {
		// create NodeClass record
		_, err := stmt.Exec(id, classDefinition.Description)
		if err != nil {
			log.Warn().Err(err).Msgf(logCannotExecuteNodeClassStmt, id, classDefinition)
		}

		// create NodeClassAttribute records
		for attributeID, attribute := range classDefinition.Attributes {
			_, err := attributeStmt.Exec(attributeID, id, attribute.Description, attribute.Type, attribute.IsRequired)
			if err != nil {
				log.Warn().Err(err).Msgf(logCannotExecuteNodeClassAttributeStmt, attributeID, id, attribute)
			}
		}
	}

	return nil
}

func StoreNodeSpecification(db *sql.DB, ns *definition.NodeSpecification) error {
	nodeStmt, err := db.Prepare(insertNodeSQL)
	if err != nil {
		log.Error().Err(err).Msg(logCannotPrepareNodeStmt)
		return err
	}

	attributeStmt, err := db.Prepare(insertNodeAttributeSQL)
	if err != nil {
		log.Error().Err(err).Msg(logCannotPrepareNodeAttributeStmt)
		return err
	}

	specClassID := strings.TrimSpace(ns.ClassID)
	for nodeID, nodeDefinition := range ns.Definitions {
		defClassID := strings.TrimSpace(nodeDefinition.ClassID)
		if defClassID == "" {
			defClassID = specClassID
		}
		_, err := nodeStmt.Exec(nodeID, defClassID)
		if err != nil {
			log.Warn().Err(err).Msgf(logCannotExecuteNodeStmt, nodeID, nodeDefinition)
		}

		// create NodeClassAttribute records
		for attributeID, attribute := range nodeDefinition.Attributes {
			_, err := attributeStmt.Exec(nodeID, defClassID, attributeID, attribute)
			if err != nil {
				log.Warn().Err(err).Msgf(logCannotExecuteNodeAttributeStmt, attributeID, nodeID, attribute)
			}
		}
	}

	return nil
}
