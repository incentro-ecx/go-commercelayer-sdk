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

// GETInventoryReturnLocations200ResponseDataInnerRelationships struct for GETInventoryReturnLocations200ResponseDataInnerRelationships
type GETInventoryReturnLocations200ResponseDataInnerRelationships struct {
	StockLocation  *GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation         `json:"stock_location,omitempty"`
	InventoryModel *GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel `json:"inventory_model,omitempty"`
}

// NewGETInventoryReturnLocations200ResponseDataInnerRelationships instantiates a new GETInventoryReturnLocations200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETInventoryReturnLocations200ResponseDataInnerRelationships() *GETInventoryReturnLocations200ResponseDataInnerRelationships {
	this := GETInventoryReturnLocations200ResponseDataInnerRelationships{}
	return &this
}

// NewGETInventoryReturnLocations200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETInventoryReturnLocations200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETInventoryReturnLocations200ResponseDataInnerRelationshipsWithDefaults() *GETInventoryReturnLocations200ResponseDataInnerRelationships {
	this := GETInventoryReturnLocations200ResponseDataInnerRelationships{}
	return &this
}

// GetStockLocation returns the StockLocation field value if set, zero value otherwise.
func (o *GETInventoryReturnLocations200ResponseDataInnerRelationships) GetStockLocation() GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation {
	if o == nil || o.StockLocation == nil {
		var ret GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation
		return ret
	}
	return *o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInventoryReturnLocations200ResponseDataInnerRelationships) GetStockLocationOk() (*GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation, bool) {
	if o == nil || o.StockLocation == nil {
		return nil, false
	}
	return o.StockLocation, true
}

// HasStockLocation returns a boolean if a field has been set.
func (o *GETInventoryReturnLocations200ResponseDataInnerRelationships) HasStockLocation() bool {
	if o != nil && o.StockLocation != nil {
		return true
	}

	return false
}

// SetStockLocation gets a reference to the given GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation and assigns it to the StockLocation field.
func (o *GETInventoryReturnLocations200ResponseDataInnerRelationships) SetStockLocation(v GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation) {
	o.StockLocation = &v
}

// GetInventoryModel returns the InventoryModel field value if set, zero value otherwise.
func (o *GETInventoryReturnLocations200ResponseDataInnerRelationships) GetInventoryModel() GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel {
	if o == nil || o.InventoryModel == nil {
		var ret GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel
		return ret
	}
	return *o.InventoryModel
}

// GetInventoryModelOk returns a tuple with the InventoryModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInventoryReturnLocations200ResponseDataInnerRelationships) GetInventoryModelOk() (*GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel, bool) {
	if o == nil || o.InventoryModel == nil {
		return nil, false
	}
	return o.InventoryModel, true
}

// HasInventoryModel returns a boolean if a field has been set.
func (o *GETInventoryReturnLocations200ResponseDataInnerRelationships) HasInventoryModel() bool {
	if o != nil && o.InventoryModel != nil {
		return true
	}

	return false
}

// SetInventoryModel gets a reference to the given GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel and assigns it to the InventoryModel field.
func (o *GETInventoryReturnLocations200ResponseDataInnerRelationships) SetInventoryModel(v GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel) {
	o.InventoryModel = &v
}

func (o GETInventoryReturnLocations200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StockLocation != nil {
		toSerialize["stock_location"] = o.StockLocation
	}
	if o.InventoryModel != nil {
		toSerialize["inventory_model"] = o.InventoryModel
	}
	return json.Marshal(toSerialize)
}

type NullableGETInventoryReturnLocations200ResponseDataInnerRelationships struct {
	value *GETInventoryReturnLocations200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETInventoryReturnLocations200ResponseDataInnerRelationships) Get() *GETInventoryReturnLocations200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETInventoryReturnLocations200ResponseDataInnerRelationships) Set(val *GETInventoryReturnLocations200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETInventoryReturnLocations200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETInventoryReturnLocations200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETInventoryReturnLocations200ResponseDataInnerRelationships(val *GETInventoryReturnLocations200ResponseDataInnerRelationships) *NullableGETInventoryReturnLocations200ResponseDataInnerRelationships {
	return &NullableGETInventoryReturnLocations200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETInventoryReturnLocations200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETInventoryReturnLocations200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
