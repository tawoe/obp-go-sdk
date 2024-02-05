# \UserApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddEntitlement**](UserApi.md#AddEntitlement) | **Post** /obp/v5.1.0/users/{USER_ID}/entitlements | Add Entitlement for a User
[**AddEntitlementRequest**](UserApi.md#AddEntitlementRequest) | **Post** /obp/v5.1.0/entitlement-requests | Create Entitlement Request for current User
[**AnswerUserAuthContextUpdateChallenge**](UserApi.md#AnswerUserAuthContextUpdateChallenge) | **Post** /obp/v5.1.0/banks/{BANK_ID}/users/current/auth-context-updates/{AUTH_CONTEXT_UPDATE_ID}/challenge | Answer User Auth Context Update Challenge
[**CreateMyPersonalUserAttribute**](UserApi.md#CreateMyPersonalUserAttribute) | **Post** /obp/v5.1.0/my/user/attributes | Create My Personal User Attribute
[**CreateNonPersonalUserAttribute**](UserApi.md#CreateNonPersonalUserAttribute) | **Post** /obp/v5.1.0/users/{USER_ID}/non-personal/attributes | Create Non Personal User Attribute
[**CreateUser**](UserApi.md#CreateUser) | **Post** /obp/v5.1.0/users | Create User
[**CreateUserAuthContext**](UserApi.md#CreateUserAuthContext) | **Post** /obp/v5.1.0/users/{USER_ID}/auth-context | Create User Auth Context
[**CreateUserAuthContextUpdateRequest**](UserApi.md#CreateUserAuthContextUpdateRequest) | **Post** /obp/v5.1.0/banks/{BANK_ID}/users/current/auth-context-updates/{SCA_METHOD} | Create User Auth Context Update Request
[**CreateUserCustomerLinks**](UserApi.md#CreateUserCustomerLinks) | **Post** /obp/v5.1.0/banks/{BANK_ID}/user_customer_links | Create User Customer Link
[**CreateUserWithAccountAccess**](UserApi.md#CreateUserWithAccountAccess) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/user-account-access | Create (DAuth) User with Account Access
[**CreateUserWithRoles**](UserApi.md#CreateUserWithRoles) | **Post** /obp/v5.1.0/user-entitlements | Create (DAuth) User with Roles
[**DeleteEntitlement**](UserApi.md#DeleteEntitlement) | **Delete** /obp/v5.1.0/users/{USER_ID}/entitlement/{ENTITLEMENT_ID} | Delete Entitlement
[**DeleteEntitlementRequest**](UserApi.md#DeleteEntitlementRequest) | **Delete** /obp/v5.1.0/entitlement-requests/{ENTITLEMENT_REQUEST_ID} | Delete Entitlement Request
[**DeleteNonPersonalUserAttribute**](UserApi.md#DeleteNonPersonalUserAttribute) | **Delete** /obp/v5.1.0/users/{USER_ID}/non-personal/attributes/USER_ATTRIBUTE_ID | Delete Non Personal User Attribute
[**DeleteUser**](UserApi.md#DeleteUser) | **Delete** /obp/v5.1.0/users/{USER_ID} | Delete a User
[**DeleteUserAuthContextById**](UserApi.md#DeleteUserAuthContextById) | **Delete** /obp/v5.1.0/users/{USER_ID}/auth-context/{USER_AUTH_CONTEXT_ID} | Delete User Auth Context
[**DeleteUserAuthContexts**](UserApi.md#DeleteUserAuthContexts) | **Delete** /obp/v5.1.0/users/{USER_ID}/auth-context | Delete User&#39;s Auth Contexts
[**GetAllEntitlementRequests**](UserApi.md#GetAllEntitlementRequests) | **Get** /obp/v5.1.0/entitlement-requests | Get all Entitlement Requests
[**GetCurrentUser**](UserApi.md#GetCurrentUser) | **Get** /obp/v5.1.0/users/current | Get User (Current)
[**GetCurrentUserId**](UserApi.md#GetCurrentUserId) | **Get** /obp/v5.1.0/users/current/user_id | Get User Id (Current)
[**GetCustomersAtAnyBank**](UserApi.md#GetCustomersAtAnyBank) | **Get** /obp/v5.1.0/customers | Get Customers at Any Bank
[**GetCustomersAtOneBank**](UserApi.md#GetCustomersAtOneBank) | **Get** /obp/v5.1.0/banks/{BANK_ID}/customers | Get Customers at Bank
[**GetCustomersForUser**](UserApi.md#GetCustomersForUser) | **Get** /obp/v5.1.0/users/current/customers | Get Customers for Current User
[**GetCustomersForUserIdsOnly**](UserApi.md#GetCustomersForUserIdsOnly) | **Get** /obp/v5.1.0/users/current/customers/customer_ids | Get Customers for Current User (IDs only)
[**GetCustomersMinimalAtAnyBank**](UserApi.md#GetCustomersMinimalAtAnyBank) | **Get** /obp/v5.1.0/customers-minimal | Get Customers Minimal at Any Bank
[**GetCustomersMinimalAtOneBank**](UserApi.md#GetCustomersMinimalAtOneBank) | **Get** /obp/v5.1.0/banks/{BANK_ID}/customers-minimal | Get Customers Minimal at Bank
[**GetEntitlementRequests**](UserApi.md#GetEntitlementRequests) | **Get** /obp/v5.1.0/users/{USER_ID}/entitlement-requests | Get Entitlement Requests for a User
[**GetEntitlementRequestsForCurrentUser**](UserApi.md#GetEntitlementRequestsForCurrentUser) | **Get** /obp/v5.1.0/my/entitlement-requests | Get Entitlement Requests for the current User
[**GetEntitlements**](UserApi.md#GetEntitlements) | **Get** /obp/v5.1.0/users/{USER_ID}/entitlements | Get Entitlements for User
[**GetEntitlementsAndPermissions**](UserApi.md#GetEntitlementsAndPermissions) | **Get** /obp/v5.1.0/users/{USER_ID}/entitlements-and-permissions | Get Entitlements and Permissions for a User
[**GetEntitlementsByBankAndUser**](UserApi.md#GetEntitlementsByBankAndUser) | **Get** /obp/v5.1.0/banks/{BANK_ID}/users/{USER_ID}/entitlements | Get Entitlements for User at Bank
[**GetEntitlementsForBank**](UserApi.md#GetEntitlementsForBank) | **Get** /obp/v5.1.0/banks/{BANK_ID}/entitlements | Get Entitlements for One Bank
[**GetEntitlementsForCurrentUser**](UserApi.md#GetEntitlementsForCurrentUser) | **Get** /obp/v5.1.0/my/entitlements | Get Entitlements for the current User
[**GetLogoutLink**](UserApi.md#GetLogoutLink) | **Get** /obp/v5.1.0/users/current/logout-link | Get Logout Link
[**GetMyCustomersAtAnyBank**](UserApi.md#GetMyCustomersAtAnyBank) | **Get** /obp/v5.1.0/my/customers | Get My Customers
[**GetMyPersonalUserAttributes**](UserApi.md#GetMyPersonalUserAttributes) | **Get** /obp/v5.1.0/my/user/attributes | Get My Personal User Attributes
[**GetMySpaces**](UserApi.md#GetMySpaces) | **Get** /obp/v5.1.0/my/spaces | Get My Spaces
[**GetNonPersonalUserAttributes**](UserApi.md#GetNonPersonalUserAttributes) | **Get** /obp/v5.1.0/users/{USER_ID}/non-personal/attributes | Get Non Personal User Attributes
[**GetPermissionForUserForBankAccount**](UserApi.md#GetPermissionForUserForBankAccount) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/permissions/{PROVIDER}/{PROVIDER_ID} | Get Account access for User
[**GetPermissionsForBankAccount**](UserApi.md#GetPermissionsForBankAccount) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/permissions | Get access
[**GetUserAuthContexts**](UserApi.md#GetUserAuthContexts) | **Get** /obp/v5.1.0/users/{USER_ID}/auth-context | Get User Auth Contexts
[**GetUserByProviderAndUsername**](UserApi.md#GetUserByProviderAndUsername) | **Get** /obp/v5.1.0/users/provider/{PROVIDER}/username/{USERNAME} | Get User by USERNAME
[**GetUserByUserId**](UserApi.md#GetUserByUserId) | **Get** /obp/v5.1.0/users/user_id/{USER_ID} | Get User by USER_ID
[**GetUserLockStatus**](UserApi.md#GetUserLockStatus) | **Get** /obp/v5.1.0/users/{PROVIDER}/{USERNAME}/lock-status | Get User Lock Status
[**GetUserWithAttributes**](UserApi.md#GetUserWithAttributes) | **Get** /obp/v5.1.0/users/{USER_ID}/attributes | Get User with Attributes by USER_ID
[**GetUsers**](UserApi.md#GetUsers) | **Get** /obp/v5.1.0/users | Get all Users
[**GetUsersByEmail**](UserApi.md#GetUsersByEmail) | **Get** /obp/v5.1.0/users/email/EMAIL/terminator | Get Users by Email Address
[**GrantUserAccessToView**](UserApi.md#GrantUserAccessToView) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/account-access/grant | Grant User access to View
[**LockUserByProviderAndUsername**](UserApi.md#LockUserByProviderAndUsername) | **Post** /obp/v5.1.0/users/{PROVIDER}/{USERNAME}/locks | Lock the user
[**RefreshUser**](UserApi.md#RefreshUser) | **Post** /obp/v5.1.0/users/{USER_ID}/refresh | Refresh User
[**ResetPasswordUrl**](UserApi.md#ResetPasswordUrl) | **Post** /obp/v5.1.0/management/user/reset-password-url | Create password reset url
[**RevokeGrantUserAccessToViews**](UserApi.md#RevokeGrantUserAccessToViews) | **Put** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/account-access | Revoke/Grant User access to View
[**RevokeUserAccessToView**](UserApi.md#RevokeUserAccessToView) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/account-access/revoke | Revoke User access to View
[**UnlockUserByProviderAndUsername**](UserApi.md#UnlockUserByProviderAndUsername) | **Put** /obp/v5.1.0/users/{PROVIDER}/{USERNAME}/lock-status | Unlock the user
[**UpdateMyPersonalUserAttribute**](UserApi.md#UpdateMyPersonalUserAttribute) | **Put** /obp/v5.1.0/my/user/attributes/USER_ATTRIBUTE_ID | Update My Personal User Attribute


# **AddEntitlement**
> EntitlementJson AddEntitlement(ctx, body, uSERID)
Add Entitlement for a User

<p>Create Entitlement. Grant Role to User.</p><p>Entitlements are used to grant System or Bank level roles to Users. (For Account level privileges, see Views)</p><p>For a System level Role (.e.g CanGetAnyUser), set bank_id to an empty string i.e. &quot;bank_id&quot;:&quot;&quot;</p><p>For a Bank level Role (e.g. CanCreateAccount), set bank_id to a valid value e.g. &quot;bank_id&quot;:&quot;my-bank-id&quot;</p><p>Authentication is required and the user needs to be a Super Admin. Super Admins are listed in the Props file.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateEntitlementJson**](CreateEntitlementJson.md)| CreateEntitlementJSON object that needs to be added. | 
  **uSERID** | **string**| The user id | 

### Return type

[**EntitlementJson**](EntitlementJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddEntitlementRequest**
> EntitlementRequestJson AddEntitlementRequest(ctx, body)
Create Entitlement Request for current User

<p>Create Entitlement Request.</p><p>Any logged in User can use this endpoint to request an Entitlement</p><p>Entitlements are used to grant System or Bank level roles to Users. (For Account level privileges, see Views)</p><p>For a System level Role (.e.g CanGetAnyUser), set bank_id to an empty string i.e. &quot;bank_id&quot;:&quot;&quot;</p><p>For a Bank level Role (e.g. CanCreateAccount), set bank_id to a valid value e.g. &quot;bank_id&quot;:&quot;my-bank-id&quot;</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateEntitlementJson**](CreateEntitlementJson.md)| CreateEntitlementJSON object that needs to be added. | 

### Return type

[**EntitlementRequestJson**](EntitlementRequestJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AnswerUserAuthContextUpdateChallenge**
> UserAuthContextUpdateJsonV500 AnswerUserAuthContextUpdateChallenge(ctx, body, aUTHCONTEXTUPDATEID, bANKID)
Answer User Auth Context Update Challenge

<p>Answer User Auth Context Update Challenge.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostUserAuthContextUpdateJsonV310**](PostUserAuthContextUpdateJsonV310.md)| PostUserAuthContextUpdateJsonV310 object that needs to be added. | 
  **aUTHCONTEXTUPDATEID** | **string**| the auth context update id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**UserAuthContextUpdateJsonV500**](UserAuthContextUpdateJsonV500.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMyPersonalUserAttribute**
> UserAttributeResponseJsonV400 CreateMyPersonalUserAttribute(ctx, body)
Create My Personal User Attribute

<p>Create My Personal User Attribute</p><p>The <code>type</code> field must be one of &quot;STRING&quot;, &quot;INTEGER&quot;, &quot;DOUBLE&quot; or DATE_WITH_DAY&quot;</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserAttributeJsonV400**](UserAttributeJsonV400.md)| UserAttributeJsonV400 object that needs to be added. | 

### Return type

[**UserAttributeResponseJsonV400**](UserAttributeResponseJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNonPersonalUserAttribute**
> UserAttributeResponseJsonV510 CreateNonPersonalUserAttribute(ctx, body, uSERID)
Create Non Personal User Attribute

<p>Create Non Personal User Attribute</p><p>The type field must be one of &quot;STRING&quot;, &quot;INTEGER&quot;, &quot;DOUBLE&quot; or DATE_WITH_DAY&quot;</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserAttributeJsonV510**](UserAttributeJsonV510.md)| UserAttributeJsonV510 object that needs to be added. | 
  **uSERID** | **string**| The user id | 

### Return type

[**UserAttributeResponseJsonV510**](UserAttributeResponseJsonV510.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUser**
> UserJsonV200 CreateUser(ctx, body)
Create User

<p>Creates OBP user.<br />No authorisation (currently) required.</p><p>Mimics current webform to Register.</p><p>Requires username(email) and password.</p><p>Returns 409 error if username not unique.</p><p>May require validation of email address.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateUserJson**](CreateUserJson.md)| CreateUserJson object that needs to be added. | 

### Return type

[**UserJsonV200**](UserJsonV200.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUserAuthContext**
> UserAuthContextJsonV500 CreateUserAuthContext(ctx, body, uSERID)
Create User Auth Context

<p>Create User Auth Context. These key value pairs will be propagated over connector to adapter. Normally used for mapping OBP user and<br />Bank User/Customer.<br />Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostUserAuthContextJson**](PostUserAuthContextJson.md)| PostUserAuthContextJson object that needs to be added. | 
  **uSERID** | **string**| The user id | 

### Return type

[**UserAuthContextJsonV500**](UserAuthContextJsonV500.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUserAuthContextUpdateRequest**
> UserAuthContextUpdateJsonV500 CreateUserAuthContextUpdateRequest(ctx, body, sCAMETHOD, bANKID)
Create User Auth Context Update Request

<p>Create User Auth Context Update Request.<br />Authentication is Mandatory</p><p>A One Time Password (OTP) (AKA security challenge) is sent Out of Band (OOB) to the User via the transport defined in SCA_METHOD<br />SCA_METHOD is typically &quot;SMS&quot; or &quot;EMAIL&quot;. &quot;EMAIL&quot; is used for testing purposes.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostUserAuthContextJson**](PostUserAuthContextJson.md)| PostUserAuthContextJson object that needs to be added. | 
  **sCAMETHOD** | **string**| the sca method | 
  **bANKID** | **string**| The bank id | 

### Return type

[**UserAuthContextUpdateJsonV500**](UserAuthContextUpdateJsonV500.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUserCustomerLinks**
> UserCustomerLinkJson CreateUserCustomerLinks(ctx, body, bANKID)
Create User Customer Link

<p>Link a User to a Customer</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateUserCustomerLinkJson**](CreateUserCustomerLinkJson.md)| CreateUserCustomerLinkJson object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**UserCustomerLinkJson**](UserCustomerLinkJson.md)

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

# **CreateUserWithRoles**
> EntitlementsJsonV400 CreateUserWithRoles(ctx, body)
Create (DAuth) User with Roles

<p>This endpoint is used as part of the DAuth solution to grant Entitlements for Roles to a smart contract on the blockchain.</p><p>Put the smart contract address in username</p><p>For provider use &quot;dauth&quot;</p><p>This endpoint will create the User with username and provider if the User does not already exist.</p><p>Then it will create Entitlements i.e. grant Roles to the User.</p><p>Entitlements are used to grant System or Bank level roles to Users. (For Account level privileges, see Views)</p><p>i.e. Entitlements are used to create / consume system or bank level resources where as views / account access are used to consume / create customer level resources.</p><p>For a System level Role (.e.g CanGetAnyUser), set bank_id to an empty string i.e. &quot;bank_id&quot;:&quot;&quot;</p><p>For a Bank level Role (e.g. CanCreateAccount), set bank_id to a valid value e.g. &quot;bank_id&quot;:&quot;my-bank-id&quot;</p><p>Note: The Roles actually granted will depend on the Roles that the calling user has.</p><p>If you try to grant Entitlements to a user that already exist (duplicate entitilements) you will get an error.</p><p>For information about DAuth see below:</p><details>  <summary style=\"display:list-item;cursor:s-resize;\">DAuth</summary>  <h3><a href=\"#dauth-introduction-setup-and-usage\" id=\"dauth-introduction-setup-and-usage\">DAuth Introduction, Setup and Usage</a></h3><p>DAuth is an experimental authentication mechanism that aims to pin an ethereum or other blockchain Smart Contract to an OBP &quot;User&quot;.</p><p>In the future, it might be possible to be more specific and pin specific actors (wallets) that are acting within the smart contract, but so far, one smart contract acts on behalf of one User.</p><p>Thus, if a smart contract &quot;X&quot; calls the OBP API using the DAuth header, OBP will get or create a user called X and the call will proceed in the context of that User &quot;X&quot;.</p><p>DAuth is invoked by the REST client (caller) including a specific header (see step 3 below) in any OBP REST call.</p><p>When OBP receives the DAuth token, it creates or gets a User with a username based on the smart_contract_address and the provider based on the network_name. The combination of username and provider is unique in OBP.</p><p>If you are calling OBP-API via an API3 Airnode, the Airnode will take care of constructing the required header.</p><p>When OBP detects a DAuth header / token it first checks if the Consumer is allowed to make such a call. OBP will validate the Consumer ip address and signature etc.</p><p>Note: The DAuth flow does <em>not</em> require an explicit POST like Direct Login to create the token.</p><p>Permissions may be assigned to an OBP User at any time, via the UserAuthContext, Views, Entitlements to Roles or Consents.</p><p>Note: <em>DAuth is NOT enabled on this instance!</em></p><p>Note: <em>The DAuth client is responsible for creating a token which will be trusted by OBP absolutely</em>!</p><p>To use DAuth:</p><h3><a href=\"#1-configure-obp-api-to-accept-dauth\" id=\"1-configure-obp-api-to-accept-dauth\">1) Configure OBP API to accept DAuth.</a></h3><p>Set up properties in your props file</p><pre><code># -- DAuth --------------------------------------# Define secret used to validate JWT token# jwt.public_key_rsa=path-to-the-pem-file# Enable/Disable DAuth communication at all# In case isn't defined default value is false# allow_dauth=false# Define comma separated list of allowed IP addresses# dauth.host=127.0.0.1# -------------------------------------- DAuth--</code></pre><p>Please keep in mind that property jwt.public_key_rsa is used to validate JWT token to check it is not changed or corrupted during transport.</p><h3><a href=\"#2-create-have-access-to-a-jwt\" id=\"2-create-have-access-to-a-jwt\">2) Create / have access to a JWT</a></h3><p>The following videos are available:<br />* <a href=\"https://vimeo.com/644315074\">DAuth in local environment</a></p><p>HEADER:ALGORITHM &amp; TOKEN TYPE</p><pre><code>{  &quot;alg&quot;: &quot;RS256&quot;,  &quot;typ&quot;: &quot;JWT&quot;}</code></pre><p>PAYLOAD:DATA</p><pre><code>{  &quot;smart_contract_address&quot;: &quot;0xe123425E7734CE288F8367e1Bb143E90bb3F051224&quot;,  &quot;network_name&quot;: &quot;AIRNODE.TESTNET.ETHEREUM&quot;,  &quot;msg_sender&quot;: &quot;0xe12340927f1725E7734CE288F8367e1Bb143E90fhku767&quot;,  &quot;consumer_key&quot;: &quot;0x1234a4ec31e89cea54d1f125db7536e874ab4a96b4d4f6438668b6bb10a6adb&quot;,  &quot;timestamp&quot;: &quot;2021-11-04T14:13:40Z&quot;,  &quot;request_id&quot;: &quot;0Xe876987694328763492876348928736497869273649&quot;}</code></pre><p>VERIFY SIGNATURE</p><pre><code>RSASHA256(  base64UrlEncode(header) + &quot;.&quot; +  base64UrlEncode(payload),<p>) your-RSA-key-pair</p></code></pre><p>Here is an example token:</p><pre><code>eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzbWFydF9jb250cmFjdF9hZGRyZXNzIjoiMHhlMTIzNDI1RTc3MzRDRTI4OEY4MzY3ZTFCYjE0M0U5MGJiM0YwNTEyMjQiLCJuZXR3b3JrX25hbWUiOiJFVEhFUkVVTSIsIm1zZ19zZW5kZXIiOiIweGUxMjM0MDkyN2YxNzI1RTc3MzRDRTI4OEY4MzY3ZTFCYjE0M0U5MGZoa3U3NjciLCJjb25zdW1lcl9rZXkiOiIweDEyMzRhNGVjMzFlODljZWE1NGQxZjEyNWRiNzUzNmU4NzRhYjRhOTZiNGQ0ZjY0Mzg2NjhiNmJiMTBhNmFkYiIsInRpbWVzdGFtcCI6IjIwMjEtMTEtMDRUMTQ6MTM6NDBaIiwicmVxdWVzdF9pZCI6IjBYZTg3Njk4NzY5NDMyODc2MzQ5Mjg3NjM0ODkyODczNjQ5Nzg2OTI3MzY0OSJ9.XSiQxjEVyCouf7zT8MubEKsbOBZuReGVhnt9uck6z6k</code></pre><h3><a href=\"#3-try-a-rest-call-using-the-header\" id=\"3-try-a-rest-call-using-the-header\">3) Try a REST call using the header</a></h3><p>Using your favorite http client:</p><p>GET <a href=\"https://apisandbox.openbankproject.com/obp/v3.0.0/users/current\">https://apisandbox.openbankproject.com/obp/v3.0.0/users/current</a></p><p>Body</p><p>Leave Empty!</p><p>Headers:</p><pre><code>   DAuth: your-jwt-from-step-above</code></pre><p>Here is it all together:</p><p>GET <a href=\"https://apisandbox.openbankproject.com/obp/v3.0.0/users/current\">https://apisandbox.openbankproject.com/obp/v3.0.0/users/current</a> HTTP/1.1<br />Host: localhost:8080<br />User-Agent: curl/7.47.0<br />Accept: <em>/</em><br />DAuth: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzbWFydF9jb250cmFjdF9hZGRyZXNzIjoiMHhlMTIzNDI1RTc3MzRDRTI4OEY4MzY3ZTFCYjE0M0U5MGJiM0YwNTEyMjQiLCJuZXR3b3JrX25hbWUiOiJFVEhFUkVVTSIsIm1zZ19zZW5kZXIiOiIweGUxMjM0MDkyN2YxNzI1RTc3MzRDRTI4OEY4MzY3ZTFCYjE0M0U5MGZoa3U3NjciLCJjb25zdW1lcl9rZXkiOiIweDEyMzRhNGVjMzFlODljZWE1NGQxZjEyNWRiNzUzNmU4NzRhYjRhOTZiNGQ0ZjY0Mzg2NjhiNmJiMTBhNmFkYiIsInRpbWVzdGFtcCI6IjIwMjEtMTEtMDRUMTQ6MTM6NDBaIiwicmVxdWVzdF9pZCI6IjBYZTg3Njk4NzY5NDMyODc2MzQ5Mjg3NjM0ODkyODczNjQ5Nzg2OTI3MzY0OSJ9.XSiQxjEVyCouf7zT8MubEKsbOBZuReGVhnt9uck6z6k</p><p>CURL example</p><pre><code>curl -v -H 'DAuth: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzbWFydF9jb250cmFjdF9hZGRyZXNzIjoiMHhlMTIzNDI1RTc3MzRDRTI4OEY4MzY3ZTFCYjE0M0U5MGJiM0YwNTEyMjQiLCJuZXR3b3JrX25hbWUiOiJFVEhFUkVVTSIsIm1zZ19zZW5kZXIiOiIweGUxMjM0MDkyN2YxNzI1RTc3MzRDRTI4OEY4MzY3ZTFCYjE0M0U5MGZoa3U3NjciLCJjb25zdW1lcl9rZXkiOiIweDEyMzRhNGVjMzFlODljZWE1NGQxZjEyNWRiNzUzNmU4NzRhYjRhOTZiNGQ0ZjY0Mzg2NjhiNmJiMTBhNmFkYiIsInRpbWVzdGFtcCI6IjIwMjEtMTEtMDRUMTQ6MTM6NDBaIiwicmVxdWVzdF9pZCI6IjBYZTg3Njk4NzY5NDMyODc2MzQ5Mjg3NjM0ODkyODczNjQ5Nzg2OTI3MzY0OSJ9.XSiQxjEVyCouf7zT8MubEKsbOBZuReGVhnt9uck6z6k' https://apisandbox.openbankproject.com/obp/v3.0.0/users/current</code></pre><p>You should receive a response like:</p><pre><code>{    &quot;user_id&quot;: &quot;4c4d3175-1e5c-4cfd-9b08-dcdc209d8221&quot;,    &quot;email&quot;: &quot;&quot;,    &quot;provider_id&quot;: &quot;0xe123425E7734CE288F8367e1Bb143E90bb3F051224&quot;,    &quot;provider&quot;: &quot;ETHEREUM&quot;,    &quot;username&quot;: &quot;0xe123425E7734CE288F8367e1Bb143E90bb3F051224&quot;,    &quot;entitlements&quot;: {        &quot;list&quot;: []    }}</code></pre><h3><a href=\"#under-the-hood\" id=\"under-the-hood\">Under the hood</a></h3><p>The file, dauth.scala handles the DAuth,</p><p>We:</p><pre><code>-&gt; Check if Props allow_dauth is true  -&gt; Check if DAuth header exists    -&gt; Check if getRemoteIpAddress is OK      -&gt; Look for &quot;token&quot;        -&gt; parse the JWT token and getOrCreate the user          -&gt; get the data of the user</code></pre><h3><a href=\"#more-information\" id=\"more-information\">More information</a></h3><p>Parameter names and values are case sensitive.<br />Each parameter MUST NOT appear more than once per request.</p></details><p><br></br></p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostCreateUserWithRolesJsonV400**](PostCreateUserWithRolesJsonV400.md)| PostCreateUserWithRolesJsonV400 object that needs to be added. | 

### Return type

[**EntitlementsJsonV400**](EntitlementsJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEntitlement**
> EmptyClassJson DeleteEntitlement(ctx, body, eNTITLEMENTID, uSERID)
Delete Entitlement

<p>Delete Entitlement specified by ENTITLEMENT_ID for an user specified by USER_ID</p><p>Authentication is required and the user needs to be a Super Admin.<br />Super Admins are listed in the Props file.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **eNTITLEMENTID** | **string**| The entitblement id | 
  **uSERID** | **string**| The user id | 

### Return type

[**EmptyClassJson**](EmptyClassJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEntitlementRequest**
> DeleteEntitlementRequest(ctx, eNTITLEMENTREQUESTID)
Delete Entitlement Request

<p>Delete the Entitlement Request specified by ENTITLEMENT_REQUEST_ID for a user specified by USER_ID</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eNTITLEMENTREQUESTID** | **string**| the entitlement request id | 

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNonPersonalUserAttribute**
> DeleteNonPersonalUserAttribute(ctx, uSERID)
Delete Non Personal User Attribute

<p>Delete the Non Personal User Attribute specified by ENTITLEMENT_REQUEST_ID for a user specified by USER_ID</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uSERID** | **string**| The user id | 

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUser**
> DeleteUser(ctx, uSERID)
Delete a User

<p>Delete a User.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uSERID** | **string**| The user id | 

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUserAuthContextById**
> DeleteUserAuthContextById(ctx, uSERAUTHCONTEXTID, uSERID)
Delete User Auth Context

<p>Delete a User AuthContext of the User specified by USER_AUTH_CONTEXT_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uSERAUTHCONTEXTID** | **string**| the user auth context id | 
  **uSERID** | **string**| The user id | 

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUserAuthContexts**
> DeleteUserAuthContexts(ctx, uSERID)
Delete User's Auth Contexts

<p>Delete the Auth Contexts of a User specified by USER_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uSERID** | **string**| The user id | 

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllEntitlementRequests**
> EntitlementRequestsJson GetAllEntitlementRequests(ctx, )
Get all Entitlement Requests

<p>Get all Entitlement Requests</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EntitlementRequestsJson**](EntitlementRequestsJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrentUser**
> UserJsonV300 GetCurrentUser(ctx, )
Get User (Current)

<p>Get the logged in user</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UserJsonV300**](UserJsonV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrentUserId**
> UserIdJsonV400 GetCurrentUserId(ctx, )
Get User Id (Current)

<p>Get the USER_ID of the logged in user</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UserIdJsonV400**](UserIdJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomersAtAnyBank**
> CustomerJsonsV300 GetCustomersAtAnyBank(ctx, )
Get Customers at Any Bank

<p>Get Customers at Any Bank.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CustomerJsonsV300**](CustomerJSONsV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomersAtOneBank**
> CustomerJsonsV300 GetCustomersAtOneBank(ctx, bANKID)
Get Customers at Bank

<p>Get Customers at Bank.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerJsonsV300**](CustomerJSONsV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomersForUser**
> CustomersWithAttributesJsonV300 GetCustomersForUser(ctx, )
Get Customers for Current User

<p>Gets all Customers that are linked to a User.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CustomersWithAttributesJsonV300**](CustomersWithAttributesJsonV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomersForUserIdsOnly**
> CustomersWithAttributesJsonV300 GetCustomersForUserIdsOnly(ctx, )
Get Customers for Current User (IDs only)

<p>Gets all Customers Ids that are linked to a User.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CustomersWithAttributesJsonV300**](CustomersWithAttributesJsonV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomersMinimalAtAnyBank**
> CustomersMinimalJsonV400 GetCustomersMinimalAtAnyBank(ctx, )
Get Customers Minimal at Any Bank

<p>Get Customers Minimal at Any Bank.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CustomersMinimalJsonV400**](CustomersMinimalJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomersMinimalAtOneBank**
> CustomersMinimalJsonV400 GetCustomersMinimalAtOneBank(ctx, bANKID)
Get Customers Minimal at Bank

<p>Get Customers Minimal at Bank.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomersMinimalJsonV400**](CustomersMinimalJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEntitlementRequests**
> EntitlementRequestsJson GetEntitlementRequests(ctx, uSERID)
Get Entitlement Requests for a User

<p>Get Entitlement Requests for a User.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uSERID** | **string**| The user id | 

### Return type

[**EntitlementRequestsJson**](EntitlementRequestsJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEntitlementRequestsForCurrentUser**
> EntitlementRequestsJson GetEntitlementRequestsForCurrentUser(ctx, )
Get Entitlement Requests for the current User

<p>Get Entitlement Requests for the current User.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EntitlementRequestsJson**](EntitlementRequestsJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEntitlements**
> EntitlementsJsonV400 GetEntitlements(ctx, uSERID)
Get Entitlements for User

<p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uSERID** | **string**| The user id | 

### Return type

[**EntitlementsJsonV400**](EntitlementsJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEntitlementsAndPermissions**
> UserJsonV300 GetEntitlementsAndPermissions(ctx, uSERID)
Get Entitlements and Permissions for a User

<p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uSERID** | **string**| The user id | 

### Return type

[**UserJsonV300**](UserJsonV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEntitlementsByBankAndUser**
> EntitlementJsons GetEntitlementsByBankAndUser(ctx, body, uSERID, bANKID)
Get Entitlements for User at Bank

<p>Get Entitlements specified by BANK_ID and USER_ID</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **uSERID** | **string**| The user id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**EntitlementJsons**](EntitlementJSONs.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEntitlementsForBank**
> EntitlementsJsonV400 GetEntitlementsForBank(ctx, bANKID)
Get Entitlements for One Bank

<p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**EntitlementsJsonV400**](EntitlementsJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEntitlementsForCurrentUser**
> EntitlementJsons GetEntitlementsForCurrentUser(ctx, )
Get Entitlements for the current User

<p>Get Entitlements for the current User.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EntitlementJsons**](EntitlementJSONs.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogoutLink**
> LogoutLinkJson GetLogoutLink(ctx, )
Get Logout Link

<p>Get the Logout Link</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**LogoutLinkJson**](LogoutLinkJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMyCustomersAtAnyBank**
> CustomerJsonV210 GetMyCustomersAtAnyBank(ctx, )
Get My Customers

<p>Gets all Customers that are linked to me.</p><p>Authentication via OAuth is required.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CustomerJsonV210**](CustomerJsonV210.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMyPersonalUserAttributes**
> UserAttributesResponseJson GetMyPersonalUserAttributes(ctx, )
Get My Personal User Attributes

<p>Get My Personal User Attributes.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UserAttributesResponseJson**](UserAttributesResponseJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMySpaces**
> MySpaces GetMySpaces(ctx, )
Get My Spaces

<p>Get My Spaces.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MySpaces**](MySpaces.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNonPersonalUserAttributes**
> GetNonPersonalUserAttributes(ctx, uSERID)
Get Non Personal User Attributes

<p>Get Non Personal User Attribute for a user specified by USER_ID</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uSERID** | **string**| The user id | 

### Return type

 (empty response body)

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

# **GetUserAuthContexts**
> UserAuthContextJsonV500 GetUserAuthContexts(ctx, uSERID)
Get User Auth Contexts

<p>Get User Auth Contexts for a User.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uSERID** | **string**| The user id | 

### Return type

[**UserAuthContextJsonV500**](UserAuthContextJsonV500.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserByProviderAndUsername**
> UserJsonV400 GetUserByProviderAndUsername(ctx, pROVIDER, uSERNAME)
Get User by USERNAME

<p>Get user by PROVIDER and USERNAME</p><p>Authentication is Mandatory</p><p>CanGetAnyUser entitlement is required,</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pROVIDER** | **string**| the user PROVIDER | 
  **uSERNAME** | **string**| the user name | 

### Return type

[**UserJsonV400**](UserJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserByUserId**
> UserJsonV400 GetUserByUserId(ctx, uSERID)
Get User by USER_ID

<p>Get user by USER_ID</p><p>Authentication is Mandatory<br />CanGetAnyUser entitlement is required,</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uSERID** | **string**| The user id | 

### Return type

[**UserJsonV400**](UserJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserLockStatus**
> BadLoginStatusJson GetUserLockStatus(ctx, pROVIDER, uSERNAME)
Get User Lock Status

<p>Get User Login Status.<br />Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pROVIDER** | **string**| the user PROVIDER | 
  **uSERNAME** | **string**| the user name | 

### Return type

[**BadLoginStatusJson**](BadLoginStatusJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserWithAttributes**
> UserWithAttributesResponseJson GetUserWithAttributes(ctx, uSERID)
Get User with Attributes by USER_ID

<p>Get User Attributes for the user defined via USER_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uSERID** | **string**| The user id | 

### Return type

[**UserWithAttributesResponseJson**](UserWithAttributesResponseJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsers**
> UsersJsonV400 GetUsers(ctx, )
Get all Users

<p>Get all users</p><p>Authentication is Mandatory</p><p>CanGetAnyUser entitlement is required,</p><p>Possible custom url parameters for pagination:</p><ul><li>limit=NUMBER ==&gt; default value: 500</li><li>offset=NUMBER ==&gt; default value: 0</li></ul><p>eg1:?limit=100&amp;offset=0</p><ul><li>sort_direction=ASC/DESC ==&gt; default value: DESC.</li></ul><p>eg2:?limit=100&amp;offset=0&amp;sort_direction=ASC</p><ul><li>locked_status (if null ignore)</li></ul>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UsersJsonV400**](UsersJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsersByEmail**
> UsersJsonV400 GetUsersByEmail(ctx, )
Get Users by Email Address

<p>Get users by email address</p><p>Authentication is Mandatory<br />CanGetAnyUser entitlement is required,</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UsersJsonV400**](UsersJsonV400.md)

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

# **LockUserByProviderAndUsername**
> UserLockStatusJson LockUserByProviderAndUsername(ctx, pROVIDER, uSERNAME)
Lock the user

<p>Lock a User.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pROVIDER** | **string**| the user PROVIDER | 
  **uSERNAME** | **string**| the user name | 

### Return type

[**UserLockStatusJson**](UserLockStatusJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefreshUser**
> RefreshUserJson RefreshUser(ctx, uSERID)
Refresh User

<p>The endpoint is used for updating the accounts, views, account holders for the user.<br />As to the Json body, you can leave it as Empty.<br />This call will get data from backend, no need to prepare the json body in api side.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uSERID** | **string**| The user id | 

### Return type

[**RefreshUserJson**](RefreshUserJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetPasswordUrl**
> ResetPasswordUrlJsonV400 ResetPasswordUrl(ctx, body)
Create password reset url

<p>Create password reset url.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostResetPasswordUrlJsonV400**](PostResetPasswordUrlJsonV400.md)| PostResetPasswordUrlJsonV400 object that needs to be added. | 

### Return type

[**ResetPasswordUrlJsonV400**](ResetPasswordUrlJsonV400.md)

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

# **UnlockUserByProviderAndUsername**
> BadLoginStatusJson UnlockUserByProviderAndUsername(ctx, pROVIDER, uSERNAME)
Unlock the user

<p>Unlock a User.</p><p>(Perhaps the user was locked due to multiple failed login attempts)</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pROVIDER** | **string**| the user PROVIDER | 
  **uSERNAME** | **string**| the user name | 

### Return type

[**BadLoginStatusJson**](BadLoginStatusJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMyPersonalUserAttribute**
> UserAttributeResponseJsonV400 UpdateMyPersonalUserAttribute(ctx, body)
Update My Personal User Attribute

<p>Update My Personal User Attribute for current user by USER_ATTRIBUTE_ID</p><p>The type field must be one of &quot;STRING&quot;, &quot;INTEGER&quot;, &quot;DOUBLE&quot; or DATE_WITH_DAY&quot;</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserAttributeJsonV400**](UserAttributeJsonV400.md)| UserAttributeJsonV400 object that needs to be added. | 

### Return type

[**UserAttributeResponseJsonV400**](UserAttributeResponseJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

