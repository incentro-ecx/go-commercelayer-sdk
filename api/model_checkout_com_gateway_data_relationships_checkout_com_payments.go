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

// checks if the CheckoutComGatewayDataRelationshipsCheckoutComPayments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckoutComGatewayDataRelationshipsCheckoutComPayments{}

// CheckoutComGatewayDataRelationshipsCheckoutComPayments struct for CheckoutComGatewayDataRelationshipsCheckoutComPayments
type CheckoutComGatewayDataRelationshipsCheckoutComPayments struct {
	Data *CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData `json:"data,omitempty"`
}

// NewCheckoutComGatewayDataRelationshipsCheckoutComPayments instantiates a new CheckoutComGatewayDataRelationshipsCheckoutComPayments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutComGatewayDataRelationshipsCheckoutComPayments() *CheckoutComGatewayDataRelationshipsCheckoutComPayments {
	this := CheckoutComGatewayDataRelationshipsCheckoutComPayments{}
	return &this
}

// NewCheckoutComGatewayDataRelationshipsCheckoutComPaymentsWithDefaults instantiates a new CheckoutComGatewayDataRelationshipsCheckoutComPayments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutComGatewayDataRelationshipsCheckoutComPaymentsWithDefaults() *CheckoutComGatewayDataRelationshipsCheckoutComPayments {
	this := CheckoutComGatewayDataRelationshipsCheckoutComPayments{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CheckoutComGatewayDataRelationshipsCheckoutComPayments) GetData() CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData {
	if o == nil || IsNil(o.Data) {
		var ret CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutComGatewayDataRelationshipsCheckoutComPayments) GetDataOk() (*CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CheckoutComGatewayDataRelationshipsCheckoutComPayments) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData and assigns it to the Data field.
func (o *CheckoutComGatewayDataRelationshipsCheckoutComPayments) SetData(v CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData) {
	o.Data = &v
}

func (o CheckoutComGatewayDataRelationshipsCheckoutComPayments) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutComGatewayDataRelationshipsCheckoutComPayments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCheckoutComGatewayDataRelationshipsCheckoutComPayments struct {
	value *CheckoutComGatewayDataRelationshipsCheckoutComPayments
	isSet bool
}

func (v NullableCheckoutComGatewayDataRelationshipsCheckoutComPayments) Get() *CheckoutComGatewayDataRelationshipsCheckoutComPayments {
	return v.value
}

func (v *NullableCheckoutComGatewayDataRelationshipsCheckoutComPayments) Set(val *CheckoutComGatewayDataRelationshipsCheckoutComPayments) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutComGatewayDataRelationshipsCheckoutComPayments) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutComGatewayDataRelationshipsCheckoutComPayments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutComGatewayDataRelationshipsCheckoutComPayments(val *CheckoutComGatewayDataRelationshipsCheckoutComPayments) *NullableCheckoutComGatewayDataRelationshipsCheckoutComPayments {
	return &NullableCheckoutComGatewayDataRelationshipsCheckoutComPayments{value: val, isSet: true}
}

func (v NullableCheckoutComGatewayDataRelationshipsCheckoutComPayments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutComGatewayDataRelationshipsCheckoutComPayments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
