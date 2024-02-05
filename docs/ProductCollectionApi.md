# \ProductCollectionApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProductCollection**](ProductCollectionApi.md#CreateProductCollection) | **Put** /obp/v5.1.0/banks/{BANK_ID}/product-collections/{COLLECTION_CODE} | Create Product Collection
[**GetProductCollection**](ProductCollectionApi.md#GetProductCollection) | **Get** /obp/v5.1.0/banks/{BANK_ID}/product-collections/{COLLECTION_CODE} | Get Product Collection


# **CreateProductCollection**
> ProductCollectionsJsonV310 CreateProductCollection(ctx, body, cOLLECTIONCODE, bANKID)
Create Product Collection

<p>Create or Update a Product Collection at the Bank.</p><p>Use Product Collections to create Product &quot;Baskets&quot;, &quot;Portfolios&quot;, &quot;Indices&quot;, &quot;Collections&quot;, &quot;Underlyings-lists&quot;, &quot;Buckets&quot; etc. etc.</p><p>There is a many to many relationship between Products and Product Collections:</p><ul><li><p>A Product can exist in many Collections</p></li><li><p>A Collection can contain many Products.</p></li></ul><p>A collection has collection code, one parent Product and one or more child Products.</p><p>Product hiearchy vs Product Collections:</p><ul><li><p>You can define a hierarchy of products - so that a child Product inherits attributes of its parent Product -  using the parent_product_code in Product.</p></li><li><p>You can define a collection (also known as baskets or buckets) of products using Product Collections.</p></li></ul><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PutProductCollectionsV310**](PutProductCollectionsV310.md)| PutProductCollectionsV310 object that needs to be added. | 
  **cOLLECTIONCODE** | **string**| the collection code | 
  **bANKID** | **string**| The bank id | 

### Return type

[**ProductCollectionsJsonV310**](ProductCollectionsJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProductCollection**
> ProductCollectionJsonTreeV310 GetProductCollection(ctx, cOLLECTIONCODE, bANKID)
Get Product Collection

<p>Returns information about the financial Product Collection specified by BANK_ID and COLLECTION_CODE:</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cOLLECTIONCODE** | **string**| the collection code | 
  **bANKID** | **string**| The bank id | 

### Return type

[**ProductCollectionJsonTreeV310**](ProductCollectionJsonTreeV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

