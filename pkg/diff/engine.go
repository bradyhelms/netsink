package diff

type DriftItem struct {
	Field    string `json:"field"`
	Desired  any    `json:"desired"`
	Observed any    `json:"observed"`
}

// Compare two structs and return a slice of DriftItems
