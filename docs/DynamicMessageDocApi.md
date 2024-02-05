# \DynamicMessageDocApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBankLevelDynamicMessageDoc**](DynamicMessageDocApi.md#CreateBankLevelDynamicMessageDoc) | **Post** /obp/v5.1.0/management/banks/{BANK_ID}/dynamic-message-docs | Create Bank Level Dynamic Message Doc
[**CreateDynamicMessageDoc**](DynamicMessageDocApi.md#CreateDynamicMessageDoc) | **Post** /obp/v5.1.0/management/dynamic-message-docs | Create Dynamic Message Doc
[**DeleteBankLevelDynamicMessageDoc**](DynamicMessageDocApi.md#DeleteBankLevelDynamicMessageDoc) | **Delete** /obp/v5.1.0/management/banks/{BANK_ID}/dynamic-message-docs/DYNAMIC_MESSAGE_DOC_ID | Delete Bank Level Dynamic Message Doc
[**DeleteDynamicMessageDoc**](DynamicMessageDocApi.md#DeleteDynamicMessageDoc) | **Delete** /obp/v5.1.0/management/dynamic-message-docs/DYNAMIC_MESSAGE_DOC_ID | Delete Dynamic Message Doc
[**GetAllBankLevelDynamicMessageDocs**](DynamicMessageDocApi.md#GetAllBankLevelDynamicMessageDocs) | **Get** /obp/v5.1.0/management/banks/{BANK_ID}/dynamic-message-docs | Get all Bank Level Dynamic Message Docs
[**GetAllDynamicMessageDocs**](DynamicMessageDocApi.md#GetAllDynamicMessageDocs) | **Get** /obp/v5.1.0/management/dynamic-message-docs | Get all Dynamic Message Docs
[**GetBankLevelDynamicMessageDoc**](DynamicMessageDocApi.md#GetBankLevelDynamicMessageDoc) | **Get** /obp/v5.1.0/management/banks/{BANK_ID}/dynamic-message-docs/DYNAMIC_MESSAGE_DOC_ID | Get Bank Level Dynamic Message Doc
[**GetDynamicMessageDoc**](DynamicMessageDocApi.md#GetDynamicMessageDoc) | **Get** /obp/v5.1.0/management/dynamic-message-docs/DYNAMIC_MESSAGE_DOC_ID | Get Dynamic Message Doc
[**UpdateBankLevelDynamicMessageDoc**](DynamicMessageDocApi.md#UpdateBankLevelDynamicMessageDoc) | **Put** /obp/v5.1.0/management/banks/{BANK_ID}/dynamic-message-docs/DYNAMIC_MESSAGE_DOC_ID | Update Bank Level Dynamic Message Doc
[**UpdateDynamicMessageDoc**](DynamicMessageDocApi.md#UpdateDynamicMessageDoc) | **Put** /obp/v5.1.0/management/dynamic-message-docs/DYNAMIC_MESSAGE_DOC_ID | Update Dynamic Message Doc


# **CreateBankLevelDynamicMessageDoc**
> JsonDynamicMessageDoc CreateBankLevelDynamicMessageDoc(ctx, body, bANKID)
Create Bank Level Dynamic Message Doc

<p>Create a Bank Level Dynamic Message Doc.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**JsonDynamicMessageDoc**](JsonDynamicMessageDoc.md)| JsonDynamicMessageDoc object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**JsonDynamicMessageDoc**](JsonDynamicMessageDoc.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDynamicMessageDoc**
> JsonDynamicMessageDoc CreateDynamicMessageDoc(ctx, body)
Create Dynamic Message Doc

<p>Create a Dynamic Message Doc.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**JsonDynamicMessageDoc**](JsonDynamicMessageDoc.md)| JsonDynamicMessageDoc object that needs to be added. | 

### Return type

[**JsonDynamicMessageDoc**](JsonDynamicMessageDoc.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBankLevelDynamicMessageDoc**
> bool DeleteBankLevelDynamicMessageDoc(ctx, bANKID)
Delete Bank Level Dynamic Message Doc

<p>Delete a Bank Level Dynamic Message Doc.</p><p>Authentication is Mandatory</p>

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

# **DeleteDynamicMessageDoc**
> bool DeleteDynamicMessageDoc(ctx, )
Delete Dynamic Message Doc

<p>Delete a Dynamic Message Doc.</p><p>Authentication is Mandatory</p>

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

# **GetAllBankLevelDynamicMessageDocs**
> InlineResponse2005 GetAllBankLevelDynamicMessageDocs(ctx, bANKID)
Get all Bank Level Dynamic Message Docs

<p>Get all Bank Level Dynamic Message Docs.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllDynamicMessageDocs**
> InlineResponse2005 GetAllDynamicMessageDocs(ctx, )
Get all Dynamic Message Docs

<p>Get all Dynamic Message Docs.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBankLevelDynamicMessageDoc**
> JsonDynamicMessageDoc GetBankLevelDynamicMessageDoc(ctx, bANKID)
Get Bank Level Dynamic Message Doc

<p>Get a Bank Level Dynamic Message Doc by DYNAMIC_MESSAGE_DOC_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**JsonDynamicMessageDoc**](JsonDynamicMessageDoc.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDynamicMessageDoc**
> JsonDynamicMessageDoc GetDynamicMessageDoc(ctx, )
Get Dynamic Message Doc

<p>Get a Dynamic Message Doc by DYNAMIC_MESSAGE_DOC_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**JsonDynamicMessageDoc**](JsonDynamicMessageDoc.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBankLevelDynamicMessageDoc**
> JsonDynamicMessageDoc UpdateBankLevelDynamicMessageDoc(ctx, body, bANKID)
Update Bank Level Dynamic Message Doc

<p>Update a Bank Level Dynamic Message Doc.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**JsonDynamicMessageDoc**](JsonDynamicMessageDoc.md)| JsonDynamicMessageDoc object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**JsonDynamicMessageDoc**](JsonDynamicMessageDoc.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDynamicMessageDoc**
> JsonDynamicMessageDoc UpdateDynamicMessageDoc(ctx, body)
Update Dynamic Message Doc

<p>Update a Dynamic Message Doc.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**JsonDynamicMessageDoc**](JsonDynamicMessageDoc.md)| JsonDynamicMessageDoc object that needs to be added. | 

### Return type

[**JsonDynamicMessageDoc**](JsonDynamicMessageDoc.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

