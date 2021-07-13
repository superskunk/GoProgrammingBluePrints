package trace

import (
	"fmt"
	"io"
)

// Tracer is the interface that describes an object capable of tracing events throught code.
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func New(w io.Writer) *tracer {
	return &tracer{
		out: w,
	}
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

type nilTracer struct {
}

func Off() *nilTracer {
	return &nilTracer{}
}

func (t *nilTracer) Trace(a ...interface{}) {

}
