package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nextmetaphor/definition-graph/definition"
	"github.com/nextmetaphor/definition-graph/model"
	"github.com/rs/zerolog/log"
	"strings"
)

const (
	insertNodeSQL                  = `INSERT INTO Node (ID, NodeClassID) values (?, ?);`
	insertNodeAttributeSQL         = `INSERT INTO NodeAttribute (NodeID, NodeClassID, NodeClassAttributeID, Value) values (?, ?, ?, ?);`
	insertNodeEdgeSQL              = `INSERT INTO NodeEdge (SourceNodeID, SourceNodeClassID, DestinationNodeID, DestinationNodeClassID, Relationship) values (?, ?, ?, ?, ?);`
	selectNodeSQL                  = `SELECT ID, NodeClassID from Node`
	selectNodeSQLByClass           = `SELECT ID, NodeClassID from Node where NodeClassID=? and NodeClassNameSpace=?`
	selectNodeSQLByID              = `SELECT ID, NodeClassID, NodeClassNamespace from Node where NodeClassNameSpace=? and NodeClassID=? and ID=?`
	selectNodeAttributeSQLByNodeID = `SELECT NodeClassAttributeID, Value from NodeAttribute where NodeID=? and NodeClassID=? and NodeClassNameSpace=?`

	selectNodeEdgeSQL = `SELECT SourceNodeID, SourceNodeClassID, DestinationNodeID, DestinationNodeClassID, Relationship from NodeEdge`

	logCannotPrepareNodeStmt              = "cannot prepare GraphNode insert statement"
	logCannotPrepareNodeAttributeStmt     = "cannot prepare NodeAttribute insert statement"
	logCannotPrepareNodeEdgeStmt          = "cannot prepare NodeEdge insert statement"
	logCannotExecuteNodeStmt              = "cannot execute GraphNode insert statement, id=[%s], [%#v]"
	logCannotExecuteNodeAttributeStmt     = "cannot execute NodeAttribute insert statement, classid=[%s], id=[%s], [%#v]"
	logCannotExecuteNodeEdgeStmt          = "cannot execute NodeEdge insert statement, classid=[%s], [%#v]"
	logAboutToCreateNode                  = "about to create GraphNode, id=[%s], [%#v]"
	logAboutToCreateNodeAttribute         = "about to create NodeAttribute, classid=[%s], id=[%s], [%#v]"
	logAboutToCreateNodeEdge              = "about to create NodeEdge, classid=[%s], nodeid=[%s], [%#v]"
	logCannotQueryNodeSelectStmt          = "cannot query Node select statement"
	logCannotQueryNodeEdgeSelectStmt      = "cannot query NodeEdge select statement"
	logCannotQueryNodeAttributeSelectStmt = "cannot query NodeAttribute select statement"
)

func SelectNodes(db *sql.DB, nodeClassID string, nodeClassNamespace string) (graph model.Nodes, err error) {
	var nodeRows *sql.Rows
	if strings.TrimSpace(nodeClassID) == "" {
		nodeRows, err = db.Query(selectNodeSQL)
	} else {
		nodeRows, err = db.Query(selectNodeSQLByClass, nodeClassID, nodeClassNamespace)
	}
	if err != nil {
		log.Error().Err(err).Msg(logCannotQueryNodeSelectStmt)
		return
	}
	defer nodeRows.Close()

	for nodeRows.Next() {
		var node model.Node

		var nodeID, classID string
		if err = nodeRows.Scan(&nodeID, &classID); err != nil {
			return
		}
		node.ID = definition.GraphNodeID(classID, nodeID)
		node.NodeClassID = classID
		graph.Nodes = append(graph.Nodes, node)
	}
	//
	//linkRows, err := db.Query(selectNodeEdgeSQL)
	//if err != nil {
	//	log.Error().Err(err).Msg(logCannotQueryNodeEdgeSelectStmt)
	//	return
	//}
	//defer linkRows.Close()
	//
	//for linkRows.Next() {
	//	var link definition.GraphLink
	//	var sourceNodeID, sourceNodeClassID, destinationNodeID, destinationNodeClassID string
	//	if err = linkRows.Scan(&sourceNodeID, &sourceNodeClassID, &destinationNodeID, &destinationNodeClassID, &link.Relationship); err != nil {
	//		return
	//	}
	//	link.Source = definition.GraphNodeID(sourceNodeClassID, sourceNodeID)
	//	link.Target = definition.GraphNodeID(destinationNodeClassID, destinationNodeID)
	//
	//	graph.Links = append(graph.Links, link)
	//}

	return
}
func ReadNodeByID(db *sql.DB, nodeClassNamespace string, nodeClassID string, nodeID string) (graph model.Nodes, err error) {
	var nodeRows *sql.Rows
	nodeRows, err = db.Query(selectNodeSQLByID, nodeClassNamespace, nodeClassID, nodeID)
	if err != nil {
		log.Error().Err(err).Msg(logCannotQueryNodeSelectStmt)
		return
	}
	defer nodeRows.Close()

	for nodeRows.Next() {
		var node model.Node
		if err = nodeRows.Scan(&node.ID, &node.NodeClassID, &node.NodeClassNamespace); err != nil {
			return
		}
		graph.Nodes = append(graph.Nodes, node)
	}

	attributeRows, err := db.Query(selectNodeAttributeSQLByNodeID, nodeID, nodeClassID, nodeClassNamespace)
	if err != nil {
		log.Error().Err(err).Msg(logCannotQueryNodeAttributeSelectStmt)
		return
	}
	defer attributeRows.Close()

	for attributeRows.Next() {
		var attribute model.NodeAttribute
		if err = attributeRows.Scan(&attribute.NodeClassAttributeID, &attribute.Value); err != nil {
			return
		}

		// TODO - lazy
		graph.Nodes[0].Attributes = append(graph.Nodes[0].Attributes, attribute)
	}

	return
}
