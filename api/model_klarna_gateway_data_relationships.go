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

// checks if the KlarnaGatewayDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KlarnaGatewayDataRelationships{}

// KlarnaGatewayDataRelationships struct for KlarnaGatewayDataRelationships
type KlarnaGatewayDataRelationships struct {
	PaymentMethods *AdyenGatewayDataRelationshipsPaymentMethods  `json:"payment_methods,omitempty"`
	Versions       *AddressDataRelationshipsVersions             `json:"versions,omitempty"`
	KlarnaPayments *KlarnaGatewayDataRelationshipsKlarnaPayments `json:"klarna_payments,omitempty"`
}

// NewKlarnaGatewayDataRelationships instantiates a new KlarnaGatewayDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKlarnaGatewayDataRelationships() *KlarnaGatewayDataRelationships {
	this := KlarnaGatewayDataRelationships{}
	return &this
}

// NewKlarnaGatewayDataRelationshipsWithDefaults instantiates a new KlarnaGatewayDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKlarnaGatewayDataRelationshipsWithDefaults() *KlarnaGatewayDataRelationships {
	this := KlarnaGatewayDataRelationships{}
	return &this
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise.
func (o *KlarnaGatewayDataRelationships) GetPaymentMethods() AdyenGatewayDataRelationshipsPaymentMethods {
	if o == nil || IsNil(o.PaymentMethods) {
		var ret AdyenGatewayDataRelationshipsPaymentMethods
		return ret
	}
	return *o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlarnaGatewayDataRelationships) GetPaymentMethodsOk() (*AdyenGatewayDataRelationshipsPaymentMethods, bool) {
	if o == nil || IsNil(o.PaymentMethods) {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *KlarnaGatewayDataRelationships) HasPaymentMethods() bool {
	if o != nil && !IsNil(o.PaymentMethods) {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given AdyenGatewayDataRelationshipsPaymentMethods and assigns it to the PaymentMethods field.
func (o *KlarnaGatewayDataRelationships) SetPaymentMethods(v AdyenGatewayDataRelationshipsPaymentMethods) {
	o.PaymentMethods = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *KlarnaGatewayDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlarnaGatewayDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *KlarnaGatewayDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *KlarnaGatewayDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

// GetKlarnaPayments returns the KlarnaPayments field value if set, zero value otherwise.
func (o *KlarnaGatewayDataRelationships) GetKlarnaPayments() KlarnaGatewayDataRelationshipsKlarnaPayments {
	if o == nil || IsNil(o.KlarnaPayments) {
		var ret KlarnaGatewayDataRelationshipsKlarnaPayments
		return ret
	}
	return *o.KlarnaPayments
}

// GetKlarnaPaymentsOk returns a tuple with the KlarnaPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlarnaGatewayDataRelationships) GetKlarnaPaymentsOk() (*KlarnaGatewayDataRelationshipsKlarnaPayments, bool) {
	if o == nil || IsNil(o.KlarnaPayments) {
		return nil, false
	}
	return o.KlarnaPayments, true
}

// HasKlarnaPayments returns a boolean if a field has been set.
func (o *KlarnaGatewayDataRelationships) HasKlarnaPayments() bool {
	if o != nil && !IsNil(o.KlarnaPayments) {
		return true
	}

	return false
}

// SetKlarnaPayments gets a reference to the given KlarnaGatewayDataRelationshipsKlarnaPayments and assigns it to the KlarnaPayments field.
func (o *KlarnaGatewayDataRelationships) SetKlarnaPayments(v KlarnaGatewayDataRelationshipsKlarnaPayments) {
	o.KlarnaPayments = &v
}

func (o KlarnaGatewayDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KlarnaGatewayDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentMethods) {
		toSerialize["payment_methods"] = o.PaymentMethods
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	if !IsNil(o.KlarnaPayments) {
		toSerialize["klarna_payments"] = o.KlarnaPayments
	}
	return toSerialize, nil
}

type NullableKlarnaGatewayDataRelationships struct {
	value *KlarnaGatewayDataRelationships
	isSet bool
}

func (v NullableKlarnaGatewayDataRelationships) Get() *KlarnaGatewayDataRelationships {
	return v.value
}

func (v *NullableKlarnaGatewayDataRelationships) Set(val *KlarnaGatewayDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableKlarnaGatewayDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableKlarnaGatewayDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlarnaGatewayDataRelationships(val *KlarnaGatewayDataRelationships) *NullableKlarnaGatewayDataRelationships {
	return &NullableKlarnaGatewayDataRelationships{value: val, isSet: true}
}

func (v NullableKlarnaGatewayDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlarnaGatewayDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
