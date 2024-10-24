# GETImportsImportId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | Pointer to **interface{}** | The type of resource being imported. | [optional] 
**Format** | Pointer to **interface{}** | The format of the import inputs one of &#39;json&#39; (default) or &#39;csv&#39;. | [optional] 
**ParentResourceId** | Pointer to **interface{}** | The ID of the parent resource to be associated with imported data. | [optional] 
**Status** | Pointer to **interface{}** | The import job status. One of &#39;pending&#39; (default), &#39;in_progress&#39;, &#39;interrupted&#39;, or &#39;completed&#39;. | [optional] 
**StartedAt** | Pointer to **interface{}** | Time at which the import was started. | [optional] 
**CompletedAt** | Pointer to **interface{}** | Time at which the import was completed. | [optional] 
**InterruptedAt** | Pointer to **interface{}** | Time at which the import was interrupted. | [optional] 
**Inputs** | Pointer to **interface{}** | Array of objects representing the resources that are being imported. | [optional] 
**InputsSize** | Pointer to **interface{}** | Indicates the size of the objects to be imported. | [optional] 
**ErrorsCount** | Pointer to **interface{}** | Indicates the number of import errors, if any. | [optional] 
**WarningsCount** | Pointer to **interface{}** | Indicates the number of import warnings, if any. | [optional] 
**ProcessedCount** | Pointer to **interface{}** | Indicates the number of records that have been processed (created or updated). | [optional] 
**ErrorsLog** | Pointer to **interface{}** | Contains the import errors, if any. | [optional] 
**WarningsLog** | Pointer to **interface{}** | Contains the import warnings, if any. | [optional] 
**AttachmentUrl** | Pointer to **interface{}** | The URL the the raw inputs file, which will be generated at import start. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETImportsImportId200ResponseDataAttributes

`func NewGETImportsImportId200ResponseDataAttributes() *GETImportsImportId200ResponseDataAttributes`

NewGETImportsImportId200ResponseDataAttributes instantiates a new GETImportsImportId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETImportsImportId200ResponseDataAttributesWithDefaults

`func NewGETImportsImportId200ResponseDataAttributesWithDefaults() *GETImportsImportId200ResponseDataAttributes`

NewGETImportsImportId200ResponseDataAttributesWithDefaults instantiates a new GETImportsImportId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *GETImportsImportId200ResponseDataAttributes) GetResourceType() interface{}`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *GETImportsImportId200ResponseDataAttributes) GetResourceTypeOk() (*interface{}, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *GETImportsImportId200ResponseDataAttributes) SetResourceType(v interface{})`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *GETImportsImportId200ResponseDataAttributes) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *GETImportsImportId200ResponseDataAttributes) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *GETImportsImportId200ResponseDataAttributes) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetFormat

`func (o *GETImportsImportId200ResponseDataAttributes) GetFormat() interface{}`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *GETImportsImportId200ResponseDataAttributes) GetFormatOk() (*interface{}, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *GETImportsImportId200ResponseDataAttributes) SetFormat(v interface{})`

SetFormat sets Format field to given value.

### HasFormat

`func (o *GETImportsImportId200ResponseDataAttributes) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *GETImportsImportId200ResponseDataAttributes) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *GETImportsImportId200ResponseDataAttributes) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetParentResourceId

`func (o *GETImportsImportId200ResponseDataAttributes) GetParentResourceId() interface{}`

GetParentResourceId returns the ParentResourceId field if non-nil, zero value otherwise.

### GetParentResourceIdOk

`func (o *GETImportsImportId200ResponseDataAttributes) GetParentResourceIdOk() (*interface{}, bool)`

GetParentResourceIdOk returns a tuple with the ParentResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentResourceId

`func (o *GETImportsImportId200ResponseDataAttributes) SetParentResourceId(v interface{})`

SetParentResourceId sets ParentResourceId field to given value.

### HasParentResourceId

`func (o *GETImportsImportId200ResponseDataAttributes) HasParentResourceId() bool`

HasParentResourceId returns a boolean if a field has been set.

### SetParentResourceIdNil

`func (o *GETImportsImportId200ResponseDataAttributes) SetParentResourceIdNil(b bool)`

 SetParentResourceIdNil sets the value for ParentResourceId to be an explicit nil

### UnsetParentResourceId
`func (o *GETImportsImportId200ResponseDataAttributes) UnsetParentResourceId()`

UnsetParentResourceId ensures that no value is present for ParentResourceId, not even an explicit nil
### GetStatus

`func (o *GETImportsImportId200ResponseDataAttributes) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETImportsImportId200ResponseDataAttributes) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETImportsImportId200ResponseDataAttributes) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETImportsImportId200ResponseDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GETImportsImportId200ResponseDataAttributes) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GETImportsImportId200ResponseDataAttributes) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStartedAt

`func (o *GETImportsImportId200ResponseDataAttributes) GetStartedAt() interface{}`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *GETImportsImportId200ResponseDataAttributes) GetStartedAtOk() (*interface{}, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *GETImportsImportId200ResponseDataAttributes) SetStartedAt(v interface{})`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *GETImportsImportId200ResponseDataAttributes) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *GETImportsImportId200ResponseDataAttributes) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *GETImportsImportId200ResponseDataAttributes) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *GETImportsImportId200ResponseDataAttributes) GetCompletedAt() interface{}`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GETImportsImportId200ResponseDataAttributes) GetCompletedAtOk() (*interface{}, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GETImportsImportId200ResponseDataAttributes) SetCompletedAt(v interface{})`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GETImportsImportId200ResponseDataAttributes) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *GETImportsImportId200ResponseDataAttributes) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *GETImportsImportId200ResponseDataAttributes) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetInterruptedAt

`func (o *GETImportsImportId200ResponseDataAttributes) GetInterruptedAt() interface{}`

GetInterruptedAt returns the InterruptedAt field if non-nil, zero value otherwise.

### GetInterruptedAtOk

`func (o *GETImportsImportId200ResponseDataAttributes) GetInterruptedAtOk() (*interface{}, bool)`

GetInterruptedAtOk returns a tuple with the InterruptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterruptedAt

`func (o *GETImportsImportId200ResponseDataAttributes) SetInterruptedAt(v interface{})`

SetInterruptedAt sets InterruptedAt field to given value.

### HasInterruptedAt

`func (o *GETImportsImportId200ResponseDataAttributes) HasInterruptedAt() bool`

HasInterruptedAt returns a boolean if a field has been set.

### SetInterruptedAtNil

`func (o *GETImportsImportId200ResponseDataAttributes) SetInterruptedAtNil(b bool)`

 SetInterruptedAtNil sets the value for InterruptedAt to be an explicit nil

### UnsetInterruptedAt
`func (o *GETImportsImportId200ResponseDataAttributes) UnsetInterruptedAt()`

UnsetInterruptedAt ensures that no value is present for InterruptedAt, not even an explicit nil
### GetInputs

`func (o *GETImportsImportId200ResponseDataAttributes) GetInputs() interface{}`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *GETImportsImportId200ResponseDataAttributes) GetInputsOk() (*interface{}, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *GETImportsImportId200ResponseDataAttributes) SetInputs(v interface{})`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *GETImportsImportId200ResponseDataAttributes) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### SetInputsNil

`func (o *GETImportsImportId200ResponseDataAttributes) SetInputsNil(b bool)`

 SetInputsNil sets the value for Inputs to be an explicit nil

### UnsetInputs
`func (o *GETImportsImportId200ResponseDataAttributes) UnsetInputs()`

UnsetInputs ensures that no value is present for Inputs, not even an explicit nil
### GetInputsSize

`func (o *GETImportsImportId200ResponseDataAttributes) GetInputsSize() interface{}`

GetInputsSize returns the InputsSize field if non-nil, zero value otherwise.

### GetInputsSizeOk

`func (o *GETImportsImportId200ResponseDataAttributes) GetInputsSizeOk() (*interface{}, bool)`

GetInputsSizeOk returns a tuple with the InputsSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputsSize

`func (o *GETImportsImportId200ResponseDataAttributes) SetInputsSize(v interface{})`

SetInputsSize sets InputsSize field to given value.

### HasInputsSize

`func (o *GETImportsImportId200ResponseDataAttributes) HasInputsSize() bool`

HasInputsSize returns a boolean if a field has been set.

### SetInputsSizeNil

`func (o *GETImportsImportId200ResponseDataAttributes) SetInputsSizeNil(b bool)`

 SetInputsSizeNil sets the value for InputsSize to be an explicit nil

### UnsetInputsSize
`func (o *GETImportsImportId200ResponseDataAttributes) UnsetInputsSize()`

UnsetInputsSize ensures that no value is present for InputsSize, not even an explicit nil
### GetErrorsCount

`func (o *GETImportsImportId200ResponseDataAttributes) GetErrorsCount() interface{}`

GetErrorsCount returns the ErrorsCount field if non-nil, zero value otherwise.

### GetErrorsCountOk

`func (o *GETImportsImportId200ResponseDataAttributes) GetErrorsCountOk() (*interface{}, bool)`

GetErrorsCountOk returns a tuple with the ErrorsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorsCount

`func (o *GETImportsImportId200ResponseDataAttributes) SetErrorsCount(v interface{})`

SetErrorsCount sets ErrorsCount field to given value.

### HasErrorsCount

`func (o *GETImportsImportId200ResponseDataAttributes) HasErrorsCount() bool`

HasErrorsCount returns a boolean if a field has been set.

### SetErrorsCountNil

`func (o *GETImportsImportId200ResponseDataAttributes) SetErrorsCountNil(b bool)`

 SetErrorsCountNil sets the value for ErrorsCount to be an explicit nil

### UnsetErrorsCount
`func (o *GETImportsImportId200ResponseDataAttributes) UnsetErrorsCount()`

UnsetErrorsCount ensures that no value is present for ErrorsCount, not even an explicit nil
### GetWarningsCount

`func (o *GETImportsImportId200ResponseDataAttributes) GetWarningsCount() interface{}`

GetWarningsCount returns the WarningsCount field if non-nil, zero value otherwise.

### GetWarningsCountOk

`func (o *GETImportsImportId200ResponseDataAttributes) GetWarningsCountOk() (*interface{}, bool)`

GetWarningsCountOk returns a tuple with the WarningsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningsCount

`func (o *GETImportsImportId200ResponseDataAttributes) SetWarningsCount(v interface{})`

SetWarningsCount sets WarningsCount field to given value.

### HasWarningsCount

`func (o *GETImportsImportId200ResponseDataAttributes) HasWarningsCount() bool`

HasWarningsCount returns a boolean if a field has been set.

### SetWarningsCountNil

`func (o *GETImportsImportId200ResponseDataAttributes) SetWarningsCountNil(b bool)`

 SetWarningsCountNil sets the value for WarningsCount to be an explicit nil

### UnsetWarningsCount
`func (o *GETImportsImportId200ResponseDataAttributes) UnsetWarningsCount()`

UnsetWarningsCount ensures that no value is present for WarningsCount, not even an explicit nil
### GetProcessedCount

`func (o *GETImportsImportId200ResponseDataAttributes) GetProcessedCount() interface{}`

GetProcessedCount returns the ProcessedCount field if non-nil, zero value otherwise.

### GetProcessedCountOk

`func (o *GETImportsImportId200ResponseDataAttributes) GetProcessedCountOk() (*interface{}, bool)`

GetProcessedCountOk returns a tuple with the ProcessedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedCount

`func (o *GETImportsImportId200ResponseDataAttributes) SetProcessedCount(v interface{})`

SetProcessedCount sets ProcessedCount field to given value.

### HasProcessedCount

`func (o *GETImportsImportId200ResponseDataAttributes) HasProcessedCount() bool`

HasProcessedCount returns a boolean if a field has been set.

### SetProcessedCountNil

`func (o *GETImportsImportId200ResponseDataAttributes) SetProcessedCountNil(b bool)`

 SetProcessedCountNil sets the value for ProcessedCount to be an explicit nil

### UnsetProcessedCount
`func (o *GETImportsImportId200ResponseDataAttributes) UnsetProcessedCount()`

UnsetProcessedCount ensures that no value is present for ProcessedCount, not even an explicit nil
### GetErrorsLog

`func (o *GETImportsImportId200ResponseDataAttributes) GetErrorsLog() interface{}`

GetErrorsLog returns the ErrorsLog field if non-nil, zero value otherwise.

### GetErrorsLogOk

`func (o *GETImportsImportId200ResponseDataAttributes) GetErrorsLogOk() (*interface{}, bool)`

GetErrorsLogOk returns a tuple with the ErrorsLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorsLog

`func (o *GETImportsImportId200ResponseDataAttributes) SetErrorsLog(v interface{})`

SetErrorsLog sets ErrorsLog field to given value.

### HasErrorsLog

`func (o *GETImportsImportId200ResponseDataAttributes) HasErrorsLog() bool`

HasErrorsLog returns a boolean if a field has been set.

### SetErrorsLogNil

`func (o *GETImportsImportId200ResponseDataAttributes) SetErrorsLogNil(b bool)`

 SetErrorsLogNil sets the value for ErrorsLog to be an explicit nil

### UnsetErrorsLog
`func (o *GETImportsImportId200ResponseDataAttributes) UnsetErrorsLog()`

UnsetErrorsLog ensures that no value is present for ErrorsLog, not even an explicit nil
### GetWarningsLog

`func (o *GETImportsImportId200ResponseDataAttributes) GetWarningsLog() interface{}`

GetWarningsLog returns the WarningsLog field if non-nil, zero value otherwise.

### GetWarningsLogOk

`func (o *GETImportsImportId200ResponseDataAttributes) GetWarningsLogOk() (*interface{}, bool)`

GetWarningsLogOk returns a tuple with the WarningsLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningsLog

`func (o *GETImportsImportId200ResponseDataAttributes) SetWarningsLog(v interface{})`

SetWarningsLog sets WarningsLog field to given value.

### HasWarningsLog

`func (o *GETImportsImportId200ResponseDataAttributes) HasWarningsLog() bool`

HasWarningsLog returns a boolean if a field has been set.

### SetWarningsLogNil

`func (o *GETImportsImportId200ResponseDataAttributes) SetWarningsLogNil(b bool)`

 SetWarningsLogNil sets the value for WarningsLog to be an explicit nil

### UnsetWarningsLog
`func (o *GETImportsImportId200ResponseDataAttributes) UnsetWarningsLog()`

UnsetWarningsLog ensures that no value is present for WarningsLog, not even an explicit nil
### GetAttachmentUrl

`func (o *GETImportsImportId200ResponseDataAttributes) GetAttachmentUrl() interface{}`

GetAttachmentUrl returns the AttachmentUrl field if non-nil, zero value otherwise.

### GetAttachmentUrlOk

`func (o *GETImportsImportId200ResponseDataAttributes) GetAttachmentUrlOk() (*interface{}, bool)`

GetAttachmentUrlOk returns a tuple with the AttachmentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentUrl

`func (o *GETImportsImportId200ResponseDataAttributes) SetAttachmentUrl(v interface{})`

SetAttachmentUrl sets AttachmentUrl field to given value.

### HasAttachmentUrl

`func (o *GETImportsImportId200ResponseDataAttributes) HasAttachmentUrl() bool`

HasAttachmentUrl returns a boolean if a field has been set.

### SetAttachmentUrlNil

`func (o *GETImportsImportId200ResponseDataAttributes) SetAttachmentUrlNil(b bool)`

 SetAttachmentUrlNil sets the value for AttachmentUrl to be an explicit nil

### UnsetAttachmentUrl
`func (o *GETImportsImportId200ResponseDataAttributes) UnsetAttachmentUrl()`

UnsetAttachmentUrl ensures that no value is present for AttachmentUrl, not even an explicit nil
### GetCreatedAt

`func (o *GETImportsImportId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETImportsImportId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETImportsImportId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETImportsImportId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETImportsImportId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETImportsImportId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETImportsImportId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETImportsImportId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETImportsImportId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETImportsImportId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETImportsImportId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETImportsImportId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETImportsImportId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETImportsImportId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETImportsImportId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETImportsImportId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETImportsImportId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETImportsImportId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETImportsImportId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETImportsImportId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETImportsImportId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETImportsImportId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETImportsImportId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETImportsImportId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETImportsImportId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETImportsImportId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETImportsImportId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETImportsImportId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETImportsImportId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETImportsImportId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


