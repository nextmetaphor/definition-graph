package model

type (
	NodeEdgeKey struct {
		SourceNodeID             string `json:"source-node-id"`
		SourceNodeClassID        string `json:"source-node-class-id"`
		SourceNodeNamespace      string `json:"source-node-namespace"`
		DestinationNodeID        string `json:"destination-node-id"`
		DestinationNodeClassID   string `json:"destination-node-class-id"`
		DestinationNodeNamespace string `json:"destination-node-namespace"`
		Relationship             string `json:"relationship"`
	}
	
	NodeEdge struct {
		NodeEdgeKey
	}
)
