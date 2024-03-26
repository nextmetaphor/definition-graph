package model

type (
	NodeClassEdges []NodeClassEdge

	NodeClassEdgeKey struct {
		SourceNodeClassID             string `json:"source-node-class-id"`
		SourceNodeClassNamespace      string `json:"source-node-class-namespace"`
		DestinationNodeClassID        string `json:"destination-node-class-id"`
		DestinationNodeClassNamespace string `json:"destination-node-class-namespace"`
		Relationship                  string `json:"relationship"`
	}

	NodeClassEdge struct {
		NodeClassEdgeKey
	}
)
