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

package api

import (
	"database/sql"
	"encoding/json"
	"github.com/nextmetaphor/definition-graph/model"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_selectNamespaceHandler(t *testing.T) {
	ns := model.Namespaces{
		{
			Namespace: "io.nextmetaphor",
		},
		{
			Namespace: "io.nextmetaphor.org",
		},
		{
			Namespace: "io.nextmetaphor.org.cloud",
		},
	}

	nsJSON, e := json.Marshal(ns)
	assert.Nil(t, e)

	selectNamespacesFunc = func(db *sql.DB) (namespaces model.Namespaces, err error) {
		namespaces = ns
		return
	}

	t.Run("selectNamespaceHandler", func(t *testing.T) {

		req := httptest.NewRequest("GET", "http://example.com/foo", nil)
		w := httptest.NewRecorder()
		selectNamespaceHandler(w, req)

		resp := w.Result()
		body, _ := io.ReadAll(resp.Body)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
		assert.Equal(t, nsJSON, body)
	})
}
