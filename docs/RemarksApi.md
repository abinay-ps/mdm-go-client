# \RemarksApi

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StandardRemarksGet**](RemarksApi.md#StandardRemarksGet) | **Get** /standard-remarks | Get Standard Remarks


# **StandardRemarksGet**
> ResponseListStandardRemarksApiResponse StandardRemarksGet(ctx, remarksType, optional)
Get Standard Remarks

Retrieves a list of standard remarks based on the remarks type (Article or EMO) and whether it's a branch office.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **remarksType** | **string**| Type of remarks | 
 **optional** | ***RemarksApiStandardRemarksGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RemarksApiStandardRemarksGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isBranchOffice** | **optional.Bool**| Indicates whether the request is for a branch office | 
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**ResponseListStandardRemarksApiResponse**](response.ListStandardRemarksAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

