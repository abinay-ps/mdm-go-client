# \ProductApi

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductsGet**](ProductApi.md#ProductsGet) | **Get** /products | List product details based on various filters
[**ProductsPost**](ProductApi.md#ProductsPost) | **Post** /products | Create a new product
[**ProductsProductIdGet**](ProductApi.md#ProductsProductIdGet) | **Get** /products/{product-id} | Get product details by Product ID
[**ProductsProductIdPut**](ProductApi.md#ProductsProductIdPut) | **Put** /products/{product-id} | Update product details


# **ProductsGet**
> ResponseListProductWithVasDetailsApiResponse ProductsGet(ctx, listType, optional)
List product details based on various filters

Retrieves a list of product details based on the specified list type (`product-code`, `product-name`, or `product-type`) and optional flags for modified or stamp status with pagination support.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listType** | **string**|  | 
 **optional** | ***ProductApiProductsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductApiProductsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | 
 **modifiedFlag** | **optional.Bool**|  | 
 **orderBy** | **optional.String**|  | 
 **productCode** | **optional.String**|  | 
 **productName** | **optional.String**|  | 
 **productTypeCode** | **optional.String**|  | 
 **skip** | **optional.Int32**|  | 
 **sortType** | **optional.String**|  | 
 **stampFlag** | **optional.Bool**|  | 

### Return type

[**ResponseListProductWithVasDetailsApiResponse**](response.ListProductWithVASDetailsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductsPost**
> ResponseCreateProductApiResponse ProductsPost(ctx, body)
Create a new product

Create a new product with the provided details, including Product Code, Product Name, Category ID, and other optional fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**HandlerCreateProductRequest**](HandlerCreateProductRequest.md)| Create product request | 

### Return type

[**ResponseCreateProductApiResponse**](response.CreateProductAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductsProductIdGet**
> ResponseGetProductWithGlDetailsApiResponse ProductsProductIdGet(ctx, productId)
Get product details by Product ID

Retrieves product details along with general ledger details using the provided Product ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | **int32**| Product ID | 

### Return type

[**ResponseGetProductWithGlDetailsApiResponse**](response.GetProductWithGLDetailsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductsProductIdPut**
> ResponseUpdateProductDetailsApiResponse ProductsProductIdPut(ctx, productId, body)
Update product details

Update details of an existing product using its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | **int32**| ID of the product to update | 
  **body** | [**HandlerUpdateProductDetailsRequest**](HandlerUpdateProductDetailsRequest.md)| Updated product details | 

### Return type

[**ResponseUpdateProductDetailsApiResponse**](response.UpdateProductDetailsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

