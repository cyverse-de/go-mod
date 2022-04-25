package deprotobuf

import (
	"github.com/cyverse-de/p/go/header"
)

type TextMapCarrier header.Header

func (d *TextMapCarrier) Get(key string) string {
	v, ok := d.Map[key]
	if !ok {
		return ""
	}
	if len(v.Value) == 0 {
		return ""
	}
	return v.Value[0]
}

func (d *TextMapCarrier) Set(key string, value string) {
	d.Map[key] = &header.Header_Value{
		Value: []string{
			value,
		},
	}
}

func (d *TextMapCarrier) Keys() []string {
	var keys []string
	for k := range d.Map {
		keys = append(keys, k)
	}
	return keys
}
