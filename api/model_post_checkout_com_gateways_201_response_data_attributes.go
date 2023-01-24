/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// POSTCheckoutComGateways201ResponseDataAttributes struct for POSTCheckoutComGateways201ResponseDataAttributes
type POSTCheckoutComGateways201ResponseDataAttributes struct {
	// The payment gateway's internal name.
	Name string `json:"name"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// The gateway secret key.
	SecretKey string `json:"secret_key"`
	// The gateway public key.
	PublicKey string `json:"public_key"`
}

// NewPOSTCheckoutComGateways201ResponseDataAttributes instantiates a new POSTCheckoutComGateways201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCheckoutComGateways201ResponseDataAttributes(name string, secretKey string, publicKey string) *POSTCheckoutComGateways201ResponseDataAttributes {
	this := POSTCheckoutComGateways201ResponseDataAttributes{}
	this.Name = name
	this.SecretKey = secretKey
	this.PublicKey = publicKey
	return &this
}

// NewPOSTCheckoutComGateways201ResponseDataAttributesWithDefaults instantiates a new POSTCheckoutComGateways201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCheckoutComGateways201ResponseDataAttributesWithDefaults() *POSTCheckoutComGateways201ResponseDataAttributes {
	this := POSTCheckoutComGateways201ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *POSTCheckoutComGateways201ResponseDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *POSTCheckoutComGateways201ResponseDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *POSTCheckoutComGateways201ResponseDataAttributes) SetName(v string) {
	o.Name = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *POSTCheckoutComGateways201ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCheckoutComGateways201ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTCheckoutComGateways201ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *POSTCheckoutComGateways201ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *POSTCheckoutComGateways201ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCheckoutComGateways201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTCheckoutComGateways201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *POSTCheckoutComGateways201ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *POSTCheckoutComGateways201ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCheckoutComGateways201ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTCheckoutComGateways201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *POSTCheckoutComGateways201ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetSecretKey returns the SecretKey field value
func (o *POSTCheckoutComGateways201ResponseDataAttributes) GetSecretKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value
// and a boolean to check if the value has been set.
func (o *POSTCheckoutComGateways201ResponseDataAttributes) GetSecretKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretKey, true
}

// SetSecretKey sets field value
func (o *POSTCheckoutComGateways201ResponseDataAttributes) SetSecretKey(v string) {
	o.SecretKey = v
}

// GetPublicKey returns the PublicKey field value
func (o *POSTCheckoutComGateways201ResponseDataAttributes) GetPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *POSTCheckoutComGateways201ResponseDataAttributes) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *POSTCheckoutComGateways201ResponseDataAttributes) SetPublicKey(v string) {
	o.PublicKey = v
}

func (o POSTCheckoutComGateways201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.ReferenceOrigin != nil {
		toSerialize["reference_origin"] = o.ReferenceOrigin
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["secret_key"] = o.SecretKey
	}
	if true {
		toSerialize["public_key"] = o.PublicKey
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTCheckoutComGateways201ResponseDataAttributes struct {
	value *POSTCheckoutComGateways201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTCheckoutComGateways201ResponseDataAttributes) Get() *POSTCheckoutComGateways201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTCheckoutComGateways201ResponseDataAttributes) Set(val *POSTCheckoutComGateways201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCheckoutComGateways201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCheckoutComGateways201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCheckoutComGateways201ResponseDataAttributes(val *POSTCheckoutComGateways201ResponseDataAttributes) *NullablePOSTCheckoutComGateways201ResponseDataAttributes {
	return &NullablePOSTCheckoutComGateways201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTCheckoutComGateways201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCheckoutComGateways201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
