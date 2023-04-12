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

// checks if the OrderSubscriptionDataRelationshipsOrderFactoriesData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderSubscriptionDataRelationshipsOrderFactoriesData{}

// OrderSubscriptionDataRelationshipsOrderFactoriesData struct for OrderSubscriptionDataRelationshipsOrderFactoriesData
type OrderSubscriptionDataRelationshipsOrderFactoriesData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource's id
	Id interface{} `json:"id,omitempty"`
}

// NewOrderSubscriptionDataRelationshipsOrderFactoriesData instantiates a new OrderSubscriptionDataRelationshipsOrderFactoriesData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSubscriptionDataRelationshipsOrderFactoriesData() *OrderSubscriptionDataRelationshipsOrderFactoriesData {
	this := OrderSubscriptionDataRelationshipsOrderFactoriesData{}
	return &this
}

// NewOrderSubscriptionDataRelationshipsOrderFactoriesDataWithDefaults instantiates a new OrderSubscriptionDataRelationshipsOrderFactoriesData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSubscriptionDataRelationshipsOrderFactoriesDataWithDefaults() *OrderSubscriptionDataRelationshipsOrderFactoriesData {
	this := OrderSubscriptionDataRelationshipsOrderFactoriesData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderSubscriptionDataRelationshipsOrderFactoriesData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderSubscriptionDataRelationshipsOrderFactoriesData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrderSubscriptionDataRelationshipsOrderFactoriesData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *OrderSubscriptionDataRelationshipsOrderFactoriesData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderSubscriptionDataRelationshipsOrderFactoriesData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderSubscriptionDataRelationshipsOrderFactoriesData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderSubscriptionDataRelationshipsOrderFactoriesData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *OrderSubscriptionDataRelationshipsOrderFactoriesData) SetId(v interface{}) {
	o.Id = v
}

func (o OrderSubscriptionDataRelationshipsOrderFactoriesData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderSubscriptionDataRelationshipsOrderFactoriesData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableOrderSubscriptionDataRelationshipsOrderFactoriesData struct {
	value *OrderSubscriptionDataRelationshipsOrderFactoriesData
	isSet bool
}

func (v NullableOrderSubscriptionDataRelationshipsOrderFactoriesData) Get() *OrderSubscriptionDataRelationshipsOrderFactoriesData {
	return v.value
}

func (v *NullableOrderSubscriptionDataRelationshipsOrderFactoriesData) Set(val *OrderSubscriptionDataRelationshipsOrderFactoriesData) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSubscriptionDataRelationshipsOrderFactoriesData) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSubscriptionDataRelationshipsOrderFactoriesData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSubscriptionDataRelationshipsOrderFactoriesData(val *OrderSubscriptionDataRelationshipsOrderFactoriesData) *NullableOrderSubscriptionDataRelationshipsOrderFactoriesData {
	return &NullableOrderSubscriptionDataRelationshipsOrderFactoriesData{value: val, isSet: true}
}

func (v NullableOrderSubscriptionDataRelationshipsOrderFactoriesData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSubscriptionDataRelationshipsOrderFactoriesData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
