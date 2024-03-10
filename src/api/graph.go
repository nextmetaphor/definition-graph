package api

import (
	"fmt"
	"github.com/nextmetaphor/definition-graph/core"
	"net/http"
)

func nodeClassGraphHandler(w http.ResponseWriter, r *http.Request) {
	b, err := core.SelectNodeClassGraph(db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		fmt.Fprintf(w, string(b))
	}
}

func nodeGraphHandler(w http.ResponseWriter, r *http.Request) {
	b, err := core.SelectNodeGraph(db)

	//TODO - sort this out
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		fmt.Fprintf(w, string(b))
	}
}
