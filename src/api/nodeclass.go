package api

import (
	"fmt"
	"github.com/nextmetaphor/definition-graph/core"
	"net/http"
)

func nodeClassHandler(w http.ResponseWriter, r *http.Request) {
	b, err := core.SelectNodeClasses(db)

	//TODO - sort this out
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		fmt.Fprintf(w, string(b))
	}
}

func selectNamespaceHandler(w http.ResponseWriter, r *http.Request) {
	b, err := core.SelectNamespaces(db)

	//TODO - sort this out
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		fmt.Fprintf(w, string(b))
	}
}
