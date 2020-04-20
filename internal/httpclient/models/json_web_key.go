// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// JSONWebKey JSONWebKey JSONWebKey JSONWebKey It is important that this model object is named JSONWebKey for
// "swagger generate spec" to generate only on definition of a
// JSONWebKey.
//
// swagger:model JSONWebKey
type JSONWebKey struct {

	// The "alg" (algorithm) parameter identifies the algorithm intended for
	// use with the key.  The values used should either be registered in the
	// IANA "JSON Web Signature and Encryption Algorithms" registry
	// established by [JWA] or be a value that contains a Collision-
	// Resistant Name.
	// Required: true
	Alg *string `json:"alg"`

	// crv
	Crv string `json:"crv,omitempty"`

	// d
	D string `json:"d,omitempty"`

	// dp
	Dp string `json:"dp,omitempty"`

	// dq
	Dq string `json:"dq,omitempty"`

	// e
	E string `json:"e,omitempty"`

	// k
	K string `json:"k,omitempty"`

	// The "kid" (key ID) parameter is used to match a specific key.  This
	// is used, for instance, to choose among a set of keys within a JWK Set
	// during key rollover.  The structure of the "kid" value is
	// unspecified.  When "kid" values are used within a JWK Set, different
	// keys within the JWK Set SHOULD use distinct "kid" values.  (One
	// example in which different keys might use the same "kid" value is if
	// they have different "kty" (key type) values but are considered to be
	// equivalent alternatives by the application using them.)  The "kid"
	// value is a case-sensitive string.
	// Required: true
	Kid *string `json:"kid"`

	// The "kty" (key type) parameter identifies the cryptographic algorithm
	// family used with the key, such as "RSA" or "EC". "kty" values should
	// either be registered in the IANA "JSON Web Key Types" registry
	// established by [JWA] or be a value that contains a Collision-
	// Resistant Name.  The "kty" value is a case-sensitive string.
	// Required: true
	Kty *string `json:"kty"`

	// n
	N string `json:"n,omitempty"`

	// p
	P string `json:"p,omitempty"`

	// q
	Q string `json:"q,omitempty"`

	// qi
	Qi string `json:"qi,omitempty"`

	// Use ("public key use") identifies the intended use of
	// the public key. The "use" parameter is employed to indicate whether
	// a public key is used for encrypting data or verifying the signature
	// on data. Values are commonly "sig" (signature) or "enc" (encryption).
	// Required: true
	Use *string `json:"use"`

	// x
	X string `json:"x,omitempty"`

	// The "x5c" (X.509 certificate chain) parameter contains a chain of one
	// or more PKIX certificates [RFC5280].  The certificate chain is
	// represented as a JSON array of certificate value strings.  Each
	// string in the array is a base64-encoded (Section 4 of [RFC4648] --
	// not base64url-encoded) DER [ITU.X690.1994] PKIX certificate value.
	// The PKIX certificate containing the key value MUST be the first
	// certificate.
	X5c []string `json:"x5c"`

	// y
	Y string `json:"y,omitempty"`
}

// Validate validates this JSON web key
func (m *JSONWebKey) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JSONWebKey) validateAlg(formats strfmt.Registry) error {

	if err := validate.Required("alg", "body", m.Alg); err != nil {
		return err
	}

	return nil
}

func (m *JSONWebKey) validateKid(formats strfmt.Registry) error {

	if err := validate.Required("kid", "body", m.Kid); err != nil {
		return err
	}

	return nil
}

func (m *JSONWebKey) validateKty(formats strfmt.Registry) error {

	if err := validate.Required("kty", "body", m.Kty); err != nil {
		return err
	}

	return nil
}

func (m *JSONWebKey) validateUse(formats strfmt.Registry) error {

	if err := validate.Required("use", "body", m.Use); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JSONWebKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JSONWebKey) UnmarshalBinary(b []byte) error {
	var res JSONWebKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
