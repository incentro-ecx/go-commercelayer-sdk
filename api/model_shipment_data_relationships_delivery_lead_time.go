/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ShipmentDataRelationshipsDeliveryLeadTime struct for ShipmentDataRelationshipsDeliveryLeadTime
type ShipmentDataRelationshipsDeliveryLeadTime struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewShipmentDataRelationshipsDeliveryLeadTime instantiates a new ShipmentDataRelationshipsDeliveryLeadTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentDataRelationshipsDeliveryLeadTime(type_ string, id string) *ShipmentDataRelationshipsDeliveryLeadTime {
	this := ShipmentDataRelationshipsDeliveryLeadTime{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewShipmentDataRelationshipsDeliveryLeadTimeWithDefaults instantiates a new ShipmentDataRelationshipsDeliveryLeadTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentDataRelationshipsDeliveryLeadTimeWithDefaults() *ShipmentDataRelationshipsDeliveryLeadTime {
	this := ShipmentDataRelationshipsDeliveryLeadTime{}
	var type_ string = "delivery_lead_times"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *ShipmentDataRelationshipsDeliveryLeadTime) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ShipmentDataRelationshipsDeliveryLeadTime) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ShipmentDataRelationshipsDeliveryLeadTime) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *ShipmentDataRelationshipsDeliveryLeadTime) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ShipmentDataRelationshipsDeliveryLeadTime) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ShipmentDataRelationshipsDeliveryLeadTime) SetId(v string) {
	o.Id = v
}

func (o ShipmentDataRelationshipsDeliveryLeadTime) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableShipmentDataRelationshipsDeliveryLeadTime struct {
	value *ShipmentDataRelationshipsDeliveryLeadTime
	isSet bool
}

func (v NullableShipmentDataRelationshipsDeliveryLeadTime) Get() *ShipmentDataRelationshipsDeliveryLeadTime {
	return v.value
}

func (v *NullableShipmentDataRelationshipsDeliveryLeadTime) Set(val *ShipmentDataRelationshipsDeliveryLeadTime) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentDataRelationshipsDeliveryLeadTime) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentDataRelationshipsDeliveryLeadTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentDataRelationshipsDeliveryLeadTime(val *ShipmentDataRelationshipsDeliveryLeadTime) *NullableShipmentDataRelationshipsDeliveryLeadTime {
	return &NullableShipmentDataRelationshipsDeliveryLeadTime{value: val, isSet: true}
}

func (v NullableShipmentDataRelationshipsDeliveryLeadTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentDataRelationshipsDeliveryLeadTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}