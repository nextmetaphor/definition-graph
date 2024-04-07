package db

import (
	"database/sql"
	"github.com/nextmetaphor/definition-graph/model"
	"github.com/rs/zerolog/log"
)

const (
	selectNodeAttributeByNodeSQL = `SELECT NodeID, NodeClassID, NodeClassNamespace, NodeClassAttributeID, Value FROM NodeAttribute WHERE NodeID=? AND NodeClassID=? AND NodeClassNamespace=? ORDER BY NodeClassAttributeID;`

	insertNodeAttributeSQL = `INSERT INTO NodeAttribute (NodeID, NodeClassID, NodeClassNamespace, NodeClassAttributeID, Value) values (?, ?, ?, ?, ?);`
	readNodeAttributeSQL   = `SELECT NodeID, NodeClassID, NodeClassNamespace, NodeClassAttributeID, Value FROM NodeAttribute WHERE NodeID=? AND NodeClassID=? AND NodeClassNamespace=? AND NodeClassAttributeID=?;`
	updateNodeAttributeSQL = `UPDATE NodeAttribute SET NodeID=?, NodeClassID=?, NodeClassNamespace=?, NodeClassAttributeID=?, Value=? WHERE NodeID=? AND NodeClassID=? AND NodeClassNamespace=? AND NodeClassAttributeID=?;`
	deleteNodeAttributeSQL = `DELETE FROM NodeAttribute WHERE NodeID=? AND NodeClassID=? AND NodeClassNamespace=? AND NodeClassAttributeID=?;`
)

func SelectNodeAttributeByNode(db *sql.DB, nodeKey model.NodeKey) (nodeAttributes model.NodeAttributes, err error) {
	nodeAttributes = model.NodeAttributes{}

	rows, err := db.Query(selectNodeAttributeByNodeSQL, nodeKey.ID, nodeKey.NodeClassID, nodeKey.NodeClassNamespace)
	if err != nil {
		log.Error().Err(err)
		return
	}
	defer rows.Close()

	var na model.NodeAttribute
	for rows.Next() {
		if err = rows.Scan(&na.NodeID, &na.NodeClassID, &na.NodeClassNamespace, &na.NodeClassAttributeID, &na.Value); err != nil {
			log.Error().Err(err)
			return
		}
		nodeAttributes = append(nodeAttributes, na)
	}

	return
}

func CreateNodeAttribute(c *sql.DB, na model.NodeAttribute) (e error) {
	s, e := c.Prepare(insertNodeAttributeSQL)
	if e != nil {
		log.Error().Err(e)
		return
	}
	_, e = s.Exec(na.NodeID, na.NodeClassID, na.NodeClassNamespace, na.NodeClassAttributeID, na.Value)
	if e != nil {
		log.Error().Err(e)
		return
	}

	return
}

func ReadNodeAttribute(c *sql.DB, key model.NodeAttributeKey) (na *model.NodeAttribute, e error) {
	rows, e := c.Query(readNodeAttributeSQL, key.NodeID, key.NodeClassID, key.NodeClassNamespace, key.NodeClassAttributeID)
	if e != nil {
		return
	}

	defer rows.Close()
	if rows.Next() {
		na = new(model.NodeAttribute)
		e = rows.Scan(&na.NodeID, &na.NodeClassID, &na.NodeClassNamespace, &na.NodeClassAttributeID, &na.Value)
	}

	return
}

func UpdateNodeAttribute(c *sql.DB, key model.NodeAttributeKey, na model.NodeAttribute) (count int64, e error) {
	s, e := c.Prepare(updateNodeAttributeSQL)
	if e != nil {
		return
	}
	r, e := s.Exec(na.NodeID, na.NodeClassID, na.NodeClassNamespace, na.NodeClassAttributeID, na.Value, key.NodeID,
		key.NodeClassID, key.NodeClassNamespace, key.NodeClassAttributeID)
	if r != nil {
		count, _ = r.RowsAffected()
	}

	return
}

func DeleteNodeAttribute(c *sql.DB, key model.NodeAttributeKey) (count int64, e error) {
	s, e := c.Prepare(deleteNodeAttributeSQL)
	if e != nil {
		return
	}
	r, e := s.Exec(key.NodeID, key.NodeClassID, key.NodeClassNamespace, key.NodeClassAttributeID)
	if r != nil {
		count, _ = r.RowsAffected()
	}

	return
}
