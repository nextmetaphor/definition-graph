package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nextmetaphor/definition-graph/definition"
	"log"
)

const (
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

	databaseName = "definition-graph.db"

	databaseDriver = "sqlite3"
)

func OpenDatabase() *sql.DB {
	db, err := sql.Open(databaseDriver, fmt.Sprintf("./%s", databaseName))
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(dropDatabaseSchemeSQL)
	if err != nil {
		log.Printf("%q: %s\n", err, dropDatabaseSchemeSQL)
		return nil
	}

	_, err = db.Exec(createDatabaseSchemaSQL)
	if err != nil {
		log.Printf("%q: %s\n", err, createDatabaseSchemaSQL)
		return nil
	}

	return db
}

func CloseDatabase(db *sql.DB) error {
	return db.Close()
}

func StoreNodeClassSpecification(db *sql.DB, ncs *definition.NodeClassSpecification) {
	stmt, err := db.Prepare(insertNodeClassSQL)
	if err != nil {
		log.Fatal(err)
	}

	attributeStmt, err := db.Prepare(insertNodeClassAttributeSQL)
	if err != nil {
		log.Fatal(err)
	}

	for id, classDefinition := range ncs.Definitions {
		// create NodeClass record
		_, err := stmt.Exec(id, classDefinition.Description)
		if err != nil {
			log.Fatal(err)
		}

		// create NodeClassAttribute records
		for attributeID, attribute := range classDefinition.Attributes {
			_, err := attributeStmt.Exec(attributeID, id, attribute.Description, attribute.Type, attribute.IsRequired)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
