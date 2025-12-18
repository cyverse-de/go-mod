// Deprecated: Provided for compatibility with code that used to use protocol buffers.

package header

type Header struct {
	Map map[string]*Header_Value `json:"map,omitempty"`
}

func (x *Header) GetMap() map[string]*Header_Value {
	if x != nil {
		return x.Map
	}
	return nil
}

type Header_Value struct {
	Value []string `json:"value,omitempty"`
}

func (x *Header_Value) GetValue() []string {
	if x != nil {
		return x.Value
	}
	return nil
}
