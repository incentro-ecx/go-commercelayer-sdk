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

// checks if the POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations{}

// POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations struct for POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations
type POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks                     `json:"links,omitempty"`
	Data  *POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocationsData `json:"data,omitempty"`
}

// NewPOSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations instantiates a new POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations() *POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations {
	this := POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations{}
	return &this
}

// NewPOSTInventoryModels201ResponseDataRelationshipsInventoryStockLocationsWithDefaults instantiates a new POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTInventoryModels201ResponseDataRelationshipsInventoryStockLocationsWithDefaults() *POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations {
	this := POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations) GetData() POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocationsData {
	if o == nil || IsNil(o.Data) {
		var ret POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocationsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations) GetDataOk() (*POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocationsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocationsData and assigns it to the Data field.
func (o *POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations) SetData(v POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocationsData) {
	o.Data = &v
}

func (o POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations struct {
	value *POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations
	isSet bool
}

func (v NullablePOSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations) Get() *POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations {
	return v.value
}

func (v *NullablePOSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations) Set(val *POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations(val *POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations) *NullablePOSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations {
	return &NullablePOSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations{value: val, isSet: true}
}

func (v NullablePOSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
