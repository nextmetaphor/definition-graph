package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nextmetaphor/definition-graph/data"
	"github.com/nextmetaphor/definition-graph/definition"
	"github.com/rs/zerolog/log"
)

const (
	insertNodeClassSQL          = `INSERT INTO NodeClass (ID, Namespace, Description) values (?, ?, ?);`
	insertNodeClassAttributeSQL = `INSERT INTO NodeClassAttribute (ID, NodeClassID, Description, Type, IsRequired) values (?, ?, ?, ?, ?);`
	insertNodeClassEdgeSQL      = `INSERT INTO NodeClassEdge (SourceNodeClassID, DestinationNodeClassID, Relationship) values (?, ?, ?);`
	selectNodeClassSQL          = `SELECT ID, Namespace, Description from NodeClass order by Namespace, ID`
	selectNodeClassEdgeSQL      = `SELECT SourceNodeClassID, DestinationNodeClassID, Relationship from NodeClassEdge`

	logCannotPrepareNodeClassStmt          = "cannot prepare NodeClass insert statement"
	logCannotPrepareNodeClassAttributeStmt = "cannot prepare NodeClassAttribute insert statement"
	logCannotPrepareNodeClassEdgeStmt      = "cannot prepare NodeClassEdge insert statement"
	logCannotExecuteNodeClassStmt          = "cannot execute NodeClass insert statement, id=[%s], [%#v]"
	logCannotExecuteNodeClassAttributeStmt = "cannot execute NodeClassAttribute insert statement, classid=[%s], id=[%s], [%#v]"
	logCannotExecuteNodeClassEdgeStmt      = "cannot execute NodeClassEdge insert statement, classid=[%s], [%#v]"
	logCannotQueryNodeClassSelectStmt      = "cannot query NodeClass select statement"
	logCannotQueryNamespaceSelectStmt      = "cannot query Namespace select statement"
	logCannotQueryNodeClassEdgeSelectStmt  = "cannot query NodeClassEdge select statement"
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
			_, err := attributeStmt.Exec(attributeID, classID, attribute.Description, attribute.Type, boolToInt(attribute.IsRequired))
			if err != nil {
				log.Warn().Err(err).Msgf(logCannotExecuteNodeClassAttributeStmt, attributeID, classID, attribute)
			}
		}

		// create NodeClassEdge records
		for _, edge := range classDefinition.Edges {
			_, err := edgeStmt.Exec(classID, edge.DestinationNodeClassID, edge.Relationship)
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

func SelectNodeClass(db *sql.DB) (nodeClasses data.NodeClassesOuter, err error) {
	nodeClassRows, err := db.Query(selectNodeClassSQL)
	if err != nil {
		log.Error().Err(err).Msg(logCannotQueryNodeClassSelectStmt)
		return
	}
	defer nodeClassRows.Close()

	for nodeClassRows.Next() {
		var nodeClass data.NodeClass
		if err = nodeClassRows.Scan(&nodeClass.ID, &nodeClass.Namespace, &nodeClass.Description); err == nil {
			nodeClasses.NodeClasses = append(nodeClasses.NodeClasses, nodeClass)
		}
	}

	return
}

func CreateNodeClass(c *sql.DB, nc data.NodeClass) (e error) {
	s, e := c.Prepare(insertNodeClassSQL)
	if e == nil {
		_, e = s.Exec(nc.ID, nc.Namespace, nc.Description)
	}

	return
}

func SelectNodeClassGraph(db *sql.DB) (graph definition.Graph, err error) {
	nodeRows, err := db.Query(selectNodeClassSQL)
	if err != nil {
		log.Error().Err(err).Msg(logCannotQueryNodeClassSelectStmt)
		return
	}
	defer nodeRows.Close()

	for nodeRows.Next() {
		var node definition.GraphNode
		if err = nodeRows.Scan(&node.ID, &node.Description); err != nil {
			return
		}
		node.Class = node.ID
		graph.Nodes = append(graph.Nodes, node)
	}

	linkRows, err := db.Query(selectNodeClassEdgeSQL)
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
