# \MethodRoutingApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMethodRouting**](MethodRoutingApi.md#CreateMethodRouting) | **Post** /obp/v5.1.0/management/method_routings | Create MethodRouting
[**DeleteMethodRouting**](MethodRoutingApi.md#DeleteMethodRouting) | **Delete** /obp/v5.1.0/management/method_routings/{METHOD_ROUTING_ID} | Delete MethodRouting
[**GetMethodRoutings**](MethodRoutingApi.md#GetMethodRoutings) | **Get** /obp/v5.1.0/management/method_routings | Get MethodRoutings
[**UpdateMethodRouting**](MethodRoutingApi.md#UpdateMethodRouting) | **Put** /obp/v5.1.0/management/method_routings/{METHOD_ROUTING_ID} | Update MethodRouting


# **CreateMethodRouting**
> MethodRoutingCommons CreateMethodRouting(ctx, body)
Create MethodRouting

<p>Create a MethodRouting.</p><p>Authentication is Mandatory</p><p>Explanation of Fields:</p><ul><li>method_name is required String value, current supported value: [mapped | internal | rest_vMar2019]</li><li>connector_name is required String value</li><li>is_bank_id_exact_match is required boolean value, if bank_id_pattern is exact bank_id value, this value is true; if bank_id_pattern is null or a regex, this value is false</li><li>bank_id_pattern is optional String value, it can be null, a exact bank_id or a regex</li><li>parameters is optional array of key value pairs. You can set some parameters for this method</li></ul><p>note and CAVEAT!:</p><ul><li>bank_id_pattern has to be empty for methods that do not take bank_id as a function parameter, otherwise might get empty result</li><li>methods that aggregate bank objects (e.g. getBankAccountsForUser) have to take any  existing method routings for these objects into consideration</li><li>so if you create e.g. a bank specific method routing for getting an account, make sure that it is also served by endpoints getting ALL accounts for ALL banks</li><li>if bank_id_pattern is regex, special characters need to do escape, for example: bank_id_pattern = &quot;some-id_pattern_\\d+&quot;</li></ul><p>If the connector name starts with rest, parameters can contain &quot;outBoundMapping&quot; and &quot;inBoundMapping&quot;, convert OutBound and InBound json structure.<br />for example:<br />outBoundMapping example, convert json from source to target:<br /><img src=\"https://user-images.githubusercontent.com/2577334/75248007-33332e00-580e-11ea-8d2a-d1856035fa24.png\" alt=\"Snipaste_outBoundMapping\" /><br />Build OutBound json value rules:<br />1 set cId value with: outboundAdapterCallContext.correlationId value<br />2 set bankId value with: concat bankId.value value with  string helloworld<br />3 set originalJson value with: whole source json, note: the field value expression is $root</p><p>inBoundMapping example, convert json from source to target:<br /><img src=\"https://user-images.githubusercontent.com/2577334/75248199-a9d02b80-580e-11ea-9238-e073264e9170.png\" alt=\"inBoundMapping\" /><br />Build InBound json value rules:<br />1 and 2 set inboundAdapterCallContext and status value: because field name ends with &quot;$default&quot;, remove &quot;$default&quot; from field name, not change the value<br />3 set fullName value with: concat string full: with result.name value<br />4 set bankRoutingScheme value: because source value is Array, but target value is not Array, the mapping field name must ends with [0].</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MethodRoutingCommons**](MethodRoutingCommons.md)| MethodRoutingCommons object that needs to be added. | 

### Return type

[**MethodRoutingCommons**](MethodRoutingCommons.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMethodRouting**
> DeleteMethodRouting(ctx, mETHODROUTINGID)
Delete MethodRouting

<p>Delete a MethodRouting specified by METHOD_ROUTING_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mETHODROUTINGID** | **string**| the method routing id  | 

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMethodRoutings**
> InlineResponse20010 GetMethodRoutings(ctx, )
Get MethodRoutings

<p>Get the all MethodRoutings.</p><p>Query url parameters:</p><ul><li>method_name: filter with method_name</li><li>active: if active = true, it will show all the webui_ props. Even if they are set yet, we will return all the default webui_ props</li></ul><p>eg:<br /><a href=\"https://apisandbox.openbankproject.com/obp/v3.1.0/management/method_routings?active=true\">https://apisandbox.openbankproject.com/obp/v3.1.0/management/method_routings?active=true</a><br /><a href=\"https://apisandbox.openbankproject.com/obp/v3.1.0/management/method_routings?method_name=getBank\">https://apisandbox.openbankproject.com/obp/v3.1.0/management/method_routings?method_name=getBank</a></p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMethodRouting**
> MethodRoutingCommons UpdateMethodRouting(ctx, body, mETHODROUTINGID)
Update MethodRouting

<p>Update a MethodRouting.</p><p>Authentication is Mandatory</p><p>Explaination of Fields:</p><ul><li>method_name is required String value, current supported value: [mapped | internal | rest_vMar2019]</li><li>connector_name is required String value</li><li>is_bank_id_exact_match is required boolean value, if bank_id_pattern is exact bank_id value, this value is true; if bank_id_pattern is null or a regex, this value is false</li><li>bank_id_pattern is optional String value, it can be null, a exact bank_id or a regex</li><li>parameters is optional array of key value pairs. You can set some paremeters for this method<br />note:</li><li><p>if bank_id_pattern is regex, special characters need to do escape, for example: bank_id_pattern = &quot;some-id_pattern_\\d+&quot;</p></li></ul><p>If connector name start with rest, parameters can contain &quot;outBoundMapping&quot; and &quot;inBoundMapping&quot;, to convert OutBound and InBound json structure.<br />for example:<br />outBoundMapping example, convert json from source to target:<br /><img src=\"https://user-images.githubusercontent.com/2577334/75248007-33332e00-580e-11ea-8d2a-d1856035fa24.png\" alt=\"Snipaste_outBoundMapping\" /><br />Build OutBound json value rules:<br />1 set cId value with: outboundAdapterCallContext.correlationId value<br />2 set bankId value with: concat bankId.value value with  string helloworld<br />3 set originalJson value with: whole source json, note: the field value expression is $root</p><p>inBoundMapping example, convert json from source to target:<br /><img src=\"https://user-images.githubusercontent.com/2577334/75248199-a9d02b80-580e-11ea-9238-e073264e9170.png\" alt=\"inBoundMapping\" /><br />Build InBound json value rules:<br />1 and 2 set inboundAdapterCallContext and status value: because field name ends with &quot;$default&quot;, remove &quot;$default&quot; from field name, not change the value<br />3 set fullName value with: concat string full: with result.name value<br />4 set bankRoutingScheme value: because source value is Array, but target value is not Array, the mapping field name must ends with [0].</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MethodRoutingCommons**](MethodRoutingCommons.md)| MethodRoutingCommons object that needs to be added. | 
  **mETHODROUTINGID** | **string**| the method routing id  | 

### Return type

[**MethodRoutingCommons**](MethodRoutingCommons.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

