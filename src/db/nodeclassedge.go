package db

import (
	"database/sql"
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

func SelectNodeClassEdgeBySourceNodeClass(db *sql.DB, nodeClassKey model.NodeClassKey) (nodeClassEdges model.NodeClassEdges, err error) {
	nodeClassEdges = model.NodeClassEdges{}

	rows, err := db.Query(selectNodeClassEdgeBySourceNodeClassSQL, nodeClassKey.ID, nodeClassKey.Namespace)
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

func CreateNodeClassEdge(c *sql.DB, nce model.NodeClassEdge) (e error) {
	s, e := c.Prepare(insertNodeClassEdgeSQL)
	if e != nil {
		log.Error().Err(e)
		return
	}
	_, e = s.Exec(nce.SourceNodeClassID, nce.SourceNodeClassNamespace, nce.DestinationNodeClassID, nce.DestinationNodeClassNamespace, nce.Relationship)
	if e != nil {
		log.Error().Err(e)
		return
	}

	return
}

func ReadNodeClassEdge(c *sql.DB, key model.NodeClassEdgeKey) (nce *model.NodeClassEdge, e error) {
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

func UpdateNodeClassEdge(c *sql.DB, key model.NodeClassEdgeKey, nce model.NodeClassEdge) (count int64, e error) {
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

func DeleteNodeClassEdge(c *sql.DB, key model.NodeClassEdgeKey) (count int64, e error) {
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
