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

// CustomerPaymentSourceResponseData struct for CustomerPaymentSourceResponseData
type CustomerPaymentSourceResponseData struct {
	// The resource's id
	Id *string `json:"id,omitempty"`
	// The resource's type
	Type          *string                                         `json:"type,omitempty"`
	Links         *AddressResponseDataLinks                       `json:"links,omitempty"`
	Attributes    *Attributes                                     `json:"attributes,omitempty"`
	Relationships *CustomerPaymentSourceResponseDataRelationships `json:"relationships,omitempty"`
}

// NewCustomerPaymentSourceResponseData instantiates a new CustomerPaymentSourceResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerPaymentSourceResponseData() *CustomerPaymentSourceResponseData {
	this := CustomerPaymentSourceResponseData{}
	return &this
}

// NewCustomerPaymentSourceResponseDataWithDefaults instantiates a new CustomerPaymentSourceResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerPaymentSourceResponseDataWithDefaults() *CustomerPaymentSourceResponseData {
	this := CustomerPaymentSourceResponseData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomerPaymentSourceResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPaymentSourceResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomerPaymentSourceResponseData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomerPaymentSourceResponseData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CustomerPaymentSourceResponseData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPaymentSourceResponseData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CustomerPaymentSourceResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CustomerPaymentSourceResponseData) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CustomerPaymentSourceResponseData) GetLinks() AddressResponseDataLinks {
	if o == nil || o.Links == nil {
		var ret AddressResponseDataLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPaymentSourceResponseData) GetLinksOk() (*AddressResponseDataLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CustomerPaymentSourceResponseData) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AddressResponseDataLinks and assigns it to the Links field.
func (o *CustomerPaymentSourceResponseData) SetLinks(v AddressResponseDataLinks) {
	o.Links = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CustomerPaymentSourceResponseData) GetAttributes() Attributes {
	if o == nil || o.Attributes == nil {
		var ret Attributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPaymentSourceResponseData) GetAttributesOk() (*Attributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CustomerPaymentSourceResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given Attributes and assigns it to the Attributes field.
func (o *CustomerPaymentSourceResponseData) SetAttributes(v Attributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CustomerPaymentSourceResponseData) GetRelationships() CustomerPaymentSourceResponseDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret CustomerPaymentSourceResponseDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPaymentSourceResponseData) GetRelationshipsOk() (*CustomerPaymentSourceResponseDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CustomerPaymentSourceResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CustomerPaymentSourceResponseDataRelationships and assigns it to the Relationships field.
func (o *CustomerPaymentSourceResponseData) SetRelationships(v CustomerPaymentSourceResponseDataRelationships) {
	o.Relationships = &v
}

func (o CustomerPaymentSourceResponseData) MarshalJSON() ([]byte, error) {
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

type NullableCustomerPaymentSourceResponseData struct {
	value *CustomerPaymentSourceResponseData
	isSet bool
}

func (v NullableCustomerPaymentSourceResponseData) Get() *CustomerPaymentSourceResponseData {
	return v.value
}

func (v *NullableCustomerPaymentSourceResponseData) Set(val *CustomerPaymentSourceResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPaymentSourceResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPaymentSourceResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPaymentSourceResponseData(val *CustomerPaymentSourceResponseData) *NullableCustomerPaymentSourceResponseData {
	return &NullableCustomerPaymentSourceResponseData{value: val, isSet: true}
}

func (v NullableCustomerPaymentSourceResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPaymentSourceResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}