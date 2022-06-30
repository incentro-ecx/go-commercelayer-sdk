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

// LineItemUpdateDataAttributes struct for LineItemUpdateDataAttributes
type LineItemUpdateDataAttributes struct {
	// The code of the associated sku.
	SkuCode *string `json:"sku_code,omitempty"`
	// The code of the associated bundle.
	BundleCode *string `json:"bundle_code,omitempty"`
	// The line item quantity.
	Quantity *int32 `json:"quantity,omitempty"`
	// The name of the line item. When blank, it gets populated with the name of the associated item (if present).
	Name *string `json:"name,omitempty"`
	// The image_url of the line item. When blank, it gets populated with the image_url of the associated item (if present, sku only).
	ImageUrl *string `json:"image_url,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewLineItemUpdateDataAttributes instantiates a new LineItemUpdateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemUpdateDataAttributes() *LineItemUpdateDataAttributes {
	this := LineItemUpdateDataAttributes{}
	return &this
}

// NewLineItemUpdateDataAttributesWithDefaults instantiates a new LineItemUpdateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemUpdateDataAttributesWithDefaults() *LineItemUpdateDataAttributes {
	this := LineItemUpdateDataAttributes{}
	return &this
}

// GetSkuCode returns the SkuCode field value if set, zero value otherwise.
func (o *LineItemUpdateDataAttributes) GetSkuCode() string {
	if o == nil || o.SkuCode == nil {
		var ret string
		return ret
	}
	return *o.SkuCode
}

// GetSkuCodeOk returns a tuple with the SkuCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemUpdateDataAttributes) GetSkuCodeOk() (*string, bool) {
	if o == nil || o.SkuCode == nil {
		return nil, false
	}
	return o.SkuCode, true
}

// HasSkuCode returns a boolean if a field has been set.
func (o *LineItemUpdateDataAttributes) HasSkuCode() bool {
	if o != nil && o.SkuCode != nil {
		return true
	}

	return false
}

// SetSkuCode gets a reference to the given string and assigns it to the SkuCode field.
func (o *LineItemUpdateDataAttributes) SetSkuCode(v string) {
	o.SkuCode = &v
}

// GetBundleCode returns the BundleCode field value if set, zero value otherwise.
func (o *LineItemUpdateDataAttributes) GetBundleCode() string {
	if o == nil || o.BundleCode == nil {
		var ret string
		return ret
	}
	return *o.BundleCode
}

// GetBundleCodeOk returns a tuple with the BundleCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemUpdateDataAttributes) GetBundleCodeOk() (*string, bool) {
	if o == nil || o.BundleCode == nil {
		return nil, false
	}
	return o.BundleCode, true
}

// HasBundleCode returns a boolean if a field has been set.
func (o *LineItemUpdateDataAttributes) HasBundleCode() bool {
	if o != nil && o.BundleCode != nil {
		return true
	}

	return false
}

// SetBundleCode gets a reference to the given string and assigns it to the BundleCode field.
func (o *LineItemUpdateDataAttributes) SetBundleCode(v string) {
	o.BundleCode = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *LineItemUpdateDataAttributes) GetQuantity() int32 {
	if o == nil || o.Quantity == nil {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemUpdateDataAttributes) GetQuantityOk() (*int32, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *LineItemUpdateDataAttributes) HasQuantity() bool {
	if o != nil && o.Quantity != nil {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *LineItemUpdateDataAttributes) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LineItemUpdateDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemUpdateDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LineItemUpdateDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LineItemUpdateDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *LineItemUpdateDataAttributes) GetImageUrl() string {
	if o == nil || o.ImageUrl == nil {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemUpdateDataAttributes) GetImageUrlOk() (*string, bool) {
	if o == nil || o.ImageUrl == nil {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *LineItemUpdateDataAttributes) HasImageUrl() bool {
	if o != nil && o.ImageUrl != nil {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *LineItemUpdateDataAttributes) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *LineItemUpdateDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemUpdateDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *LineItemUpdateDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *LineItemUpdateDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *LineItemUpdateDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemUpdateDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *LineItemUpdateDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *LineItemUpdateDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *LineItemUpdateDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemUpdateDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *LineItemUpdateDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *LineItemUpdateDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o LineItemUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SkuCode != nil {
		toSerialize["sku_code"] = o.SkuCode
	}
	if o.BundleCode != nil {
		toSerialize["bundle_code"] = o.BundleCode
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ImageUrl != nil {
		toSerialize["image_url"] = o.ImageUrl
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

type NullableLineItemUpdateDataAttributes struct {
	value *LineItemUpdateDataAttributes
	isSet bool
}

func (v NullableLineItemUpdateDataAttributes) Get() *LineItemUpdateDataAttributes {
	return v.value
}

func (v *NullableLineItemUpdateDataAttributes) Set(val *LineItemUpdateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemUpdateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemUpdateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemUpdateDataAttributes(val *LineItemUpdateDataAttributes) *NullableLineItemUpdateDataAttributes {
	return &NullableLineItemUpdateDataAttributes{value: val, isSet: true}
}

func (v NullableLineItemUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemUpdateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
