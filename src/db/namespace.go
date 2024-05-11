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
	selectNamespacesSQL = `SELECT DISTINCT Namespace from NodeClass order by Namespace`
)

func SelectNamespaces() (namespaces model.Namespaces, err error) {
	c := getDBConn()
	namespaces = model.Namespaces{}

	namespaceRows, err := c.Query(selectNamespacesSQL)
	if err != nil {
		log.Error().Err(err).Msg(logCannotQueryNamespaceSelectStmt)
		return
	}
	defer namespaceRows.Close()

	for namespaceRows.Next() {
		var nodeClass model.Namespace
		if err = namespaceRows.Scan(&nodeClass.Namespace); err == nil {
			namespaces = append(namespaces, nodeClass)
		}
	}

	return
}
