package model

type (
	NodeClasses []NodeClass

	NodeClassKey struct {
		ID        string `json:"id"`
		Namespace string `json:"namespace"`
	}

	NodeClass struct {
		NodeClassKey
		Description string `json:"description"`
	}
)
