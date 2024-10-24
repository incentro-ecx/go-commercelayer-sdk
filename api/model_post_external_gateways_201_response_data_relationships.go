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

// checks if the POSTExternalGateways201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTExternalGateways201ResponseDataRelationships{}

// POSTExternalGateways201ResponseDataRelationships struct for POSTExternalGateways201ResponseDataRelationships
type POSTExternalGateways201ResponseDataRelationships struct {
	PaymentMethods   *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods      `json:"payment_methods,omitempty"`
	Versions         *POSTAddresses201ResponseDataRelationshipsVersions                `json:"versions,omitempty"`
	ExternalPayments *POSTExternalGateways201ResponseDataRelationshipsExternalPayments `json:"external_payments,omitempty"`
}

// NewPOSTExternalGateways201ResponseDataRelationships instantiates a new POSTExternalGateways201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTExternalGateways201ResponseDataRelationships() *POSTExternalGateways201ResponseDataRelationships {
	this := POSTExternalGateways201ResponseDataRelationships{}
	return &this
}

// NewPOSTExternalGateways201ResponseDataRelationshipsWithDefaults instantiates a new POSTExternalGateways201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTExternalGateways201ResponseDataRelationshipsWithDefaults() *POSTExternalGateways201ResponseDataRelationships {
	this := POSTExternalGateways201ResponseDataRelationships{}
	return &this
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise.
func (o *POSTExternalGateways201ResponseDataRelationships) GetPaymentMethods() POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods {
	if o == nil || IsNil(o.PaymentMethods) {
		var ret POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods
		return ret
	}
	return *o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTExternalGateways201ResponseDataRelationships) GetPaymentMethodsOk() (*POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods, bool) {
	if o == nil || IsNil(o.PaymentMethods) {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *POSTExternalGateways201ResponseDataRelationships) HasPaymentMethods() bool {
	if o != nil && !IsNil(o.PaymentMethods) {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods and assigns it to the PaymentMethods field.
func (o *POSTExternalGateways201ResponseDataRelationships) SetPaymentMethods(v POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods) {
	o.PaymentMethods = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *POSTExternalGateways201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret POSTAddresses201ResponseDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTExternalGateways201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *POSTExternalGateways201ResponseDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given POSTAddresses201ResponseDataRelationshipsVersions and assigns it to the Versions field.
func (o *POSTExternalGateways201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions) {
	o.Versions = &v
}

// GetExternalPayments returns the ExternalPayments field value if set, zero value otherwise.
func (o *POSTExternalGateways201ResponseDataRelationships) GetExternalPayments() POSTExternalGateways201ResponseDataRelationshipsExternalPayments {
	if o == nil || IsNil(o.ExternalPayments) {
		var ret POSTExternalGateways201ResponseDataRelationshipsExternalPayments
		return ret
	}
	return *o.ExternalPayments
}

// GetExternalPaymentsOk returns a tuple with the ExternalPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTExternalGateways201ResponseDataRelationships) GetExternalPaymentsOk() (*POSTExternalGateways201ResponseDataRelationshipsExternalPayments, bool) {
	if o == nil || IsNil(o.ExternalPayments) {
		return nil, false
	}
	return o.ExternalPayments, true
}

// HasExternalPayments returns a boolean if a field has been set.
func (o *POSTExternalGateways201ResponseDataRelationships) HasExternalPayments() bool {
	if o != nil && !IsNil(o.ExternalPayments) {
		return true
	}

	return false
}

// SetExternalPayments gets a reference to the given POSTExternalGateways201ResponseDataRelationshipsExternalPayments and assigns it to the ExternalPayments field.
func (o *POSTExternalGateways201ResponseDataRelationships) SetExternalPayments(v POSTExternalGateways201ResponseDataRelationshipsExternalPayments) {
	o.ExternalPayments = &v
}

func (o POSTExternalGateways201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTExternalGateways201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentMethods) {
		toSerialize["payment_methods"] = o.PaymentMethods
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	if !IsNil(o.ExternalPayments) {
		toSerialize["external_payments"] = o.ExternalPayments
	}
	return toSerialize, nil
}

type NullablePOSTExternalGateways201ResponseDataRelationships struct {
	value *POSTExternalGateways201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTExternalGateways201ResponseDataRelationships) Get() *POSTExternalGateways201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTExternalGateways201ResponseDataRelationships) Set(val *POSTExternalGateways201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTExternalGateways201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTExternalGateways201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTExternalGateways201ResponseDataRelationships(val *POSTExternalGateways201ResponseDataRelationships) *NullablePOSTExternalGateways201ResponseDataRelationships {
	return &NullablePOSTExternalGateways201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTExternalGateways201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTExternalGateways201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
