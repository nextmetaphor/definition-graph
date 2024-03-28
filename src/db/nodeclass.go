package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nextmetaphor/definition-graph/model"
	"github.com/rs/zerolog/log"
)

const (
	selectNodeClassSQL = `SELECT ID, Namespace, Description FROM NodeClass ORDER BY Namespace, ID`
	createNodeClassSQL = `INSERT INTO NodeClass (ID, Namespace, Description) VALUES (?, ?, ?);`
	readNodeClassSQL   = `SELECT ID, Namespace, Description FROM NodeClass WHERE ID=? AND Namespace=?;`
	updateNodeClassSQL = `UPDATE NodeClass SET ID=?, Namespace=?, Description=? WHERE ID=? AND Namespace=?;`
	deleteNodeClassSQL = `DELETE FROM NodeClass WHERE ID=? AND Namespace=?;`

	logCannotPrepareNodeClassStmt         = "cannot prepare NodeClass insert statement"
	logCannotPrepareNodeClassEdgeStmt     = "cannot prepare NodeClassEdge insert statement"
	logCannotExecuteNodeClassStmt         = "cannot execute NodeClass insert statement, id=[%s], [%#v]"
	logCannotExecuteNodeClassEdgeStmt     = "cannot execute NodeClassEdge insert statement, classid=[%s], [%#v]"
	logCannotQueryNodeClassSelectStmt     = "cannot query NodeClass select statement"
	logCannotQueryNamespaceSelectStmt     = "cannot query Namespace select statement"
	logCannotQueryNodeClassEdgeSelectStmt = "cannot query NodeClassEdge select statement"
)

// SelectNodeClass selects all NodeClass records from the database.
func SelectNodeClass(db *sql.DB) (nodeClasses model.NodeClasses, err error) {
	nodeClassRows, err := db.Query(selectNodeClassSQL)
	if err != nil {
		log.Error().Err(err).Msg(logCannotQueryNodeClassSelectStmt)
		return
	}
	defer nodeClassRows.Close()

	for nodeClassRows.Next() {
		var nodeClass model.NodeClass
		if err = nodeClassRows.Scan(&nodeClass.ID, &nodeClass.Namespace, &nodeClass.Description); err == nil {
			nodeClasses = append(nodeClasses, nodeClass)
		}
	}

	return
}

func CreateNodeClass(c *sql.DB, nc model.NodeClass) (e error) {
	s, e := c.Prepare(createNodeClassSQL)
	if e != nil {
		return
	}
	_, e = s.Exec(nc.ID, nc.Namespace, nc.Description)

	return
}

func ReadNodeClass(c *sql.DB, key model.NodeClassKey) (nc *model.NodeClass, e error) {
	rows, e := c.Query(readNodeClassSQL, key.ID, key.Namespace)
	if e != nil {
		return
	}

	defer rows.Close()
	if rows.Next() {
		nc = new(model.NodeClass)
		e = rows.Scan(&nc.ID, &nc.Namespace, &nc.Description)
	}

	return
}

func UpdateNodeClass(c *sql.DB, key model.NodeClassKey, nc model.NodeClass) (count int64, e error) {
	s, e := c.Prepare(updateNodeClassSQL)
	if e != nil {
		return
	}
	r, e := s.Exec(nc.ID, nc.Namespace, nc.Description, key.ID, key.Namespace)
	count, _ = r.RowsAffected()

	return
}

func DeleteNodeClass(c *sql.DB, key model.NodeClassKey) (count int64, e error) {
	s, e := c.Prepare(deleteNodeClassSQL)
	if e != nil {
		return
	}
	r, e := s.Exec(key.ID, key.Namespace)
	if e != nil {
		return
	}

	count, _ = r.RowsAffected()

	return
}
