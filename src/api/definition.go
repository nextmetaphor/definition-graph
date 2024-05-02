package api

import (
	"github.com/nextmetaphor/definition-graph/core"
	"net/http"
)

const (
	logCannotLoadDefinitions = "cannot load definitions"
	logCannotSaveDefinitions = "cannot save definitions"
)

func loadDefinitionsHandler(w http.ResponseWriter, r *http.Request) {
	var e error
	if e = core.LoadNodeClassDefinitions([]string{"../definition/nodeClass"}, "yaml", dbConn); e == nil {
		if e = core.LoadNodeDefinitionsWithoutEdges([]string{"../definition/node"}, "yaml", dbConn); e == nil {
			e = core.LoadNodeDefinitionsOnlyEdges([]string{"../definition/node"}, "yaml", dbConn)
		}
	}

	writeHTTPResponse(http.StatusOK, nil, e, w, logCannotLoadDefinitions)
}

func saveDefinitionsHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
}
