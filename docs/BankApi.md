# \BankApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountWebhook**](BankApi.md#CreateAccountWebhook) | **Post** /obp/v5.1.0/banks/{BANK_ID}/account-web-hooks | Create an Account Webhook
[**CreateBank**](BankApi.md#CreateBank) | **Post** /obp/v5.1.0/banks | Create Bank
[**CreateBankAccountNotificationWebhook**](BankApi.md#CreateBankAccountNotificationWebhook) | **Post** /obp/v5.1.0/banks/{BANK_ID}/web-hooks/account/notifications/on-create-transaction | Create bank level Account Notification Webhook
[**CreateBankAttribute**](BankApi.md#CreateBankAttribute) | **Post** /obp/v5.1.0/banks/{BANK_ID}/attribute | Create Bank Attribute
[**CreateOrUpdateBankAttributeDefinition**](BankApi.md#CreateOrUpdateBankAttributeDefinition) | **Put** /obp/v5.1.0/banks/{BANK_ID}/attribute-definitions/bank | Create or Update Bank Attribute Definition
[**CreateSettlementAccount**](BankApi.md#CreateSettlementAccount) | **Post** /obp/v5.1.0/banks/{BANK_ID}/settlement-accounts | Create Settlement Account
[**CreateSystemAccountNotificationWebhook**](BankApi.md#CreateSystemAccountNotificationWebhook) | **Post** /obp/v5.1.0/web-hooks/account/notifications/on-create-transaction | Create system level Account Notification Webhook
[**CreateTransactionType**](BankApi.md#CreateTransactionType) | **Put** /obp/v5.1.0/banks/{BANK_ID}/transaction-types | Create Transaction Type at bank
[**DeleteBankAttribute**](BankApi.md#DeleteBankAttribute) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/attributes/BANK_ATTRIBUTE_ID | Delete Bank Attribute
[**DeleteBankCascade**](BankApi.md#DeleteBankCascade) | **Delete** /obp/v5.1.0/management/cascading/banks/{BANK_ID} | Delete Bank Cascade
[**EnableDisableAccountWebhook**](BankApi.md#EnableDisableAccountWebhook) | **Put** /obp/v5.1.0/banks/{BANK_ID}/account-web-hooks | Enable/Disable an Account Webhook
[**GetAccountWebhooks**](BankApi.md#GetAccountWebhooks) | **Get** /obp/v5.1.0/management/banks/{BANK_ID}/account-web-hooks | Get Account Webhooks
[**GetBank**](BankApi.md#GetBank) | **Get** /obp/v5.1.0/banks/{BANK_ID} | Get Bank
[**GetBankAttribute**](BankApi.md#GetBankAttribute) | **Get** /obp/v5.1.0/banks/{BANK_ID}/attributes/BANK_ATTRIBUTE_ID | Get Bank Attribute By BANK_ATTRIBUTE_ID
[**GetBankAttributes**](BankApi.md#GetBankAttributes) | **Get** /obp/v5.1.0/banks/{BANK_ID}/attributes | Get Bank Attributes
[**GetBanks**](BankApi.md#GetBanks) | **Get** /obp/v5.1.0/banks | Get Banks
[**GetBranch**](BankApi.md#GetBranch) | **Get** /obp/v5.1.0/banks/{BANK_ID}/branches/{BRANCH_ID} | Get Branch
[**GetBranches**](BankApi.md#GetBranches) | **Get** /obp/v5.1.0/banks/{BANK_ID}/branches | Get Branches for a Bank
[**GetSettlementAccounts**](BankApi.md#GetSettlementAccounts) | **Get** /obp/v5.1.0/banks/{BANK_ID}/settlement-accounts | Get Settlement accounts at Bank
[**GetTransactionRequestTypesSupportedByBank**](BankApi.md#GetTransactionRequestTypesSupportedByBank) | **Get** /obp/v5.1.0/banks/{BANK_ID}/transaction-request-types | Get Transaction Request Types at Bank
[**GetTransactionTypes**](BankApi.md#GetTransactionTypes) | **Get** /obp/v5.1.0/banks/{BANK_ID}/transaction-types | Get Transaction Types at Bank
[**UpdateBank**](BankApi.md#UpdateBank) | **Put** /obp/v5.1.0/banks | Update Bank
[**UpdateBankAttribute**](BankApi.md#UpdateBankAttribute) | **Put** /obp/v5.1.0/banks/{BANK_ID}/attributes/BANK_ATTRIBUTE_ID | Update Bank Attribute


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

# **CreateBank**
> BankJson500 CreateBank(ctx, body)
Create Bank

<p>Create a new bank (Authenticated access).</p><p>The user creating this will be automatically assigned the Role CanCreateEntitlementAtOneBank.<br />Thus the User can manage the bank they create and assign Roles to other Users.</p><p>Only SANDBOX mode<br />The settlement accounts are created specified by the bank in the POST body.<br />Name and account id are created in accordance to the next rules:<br />- Incoming account (name: Default incoming settlement account, Account ID: OBP_DEFAULT_INCOMING_ACCOUNT_ID, currency: EUR)<br />- Outgoing account (name: Default outgoing settlement account, Account ID: OBP_DEFAULT_OUTGOING_ACCOUNT_ID, currency: EUR)</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostBankJson500**](PostBankJson500.md)| PostBankJson500 object that needs to be added. | 

### Return type

[**BankJson500**](BankJson500.md)

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

# **CreateBankAttribute**
> BankAttributeResponseJsonV400 CreateBankAttribute(ctx, body, bANKID)
Create Bank Attribute

<p>Create Bank Attribute</p><p>Typical product attributes might be:</p><p>ISIN (for International bonds)<br />VKN (for German bonds)<br />REDCODE (markit short code for credit derivative)<br />LOAN_ID (e.g. used for Anacredit reporting)</p><p>ISSUE_DATE (When the bond was issued in the market)<br />MATURITY_DATE (End of life time of a product)<br />TRADABLE</p><p>See <a href=\"http://www.fpml.org/\">FPML</a> for more examples.</p><p>The type field must be one of &quot;STRING&quot;, &quot;INTEGER&quot;, &quot;DOUBLE&quot; or DATE_WITH_DAY&quot;</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BankAttributeJsonV400**](BankAttributeJsonV400.md)| BankAttributeJsonV400 object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**BankAttributeResponseJsonV400**](BankAttributeResponseJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateBankAttributeDefinition**
> AttributeDefinitionResponseJsonV400 CreateOrUpdateBankAttributeDefinition(ctx, body, bANKID)
Create or Update Bank Attribute Definition

<p>Create or Update Bank Attribute Definition</p><p>The category field must be Bank</p><p>The type field must be one of; DOUBLE, STRING, INTEGER and DATE_WITH_DAY</p><p>Authentication is Mandatory</p>

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

# **CreateSettlementAccount**
> SettlementAccountResponseJson CreateSettlementAccount(ctx, body, bANKID)
Create Settlement Account

<p>Create a new settlement account at a bank.</p><p>The created settlement account id will be the concatenation of the payment system and the account currency.<br />For examples: SEPA_SETTLEMENT_ACCOUNT_EUR, CARD_SETTLEMENT_ACCOUNT_USD</p><p>By default, when you create a new bank, two settlements accounts are created automatically: OBP_DEFAULT_INCOMING_ACCOUNT_ID and OBP_DEFAULT_OUTGOING_ACCOUNT_ID<br />Those two accounts have EUR as default currency.</p><p>If you want to create default settlement account for a specific currency, you can fill the <code>payment_system</code> field with the <code>DEFAULT</code> value.</p><p>When a transaction is saved in OBP through the mapped connector, OBP-API look for the account to save the double-entry transaction.<br />If no OBP account can be found from the counterparty, the double-entry transaction will be saved on a bank settlement account.<br />- First, the mapped connector looks for a settlement account specific to the payment system and currency. E.g SEPA_SETTLEMENT_ACCOUNT_EUR.<br />- If we don't find any specific settlement account with the payment system, we look for a default settlement account for the counterparty currency. E.g DEFAULT_SETTLEMENT_ACCOUNT_EUR.<br />- Else, we select one of the two OBP default settlement accounts (OBP_DEFAULT_INCOMING_ACCOUNT_ID/OBP_DEFAULT_OUTGOING_ACCOUNT_ID) according to the transaction direction.</p><p>If the POST body USER_ID <em>is</em> specified, the logged in user must have the Role CanCreateAccount. Once created, the Account will be owned by the User specified by USER_ID.</p><p>If the POST body USER_ID is <em>not</em> specified, the account will be owned by the logged in User.</p><p>Note: The Amount MUST be zero.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SettlementAccountRequestJson**](SettlementAccountRequestJson.md)| SettlementAccountRequestJson object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**SettlementAccountResponseJson**](SettlementAccountResponseJson.md)

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

# **CreateTransactionType**
> TransactionType CreateTransactionType(ctx, body, bANKID)
Create Transaction Type at bank

<p>Create Transaction Types for the bank specified by BANK_ID:</p><ul><li>id : Unique transaction type id across the API instance. SHOULD be a UUID. MUST be unique.</li><li>bank_id : The bank that supports this TransactionType</li><li>short_code : A short code (SHOULD have no-spaces) which MUST be unique across the bank. May be stored with Transactions to link here</li><li>summary : A succinct summary</li><li>description : A longer description</li><li>charge : The charge to the customer for each one of these</li></ul><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransactionTypeJsonV200**](TransactionTypeJsonV200.md)| TransactionTypeJsonV200 object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**TransactionType**](TransactionType.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBankAttribute**
> DeleteBankAttribute(ctx, bANKID)
Delete Bank Attribute

<p>Delete Bank Attribute</p><p>Delete a Bank Attribute by its id.</p><p>Authentication is Mandatory</p>

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

# **DeleteBankCascade**
> DeleteBankCascade(ctx, bANKID)
Delete Bank Cascade

<p>Delete a Bank Cascade specified by BANK_ID.</p><p>Authentication is Mandatory</p>

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

# **GetBank**
> BankJson500 GetBank(ctx, bANKID)
Get Bank

<p>Get the bank specified by BANK_ID<br />Returns information about a single bank specified by BANK_ID including:</p><ul><li>Bank code and full name of bank</li><li>Logo URL</li><li>Website</li></ul><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**BankJson500**](BankJson500.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBankAttribute**
> BankAttributeResponseJsonV400 GetBankAttribute(ctx, bANKID)
Get Bank Attribute By BANK_ATTRIBUTE_ID

<p>Get Bank Attribute By BANK_ATTRIBUTE_ID</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**BankAttributeResponseJsonV400**](BankAttributeResponseJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBankAttributes**
> BankAttributesResponseJsonV400 GetBankAttributes(ctx, bANKID)
Get Bank Attributes

<p>Get Bank Attributes</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**BankAttributesResponseJsonV400**](BankAttributesResponseJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBanks**
> BanksJson400 GetBanks(ctx, )
Get Banks

<p>Get banks on this API instance<br />Returns a list of banks supported on this server:</p><ul><li>ID used as parameter in URLs</li><li>Short and full name of bank</li><li>Logo URL</li><li>Website</li></ul><p>Authentication is Optional</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**BanksJson400**](BanksJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBranch**
> BranchJsonV300 GetBranch(ctx, bRANCHID, bANKID)
Get Branch

<p>Returns information about a single Branch specified by BANK_ID and BRANCH_ID including:</p><ul><li>Name</li><li>Address</li><li>Geo Location</li><li>License the data under this endpoint is released under.</li></ul><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bRANCHID** | **string**| The branch id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**BranchJsonV300**](BranchJsonV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBranches**
> BranchesJsonV300 GetBranches(ctx, bANKID)
Get Branches for a Bank

<p>Returns information about branches for a single bank specified by BANK_ID including:</p><ul><li>Name</li><li>Address</li><li>Geo Location</li><li>License the data under this endpoint is released under</li><li>Structured opening hours</li><li>Accessible flag</li><li>Branch Type</li><li>More Info</li></ul><p>Pagination:</p><p>By default, 50 records are returned.</p><p>You can use the url query parameters <em>limit</em> and <em>offset</em> for pagination<br />You can also use the follow url query parameters:</p><ul><li><p>city - string, find Branches those in this city, optional</p></li><li><p>withinMetersOf - number, find Branches within given meters distance, optional</p></li><li>nearLatitude - number, a position of latitude value, cooperate with withMetersOf do query filter, optional</li><li>nearLongitude - number, a position of longitude value, cooperate with withMetersOf do query filter, optional</li></ul><p>note: withinMetersOf, nearLatitude and nearLongitude either all empty or all have value.</p><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**BranchesJsonV300**](BranchesJsonV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSettlementAccounts**
> SettlementAccountsJson GetSettlementAccounts(ctx, bANKID)
Get Settlement accounts at Bank

<p>Get settlement accounts on this API instance<br />Returns a list of settlement accounts at this Bank</p><p>Note: a settlement account is considered as a bank account.<br />So you can update it and add account attributes to it using the regular account endpoints</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**SettlementAccountsJson**](SettlementAccountsJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransactionRequestTypesSupportedByBank**
> TransactionRequestTypesJson GetTransactionRequestTypesSupportedByBank(ctx, body, bANKID)
Get Transaction Request Types at Bank

<p>Get the list of the Transaction Request Types supported by the bank.</p><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**TransactionRequestTypesJson**](TransactionRequestTypesJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransactionTypes**
> TransactionTypesJsonV200 GetTransactionTypes(ctx, body, bANKID)
Get Transaction Types at Bank

<p>Get Transaction Types for the bank specified by BANK_ID:</p><p>Lists the possible Transaction Types available at the bank (as opposed to Transaction Request Types which are the possible ways Transactions can be created by this API Server).</p><ul><li>id : Unique transaction type id across the API instance. SHOULD be a UUID. MUST be unique.</li><li>bank_id : The bank that supports this TransactionType</li><li>short_code : A short code (SHOULD have no-spaces) which MUST be unique across the bank. May be stored with Transactions to link here</li><li>summary : A succinct summary</li><li>description : A longer description</li><li>charge : The charge to the customer for each one of these</li></ul><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**TransactionTypesJsonV200**](TransactionTypesJsonV200.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBank**
> BankJson500 UpdateBank(ctx, body)
Update Bank

<p>Update an existing bank (Authenticated access).</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostBankJson500**](PostBankJson500.md)| PostBankJson500 object that needs to be added. | 

### Return type

[**BankJson500**](BankJson500.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBankAttribute**
> AttributeDefinitionJsonV400 UpdateBankAttribute(ctx, body, bANKID)
Update Bank Attribute

<p>Update Bank Attribute.</p><p>Update one Bak Attribute by its id.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BankAttributeJsonV400**](BankAttributeJsonV400.md)| BankAttributeJsonV400 object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**AttributeDefinitionJsonV400**](AttributeDefinitionJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

