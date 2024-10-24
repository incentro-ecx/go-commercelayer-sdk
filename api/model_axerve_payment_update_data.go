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

// checks if the AxervePaymentUpdateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AxervePaymentUpdateData{}

// AxervePaymentUpdateData struct for AxervePaymentUpdateData
type AxervePaymentUpdateData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                                 `json:"id"`
	Attributes    PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes `json:"attributes"`
	Relationships *AdyenPaymentUpdateDataRelationships                        `json:"relationships,omitempty"`
}

// NewAxervePaymentUpdateData instantiates a new AxervePaymentUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAxervePaymentUpdateData(type_ interface{}, id interface{}, attributes PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes) *AxervePaymentUpdateData {
	this := AxervePaymentUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewAxervePaymentUpdateDataWithDefaults instantiates a new AxervePaymentUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAxervePaymentUpdateDataWithDefaults() *AxervePaymentUpdateData {
	this := AxervePaymentUpdateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *AxervePaymentUpdateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AxervePaymentUpdateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AxervePaymentUpdateData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *AxervePaymentUpdateData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AxervePaymentUpdateData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AxervePaymentUpdateData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *AxervePaymentUpdateData) GetAttributes() PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AxervePaymentUpdateData) GetAttributesOk() (*PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AxervePaymentUpdateData) SetAttributes(v PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AxervePaymentUpdateData) GetRelationships() AdyenPaymentUpdateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AdyenPaymentUpdateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AxervePaymentUpdateData) GetRelationshipsOk() (*AdyenPaymentUpdateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AxervePaymentUpdateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AdyenPaymentUpdateDataRelationships and assigns it to the Relationships field.
func (o *AxervePaymentUpdateData) SetRelationships(v AdyenPaymentUpdateDataRelationships) {
	o.Relationships = &v
}

func (o AxervePaymentUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AxervePaymentUpdateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableAxervePaymentUpdateData struct {
	value *AxervePaymentUpdateData
	isSet bool
}

func (v NullableAxervePaymentUpdateData) Get() *AxervePaymentUpdateData {
	return v.value
}

func (v *NullableAxervePaymentUpdateData) Set(val *AxervePaymentUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableAxervePaymentUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableAxervePaymentUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAxervePaymentUpdateData(val *AxervePaymentUpdateData) *NullableAxervePaymentUpdateData {
	return &NullableAxervePaymentUpdateData{value: val, isSet: true}
}

func (v NullableAxervePaymentUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAxervePaymentUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
