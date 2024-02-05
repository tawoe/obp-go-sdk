# \ObpActivityObpTesting01Api

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DynamicEntityCreateObpActivityObpTesting01**](ObpActivityObpTesting01Api.md#DynamicEntityCreateObpActivityObpTesting01) | **Post** /banks/obp.testing.01/ObpActivity | Create new Obp Activity
[**DynamicEntityDeleteObpActivityObpTesting01**](ObpActivityObpTesting01Api.md#DynamicEntityDeleteObpActivityObpTesting01) | **Delete** /banks/obp.testing.01/ObpActivity/OBP_ACTIVITY_ID | Delete Obp Activity by id
[**DynamicEntityGetObpActivityListObpTesting01**](ObpActivityObpTesting01Api.md#DynamicEntityGetObpActivityListObpTesting01) | **Get** /banks/obp.testing.01/ObpActivity | Get Obp Activity List
[**DynamicEntityGetSingleObpActivityObpTesting01**](ObpActivityObpTesting01Api.md#DynamicEntityGetSingleObpActivityObpTesting01) | **Get** /banks/obp.testing.01/ObpActivity/OBP_ACTIVITY_ID | Get Obp Activity by id
[**DynamicEntityUpdateObpActivityObpTesting01**](ObpActivityObpTesting01Api.md#DynamicEntityUpdateObpActivityObpTesting01) | **Put** /banks/obp.testing.01/ObpActivity/OBP_ACTIVITY_ID | Update Obp Activity


# **DynamicEntityCreateObpActivityObpTesting01**
> InlineResponse20022ObpActivityList DynamicEntityCreateObpActivityObpTesting01(ctx, body)
Create new Obp Activity

<p>Create new Obp Activity.</p><p>Description of this entity, can be markdown text.</p><p><strong>Property List:</strong></p><ul><li>name: description of <strong>name</strong> field, can be markdown text.</li><li>username: description of <strong>name</strong> field, can be markdown text.</li><li>created_date: description of <strong>number</strong> field, can be markdown text.</li></ul><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;ObpActivity&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body33**](Body33.md)| JObject object that needs to be added. | 

### Return type

[**InlineResponse20022ObpActivityList**](inline_response_200_22_obp_activity_list.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityDeleteObpActivityObpTesting01**
> InlineResponse20022ObpActivityList DynamicEntityDeleteObpActivityObpTesting01(ctx, body)
Delete Obp Activity by id

<p>Delete Obp Activity by id</p><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;ObpActivity&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body35**](Body35.md)| JObject object that needs to be added. | 

### Return type

[**InlineResponse20022ObpActivityList**](inline_response_200_22_obp_activity_list.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityGetObpActivityListObpTesting01**
> InlineResponse20022 DynamicEntityGetObpActivityListObpTesting01(ctx, )
Get Obp Activity List

<p>Get Obp Activity List.</p><p>Description of this entity, can be markdown text.</p><p><strong>Property List:</strong></p><ul><li>name: description of <strong>name</strong> field, can be markdown text.</li><li>username: description of <strong>name</strong> field, can be markdown text.</li><li>created_date: description of <strong>number</strong> field, can be markdown text.</li></ul><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;ObpActivity&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p><p>Can do filter on the fields<br />e.g: /ObpActivity?name=James%20Brown&amp;number=123.456&amp;number=11.11<br />Will do filter by this rule: name == &quot;James Brown&quot; &amp;&amp; (number==123.456 || number=11.11)</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityGetSingleObpActivityObpTesting01**
> InlineResponse20022ObpActivityList DynamicEntityGetSingleObpActivityObpTesting01(ctx, )
Get Obp Activity by id

<p>Get Obp Activity by id.</p><p>Description of this entity, can be markdown text.</p><p><strong>Property List:</strong></p><ul><li>name: description of <strong>name</strong> field, can be markdown text.</li><li>username: description of <strong>name</strong> field, can be markdown text.</li><li>created_date: description of <strong>number</strong> field, can be markdown text.</li></ul><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;ObpActivity&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20022ObpActivityList**](inline_response_200_22_obp_activity_list.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityUpdateObpActivityObpTesting01**
> InlineResponse20022ObpActivityList DynamicEntityUpdateObpActivityObpTesting01(ctx, body)
Update Obp Activity

<p>Update Obp Activity.</p><p>Description of this entity, can be markdown text.</p><p><strong>Property List:</strong></p><ul><li>name: description of <strong>name</strong> field, can be markdown text.</li><li>username: description of <strong>name</strong> field, can be markdown text.</li><li>created_date: description of <strong>number</strong> field, can be markdown text.</li></ul><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;ObpActivity&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body34**](Body34.md)| JObject object that needs to be added. | 

### Return type

[**InlineResponse20022ObpActivityList**](inline_response_200_22_obp_activity_list.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

