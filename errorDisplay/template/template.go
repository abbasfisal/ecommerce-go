package template

import (
	"github.com/abbasfisal/ecommerce-go/errorDisplay/forms"
	"net/http"
)

type Data struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
	req       *http.Request
}
