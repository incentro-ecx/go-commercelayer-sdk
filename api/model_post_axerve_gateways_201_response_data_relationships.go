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

// checks if the POSTAxerveGateways201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTAxerveGateways201ResponseDataRelationships{}

// POSTAxerveGateways201ResponseDataRelationships struct for POSTAxerveGateways201ResponseDataRelationships
type POSTAxerveGateways201ResponseDataRelationships struct {
	PaymentMethods *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods  `json:"payment_methods,omitempty"`
	AxervePayments *POSTAxerveGateways201ResponseDataRelationshipsAxervePayments `json:"axerve_payments,omitempty"`
}

// NewPOSTAxerveGateways201ResponseDataRelationships instantiates a new POSTAxerveGateways201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAxerveGateways201ResponseDataRelationships() *POSTAxerveGateways201ResponseDataRelationships {
	this := POSTAxerveGateways201ResponseDataRelationships{}
	return &this
}

// NewPOSTAxerveGateways201ResponseDataRelationshipsWithDefaults instantiates a new POSTAxerveGateways201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAxerveGateways201ResponseDataRelationshipsWithDefaults() *POSTAxerveGateways201ResponseDataRelationships {
	this := POSTAxerveGateways201ResponseDataRelationships{}
	return &this
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise.
func (o *POSTAxerveGateways201ResponseDataRelationships) GetPaymentMethods() POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods {
	if o == nil || IsNil(o.PaymentMethods) {
		var ret POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods
		return ret
	}
	return *o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAxerveGateways201ResponseDataRelationships) GetPaymentMethodsOk() (*POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods, bool) {
	if o == nil || IsNil(o.PaymentMethods) {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *POSTAxerveGateways201ResponseDataRelationships) HasPaymentMethods() bool {
	if o != nil && !IsNil(o.PaymentMethods) {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods and assigns it to the PaymentMethods field.
func (o *POSTAxerveGateways201ResponseDataRelationships) SetPaymentMethods(v POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods) {
	o.PaymentMethods = &v
}

// GetAxervePayments returns the AxervePayments field value if set, zero value otherwise.
func (o *POSTAxerveGateways201ResponseDataRelationships) GetAxervePayments() POSTAxerveGateways201ResponseDataRelationshipsAxervePayments {
	if o == nil || IsNil(o.AxervePayments) {
		var ret POSTAxerveGateways201ResponseDataRelationshipsAxervePayments
		return ret
	}
	return *o.AxervePayments
}

// GetAxervePaymentsOk returns a tuple with the AxervePayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAxerveGateways201ResponseDataRelationships) GetAxervePaymentsOk() (*POSTAxerveGateways201ResponseDataRelationshipsAxervePayments, bool) {
	if o == nil || IsNil(o.AxervePayments) {
		return nil, false
	}
	return o.AxervePayments, true
}

// HasAxervePayments returns a boolean if a field has been set.
func (o *POSTAxerveGateways201ResponseDataRelationships) HasAxervePayments() bool {
	if o != nil && !IsNil(o.AxervePayments) {
		return true
	}

	return false
}

// SetAxervePayments gets a reference to the given POSTAxerveGateways201ResponseDataRelationshipsAxervePayments and assigns it to the AxervePayments field.
func (o *POSTAxerveGateways201ResponseDataRelationships) SetAxervePayments(v POSTAxerveGateways201ResponseDataRelationshipsAxervePayments) {
	o.AxervePayments = &v
}

func (o POSTAxerveGateways201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTAxerveGateways201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentMethods) {
		toSerialize["payment_methods"] = o.PaymentMethods
	}
	if !IsNil(o.AxervePayments) {
		toSerialize["axerve_payments"] = o.AxervePayments
	}
	return toSerialize, nil
}

type NullablePOSTAxerveGateways201ResponseDataRelationships struct {
	value *POSTAxerveGateways201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTAxerveGateways201ResponseDataRelationships) Get() *POSTAxerveGateways201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTAxerveGateways201ResponseDataRelationships) Set(val *POSTAxerveGateways201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAxerveGateways201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAxerveGateways201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAxerveGateways201ResponseDataRelationships(val *POSTAxerveGateways201ResponseDataRelationships) *NullablePOSTAxerveGateways201ResponseDataRelationships {
	return &NullablePOSTAxerveGateways201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTAxerveGateways201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAxerveGateways201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}