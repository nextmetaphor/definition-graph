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

	databaseName = "definition-graph.db"

	databaseDriver = "sqlite3"

	logCannotOpenDatabase      = "cannot open database [%s]"
	logCannotDropDatabase      = "cannot drop database schema"
	logCannotCreateDatabase    = "cannot create database schema"
	logCannotCloseDatabase     = "cannot close database"
	logCannotEnableForeignKeys = "cannot enable foreign keys"
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