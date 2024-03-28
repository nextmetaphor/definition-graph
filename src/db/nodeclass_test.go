package db

import (
	"github.com/nextmetaphor/definition-graph/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SelectNodeClass(t *testing.T) {
	conn, _ := SetupCleanDatabase()
	_ = PopulateDatabaseWithSampleData(conn)

	t.Run("SelectNodeClass", func(t *testing.T) {
		classes, err := SelectNodeClass(conn)
		assert.Nil(t, err)
		assert.Equal(t, len(classes), 4)

		assert.Equal(t, classes, model.NodeClasses{
			{
				NodeClassKey: model.NodeClassKey{
					ID:        "person",
					Namespace: "io.nextmetaphor",
				},
				Description: "A person",
			},
			{
				NodeClassKey: model.NodeClassKey{
					ID:        "bu",
					Namespace: "io.nextmetaphor.org",
				},
				Description: "A business unit",
			},
			{
				NodeClassKey: model.NodeClassKey{
					ID:        "company",
					Namespace: "io.nextmetaphor.org",
				},
				Description: "A company",
			},
			{
				NodeClassKey: model.NodeClassKey{
					ID:        "workload",
					Namespace: "io.nextmetaphor.org.cloud",
				},
				Description: "A workload",
			},
		})
	})
}

func Test_CreateNodeClass(t *testing.T) {
	conn, _ := SetupCleanDatabase()

	nc := model.NodeClass{
		NodeClassKey: model.NodeClassKey{
			ID:        "nc001",
			Namespace: "io.nextmetaphor",
		},
		Description: "NodeClass 001",
	}

	t.Run("CreateNodeClass", func(t *testing.T) {
		err := CreateNodeClass(conn, nc)
		assert.Nil(t, err)

		rows, _ := conn.Query("select ID, Namespace, Description from NodeClass where ID=? and Namespace=?", nc.ID, nc.Namespace)

		assert.True(t, rows.Next())
		var nc2 model.NodeClass
		_ = rows.Scan(&nc2.ID, &nc2.Namespace, &nc2.Description)

		assert.Equal(t, nc.ID, nc2.ID)
		assert.Equal(t, nc.Namespace, nc2.Namespace)
		assert.Equal(t, nc.Description, nc2.Description)
		assert.False(t, rows.Next())
	})
}

func Test_ReadNodeClass(t *testing.T) {
	conn, _ := SetupCleanDatabase()
	_ = PopulateDatabaseWithSampleData(conn)

	t.Run("ReadNodeClass", func(t *testing.T) {
		nc, err := ReadNodeClass(conn, model.NodeClassKey{
			ID:        "person",
			Namespace: "io.nextmetaphor",
		})
		assert.Nil(t, err)

		assert.Equal(t, "person", nc.ID)
		assert.Equal(t, "io.nextmetaphor", nc.Namespace)
		assert.Equal(t, "A person", nc.Description)
	})
}

func Test_UpdateNodeClass(t *testing.T) {
	conn, _ := SetupCleanDatabase()
	_ = PopulateDatabaseWithSampleData(conn)

	t.Run("UpdateNodeClass-NotKey", func(t *testing.T) {
		key := model.NodeClassKey{
			ID:        "person",
			Namespace: "io.nextmetaphor",
		}
		newNodeClass := model.NodeClass{
			NodeClassKey: key,
			Description:  "NEW DESCRIPTION",
		}
		count, err := UpdateNodeClass(conn, key, newNodeClass)
		assert.Nil(t, err)
		assert.Equal(t, int64(1), count)

		rows, _ := conn.Query("select ID, Namespace, Description from NodeClass where ID=? and Namespace=?", key.ID, key.Namespace)

		assert.True(t, rows.Next())
		var updatedNodeClass model.NodeClass
		_ = rows.Scan(&updatedNodeClass.ID, &updatedNodeClass.Namespace, &updatedNodeClass.Description)

		assert.Equal(t, newNodeClass.ID, updatedNodeClass.ID)
		assert.Equal(t, newNodeClass.Namespace, updatedNodeClass.Namespace)
		assert.Equal(t, newNodeClass.Description, updatedNodeClass.Description)
		assert.False(t, rows.Next())
	})

	t.Run("UpdateNodeClass-Key", func(t *testing.T) {
		key := model.NodeClassKey{
			ID:        "person",
			Namespace: "io.nextmetaphor",
		}
		newNodeClass := model.NodeClass{
			NodeClassKey: model.NodeClassKey{
				ID:        "person2",
				Namespace: "io.nextmetaphor2",
			},
			Description: "Description 2",
		}
		count, err := UpdateNodeClass(conn, key, newNodeClass)
		assert.Nil(t, err)
		assert.Equal(t, int64(1), count)

		rows, _ := conn.Query("select ID, Namespace, Description from NodeClass where ID=? and Namespace=?", newNodeClass.ID, newNodeClass.Namespace)
		defer rows.Close()

		assert.True(t, rows.Next())
		var updatedNodeClass model.NodeClass
		_ = rows.Scan(&updatedNodeClass.ID, &updatedNodeClass.Namespace, &updatedNodeClass.Description)

		assert.Equal(t, newNodeClass.ID, updatedNodeClass.ID)
		assert.Equal(t, newNodeClass.Namespace, updatedNodeClass.Namespace)
		assert.Equal(t, newNodeClass.Description, updatedNodeClass.Description)
		assert.False(t, rows.Next())
	})
}

func Test_DeleteNodeClass(t *testing.T) {
	conn, _ := SetupCleanDatabase()
	_ = PopulateDatabaseWithSampleData(conn)

	t.Run("DeleteNodeClass", func(t *testing.T) {
		key := model.NodeClassKey{
			ID:        "workload",
			Namespace: "io.nextmetaphor.org.cloud",
		}

		rows, _ := conn.Query("select ID, Namespace, Description from NodeClass where ID=? and Namespace=?", key.ID, key.Namespace)
		assert.True(t, rows.Next())
		rows.Close()

		count, err := DeleteNodeClass(conn, key)
		assert.Nil(t, err)
		assert.Equal(t, int64(1), count)

		rows, _ = conn.Query("select ID, Namespace, Description from NodeClass where ID=? and Namespace=?", key.ID, key.Namespace)
		assert.False(t, rows.Next())
		rows.Close()
	})
}
