# POSTLineItems201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkuCode** | Pointer to **interface{}** | The code of the associated SKU. | [optional] 
**BundleCode** | Pointer to **interface{}** | The code of the associated bundle. | [optional] 
**Quantity** | **interface{}** | The line item quantity. | 
**ExternalPrice** | Pointer to **interface{}** | When creating or updating a new line item, set this attribute to &#39;1&#39; if you want to inject the unit_amount_cents price from an external source. Any successive price computation will be done externally, until the attribute is reset to &#39;0&#39;. | [optional] 
**UpdateQuantity** | Pointer to **interface{}** | When creating a new line item, set this attribute to &#39;1&#39; if you want to update the line item quantity (if present) instead of creating a new line item for the same SKU. | [optional] 
**ReserveStock** | Pointer to **interface{}** | Send this attribute if you want to reserve the stock for the line item&#39;s SKUs quantity. Stock reservations expiration depends on the inventory model&#39;s cutoff. When used on update the existing active stock reservations are renewed. Cannot be passed by sales channels. | [optional] 
**UnitAmountCents** | Pointer to **interface{}** | The unit amount of the line item, in cents. Can be specified only via an integration application, or when the item is missing, otherwise is automatically computed by using one of the available methods. Cannot be passed by sales channels. | [optional] 
**CompareAtAmountCents** | Pointer to **interface{}** | The compared price amount, in cents. Useful to display a percentage discount. | [optional] 
**Name** | Pointer to **interface{}** | The name of the line item. When blank, it gets populated with the name of the associated item (if present). | [optional] 
**ImageUrl** | Pointer to **interface{}** | The image_url of the line item. When blank, it gets populated with the image_url of the associated item (if present, SKU only). | [optional] 
**ItemType** | Pointer to **interface{}** | The type of the associated item. One of &#39;skus&#39;, &#39;bundles&#39;, &#39;gift_cards&#39;, &#39;shipments&#39;, &#39;payment_methods&#39;, &#39;adjustments&#39;, &#39;discount_engine_items&#39;, &#39;percentage_discount_promotions&#39;, &#39;free_shipping_promotions&#39;, &#39;buy_x_pay_y_promotions&#39;, &#39;free_gift_promotions&#39;, &#39;fixed_price_promotions&#39;, &#39;external_promotions&#39;, &#39;fixed_amount_promotions&#39;, or &#39;flex_promotions&#39;. | [optional] 
**Frequency** | Pointer to **interface{}** | The frequency which generates a subscription. Must be supported by existing associated subscription_model. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTLineItems201ResponseDataAttributes

`func NewPOSTLineItems201ResponseDataAttributes(quantity interface{}, ) *POSTLineItems201ResponseDataAttributes`

NewPOSTLineItems201ResponseDataAttributes instantiates a new POSTLineItems201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTLineItems201ResponseDataAttributesWithDefaults

`func NewPOSTLineItems201ResponseDataAttributesWithDefaults() *POSTLineItems201ResponseDataAttributes`

NewPOSTLineItems201ResponseDataAttributesWithDefaults instantiates a new POSTLineItems201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkuCode

`func (o *POSTLineItems201ResponseDataAttributes) GetSkuCode() interface{}`

GetSkuCode returns the SkuCode field if non-nil, zero value otherwise.

### GetSkuCodeOk

`func (o *POSTLineItems201ResponseDataAttributes) GetSkuCodeOk() (*interface{}, bool)`

GetSkuCodeOk returns a tuple with the SkuCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCode

`func (o *POSTLineItems201ResponseDataAttributes) SetSkuCode(v interface{})`

SetSkuCode sets SkuCode field to given value.

### HasSkuCode

`func (o *POSTLineItems201ResponseDataAttributes) HasSkuCode() bool`

HasSkuCode returns a boolean if a field has been set.

### SetSkuCodeNil

`func (o *POSTLineItems201ResponseDataAttributes) SetSkuCodeNil(b bool)`

 SetSkuCodeNil sets the value for SkuCode to be an explicit nil

### UnsetSkuCode
`func (o *POSTLineItems201ResponseDataAttributes) UnsetSkuCode()`

UnsetSkuCode ensures that no value is present for SkuCode, not even an explicit nil
### GetBundleCode

`func (o *POSTLineItems201ResponseDataAttributes) GetBundleCode() interface{}`

GetBundleCode returns the BundleCode field if non-nil, zero value otherwise.

### GetBundleCodeOk

`func (o *POSTLineItems201ResponseDataAttributes) GetBundleCodeOk() (*interface{}, bool)`

GetBundleCodeOk returns a tuple with the BundleCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleCode

`func (o *POSTLineItems201ResponseDataAttributes) SetBundleCode(v interface{})`

SetBundleCode sets BundleCode field to given value.

### HasBundleCode

`func (o *POSTLineItems201ResponseDataAttributes) HasBundleCode() bool`

HasBundleCode returns a boolean if a field has been set.

### SetBundleCodeNil

`func (o *POSTLineItems201ResponseDataAttributes) SetBundleCodeNil(b bool)`

 SetBundleCodeNil sets the value for BundleCode to be an explicit nil

### UnsetBundleCode
`func (o *POSTLineItems201ResponseDataAttributes) UnsetBundleCode()`

UnsetBundleCode ensures that no value is present for BundleCode, not even an explicit nil
### GetQuantity

`func (o *POSTLineItems201ResponseDataAttributes) GetQuantity() interface{}`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *POSTLineItems201ResponseDataAttributes) GetQuantityOk() (*interface{}, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *POSTLineItems201ResponseDataAttributes) SetQuantity(v interface{})`

SetQuantity sets Quantity field to given value.


### SetQuantityNil

`func (o *POSTLineItems201ResponseDataAttributes) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *POSTLineItems201ResponseDataAttributes) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetExternalPrice

`func (o *POSTLineItems201ResponseDataAttributes) GetExternalPrice() interface{}`

GetExternalPrice returns the ExternalPrice field if non-nil, zero value otherwise.

### GetExternalPriceOk

`func (o *POSTLineItems201ResponseDataAttributes) GetExternalPriceOk() (*interface{}, bool)`

GetExternalPriceOk returns a tuple with the ExternalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPrice

`func (o *POSTLineItems201ResponseDataAttributes) SetExternalPrice(v interface{})`

SetExternalPrice sets ExternalPrice field to given value.

### HasExternalPrice

`func (o *POSTLineItems201ResponseDataAttributes) HasExternalPrice() bool`

HasExternalPrice returns a boolean if a field has been set.

### SetExternalPriceNil

`func (o *POSTLineItems201ResponseDataAttributes) SetExternalPriceNil(b bool)`

 SetExternalPriceNil sets the value for ExternalPrice to be an explicit nil

### UnsetExternalPrice
`func (o *POSTLineItems201ResponseDataAttributes) UnsetExternalPrice()`

UnsetExternalPrice ensures that no value is present for ExternalPrice, not even an explicit nil
### GetUpdateQuantity

`func (o *POSTLineItems201ResponseDataAttributes) GetUpdateQuantity() interface{}`

GetUpdateQuantity returns the UpdateQuantity field if non-nil, zero value otherwise.

### GetUpdateQuantityOk

`func (o *POSTLineItems201ResponseDataAttributes) GetUpdateQuantityOk() (*interface{}, bool)`

GetUpdateQuantityOk returns a tuple with the UpdateQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateQuantity

`func (o *POSTLineItems201ResponseDataAttributes) SetUpdateQuantity(v interface{})`

SetUpdateQuantity sets UpdateQuantity field to given value.

### HasUpdateQuantity

`func (o *POSTLineItems201ResponseDataAttributes) HasUpdateQuantity() bool`

HasUpdateQuantity returns a boolean if a field has been set.

### SetUpdateQuantityNil

`func (o *POSTLineItems201ResponseDataAttributes) SetUpdateQuantityNil(b bool)`

 SetUpdateQuantityNil sets the value for UpdateQuantity to be an explicit nil

### UnsetUpdateQuantity
`func (o *POSTLineItems201ResponseDataAttributes) UnsetUpdateQuantity()`

UnsetUpdateQuantity ensures that no value is present for UpdateQuantity, not even an explicit nil
### GetReserveStock

`func (o *POSTLineItems201ResponseDataAttributes) GetReserveStock() interface{}`

GetReserveStock returns the ReserveStock field if non-nil, zero value otherwise.

### GetReserveStockOk

`func (o *POSTLineItems201ResponseDataAttributes) GetReserveStockOk() (*interface{}, bool)`

GetReserveStockOk returns a tuple with the ReserveStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserveStock

`func (o *POSTLineItems201ResponseDataAttributes) SetReserveStock(v interface{})`

SetReserveStock sets ReserveStock field to given value.

### HasReserveStock

`func (o *POSTLineItems201ResponseDataAttributes) HasReserveStock() bool`

HasReserveStock returns a boolean if a field has been set.

### SetReserveStockNil

`func (o *POSTLineItems201ResponseDataAttributes) SetReserveStockNil(b bool)`

 SetReserveStockNil sets the value for ReserveStock to be an explicit nil

### UnsetReserveStock
`func (o *POSTLineItems201ResponseDataAttributes) UnsetReserveStock()`

UnsetReserveStock ensures that no value is present for ReserveStock, not even an explicit nil
### GetUnitAmountCents

`func (o *POSTLineItems201ResponseDataAttributes) GetUnitAmountCents() interface{}`

GetUnitAmountCents returns the UnitAmountCents field if non-nil, zero value otherwise.

### GetUnitAmountCentsOk

`func (o *POSTLineItems201ResponseDataAttributes) GetUnitAmountCentsOk() (*interface{}, bool)`

GetUnitAmountCentsOk returns a tuple with the UnitAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountCents

`func (o *POSTLineItems201ResponseDataAttributes) SetUnitAmountCents(v interface{})`

SetUnitAmountCents sets UnitAmountCents field to given value.

### HasUnitAmountCents

`func (o *POSTLineItems201ResponseDataAttributes) HasUnitAmountCents() bool`

HasUnitAmountCents returns a boolean if a field has been set.

### SetUnitAmountCentsNil

`func (o *POSTLineItems201ResponseDataAttributes) SetUnitAmountCentsNil(b bool)`

 SetUnitAmountCentsNil sets the value for UnitAmountCents to be an explicit nil

### UnsetUnitAmountCents
`func (o *POSTLineItems201ResponseDataAttributes) UnsetUnitAmountCents()`

UnsetUnitAmountCents ensures that no value is present for UnitAmountCents, not even an explicit nil
### GetCompareAtAmountCents

`func (o *POSTLineItems201ResponseDataAttributes) GetCompareAtAmountCents() interface{}`

GetCompareAtAmountCents returns the CompareAtAmountCents field if non-nil, zero value otherwise.

### GetCompareAtAmountCentsOk

`func (o *POSTLineItems201ResponseDataAttributes) GetCompareAtAmountCentsOk() (*interface{}, bool)`

GetCompareAtAmountCentsOk returns a tuple with the CompareAtAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareAtAmountCents

`func (o *POSTLineItems201ResponseDataAttributes) SetCompareAtAmountCents(v interface{})`

SetCompareAtAmountCents sets CompareAtAmountCents field to given value.

### HasCompareAtAmountCents

`func (o *POSTLineItems201ResponseDataAttributes) HasCompareAtAmountCents() bool`

HasCompareAtAmountCents returns a boolean if a field has been set.

### SetCompareAtAmountCentsNil

`func (o *POSTLineItems201ResponseDataAttributes) SetCompareAtAmountCentsNil(b bool)`

 SetCompareAtAmountCentsNil sets the value for CompareAtAmountCents to be an explicit nil

### UnsetCompareAtAmountCents
`func (o *POSTLineItems201ResponseDataAttributes) UnsetCompareAtAmountCents()`

UnsetCompareAtAmountCents ensures that no value is present for CompareAtAmountCents, not even an explicit nil
### GetName

`func (o *POSTLineItems201ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTLineItems201ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTLineItems201ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *POSTLineItems201ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *POSTLineItems201ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *POSTLineItems201ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetImageUrl

`func (o *POSTLineItems201ResponseDataAttributes) GetImageUrl() interface{}`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *POSTLineItems201ResponseDataAttributes) GetImageUrlOk() (*interface{}, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *POSTLineItems201ResponseDataAttributes) SetImageUrl(v interface{})`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *POSTLineItems201ResponseDataAttributes) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *POSTLineItems201ResponseDataAttributes) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *POSTLineItems201ResponseDataAttributes) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetItemType

`func (o *POSTLineItems201ResponseDataAttributes) GetItemType() interface{}`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *POSTLineItems201ResponseDataAttributes) GetItemTypeOk() (*interface{}, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *POSTLineItems201ResponseDataAttributes) SetItemType(v interface{})`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *POSTLineItems201ResponseDataAttributes) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### SetItemTypeNil

`func (o *POSTLineItems201ResponseDataAttributes) SetItemTypeNil(b bool)`

 SetItemTypeNil sets the value for ItemType to be an explicit nil

### UnsetItemType
`func (o *POSTLineItems201ResponseDataAttributes) UnsetItemType()`

UnsetItemType ensures that no value is present for ItemType, not even an explicit nil
### GetFrequency

`func (o *POSTLineItems201ResponseDataAttributes) GetFrequency() interface{}`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *POSTLineItems201ResponseDataAttributes) GetFrequencyOk() (*interface{}, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *POSTLineItems201ResponseDataAttributes) SetFrequency(v interface{})`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *POSTLineItems201ResponseDataAttributes) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### SetFrequencyNil

`func (o *POSTLineItems201ResponseDataAttributes) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *POSTLineItems201ResponseDataAttributes) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
### GetReference

`func (o *POSTLineItems201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTLineItems201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTLineItems201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTLineItems201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTLineItems201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTLineItems201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTLineItems201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTLineItems201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTLineItems201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTLineItems201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTLineItems201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTLineItems201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTLineItems201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTLineItems201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTLineItems201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTLineItems201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTLineItems201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTLineItems201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


