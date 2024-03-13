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
