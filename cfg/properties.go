package cfg

import (
	"bytes"
	"fmt"

	"github.com/magiconair/properties"
)

type Properties struct{}

func PropertiesParser() *Properties {
	return &Properties{}
}

func (p *Properties) Unmarshal(b []byte) (map[string]interface{}, error) {
	props, err := properties.Load(b, properties.UTF8)
	if err != nil {
		return nil, err
	}

	m := props.Map()

	out := make(map[string]interface{})

	for key, value := range m {
		out[key] = value
	}

	return out, nil
}

func (p *Properties) Marshal(m map[string]interface{}) ([]byte, error) {
	var err error

	props := properties.NewProperties()

	for key, value := range m {
		// turn the value into a string.
		v := fmt.Sprint(value) // Got this from the dotenv parser in koanf.

		// load the key and value into the properties instance. ignoring all
		// of the return values should be okay since props is empty.
		if _, _, err = props.Set(key, v); err != nil {
			return nil, err
		}
	}

	// Write out the properties to a []byte.
	buf := bytes.NewBuffer(make([]byte, 0))
	if _, err = props.Write(buf, properties.UTF8); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
