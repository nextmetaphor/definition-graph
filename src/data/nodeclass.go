package data

type (
	NodeClasses struct {
		NodeClasses []NodeClass `json:"nodeClasses"`
	}

	NodeClass struct {
		ID          string `json:"id"`
		Namespace   string `json:"namespace"`
		Description string `json:"description"`
	}
)
