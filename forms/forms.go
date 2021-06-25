//this package will hold all form information

package forms

import (
	"net/http"
	"net/url"
)

//creates custum form struct, embeds a url.values object
type Form struct {
	url.Values
	Errors errors
}

//initalizes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

//checking if form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	//check request
	x := r.Form.Get(field)
	if x == "" {
		return false
	}
	return true
}
