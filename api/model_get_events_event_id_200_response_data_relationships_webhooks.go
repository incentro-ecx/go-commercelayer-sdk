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

// checks if the GETEventsEventId200ResponseDataRelationshipsWebhooks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETEventsEventId200ResponseDataRelationshipsWebhooks{}

// GETEventsEventId200ResponseDataRelationshipsWebhooks struct for GETEventsEventId200ResponseDataRelationshipsWebhooks
type GETEventsEventId200ResponseDataRelationshipsWebhooks struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks   `json:"links,omitempty"`
	Data  *GETEventsEventId200ResponseDataRelationshipsWebhooksData `json:"data,omitempty"`
}

// NewGETEventsEventId200ResponseDataRelationshipsWebhooks instantiates a new GETEventsEventId200ResponseDataRelationshipsWebhooks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETEventsEventId200ResponseDataRelationshipsWebhooks() *GETEventsEventId200ResponseDataRelationshipsWebhooks {
	this := GETEventsEventId200ResponseDataRelationshipsWebhooks{}
	return &this
}

// NewGETEventsEventId200ResponseDataRelationshipsWebhooksWithDefaults instantiates a new GETEventsEventId200ResponseDataRelationshipsWebhooks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETEventsEventId200ResponseDataRelationshipsWebhooksWithDefaults() *GETEventsEventId200ResponseDataRelationshipsWebhooks {
	this := GETEventsEventId200ResponseDataRelationshipsWebhooks{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETEventsEventId200ResponseDataRelationshipsWebhooks) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETEventsEventId200ResponseDataRelationshipsWebhooks) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETEventsEventId200ResponseDataRelationshipsWebhooks) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETEventsEventId200ResponseDataRelationshipsWebhooks) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETEventsEventId200ResponseDataRelationshipsWebhooks) GetData() GETEventsEventId200ResponseDataRelationshipsWebhooksData {
	if o == nil || IsNil(o.Data) {
		var ret GETEventsEventId200ResponseDataRelationshipsWebhooksData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETEventsEventId200ResponseDataRelationshipsWebhooks) GetDataOk() (*GETEventsEventId200ResponseDataRelationshipsWebhooksData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETEventsEventId200ResponseDataRelationshipsWebhooks) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETEventsEventId200ResponseDataRelationshipsWebhooksData and assigns it to the Data field.
func (o *GETEventsEventId200ResponseDataRelationshipsWebhooks) SetData(v GETEventsEventId200ResponseDataRelationshipsWebhooksData) {
	o.Data = &v
}

func (o GETEventsEventId200ResponseDataRelationshipsWebhooks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETEventsEventId200ResponseDataRelationshipsWebhooks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETEventsEventId200ResponseDataRelationshipsWebhooks struct {
	value *GETEventsEventId200ResponseDataRelationshipsWebhooks
	isSet bool
}

func (v NullableGETEventsEventId200ResponseDataRelationshipsWebhooks) Get() *GETEventsEventId200ResponseDataRelationshipsWebhooks {
	return v.value
}

func (v *NullableGETEventsEventId200ResponseDataRelationshipsWebhooks) Set(val *GETEventsEventId200ResponseDataRelationshipsWebhooks) {
	v.value = val
	v.isSet = true
}

func (v NullableGETEventsEventId200ResponseDataRelationshipsWebhooks) IsSet() bool {
	return v.isSet
}

func (v *NullableGETEventsEventId200ResponseDataRelationshipsWebhooks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETEventsEventId200ResponseDataRelationshipsWebhooks(val *GETEventsEventId200ResponseDataRelationshipsWebhooks) *NullableGETEventsEventId200ResponseDataRelationshipsWebhooks {
	return &NullableGETEventsEventId200ResponseDataRelationshipsWebhooks{value: val, isSet: true}
}

func (v NullableGETEventsEventId200ResponseDataRelationshipsWebhooks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETEventsEventId200ResponseDataRelationshipsWebhooks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
