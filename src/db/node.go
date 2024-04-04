package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nextmetaphor/definition-graph/model"
	"github.com/rs/zerolog/log"
	"strings"
)

const (
	insertNodeSQL        = `INSERT INTO Node (ID, NodeClassID) values (?, ?);`
	selectNodeSQL        = `SELECT ID, NodeClassID, NodeClassNamespace from Node`
	selectNodeSQLByClass = `SELECT ID, NodeClassID, NodeClassNamespace from Node where NodeClassID=? and NodeClassNameSpace=?`
	selectNodeSQLByID    = `SELECT ID, NodeClassID, NodeClassNamespace from Node where NodeClassNameSpace=? and NodeClassID=? and ID=?`

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

func SelectNodes(db *sql.DB, nodeClassKey model.NodeClassKey) (nodes model.Nodes, err error) {
	nodes = model.Nodes{}

	var nodeRows *sql.Rows
	if strings.TrimSpace(nodeClassKey.Namespace) == "" {
		nodeRows, err = db.Query(selectNodeSQL)
	} else {
		nodeRows, err = db.Query(selectNodeSQLByClass, nodeClassKey.ID, nodeClassKey.Namespace)
	}
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
		nodes = append(nodes, node)
	}

	return
}
func ReadNodeByID(db *sql.DB, nodeKey model.NodeKey) (nodes model.Nodes, err error) {
	var nodeRows *sql.Rows
	nodeRows, err = db.Query(selectNodeSQLByID, nodeKey.NodeClassNamespace, nodeKey.NodeClassID, nodeKey.ID)
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
		nodes = append(nodes, node)
	}

	return
}
