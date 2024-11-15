# \ForeignSALAMSProductApi

All URIs are relative to *https://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ForeignsalamsproductmasterForeignsalamsallproductcodedetailsCountryCodeGet**](ForeignSALAMSProductApi.md#ForeignsalamsproductmasterForeignsalamsallproductcodedetailsCountryCodeGet) | **Get** /foreignsalamsproductmaster/foreignsalamsallproductcodedetails/{CountryCode} | Get all ForeignSalamsProduct details with all Product as an array within country
[**ForeignsalamsproductmasterForeignsalamsproductcodedetailsCountryCodeProductCodeGet**](ForeignSALAMSProductApi.md#ForeignsalamsproductmasterForeignsalamsproductcodedetailsCountryCodeProductCodeGet) | **Get** /foreignsalamsproductmaster/foreignsalamsproductcodedetails/{CountryCode}/{ProductCode} | Get ForeignSalamsProduct Code by Country Code
[**ForeignsalamsproductmasterForeignsalamsproductdetailsCountryproductmappingidGet**](ForeignSALAMSProductApi.md#ForeignsalamsproductmasterForeignsalamsproductdetailsCountryproductmappingidGet) | **Get** /foreignsalamsproductmaster/foreignsalamsproductdetails/{countryproductmappingid} | Get Foreign Salams Product
[**ForeignsalamsproductmasterPost**](ForeignSALAMSProductApi.md#ForeignsalamsproductmasterPost) | **Post** /foreignsalamsproductmaster/ | Create a new ForeignSalamsProduct
[**ForeignsalamsproductmasterUpdateforeignsalamsprodcutdetailsCountryproductmappingidPut**](ForeignSALAMSProductApi.md#ForeignsalamsproductmasterUpdateforeignsalamsprodcutdetailsCountryproductmappingidPut) | **Put** /foreignsalamsproductmaster/updateforeignsalamsprodcutdetails/{countryproductmappingid} | Update a ForeignSalamsProduct


# **ForeignsalamsproductmasterForeignsalamsallproductcodedetailsCountryCodeGet**
> HandlerResponse ForeignsalamsproductmasterForeignsalamsallproductcodedetailsCountryCodeGet(ctx, countryCode)
Get all ForeignSalamsProduct details with all Product as an array within country

Fetches all the details of all ForeignSalamsProducts as an array within country by countrycode

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countryCode** | **string**| Country Code | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForeignsalamsproductmasterForeignsalamsproductcodedetailsCountryCodeProductCodeGet**
> HandlerResponse ForeignsalamsproductmasterForeignsalamsproductcodedetailsCountryCodeProductCodeGet(ctx, countryCode, productCode)
Get ForeignSalamsProduct Code by Country Code

Fetches the details of ForeignSalamsProducts with all mode of transmission by Country Code and Product Code

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countryCode** | **string**| Country Code | 
  **productCode** | **string**| Product Code | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForeignsalamsproductmasterForeignsalamsproductdetailsCountryproductmappingidGet**
> HandlerResponse ForeignsalamsproductmasterForeignsalamsproductdetailsCountryproductmappingidGet(ctx, countryproductmappingid)
Get Foreign Salams Product

Get Foreign SAL AMS Product details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countryproductmappingid** | **int32**| Country Product Mapping ID | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForeignsalamsproductmasterPost**
> HandlerResponse ForeignsalamsproductmasterPost(ctx, createInternationalProductCountryMappingRequest)
Create a new ForeignSalamsProduct

To Create a new ForeignSalamsProduct

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createInternationalProductCountryMappingRequest** | [**HandlerCreateInternationalProductCountryMappingRequest**](HandlerCreateInternationalProductCountryMappingRequest.md)| Input Field  | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForeignsalamsproductmasterUpdateforeignsalamsprodcutdetailsCountryproductmappingidPut**
> HandlerResponse ForeignsalamsproductmasterUpdateforeignsalamsprodcutdetailsCountryproductmappingidPut(ctx, countryproductmappingid)
Update a ForeignSalamsProduct

Updates Foreign SAL AMS Product details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countryproductmappingid** | **int32**| Country Product Mapping ID | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

