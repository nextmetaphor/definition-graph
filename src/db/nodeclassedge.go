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
	"github.com/rs/zerolog/log"
)

const (
	selectNodeClassEdgeSQL = `SELECT SourceNodeClassID, DestinationNodeClassID, Relationship from NodeClassEdge ORDER BY DestinationNodeClassNamespace, DestinationNodeClassID;`

	selectNodeClassEdgeBySourceNodeClassSQL = `SELECT SourceNodeClassID, SourceNodeClassNamespace, DestinationNodeClassID, DestinationNodeClassNamespace, Relationship FROM NodeClassEdge WHERE SourceNodeClassID=? AND SourceNodeClassNamespace=? ORDER BY DestinationNodeClassNamespace, DestinationNodeClassID;`

	insertNodeClassEdgeSQL = `INSERT INTO NodeClassEdge (SourceNodeClassID, SourceNodeClassNamespace, DestinationNodeClassID, DestinationNodeClassNamespace, Relationship) values (?, ?, ?, ?, ?);`
	readNodeClassEdgeSQL   = `SELECT SourceNodeClassID, SourceNodeClassNamespace, DestinationNodeClassID, DestinationNodeClassNamespace, Relationship FROM NodeClassEdge WHERE SourceNodeClassID=? AND SourceNodeClassNamespace=? AND DestinationNodeClassID=? AND DestinationNodeClassNamespace=? AND Relationship=?;`
	updateNodeClassEdgeSQL = `UPDATE NodeClassEdge SET SourceNodeClassID=?, SourceNodeClassNamespace=?, DestinationNodeClassID=?, DestinationNodeClassNamespace=?, Relationship=? WHERE SourceNodeClassID=? AND SourceNodeClassNamespace=? AND DestinationNodeClassID=? AND DestinationNodeClassNamespace=? AND Relationship=?;`
	deleteNodeClassEdgeSQL = `DELETE FROM NodeClassEdge WHERE SourceNodeClassID=? AND SourceNodeClassNamespace=? AND DestinationNodeClassID=? AND DestinationNodeClassNamespace=? AND Relationship=?;`
)

func SelectNodeClassEdgeBySourceNodeClass(nodeClassKey model.NodeClassKey) (nodeClassEdges model.NodeClassEdges, err error) {
	c := getDBConn()
	nodeClassEdges = model.NodeClassEdges{}

	rows, err := c.Query(selectNodeClassEdgeBySourceNodeClassSQL, nodeClassKey.ID, nodeClassKey.Namespace)
	if err != nil {
		log.Error().Err(err)
		return
	}
	defer rows.Close()

	var nce model.NodeClassEdge
	for rows.Next() {
		if err = rows.Scan(&nce.SourceNodeClassID, &nce.SourceNodeClassNamespace, &nce.DestinationNodeClassID, &nce.DestinationNodeClassNamespace, &nce.Relationship); err != nil {
			log.Error().Err(err)
			return
		}
		nodeClassEdges = append(nodeClassEdges, nce)
	}

	return
}

func CreateNodeClassEdge(nce model.NodeClassEdge) (e error) {
	c := getDBConn()
	s, e := c.Prepare(insertNodeClassEdgeSQL)
	if e != nil {
		log.Error().Err(e)
		return
	}
	_, e = s.Exec(nce.SourceNodeClassID, nce.SourceNodeClassNamespace, nce.DestinationNodeClassID, nce.DestinationNodeClassNamespace, nce.Relationship)
	if e != nil {
		return
	}

	return
}

func ReadNodeClassEdge(key model.NodeClassEdgeKey) (nce *model.NodeClassEdge, e error) {
	c := getDBConn()
	rows, e := c.Query(readNodeClassEdgeSQL, key.SourceNodeClassID, key.SourceNodeClassNamespace, key.DestinationNodeClassID, key.DestinationNodeClassNamespace, key.Relationship)
	if e != nil {
		return
	}

	defer rows.Close()
	if rows.Next() {
		nce = new(model.NodeClassEdge)
		e = rows.Scan(&nce.SourceNodeClassID, &nce.SourceNodeClassNamespace, &nce.DestinationNodeClassID, &nce.DestinationNodeClassNamespace, &nce.Relationship)
	}

	return
}

func UpdateNodeClassEdge(key model.NodeClassEdgeKey, nce model.NodeClassEdge) (count int64, e error) {
	c := getDBConn()
	s, e := c.Prepare(updateNodeClassEdgeSQL)
	if e != nil {
		return
	}
	r, e := s.Exec(nce.SourceNodeClassID, nce.SourceNodeClassNamespace, nce.DestinationNodeClassID, nce.DestinationNodeClassNamespace, nce.Relationship, key.SourceNodeClassID, key.SourceNodeClassNamespace, key.DestinationNodeClassID, key.DestinationNodeClassNamespace, key.Relationship)
	if r != nil {
		count, _ = r.RowsAffected()
	}

	return
}

func DeleteNodeClassEdge(key model.NodeClassEdgeKey) (count int64, e error) {
	c := getDBConn()
	s, e := c.Prepare(deleteNodeClassEdgeSQL)
	if e != nil {
		return
	}
	r, e := s.Exec(key.SourceNodeClassID, key.SourceNodeClassNamespace, key.DestinationNodeClassID, key.DestinationNodeClassNamespace, key.Relationship)
	if r != nil {
		count, _ = r.RowsAffected()
	}

	return
}
