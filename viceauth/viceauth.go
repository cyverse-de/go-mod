// Package viceauth provides a signed-state codec for the VICE OAuth login
// flow. The OAuth "state" parameter must round-trip from vice-proxy through
// Keycloak and a vice-operator relay, so it carries both a CSRF token and the
// original app URL the browser should be returned to. An HMAC-SHA256 signature
// lets the operator trust the embedded URL without sharing vice-proxy's OIDC
// client secret.
package viceauth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

// StateClaims is the payload carried in the OAuth state parameter.
type StateClaims struct {
	// StateID is the CSRF token. vice-proxy stores it in its state-session
	// cookie and compares it against this value on the callback.
	StateID string `json:"sid"`

	// Origin is the full app URL (scheme, host, path, query) the browser
	// should be returned to once the operator relays the authorization code
	// back. Consumers MUST validate the host against an allowlist before
	// redirecting to it — a valid signature only proves the value was not
	// tampered with, not that the target is safe.
	Origin string `json:"origin"`
}

// Codec encodes and decodes signed OAuth state values. Construct one per
// process with NewCodec and reuse it; it is safe for concurrent use.
type Codec struct {
	secret []byte
}

// NewCodec returns a Codec that signs and verifies state values with the given
// HMAC secret. The same secret must be configured on every component that
// encodes or decodes state (vice-proxy and each vice-operator).
func NewCodec(secret []byte) *Codec {
	return &Codec{secret: secret}
}

// Errors returned by Decode. Callers should treat any error as "untrusted
// state" and refuse to act on it, rather than distinguishing between them.
var (
	ErrMalformedState = errors.New("viceauth: malformed state")
	ErrBadSignature   = errors.New("viceauth: state signature mismatch")
)

// Encode serializes claims to JSON and returns a signed, URL-safe state value
// of the form "<base64url(payload)>.<base64url(signature)>".
func (c *Codec) Encode(claims StateClaims) (string, error) {
	payload, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}
	b64 := base64.RawURLEncoding.EncodeToString(payload)
	return b64 + "." + c.sign(b64), nil
}

// Decode verifies the signature on a state value produced by Encode and
// returns the embedded claims. It returns ErrMalformedState for structurally
// invalid input and ErrBadSignature when the signature does not match; in
// either case the state must not be trusted.
func (c *Codec) Decode(state string) (StateClaims, error) {
	var claims StateClaims

	b64, sig, ok := strings.Cut(state, ".")
	if !ok {
		return claims, ErrMalformedState
	}

	// Constant-time comparison guards against signature-timing attacks.
	if !hmac.Equal([]byte(c.sign(b64)), []byte(sig)) {
		return claims, ErrBadSignature
	}

	payload, err := base64.RawURLEncoding.DecodeString(b64)
	if err != nil {
		return claims, ErrMalformedState
	}
	if err := json.Unmarshal(payload, &claims); err != nil {
		return claims, ErrMalformedState
	}
	return claims, nil
}

// sign returns the base64url-encoded HMAC-SHA256 of msg.
func (c *Codec) sign(msg string) string {
	mac := hmac.New(sha256.New, c.secret)
	mac.Write([]byte(msg))
	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}
