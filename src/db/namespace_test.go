package db

import (
	"github.com/nextmetaphor/definition-graph/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SelectNamespaces(t *testing.T) {
	conn, _ := SetupCleanDatabase()
	_ = PopulateDatabaseWithSampleData(conn)

	t.Run("SelectNamespaces", func(t *testing.T) {
		namespaces, err := SelectNamespaces(conn)

		assert.Nil(t, err)
		assert.Equal(t, len(namespaces.Namespace), 3)

		assert.Equal(t, namespaces, data.Namespaces{Namespace: []data.Namespace{
			{
				Namespace: "io.nextmetaphor",
			},
			{
				Namespace: "io.nextmetaphor.org",
			},
			{
				Namespace: "io.nextmetaphor.org.cloud",
			}},
		})
	})
}
