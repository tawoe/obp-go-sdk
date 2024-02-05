# \AccountApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAccount**](AccountApi.md#AddAccount) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts | Create Account (POST)
[**AddTagForViewOnAccount**](AccountApi.md#AddTagForViewOnAccount) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/metadata/tags | Create a tag on account
[**CheckFundsAvailable**](AccountApi.md#CheckFundsAvailable) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/funds-available | Check Available Funds
[**CorePrivateAccountsAllBanks**](AccountApi.md#CorePrivateAccountsAllBanks) | **Get** /obp/v5.1.0/my/accounts | Get Accounts at all Banks (private)
[**CreateAccount**](AccountApi.md#CreateAccount) | **Put** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID} | Create Account (PUT)
[**CreateAccountApplication**](AccountApi.md#CreateAccountApplication) | **Post** /obp/v5.1.0/banks/{BANK_ID}/account-applications | Create Account Application
[**CreateAccountAttribute**](AccountApi.md#CreateAccountAttribute) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/products/{PRODUCT_CODE}/attribute | Create Account Attribute
[**CreateCounterparty**](AccountApi.md#CreateCounterparty) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/counterparties | Create Counterparty (Explicit)
[**CreateCounterpartyForAnyAccount**](AccountApi.md#CreateCounterpartyForAnyAccount) | **Post** /obp/v5.1.0/management/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/counterparties | Create Counterparty for any account (Explicit)
[**CreateCustomerAccountLink**](AccountApi.md#CreateCustomerAccountLink) | **Post** /obp/v5.1.0/banks/{BANK_ID}/customer-account-links | Create Customer Account Link
[**CreateDirectDebit**](AccountApi.md#CreateDirectDebit) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/direct-debit | Create Direct Debit
[**CreateDirectDebitManagement**](AccountApi.md#CreateDirectDebitManagement) | **Post** /obp/v5.1.0/management/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/direct-debit | Create Direct Debit (management)
[**CreateOrUpdateAccountAttributeDefinition**](AccountApi.md#CreateOrUpdateAccountAttributeDefinition) | **Put** /obp/v5.1.0/banks/{BANK_ID}/attribute-definitions/account | Create or Update Account Attribute Definition
[**CreateStandingOrder**](AccountApi.md#CreateStandingOrder) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/standing-order | Create Standing Order
[**CreateStandingOrderManagement**](AccountApi.md#CreateStandingOrderManagement) | **Post** /obp/v5.1.0/management/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/standing-order | Create Standing Order (management)
[**CreateUserWithAccountAccess**](AccountApi.md#CreateUserWithAccountAccess) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/user-account-access | Create (DAuth) User with Account Access
[**CreateViewForBankAccount**](AccountApi.md#CreateViewForBankAccount) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/views | Create Custom View
[**DeleteAccountAttributeDefinition**](AccountApi.md#DeleteAccountAttributeDefinition) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/attribute-definitions/ATTRIBUTE_DEFINITION_ID/account | Delete Account Attribute Definition
[**DeleteAccountCascade**](AccountApi.md#DeleteAccountCascade) | **Delete** /obp/v5.1.0/management/cascading/banks/{BANK_ID}/accounts/{ACCOUNT_ID} | Delete Account Cascade
[**DeleteCounterpartyForAnyAccount**](AccountApi.md#DeleteCounterpartyForAnyAccount) | **Delete** /obp/v5.1.0/management/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/counterparties/{COUNTERPARTY_ID} | Delete Counterparty for any account (Explicit)
[**DeleteExplicitCounterparty**](AccountApi.md#DeleteExplicitCounterparty) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/counterparties/{COUNTERPARTY_ID} | Delete Counterparty (Explicit)
[**DeleteTagForViewOnAccount**](AccountApi.md#DeleteTagForViewOnAccount) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/metadata/tags/{TAG_ID} | Delete a tag on account
[**DeleteViewForBankAccount**](AccountApi.md#DeleteViewForBankAccount) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/views/{VIEW_ID} | Delete Custom View
[**GetAccountApplication**](AccountApi.md#GetAccountApplication) | **Get** /obp/v5.1.0/banks/{BANK_ID}/account-applications/{ACCOUNT_APPLICATION_ID} | Get Account Application by Id
[**GetAccountApplications**](AccountApi.md#GetAccountApplications) | **Get** /obp/v5.1.0/banks/{BANK_ID}/account-applications | Get Account Applications
[**GetAccountAttributeDefinition**](AccountApi.md#GetAccountAttributeDefinition) | **Get** /obp/v5.1.0/banks/{BANK_ID}/attribute-definitions/account | Get Account Attribute Definition
[**GetAccountByAccountRouting**](AccountApi.md#GetAccountByAccountRouting) | **Post** /obp/v5.1.0/management/accounts/account-routing-query | Get Account by Account Routing
[**GetAccountsByAccountRoutingRegex**](AccountApi.md#GetAccountsByAccountRoutingRegex) | **Post** /obp/v5.1.0/management/accounts/account-routing-regex-query | Get Accounts by Account Routing Regex
[**GetAccountsHeld**](AccountApi.md#GetAccountsHeld) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts-held | Get Accounts Held
[**GetAccountsMinimalByCustomerId**](AccountApi.md#GetAccountsMinimalByCustomerId) | **Get** /obp/v5.1.0/customers/{CUSTOMER_ID}/accounts-minimal | Get Accounts Minimal for a Customer
[**GetBankAccountBalances**](AccountApi.md#GetBankAccountBalances) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/balances | Get Account Balances
[**GetBankAccountsBalances**](AccountApi.md#GetBankAccountsBalances) | **Get** /obp/v5.1.0/banks/{BANK_ID}/balances | Get Accounts Balances
[**GetCheckbookOrders**](AccountApi.md#GetCheckbookOrders) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/checkbook/orders | Get Checkbook orders
[**GetCoreAccountById**](AccountApi.md#GetCoreAccountById) | **Get** /obp/v5.1.0/my/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/account | Get Account by Id (Core)
[**GetCoreTransactionsForBankAccount**](AccountApi.md#GetCoreTransactionsForBankAccount) | **Get** /obp/v5.1.0/my/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/transactions | Get Transactions for Account (Core)
[**GetCounterpartiesForAnyAccount**](AccountApi.md#GetCounterpartiesForAnyAccount) | **Get** /obp/v5.1.0/management/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/counterparties | Get Counterparties for any account (Explicit)
[**GetCounterpartyByIdForAnyAccount**](AccountApi.md#GetCounterpartyByIdForAnyAccount) | **Get** /obp/v5.1.0/management/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/counterparties/{COUNTERPARTY_ID} | Get Counterparty by Id for any account (Explicit) 
[**GetCounterpartyByNameForAnyAccount**](AccountApi.md#GetCounterpartyByNameForAnyAccount) | **Get** /obp/v5.1.0/management/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/counterparty-names/{COUNTERPARTY_NAME} | Get Counterparty by name for any account (Explicit) 
[**GetExplictCounterpartiesForAccount**](AccountApi.md#GetExplictCounterpartiesForAccount) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/counterparties | Get Counterparties (Explicit)
[**GetFastFirehoseAccountsAtOneBank**](AccountApi.md#GetFastFirehoseAccountsAtOneBank) | **Get** /obp/v5.1.0/management/banks/{BANK_ID}/fast-firehose/accounts | Get Fast Firehose Accounts at Bank
[**GetFirehoseAccountsAtOneBank**](AccountApi.md#GetFirehoseAccountsAtOneBank) | **Get** /obp/v5.1.0/banks/{BANK_ID}/firehose/accounts/views/{VIEW_ID} | Get Firehose Accounts at Bank
[**GetOtherAccountByIdForBankAccount**](AccountApi.md#GetOtherAccountByIdForBankAccount) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/other_accounts/{OTHER_ACCOUNT_ID} | Get Other Account by Id
[**GetOtherAccountsForBankAccount**](AccountApi.md#GetOtherAccountsForBankAccount) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/other_accounts | Get Other Accounts of one Account
[**GetPermissionForUserForBankAccount**](AccountApi.md#GetPermissionForUserForBankAccount) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/permissions/{PROVIDER}/{PROVIDER_ID} | Get Account access for User
[**GetPermissionsForBankAccount**](AccountApi.md#GetPermissionsForBankAccount) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/permissions | Get access
[**GetPrivateAccountByIdFull**](AccountApi.md#GetPrivateAccountByIdFull) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/account | Get Account by Id (Full)
[**GetPrivateAccountIdsbyBankId**](AccountApi.md#GetPrivateAccountIdsbyBankId) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/account_ids/private | Get Accounts at Bank (IDs only)
[**GetPrivateAccountsAtOneBank**](AccountApi.md#GetPrivateAccountsAtOneBank) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts | Get Accounts at Bank
[**GetPublicAccountById**](AccountApi.md#GetPublicAccountById) | **Get** /obp/v5.1.0/banks/{BANK_ID}/public/accounts/{ACCOUNT_ID}/{VIEW_ID}/account | Get Public Account by Id
[**GetTagsForViewOnAccount**](AccountApi.md#GetTagsForViewOnAccount) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/metadata/tags | Get tags on account
[**GetTransactionsForBankAccount**](AccountApi.md#GetTransactionsForBankAccount) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions | Get Transactions for Account (Full)
[**GetViewsForBankAccount**](AccountApi.md#GetViewsForBankAccount) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/views | Get Views for Account
[**GrantUserAccessToView**](AccountApi.md#GrantUserAccessToView) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/account-access/grant | Grant User access to View
[**IbanChecker**](AccountApi.md#IbanChecker) | **Post** /obp/v5.1.0/account/check/scheme/iban | Validate and check IBAN
[**PrivateAccountsAtOneBank**](AccountApi.md#PrivateAccountsAtOneBank) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/private | Get Accounts at Bank (Minimal)
[**PublicAccountsAllBanks**](AccountApi.md#PublicAccountsAllBanks) | **Get** /obp/v5.1.0/accounts/public | Get Public Accounts at all Banks
[**PublicAccountsAtOneBank**](AccountApi.md#PublicAccountsAtOneBank) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/public | Get Public Accounts at Bank
[**RevokeGrantUserAccessToViews**](AccountApi.md#RevokeGrantUserAccessToViews) | **Put** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/account-access | Revoke/Grant User access to View
[**RevokeUserAccessToView**](AccountApi.md#RevokeUserAccessToView) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/account-access/revoke | Revoke User access to View
[**UpdateAccount**](AccountApi.md#UpdateAccount) | **Put** /obp/v5.1.0/management/banks/{BANK_ID}/accounts/{ACCOUNT_ID} | Update Account
[**UpdateAccountApplicationStatus**](AccountApi.md#UpdateAccountApplicationStatus) | **Put** /obp/v5.1.0/banks/{BANK_ID}/account-applications/{ACCOUNT_APPLICATION_ID} | Update Account Application Status
[**UpdateAccountAttribute**](AccountApi.md#UpdateAccountAttribute) | **Put** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/products/{PRODUCT_CODE}/attributes/{ACCOUNT_ATTRIBUTE_ID} | Update Account Attribute
[**UpdateAccountLabel**](AccountApi.md#UpdateAccountLabel) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID} | Update Account Label
[**UpdateViewForBankAccount**](AccountApi.md#UpdateViewForBankAccount) | **Put** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/views/{VIEW_ID} | Update Custom View


# **AddAccount**
> CreateAccountResponseJsonV310 AddAccount(ctx, body, bANKID)
Create Account (POST)

<p>Create Account at bank specified by BANK_ID.</p><p>The User can create an Account for themself  - or -  the User that has the USER_ID specified in the POST body.</p><p>If the POST body USER_ID <em>is</em> specified, the logged in user must have the Role CanCreateAccount. Once created, the Account will be owned by the User specified by USER_ID.</p><p>If the POST body USER_ID is <em>not</em> specified, the account will be owned by the logged in User.</p><p>The 'product_code' field SHOULD be a product_code from Product.<br />If the product_code matches a product_code from Product, account attributes will be created that match the Product Attributes.</p><p>Note: The Amount MUST be zero.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateAccountRequestJsonV310**](CreateAccountRequestJsonV310.md)| CreateAccountRequestJsonV310 object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CreateAccountResponseJsonV310**](CreateAccountResponseJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

# **CheckFundsAvailable**
> CheckFundsAvailableJson CheckFundsAvailable(ctx, vIEWID, aCCOUNTID, bANKID)
Check Available Funds

<p>Check Available Funds<br />Mandatory URL parameters:</p><ul><li>amount=NUMBER</li><li>currency=STRING</li></ul><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CheckFundsAvailableJson**](CheckFundsAvailableJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CorePrivateAccountsAllBanks**
> CoreAccountsJsonV300 CorePrivateAccountsAllBanks(ctx, )
Get Accounts at all Banks (private)

<p>Returns the list of accounts containing private views for the user.<br />Each account lists the views available to the user.</p><p>optional request parameters:</p><ul><li>account_type_filter: one or many accountType value, split by comma</li><li>account_type_filter_operation: the filter type of account_type_filter, value must be INCLUDE or EXCLUDE</li></ul><p>whole url example:<br />/my/accounts?account_type_filter=330,CURRENT+PLUS&amp;account_type_filter_operation=INCLUDE</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CoreAccountsJsonV300**](CoreAccountsJsonV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAccount**
> CreateAccountResponseJsonV310 CreateAccount(ctx, body, aCCOUNTID, bANKID)
Create Account (PUT)

<p>Create Account at bank specified by BANK_ID with Id specified by ACCOUNT_ID.</p><p>The User can create an Account for themself  - or -  the User that has the USER_ID specified in the POST body.</p><p>If the PUT body USER_ID <em>is</em> specified, the logged in user must have the Role canCreateAccount. Once created, the Account will be owned by the User specified by USER_ID.</p><p>If the PUT body USER_ID is <em>not</em> specified, the account will be owned by the logged in User.</p><p>The 'product_code' field SHOULD be a product_code from Product.<br />If the 'product_code' matches a product_code from Product, account attributes will be created that match the Product Attributes.</p><p>Note: The Amount MUST be zero.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateAccountRequestJsonV500**](CreateAccountRequestJsonV500.md)| CreateAccountRequestJsonV500 object that needs to be added. | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CreateAccountResponseJsonV310**](CreateAccountResponseJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAccountApplication**
> AccountApplicationResponseJson CreateAccountApplication(ctx, body, bANKID)
Create Account Application

<p>Create Account Application</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountApplicationJson**](AccountApplicationJson.md)| AccountApplicationJson object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**AccountApplicationResponseJson**](AccountApplicationResponseJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAccountAttribute**
> AccountAttributeResponseJson CreateAccountAttribute(ctx, body, pRODUCTCODE, aCCOUNTID, bANKID)
Create Account Attribute

<p>Create Account Attribute</p><p>Account Attributes are used to describe a financial Product with a list of typed key value pairs.</p><p>Each Account Attribute is linked to its Account by ACCOUNT_ID</p><p>Typical account attributes might be:</p><p>ISIN (for International bonds)<br />VKN (for German bonds)<br />REDCODE (markit short code for credit derivative)<br />LOAN_ID (e.g. used for Anacredit reporting)</p><p>ISSUE_DATE (When the bond was issued in the market)<br />MATURITY_DATE (End of life time of a product)<br />TRADABLE</p><p>See <a href=\"http://www.fpml.org/\">FPML</a> for more examples.</p><p>The type field must be one of &quot;STRING&quot;, &quot;INTEGER&quot;, &quot;DOUBLE&quot; or DATE_WITH_DAY&quot;</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountAttributeJson**](AccountAttributeJson.md)| AccountAttributeJson object that needs to be added. | 
  **pRODUCTCODE** | **string**| the product code | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**AccountAttributeResponseJson**](AccountAttributeResponseJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCounterparty**
> CounterpartyWithMetadataJson400 CreateCounterparty(ctx, body, vIEWID, aCCOUNTID, bANKID)
Create Counterparty (Explicit)

<p>Create Counterparty (Explicit) for an Account.</p><p>In OBP, there are two types of Counterparty.</p><ul><li><p>Explicit Counterparties (those here) which we create explicitly and are used in COUNTERPARTY Transaction Requests</p></li><li><p>Implicit Counterparties (AKA Other Accounts) which are generated automatically from the other sides of Transactions.</p></li></ul><p>Explicit Counterparties are created for the account / view<br />They are how the user of the view (e.g. account owner) refers to the other side of the transaction</p><p>name : the human readable name (e.g. Piano teacher, Miss Nipa)</p><p>description : the human readable name (e.g. Piano teacher, Miss Nipa)</p><p>currency : counterparty account currency (e.g. EUR, GBP, USD, ...)</p><p>bank_routing_scheme : eg: bankId or bankCode or any other strings</p><p>bank_routing_address : eg: <code>gh.29.uk</code>, must be valid sandbox bankIds</p><p>account_routing_scheme : eg: AccountId or AccountNumber or any other strings</p><p>account_routing_address : eg: <code>1d65db7c-a7b2-4839-af41-95</code>, must be valid accountIds</p><p>other_account_secondary_routing_scheme : eg: IBAN or any other strings</p><p>other_account_secondary_routing_address : if it is an IBAN, it should be unique for each counterparty.</p><p>other_branch_routing_scheme : eg: branchId or any other strings or you can leave it empty, not useful in sandbox mode.</p><p>other_branch_routing_address : eg: <code>branch-id-123</code> or you can leave it empty, not useful in sandbox mode.</p><p>is_beneficiary : must be set to <code>true</code> in order to send payments to this counterparty</p><p>bespoke: It supports a list of key-value, you can add it to the counterparty.</p><p>bespoke.key : any info-key you want to add to this counterparty</p><p>bespoke.value : any info-value you want to add to this counterparty</p><p>The view specified by VIEW_ID must have the canAddCounterparty permission</p><p>A minimal example for TransactionRequestType == COUNTERPARTY<br />{<br />&quot;name&quot;: &quot;Tesobe1&quot;,<br />&quot;description&quot;: &quot;Good Company&quot;,<br />&quot;currency&quot;: &quot;EUR&quot;,<br />&quot;other_bank_routing_scheme&quot;: &quot;OBP&quot;,<br />&quot;other_bank_routing_address&quot;: &quot;gh.29.uk&quot;,<br />&quot;other_account_routing_scheme&quot;: &quot;OBP&quot;,<br />&quot;other_account_routing_address&quot;: &quot;8ca8a7e4-6d02-48e3-a029-0b2bf89de9f0&quot;,<br />&quot;is_beneficiary&quot;: true,<br />&quot;other_account_secondary_routing_scheme&quot;: &quot;&quot;,<br />&quot;other_account_secondary_routing_address&quot;: &quot;&quot;,<br />&quot;other_branch_routing_scheme&quot;: &quot;&quot;,<br />&quot;other_branch_routing_address&quot;: &quot;&quot;,<br />&quot;bespoke&quot;: []<br />}</p><p>A minimal example for TransactionRequestType == SEPA</p><p>{<br />&quot;name&quot;: &quot;Tesobe2&quot;,<br />&quot;description&quot;: &quot;Good Company&quot;,<br />&quot;currency&quot;: &quot;EUR&quot;,<br />&quot;other_bank_routing_scheme&quot;: &quot;OBP&quot;,<br />&quot;other_bank_routing_address&quot;: &quot;gh.29.uk&quot;,<br />&quot;other_account_routing_scheme&quot;: &quot;OBP&quot;,<br />&quot;other_account_routing_address&quot;: &quot;8ca8a7e4-6d02-48e3-a029-0b2bf89de9f0&quot;,<br />&quot;other_account_secondary_routing_scheme&quot;: &quot;IBAN&quot;,<br />&quot;other_account_secondary_routing_address&quot;: &quot;DE89 3704 0044 0532 0130 00&quot;,<br />&quot;is_beneficiary&quot;: true,<br />&quot;other_branch_routing_scheme&quot;: &quot;&quot;,<br />&quot;other_branch_routing_address&quot;: &quot;&quot;,<br />&quot;bespoke&quot;: []<br />}</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostCounterpartyJson400**](PostCounterpartyJson400.md)| PostCounterpartyJson400 object that needs to be added. | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CounterpartyWithMetadataJson400**](CounterpartyWithMetadataJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCounterpartyForAnyAccount**
> CounterpartyWithMetadataJson400 CreateCounterpartyForAnyAccount(ctx, body, vIEWID, aCCOUNTID, bANKID)
Create Counterparty for any account (Explicit)

<p>Create Counterparty for any Account. (Explicit)</p><p>In OBP, there are two types of Counterparty.</p><ul><li><p>Explicit Counterparties (those here) which we create explicitly and are used in COUNTERPARTY Transaction Requests</p></li><li><p>Implicit Counterparties (AKA Other Accounts) which are generated automatically from the other sides of Transactions.</p></li></ul><p>Explicit Counterparties are created for the account / view<br />They are how the user of the view (e.g. account owner) refers to the other side of the transaction</p><p>name : the human readable name (e.g. Piano teacher, Miss Nipa)</p><p>description : the human readable name (e.g. Piano teacher, Miss Nipa)</p><p>currency : counterparty account currency (e.g. EUR, GBP, USD, ...)</p><p>bank_routing_scheme : eg: bankId or bankCode or any other strings</p><p>bank_routing_address : eg: <code>gh.29.uk</code>, must be valid sandbox bankIds</p><p>account_routing_scheme : eg: AccountId or AccountNumber or any other strings</p><p>account_routing_address : eg: <code>1d65db7c-a7b2-4839-af41-95</code>, must be valid accountIds</p><p>other_account_secondary_routing_scheme : eg: IBAN or any other strings</p><p>other_account_secondary_routing_address : if it is an IBAN, it should be unique for each counterparty.</p><p>other_branch_routing_scheme : eg: branchId or any other strings or you can leave it empty, not useful in sandbox mode.</p><p>other_branch_routing_address : eg: <code>branch-id-123</code> or you can leave it empty, not useful in sandbox mode.</p><p>is_beneficiary : must be set to <code>true</code> in order to send payments to this counterparty</p><p>bespoke: It supports a list of key-value, you can add it to the counterparty.</p><p>bespoke.key : any info-key you want to add to this counterparty</p><p>bespoke.value : any info-value you want to add to this counterparty</p><p>The view specified by VIEW_ID must have the canAddCounterparty permission</p><p>A minimal example for TransactionRequestType == COUNTERPARTY<br />{<br />&quot;name&quot;: &quot;Tesobe1&quot;,<br />&quot;description&quot;: &quot;Good Company&quot;,<br />&quot;currency&quot;: &quot;EUR&quot;,<br />&quot;other_bank_routing_scheme&quot;: &quot;OBP&quot;,<br />&quot;other_bank_routing_address&quot;: &quot;gh.29.uk&quot;,<br />&quot;other_account_routing_scheme&quot;: &quot;OBP&quot;,<br />&quot;other_account_routing_address&quot;: &quot;8ca8a7e4-6d02-48e3-a029-0b2bf89de9f0&quot;,<br />&quot;is_beneficiary&quot;: true,<br />&quot;other_account_secondary_routing_scheme&quot;: &quot;&quot;,<br />&quot;other_account_secondary_routing_address&quot;: &quot;&quot;,<br />&quot;other_branch_routing_scheme&quot;: &quot;&quot;,<br />&quot;other_branch_routing_address&quot;: &quot;&quot;,<br />&quot;bespoke&quot;: []<br />}</p><p>A minimal example for TransactionRequestType == SEPA</p><p>{<br />&quot;name&quot;: &quot;Tesobe2&quot;,<br />&quot;description&quot;: &quot;Good Company&quot;,<br />&quot;currency&quot;: &quot;EUR&quot;,<br />&quot;other_bank_routing_scheme&quot;: &quot;OBP&quot;,<br />&quot;other_bank_routing_address&quot;: &quot;gh.29.uk&quot;,<br />&quot;other_account_routing_scheme&quot;: &quot;OBP&quot;,<br />&quot;other_account_routing_address&quot;: &quot;8ca8a7e4-6d02-48e3-a029-0b2bf89de9f0&quot;,<br />&quot;other_account_secondary_routing_scheme&quot;: &quot;IBAN&quot;,<br />&quot;other_account_secondary_routing_address&quot;: &quot;DE89 3704 0044 0532 0130 00&quot;,<br />&quot;is_beneficiary&quot;: true,<br />&quot;other_branch_routing_scheme&quot;: &quot;&quot;,<br />&quot;other_branch_routing_address&quot;: &quot;&quot;,<br />&quot;bespoke&quot;: []<br />}</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostCounterpartyJson400**](PostCounterpartyJson400.md)| PostCounterpartyJson400 object that needs to be added. | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CounterpartyWithMetadataJson400**](CounterpartyWithMetadataJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCustomerAccountLink**
> CustomerAccountLinkJson CreateCustomerAccountLink(ctx, body, bANKID)
Create Customer Account Link

<p>Link a Customer to a Account</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateCustomerAccountLinkJson**](CreateCustomerAccountLinkJson.md)| CreateCustomerAccountLinkJson object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerAccountLinkJson**](CustomerAccountLinkJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDirectDebit**
> DirectDebitJsonV400 CreateDirectDebit(ctx, body, vIEWID, aCCOUNTID, bANKID)
Create Direct Debit

<p>Create direct debit for an account.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostDirectDebitJsonV400**](PostDirectDebitJsonV400.md)| PostDirectDebitJsonV400 object that needs to be added. | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**DirectDebitJsonV400**](DirectDebitJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDirectDebitManagement**
> DirectDebitJsonV400 CreateDirectDebitManagement(ctx, body, aCCOUNTID, bANKID)
Create Direct Debit (management)

<p>Create direct debit for an account.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostDirectDebitJsonV400**](PostDirectDebitJsonV400.md)| PostDirectDebitJsonV400 object that needs to be added. | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**DirectDebitJsonV400**](DirectDebitJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateAccountAttributeDefinition**
> AttributeDefinitionResponseJsonV400 CreateOrUpdateAccountAttributeDefinition(ctx, body, bANKID)
Create or Update Account Attribute Definition

<p>Create or Update Account Attribute Definition</p><p>The category field must be Account</p><p>The type field must be one of; DOUBLE, STRING, INTEGER and DATE_WITH_DAY</p><p>Authentication is Mandatory</p>

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

# **CreateStandingOrder**
> StandingOrderJsonV400 CreateStandingOrder(ctx, body, vIEWID, aCCOUNTID, bANKID)
Create Standing Order

<p>Create standing order for an account.</p><p>when -&gt; frequency = {‘YEARLY’,’MONTHLY, ‘WEEKLY’, ‘BI-WEEKLY’, DAILY’}<br />when -&gt; detail = { ‘FIRST_MONDAY’, ‘FIRST_DAY’, ‘LAST_DAY’}}</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostStandingOrderJsonV400**](PostStandingOrderJsonV400.md)| PostStandingOrderJsonV400 object that needs to be added. | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**StandingOrderJsonV400**](StandingOrderJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateStandingOrderManagement**
> StandingOrderJsonV400 CreateStandingOrderManagement(ctx, body, aCCOUNTID, bANKID)
Create Standing Order (management)

<p>Create standing order for an account.</p><p>when -&gt; frequency = {‘YEARLY’,’MONTHLY, ‘WEEKLY’, ‘BI-WEEKLY’, DAILY’}<br />when -&gt; detail = { ‘FIRST_MONDAY’, ‘FIRST_DAY’, ‘LAST_DAY’}}</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostStandingOrderJsonV400**](PostStandingOrderJsonV400.md)| PostStandingOrderJsonV400 object that needs to be added. | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**StandingOrderJsonV400**](StandingOrderJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUserWithAccountAccess**
> NotSupportedYet CreateUserWithAccountAccess(ctx, body, aCCOUNTID, bANKID)
Create (DAuth) User with Account Access

<p>This endpoint is used as part of the DAuth solution to grant access to account and transaction data to a smart contract on the blockchain.</p><p>Put the smart contract address in username</p><p>For provider use &quot;dauth&quot;</p><p>This endpoint will create the (DAuth) User with username and provider if the User does not already exist.</p><p>Authentication is Mandatory and the logged in user needs to be account holder.</p><p>For information about DAuth see below:</p><details>  <summary style=\"display:list-item;cursor:s-resize;\">DAuth</summary>  <h3><a href=\"#dauth-introduction-setup-and-usage\" id=\"dauth-introduction-setup-and-usage\">DAuth Introduction, Setup and Usage</a></h3><p>DAuth is an experimental authentication mechanism that aims to pin an ethereum or other blockchain Smart Contract to an OBP &quot;User&quot;.</p><p>In the future, it might be possible to be more specific and pin specific actors (wallets) that are acting within the smart contract, but so far, one smart contract acts on behalf of one User.</p><p>Thus, if a smart contract &quot;X&quot; calls the OBP API using the DAuth header, OBP will get or create a user called X and the call will proceed in the context of that User &quot;X&quot;.</p><p>DAuth is invoked by the REST client (caller) including a specific header (see step 3 below) in any OBP REST call.</p><p>When OBP receives the DAuth token, it creates or gets a User with a username based on the smart_contract_address and the provider based on the network_name. The combination of username and provider is unique in OBP.</p><p>If you are calling OBP-API via an API3 Airnode, the Airnode will take care of constructing the required header.</p><p>When OBP detects a DAuth header / token it first checks if the Consumer is allowed to make such a call. OBP will validate the Consumer ip address and signature etc.</p><p>Note: The DAuth flow does <em>not</em> require an explicit POST like Direct Login to create the token.</p><p>Permissions may be assigned to an OBP User at any time, via the UserAuthContext, Views, Entitlements to Roles or Consents.</p><p>Note: <em>DAuth is NOT enabled on this instance!</em></p><p>Note: <em>The DAuth client is responsible for creating a token which will be trusted by OBP absolutely</em>!</p><p>To use DAuth:</p><h3><a href=\"#1-configure-obp-api-to-accept-dauth\" id=\"1-configure-obp-api-to-accept-dauth\">1) Configure OBP API to accept DAuth.</a></h3><p>Set up properties in your props file</p><pre><code># -- DAuth --------------------------------------# Define secret used to validate JWT token# jwt.public_key_rsa=path-to-the-pem-file# Enable/Disable DAuth communication at all# In case isn't defined default value is false# allow_dauth=false# Define comma separated list of allowed IP addresses# dauth.host=127.0.0.1# -------------------------------------- DAuth--</code></pre><p>Please keep in mind that property jwt.public_key_rsa is used to validate JWT token to check it is not changed or corrupted during transport.</p><h3><a href=\"#2-create-have-access-to-a-jwt\" id=\"2-create-have-access-to-a-jwt\">2) Create / have access to a JWT</a></h3><p>The following videos are available:<br />* <a href=\"https://vimeo.com/644315074\">DAuth in local environment</a></p><p>HEADER:ALGORITHM &amp; TOKEN TYPE</p><pre><code>{  &quot;alg&quot;: &quot;RS256&quot;,  &quot;typ&quot;: &quot;JWT&quot;}</code></pre><p>PAYLOAD:DATA</p><pre><code>{  &quot;smart_contract_address&quot;: &quot;0xe123425E7734CE288F8367e1Bb143E90bb3F051224&quot;,  &quot;network_name&quot;: &quot;AIRNODE.TESTNET.ETHEREUM&quot;,  &quot;msg_sender&quot;: &quot;0xe12340927f1725E7734CE288F8367e1Bb143E90fhku767&quot;,  &quot;consumer_key&quot;: &quot;0x1234a4ec31e89cea54d1f125db7536e874ab4a96b4d4f6438668b6bb10a6adb&quot;,  &quot;timestamp&quot;: &quot;2021-11-04T14:13:40Z&quot;,  &quot;request_id&quot;: &quot;0Xe876987694328763492876348928736497869273649&quot;}</code></pre><p>VERIFY SIGNATURE</p><pre><code>RSASHA256(  base64UrlEncode(header) + &quot;.&quot; +  base64UrlEncode(payload),<p>) your-RSA-key-pair</p></code></pre><p>Here is an example token:</p><pre><code>eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzbWFydF9jb250cmFjdF9hZGRyZXNzIjoiMHhlMTIzNDI1RTc3MzRDRTI4OEY4MzY3ZTFCYjE0M0U5MGJiM0YwNTEyMjQiLCJuZXR3b3JrX25hbWUiOiJFVEhFUkVVTSIsIm1zZ19zZW5kZXIiOiIweGUxMjM0MDkyN2YxNzI1RTc3MzRDRTI4OEY4MzY3ZTFCYjE0M0U5MGZoa3U3NjciLCJjb25zdW1lcl9rZXkiOiIweDEyMzRhNGVjMzFlODljZWE1NGQxZjEyNWRiNzUzNmU4NzRhYjRhOTZiNGQ0ZjY0Mzg2NjhiNmJiMTBhNmFkYiIsInRpbWVzdGFtcCI6IjIwMjEtMTEtMDRUMTQ6MTM6NDBaIiwicmVxdWVzdF9pZCI6IjBYZTg3Njk4NzY5NDMyODc2MzQ5Mjg3NjM0ODkyODczNjQ5Nzg2OTI3MzY0OSJ9.XSiQxjEVyCouf7zT8MubEKsbOBZuReGVhnt9uck6z6k</code></pre><h3><a href=\"#3-try-a-rest-call-using-the-header\" id=\"3-try-a-rest-call-using-the-header\">3) Try a REST call using the header</a></h3><p>Using your favorite http client:</p><p>GET <a href=\"https://apisandbox.openbankproject.com/obp/v3.0.0/users/current\">https://apisandbox.openbankproject.com/obp/v3.0.0/users/current</a></p><p>Body</p><p>Leave Empty!</p><p>Headers:</p><pre><code>   DAuth: your-jwt-from-step-above</code></pre><p>Here is it all together:</p><p>GET <a href=\"https://apisandbox.openbankproject.com/obp/v3.0.0/users/current\">https://apisandbox.openbankproject.com/obp/v3.0.0/users/current</a> HTTP/1.1<br />Host: localhost:8080<br />User-Agent: curl/7.47.0<br />Accept: <em>/</em><br />DAuth: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzbWFydF9jb250cmFjdF9hZGRyZXNzIjoiMHhlMTIzNDI1RTc3MzRDRTI4OEY4MzY3ZTFCYjE0M0U5MGJiM0YwNTEyMjQiLCJuZXR3b3JrX25hbWUiOiJFVEhFUkVVTSIsIm1zZ19zZW5kZXIiOiIweGUxMjM0MDkyN2YxNzI1RTc3MzRDRTI4OEY4MzY3ZTFCYjE0M0U5MGZoa3U3NjciLCJjb25zdW1lcl9rZXkiOiIweDEyMzRhNGVjMzFlODljZWE1NGQxZjEyNWRiNzUzNmU4NzRhYjRhOTZiNGQ0ZjY0Mzg2NjhiNmJiMTBhNmFkYiIsInRpbWVzdGFtcCI6IjIwMjEtMTEtMDRUMTQ6MTM6NDBaIiwicmVxdWVzdF9pZCI6IjBYZTg3Njk4NzY5NDMyODc2MzQ5Mjg3NjM0ODkyODczNjQ5Nzg2OTI3MzY0OSJ9.XSiQxjEVyCouf7zT8MubEKsbOBZuReGVhnt9uck6z6k</p><p>CURL example</p><pre><code>curl -v -H 'DAuth: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzbWFydF9jb250cmFjdF9hZGRyZXNzIjoiMHhlMTIzNDI1RTc3MzRDRTI4OEY4MzY3ZTFCYjE0M0U5MGJiM0YwNTEyMjQiLCJuZXR3b3JrX25hbWUiOiJFVEhFUkVVTSIsIm1zZ19zZW5kZXIiOiIweGUxMjM0MDkyN2YxNzI1RTc3MzRDRTI4OEY4MzY3ZTFCYjE0M0U5MGZoa3U3NjciLCJjb25zdW1lcl9rZXkiOiIweDEyMzRhNGVjMzFlODljZWE1NGQxZjEyNWRiNzUzNmU4NzRhYjRhOTZiNGQ0ZjY0Mzg2NjhiNmJiMTBhNmFkYiIsInRpbWVzdGFtcCI6IjIwMjEtMTEtMDRUMTQ6MTM6NDBaIiwicmVxdWVzdF9pZCI6IjBYZTg3Njk4NzY5NDMyODc2MzQ5Mjg3NjM0ODkyODczNjQ5Nzg2OTI3MzY0OSJ9.XSiQxjEVyCouf7zT8MubEKsbOBZuReGVhnt9uck6z6k' https://apisandbox.openbankproject.com/obp/v3.0.0/users/current</code></pre><p>You should receive a response like:</p><pre><code>{    &quot;user_id&quot;: &quot;4c4d3175-1e5c-4cfd-9b08-dcdc209d8221&quot;,    &quot;email&quot;: &quot;&quot;,    &quot;provider_id&quot;: &quot;0xe123425E7734CE288F8367e1Bb143E90bb3F051224&quot;,    &quot;provider&quot;: &quot;ETHEREUM&quot;,    &quot;username&quot;: &quot;0xe123425E7734CE288F8367e1Bb143E90bb3F051224&quot;,    &quot;entitlements&quot;: {        &quot;list&quot;: []    }}</code></pre><h3><a href=\"#under-the-hood\" id=\"under-the-hood\">Under the hood</a></h3><p>The file, dauth.scala handles the DAuth,</p><p>We:</p><pre><code>-&gt; Check if Props allow_dauth is true  -&gt; Check if DAuth header exists    -&gt; Check if getRemoteIpAddress is OK      -&gt; Look for &quot;token&quot;        -&gt; parse the JWT token and getOrCreate the user          -&gt; get the data of the user</code></pre><h3><a href=\"#more-information\" id=\"more-information\">More information</a></h3><p>Parameter names and values are case sensitive.<br />Each parameter MUST NOT appear more than once per request.</p></details><p><br></br></p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostCreateUserAccountAccessJsonV400**](PostCreateUserAccountAccessJsonV400.md)| PostCreateUserAccountAccessJsonV400 object that needs to be added. | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**NotSupportedYet**](NotSupportedYet.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateViewForBankAccount**
> ViewJsonV300 CreateViewForBankAccount(ctx, body, aCCOUNTID, bANKID)
Create Custom View

<p>Create a custom view on bank account</p><p>Authentication is Mandatory and the user needs to have access to the owner view.<br />The 'alias' field in the JSON can take one of three values:</p><ul><li><em>public</em>: to use the public alias if there is one specified for the other account.</li><li><em>private</em>: to use the public alias if there is one specified for the other account.</li><li><p><em>''(empty string)</em>: to use no alias; the view shows the real name of the other account.</p></li></ul><p>The 'hide_metadata_if_alias_used' field in the JSON can take boolean values. If it is set to <code>true</code> and there is an alias on the other account then the other accounts' metadata (like more_info, url, image_url, open_corporates_url, etc.) will be hidden. Otherwise the metadata will be shown.</p><p>The 'allowed_actions' field is a list containing the name of the actions allowed on this view, all the actions contained will be set to <code>true</code> on the view creation, the rest will be set to <code>false</code>.</p><p>You MUST use a leading _ (underscore) in the view name because other view names are reserved for OBP <a href=\"/index#group-View-System\">system views</a>.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateViewJsonV300**](CreateViewJsonV300.md)| CreateViewJsonV300 object that needs to be added. | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**ViewJsonV300**](ViewJsonV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccountAttributeDefinition**
> DeleteAccountAttributeDefinition(ctx, bANKID)
Delete Account Attribute Definition

<p>Delete Account Attribute Definition by ATTRIBUTE_DEFINITION_ID</p><p>Authentication is Mandatory</p>

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

# **DeleteAccountCascade**
> DeleteAccountCascade(ctx, aCCOUNTID, bANKID)
Delete Account Cascade

<p>Delete an Account Cascade specified by ACCOUNT_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

# **DeleteCounterpartyForAnyAccount**
> DeleteCounterpartyForAnyAccount(ctx, cOUNTERPARTYID, vIEWID, aCCOUNTID, bANKID)
Delete Counterparty for any account (Explicit)

<p>Delete Counterparty (Explicit) for any account<br />and also delete the Metadata for its counterparty.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cOUNTERPARTYID** | **string**| the counterparty id | 
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

# **DeleteExplicitCounterparty**
> DeleteExplicitCounterparty(ctx, cOUNTERPARTYID, vIEWID, aCCOUNTID, bANKID)
Delete Counterparty (Explicit)

<p>Delete Counterparty (Explicit) for an Account.<br />and also delete the Metadata for its counterparty.</p><p>need the view permission <code>can_delete_counterparty</code><br />Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cOUNTERPARTYID** | **string**| the counterparty id | 
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

# **DeleteViewForBankAccount**
> EmptyClassJson DeleteViewForBankAccount(ctx, body, vIEWID, aCCOUNTID, bANKID)
Delete Custom View

<p>Deletes the custom view specified by VIEW_ID on the bank account specified by ACCOUNT_ID at bank BANK_ID</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
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

# **GetAccountApplication**
> AccountApplicationResponseJson GetAccountApplication(ctx, aCCOUNTAPPLICATIONID, bANKID)
Get Account Application by Id

<p>Get the Account Application.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aCCOUNTAPPLICATIONID** | **string**| the account application id  | 
  **bANKID** | **string**| The bank id | 

### Return type

[**AccountApplicationResponseJson**](AccountApplicationResponseJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountApplications**
> AccountApplicationsJsonV310 GetAccountApplications(ctx, bANKID)
Get Account Applications

<p>Get the Account Applications.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**AccountApplicationsJsonV310**](AccountApplicationsJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountAttributeDefinition**
> AttributeDefinitionsResponseJsonV400 GetAccountAttributeDefinition(ctx, bANKID)
Get Account Attribute Definition

<p>Get Account Attribute Definition</p><p>Authentication is Mandatory</p>

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

# **GetAccountByAccountRouting**
> ModeratedAccountJson400 GetAccountByAccountRouting(ctx, body)
Get Account by Account Routing

<p>This endpoint returns the account (if it exists) linked with the provided scheme and address.</p><p>The <code>bank_id</code> field is optional, but if it's not provided, we don't guarantee that the returned account is unique across all the banks.</p><p>Example of account routing scheme: <code>IBAN</code>, &quot;OBP&quot;, &quot;AccountNumber&quot;, ...<br />Example of account routing address: <code>DE17500105178275645584</code>, &quot;321774cc-fccd-11ea-adc1-0242ac120002&quot;, &quot;55897106215&quot;, ...</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BankAccountRoutingJson**](BankAccountRoutingJson.md)| BankAccountRoutingJson object that needs to be added. | 

### Return type

[**ModeratedAccountJson400**](ModeratedAccountJSON400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountsByAccountRoutingRegex**
> ModeratedAccountsJson400 GetAccountsByAccountRoutingRegex(ctx, body)
Get Accounts by Account Routing Regex

<p>This endpoint returns an array of accounts matching the provided routing scheme and the routing address regex.</p><p>The <code>bank_id</code> field is optional.</p><p>Example of account routing scheme: <code>IBAN</code>, <code>OBP</code>, <code>AccountNumber</code>, ...<br />Example of account routing address regex: <code>DE175.*</code>, <code>55897106215-[A-Z]{3}</code>, ...</p><p>This endpoint can be used to retrieve multiples accounts matching a same account routing address pattern.<br />For example, if you want to link multiple accounts having different currencies, you can create an account<br />with <code>123456789-EUR</code> as Account Number and an other account with <code>123456789-USD</code> as Account Number.<br />So we can identify the Account Number as <code>123456789</code>, so to get all the accounts with the same account number<br />and the different currencies, we can use this body in the request :</p><pre><code>{   &quot;bank_id&quot;: &quot;BANK_ID&quot;,   &quot;account_routing&quot;: {       &quot;scheme&quot;: &quot;AccountNumber&quot;,       &quot;address&quot;: &quot;123456789-[A-Z]{3}&quot;   }}</code></pre><p>This request will returns the accounts matching the routing address regex (<code>123456789-EUR</code> and <code>123456789-USD</code>).</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BankAccountRoutingJson**](BankAccountRoutingJson.md)| BankAccountRoutingJson object that needs to be added. | 

### Return type

[**ModeratedAccountsJson400**](ModeratedAccountsJSON400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountsHeld**
> CoreAccountsHeldJsonV300 GetAccountsHeld(ctx, bANKID)
Get Accounts Held

<p>Get Accounts held by the current User if even the User has not been assigned the owner View yet.</p><p>Can be used to onboard the account to the API - since all other account and transaction endpoints require views to be assigned.</p><p>optional request parameters:</p><ul><li>account_type_filter: one or many accountType value, split by comma</li><li>account_type_filter_operation: the filter type of account_type_filter, value must be INCLUDE or EXCLUDE</li></ul><p>whole url example:<br />/banks/BANK_ID/accounts-held?account_type_filter=330,CURRENT+PLUS&amp;account_type_filter_operation=INCLUDE</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**CoreAccountsHeldJsonV300**](CoreAccountsHeldJsonV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountsMinimalByCustomerId**
> AccountsMinimalJson400 GetAccountsMinimalByCustomerId(ctx, cUSTOMERID)
Get Accounts Minimal for a Customer

<p>Get Accounts Minimal by CUSTOMER_ID</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cUSTOMERID** | **string**| The customer id | 

### Return type

[**AccountsMinimalJson400**](AccountsMinimalJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBankAccountBalances**
> AccountBalanceJsonV400 GetBankAccountBalances(ctx, aCCOUNTID, bANKID)
Get Account Balances

<p>Get the Balances for one Account of the current User at one bank.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**AccountBalanceJsonV400**](AccountBalanceJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBankAccountsBalances**
> AccountsBalancesJsonV400 GetBankAccountsBalances(ctx, bANKID)
Get Accounts Balances

<p>Get the Balances for the Accounts of the current User at one bank.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**AccountsBalancesJsonV400**](AccountsBalancesJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCheckbookOrders**
> CheckbookOrdersJson GetCheckbookOrders(ctx, vIEWID, aCCOUNTID, bANKID)
Get Checkbook orders

<pre><code>  Get all checkbook orders</code></pre><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CheckbookOrdersJson**](CheckbookOrdersJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCoreAccountById**
> ModeratedCoreAccountJsonV400 GetCoreAccountById(ctx, aCCOUNTID, bANKID)
Get Account by Id (Core)

<p>Information returned about the account specified by ACCOUNT_ID:</p><ul><li>Number - The human readable account number given by the bank that identifies the account.</li><li>Label - A label given by the owner of the account</li><li>Owners - Users that own this account</li><li>Type - The type of account</li><li>Balance - Currency and Value</li><li>Account Routings - A list that might include IBAN or national account identifiers</li><li>Account Rules - A list that might include Overdraft and other bank specific rules</li><li>Tags - A list of Tags assigned to this account</li></ul><p>This call returns the owner view and requires access to that view.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**ModeratedCoreAccountJsonV400**](ModeratedCoreAccountJsonV400.md)

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

# **GetCounterpartiesForAnyAccount**
> CounterpartiesJson400 GetCounterpartiesForAnyAccount(ctx, vIEWID, aCCOUNTID, bANKID)
Get Counterparties for any account (Explicit)

<p>Get the Counterparties (Explicit) for any account .</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CounterpartiesJson400**](CounterpartiesJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCounterpartyByIdForAnyAccount**
> CounterpartyWithMetadataJson400 GetCounterpartyByIdForAnyAccount(ctx, cOUNTERPARTYID, vIEWID, aCCOUNTID, bANKID)
Get Counterparty by Id for any account (Explicit) 

<p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cOUNTERPARTYID** | **string**| the counterparty id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CounterpartyWithMetadataJson400**](CounterpartyWithMetadataJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCounterpartyByNameForAnyAccount**
> CounterpartyWithMetadataJson400 GetCounterpartyByNameForAnyAccount(ctx, cOUNTERPARTYNAME, vIEWID, aCCOUNTID, bANKID)
Get Counterparty by name for any account (Explicit) 

<p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cOUNTERPARTYNAME** | **string**| the counterparty name | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CounterpartyWithMetadataJson400**](CounterpartyWithMetadataJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExplictCounterpartiesForAccount**
> CounterpartiesJson400 GetExplictCounterpartiesForAccount(ctx, vIEWID, aCCOUNTID, bANKID)
Get Counterparties (Explicit)

<p>Get the Counterparties (Explicit) for the account / view.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CounterpartiesJson400**](CounterpartiesJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFastFirehoseAccountsAtOneBank**
> FastFirehoseAccountsJsonV400 GetFastFirehoseAccountsAtOneBank(ctx, bANKID)
Get Fast Firehose Accounts at Bank

<p>This endpoint allows bulk access to accounts.</p><p>optional pagination parameters for filter with accounts</p><p>Possible custom url parameters for pagination:</p><ul><li>limit=NUMBER ==&gt; default value: 500</li><li>offset=NUMBER ==&gt; default value: 0</li></ul><p>eg1:?limit=100&amp;offset=0</p><ul><li>sort_direction=ASC/DESC ==&gt; default value: DESC.</li></ul><p>eg2:?limit=100&amp;offset=0&amp;sort_direction=ASC</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**FastFirehoseAccountsJsonV400**](FastFirehoseAccountsJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFirehoseAccountsAtOneBank**
> ModeratedFirehoseAccountsJsonV400 GetFirehoseAccountsAtOneBank(ctx, vIEWID, bANKID)
Get Firehose Accounts at Bank

<p>Get Accounts which have a firehose view assigned to them.</p><p>This endpoint allows bulk access to accounts.</p><p>Requires the CanUseFirehoseAtAnyBank Role</p><p>To be shown on the list, each Account must have a firehose View linked to it.</p><p>A firehose view has is_firehose = true</p><p>For VIEW_ID try 'owner'</p><p>optional request parameters for filter with attributes<br />URL params example:<br />/banks/some-bank-id/firehose/accounts/views/owner?manager=John&amp;count=8</p><p>to invalid Browser cache, add timestamp query parameter as follow, the parameter name must be <code>_timestamp_</code><br />URL params example:<br /><code>/banks/some-bank-id/firehose/accounts/views/owner?manager=John&amp;count=8&amp;_timestamp_=1596762180358</code></p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vIEWID** | **string**| The view id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**ModeratedFirehoseAccountsJsonV400**](ModeratedFirehoseAccountsJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOtherAccountByIdForBankAccount**
> OtherAccountJsonV300 GetOtherAccountByIdForBankAccount(ctx, oTHERACCOUNTID, vIEWID, aCCOUNTID, bANKID)
Get Other Account by Id

<p>Returns data about the Other Account that has shared at least one transaction with ACCOUNT_ID at BANK_ID.<br />Authentication is Optional</p><p>Authentication is required if the view is not public.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oTHERACCOUNTID** | **string**| The other account id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**OtherAccountJsonV300**](OtherAccountJsonV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOtherAccountsForBankAccount**
> OtherAccountsJsonV300 GetOtherAccountsForBankAccount(ctx, vIEWID, aCCOUNTID, bANKID)
Get Other Accounts of one Account

<p>Returns data about all the other accounts that have shared at least one transaction with the ACCOUNT_ID at BANK_ID.<br />Authentication is Optional</p><p>Authentication is required if the view VIEW_ID is not public.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**OtherAccountsJsonV300**](OtherAccountsJsonV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPermissionForUserForBankAccount**
> ViewsJsonV300 GetPermissionForUserForBankAccount(ctx, pROVIDER, pROVIDERID, aCCOUNTID, bANKID)
Get Account access for User

<p>Returns the list of the views at BANK_ID for account ACCOUNT_ID that a user identified by PROVIDER_ID at their provider PROVIDER has access to.<br />All url parameters must be <a href=\"http://en.wikipedia.org/wiki/Percent-encoding\">%-encoded</a>, which is often especially relevant for USER_ID and PROVIDER.</p><p>Authentication is Mandatory</p><p>The user needs to have access to the owner view.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pROVIDER** | **string**| the user PROVIDER | 
  **pROVIDERID** | **string**| The provider id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**ViewsJsonV300**](ViewsJsonV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPermissionsForBankAccount**
> PermissionsJson GetPermissionsForBankAccount(ctx, body, aCCOUNTID, bANKID)
Get access

<p>Returns the list of the permissions at BANK_ID for account ACCOUNT_ID, with each time a pair composed of the user and the views that he has access to.</p><p>Authentication is Mandatory<br />and the user needs to have access to the owner view.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**PermissionsJson**](PermissionsJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPrivateAccountByIdFull**
> ModeratedAccountJson400 GetPrivateAccountByIdFull(ctx, vIEWID, aCCOUNTID, bANKID)
Get Account by Id (Full)

<p>Information returned about an account specified by ACCOUNT_ID as moderated by the view (VIEW_ID):</p><ul><li>Number</li><li>Owners</li><li>Type</li><li>Balance</li><li>IBAN</li><li>Available views (sorted by short_name)</li></ul><p>More details about the data moderation by the view <a href=\"#1_2_1-getViewsForBankAccount\">here</a>.</p><p>PSD2 Context: PSD2 requires customers to have access to their account information via third party applications.<br />This call provides balance and other account information via delegated authentication using OAuth.</p><p>Authentication is required if the 'is_public' field in view (VIEW_ID) is not set to <code>true</code>.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**ModeratedAccountJson400**](ModeratedAccountJSON400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPrivateAccountIdsbyBankId**
> AccountsIdsJsonV300 GetPrivateAccountIdsbyBankId(ctx, bANKID)
Get Accounts at Bank (IDs only)

<p>Returns only the list of accounts ids at BANK_ID that the user has access to.</p><p>Each account must have at least one private View.</p><p>For each account the API returns its account ID.</p><p>If you want to see more information on the Views, use the Account Detail call.</p><p>optional request parameters:</p><ul><li>account_type_filter: one or many accountType value, split by comma</li><li>account_type_filter_operation: the filter type of account_type_filter, value must be INCLUDE or EXCLUDE</li></ul><p>whole url example:<br />/banks/BANK_ID/accounts/account_ids/private?account_type_filter=330,CURRENT+PLUS&amp;account_type_filter_operation=INCLUDE</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**AccountsIdsJsonV300**](AccountsIdsJsonV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPrivateAccountsAtOneBank**
> BasicAccountsJson GetPrivateAccountsAtOneBank(ctx, bANKID)
Get Accounts at Bank

<p>Returns the list of accounts at BANK_ID that the user has access to.<br />For each account the API returns the account ID and the views available to the user..<br />Each account must have at least one private View.</p><p>optional request parameters for filter with attributes<br />URL params example: /banks/some-bank-id/accounts?manager=John&amp;count=8</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**BasicAccountsJson**](BasicAccountsJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPublicAccountById**
> ModeratedCoreAccountJsonV300 GetPublicAccountById(ctx, vIEWID, aCCOUNTID, bANKID)
Get Public Account by Id

<p>Returns information about an account that has a public view.</p><p>The account is specified by ACCOUNT_ID. The information is moderated by the view specified by VIEW_ID.</p><ul><li>Number</li><li>Owners</li><li>Type</li><li>Balance</li><li>Routing</li></ul><p>PSD2 Context: PSD2 requires customers to have access to their account information via third party applications.<br />This call provides balance and other account information via delegated authentication using OAuth.</p><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**ModeratedCoreAccountJsonV300**](ModeratedCoreAccountJsonV300.md)

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

# **GetViewsForBankAccount**
> ViewsJsonV500 GetViewsForBankAccount(ctx, aCCOUNTID, bANKID)
Get Views for Account

<h1><a href=\"#views\" id=\"views\">Views</a></h1><p>Views in Open Bank Project provide a mechanism for fine grained access control and delegation to Accounts and Transactions. Account holders use the 'owner' view by default. Delegated access is made through other views for example 'accountants', 'share-holders' or 'tagging-application'. Views can be created via the API and each view has a list of entitlements.</p><p>Views on accounts and transactions filter the underlying data to redact certain fields for certain users. For instance the balance on an account may be hidden from the public. The way to know what is possible on a view is determined in the following JSON.</p><p><strong>Data:</strong> When a view moderates a set of data, some fields my contain the value <code>null</code> rather than the original value. This indicates either that the user is not allowed to see the original data or the field is empty.</p><p>There is currently one exception to this rule; the 'holder' field in the JSON contains always a value which is either an alias or the real name - indicated by the 'is_alias' field.</p><p><strong>Action:</strong> When a user performs an action like trying to post a comment (with POST API call), if he is not allowed, the body response will contain an error message.</p><p><strong>Metadata:</strong><br />Transaction metadata (like images, tags, comments, etc.) will appears <em>ONLY</em> on the view where they have been created e.g. comments posted to the public view only appear on the public view.</p><p>The other account metadata fields (like image_URL, more_info, etc.) are unique through all the views. Example, if a user edits the 'more_info' field in the 'team' view, then the view 'authorities' will show the new value (if it is allowed to do it).</p><h1><a href=\"#all\" id=\"all\">All</a></h1><p><em>Optional</em></p><p>Returns the list of the views created for account ACCOUNT_ID at BANK_ID.</p><p>Authentication is Mandatory and the user needs to have access to the owner view.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**ViewsJsonV500**](ViewsJsonV500.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GrantUserAccessToView**
> ViewJsonV300 GrantUserAccessToView(ctx, body, aCCOUNTID, bANKID)
Grant User access to View

<p>Grants the User identified by USER_ID access to the view identified by VIEW_ID.</p><p>Authentication is Mandatory and the user needs to be account holder.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostAccountAccessJsonV400**](PostAccountAccessJsonV400.md)| PostAccountAccessJsonV400 object that needs to be added. | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**ViewJsonV300**](ViewJsonV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IbanChecker**
> IbanCheckerJsonV400 IbanChecker(ctx, body)
Validate and check IBAN

<p>Validate and check IBAN for errors</p><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IbanAddress**](IbanAddress.md)| IbanAddress object that needs to be added. | 

### Return type

[**IbanCheckerJsonV400**](IbanCheckerJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateAccountsAtOneBank**
> CoreAccountsJsonV300 PrivateAccountsAtOneBank(ctx, bANKID)
Get Accounts at Bank (Minimal)

<p>Returns the minimal list of private accounts at BANK_ID that the user has access to.<br />For each account, the API returns the ID, routing addresses and the views available to the current user.</p><p>If you want to see more information on the Views, use the Account Detail call.</p><p>optional request parameters:</p><ul><li>account_type_filter: one or many accountType value, split by comma</li><li>account_type_filter_operation: the filter type of account_type_filter, value must be INCLUDE or EXCLUDE</li></ul><p>whole url example:<br />/banks/BANK_ID/accounts/private?account_type_filter=330,CURRENT+PLUS&amp;account_type_filter_operation=INCLUDE</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**CoreAccountsJsonV300**](CoreAccountsJsonV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublicAccountsAllBanks**
> BasicAccountsJson PublicAccountsAllBanks(ctx, body)
Get Public Accounts at all Banks

<p>Get public accounts at all banks (Anonymous access).<br />Returns accounts that contain at least one public view (a view where is_public is true)<br />For each account the API returns the ID and the available views.</p><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 

### Return type

[**BasicAccountsJson**](BasicAccountsJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublicAccountsAtOneBank**
> BasicAccountsJson PublicAccountsAtOneBank(ctx, body, bANKID)
Get Public Accounts at Bank

<p>Returns a list of the public accounts (Anonymous access) at BANK_ID. For each account the API returns the ID and the available views.</p><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**BasicAccountsJson**](BasicAccountsJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeGrantUserAccessToViews**
> RevokedJsonV400 RevokeGrantUserAccessToViews(ctx, body, aCCOUNTID, bANKID)
Revoke/Grant User access to View

<p>Revoke/Grant the logged in User access to the views identified by json.</p><p>Authentication is Mandatory and the user needs to be an account holder or has owner view access.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostRevokeGrantAccountAccessJsonV400**](PostRevokeGrantAccountAccessJsonV400.md)| PostRevokeGrantAccountAccessJsonV400 object that needs to be added. | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**RevokedJsonV400**](RevokedJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeUserAccessToView**
> RevokedJsonV400 RevokeUserAccessToView(ctx, body, aCCOUNTID, bANKID)
Revoke User access to View

<p>Revoke the User identified by USER_ID access to the view identified by VIEW_ID.</p><p>Authentication is Mandatory and the user needs to be account holder.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostAccountAccessJsonV400**](PostAccountAccessJsonV400.md)| PostAccountAccessJsonV400 object that needs to be added. | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**RevokedJsonV400**](RevokedJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccount**
> UpdateAccountResponseJsonV310 UpdateAccount(ctx, body, aCCOUNTID, bANKID)
Update Account

<p>Update the account.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateAccountRequestJsonV310**](UpdateAccountRequestJsonV310.md)| UpdateAccountRequestJsonV310 object that needs to be added. | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**UpdateAccountResponseJsonV310**](UpdateAccountResponseJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccountApplicationStatus**
> AccountApplicationResponseJson UpdateAccountApplicationStatus(ctx, body, aCCOUNTAPPLICATIONID, bANKID)
Update Account Application Status

<p>Update an Account Application status</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountApplicationUpdateStatusJson**](AccountApplicationUpdateStatusJson.md)| AccountApplicationUpdateStatusJson object that needs to be added. | 
  **aCCOUNTAPPLICATIONID** | **string**| the account application id  | 
  **bANKID** | **string**| The bank id | 

### Return type

[**AccountApplicationResponseJson**](AccountApplicationResponseJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccountAttribute**
> AccountAttributeResponseJson UpdateAccountAttribute(ctx, body, aCCOUNTATTRIBUTEID, pRODUCTCODE, aCCOUNTID, bANKID)
Update Account Attribute

<p>Update Account Attribute</p><p>Account Attributes are used to describe a financial Product with a list of typed key value pairs.</p><p>Each Account Attribute is linked to its Account by ACCOUNT_ID</p><p>Typical account attributes might be:</p><p>ISIN (for International bonds)<br />VKN (for German bonds)<br />REDCODE (markit short code for credit derivative)<br />LOAN_ID (e.g. used for Anacredit reporting)</p><p>ISSUE_DATE (When the bond was issued in the market)<br />MATURITY_DATE (End of life time of a product)<br />TRADABLE</p><p>See <a href=\"http://www.fpml.org/\">FPML</a> for more examples.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountAttributeJson**](AccountAttributeJson.md)| AccountAttributeJson object that needs to be added. | 
  **aCCOUNTATTRIBUTEID** | **string**| the account attribute id  | 
  **pRODUCTCODE** | **string**| the product code | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**AccountAttributeResponseJson**](AccountAttributeResponseJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccountLabel**
> SuccessMessage UpdateAccountLabel(ctx, body, aCCOUNTID, bANKID)
Update Account Label

<p>Update the label for the account. The label is how the account is known to the account owner e.g. 'My savings account'</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateAccountJsonV400**](UpdateAccountJsonV400.md)| UpdateAccountJsonV400 object that needs to be added. | 
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

# **UpdateViewForBankAccount**
> ViewJsonV300 UpdateViewForBankAccount(ctx, body, vIEWID, aCCOUNTID, bANKID)
Update Custom View

<p>Update an existing custom view on a bank account</p><p>Authentication is Mandatory and the user needs to have access to the owner view.</p><p>The json sent is the same as during view creation (above), with one difference: the 'name' field<br />of a view is not editable (it is only set when a view is created)</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateViewJsonV300**](UpdateViewJsonV300.md)| UpdateViewJsonV300 object that needs to be added. | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**ViewJsonV300**](ViewJsonV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

