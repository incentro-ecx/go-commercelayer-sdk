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

// GETEventsEventId200Response struct for GETEventsEventId200Response
type GETEventsEventId200Response struct {
	Data *GETEvents200ResponseDataInner `json:"data,omitempty"`
}

// NewGETEventsEventId200Response instantiates a new GETEventsEventId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETEventsEventId200Response() *GETEventsEventId200Response {
	this := GETEventsEventId200Response{}
	return &this
}

// NewGETEventsEventId200ResponseWithDefaults instantiates a new GETEventsEventId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETEventsEventId200ResponseWithDefaults() *GETEventsEventId200Response {
	this := GETEventsEventId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETEventsEventId200Response) GetData() GETEvents200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETEvents200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETEventsEventId200Response) GetDataOk() (*GETEvents200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETEventsEventId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETEvents200ResponseDataInner and assigns it to the Data field.
func (o *GETEventsEventId200Response) SetData(v GETEvents200ResponseDataInner) {
	o.Data = &v
}

func (o GETEventsEventId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETEventsEventId200Response struct {
	value *GETEventsEventId200Response
	isSet bool
}

func (v NullableGETEventsEventId200Response) Get() *GETEventsEventId200Response {
	return v.value
}

func (v *NullableGETEventsEventId200Response) Set(val *GETEventsEventId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETEventsEventId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETEventsEventId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETEventsEventId200Response(val *GETEventsEventId200Response) *NullableGETEventsEventId200Response {
	return &NullableGETEventsEventId200Response{value: val, isSet: true}
}

func (v NullableGETEventsEventId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETEventsEventId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}