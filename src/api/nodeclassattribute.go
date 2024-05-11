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
	logCannotSelectNodeClassAttribute = "cannot select nodeClassAttribute"
	logCannotCreateNodeClassAttribute = "cannot create nodeClassAttribute"
	logCannotReadNodeClassAttribute   = "cannot read nodeClassAttribute"
	logCannotUpdateNodeClassAttribute = "cannot update node class attribute"
	logCannotDeleteNodeClassAttribute = "cannot delete nodeClassAttribute"
)

// function indirection to allow unit test stubs to be created
var (
	selectNodeClassAttributeFunc = db.SelectNodeClassAttributeByNodeClass
	createNodeClassAttributeFunc = db.CreateNodeClassAttribute
	readNodeClassAttributeFunc   = db.ReadNodeClassAttribute
	updateNodeClassAttributeFunc = db.UpdateNodeClassAttribute
	deleteNodeClassAttributeFunc = db.DeleteNodeClassAttribute
)

func selectNodeClassAttributeHandler(w http.ResponseWriter, r *http.Request) {
	data, err := selectNodeClassAttributeFunc(model.NodeClassKey{
		ID:        r.Header.Get(paramNodeClassID),
		Namespace: r.Header.Get(paramNodeClassNamespace),
	})
	writeHTTPResponse(http.StatusOK, data, err, w, logCannotSelectNodeClassAttribute)
}

func createNodeClassAttributeHandler(w http.ResponseWriter, r *http.Request) {
	var nca model.NodeClassAttribute
	err := json.NewDecoder(r.Body).Decode(&nca)

	if err == nil {
		err = createNodeClassAttributeFunc(nca)
	}
	writeHTTPResponse(http.StatusOK, nil, err, w, logCannotCreateNodeClassAttribute)
}

func readNodeClassAttributeHandler(w http.ResponseWriter, r *http.Request) {
	nca, err := readNodeClassAttributeFunc(model.NodeClassAttributeKey{
		ID:                 r.Header.Get(paramID),
		NodeClassID:        r.Header.Get(paramNodeClassID),
		NodeClassNamespace: r.Header.Get(paramNodeClassNamespace),
	})
	if nca == nil {
		writeHTTPResponse(http.StatusNotFound, nil, err, w, logCannotReadNodeClassAttribute)
	} else {
		writeHTTPResponse(http.StatusOK, nca, err, w, logCannotReadNodeClassAttribute)
	}
}

func updateNodeClassAttributeHandler(w http.ResponseWriter, r *http.Request) {
	var nce model.NodeClassAttribute
	err := json.NewDecoder(r.Body).Decode(&nce)

	if err != nil {
		writeHTTPResponse(http.StatusOK, nil, err, w, logCannotUpdateNodeClassAttribute)
	} else {
		count, err := updateNodeClassAttributeFunc(model.NodeClassAttributeKey{
			ID:                 r.Header.Get(paramID),
			NodeClassID:        r.Header.Get(paramNodeClassID),
			NodeClassNamespace: r.Header.Get(paramNodeClassNamespace),
		}, nce)
		if count == 0 {
			writeHTTPResponse(http.StatusNotFound, nil, err, w, logCannotUpdateNodeClassAttribute)
		} else {
			writeHTTPResponse(http.StatusOK, nil, err, w, logCannotUpdateNodeClassAttribute)
		}
	}
}

func deleteNodeClassAttributeHandler(w http.ResponseWriter, r *http.Request) {
	count, err := deleteNodeClassAttributeFunc(model.NodeClassAttributeKey{
		ID:                 r.Header.Get(paramID),
		NodeClassID:        r.Header.Get(paramNodeClassID),
		NodeClassNamespace: r.Header.Get(paramNodeClassNamespace),
	})
	if count == 0 {
		writeHTTPResponse(http.StatusNotFound, nil, err, w, logCannotDeleteNodeClassAttribute)
	} else {
		writeHTTPResponse(http.StatusOK, nil, err, w, logCannotDeleteNodeClassAttribute)
	}
}
