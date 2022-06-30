/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.5
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// BraintreePaymentCreateData struct for BraintreePaymentCreateData
type BraintreePaymentCreateData struct {
	// The resource's type
	Type          string                               `json:"type"`
	Attributes    BraintreePaymentCreateDataAttributes `json:"attributes"`
	Relationships *AdyenPaymentCreateDataRelationships `json:"relationships,omitempty"`
}

// NewBraintreePaymentCreateData instantiates a new BraintreePaymentCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBraintreePaymentCreateData(type_ string, attributes BraintreePaymentCreateDataAttributes) *BraintreePaymentCreateData {
	this := BraintreePaymentCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewBraintreePaymentCreateDataWithDefaults instantiates a new BraintreePaymentCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBraintreePaymentCreateDataWithDefaults() *BraintreePaymentCreateData {
	this := BraintreePaymentCreateData{}
	var type_ string = "braintree_payments"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *BraintreePaymentCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BraintreePaymentCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BraintreePaymentCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *BraintreePaymentCreateData) GetAttributes() BraintreePaymentCreateDataAttributes {
	if o == nil {
		var ret BraintreePaymentCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *BraintreePaymentCreateData) GetAttributesOk() (*BraintreePaymentCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *BraintreePaymentCreateData) SetAttributes(v BraintreePaymentCreateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BraintreePaymentCreateData) GetRelationships() AdyenPaymentCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AdyenPaymentCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BraintreePaymentCreateData) GetRelationshipsOk() (*AdyenPaymentCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BraintreePaymentCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AdyenPaymentCreateDataRelationships and assigns it to the Relationships field.
func (o *BraintreePaymentCreateData) SetRelationships(v AdyenPaymentCreateDataRelationships) {
	o.Relationships = &v
}

func (o BraintreePaymentCreateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableBraintreePaymentCreateData struct {
	value *BraintreePaymentCreateData
	isSet bool
}

func (v NullableBraintreePaymentCreateData) Get() *BraintreePaymentCreateData {
	return v.value
}

func (v *NullableBraintreePaymentCreateData) Set(val *BraintreePaymentCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableBraintreePaymentCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableBraintreePaymentCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBraintreePaymentCreateData(val *BraintreePaymentCreateData) *NullableBraintreePaymentCreateData {
	return &NullableBraintreePaymentCreateData{value: val, isSet: true}
}

func (v NullableBraintreePaymentCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBraintreePaymentCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
