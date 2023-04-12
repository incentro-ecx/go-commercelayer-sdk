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

// checks if the POSTOrders201ResponseDataRelationshipsShipments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrders201ResponseDataRelationshipsShipments{}

// POSTOrders201ResponseDataRelationshipsShipments struct for POSTOrders201ResponseDataRelationshipsShipments
type POSTOrders201ResponseDataRelationshipsShipments struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks `json:"links,omitempty"`
	Data  *POSTOrders201ResponseDataRelationshipsShipmentsData    `json:"data,omitempty"`
}

// NewPOSTOrders201ResponseDataRelationshipsShipments instantiates a new POSTOrders201ResponseDataRelationshipsShipments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrders201ResponseDataRelationshipsShipments() *POSTOrders201ResponseDataRelationshipsShipments {
	this := POSTOrders201ResponseDataRelationshipsShipments{}
	return &this
}

// NewPOSTOrders201ResponseDataRelationshipsShipmentsWithDefaults instantiates a new POSTOrders201ResponseDataRelationshipsShipments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrders201ResponseDataRelationshipsShipmentsWithDefaults() *POSTOrders201ResponseDataRelationshipsShipments {
	this := POSTOrders201ResponseDataRelationshipsShipments{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTOrders201ResponseDataRelationshipsShipments) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrders201ResponseDataRelationshipsShipments) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataRelationshipsShipments) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTOrders201ResponseDataRelationshipsShipments) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTOrders201ResponseDataRelationshipsShipments) GetData() POSTOrders201ResponseDataRelationshipsShipmentsData {
	if o == nil || IsNil(o.Data) {
		var ret POSTOrders201ResponseDataRelationshipsShipmentsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrders201ResponseDataRelationshipsShipments) GetDataOk() (*POSTOrders201ResponseDataRelationshipsShipmentsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataRelationshipsShipments) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTOrders201ResponseDataRelationshipsShipmentsData and assigns it to the Data field.
func (o *POSTOrders201ResponseDataRelationshipsShipments) SetData(v POSTOrders201ResponseDataRelationshipsShipmentsData) {
	o.Data = &v
}

func (o POSTOrders201ResponseDataRelationshipsShipments) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrders201ResponseDataRelationshipsShipments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTOrders201ResponseDataRelationshipsShipments struct {
	value *POSTOrders201ResponseDataRelationshipsShipments
	isSet bool
}

func (v NullablePOSTOrders201ResponseDataRelationshipsShipments) Get() *POSTOrders201ResponseDataRelationshipsShipments {
	return v.value
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsShipments) Set(val *POSTOrders201ResponseDataRelationshipsShipments) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrders201ResponseDataRelationshipsShipments) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsShipments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrders201ResponseDataRelationshipsShipments(val *POSTOrders201ResponseDataRelationshipsShipments) *NullablePOSTOrders201ResponseDataRelationshipsShipments {
	return &NullablePOSTOrders201ResponseDataRelationshipsShipments{value: val, isSet: true}
}

func (v NullablePOSTOrders201ResponseDataRelationshipsShipments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsShipments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}