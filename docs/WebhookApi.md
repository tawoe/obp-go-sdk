# \WebhookApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountWebhook**](WebhookApi.md#CreateAccountWebhook) | **Post** /obp/v5.1.0/banks/{BANK_ID}/account-web-hooks | Create an Account Webhook
[**CreateBankAccountNotificationWebhook**](WebhookApi.md#CreateBankAccountNotificationWebhook) | **Post** /obp/v5.1.0/banks/{BANK_ID}/web-hooks/account/notifications/on-create-transaction | Create bank level Account Notification Webhook
[**CreateSystemAccountNotificationWebhook**](WebhookApi.md#CreateSystemAccountNotificationWebhook) | **Post** /obp/v5.1.0/web-hooks/account/notifications/on-create-transaction | Create system level Account Notification Webhook
[**EnableDisableAccountWebhook**](WebhookApi.md#EnableDisableAccountWebhook) | **Put** /obp/v5.1.0/banks/{BANK_ID}/account-web-hooks | Enable/Disable an Account Webhook
[**GetAccountWebhooks**](WebhookApi.md#GetAccountWebhooks) | **Get** /obp/v5.1.0/management/banks/{BANK_ID}/account-web-hooks | Get Account Webhooks


# **CreateAccountWebhook**
> AccountWebhookJson CreateAccountWebhook(ctx, body, bANKID)
Create an Account Webhook

<p>Create an Account Webhook</p><p>Webhooks are used to call external URLs when certain events happen.</p><p>Account Webhooks focus on events around accounts.</p><p>For instance, a webhook could be used to notify an external service if a balance changes on an account.</p><p>This functionality is work in progress! Please note that only implemented trigger is: OnBalanceChange</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountWebhookPostJson**](AccountWebhookPostJson.md)| AccountWebhookPostJson object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**AccountWebhookJson**](AccountWebhookJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBankAccountNotificationWebhook**
> BankAccountNotificationWebhookJson CreateBankAccountNotificationWebhook(ctx, body, bANKID)
Create bank level Account Notification Webhook

<p>Create a notification Webhook that will fire for all accounts on the specified Bank.</p><p>Webhooks are used to call external web services when certain events happen.</p><p>For instance, a webhook can be used to notify an external service if a transaction is created on an account.</p><p>When an account notification webhook fires it will POST to the URL you specify during the creation of the webhook.</p><p>Inside the payload you will find account_id and transaction_id and also user_ids and customer_ids of the Users / Customers linked to the Account.</p><p>The webhook will POST the following structure to your service:</p><p>{<br />&quot;event_name&quot;: &quot;OnCreateTransaction&quot;,<br />&quot;event_id&quot;: &quot;9ca9a7e4-6d02-40e3-a129-0b2bf89de9b1&quot;,<br />&quot;bank_id&quot;: &quot;gh.29.uk&quot;,<br />&quot;account_id&quot;: &quot;8ca9a7e4-6d02-40e3-a129-0b2bf89de9b1&quot;,<br />&quot;transaction_id&quot;: &quot;7ca9a7e4-6d02-40e3-a129-0b2bf89de9b1&quot;,<br />&quot;related_entities&quot;: [<br />{<br />&quot;user_id&quot;: &quot;8ca9a7e4-6d02-40e3-a129-0b2bf89de9b1&quot;,<br />&quot;customer_ids&quot;: [&quot;3ca9a7e4-6d02-40e3-a129-0b2bf89de9b1&quot;]<br />}<br />]<br />}</p><p>Thus, your service should accept the above POST body structure.</p><p>In this way, your web service can be informed about an event on an account and act accordingly.</p><p>Further information about the account, transaction or related entities can then be retrieved using the standard REST APIs.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountNotificationWebhookPostJson**](AccountNotificationWebhookPostJson.md)| AccountNotificationWebhookPostJson object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**BankAccountNotificationWebhookJson**](BankAccountNotificationWebhookJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSystemAccountNotificationWebhook**
> SystemAccountNotificationWebhookJson CreateSystemAccountNotificationWebhook(ctx, body)
Create system level Account Notification Webhook

<p>Create a notification Webhook that will fire for all accounts on the system.</p><p>Webhooks are used to call external web services when certain events happen.</p><p>For instance, a webhook can be used to notify an external service if a transaction is created on an account.</p><p>When an account notification webhook fires it will POST to the URL you specify during the creation of the webhook.</p><p>Inside the payload you will find account_id and transaction_id and also user_ids and customer_ids of the Users / Customers linked to the Account.</p><p>The webhook will POST the following structure to your service:</p><p>{<br />&quot;event_name&quot;: &quot;OnCreateTransaction&quot;,<br />&quot;event_id&quot;: &quot;9ca9a7e4-6d02-40e3-a129-0b2bf89de9b1&quot;,<br />&quot;bank_id&quot;: &quot;gh.29.uk&quot;,<br />&quot;account_id&quot;: &quot;8ca9a7e4-6d02-40e3-a129-0b2bf89de9b1&quot;,<br />&quot;transaction_id&quot;: &quot;7ca9a7e4-6d02-40e3-a129-0b2bf89de9b1&quot;,<br />&quot;related_entities&quot;: [<br />{<br />&quot;user_id&quot;: &quot;8ca9a7e4-6d02-40e3-a129-0b2bf89de9b1&quot;,<br />&quot;customer_ids&quot;: [&quot;3ca9a7e4-6d02-40e3-a129-0b2bf89de9b1&quot;]<br />}<br />]<br />}</p><p>Thus, your service should accept the above POST body structure.</p><p>In this way, your web service can be informed about an event on an account and act accordingly.</p><p>Further information about the account, transaction or related entities can then be retrieved using the standard REST APIs.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountNotificationWebhookPostJson**](AccountNotificationWebhookPostJson.md)| AccountNotificationWebhookPostJson object that needs to be added. | 

### Return type

[**SystemAccountNotificationWebhookJson**](SystemAccountNotificationWebhookJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableDisableAccountWebhook**
> AccountWebhookJson EnableDisableAccountWebhook(ctx, body, bANKID)
Enable/Disable an Account Webhook

<p>Enable/Disable an Account Webhook</p><p>Webhooks are used to call external URLs when certain events happen.</p><p>Account Webhooks focus on events around accounts.</p><p>For instance, a webhook could be used to notify an external service if a balance changes on an account.</p><p>This functionality is work in progress! Please note that only implemented trigger is: OnBalanceChange</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountWebhookPutJson**](AccountWebhookPutJson.md)| AccountWebhookPutJson object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**AccountWebhookJson**](AccountWebhookJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountWebhooks**
> AccountWebhooksJson GetAccountWebhooks(ctx, bANKID)
Get Account Webhooks

<p>Get Account Webhooks.</p><p>Possible custom URL parameters for pagination:</p><p>Possible custom url parameters for pagination:</p><ul><li>limit=NUMBER ==&gt; default value: 500</li><li>offset=NUMBER ==&gt; default value: 0</li></ul><p>eg1:?limit=100&amp;offset=0</p><ul><li>sort_direction=ASC/DESC ==&gt; default value: DESC.</li></ul><p>eg2:?limit=100&amp;offset=0&amp;sort_direction=ASC</p><ul><li>account_id=STRING (if null ignore)</li><li>user_id=STRING (if null ignore)</li></ul><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**AccountWebhooksJson**](AccountWebhooksJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

