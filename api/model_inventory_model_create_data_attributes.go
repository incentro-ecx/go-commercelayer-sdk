/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.5
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// InventoryModelCreateDataAttributes struct for InventoryModelCreateDataAttributes
type InventoryModelCreateDataAttributes struct {
	// The inventory model's internal name.
	Name string `json:"name"`
	// The inventory model's shipping strategy: one between 'split_shipments' (default), 'ship_from_primary' and 'ship_from_first_available_or_primary'.
	Strategy string `json:"strategy"`
	// The maximum number of stock locations used for inventory computation
	StockLocationsCutoff *int32 `json:"stock_locations_cutoff,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewInventoryModelCreateDataAttributes instantiates a new InventoryModelCreateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryModelCreateDataAttributes(name string, strategy string) *InventoryModelCreateDataAttributes {
	this := InventoryModelCreateDataAttributes{}
	this.Name = name
	this.Strategy = strategy
	return &this
}

// NewInventoryModelCreateDataAttributesWithDefaults instantiates a new InventoryModelCreateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryModelCreateDataAttributesWithDefaults() *InventoryModelCreateDataAttributes {
	this := InventoryModelCreateDataAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *InventoryModelCreateDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InventoryModelCreateDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InventoryModelCreateDataAttributes) SetName(v string) {
	o.Name = v
}

// GetStrategy returns the Strategy field value
func (o *InventoryModelCreateDataAttributes) GetStrategy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value
// and a boolean to check if the value has been set.
func (o *InventoryModelCreateDataAttributes) GetStrategyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Strategy, true
}

// SetStrategy sets field value
func (o *InventoryModelCreateDataAttributes) SetStrategy(v string) {
	o.Strategy = v
}

// GetStockLocationsCutoff returns the StockLocationsCutoff field value if set, zero value otherwise.
func (o *InventoryModelCreateDataAttributes) GetStockLocationsCutoff() int32 {
	if o == nil || o.StockLocationsCutoff == nil {
		var ret int32
		return ret
	}
	return *o.StockLocationsCutoff
}

// GetStockLocationsCutoffOk returns a tuple with the StockLocationsCutoff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryModelCreateDataAttributes) GetStockLocationsCutoffOk() (*int32, bool) {
	if o == nil || o.StockLocationsCutoff == nil {
		return nil, false
	}
	return o.StockLocationsCutoff, true
}

// HasStockLocationsCutoff returns a boolean if a field has been set.
func (o *InventoryModelCreateDataAttributes) HasStockLocationsCutoff() bool {
	if o != nil && o.StockLocationsCutoff != nil {
		return true
	}

	return false
}

// SetStockLocationsCutoff gets a reference to the given int32 and assigns it to the StockLocationsCutoff field.
func (o *InventoryModelCreateDataAttributes) SetStockLocationsCutoff(v int32) {
	o.StockLocationsCutoff = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *InventoryModelCreateDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryModelCreateDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *InventoryModelCreateDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *InventoryModelCreateDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *InventoryModelCreateDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryModelCreateDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *InventoryModelCreateDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *InventoryModelCreateDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *InventoryModelCreateDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryModelCreateDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *InventoryModelCreateDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *InventoryModelCreateDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o InventoryModelCreateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["strategy"] = o.Strategy
	}
	if o.StockLocationsCutoff != nil {
		toSerialize["stock_locations_cutoff"] = o.StockLocationsCutoff
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.ReferenceOrigin != nil {
		toSerialize["reference_origin"] = o.ReferenceOrigin
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryModelCreateDataAttributes struct {
	value *InventoryModelCreateDataAttributes
	isSet bool
}

func (v NullableInventoryModelCreateDataAttributes) Get() *InventoryModelCreateDataAttributes {
	return v.value
}

func (v *NullableInventoryModelCreateDataAttributes) Set(val *InventoryModelCreateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryModelCreateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryModelCreateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryModelCreateDataAttributes(val *InventoryModelCreateDataAttributes) *NullableInventoryModelCreateDataAttributes {
	return &NullableInventoryModelCreateDataAttributes{value: val, isSet: true}
}

func (v NullableInventoryModelCreateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryModelCreateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
