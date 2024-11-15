# \OfficeShiftsApi

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OfficeShiftsBeatShiftDetailsGet**](OfficeShiftsApi.md#OfficeShiftsBeatShiftDetailsGet) | **Get** /office-shifts/beat-shift-details | Retrieves the details for the beat end.
[**OfficeShiftsForceShiftEndStatisticsGet**](OfficeShiftsApi.md#OfficeShiftsForceShiftEndStatisticsGet) | **Get** /office-shifts/force-shift-end-statistics | Retrieves statistics for the end of a force shift.
[**OfficeShiftsOfficeIdBeatPut**](OfficeShiftsApi.md#OfficeShiftsOfficeIdBeatPut) | **Put** /office-shifts/{office-id}/beat | Update Beat End of Day Status
[**OfficeShiftsOfficeIdRevertPut**](OfficeShiftsApi.md#OfficeShiftsOfficeIdRevertPut) | **Put** /office-shifts/{office-id}/revert | Reverts the shift based on the provided reversion type.
[**OfficeShiftsOfficeIdShiftStatusGet**](OfficeShiftsApi.md#OfficeShiftsOfficeIdShiftStatusGet) | **Get** /office-shifts/{office-id}/shift-status | Checks the status of the shift for the specified office based on state.
[**OfficeShiftsPost**](OfficeShiftsApi.md#OfficeShiftsPost) | **Post** /office-shifts | Create Office Shift Record
[**OfficeShiftsShiftEndStatisticsGet**](OfficeShiftsApi.md#OfficeShiftsShiftEndStatisticsGet) | **Get** /office-shifts/shift-end-statistics | Retrieves statistics for the end of a shift.


# **OfficeShiftsBeatShiftDetailsGet**
> []ResponseListOfficeShiftsBeatsToBeatEndApiResponse OfficeShiftsBeatShiftDetailsGet(ctx, officeId, optional)
Retrieves the details for the beat end.

Gets beat end details based on office ID, batch name, and beat name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeId** | **int32**| Office ID | 
 **optional** | ***OfficeShiftsApiOfficeShiftsBeatShiftDetailsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OfficeShiftsApiOfficeShiftsBeatShiftDetailsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchName** | **optional.String**| Batch Name | 
 **beatName** | **optional.String**| Beat Name | 
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**[]ResponseListOfficeShiftsBeatsToBeatEndApiResponse**](response.ListOfficeShiftsBeatsToBeatEndAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficeShiftsForceShiftEndStatisticsGet**
> []ResponseListOfficeShiftsForceShiftEndStatisticsApiResponse OfficeShiftsForceShiftEndStatisticsGet(ctx, officeId, date, optional)
Retrieves statistics for the end of a force shift.

Gets shift end statistics based on office ID, batch name, beat name, and date.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeId** | **int32**| Office ID | 
  **date** | **string**| Date | 
 **optional** | ***OfficeShiftsApiOfficeShiftsForceShiftEndStatisticsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OfficeShiftsApiOfficeShiftsForceShiftEndStatisticsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchName** | **optional.String**| Batch Name | 
 **beatName** | **optional.String**| Beat Name | 
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**[]ResponseListOfficeShiftsForceShiftEndStatisticsApiResponse**](response.ListOfficeShiftsForceShiftEndStatisticsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficeShiftsOfficeIdBeatPut**
> string OfficeShiftsOfficeIdBeatPut(ctx, officeId, batchName, beatName, optional)
Update Beat End of Day Status

Updates the end of day status for office shifts based on the provided office ID, batch name, and beat name. Optionally forces the shift to end.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeId** | **int32**| Office ID | 
  **batchName** | **string**| Batch Name | 
  **beatName** | **string**| Beat Name | 
 **optional** | ***OfficeShiftsApiOfficeShiftsOfficeIdBeatPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OfficeShiftsApiOfficeShiftsOfficeIdBeatPutOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **forceShiftEnd** | **optional.Bool**| Force Shift End | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficeShiftsOfficeIdRevertPut**
> string OfficeShiftsOfficeIdRevertPut(ctx, reversionType, officeId, body)
Reverts the shift based on the provided reversion type.

Reverts the shift either at the beginning or end based on office ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reversionType** | **string**| ReversionType | 
  **officeId** | **int32**| Office ID | 
  **body** | [**HandlerUpdateOfficeShiftsRevertRequest**](HandlerUpdateOfficeShiftsRevertRequest.md)| Shift Reversion Request | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficeShiftsOfficeIdShiftStatusGet**
> ResponseFetchOfficeShiftsStatusApiResponse OfficeShiftsOfficeIdShiftStatusGet(ctx, officeId, shiftState, optional)
Checks the status of the shift for the specified office based on state.

Retrieves the shift status based on office ID, optional shift start date, and shift state.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeId** | **int32**| Office ID | 
  **shiftState** | **string**| Shift State (current, last) | 
 **optional** | ***OfficeShiftsApiOfficeShiftsOfficeIdShiftStatusGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OfficeShiftsApiOfficeShiftsOfficeIdShiftStatusGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **shiftStartDate** | **optional.String**| Shift Start Date | 

### Return type

[**ResponseFetchOfficeShiftsStatusApiResponse**](response.FetchOfficeShiftsStatusAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficeShiftsPost**
> string OfficeShiftsPost(ctx, shift)
Create Office Shift Record

Creates a new office shift record based on the provided details. Supports beginning, ending, or force-ending a shift.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shift** | [**HandlerCreateOfficeShiftsRequest**](HandlerCreateOfficeShiftsRequest.md)| Shift Details | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficeShiftsShiftEndStatisticsGet**
> []ResponseListOfficeShiftsShiftEndStatisticsApiResponse OfficeShiftsShiftEndStatisticsGet(ctx, officeId, date, optional)
Retrieves statistics for the end of a shift.

Gets shift end statistics based on office ID, date, batch name, and beat name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeId** | **int32**| Office ID | 
  **date** | **string**| Date | 
 **optional** | ***OfficeShiftsApiOfficeShiftsShiftEndStatisticsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OfficeShiftsApiOfficeShiftsShiftEndStatisticsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchName** | **optional.String**| Batch Name | 
 **beatName** | **optional.String**| Beat Name | 

### Return type

[**[]ResponseListOfficeShiftsShiftEndStatisticsApiResponse**](response.ListOfficeShiftsShiftEndStatisticsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

