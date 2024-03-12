package db

import (
	"database/sql"
	"github.com/nextmetaphor/definition-graph/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func nodeClassInsert(conn *sql.DB, ID string, Namespace string, Description string) (err error) {
	stmt, err := conn.Prepare("insert into NodeClass (ID, Namespace, Description) values (?, ?, ?)")
	if err == nil {
		_, err = stmt.Exec(ID, Namespace, Description)
	}

	return
}

func SetupCleanDatabase() (*sql.DB, error) {
	return OpenDatabase()
}

func PopulateDatabaseWithSampleData(conn *sql.DB) error {
	_ = nodeClassInsert(conn, "person", "io.nextmetaphor", "A person")
	_ = nodeClassInsert(conn, "company", "io.nextmetaphor.org", "A company")
	_ = nodeClassInsert(conn, "bu", "io.nextmetaphor.org", "A business unit")
	_ = nodeClassInsert(conn, "workload", "io.nextmetaphor.org.cloud", "A workload")

	return nil
}

func Test_SelectNodeClass(t *testing.T) {
	conn, _ := SetupCleanDatabase()
	_ = PopulateDatabaseWithSampleData(conn)

	t.Run("SelectNodeClass", func(t *testing.T) {
		classes, err := SelectNodeClass(conn)
		assert.Nil(t, err)
		assert.Equal(t, len(classes.NodeClasses), 4)

		assert.Equal(t, classes.NodeClasses[0], data.NodeClass{
			ID:          "person",
			Namespace:   "io.nextmetaphor",
			Description: "A person",
		})
		assert.Equal(t, classes.NodeClasses[1], data.NodeClass{
			ID:          "bu",
			Namespace:   "io.nextmetaphor.org",
			Description: "A business unit",
		})
		assert.Equal(t, classes.NodeClasses[2], data.NodeClass{
			ID:          "company",
			Namespace:   "io.nextmetaphor.org",
			Description: "A company",
		})
		assert.Equal(t, classes.NodeClasses[3], data.NodeClass{
			ID:          "workload",
			Namespace:   "io.nextmetaphor.org.cloud",
			Description: "A workload",
		})
	})
}
