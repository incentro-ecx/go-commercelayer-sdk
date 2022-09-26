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

// GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation struct for GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation
type GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks             `json:"links,omitempty"`
	Data  *GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocationData `json:"data,omitempty"`
}

// NewGETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation instantiates a new GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation() *GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation {
	this := GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation{}
	return &this
}

// NewGETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocationWithDefaults instantiates a new GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocationWithDefaults() *GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation {
	this := GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation) GetData() GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocationData {
	if o == nil || o.Data == nil {
		var ret GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocationData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation) GetDataOk() (*GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocationData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocationData and assigns it to the Data field.
func (o *GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation) SetData(v GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocationData) {
	o.Data = &v
}

func (o GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation struct {
	value *GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation
	isSet bool
}

func (v NullableGETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation) Get() *GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation {
	return v.value
}

func (v *NullableGETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation) Set(val *GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableGETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableGETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation(val *GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation) *NullableGETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation {
	return &NullableGETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation{value: val, isSet: true}
}

func (v NullableGETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}