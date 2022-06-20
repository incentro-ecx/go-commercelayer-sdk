/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ReturnDataRelationships struct for ReturnDataRelationships
type ReturnDataRelationships struct {
	Order              *AdyenPaymentDataRelationshipsOrder             `json:"order,omitempty"`
	Customer           *CouponRecipientDataRelationshipsCustomer       `json:"customer,omitempty"`
	StockLocation      *DeliveryLeadTimeDataRelationshipsStockLocation `json:"stock_location,omitempty"`
	OriginAddress      *BingGeocoderDataRelationshipsAddresses         `json:"origin_address,omitempty"`
	DestinationAddress *BingGeocoderDataRelationshipsAddresses         `json:"destination_address,omitempty"`
	ReturnLineItems    *ReturnDataRelationshipsReturnLineItems         `json:"return_line_items,omitempty"`
	Attachments        *AvalaraAccountDataRelationshipsAttachments     `json:"attachments,omitempty"`
}

// NewReturnDataRelationships instantiates a new ReturnDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnDataRelationships() *ReturnDataRelationships {
	this := ReturnDataRelationships{}
	return &this
}

// NewReturnDataRelationshipsWithDefaults instantiates a new ReturnDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnDataRelationshipsWithDefaults() *ReturnDataRelationships {
	this := ReturnDataRelationships{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ReturnDataRelationships) GetOrder() AdyenPaymentDataRelationshipsOrder {
	if o == nil || o.Order == nil {
		var ret AdyenPaymentDataRelationshipsOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataRelationships) GetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ReturnDataRelationships) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given AdyenPaymentDataRelationshipsOrder and assigns it to the Order field.
func (o *ReturnDataRelationships) SetOrder(v AdyenPaymentDataRelationshipsOrder) {
	o.Order = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *ReturnDataRelationships) GetCustomer() CouponRecipientDataRelationshipsCustomer {
	if o == nil || o.Customer == nil {
		var ret CouponRecipientDataRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataRelationships) GetCustomerOk() (*CouponRecipientDataRelationshipsCustomer, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *ReturnDataRelationships) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given CouponRecipientDataRelationshipsCustomer and assigns it to the Customer field.
func (o *ReturnDataRelationships) SetCustomer(v CouponRecipientDataRelationshipsCustomer) {
	o.Customer = &v
}

// GetStockLocation returns the StockLocation field value if set, zero value otherwise.
func (o *ReturnDataRelationships) GetStockLocation() DeliveryLeadTimeDataRelationshipsStockLocation {
	if o == nil || o.StockLocation == nil {
		var ret DeliveryLeadTimeDataRelationshipsStockLocation
		return ret
	}
	return *o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeDataRelationshipsStockLocation, bool) {
	if o == nil || o.StockLocation == nil {
		return nil, false
	}
	return o.StockLocation, true
}

// HasStockLocation returns a boolean if a field has been set.
func (o *ReturnDataRelationships) HasStockLocation() bool {
	if o != nil && o.StockLocation != nil {
		return true
	}

	return false
}

// SetStockLocation gets a reference to the given DeliveryLeadTimeDataRelationshipsStockLocation and assigns it to the StockLocation field.
func (o *ReturnDataRelationships) SetStockLocation(v DeliveryLeadTimeDataRelationshipsStockLocation) {
	o.StockLocation = &v
}

// GetOriginAddress returns the OriginAddress field value if set, zero value otherwise.
func (o *ReturnDataRelationships) GetOriginAddress() BingGeocoderDataRelationshipsAddresses {
	if o == nil || o.OriginAddress == nil {
		var ret BingGeocoderDataRelationshipsAddresses
		return ret
	}
	return *o.OriginAddress
}

// GetOriginAddressOk returns a tuple with the OriginAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataRelationships) GetOriginAddressOk() (*BingGeocoderDataRelationshipsAddresses, bool) {
	if o == nil || o.OriginAddress == nil {
		return nil, false
	}
	return o.OriginAddress, true
}

// HasOriginAddress returns a boolean if a field has been set.
func (o *ReturnDataRelationships) HasOriginAddress() bool {
	if o != nil && o.OriginAddress != nil {
		return true
	}

	return false
}

// SetOriginAddress gets a reference to the given BingGeocoderDataRelationshipsAddresses and assigns it to the OriginAddress field.
func (o *ReturnDataRelationships) SetOriginAddress(v BingGeocoderDataRelationshipsAddresses) {
	o.OriginAddress = &v
}

// GetDestinationAddress returns the DestinationAddress field value if set, zero value otherwise.
func (o *ReturnDataRelationships) GetDestinationAddress() BingGeocoderDataRelationshipsAddresses {
	if o == nil || o.DestinationAddress == nil {
		var ret BingGeocoderDataRelationshipsAddresses
		return ret
	}
	return *o.DestinationAddress
}

// GetDestinationAddressOk returns a tuple with the DestinationAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataRelationships) GetDestinationAddressOk() (*BingGeocoderDataRelationshipsAddresses, bool) {
	if o == nil || o.DestinationAddress == nil {
		return nil, false
	}
	return o.DestinationAddress, true
}

// HasDestinationAddress returns a boolean if a field has been set.
func (o *ReturnDataRelationships) HasDestinationAddress() bool {
	if o != nil && o.DestinationAddress != nil {
		return true
	}

	return false
}

// SetDestinationAddress gets a reference to the given BingGeocoderDataRelationshipsAddresses and assigns it to the DestinationAddress field.
func (o *ReturnDataRelationships) SetDestinationAddress(v BingGeocoderDataRelationshipsAddresses) {
	o.DestinationAddress = &v
}

// GetReturnLineItems returns the ReturnLineItems field value if set, zero value otherwise.
func (o *ReturnDataRelationships) GetReturnLineItems() ReturnDataRelationshipsReturnLineItems {
	if o == nil || o.ReturnLineItems == nil {
		var ret ReturnDataRelationshipsReturnLineItems
		return ret
	}
	return *o.ReturnLineItems
}

// GetReturnLineItemsOk returns a tuple with the ReturnLineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataRelationships) GetReturnLineItemsOk() (*ReturnDataRelationshipsReturnLineItems, bool) {
	if o == nil || o.ReturnLineItems == nil {
		return nil, false
	}
	return o.ReturnLineItems, true
}

// HasReturnLineItems returns a boolean if a field has been set.
func (o *ReturnDataRelationships) HasReturnLineItems() bool {
	if o != nil && o.ReturnLineItems != nil {
		return true
	}

	return false
}

// SetReturnLineItems gets a reference to the given ReturnDataRelationshipsReturnLineItems and assigns it to the ReturnLineItems field.
func (o *ReturnDataRelationships) SetReturnLineItems(v ReturnDataRelationshipsReturnLineItems) {
	o.ReturnLineItems = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *ReturnDataRelationships) GetAttachments() AvalaraAccountDataRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret AvalaraAccountDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataRelationships) GetAttachmentsOk() (*AvalaraAccountDataRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *ReturnDataRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AvalaraAccountDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *ReturnDataRelationships) SetAttachments(v AvalaraAccountDataRelationshipsAttachments) {
	o.Attachments = &v
}

func (o ReturnDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
	}
	if o.StockLocation != nil {
		toSerialize["stock_location"] = o.StockLocation
	}
	if o.OriginAddress != nil {
		toSerialize["origin_address"] = o.OriginAddress
	}
	if o.DestinationAddress != nil {
		toSerialize["destination_address"] = o.DestinationAddress
	}
	if o.ReturnLineItems != nil {
		toSerialize["return_line_items"] = o.ReturnLineItems
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullableReturnDataRelationships struct {
	value *ReturnDataRelationships
	isSet bool
}

func (v NullableReturnDataRelationships) Get() *ReturnDataRelationships {
	return v.value
}

func (v *NullableReturnDataRelationships) Set(val *ReturnDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnDataRelationships(val *ReturnDataRelationships) *NullableReturnDataRelationships {
	return &NullableReturnDataRelationships{value: val, isSet: true}
}

func (v NullableReturnDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}