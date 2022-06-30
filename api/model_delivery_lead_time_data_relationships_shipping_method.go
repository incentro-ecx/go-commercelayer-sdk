/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.5
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DeliveryLeadTimeDataRelationshipsShippingMethod struct for DeliveryLeadTimeDataRelationshipsShippingMethod
type DeliveryLeadTimeDataRelationshipsShippingMethod struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewDeliveryLeadTimeDataRelationshipsShippingMethod instantiates a new DeliveryLeadTimeDataRelationshipsShippingMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryLeadTimeDataRelationshipsShippingMethod(type_ string, id string) *DeliveryLeadTimeDataRelationshipsShippingMethod {
	this := DeliveryLeadTimeDataRelationshipsShippingMethod{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewDeliveryLeadTimeDataRelationshipsShippingMethodWithDefaults instantiates a new DeliveryLeadTimeDataRelationshipsShippingMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryLeadTimeDataRelationshipsShippingMethodWithDefaults() *DeliveryLeadTimeDataRelationshipsShippingMethod {
	this := DeliveryLeadTimeDataRelationshipsShippingMethod{}
	var type_ string = "shipping_methods"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *DeliveryLeadTimeDataRelationshipsShippingMethod) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTimeDataRelationshipsShippingMethod) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DeliveryLeadTimeDataRelationshipsShippingMethod) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *DeliveryLeadTimeDataRelationshipsShippingMethod) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTimeDataRelationshipsShippingMethod) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeliveryLeadTimeDataRelationshipsShippingMethod) SetId(v string) {
	o.Id = v
}

func (o DeliveryLeadTimeDataRelationshipsShippingMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableDeliveryLeadTimeDataRelationshipsShippingMethod struct {
	value *DeliveryLeadTimeDataRelationshipsShippingMethod
	isSet bool
}

func (v NullableDeliveryLeadTimeDataRelationshipsShippingMethod) Get() *DeliveryLeadTimeDataRelationshipsShippingMethod {
	return v.value
}

func (v *NullableDeliveryLeadTimeDataRelationshipsShippingMethod) Set(val *DeliveryLeadTimeDataRelationshipsShippingMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryLeadTimeDataRelationshipsShippingMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryLeadTimeDataRelationshipsShippingMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryLeadTimeDataRelationshipsShippingMethod(val *DeliveryLeadTimeDataRelationshipsShippingMethod) *NullableDeliveryLeadTimeDataRelationshipsShippingMethod {
	return &NullableDeliveryLeadTimeDataRelationshipsShippingMethod{value: val, isSet: true}
}

func (v NullableDeliveryLeadTimeDataRelationshipsShippingMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryLeadTimeDataRelationshipsShippingMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
