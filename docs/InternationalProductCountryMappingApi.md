# \InternationalProductCountryMappingApi

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InternationalProductCountryMappingsCountryProductMappingIdGet**](InternationalProductCountryMappingApi.md#InternationalProductCountryMappingsCountryProductMappingIdGet) | **Get** /international-product-country-mappings/{country-product-mapping-id} | Retrieve an international product country mapping by its ID
[**InternationalProductCountryMappingsCountryProductMappingIdPut**](InternationalProductCountryMappingApi.md#InternationalProductCountryMappingsCountryProductMappingIdPut) | **Put** /international-product-country-mappings/{country-product-mapping-id} | Update an international product country mapping
[**InternationalProductCountryMappingsGet**](InternationalProductCountryMappingApi.md#InternationalProductCountryMappingsGet) | **Get** /international-product-country-mappings | List international product country mappings based on query parameters
[**InternationalProductCountryMappingsPost**](InternationalProductCountryMappingApi.md#InternationalProductCountryMappingsPost) | **Post** /international-product-country-mappings | Create a new international product country mapping


# **InternationalProductCountryMappingsCountryProductMappingIdGet**
> ResponseGetIntProductCountryMappingApiResponse InternationalProductCountryMappingsCountryProductMappingIdGet(ctx, countryProductMappingId)
Retrieve an international product country mapping by its ID

Retrieves the details of an international product country mapping based on the country product mapping ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countryProductMappingId** | **int32**| Country Product Mapping ID | 

### Return type

[**ResponseGetIntProductCountryMappingApiResponse**](response.GetIntProductCountryMappingAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InternationalProductCountryMappingsCountryProductMappingIdPut**
> ResponseUpdateIntProductCountryMappingApiResponse InternationalProductCountryMappingsCountryProductMappingIdPut(ctx, countryProductMappingId, requestBody)
Update an international product country mapping

Updates the details of an international product country mapping using its country product mapping ID and provided details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countryProductMappingId** | **int32**| Country Product Mapping ID | 
  **requestBody** | [**HandlerUpdateIntProductCountryMappingRequest**](HandlerUpdateIntProductCountryMappingRequest.md)| Request body containing product details | 

### Return type

[**ResponseUpdateIntProductCountryMappingApiResponse**](response.UpdateIntProductCountryMappingAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InternationalProductCountryMappingsGet**
> ResponseListIntProductCountryMappingApiResponse InternationalProductCountryMappingsGet(ctx, optional)
List international product country mappings based on query parameters

Retrieves a list of international product country mappings filtered by a combination of optional parameters, including: 1. Product Code 2. Country Code 3. Country Code and Product Code 4. Country Code and Article Type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InternationalProductCountryMappingApiInternationalProductCountryMappingsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InternationalProductCountryMappingApiInternationalProductCountryMappingsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **articleType** | **optional.String**|  | 
 **countryCode** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **productCode** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**ResponseListIntProductCountryMappingApiResponse**](response.ListIntProductCountryMappingAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InternationalProductCountryMappingsPost**
> ResponseCreateIntProductCountryMappingApiResponse InternationalProductCountryMappingsPost(ctx, body)
Create a new international product country mapping

Creates a new product in the international product country mapping.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**HandlerCreateIntProductCountryMappingRequest**](HandlerCreateIntProductCountryMappingRequest.md)| Create International Product Country Mapping Request | 

### Return type

[**ResponseCreateIntProductCountryMappingApiResponse**](response.CreateIntProductCountryMappingAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

