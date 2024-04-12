package model

type (
	NodeAttributes []NodeAttribute

	NodeAttributeKey struct {
		NodeID               string `json:"nodeID"`
		NodeClassID          string `json:"nodeClassID"`
		NodeClassNamespace   string `json:"nodeClassNamespace"`
		NodeClassAttributeID string `json:"nodeClassAttributeID"`
	}

	NodeAttribute struct {
		NodeAttributeKey
		Value string `json:"value"`
	}
)
