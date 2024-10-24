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

// checks if the AdyenPaymentUpdateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdyenPaymentUpdateDataRelationships{}

// AdyenPaymentUpdateDataRelationships struct for AdyenPaymentUpdateDataRelationships
type AdyenPaymentUpdateDataRelationships struct {
	Order *AdyenPaymentCreateDataRelationshipsOrder `json:"order,omitempty"`
}

// NewAdyenPaymentUpdateDataRelationships instantiates a new AdyenPaymentUpdateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdyenPaymentUpdateDataRelationships() *AdyenPaymentUpdateDataRelationships {
	this := AdyenPaymentUpdateDataRelationships{}
	return &this
}

// NewAdyenPaymentUpdateDataRelationshipsWithDefaults instantiates a new AdyenPaymentUpdateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdyenPaymentUpdateDataRelationshipsWithDefaults() *AdyenPaymentUpdateDataRelationships {
	this := AdyenPaymentUpdateDataRelationships{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *AdyenPaymentUpdateDataRelationships) GetOrder() AdyenPaymentCreateDataRelationshipsOrder {
	if o == nil || IsNil(o.Order) {
		var ret AdyenPaymentCreateDataRelationshipsOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenPaymentUpdateDataRelationships) GetOrderOk() (*AdyenPaymentCreateDataRelationshipsOrder, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *AdyenPaymentUpdateDataRelationships) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given AdyenPaymentCreateDataRelationshipsOrder and assigns it to the Order field.
func (o *AdyenPaymentUpdateDataRelationships) SetOrder(v AdyenPaymentCreateDataRelationshipsOrder) {
	o.Order = &v
}

func (o AdyenPaymentUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdyenPaymentUpdateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	return toSerialize, nil
}

type NullableAdyenPaymentUpdateDataRelationships struct {
	value *AdyenPaymentUpdateDataRelationships
	isSet bool
}

func (v NullableAdyenPaymentUpdateDataRelationships) Get() *AdyenPaymentUpdateDataRelationships {
	return v.value
}

func (v *NullableAdyenPaymentUpdateDataRelationships) Set(val *AdyenPaymentUpdateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAdyenPaymentUpdateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAdyenPaymentUpdateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdyenPaymentUpdateDataRelationships(val *AdyenPaymentUpdateDataRelationships) *NullableAdyenPaymentUpdateDataRelationships {
	return &NullableAdyenPaymentUpdateDataRelationships{value: val, isSet: true}
}

func (v NullableAdyenPaymentUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdyenPaymentUpdateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
