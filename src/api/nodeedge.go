package api

import (
	"encoding/json"
	"github.com/nextmetaphor/definition-graph/db"
	"github.com/nextmetaphor/definition-graph/model"
	"net/http"
)

const (
	logSelectNodeEdgeBySourceNode = "cannot select node edge by source node class"
	logCannotCreateNodeEdge       = "cannot create node edge"
	logCannotReadNodeEdge         = "cannot read node edge"
	logCannotUpdateNodeEdge       = "cannot update node edge"
	logCannotDeleteNodeEdge       = "cannot delete node edge"
)

// function indirection to allow unit test stubs to be created
var (
	selectNodeEdgeBySourceNodeFunc = db.SelectNodeEdgeBySourceNode
	createNodeEdgeFunc             = db.CreateNodeEdge
	readNodeEdgeFunc               = db.ReadNodeEdge
	updateNodeEdgeFunc             = db.UpdateNodeEdge
	deleteNodeEdgeFunc             = db.DeleteNodeEdge
)

func selectNodeEdgeBySourceNodeHandler(w http.ResponseWriter, r *http.Request) {
	data, err := selectNodeEdgeBySourceNodeFunc(dbConn, model.NodeKey{
		ID:                 r.Header.Get(paramSourceNodeID),
		NodeClassID:        r.Header.Get(paramSourceNodeClassID),
		NodeClassNamespace: r.Header.Get(paramSourceNodeClassNamespace),
	})
	writeHTTPResponse(http.StatusOK, data, err, w, logSelectNodeEdgeBySourceNode)
}

func createNodeEdgeHandler(w http.ResponseWriter, r *http.Request) {
	var ne model.NodeEdge
	err := json.NewDecoder(r.Body).Decode(&ne)

	if err == nil {
		err = createNodeEdgeFunc(dbConn, ne)
	}
	writeHTTPResponse(http.StatusOK, nil, err, w, logCannotCreateNodeEdge)
}

func readNodeEdgeHandler(w http.ResponseWriter, r *http.Request) {
	nce, err := readNodeEdgeFunc(dbConn, model.NodeEdgeKey{
		SourceNodeID:                  r.Header.Get(paramSourceNodeID),
		SourceNodeClassID:             r.Header.Get(paramSourceNodeClassID),
		SourceNodeClassNamespace:      r.Header.Get(paramSourceNodeClassNamespace),
		DestinationNodeID:             r.Header.Get(paramDestinationNodeID),
		DestinationNodeClassID:        r.Header.Get(paramDestinationNodeClassID),
		DestinationNodeClassNamespace: r.Header.Get(paramDestinationNodeClassNamespace),
		Relationship:                  r.Header.Get(paramRelationship),
	})
	if nce == nil {
		writeHTTPResponse(http.StatusNotFound, nil, err, w, logCannotReadNodeEdge)
	} else {
		writeHTTPResponse(http.StatusOK, nce, err, w, logCannotReadNodeEdge)
	}
}

func updateNodeEdgeHandler(w http.ResponseWriter, r *http.Request) {
	var nce model.NodeEdge
	err := json.NewDecoder(r.Body).Decode(&nce)

	if err != nil {
		writeHTTPResponse(http.StatusOK, nil, err, w, logCannotUpdateNodeEdge)
	} else {
		count, err := updateNodeEdgeFunc(dbConn, model.NodeEdgeKey{
			SourceNodeID:                  r.Header.Get(paramSourceNodeID),
			SourceNodeClassID:             r.Header.Get(paramSourceNodeClassID),
			SourceNodeClassNamespace:      r.Header.Get(paramSourceNodeClassNamespace),
			DestinationNodeID:             r.Header.Get(paramDestinationNodeID),
			DestinationNodeClassID:        r.Header.Get(paramDestinationNodeClassID),
			DestinationNodeClassNamespace: r.Header.Get(paramDestinationNodeClassNamespace),
			Relationship:                  r.Header.Get(paramRelationship),
		}, nce)
		if count == 0 {
			writeHTTPResponse(http.StatusNotFound, nil, err, w, logCannotUpdateNodeEdge)
		} else {
			writeHTTPResponse(http.StatusOK, nil, err, w, logCannotUpdateNodeEdge)
		}
	}
}

func deleteNodeEdgeHandler(w http.ResponseWriter, r *http.Request) {
	count, err := deleteNodeEdgeFunc(dbConn, model.NodeEdgeKey{
		SourceNodeID:                  r.Header.Get(paramSourceNodeID),
		SourceNodeClassID:             r.Header.Get(paramSourceNodeClassID),
		SourceNodeClassNamespace:      r.Header.Get(paramSourceNodeClassNamespace),
		DestinationNodeID:             r.Header.Get(paramDestinationNodeID),
		DestinationNodeClassID:        r.Header.Get(paramDestinationNodeClassID),
		DestinationNodeClassNamespace: r.Header.Get(paramDestinationNodeClassNamespace),
		Relationship:                  r.Header.Get(paramRelationship),
	})
	if count == 0 {
		writeHTTPResponse(http.StatusNotFound, nil, err, w, logCannotDeleteNodeEdge)
	} else {
		writeHTTPResponse(http.StatusOK, nil, err, w, logCannotDeleteNodeEdge)
	}
}
