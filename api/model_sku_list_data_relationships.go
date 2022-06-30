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

// SkuListDataRelationships struct for SkuListDataRelationships
type SkuListDataRelationships struct {
	Skus         *BundleDataRelationshipsSkus                `json:"skus,omitempty"`
	SkuListItems *SkuListDataRelationshipsSkuListItems       `json:"sku_list_items,omitempty"`
	Bundles      *OrderDataRelationshipsAvailableFreeBundles `json:"bundles,omitempty"`
}

// NewSkuListDataRelationships instantiates a new SkuListDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuListDataRelationships() *SkuListDataRelationships {
	this := SkuListDataRelationships{}
	return &this
}

// NewSkuListDataRelationshipsWithDefaults instantiates a new SkuListDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuListDataRelationshipsWithDefaults() *SkuListDataRelationships {
	this := SkuListDataRelationships{}
	return &this
}

// GetSkus returns the Skus field value if set, zero value otherwise.
func (o *SkuListDataRelationships) GetSkus() BundleDataRelationshipsSkus {
	if o == nil || o.Skus == nil {
		var ret BundleDataRelationshipsSkus
		return ret
	}
	return *o.Skus
}

// GetSkusOk returns a tuple with the Skus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListDataRelationships) GetSkusOk() (*BundleDataRelationshipsSkus, bool) {
	if o == nil || o.Skus == nil {
		return nil, false
	}
	return o.Skus, true
}

// HasSkus returns a boolean if a field has been set.
func (o *SkuListDataRelationships) HasSkus() bool {
	if o != nil && o.Skus != nil {
		return true
	}

	return false
}

// SetSkus gets a reference to the given BundleDataRelationshipsSkus and assigns it to the Skus field.
func (o *SkuListDataRelationships) SetSkus(v BundleDataRelationshipsSkus) {
	o.Skus = &v
}

// GetSkuListItems returns the SkuListItems field value if set, zero value otherwise.
func (o *SkuListDataRelationships) GetSkuListItems() SkuListDataRelationshipsSkuListItems {
	if o == nil || o.SkuListItems == nil {
		var ret SkuListDataRelationshipsSkuListItems
		return ret
	}
	return *o.SkuListItems
}

// GetSkuListItemsOk returns a tuple with the SkuListItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListDataRelationships) GetSkuListItemsOk() (*SkuListDataRelationshipsSkuListItems, bool) {
	if o == nil || o.SkuListItems == nil {
		return nil, false
	}
	return o.SkuListItems, true
}

// HasSkuListItems returns a boolean if a field has been set.
func (o *SkuListDataRelationships) HasSkuListItems() bool {
	if o != nil && o.SkuListItems != nil {
		return true
	}

	return false
}

// SetSkuListItems gets a reference to the given SkuListDataRelationshipsSkuListItems and assigns it to the SkuListItems field.
func (o *SkuListDataRelationships) SetSkuListItems(v SkuListDataRelationshipsSkuListItems) {
	o.SkuListItems = &v
}

// GetBundles returns the Bundles field value if set, zero value otherwise.
func (o *SkuListDataRelationships) GetBundles() OrderDataRelationshipsAvailableFreeBundles {
	if o == nil || o.Bundles == nil {
		var ret OrderDataRelationshipsAvailableFreeBundles
		return ret
	}
	return *o.Bundles
}

// GetBundlesOk returns a tuple with the Bundles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListDataRelationships) GetBundlesOk() (*OrderDataRelationshipsAvailableFreeBundles, bool) {
	if o == nil || o.Bundles == nil {
		return nil, false
	}
	return o.Bundles, true
}

// HasBundles returns a boolean if a field has been set.
func (o *SkuListDataRelationships) HasBundles() bool {
	if o != nil && o.Bundles != nil {
		return true
	}

	return false
}

// SetBundles gets a reference to the given OrderDataRelationshipsAvailableFreeBundles and assigns it to the Bundles field.
func (o *SkuListDataRelationships) SetBundles(v OrderDataRelationshipsAvailableFreeBundles) {
	o.Bundles = &v
}

func (o SkuListDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Skus != nil {
		toSerialize["skus"] = o.Skus
	}
	if o.SkuListItems != nil {
		toSerialize["sku_list_items"] = o.SkuListItems
	}
	if o.Bundles != nil {
		toSerialize["bundles"] = o.Bundles
	}
	return json.Marshal(toSerialize)
}

type NullableSkuListDataRelationships struct {
	value *SkuListDataRelationships
	isSet bool
}

func (v NullableSkuListDataRelationships) Get() *SkuListDataRelationships {
	return v.value
}

func (v *NullableSkuListDataRelationships) Set(val *SkuListDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuListDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuListDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuListDataRelationships(val *SkuListDataRelationships) *NullableSkuListDataRelationships {
	return &NullableSkuListDataRelationships{value: val, isSet: true}
}

func (v NullableSkuListDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuListDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
