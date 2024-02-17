package db

import (
	"database/sql"
	"fmt"
	"github.com/rs/zerolog/log"
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
			primary key (SourceNodeClassID, DestinationNodeClassID, Relationship)
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
			primary key (SourceNodeID, SourceNodeClassID, DestinationNodeID, DestinationNodeClassID, Relationship),
			foreign key (SourceNodeID, SourceNodeClassID) references Node (ID, NodeClassID),
			foreign key (DestinationNodeID, DestinationNodeClassID) references Node (ID, NodeClassID),
			foreign key (SourceNodeClassID, DestinationNodeClassID, Relationship) references NodeClassEdge (SourceNodeClassID, DestinationNodeClassID, Relationship)
		)
	`

	dropDatabaseSchemeSQL = `
		DROP TABLE IF EXISTS NodeEdge;
		DROP TABLE IF EXISTS NodeAttribute;
		DROP TABLE IF EXISTS Node;
		DROP TABLE IF EXISTS NodeClassEdge;
		DROP TABLE IF EXISTS NodeClassAttribute;
		DROP TABLE IF EXISTS NodeClass;
	`

	databaseName = "definition-graph.db"

	databaseDriver = "sqlite3"

	logCannotOpenDatabase      = "cannot open database [%s]"
	logCannotDropDatabase      = "cannot drop database schema"
	logCannotCreateDatabase    = "cannot create database schema"
	logCannotCloseDatabase     = "cannot close database"
	logCannotEnableForeignKeys = "cannot enable foreign keys"
)

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

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
