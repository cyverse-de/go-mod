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
type Codec struct{}

func (c *Codec) Encode(subject string, v interface{}) ([]byte, error) {
	return Marshal(v)
}

func (c *Codec) Decode(subject string, data []byte, vPtr interface{}) error {
	return Unmarshal(data, vPtr)
}
