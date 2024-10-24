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

// checks if the AdyenGatewayCreateDataRelationshipsAdyenPayments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdyenGatewayCreateDataRelationshipsAdyenPayments{}

// AdyenGatewayCreateDataRelationshipsAdyenPayments struct for AdyenGatewayCreateDataRelationshipsAdyenPayments
type AdyenGatewayCreateDataRelationshipsAdyenPayments struct {
	Data AdyenGatewayDataRelationshipsAdyenPaymentsData `json:"data"`
}

// NewAdyenGatewayCreateDataRelationshipsAdyenPayments instantiates a new AdyenGatewayCreateDataRelationshipsAdyenPayments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdyenGatewayCreateDataRelationshipsAdyenPayments(data AdyenGatewayDataRelationshipsAdyenPaymentsData) *AdyenGatewayCreateDataRelationshipsAdyenPayments {
	this := AdyenGatewayCreateDataRelationshipsAdyenPayments{}
	this.Data = data
	return &this
}

// NewAdyenGatewayCreateDataRelationshipsAdyenPaymentsWithDefaults instantiates a new AdyenGatewayCreateDataRelationshipsAdyenPayments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdyenGatewayCreateDataRelationshipsAdyenPaymentsWithDefaults() *AdyenGatewayCreateDataRelationshipsAdyenPayments {
	this := AdyenGatewayCreateDataRelationshipsAdyenPayments{}
	return &this
}

// GetData returns the Data field value
func (o *AdyenGatewayCreateDataRelationshipsAdyenPayments) GetData() AdyenGatewayDataRelationshipsAdyenPaymentsData {
	if o == nil {
		var ret AdyenGatewayDataRelationshipsAdyenPaymentsData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AdyenGatewayCreateDataRelationshipsAdyenPayments) GetDataOk() (*AdyenGatewayDataRelationshipsAdyenPaymentsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AdyenGatewayCreateDataRelationshipsAdyenPayments) SetData(v AdyenGatewayDataRelationshipsAdyenPaymentsData) {
	o.Data = v
}

func (o AdyenGatewayCreateDataRelationshipsAdyenPayments) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdyenGatewayCreateDataRelationshipsAdyenPayments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAdyenGatewayCreateDataRelationshipsAdyenPayments struct {
	value *AdyenGatewayCreateDataRelationshipsAdyenPayments
	isSet bool
}

func (v NullableAdyenGatewayCreateDataRelationshipsAdyenPayments) Get() *AdyenGatewayCreateDataRelationshipsAdyenPayments {
	return v.value
}

func (v *NullableAdyenGatewayCreateDataRelationshipsAdyenPayments) Set(val *AdyenGatewayCreateDataRelationshipsAdyenPayments) {
	v.value = val
	v.isSet = true
}

func (v NullableAdyenGatewayCreateDataRelationshipsAdyenPayments) IsSet() bool {
	return v.isSet
}

func (v *NullableAdyenGatewayCreateDataRelationshipsAdyenPayments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdyenGatewayCreateDataRelationshipsAdyenPayments(val *AdyenGatewayCreateDataRelationshipsAdyenPayments) *NullableAdyenGatewayCreateDataRelationshipsAdyenPayments {
	return &NullableAdyenGatewayCreateDataRelationshipsAdyenPayments{value: val, isSet: true}
}

func (v NullableAdyenGatewayCreateDataRelationshipsAdyenPayments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdyenGatewayCreateDataRelationshipsAdyenPayments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
