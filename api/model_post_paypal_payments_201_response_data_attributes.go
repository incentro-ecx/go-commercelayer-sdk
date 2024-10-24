/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTPaypalPayments201ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTPaypalPayments201ResponseDataAttributes{}

// POSTPaypalPayments201ResponseDataAttributes struct for POSTPaypalPayments201ResponseDataAttributes
type POSTPaypalPayments201ResponseDataAttributes struct {
	// The URL where the payer is redirected after they approve the payment.
	ReturnUrl interface{} `json:"return_url"`
	// The URL where the payer is redirected after they cancel the payment.
	CancelUrl interface{} `json:"cancel_url"`
	// A free-form field that you can use to send a note to the payer on PayPal.
	NoteToPayer interface{} `json:"note_to_payer,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPOSTPaypalPayments201ResponseDataAttributes instantiates a new POSTPaypalPayments201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTPaypalPayments201ResponseDataAttributes(returnUrl interface{}, cancelUrl interface{}) *POSTPaypalPayments201ResponseDataAttributes {
	this := POSTPaypalPayments201ResponseDataAttributes{}
	this.ReturnUrl = returnUrl
	this.CancelUrl = cancelUrl
	return &this
}

// NewPOSTPaypalPayments201ResponseDataAttributesWithDefaults instantiates a new POSTPaypalPayments201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTPaypalPayments201ResponseDataAttributesWithDefaults() *POSTPaypalPayments201ResponseDataAttributes {
	this := POSTPaypalPayments201ResponseDataAttributes{}
	return &this
}

// GetReturnUrl returns the ReturnUrl field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTPaypalPayments201ResponseDataAttributes) GetReturnUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.ReturnUrl
}

// GetReturnUrlOk returns a tuple with the ReturnUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPaypalPayments201ResponseDataAttributes) GetReturnUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReturnUrl) {
		return nil, false
	}
	return &o.ReturnUrl, true
}

// SetReturnUrl sets field value
func (o *POSTPaypalPayments201ResponseDataAttributes) SetReturnUrl(v interface{}) {
	o.ReturnUrl = v
}

// GetCancelUrl returns the CancelUrl field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTPaypalPayments201ResponseDataAttributes) GetCancelUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.CancelUrl
}

// GetCancelUrlOk returns a tuple with the CancelUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPaypalPayments201ResponseDataAttributes) GetCancelUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CancelUrl) {
		return nil, false
	}
	return &o.CancelUrl, true
}

// SetCancelUrl sets field value
func (o *POSTPaypalPayments201ResponseDataAttributes) SetCancelUrl(v interface{}) {
	o.CancelUrl = v
}

// GetNoteToPayer returns the NoteToPayer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPaypalPayments201ResponseDataAttributes) GetNoteToPayer() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.NoteToPayer
}

// GetNoteToPayerOk returns a tuple with the NoteToPayer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPaypalPayments201ResponseDataAttributes) GetNoteToPayerOk() (*interface{}, bool) {
	if o == nil || IsNil(o.NoteToPayer) {
		return nil, false
	}
	return &o.NoteToPayer, true
}

// HasNoteToPayer returns a boolean if a field has been set.
func (o *POSTPaypalPayments201ResponseDataAttributes) HasNoteToPayer() bool {
	if o != nil && IsNil(o.NoteToPayer) {
		return true
	}

	return false
}

// SetNoteToPayer gets a reference to the given interface{} and assigns it to the NoteToPayer field.
func (o *POSTPaypalPayments201ResponseDataAttributes) SetNoteToPayer(v interface{}) {
	o.NoteToPayer = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPaypalPayments201ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPaypalPayments201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTPaypalPayments201ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTPaypalPayments201ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPaypalPayments201ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPaypalPayments201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTPaypalPayments201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTPaypalPayments201ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPaypalPayments201ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPaypalPayments201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTPaypalPayments201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTPaypalPayments201ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o POSTPaypalPayments201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTPaypalPayments201ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ReturnUrl != nil {
		toSerialize["return_url"] = o.ReturnUrl
	}
	if o.CancelUrl != nil {
		toSerialize["cancel_url"] = o.CancelUrl
	}
	if o.NoteToPayer != nil {
		toSerialize["note_to_payer"] = o.NoteToPayer
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

type NullablePOSTPaypalPayments201ResponseDataAttributes struct {
	value *POSTPaypalPayments201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTPaypalPayments201ResponseDataAttributes) Get() *POSTPaypalPayments201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTPaypalPayments201ResponseDataAttributes) Set(val *POSTPaypalPayments201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTPaypalPayments201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTPaypalPayments201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTPaypalPayments201ResponseDataAttributes(val *POSTPaypalPayments201ResponseDataAttributes) *NullablePOSTPaypalPayments201ResponseDataAttributes {
	return &NullablePOSTPaypalPayments201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTPaypalPayments201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTPaypalPayments201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
