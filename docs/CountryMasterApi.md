# \CountryMasterApi

All URIs are relative to *https://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountrymasterApiAddcountriesPost**](CountryMasterApi.md#CountrymasterApiAddcountriesPost) | **Post** /countrymaster/api/addcountries | Add Countries
[**CountrymasterCountrycodeModifyCountriesCountrycodePut**](CountryMasterApi.md#CountrymasterCountrycodeModifyCountriesCountrycodePut) | **Put** /countrymaster/countrycode/modifyCountries/{countrycode} | Modify Country by Country Code
[**CountrymasterGetAllCountriesCountrycodeGet**](CountryMasterApi.md#CountrymasterGetAllCountriesCountrycodeGet) | **Get** /countrymaster/getAllCountries/{countrycode} | Get All Countries by Country Code
[**CountrymasterGetAllCountriesGet**](CountryMasterApi.md#CountrymasterGetAllCountriesGet) | **Get** /countrymaster/getAllCountries/ | Get All Countries


# **CountrymasterApiAddcountriesPost**
> HandlerResponse CountrymasterApiAddcountriesPost(ctx, addCountryRequest)
Add Countries

Add a new Country

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **addCountryRequest** | [**HandlerAddCountryRequest**](HandlerAddCountryRequest.md)| Add New Country Request  | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CountrymasterCountrycodeModifyCountriesCountrycodePut**
> HandlerResponse CountrymasterCountrycodeModifyCountriesCountrycodePut(ctx, countrycode)
Modify Country by Country Code

Modify a Country by Country Code

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countrycode** | **string**| Country Code | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CountrymasterGetAllCountriesCountrycodeGet**
> HandlerResponse CountrymasterGetAllCountriesCountrycodeGet(ctx, countrycode)
Get All Countries by Country Code

Fetches a Country Details by Country Code

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countrycode** | **string**| Country Code | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CountrymasterGetAllCountriesGet**
> HandlerResponse CountrymasterGetAllCountriesGet(ctx, )
Get All Countries

Fetches all Countries

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

