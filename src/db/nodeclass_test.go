package db

import (
	"github.com/nextmetaphor/definition-graph/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SelectNodeClass(t *testing.T) {
	conn, _ := SetupCleanDatabase()
	_ = PopulateDatabaseWithSampleData(conn)

	t.Run("SelectNodeClass", func(t *testing.T) {
		classes, err := SelectNodeClass(conn)
		assert.Nil(t, err)
		assert.Equal(t, len(classes.NodeClasses), 4)

		assert.Equal(t, classes, data.NodeClassesOuter{NodeClasses: []data.NodeClass{
			{
				ID:          "person",
				Namespace:   "io.nextmetaphor",
				Description: "A person",
			},
			{
				ID:          "bu",
				Namespace:   "io.nextmetaphor.org",
				Description: "A business unit",
			},
			{
				ID:          "company",
				Namespace:   "io.nextmetaphor.org",
				Description: "A company",
			},
			{
				ID:          "workload",
				Namespace:   "io.nextmetaphor.org.cloud",
				Description: "A workload",
			}},
		})
	})
}

func Test_CreateNodeClass(t *testing.T) {
	conn, _ := SetupCleanDatabase()

	nc := data.NodeClass{
		ID:          "nc001",
		Namespace:   "io.nextmetaphor",
		Description: "NodeClass 001",
		Attributes:  nil,
		Edges:       nil,
	}

	t.Run("CreateNodeClass", func(t *testing.T) {
		err := CreateNodeClass(conn, nc)
		assert.Nil(t, err)

		rows, _ := conn.Query("select ID, Namespace, Description from NodeClass where ID=? and Namespace = ?", nc.ID, nc.Namespace)

		assert.True(t, rows.Next())
		var nc2 data.NodeClass
		_ = rows.Scan(&nc2.ID, &nc2.Namespace, &nc2.Description)

		assert.Equal(t, nc, nc2)
		assert.False(t, rows.Next())
	})
}
