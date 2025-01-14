# PATCHLineItemsLineItemId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkuCode** | Pointer to **interface{}** | The code of the associated SKU. | [optional] 
**BundleCode** | Pointer to **interface{}** | The code of the associated bundle. | [optional] 
**Quantity** | Pointer to **interface{}** | The line item quantity. | [optional] 
**ExternalPrice** | Pointer to **interface{}** | When creating or updating a new line item, set this attribute to &#39;1&#39; if you want to inject the unit_amount_cents price from an external source. Any successive price computation will be done externally, until the attribute is reset to &#39;0&#39;. | [optional] 
**ReserveStock** | Pointer to **interface{}** | Send this attribute if you want to reserve the stock for the line item&#39;s SKUs quantity. Stock reservations expiration depends on the inventory model&#39;s cutoff. When used on update the existing active stock reservations are renewed. Cannot be passed by sales channels. | [optional] 
**UnitAmountCents** | Pointer to **interface{}** | The unit amount of the line item, in cents. Can be specified only via an integration application, or when the item is missing, otherwise is automatically computed by using one of the available methods. Cannot be passed by sales channels. | [optional] 
**CompareAtAmountCents** | Pointer to **interface{}** | The compared price amount, in cents. Useful to display a percentage discount. | [optional] 
**OptionsAmountCents** | Pointer to **interface{}** | The options amount of the line item, in cents. Cannot be passed by sales channels. | [optional] 
**Name** | Pointer to **interface{}** | The name of the line item. When blank, it gets populated with the name of the associated item (if present). | [optional] 
**ImageUrl** | Pointer to **interface{}** | The image_url of the line item. When blank, it gets populated with the image_url of the associated item (if present, SKU only). | [optional] 
**Frequency** | Pointer to **interface{}** | The frequency which generates a subscription. Must be supported by existing associated subscription_model. | [optional] 
**ResetCircuit** | Pointer to **interface{}** | Send this attribute if you want to reset the circuit breaker associated to this resource to &#39;closed&#39; state and zero failures count. Cannot be passed by sales channels. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHLineItemsLineItemId200ResponseDataAttributes

`func NewPATCHLineItemsLineItemId200ResponseDataAttributes() *PATCHLineItemsLineItemId200ResponseDataAttributes`

NewPATCHLineItemsLineItemId200ResponseDataAttributes instantiates a new PATCHLineItemsLineItemId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHLineItemsLineItemId200ResponseDataAttributesWithDefaults

`func NewPATCHLineItemsLineItemId200ResponseDataAttributesWithDefaults() *PATCHLineItemsLineItemId200ResponseDataAttributes`

NewPATCHLineItemsLineItemId200ResponseDataAttributesWithDefaults instantiates a new PATCHLineItemsLineItemId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkuCode

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetSkuCode() interface{}`

GetSkuCode returns the SkuCode field if non-nil, zero value otherwise.

### GetSkuCodeOk

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetSkuCodeOk() (*interface{}, bool)`

GetSkuCodeOk returns a tuple with the SkuCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCode

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetSkuCode(v interface{})`

SetSkuCode sets SkuCode field to given value.

### HasSkuCode

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasSkuCode() bool`

HasSkuCode returns a boolean if a field has been set.

### SetSkuCodeNil

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetSkuCodeNil(b bool)`

 SetSkuCodeNil sets the value for SkuCode to be an explicit nil

### UnsetSkuCode
`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) UnsetSkuCode()`

UnsetSkuCode ensures that no value is present for SkuCode, not even an explicit nil
### GetBundleCode

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetBundleCode() interface{}`

GetBundleCode returns the BundleCode field if non-nil, zero value otherwise.

### GetBundleCodeOk

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetBundleCodeOk() (*interface{}, bool)`

GetBundleCodeOk returns a tuple with the BundleCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleCode

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetBundleCode(v interface{})`

SetBundleCode sets BundleCode field to given value.

### HasBundleCode

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasBundleCode() bool`

HasBundleCode returns a boolean if a field has been set.

### SetBundleCodeNil

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetBundleCodeNil(b bool)`

 SetBundleCodeNil sets the value for BundleCode to be an explicit nil

### UnsetBundleCode
`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) UnsetBundleCode()`

UnsetBundleCode ensures that no value is present for BundleCode, not even an explicit nil
### GetQuantity

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetQuantity() interface{}`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetQuantityOk() (*interface{}, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetQuantity(v interface{})`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetExternalPrice

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetExternalPrice() interface{}`

GetExternalPrice returns the ExternalPrice field if non-nil, zero value otherwise.

### GetExternalPriceOk

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetExternalPriceOk() (*interface{}, bool)`

GetExternalPriceOk returns a tuple with the ExternalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPrice

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetExternalPrice(v interface{})`

SetExternalPrice sets ExternalPrice field to given value.

### HasExternalPrice

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasExternalPrice() bool`

HasExternalPrice returns a boolean if a field has been set.

### SetExternalPriceNil

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetExternalPriceNil(b bool)`

 SetExternalPriceNil sets the value for ExternalPrice to be an explicit nil

### UnsetExternalPrice
`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) UnsetExternalPrice()`

UnsetExternalPrice ensures that no value is present for ExternalPrice, not even an explicit nil
### GetReserveStock

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetReserveStock() interface{}`

GetReserveStock returns the ReserveStock field if non-nil, zero value otherwise.

### GetReserveStockOk

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetReserveStockOk() (*interface{}, bool)`

GetReserveStockOk returns a tuple with the ReserveStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserveStock

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetReserveStock(v interface{})`

SetReserveStock sets ReserveStock field to given value.

### HasReserveStock

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasReserveStock() bool`

HasReserveStock returns a boolean if a field has been set.

### SetReserveStockNil

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetReserveStockNil(b bool)`

 SetReserveStockNil sets the value for ReserveStock to be an explicit nil

### UnsetReserveStock
`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) UnsetReserveStock()`

UnsetReserveStock ensures that no value is present for ReserveStock, not even an explicit nil
### GetUnitAmountCents

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetUnitAmountCents() interface{}`

GetUnitAmountCents returns the UnitAmountCents field if non-nil, zero value otherwise.

### GetUnitAmountCentsOk

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetUnitAmountCentsOk() (*interface{}, bool)`

GetUnitAmountCentsOk returns a tuple with the UnitAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountCents

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetUnitAmountCents(v interface{})`

SetUnitAmountCents sets UnitAmountCents field to given value.

### HasUnitAmountCents

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasUnitAmountCents() bool`

HasUnitAmountCents returns a boolean if a field has been set.

### SetUnitAmountCentsNil

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetUnitAmountCentsNil(b bool)`

 SetUnitAmountCentsNil sets the value for UnitAmountCents to be an explicit nil

### UnsetUnitAmountCents
`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) UnsetUnitAmountCents()`

UnsetUnitAmountCents ensures that no value is present for UnitAmountCents, not even an explicit nil
### GetCompareAtAmountCents

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetCompareAtAmountCents() interface{}`

GetCompareAtAmountCents returns the CompareAtAmountCents field if non-nil, zero value otherwise.

### GetCompareAtAmountCentsOk

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetCompareAtAmountCentsOk() (*interface{}, bool)`

GetCompareAtAmountCentsOk returns a tuple with the CompareAtAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareAtAmountCents

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetCompareAtAmountCents(v interface{})`

SetCompareAtAmountCents sets CompareAtAmountCents field to given value.

### HasCompareAtAmountCents

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasCompareAtAmountCents() bool`

HasCompareAtAmountCents returns a boolean if a field has been set.

### SetCompareAtAmountCentsNil

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetCompareAtAmountCentsNil(b bool)`

 SetCompareAtAmountCentsNil sets the value for CompareAtAmountCents to be an explicit nil

### UnsetCompareAtAmountCents
`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) UnsetCompareAtAmountCents()`

UnsetCompareAtAmountCents ensures that no value is present for CompareAtAmountCents, not even an explicit nil
### GetOptionsAmountCents

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetOptionsAmountCents() interface{}`

GetOptionsAmountCents returns the OptionsAmountCents field if non-nil, zero value otherwise.

### GetOptionsAmountCentsOk

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetOptionsAmountCentsOk() (*interface{}, bool)`

GetOptionsAmountCentsOk returns a tuple with the OptionsAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionsAmountCents

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetOptionsAmountCents(v interface{})`

SetOptionsAmountCents sets OptionsAmountCents field to given value.

### HasOptionsAmountCents

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasOptionsAmountCents() bool`

HasOptionsAmountCents returns a boolean if a field has been set.

### SetOptionsAmountCentsNil

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetOptionsAmountCentsNil(b bool)`

 SetOptionsAmountCentsNil sets the value for OptionsAmountCents to be an explicit nil

### UnsetOptionsAmountCents
`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) UnsetOptionsAmountCents()`

UnsetOptionsAmountCents ensures that no value is present for OptionsAmountCents, not even an explicit nil
### GetName

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetImageUrl

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetImageUrl() interface{}`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetImageUrlOk() (*interface{}, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetImageUrl(v interface{})`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetFrequency

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetFrequency() interface{}`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetFrequencyOk() (*interface{}, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetFrequency(v interface{})`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### SetFrequencyNil

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
### GetResetCircuit

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetResetCircuit() interface{}`

GetResetCircuit returns the ResetCircuit field if non-nil, zero value otherwise.

### GetResetCircuitOk

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetResetCircuitOk() (*interface{}, bool)`

GetResetCircuitOk returns a tuple with the ResetCircuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetCircuit

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetResetCircuit(v interface{})`

SetResetCircuit sets ResetCircuit field to given value.

### HasResetCircuit

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasResetCircuit() bool`

HasResetCircuit returns a boolean if a field has been set.

### SetResetCircuitNil

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetResetCircuitNil(b bool)`

 SetResetCircuitNil sets the value for ResetCircuit to be an explicit nil

### UnsetResetCircuit
`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) UnsetResetCircuit()`

UnsetResetCircuit ensures that no value is present for ResetCircuit, not even an explicit nil
### GetReference

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


