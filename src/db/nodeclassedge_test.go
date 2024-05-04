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

func Test_SelectNodeClassEdgeBySourceNodeClass(t *testing.T) {
	conn, _ := SetupCleanDatabase()
	_ = PopulateDatabaseWithSampleData(conn)

	t.Run("SelectNodeClassEdgeBySourceNodeClass", func(t *testing.T) {
		attributes, err := SelectNodeClassEdgeBySourceNodeClass(conn, model.NodeClassKey{
			ID:        "company",
			Namespace: "io.nextmetaphor.org"})
		assert.Nil(t, err)
		assert.Equal(t, 2, len(attributes))

		assert.Equal(t, model.NodeClassEdges{
			{
				NodeClassEdgeKey: model.NodeClassEdgeKey{
					SourceNodeClassID:             "company",
					SourceNodeClassNamespace:      "io.nextmetaphor.org",
					DestinationNodeClassID:        "person",
					DestinationNodeClassNamespace: "io.nextmetaphor",
					Relationship:                  "EMPLOYS",
				},
			}, {
				NodeClassEdgeKey: model.NodeClassEdgeKey{
					SourceNodeClassID:             "company",
					SourceNodeClassNamespace:      "io.nextmetaphor.org",
					DestinationNodeClassID:        "bu",
					DestinationNodeClassNamespace: "io.nextmetaphor.org",
					Relationship:                  "MADE_UP_OF",
				},
			},
		}, attributes)
	})
}

//func Test_CreateNodeClassAttribute(t *testing.T) {
//	conn, _ := SetupCleanDatabase()
//	_ = PopulateDatabaseWithSampleData(conn)
//
//	t.Run("CreateNodeClassAttribute", func(t *testing.T) {
//
//		err := CreateNodeClassAttribute(conn, model.NodeClassAttribute{
//			NodeClassAttributeKey: model.NodeClassAttributeKey{
//				ID:                 "dob",
//				NodeClassID:        "person",
//				NodeClassNamespace: "io.nextmetaphor",
//			},
//			Type:        "date",
//			IsRequired:  0,
//			Description: stringPointer("date of birth"),
//		})
//		err = CreateNodeClassAttribute(conn, model.NodeClassAttribute{
//			NodeClassAttributeKey: model.NodeClassAttributeKey{
//				ID:                 "title",
//				NodeClassID:        "person",
//				NodeClassNamespace: "io.nextmetaphor",
//			},
//			Type:        "string",
//			IsRequired:  1,
//			Description: nil,
//		})
//		attributes, err := SelectNodeClassAttributeByNodeClass(conn, model.NodeClassKey{
//			ID:        "person",
//			Namespace: "io.nextmetaphor",
//		})
//		assert.Nil(t, err)
//		assert.Equal(t, 5, len(attributes))
//
//		assert.Equal(t, model.NodeClassAttributes{
//			{
//				NodeClassAttributeKey: model.NodeClassAttributeKey{
//					ID:                 "dob",
//					NodeClassID:        "person",
//					NodeClassNamespace: "io.nextmetaphor",
//				},
//				Type:        "date",
//				IsRequired:  0,
//				Description: stringPointer("date of birth"),
//			},
//			{
//				NodeClassAttributeKey: model.NodeClassAttributeKey{
//					ID:                 "firstname",
//					NodeClassID:        "person",
//					NodeClassNamespace: "io.nextmetaphor",
//				},
//				Type:        "string",
//				IsRequired:  0,
//				Description: stringPointer("first name"),
//			}, {
//				NodeClassAttributeKey: model.NodeClassAttributeKey{
//					ID:                 "middle-name",
//					NodeClassID:        "person",
//					NodeClassNamespace: "io.nextmetaphor",
//				},
//				Type:        "string",
//				IsRequired:  0,
//				Description: nil,
//			}, {
//				NodeClassAttributeKey: model.NodeClassAttributeKey{
//					ID:                 "surname",
//					NodeClassID:        "person",
//					NodeClassNamespace: "io.nextmetaphor",
//				},
//				Type:        "string",
//				IsRequired:  1,
//				Description: stringPointer("second name"),
//			}, {
//				NodeClassAttributeKey: model.NodeClassAttributeKey{
//					ID:                 "title",
//					NodeClassID:        "person",
//					NodeClassNamespace: "io.nextmetaphor",
//				},
//				Type:        "string",
//				IsRequired:  1,
//				Description: nil,
//			},
//		}, attributes)
//	})
//}
//
//func Test_ReadNodeAttributeClass(t *testing.T) {
//	conn, _ := SetupCleanDatabase()
//	_ = PopulateDatabaseWithSampleData(conn)
//
//	t.Run("ReadNodeClassAttribute_AllFields", func(t *testing.T) {
//		nca, err := ReadNodeClassAttribute(conn, model.NodeClassAttributeKey{
//			ID:                 "firstname",
//			NodeClassID:        "person",
//			NodeClassNamespace: "io.nextmetaphor",
//		})
//		assert.Nil(t, err)
//
//		assert.Equal(t, "firstname", nca.ID)
//		assert.Equal(t, "person", nca.NodeClassID)
//		assert.Equal(t, "io.nextmetaphor", nca.NodeClassNamespace)
//		assert.Equal(t, "string", nca.Type)
//		assert.Equal(t, 0, nca.IsRequired)
//		assert.Equal(t, stringPointer("first name"), nca.Description)
//	})
//
//	t.Run("ReadNodeClassAttribute_IncludingNulls", func(t *testing.T) {
//		nca, err := ReadNodeClassAttribute(conn, model.NodeClassAttributeKey{
//			ID:                 "middle-name",
//			NodeClassID:        "person",
//			NodeClassNamespace: "io.nextmetaphor",
//		})
//		assert.Nil(t, err)
//
//		assert.Equal(t, "middle-name", nca.ID)
//		assert.Equal(t, "person", nca.NodeClassID)
//		assert.Equal(t, "io.nextmetaphor", nca.NodeClassNamespace)
//		assert.Equal(t, "string", nca.Type)
//		assert.Equal(t, 0, nca.IsRequired)
//		assert.Nil(t, nca.Description)
//	})
//}
//
//func Test_UpdateNodeClassAttribute(t *testing.T) {
//	conn, _ := SetupCleanDatabase()
//	_ = PopulateDatabaseWithSampleData(conn)
//
//	t.Run("UpdateNodeClassAttribute-NotKey", func(t *testing.T) {
//		key := model.NodeClassAttributeKey{
//			ID:                 "firstname",
//			NodeClassID:        "person",
//			NodeClassNamespace: "io.nextmetaphor",
//		}
//		newNodeClassAttribute := model.NodeClassAttribute{
//			NodeClassAttributeKey: key,
//			Type:                  "integer",
//			IsRequired:            1,
//			Description:           stringPointer("NEW DESCRIPTION"),
//		}
//		err := UpdateNodeClassAttribute(conn, key, newNodeClassAttribute)
//		assert.Nil(t, err)
//
//		rows, _ := conn.Query("SELECT ID, NodeClassID, NodeClassNamespace, Type, IsRequired, Description from NodeClassAttribute WHERE ID=? AND NodeClassID=? AND NodeClassNamespace=?", key.ID, key.NodeClassID, key.NodeClassNamespace)
//
//		assert.True(t, rows.Next())
//		var updatedNodeClassAttribute model.NodeClassAttribute
//		_ = rows.Scan(&updatedNodeClassAttribute.ID, &updatedNodeClassAttribute.NodeClassID,
//			&updatedNodeClassAttribute.NodeClassNamespace, &updatedNodeClassAttribute.Type,
//			&updatedNodeClassAttribute.IsRequired, &updatedNodeClassAttribute.Description)
//
//		assert.Equal(t, newNodeClassAttribute.ID, updatedNodeClassAttribute.ID)
//		assert.Equal(t, newNodeClassAttribute.NodeClassID, updatedNodeClassAttribute.NodeClassID)
//		assert.Equal(t, newNodeClassAttribute.NodeClassNamespace, updatedNodeClassAttribute.NodeClassNamespace)
//		assert.Equal(t, newNodeClassAttribute.Type, updatedNodeClassAttribute.Type)
//		assert.Equal(t, newNodeClassAttribute.IsRequired, updatedNodeClassAttribute.IsRequired)
//		assert.Equal(t, newNodeClassAttribute.Description, updatedNodeClassAttribute.Description)
//		assert.False(t, rows.Next())
//	})
//
//	t.Run("UpdateNodeClassAttribute-NotKeyWithNil", func(t *testing.T) {
//		key := model.NodeClassAttributeKey{
//			ID:                 "firstname",
//			NodeClassID:        "person",
//			NodeClassNamespace: "io.nextmetaphor",
//		}
//		newNodeClassAttribute := model.NodeClassAttribute{
//			NodeClassAttributeKey: key,
//			Type:                  "integer",
//			IsRequired:            1,
//			Description:           nil,
//		}
//		err := UpdateNodeClassAttribute(conn, key, newNodeClassAttribute)
//		assert.Nil(t, err)
//
//		rows, _ := conn.Query("SELECT ID, NodeClassID, NodeClassNamespace, Type, IsRequired, Description from NodeClassAttribute WHERE ID=? AND NodeClassID=? AND NodeClassNamespace=?", key.ID, key.NodeClassID, key.NodeClassNamespace)
//
//		assert.True(t, rows.Next())
//		var updatedNodeClassAttribute model.NodeClassAttribute
//		_ = rows.Scan(&updatedNodeClassAttribute.ID, &updatedNodeClassAttribute.NodeClassID,
//			&updatedNodeClassAttribute.NodeClassNamespace, &updatedNodeClassAttribute.Type,
//			&updatedNodeClassAttribute.IsRequired, &updatedNodeClassAttribute.Description)
//
//		assert.Equal(t, newNodeClassAttribute.ID, updatedNodeClassAttribute.ID)
//		assert.Equal(t, newNodeClassAttribute.NodeClassID, updatedNodeClassAttribute.NodeClassID)
//		assert.Equal(t, newNodeClassAttribute.NodeClassNamespace, updatedNodeClassAttribute.NodeClassNamespace)
//		assert.Equal(t, newNodeClassAttribute.Type, updatedNodeClassAttribute.Type)
//		assert.Equal(t, newNodeClassAttribute.IsRequired, updatedNodeClassAttribute.IsRequired)
//		assert.Nil(t, updatedNodeClassAttribute.Description)
//		assert.False(t, rows.Next())
//	})
//
//	t.Run("UpdateNodeClass-Key", func(t *testing.T) {
//		key := model.NodeClassAttributeKey{
//			ID:                 "firstname",
//			NodeClassID:        "person",
//			NodeClassNamespace: "io.nextmetaphor",
//		}
//		newNodeClassAttribute := model.NodeClassAttribute{
//			NodeClassAttributeKey: model.NodeClassAttributeKey{
//				ID:                 "firstname2",
//				NodeClassID:        "company",
//				NodeClassNamespace: "io.nextmetaphor.org",
//			},
//			Type:        "integer",
//			IsRequired:  1,
//			Description: stringPointer("NEW DESCRIPTION"),
//		}
//		err := UpdateNodeClassAttribute(conn, key, newNodeClassAttribute)
//		assert.Nil(t, err)
//
//		rows, _ := conn.Query("SELECT ID, NodeClassID, NodeClassNamespace, Type, IsRequired, Description from NodeClassAttribute WHERE ID=? AND NodeClassID=? AND NodeClassNamespace=?", newNodeClassAttribute.ID, newNodeClassAttribute.NodeClassID, newNodeClassAttribute.NodeClassNamespace)
//
//		assert.True(t, rows.Next())
//		var updatedNodeClassAttribute model.NodeClassAttribute
//		_ = rows.Scan(&updatedNodeClassAttribute.ID, &updatedNodeClassAttribute.NodeClassID,
//			&updatedNodeClassAttribute.NodeClassNamespace, &updatedNodeClassAttribute.Type,
//			&updatedNodeClassAttribute.IsRequired, &updatedNodeClassAttribute.Description)
//
//		assert.Equal(t, newNodeClassAttribute.ID, updatedNodeClassAttribute.ID)
//		assert.Equal(t, newNodeClassAttribute.NodeClassID, updatedNodeClassAttribute.NodeClassID)
//		assert.Equal(t, newNodeClassAttribute.NodeClassNamespace, updatedNodeClassAttribute.NodeClassNamespace)
//		assert.Equal(t, newNodeClassAttribute.Type, updatedNodeClassAttribute.Type)
//		assert.Equal(t, newNodeClassAttribute.IsRequired, updatedNodeClassAttribute.IsRequired)
//		assert.Equal(t, newNodeClassAttribute.Description, updatedNodeClassAttribute.Description)
//		assert.False(t, rows.Next())
//	})
//}
//
//func Test_DeleteNodeClassAttribute(t *testing.T) {
//	conn, _ := SetupCleanDatabase()
//	_ = PopulateDatabaseWithSampleData(conn)
//
//	t.Run("DeleteNodeClassAttribute", func(t *testing.T) {
//		key := model.NodeClassAttributeKey{
//			ID:                 "middle-name",
//			NodeClassID:        "person",
//			NodeClassNamespace: "io.nextmetaphor",
//		}
//
//		rows, _ := conn.Query("select * from NodeClassAttribute where ID=? and NodeClassID=? and NodeClassNamespace=?", key.ID, key.NodeClassID, key.NodeClassNamespace)
//		assert.True(t, rows.Next())
//		rows.Close()
//
//		err := DeleteNodeClassAttribute(conn, key)
//		assert.Nil(t, err)
//
//		rows, _ = conn.Query("select * from NodeClassAttribute where ID=? and NodeClassID=? and NodeClassNamespace=?", key.ID, key.NodeClassID, key.NodeClassNamespace)
//		assert.False(t, rows.Next())
//		rows.Close()
//	})
//}
