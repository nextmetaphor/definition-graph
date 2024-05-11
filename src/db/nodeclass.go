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
func SelectNodeClass() (nodeClasses model.NodeClasses, err error) {
	c := getDBConn()
	nodeClasses = model.NodeClasses{}
	nodeClassRows, err := c.Query(selectNodeClassSQL)
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

func CreateNodeClass(nc model.NodeClass) (e error) {
	c := getDBConn()
	s, e := c.Prepare(createNodeClassSQL)
	if e != nil {
		return
	}
	_, e = s.Exec(nc.ID, nc.Namespace, nc.Description)

	return
}

func ReadNodeClass(key model.NodeClassKey) (nc *model.NodeClass, e error) {
	c := getDBConn()
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

func UpdateNodeClass(key model.NodeClassKey, nc model.NodeClass) (count int64, e error) {
	c := getDBConn()
	s, e := c.Prepare(updateNodeClassSQL)
	if e != nil {
		return
	}
	r, e := s.Exec(nc.ID, nc.Namespace, nc.Description, key.ID, key.Namespace)
	count, _ = r.RowsAffected()

	return
}

func DeleteNodeClass(key model.NodeClassKey) (count int64, e error) {
	c := getDBConn()
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
