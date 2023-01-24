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

// PATCHExternalPromotionsExternalPromotionId200Response struct for PATCHExternalPromotionsExternalPromotionId200Response
type PATCHExternalPromotionsExternalPromotionId200Response struct {
	Data *PATCHExternalPromotionsExternalPromotionId200ResponseData `json:"data,omitempty"`
}

// NewPATCHExternalPromotionsExternalPromotionId200Response instantiates a new PATCHExternalPromotionsExternalPromotionId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHExternalPromotionsExternalPromotionId200Response() *PATCHExternalPromotionsExternalPromotionId200Response {
	this := PATCHExternalPromotionsExternalPromotionId200Response{}
	return &this
}

// NewPATCHExternalPromotionsExternalPromotionId200ResponseWithDefaults instantiates a new PATCHExternalPromotionsExternalPromotionId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHExternalPromotionsExternalPromotionId200ResponseWithDefaults() *PATCHExternalPromotionsExternalPromotionId200Response {
	this := PATCHExternalPromotionsExternalPromotionId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHExternalPromotionsExternalPromotionId200Response) GetData() PATCHExternalPromotionsExternalPromotionId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHExternalPromotionsExternalPromotionId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHExternalPromotionsExternalPromotionId200Response) GetDataOk() (*PATCHExternalPromotionsExternalPromotionId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHExternalPromotionsExternalPromotionId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHExternalPromotionsExternalPromotionId200ResponseData and assigns it to the Data field.
func (o *PATCHExternalPromotionsExternalPromotionId200Response) SetData(v PATCHExternalPromotionsExternalPromotionId200ResponseData) {
	o.Data = &v
}

func (o PATCHExternalPromotionsExternalPromotionId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHExternalPromotionsExternalPromotionId200Response struct {
	value *PATCHExternalPromotionsExternalPromotionId200Response
	isSet bool
}

func (v NullablePATCHExternalPromotionsExternalPromotionId200Response) Get() *PATCHExternalPromotionsExternalPromotionId200Response {
	return v.value
}

func (v *NullablePATCHExternalPromotionsExternalPromotionId200Response) Set(val *PATCHExternalPromotionsExternalPromotionId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHExternalPromotionsExternalPromotionId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHExternalPromotionsExternalPromotionId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHExternalPromotionsExternalPromotionId200Response(val *PATCHExternalPromotionsExternalPromotionId200Response) *NullablePATCHExternalPromotionsExternalPromotionId200Response {
	return &NullablePATCHExternalPromotionsExternalPromotionId200Response{value: val, isSet: true}
}

func (v NullablePATCHExternalPromotionsExternalPromotionId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHExternalPromotionsExternalPromotionId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
