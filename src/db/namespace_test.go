/*
 * Copyright (C) 2024 Paul Tatham <paul@nextmetaphor.io>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package db

import (
	"github.com/nextmetaphor/definition-graph/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SelectNamespaces(t *testing.T) {
	conn, _ := SetupCleanDatabase()
	_ = PopulateDatabaseWithSampleData(conn)

	t.Run("SelectNamespaces", func(t *testing.T) {
		namespaces, err := SelectNamespaces(conn)

		assert.Nil(t, err)
		assert.Equal(t, len(namespaces), 3)

		assert.Equal(t, namespaces, model.Namespaces{
			{
				Namespace: "io.nextmetaphor",
			},
			{
				Namespace: "io.nextmetaphor.org",
			},
			{
				Namespace: "io.nextmetaphor.org.cloud",
			},
		})
	})
}
