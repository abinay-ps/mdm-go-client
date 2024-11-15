# \PincodeMasterApi

All URIs are relative to *https://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LocalpinmasterModifylocalpincodeSourcepincodePut**](PincodeMasterApi.md#LocalpinmasterModifylocalpincodeSourcepincodePut) | **Put** /localpinmaster/modifylocalpincode/{sourcepincode} | Update Local Pin Details By Source Pincode
[**LocalpinmasterPost**](PincodeMasterApi.md#LocalpinmasterPost) | **Post** /localpinmaster/ | Create Local Pin List in One Source Pin
[**LocalpinmasterSearchlocalpincodelistGet**](PincodeMasterApi.md#LocalpinmasterSearchlocalpincodelistGet) | **Get** /localpinmaster/searchlocalpincodelist | Search all local pincode list


# **LocalpinmasterModifylocalpincodeSourcepincodePut**
> HandlerResponse LocalpinmasterModifylocalpincodeSourcepincodePut(ctx, sourcepincode, optional)
Update Local Pin Details By Source Pincode

Modifies Local Pincode details by Source Pincode

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sourcepincode** | **float32**| Source Pincode | 
 **optional** | ***PincodeMasterApiLocalpinmasterModifylocalpincodeSourcepincodePutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PincodeMasterApiLocalpinmasterModifylocalpincodeSourcepincodePutOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | [**optional.Interface of HandlerLocalPinRequest**](HandlerLocalPinRequest.md)| Local Pin Request | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocalpinmasterPost**
> HandlerResponse LocalpinmasterPost(ctx, localPin)
Create Local Pin List in One Source Pin

Creates a new Local Pincode under a Source Pincode

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **localPin** | [**HandlerLocalPinRequest**](HandlerLocalPinRequest.md)| Local Pincode Details | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocalpinmasterSearchlocalpincodelistGet**
> HandlerResponse LocalpinmasterSearchlocalpincodelistGet(ctx, sourcepincode, officeid)
Search all local pincode list

Searches all local pincode list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sourcepincode** | **float32**| Source Pincode | 
  **officeid** | **int32**| Office ID | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

