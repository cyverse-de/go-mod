package gotelnats

import (
	"context"
	"fmt"

	"github.com/cyverse-de/p/go/header"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
)

type PBTextMapCarrier struct {
	*header.Header
}

func (p *PBTextMapCarrier) Get(key string) string {
	v, ok := p.Map[key]
	if !ok {
		return ""
	}
	if len(v.Value) == 0 {
		return ""
	}
	return v.Value[0]
}

func (p *PBTextMapCarrier) Set(key string, value string) {
	p.Map[key] = &header.Header_Value{
		Value: []string{
			value,
		},
	}
}

func (p *PBTextMapCarrier) Keys() []string {
	var keys []string
	for k := range p.Map {
		keys = append(keys, k)
	}
	return keys
}

const natsName = "nats"
const tracerModule = "github.com/cyverse-de/go-mod/gotelnats"

type Operation int

const (
	Process Operation = iota
	Send
)

func (o Operation) String() string {
	switch o {
	case Process:
		return "process"
	case Send:
		return "send"
	}
	return "unknown"
}

// StartSpan creates a new context and populate it with information from
// the carrier. You will need to call defer span.End() in the calling code.
// Returns a new context based on context.Background() that can be passed to
// InjectSpan().
func StartSpan(carrier propagation.TextMapCarrier, subject string, op Operation) (context.Context, trace.Span) {
	ctx := otel.GetTextMapPropagator().Extract(context.Background(), carrier)
	tracer := otel.GetTracerProvider().Tracer(tracerModule)
	ctx, span := tracer.Start(ctx, fmt.Sprintf("%s %s", subject, op.String()), trace.WithSpanKind(trace.SpanKindConsumer))

	span.SetAttributes(
		semconv.MessagingSystemKey.String(natsName),
		semconv.MessagingProtocolKey.String(natsName),
		semconv.MessagingProtocolVersionKey.String(nats.Version),
		semconv.MessagingDestinationKindTopic,
		semconv.MessagingDestinationKey.String(subject),
	)

	if op != Send {
		span.SetAttributes(
			semconv.MessagingOperationKey.String(op.String()),
		)
	}

	return ctx, span
}

// InjectSpan adds information from the context to the carrier so that it can
// be serialized and sent along to other services. You will need to call
// 'defer span.End()' in the calling code. Returns a context based on the one
// passed into the function, plus a new span.
func InjectSpan(ctx context.Context, carrier propagation.TextMapCarrier, subject string, op Operation) (context.Context, trace.Span) {
	var tracer trace.Tracer

	if span := trace.SpanFromContext(ctx); span.SpanContext().IsValid() {
		tracer = span.TracerProvider().Tracer(tracerModule)
	} else {
		tracer = otel.GetTracerProvider().Tracer(tracerModule)
	}

	ctx, span := tracer.Start(ctx, fmt.Sprintf("%s %s", subject, op.String()), trace.WithSpanKind(trace.SpanKindProducer))

	span.SetAttributes(
		semconv.MessagingSystemKey.String(natsName),
		semconv.MessagingProtocolKey.String(natsName),
		semconv.MessagingProtocolVersionKey.String(nats.Version),
		semconv.MessagingDestinationKindTopic,
		semconv.MessagingDestinationKey.String(subject),
	)

	if op != Send {
		span.SetAttributes(
			semconv.MessagingOperationKey.String(op.String()),
		)
	}

	otel.GetTextMapPropagator().Inject(ctx, carrier)

	return ctx, span
}
