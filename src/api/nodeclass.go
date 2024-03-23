package api

import (
	db2 "github.com/nextmetaphor/definition-graph/db"
	"net/http"
)

const (
	logCannotSelectNamespaces = "cannot select namespaces"
	logCannotSelectNodeClass  = "cannot select nodeClass"
)

func nodeClassHandler(w http.ResponseWriter, r *http.Request) {
	data, err := db2.SelectNodeClass(db)
	writeHTTPResponse(data, err, w, logCannotSelectNodeClass)
}

func selectNamespaceHandler(w http.ResponseWriter, r *http.Request) {
	data, err := db2.SelectNamespaces(db)
	writeHTTPResponse(data, err, w, logCannotSelectNamespaces)
}
