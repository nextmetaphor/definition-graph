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
