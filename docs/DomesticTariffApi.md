# \DomesticTariffApi

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomesticTariffsApsDaplGet**](DomesticTariffApi.md#DomesticTariffsApsDaplGet) | **Get** /domestic-tariffs/aps-dapl | Get DAPL tariff based on product code, weight, dimensions, and pincodes
[**DomesticTariffsApsSdsGet**](DomesticTariffApi.md#DomesticTariffsApsSdsGet) | **Get** /domestic-tariffs/aps-sds | Get SDS tariff based on product code, weight, and optional parameters
[**DomesticTariffsBillMailsGet**](DomesticTariffApi.md#DomesticTariffsBillMailsGet) | **Get** /domestic-tariffs/bill-mails | Get Bill Mail tariff based on product code, weight, and quantity
[**DomesticTariffsBlindLiteraturesGet**](DomesticTariffApi.md#DomesticTariffsBlindLiteraturesGet) | **Get** /domestic-tariffs/blind-literatures | Get blind literature tariff based on various parameters
[**DomesticTariffsBookPacketPeriodicalsGet**](DomesticTariffApi.md#DomesticTariffsBookPacketPeriodicalsGet) | **Get** /domestic-tariffs/book-packet-periodicals | Get the book packet periodical tariff based on various parameters
[**DomesticTariffsBookPacketsGet**](DomesticTariffApi.md#DomesticTariffsBookPacketsGet) | **Get** /domestic-tariffs/book-packets | Get book tariff based on various parameters
[**DomesticTariffsBusinessParcelsGet**](DomesticTariffApi.md#DomesticTariffsBusinessParcelsGet) | **Get** /domestic-tariffs/business-parcels | Get business parcel tariff based on various parameters
[**DomesticTariffsDirectPostsGet**](DomesticTariffApi.md#DomesticTariffsDirectPostsGet) | **Get** /domestic-tariffs/direct-posts | Get direct post tariff based on various parameters
[**DomesticTariffsEmoGet**](DomesticTariffApi.md#DomesticTariffsEmoGet) | **Get** /domestic-tariffs/emo | Get EMO tariff based on product code and EMO amount
[**DomesticTariffsIpoGet**](DomesticTariffApi.md#DomesticTariffsIpoGet) | **Get** /domestic-tariffs/ipo | Get IPO tariff based on product code and IPO amount
[**DomesticTariffsLetterCardsGet**](DomesticTariffApi.md#DomesticTariffsLetterCardsGet) | **Get** /domestic-tariffs/letter-cards | Get lettercard tariff based on various parameters
[**DomesticTariffsLettersGet**](DomesticTariffApi.md#DomesticTariffsLettersGet) | **Get** /domestic-tariffs/letters | Get letter tariff based on various parameters
[**DomesticTariffsMagazinePostsGet**](DomesticTariffApi.md#DomesticTariffsMagazinePostsGet) | **Get** /domestic-tariffs/magazine-posts | Get Magazine Post tariff based on product code, weight, dimensions, and pincodes
[**DomesticTariffsNationalBillMailsGet**](DomesticTariffApi.md#DomesticTariffsNationalBillMailsGet) | **Get** /domestic-tariffs/national-bill-mails | Get National Bill Mail tariff based on product code, weight, and quantity
[**DomesticTariffsNewspaperBundlesGet**](DomesticTariffApi.md#DomesticTariffsNewspaperBundlesGet) | **Get** /domestic-tariffs/newspaper-bundles | Get newspaper bundle tariff based on various parameters
[**DomesticTariffsNewspapersGet**](DomesticTariffApi.md#DomesticTariffsNewspapersGet) | **Get** /domestic-tariffs/newspapers | Get newspaper tariff based on various parameters
[**DomesticTariffsNonVariantsGet**](DomesticTariffApi.md#DomesticTariffsNonVariantsGet) | **Get** /domestic-tariffs/non-variants | Get the non-variant product tariff based on the product code
[**DomesticTariffsParcelsGet**](DomesticTariffApi.md#DomesticTariffsParcelsGet) | **Get** /domestic-tariffs/parcels | Get parcel tariff based on various parameters
[**DomesticTariffsPatternSamplePacketsGet**](DomesticTariffApi.md#DomesticTariffsPatternSamplePacketsGet) | **Get** /domestic-tariffs/pattern-sample-packets | Get pattern sample packet tariff based on various parameters
[**DomesticTariffsPostCardsGet**](DomesticTariffApi.md#DomesticTariffsPostCardsGet) | **Get** /domestic-tariffs/post-cards | Get Post Card tariff based on product code, weight, length, and width
[**DomesticTariffsPrintedBooksGet**](DomesticTariffApi.md#DomesticTariffsPrintedBooksGet) | **Get** /domestic-tariffs/printed-books | Get printed book tariff based on various parameters
[**DomesticTariffsRetailsGet**](DomesticTariffApi.md#DomesticTariffsRetailsGet) | **Get** /domestic-tariffs/retails | Get the retail post product tariff based on the product code
[**DomesticTariffsSpeedPostsGet**](DomesticTariffApi.md#DomesticTariffsSpeedPostsGet) | **Get** /domestic-tariffs/speed-posts | Get speed letter tariff based on various parameters


# **DomesticTariffsApsDaplGet**
> ResponseGetDaplTariffApiResponse DomesticTariffsApsDaplGet(ctx, destinationPincode, productCode, sourcePincode, weight, optional)
Get DAPL tariff based on product code, weight, dimensions, and pincodes

Calculate the DAPL tariff using the provided details, including product code, weight, dimensions (length, width, height, diameter),

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **destinationPincode** | **int32**|  | 
  **productCode** | **string**|  | 
  **sourcePincode** | **int32**|  | 
  **weight** | **float32**|  | 
 **optional** | ***DomesticTariffApiDomesticTariffsApsDaplGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomesticTariffApiDomesticTariffsApsDaplGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **aCK** | **optional.String**|  | 
 **aMS** | **optional.String**|  | 
 **iNS** | **optional.String**|  | 
 **iNSC** | **optional.String**|  | 
 **pOD** | **optional.String**|  | 
 **rEG** | **optional.String**|  | 
 **vPP** | **optional.String**|  | 
 **diameter** | **optional.Float32**|  | 
 **height** | **optional.Float32**|  | 
 **length** | **optional.Float32**|  | 
 **width** | **optional.Float32**|  | 

### Return type

[**ResponseGetDaplTariffApiResponse**](response.GetDAPLTariffAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticTariffsApsSdsGet**
> ResponseGetSdsTariffApiResponse DomesticTariffsApsSdsGet(ctx, productCode, weight, optional)
Get SDS tariff based on product code, weight, and optional parameters

Calculate the SDS tariff using the product code, weight, and optional parameters such as REG, INS, ACK, VPP, AMS, and dimensions (length, width, height).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**|  | 
  **weight** | **float32**|  | 
 **optional** | ***DomesticTariffApiDomesticTariffsApsSdsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomesticTariffApiDomesticTariffsApsSdsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aCK** | **optional.String**|  | 
 **aMS** | **optional.String**|  | 
 **iNS** | **optional.String**|  | 
 **rEG** | **optional.String**|  | 
 **vPP** | **optional.String**|  | 
 **height** | **optional.Float32**|  | 
 **length** | **optional.Float32**|  | 
 **width** | **optional.Float32**|  | 

### Return type

[**ResponseGetSdsTariffApiResponse**](response.GetSDSTariffAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticTariffsBillMailsGet**
> ResponseGetBillMailTariffApiResponse DomesticTariffsBillMailsGet(ctx, productCode, quantity, weight)
Get Bill Mail tariff based on product code, weight, and quantity

Calculate the Bill Mail tariff using the product code, weight, and quantity provided.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**|  | 
  **quantity** | **int32**|  | 
  **weight** | **float32**|  | 

### Return type

[**ResponseGetBillMailTariffApiResponse**](response.GetBillMailTariffAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticTariffsBlindLiteraturesGet**
> ResponseGetBlindLiteratureTariffApiResponse DomesticTariffsBlindLiteraturesGet(ctx, productCode, weight, optional)
Get blind literature tariff based on various parameters

Calculate the blind literature tariff using product details and additional optional parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**|  | 
  **weight** | **float32**|  | 
 **optional** | ***DomesticTariffApiDomesticTariffsBlindLiteraturesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomesticTariffApiDomesticTariffsBlindLiteraturesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aCKBL** | **optional.String**|  | 
 **aMS** | **optional.String**|  | 
 **rEGBL** | **optional.String**|  | 
 **diameter** | **optional.Float32**|  | 
 **height** | **optional.Float32**|  | 
 **length** | **optional.Float32**|  | 
 **width** | **optional.Float32**|  | 

### Return type

[**ResponseGetBlindLiteratureTariffApiResponse**](response.GetBlindLiteratureTariffAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticTariffsBookPacketPeriodicalsGet**
> ResponseGetBookTariffApiResponse DomesticTariffsBookPacketPeriodicalsGet(ctx, productCode, weight, optional)
Get the book packet periodical tariff based on various parameters

Calculate the tariff for a book packet using product details and additional optional parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**|  | 
  **weight** | **float32**|  | 
 **optional** | ***DomesticTariffApiDomesticTariffsBookPacketPeriodicalsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomesticTariffApiDomesticTariffsBookPacketPeriodicalsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aCK** | **optional.String**|  | 
 **aMS** | **optional.String**|  | 
 **iNS** | **optional.String**|  | 
 **rEG** | **optional.String**|  | 
 **vPP** | **optional.String**|  | 
 **diameter** | **optional.Float32**|  | 
 **height** | **optional.Float32**|  | 
 **length** | **optional.Float32**|  | 
 **value** | **optional.Float32**|  | 
 **width** | **optional.Float32**|  | 

### Return type

[**ResponseGetBookTariffApiResponse**](response.GetBookTariffAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticTariffsBookPacketsGet**
> ResponseGetBookTariffApiResponse DomesticTariffsBookPacketsGet(ctx, productCode, weight, optional)
Get book tariff based on various parameters

Calculate the book tariff using product details and various additional parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**|  | 
  **weight** | **float32**|  | 
 **optional** | ***DomesticTariffApiDomesticTariffsBookPacketsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomesticTariffApiDomesticTariffsBookPacketsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aCK** | **optional.String**|  | 
 **aMS** | **optional.String**|  | 
 **iNS** | **optional.String**|  | 
 **rEG** | **optional.String**|  | 
 **vPP** | **optional.String**|  | 
 **diameter** | **optional.Float32**|  | 
 **height** | **optional.Float32**|  | 
 **length** | **optional.Float32**|  | 
 **value** | **optional.Float32**|  | 
 **width** | **optional.Float32**|  | 

### Return type

[**ResponseGetBookTariffApiResponse**](response.GetBookTariffAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticTariffsBusinessParcelsGet**
> ResponseGetBusinessParcelTariffApiResponse DomesticTariffsBusinessParcelsGet(ctx, destinationPincode, productCode, sourcePincode, weight, optional)
Get business parcel tariff based on various parameters

Calculate the business parcel tariff using product details and additional optional parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **destinationPincode** | **string**|  | 
  **productCode** | **string**|  | 
  **sourcePincode** | **string**|  | 
  **weight** | **float32**|  | 
 **optional** | ***DomesticTariffApiDomesticTariffsBusinessParcelsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomesticTariffApiDomesticTariffsBusinessParcelsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **aCK** | **optional.String**|  | 
 **aMS** | **optional.String**|  | 
 **iNS** | **optional.String**|  | 
 **iNSC** | **optional.String**|  | 
 **rEG** | **optional.String**|  | 
 **vPP** | **optional.String**|  | 
 **diameter** | **optional.Float32**|  | 
 **height** | **optional.Float32**|  | 
 **length** | **optional.Float32**|  | 
 **width** | **optional.Float32**|  | 

### Return type

[**ResponseGetBusinessParcelTariffApiResponse**](response.GetBusinessParcelTariffAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticTariffsDirectPostsGet**
> ResponseGetDirectPostTariffApiResponse DomesticTariffsDirectPostsGet(ctx, destinationPincode, productCode, quantity, sourcePincode, weight)
Get direct post tariff based on various parameters

Calculate the direct post tariff using the provided product details, weight, and pincode information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **destinationPincode** | **int32**|  | 
  **productCode** | **string**|  | 
  **quantity** | **int32**|  | 
  **sourcePincode** | **int32**|  | 
  **weight** | **float32**|  | 

### Return type

[**ResponseGetDirectPostTariffApiResponse**](response.GetDirectPostTariffAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticTariffsEmoGet**
> ResponseGetEmoTariffApiResponse DomesticTariffsEmoGet(ctx, productCode, emoAmount)
Get EMO tariff based on product code and EMO amount

Calculate the EMO tariff using the product code and the EMO amount provided.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product code | 
  **emoAmount** | **float32**| EMO Amount | 

### Return type

[**ResponseGetEmoTariffApiResponse**](response.GetEMOTariffAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticTariffsIpoGet**
> ResponseGetIpoTariffApiResponse DomesticTariffsIpoGet(ctx, productCode, ipoAmount)
Get IPO tariff based on product code and IPO amount

Calculate the IPO tariff using the product code and the IPO amount provided.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product code | 
  **ipoAmount** | **float32**| IPO Amount | 

### Return type

[**ResponseGetIpoTariffApiResponse**](response.GetIPOTariffAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticTariffsLetterCardsGet**
> ResponseGetLetterCardTariffApiResponse DomesticTariffsLetterCardsGet(ctx, productCode, weight, optional)
Get lettercard tariff based on various parameters

Calculate the lettercard tariff using product details and additional optional parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**|  | 
  **weight** | **float32**|  | 
 **optional** | ***DomesticTariffApiDomesticTariffsLetterCardsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomesticTariffApiDomesticTariffsLetterCardsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aCK** | **optional.String**|  | 
 **iNS** | **optional.String**|  | 
 **rEG** | **optional.String**|  | 
 **vPP** | **optional.String**|  | 
 **height** | **optional.Float32**|  | 
 **length** | **optional.Float32**|  | 
 **width** | **optional.Float32**|  | 

### Return type

[**ResponseGetLetterCardTariffApiResponse**](response.GetLetterCardTariffAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticTariffsLettersGet**
> ResponseGetLetterTariffApiResponse DomesticTariffsLettersGet(ctx, productCode, weight, optional)
Get letter tariff based on various parameters

Calculate the letter tariff using product details and additional parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**|  | 
  **weight** | **float32**|  | 
 **optional** | ***DomesticTariffApiDomesticTariffsLettersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomesticTariffApiDomesticTariffsLettersGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aCK** | **optional.String**|  | 
 **iNS** | **optional.String**|  | 
 **rEG** | **optional.String**|  | 
 **vPP** | **optional.String**|  | 
 **height** | **optional.Float32**|  | 
 **length** | **optional.Float32**|  | 
 **width** | **optional.Float32**|  | 

### Return type

[**ResponseGetLetterTariffApiResponse**](response.GetLetterTariffAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticTariffsMagazinePostsGet**
> ResponseGetMagazinePostTariffApiResponse DomesticTariffsMagazinePostsGet(ctx, destinationPincode, productCode, sourcePincode, weight, optional)
Get Magazine Post tariff based on product code, weight, dimensions, and pincodes

Calculate the Magazine Post tariff using the product code, weight, dimensions (length, width, height), and source/destination pincodes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **destinationPincode** | **int32**|  | 
  **productCode** | **string**|  | 
  **sourcePincode** | **int32**|  | 
  **weight** | **float32**| Accepts int strings and floats | 
 **optional** | ***DomesticTariffApiDomesticTariffsMagazinePostsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomesticTariffApiDomesticTariffsMagazinePostsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **height** | **optional.Float32**|  | 
 **length** | **optional.Float32**|  | 
 **width** | **optional.Float32**|  | 

### Return type

[**ResponseGetMagazinePostTariffApiResponse**](response.GetMagazinePostTariffAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticTariffsNationalBillMailsGet**
> ResponseGetBillMailTariffApiResponse DomesticTariffsNationalBillMailsGet(ctx, productCode, quantity, weight)
Get National Bill Mail tariff based on product code, weight, and quantity

Calculate the National Bill Mail tariff using the product code, weight, and quantity provided.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**|  | 
  **quantity** | **int32**|  | 
  **weight** | **float32**|  | 

### Return type

[**ResponseGetBillMailTariffApiResponse**](response.GetBillMailTariffAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticTariffsNewspaperBundlesGet**
> ResponseGetNewspaperBundleTariffApiResponse DomesticTariffsNewspaperBundlesGet(ctx, productCode, weight, optional)
Get newspaper bundle tariff based on various parameters

Calculate the newspaper bundle tariff using product details and additional optional parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**|  | 
  **weight** | **float32**|  | 
 **optional** | ***DomesticTariffApiDomesticTariffsNewspaperBundlesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomesticTariffApiDomesticTariffsNewspaperBundlesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aCK** | **optional.String**|  | 
 **aMS** | **optional.String**|  | 
 **iNS** | **optional.String**|  | 
 **rEG** | **optional.String**|  | 
 **vPP** | **optional.String**|  | 
 **height** | **optional.Float32**|  | 
 **length** | **optional.Float32**|  | 
 **width** | **optional.Float32**|  | 

### Return type

[**ResponseGetNewspaperBundleTariffApiResponse**](response.GetNewspaperBundleTariffAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticTariffsNewspapersGet**
> ResponseGetNewspaperTariffApiResponse DomesticTariffsNewspapersGet(ctx, productCode, weight, optional)
Get newspaper tariff based on various parameters

Calculate the newspaper tariff using product details and additional optional parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**|  | 
  **weight** | **float32**|  | 
 **optional** | ***DomesticTariffApiDomesticTariffsNewspapersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomesticTariffApiDomesticTariffsNewspapersGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aCK** | **optional.String**|  | 
 **aMS** | **optional.String**|  | 
 **iNS** | **optional.String**|  | 
 **rEG** | **optional.String**|  | 
 **vPP** | **optional.String**|  | 
 **height** | **optional.Float32**|  | 
 **length** | **optional.Float32**|  | 
 **width** | **optional.Float32**|  | 

### Return type

[**ResponseGetNewspaperTariffApiResponse**](response.GetNewspaperTariffAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticTariffsNonVariantsGet**
> ResponseListNonVariantProductApiResponse DomesticTariffsNonVariantsGet(ctx, productCode)
Get the non-variant product tariff based on the product code

Retrieve the tariff information for a non-variant product using its product code.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product code | 

### Return type

[**ResponseListNonVariantProductApiResponse**](response.ListNonVariantProductAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticTariffsParcelsGet**
> ResponseGetParcelTariffApiResponse DomesticTariffsParcelsGet(ctx, productCode, weight, optional)
Get parcel tariff based on various parameters

Calculate the parcel tariff using product details and various additional parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**|  | 
  **weight** | **float32**|  | 
 **optional** | ***DomesticTariffApiDomesticTariffsParcelsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomesticTariffApiDomesticTariffsParcelsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aCK** | **optional.String**|  | 
 **aMS** | **optional.String**|  | 
 **iNS** | **optional.String**|  | 
 **rEG** | **optional.String**|  | 
 **vPP** | **optional.String**|  | 
 **diameter** | **optional.Float32**|  | 
 **height** | **optional.Float32**|  | 
 **length** | **optional.Float32**|  | 
 **width** | **optional.Float32**|  | 

### Return type

[**ResponseGetParcelTariffApiResponse**](response.GetParcelTariffAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticTariffsPatternSamplePacketsGet**
> ResponseGetPatternSamplePacketTariffApiResponse DomesticTariffsPatternSamplePacketsGet(ctx, productCode, weight, optional)
Get pattern sample packet tariff based on various parameters

Calculate the pattern sample packet tariff using product details and additional optional parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**|  | 
  **weight** | **float32**|  | 
 **optional** | ***DomesticTariffApiDomesticTariffsPatternSamplePacketsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomesticTariffApiDomesticTariffsPatternSamplePacketsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aCK** | **optional.String**|  | 
 **aMS** | **optional.String**|  | 
 **iNS** | **optional.String**|  | 
 **rEG** | **optional.String**|  | 
 **vPP** | **optional.String**|  | 
 **diameter** | **optional.Float32**|  | 
 **height** | **optional.Float32**|  | 
 **length** | **optional.Float32**|  | 
 **width** | **optional.Float32**|  | 

### Return type

[**ResponseGetPatternSamplePacketTariffApiResponse**](response.GetPatternSamplePacketTariffAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticTariffsPostCardsGet**
> ResponseGetPostCardTariffApiResponse DomesticTariffsPostCardsGet(ctx, productCode, optional)
Get Post Card tariff based on product code, weight, length, and width

Calculate the Post Card tariff using the product code, weight, length, and width provided.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**|  | 
 **optional** | ***DomesticTariffApiDomesticTariffsPostCardsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomesticTariffApiDomesticTariffsPostCardsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aCK** | **optional.String**|  | 
 **aMS** | **optional.String**|  | 
 **iNS** | **optional.String**|  | 
 **rEG** | **optional.String**|  | 
 **vPP** | **optional.String**|  | 
 **length** | **optional.Float32**|  | 
 **weight** | **optional.Float32**|  | 
 **width** | **optional.Float32**|  | 

### Return type

[**ResponseGetPostCardTariffApiResponse**](response.GetPostCardTariffAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticTariffsPrintedBooksGet**
> ResponseGetPrintedBookTariffApiResponse DomesticTariffsPrintedBooksGet(ctx, productCode, weight, optional)
Get printed book tariff based on various parameters

Calculate the printed book tariff using product details and additional optional parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**|  | 
  **weight** | **float32**|  | 
 **optional** | ***DomesticTariffApiDomesticTariffsPrintedBooksGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomesticTariffApiDomesticTariffsPrintedBooksGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aCK** | **optional.String**|  | 
 **aMS** | **optional.String**|  | 
 **iNS** | **optional.String**|  | 
 **rEG** | **optional.String**|  | 
 **rEGC** | **optional.String**|  | 
 **vPP** | **optional.String**|  | 
 **diameter** | **optional.Float32**|  | 
 **height** | **optional.Float32**|  | 
 **length** | **optional.Float32**|  | 
 **width** | **optional.Float32**|  | 

### Return type

[**ResponseGetPrintedBookTariffApiResponse**](response.GetPrintedBookTariffAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticTariffsRetailsGet**
> ResponseListRetailPostApiResponse DomesticTariffsRetailsGet(ctx, productCode)
Get the retail post product tariff based on the product code

Retrieve the tariff information for a retail post product using its product code.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productCode** | **string**| Product code | 

### Return type

[**ResponseListRetailPostApiResponse**](response.ListRetailPostAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticTariffsSpeedPostsGet**
> ResponseGetSpeedLetterTariffApiResponse DomesticTariffsSpeedPostsGet(ctx, destinationPincode, productCode, sourcePincode, weight, optional)
Get speed letter tariff based on various parameters

Calculate the speed letter tariff using product details and additional optional parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **destinationPincode** | **int32**|  | 
  **productCode** | **string**|  | 
  **sourcePincode** | **int32**|  | 
  **weight** | **float32**|  | 
 **optional** | ***DomesticTariffApiDomesticTariffsSpeedPostsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomesticTariffApiDomesticTariffsSpeedPostsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **aCK** | **optional.String**|  | 
 **aMS** | **optional.String**|  | 
 **iNS** | **optional.String**|  | 
 **iNSC** | **optional.String**|  | 
 **pOD** | **optional.String**|  | 
 **rEG** | **optional.String**|  | 
 **vPP** | **optional.String**|  | 
 **diameter** | **optional.Float32**|  | 
 **height** | **optional.Float32**|  | 
 **length** | **optional.Float32**|  | 
 **width** | **optional.Float32**|  | 

### Return type

[**ResponseGetSpeedLetterTariffApiResponse**](response.GetSpeedLetterTariffAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

