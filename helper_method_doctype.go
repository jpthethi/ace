package ace

import "io"

// helperMethodDoctype represents a helper method doctype.
type helperMethodDoctype struct {
	elementBase
}

// WriteTo writes data to w.
func (e *helperMethodDoctype) WriteTo(w io.Writer) (int64, error) {
	return 0, nil
}

// newHelperMethodDoctype creates and returns a helper method doctype.
func newHelperMethodDoctype(ln *line, rslt *result, parent element) *helperMethodDoctype {
	return &helperMethodDoctype{
		elementBase: newElementBase(ln, rslt, parent),
	}
}
