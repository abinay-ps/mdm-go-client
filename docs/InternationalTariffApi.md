# \InternationalTariffApi

All URIs are relative to *https://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InternationalproductsTariffBlindliteratureGet**](InternationalTariffApi.md#InternationalproductsTariffBlindliteratureGet) | **Get** /internationalproducts/tariff/blindliterature | Get International Blind Literature Tariff
[**InternationalproductsTariffBulkbagGet**](InternationalTariffApi.md#InternationalproductsTariffBulkbagGet) | **Get** /internationalproducts/tariff/bulkbag | Get International Bulk Bag Tariff
[**InternationalproductsTariffEmsGet**](InternationalTariffApi.md#InternationalproductsTariffEmsGet) | **Get** /internationalproducts/tariff/ems | Get International Speed Document Tariff
[**InternationalproductsTariffItpsGet**](InternationalTariffApi.md#InternationalproductsTariffItpsGet) | **Get** /internationalproducts/tariff/itps | Get International Tracked Packets Tariff
[**InternationalproductsTariffLetterGet**](InternationalTariffApi.md#InternationalproductsTariffLetterGet) | **Get** /internationalproducts/tariff/letter | Get International Letter Tariff
[**InternationalproductsTariffParcelGet**](InternationalTariffApi.md#InternationalproductsTariffParcelGet) | **Get** /internationalproducts/tariff/parcel | Get International Parcel Tariff
[**InternationalproductsTariffPrintedpapersGet**](InternationalTariffApi.md#InternationalproductsTariffPrintedpapersGet) | **Get** /internationalproducts/tariff/printedpapers | Get International Printed Papers Tariff
[**InternationalproductsTariffSmallpacketsGet**](InternationalTariffApi.md#InternationalproductsTariffSmallpacketsGet) | **Get** /internationalproducts/tariff/smallpackets | Get International Small Packets Tariff


# **InternationalproductsTariffBlindliteratureGet**
> HandlerResponse InternationalproductsTariffBlindliteratureGet(ctx, productCode, weight, countryCode, optional)
Get International Blind Literature Tariff

Gives tariff for International Blind Literature Tariff

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product Code | 
  **weight** | **string**| Weight | 
  **countryCode** | **string**| Country Code | 
 **optional** | ***InternationalTariffApiInternationalproductsTariffBlindliteratureGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InternationalTariffApiInternationalproductsTariffBlindliteratureGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **length** | **optional.String**| Length | 
 **width** | **optional.String**| Width | 
 **height** | **optional.String**| Height | 
 **diameter** | **optional.String**| Diameter | 
 **modeOfTransmission** | **optional.String**| Mode of Transmission | 
 **vasCode** | [**optional.Interface of []string**](string.md)| VAS Codes | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InternationalproductsTariffBulkbagGet**
> HandlerResponse InternationalproductsTariffBulkbagGet(ctx, productCode, weight, countryCode, optional)
Get International Bulk Bag Tariff

Gives tariff for International Bulk Bag Tariff

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product Code | 
  **weight** | **string**| Weight | 
  **countryCode** | **string**| Country Code | 
 **optional** | ***InternationalTariffApiInternationalproductsTariffBulkbagGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InternationalTariffApiInternationalproductsTariffBulkbagGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **length** | **optional.String**| Length | 
 **width** | **optional.String**| Width | 
 **height** | **optional.String**| Height | 
 **diameter** | **optional.String**| Diameter | 
 **iNSAMOUNT** | **optional.String**| Ins Amount | 
 **modeOfTransmission** | **optional.String**| Mode of Transmission | 
 **vasCode** | [**optional.Interface of []string**](string.md)| VAS Codes | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InternationalproductsTariffEmsGet**
> HandlerResponse InternationalproductsTariffEmsGet(ctx, productCode, weight, countryCode, optional)
Get International Speed Document Tariff

Gives tariff for International Speed Document Tariff

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product Code | 
  **weight** | **string**| Weight | 
  **countryCode** | **string**| Country Code | 
 **optional** | ***InternationalTariffApiInternationalproductsTariffEmsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InternationalTariffApiInternationalproductsTariffEmsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **length** | **optional.String**| Length | 
 **width** | **optional.String**| Width | 
 **height** | **optional.String**| Height | 
 **diameter** | **optional.String**| Diameter | 
 **modeOfTransmission** | **optional.String**| Mode of Transmission | 
 **vasCode** | [**optional.Interface of []string**](string.md)| VAS Codes | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InternationalproductsTariffItpsGet**
> HandlerResponse InternationalproductsTariffItpsGet(ctx, productCode, weight, countryCode, optional)
Get International Tracked Packets Tariff

Gives tariff for International Tracked Packets Tariff

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product Code | 
  **weight** | **string**| Weight | 
  **countryCode** | **string**| Country Code | 
 **optional** | ***InternationalTariffApiInternationalproductsTariffItpsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InternationalTariffApiInternationalproductsTariffItpsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **length** | **optional.String**| Length | 
 **width** | **optional.String**| Width | 
 **height** | **optional.String**| Height | 
 **diameter** | **optional.String**| Diameter | 
 **modeOfTransmission** | **optional.String**| Mode of Transmission | 
 **vasCode** | [**optional.Interface of []string**](string.md)| VAS Codes | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InternationalproductsTariffLetterGet**
> HandlerResponse InternationalproductsTariffLetterGet(ctx, productCode, weight, countryCode, optional)
Get International Letter Tariff

Gives tariff for International Letter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product Code | 
  **weight** | **string**| Weight | 
  **countryCode** | **string**| Country Code | 
 **optional** | ***InternationalTariffApiInternationalproductsTariffLetterGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InternationalTariffApiInternationalproductsTariffLetterGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **length** | **optional.String**| Length | 
 **width** | **optional.String**| Width | 
 **height** | **optional.String**| Height | 
 **diameter** | **optional.String**| Diameter | 
 **modeOfTransmission** | **optional.String**| Mode of Transmission | 
 **vasCode** | [**optional.Interface of []string**](string.md)| VAS Codes | 
 **iNSAMOUNT** | **optional.String**| Ins Amount | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InternationalproductsTariffParcelGet**
> HandlerResponse InternationalproductsTariffParcelGet(ctx, productCode, weight, countryCode, optional)
Get International Parcel Tariff

Gives tariff for International Parcel Tariff

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product Code | 
  **weight** | **string**| Weight | 
  **countryCode** | **string**| Country Code | 
 **optional** | ***InternationalTariffApiInternationalproductsTariffParcelGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InternationalTariffApiInternationalproductsTariffParcelGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **length** | **optional.String**| Length | 
 **width** | **optional.String**| Width | 
 **height** | **optional.String**| Height | 
 **diameter** | **optional.String**| Diameter | 
 **modeOfTransmission** | **optional.String**| Mode of Transmission | 
 **vasCode** | [**optional.Interface of []string**](string.md)| VAS Codes | 
 **iNSAMOUNT** | **optional.String**| INS Amount | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InternationalproductsTariffPrintedpapersGet**
> HandlerResponse InternationalproductsTariffPrintedpapersGet(ctx, productCode, weight, countryCode, optional)
Get International Printed Papers Tariff

Gives tariff for International Printed Papers Tariff

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product Code | 
  **weight** | **string**| Weight | 
  **countryCode** | **string**| Country Code | 
 **optional** | ***InternationalTariffApiInternationalproductsTariffPrintedpapersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InternationalTariffApiInternationalproductsTariffPrintedpapersGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **length** | **optional.String**| Length | 
 **width** | **optional.String**| Width | 
 **height** | **optional.String**| Height | 
 **diameter** | **optional.String**| Diameter | 
 **iNSAMOUNT** | **optional.String**| INS Amount | 
 **modeOfTransmission** | **optional.String**| Mode of Transmission | 
 **vasCode** | [**optional.Interface of []string**](string.md)| VAS Codes | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InternationalproductsTariffSmallpacketsGet**
> HandlerResponse InternationalproductsTariffSmallpacketsGet(ctx, productCode, weight, countryCode, optional)
Get International Small Packets Tariff

Gives tariff for International Small Packets Tariff

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product Code | 
  **weight** | **string**| Weight | 
  **countryCode** | **string**| Country Code | 
 **optional** | ***InternationalTariffApiInternationalproductsTariffSmallpacketsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InternationalTariffApiInternationalproductsTariffSmallpacketsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **length** | **optional.String**| Length | 
 **width** | **optional.String**| Width | 
 **height** | **optional.String**| Height | 
 **diameter** | **optional.String**| Diameter | 
 **iNSAMOUNT** | **optional.String**| INS Amount | 
 **modeOfTransmission** | **optional.String**| Mode of Transmission | 
 **vasCode** | [**optional.Interface of []string**](string.md)| VAS Codes | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

