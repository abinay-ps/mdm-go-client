# \PinocodeApi

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PincodeMappingsGet**](PinocodeApi.md#PincodeMappingsGet) | **Get** /pincode-mappings | List local pincodes
[**PincodeMappingsLocalPinIdPut**](PinocodeApi.md#PincodeMappingsLocalPinIdPut) | **Put** /pincode-mappings/{local-pin-id} | Update local pin codes
[**PincodeMappingsPost**](PinocodeApi.md#PincodeMappingsPost) | **Post** /pincode-mappings | Create multiple local pincodes


# **PincodeMappingsGet**
> ResponseListLocalPincodesApiResponse PincodeMappingsGet(ctx, optional)
List local pincodes

Retrieves a list of local pincodes based on the provided source pincode or office ID with pagination.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PinocodeApiPincodeMappingsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PinocodeApiPincodeMappingsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**|  | 
 **officeId** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 
 **sourcePincode** | **optional.Float32**|  | 

### Return type

[**ResponseListLocalPincodesApiResponse**](response.ListLocalPincodesAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PincodeMappingsLocalPinIdPut**
> ResponseUpdatePincodesApiResponse PincodeMappingsLocalPinIdPut(ctx, localPinId, body)
Update local pin codes

Updates the details of a local pin code based on the provided local pin ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **localPinId** | **int32**| Local pin ID | 
  **body** | [**HandlerUpdatePincodesRequest**](HandlerUpdatePincodesRequest.md)| Request body containing updated pin code details | 

### Return type

[**ResponseUpdatePincodesApiResponse**](response.UpdatePincodesAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PincodeMappingsPost**
> ResponseCreatePincodesApiResponse PincodeMappingsPost(ctx, requestBody)
Create multiple local pincodes

Creates multiple local pincodes based on the provided JSON payload, which contains an array of source and local pincodes with optional flags.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestBody** | [**HandlerCreateLocalPincodesRequest**](HandlerCreateLocalPincodesRequest.md)| Request body containing an array of local pincodes | 

### Return type

[**ResponseCreatePincodesApiResponse**](response.CreatePincodesAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

