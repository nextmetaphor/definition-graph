package db

import (
	"database/sql"
	"github.com/nextmetaphor/definition-graph/model"
	"github.com/rs/zerolog/log"
)

const (
	selectNodeEdgeSQL = `SELECT SourceNodeID, SourceNodeClassID, SourceNodeClassNamespace, DestinationNodeID, DestinationNodeClassID, DestinationNodeClassNamespace, Relationship from NodeEdge;`

	selectNodeEdgeBySourceNodeSQL = `SELECT SourceNodeID, SourceNodeClassID, SourceNodeClassNamespace, DestinationNodeID, DestinationNodeClassID, DestinationNodeClassNamespace, Relationship FROM NodeEdge WHERE SourceNodeID=? AND SourceNodeClassID=? AND SourceNodeClassNamespace=? ORDER BY DestinationNodeID, DestinationNodeClassNamespace, DestinationNodeClassID;`

	insertNodeEdgeSQL = `INSERT INTO NodeEdge (SourceNodeID, SourceNodeClassID, SourceNodeClassNamespace, DestinationNodeID, DestinationNodeClassID, DestinationNodeClassNamespace, Relationship) values (?, ?, ?, ?, ?, ?, ?);`
	readNodeEdgeSQL   = `SELECT SourceNodeID, SourceNodeClassID, SourceNodeClassNamespace, DestinationNodeID, DestinationNodeClassID, DestinationNodeClassNamespace, Relationship FROM NodeEdge WHERE SourceNodeID=? AND SourceNodeClassID=? AND SourceNodeClassNamespace=? AND DestinationNodeID=? AND DestinationNodeClassID=? AND DestinationNodeClassNamespace=? AND Relationship=?;`
	updateNodeEdgeSQL = `UPDATE NodeEdge SET SourceNodeID=?, SourceNodeClassID=?, SourceNodeClassNamespace=?, DestinationNodeID=?, DestinationNodeClassID=?, DestinationNodeClassNamespace=?, Relationship=? WHERE SourceNodeID=? AND SourceNodeClassID=? AND SourceNodeClassNamespace=? AND DestinationNodeID=? AND DestinationNodeClassID=? AND DestinationNodeClassNamespace=? AND Relationship=?;`
	deleteNodeEdgeSQL = `DELETE FROM NodeEdge WHERE SourceNodeID=? AND SourceNodeClassID=? AND SourceNodeClassNamespace=? AND DestinationNodeID=? AND DestinationNodeClassID=? AND DestinationNodeClassNamespace=? AND Relationship=?;`
)

func SelectNodeEdgeBySourceNode(db *sql.DB, nodeClassKey model.NodeClassKey) (nodeClassEdges model.NodeClassEdges, err error) {
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

func CreateNodeEdge(c *sql.DB, nce model.NodeClassEdge) (e error) {
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

func ReadNodeEdge(c *sql.DB, key model.NodeClassEdgeKey) (nce model.NodeClassEdge, e error) {
	rows, e := c.Query(readNodeClassEdgeSQL, key.SourceNodeClassID, key.SourceNodeClassNamespace, key.DestinationNodeClassID, key.DestinationNodeClassNamespace, key.Relationship)
	if e != nil {
		return
	}

	defer rows.Close()
	if rows.Next() {
		e = rows.Scan(&nce.SourceNodeClassID, &nce.SourceNodeClassNamespace, &nce.DestinationNodeClassID, &nce.DestinationNodeClassNamespace, &nce.Relationship)
	}

	return
}

func UpdateNodeEdge(c *sql.DB, key model.NodeClassEdgeKey, nce model.NodeClassEdge) (e error) {
	s, e := c.Prepare(updateNodeClassEdgeSQL)
	if e != nil {
		return
	}
	_, e = s.Exec(nce.SourceNodeClassID, nce.SourceNodeClassNamespace, nce.DestinationNodeClassID, nce.DestinationNodeClassNamespace, nce.Relationship, key.SourceNodeClassID, key.SourceNodeClassNamespace, key.DestinationNodeClassID, key.DestinationNodeClassNamespace, key.Relationship)

	return
}

func DeleteNodeEdge(c *sql.DB, key model.NodeClassEdgeKey) (e error) {
	s, e := c.Prepare(deleteNodeClassEdgeSQL)
	if e != nil {
		return
	}
	_, e = s.Exec(key.SourceNodeClassID, key.SourceNodeClassNamespace, key.DestinationNodeClassID, key.DestinationNodeClassNamespace, key.Relationship)

	return
}
