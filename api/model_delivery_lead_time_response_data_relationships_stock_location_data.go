/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DeliveryLeadTimeResponseDataRelationshipsStockLocationData struct for DeliveryLeadTimeResponseDataRelationshipsStockLocationData
type DeliveryLeadTimeResponseDataRelationshipsStockLocationData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewDeliveryLeadTimeResponseDataRelationshipsStockLocationData instantiates a new DeliveryLeadTimeResponseDataRelationshipsStockLocationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryLeadTimeResponseDataRelationshipsStockLocationData() *DeliveryLeadTimeResponseDataRelationshipsStockLocationData {
	this := DeliveryLeadTimeResponseDataRelationshipsStockLocationData{}
	return &this
}

// NewDeliveryLeadTimeResponseDataRelationshipsStockLocationDataWithDefaults instantiates a new DeliveryLeadTimeResponseDataRelationshipsStockLocationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryLeadTimeResponseDataRelationshipsStockLocationDataWithDefaults() *DeliveryLeadTimeResponseDataRelationshipsStockLocationData {
	this := DeliveryLeadTimeResponseDataRelationshipsStockLocationData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DeliveryLeadTimeResponseDataRelationshipsStockLocationData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTimeResponseDataRelationshipsStockLocationData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DeliveryLeadTimeResponseDataRelationshipsStockLocationData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DeliveryLeadTimeResponseDataRelationshipsStockLocationData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeliveryLeadTimeResponseDataRelationshipsStockLocationData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTimeResponseDataRelationshipsStockLocationData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeliveryLeadTimeResponseDataRelationshipsStockLocationData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeliveryLeadTimeResponseDataRelationshipsStockLocationData) SetId(v string) {
	o.Id = &v
}

func (o DeliveryLeadTimeResponseDataRelationshipsStockLocationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableDeliveryLeadTimeResponseDataRelationshipsStockLocationData struct {
	value *DeliveryLeadTimeResponseDataRelationshipsStockLocationData
	isSet bool
}

func (v NullableDeliveryLeadTimeResponseDataRelationshipsStockLocationData) Get() *DeliveryLeadTimeResponseDataRelationshipsStockLocationData {
	return v.value
}

func (v *NullableDeliveryLeadTimeResponseDataRelationshipsStockLocationData) Set(val *DeliveryLeadTimeResponseDataRelationshipsStockLocationData) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryLeadTimeResponseDataRelationshipsStockLocationData) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryLeadTimeResponseDataRelationshipsStockLocationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryLeadTimeResponseDataRelationshipsStockLocationData(val *DeliveryLeadTimeResponseDataRelationshipsStockLocationData) *NullableDeliveryLeadTimeResponseDataRelationshipsStockLocationData {
	return &NullableDeliveryLeadTimeResponseDataRelationshipsStockLocationData{value: val, isSet: true}
}

func (v NullableDeliveryLeadTimeResponseDataRelationshipsStockLocationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryLeadTimeResponseDataRelationshipsStockLocationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}