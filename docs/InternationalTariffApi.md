# \InternationalTariffApi

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InternationalTariffsBlindLiteraturesGet**](InternationalTariffApi.md#InternationalTariffsBlindLiteraturesGet) | **Get** /international-tariffs/blind-literatures | Calculate the tariff for international blind literature
[**InternationalTariffsBulkBagsGet**](InternationalTariffApi.md#InternationalTariffsBulkBagsGet) | **Get** /international-tariffs/bulk-bags | Calculate the tariff for international bulk bags
[**InternationalTariffsEmsGet**](InternationalTariffApi.md#InternationalTariffsEmsGet) | **Get** /international-tariffs/ems | Calculate the tariff for international speed documents
[**InternationalTariffsItpsGet**](InternationalTariffApi.md#InternationalTariffsItpsGet) | **Get** /international-tariffs/itps | Calculate the tariff for international tracked packets
[**InternationalTariffsLettersGet**](InternationalTariffApi.md#InternationalTariffsLettersGet) | **Get** /international-tariffs/letters | Calculate the tariff for foreign letters
[**InternationalTariffsParcelsGet**](InternationalTariffApi.md#InternationalTariffsParcelsGet) | **Get** /international-tariffs/parcels | Calculate the tariff for foreign parcels
[**InternationalTariffsPrintedPapersGet**](InternationalTariffApi.md#InternationalTariffsPrintedPapersGet) | **Get** /international-tariffs/printed-papers | Calculate the tariff for international printed papers
[**InternationalTariffsSmallPacketsGet**](InternationalTariffApi.md#InternationalTariffsSmallPacketsGet) | **Get** /international-tariffs/small-packets | Calculate the tariff for international small packets


# **InternationalTariffsBlindLiteraturesGet**
> ResponseGetFgnBlindLiteratureTariffApiResponse InternationalTariffsBlindLiteraturesGet(ctx, countryCode, productCode, weight, optional)
Calculate the tariff for international blind literature

Calculates the tariff information for international blind literature based on various parameters, including product code, country code, weight, dimensions, and value-added services (VAS).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countryCode** | **string**|  | 
  **productCode** | **string**|  | 
  **weight** | **float32**|  | 
 **optional** | ***InternationalTariffApiInternationalTariffsBlindLiteraturesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InternationalTariffApiInternationalTariffsBlindLiteraturesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **diameter** | **optional.Float32**|  | 
 **height** | **optional.Float32**|  | 
 **length** | **optional.Float32**|  | 
 **modeOfTransmission** | **optional.String**|  | 
 **vasCode** | [**optional.Interface of []string**](string.md)|  | 
 **width** | **optional.Float32**|  | 

### Return type

[**ResponseGetFgnBlindLiteratureTariffApiResponse**](response.GetFgnBlindLiteratureTariffAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InternationalTariffsBulkBagsGet**
> ResponseGetFgnBulkBagTariffApiResponse InternationalTariffsBulkBagsGet(ctx, countryCode, productCode, weight, optional)
Calculate the tariff for international bulk bags

Calculates the tariff information for international bulk bags based on various parameters, including product code, country code, weight, dimensions, and value-added services (VAS).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countryCode** | **string**|  | 
  **productCode** | **string**|  | 
  **weight** | **float32**|  | 
 **optional** | ***InternationalTariffApiInternationalTariffsBulkBagsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InternationalTariffApiInternationalTariffsBulkBagsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **diameter** | **optional.Float32**|  | 
 **height** | **optional.Float32**|  | 
 **insAmount** | **optional.String**|  | 
 **length** | **optional.Float32**|  | 
 **modeOfTransmission** | **optional.String**|  | 
 **vasCode** | [**optional.Interface of []string**](string.md)|  | 
 **width** | **optional.Float32**|  | 

### Return type

[**ResponseGetFgnBulkBagTariffApiResponse**](response.GetFgnBulkBagTariffAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InternationalTariffsEmsGet**
> ResponseGetFgnSpeedDocumentTariffApiResponse InternationalTariffsEmsGet(ctx, countryCode, productCode, weight, optional)
Calculate the tariff for international speed documents

Calculate the tariff information for international speed documents based on various parameters, including product code, country code, weight, dimensions, and value-added services (VAS).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countryCode** | **string**|  | 
  **productCode** | **string**|  | 
  **weight** | **float32**|  | 
 **optional** | ***InternationalTariffApiInternationalTariffsEmsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InternationalTariffApiInternationalTariffsEmsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **diameter** | **optional.Float32**|  | 
 **height** | **optional.Float32**|  | 
 **insAmount** | **optional.Float32**|  | 
 **length** | **optional.Float32**|  | 
 **modeOfTransmission** | **optional.String**|  | 
 **vasCode** | [**optional.Interface of []string**](string.md)|  | 
 **width** | **optional.Float32**|  | 

### Return type

[**ResponseGetFgnSpeedDocumentTariffApiResponse**](response.GetFgnSpeedDocumentTariffAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InternationalTariffsItpsGet**
> ResponseGetFgnItpsTariffApiResponse InternationalTariffsItpsGet(ctx, countryCode, productCode, weight, optional)
Calculate the tariff for international tracked packets

Calculate the tariff information for international tracked packets based on various parameters, including product code, country code, weight, dimensions, and value-added services (VAS).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countryCode** | **string**|  | 
  **productCode** | **string**|  | 
  **weight** | **float32**|  | 
 **optional** | ***InternationalTariffApiInternationalTariffsItpsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InternationalTariffApiInternationalTariffsItpsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **diameter** | **optional.Float32**|  | 
 **height** | **optional.Float32**|  | 
 **insAmount** | **optional.Float32**|  | 
 **length** | **optional.Float32**|  | 
 **modeOfTransmission** | **optional.String**|  | 
 **vasCode** | [**optional.Interface of []string**](string.md)|  | 
 **width** | **optional.Float32**|  | 

### Return type

[**ResponseGetFgnItpsTariffApiResponse**](response.GetFgnITPSTariffAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InternationalTariffsLettersGet**
> ResponseGetFgnLetterTariffApiResponse InternationalTariffsLettersGet(ctx, countryCode, productCode, weight, optional)
Calculate the tariff for foreign letters

Calculates the tariff information for a foreign letter based on various parameters, including product code, country code, weight, dimensions, and value-added services (VAS).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countryCode** | **string**|  | 
  **productCode** | **string**|  | 
  **weight** | **float32**|  | 
 **optional** | ***InternationalTariffApiInternationalTariffsLettersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InternationalTariffApiInternationalTariffsLettersGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **height** | **optional.Float32**|  | 
 **insAmount** | **optional.Float32**|  | 
 **length** | **optional.Float32**|  | 
 **modeOfTransmission** | **optional.String**|  | 
 **vasCode** | [**optional.Interface of []string**](string.md)|  | 
 **width** | **optional.Float32**|  | 

### Return type

[**ResponseGetFgnLetterTariffApiResponse**](response.GetFgnLetterTariffAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InternationalTariffsParcelsGet**
> ResponseGetFgnParcelTariffApiResponse InternationalTariffsParcelsGet(ctx, countryCode, productCode, weight, optional)
Calculate the tariff for foreign parcels

Calculate the tariff information for a foreign parcel based on various parameters, including product code, country code, weight, dimensions, and value-added services (VAS).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countryCode** | **string**|  | 
  **productCode** | **string**|  | 
  **weight** | **float32**|  | 
 **optional** | ***InternationalTariffApiInternationalTariffsParcelsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InternationalTariffApiInternationalTariffsParcelsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **diameter** | **optional.Float32**|  | 
 **height** | **optional.Float32**|  | 
 **insAmount** | **optional.Float32**|  | 
 **length** | **optional.Float32**|  | 
 **modeOfTransmission** | **optional.String**|  | 
 **vasCode** | [**optional.Interface of []string**](string.md)|  | 
 **width** | **optional.Float32**|  | 

### Return type

[**ResponseGetFgnParcelTariffApiResponse**](response.GetFgnParcelTariffAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InternationalTariffsPrintedPapersGet**
> ResponseGetFgnPrintedPapersTariffApiResponse InternationalTariffsPrintedPapersGet(ctx, countryCode, productCode, weight, optional)
Calculate the tariff for international printed papers

Calculate the tariff information for international printed papers based on various parameters, including product code, country code, weight, dimensions, and value-added services (VAS).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countryCode** | **string**|  | 
  **productCode** | **string**|  | 
  **weight** | **float32**|  | 
 **optional** | ***InternationalTariffApiInternationalTariffsPrintedPapersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InternationalTariffApiInternationalTariffsPrintedPapersGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **diameter** | **optional.Float32**|  | 
 **height** | **optional.Float32**|  | 
 **insAmount** | **optional.Float32**|  | 
 **length** | **optional.Float32**|  | 
 **modeOfTransmission** | **optional.String**|  | 
 **vasCode** | [**optional.Interface of []string**](string.md)|  | 
 **width** | **optional.Float32**|  | 

### Return type

[**ResponseGetFgnPrintedPapersTariffApiResponse**](response.GetFgnPrintedPapersTariffAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InternationalTariffsSmallPacketsGet**
> ResponseGetFgnSmallPacketsTariffApiResponse InternationalTariffsSmallPacketsGet(ctx, countryCode, productCode, weight, optional)
Calculate the tariff for international small packets

Calculate the tariff information for international small packets based on various parameters, including product code, country code, weight, dimensions, and value-added services (VAS).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countryCode** | **string**|  | 
  **productCode** | **string**|  | 
  **weight** | **float32**|  | 
 **optional** | ***InternationalTariffApiInternationalTariffsSmallPacketsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InternationalTariffApiInternationalTariffsSmallPacketsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **diameter** | **optional.Float32**|  | 
 **height** | **optional.Float32**|  | 
 **insAmount** | **optional.Float32**|  | 
 **length** | **optional.Float32**|  | 
 **modeOfTransmission** | **optional.String**|  | 
 **vasCode** | [**optional.Interface of []string**](string.md)|  | 
 **width** | **optional.Float32**|  | 

### Return type

[**ResponseGetFgnSmallPacketsTariffApiResponse**](response.GetFgnSmallPacketsTariffAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

