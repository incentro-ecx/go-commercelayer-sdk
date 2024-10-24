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

// checks if the PaymentOptionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentOptionData{}

// PaymentOptionData struct for PaymentOptionData
type PaymentOptionData struct {
	// The resource's type
	Type          interface{}                                               `json:"type"`
	Attributes    GETPaymentOptionsPaymentOptionId200ResponseDataAttributes `json:"attributes"`
	Relationships *PaymentOptionDataRelationships                           `json:"relationships,omitempty"`
}

// NewPaymentOptionData instantiates a new PaymentOptionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentOptionData(type_ interface{}, attributes GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) *PaymentOptionData {
	this := PaymentOptionData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPaymentOptionDataWithDefaults instantiates a new PaymentOptionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentOptionDataWithDefaults() *PaymentOptionData {
	this := PaymentOptionData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PaymentOptionData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentOptionData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PaymentOptionData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *PaymentOptionData) GetAttributes() GETPaymentOptionsPaymentOptionId200ResponseDataAttributes {
	if o == nil {
		var ret GETPaymentOptionsPaymentOptionId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PaymentOptionData) GetAttributesOk() (*GETPaymentOptionsPaymentOptionId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PaymentOptionData) SetAttributes(v GETPaymentOptionsPaymentOptionId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PaymentOptionData) GetRelationships() PaymentOptionDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret PaymentOptionDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentOptionData) GetRelationshipsOk() (*PaymentOptionDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PaymentOptionData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given PaymentOptionDataRelationships and assigns it to the Relationships field.
func (o *PaymentOptionData) SetRelationships(v PaymentOptionDataRelationships) {
	o.Relationships = &v
}

func (o PaymentOptionData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentOptionData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullablePaymentOptionData struct {
	value *PaymentOptionData
	isSet bool
}

func (v NullablePaymentOptionData) Get() *PaymentOptionData {
	return v.value
}

func (v *NullablePaymentOptionData) Set(val *PaymentOptionData) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentOptionData) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentOptionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentOptionData(val *PaymentOptionData) *NullablePaymentOptionData {
	return &NullablePaymentOptionData{value: val, isSet: true}
}

func (v NullablePaymentOptionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentOptionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
