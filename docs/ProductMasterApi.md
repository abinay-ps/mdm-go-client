# \ProductMasterApi

All URIs are relative to *https://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductCheckproductavailabilityinanofficeGet**](ProductMasterApi.md#ProductCheckproductavailabilityinanofficeGet) | **Get** /product/checkproductavailabilityinanoffice | Get Product availability in an Office
[**ProductCreateproductPost**](ProductMasterApi.md#ProductCreateproductPost) | **Post** /product/createproduct | Create a Product
[**ProductMapofficesandprodcutsPost**](ProductMasterApi.md#ProductMapofficesandprodcutsPost) | **Post** /product/mapofficesandprodcuts | Create a Product and Office mapping
[**ProductProductdetailsGet**](ProductMasterApi.md#ProductProductdetailsGet) | **Get** /product/productdetails | Get Product Details
[**ProductProducttypeGet**](ProductMasterApi.md#ProductProducttypeGet) | **Get** /product/producttype | Get Product Types
[**ProductUpdateproductProductidProductidPut**](ProductMasterApi.md#ProductUpdateproductProductidProductidPut) | **Put** /product/updateproduct/productid/{productid} | Update Product Details


# **ProductCheckproductavailabilityinanofficeGet**
> HandlerResponse ProductCheckproductavailabilityinanofficeGet(ctx, optional)
Get Product availability in an Office

Get Product availability in an Office

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductMasterApiProductCheckproductavailabilityinanofficeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductMasterApiProductCheckproductavailabilityinanofficeGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productID** | **optional.Int32**| Product ID | 
 **officeID** | **optional.Int32**| Office ID | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductCreateproductPost**
> HandlerResponse ProductCreateproductPost(ctx, productMaster)
Create a Product

Creates a new Product

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productMaster** | [**DomainProductMaster**](DomainProductMaster.md)| New Product details | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductMapofficesandprodcutsPost**
> HandlerResponse ProductMapofficesandprodcutsPost(ctx, offices)
Create a Product and Office mapping

Creates a Product and Office mapping

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **offices** | [**HandlerOffices**](HandlerOffices.md)| Office and Product details | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductProductdetailsGet**
> HandlerResponse ProductProductdetailsGet(ctx, optional)
Get Product Details

Get Product Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductMasterApiProductProductdetailsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductMasterApiProductProductdetailsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productID** | **optional.Int32**| Product ID | 
 **productCode** | **optional.String**| Product Code | 
 **productName** | **optional.String**| Product Name | 
 **productTypeCode** | **optional.String**| Product Type Code | 
 **modifiedFlag** | **optional.Bool**| Modified Flag | 
 **stampFlag** | **optional.Bool**| Stamp Flag | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductProducttypeGet**
> HandlerResponse ProductProducttypeGet(ctx, optional)
Get Product Types

Get Product Types

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductMasterApiProductProducttypeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductMasterApiProductProducttypeGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productTypeCode** | **optional.String**| Product Type Code | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductUpdateproductProductidProductidPut**
> HandlerResponse ProductUpdateproductProductidProductidPut(ctx, productid, optional)
Update Product Details

Update Product Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productid** | **int32**| Product ID | 
 **optional** | ***ProductMasterApiProductUpdateproductProductidProductidPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductMasterApiProductUpdateproductProductidProductidPutOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **productMaster** | **optional.Interface{}**| Product Details | 

### Return type

[**HandlerResponse**](handler.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

