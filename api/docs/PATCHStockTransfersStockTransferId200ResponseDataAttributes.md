# PATCHStockTransfersStockTransferId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **interface{}** | Unique identifier for the stock transfer (numeric). | [optional] 
**SkuCode** | Pointer to **interface{}** | The code of the associated SKU. | [optional] 
**Quantity** | Pointer to **interface{}** | The stock quantity to be transferred from the origin stock location to destination one. Updatable unless stock transfer is completed or cancelled and depending on origin stock availability. | [optional] 
**Upcoming** | Pointer to **interface{}** | Send this attribute if you want to mark this stock transfer as upcoming. | [optional] 
**OnHold** | Pointer to **interface{}** | Send this attribute if you want to put this stock transfer on hold. | [optional] 
**Picking** | Pointer to **interface{}** | Send this attribute if you want to start picking this stock transfer. | [optional] 
**InTransit** | Pointer to **interface{}** | Send this attribute if you want to mark this stock transfer as in transit. | [optional] 
**Complete** | Pointer to **interface{}** | Send this attribute if you want to complete this stock transfer. | [optional] 
**Cancel** | Pointer to **interface{}** | Send this attribute if you want to cancel this stock transfer. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHStockTransfersStockTransferId200ResponseDataAttributes

`func NewPATCHStockTransfersStockTransferId200ResponseDataAttributes() *PATCHStockTransfersStockTransferId200ResponseDataAttributes`

NewPATCHStockTransfersStockTransferId200ResponseDataAttributes instantiates a new PATCHStockTransfersStockTransferId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHStockTransfersStockTransferId200ResponseDataAttributesWithDefaults

`func NewPATCHStockTransfersStockTransferId200ResponseDataAttributesWithDefaults() *PATCHStockTransfersStockTransferId200ResponseDataAttributes`

NewPATCHStockTransfersStockTransferId200ResponseDataAttributesWithDefaults instantiates a new PATCHStockTransfersStockTransferId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetNumber() interface{}`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetNumberOk() (*interface{}, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetNumber(v interface{})`

SetNumber sets Number field to given value.

### HasNumber

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetSkuCode

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetSkuCode() interface{}`

GetSkuCode returns the SkuCode field if non-nil, zero value otherwise.

### GetSkuCodeOk

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetSkuCodeOk() (*interface{}, bool)`

GetSkuCodeOk returns a tuple with the SkuCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCode

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetSkuCode(v interface{})`

SetSkuCode sets SkuCode field to given value.

### HasSkuCode

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) HasSkuCode() bool`

HasSkuCode returns a boolean if a field has been set.

### SetSkuCodeNil

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetSkuCodeNil(b bool)`

 SetSkuCodeNil sets the value for SkuCode to be an explicit nil

### UnsetSkuCode
`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) UnsetSkuCode()`

UnsetSkuCode ensures that no value is present for SkuCode, not even an explicit nil
### GetQuantity

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetQuantity() interface{}`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetQuantityOk() (*interface{}, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetQuantity(v interface{})`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetUpcoming

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetUpcoming() interface{}`

GetUpcoming returns the Upcoming field if non-nil, zero value otherwise.

### GetUpcomingOk

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetUpcomingOk() (*interface{}, bool)`

GetUpcomingOk returns a tuple with the Upcoming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpcoming

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetUpcoming(v interface{})`

SetUpcoming sets Upcoming field to given value.

### HasUpcoming

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) HasUpcoming() bool`

HasUpcoming returns a boolean if a field has been set.

### SetUpcomingNil

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetUpcomingNil(b bool)`

 SetUpcomingNil sets the value for Upcoming to be an explicit nil

### UnsetUpcoming
`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) UnsetUpcoming()`

UnsetUpcoming ensures that no value is present for Upcoming, not even an explicit nil
### GetOnHold

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetOnHold() interface{}`

GetOnHold returns the OnHold field if non-nil, zero value otherwise.

### GetOnHoldOk

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetOnHoldOk() (*interface{}, bool)`

GetOnHoldOk returns a tuple with the OnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnHold

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetOnHold(v interface{})`

SetOnHold sets OnHold field to given value.

### HasOnHold

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) HasOnHold() bool`

HasOnHold returns a boolean if a field has been set.

### SetOnHoldNil

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetOnHoldNil(b bool)`

 SetOnHoldNil sets the value for OnHold to be an explicit nil

### UnsetOnHold
`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) UnsetOnHold()`

UnsetOnHold ensures that no value is present for OnHold, not even an explicit nil
### GetPicking

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetPicking() interface{}`

GetPicking returns the Picking field if non-nil, zero value otherwise.

### GetPickingOk

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetPickingOk() (*interface{}, bool)`

GetPickingOk returns a tuple with the Picking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicking

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetPicking(v interface{})`

SetPicking sets Picking field to given value.

### HasPicking

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) HasPicking() bool`

HasPicking returns a boolean if a field has been set.

### SetPickingNil

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetPickingNil(b bool)`

 SetPickingNil sets the value for Picking to be an explicit nil

### UnsetPicking
`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) UnsetPicking()`

UnsetPicking ensures that no value is present for Picking, not even an explicit nil
### GetInTransit

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetInTransit() interface{}`

GetInTransit returns the InTransit field if non-nil, zero value otherwise.

### GetInTransitOk

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetInTransitOk() (*interface{}, bool)`

GetInTransitOk returns a tuple with the InTransit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInTransit

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetInTransit(v interface{})`

SetInTransit sets InTransit field to given value.

### HasInTransit

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) HasInTransit() bool`

HasInTransit returns a boolean if a field has been set.

### SetInTransitNil

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetInTransitNil(b bool)`

 SetInTransitNil sets the value for InTransit to be an explicit nil

### UnsetInTransit
`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) UnsetInTransit()`

UnsetInTransit ensures that no value is present for InTransit, not even an explicit nil
### GetComplete

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetComplete() interface{}`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetCompleteOk() (*interface{}, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetComplete(v interface{})`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### SetCompleteNil

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetCompleteNil(b bool)`

 SetCompleteNil sets the value for Complete to be an explicit nil

### UnsetComplete
`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) UnsetComplete()`

UnsetComplete ensures that no value is present for Complete, not even an explicit nil
### GetCancel

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetCancel() interface{}`

GetCancel returns the Cancel field if non-nil, zero value otherwise.

### GetCancelOk

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetCancelOk() (*interface{}, bool)`

GetCancelOk returns a tuple with the Cancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancel

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetCancel(v interface{})`

SetCancel sets Cancel field to given value.

### HasCancel

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) HasCancel() bool`

HasCancel returns a boolean if a field has been set.

### SetCancelNil

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetCancelNil(b bool)`

 SetCancelNil sets the value for Cancel to be an explicit nil

### UnsetCancel
`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) UnsetCancel()`

UnsetCancel ensures that no value is present for Cancel, not even an explicit nil
### GetReference

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHStockTransfersStockTransferId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


