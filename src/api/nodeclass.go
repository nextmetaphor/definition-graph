package api

import (
	"encoding/json"
	db2 "github.com/nextmetaphor/definition-graph/db"
	"github.com/nextmetaphor/definition-graph/model"
	"net/http"
)

const (
	logCannotSelectNamespaces = "cannot select namespaces"
	logCannotSelectNodeClass  = "cannot select nodeClass"
	logCannotCreateNodeClass  = "cannot create nodeClass"
	logCannotReadNodeClass    = "cannot read nodeClass"
	logCannotDeleteNodeClass  = "cannot delete nodeClass"
)

func selectNamespaceHandler(w http.ResponseWriter, r *http.Request) {
	data, err := db2.SelectNamespaces(db)
	writeHTTPResponse(data, err, w, logCannotSelectNamespaces)
}

func selectNodeClassHandler(w http.ResponseWriter, r *http.Request) {
	data, err := db2.SelectNodeClass(db)
	writeHTTPResponse(data, err, w, logCannotSelectNodeClass)
}

func createNodeClassHandler(w http.ResponseWriter, r *http.Request) {
	var nc model.NodeClass
	err := json.NewDecoder(r.Body).Decode(&nc)

	if err == nil {
		err = db2.CreateNodeClass(db, nc)
	}
	writeHTTPResponse(nil, err, w, logCannotCreateNodeClass)
}

func readNodeClassHandler(w http.ResponseWriter, r *http.Request) {
	ns := r.PathValue(entityNamespace)
	id := r.PathValue(entityNodeClass)

	nc, err := db2.ReadNodeClass(db, model.NodeClassKey{
		ID:        id,
		Namespace: ns,
	})
	writeHTTPResponse(nc, err, w, logCannotReadNodeClass)
}

func updateNodeClassHandler(w http.ResponseWriter, r *http.Request) {
	ns := r.PathValue(entityNamespace)
	id := r.PathValue(entityNodeClass)

	var nc model.NodeClass
	err := json.NewDecoder(r.Body).Decode(&nc)

	if err == nil {
		err = db2.CreateNodeClass(db, nc)
	}

	err = db2.UpdateNodeClass(db, model.NodeClassKey{
		ID:        id,
		Namespace: ns,
	}, nc)
	writeHTTPResponse(nil, err, w, logCannotReadNodeClass)
}

func deleteNodeClassHandler(w http.ResponseWriter, r *http.Request) {
	ns := r.PathValue(entityNamespace)
	id := r.PathValue(entityNodeClass)

	err := db2.DeleteNodeClass(db, model.NodeClassKey{
		ID:        id,
		Namespace: ns,
	})
	writeHTTPResponse(nil, err, w, logCannotDeleteNodeClass)
}
