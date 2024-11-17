# \OfficeApi

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OfficesCirclesGet**](OfficeApi.md#OfficesCirclesGet) | **Get** /offices/circles | List all circles
[**OfficesDdoDetailsGet**](OfficeApi.md#OfficesDdoDetailsGet) | **Get** /offices/ddo-details | List DDO office details based on Admin Office ID
[**OfficesDetailsGet**](OfficeApi.md#OfficesDetailsGet) | **Get** /offices/details | List details of multiple offices
[**OfficesDivisionsGet**](OfficeApi.md#OfficesDivisionsGet) | **Get** /offices/divisions | List divisions by Region ID or Region Name
[**OfficesGet**](OfficeApi.md#OfficesGet) | **Get** /offices | List offices under a specified office ID
[**OfficesLimitedDetailsGet**](OfficeApi.md#OfficesLimitedDetailsGet) | **Get** /offices/limited-details | List partial office details based on Pincode or Office Name
[**OfficesMergePost**](OfficeApi.md#OfficesMergePost) | **Post** /offices/merge | Merges two offices by updating their details
[**OfficesOfficeIdApprovePut**](OfficeApi.md#OfficesOfficeIdApprovePut) | **Put** /offices/{office-id}/approve | Updates the approval status of an office
[**OfficesOfficeIdChangeStatusPut**](OfficeApi.md#OfficesOfficeIdChangeStatusPut) | **Put** /offices/{office-id}/change-status | Updates the active status of an office
[**OfficesOfficeIdDetailsGet**](OfficeApi.md#OfficesOfficeIdDetailsGet) | **Get** /offices/{office-id}/details | Get details of a specific office
[**OfficesOfficeIdPartialDetailsGet**](OfficeApi.md#OfficesOfficeIdPartialDetailsGet) | **Get** /offices/{office-id}/partial-details | Get partial office details by Office ID
[**OfficesOfficeIdPut**](OfficeApi.md#OfficesOfficeIdPut) | **Put** /offices/{office-id} | Update office details by office ID
[**OfficesPaoCodesGet**](OfficeApi.md#OfficesPaoCodesGet) | **Get** /offices/pao-codes | List the distinct PAO codes
[**OfficesPartialDetailsGet**](OfficeApi.md#OfficesPartialDetailsGet) | **Get** /offices/partial-details | List partial office details based on various office IDs
[**OfficesPost**](OfficeApi.md#OfficesPost) | **Post** /offices | Create a new office
[**OfficesRegionsGet**](OfficeApi.md#OfficesRegionsGet) | **Get** /offices/regions | List regions by Circle ID or Circle Name
[**OfficesStatusGet**](OfficeApi.md#OfficesStatusGet) | **Get** /offices/status | List the status of multiple offices
[**OfficesSubDivisionsGet**](OfficeApi.md#OfficesSubDivisionsGet) | **Get** /offices/sub-divisions | List subdivisions by Division ID or Division Name
[**OfficesWorkingHoursGet**](OfficeApi.md#OfficesWorkingHoursGet) | **Get** /offices/working-hours | List the working hours of multiple offices
[**V1OfficesValidationWithPincodeGet**](OfficeApi.md#V1OfficesValidationWithPincodeGet) | **Get** /v1/offices/validation-with-pincode | Validate office ID with pin code and delivery flag


# **OfficesCirclesGet**
> ResponseListCirclesApiResponse OfficesCirclesGet(ctx, optional)
List all circles

Retrieves a list of circles with pagination support.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OfficeApiOfficesCirclesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OfficeApiOfficesCirclesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**ResponseListCirclesApiResponse**](response.ListCirclesAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficesDdoDetailsGet**
> ResponseListDdoDetailsApiResponse OfficesDdoDetailsGet(ctx, adminOfficeId, optional)
List DDO office details based on Admin Office ID

Retrieves DDO office details using Admin Office ID. First checks Redis cache for existing data. If not found, queries the service layer. Supports pagination.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **adminOfficeId** | **int32**|  | 
 **optional** | ***OfficeApiOfficesDdoDetailsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OfficeApiOfficesDdoDetailsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**ResponseListDdoDetailsApiResponse**](response.ListDDODetailsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficesDetailsGet**
> ResponseListOfficeDetailsApiResponse OfficesDetailsGet(ctx, optional)
List details of multiple offices

Retrieves details of offices based on various filters such as DivisionOfficeID, ReportingOfficeID, AccountingOfficeID, AdminOfficeID, location coordinates, office name, pincode, and state/district names. Optionally, include a ModifiedFlag to filter offices(verification) that have been modified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OfficeApiOfficesDetailsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OfficeApiOfficesDetailsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountingOfficeId** | **optional.Int32**|  | 
 **adminOfficeId** | **optional.Int32**|  | 
 **districtName** | **optional.String**|  | 
 **divisionOfficeId** | **optional.Int32**|  | 
 **latitude** | **optional.Float32**|  | 
 **limit** | **optional.Int32**|  | 
 **longitude** | **optional.Float32**|  | 
 **modifiedFlag** | **optional.Bool**|  | 
 **officeName** | **optional.String**|  | 
 **orderBy** | **optional.String**|  | 
 **pincode** | **optional.Int32**|  | 
 **reportingOfficeId** | **optional.Int32**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 
 **stateName** | **optional.String**|  | 

### Return type

[**ResponseListOfficeDetailsApiResponse**](response.ListOfficeDetailsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficesDivisionsGet**
> ResponseListDivisionsApiResponse OfficesDivisionsGet(ctx, optional)
List divisions by Region ID or Region Name

Retrieves a list of divisions based on Region ID or Region Name with pagination support

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OfficeApiOfficesDivisionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OfficeApiOfficesDivisionsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **regionId** | **optional.Int32**|  | 
 **regionName** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**ResponseListDivisionsApiResponse**](response.ListDivisionsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficesGet**
> ResponseListOfficesApiResponse OfficesGet(ctx, optional)
List offices under a specified office ID

Retrieves a list of offices under a given office ID. The endpoint checks for various office IDs and fetches the corresponding office details from Redis or the service layer if not cached. Supports pagination.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OfficeApiOfficesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OfficeApiOfficesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountingOfficeId** | **optional.String**|  | 
 **circleOfficeId** | **optional.String**|  | 
 **divisionOfficeId** | **optional.String**|  | 
 **hoId** | **optional.String**|  | 
 **hroId** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **regionOfficeId** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **soId** | **optional.String**|  | 
 **sortType** | **optional.String**|  | 
 **sroId** | **optional.String**|  | 
 **subDivisionOfficeId** | **optional.String**|  | 

### Return type

[**ResponseListOfficesApiResponse**](response.ListOfficesAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficesLimitedDetailsGet**
> ResponseListLimitedOfficeDetailsApiResponse OfficesLimitedDetailsGet(ctx, officeType, optional)
List partial office details based on Pincode or Office Name

Retrieves partial office details based on a given Pincode (3 to 6 digits) or Office Name (minimum 3 characters). The function first attempts to fetch data from Redis cache, and if not found, it queries the service layer. Supports pagination and allows filtering by office type (specify 'all' for all offices or 'post' for post offices only).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeType** | **string**| OfficeType specifies whether to fetch all offices or only post offices. | 
 **optional** | ***OfficeApiOfficesLimitedDetailsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OfficeApiOfficesLimitedDetailsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | 
 **officeName** | **optional.String**| OfficeName is a partial office name with a minimum length of 3 characters. | 
 **orderBy** | **optional.String**|  | 
 **pincode** | **optional.Int32**| Pincode is a partial pincode with a minimum length of 3 digits and a maximum length of 6 digits. | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**ResponseListLimitedOfficeDetailsApiResponse**](response.ListLimitedOfficeDetailsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficesMergePost**
> ResponseUpdateMergeOfficesApiResponse OfficesMergePost(ctx, sourceOfficeId, targetOfficeId, body)
Merges two offices by updating their details

Merges the source and target offices specified by their IDs. Updates both offices with the provided details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sourceOfficeId** | **int32**| ID of the source office to update | 
  **targetOfficeId** | **int32**| ID of the target office to update | 
  **body** | [**HandlerUpdateMergeOfficesRequest**](HandlerUpdateMergeOfficesRequest.md)| Details for merging offices | 

### Return type

[**ResponseUpdateMergeOfficesApiResponse**](response.UpdateMergeOfficesAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficesOfficeIdApprovePut**
> ResponseUpdateOfficeApprovalApiResponse OfficesOfficeIdApprovePut(ctx, officeId, body)
Updates the approval status of an office

Updates the approval information of an office specified by its ID. Requires the office ID in the URI and details in the JSON body.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeId** | **int32**| ID of the office to update | 
  **body** | [**HandlerUpdateOfficeApprovalRequest**](HandlerUpdateOfficeApprovalRequest.md)| New approval details for the office | 

### Return type

[**ResponseUpdateOfficeApprovalApiResponse**](response.UpdateOfficeApprovalAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficesOfficeIdChangeStatusPut**
> ResponseUpdateOfficeActiveStatusApiResponse OfficesOfficeIdChangeStatusPut(ctx, officeId, body)
Updates the active status of an office

Updates the status and related information of an office specified by its ID. Requires the office ID in the URI and details in the JSON body.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeId** | **int32**| ID of the office to update | 
  **body** | [**HandlerUpdateOfficeActiveStatusRequest**](HandlerUpdateOfficeActiveStatusRequest.md)| New active status details for the office | 

### Return type

[**ResponseUpdateOfficeActiveStatusApiResponse**](response.UpdateOfficeActiveStatusAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficesOfficeIdDetailsGet**
> ResponseGetOfficeDetailsApiResponse OfficesOfficeIdDetailsGet(ctx, officeId, optional)
Get details of a specific office

Retrieves details of an office by its OfficeID, with optional filters like office status, user office ID, and modified flag.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeId** | **int32**| Office ID | 
 **optional** | ***OfficeApiOfficesOfficeIdDetailsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OfficeApiOfficesOfficeIdDetailsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userOfficeId** | **optional.Int32**| User Office ID | 
 **modifiedFlag** | **optional.Bool**| Modified Flag | 
 **officeStatus** | **optional.String**| Office Status | 

### Return type

[**ResponseGetOfficeDetailsApiResponse**](response.GetOfficeDetailsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficesOfficeIdPartialDetailsGet**
> ResponseListPartialOfficeDetailsApiResponse OfficesOfficeIdPartialDetailsGet(ctx, officeId)
Get partial office details by Office ID

Retrieves partial office details for a specified office ID, checking Redis cache first. If the data is not cached, it fetches from the service layer. This endpoint supports partial retrieval of office details based on various IDs like reporting, accounting, admin, or DDO office IDs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeId** | **int32**| ID of the office for partial details | 

### Return type

[**ResponseListPartialOfficeDetailsApiResponse**](response.ListPartialOfficeDetailsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficesOfficeIdPut**
> ResponseUpdateOfficeDetailsByOfficeIdapiResponse OfficesOfficeIdPut(ctx, officeId)
Update office details by office ID

Updates the details of an office specified by its ID. Requires a JSON body with the fields to be updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeId** | **int32**| ID of the office to update | 

### Return type

[**ResponseUpdateOfficeDetailsByOfficeIdapiResponse**](response.UpdateOfficeDetailsByOfficeIDAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficesPaoCodesGet**
> ResponseListPaoCodesApiResponse OfficesPaoCodesGet(ctx, optional)
List the distinct PAO codes

Retrieves distinct PAO codes with pagination support. It attempts to fetch the codes from Redis first; if not available, it retrieves them from the service layer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OfficeApiOfficesPaoCodesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OfficeApiOfficesPaoCodesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**ResponseListPaoCodesApiResponse**](response.ListPAOCodesAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficesPartialDetailsGet**
> ResponseListPartialOfficeDetailsApiResponse OfficesPartialDetailsGet(ctx, optional)
List partial office details based on various office IDs

Retrieves partial office details based on reporting, accounting, admin, or DDO office IDs. The function first attempts to fetch data from Redis cache, and if not found, it queries the service layer. Supports pagination and allows filtering by office name (minimum 3 characters) or pincode (between 3 to 6 digits).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OfficeApiOfficesPartialDetailsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OfficeApiOfficesPartialDetailsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountingOfficeId** | **optional.Int32**|  | 
 **adminOfficeId** | **optional.Int32**|  | 
 **ddoOfficeId** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **reportingOfficeId** | **optional.Int32**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**ResponseListPartialOfficeDetailsApiResponse**](response.ListPartialOfficeDetailsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficesPost**
> ResponseCreateOfficeApiResponse OfficesPost(ctx, body)
Create a new office

Add a new office to the system with office details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**HandlerCreateOfficeRequest**](HandlerCreateOfficeRequest.md)| Create Office Request | 

### Return type

[**ResponseCreateOfficeApiResponse**](response.CreateOfficeAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficesRegionsGet**
> ResponseListRegionsApiResponse OfficesRegionsGet(ctx, optional)
List regions by Circle ID or Circle Name

Retrieves a list of regions based on Circle ID or Circle Name with pagination support

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OfficeApiOfficesRegionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OfficeApiOfficesRegionsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **circleId** | **optional.Int32**|  | 
 **circleName** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**ResponseListRegionsApiResponse**](response.ListRegionsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficesStatusGet**
> ResponseListOfficeStatusApiResponse OfficesStatusGet(ctx, optional)
List the status of multiple offices

Retrieves the status of offices with pagination support.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OfficeApiOfficesStatusGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OfficeApiOfficesStatusGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**ResponseListOfficeStatusApiResponse**](response.ListOfficeStatusAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficesSubDivisionsGet**
> ResponseListSubDivisionsApiResponse OfficesSubDivisionsGet(ctx, optional)
List subdivisions by Division ID or Division Name

Retrieves a list of subdivisions based on Division ID or Division Name with pagination support

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OfficeApiOfficesSubDivisionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OfficeApiOfficesSubDivisionsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **divisionId** | **optional.Int32**|  | 
 **divisionName** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**ResponseListSubDivisionsApiResponse**](response.ListSubDivisionsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficesWorkingHoursGet**
> ResponseListWorkingHoursApiResponse OfficesWorkingHoursGet(ctx, optional)
List the working hours of multiple offices

Retrieves the working hours of offices with pagination support.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OfficeApiOfficesWorkingHoursGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OfficeApiOfficesWorkingHoursGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**ResponseListWorkingHoursApiResponse**](response.ListWorkingHoursAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1OfficesValidationWithPincodeGet**
> ResponseGetOfficeIdValidationApiResponse V1OfficesValidationWithPincodeGet(ctx, deliveryFlag, pincode)
Validate office ID with pin code and delivery flag

Validates the office ID based on the provided pin code and delivery flag.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deliveryFlag** | **bool**|  | 
  **pincode** | **int32**|  | 

### Return type

[**ResponseGetOfficeIdValidationApiResponse**](response.GetOfficeIDValidationAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

