# \TransactionFirehoseApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFirehoseTransactionsForBankAccount**](TransactionFirehoseApi.md#GetFirehoseTransactionsForBankAccount) | **Get** /obp/v5.1.0/banks/{BANK_ID}/firehose/accounts/{ACCOUNT_ID}/views/{VIEW_ID}/transactions | Get Firehose Transactions for Account


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

