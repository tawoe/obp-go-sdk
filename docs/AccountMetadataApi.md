# \AccountMetadataApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTagForViewOnAccount**](AccountMetadataApi.md#AddTagForViewOnAccount) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/metadata/tags | Create a tag on account
[**DeleteTagForViewOnAccount**](AccountMetadataApi.md#DeleteTagForViewOnAccount) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/metadata/tags/{TAG_ID} | Delete a tag on account
[**GetTagsForViewOnAccount**](AccountMetadataApi.md#GetTagsForViewOnAccount) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/metadata/tags | Get tags on account


# **AddTagForViewOnAccount**
> AccountTagJson AddTagForViewOnAccount(ctx, body, vIEWID, aCCOUNTID, bANKID)
Create a tag on account

<p>Posts a tag about an account ACCOUNT_ID on a <a href=\"#1_2_1-getViewsForBankAccount\">view</a> VIEW_ID.</p><p>Authentication is Mandatory</p><p>Authentication is required as the tag is linked with the user.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostAccountTagJson**](PostAccountTagJson.md)| PostAccountTagJSON object that needs to be added. | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**AccountTagJson**](AccountTagJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTagForViewOnAccount**
> DeleteTagForViewOnAccount(ctx, tAGID, vIEWID, aCCOUNTID, bANKID)
Delete a tag on account

<p>Deletes the tag TAG_ID about the account ACCOUNT_ID made on <a href=\"#1_2_1-getViewsForBankAccount\">view</a>.</p><p>Authentication is Mandatory</p><p>Authentication is required as the tag is linked with the user.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tAGID** | **string**| The tag id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTagsForViewOnAccount**
> AccountTagsJson GetTagsForViewOnAccount(ctx, vIEWID, aCCOUNTID, bANKID)
Get tags on account

<p>Returns the account ACCOUNT_ID tags made on a <a href=\"#1_2_1-getViewsForBankAccount\">view</a> (VIEW_ID).<br />Authentication is Mandatory</p><p>Authentication is required as the tag is linked with the user.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**AccountTagsJson**](AccountTagsJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

