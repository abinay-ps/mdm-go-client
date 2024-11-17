# \CountryApi

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountriesCountryIdGet**](CountryApi.md#CountriesCountryIdGet) | **Get** /countries/{country-id} | Get country details by country id
[**CountriesCountryIdPut**](CountryApi.md#CountriesCountryIdPut) | **Put** /countries/{country-id} | Update country details
[**CountriesGet**](CountryApi.md#CountriesGet) | **Get** /countries | List all countries with pagination
[**CountriesPost**](CountryApi.md#CountriesPost) | **Post** /countries | Create a new country entry


# **CountriesCountryIdGet**
> ResponseGetCountryApiResponse CountriesCountryIdGet(ctx, countryId)
Get country details by country id

Retrieve country information based on the provided country ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countryId** | **int32**| ID of the country to retrieve | 

### Return type

[**ResponseGetCountryApiResponse**](response.GetCountryAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CountriesCountryIdPut**
> ResponseUpdateCountryApiResponse CountriesCountryIdPut(ctx, countryId, body)
Update country details

Update country information using the provided country ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countryId** | **int32**| Country ID | 
  **body** | [**HandlerUpdateCountryRequest**](HandlerUpdateCountryRequest.md)| Update Country Request | 

### Return type

[**ResponseUpdateCountryApiResponse**](response.UpdateCountryAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CountriesGet**
> ResponseListCountriesApiResponse CountriesGet(ctx, optional)
List all countries with pagination

Retrieve a list of countries with pagination options.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CountryApiCountriesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CountryApiCountriesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**ResponseListCountriesApiResponse**](response.ListCountriesAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CountriesPost**
> ResponseCreateCountryApiResponse CountriesPost(ctx, body)
Create a new country entry

Add a new country with its details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**HandlerCreateCountryRequest**](HandlerCreateCountryRequest.md)| Create Country Request | 

### Return type

[**ResponseCreateCountryApiResponse**](response.CreateCountryAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

