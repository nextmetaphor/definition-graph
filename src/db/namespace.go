package db

import (
	"database/sql"
	"github.com/nextmetaphor/definition-graph/data"
	"github.com/rs/zerolog/log"
)

const (
	selectNamespacesSQL = `SELECT DISTINCT Namespace from NodeClass order by Namespace`
)

func SelectNamespaces(db *sql.DB) (namespaces data.Namespaces, err error) {
	namespaceRows, err := db.Query(selectNamespacesSQL)
	if err != nil {
		log.Error().Err(err).Msg(logCannotQueryNamespaceSelectStmt)
		return
	}
	defer namespaceRows.Close()

	for namespaceRows.Next() {
		var nodeClass data.Namespace
		if err = namespaceRows.Scan(&nodeClass.Namespace); err == nil {
			namespaces.Namespace = append(namespaces.Namespace, nodeClass)
		}
	}

	return
}
