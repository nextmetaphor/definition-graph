package db

import (
	"database/sql"
	"github.com/nextmetaphor/definition-graph/model"
)

const (
	selectNodeEdgeBySourceNodeSQL = `SELECT SourceNodeID, SourceNodeClassID, SourceNodeClassNamespace, DestinationNodeID, DestinationNodeClassID, DestinationNodeClassNamespace, Relationship FROM NodeEdge WHERE SourceNodeID=? AND SourceNodeClassID=? AND SourceNodeClassNamespace=? ORDER BY DestinationNodeID, DestinationNodeClassNamespace, DestinationNodeClassID;`

	insertNodeEdgeSQL = `INSERT INTO NodeEdge (SourceNodeID, SourceNodeClassID, SourceNodeClassNamespace, DestinationNodeID, DestinationNodeClassID, DestinationNodeClassNamespace, Relationship) values (?, ?, ?, ?, ?, ?, ?);`
	readNodeEdgeSQL   = `SELECT SourceNodeID, SourceNodeClassID, SourceNodeClassNamespace, DestinationNodeID, DestinationNodeClassID, DestinationNodeClassNamespace, Relationship FROM NodeEdge WHERE SourceNodeID=? AND SourceNodeClassID=? AND SourceNodeClassNamespace=? AND DestinationNodeID=? AND DestinationNodeClassID=? AND DestinationNodeClassNamespace=? AND Relationship=?;`
	updateNodeEdgeSQL = `UPDATE NodeEdge SET SourceNodeID=?, SourceNodeClassID=?, SourceNodeClassNamespace=?, DestinationNodeID=?, DestinationNodeClassID=?, DestinationNodeClassNamespace=?, Relationship=? WHERE SourceNodeID=? AND SourceNodeClassID=? AND SourceNodeClassNamespace=? AND DestinationNodeID=? AND DestinationNodeClassID=? AND DestinationNodeClassNamespace=? AND Relationship=?;`
	deleteNodeEdgeSQL = `DELETE FROM NodeEdge WHERE SourceNodeID=? AND SourceNodeClassID=? AND SourceNodeClassNamespace=? AND DestinationNodeID=? AND DestinationNodeClassID=? AND DestinationNodeClassNamespace=? AND Relationship=?;`
)

func SelectNodeEdgeBySourceNode(db *sql.DB, nodeKey model.NodeKey) (nodeEdges model.NodeEdges, err error) {
	nodeEdges = model.NodeEdges{}

	rows, err := db.Query(selectNodeEdgeBySourceNodeSQL, nodeKey.ID, nodeKey.NodeClassID, nodeKey.NodeClassNamespace)
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

func CreateNodeEdge(c *sql.DB, ne model.NodeEdge) (e error) {
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

func ReadNodeEdge(c *sql.DB, key model.NodeEdgeKey) (ne *model.NodeEdge, e error) {
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

func UpdateNodeEdge(c *sql.DB, key model.NodeEdgeKey, ne model.NodeEdge) (count int64, e error) {
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

func DeleteNodeEdge(c *sql.DB, key model.NodeEdgeKey) (count int64, e error) {
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
