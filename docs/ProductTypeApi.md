# \ProductTypeApi

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductsTypesGet**](ProductTypeApi.md#ProductsTypesGet) | **Get** /products-types | List all product types
[**ProductsTypesProductTypeIdGet**](ProductTypeApi.md#ProductsTypesProductTypeIdGet) | **Get** /products-types/{product-type-id} | Retrieve product type by ID


# **ProductsTypesGet**
> ResponseListProductTypeApiResponse ProductsTypesGet(ctx, optional)
List all product types

Fetches a paginated list of all available product types.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductTypeApiProductsTypesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductTypeApiProductsTypesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 

### Return type

[**ResponseListProductTypeApiResponse**](response.ListProductTypeAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductsTypesProductTypeIdGet**
> ResponseGetProductTypeApiResponse ProductsTypesProductTypeIdGet(ctx, productTypeId)
Retrieve product type by ID

Fetches a specific product type based on the provided product type ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productTypeId** | **int32**| Product Type ID | 

### Return type

[**ResponseGetProductTypeApiResponse**](response.GetProductTypeAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

