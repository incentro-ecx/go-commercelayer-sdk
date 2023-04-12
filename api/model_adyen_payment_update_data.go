/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AdyenPaymentUpdateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdyenPaymentUpdateData{}

// AdyenPaymentUpdateData struct for AdyenPaymentUpdateData
type AdyenPaymentUpdateData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                               `json:"id"`
	Attributes    PATCHAdyenPaymentsAdyenPaymentId200ResponseDataAttributes `json:"attributes"`
	Relationships *AdyenPaymentUpdateDataRelationships                      `json:"relationships,omitempty"`
}

// NewAdyenPaymentUpdateData instantiates a new AdyenPaymentUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdyenPaymentUpdateData(type_ interface{}, id interface{}, attributes PATCHAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) *AdyenPaymentUpdateData {
	this := AdyenPaymentUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewAdyenPaymentUpdateDataWithDefaults instantiates a new AdyenPaymentUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdyenPaymentUpdateDataWithDefaults() *AdyenPaymentUpdateData {
	this := AdyenPaymentUpdateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *AdyenPaymentUpdateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdyenPaymentUpdateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AdyenPaymentUpdateData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *AdyenPaymentUpdateData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdyenPaymentUpdateData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AdyenPaymentUpdateData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *AdyenPaymentUpdateData) GetAttributes() PATCHAdyenPaymentsAdyenPaymentId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHAdyenPaymentsAdyenPaymentId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AdyenPaymentUpdateData) GetAttributesOk() (*PATCHAdyenPaymentsAdyenPaymentId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AdyenPaymentUpdateData) SetAttributes(v PATCHAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AdyenPaymentUpdateData) GetRelationships() AdyenPaymentUpdateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AdyenPaymentUpdateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenPaymentUpdateData) GetRelationshipsOk() (*AdyenPaymentUpdateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AdyenPaymentUpdateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AdyenPaymentUpdateDataRelationships and assigns it to the Relationships field.
func (o *AdyenPaymentUpdateData) SetRelationships(v AdyenPaymentUpdateDataRelationships) {
	o.Relationships = &v
}

func (o AdyenPaymentUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdyenPaymentUpdateData) ToMap() (map[string]interface{}, error) {
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

type NullableAdyenPaymentUpdateData struct {
	value *AdyenPaymentUpdateData
	isSet bool
}

func (v NullableAdyenPaymentUpdateData) Get() *AdyenPaymentUpdateData {
	return v.value
}

func (v *NullableAdyenPaymentUpdateData) Set(val *AdyenPaymentUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableAdyenPaymentUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableAdyenPaymentUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdyenPaymentUpdateData(val *AdyenPaymentUpdateData) *NullableAdyenPaymentUpdateData {
	return &NullableAdyenPaymentUpdateData{value: val, isSet: true}
}

func (v NullableAdyenPaymentUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdyenPaymentUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
