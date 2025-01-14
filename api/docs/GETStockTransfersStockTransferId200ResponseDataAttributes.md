# GETStockTransfersStockTransferId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **interface{}** | Unique identifier for the stock transfer (numeric). | [optional] 
**SkuCode** | Pointer to **interface{}** | The code of the associated SKU. | [optional] 
**Status** | Pointer to **interface{}** | The stock transfer status. One of &#39;draft&#39; (default), &#39;upcoming&#39;, &#39;on_hold&#39;, &#39;picking&#39;, &#39;in_transit&#39;, &#39;completed&#39;, or &#39;cancelled&#39;. | [optional] 
**Quantity** | Pointer to **interface{}** | The stock quantity to be transferred from the origin stock location to destination one. Updatable unless stock transfer is completed or cancelled and depending on origin stock availability. | [optional] 
**OnHoldAt** | Pointer to **interface{}** | Time at which the stock transfer was put on hold. | [optional] 
**PickingAt** | Pointer to **interface{}** | Time at which the stock transfer was picking. | [optional] 
**InTransitAt** | Pointer to **interface{}** | Time at which the stock transfer was in transit. | [optional] 
**CompletedAt** | Pointer to **interface{}** | Time at which the stock transfer was completed. | [optional] 
**CancelledAt** | Pointer to **interface{}** | Time at which the stock transfer was cancelled. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETStockTransfersStockTransferId200ResponseDataAttributes

`func NewGETStockTransfersStockTransferId200ResponseDataAttributes() *GETStockTransfersStockTransferId200ResponseDataAttributes`

NewGETStockTransfersStockTransferId200ResponseDataAttributes instantiates a new GETStockTransfersStockTransferId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETStockTransfersStockTransferId200ResponseDataAttributesWithDefaults

`func NewGETStockTransfersStockTransferId200ResponseDataAttributesWithDefaults() *GETStockTransfersStockTransferId200ResponseDataAttributes`

NewGETStockTransfersStockTransferId200ResponseDataAttributesWithDefaults instantiates a new GETStockTransfersStockTransferId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) GetNumber() interface{}`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) GetNumberOk() (*interface{}, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) SetNumber(v interface{})`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetSkuCode

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) GetSkuCode() interface{}`

GetSkuCode returns the SkuCode field if non-nil, zero value otherwise.

### GetSkuCodeOk

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) GetSkuCodeOk() (*interface{}, bool)`

GetSkuCodeOk returns a tuple with the SkuCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCode

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) SetSkuCode(v interface{})`

SetSkuCode sets SkuCode field to given value.

### HasSkuCode

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) HasSkuCode() bool`

HasSkuCode returns a boolean if a field has been set.

### SetSkuCodeNil

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) SetSkuCodeNil(b bool)`

 SetSkuCodeNil sets the value for SkuCode to be an explicit nil

### UnsetSkuCode
`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) UnsetSkuCode()`

UnsetSkuCode ensures that no value is present for SkuCode, not even an explicit nil
### GetStatus

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetQuantity

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) GetQuantity() interface{}`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) GetQuantityOk() (*interface{}, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) SetQuantity(v interface{})`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetOnHoldAt

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) GetOnHoldAt() interface{}`

GetOnHoldAt returns the OnHoldAt field if non-nil, zero value otherwise.

### GetOnHoldAtOk

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) GetOnHoldAtOk() (*interface{}, bool)`

GetOnHoldAtOk returns a tuple with the OnHoldAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnHoldAt

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) SetOnHoldAt(v interface{})`

SetOnHoldAt sets OnHoldAt field to given value.

### HasOnHoldAt

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) HasOnHoldAt() bool`

HasOnHoldAt returns a boolean if a field has been set.

### SetOnHoldAtNil

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) SetOnHoldAtNil(b bool)`

 SetOnHoldAtNil sets the value for OnHoldAt to be an explicit nil

### UnsetOnHoldAt
`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) UnsetOnHoldAt()`

UnsetOnHoldAt ensures that no value is present for OnHoldAt, not even an explicit nil
### GetPickingAt

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) GetPickingAt() interface{}`

GetPickingAt returns the PickingAt field if non-nil, zero value otherwise.

### GetPickingAtOk

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) GetPickingAtOk() (*interface{}, bool)`

GetPickingAtOk returns a tuple with the PickingAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickingAt

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) SetPickingAt(v interface{})`

SetPickingAt sets PickingAt field to given value.

### HasPickingAt

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) HasPickingAt() bool`

HasPickingAt returns a boolean if a field has been set.

### SetPickingAtNil

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) SetPickingAtNil(b bool)`

 SetPickingAtNil sets the value for PickingAt to be an explicit nil

### UnsetPickingAt
`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) UnsetPickingAt()`

UnsetPickingAt ensures that no value is present for PickingAt, not even an explicit nil
### GetInTransitAt

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) GetInTransitAt() interface{}`

GetInTransitAt returns the InTransitAt field if non-nil, zero value otherwise.

### GetInTransitAtOk

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) GetInTransitAtOk() (*interface{}, bool)`

GetInTransitAtOk returns a tuple with the InTransitAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInTransitAt

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) SetInTransitAt(v interface{})`

SetInTransitAt sets InTransitAt field to given value.

### HasInTransitAt

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) HasInTransitAt() bool`

HasInTransitAt returns a boolean if a field has been set.

### SetInTransitAtNil

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) SetInTransitAtNil(b bool)`

 SetInTransitAtNil sets the value for InTransitAt to be an explicit nil

### UnsetInTransitAt
`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) UnsetInTransitAt()`

UnsetInTransitAt ensures that no value is present for InTransitAt, not even an explicit nil
### GetCompletedAt

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) GetCompletedAt() interface{}`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) GetCompletedAtOk() (*interface{}, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) SetCompletedAt(v interface{})`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetCancelledAt

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) GetCancelledAt() interface{}`

GetCancelledAt returns the CancelledAt field if non-nil, zero value otherwise.

### GetCancelledAtOk

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) GetCancelledAtOk() (*interface{}, bool)`

GetCancelledAtOk returns a tuple with the CancelledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledAt

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) SetCancelledAt(v interface{})`

SetCancelledAt sets CancelledAt field to given value.

### HasCancelledAt

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) HasCancelledAt() bool`

HasCancelledAt returns a boolean if a field has been set.

### SetCancelledAtNil

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) SetCancelledAtNil(b bool)`

 SetCancelledAtNil sets the value for CancelledAt to be an explicit nil

### UnsetCancelledAt
`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) UnsetCancelledAt()`

UnsetCancelledAt ensures that no value is present for CancelledAt, not even an explicit nil
### GetCreatedAt

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETStockTransfersStockTransferId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


