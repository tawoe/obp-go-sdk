# \ProductApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateProductAttributeDefinition**](ProductApi.md#CreateOrUpdateProductAttributeDefinition) | **Put** /obp/v5.1.0/banks/{BANK_ID}/attribute-definitions/product | Create or Update Product Attribute Definition
[**CreateProduct**](ProductApi.md#CreateProduct) | **Put** /obp/v5.1.0/banks/{BANK_ID}/products/{PRODUCT_CODE} | Create Product
[**CreateProductAttribute**](ProductApi.md#CreateProductAttribute) | **Post** /obp/v5.1.0/banks/{BANK_ID}/products/{PRODUCT_CODE}/attribute | Create Product Attribute
[**CreateProductCollection**](ProductApi.md#CreateProductCollection) | **Put** /obp/v5.1.0/banks/{BANK_ID}/product-collections/{COLLECTION_CODE} | Create Product Collection
[**CreateProductFee**](ProductApi.md#CreateProductFee) | **Post** /obp/v5.1.0/banks/{BANK_ID}/products/{PRODUCT_CODE}/fee | Create Product Fee
[**DeleteProductAttribute**](ProductApi.md#DeleteProductAttribute) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/products/{PRODUCT_CODE}/attributes/{PRODUCT_ATTRIBUTE_ID} | Delete Product Attribute
[**DeleteProductAttributeDefinition**](ProductApi.md#DeleteProductAttributeDefinition) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/attribute-definitions/ATTRIBUTE_DEFINITION_ID/product | Delete Product Attribute Definition
[**DeleteProductCascade**](ProductApi.md#DeleteProductCascade) | **Delete** /obp/v5.1.0/management/cascading/banks/{BANK_ID}/products/{PRODUCT_CODE} | Delete Product Cascade
[**DeleteProductFee**](ProductApi.md#DeleteProductFee) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/products/{PRODUCT_CODE}/fees/PRODUCT_FEE_ID | Delete Product Fee
[**GetProduct**](ProductApi.md#GetProduct) | **Get** /obp/v5.1.0/banks/{BANK_ID}/products/{PRODUCT_CODE} | Get Bank Product
[**GetProductAttribute**](ProductApi.md#GetProductAttribute) | **Get** /obp/v5.1.0/banks/{BANK_ID}/products/{PRODUCT_CODE}/attributes/{PRODUCT_ATTRIBUTE_ID} | Get Product Attribute
[**GetProductAttributeDefinition**](ProductApi.md#GetProductAttributeDefinition) | **Get** /obp/v5.1.0/banks/{BANK_ID}/attribute-definitions/product | Get Product Attribute Definition
[**GetProductCollection**](ProductApi.md#GetProductCollection) | **Get** /obp/v5.1.0/banks/{BANK_ID}/product-collections/{COLLECTION_CODE} | Get Product Collection
[**GetProductFee**](ProductApi.md#GetProductFee) | **Get** /obp/v5.1.0/banks/{BANK_ID}/products/{PRODUCT_CODE}/fees/PRODUCT_FEE_ID | Get Product Fee
[**GetProductFees**](ProductApi.md#GetProductFees) | **Get** /obp/v5.1.0/banks/{BANK_ID}/products/{PRODUCT_CODE}/fees | Get Product Fees
[**GetProductTree**](ProductApi.md#GetProductTree) | **Get** /obp/v5.1.0/banks/{BANK_ID}/product-tree/{PRODUCT_CODE} | Get Product Tree
[**GetProducts**](ProductApi.md#GetProducts) | **Get** /obp/v5.1.0/banks/{BANK_ID}/products | Get Products
[**UpdateProductAttribute**](ProductApi.md#UpdateProductAttribute) | **Put** /obp/v5.1.0/banks/{BANK_ID}/products/{PRODUCT_CODE}/attributes/{PRODUCT_ATTRIBUTE_ID} | Update Product Attribute
[**UpdateProductFee**](ProductApi.md#UpdateProductFee) | **Put** /obp/v5.1.0/banks/{BANK_ID}/products/{PRODUCT_CODE}/fees/PRODUCT_FEE_ID | Update Product Fee


# **CreateOrUpdateProductAttributeDefinition**
> AttributeDefinitionResponseJsonV400 CreateOrUpdateProductAttributeDefinition(ctx, body, bANKID)
Create or Update Product Attribute Definition

<p>Create or Update Product Attribute Definition</p><p>The category field must be Product</p><p>The type field must be one of; DOUBLE, STRING, INTEGER and DATE_WITH_DAY</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AttributeDefinitionJsonV400**](AttributeDefinitionJsonV400.md)| AttributeDefinitionJsonV400 object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**AttributeDefinitionResponseJsonV400**](AttributeDefinitionResponseJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProduct**
> ProductJsonV400 CreateProduct(ctx, body, pRODUCTCODE, bANKID)
Create Product

<p>Create or Update Product for the Bank.</p><p>Typical Super Family values / Asset classes are:</p><p>Debt<br />Equity<br />FX<br />Commodity<br />Derivative</p><p>Product hiearchy vs Product Collections:</p><ul><li><p>You can define a hierarchy of products - so that a child Product inherits attributes of its parent Product -  using the parent_product_code in Product.</p></li><li><p>You can define a collection (also known as baskets or buckets) of products using Product Collections.</p></li></ul><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PutProductJsonV500**](PutProductJsonV500.md)| PutProductJsonV500 object that needs to be added. | 
  **pRODUCTCODE** | **string**| the product code | 
  **bANKID** | **string**| The bank id | 

### Return type

[**ProductJsonV400**](ProductJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProductAttribute**
> ProductAttributeResponseJsonV400 CreateProductAttribute(ctx, body, pRODUCTCODE, bANKID)
Create Product Attribute

<p>Create Product Attribute</p><p>Product Attributes are used to describe a financial Product with a list of typed key value pairs.</p><p>Each Product Attribute is linked to its Product by PRODUCT_CODE</p><p>Typical product attributes might be:</p><p>ISIN (for International bonds)<br />VKN (for German bonds)<br />REDCODE (markit short code for credit derivative)<br />LOAN_ID (e.g. used for Anacredit reporting)</p><p>ISSUE_DATE (When the bond was issued in the market)<br />MATURITY_DATE (End of life time of a product)<br />TRADABLE</p><p>See <a href=\"http://www.fpml.org/\">FPML</a> for more examples.</p><p>The type field must be one of &quot;STRING&quot;, &quot;INTEGER&quot;, &quot;DOUBLE&quot; or DATE_WITH_DAY&quot;</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProductAttributeJsonV400**](ProductAttributeJsonV400.md)| ProductAttributeJsonV400 object that needs to be added. | 
  **pRODUCTCODE** | **string**| the product code | 
  **bANKID** | **string**| The bank id | 

### Return type

[**ProductAttributeResponseJsonV400**](ProductAttributeResponseJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

# **CreateProductFee**
> ProductFeeResponseJsonV400 CreateProductFee(ctx, body, pRODUCTCODE, bANKID)
Create Product Fee

<p>Create Product Fee</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProductFeeJsonV400**](ProductFeeJsonV400.md)| ProductFeeJsonV400 object that needs to be added. | 
  **pRODUCTCODE** | **string**| the product code | 
  **bANKID** | **string**| The bank id | 

### Return type

[**ProductFeeResponseJsonV400**](ProductFeeResponseJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProductAttribute**
> DeleteProductAttribute(ctx, pRODUCTATTRIBUTEID, pRODUCTCODE, bANKID)
Delete Product Attribute

<p>Delete Product Attribute</p><p>Product Attributes are used to describe a financial Product with a list of typed key value pairs.</p><p>Each Product Attribute is linked to its Product by PRODUCT_CODE</p><p>Delete a Product Attribute by its id.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pRODUCTATTRIBUTEID** | **string**| the product attribute id | 
  **pRODUCTCODE** | **string**| the product code | 
  **bANKID** | **string**| The bank id | 

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProductAttributeDefinition**
> DeleteProductAttributeDefinition(ctx, bANKID)
Delete Product Attribute Definition

<p>Delete Product Attribute Definition by ATTRIBUTE_DEFINITION_ID</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProductCascade**
> DeleteProductCascade(ctx, pRODUCTCODE, bANKID)
Delete Product Cascade

<p>Delete a Product Cascade specified by PRODUCT_CODE.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pRODUCTCODE** | **string**| the product code | 
  **bANKID** | **string**| The bank id | 

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProductFee**
> bool DeleteProductFee(ctx, pRODUCTCODE, bANKID)
Delete Product Fee

<p>Delete Product Fee</p><p>Delete one product fee by its id.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pRODUCTCODE** | **string**| the product code | 
  **bANKID** | **string**| The bank id | 

### Return type

**bool**

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProduct**
> ProductJsonV400 GetProduct(ctx, pRODUCTCODE, bANKID)
Get Bank Product

<p>Returns information about a financial Product offered by the bank specified by BANK_ID and PRODUCT_CODE including:</p><ul><li>Name</li><li>Code</li><li>Parent Product Code</li><li>More info URL</li><li>Description</li><li>Terms and Conditions</li><li>Description</li><li>Meta</li><li>Attributes</li><li>Fees</li></ul><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pRODUCTCODE** | **string**| the product code | 
  **bANKID** | **string**| The bank id | 

### Return type

[**ProductJsonV400**](ProductJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProductAttribute**
> ProductAttributeResponseJsonV400 GetProductAttribute(ctx, pRODUCTATTRIBUTEID, pRODUCTCODE, bANKID)
Get Product Attribute

<p>Get Product Attribute</p><p>Product Attributes are used to describe a financial Product with a list of typed key value pairs.</p><p>Each Product Attribute is linked to its Product by PRODUCT_CODE</p><p>Get one product attribute by its id.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pRODUCTATTRIBUTEID** | **string**| the product attribute id | 
  **pRODUCTCODE** | **string**| the product code | 
  **bANKID** | **string**| The bank id | 

### Return type

[**ProductAttributeResponseJsonV400**](ProductAttributeResponseJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProductAttributeDefinition**
> AttributeDefinitionsResponseJsonV400 GetProductAttributeDefinition(ctx, bANKID)
Get Product Attribute Definition

<p>Get Product Attribute Definition</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**AttributeDefinitionsResponseJsonV400**](AttributeDefinitionsResponseJsonV400.md)

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

# **GetProductFee**
> ProductFeeResponseJsonV400 GetProductFee(ctx, pRODUCTCODE, bANKID)
Get Product Fee

<p>Get Product Fee</p><p>Get one product fee by its id.</p><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pRODUCTCODE** | **string**| the product code | 
  **bANKID** | **string**| The bank id | 

### Return type

[**ProductFeeResponseJsonV400**](ProductFeeResponseJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProductFees**
> ProductFeesResponseJsonV400 GetProductFees(ctx, pRODUCTCODE, bANKID)
Get Product Fees

<p>Get Product Fees</p><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pRODUCTCODE** | **string**| the product code | 
  **bANKID** | **string**| The bank id | 

### Return type

[**ProductFeesResponseJsonV400**](ProductFeesResponseJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProductTree**
> ProductTreeJsonV310 GetProductTree(ctx, pRODUCTCODE, bANKID)
Get Product Tree

<p>Returns information about a particular financial product specified by BANK_ID and PRODUCT_CODE<br />and it's parent product(s) recursively as specified by parent_product_code.</p><p>Each product includes the following information.</p><ul><li>Name</li><li>Code</li><li>Parent Product Code</li><li>Category</li><li>Family</li><li>Super Family</li><li>More info URL</li><li>Description</li><li>Terms and Conditions</li><li>License: The licence under which this product data is released. Licence can be an Open Data licence such as Open Data Commons Public Domain Dedication and License (PDDL) or Copyright etc.</li></ul><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pRODUCTCODE** | **string**| the product code | 
  **bANKID** | **string**| The bank id | 

### Return type

[**ProductTreeJsonV310**](ProductTreeJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProducts**
> ProductsJsonV400 GetProducts(ctx, bANKID)
Get Products

<p>Returns information about the financial products offered by a bank specified by BANK_ID including:</p><ul><li>Name</li><li>Code</li><li>Parent Product Code</li><li>More info URL</li><li>Terms And Conditions URL</li><li>Description</li><li>Terms and Conditions</li><li>License the data under this endpoint is released under</li></ul><p>Can filter with attributes name and values.<br />URL params example: /banks/some-bank-id/products?manager=John&amp;count=8</p><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**ProductsJsonV400**](ProductsJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProductAttribute**
> ProductAttributeResponseJsonV400 UpdateProductAttribute(ctx, body, pRODUCTATTRIBUTEID, pRODUCTCODE, bANKID)
Update Product Attribute

<p>Update Product Attribute.</p><p>Product Attributes are used to describe a financial Product with a list of typed key value pairs.</p><p>Each Product Attribute is linked to its Product by PRODUCT_CODE</p><p>Update one Product Attribute by its id.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProductAttributeJsonV400**](ProductAttributeJsonV400.md)| ProductAttributeJsonV400 object that needs to be added. | 
  **pRODUCTATTRIBUTEID** | **string**| the product attribute id | 
  **pRODUCTCODE** | **string**| the product code | 
  **bANKID** | **string**| The bank id | 

### Return type

[**ProductAttributeResponseJsonV400**](ProductAttributeResponseJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProductFee**
> ProductFeeResponseJsonV400 UpdateProductFee(ctx, body, pRODUCTCODE, bANKID)
Update Product Fee

<p>Update Product Fee.</p><p>Update one Product Fee by its id.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProductFeeJsonV400**](ProductFeeJsonV400.md)| ProductFeeJsonV400 object that needs to be added. | 
  **pRODUCTCODE** | **string**| the product code | 
  **bANKID** | **string**| The bank id | 

### Return type

[**ProductFeeResponseJsonV400**](ProductFeeResponseJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

