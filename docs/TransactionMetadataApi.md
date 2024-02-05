# \TransactionMetadataApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCommentForViewOnTransaction**](TransactionMetadataApi.md#AddCommentForViewOnTransaction) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/comments | Add a Transaction Comment
[**AddImageForViewOnTransaction**](TransactionMetadataApi.md#AddImageForViewOnTransaction) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/images | Add a Transaction Image
[**AddTagForViewOnTransaction**](TransactionMetadataApi.md#AddTagForViewOnTransaction) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/tags | Add a Transaction Tag
[**AddTransactionNarrative**](TransactionMetadataApi.md#AddTransactionNarrative) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/narrative | Add a Transaction Narrative
[**AddWhereTagForViewOnTransaction**](TransactionMetadataApi.md#AddWhereTagForViewOnTransaction) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/where | Add a Transaction where Tag
[**DeleteCommentForViewOnTransaction**](TransactionMetadataApi.md#DeleteCommentForViewOnTransaction) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/comments/{COMMENT_ID} | Delete a Transaction Comment
[**DeleteImageForViewOnTransaction**](TransactionMetadataApi.md#DeleteImageForViewOnTransaction) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/images/{IMAGE_ID} | Delete a Transaction Image
[**DeleteTagForViewOnTransaction**](TransactionMetadataApi.md#DeleteTagForViewOnTransaction) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/tags/{TAG_ID} | Delete a Transaction Tag
[**DeleteTransactionNarrative**](TransactionMetadataApi.md#DeleteTransactionNarrative) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/narrative | Delete a Transaction Narrative
[**DeleteWhereTagForViewOnTransaction**](TransactionMetadataApi.md#DeleteWhereTagForViewOnTransaction) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/where | Delete a Transaction Tag
[**GetCommentsForViewOnTransaction**](TransactionMetadataApi.md#GetCommentsForViewOnTransaction) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/comments | Get Transaction Comments
[**GetImagesForViewOnTransaction**](TransactionMetadataApi.md#GetImagesForViewOnTransaction) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/images | Get Transaction Images
[**GetTagsForViewOnTransaction**](TransactionMetadataApi.md#GetTagsForViewOnTransaction) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/tags | Get Transaction Tags
[**GetTransactionNarrative**](TransactionMetadataApi.md#GetTransactionNarrative) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/narrative | Get a Transaction Narrative
[**GetWhereTagForViewOnTransaction**](TransactionMetadataApi.md#GetWhereTagForViewOnTransaction) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/where | Get a Transaction where Tag
[**UpdateTransactionNarrative**](TransactionMetadataApi.md#UpdateTransactionNarrative) | **Put** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/narrative | Update a Transaction Narrative
[**UpdateWhereTagForViewOnTransaction**](TransactionMetadataApi.md#UpdateWhereTagForViewOnTransaction) | **Put** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/where | Update a Transaction where Tag


# **AddCommentForViewOnTransaction**
> TransactionCommentJson AddCommentForViewOnTransaction(ctx, body, tRANSACTIONID, vIEWID, aCCOUNTID, bANKID)
Add a Transaction Comment

<p>Posts a comment about a transaction TRANSACTION_ID on a <a href=\"#1_2_1-getViewsForBankAccount\">view</a> VIEW_ID.</p><p>${authenticationRequiredMessage(false)}</p><p>Authentication is required since the comment is linked with the user.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostTransactionCommentJson**](PostTransactionCommentJson.md)| PostTransactionCommentJSON object that needs to be added. | 
  **tRANSACTIONID** | **string**| The transaction id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**TransactionCommentJson**](TransactionCommentJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddImageForViewOnTransaction**
> TransactionImageJson AddImageForViewOnTransaction(ctx, body, tRANSACTIONID, vIEWID, aCCOUNTID, bANKID)
Add a Transaction Image

<p>Posts an image about a transaction TRANSACTION_ID on a <a href=\"#1_2_1-getViewsForBankAccount\">view</a> VIEW_ID.</p><p>Authentication is Mandatory</p><p>The image is linked with the user.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostTransactionImageJson**](PostTransactionImageJson.md)| PostTransactionImageJSON object that needs to be added. | 
  **tRANSACTIONID** | **string**| The transaction id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**TransactionImageJson**](TransactionImageJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddTagForViewOnTransaction**
> TransactionTagJson AddTagForViewOnTransaction(ctx, body, tRANSACTIONID, vIEWID, aCCOUNTID, bANKID)
Add a Transaction Tag

<p>Posts a tag about a transaction TRANSACTION_ID on a <a href=\"#1_2_1-getViewsForBankAccount\">view</a> VIEW_ID.</p><p>Authentication is Mandatory</p><p>Authentication is required as the tag is linked with the user.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostTransactionTagJson**](PostTransactionTagJson.md)| PostTransactionTagJSON object that needs to be added. | 
  **tRANSACTIONID** | **string**| The transaction id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**TransactionTagJson**](TransactionTagJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddTransactionNarrative**
> SuccessMessage AddTransactionNarrative(ctx, body, tRANSACTIONID, vIEWID, aCCOUNTID, bANKID)
Add a Transaction Narrative

<p>Creates a description of the transaction TRANSACTION_ID.</p><p>Note: Unlike other items of metadata, there is only one &quot;narrative&quot; per transaction accross all views.<br />If you set narrative via a view e.g. view-x it will be seen via view-y (as long as view-y has permission to see the narrative).</p><p>Authentication is Optional<br />Authentication is required if the view is not public.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransactionNarrativeJson**](TransactionNarrativeJson.md)| TransactionNarrativeJSON object that needs to be added. | 
  **tRANSACTIONID** | **string**| The transaction id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddWhereTagForViewOnTransaction**
> SuccessMessage AddWhereTagForViewOnTransaction(ctx, body, tRANSACTIONID, vIEWID, aCCOUNTID, bANKID)
Add a Transaction where Tag

<p>Creates a &quot;where&quot; Geo tag on a transaction TRANSACTION_ID in a <a href=\"#1_2_1-getViewsForBankAccount\">view</a>.</p><p>Authentication is Mandatory</p><p>The geo tag is linked with the user.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostTransactionWhereJson**](PostTransactionWhereJson.md)| PostTransactionWhereJSON object that needs to be added. | 
  **tRANSACTIONID** | **string**| The transaction id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCommentForViewOnTransaction**
> EmptyClassJson DeleteCommentForViewOnTransaction(ctx, body, cOMMENTID, tRANSACTIONID, vIEWID, aCCOUNTID, bANKID)
Delete a Transaction Comment

<p>Delete the comment COMMENT_ID about the transaction TRANSACTION_ID made on <a href=\"#1_2_1-getViewsForBankAccount\">view</a>.</p><p>Authentication via OAuth is required. The user must either have owner privileges for this account, or must be the user that posted the comment.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **cOMMENTID** | **string**| The comment id | 
  **tRANSACTIONID** | **string**| The transaction id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**EmptyClassJson**](EmptyClassJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteImageForViewOnTransaction**
> EmptyClassJson DeleteImageForViewOnTransaction(ctx, body, iMAGEID, tRANSACTIONID, vIEWID, aCCOUNTID, bANKID)
Delete a Transaction Image

<p>Deletes the image IMAGE_ID about the transaction TRANSACTION_ID made on <a href=\"#1_2_1-getViewsForBankAccount\">view</a>.</p><p>Authentication via OAuth is required. The user must either have owner privileges for this account, or must be the user that posted the image.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **iMAGEID** | **string**| The image id | 
  **tRANSACTIONID** | **string**| The transaction id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**EmptyClassJson**](EmptyClassJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTagForViewOnTransaction**
> EmptyClassJson DeleteTagForViewOnTransaction(ctx, body, tAGID, tRANSACTIONID, vIEWID, aCCOUNTID, bANKID)
Delete a Transaction Tag

<p>Deletes the tag TAG_ID about the transaction TRANSACTION_ID made on <a href=\"#1_2_1-getViewsForBankAccount\">view</a>.<br />Authentication via OAuth is required. The user must either have owner privileges for this account,<br />or must be the user that posted the tag.</p><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **tAGID** | **string**| The tag id | 
  **tRANSACTIONID** | **string**| The transaction id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**EmptyClassJson**](EmptyClassJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTransactionNarrative**
> EmptyClassJson DeleteTransactionNarrative(ctx, body, tRANSACTIONID, vIEWID, aCCOUNTID, bANKID)
Delete a Transaction Narrative

<p>Deletes the description of the transaction TRANSACTION_ID.</p><p>Authentication via OAuth is required if the view is not public.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **tRANSACTIONID** | **string**| The transaction id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**EmptyClassJson**](EmptyClassJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteWhereTagForViewOnTransaction**
> EmptyClassJson DeleteWhereTagForViewOnTransaction(ctx, body, tRANSACTIONID, vIEWID, aCCOUNTID, bANKID)
Delete a Transaction Tag

<p>Deletes the where tag of the transaction TRANSACTION_ID made on <a href=\"#1_2_1-getViewsForBankAccount\">view</a>.</p><p>Authentication is Mandatory</p><p>The user must either have owner privileges for this account, or must be the user that posted the geo tag.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **tRANSACTIONID** | **string**| The transaction id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**EmptyClassJson**](EmptyClassJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCommentsForViewOnTransaction**
> TransactionCommentsJson GetCommentsForViewOnTransaction(ctx, body, tRANSACTIONID, vIEWID, aCCOUNTID, bANKID)
Get Transaction Comments

<p>Returns the transaction TRANSACTION_ID comments made on a <a href=\"#1_2_1-getViewsForBankAccount\">view</a> (VIEW_ID).</p><p>Authentication via OAuth is required if the view is not public.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **tRANSACTIONID** | **string**| The transaction id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**TransactionCommentsJson**](TransactionCommentsJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImagesForViewOnTransaction**
> TransactionImagesJson GetImagesForViewOnTransaction(ctx, body, tRANSACTIONID, vIEWID, aCCOUNTID, bANKID)
Get Transaction Images

<p>Returns the transaction TRANSACTION_ID images made on a <a href=\"#1_2_1-getViewsForBankAccount\">view</a> (VIEW_ID).<br />Authentication via OAuth is required if the view is not public.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **tRANSACTIONID** | **string**| The transaction id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**TransactionImagesJson**](TransactionImagesJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTagsForViewOnTransaction**
> TransactionTagJson GetTagsForViewOnTransaction(ctx, body, tRANSACTIONID, vIEWID, aCCOUNTID, bANKID)
Get Transaction Tags

<p>Returns the transaction TRANSACTION_ID tags made on a <a href=\"#1_2_1-getViewsForBankAccount\">view</a> (VIEW_ID).<br />Authentication via OAuth is required if the view is not public.</p><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **tRANSACTIONID** | **string**| The transaction id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**TransactionTagJson**](TransactionTagJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransactionNarrative**
> TransactionNarrativeJson GetTransactionNarrative(ctx, body, tRANSACTIONID, vIEWID, aCCOUNTID, bANKID)
Get a Transaction Narrative

<p>Returns the account owner description of the transaction <a href=\"#1_2_1-getViewsForBankAccount\">moderated</a> by the view.</p><p>Authentication via OAuth is required if the view is not public.</p><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **tRANSACTIONID** | **string**| The transaction id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**TransactionNarrativeJson**](TransactionNarrativeJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWhereTagForViewOnTransaction**
> TransactionWhereJson GetWhereTagForViewOnTransaction(ctx, body, tRANSACTIONID, vIEWID, aCCOUNTID, bANKID)
Get a Transaction where Tag

<p>Returns the &quot;where&quot; Geo tag added to the transaction TRANSACTION_ID made on a <a href=\"#1_2_1-getViewsForBankAccount\">view</a> (VIEW_ID).<br />It represents the location where the transaction has been initiated.</p><p>Authentication via OAuth is required if the view is not public.</p><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **tRANSACTIONID** | **string**| The transaction id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**TransactionWhereJson**](TransactionWhereJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTransactionNarrative**
> SuccessMessage UpdateTransactionNarrative(ctx, body, tRANSACTIONID, vIEWID, aCCOUNTID, bANKID)
Update a Transaction Narrative

<p>Updates the description of the transaction TRANSACTION_ID.</p><p>Authentication via OAuth is required if the view is not public.</p><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransactionNarrativeJson**](TransactionNarrativeJson.md)| TransactionNarrativeJSON object that needs to be added. | 
  **tRANSACTIONID** | **string**| The transaction id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWhereTagForViewOnTransaction**
> SuccessMessage UpdateWhereTagForViewOnTransaction(ctx, body, tRANSACTIONID, vIEWID, aCCOUNTID, bANKID)
Update a Transaction where Tag

<p>Updates the &quot;where&quot; Geo tag on a transaction TRANSACTION_ID in a <a href=\"#1_2_1-getViewsForBankAccount\">view</a>.</p><p>Authentication is Mandatory</p><p>The geo tag is linked with the user.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostTransactionWhereJson**](PostTransactionWhereJson.md)| PostTransactionWhereJSON object that needs to be added. | 
  **tRANSACTIONID** | **string**| The transaction id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

