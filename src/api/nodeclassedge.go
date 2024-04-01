package api

import (
	"encoding/json"
	"github.com/nextmetaphor/definition-graph/db"
	"github.com/nextmetaphor/definition-graph/model"
	"net/http"
)

const (
	logSelectNodeClassEdgeBySourceNodeClass = "cannot select node class edge by source node class"
	logCannotCreateNodeClassEdge            = "cannot create node class edge"
)

// function indirection to allow unit test stubs to be created
var (
	selectNodeClassEdgeBySourceNodeClassFunc = db.SelectNodeClassEdgeBySourceNodeClass
	createNodeClassEdgeFunc                  = db.CreateNodeClassEdge
)

func selectNodeClassEdgeBySourceNodeClassHandler(w http.ResponseWriter, r *http.Request) {
	nodeClassKey := model.NodeClassKey{
		ID:        r.Header.Get("source-node-class-id"),
		Namespace: r.Header.Get("source-node-class-namespace"),
	}

	data, err := selectNodeClassEdgeBySourceNodeClassFunc(dbConn, nodeClassKey)
	writeHTTPResponse(http.StatusOK, data, err, w, logSelectNodeClassEdgeBySourceNodeClass)
}

func createNodeClassEdgeHandler(w http.ResponseWriter, r *http.Request) {
	var nce model.NodeClassEdge
	err := json.NewDecoder(r.Body).Decode(&nce)

	if err == nil {
		err = createNodeClassEdgeFunc(dbConn, nce)
	}
	writeHTTPResponse(http.StatusOK, nil, err, w, logCannotCreateNodeClassEdge)
}
