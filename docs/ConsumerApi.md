# \ConsumerApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddScope**](ConsumerApi.md#AddScope) | **Post** /obp/v5.1.0/consumers/{CONSUMER_ID}/scopes | Create Scope for a Consumer
[**CallsLimit**](ConsumerApi.md#CallsLimit) | **Put** /obp/v5.1.0/management/consumers/{CONSUMER_ID}/consumer/call-limits | Set Rate Limits / Call Limits per Consumer
[**CreateConsumer**](ConsumerApi.md#CreateConsumer) | **Post** /obp/v5.1.0/dynamic-registration/consumers | Create a Consumer
[**CreateConsumer_0**](ConsumerApi.md#CreateConsumer_0) | **Post** /obp/v5.1.0/management/consumers | Post a Consumer
[**DeleteScope**](ConsumerApi.md#DeleteScope) | **Delete** /obp/v5.1.0/consumers/{CONSUMER_ID}/scope/{SCOPE_ID} | Delete Consumer Scope
[**EnableDisableConsumers**](ConsumerApi.md#EnableDisableConsumers) | **Put** /obp/v5.1.0/management/consumers/{CONSUMER_ID} | Enable or Disable Consumers
[**GetCallsLimit**](ConsumerApi.md#GetCallsLimit) | **Get** /obp/v5.1.0/management/consumers/{CONSUMER_ID}/consumer/call-limits | Get Call Limits for a Consumer
[**GetConsumer**](ConsumerApi.md#GetConsumer) | **Get** /obp/v5.1.0/management/consumers/{CONSUMER_ID} | Get Consumer
[**GetConsumers**](ConsumerApi.md#GetConsumers) | **Get** /obp/v5.1.0/management/consumers | Get Consumers
[**GetConsumersForCurrentUser**](ConsumerApi.md#GetConsumersForCurrentUser) | **Get** /obp/v5.1.0/management/users/current/consumers | Get Consumers (logged in User)
[**GetScopes**](ConsumerApi.md#GetScopes) | **Get** /obp/v5.1.0/consumers/{CONSUMER_ID}/scopes | Get Scopes for Consumer
[**UpdateConsumerRedirectUrl**](ConsumerApi.md#UpdateConsumerRedirectUrl) | **Put** /obp/v5.1.0/management/consumers/{CONSUMER_ID}/consumer/redirect_url | Update Consumer RedirectUrl


# **AddScope**
> ScopeJson AddScope(ctx, body, cONSUMERID)
Create Scope for a Consumer

<p>Create Scope. Grant Role to Consumer.</p><p>Scopes are used to grant System or Bank level roles to the Consumer (App). (For Account level privileges, see Views)</p><p>For a System level Role (.e.g CanGetAnyUser), set bank_id to an empty string i.e. &quot;bank_id&quot;:&quot;&quot;</p><p>For a Bank level Role (e.g. CanCreateAccount), set bank_id to a valid value e.g. &quot;bank_id&quot;:&quot;my-bank-id&quot;</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateScopeJson**](CreateScopeJson.md)| CreateScopeJson object that needs to be added. | 
  **cONSUMERID** | **string**| new consumer id | 

### Return type

[**ScopeJson**](ScopeJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CallsLimit**
> CallLimitPostJsonV400 CallsLimit(ctx, body, cONSUMERID)
Set Rate Limits / Call Limits per Consumer

<p>Set the API rate limits / call limits for a Consumer:</p><p>Rate limiting can be set:</p><p>Per Second<br />Per Minute<br />Per Hour<br />Per Week<br />Per Month</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CallLimitPostJsonV400**](CallLimitPostJsonV400.md)| CallLimitPostJsonV400 object that needs to be added. | 
  **cONSUMERID** | **string**| new consumer id | 

### Return type

[**CallLimitPostJsonV400**](CallLimitPostJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateConsumer**
> ConsumerJsonV510 CreateConsumer(ctx, body)
Create a Consumer

<p>Create a Consumer (mTLS access).</p><p>JWT payload:<br />- minimal<br />{ &quot;description&quot;:&quot;Description&quot; }<br />- full<br />{<br />&quot;description&quot;: &quot;Description&quot;,<br />&quot;app_name&quot;: &quot;Tesobe GmbH&quot;,<br />&quot;app_type&quot;: &quot;Sofit&quot;,<br />&quot;developer_email&quot;: &quot;<a href=\"m&#97;i&#x6c;&#x74;&#111;&#x3a;&#109;&#x61;&#x72;&#x6b;&#111;@&#116;&#x65;&#115;&#x6f;&#98;e&#x2e;&#x63;o&#109;\">&#109;&#x61;&#114;ko&#x40;te&#x73;&#x6f;&#98;&#101;.&#99;&#111;m</a>&quot;,<br />&quot;redirect_url&quot;: &quot;<a href=\"http://localhost:8082\">http://localhost:8082</a>&quot;<br />}<br />Please note that JWT must be signed with the counterpart private kew of the public key used to establish mTLS</p><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConsumerJwtPostJsonV510**](ConsumerJwtPostJsonV510.md)| ConsumerJwtPostJsonV510 object that needs to be added. | 

### Return type

[**ConsumerJsonV510**](ConsumerJsonV510.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateConsumer_0**
> ConsumerJson CreateConsumer_0(ctx, body)
Post a Consumer

<p>Create a Consumer (Authenticated access).</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConsumerPostJson**](ConsumerPostJson.md)| ConsumerPostJSON object that needs to be added. | 

### Return type

[**ConsumerJson**](ConsumerJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteScope**
> DeleteScope(ctx, sCOPEID, cONSUMERID)
Delete Consumer Scope

<p>Delete Consumer Scope specified by SCOPE_ID for an consumer specified by CONSUMER_ID</p><p>Authentication is required and the user needs to be a Super Admin.<br />Super Admins are listed in the Props file.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sCOPEID** | **string**| the scope id | 
  **cONSUMERID** | **string**| new consumer id | 

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableDisableConsumers**
> PutEnabledJson EnableDisableConsumers(ctx, body, cONSUMERID)
Enable or Disable Consumers

<p>Enable/Disable a Consumer specified by CONSUMER_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PutEnabledJson**](PutEnabledJson.md)| PutEnabledJSON object that needs to be added. | 
  **cONSUMERID** | **string**| new consumer id | 

### Return type

[**PutEnabledJson**](PutEnabledJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCallsLimit**
> CallLimitJson GetCallsLimit(ctx, cONSUMERID)
Get Call Limits for a Consumer

<p>Get Calls limits per Consumer.<br />Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cONSUMERID** | **string**| new consumer id | 

### Return type

[**CallLimitJson**](CallLimitJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConsumer**
> ConsumerJsonV210 GetConsumer(ctx, cONSUMERID)
Get Consumer

<p>Get the Consumer specified by CONSUMER_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cONSUMERID** | **string**| new consumer id | 

### Return type

[**ConsumerJsonV210**](ConsumerJsonV210.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConsumers**
> ConsumersJsonV310 GetConsumers(ctx, )
Get Consumers

<p>Get the all Consumers.</p><p>Authentication is Mandatory</p><p>Possible custom url parameters for pagination:</p><ul><li>limit=NUMBER ==&gt; default value: 500</li><li>offset=NUMBER ==&gt; default value: 0</li></ul><p>eg1:?limit=100&amp;offset=0</p><ul><li>sort_direction=ASC/DESC ==&gt; default value: DESC.</li></ul><p>eg2:?limit=100&amp;offset=0&amp;sort_direction=ASC</p><ul><li>from_date=DATE =&gt; example value: 1970-01-01T00:00:00.000Z. NOTE! The default value is one year ago (1970-01-01T00:00:00.000Z).</li><li>to_date=DATE =&gt; example value: 2024-02-05T14:15:55.293Z. NOTE! The default value is now (2024-02-05T14:15:55.293Z).</li></ul><p>Date format parameter: yyyy-MM-dd'T'HH:mm:ss.SSS'Z'(1100-01-01T01:01:01.000Z) ==&gt; time zone is UTC.</p><p>eg3:?sort_direction=ASC&amp;limit=100&amp;offset=0&amp;from_date=1100-01-01T01:01:01.000Z&amp;to_date=1100-01-01T01:01:01.000Z</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ConsumersJsonV310**](ConsumersJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConsumersForCurrentUser**
> ConsumersJsonV310 GetConsumersForCurrentUser(ctx, )
Get Consumers (logged in User)

<p>Get the Consumers for logged in User.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ConsumersJsonV310**](ConsumersJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetScopes**
> ScopeJsons GetScopes(ctx, cONSUMERID)
Get Scopes for Consumer

<p>Get all the scopes for an consumer specified by CONSUMER_ID</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cONSUMERID** | **string**| new consumer id | 

### Return type

[**ScopeJsons**](ScopeJsons.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConsumerRedirectUrl**
> ConsumerJsonV210 UpdateConsumerRedirectUrl(ctx, body, cONSUMERID)
Update Consumer RedirectUrl

<p>Update an existing redirectUrl for a Consumer specified by CONSUMER_ID.</p><p>Please note: Your consumer may be disabled as a result of this action.</p><p>CONSUMER_ID can be obtained after you register the application.</p><p>Or use the endpoint 'Get Consumers' to get it</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConsumerRedirectUrlJson**](ConsumerRedirectUrlJson.md)| ConsumerRedirectUrlJSON object that needs to be added. | 
  **cONSUMERID** | **string**| new consumer id | 

### Return type

[**ConsumerJsonV210**](ConsumerJsonV210.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

