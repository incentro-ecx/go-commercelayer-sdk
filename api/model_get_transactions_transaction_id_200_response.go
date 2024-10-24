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

// checks if the GETTransactionsTransactionId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETTransactionsTransactionId200Response{}

// GETTransactionsTransactionId200Response struct for GETTransactionsTransactionId200Response
type GETTransactionsTransactionId200Response struct {
	Data *GETTransactionsTransactionId200ResponseData `json:"data,omitempty"`
}

// NewGETTransactionsTransactionId200Response instantiates a new GETTransactionsTransactionId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETTransactionsTransactionId200Response() *GETTransactionsTransactionId200Response {
	this := GETTransactionsTransactionId200Response{}
	return &this
}

// NewGETTransactionsTransactionId200ResponseWithDefaults instantiates a new GETTransactionsTransactionId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETTransactionsTransactionId200ResponseWithDefaults() *GETTransactionsTransactionId200Response {
	this := GETTransactionsTransactionId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETTransactionsTransactionId200Response) GetData() GETTransactionsTransactionId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETTransactionsTransactionId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTransactionsTransactionId200Response) GetDataOk() (*GETTransactionsTransactionId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETTransactionsTransactionId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETTransactionsTransactionId200ResponseData and assigns it to the Data field.
func (o *GETTransactionsTransactionId200Response) SetData(v GETTransactionsTransactionId200ResponseData) {
	o.Data = &v
}

func (o GETTransactionsTransactionId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETTransactionsTransactionId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETTransactionsTransactionId200Response struct {
	value *GETTransactionsTransactionId200Response
	isSet bool
}

func (v NullableGETTransactionsTransactionId200Response) Get() *GETTransactionsTransactionId200Response {
	return v.value
}

func (v *NullableGETTransactionsTransactionId200Response) Set(val *GETTransactionsTransactionId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETTransactionsTransactionId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETTransactionsTransactionId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETTransactionsTransactionId200Response(val *GETTransactionsTransactionId200Response) *NullableGETTransactionsTransactionId200Response {
	return &NullableGETTransactionsTransactionId200Response{value: val, isSet: true}
}

func (v NullableGETTransactionsTransactionId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETTransactionsTransactionId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
