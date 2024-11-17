# \CurrencyExchangeApi

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CurrencyExchangeRatesCurrencyExchangeIdPut**](CurrencyExchangeApi.md#CurrencyExchangeRatesCurrencyExchangeIdPut) | **Put** /currency-exchange-rates/{currency-exchange-id} | Update a currency exchange rate
[**CurrencyExchangeRatesGet**](CurrencyExchangeApi.md#CurrencyExchangeRatesGet) | **Get** /currency-exchange-rates | List currency exchange rates
[**CurrencyExchangeRatesPost**](CurrencyExchangeApi.md#CurrencyExchangeRatesPost) | **Post** /currency-exchange-rates | Add a new currency exchange rate
[**CurrencyExchangeRatesToggleStatusCurrencyExchangeIdPut**](CurrencyExchangeApi.md#CurrencyExchangeRatesToggleStatusCurrencyExchangeIdPut) | **Put** /currency-exchange-rates/toggle-status/{currency-exchange-id} | Update the status of a currency exchange rate


# **CurrencyExchangeRatesCurrencyExchangeIdPut**
> ResponseUpdateCurrencyExchangeRateApiResponse CurrencyExchangeRatesCurrencyExchangeIdPut(ctx, currencyExchangeId, body)
Update a currency exchange rate

Updates the details of a specified currency exchange rate identified by its unique ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **currencyExchangeId** | **int32**| Unique identifier of the currency exchange rate | 
  **body** | [**HandlerUpdateCurrencyExchangeRatesRequest**](HandlerUpdateCurrencyExchangeRatesRequest.md)| Details of the currency exchange rate to update | 

### Return type

[**ResponseUpdateCurrencyExchangeRateApiResponse**](response.UpdateCurrencyExchangeRateAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CurrencyExchangeRatesGet**
> ResponseListCurrencyExchangeRateApiResponse CurrencyExchangeRatesGet(ctx, optional)
List currency exchange rates

Retrieves a paginated list of currency exchange rates based on the specified skip and limit parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CurrencyExchangeApiCurrencyExchangeRatesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CurrencyExchangeApiCurrencyExchangeRatesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**ResponseListCurrencyExchangeRateApiResponse**](response.ListCurrencyExchangeRateAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CurrencyExchangeRatesPost**
> ResponseCreateCurrencyExchangeRateApiResponse CurrencyExchangeRatesPost(ctx, body)
Add a new currency exchange rate

Creates a new currency exchange rate record with the provided details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**HandlerCreateCurrencyExchangeRateRequest**](HandlerCreateCurrencyExchangeRateRequest.md)| Request body containing the currency exchange rate details | 

### Return type

[**ResponseCreateCurrencyExchangeRateApiResponse**](response.CreateCurrencyExchangeRateAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CurrencyExchangeRatesToggleStatusCurrencyExchangeIdPut**
> ResponseUpdateCurrencyExchangeRateStatusApiResponse CurrencyExchangeRatesToggleStatusCurrencyExchangeIdPut(ctx, currencyExchangeId)
Update the status of a currency exchange rate

Updates the status of a specified currency exchange rate using its unique identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **currencyExchangeId** | **int32**| Unique identifier of the currency exchange rate | 

### Return type

[**ResponseUpdateCurrencyExchangeRateStatusApiResponse**](response.UpdateCurrencyExchangeRateStatusAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

