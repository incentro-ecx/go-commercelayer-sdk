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

// checks if the ExternalPromotionUpdateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalPromotionUpdateData{}

// ExternalPromotionUpdateData struct for ExternalPromotionUpdateData
type ExternalPromotionUpdateData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                                         `json:"id"`
	Attributes    PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes `json:"attributes"`
	Relationships *BuyXPayYPromotionUpdateDataRelationships                           `json:"relationships,omitempty"`
}

// NewExternalPromotionUpdateData instantiates a new ExternalPromotionUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalPromotionUpdateData(type_ interface{}, id interface{}, attributes PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) *ExternalPromotionUpdateData {
	this := ExternalPromotionUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewExternalPromotionUpdateDataWithDefaults instantiates a new ExternalPromotionUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalPromotionUpdateDataWithDefaults() *ExternalPromotionUpdateData {
	this := ExternalPromotionUpdateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExternalPromotionUpdateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalPromotionUpdateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExternalPromotionUpdateData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExternalPromotionUpdateData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalPromotionUpdateData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExternalPromotionUpdateData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *ExternalPromotionUpdateData) GetAttributes() PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ExternalPromotionUpdateData) GetAttributesOk() (*PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ExternalPromotionUpdateData) SetAttributes(v PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ExternalPromotionUpdateData) GetRelationships() BuyXPayYPromotionUpdateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret BuyXPayYPromotionUpdateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPromotionUpdateData) GetRelationshipsOk() (*BuyXPayYPromotionUpdateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ExternalPromotionUpdateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given BuyXPayYPromotionUpdateDataRelationships and assigns it to the Relationships field.
func (o *ExternalPromotionUpdateData) SetRelationships(v BuyXPayYPromotionUpdateDataRelationships) {
	o.Relationships = &v
}

func (o ExternalPromotionUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalPromotionUpdateData) ToMap() (map[string]interface{}, error) {
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

type NullableExternalPromotionUpdateData struct {
	value *ExternalPromotionUpdateData
	isSet bool
}

func (v NullableExternalPromotionUpdateData) Get() *ExternalPromotionUpdateData {
	return v.value
}

func (v *NullableExternalPromotionUpdateData) Set(val *ExternalPromotionUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalPromotionUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalPromotionUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalPromotionUpdateData(val *ExternalPromotionUpdateData) *NullableExternalPromotionUpdateData {
	return &NullableExternalPromotionUpdateData{value: val, isSet: true}
}

func (v NullableExternalPromotionUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalPromotionUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
