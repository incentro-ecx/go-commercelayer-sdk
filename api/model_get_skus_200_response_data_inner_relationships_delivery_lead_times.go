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

// GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes struct for GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes
type GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks    `json:"links,omitempty"`
	Data  *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData `json:"data,omitempty"`
}

// NewGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes instantiates a new GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes() *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes {
	this := GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes{}
	return &this
}

// NewGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesWithDefaults instantiates a new GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesWithDefaults() *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes {
	this := GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes) GetData() GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData {
	if o == nil || o.Data == nil {
		var ret GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes) GetDataOk() (*GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData and assigns it to the Data field.
func (o *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes) SetData(v GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData) {
	o.Data = &v
}

func (o GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes struct {
	value *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes
	isSet bool
}

func (v NullableGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes) Get() *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes {
	return v.value
}

func (v *NullableGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes) Set(val *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes(val *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes) *NullableGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes {
	return &NullableGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes{value: val, isSet: true}
}

func (v NullableGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
