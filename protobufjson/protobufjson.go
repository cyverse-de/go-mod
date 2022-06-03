package protobufjson

import (
	"errors"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func Unmarshal(data []byte, ptr interface{}) error {
	if _, ok := ptr.(*interface{}); ok {
		return nil
	}

	msg, ok := ptr.(proto.Message)
	if !ok {
		return errors.New("invalid protocol buffer message passed to Unmarshal")
	}

	return protojson.Unmarshal(data, msg)
}

func Marshal(ptr interface{}) ([]byte, error) {
	msg, ok := ptr.(proto.Message)
	if !ok {
		return nil, errors.New("invalid protocol buffer message passed to Marshal")
	}
	b, err := protojson.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Codec is a an implmentation of the NATS Encoder interface that
// can serialize/deserialize messages using protojson. Some logic is borrowed
// from the protocol buffer encoder included in NATS.
// See https://github.com/nats-io/nats.go/blob/main/encoders/protobuf/protobuf_enc.go
// See https://www.sohamkamani.com/golang/options-pattern/ for the options
// pattern.
type Codec struct {
	marshaller *protojson.MarshalOptions
	parser     *protojson.UnmarshalOptions
}

type CodecOption func(*Codec)

// WithAllowMarshalPartial corresponds to protojson.MarshalOption's AllowPartial
// field. See https://pkg.go.dev/google.golang.org/protobuf/encoding/protojson#MarshalOptions
func WithAllowMarshalPartial() CodecOption {
	return func(c *Codec) {
		c.marshaller.AllowPartial = true
	}
}

// WithAllowUnmarshalPartial() corresponds to protojson.UnmarshalOption's
// AllowPartial field. See https://pkg.go.dev/google.golang.org/protobuf/encoding/protojson#UnmarshalOptions
func WithAllowUnmarshalPartial() CodecOption {
	return func(c *Codec) {
		c.parser.AllowPartial = true
	}
}

// WithIndent corrsponds to protojson.MarshalOption's Indent field. See
// https://pkg.go.dev/google.golang.org/protobuf/encoding/protojson#MarshalOptions
func WithIndent(i string) CodecOption {
	return func(c *Codec) {
		c.marshaller.Indent = i
	}
}

// WithMultiline corresponds to protojson.MarshalOption's Multiline field. See
// https://pkg.go.dev/google.golang.org/protobuf/encoding/protojson#MarshalOptions
func WithMultiline() CodecOption {
	return func(c *Codec) {
		c.marshaller.Multiline = true
	}
}

// WithProtoNames corresponds to protojson.MarshalOption's UseProtoNames field.
// See https://pkg.go.dev/google.golang.org/protobuf/encoding/protojson#MarshalOptions
func WithProtoNames() CodecOption {
	return func(c *Codec) {
		c.marshaller.UseProtoNames = true
	}
}

// WithEnumNumbers corresponds to protojson.MarshalOption's UseEnumNumbers field.
// See https://pkg.go.dev/google.golang.org/protobuf/encoding/protojson#MarshalOptions
func WithEnumNumbers() CodecOption {
	return func(c *Codec) {
		c.marshaller.UseEnumNumbers = true
	}
}

// WithEmitUnpopulated corresponds to protojson.MarshalOption's EmitUnpopulated
// field. See https://pkg.go.dev/google.golang.org/protobuf/encoding/protojson#MarshalOptions
func WithEmitUnpopulated() CodecOption {
	return func(c *Codec) {
		c.marshaller.EmitUnpopulated = true
	}
}

// WithDiscardUnknown corresponds to protojson.UnmarshalOption's DiscardUnknown
// field. See https://pkg.go.dev/google.golang.org/protobuf/encoding/protojson#UnmarshalOptions
func WithDiscardUnknown() CodecOption {
	return func(c *Codec) {
		c.parser.DiscardUnknown = true
	}
}

// NewCodec returns a *Codec instance configured with the options defined by the
// With* functions defined in this package.
func NewCodec(opts ...CodecOption) *Codec {
	c := &Codec{
		marshaller: &protojson.MarshalOptions{},
		parser:     &protojson.UnmarshalOptions{},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *Codec) Encode(subject string, ptr interface{}) ([]byte, error) {
	msg, ok := ptr.(proto.Message)
	if !ok {
		return nil, errors.New("invalid protocol buffer message passed to Marshal")
	}
	return c.marshaller.Marshal(msg)
}

func (c *Codec) Decode(subject string, data []byte, vPtr interface{}) error {
	if _, ok := vPtr.(*interface{}); ok {
		return nil
	}

	msg, ok := vPtr.(proto.Message)
	if !ok {
		return errors.New("invalid protocol buffer message passed to Unmarshal")
	}

	return c.parser.Unmarshal(data, msg)
}
