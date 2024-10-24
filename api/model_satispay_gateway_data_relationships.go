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

// checks if the SatispayGatewayDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SatispayGatewayDataRelationships{}

// SatispayGatewayDataRelationships struct for SatispayGatewayDataRelationships
type SatispayGatewayDataRelationships struct {
	PaymentMethods   *AdyenGatewayDataRelationshipsPaymentMethods      `json:"payment_methods,omitempty"`
	Versions         *AddressDataRelationshipsVersions                 `json:"versions,omitempty"`
	SatispayPayments *SatispayGatewayDataRelationshipsSatispayPayments `json:"satispay_payments,omitempty"`
}

// NewSatispayGatewayDataRelationships instantiates a new SatispayGatewayDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSatispayGatewayDataRelationships() *SatispayGatewayDataRelationships {
	this := SatispayGatewayDataRelationships{}
	return &this
}

// NewSatispayGatewayDataRelationshipsWithDefaults instantiates a new SatispayGatewayDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSatispayGatewayDataRelationshipsWithDefaults() *SatispayGatewayDataRelationships {
	this := SatispayGatewayDataRelationships{}
	return &this
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise.
func (o *SatispayGatewayDataRelationships) GetPaymentMethods() AdyenGatewayDataRelationshipsPaymentMethods {
	if o == nil || IsNil(o.PaymentMethods) {
		var ret AdyenGatewayDataRelationshipsPaymentMethods
		return ret
	}
	return *o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SatispayGatewayDataRelationships) GetPaymentMethodsOk() (*AdyenGatewayDataRelationshipsPaymentMethods, bool) {
	if o == nil || IsNil(o.PaymentMethods) {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *SatispayGatewayDataRelationships) HasPaymentMethods() bool {
	if o != nil && !IsNil(o.PaymentMethods) {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given AdyenGatewayDataRelationshipsPaymentMethods and assigns it to the PaymentMethods field.
func (o *SatispayGatewayDataRelationships) SetPaymentMethods(v AdyenGatewayDataRelationshipsPaymentMethods) {
	o.PaymentMethods = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *SatispayGatewayDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SatispayGatewayDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *SatispayGatewayDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *SatispayGatewayDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

// GetSatispayPayments returns the SatispayPayments field value if set, zero value otherwise.
func (o *SatispayGatewayDataRelationships) GetSatispayPayments() SatispayGatewayDataRelationshipsSatispayPayments {
	if o == nil || IsNil(o.SatispayPayments) {
		var ret SatispayGatewayDataRelationshipsSatispayPayments
		return ret
	}
	return *o.SatispayPayments
}

// GetSatispayPaymentsOk returns a tuple with the SatispayPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SatispayGatewayDataRelationships) GetSatispayPaymentsOk() (*SatispayGatewayDataRelationshipsSatispayPayments, bool) {
	if o == nil || IsNil(o.SatispayPayments) {
		return nil, false
	}
	return o.SatispayPayments, true
}

// HasSatispayPayments returns a boolean if a field has been set.
func (o *SatispayGatewayDataRelationships) HasSatispayPayments() bool {
	if o != nil && !IsNil(o.SatispayPayments) {
		return true
	}

	return false
}

// SetSatispayPayments gets a reference to the given SatispayGatewayDataRelationshipsSatispayPayments and assigns it to the SatispayPayments field.
func (o *SatispayGatewayDataRelationships) SetSatispayPayments(v SatispayGatewayDataRelationshipsSatispayPayments) {
	o.SatispayPayments = &v
}

func (o SatispayGatewayDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SatispayGatewayDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentMethods) {
		toSerialize["payment_methods"] = o.PaymentMethods
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	if !IsNil(o.SatispayPayments) {
		toSerialize["satispay_payments"] = o.SatispayPayments
	}
	return toSerialize, nil
}

type NullableSatispayGatewayDataRelationships struct {
	value *SatispayGatewayDataRelationships
	isSet bool
}

func (v NullableSatispayGatewayDataRelationships) Get() *SatispayGatewayDataRelationships {
	return v.value
}

func (v *NullableSatispayGatewayDataRelationships) Set(val *SatispayGatewayDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableSatispayGatewayDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableSatispayGatewayDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSatispayGatewayDataRelationships(val *SatispayGatewayDataRelationships) *NullableSatispayGatewayDataRelationships {
	return &NullableSatispayGatewayDataRelationships{value: val, isSet: true}
}

func (v NullableSatispayGatewayDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSatispayGatewayDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
