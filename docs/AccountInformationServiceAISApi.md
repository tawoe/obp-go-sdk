# \AccountInformationServiceAISApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddConsentUser**](AccountInformationServiceAISApi.md#AddConsentUser) | **Put** /obp/v5.1.0/banks/{BANK_ID}/consents/{CONSENT_ID}/user-update-request | Add User to a Consent
[**AnswerConsentChallenge**](AccountInformationServiceAISApi.md#AnswerConsentChallenge) | **Post** /obp/v5.1.0/banks/{BANK_ID}/consents/{CONSENT_ID}/challenge | Answer Consent Challenge
[**CorePrivateAccountsAllBanks**](AccountInformationServiceAISApi.md#CorePrivateAccountsAllBanks) | **Get** /obp/v5.1.0/my/accounts | Get Accounts at all Banks (private)
[**CreateConsentByConsentRequestIdEmail**](AccountInformationServiceAISApi.md#CreateConsentByConsentRequestIdEmail) | **Post** /obp/v5.1.0/consumer/consent-requests/CONSENT_REQUEST_ID/EMAIL/consents | Create Consent By CONSENT_REQUEST_ID (EMAIL)
[**CreateConsentByConsentRequestIdImplicit**](AccountInformationServiceAISApi.md#CreateConsentByConsentRequestIdImplicit) | **Post** /obp/v5.1.0/consumer/consent-requests/CONSENT_REQUEST_ID/IMPLICIT/consents | Create Consent By CONSENT_REQUEST_ID (IMPLICIT)
[**CreateConsentByConsentRequestIdSms**](AccountInformationServiceAISApi.md#CreateConsentByConsentRequestIdSms) | **Post** /obp/v5.1.0/consumer/consent-requests/CONSENT_REQUEST_ID/SMS/consents | Create Consent By CONSENT_REQUEST_ID (SMS)
[**CreateConsentEmail**](AccountInformationServiceAISApi.md#CreateConsentEmail) | **Post** /obp/v5.1.0/banks/{BANK_ID}/my/consents/EMAIL | Create Consent (EMAIL)
[**CreateConsentImplicit**](AccountInformationServiceAISApi.md#CreateConsentImplicit) | **Post** /obp/v5.1.0/banks/{BANK_ID}/my/consents/IMPLICIT | Create Consent (IMPLICIT)
[**CreateConsentRequest**](AccountInformationServiceAISApi.md#CreateConsentRequest) | **Post** /obp/v5.1.0/consumer/consent-requests | Create Consent Request
[**CreateConsentSms**](AccountInformationServiceAISApi.md#CreateConsentSms) | **Post** /obp/v5.1.0/banks/{BANK_ID}/my/consents/SMS | Create Consent (SMS)
[**GetAccountsHeld**](AccountInformationServiceAISApi.md#GetAccountsHeld) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts-held | Get Accounts Held
[**GetBank**](AccountInformationServiceAISApi.md#GetBank) | **Get** /obp/v5.1.0/banks/{BANK_ID} | Get Bank
[**GetBankAccountBalances**](AccountInformationServiceAISApi.md#GetBankAccountBalances) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/balances | Get Account Balances
[**GetBankAccountsBalances**](AccountInformationServiceAISApi.md#GetBankAccountsBalances) | **Get** /obp/v5.1.0/banks/{BANK_ID}/balances | Get Accounts Balances
[**GetBanks**](AccountInformationServiceAISApi.md#GetBanks) | **Get** /obp/v5.1.0/banks | Get Banks
[**GetConsentByConsentId**](AccountInformationServiceAISApi.md#GetConsentByConsentId) | **Get** /obp/v5.1.0/consumer/consents/{CONSENT_ID} | Get Consent By Consent Id
[**GetConsentByConsentRequestId**](AccountInformationServiceAISApi.md#GetConsentByConsentRequestId) | **Get** /obp/v5.1.0/consumer/consent-requests/CONSENT_REQUEST_ID/consents | Get Consent By Consent Request Id
[**GetConsentInfos**](AccountInformationServiceAISApi.md#GetConsentInfos) | **Get** /obp/v5.1.0/banks/{BANK_ID}/my/consent-infos | Get Consents Info
[**GetConsentRequest**](AccountInformationServiceAISApi.md#GetConsentRequest) | **Get** /obp/v5.1.0/consumer/consent-requests/CONSENT_REQUEST_ID | Get Consent Request
[**GetConsents**](AccountInformationServiceAISApi.md#GetConsents) | **Get** /obp/v5.1.0/banks/{BANK_ID}/my/consents | Get Consents
[**GetCoreAccountById**](AccountInformationServiceAISApi.md#GetCoreAccountById) | **Get** /obp/v5.1.0/my/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/account | Get Account by Id (Core)
[**GetCoreTransactionsForBankAccount**](AccountInformationServiceAISApi.md#GetCoreTransactionsForBankAccount) | **Get** /obp/v5.1.0/my/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/transactions | Get Transactions for Account (Core)
[**GetPrivateAccountIdsbyBankId**](AccountInformationServiceAISApi.md#GetPrivateAccountIdsbyBankId) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/account_ids/private | Get Accounts at Bank (IDs only)
[**GetServerJWK**](AccountInformationServiceAISApi.md#GetServerJWK) | **Get** /obp/v5.1.0/certs | Get JSON Web Key (JWK)
[**GetTransactionTypes**](AccountInformationServiceAISApi.md#GetTransactionTypes) | **Get** /obp/v5.1.0/banks/{BANK_ID}/transaction-types | Get Transaction Types at Bank
[**MtlsClientCertificateInfo**](AccountInformationServiceAISApi.md#MtlsClientCertificateInfo) | **Get** /obp/v5.1.0/my/mtls/certificate/current | Provide client&#39;s certificate info of a current call
[**PrivateAccountsAtOneBank**](AccountInformationServiceAISApi.md#PrivateAccountsAtOneBank) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/private | Get Accounts at Bank (Minimal)
[**RevokeConsent**](AccountInformationServiceAISApi.md#RevokeConsent) | **Get** /obp/v5.1.0/banks/{BANK_ID}/my/consents/{CONSENT_ID}/revoke | Revoke Consent
[**RevokeConsentAtBank**](AccountInformationServiceAISApi.md#RevokeConsentAtBank) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/consents/{CONSENT_ID} | Revoke Consent at Bank
[**SelfRevokeConsent**](AccountInformationServiceAISApi.md#SelfRevokeConsent) | **Delete** /obp/v5.1.0/my/consent/current | Revoke Consent used in the Current Call
[**UpdateConsentStatus**](AccountInformationServiceAISApi.md#UpdateConsentStatus) | **Put** /obp/v5.1.0/banks/{BANK_ID}/consents/{CONSENT_ID} | Update Consent Status


# **AddConsentUser**
> ConsentChallengeJsonV310 AddConsentUser(ctx, body, cONSENTID, bANKID)
Add User to a Consent

<p>This endpoint is used to add the User of Consent.</p><p>Each Consent has one of the following states: INITIATED, ACCEPTED, REJECTED, REVOKED, RECEIVED, VALID, REVOKEDBYPSU, EXPIRED, TERMINATEDBYTPP, AUTHORISED, AWAITINGAUTHORISATION.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PutConsentUserJsonV400**](PutConsentUserJsonV400.md)| PutConsentUserJsonV400 object that needs to be added. | 
  **cONSENTID** | **string**| the consent id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**ConsentChallengeJsonV310**](ConsentChallengeJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AnswerConsentChallenge**
> ConsentChallengeJsonV310 AnswerConsentChallenge(ctx, body, cONSENTID, bANKID)
Answer Consent Challenge

<p>An OBP Consent allows the holder of the Consent to call one or more endpoints.</p><p>Consents must be created and authorisied using SCA (Strong Customer Authentication).</p><p>That is, Consents can be created by an authorised User via the OBP REST API but they must be confirmed via an out of band (OOB) mechanism such as a code sent to a mobile phone.</p><p>Each Consent has one of the following states: INITIATED, ACCEPTED, REJECTED, REVOKED, RECEIVED, VALID, REVOKEDBYPSU, EXPIRED, TERMINATEDBYTPP, AUTHORISED, AWAITINGAUTHORISATION.</p><p>Each Consent is bound to a consumer i.e. you need to identify yourself over request header value Consumer-Key.<br />For example:<br />GET /obp/v4.0.0/users/current HTTP/1.1<br />Host: 127.0.0.1:8080<br />Consent-JWT: eyJhbGciOiJIUzI1NiJ9.eyJlbnRpdGxlbWVudHMiOlt7InJvbGVfbmFtZSI6IkNhbkdldEFueVVzZXIiLCJiYW5rX2lkIjoiIn<br />1dLCJjcmVhdGVkQnlVc2VySWQiOiJhYjY1MzlhOS1iMTA1LTQ0ODktYTg4My0wYWQ4ZDZjNjE2NTciLCJzdWIiOiIzNDc1MDEzZi03YmY5LTQyNj<br />EtOWUxYy0xZTdlNWZjZTJlN2UiLCJhdWQiOiI4MTVhMGVmMS00YjZhLTQyMDUtYjExMi1lNDVmZDZmNGQzYWQiLCJuYmYiOjE1ODA3NDE2NjcsIml<br />zcyI6Imh0dHA6XC9cLzEyNy4wLjAuMTo4MDgwIiwiZXhwIjoxNTgwNzQ1MjY3LCJpYXQiOjE1ODA3NDE2NjcsImp0aSI6ImJkYzVjZTk5LTE2ZTY<br />tNDM4Yi1hNjllLTU3MTAzN2RhMTg3OCIsInZpZXdzIjpbXX0.L3fEEEhdCVr3qnmyRKBBUaIQ7dk1VjiFaEBW8hUNjfg</p><p>Consumer-Key: ejznk505d132ryomnhbx1qmtohurbsbb0kijajsk<br />cache-control: no-cache</p><p>Maximum time to live of the token is specified over props value consents.max_time_to_live. In case isn't defined default value is 3600 seconds.</p><p>Example of POST JSON:<br />{<br />&quot;everything&quot;: false,<br />&quot;views&quot;: [<br />{<br />&quot;bank_id&quot;: &quot;GENODEM1GLS&quot;,<br />&quot;account_id&quot;: &quot;8ca8a7e4-6d02-40e3-a129-0b2bf89de9f0&quot;,<br />&quot;view_id&quot;: &quot;owner&quot;<br />}<br />],<br />&quot;entitlements&quot;: [<br />{<br />&quot;bank_id&quot;: &quot;GENODEM1GLS&quot;,<br />&quot;role_name&quot;: &quot;CanGetCustomer&quot;<br />}<br />],<br />&quot;consumer_id&quot;: &quot;7uy8a7e4-6d02-40e3-a129-0b2bf89de8uh&quot;,<br />&quot;email&quot;: &quot;<a href=\"&#109;a&#105;&#x6c;&#x74;&#x6f;&#x3a;&#101;&#x76;&#x65;l&#105;&#x6e;&#101;@&#101;&#120;&#97;m&#112;l&#x65;&#46;&#x63;&#x6f;&#109;\">&#101;v&#x65;&#x6c;i&#110;&#x65;&#x40;&#101;&#120;&#97;&#109;&#112;&#x6c;&#x65;&#46;&#99;&#111;&#109;</a>&quot;,<br />&quot;valid_from&quot;: &quot;2020-02-07T08:43:34Z&quot;,<br />&quot;time_to_live&quot;: 3600<br />}<br />Please note that only optional fields are: consumer_id, valid_from and time_to_live.<br />In case you omit they the default values are used:<br />consumer_id = consumer of current user<br />valid_from = current time<br />time_to_live = consents.max_time_to_live</p><p>This endpoint is used to confirm a Consent previously created.</p><p>The User must supply a code that was sent out of band (OOB) for example via an SMS.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostConsentChallengeJsonV310**](PostConsentChallengeJsonV310.md)| PostConsentChallengeJsonV310 object that needs to be added. | 
  **cONSENTID** | **string**| the consent id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**ConsentChallengeJsonV310**](ConsentChallengeJsonV310.md)

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

# **CreateConsentByConsentRequestIdEmail**
> ConsentJsonV500 CreateConsentByConsentRequestIdEmail(ctx, )
Create Consent By CONSENT_REQUEST_ID (EMAIL)

<p>This endpoint continues the process of creating a Consent. It starts the SCA flow which changes the status of the consent from INITIATED to ACCEPTED or REJECTED.<br />Please note that the Consent cannot elevate the privileges logged in user already have.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ConsentJsonV500**](ConsentJsonV500.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateConsentByConsentRequestIdImplicit**
> ConsentJsonV500 CreateConsentByConsentRequestIdImplicit(ctx, )
Create Consent By CONSENT_REQUEST_ID (IMPLICIT)

<p>This endpoint continues the process of creating a Consent. It starts the SCA flow which changes the status of the consent from INITIATED to ACCEPTED or REJECTED.<br />Please note that the Consent cannot elevate the privileges logged in user already have.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ConsentJsonV500**](ConsentJsonV500.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateConsentByConsentRequestIdSms**
> ConsentJsonV500 CreateConsentByConsentRequestIdSms(ctx, )
Create Consent By CONSENT_REQUEST_ID (SMS)

<p>This endpoint continues the process of creating a Consent. It starts the SCA flow which changes the status of the consent from INITIATED to ACCEPTED or REJECTED.<br />Please note that the Consent cannot elevate the privileges logged in user already have.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ConsentJsonV500**](ConsentJsonV500.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateConsentEmail**
> ConsentJsonV310 CreateConsentEmail(ctx, body, bANKID)
Create Consent (EMAIL)

<p>This endpoint starts the process of creating a Consent.</p><p>The Consent is created in an INITIATED state.</p><p>A One Time Password (OTP) (AKA security challenge) is sent Out of Band (OOB) to the User via the transport defined in SCA_METHOD<br />SCA_METHOD is typically &quot;SMS&quot;,&quot;EMAIL&quot; or &quot;IMPLICIT&quot;. &quot;EMAIL&quot; is used for testing purposes. OBP mapped mode &quot;IMPLICIT&quot; is &quot;EMAIL&quot;.<br />Other mode, bank can decide it in the connector method 'getConsentImplicitSCA'.</p><p>When the Consent is created, OBP (or a backend system) stores the challenge so it can be checked later against the value supplied by the User with the Answer Consent Challenge endpoint.</p><p>An OBP Consent allows the holder of the Consent to call one or more endpoints.</p><p>Consents must be created and authorisied using SCA (Strong Customer Authentication).</p><p>That is, Consents can be created by an authorised User via the OBP REST API but they must be confirmed via an out of band (OOB) mechanism such as a code sent to a mobile phone.</p><p>Each Consent has one of the following states: INITIATED, ACCEPTED, REJECTED, REVOKED, RECEIVED, VALID, REVOKEDBYPSU, EXPIRED, TERMINATEDBYTPP, AUTHORISED, AWAITINGAUTHORISATION.</p><p>Each Consent is bound to a consumer i.e. you need to identify yourself over request header value Consumer-Key.<br />For example:<br />GET /obp/v4.0.0/users/current HTTP/1.1<br />Host: 127.0.0.1:8080<br />Consent-JWT: eyJhbGciOiJIUzI1NiJ9.eyJlbnRpdGxlbWVudHMiOlt7InJvbGVfbmFtZSI6IkNhbkdldEFueVVzZXIiLCJiYW5rX2lkIjoiIn<br />1dLCJjcmVhdGVkQnlVc2VySWQiOiJhYjY1MzlhOS1iMTA1LTQ0ODktYTg4My0wYWQ4ZDZjNjE2NTciLCJzdWIiOiIzNDc1MDEzZi03YmY5LTQyNj<br />EtOWUxYy0xZTdlNWZjZTJlN2UiLCJhdWQiOiI4MTVhMGVmMS00YjZhLTQyMDUtYjExMi1lNDVmZDZmNGQzYWQiLCJuYmYiOjE1ODA3NDE2NjcsIml<br />zcyI6Imh0dHA6XC9cLzEyNy4wLjAuMTo4MDgwIiwiZXhwIjoxNTgwNzQ1MjY3LCJpYXQiOjE1ODA3NDE2NjcsImp0aSI6ImJkYzVjZTk5LTE2ZTY<br />tNDM4Yi1hNjllLTU3MTAzN2RhMTg3OCIsInZpZXdzIjpbXX0.L3fEEEhdCVr3qnmyRKBBUaIQ7dk1VjiFaEBW8hUNjfg</p><p>Consumer-Key: ejznk505d132ryomnhbx1qmtohurbsbb0kijajsk<br />cache-control: no-cache</p><p>Maximum time to live of the token is specified over props value consents.max_time_to_live. In case isn't defined default value is 3600 seconds.</p><p>Example of POST JSON:<br />{<br />&quot;everything&quot;: false,<br />&quot;views&quot;: [<br />{<br />&quot;bank_id&quot;: &quot;GENODEM1GLS&quot;,<br />&quot;account_id&quot;: &quot;8ca8a7e4-6d02-40e3-a129-0b2bf89de9f0&quot;,<br />&quot;view_id&quot;: &quot;owner&quot;<br />}<br />],<br />&quot;entitlements&quot;: [<br />{<br />&quot;bank_id&quot;: &quot;GENODEM1GLS&quot;,<br />&quot;role_name&quot;: &quot;CanGetCustomer&quot;<br />}<br />],<br />&quot;consumer_id&quot;: &quot;7uy8a7e4-6d02-40e3-a129-0b2bf89de8uh&quot;,<br />&quot;email&quot;: &quot;<a href=\"&#x6d;&#97;&#105;&#108;&#x74;&#111;&#58;&#x65;&#118;&#x65;&#x6c;&#x69;&#x6e;&#101;&#x40;&#101;&#120;a&#109;&#x70;&#x6c;e&#x2e;&#99;&#111;&#109;\">&#x65;&#118;e&#x6c;&#105;&#110;&#x65;&#64;&#x65;&#120;&#97;&#109;&#x70;&#x6c;&#x65;&#46;&#99;&#111;&#x6d;</a>&quot;,<br />&quot;valid_from&quot;: &quot;2020-02-07T08:43:34Z&quot;,<br />&quot;time_to_live&quot;: 3600<br />}<br />Please note that only optional fields are: consumer_id, valid_from and time_to_live.<br />In case you omit they the default values are used:<br />consumer_id = consumer of current user<br />valid_from = current time<br />time_to_live = consents.max_time_to_live</p><p>Authentication is Mandatory</p><p>Example 1:<br />{<br />&quot;everything&quot;: true,<br />&quot;views&quot;: [],<br />&quot;entitlements&quot;: [],<br />&quot;consumer_id&quot;: &quot;7uy8a7e4-6d02-40e3-a129-0b2bf89de8uh&quot;,<br />&quot;phone_number&quot;: &quot;+49 170 1234567&quot;<br />}</p><p>Please note that consumer_id is optional field<br />Example 2:<br />{<br />&quot;everything&quot;: true,<br />&quot;views&quot;: [],<br />&quot;entitlements&quot;: [],<br />&quot;phone_number&quot;: &quot;+49 170 1234567&quot;<br />}</p><p>Please note if everything=false you need to explicitly specify views and entitlements<br />Example 3:<br />{<br />&quot;everything&quot;: false,<br />&quot;views&quot;: [<br />{<br />&quot;bank_id&quot;: &quot;GENODEM1GLS&quot;,<br />&quot;account_id&quot;: &quot;8ca8a7e4-6d02-40e3-a129-0b2bf89de9f0&quot;,<br />&quot;view_id&quot;: &quot;owner&quot;<br />}<br />],<br />&quot;entitlements&quot;: [<br />{<br />&quot;bank_id&quot;: &quot;GENODEM1GLS&quot;,<br />&quot;role_name&quot;: &quot;CanGetCustomer&quot;<br />}<br />],<br />&quot;consumer_id&quot;: &quot;7uy8a7e4-6d02-40e3-a129-0b2bf89de8uh&quot;,<br />&quot;phone_number&quot;: &quot;+49 170 1234567&quot;<br />}</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostConsentEmailJsonV310**](PostConsentEmailJsonV310.md)| PostConsentEmailJsonV310 object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**ConsentJsonV310**](ConsentJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateConsentImplicit**
> ConsentJsonV310 CreateConsentImplicit(ctx, body, bANKID)
Create Consent (IMPLICIT)

<p>This endpoint starts the process of creating a Consent.</p><p>The Consent is created in an INITIATED state.</p><p>A One Time Password (OTP) (AKA security challenge) is sent Out of Band (OOB) to the User via the transport defined in SCA_METHOD<br />SCA_METHOD is typically &quot;SMS&quot;,&quot;EMAIL&quot; or &quot;IMPLICIT&quot;. &quot;EMAIL&quot; is used for testing purposes. OBP mapped mode &quot;IMPLICIT&quot; is &quot;EMAIL&quot;.<br />Other mode, bank can decide it in the connector method 'getConsentImplicitSCA'.</p><p>When the Consent is created, OBP (or a backend system) stores the challenge so it can be checked later against the value supplied by the User with the Answer Consent Challenge endpoint.</p><p>An OBP Consent allows the holder of the Consent to call one or more endpoints.</p><p>Consents must be created and authorisied using SCA (Strong Customer Authentication).</p><p>That is, Consents can be created by an authorised User via the OBP REST API but they must be confirmed via an out of band (OOB) mechanism such as a code sent to a mobile phone.</p><p>Each Consent has one of the following states: INITIATED, ACCEPTED, REJECTED, REVOKED, RECEIVED, VALID, REVOKEDBYPSU, EXPIRED, TERMINATEDBYTPP, AUTHORISED, AWAITINGAUTHORISATION.</p><p>Each Consent is bound to a consumer i.e. you need to identify yourself over request header value Consumer-Key.<br />For example:<br />GET /obp/v4.0.0/users/current HTTP/1.1<br />Host: 127.0.0.1:8080<br />Consent-JWT: eyJhbGciOiJIUzI1NiJ9.eyJlbnRpdGxlbWVudHMiOlt7InJvbGVfbmFtZSI6IkNhbkdldEFueVVzZXIiLCJiYW5rX2lkIjoiIn<br />1dLCJjcmVhdGVkQnlVc2VySWQiOiJhYjY1MzlhOS1iMTA1LTQ0ODktYTg4My0wYWQ4ZDZjNjE2NTciLCJzdWIiOiIzNDc1MDEzZi03YmY5LTQyNj<br />EtOWUxYy0xZTdlNWZjZTJlN2UiLCJhdWQiOiI4MTVhMGVmMS00YjZhLTQyMDUtYjExMi1lNDVmZDZmNGQzYWQiLCJuYmYiOjE1ODA3NDE2NjcsIml<br />zcyI6Imh0dHA6XC9cLzEyNy4wLjAuMTo4MDgwIiwiZXhwIjoxNTgwNzQ1MjY3LCJpYXQiOjE1ODA3NDE2NjcsImp0aSI6ImJkYzVjZTk5LTE2ZTY<br />tNDM4Yi1hNjllLTU3MTAzN2RhMTg3OCIsInZpZXdzIjpbXX0.L3fEEEhdCVr3qnmyRKBBUaIQ7dk1VjiFaEBW8hUNjfg</p><p>Consumer-Key: ejznk505d132ryomnhbx1qmtohurbsbb0kijajsk<br />cache-control: no-cache</p><p>Maximum time to live of the token is specified over props value consents.max_time_to_live. In case isn't defined default value is 3600 seconds.</p><p>Example of POST JSON:<br />{<br />&quot;everything&quot;: false,<br />&quot;views&quot;: [<br />{<br />&quot;bank_id&quot;: &quot;GENODEM1GLS&quot;,<br />&quot;account_id&quot;: &quot;8ca8a7e4-6d02-40e3-a129-0b2bf89de9f0&quot;,<br />&quot;view_id&quot;: &quot;owner&quot;<br />}<br />],<br />&quot;entitlements&quot;: [<br />{<br />&quot;bank_id&quot;: &quot;GENODEM1GLS&quot;,<br />&quot;role_name&quot;: &quot;CanGetCustomer&quot;<br />}<br />],<br />&quot;consumer_id&quot;: &quot;7uy8a7e4-6d02-40e3-a129-0b2bf89de8uh&quot;,<br />&quot;email&quot;: &quot;<a href=\"&#x6d;&#97;&#105;l&#x74;o&#58;&#101;v&#x65;&#x6c;&#x69;&#110;&#101;@&#101;&#120;&#x61;m&#112;&#108;&#101;&#x2e;&#99;o&#109;\">eve&#108;&#x69;&#x6e;&#x65;@&#x65;xa&#109;&#112;&#108;&#101;.&#99;&#111;&#109;</a>&quot;,<br />&quot;valid_from&quot;: &quot;2020-02-07T08:43:34Z&quot;,<br />&quot;time_to_live&quot;: 3600<br />}<br />Please note that only optional fields are: consumer_id, valid_from and time_to_live.<br />In case you omit they the default values are used:<br />consumer_id = consumer of current user<br />valid_from = current time<br />time_to_live = consents.max_time_to_live</p><p>Authentication is Mandatory</p><p>Example 1:<br />{<br />&quot;everything&quot;: true,<br />&quot;views&quot;: [],<br />&quot;entitlements&quot;: [],<br />&quot;consumer_id&quot;: &quot;7uy8a7e4-6d02-40e3-a129-0b2bf89de8uh&quot;,<br />}</p><p>Please note that consumer_id is optional field<br />Example 2:<br />{<br />&quot;everything&quot;: true,<br />&quot;views&quot;: [],<br />&quot;entitlements&quot;: [],<br />}</p><p>Please note if everything=false you need to explicitly specify views and entitlements<br />Example 3:<br />{<br />&quot;everything&quot;: false,<br />&quot;views&quot;: [<br />{<br />&quot;bank_id&quot;: &quot;GENODEM1GLS&quot;,<br />&quot;account_id&quot;: &quot;8ca8a7e4-6d02-40e3-a129-0b2bf89de9f0&quot;,<br />&quot;view_id&quot;: &quot;owner&quot;<br />}<br />],<br />&quot;entitlements&quot;: [<br />{<br />&quot;bank_id&quot;: &quot;GENODEM1GLS&quot;,<br />&quot;role_name&quot;: &quot;CanGetCustomer&quot;<br />}<br />],<br />&quot;consumer_id&quot;: &quot;7uy8a7e4-6d02-40e3-a129-0b2bf89de8uh&quot;,<br />}</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostConsentImplicitJsonV310**](PostConsentImplicitJsonV310.md)| PostConsentImplicitJsonV310 object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**ConsentJsonV310**](ConsentJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateConsentRequest**
> ConsentRequestResponseJson CreateConsentRequest(ctx, body)
Create Consent Request

<p>Client Authentication (mandatory)</p><p>It is used when applications request an access token to access their own resources, not on behalf of a user.</p><p>The client needs to authenticate themselves for this request.<br />In case of public client we use client_id and private kew to obtain access token, otherwise we use client_id and client_secret.<br />The obtained access token is used in the HTTP Bearer auth header of our request.</p><p>Example:<br />Authorization: Bearer eXtneO-THbQtn3zvK_kQtXXfvOZyZFdBCItlPDbR2Bk.dOWqtXCtFX-tqGTVR0YrIjvAolPIVg7GZ-jz83y6nA0</p><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostConsentRequestJsonV500**](PostConsentRequestJsonV500.md)| PostConsentRequestJsonV500 object that needs to be added. | 

### Return type

[**ConsentRequestResponseJson**](ConsentRequestResponseJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateConsentSms**
> ConsentJsonV310 CreateConsentSms(ctx, body, bANKID)
Create Consent (SMS)

<p>This endpoint starts the process of creating a Consent.</p><p>The Consent is created in an INITIATED state.</p><p>A One Time Password (OTP) (AKA security challenge) is sent Out of Band (OOB) to the User via the transport defined in SCA_METHOD<br />SCA_METHOD is typically &quot;SMS&quot;,&quot;EMAIL&quot; or &quot;IMPLICIT&quot;. &quot;EMAIL&quot; is used for testing purposes. OBP mapped mode &quot;IMPLICIT&quot; is &quot;EMAIL&quot;.<br />Other mode, bank can decide it in the connector method 'getConsentImplicitSCA'.</p><p>When the Consent is created, OBP (or a backend system) stores the challenge so it can be checked later against the value supplied by the User with the Answer Consent Challenge endpoint.</p><p>An OBP Consent allows the holder of the Consent to call one or more endpoints.</p><p>Consents must be created and authorisied using SCA (Strong Customer Authentication).</p><p>That is, Consents can be created by an authorised User via the OBP REST API but they must be confirmed via an out of band (OOB) mechanism such as a code sent to a mobile phone.</p><p>Each Consent has one of the following states: INITIATED, ACCEPTED, REJECTED, REVOKED, RECEIVED, VALID, REVOKEDBYPSU, EXPIRED, TERMINATEDBYTPP, AUTHORISED, AWAITINGAUTHORISATION.</p><p>Each Consent is bound to a consumer i.e. you need to identify yourself over request header value Consumer-Key.<br />For example:<br />GET /obp/v4.0.0/users/current HTTP/1.1<br />Host: 127.0.0.1:8080<br />Consent-JWT: eyJhbGciOiJIUzI1NiJ9.eyJlbnRpdGxlbWVudHMiOlt7InJvbGVfbmFtZSI6IkNhbkdldEFueVVzZXIiLCJiYW5rX2lkIjoiIn<br />1dLCJjcmVhdGVkQnlVc2VySWQiOiJhYjY1MzlhOS1iMTA1LTQ0ODktYTg4My0wYWQ4ZDZjNjE2NTciLCJzdWIiOiIzNDc1MDEzZi03YmY5LTQyNj<br />EtOWUxYy0xZTdlNWZjZTJlN2UiLCJhdWQiOiI4MTVhMGVmMS00YjZhLTQyMDUtYjExMi1lNDVmZDZmNGQzYWQiLCJuYmYiOjE1ODA3NDE2NjcsIml<br />zcyI6Imh0dHA6XC9cLzEyNy4wLjAuMTo4MDgwIiwiZXhwIjoxNTgwNzQ1MjY3LCJpYXQiOjE1ODA3NDE2NjcsImp0aSI6ImJkYzVjZTk5LTE2ZTY<br />tNDM4Yi1hNjllLTU3MTAzN2RhMTg3OCIsInZpZXdzIjpbXX0.L3fEEEhdCVr3qnmyRKBBUaIQ7dk1VjiFaEBW8hUNjfg</p><p>Consumer-Key: ejznk505d132ryomnhbx1qmtohurbsbb0kijajsk<br />cache-control: no-cache</p><p>Maximum time to live of the token is specified over props value consents.max_time_to_live. In case isn't defined default value is 3600 seconds.</p><p>Example of POST JSON:<br />{<br />&quot;everything&quot;: false,<br />&quot;views&quot;: [<br />{<br />&quot;bank_id&quot;: &quot;GENODEM1GLS&quot;,<br />&quot;account_id&quot;: &quot;8ca8a7e4-6d02-40e3-a129-0b2bf89de9f0&quot;,<br />&quot;view_id&quot;: &quot;owner&quot;<br />}<br />],<br />&quot;entitlements&quot;: [<br />{<br />&quot;bank_id&quot;: &quot;GENODEM1GLS&quot;,<br />&quot;role_name&quot;: &quot;CanGetCustomer&quot;<br />}<br />],<br />&quot;consumer_id&quot;: &quot;7uy8a7e4-6d02-40e3-a129-0b2bf89de8uh&quot;,<br />&quot;email&quot;: &quot;<a href=\"&#x6d;&#x61;&#x69;&#x6c;&#116;&#111;:&#x65;&#118;e&#x6c;&#x69;n&#101;&#64;&#101;x&#97;&#x6d;p&#x6c;&#x65;.co&#x6d;\">&#x65;&#118;e&#108;&#x69;&#110;&#101;&#x40;e&#120;&#97;&#x6d;&#x70;l&#101;&#x2e;&#99;&#x6f;m</a>&quot;,<br />&quot;valid_from&quot;: &quot;2020-02-07T08:43:34Z&quot;,<br />&quot;time_to_live&quot;: 3600<br />}<br />Please note that only optional fields are: consumer_id, valid_from and time_to_live.<br />In case you omit they the default values are used:<br />consumer_id = consumer of current user<br />valid_from = current time<br />time_to_live = consents.max_time_to_live</p><p>Authentication is Mandatory</p><p>Example 1:<br />{<br />&quot;everything&quot;: true,<br />&quot;views&quot;: [],<br />&quot;entitlements&quot;: [],<br />&quot;consumer_id&quot;: &quot;7uy8a7e4-6d02-40e3-a129-0b2bf89de8uh&quot;,<br />&quot;email&quot;: &quot;<a href=\"&#109;ai&#108;&#x74;&#x6f;:&#x65;&#x76;&#101;&#x6c;i&#110;e&#x40;&#101;xa&#109;p&#x6c;&#101;.&#x63;&#111;&#x6d;\">&#101;veli&#x6e;&#101;&#64;e&#120;a&#109;&#112;&#x6c;&#101;&#46;&#x63;&#x6f;&#x6d;</a>&quot;<br />}</p><p>Please note that consumer_id is optional field<br />Example 2:<br />{<br />&quot;everything&quot;: true,<br />&quot;views&quot;: [],<br />&quot;entitlements&quot;: [],<br />&quot;email&quot;: &quot;<a href=\"&#109;&#97;&#105;&#108;&#x74;&#111;&#58;&#x65;v&#x65;&#108;i&#x6e;&#x65;&#64;&#x65;&#x78;&#97;&#109;&#x70;l&#101;&#46;&#x63;&#x6f;&#x6d;\">&#x65;&#118;&#x65;&#108;i&#x6e;&#x65;&#x40;&#101;xa&#x6d;p&#108;&#101;.c&#111;&#x6d;</a>&quot;<br />}</p><p>Please note if everything=false you need to explicitly specify views and entitlements<br />Example 3:<br />{<br />&quot;everything&quot;: false,<br />&quot;views&quot;: [<br />{<br />&quot;bank_id&quot;: &quot;GENODEM1GLS&quot;,<br />&quot;account_id&quot;: &quot;8ca8a7e4-6d02-40e3-a129-0b2bf89de9f0&quot;,<br />&quot;view_id&quot;: &quot;owner&quot;<br />}<br />],<br />&quot;entitlements&quot;: [<br />{<br />&quot;bank_id&quot;: &quot;GENODEM1GLS&quot;,<br />&quot;role_name&quot;: &quot;CanGetCustomer&quot;<br />}<br />],<br />&quot;consumer_id&quot;: &quot;7uy8a7e4-6d02-40e3-a129-0b2bf89de8uh&quot;,<br />&quot;email&quot;: &quot;<a href=\"&#x6d;&#x61;&#x69;&#108;&#116;&#x6f;&#x3a;&#x65;&#x76;&#x65;l&#105;&#x6e;&#x65;&#64;&#101;&#x78;&#97;mp&#108;&#x65;.&#99;&#x6f;&#109;\">&#x65;&#x76;&#x65;&#108;i&#110;&#x65;&#x40;&#101;&#x78;am&#112;l&#x65;&#46;&#x63;&#x6f;&#x6d;</a>&quot;<br />}</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostConsentPhoneJsonV310**](PostConsentPhoneJsonV310.md)| PostConsentPhoneJsonV310 object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**ConsentJsonV310**](ConsentJsonV310.md)

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

# **GetConsentByConsentId**
> ConsentJsonV500 GetConsentByConsentId(ctx, cONSENTID)
Get Consent By Consent Id

<p>This endpoint gets the Consent By consent id.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cONSENTID** | **string**| the consent id | 

### Return type

[**ConsentJsonV500**](ConsentJsonV500.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConsentByConsentRequestId**
> ConsentJsonV500 GetConsentByConsentRequestId(ctx, )
Get Consent By Consent Request Id

<p>This endpoint gets the Consent By consent request id.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ConsentJsonV500**](ConsentJsonV500.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConsentInfos**
> ConsentInfosJsonV400 GetConsentInfos(ctx, bANKID)
Get Consents Info

<p>This endpoint gets the Consents that the current User created.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**ConsentInfosJsonV400**](ConsentInfosJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConsentRequest**
> ConsentRequestResponseJson GetConsentRequest(ctx, )
Get Consent Request

<p>Authentication is Optional</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ConsentRequestResponseJson**](ConsentRequestResponseJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConsents**
> ConsentsJsonV400 GetConsents(ctx, bANKID)
Get Consents

<p>This endpoint gets the Consents that the current User created.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**ConsentsJsonV400**](ConsentsJsonV400.md)

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

# **GetServerJWK**
> SeverJwk GetServerJWK(ctx, )
Get JSON Web Key (JWK)

<p>Get the server's public JSON Web Key (JWK) set and certificate chain.<br />It is required by client applications to validate ID tokens, self-contained access tokens and other issued objects.</p><p>Authentication is Optional</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SeverJwk**](SeverJWK.md)

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

# **MtlsClientCertificateInfo**
> CertificateInfoJsonV510 MtlsClientCertificateInfo(ctx, )
Provide client's certificate info of a current call

<p>Provide client's certificate info of a current call specified by PSD2-CERT value at Request Header</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CertificateInfoJsonV510**](CertificateInfoJsonV510.md)

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

# **RevokeConsent**
> ConsentJsonV310 RevokeConsent(ctx, cONSENTID, bANKID)
Revoke Consent

<p>Revoke Consent for current user specified by CONSENT_ID</p><p>There are a few reasons you might need to revoke an application’s access to a user’s account:<br />- The user explicitly wishes to revoke the application’s access<br />- You as the service provider have determined an application is compromised or malicious, and want to disable it<br />- etc.</p><p>Please note that this endpoint only supports the case:: &quot;The user explicitly wishes to revoke the application’s access&quot;</p><p>OBP as a resource server stores access tokens in a database, then it is relatively easy to revoke some token that belongs to a particular user.<br />The status of the token is changed to &quot;REVOKED&quot; so the next time the revoked client makes a request, their token will fail to validate.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cONSENTID** | **string**| the consent id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**ConsentJsonV310**](ConsentJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeConsentAtBank**
> ConsentJsonV310 RevokeConsentAtBank(ctx, cONSENTID, bANKID)
Revoke Consent at Bank

<p>Revoke Consent specified by CONSENT_ID</p><p>There are a few reasons you might need to revoke an application’s access to a user’s account:<br />- The user explicitly wishes to revoke the application’s access<br />- You as the service provider have determined an application is compromised or malicious, and want to disable it<br />- etc.</p><p>OBP as a resource server stores access tokens in a database, then it is relatively easy to revoke some token that belongs to a particular user.<br />The status of the token is changed to &quot;REVOKED&quot; so the next time the revoked client makes a request, their token will fail to validate.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cONSENTID** | **string**| the consent id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**ConsentJsonV310**](ConsentJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SelfRevokeConsent**
> ConsentJsonV310 SelfRevokeConsent(ctx, )
Revoke Consent used in the Current Call

<p>Revoke Consent specified by Consent-Id at Request Header</p><p>There are a few reasons you might need to revoke an application’s access to a user’s account:<br />- The user explicitly wishes to revoke the application’s access<br />- You as the service provider have determined an application is compromised or malicious, and want to disable it<br />- etc.</p><p>OBP as a resource server stores access tokens in a database, then it is relatively easy to revoke some token that belongs to a particular user.<br />The status of the token is changed to &quot;REVOKED&quot; so the next time the revoked client makes a request, their token will fail to validate.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ConsentJsonV310**](ConsentJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConsentStatus**
> ConsentChallengeJsonV310 UpdateConsentStatus(ctx, body, cONSENTID, bANKID)
Update Consent Status

<p>This endpoint is used to update the Status of Consent.</p><p>Each Consent has one of the following states: INITIATED, ACCEPTED, REJECTED, REVOKED, RECEIVED, VALID, REVOKEDBYPSU, EXPIRED, TERMINATEDBYTPP, AUTHORISED, AWAITINGAUTHORISATION.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PutConsentStatusJsonV400**](PutConsentStatusJsonV400.md)| PutConsentStatusJsonV400 object that needs to be added. | 
  **cONSENTID** | **string**| the consent id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**ConsentChallengeJsonV310**](ConsentChallengeJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

