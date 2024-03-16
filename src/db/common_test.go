package db

import (
	"database/sql"
	"fmt"
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
		fmt.Println(err)
		err.Error()
	}

	if _, err = stmt.Exec(ID, NodeClassID, NodeClassNamespace, Type, IsRequired, Description); err != nil {
		fmt.Println(err)
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

	return nil
}
