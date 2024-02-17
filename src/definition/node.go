package definition

type (
	// Attributes TODO
	Attributes map[string]interface{}

	// NodeEdges TODO
	NodeEdges []NodeEdge

	// NodeEdge TODO
	NodeEdge struct {
		DestinationNodeID      string `yaml:"DestinationNode"`
		DestinationNodeClassID string `yaml:"DestinationNodeClass"`
		Relationship           string `yaml:"Relationship"`
		IsToDestination        int    `yaml:"IsToDestination"`
		IsFromDestination      int    `yaml:"IsFromDestination"`
	}

	// NodeDefinition TODO
	NodeDefinition struct {
		// ClassID TODO
		ClassID string `yaml:"Class"`

		// Attributes TODO
		Attributes Attributes `yaml:"Attributes"`

		// Edges TODO
		Edges NodeEdges `yaml:"Edges"`
	}

	// NodeSpecification TODO
	NodeSpecification struct {
		// Class allows the class for all the definitions within the document to be specified.
		ClassID string `yaml:"Class,omitempty"`

		// Definitions TODO
		Definitions map[string]NodeDefinition `yaml:"Definitions,omitempty"`
	}
)
