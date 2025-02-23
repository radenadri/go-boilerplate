package response

type ValidationError struct {
	Field  string `json:"field"`
	Rule   string `json:"rule"`
	Value  string `json:"value,omitempty"`
	Reason string `json:"reason"`
}
