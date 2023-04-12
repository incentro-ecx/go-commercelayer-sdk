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

// checks if the GETTransactionsTransactionId200ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETTransactionsTransactionId200ResponseDataRelationships{}

// GETTransactionsTransactionId200ResponseDataRelationships struct for GETTransactionsTransactionId200ResponseDataRelationships
type GETTransactionsTransactionId200ResponseDataRelationships struct {
	Order *POSTAdyenPayments201ResponseDataRelationshipsOrder `json:"order,omitempty"`
}

// NewGETTransactionsTransactionId200ResponseDataRelationships instantiates a new GETTransactionsTransactionId200ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETTransactionsTransactionId200ResponseDataRelationships() *GETTransactionsTransactionId200ResponseDataRelationships {
	this := GETTransactionsTransactionId200ResponseDataRelationships{}
	return &this
}

// NewGETTransactionsTransactionId200ResponseDataRelationshipsWithDefaults instantiates a new GETTransactionsTransactionId200ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETTransactionsTransactionId200ResponseDataRelationshipsWithDefaults() *GETTransactionsTransactionId200ResponseDataRelationships {
	this := GETTransactionsTransactionId200ResponseDataRelationships{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *GETTransactionsTransactionId200ResponseDataRelationships) GetOrder() POSTAdyenPayments201ResponseDataRelationshipsOrder {
	if o == nil || IsNil(o.Order) {
		var ret POSTAdyenPayments201ResponseDataRelationshipsOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTransactionsTransactionId200ResponseDataRelationships) GetOrderOk() (*POSTAdyenPayments201ResponseDataRelationshipsOrder, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *GETTransactionsTransactionId200ResponseDataRelationships) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given POSTAdyenPayments201ResponseDataRelationshipsOrder and assigns it to the Order field.
func (o *GETTransactionsTransactionId200ResponseDataRelationships) SetOrder(v POSTAdyenPayments201ResponseDataRelationshipsOrder) {
	o.Order = &v
}

func (o GETTransactionsTransactionId200ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETTransactionsTransactionId200ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	return toSerialize, nil
}

type NullableGETTransactionsTransactionId200ResponseDataRelationships struct {
	value *GETTransactionsTransactionId200ResponseDataRelationships
	isSet bool
}

func (v NullableGETTransactionsTransactionId200ResponseDataRelationships) Get() *GETTransactionsTransactionId200ResponseDataRelationships {
	return v.value
}

func (v *NullableGETTransactionsTransactionId200ResponseDataRelationships) Set(val *GETTransactionsTransactionId200ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETTransactionsTransactionId200ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETTransactionsTransactionId200ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETTransactionsTransactionId200ResponseDataRelationships(val *GETTransactionsTransactionId200ResponseDataRelationships) *NullableGETTransactionsTransactionId200ResponseDataRelationships {
	return &NullableGETTransactionsTransactionId200ResponseDataRelationships{value: val, isSet: true}
}

func (v NullableGETTransactionsTransactionId200ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETTransactionsTransactionId200ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}