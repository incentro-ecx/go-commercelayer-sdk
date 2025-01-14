# GETExportsExportId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | Pointer to **interface{}** | The type of resource being exported. | [optional] 
**Format** | Pointer to **interface{}** | The format of the export one of &#39;json&#39; (default) or &#39;csv&#39;. | [optional] 
**Status** | Pointer to **interface{}** | The export job status. One of &#39;pending&#39; (default), &#39;in_progress&#39;, &#39;interrupted&#39;, or &#39;completed&#39;. | [optional] 
**Includes** | Pointer to **interface{}** | List of related resources that should be included in the export (redundant when &#39;fields&#39; are specified). | [optional] 
**Fields** | Pointer to **interface{}** | List of fields to export for the main and related resources (automatically included). Pass the asterisk &#39;*&#39; to include all exportable fields for the main and related resources. | [optional] 
**Filters** | Pointer to **interface{}** | The filters used to select the records to be exported. | [optional] 
**DryData** | Pointer to **interface{}** | Send this attribute if you want to skip exporting redundant attributes (IDs, timestamps, blanks, etc.), useful when combining export and import to duplicate your dataset. | [optional] 
**StartedAt** | Pointer to **interface{}** | Time at which the export was started. | [optional] 
**CompletedAt** | Pointer to **interface{}** | Time at which the export was completed. | [optional] 
**InterruptedAt** | Pointer to **interface{}** | Time at which the export was interrupted. | [optional] 
**RecordsCount** | Pointer to **interface{}** | Indicates the number of records to be exported. | [optional] 
**AttachmentUrl** | Pointer to **interface{}** | The URL to the output file, which will be generated upon export completion. | [optional] 
**ErrorsLog** | Pointer to **interface{}** | Contains the exports errors, if any. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETExportsExportId200ResponseDataAttributes

`func NewGETExportsExportId200ResponseDataAttributes() *GETExportsExportId200ResponseDataAttributes`

NewGETExportsExportId200ResponseDataAttributes instantiates a new GETExportsExportId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETExportsExportId200ResponseDataAttributesWithDefaults

`func NewGETExportsExportId200ResponseDataAttributesWithDefaults() *GETExportsExportId200ResponseDataAttributes`

NewGETExportsExportId200ResponseDataAttributesWithDefaults instantiates a new GETExportsExportId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *GETExportsExportId200ResponseDataAttributes) GetResourceType() interface{}`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *GETExportsExportId200ResponseDataAttributes) GetResourceTypeOk() (*interface{}, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *GETExportsExportId200ResponseDataAttributes) SetResourceType(v interface{})`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *GETExportsExportId200ResponseDataAttributes) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *GETExportsExportId200ResponseDataAttributes) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *GETExportsExportId200ResponseDataAttributes) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetFormat

`func (o *GETExportsExportId200ResponseDataAttributes) GetFormat() interface{}`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *GETExportsExportId200ResponseDataAttributes) GetFormatOk() (*interface{}, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *GETExportsExportId200ResponseDataAttributes) SetFormat(v interface{})`

SetFormat sets Format field to given value.

### HasFormat

`func (o *GETExportsExportId200ResponseDataAttributes) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *GETExportsExportId200ResponseDataAttributes) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *GETExportsExportId200ResponseDataAttributes) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetStatus

`func (o *GETExportsExportId200ResponseDataAttributes) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETExportsExportId200ResponseDataAttributes) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETExportsExportId200ResponseDataAttributes) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETExportsExportId200ResponseDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GETExportsExportId200ResponseDataAttributes) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GETExportsExportId200ResponseDataAttributes) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetIncludes

`func (o *GETExportsExportId200ResponseDataAttributes) GetIncludes() interface{}`

GetIncludes returns the Includes field if non-nil, zero value otherwise.

### GetIncludesOk

`func (o *GETExportsExportId200ResponseDataAttributes) GetIncludesOk() (*interface{}, bool)`

GetIncludesOk returns a tuple with the Includes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludes

`func (o *GETExportsExportId200ResponseDataAttributes) SetIncludes(v interface{})`

SetIncludes sets Includes field to given value.

### HasIncludes

`func (o *GETExportsExportId200ResponseDataAttributes) HasIncludes() bool`

HasIncludes returns a boolean if a field has been set.

### SetIncludesNil

`func (o *GETExportsExportId200ResponseDataAttributes) SetIncludesNil(b bool)`

 SetIncludesNil sets the value for Includes to be an explicit nil

### UnsetIncludes
`func (o *GETExportsExportId200ResponseDataAttributes) UnsetIncludes()`

UnsetIncludes ensures that no value is present for Includes, not even an explicit nil
### GetFields

`func (o *GETExportsExportId200ResponseDataAttributes) GetFields() interface{}`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *GETExportsExportId200ResponseDataAttributes) GetFieldsOk() (*interface{}, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *GETExportsExportId200ResponseDataAttributes) SetFields(v interface{})`

SetFields sets Fields field to given value.

### HasFields

`func (o *GETExportsExportId200ResponseDataAttributes) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *GETExportsExportId200ResponseDataAttributes) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *GETExportsExportId200ResponseDataAttributes) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil
### GetFilters

`func (o *GETExportsExportId200ResponseDataAttributes) GetFilters() interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *GETExportsExportId200ResponseDataAttributes) GetFiltersOk() (*interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *GETExportsExportId200ResponseDataAttributes) SetFilters(v interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *GETExportsExportId200ResponseDataAttributes) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *GETExportsExportId200ResponseDataAttributes) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *GETExportsExportId200ResponseDataAttributes) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetDryData

`func (o *GETExportsExportId200ResponseDataAttributes) GetDryData() interface{}`

GetDryData returns the DryData field if non-nil, zero value otherwise.

### GetDryDataOk

`func (o *GETExportsExportId200ResponseDataAttributes) GetDryDataOk() (*interface{}, bool)`

GetDryDataOk returns a tuple with the DryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryData

`func (o *GETExportsExportId200ResponseDataAttributes) SetDryData(v interface{})`

SetDryData sets DryData field to given value.

### HasDryData

`func (o *GETExportsExportId200ResponseDataAttributes) HasDryData() bool`

HasDryData returns a boolean if a field has been set.

### SetDryDataNil

`func (o *GETExportsExportId200ResponseDataAttributes) SetDryDataNil(b bool)`

 SetDryDataNil sets the value for DryData to be an explicit nil

### UnsetDryData
`func (o *GETExportsExportId200ResponseDataAttributes) UnsetDryData()`

UnsetDryData ensures that no value is present for DryData, not even an explicit nil
### GetStartedAt

`func (o *GETExportsExportId200ResponseDataAttributes) GetStartedAt() interface{}`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *GETExportsExportId200ResponseDataAttributes) GetStartedAtOk() (*interface{}, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *GETExportsExportId200ResponseDataAttributes) SetStartedAt(v interface{})`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *GETExportsExportId200ResponseDataAttributes) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *GETExportsExportId200ResponseDataAttributes) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *GETExportsExportId200ResponseDataAttributes) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *GETExportsExportId200ResponseDataAttributes) GetCompletedAt() interface{}`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GETExportsExportId200ResponseDataAttributes) GetCompletedAtOk() (*interface{}, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GETExportsExportId200ResponseDataAttributes) SetCompletedAt(v interface{})`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GETExportsExportId200ResponseDataAttributes) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *GETExportsExportId200ResponseDataAttributes) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *GETExportsExportId200ResponseDataAttributes) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetInterruptedAt

`func (o *GETExportsExportId200ResponseDataAttributes) GetInterruptedAt() interface{}`

GetInterruptedAt returns the InterruptedAt field if non-nil, zero value otherwise.

### GetInterruptedAtOk

`func (o *GETExportsExportId200ResponseDataAttributes) GetInterruptedAtOk() (*interface{}, bool)`

GetInterruptedAtOk returns a tuple with the InterruptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterruptedAt

`func (o *GETExportsExportId200ResponseDataAttributes) SetInterruptedAt(v interface{})`

SetInterruptedAt sets InterruptedAt field to given value.

### HasInterruptedAt

`func (o *GETExportsExportId200ResponseDataAttributes) HasInterruptedAt() bool`

HasInterruptedAt returns a boolean if a field has been set.

### SetInterruptedAtNil

`func (o *GETExportsExportId200ResponseDataAttributes) SetInterruptedAtNil(b bool)`

 SetInterruptedAtNil sets the value for InterruptedAt to be an explicit nil

### UnsetInterruptedAt
`func (o *GETExportsExportId200ResponseDataAttributes) UnsetInterruptedAt()`

UnsetInterruptedAt ensures that no value is present for InterruptedAt, not even an explicit nil
### GetRecordsCount

`func (o *GETExportsExportId200ResponseDataAttributes) GetRecordsCount() interface{}`

GetRecordsCount returns the RecordsCount field if non-nil, zero value otherwise.

### GetRecordsCountOk

`func (o *GETExportsExportId200ResponseDataAttributes) GetRecordsCountOk() (*interface{}, bool)`

GetRecordsCountOk returns a tuple with the RecordsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsCount

`func (o *GETExportsExportId200ResponseDataAttributes) SetRecordsCount(v interface{})`

SetRecordsCount sets RecordsCount field to given value.

### HasRecordsCount

`func (o *GETExportsExportId200ResponseDataAttributes) HasRecordsCount() bool`

HasRecordsCount returns a boolean if a field has been set.

### SetRecordsCountNil

`func (o *GETExportsExportId200ResponseDataAttributes) SetRecordsCountNil(b bool)`

 SetRecordsCountNil sets the value for RecordsCount to be an explicit nil

### UnsetRecordsCount
`func (o *GETExportsExportId200ResponseDataAttributes) UnsetRecordsCount()`

UnsetRecordsCount ensures that no value is present for RecordsCount, not even an explicit nil
### GetAttachmentUrl

`func (o *GETExportsExportId200ResponseDataAttributes) GetAttachmentUrl() interface{}`

GetAttachmentUrl returns the AttachmentUrl field if non-nil, zero value otherwise.

### GetAttachmentUrlOk

`func (o *GETExportsExportId200ResponseDataAttributes) GetAttachmentUrlOk() (*interface{}, bool)`

GetAttachmentUrlOk returns a tuple with the AttachmentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentUrl

`func (o *GETExportsExportId200ResponseDataAttributes) SetAttachmentUrl(v interface{})`

SetAttachmentUrl sets AttachmentUrl field to given value.

### HasAttachmentUrl

`func (o *GETExportsExportId200ResponseDataAttributes) HasAttachmentUrl() bool`

HasAttachmentUrl returns a boolean if a field has been set.

### SetAttachmentUrlNil

`func (o *GETExportsExportId200ResponseDataAttributes) SetAttachmentUrlNil(b bool)`

 SetAttachmentUrlNil sets the value for AttachmentUrl to be an explicit nil

### UnsetAttachmentUrl
`func (o *GETExportsExportId200ResponseDataAttributes) UnsetAttachmentUrl()`

UnsetAttachmentUrl ensures that no value is present for AttachmentUrl, not even an explicit nil
### GetErrorsLog

`func (o *GETExportsExportId200ResponseDataAttributes) GetErrorsLog() interface{}`

GetErrorsLog returns the ErrorsLog field if non-nil, zero value otherwise.

### GetErrorsLogOk

`func (o *GETExportsExportId200ResponseDataAttributes) GetErrorsLogOk() (*interface{}, bool)`

GetErrorsLogOk returns a tuple with the ErrorsLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorsLog

`func (o *GETExportsExportId200ResponseDataAttributes) SetErrorsLog(v interface{})`

SetErrorsLog sets ErrorsLog field to given value.

### HasErrorsLog

`func (o *GETExportsExportId200ResponseDataAttributes) HasErrorsLog() bool`

HasErrorsLog returns a boolean if a field has been set.

### SetErrorsLogNil

`func (o *GETExportsExportId200ResponseDataAttributes) SetErrorsLogNil(b bool)`

 SetErrorsLogNil sets the value for ErrorsLog to be an explicit nil

### UnsetErrorsLog
`func (o *GETExportsExportId200ResponseDataAttributes) UnsetErrorsLog()`

UnsetErrorsLog ensures that no value is present for ErrorsLog, not even an explicit nil
### GetCreatedAt

`func (o *GETExportsExportId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETExportsExportId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETExportsExportId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETExportsExportId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETExportsExportId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETExportsExportId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETExportsExportId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETExportsExportId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETExportsExportId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETExportsExportId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETExportsExportId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETExportsExportId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETExportsExportId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETExportsExportId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETExportsExportId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETExportsExportId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETExportsExportId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETExportsExportId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETExportsExportId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETExportsExportId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETExportsExportId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETExportsExportId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETExportsExportId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETExportsExportId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETExportsExportId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETExportsExportId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETExportsExportId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETExportsExportId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETExportsExportId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETExportsExportId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


