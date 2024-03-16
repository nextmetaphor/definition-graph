package db

import (
	"github.com/nextmetaphor/definition-graph/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SelectNodeClassAttributeByNodeClass(t *testing.T) {
	conn, _ := SetupCleanDatabase()
	_ = PopulateDatabaseWithSampleData(conn)

	t.Run("SelectNodeClassAttributeByNodeClass", func(t *testing.T) {
		attributes, err := SelectNodeClassAttributeByNodeClass(conn, "person", "io.nextmetaphor")
		assert.Nil(t, err)
		assert.Equal(t, 3, len(attributes))

		assert.Equal(t, data.NodeClassAttributes{
			{
				ID:                 "firstname",
				NodeClassID:        "person",
				NodeClassNamespace: "io.nextmetaphor",
				Type:               "string",
				IsRequired:         0,
				Description:        stringPointer("first name"),
			}, {
				ID:                 "middle-name",
				NodeClassID:        "person",
				NodeClassNamespace: "io.nextmetaphor",
				Type:               "string",
				IsRequired:         0,
				Description:        nil,
			}, {
				ID:                 "surname",
				NodeClassID:        "person",
				NodeClassNamespace: "io.nextmetaphor",
				Type:               "string",
				IsRequired:         1,
				Description:        stringPointer("second name"),
			},
		}, attributes)
	})
}

func Test_CreateNodeClassAttribute(t *testing.T) {
	conn, _ := SetupCleanDatabase()
	_ = PopulateDatabaseWithSampleData(conn)

	t.Run("CreateNodeClassAttribute", func(t *testing.T) {

		err := CreateNodeClassAttribute(conn, data.NodeClassAttribute{
			ID:                 "dob",
			NodeClassID:        "person",
			NodeClassNamespace: "io.nextmetaphor",
			Type:               "date",
			IsRequired:         0,
			Description:        stringPointer("date of birth"),
		})
		err = CreateNodeClassAttribute(conn, data.NodeClassAttribute{
			ID:                 "title",
			NodeClassID:        "person",
			NodeClassNamespace: "io.nextmetaphor",
			Type:               "string",
			IsRequired:         1,
			Description:        nil,
		})
		attributes, err := SelectNodeClassAttributeByNodeClass(conn, "person", "io.nextmetaphor")
		assert.Nil(t, err)
		assert.Equal(t, 5, len(attributes))

		assert.Equal(t, data.NodeClassAttributes{
			{
				ID:                 "dob",
				NodeClassID:        "person",
				NodeClassNamespace: "io.nextmetaphor",
				Type:               "date",
				IsRequired:         0,
				Description:        stringPointer("date of birth"),
			},
			{
				ID:                 "firstname",
				NodeClassID:        "person",
				NodeClassNamespace: "io.nextmetaphor",
				Type:               "string",
				IsRequired:         0,
				Description:        stringPointer("first name"),
			}, {
				ID:                 "middle-name",
				NodeClassID:        "person",
				NodeClassNamespace: "io.nextmetaphor",
				Type:               "string",
				IsRequired:         0,
				Description:        nil,
			}, {
				ID:                 "surname",
				NodeClassID:        "person",
				NodeClassNamespace: "io.nextmetaphor",
				Type:               "string",
				IsRequired:         1,
				Description:        stringPointer("second name"),
			}, {
				ID:                 "title",
				NodeClassID:        "person",
				NodeClassNamespace: "io.nextmetaphor",
				Type:               "string",
				IsRequired:         1,
				Description:        nil,
			},
		}, attributes)
	})
}
