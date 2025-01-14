# POSTShippingMethods201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**POSTBundles201ResponseDataRelationshipsMarket**](POSTBundles201ResponseDataRelationshipsMarket.md) |  | [optional] 
**ShippingZone** | Pointer to [**POSTShippingMethods201ResponseDataRelationshipsShippingZone**](POSTShippingMethods201ResponseDataRelationshipsShippingZone.md) |  | [optional] 
**ShippingCategory** | Pointer to [**POSTShipments201ResponseDataRelationshipsShippingCategory**](POSTShipments201ResponseDataRelationshipsShippingCategory.md) |  | [optional] 
**StockLocation** | Pointer to [**POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation**](POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation.md) |  | [optional] 
**DeliveryLeadTimeForShipment** | Pointer to [**POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment**](POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment.md) |  | [optional] 
**ShippingMethodTiers** | Pointer to [**POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers**](POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers.md) |  | [optional] 
**ShippingWeightTiers** | Pointer to [**POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers**](POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers.md) |  | [optional] 
**Attachments** | Pointer to [**GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments**](GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments.md) |  | [optional] 
**Versions** | Pointer to [**POSTAddresses201ResponseDataRelationshipsVersions**](POSTAddresses201ResponseDataRelationshipsVersions.md) |  | [optional] 

## Methods

### NewPOSTShippingMethods201ResponseDataRelationships

`func NewPOSTShippingMethods201ResponseDataRelationships() *POSTShippingMethods201ResponseDataRelationships`

NewPOSTShippingMethods201ResponseDataRelationships instantiates a new POSTShippingMethods201ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTShippingMethods201ResponseDataRelationshipsWithDefaults

`func NewPOSTShippingMethods201ResponseDataRelationshipsWithDefaults() *POSTShippingMethods201ResponseDataRelationships`

NewPOSTShippingMethods201ResponseDataRelationshipsWithDefaults instantiates a new POSTShippingMethods201ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *POSTShippingMethods201ResponseDataRelationships) GetMarket() POSTBundles201ResponseDataRelationshipsMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *POSTShippingMethods201ResponseDataRelationships) GetMarketOk() (*POSTBundles201ResponseDataRelationshipsMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *POSTShippingMethods201ResponseDataRelationships) SetMarket(v POSTBundles201ResponseDataRelationshipsMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *POSTShippingMethods201ResponseDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetShippingZone

`func (o *POSTShippingMethods201ResponseDataRelationships) GetShippingZone() POSTShippingMethods201ResponseDataRelationshipsShippingZone`

GetShippingZone returns the ShippingZone field if non-nil, zero value otherwise.

### GetShippingZoneOk

`func (o *POSTShippingMethods201ResponseDataRelationships) GetShippingZoneOk() (*POSTShippingMethods201ResponseDataRelationshipsShippingZone, bool)`

GetShippingZoneOk returns a tuple with the ShippingZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingZone

`func (o *POSTShippingMethods201ResponseDataRelationships) SetShippingZone(v POSTShippingMethods201ResponseDataRelationshipsShippingZone)`

SetShippingZone sets ShippingZone field to given value.

### HasShippingZone

`func (o *POSTShippingMethods201ResponseDataRelationships) HasShippingZone() bool`

HasShippingZone returns a boolean if a field has been set.

### GetShippingCategory

`func (o *POSTShippingMethods201ResponseDataRelationships) GetShippingCategory() POSTShipments201ResponseDataRelationshipsShippingCategory`

GetShippingCategory returns the ShippingCategory field if non-nil, zero value otherwise.

### GetShippingCategoryOk

`func (o *POSTShippingMethods201ResponseDataRelationships) GetShippingCategoryOk() (*POSTShipments201ResponseDataRelationshipsShippingCategory, bool)`

GetShippingCategoryOk returns a tuple with the ShippingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCategory

`func (o *POSTShippingMethods201ResponseDataRelationships) SetShippingCategory(v POSTShipments201ResponseDataRelationshipsShippingCategory)`

SetShippingCategory sets ShippingCategory field to given value.

### HasShippingCategory

`func (o *POSTShippingMethods201ResponseDataRelationships) HasShippingCategory() bool`

HasShippingCategory returns a boolean if a field has been set.

### GetStockLocation

`func (o *POSTShippingMethods201ResponseDataRelationships) GetStockLocation() POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation`

GetStockLocation returns the StockLocation field if non-nil, zero value otherwise.

### GetStockLocationOk

`func (o *POSTShippingMethods201ResponseDataRelationships) GetStockLocationOk() (*POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation, bool)`

GetStockLocationOk returns a tuple with the StockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocation

`func (o *POSTShippingMethods201ResponseDataRelationships) SetStockLocation(v POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation)`

SetStockLocation sets StockLocation field to given value.

### HasStockLocation

`func (o *POSTShippingMethods201ResponseDataRelationships) HasStockLocation() bool`

HasStockLocation returns a boolean if a field has been set.

### GetDeliveryLeadTimeForShipment

`func (o *POSTShippingMethods201ResponseDataRelationships) GetDeliveryLeadTimeForShipment() POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment`

GetDeliveryLeadTimeForShipment returns the DeliveryLeadTimeForShipment field if non-nil, zero value otherwise.

### GetDeliveryLeadTimeForShipmentOk

`func (o *POSTShippingMethods201ResponseDataRelationships) GetDeliveryLeadTimeForShipmentOk() (*POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment, bool)`

GetDeliveryLeadTimeForShipmentOk returns a tuple with the DeliveryLeadTimeForShipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryLeadTimeForShipment

`func (o *POSTShippingMethods201ResponseDataRelationships) SetDeliveryLeadTimeForShipment(v POSTShippingMethods201ResponseDataRelationshipsDeliveryLeadTimeForShipment)`

SetDeliveryLeadTimeForShipment sets DeliveryLeadTimeForShipment field to given value.

### HasDeliveryLeadTimeForShipment

`func (o *POSTShippingMethods201ResponseDataRelationships) HasDeliveryLeadTimeForShipment() bool`

HasDeliveryLeadTimeForShipment returns a boolean if a field has been set.

### GetShippingMethodTiers

`func (o *POSTShippingMethods201ResponseDataRelationships) GetShippingMethodTiers() POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers`

GetShippingMethodTiers returns the ShippingMethodTiers field if non-nil, zero value otherwise.

### GetShippingMethodTiersOk

`func (o *POSTShippingMethods201ResponseDataRelationships) GetShippingMethodTiersOk() (*POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers, bool)`

GetShippingMethodTiersOk returns a tuple with the ShippingMethodTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethodTiers

`func (o *POSTShippingMethods201ResponseDataRelationships) SetShippingMethodTiers(v POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers)`

SetShippingMethodTiers sets ShippingMethodTiers field to given value.

### HasShippingMethodTiers

`func (o *POSTShippingMethods201ResponseDataRelationships) HasShippingMethodTiers() bool`

HasShippingMethodTiers returns a boolean if a field has been set.

### GetShippingWeightTiers

`func (o *POSTShippingMethods201ResponseDataRelationships) GetShippingWeightTiers() POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers`

GetShippingWeightTiers returns the ShippingWeightTiers field if non-nil, zero value otherwise.

### GetShippingWeightTiersOk

`func (o *POSTShippingMethods201ResponseDataRelationships) GetShippingWeightTiersOk() (*POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers, bool)`

GetShippingWeightTiersOk returns a tuple with the ShippingWeightTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingWeightTiers

`func (o *POSTShippingMethods201ResponseDataRelationships) SetShippingWeightTiers(v POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers)`

SetShippingWeightTiers sets ShippingWeightTiers field to given value.

### HasShippingWeightTiers

`func (o *POSTShippingMethods201ResponseDataRelationships) HasShippingWeightTiers() bool`

HasShippingWeightTiers returns a boolean if a field has been set.

### GetAttachments

`func (o *POSTShippingMethods201ResponseDataRelationships) GetAttachments() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *POSTShippingMethods201ResponseDataRelationships) GetAttachmentsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *POSTShippingMethods201ResponseDataRelationships) SetAttachments(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *POSTShippingMethods201ResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetVersions

`func (o *POSTShippingMethods201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *POSTShippingMethods201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *POSTShippingMethods201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *POSTShippingMethods201ResponseDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


