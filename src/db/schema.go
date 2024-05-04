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
	"database/sql"
	"fmt"
	"github.com/rs/zerolog/log"
)

const (
	enableForeignKeysSQL    = `PRAGMA foreign_keys = ON`
	createDatabaseSchemaSQL = `
		CREATE TABLE NodeClass
		(
			ID          TEXT NOT NULL,
			Namespace	TEXT NOT NULL  DEFAULT "default",
			Description TEXT,
			primary key (ID, Namespace)
		);
		
		CREATE TABLE NodeClassAttribute
		(
			ID          		TEXT    NOT NULL,
			NodeClassID 		TEXT    NOT NULL,
			NodeClassNamespace 	TEXT	NOT NULL  DEFAULT "default", 
			Type        		TEXT    NOT NULL,
			IsRequired  		INTEGER NOT NULL,
			Description 		TEXT,
			primary key (ID, NodeClassID, NodeClassNamespace),
			foreign key (NodeClassID, NodeClassNamespace) references NodeClass (ID, Namespace) ON UPDATE CASCADE
		);
		
		CREATE TABLE NodeClassEdge
		(
			SourceNodeClassID      			TEXT    NOT NULL,
			SourceNodeClassNamespace 		TEXT	NOT NULL  DEFAULT "default",
			DestinationNodeClassID 			TEXT    NOT NULL,
			DestinationNodeClassNamespace	TEXT	NOT NULL  DEFAULT "default",
			Relationship           			TEXT    NOT NULL,
			primary key (SourceNodeClassID, SourceNodeClassNamespace, DestinationNodeClassID, DestinationNodeClassNamespace, Relationship),
			foreign key (SourceNodeClassID, SourceNodeClassNamespace) references NodeClass (ID, Namespace) ON UPDATE CASCADE,
			foreign key (DestinationNodeClassID, DestinationNodeClassNamespace) references NodeClass (ID, Namespace) ON UPDATE CASCADE
		);
		
		CREATE TABLE Node
		(
			ID          		TEXT NOT NULL,
			NodeClassID 		TEXT NOT NULL,
			NodeClassNamespace	TEXT NOT NULL  DEFAULT "default",
			primary key (ID, NodeClassID, NodeClassNamespace),
			foreign key (NodeClassID, NodeClassNamespace) references NodeClass (ID, Namespace) ON UPDATE CASCADE
		);
		
		CREATE TABLE NodeAttribute
		(
			NodeID					TEXT NOT NULL,
			NodeClassID          	TEXT NOT NULL,
			NodeClassNamespace		TEXT NOT NULL DEFAULT "default",
			NodeClassAttributeID 	TEXT NOT NULL,
			Value                	TEXT NOT NULL,
			primary key (NodeID, NodeClassID, NodeClassNamespace, NodeClassAttributeID),
			foreign key (NodeClassID, NodeClassNamespace) references NodeClass (ID, Namespace) ON UPDATE CASCADE,
			foreign key (NodeID, NodeClassID, NodeClassNamespace) references Node (ID, NodeClassID, NodeClassNamespace) ON UPDATE CASCADE,
			foreign key (NodeClassAttributeID, NodeClassID, NodeClassNamespace) references NodeClassAttribute (ID, NodeClassID, NodeClassNamespace) ON UPDATE CASCADE
		);
		
		CREATE TABLE NodeEdge
		(
			SourceNodeID					TEXT    NOT NULL,
			SourceNodeClassID				TEXT    NOT NULL,
			SourceNodeClassNamespace 		TEXT	NOT NULL DEFAULT "default",
			DestinationNodeID				TEXT    NOT NULL,
			DestinationNodeClassID			TEXT    NOT NULL,
			DestinationNodeClassNamespace	TEXT	NOT NULL DEFAULT "default",
			Relationship           			TEXT    NOT NULL,
			primary key (SourceNodeID, SourceNodeClassID, SourceNodeClassNamespace, DestinationNodeID, DestinationNodeClassID, DestinationNodeClassNamespace, Relationship),
			foreign key (SourceNodeClassID, SourceNodeClassNamespace) references NodeClass (ID, Namespace) ON UPDATE CASCADE,
			foreign key (DestinationNodeClassID, DestinationNodeClassNamespace) references NodeClass (ID, Namespace) ON UPDATE CASCADE,
			foreign key (SourceNodeID, SourceNodeClassID, SourceNodeClassNamespace) references Node (ID, NodeClassID, NodeClassNamespace) ON UPDATE CASCADE,
			foreign key (DestinationNodeID, DestinationNodeClassID, DestinationNodeClassNamespace) references Node (ID, NodeClassID, NodeClassNamespace) ON UPDATE CASCADE,
			foreign key (SourceNodeClassID, SourceNodeClassNamespace, DestinationNodeClassID, DestinationNodeClassNamespace, Relationship) references NodeClassEdge (SourceNodeClassID, SourceNodeClassNamespace, DestinationNodeClassID, DestinationNodeClassNamespace, Relationship) ON UPDATE CASCADE
		
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
