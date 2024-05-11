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
	"encoding/json"
	"github.com/nextmetaphor/definition-graph/db"
	"github.com/nextmetaphor/definition-graph/model"
	"net/http"
)

const (
	logCannotSelectNodeAttribute = "cannot select node attribute"
	logCannotCreateNodeAttribute = "cannot create node attribute"
	logCannotReadNodeAttribute   = "cannot read node attribute"
	logCannotUpdateNodeAttribute = "cannot update node attribute"
	logCannotDeleteNodeAttribute = "cannot delete node attribute"
)

// function indirection to allow unit test stubs to be created
var (
	selectNodeAttributeFunc = db.SelectNodeAttributeByNode
	createNodeAttributeFunc = db.CreateNodeAttribute
	readNodeAttributeFunc   = db.ReadNodeAttribute
	updateNodeAttributeFunc = db.UpdateNodeAttribute
	deleteNodeAttributeFunc = db.DeleteNodeAttribute
)

func selectNodeAttributeHandler(w http.ResponseWriter, r *http.Request) {
	data, err := selectNodeAttributeFunc(model.NodeKey{
		ID:                 r.Header.Get(paramNodeID),
		NodeClassID:        r.Header.Get(paramNodeClassID),
		NodeClassNamespace: r.Header.Get(paramNodeClassNamespace),
	})
	writeHTTPResponse(http.StatusOK, data, err, w, logCannotSelectNodeAttribute)
}

func createNodeAttributeHandler(w http.ResponseWriter, r *http.Request) {
	var na model.NodeAttribute
	err := json.NewDecoder(r.Body).Decode(&na)

	if err == nil {
		err = createNodeAttributeFunc(na)
	}
	writeHTTPResponse(http.StatusOK, nil, err, w, logCannotCreateNodeAttribute)
}

func readNodeAttributeHandler(w http.ResponseWriter, r *http.Request) {
	na, err := readNodeAttributeFunc(model.NodeAttributeKey{
		NodeID:               r.Header.Get(paramNodeID),
		NodeClassID:          r.Header.Get(paramNodeClassID),
		NodeClassNamespace:   r.Header.Get(paramNodeClassNamespace),
		NodeClassAttributeID: r.Header.Get(paramNodeClassAttributeID),
	})
	if na == nil {
		writeHTTPResponse(http.StatusNotFound, nil, err, w, logCannotReadNodeAttribute)
	} else {
		writeHTTPResponse(http.StatusOK, na, err, w, logCannotReadNodeAttribute)
	}
}

func updateNodeAttributeHandler(w http.ResponseWriter, r *http.Request) {
	var na model.NodeAttribute
	err := json.NewDecoder(r.Body).Decode(&na)

	if err != nil {
		writeHTTPResponse(http.StatusOK, nil, err, w, logCannotUpdateNodeAttribute)
	} else {
		count, err := updateNodeAttributeFunc(model.NodeAttributeKey{
			NodeID:               r.Header.Get(paramNodeID),
			NodeClassID:          r.Header.Get(paramNodeClassID),
			NodeClassNamespace:   r.Header.Get(paramNodeClassNamespace),
			NodeClassAttributeID: r.Header.Get(paramNodeClassAttributeID),
		}, na)
		if count == 0 {
			writeHTTPResponse(http.StatusNotFound, nil, err, w, logCannotUpdateNodeAttribute)
		} else {
			writeHTTPResponse(http.StatusOK, nil, err, w, logCannotUpdateNodeAttribute)
		}
	}
}

func deleteNodeAttributeHandler(w http.ResponseWriter, r *http.Request) {
	count, err := deleteNodeAttributeFunc(model.NodeAttributeKey{
		NodeID:               r.Header.Get(paramNodeID),
		NodeClassID:          r.Header.Get(paramNodeClassID),
		NodeClassNamespace:   r.Header.Get(paramNodeClassNamespace),
		NodeClassAttributeID: r.Header.Get(paramNodeClassAttributeID),
	})
	if count == 0 {
		writeHTTPResponse(http.StatusNotFound, nil, err, w, logCannotDeleteNodeAttribute)
	} else {
		writeHTTPResponse(http.StatusOK, nil, err, w, logCannotDeleteNodeAttribute)
	}
}
