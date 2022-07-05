package gitlab

import "io"

const (
	EventMr       = "merge_request"
	PRIVATE_TOKEN = "glpat-EJVGcz-RcEwemdYJ4btg"
)

type Engine struct {
}

func NewEngine() Engine {
	return Engine{}
}

func (e Engine) RecieveHook(reader io.Reader) {
	// TODO
}
