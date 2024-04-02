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
	logCannotReadNodeClassEdge              = "cannot read node class edge"
	logCannotUpdateNodeClassEdge            = "cannot update node class edge"
	logCannotDeleteNodeClassEdge            = "cannot delete node class edge"
)

// function indirection to allow unit test stubs to be created
var (
	selectNodeClassEdgeBySourceNodeClassFunc = db.SelectNodeClassEdgeBySourceNodeClass
	createNodeClassEdgeFunc                  = db.CreateNodeClassEdge
	readNodeClassEdgeFunc                    = db.ReadNodeClassEdge
	updateNodeClassEdgeFunc                  = db.UpdateNodeClassEdge
	deleteNodeClassEdgeFunc                  = db.DeleteNodeClassEdge
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

func readNodeClassEdgeHandler(w http.ResponseWriter, r *http.Request) {
	nce, err := readNodeClassEdgeFunc(dbConn, model.NodeClassEdgeKey{
		SourceNodeClassID:             r.Header.Get("source-node-class-id"),
		SourceNodeClassNamespace:      r.Header.Get("source-node-class-namespace"),
		DestinationNodeClassID:        r.Header.Get("destination-node-class-id"),
		DestinationNodeClassNamespace: r.Header.Get("destination-node-class-namespace"),
		Relationship:                  r.Header.Get("relationship"),
	})
	if nce == nil {
		writeHTTPResponse(http.StatusNotFound, nil, err, w, logCannotReadNodeClassEdge)
	} else {
		writeHTTPResponse(http.StatusOK, nce, err, w, logCannotReadNodeClassEdge)
	}
}

func updateNodeClassEdgeHandler(w http.ResponseWriter, r *http.Request) {
	var nce model.NodeClassEdge
	err := json.NewDecoder(r.Body).Decode(&nce)

	if err != nil {
		writeHTTPResponse(http.StatusOK, nil, err, w, logCannotUpdateNodeClassEdge)
	} else {
		count, err := updateNodeClassEdgeFunc(dbConn, model.NodeClassEdgeKey{
			SourceNodeClassID:             r.Header.Get("source-node-class-id"),
			SourceNodeClassNamespace:      r.Header.Get("source-node-class-namespace"),
			DestinationNodeClassID:        r.Header.Get("destination-node-class-id"),
			DestinationNodeClassNamespace: r.Header.Get("destination-node-class-namespace"),
			Relationship:                  r.Header.Get("relationship"),
		}, nce)
		if count == 0 {
			writeHTTPResponse(http.StatusNotFound, nil, err, w, logCannotUpdateNodeClassEdge)
		} else {
			writeHTTPResponse(http.StatusOK, nil, err, w, logCannotUpdateNodeClassEdge)
		}
	}
}

func deleteNodeClassEdgeHandler(w http.ResponseWriter, r *http.Request) {
	count, err := deleteNodeClassEdgeFunc(dbConn, model.NodeClassEdgeKey{
		SourceNodeClassID:             r.Header.Get("source-node-class-id"),
		SourceNodeClassNamespace:      r.Header.Get("source-node-class-namespace"),
		DestinationNodeClassID:        r.Header.Get("destination-node-class-id"),
		DestinationNodeClassNamespace: r.Header.Get("destination-node-class-namespace"),
		Relationship:                  r.Header.Get("relationship"),
	})
	if count == 0 {
		writeHTTPResponse(http.StatusNotFound, nil, err, w, logCannotDeleteNodeClassEdge)
	} else {
		writeHTTPResponse(http.StatusOK, nil, err, w, logCannotDeleteNodeClassEdge)
	}
}
