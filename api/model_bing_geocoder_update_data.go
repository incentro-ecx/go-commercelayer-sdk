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

// BingGeocoderUpdateData struct for BingGeocoderUpdateData
type BingGeocoderUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                           `json:"id"`
	Attributes    BingGeocoderUpdateDataAttributes `json:"attributes"`
	Relationships map[string]interface{}           `json:"relationships,omitempty"`
}

// NewBingGeocoderUpdateData instantiates a new BingGeocoderUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBingGeocoderUpdateData(type_ string, id string, attributes BingGeocoderUpdateDataAttributes) *BingGeocoderUpdateData {
	this := BingGeocoderUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewBingGeocoderUpdateDataWithDefaults instantiates a new BingGeocoderUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBingGeocoderUpdateDataWithDefaults() *BingGeocoderUpdateData {
	this := BingGeocoderUpdateData{}
	var type_ string = "bing_geocoders"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *BingGeocoderUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BingGeocoderUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BingGeocoderUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *BingGeocoderUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BingGeocoderUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BingGeocoderUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *BingGeocoderUpdateData) GetAttributes() BingGeocoderUpdateDataAttributes {
	if o == nil {
		var ret BingGeocoderUpdateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *BingGeocoderUpdateData) GetAttributesOk() (*BingGeocoderUpdateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *BingGeocoderUpdateData) SetAttributes(v BingGeocoderUpdateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BingGeocoderUpdateData) GetRelationships() map[string]interface{} {
	if o == nil || o.Relationships == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BingGeocoderUpdateData) GetRelationshipsOk() (map[string]interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BingGeocoderUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]interface{} and assigns it to the Relationships field.
func (o *BingGeocoderUpdateData) SetRelationships(v map[string]interface{}) {
	o.Relationships = v
}

func (o BingGeocoderUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableBingGeocoderUpdateData struct {
	value *BingGeocoderUpdateData
	isSet bool
}

func (v NullableBingGeocoderUpdateData) Get() *BingGeocoderUpdateData {
	return v.value
}

func (v *NullableBingGeocoderUpdateData) Set(val *BingGeocoderUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableBingGeocoderUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableBingGeocoderUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBingGeocoderUpdateData(val *BingGeocoderUpdateData) *NullableBingGeocoderUpdateData {
	return &NullableBingGeocoderUpdateData{value: val, isSet: true}
}

func (v NullableBingGeocoderUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBingGeocoderUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}