# \DynamicEndpointApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DynamicEndpointGETAccountsACCOUNTID**](DynamicEndpointApi.md#DynamicEndpointGETAccountsACCOUNTID) | **Get** /accounts/{ACCOUNT_ID} | Get Bank Account By Id
[**DynamicEndpointPOSTAccounts**](DynamicEndpointApi.md#DynamicEndpointPOSTAccounts) | **Post** /accounts | Post Accounts


# **DynamicEndpointGETAccountsACCOUNTID**
> InlineResponse2012 DynamicEndpointGETAccountsACCOUNTID(ctx, aCCOUNTID)
Get Bank Account By Id

<p>Get Bank Account</p><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEndpointProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;url_pattern&quot;,        &quot;value&quot;:&quot;http://obp_mock//accounts/{account_id}&quot;    },    {        &quot;key&quot;:&quot;http_method&quot;,        &quot;value&quot;:&quot;GET&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aCCOUNTID** | **string**| The account id | 

### Return type

[**InlineResponse2012**](inline_response_201_2.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEndpointPOSTAccounts**
> InlineResponse2012 DynamicEndpointPOSTAccounts(ctx, )
Post Accounts

<p>POST Accounts</p><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEndpointProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;url_pattern&quot;,        &quot;value&quot;:&quot;http://obp_mock//accounts&quot;    },    {        &quot;key&quot;:&quot;http_method&quot;,        &quot;value&quot;:&quot;POST&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2012**](inline_response_201_2.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

