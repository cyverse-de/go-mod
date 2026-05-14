package viceauth

import (
	"errors"
	"strings"
	"testing"
)

func TestCodecRoundTrip(t *testing.T) {
	c := NewCodec([]byte("test-secret"))

	cases := []StateClaims{
		{StateID: "81d5cc3f-4fa8-11f1-a186-1a46dc0baeb8", Origin: "https://a44abe2d0.cyverse.run:4343/foo?bar=baz"},
		{StateID: "sid-only", Origin: ""},
		{StateID: "", Origin: ""},
	}
	for _, want := range cases {
		state, err := c.Encode(want)
		if err != nil {
			t.Fatalf("Encode(%+v): %v", want, err)
		}
		got, err := c.Decode(state)
		if err != nil {
			t.Fatalf("Decode(%q): %v", state, err)
		}
		if got != want {
			t.Errorf("round trip mismatch: got %+v, want %+v", got, want)
		}
	}
}

func TestDecodeRejectsUntrustedState(t *testing.T) {
	c := NewCodec([]byte("test-secret"))

	valid, err := c.Encode(StateClaims{StateID: "sid", Origin: "https://x.cyverse.run:4343/"})
	if err != nil {
		t.Fatalf("Encode: %v", err)
	}
	payload, sig, _ := strings.Cut(valid, ".")

	otherSecret, err := NewCodec([]byte("other-secret")).Encode(StateClaims{StateID: "sid"})
	if err != nil {
		t.Fatalf("Encode with other secret: %v", err)
	}

	// A validly-signed value whose payload is not base64 — exercises the
	// decode path past the signature check.
	badPayload := "!!!not-base64!!!"
	signedBadPayload := badPayload + "." + c.sign(badPayload)

	tests := []struct {
		name    string
		state   string
		wantErr error
	}{
		{"empty", "", ErrMalformedState},
		{"no signature delimiter", payload, ErrMalformedState},
		{"tampered payload", payload + "x." + sig, ErrBadSignature},
		{"tampered signature", payload + "." + sig + "x", ErrBadSignature},
		{"signed with a different secret", otherSecret, ErrBadSignature},
		{"valid signature, non-base64 payload", signedBadPayload, ErrMalformedState},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.Decode(tt.state)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Decode(%q) = %v, want %v", tt.state, err, tt.wantErr)
			}
		})
	}
}
