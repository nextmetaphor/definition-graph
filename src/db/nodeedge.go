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
	"github.com/nextmetaphor/definition-graph/model"
)

const (
	selectNodeEdgeBySourceNodeSQL = `SELECT SourceNodeID, SourceNodeClassID, SourceNodeClassNamespace, DestinationNodeID, DestinationNodeClassID, DestinationNodeClassNamespace, Relationship FROM NodeEdge WHERE SourceNodeID=? AND SourceNodeClassID=? AND SourceNodeClassNamespace=? ORDER BY DestinationNodeID, DestinationNodeClassNamespace, DestinationNodeClassID;`

	insertNodeEdgeSQL = `INSERT INTO NodeEdge (SourceNodeID, SourceNodeClassID, SourceNodeClassNamespace, DestinationNodeID, DestinationNodeClassID, DestinationNodeClassNamespace, Relationship) values (?, ?, ?, ?, ?, ?, ?);`
	readNodeEdgeSQL   = `SELECT SourceNodeID, SourceNodeClassID, SourceNodeClassNamespace, DestinationNodeID, DestinationNodeClassID, DestinationNodeClassNamespace, Relationship FROM NodeEdge WHERE SourceNodeID=? AND SourceNodeClassID=? AND SourceNodeClassNamespace=? AND DestinationNodeID=? AND DestinationNodeClassID=? AND DestinationNodeClassNamespace=? AND Relationship=?;`
	updateNodeEdgeSQL = `UPDATE NodeEdge SET SourceNodeID=?, SourceNodeClassID=?, SourceNodeClassNamespace=?, DestinationNodeID=?, DestinationNodeClassID=?, DestinationNodeClassNamespace=?, Relationship=? WHERE SourceNodeID=? AND SourceNodeClassID=? AND SourceNodeClassNamespace=? AND DestinationNodeID=? AND DestinationNodeClassID=? AND DestinationNodeClassNamespace=? AND Relationship=?;`
	deleteNodeEdgeSQL = `DELETE FROM NodeEdge WHERE SourceNodeID=? AND SourceNodeClassID=? AND SourceNodeClassNamespace=? AND DestinationNodeID=? AND DestinationNodeClassID=? AND DestinationNodeClassNamespace=? AND Relationship=?;`
)

func SelectNodeEdgeBySourceNode(nodeKey model.NodeKey) (nodeEdges model.NodeEdges, err error) {
	c := getDBConn()
	nodeEdges = model.NodeEdges{}

	rows, err := c.Query(selectNodeEdgeBySourceNodeSQL, nodeKey.ID, nodeKey.NodeClassID, nodeKey.NodeClassNamespace)
	if err != nil {
		return
	}
	defer rows.Close()

	var ne model.NodeEdge
	for rows.Next() {
		if err = rows.Scan(&ne.SourceNodeID, &ne.SourceNodeClassID, &ne.SourceNodeClassNamespace, &ne.DestinationNodeID, &ne.DestinationNodeClassID, &ne.DestinationNodeClassNamespace, &ne.Relationship); err != nil {
			return
		}
		nodeEdges = append(nodeEdges, ne)
	}

	return
}

func CreateNodeEdge(ne model.NodeEdge) (e error) {
	c := getDBConn()
	s, e := c.Prepare(insertNodeEdgeSQL)
	if e != nil {
		return
	}
	_, e = s.Exec(ne.SourceNodeID, ne.SourceNodeClassID, ne.SourceNodeClassNamespace, ne.DestinationNodeID, ne.DestinationNodeClassID, ne.DestinationNodeClassNamespace, ne.Relationship)
	if e != nil {
		return
	}

	return
}

func ReadNodeEdge(key model.NodeEdgeKey) (ne *model.NodeEdge, e error) {
	c := getDBConn()
	rows, e := c.Query(readNodeEdgeSQL, key.SourceNodeID, key.SourceNodeClassID, key.SourceNodeClassNamespace, key.DestinationNodeID, key.DestinationNodeClassID, key.DestinationNodeClassNamespace, key.Relationship)
	if e != nil {
		return
	}

	defer rows.Close()
	if rows.Next() {
		ne = new(model.NodeEdge)
		e = rows.Scan(&ne.SourceNodeID, &ne.SourceNodeClassID, &ne.SourceNodeClassNamespace, &ne.DestinationNodeID, &ne.DestinationNodeClassID, &ne.DestinationNodeClassNamespace, &ne.Relationship)
	}

	return
}

func UpdateNodeEdge(key model.NodeEdgeKey, ne model.NodeEdge) (count int64, e error) {
	c := getDBConn()
	s, e := c.Prepare(updateNodeEdgeSQL)
	if e != nil {
		return
	}
	r, e := s.Exec(ne.SourceNodeID, ne.SourceNodeClassID, ne.SourceNodeClassNamespace, ne.DestinationNodeID, ne.DestinationNodeClassID, ne.DestinationNodeClassNamespace, ne.Relationship, key.SourceNodeID, key.SourceNodeClassID, key.SourceNodeClassNamespace, key.DestinationNodeID, key.DestinationNodeClassID, key.DestinationNodeClassNamespace, key.Relationship)
	if r != nil {
		count, _ = r.RowsAffected()
	}

	return
}

func DeleteNodeEdge(key model.NodeEdgeKey) (count int64, e error) {
	c := getDBConn()
	s, e := c.Prepare(deleteNodeEdgeSQL)
	if e != nil {
		return
	}
	r, e := s.Exec(key.SourceNodeID, key.SourceNodeClassID, key.SourceNodeClassNamespace, key.DestinationNodeID, key.DestinationNodeClassID, key.DestinationNodeClassNamespace, key.Relationship)
	if r != nil {
		count, _ = r.RowsAffected()
	}

	return
}
