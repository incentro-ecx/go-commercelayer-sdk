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

// GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization struct for GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization
type GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks             `json:"links,omitempty"`
	Data  *GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorizationData `json:"data,omitempty"`
}

// NewGETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization instantiates a new GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization() *GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization {
	this := GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization{}
	return &this
}

// NewGETCaptures200ResponseDataInnerRelationshipsReferenceAuthorizationWithDefaults instantiates a new GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCaptures200ResponseDataInnerRelationshipsReferenceAuthorizationWithDefaults() *GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization {
	this := GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization) GetData() GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorizationData {
	if o == nil || o.Data == nil {
		var ret GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorizationData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization) GetDataOk() (*GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorizationData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorizationData and assigns it to the Data field.
func (o *GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization) SetData(v GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorizationData) {
	o.Data = &v
}

func (o GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization struct {
	value *GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization
	isSet bool
}

func (v NullableGETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization) Get() *GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization {
	return v.value
}

func (v *NullableGETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization) Set(val *GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization(val *GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization) *NullableGETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization {
	return &NullableGETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization{value: val, isSet: true}
}

func (v NullableGETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
