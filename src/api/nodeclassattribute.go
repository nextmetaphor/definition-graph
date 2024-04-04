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
	data, err := selectNodeClassAttributeFunc(dbConn, model.NodeClassKey{
		ID:        r.Header.Get("node-class-id"),
		Namespace: r.Header.Get("node-class-namespace"),
	})
	writeHTTPResponse(http.StatusOK, data, err, w, logCannotSelectNodeClassAttribute)
}

func createNodeClassAttributeHandler(w http.ResponseWriter, r *http.Request) {
	var nca model.NodeClassAttribute
	err := json.NewDecoder(r.Body).Decode(&nca)

	if err == nil {
		err = createNodeClassAttributeFunc(dbConn, nca)
	}
	writeHTTPResponse(http.StatusOK, nil, err, w, logCannotCreateNodeClassAttribute)
}

func readNodeClassAttributeHandler(w http.ResponseWriter, r *http.Request) {
	nca, err := readNodeClassAttributeFunc(dbConn, model.NodeClassAttributeKey{
		ID:                 r.Header.Get("id"),
		NodeClassID:        r.Header.Get("node-class-id"),
		NodeClassNamespace: r.Header.Get("node-class-namespace"),
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
		count, err := updateNodeClassAttributeFunc(dbConn, model.NodeClassAttributeKey{
			ID:                 r.Header.Get("id"),
			NodeClassID:        r.Header.Get("node-class-id"),
			NodeClassNamespace: r.Header.Get("node-class-namespace"),
		}, nce)
		if count == 0 {
			writeHTTPResponse(http.StatusNotFound, nil, err, w, logCannotUpdateNodeClassAttribute)
		} else {
			writeHTTPResponse(http.StatusOK, nil, err, w, logCannotUpdateNodeClassAttribute)
		}
	}
}

func deleteNodeClassAttributeHandler(w http.ResponseWriter, r *http.Request) {
	count, err := deleteNodeClassAttributeFunc(dbConn, model.NodeClassAttributeKey{
		ID:                 r.Header.Get("id"),
		NodeClassID:        r.Header.Get("node-class-id"),
		NodeClassNamespace: r.Header.Get("node-class-namespace"),
	})
	if count == 0 {
		writeHTTPResponse(http.StatusNotFound, nil, err, w, logCannotDeleteNodeClassAttribute)
	} else {
		writeHTTPResponse(http.StatusOK, nil, err, w, logCannotDeleteNodeClassAttribute)
	}
}
