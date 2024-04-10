package model

type (
	NodeClassEdges []NodeClassEdge

	NodeClassEdgeKey struct {
		SourceNodeClassID             string `json:"sourceNodeClassID"`
		SourceNodeClassNamespace      string `json:"sourceNodeClassNamespace"`
		DestinationNodeClassID        string `json:"destinationNodeClassID"`
		DestinationNodeClassNamespace string `json:"destinationNodeClassNamespace"`
		Relationship                  string `json:"relationship"`
	}

	NodeClassEdge struct {
		NodeClassEdgeKey
	}
)
