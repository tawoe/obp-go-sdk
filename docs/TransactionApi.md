# \TransactionApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCommentForViewOnTransaction**](TransactionApi.md#AddCommentForViewOnTransaction) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/comments | Add a Transaction Comment
[**AddImageForViewOnTransaction**](TransactionApi.md#AddImageForViewOnTransaction) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/images | Add a Transaction Image
[**AddTagForViewOnTransaction**](TransactionApi.md#AddTagForViewOnTransaction) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/tags | Add a Transaction Tag
[**AddTransactionNarrative**](TransactionApi.md#AddTransactionNarrative) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/narrative | Add a Transaction Narrative
[**AddWhereTagForViewOnTransaction**](TransactionApi.md#AddWhereTagForViewOnTransaction) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/where | Add a Transaction where Tag
[**CreateOrUpdateTransactionAttributeDefinition**](TransactionApi.md#CreateOrUpdateTransactionAttributeDefinition) | **Put** /obp/v5.1.0/banks/{BANK_ID}/attribute-definitions/transaction | Create or Update Transaction Attribute Definition
[**CreateTransactionAttribute**](TransactionApi.md#CreateTransactionAttribute) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/transactions/{TRANSACTION_ID}/attribute | Create Transaction Attribute
[**DeleteCommentForViewOnTransaction**](TransactionApi.md#DeleteCommentForViewOnTransaction) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/comments/{COMMENT_ID} | Delete a Transaction Comment
[**DeleteImageForViewOnTransaction**](TransactionApi.md#DeleteImageForViewOnTransaction) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/images/{IMAGE_ID} | Delete a Transaction Image
[**DeleteTagForViewOnTransaction**](TransactionApi.md#DeleteTagForViewOnTransaction) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/tags/{TAG_ID} | Delete a Transaction Tag
[**DeleteTransactionAttributeDefinition**](TransactionApi.md#DeleteTransactionAttributeDefinition) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/attribute-definitions/ATTRIBUTE_DEFINITION_ID/transaction | Delete Transaction Attribute Definition
[**DeleteTransactionCascade**](TransactionApi.md#DeleteTransactionCascade) | **Delete** /obp/v5.1.0/management/cascading/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/transactions/{TRANSACTION_ID} | Delete Transaction Cascade
[**DeleteTransactionNarrative**](TransactionApi.md#DeleteTransactionNarrative) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/narrative | Delete a Transaction Narrative
[**DeleteWhereTagForViewOnTransaction**](TransactionApi.md#DeleteWhereTagForViewOnTransaction) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/where | Delete a Transaction Tag
[**GetBalancingTransaction**](TransactionApi.md#GetBalancingTransaction) | **Get** /obp/v5.1.0/transactions/{TRANSACTION_ID}/balancing-transaction | Get Balancing Transaction
[**GetCommentsForViewOnTransaction**](TransactionApi.md#GetCommentsForViewOnTransaction) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/comments | Get Transaction Comments
[**GetCoreTransactionsForBankAccount**](TransactionApi.md#GetCoreTransactionsForBankAccount) | **Get** /obp/v5.1.0/my/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/transactions | Get Transactions for Account (Core)
[**GetDoubleEntryTransaction**](TransactionApi.md#GetDoubleEntryTransaction) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/double-entry-transaction | Get Double Entry Transaction
[**GetFirehoseTransactionsForBankAccount**](TransactionApi.md#GetFirehoseTransactionsForBankAccount) | **Get** /obp/v5.1.0/banks/{BANK_ID}/firehose/accounts/{ACCOUNT_ID}/views/{VIEW_ID}/transactions | Get Firehose Transactions for Account
[**GetImagesForViewOnTransaction**](TransactionApi.md#GetImagesForViewOnTransaction) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/images | Get Transaction Images
[**GetOtherAccountForTransaction**](TransactionApi.md#GetOtherAccountForTransaction) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/other_account | Get Other Account of Transaction
[**GetTagsForViewOnTransaction**](TransactionApi.md#GetTagsForViewOnTransaction) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/tags | Get Transaction Tags
[**GetTransactionAttributeById**](TransactionApi.md#GetTransactionAttributeById) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/transactions/{TRANSACTION_ID}/attributes/ATTRIBUTE_ID | Get Transaction Attribute By Id
[**GetTransactionAttributeDefinition**](TransactionApi.md#GetTransactionAttributeDefinition) | **Get** /obp/v5.1.0/banks/{BANK_ID}/attribute-definitions/transaction | Get Transaction Attribute Definition
[**GetTransactionAttributes**](TransactionApi.md#GetTransactionAttributes) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/transactions/{TRANSACTION_ID}/attributes | Get Transaction Attributes
[**GetTransactionByIdForBankAccount**](TransactionApi.md#GetTransactionByIdForBankAccount) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/transaction | Get Transaction by Id
[**GetTransactionNarrative**](TransactionApi.md#GetTransactionNarrative) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/narrative | Get a Transaction Narrative
[**GetTransactionsForBankAccount**](TransactionApi.md#GetTransactionsForBankAccount) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions | Get Transactions for Account (Full)
[**GetWhereTagForViewOnTransaction**](TransactionApi.md#GetWhereTagForViewOnTransaction) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/where | Get a Transaction where Tag
[**UpdateTransactionAttribute**](TransactionApi.md#UpdateTransactionAttribute) | **Put** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/transactions/{TRANSACTION_ID}/attributes/{ACCOUNT_ATTRIBUTE_ID} | Update Transaction Attribute
[**UpdateTransactionNarrative**](TransactionApi.md#UpdateTransactionNarrative) | **Put** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/narrative | Update a Transaction Narrative
[**UpdateWhereTagForViewOnTransaction**](TransactionApi.md#UpdateWhereTagForViewOnTransaction) | **Put** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/metadata/where | Update a Transaction where Tag


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

# **CreateOrUpdateTransactionAttributeDefinition**
> AttributeDefinitionResponseJsonV400 CreateOrUpdateTransactionAttributeDefinition(ctx, body, bANKID)
Create or Update Transaction Attribute Definition

<p>Create or Update Transaction Attribute Definition</p><p>The category field must be Transaction</p><p>The type field must be one of; DOUBLE, STRING, INTEGER and DATE_WITH_DAY</p><p>Authentication is Mandatory</p>

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

# **CreateTransactionAttribute**
> TransactionAttributeResponseJson CreateTransactionAttribute(ctx, body, tRANSACTIONID, aCCOUNTID, bANKID)
Create Transaction Attribute

<p>Create Transaction Attribute</p><p>The type field must be one of &quot;STRING&quot;, &quot;INTEGER&quot;, &quot;DOUBLE&quot; or DATE_WITH_DAY&quot;</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransactionAttributeJsonV400**](TransactionAttributeJsonV400.md)| TransactionAttributeJsonV400 object that needs to be added. | 
  **tRANSACTIONID** | **string**| The transaction id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**TransactionAttributeResponseJson**](TransactionAttributeResponseJson.md)

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

# **DeleteTransactionAttributeDefinition**
> DeleteTransactionAttributeDefinition(ctx, bANKID)
Delete Transaction Attribute Definition

<p>Delete Transaction Attribute Definition by ATTRIBUTE_DEFINITION_ID</p><p>Authentication is Mandatory</p>

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

# **DeleteTransactionCascade**
> DeleteTransactionCascade(ctx, tRANSACTIONID, aCCOUNTID, bANKID)
Delete Transaction Cascade

<p>Delete a Transaction Cascade specified by TRANSACTION_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tRANSACTIONID** | **string**| The transaction id | 
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

# **GetBalancingTransaction**
> DoubleEntryTransactionJson GetBalancingTransaction(ctx, tRANSACTIONID)
Get Balancing Transaction

<p>Get Balancing Transaction</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tRANSACTIONID** | **string**| The transaction id | 

### Return type

[**DoubleEntryTransactionJson**](DoubleEntryTransactionJson.md)

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

# **GetCoreTransactionsForBankAccount**
> CoreTransactionsJsonV300 GetCoreTransactionsForBankAccount(ctx, aCCOUNTID, bANKID)
Get Transactions for Account (Core)

<p>Returns transactions list (Core info) of the account specified by ACCOUNT_ID.</p><p>Authentication is Mandatory</p><p>Possible custom url parameters for pagination:</p><ul><li>limit=NUMBER ==&gt; default value: 500</li><li>offset=NUMBER ==&gt; default value: 0</li></ul><p>eg1:?limit=100&amp;offset=0</p><ul><li>sort_direction=ASC/DESC ==&gt; default value: DESC.</li></ul><p>eg2:?limit=100&amp;offset=0&amp;sort_direction=ASC</p><ul><li>from_date=DATE =&gt; example value: 1970-01-01T00:00:00.000Z. NOTE! The default value is one year ago (1970-01-01T00:00:00.000Z).</li><li>to_date=DATE =&gt; example value: 2024-02-05T14:15:55.255Z. NOTE! The default value is now (2024-02-05T14:15:55.255Z).</li></ul><p>Date format parameter: yyyy-MM-dd'T'HH:mm:ss.SSS'Z'(1100-01-01T01:01:01.000Z) ==&gt; time zone is UTC.</p><p>eg3:?sort_direction=ASC&amp;limit=100&amp;offset=0&amp;from_date=1100-01-01T01:01:01.000Z&amp;to_date=1100-01-01T01:01:01.000Z</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CoreTransactionsJsonV300**](CoreTransactionsJsonV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDoubleEntryTransaction**
> DoubleEntryTransactionJson GetDoubleEntryTransaction(ctx, tRANSACTIONID, vIEWID, aCCOUNTID, bANKID)
Get Double Entry Transaction

<p>Get Double Entry Transaction</p><p>This endpoint can be used to see the double entry transactions. It returns the <code>bank_id</code>, <code>account_id</code> and <code>transaction_id</code><br />for the debit end the credit transaction. The other side account can be a settlement account or an OBP account.</p><p>The endpoint also provide the <code>transaction_request</code> object which contains the <code>bank_id</code>, <code>account_id</code> and<br /><code>transaction_request_id</code> of the transaction request at the origin of the transaction. Please note that if none<br />transaction request is at the origin of the transaction, the <code>transaction_request</code> object will be <code>null</code>.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tRANSACTIONID** | **string**| The transaction id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**DoubleEntryTransactionJson**](DoubleEntryTransactionJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFirehoseTransactionsForBankAccount**
> TransactionsJsonV300 GetFirehoseTransactionsForBankAccount(ctx, vIEWID, aCCOUNTID, bANKID)
Get Firehose Transactions for Account

<p>Get Transactions for an Account that has a firehose View.</p><p>Allows bulk access to an account's transactions.<br />User must have the CanUseFirehoseAtAnyBank Role</p><p>To find ACCOUNT_IDs, use the getFirehoseAccountsAtOneBank call.</p><p>For VIEW_ID try 'owner'</p><p>Possible custom url parameters for pagination:</p><ul><li>limit=NUMBER ==&gt; default value: 500</li><li>offset=NUMBER ==&gt; default value: 0</li></ul><p>eg1:?limit=100&amp;offset=0</p><ul><li>sort_direction=ASC/DESC ==&gt; default value: DESC.</li></ul><p>eg2:?limit=100&amp;offset=0&amp;sort_direction=ASC</p><ul><li>from_date=DATE =&gt; example value: 1970-01-01T00:00:00.000Z. NOTE! The default value is one year ago (1970-01-01T00:00:00.000Z).</li><li>to_date=DATE =&gt; example value: 2024-02-05T14:15:55.255Z. NOTE! The default value is now (2024-02-05T14:15:55.255Z).</li></ul><p>Date format parameter: yyyy-MM-dd'T'HH:mm:ss.SSS'Z'(1100-01-01T01:01:01.000Z) ==&gt; time zone is UTC.</p><p>eg3:?sort_direction=ASC&amp;limit=100&amp;offset=0&amp;from_date=1100-01-01T01:01:01.000Z&amp;to_date=1100-01-01T01:01:01.000Z</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**TransactionsJsonV300**](TransactionsJsonV300.md)

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

# **GetOtherAccountForTransaction**
> OtherAccountJson GetOtherAccountForTransaction(ctx, body, tRANSACTIONID, vIEWID, aCCOUNTID, bANKID)
Get Other Account of Transaction

<p>Get other account of a transaction.<br />Returns details of the other party involved in the transaction, moderated by the <a href=\"#1_2_1-getViewsForBankAccount\">view</a> (VIEW_ID).<br />Authentication via OAuth is required if the view is not public.</p><p>Authentication is Optional</p>

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

[**OtherAccountJson**](OtherAccountJSON.md)

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

# **GetTransactionAttributeById**
> TransactionAttributeResponseJson GetTransactionAttributeById(ctx, tRANSACTIONID, aCCOUNTID, bANKID)
Get Transaction Attribute By Id

<p>Get Transaction Attribute By Id</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tRANSACTIONID** | **string**| The transaction id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**TransactionAttributeResponseJson**](TransactionAttributeResponseJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransactionAttributeDefinition**
> AttributeDefinitionsResponseJsonV400 GetTransactionAttributeDefinition(ctx, bANKID)
Get Transaction Attribute Definition

<p>Get Transaction Attribute Definition</p><p>Authentication is Mandatory</p>

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

# **GetTransactionAttributes**
> TransactionAttributesResponseJson GetTransactionAttributes(ctx, tRANSACTIONID, aCCOUNTID, bANKID)
Get Transaction Attributes

<p>Get Transaction Attributes</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tRANSACTIONID** | **string**| The transaction id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**TransactionAttributesResponseJson**](TransactionAttributesResponseJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransactionByIdForBankAccount**
> TransactionJsonV300 GetTransactionByIdForBankAccount(ctx, tRANSACTIONID, vIEWID, aCCOUNTID, bANKID)
Get Transaction by Id

<p>Returns one transaction specified by TRANSACTION_ID of the account ACCOUNT_ID and <a href=\"#1_2_1-getViewsForBankAccount\">moderated</a> by the view (VIEW_ID).</p><p>Authentication is Optional<br />Authentication is required if the view is not public.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tRANSACTIONID** | **string**| The transaction id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**TransactionJsonV300**](TransactionJsonV300.md)

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

# **GetTransactionsForBankAccount**
> TransactionsJsonV300 GetTransactionsForBankAccount(ctx, vIEWID, aCCOUNTID, bANKID)
Get Transactions for Account (Full)

<p>Returns transactions list of the account specified by ACCOUNT_ID and <a href=\"#1_2_1-getViewsForBankAccount\">moderated</a> by the view (VIEW_ID).</p><p>Authentication is Optional</p><p>Authentication is required if the view is not public.</p><p>Possible custom url parameters for pagination:</p><ul><li>limit=NUMBER ==&gt; default value: 500</li><li>offset=NUMBER ==&gt; default value: 0</li></ul><p>eg1:?limit=100&amp;offset=0</p><ul><li>sort_direction=ASC/DESC ==&gt; default value: DESC.</li></ul><p>eg2:?limit=100&amp;offset=0&amp;sort_direction=ASC</p><ul><li>from_date=DATE =&gt; example value: 1970-01-01T00:00:00.000Z. NOTE! The default value is one year ago (1970-01-01T00:00:00.000Z).</li><li>to_date=DATE =&gt; example value: 2024-02-05T14:15:55.255Z. NOTE! The default value is now (2024-02-05T14:15:55.255Z).</li></ul><p>Date format parameter: yyyy-MM-dd'T'HH:mm:ss.SSS'Z'(1100-01-01T01:01:01.000Z) ==&gt; time zone is UTC.</p><p>eg3:?sort_direction=ASC&amp;limit=100&amp;offset=0&amp;from_date=1100-01-01T01:01:01.000Z&amp;to_date=1100-01-01T01:01:01.000Z</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**TransactionsJsonV300**](TransactionsJsonV300.md)

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

# **UpdateTransactionAttribute**
> TransactionAttributeResponseJson UpdateTransactionAttribute(ctx, body, aCCOUNTATTRIBUTEID, tRANSACTIONID, aCCOUNTID, bANKID)
Update Transaction Attribute

<p>Update Transaction Attribute</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransactionAttributeJsonV400**](TransactionAttributeJsonV400.md)| TransactionAttributeJsonV400 object that needs to be added. | 
  **aCCOUNTATTRIBUTEID** | **string**| the account attribute id  | 
  **tRANSACTIONID** | **string**| The transaction id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**TransactionAttributeResponseJson**](TransactionAttributeResponseJson.md)

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

