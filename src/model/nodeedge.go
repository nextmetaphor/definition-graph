package model

type (
	NodeEdges []NodeEdge

	NodeEdgeKey struct {
		SourceNodeID                  string `json:"source-node-id"`
		SourceNodeClassID             string `json:"source-node-class-id"`
		SourceNodeClassNamespace      string `json:"source-node-class-namespace"`
		DestinationNodeID             string `json:"destination-node-id"`
		DestinationNodeClassID        string `json:"destination-node-class-id"`
		DestinationNodeClassNamespace string `json:"destination-node-class-namespace"`
		Relationship                  string `json:"relationship"`
	}

	NodeEdge struct {
		NodeEdgeKey
	}
)
