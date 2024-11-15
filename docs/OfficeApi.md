# \OfficeApi

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OfficesBoDetailsGet**](OfficeApi.md#OfficesBoDetailsGet) | **Get** /offices/bo-details | Retrieve BO details by reporting office ID
[**OfficesGetBoDetailsOfficetypeReportingofficeidGet**](OfficeApi.md#OfficesGetBoDetailsOfficetypeReportingofficeidGet) | **Get** /offices/get-bo-details/{officetype}/{reportingofficeid} | Retrieve BO details by reporting office ID
[**OfficesGetOfficeDetailsByOfficeIdOfficeidGet**](OfficeApi.md#OfficesGetOfficeDetailsByOfficeIdOfficeidGet) | **Get** /offices/get-office-details-by-office-id/{officeid} | Retrieve office details by officeID
[**OfficesGetOfficeDetailsByOfficeNameOfficenameGet**](OfficeApi.md#OfficesGetOfficeDetailsByOfficeNameOfficenameGet) | **Get** /offices/get-office-details-by-office-name/{officename} | Retrieve office details by officeID
[**OfficesGetOfficeDetailsByPincodePincodeGet**](OfficeApi.md#OfficesGetOfficeDetailsByPincodePincodeGet) | **Get** /offices/get-office-details-by-pincode/{pincode} | Retrieve office details by Pincode
[**OfficesOfficeDetailsByOfficeIdGet**](OfficeApi.md#OfficesOfficeDetailsByOfficeIdGet) | **Get** /offices/office-details-by-office-id | Retrieve office details by officeID
[**OfficesOfficeDetailsByOfficeNameGet**](OfficeApi.md#OfficesOfficeDetailsByOfficeNameGet) | **Get** /offices/office-details-by-office-name | Retrieve office details by officeID
[**OfficesOfficeDetailsByPincodeGet**](OfficeApi.md#OfficesOfficeDetailsByPincodeGet) | **Get** /offices/office-details-by-pincode | Retrieve office details by Pincode


# **OfficesBoDetailsGet**
> []HandlerOfficeDataResponse OfficesBoDetailsGet(ctx, officetype, reportingofficeid)
Retrieve BO details by reporting office ID

Retrieves details of branch offices associated with a specific reporting office ID and office type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officetype** | **string**| Office Type | 
  **reportingofficeid** | **string**| Reporting Office ID | 

### Return type

[**[]HandlerOfficeDataResponse**](handler.OfficeDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficesGetBoDetailsOfficetypeReportingofficeidGet**
> []HandlerOfficeDataResponse OfficesGetBoDetailsOfficetypeReportingofficeidGet(ctx, officetype, reportingofficeid)
Retrieve BO details by reporting office ID

Retrieves details of branch offices associated with a specific reporting office ID and office type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officetype** | **string**| Office Type | 
  **reportingofficeid** | **string**| Reporting Office ID | 

### Return type

[**[]HandlerOfficeDataResponse**](handler.OfficeDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficesGetOfficeDetailsByOfficeIdOfficeidGet**
> []HandlerOfficesResponse OfficesGetOfficeDetailsByOfficeIdOfficeidGet(ctx, officeid)
Retrieve office details by officeID

Retrieves details of office associated with a specific officeID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeid** | **string**| Office ID | 

### Return type

[**[]HandlerOfficesResponse**](handler.OfficesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficesGetOfficeDetailsByOfficeNameOfficenameGet**
> []HandlerOfficesResponse OfficesGetOfficeDetailsByOfficeNameOfficenameGet(ctx, officename)
Retrieve office details by officeID

Retrieves details of office associated with a specific officeID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officename** | **string**| Office Name | 

### Return type

[**[]HandlerOfficesResponse**](handler.OfficesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficesGetOfficeDetailsByPincodePincodeGet**
> []HandlerOfficesResponse OfficesGetOfficeDetailsByPincodePincodeGet(ctx, pincode)
Retrieve office details by Pincode

Retrieves details of offices associated with a specific Pincode.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pincode** | **string**| Pincode | 

### Return type

[**[]HandlerOfficesResponse**](handler.OfficesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficesOfficeDetailsByOfficeIdGet**
> []HandlerOfficesResponse OfficesOfficeDetailsByOfficeIdGet(ctx, officeid)
Retrieve office details by officeID

Retrieves details of office associated with a specific officeID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officeid** | **string**| Office ID | 

### Return type

[**[]HandlerOfficesResponse**](handler.OfficesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficesOfficeDetailsByOfficeNameGet**
> []HandlerOfficesResponse OfficesOfficeDetailsByOfficeNameGet(ctx, officename)
Retrieve office details by officeID

Retrieves details of office associated with a specific officeID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **officename** | **string**| Office Name | 

### Return type

[**[]HandlerOfficesResponse**](handler.OfficesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfficesOfficeDetailsByPincodeGet**
> []HandlerOfficesResponse OfficesOfficeDetailsByPincodeGet(ctx, pincode)
Retrieve office details by Pincode

Retrieves details of offices associated with a specific Pincode.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pincode** | **string**| Pincode | 

### Return type

[**[]HandlerOfficesResponse**](handler.OfficesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

