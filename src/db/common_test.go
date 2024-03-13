package db

import (
	"database/sql"
)

func nodeClassInsert(conn *sql.DB, ID string, Namespace string, Description string) {
	stmt, err := conn.Prepare("insert into NodeClass (ID, Namespace, Description) values (?, ?, ?)")
	if err != nil {
	}
	if _, err = stmt.Exec(ID, Namespace, Description); err != nil {
		err.Error()
	}
}

func nodeClassAttributeInsert(conn *sql.DB, ID string, NodeClassID string, NodeClassNamespace string, Type string, IsRequired string, Description string) {
	stmt, err := conn.Prepare("insert into NodeClassAttribute (ID, NodeClassID, NodeClassNamespace, Type, IsRequired, Description) values (?, ?, ?, ?, ?, ?)")
	if err != nil {
		err.Error()
	}

	if _, err = stmt.Exec(ID, NodeClassID, NodeClassNamespace, Type, IsRequired, Description); err != nil {
		err.Error()
	}
}

func SetupCleanDatabase() (*sql.DB, error) {
	return OpenDatabase()
}

func PopulateDatabaseWithSampleData(conn *sql.DB) error {
	nodeClassInsert(conn, "person", "io.nextmetaphor", "A person")
	nodeClassAttributeInsert(conn, "firstname", "person", "io.nextmetaphor", "string", "1", "first name")
	nodeClassAttributeInsert(conn, "surname", "person", "io.nextmetaphor", "string", "1", "second name")

	nodeClassInsert(conn, "company", "io.nextmetaphor.org", "A company")
	nodeClassInsert(conn, "bu", "io.nextmetaphor.org", "A business unit")
	nodeClassInsert(conn, "workload", "io.nextmetaphor.org.cloud", "A workload")

	return nil
}
