package sampler

import (
	"fmt"
	"sync/atomic"

	"go.opentelemetry.io/otel/sdk/trace"
	otelTrace "go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/metadata"
)

const desc = `This sampler samples if TracdID.Valid(), gRPC metadata contains key "trace" or the child sampler decides to sample`

// Sampler decides whether a trace should be sampled and exported. This sampler will sample if
// paramters TraceID.Valid() == true or the Context contains gRPC metadata that has key "trace" (it doesn't care about values).
type Sampler struct {
	// child stores a *trace.Sampler. trace.Sampler is an interface. Because atomic.Value cares about
	// the underlying type, you can't just store trace.Sampler. So we do a pointer, which is the only valid
	// use of *interface I've ever seen.
	child atomic.Value // *trace.Sampler
}

// New creates a new Sampler with the child Sampler used if TraceID.Valid() == false and gRPC metadata does not contain
// key "trace".
func New(child trace.Sampler) (*Sampler, error) {
	if child == nil {
		return nil, fmt.Errorf("child cannot == nil")
	}

	s := &Sampler{}
	s.child.Store(&child)
	return s, nil
}

// ShouldSample implements trace.Sampler.ShouldSample.
func (s *Sampler) ShouldSample(p trace.SamplingParameters) trace.SamplingResult {
	psc := otelTrace.SpanContextFromContext(p.ParentContext)
	if psc.IsValid() {
		if psc.IsRemote() {
			if psc.IsSampled() {
				return trace.SamplingResult{
					Decision:   trace.RecordAndSample,
					Tracestate: psc.TraceState(),
				}
			}
		}
		if psc.IsSampled() {
			return trace.SamplingResult{
				Decision:   trace.RecordAndSample,
				Tracestate: psc.TraceState(),
			}
		}
	}
	md, ok := metadata.FromIncomingContext(p.ParentContext)
	if !ok {
		return (*s.child.Load().(*trace.Sampler)).ShouldSample(p)
	}

	if _, ok := md["trace"]; ok {
		psc := otelTrace.SpanContextFromContext(p.ParentContext)
		return trace.SamplingResult{
			Decision:   trace.RecordAndSample,
			Tracestate: psc.TraceState(),
		}
	}

	return (*s.child.Load().(*trace.Sampler)).ShouldSample(p)
}

// Description implements trace.Sampler.Description().
func (s *Sampler) Description() string {
	return desc
}

// Switch switches the underlying trace.Sampler.
func (s *Sampler) Switch(sampler trace.Sampler) {
	if sampler == nil {
		panic("cannot call Switch() with a nil Sampler")
	}
	s.child.Store(&sampler)
}