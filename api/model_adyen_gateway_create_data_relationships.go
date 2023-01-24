/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AdyenGatewayCreateDataRelationships struct for AdyenGatewayCreateDataRelationships
type AdyenGatewayCreateDataRelationships struct {
	AdyenPayments *AdyenGatewayCreateDataRelationshipsAdyenPayments `json:"adyen_payments,omitempty"`
}

// NewAdyenGatewayCreateDataRelationships instantiates a new AdyenGatewayCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdyenGatewayCreateDataRelationships() *AdyenGatewayCreateDataRelationships {
	this := AdyenGatewayCreateDataRelationships{}
	return &this
}

// NewAdyenGatewayCreateDataRelationshipsWithDefaults instantiates a new AdyenGatewayCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdyenGatewayCreateDataRelationshipsWithDefaults() *AdyenGatewayCreateDataRelationships {
	this := AdyenGatewayCreateDataRelationships{}
	return &this
}

// GetAdyenPayments returns the AdyenPayments field value if set, zero value otherwise.
func (o *AdyenGatewayCreateDataRelationships) GetAdyenPayments() AdyenGatewayCreateDataRelationshipsAdyenPayments {
	if o == nil || o.AdyenPayments == nil {
		var ret AdyenGatewayCreateDataRelationshipsAdyenPayments
		return ret
	}
	return *o.AdyenPayments
}

// GetAdyenPaymentsOk returns a tuple with the AdyenPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenGatewayCreateDataRelationships) GetAdyenPaymentsOk() (*AdyenGatewayCreateDataRelationshipsAdyenPayments, bool) {
	if o == nil || o.AdyenPayments == nil {
		return nil, false
	}
	return o.AdyenPayments, true
}

// HasAdyenPayments returns a boolean if a field has been set.
func (o *AdyenGatewayCreateDataRelationships) HasAdyenPayments() bool {
	if o != nil && o.AdyenPayments != nil {
		return true
	}

	return false
}

// SetAdyenPayments gets a reference to the given AdyenGatewayCreateDataRelationshipsAdyenPayments and assigns it to the AdyenPayments field.
func (o *AdyenGatewayCreateDataRelationships) SetAdyenPayments(v AdyenGatewayCreateDataRelationshipsAdyenPayments) {
	o.AdyenPayments = &v
}

func (o AdyenGatewayCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdyenPayments != nil {
		toSerialize["adyen_payments"] = o.AdyenPayments
	}
	return json.Marshal(toSerialize)
}

type NullableAdyenGatewayCreateDataRelationships struct {
	value *AdyenGatewayCreateDataRelationships
	isSet bool
}

func (v NullableAdyenGatewayCreateDataRelationships) Get() *AdyenGatewayCreateDataRelationships {
	return v.value
}

func (v *NullableAdyenGatewayCreateDataRelationships) Set(val *AdyenGatewayCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAdyenGatewayCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAdyenGatewayCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdyenGatewayCreateDataRelationships(val *AdyenGatewayCreateDataRelationships) *NullableAdyenGatewayCreateDataRelationships {
	return &NullableAdyenGatewayCreateDataRelationships{value: val, isSet: true}
}

func (v NullableAdyenGatewayCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdyenGatewayCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
