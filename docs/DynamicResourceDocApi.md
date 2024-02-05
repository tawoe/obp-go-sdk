# \DynamicResourceDocApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BuildDynamicEndpointTemplate**](DynamicResourceDocApi.md#BuildDynamicEndpointTemplate) | **Post** /obp/v5.1.0/management/dynamic-resource-docs/endpoint-code | Create Dynamic Resource Doc endpoint code
[**CreateBankLevelDynamicResourceDoc**](DynamicResourceDocApi.md#CreateBankLevelDynamicResourceDoc) | **Post** /obp/v5.1.0/management/banks/{BANK_ID}/dynamic-resource-docs | Create Bank Level Dynamic Resource Doc
[**CreateDynamicResourceDoc**](DynamicResourceDocApi.md#CreateDynamicResourceDoc) | **Post** /obp/v5.1.0/management/dynamic-resource-docs | Create Dynamic Resource Doc
[**DeleteBankLevelDynamicResourceDoc**](DynamicResourceDocApi.md#DeleteBankLevelDynamicResourceDoc) | **Delete** /obp/v5.1.0/management/banks/{BANK_ID}/dynamic-resource-docs/DYNAMIC-RESOURCE-DOC-ID | Delete Bank Level Dynamic Resource Doc
[**DeleteDynamicResourceDoc**](DynamicResourceDocApi.md#DeleteDynamicResourceDoc) | **Delete** /obp/v5.1.0/management/dynamic-resource-docs/DYNAMIC-RESOURCE-DOC-ID | Delete Dynamic Resource Doc
[**GetAllBankLevelDynamicResourceDocs**](DynamicResourceDocApi.md#GetAllBankLevelDynamicResourceDocs) | **Get** /obp/v5.1.0/management/banks/{BANK_ID}/dynamic-resource-docs | Get all Bank Level Dynamic Resource Docs
[**GetAllDynamicResourceDocs**](DynamicResourceDocApi.md#GetAllDynamicResourceDocs) | **Get** /obp/v5.1.0/management/dynamic-resource-docs | Get all Dynamic Resource Docs
[**GetBankLevelDynamicResourceDoc**](DynamicResourceDocApi.md#GetBankLevelDynamicResourceDoc) | **Get** /obp/v5.1.0/management/banks/{BANK_ID}/dynamic-resource-docs/DYNAMIC-RESOURCE-DOC-ID | Get Bank Level Dynamic Resource Doc by Id
[**GetDynamicResourceDoc**](DynamicResourceDocApi.md#GetDynamicResourceDoc) | **Get** /obp/v5.1.0/management/dynamic-resource-docs/DYNAMIC-RESOURCE-DOC-ID | Get Dynamic Resource Doc by Id
[**TestDynamicResourceDoc**](DynamicResourceDocApi.md#TestDynamicResourceDoc) | **Post** /test-dynamic-resource-doc/my_user/MY_USER_ID | A test endpoint
[**UpdateBankLevelDynamicResourceDoc**](DynamicResourceDocApi.md#UpdateBankLevelDynamicResourceDoc) | **Put** /obp/v5.1.0/management/banks/{BANK_ID}/dynamic-resource-docs/DYNAMIC-RESOURCE-DOC-ID | Update Bank Level Dynamic Resource Doc
[**UpdateDynamicResourceDoc**](DynamicResourceDocApi.md#UpdateDynamicResourceDoc) | **Put** /obp/v5.1.0/management/dynamic-resource-docs/DYNAMIC-RESOURCE-DOC-ID | Update Dynamic Resource Doc


# **BuildDynamicEndpointTemplate**
> JsonCodeTemplateJson BuildDynamicEndpointTemplate(ctx, body)
Create Dynamic Resource Doc endpoint code

<p>Create a Dynamic Resource Doc endpoint code.</p><p>copy the response and past to PractiseEndpoint, So you can have the benefits of<br />auto compilation and debug</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ResourceDocFragment**](ResourceDocFragment.md)| ResourceDocFragment object that needs to be added. | 

### Return type

[**JsonCodeTemplateJson**](JsonCodeTemplateJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBankLevelDynamicResourceDoc**
> JsonDynamicResourceDoc CreateBankLevelDynamicResourceDoc(ctx, body, bANKID)
Create Bank Level Dynamic Resource Doc

<p>Create a Bank Level Dynamic Resource Doc.</p><p>The connector_method_body is URL-encoded format String</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**JsonDynamicResourceDoc**](JsonDynamicResourceDoc.md)| JsonDynamicResourceDoc object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**JsonDynamicResourceDoc**](JsonDynamicResourceDoc.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDynamicResourceDoc**
> JsonDynamicResourceDoc CreateDynamicResourceDoc(ctx, body)
Create Dynamic Resource Doc

<p>Create a Dynamic Resource Doc.</p><p>The connector_method_body is URL-encoded format String</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**JsonDynamicResourceDoc**](JsonDynamicResourceDoc.md)| JsonDynamicResourceDoc object that needs to be added. | 

### Return type

[**JsonDynamicResourceDoc**](JsonDynamicResourceDoc.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBankLevelDynamicResourceDoc**
> bool DeleteBankLevelDynamicResourceDoc(ctx, bANKID)
Delete Bank Level Dynamic Resource Doc

<p>Delete a Bank Level Dynamic Resource Doc.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

**bool**

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDynamicResourceDoc**
> bool DeleteDynamicResourceDoc(ctx, )
Delete Dynamic Resource Doc

<p>Delete a Dynamic Resource Doc.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

**bool**

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllBankLevelDynamicResourceDocs**
> InlineResponse2006 GetAllBankLevelDynamicResourceDocs(ctx, bANKID)
Get all Bank Level Dynamic Resource Docs

<p>Get all Bank Level Dynamic Resource Docs.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllDynamicResourceDocs**
> InlineResponse2006 GetAllDynamicResourceDocs(ctx, )
Get all Dynamic Resource Docs

<p>Get all Dynamic Resource Docs.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBankLevelDynamicResourceDoc**
> JsonDynamicResourceDoc GetBankLevelDynamicResourceDoc(ctx, bANKID)
Get Bank Level Dynamic Resource Doc by Id

<p>Get a Bank Level Dynamic Resource Doc by DYNAMIC-RESOURCE-DOC-ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**JsonDynamicResourceDoc**](JsonDynamicResourceDoc.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDynamicResourceDoc**
> JsonDynamicResourceDoc GetDynamicResourceDoc(ctx, )
Get Dynamic Resource Doc by Id

<p>Get a Dynamic Resource Doc by DYNAMIC-RESOURCE-DOC-ID.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**JsonDynamicResourceDoc**](JsonDynamicResourceDoc.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestDynamicResourceDoc**
> RequestRootJsonClass TestDynamicResourceDoc(ctx, body)
A test endpoint

<p>A test endpoint.</p><p>Just for debug method body of dynamic resource doc.<br />better watch the following introduction video first<br />* <a href=\"https://vimeo.com/623381607\">Dynamic resourceDoc version1</a></p><p>The endpoint return the response from PractiseEndpoint code.<br />Here, code.api.DynamicEndpoints.dynamic.practise.PractiseEndpoint.process<br />You can test the method body grammar, and try the business logic, but need to restart the OBP-API code .</p><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestRootJsonClass**](RequestRootJsonClass.md)| RequestRootJsonClass object that needs to be added. | 

### Return type

[**RequestRootJsonClass**](RequestRootJsonClass.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBankLevelDynamicResourceDoc**
> JsonDynamicResourceDoc UpdateBankLevelDynamicResourceDoc(ctx, body, bANKID)
Update Bank Level Dynamic Resource Doc

<p>Update a Bank Level Dynamic Resource Doc.</p><p>The connector_method_body is URL-encoded format String</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**JsonDynamicResourceDoc**](JsonDynamicResourceDoc.md)| JsonDynamicResourceDoc object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**JsonDynamicResourceDoc**](JsonDynamicResourceDoc.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDynamicResourceDoc**
> JsonDynamicResourceDoc UpdateDynamicResourceDoc(ctx, body)
Update Dynamic Resource Doc

<p>Update a Dynamic Resource Doc.</p><p>The connector_method_body is URL-encoded format String</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**JsonDynamicResourceDoc**](JsonDynamicResourceDoc.md)| JsonDynamicResourceDoc object that needs to be added. | 

### Return type

[**JsonDynamicResourceDoc**](JsonDynamicResourceDoc.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

