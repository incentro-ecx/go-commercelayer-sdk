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

// checks if the SkuListItemDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkuListItemDataRelationships{}

// SkuListItemDataRelationships struct for SkuListItemDataRelationships
type SkuListItemDataRelationships struct {
	SkuList  *BundleDataRelationshipsSkuList   `json:"sku_list,omitempty"`
	Sku      *BundleDataRelationshipsSkus      `json:"sku,omitempty"`
	Versions *AddressDataRelationshipsVersions `json:"versions,omitempty"`
}

// NewSkuListItemDataRelationships instantiates a new SkuListItemDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuListItemDataRelationships() *SkuListItemDataRelationships {
	this := SkuListItemDataRelationships{}
	return &this
}

// NewSkuListItemDataRelationshipsWithDefaults instantiates a new SkuListItemDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuListItemDataRelationshipsWithDefaults() *SkuListItemDataRelationships {
	this := SkuListItemDataRelationships{}
	return &this
}

// GetSkuList returns the SkuList field value if set, zero value otherwise.
func (o *SkuListItemDataRelationships) GetSkuList() BundleDataRelationshipsSkuList {
	if o == nil || IsNil(o.SkuList) {
		var ret BundleDataRelationshipsSkuList
		return ret
	}
	return *o.SkuList
}

// GetSkuListOk returns a tuple with the SkuList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListItemDataRelationships) GetSkuListOk() (*BundleDataRelationshipsSkuList, bool) {
	if o == nil || IsNil(o.SkuList) {
		return nil, false
	}
	return o.SkuList, true
}

// HasSkuList returns a boolean if a field has been set.
func (o *SkuListItemDataRelationships) HasSkuList() bool {
	if o != nil && !IsNil(o.SkuList) {
		return true
	}

	return false
}

// SetSkuList gets a reference to the given BundleDataRelationshipsSkuList and assigns it to the SkuList field.
func (o *SkuListItemDataRelationships) SetSkuList(v BundleDataRelationshipsSkuList) {
	o.SkuList = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *SkuListItemDataRelationships) GetSku() BundleDataRelationshipsSkus {
	if o == nil || IsNil(o.Sku) {
		var ret BundleDataRelationshipsSkus
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListItemDataRelationships) GetSkuOk() (*BundleDataRelationshipsSkus, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *SkuListItemDataRelationships) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given BundleDataRelationshipsSkus and assigns it to the Sku field.
func (o *SkuListItemDataRelationships) SetSku(v BundleDataRelationshipsSkus) {
	o.Sku = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *SkuListItemDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListItemDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *SkuListItemDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *SkuListItemDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

func (o SkuListItemDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkuListItemDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SkuList) {
		toSerialize["sku_list"] = o.SkuList
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullableSkuListItemDataRelationships struct {
	value *SkuListItemDataRelationships
	isSet bool
}

func (v NullableSkuListItemDataRelationships) Get() *SkuListItemDataRelationships {
	return v.value
}

func (v *NullableSkuListItemDataRelationships) Set(val *SkuListItemDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuListItemDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuListItemDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuListItemDataRelationships(val *SkuListItemDataRelationships) *NullableSkuListItemDataRelationships {
	return &NullableSkuListItemDataRelationships{value: val, isSet: true}
}

func (v NullableSkuListItemDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuListItemDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
