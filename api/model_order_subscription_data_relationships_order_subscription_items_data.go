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

// checks if the OrderSubscriptionDataRelationshipsOrderSubscriptionItemsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderSubscriptionDataRelationshipsOrderSubscriptionItemsData{}

// OrderSubscriptionDataRelationshipsOrderSubscriptionItemsData struct for OrderSubscriptionDataRelationshipsOrderSubscriptionItemsData
type OrderSubscriptionDataRelationshipsOrderSubscriptionItemsData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource's id
	Id interface{} `json:"id,omitempty"`
}

// NewOrderSubscriptionDataRelationshipsOrderSubscriptionItemsData instantiates a new OrderSubscriptionDataRelationshipsOrderSubscriptionItemsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSubscriptionDataRelationshipsOrderSubscriptionItemsData() *OrderSubscriptionDataRelationshipsOrderSubscriptionItemsData {
	this := OrderSubscriptionDataRelationshipsOrderSubscriptionItemsData{}
	return &this
}

// NewOrderSubscriptionDataRelationshipsOrderSubscriptionItemsDataWithDefaults instantiates a new OrderSubscriptionDataRelationshipsOrderSubscriptionItemsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSubscriptionDataRelationshipsOrderSubscriptionItemsDataWithDefaults() *OrderSubscriptionDataRelationshipsOrderSubscriptionItemsData {
	this := OrderSubscriptionDataRelationshipsOrderSubscriptionItemsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderSubscriptionDataRelationshipsOrderSubscriptionItemsData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderSubscriptionDataRelationshipsOrderSubscriptionItemsData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrderSubscriptionDataRelationshipsOrderSubscriptionItemsData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *OrderSubscriptionDataRelationshipsOrderSubscriptionItemsData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderSubscriptionDataRelationshipsOrderSubscriptionItemsData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderSubscriptionDataRelationshipsOrderSubscriptionItemsData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderSubscriptionDataRelationshipsOrderSubscriptionItemsData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *OrderSubscriptionDataRelationshipsOrderSubscriptionItemsData) SetId(v interface{}) {
	o.Id = v
}

func (o OrderSubscriptionDataRelationshipsOrderSubscriptionItemsData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderSubscriptionDataRelationshipsOrderSubscriptionItemsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableOrderSubscriptionDataRelationshipsOrderSubscriptionItemsData struct {
	value *OrderSubscriptionDataRelationshipsOrderSubscriptionItemsData
	isSet bool
}

func (v NullableOrderSubscriptionDataRelationshipsOrderSubscriptionItemsData) Get() *OrderSubscriptionDataRelationshipsOrderSubscriptionItemsData {
	return v.value
}

func (v *NullableOrderSubscriptionDataRelationshipsOrderSubscriptionItemsData) Set(val *OrderSubscriptionDataRelationshipsOrderSubscriptionItemsData) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSubscriptionDataRelationshipsOrderSubscriptionItemsData) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSubscriptionDataRelationshipsOrderSubscriptionItemsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSubscriptionDataRelationshipsOrderSubscriptionItemsData(val *OrderSubscriptionDataRelationshipsOrderSubscriptionItemsData) *NullableOrderSubscriptionDataRelationshipsOrderSubscriptionItemsData {
	return &NullableOrderSubscriptionDataRelationshipsOrderSubscriptionItemsData{value: val, isSet: true}
}

func (v NullableOrderSubscriptionDataRelationshipsOrderSubscriptionItemsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSubscriptionDataRelationshipsOrderSubscriptionItemsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
