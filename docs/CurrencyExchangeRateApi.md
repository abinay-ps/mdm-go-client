# \CurrencyExchangeRateApi

All URIs are relative to *https://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CurrencyexchangeAddCurrencyExchangeRatePost**](CurrencyExchangeRateApi.md#CurrencyexchangeAddCurrencyExchangeRatePost) | **Post** /currencyexchange/add_currency_exchange_rate | Add Currency Exchange Rate
[**CurrencyexchangeGetCurrencyExchangeRatesGet**](CurrencyExchangeRateApi.md#CurrencyexchangeGetCurrencyExchangeRatesGet) | **Get** /currencyexchange/get_currency_exchange_rates | Get all Currency Exchange Rates
[**CurrencyexchangeModifyCurrencyExchangeRatePut**](CurrencyExchangeRateApi.md#CurrencyexchangeModifyCurrencyExchangeRatePut) | **Put** /currencyexchange/modify_currency_exchange_rate | Modify Currency Exchange Rate
[**CurrencyexchangeStatustogglerGet**](CurrencyExchangeRateApi.md#CurrencyexchangeStatustogglerGet) | **Get** /currencyexchange/statustoggler | Currency Exchange Rate Status Toggler


# **CurrencyexchangeAddCurrencyExchangeRatePost**
> HandlerResponse CurrencyexchangeAddCurrencyExchangeRatePost(ctx, currencyExchangeRate)
Add Currency Exchange Rate

Creates a new Currency Exchange Rate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **currencyExchangeRate** | [**HandlerCreateCurrencyExchangeRate**](HandlerCreateCurrencyExchangeRate.md)| Currency Exchange Rate details | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CurrencyexchangeGetCurrencyExchangeRatesGet**
> HandlerResponse CurrencyexchangeGetCurrencyExchangeRatesGet(ctx, )
Get all Currency Exchange Rates

Fetches all Currency Exchange Rates

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

# **CurrencyexchangeModifyCurrencyExchangeRatePut**
> HandlerResponse CurrencyexchangeModifyCurrencyExchangeRatePut(ctx, currencyExchangeRate)
Modify Currency Exchange Rate

Modifies Currency Exchange Rate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **currencyExchangeRate** | [**HandlerModifyCurrencyExchangeRate**](HandlerModifyCurrencyExchangeRate.md)| Modify Currency Exchange Rate request | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CurrencyexchangeStatustogglerGet**
> HandlerResponse CurrencyexchangeStatustogglerGet(ctx, currencyName)
Currency Exchange Rate Status Toggler

Toggles Currency Exchange Rate Status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **currencyName** | **string**| Currency Name | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

