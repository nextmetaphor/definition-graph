package db

import (
	"database/sql"
	"github.com/nextmetaphor/definition-graph/model"
	"github.com/rs/zerolog/log"
)

const (
	selectNodeClassAttributeByNodeClassSQL = `SELECT ID, NodeClassID, NodeClassNamespace, Description, Type, IsRequired FROM NodeClassAttribute WHERE NodeClassID=? AND NodeClassNamespace=? ORDER BY ID, Description;`

	insertNodeClassAttributeSQL = `INSERT INTO NodeClassAttribute (ID, NodeClassID, NodeClassNamespace, Description, Type, IsRequired) values (?, ?, ?, ?, ?, ?);`
	readNodeClassAttributeSQL   = `SELECT ID, NodeClassID, NodeClassNamespace, Description, Type, IsRequired FROM NodeClassAttribute WHERE ID=? AND NodeClassID=? AND NodeClassNamespace=?;`
	updateNodeClassAttributeSQL = `UPDATE NodeClassAttribute SET ID=?, NodeClassID=?, NodeClassNamespace=?, Description=?, Type=?, IsRequired=? WHERE ID=? AND NodeClassID=? AND NodeClassNamespace=?;`
	deleteNodeClassAttributeSQL = `DELETE FROM NodeClassAttribute WHERE ID=? AND NodeClassID=? AND NodeClassNamespace=?;`
)

func SelectNodeClassAttributeByNodeClass(db *sql.DB, nodeClassKey model.NodeClassKey) (nodeClassAttributes model.NodeClassAttributes, err error) {
	rows, err := db.Query(selectNodeClassAttributeByNodeClassSQL, nodeClassKey.ID, nodeClassKey.Namespace)
	if err != nil {
		log.Error().Err(err)
		return
	}
	defer rows.Close()

	nodeClassAttributes = []model.NodeClassAttribute{}
	var nca model.NodeClassAttribute
	for rows.Next() {
		if err = rows.Scan(&nca.ID, &nca.NodeClassID, &nca.NodeClassNamespace, &nca.Description, &nca.Type, &nca.IsRequired); err != nil {
			log.Error().Err(err)
			return
		}
		nodeClassAttributes = append(nodeClassAttributes, nca)
	}

	return
}

func CreateNodeClassAttribute(c *sql.DB, nca model.NodeClassAttribute) (e error) {
	s, e := c.Prepare(insertNodeClassAttributeSQL)
	if e != nil {
		log.Error().Err(e)
		return
	}
	_, e = s.Exec(nca.ID, nca.NodeClassID, nca.NodeClassNamespace, nca.Description, nca.Type, nca.IsRequired)
	if e != nil {
		log.Error().Err(e)
		return
	}

	return
}

func ReadNodeClassAttribute(c *sql.DB, key model.NodeClassAttributeKey) (nca *model.NodeClassAttribute, e error) {
	rows, e := c.Query(readNodeClassAttributeSQL, key.ID, key.NodeClassID, key.NodeClassNamespace)
	if e != nil {
		return
	}

	defer rows.Close()
	if rows.Next() {
		nca = new(model.NodeClassAttribute)
		e = rows.Scan(&nca.ID, &nca.NodeClassID, &nca.NodeClassNamespace, &nca.Description, &nca.Type, &nca.IsRequired)
	}

	return
}

func UpdateNodeClassAttribute(c *sql.DB, key model.NodeClassAttributeKey, nca model.NodeClassAttribute) (count int64, e error) {
	s, e := c.Prepare(updateNodeClassAttributeSQL)
	if e != nil {
		return
	}
	r, e := s.Exec(nca.ID, nca.NodeClassID, nca.NodeClassNamespace, nca.Description, nca.Type, nca.IsRequired, key.ID,
		key.NodeClassID, key.NodeClassNamespace)
	if r != nil {
		count, _ = r.RowsAffected()
	}

	return
}

func DeleteNodeClassAttribute(c *sql.DB, key model.NodeClassAttributeKey) (count int64, e error) {
	s, e := c.Prepare(deleteNodeClassAttributeSQL)
	if e != nil {
		return
	}
	r, e := s.Exec(key.ID, key.NodeClassID, key.NodeClassNamespace)
	if r != nil {
		count, _ = r.RowsAffected()
	}

	return
}
