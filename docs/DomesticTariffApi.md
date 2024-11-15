# \DomesticTariffApi

All URIs are relative to *https://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomesticproductsTariffApsDapltariffGet**](DomesticTariffApi.md#DomesticproductsTariffApsDapltariffGet) | **Get** /domesticproducts/tariff/aps/dapltariff | Get DAPL Tariff
[**DomesticproductsTariffApsSdstariffGet**](DomesticTariffApi.md#DomesticproductsTariffApsSdstariffGet) | **Get** /domesticproducts/tariff/aps/sdstariff | Get SDS Tariff
[**DomesticproductsTariffBillmailGet**](DomesticTariffApi.md#DomesticproductsTariffBillmailGet) | **Get** /domesticproducts/tariff/billmail | Get Bill Mail Tariff
[**DomesticproductsTariffBlindliteratureGet**](DomesticTariffApi.md#DomesticproductsTariffBlindliteratureGet) | **Get** /domesticproducts/tariff/blindliterature | Get Blind Literature Tariff
[**DomesticproductsTariffBookpacketGet**](DomesticTariffApi.md#DomesticproductsTariffBookpacketGet) | **Get** /domesticproducts/tariff/bookpacket | Get Book Packet Tariff
[**DomesticproductsTariffBpperiodicalGet**](DomesticTariffApi.md#DomesticproductsTariffBpperiodicalGet) | **Get** /domesticproducts/tariff/bpperiodical | Get Book Packet containing Periodical Tariff
[**DomesticproductsTariffBusinessparcelGet**](DomesticTariffApi.md#DomesticproductsTariffBusinessparcelGet) | **Get** /domesticproducts/tariff/businessparcel | Get Business Parcel Tariff
[**DomesticproductsTariffDirectpostGet**](DomesticTariffApi.md#DomesticproductsTariffDirectpostGet) | **Get** /domesticproducts/tariff/directpost | Get Direct Post Tariff
[**DomesticproductsTariffEmoGet**](DomesticTariffApi.md#DomesticproductsTariffEmoGet) | **Get** /domesticproducts/tariff/emo/ | Get Emo Tariff
[**DomesticproductsTariffIpoGet**](DomesticTariffApi.md#DomesticproductsTariffIpoGet) | **Get** /domesticproducts/tariff/ipo/ | Get IPO Tariff
[**DomesticproductsTariffLetterGet**](DomesticTariffApi.md#DomesticproductsTariffLetterGet) | **Get** /domesticproducts/tariff/letter | Get Letter Tariff
[**DomesticproductsTariffLettercardGet**](DomesticTariffApi.md#DomesticproductsTariffLettercardGet) | **Get** /domesticproducts/tariff/lettercard | Get Letter Card Tariff
[**DomesticproductsTariffMagazinepostGet**](DomesticTariffApi.md#DomesticproductsTariffMagazinepostGet) | **Get** /domesticproducts/tariff/magazinepost | Get Magazine Post Tariff
[**DomesticproductsTariffNationalbillmailGet**](DomesticTariffApi.md#DomesticproductsTariffNationalbillmailGet) | **Get** /domesticproducts/tariff/nationalbillmail | Get National Bill Mail Tariff
[**DomesticproductsTariffNewspaperGet**](DomesticTariffApi.md#DomesticproductsTariffNewspaperGet) | **Get** /domesticproducts/tariff/newspaper | Get Newspaper Tariff
[**DomesticproductsTariffNewspaperbundleGet**](DomesticTariffApi.md#DomesticproductsTariffNewspaperbundleGet) | **Get** /domesticproducts/tariff/newspaperbundle | Get Newspaper Bundle Tariff
[**DomesticproductsTariffNonvariantproductsProductCodeGet**](DomesticTariffApi.md#DomesticproductsTariffNonvariantproductsProductCodeGet) | **Get** /domesticproducts/tariff/nonvariantproducts/{ProductCode} | Get Non Variant Tariff
[**DomesticproductsTariffParcelGet**](DomesticTariffApi.md#DomesticproductsTariffParcelGet) | **Get** /domesticproducts/tariff/parcel | Get Parcel Tariff
[**DomesticproductsTariffPatternsamplepacketGet**](DomesticTariffApi.md#DomesticproductsTariffPatternsamplepacketGet) | **Get** /domesticproducts/tariff/patternsamplepacket | Get Pattern Sample Packet Tariff
[**DomesticproductsTariffPostcardGet**](DomesticTariffApi.md#DomesticproductsTariffPostcardGet) | **Get** /domesticproducts/tariff/postcard | Get Post Card Tariff
[**DomesticproductsTariffPrintedbookGet**](DomesticTariffApi.md#DomesticproductsTariffPrintedbookGet) | **Get** /domesticproducts/tariff/printedbook | Get Printed Book Tariff
[**DomesticproductsTariffRetailproductsProductCodeGet**](DomesticTariffApi.md#DomesticproductsTariffRetailproductsProductCodeGet) | **Get** /domesticproducts/tariff/retailproducts/{ProductCode} | Get Retail Post Tariff
[**DomesticproductsTariffSpeedpostGet**](DomesticTariffApi.md#DomesticproductsTariffSpeedpostGet) | **Get** /domesticproducts/tariff/speedpost | Get Domestic Speed Post Tariff


# **DomesticproductsTariffApsDapltariffGet**
> HandlerResponse DomesticproductsTariffApsDapltariffGet(ctx, productCode, weight, sourcePincode, destinationPincode, optional)
Get DAPL Tariff

Gives tariff for DAPL product

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product Code | 
  **weight** | **string**| Weight | 
  **sourcePincode** | **string**| Source Pincode | 
  **destinationPincode** | **string**| Destination Pincode | 
 **optional** | ***DomesticTariffApiDomesticproductsTariffApsDapltariffGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomesticTariffApiDomesticproductsTariffApsDapltariffGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **rEG** | **optional.String**| Reg | 
 **iNS** | **optional.String**| Ins | 
 **iNSC** | **optional.String**| InsC | 
 **pOD** | **optional.String**| Pod | 
 **aCK** | **optional.String**| Ack | 
 **vPP** | **optional.String**| Vpp | 
 **aMS** | **optional.String**| Ams | 
 **length** | **optional.String**| Length | 
 **width** | **optional.String**| Width | 
 **height** | **optional.String**| Height | 
 **diameter** | **optional.String**| Diameter | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticproductsTariffApsSdstariffGet**
> HandlerResponse DomesticproductsTariffApsSdstariffGet(ctx, productCode, weight, optional)
Get SDS Tariff

Gives tariff for APS SDS Product

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product Code | 
  **weight** | **string**| Weight | 
 **optional** | ***DomesticTariffApiDomesticproductsTariffApsSdstariffGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomesticTariffApiDomesticproductsTariffApsSdstariffGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rEG** | **optional.String**| Reg | 
 **iNS** | **optional.String**| Ins | 
 **aCK** | **optional.String**| Ack | 
 **vPP** | **optional.String**| Vpp | 
 **aMS** | **optional.String**| Ams | 
 **length** | **optional.String**| Length | 
 **width** | **optional.String**| Width | 
 **height** | **optional.String**| Height | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticproductsTariffBillmailGet**
> HandlerResponse DomesticproductsTariffBillmailGet(ctx, productCode, weight, quantity)
Get Bill Mail Tariff

Gives tariff for Bill Mail Service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product Code | 
  **weight** | **float32**| Weight | 
  **quantity** | **int32**| Quantity | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticproductsTariffBlindliteratureGet**
> HandlerResponse DomesticproductsTariffBlindliteratureGet(ctx, productCode, weight, optional)
Get Blind Literature Tariff

Gives tariff for Blind Literature

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product Code | 
  **weight** | **string**| Weight | 
 **optional** | ***DomesticTariffApiDomesticproductsTariffBlindliteratureGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomesticTariffApiDomesticproductsTariffBlindliteratureGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rEGBL** | **optional.String**| RegBl | 
 **iNS** | **optional.String**| Ins | 
 **aCKBL** | **optional.String**| AckBl | 
 **vPP** | **optional.String**| Vpp | 
 **aMS** | **optional.String**| Ams | 
 **length** | **optional.String**| Length | 
 **width** | **optional.String**| Width | 
 **height** | **optional.String**| Height | 
 **diameter** | **optional.String**| Diameter | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticproductsTariffBookpacketGet**
> HandlerResponse DomesticproductsTariffBookpacketGet(ctx, productCode, weight, optional)
Get Book Packet Tariff

Gives tariff for Book Packet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product Code | 
  **weight** | **string**| Weight | 
 **optional** | ***DomesticTariffApiDomesticproductsTariffBookpacketGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomesticTariffApiDomesticproductsTariffBookpacketGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rEG** | **optional.String**| Reg | 
 **iNS** | **optional.String**| Ins | 
 **aCK** | **optional.String**| Ack | 
 **vPP** | **optional.String**| Vpp | 
 **aMS** | **optional.String**| Ams | 
 **length** | **optional.String**| Length | 
 **width** | **optional.String**| Width | 
 **height** | **optional.String**| Height | 
 **diameter** | **optional.String**| Diameter | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticproductsTariffBpperiodicalGet**
> HandlerResponse DomesticproductsTariffBpperiodicalGet(ctx, productCode, weight, optional)
Get Book Packet containing Periodical Tariff

Gives tariff for Book Packet containing Periodicals

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product Code | 
  **weight** | **string**| Weight | 
 **optional** | ***DomesticTariffApiDomesticproductsTariffBpperiodicalGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomesticTariffApiDomesticproductsTariffBpperiodicalGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rEG** | **optional.String**| Reg | 
 **iNS** | **optional.String**| Ins | 
 **aCK** | **optional.String**| Ack | 
 **vPP** | **optional.String**| Vpp | 
 **aMS** | **optional.String**| Ams | 
 **value** | **optional.String**| Value | 
 **length** | **optional.String**| Length | 
 **width** | **optional.String**| Width | 
 **height** | **optional.String**| Height | 
 **diameter** | **optional.String**| Diameter | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticproductsTariffBusinessparcelGet**
> HandlerResponse DomesticproductsTariffBusinessparcelGet(ctx, productCode, weight, sourcePincode, destinationPincode, optional)
Get Business Parcel Tariff

Gives tariff for Business Parcel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product Code | 
  **weight** | **string**| Weight | 
  **sourcePincode** | **string**| Source Pincode | 
  **destinationPincode** | **string**| Destination Pincode | 
 **optional** | ***DomesticTariffApiDomesticproductsTariffBusinessparcelGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomesticTariffApiDomesticproductsTariffBusinessparcelGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **rEG** | **optional.String**| Reg | 
 **iNS** | **optional.String**| Ins | 
 **iNSC** | **optional.String**| InsC | 
 **aCK** | **optional.String**| Ack | 
 **vPP** | **optional.String**| Vpp | 
 **aMS** | **optional.String**| Ams | 
 **length** | **optional.String**| Length | 
 **width** | **optional.String**| Width | 
 **height** | **optional.String**| Height | 
 **diameter** | **optional.String**| Diameter | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticproductsTariffDirectpostGet**
> HandlerResponse DomesticproductsTariffDirectpostGet(ctx, productCode, weight, quantity, sourcePincode, destinationPincode)
Get Direct Post Tariff

Gives tariff for Direct Post

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product Code | 
  **weight** | **string**| Weight | 
  **quantity** | **string**| Quantity | 
  **sourcePincode** | **string**| Source Pincode | 
  **destinationPincode** | **string**| Destination Pincode | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticproductsTariffEmoGet**
> HandlerResponse DomesticproductsTariffEmoGet(ctx, productCode, eMOAmount)
Get Emo Tariff

Gives tariff for Emo

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product Code | 
  **eMOAmount** | **float32**| EMOAmount | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticproductsTariffIpoGet**
> HandlerResponse DomesticproductsTariffIpoGet(ctx, productCode, iPOAmount)
Get IPO Tariff

Gives tariff for IPO

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product Code | 
  **iPOAmount** | **float32**| IPO Amount | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticproductsTariffLetterGet**
> HandlerResponse DomesticproductsTariffLetterGet(ctx, productCode, weight, optional)
Get Letter Tariff

Gives tariff for Letter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product Code | 
  **weight** | **string**| Weight | 
 **optional** | ***DomesticTariffApiDomesticproductsTariffLetterGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomesticTariffApiDomesticproductsTariffLetterGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rEG** | **optional.String**| Reg | 
 **iNS** | **optional.String**| Ins | 
 **aCK** | **optional.String**| Ack | 
 **vPP** | **optional.String**| Vpp | 
 **length** | **optional.String**| Length | 
 **width** | **optional.String**| Width | 
 **height** | **optional.String**| Height | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticproductsTariffLettercardGet**
> HandlerResponse DomesticproductsTariffLettercardGet(ctx, productCode, weight, optional)
Get Letter Card Tariff

Gives tariff for Letter Card

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product Code | 
  **weight** | **string**| Weight | 
 **optional** | ***DomesticTariffApiDomesticproductsTariffLettercardGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomesticTariffApiDomesticproductsTariffLettercardGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rEG** | **optional.String**| Reg | 
 **aCK** | **optional.String**| Ack | 
 **length** | **optional.String**| Length | 
 **width** | **optional.String**| Width | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticproductsTariffMagazinepostGet**
> HandlerResponse DomesticproductsTariffMagazinepostGet(ctx, productCode, weight, sourcePincode, destinationPincode, optional)
Get Magazine Post Tariff

Gives tariff for Magazine Post

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product Code | 
  **weight** | **string**| Weight | 
  **sourcePincode** | **string**| Source Pincode | 
  **destinationPincode** | **string**| Destination Pincode | 
 **optional** | ***DomesticTariffApiDomesticproductsTariffMagazinepostGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomesticTariffApiDomesticproductsTariffMagazinepostGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **length** | **optional.String**| Length | 
 **width** | **optional.String**| Width | 
 **height** | **optional.String**| Height | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticproductsTariffNationalbillmailGet**
> HandlerResponse DomesticproductsTariffNationalbillmailGet(ctx, productCode, weight, quantity)
Get National Bill Mail Tariff

Gives tariff for National Bill Mail Service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product Code | 
  **weight** | **float32**| Weight | 
  **quantity** | **int32**| Quantity | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticproductsTariffNewspaperGet**
> HandlerResponse DomesticproductsTariffNewspaperGet(ctx, productCode, weight, optional)
Get Newspaper Tariff

Gives tariff for Newspaper

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product Code | 
  **weight** | **string**| Weight | 
 **optional** | ***DomesticTariffApiDomesticproductsTariffNewspaperGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomesticTariffApiDomesticproductsTariffNewspaperGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rEG** | **optional.String**| Reg | 
 **iNS** | **optional.String**| Ins | 
 **aCK** | **optional.String**| Ack | 
 **vPP** | **optional.String**| Vpp | 
 **aMS** | **optional.String**| Ams | 
 **length** | **optional.String**| Length | 
 **width** | **optional.String**| Width | 
 **height** | **optional.String**| Height | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticproductsTariffNewspaperbundleGet**
> HandlerResponse DomesticproductsTariffNewspaperbundleGet(ctx, productCode, weight, optional)
Get Newspaper Bundle Tariff

Gives tariff for Newspaper Bundle

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product Code | 
  **weight** | **string**| Weight | 
 **optional** | ***DomesticTariffApiDomesticproductsTariffNewspaperbundleGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomesticTariffApiDomesticproductsTariffNewspaperbundleGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rEG** | **optional.String**| Reg | 
 **iNS** | **optional.String**| Ins | 
 **aCK** | **optional.String**| Ack | 
 **vPP** | **optional.String**| Vpp | 
 **aMS** | **optional.String**| Ams | 
 **length** | **optional.String**| Length | 
 **width** | **optional.String**| Width | 
 **height** | **optional.String**| Height | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticproductsTariffNonvariantproductsProductCodeGet**
> HandlerResponse DomesticproductsTariffNonvariantproductsProductCodeGet(ctx, productCode)
Get Non Variant Tariff

Gives tariff for Non Variant Products

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product Code | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticproductsTariffParcelGet**
> HandlerResponse DomesticproductsTariffParcelGet(ctx, productCode, weight, optional)
Get Parcel Tariff

Gives tariff for Parcel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product Code | 
  **weight** | **string**| Weight | 
 **optional** | ***DomesticTariffApiDomesticproductsTariffParcelGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomesticTariffApiDomesticproductsTariffParcelGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rEG** | **optional.String**| Reg | 
 **iNS** | **optional.String**| Ins | 
 **aCK** | **optional.String**| Ack | 
 **vPP** | **optional.String**| Vpp | 
 **aMS** | **optional.String**| Ams | 
 **length** | **optional.String**| Length | 
 **width** | **optional.String**| Width | 
 **height** | **optional.String**| Height | 
 **diameter** | **optional.String**| Diameter | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticproductsTariffPatternsamplepacketGet**
> HandlerResponse DomesticproductsTariffPatternsamplepacketGet(ctx, productCode, weight, optional)
Get Pattern Sample Packet Tariff

Gives tariff for Pattern Sample Packet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product Code | 
  **weight** | **string**| Weight | 
 **optional** | ***DomesticTariffApiDomesticproductsTariffPatternsamplepacketGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomesticTariffApiDomesticproductsTariffPatternsamplepacketGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rEG** | **optional.String**| Reg | 
 **iNS** | **optional.String**| Ins | 
 **aCK** | **optional.String**| Ack | 
 **vPP** | **optional.String**| Vpp | 
 **aMS** | **optional.String**| Ams | 
 **length** | **optional.String**| Length | 
 **width** | **optional.String**| Width | 
 **height** | **optional.String**| Height | 
 **diameter** | **optional.String**| Diameter | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticproductsTariffPostcardGet**
> HandlerResponse DomesticproductsTariffPostcardGet(ctx, productCode, optional)
Get Post Card Tariff

Gives tariff for Post Card

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product Code | 
 **optional** | ***DomesticTariffApiDomesticproductsTariffPostcardGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomesticTariffApiDomesticproductsTariffPostcardGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rEG** | **optional.String**| Reg | 
 **aCK** | **optional.String**| Ack | 
 **length** | **optional.String**| Length | 
 **width** | **optional.String**| Width | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticproductsTariffPrintedbookGet**
> HandlerResponse DomesticproductsTariffPrintedbookGet(ctx, productCode, weight, optional)
Get Printed Book Tariff

Gives tariff for Printed Book

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product Code | 
  **weight** | **string**| Weight | 
 **optional** | ***DomesticTariffApiDomesticproductsTariffPrintedbookGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomesticTariffApiDomesticproductsTariffPrintedbookGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rEG** | **optional.String**| Reg | 
 **rEGC** | **optional.String**| RegC | 
 **iNS** | **optional.String**| Ins | 
 **aCK** | **optional.String**| Ack | 
 **vPP** | **optional.String**| Vpp | 
 **aMS** | **optional.String**| Ams | 
 **length** | **optional.String**| Length | 
 **width** | **optional.String**| Width | 
 **height** | **optional.String**| Height | 
 **diameter** | **optional.String**| Diameter | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticproductsTariffRetailproductsProductCodeGet**
> HandlerResponse DomesticproductsTariffRetailproductsProductCodeGet(ctx, productCode)
Get Retail Post Tariff

Gives tariff for Retail Post

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product Code | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticproductsTariffSpeedpostGet**
> HandlerResponse DomesticproductsTariffSpeedpostGet(ctx, productCode, weight, sourcePincode, destinationPincode, optional)
Get Domestic Speed Post Tariff

Gives tariff for Domestic Speed Post

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product Code | 
  **weight** | **string**| Weight | 
  **sourcePincode** | **string**| Source Pincode | 
  **destinationPincode** | **string**| Destination Pincode | 
 **optional** | ***DomesticTariffApiDomesticproductsTariffSpeedpostGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomesticTariffApiDomesticproductsTariffSpeedpostGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **rEG** | **optional.String**| Reg | 
 **iNS** | **optional.String**| Ins | 
 **iNSC** | **optional.String**| InsC | 
 **pOD** | **optional.String**| Pod | 
 **aCK** | **optional.String**| Ack | 
 **vPP** | **optional.String**| Vpp | 
 **aMS** | **optional.String**| Ams | 
 **length** | **optional.String**| Length | 
 **width** | **optional.String**| Width | 
 **height** | **optional.String**| Height | 
 **diameter** | **optional.String**| Diameter | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

