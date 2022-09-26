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

// InStockSubscriptionResponseData struct for InStockSubscriptionResponseData
type InStockSubscriptionResponseData struct {
	// The resource's id
	Id *string `json:"id,omitempty"`
	// The resource's type
	Type          *string                                       `json:"type,omitempty"`
	Links         *AddressResponseDataLinks                     `json:"links,omitempty"`
	Attributes    *Attributes                                   `json:"attributes,omitempty"`
	Relationships *InStockSubscriptionResponseDataRelationships `json:"relationships,omitempty"`
}

// NewInStockSubscriptionResponseData instantiates a new InStockSubscriptionResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInStockSubscriptionResponseData() *InStockSubscriptionResponseData {
	this := InStockSubscriptionResponseData{}
	return &this
}

// NewInStockSubscriptionResponseDataWithDefaults instantiates a new InStockSubscriptionResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInStockSubscriptionResponseDataWithDefaults() *InStockSubscriptionResponseData {
	this := InStockSubscriptionResponseData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InStockSubscriptionResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InStockSubscriptionResponseData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InStockSubscriptionResponseData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InStockSubscriptionResponseData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionResponseData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InStockSubscriptionResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InStockSubscriptionResponseData) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *InStockSubscriptionResponseData) GetLinks() AddressResponseDataLinks {
	if o == nil || o.Links == nil {
		var ret AddressResponseDataLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionResponseData) GetLinksOk() (*AddressResponseDataLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *InStockSubscriptionResponseData) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AddressResponseDataLinks and assigns it to the Links field.
func (o *InStockSubscriptionResponseData) SetLinks(v AddressResponseDataLinks) {
	o.Links = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *InStockSubscriptionResponseData) GetAttributes() Attributes {
	if o == nil || o.Attributes == nil {
		var ret Attributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionResponseData) GetAttributesOk() (*Attributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *InStockSubscriptionResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given Attributes and assigns it to the Attributes field.
func (o *InStockSubscriptionResponseData) SetAttributes(v Attributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *InStockSubscriptionResponseData) GetRelationships() InStockSubscriptionResponseDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret InStockSubscriptionResponseDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionResponseData) GetRelationshipsOk() (*InStockSubscriptionResponseDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *InStockSubscriptionResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given InStockSubscriptionResponseDataRelationships and assigns it to the Relationships field.
func (o *InStockSubscriptionResponseData) SetRelationships(v InStockSubscriptionResponseDataRelationships) {
	o.Relationships = &v
}

func (o InStockSubscriptionResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableInStockSubscriptionResponseData struct {
	value *InStockSubscriptionResponseData
	isSet bool
}

func (v NullableInStockSubscriptionResponseData) Get() *InStockSubscriptionResponseData {
	return v.value
}

func (v *NullableInStockSubscriptionResponseData) Set(val *InStockSubscriptionResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableInStockSubscriptionResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableInStockSubscriptionResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInStockSubscriptionResponseData(val *InStockSubscriptionResponseData) *NullableInStockSubscriptionResponseData {
	return &NullableInStockSubscriptionResponseData{value: val, isSet: true}
}

func (v NullableInStockSubscriptionResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInStockSubscriptionResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}