package db

import (
	"database/sql"
	"github.com/nextmetaphor/definition-graph/model"
	"github.com/rs/zerolog/log"
)

const (
	selectNodeClassAttributeByNodeClassSQL = `SELECT ID, NodeClassID, NodeClassNamespace, Description, Type, IsRequired FROM NodeClassAttribute WHERE NodeClassID=? AND NodeClassNamespace=? ORDER BY ID, Description;`

	insertNodeClassAttributeSQL = `INSERT INTO NodeClassAttribute (ID, NodeClassID, NodeClassNamespace, Description, Type, IsRequired) values (?, ?, ?, ?, ?, ?);`
)

func SelectNodeClassAttributeByNodeClass(db *sql.DB, nodeClassID, nodeClassNamespace string) (nodeClassAttributes model.NodeClassAttributes, err error) {
	rows, err := db.Query(selectNodeClassAttributeByNodeClassSQL, nodeClassID, nodeClassNamespace)
	if err != nil {
		log.Error().Err(err)
		return
	}
	defer rows.Close()

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
