package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nextmetaphor/definition-graph/model"
	"github.com/rs/zerolog/log"
	"strings"
)

const (
	// TODO used by definitions - refactor
	insertNodeSQL        = `INSERT INTO Node (ID, NodeClassID) values (?, ?);`
	createNodeSQL        = `INSERT INTO Node (ID, NodeClassID, NodeClassNamespace) values (?, ?, ?);`
	selectNodeSQL        = `SELECT ID, NodeClassID, NodeClassNamespace FROM Node ORDER BY NodeClassNamespace, ID`
	selectNodeSQLByClass = `SELECT ID, NodeClassID, NodeClassNamespace FROM Node WHERE NodeClassID=? AND NodeClassNameSpace=? ORDER BY NodeClassNamespace, ID`
	selectNodeSQLByID    = `SELECT ID, NodeClassID, NodeClassNamespace FROM Node WHERE ID=? AND NodeClassID=? AND NodeClassNameSpace=?`
	updateNodeSQL        = `UPDATE Node SET ID=?, NodeClassID=?, NodeClassNamespace=? WHERE ID=? AND NodeClassID=? AND NodeClassNamespace=?;`
	deleteNodeSQL        = `DELETE FROM Node WHERE ID=? AND NodeClassID=? AND NodeClassNamespace=?;`

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

func CreateNode(c *sql.DB, nc model.Node) (e error) {
	s, e := c.Prepare(createNodeSQL)
	if e != nil {
		return
	}
	_, e = s.Exec(nc.ID, nc.NodeClassID, nc.NodeClassNamespace)

	return
}
func ReadNode(db *sql.DB, nodeKey model.NodeKey) (node *model.Node, err error) {
	var nodeRows *sql.Rows
	nodeRows, err = db.Query(selectNodeSQLByID, nodeKey.ID, nodeKey.NodeClassID, nodeKey.NodeClassNamespace)
	if err != nil {
		log.Error().Err(err).Msg(logCannotQueryNodeSelectStmt)
		return
	}
	defer nodeRows.Close()

	if nodeRows.Next() {
		node = new(model.Node)
		err = nodeRows.Scan(&node.ID, &node.NodeClassID, &node.NodeClassNamespace)
	}

	return
}

func UpdateNode(c *sql.DB, key model.NodeKey, nc model.Node) (count int64, e error) {
	s, e := c.Prepare(updateNodeSQL)
	if e != nil {
		return
	}
	r, e := s.Exec(nc.ID, nc.NodeClassID, nc.NodeClassNamespace, key.ID, key.NodeClassID, key.NodeClassNamespace)
	count, _ = r.RowsAffected()

	return
}

func DeleteNode(c *sql.DB, key model.NodeKey) (count int64, e error) {
	s, e := c.Prepare(deleteNodeSQL)
	if e != nil {
		return
	}
	r, e := s.Exec(key.ID, key.NodeClassID, key.NodeClassNamespace)
	if e != nil {
		return
	}

	count, _ = r.RowsAffected()

	return
}
