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

// checks if the BraintreeGatewayDataRelationshipsBraintreePayments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BraintreeGatewayDataRelationshipsBraintreePayments{}

// BraintreeGatewayDataRelationshipsBraintreePayments struct for BraintreeGatewayDataRelationshipsBraintreePayments
type BraintreeGatewayDataRelationshipsBraintreePayments struct {
	Data *BraintreeGatewayDataRelationshipsBraintreePaymentsData `json:"data,omitempty"`
}

// NewBraintreeGatewayDataRelationshipsBraintreePayments instantiates a new BraintreeGatewayDataRelationshipsBraintreePayments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBraintreeGatewayDataRelationshipsBraintreePayments() *BraintreeGatewayDataRelationshipsBraintreePayments {
	this := BraintreeGatewayDataRelationshipsBraintreePayments{}
	return &this
}

// NewBraintreeGatewayDataRelationshipsBraintreePaymentsWithDefaults instantiates a new BraintreeGatewayDataRelationshipsBraintreePayments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBraintreeGatewayDataRelationshipsBraintreePaymentsWithDefaults() *BraintreeGatewayDataRelationshipsBraintreePayments {
	this := BraintreeGatewayDataRelationshipsBraintreePayments{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BraintreeGatewayDataRelationshipsBraintreePayments) GetData() BraintreeGatewayDataRelationshipsBraintreePaymentsData {
	if o == nil || IsNil(o.Data) {
		var ret BraintreeGatewayDataRelationshipsBraintreePaymentsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BraintreeGatewayDataRelationshipsBraintreePayments) GetDataOk() (*BraintreeGatewayDataRelationshipsBraintreePaymentsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BraintreeGatewayDataRelationshipsBraintreePayments) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given BraintreeGatewayDataRelationshipsBraintreePaymentsData and assigns it to the Data field.
func (o *BraintreeGatewayDataRelationshipsBraintreePayments) SetData(v BraintreeGatewayDataRelationshipsBraintreePaymentsData) {
	o.Data = &v
}

func (o BraintreeGatewayDataRelationshipsBraintreePayments) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BraintreeGatewayDataRelationshipsBraintreePayments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableBraintreeGatewayDataRelationshipsBraintreePayments struct {
	value *BraintreeGatewayDataRelationshipsBraintreePayments
	isSet bool
}

func (v NullableBraintreeGatewayDataRelationshipsBraintreePayments) Get() *BraintreeGatewayDataRelationshipsBraintreePayments {
	return v.value
}

func (v *NullableBraintreeGatewayDataRelationshipsBraintreePayments) Set(val *BraintreeGatewayDataRelationshipsBraintreePayments) {
	v.value = val
	v.isSet = true
}

func (v NullableBraintreeGatewayDataRelationshipsBraintreePayments) IsSet() bool {
	return v.isSet
}

func (v *NullableBraintreeGatewayDataRelationshipsBraintreePayments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBraintreeGatewayDataRelationshipsBraintreePayments(val *BraintreeGatewayDataRelationshipsBraintreePayments) *NullableBraintreeGatewayDataRelationshipsBraintreePayments {
	return &NullableBraintreeGatewayDataRelationshipsBraintreePayments{value: val, isSet: true}
}

func (v NullableBraintreeGatewayDataRelationshipsBraintreePayments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBraintreeGatewayDataRelationshipsBraintreePayments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
