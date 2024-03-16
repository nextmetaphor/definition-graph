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
		assert.Equal(t, len(classes), 4)

		assert.Equal(t, classes, data.NodeClasses{
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
			},
		})
	})
}

func Test_CreateNodeClass(t *testing.T) {
	conn, _ := SetupCleanDatabase()

	nc := data.NodeClass{
		ID:          "nc001",
		Namespace:   "io.nextmetaphor",
		Description: "NodeClass 001",
		//Attributes: []data.NodeClassAttribute{
		//	{
		//		ID:          "att1",
		//		Type:        "string",
		//		IsRequired:  0,
		//		Description: "Attribute 1",
		//	},
		//	{
		//		ID:          "att2",
		//		Type:        "int",
		//		IsRequired:  1,
		//		Description: "Attribute 2",
		//	},
		//},
		//Edges: nil,
	}

	t.Run("CreateNodeClass", func(t *testing.T) {
		err := CreateNodeClass(conn, nc)
		assert.Nil(t, err)

		rows, _ := conn.Query("select ID, Namespace, Description from NodeClass where ID=? and Namespace=?", nc.ID, nc.Namespace)

		assert.True(t, rows.Next())
		var nc2 data.NodeClass
		_ = rows.Scan(&nc2.ID, &nc2.Namespace, &nc2.Description)

		assert.Equal(t, nc.ID, nc2.ID)
		assert.Equal(t, nc.Namespace, nc2.Namespace)
		assert.Equal(t, nc.Description, nc2.Description)
		assert.False(t, rows.Next())

		rows, _ = conn.Query("select ID, NodeClassID, NodeClassNamespace, Type, IsRequired, Description from NodeClassAttribute where NodeClassID=? and NodeClassNamespace=? order by NodeClassID, NodeClassNamespace, ID", nc.ID, nc.Namespace)

		//var att data.NodeClassAttribute
		//assert.True(t, rows.Next())
		//_ = rows.Scan(&att.ID, &att.NodeClassID, &att.NodeClassNamespace, &att.Type, &att.IsRequired, &att.Description)
		//
		//assert.Equal(t, nc.Attributes[0].ID, att.ID)
		//assert.Equal(t, nc.ID, att.NodeClassID)
		//assert.Equal(t, nc.Namespace, att.NodeClassNamespace)
		//assert.Equal(t, nc.Attributes[0].Type, att.Type)
		//assert.Equal(t, nc.Attributes[0].IsRequired, att.IsRequired)
		//assert.Equal(t, nc.Attributes[0].Description, att.Description)
		//
		//assert.True(t, rows.Next())
		//_ = rows.Scan(&att.ID, &att.NodeClassID, &att.NodeClassNamespace, &att.Type, &att.IsRequired, &att.Description)
		//
		//assert.Equal(t, nc.Attributes[1].ID, att.ID)
		//assert.Equal(t, nc.ID, att.NodeClassID)
		//assert.Equal(t, nc.Namespace, att.NodeClassNamespace)
		//assert.Equal(t, nc.Attributes[1].Type, att.Type)
		//assert.Equal(t, nc.Attributes[1].IsRequired, att.IsRequired)
		//assert.Equal(t, nc.Attributes[1].Description, att.Description)
		//assert.False(t, rows.Next())
	})
}
