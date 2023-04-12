/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTSatispayPayments201ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTSatispayPayments201ResponseDataAttributes{}

// POSTSatispayPayments201ResponseDataAttributes struct for POSTSatispayPayments201ResponseDataAttributes
type POSTSatispayPayments201ResponseDataAttributes struct {
	// Activation code generated from the Satispay Dashboard.
	Token interface{} `json:"token,omitempty"`
	// The url to redirect the customer after the payment flow is completed.
	RedirectUrl interface{} `json:"redirect_url,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPOSTSatispayPayments201ResponseDataAttributes instantiates a new POSTSatispayPayments201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTSatispayPayments201ResponseDataAttributes() *POSTSatispayPayments201ResponseDataAttributes {
	this := POSTSatispayPayments201ResponseDataAttributes{}
	return &this
}

// NewPOSTSatispayPayments201ResponseDataAttributesWithDefaults instantiates a new POSTSatispayPayments201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTSatispayPayments201ResponseDataAttributesWithDefaults() *POSTSatispayPayments201ResponseDataAttributes {
	this := POSTSatispayPayments201ResponseDataAttributes{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSatispayPayments201ResponseDataAttributes) GetToken() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSatispayPayments201ResponseDataAttributes) GetTokenOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return &o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *POSTSatispayPayments201ResponseDataAttributes) HasToken() bool {
	if o != nil && IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given interface{} and assigns it to the Token field.
func (o *POSTSatispayPayments201ResponseDataAttributes) SetToken(v interface{}) {
	o.Token = v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSatispayPayments201ResponseDataAttributes) GetRedirectUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSatispayPayments201ResponseDataAttributes) GetRedirectUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RedirectUrl) {
		return nil, false
	}
	return &o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *POSTSatispayPayments201ResponseDataAttributes) HasRedirectUrl() bool {
	if o != nil && IsNil(o.RedirectUrl) {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given interface{} and assigns it to the RedirectUrl field.
func (o *POSTSatispayPayments201ResponseDataAttributes) SetRedirectUrl(v interface{}) {
	o.RedirectUrl = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSatispayPayments201ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSatispayPayments201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTSatispayPayments201ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTSatispayPayments201ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSatispayPayments201ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSatispayPayments201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTSatispayPayments201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTSatispayPayments201ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSatispayPayments201ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSatispayPayments201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTSatispayPayments201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTSatispayPayments201ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o POSTSatispayPayments201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTSatispayPayments201ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.RedirectUrl != nil {
		toSerialize["redirect_url"] = o.RedirectUrl
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
	return toSerialize, nil
}

type NullablePOSTSatispayPayments201ResponseDataAttributes struct {
	value *POSTSatispayPayments201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTSatispayPayments201ResponseDataAttributes) Get() *POSTSatispayPayments201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTSatispayPayments201ResponseDataAttributes) Set(val *POSTSatispayPayments201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTSatispayPayments201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTSatispayPayments201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTSatispayPayments201ResponseDataAttributes(val *POSTSatispayPayments201ResponseDataAttributes) *NullablePOSTSatispayPayments201ResponseDataAttributes {
	return &NullablePOSTSatispayPayments201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTSatispayPayments201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTSatispayPayments201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
