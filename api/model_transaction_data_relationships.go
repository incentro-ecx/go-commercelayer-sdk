/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.4.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// TransactionDataRelationships struct for TransactionDataRelationships
type TransactionDataRelationships struct {
	Order *AdyenPaymentDataRelationshipsOrder `json:"order,omitempty"`
}

// NewTransactionDataRelationships instantiates a new TransactionDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionDataRelationships() *TransactionDataRelationships {
	this := TransactionDataRelationships{}
	return &this
}

// NewTransactionDataRelationshipsWithDefaults instantiates a new TransactionDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionDataRelationshipsWithDefaults() *TransactionDataRelationships {
	this := TransactionDataRelationships{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *TransactionDataRelationships) GetOrder() AdyenPaymentDataRelationshipsOrder {
	if o == nil || o.Order == nil {
		var ret AdyenPaymentDataRelationshipsOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDataRelationships) GetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *TransactionDataRelationships) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given AdyenPaymentDataRelationshipsOrder and assigns it to the Order field.
func (o *TransactionDataRelationships) SetOrder(v AdyenPaymentDataRelationshipsOrder) {
	o.Order = &v
}

func (o TransactionDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionDataRelationships struct {
	value *TransactionDataRelationships
	isSet bool
}

func (v NullableTransactionDataRelationships) Get() *TransactionDataRelationships {
	return v.value
}

func (v *NullableTransactionDataRelationships) Set(val *TransactionDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionDataRelationships(val *TransactionDataRelationships) *NullableTransactionDataRelationships {
	return &NullableTransactionDataRelationships{value: val, isSet: true}
}

func (v NullableTransactionDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}