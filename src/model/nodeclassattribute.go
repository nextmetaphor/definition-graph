package model

type (
	NodeClassAttributes []NodeClassAttribute

	NodeClassAttributeKey struct {
		ID                 string `json:"ID"`
		NodeClassID        string `json:"nodeClassID"`
		NodeClassNamespace string `json:"nodeClassNamespace"`
	}

	NodeClassAttribute struct {
		NodeClassAttributeKey
		Type        string  `json:"type"`
		IsRequired  int     `json:"isRequired"`
		Description *string `json:"description,omitempty"`
	}
)
