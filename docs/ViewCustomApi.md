# \ViewCustomApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserWithAccountAccess**](ViewCustomApi.md#CreateUserWithAccountAccess) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/user-account-access | Create (DAuth) User with Account Access
[**CreateViewForBankAccount**](ViewCustomApi.md#CreateViewForBankAccount) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/views | Create Custom View
[**DeleteViewForBankAccount**](ViewCustomApi.md#DeleteViewForBankAccount) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/views/{VIEW_ID} | Delete Custom View
[**GetAccountsHeld**](ViewCustomApi.md#GetAccountsHeld) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts-held | Get Accounts Held
[**GetPermissionForUserForBankAccount**](ViewCustomApi.md#GetPermissionForUserForBankAccount) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/permissions/{PROVIDER}/{PROVIDER_ID} | Get Account access for User
[**GetPermissionsForBankAccount**](ViewCustomApi.md#GetPermissionsForBankAccount) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/permissions | Get access
[**GetViewsForBankAccount**](ViewCustomApi.md#GetViewsForBankAccount) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/views | Get Views for Account
[**GrantUserAccessToView**](ViewCustomApi.md#GrantUserAccessToView) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/account-access/grant | Grant User access to View
[**RevokeGrantUserAccessToViews**](ViewCustomApi.md#RevokeGrantUserAccessToViews) | **Put** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/account-access | Revoke/Grant User access to View
[**RevokeUserAccessToView**](ViewCustomApi.md#RevokeUserAccessToView) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/account-access/revoke | Revoke User access to View
[**UpdateViewForBankAccount**](ViewCustomApi.md#UpdateViewForBankAccount) | **Put** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/views/{VIEW_ID} | Update Custom View


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

