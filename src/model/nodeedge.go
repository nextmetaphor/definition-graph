package model

type (
	NodeEdges []NodeEdge

	NodeEdgeKey struct {
		SourceNodeID                  string `json:"sourceNodeID"`
		SourceNodeClassID             string `json:"sourceNodeClassID"`
		SourceNodeClassNamespace      string `json:"sourceNodeClassNamespace"`
		DestinationNodeID             string `json:"destinationNodeID"`
		DestinationNodeClassID        string `json:"destinationNodeClassID"`
		DestinationNodeClassNamespace string `json:"destinationNodeClassNamespace"`
		Relationship                  string `json:"relationship"`
	}

	NodeEdge struct {
		NodeEdgeKey
	}
)
