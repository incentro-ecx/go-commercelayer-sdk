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

// checks if the POSTShippingMethods201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTShippingMethods201ResponseDataRelationships{}

// POSTShippingMethods201ResponseDataRelationships struct for POSTShippingMethods201ResponseDataRelationships
type POSTShippingMethods201ResponseDataRelationships struct {
	Market                      *POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket           `json:"market,omitempty"`
	ShippingZone                *POSTShippingMethods201ResponseDataRelationshipsShippingZone                `json:"shipping_zone,omitempty"`
	ShippingCategory            *GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory         `json:"shipping_category,omitempty"`
	StockLocation               *POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation             `json:"stock_location,omitempty"`
	DeliveryLeadTimeForShipment *POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment `json:"delivery_lead_time_for_shipment,omitempty"`
	ShippingMethodTiers         *POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers         `json:"shipping_method_tiers,omitempty"`
	ShippingWeightTiers         *POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers         `json:"shipping_weight_tiers,omitempty"`
	Attachments                 *POSTAvalaraAccounts201ResponseDataRelationshipsAttachments                 `json:"attachments,omitempty"`
}

// NewPOSTShippingMethods201ResponseDataRelationships instantiates a new POSTShippingMethods201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTShippingMethods201ResponseDataRelationships() *POSTShippingMethods201ResponseDataRelationships {
	this := POSTShippingMethods201ResponseDataRelationships{}
	return &this
}

// NewPOSTShippingMethods201ResponseDataRelationshipsWithDefaults instantiates a new POSTShippingMethods201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTShippingMethods201ResponseDataRelationshipsWithDefaults() *POSTShippingMethods201ResponseDataRelationships {
	this := POSTShippingMethods201ResponseDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *POSTShippingMethods201ResponseDataRelationships) GetMarket() POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket {
	if o == nil || IsNil(o.Market) {
		var ret POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingMethods201ResponseDataRelationships) GetMarketOk() (*POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket, bool) {
	if o == nil || IsNil(o.Market) {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *POSTShippingMethods201ResponseDataRelationships) HasMarket() bool {
	if o != nil && !IsNil(o.Market) {
		return true
	}

	return false
}

// SetMarket gets a reference to the given POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket and assigns it to the Market field.
func (o *POSTShippingMethods201ResponseDataRelationships) SetMarket(v POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket) {
	o.Market = &v
}

// GetShippingZone returns the ShippingZone field value if set, zero value otherwise.
func (o *POSTShippingMethods201ResponseDataRelationships) GetShippingZone() POSTShippingMethods201ResponseDataRelationshipsShippingZone {
	if o == nil || IsNil(o.ShippingZone) {
		var ret POSTShippingMethods201ResponseDataRelationshipsShippingZone
		return ret
	}
	return *o.ShippingZone
}

// GetShippingZoneOk returns a tuple with the ShippingZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingMethods201ResponseDataRelationships) GetShippingZoneOk() (*POSTShippingMethods201ResponseDataRelationshipsShippingZone, bool) {
	if o == nil || IsNil(o.ShippingZone) {
		return nil, false
	}
	return o.ShippingZone, true
}

// HasShippingZone returns a boolean if a field has been set.
func (o *POSTShippingMethods201ResponseDataRelationships) HasShippingZone() bool {
	if o != nil && !IsNil(o.ShippingZone) {
		return true
	}

	return false
}

// SetShippingZone gets a reference to the given POSTShippingMethods201ResponseDataRelationshipsShippingZone and assigns it to the ShippingZone field.
func (o *POSTShippingMethods201ResponseDataRelationships) SetShippingZone(v POSTShippingMethods201ResponseDataRelationshipsShippingZone) {
	o.ShippingZone = &v
}

// GetShippingCategory returns the ShippingCategory field value if set, zero value otherwise.
func (o *POSTShippingMethods201ResponseDataRelationships) GetShippingCategory() GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory {
	if o == nil || IsNil(o.ShippingCategory) {
		var ret GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory
		return ret
	}
	return *o.ShippingCategory
}

// GetShippingCategoryOk returns a tuple with the ShippingCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingMethods201ResponseDataRelationships) GetShippingCategoryOk() (*GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory, bool) {
	if o == nil || IsNil(o.ShippingCategory) {
		return nil, false
	}
	return o.ShippingCategory, true
}

// HasShippingCategory returns a boolean if a field has been set.
func (o *POSTShippingMethods201ResponseDataRelationships) HasShippingCategory() bool {
	if o != nil && !IsNil(o.ShippingCategory) {
		return true
	}

	return false
}

// SetShippingCategory gets a reference to the given GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory and assigns it to the ShippingCategory field.
func (o *POSTShippingMethods201ResponseDataRelationships) SetShippingCategory(v GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory) {
	o.ShippingCategory = &v
}

// GetStockLocation returns the StockLocation field value if set, zero value otherwise.
func (o *POSTShippingMethods201ResponseDataRelationships) GetStockLocation() POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation {
	if o == nil || IsNil(o.StockLocation) {
		var ret POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation
		return ret
	}
	return *o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingMethods201ResponseDataRelationships) GetStockLocationOk() (*POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation, bool) {
	if o == nil || IsNil(o.StockLocation) {
		return nil, false
	}
	return o.StockLocation, true
}

// HasStockLocation returns a boolean if a field has been set.
func (o *POSTShippingMethods201ResponseDataRelationships) HasStockLocation() bool {
	if o != nil && !IsNil(o.StockLocation) {
		return true
	}

	return false
}

// SetStockLocation gets a reference to the given POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation and assigns it to the StockLocation field.
func (o *POSTShippingMethods201ResponseDataRelationships) SetStockLocation(v POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation) {
	o.StockLocation = &v
}

// GetDeliveryLeadTimeForShipment returns the DeliveryLeadTimeForShipment field value if set, zero value otherwise.
func (o *POSTShippingMethods201ResponseDataRelationships) GetDeliveryLeadTimeForShipment() POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment {
	if o == nil || IsNil(o.DeliveryLeadTimeForShipment) {
		var ret POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment
		return ret
	}
	return *o.DeliveryLeadTimeForShipment
}

// GetDeliveryLeadTimeForShipmentOk returns a tuple with the DeliveryLeadTimeForShipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingMethods201ResponseDataRelationships) GetDeliveryLeadTimeForShipmentOk() (*POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment, bool) {
	if o == nil || IsNil(o.DeliveryLeadTimeForShipment) {
		return nil, false
	}
	return o.DeliveryLeadTimeForShipment, true
}

// HasDeliveryLeadTimeForShipment returns a boolean if a field has been set.
func (o *POSTShippingMethods201ResponseDataRelationships) HasDeliveryLeadTimeForShipment() bool {
	if o != nil && !IsNil(o.DeliveryLeadTimeForShipment) {
		return true
	}

	return false
}

// SetDeliveryLeadTimeForShipment gets a reference to the given POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment and assigns it to the DeliveryLeadTimeForShipment field.
func (o *POSTShippingMethods201ResponseDataRelationships) SetDeliveryLeadTimeForShipment(v POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment) {
	o.DeliveryLeadTimeForShipment = &v
}

// GetShippingMethodTiers returns the ShippingMethodTiers field value if set, zero value otherwise.
func (o *POSTShippingMethods201ResponseDataRelationships) GetShippingMethodTiers() POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers {
	if o == nil || IsNil(o.ShippingMethodTiers) {
		var ret POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers
		return ret
	}
	return *o.ShippingMethodTiers
}

// GetShippingMethodTiersOk returns a tuple with the ShippingMethodTiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingMethods201ResponseDataRelationships) GetShippingMethodTiersOk() (*POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers, bool) {
	if o == nil || IsNil(o.ShippingMethodTiers) {
		return nil, false
	}
	return o.ShippingMethodTiers, true
}

// HasShippingMethodTiers returns a boolean if a field has been set.
func (o *POSTShippingMethods201ResponseDataRelationships) HasShippingMethodTiers() bool {
	if o != nil && !IsNil(o.ShippingMethodTiers) {
		return true
	}

	return false
}

// SetShippingMethodTiers gets a reference to the given POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers and assigns it to the ShippingMethodTiers field.
func (o *POSTShippingMethods201ResponseDataRelationships) SetShippingMethodTiers(v POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers) {
	o.ShippingMethodTiers = &v
}

// GetShippingWeightTiers returns the ShippingWeightTiers field value if set, zero value otherwise.
func (o *POSTShippingMethods201ResponseDataRelationships) GetShippingWeightTiers() POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers {
	if o == nil || IsNil(o.ShippingWeightTiers) {
		var ret POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers
		return ret
	}
	return *o.ShippingWeightTiers
}

// GetShippingWeightTiersOk returns a tuple with the ShippingWeightTiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingMethods201ResponseDataRelationships) GetShippingWeightTiersOk() (*POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers, bool) {
	if o == nil || IsNil(o.ShippingWeightTiers) {
		return nil, false
	}
	return o.ShippingWeightTiers, true
}

// HasShippingWeightTiers returns a boolean if a field has been set.
func (o *POSTShippingMethods201ResponseDataRelationships) HasShippingWeightTiers() bool {
	if o != nil && !IsNil(o.ShippingWeightTiers) {
		return true
	}

	return false
}

// SetShippingWeightTiers gets a reference to the given POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers and assigns it to the ShippingWeightTiers field.
func (o *POSTShippingMethods201ResponseDataRelationships) SetShippingWeightTiers(v POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers) {
	o.ShippingWeightTiers = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *POSTShippingMethods201ResponseDataRelationships) GetAttachments() POSTAvalaraAccounts201ResponseDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret POSTAvalaraAccounts201ResponseDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingMethods201ResponseDataRelationships) GetAttachmentsOk() (*POSTAvalaraAccounts201ResponseDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *POSTShippingMethods201ResponseDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given POSTAvalaraAccounts201ResponseDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *POSTShippingMethods201ResponseDataRelationships) SetAttachments(v POSTAvalaraAccounts201ResponseDataRelationshipsAttachments) {
	o.Attachments = &v
}

func (o POSTShippingMethods201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTShippingMethods201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Market) {
		toSerialize["market"] = o.Market
	}
	if !IsNil(o.ShippingZone) {
		toSerialize["shipping_zone"] = o.ShippingZone
	}
	if !IsNil(o.ShippingCategory) {
		toSerialize["shipping_category"] = o.ShippingCategory
	}
	if !IsNil(o.StockLocation) {
		toSerialize["stock_location"] = o.StockLocation
	}
	if !IsNil(o.DeliveryLeadTimeForShipment) {
		toSerialize["delivery_lead_time_for_shipment"] = o.DeliveryLeadTimeForShipment
	}
	if !IsNil(o.ShippingMethodTiers) {
		toSerialize["shipping_method_tiers"] = o.ShippingMethodTiers
	}
	if !IsNil(o.ShippingWeightTiers) {
		toSerialize["shipping_weight_tiers"] = o.ShippingWeightTiers
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	return toSerialize, nil
}

type NullablePOSTShippingMethods201ResponseDataRelationships struct {
	value *POSTShippingMethods201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTShippingMethods201ResponseDataRelationships) Get() *POSTShippingMethods201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTShippingMethods201ResponseDataRelationships) Set(val *POSTShippingMethods201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTShippingMethods201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTShippingMethods201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTShippingMethods201ResponseDataRelationships(val *POSTShippingMethods201ResponseDataRelationships) *NullablePOSTShippingMethods201ResponseDataRelationships {
	return &NullablePOSTShippingMethods201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTShippingMethods201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTShippingMethods201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
