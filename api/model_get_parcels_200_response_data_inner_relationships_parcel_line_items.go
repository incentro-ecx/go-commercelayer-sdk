/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETParcels200ResponseDataInnerRelationshipsParcelLineItems struct for GETParcels200ResponseDataInnerRelationshipsParcelLineItems
type GETParcels200ResponseDataInnerRelationshipsParcelLineItems struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks           `json:"links,omitempty"`
	Data  []GETParcels200ResponseDataInnerRelationshipsParcelLineItemsDataInner `json:"data,omitempty"`
}

// NewGETParcels200ResponseDataInnerRelationshipsParcelLineItems instantiates a new GETParcels200ResponseDataInnerRelationshipsParcelLineItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETParcels200ResponseDataInnerRelationshipsParcelLineItems() *GETParcels200ResponseDataInnerRelationshipsParcelLineItems {
	this := GETParcels200ResponseDataInnerRelationshipsParcelLineItems{}
	return &this
}

// NewGETParcels200ResponseDataInnerRelationshipsParcelLineItemsWithDefaults instantiates a new GETParcels200ResponseDataInnerRelationshipsParcelLineItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETParcels200ResponseDataInnerRelationshipsParcelLineItemsWithDefaults() *GETParcels200ResponseDataInnerRelationshipsParcelLineItems {
	this := GETParcels200ResponseDataInnerRelationshipsParcelLineItems{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETParcels200ResponseDataInnerRelationshipsParcelLineItems) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETParcels200ResponseDataInnerRelationshipsParcelLineItems) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETParcels200ResponseDataInnerRelationshipsParcelLineItems) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETParcels200ResponseDataInnerRelationshipsParcelLineItems) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETParcels200ResponseDataInnerRelationshipsParcelLineItems) GetData() []GETParcels200ResponseDataInnerRelationshipsParcelLineItemsDataInner {
	if o == nil || o.Data == nil {
		var ret []GETParcels200ResponseDataInnerRelationshipsParcelLineItemsDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETParcels200ResponseDataInnerRelationshipsParcelLineItems) GetDataOk() ([]GETParcels200ResponseDataInnerRelationshipsParcelLineItemsDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETParcels200ResponseDataInnerRelationshipsParcelLineItems) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETParcels200ResponseDataInnerRelationshipsParcelLineItemsDataInner and assigns it to the Data field.
func (o *GETParcels200ResponseDataInnerRelationshipsParcelLineItems) SetData(v []GETParcels200ResponseDataInnerRelationshipsParcelLineItemsDataInner) {
	o.Data = v
}

func (o GETParcels200ResponseDataInnerRelationshipsParcelLineItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETParcels200ResponseDataInnerRelationshipsParcelLineItems struct {
	value *GETParcels200ResponseDataInnerRelationshipsParcelLineItems
	isSet bool
}

func (v NullableGETParcels200ResponseDataInnerRelationshipsParcelLineItems) Get() *GETParcels200ResponseDataInnerRelationshipsParcelLineItems {
	return v.value
}

func (v *NullableGETParcels200ResponseDataInnerRelationshipsParcelLineItems) Set(val *GETParcels200ResponseDataInnerRelationshipsParcelLineItems) {
	v.value = val
	v.isSet = true
}

func (v NullableGETParcels200ResponseDataInnerRelationshipsParcelLineItems) IsSet() bool {
	return v.isSet
}

func (v *NullableGETParcels200ResponseDataInnerRelationshipsParcelLineItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETParcels200ResponseDataInnerRelationshipsParcelLineItems(val *GETParcels200ResponseDataInnerRelationshipsParcelLineItems) *NullableGETParcels200ResponseDataInnerRelationshipsParcelLineItems {
	return &NullableGETParcels200ResponseDataInnerRelationshipsParcelLineItems{value: val, isSet: true}
}

func (v NullableGETParcels200ResponseDataInnerRelationshipsParcelLineItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETParcels200ResponseDataInnerRelationshipsParcelLineItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
