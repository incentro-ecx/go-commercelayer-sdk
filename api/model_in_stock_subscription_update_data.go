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

// InStockSubscriptionUpdateData struct for InStockSubscriptionUpdateData
type InStockSubscriptionUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                                                  `json:"id"`
	Attributes    PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes `json:"attributes"`
	Relationships *InStockSubscriptionUpdateDataRelationships                             `json:"relationships,omitempty"`
}

// NewInStockSubscriptionUpdateData instantiates a new InStockSubscriptionUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInStockSubscriptionUpdateData(type_ string, id string, attributes PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) *InStockSubscriptionUpdateData {
	this := InStockSubscriptionUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewInStockSubscriptionUpdateDataWithDefaults instantiates a new InStockSubscriptionUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInStockSubscriptionUpdateDataWithDefaults() *InStockSubscriptionUpdateData {
	this := InStockSubscriptionUpdateData{}
	return &this
}

// GetType returns the Type field value
func (o *InStockSubscriptionUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InStockSubscriptionUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *InStockSubscriptionUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InStockSubscriptionUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *InStockSubscriptionUpdateData) GetAttributes() PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionUpdateData) GetAttributesOk() (*PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *InStockSubscriptionUpdateData) SetAttributes(v PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *InStockSubscriptionUpdateData) GetRelationships() InStockSubscriptionUpdateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret InStockSubscriptionUpdateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionUpdateData) GetRelationshipsOk() (*InStockSubscriptionUpdateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *InStockSubscriptionUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given InStockSubscriptionUpdateDataRelationships and assigns it to the Relationships field.
func (o *InStockSubscriptionUpdateData) SetRelationships(v InStockSubscriptionUpdateDataRelationships) {
	o.Relationships = &v
}

func (o InStockSubscriptionUpdateData) MarshalJSON() ([]byte, error) {
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

type NullableInStockSubscriptionUpdateData struct {
	value *InStockSubscriptionUpdateData
	isSet bool
}

func (v NullableInStockSubscriptionUpdateData) Get() *InStockSubscriptionUpdateData {
	return v.value
}

func (v *NullableInStockSubscriptionUpdateData) Set(val *InStockSubscriptionUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableInStockSubscriptionUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableInStockSubscriptionUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInStockSubscriptionUpdateData(val *InStockSubscriptionUpdateData) *NullableInStockSubscriptionUpdateData {
	return &NullableInStockSubscriptionUpdateData{value: val, isSet: true}
}

func (v NullableInStockSubscriptionUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInStockSubscriptionUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
