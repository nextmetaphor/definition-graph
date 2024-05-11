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
	"github.com/nextmetaphor/definition-graph/core"
	"net/http"
)

const (
	logCannotLoadNodeClassDefinitions = "cannot load nodeClass definitions"
	logCannotSaveNodeClassDefinitions = "cannot save nodeClass definitions"
	logCannotLoadNodeDefinitions      = "cannot load node definitions"
	logCannotSaveNodeDefinitions      = "cannot save node definitions"
)

func loadNodeClassDefinitionsHandler(w http.ResponseWriter, r *http.Request) {
	e := core.LoadNodeClassDefinitions(r.Header.Values(paramNodeClassDefinitionsDirectory), r.Header.Get(paramDefinitionFormat))

	writeHTTPResponse(http.StatusOK, nil, e, w, logCannotLoadNodeClassDefinitions)
}

func saveNodeClassDefinitionsHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func loadNodeDefinitionsHandler(w http.ResponseWriter, r *http.Request) {
	var e error
	if e = core.LoadNodeDefinitionsWithoutEdges(r.Header.Values(paramNodeDefinitionsDirectory), r.Header.Get(paramDefinitionFormat)); e == nil {
		e = core.LoadNodeDefinitionsOnlyEdges(r.Header.Values(paramNodeDefinitionsDirectory), r.Header.Get(paramDefinitionFormat))
	}

	writeHTTPResponse(http.StatusOK, nil, e, w, logCannotLoadNodeDefinitions)
}

func saveNodeDefinitionsHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
}
