package protobufjson

import (
	"errors"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// Codec is a an implmentation of the NATS Encoder interface that
// can serialize/deserialize messages using protojson. Some logic is borrowed
// from the protocol buffer encoder included in NATS.
// See https://github.com/nats-io/nats.go/blob/main/encoders/protobuf/protobuf_enc.go
type Codec struct{}

func (c *Codec) Encode(subject string, v interface{}) ([]byte, error) {
	msg, ok := v.(proto.Message)
	if !ok {
		return nil, errors.New("invalid protocol buffer message passed to Encode()")
	}
	b, err := protojson.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (c *Codec) Decode(subject string, data []byte, vPtr interface{}) error {
	if _, ok := vPtr.(*interface{}); ok {
		return nil
	}

	msg, ok := vPtr.(proto.Message)
	if !ok {
		return errors.New("invalid protocol buffer message passed to Decode()")
	}

	return protojson.Unmarshal(data, msg)
}
