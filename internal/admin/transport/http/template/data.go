package template

type Data struct {
	Message         string
	Error           string
	ValidationError map[string]any
	StatusCode      int
	Data            map[string]any
	Meta            map[string]any
	OldData         map[string]interface{}
}
