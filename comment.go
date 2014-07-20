package ace

import "io"

// comment represents a comment.
type comment struct {
	elementBase
}

// WriteTo writes data to w.
func (e *comment) WriteTo(w io.Writer) (int64, error) {
	return 0, nil
}

// newComment creates and returns a comment.
func newComment(ln *line, rslt *result, parent element) *comment {
	return &comment{
		elementBase: newElementBase(ln, rslt, parent),
	}
}
