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

//return true if form is valid
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
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
		f.Errors.Add(field, "This field can not be blank ")
		return false
	}
	return true
}
