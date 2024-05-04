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
	"database/sql"
	"github.com/nextmetaphor/definition-graph/model"
)

func nodeClassInsert(conn *sql.DB, ID string, Namespace string, Description string) {
	stmt, err := conn.Prepare("insert into NodeClass (ID, Namespace, Description) values (?, ?, ?)")
	if err != nil {
	}
	if _, err = stmt.Exec(ID, Namespace, Description); err != nil {
		err.Error()
	}
}

func nodeClassAttributeInsert(conn *sql.DB, ID string, NodeClassID string, NodeClassNamespace string, Type string, IsRequired int, Description *string) {
	stmt, err := conn.Prepare("insert into NodeClassAttribute (ID, NodeClassID, NodeClassNamespace, Type, IsRequired, Description) values (?, ?, ?, ?, ?, ?)")
	if err != nil {
		err.Error()
	}

	if _, err = stmt.Exec(ID, NodeClassID, NodeClassNamespace, Type, IsRequired, Description); err != nil {
		err.Error()
	}
}

func nodeClassEdgeInsert(conn *sql.DB, edge model.NodeClassEdge) {
	stmt, err := conn.Prepare("insert into NodeClassEdge (SourceNodeClassID, SourceNodeClassNamespace, DestinationNodeClassID, DestinationNodeClassNamespace, Relationship) values (?, ?, ?, ?, ?)")
	if err != nil {
		err.Error()
	}

	if _, err = stmt.Exec(edge.SourceNodeClassID, edge.SourceNodeClassNamespace, edge.DestinationNodeClassID, edge.DestinationNodeClassNamespace, edge.Relationship); err != nil {
		err.Error()
	}
}

func stringPointer(s string) *string {
	return &s
}

func intPointer(i int) *int {
	return &i
}

func SetupCleanDatabase() (*sql.DB, error) {
	return OpenDatabase()
}

func PopulateDatabaseWithSampleData(conn *sql.DB) error {
	nodeClassInsert(conn, "person", "io.nextmetaphor", "A person")
	nodeClassAttributeInsert(conn, "firstname", "person", "io.nextmetaphor", "string", 0, stringPointer("first name"))
	nodeClassAttributeInsert(conn, "surname", "person", "io.nextmetaphor", "string", 1, stringPointer("second name"))
	nodeClassAttributeInsert(conn, "middle-name", "person", "io.nextmetaphor", "string", 0, nil)

	nodeClassInsert(conn, "company", "io.nextmetaphor.org", "A company")
	nodeClassInsert(conn, "bu", "io.nextmetaphor.org", "A business unit")
	nodeClassInsert(conn, "workload", "io.nextmetaphor.org.cloud", "A workload")

	nodeClassEdgeInsert(conn, model.NodeClassEdge{
		NodeClassEdgeKey: model.NodeClassEdgeKey{
			SourceNodeClassID:             "company",
			SourceNodeClassNamespace:      "io.nextmetaphor.org",
			DestinationNodeClassID:        "person",
			DestinationNodeClassNamespace: "io.nextmetaphor",
			Relationship:                  "EMPLOYS",
		}})
	nodeClassEdgeInsert(conn, model.NodeClassEdge{
		NodeClassEdgeKey: model.NodeClassEdgeKey{
			SourceNodeClassID:             "company",
			SourceNodeClassNamespace:      "io.nextmetaphor.org",
			DestinationNodeClassID:        "bu",
			DestinationNodeClassNamespace: "io.nextmetaphor.org",
			Relationship:                  "MADE_UP_OF",
		}})

	return nil
}
