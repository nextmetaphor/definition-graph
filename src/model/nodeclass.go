package model

type (
	NodeClasses []NodeClass

	NodeClass struct {
		ID          string `json:"id"`
		Namespace   string `json:"namespace"`
		Description string `json:"description"`
	}
)
