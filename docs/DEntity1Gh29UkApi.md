# \DEntity1Gh29UkApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DynamicEntityCreateDEntity1Gh29Uk**](DEntity1Gh29UkApi.md#DynamicEntityCreateDEntity1Gh29Uk) | **Post** /banks/gh.29.uk/D-Entity1 | Create new D Entity1
[**DynamicEntityDeleteDEntity1Gh29Uk**](DEntity1Gh29UkApi.md#DynamicEntityDeleteDEntity1Gh29Uk) | **Delete** /banks/gh.29.uk/D-Entity1/D-ENTITY1_ID | Delete D Entity1 by id
[**DynamicEntityGetDEntity1ListGh29Uk**](DEntity1Gh29UkApi.md#DynamicEntityGetDEntity1ListGh29Uk) | **Get** /banks/gh.29.uk/D-Entity1 | Get D Entity1 List
[**DynamicEntityGetSingleDEntity1Gh29Uk**](DEntity1Gh29UkApi.md#DynamicEntityGetSingleDEntity1Gh29Uk) | **Get** /banks/gh.29.uk/D-Entity1/D-ENTITY1_ID | Get D Entity1 by id
[**DynamicEntityUpdateDEntity1Gh29Uk**](DEntity1Gh29UkApi.md#DynamicEntityUpdateDEntity1Gh29Uk) | **Put** /banks/gh.29.uk/D-Entity1/D-ENTITY1_ID | Update D Entity1


# **DynamicEntityCreateDEntity1Gh29Uk**
> InlineResponse20017Dentity1List DynamicEntityCreateDEntity1Gh29Uk(ctx, body)
Create new D Entity1

<p>Create new D Entity1.</p><p>Description of this entity, can be markdown text.</p><p><strong>Property List:</strong></p><ul><li>name: description of <strong>name</strong> field, can be markdown text.</li><li>number: description of <strong>number</strong> field, can be markdown text.</li></ul><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;D-Entity1&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body18**](Body18.md)| JObject object that needs to be added. | 

### Return type

[**InlineResponse20017Dentity1List**](inline_response_200_17_dentity1_list.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityDeleteDEntity1Gh29Uk**
> InlineResponse20017Dentity1List DynamicEntityDeleteDEntity1Gh29Uk(ctx, body)
Delete D Entity1 by id

<p>Delete D Entity1 by id</p><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;D-Entity1&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body20**](Body20.md)| JObject object that needs to be added. | 

### Return type

[**InlineResponse20017Dentity1List**](inline_response_200_17_dentity1_list.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityGetDEntity1ListGh29Uk**
> InlineResponse20017 DynamicEntityGetDEntity1ListGh29Uk(ctx, )
Get D Entity1 List

<p>Get D Entity1 List.</p><p>Description of this entity, can be markdown text.</p><p><strong>Property List:</strong></p><ul><li>name: description of <strong>name</strong> field, can be markdown text.</li><li>number: description of <strong>number</strong> field, can be markdown text.</li></ul><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;D-Entity1&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p><p>Can do filter on the fields<br />e.g: /D-Entity1?name=James%20Brown&amp;number=123.456&amp;number=11.11<br />Will do filter by this rule: name == &quot;James Brown&quot; &amp;&amp; (number==123.456 || number=11.11)</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityGetSingleDEntity1Gh29Uk**
> InlineResponse20017Dentity1List DynamicEntityGetSingleDEntity1Gh29Uk(ctx, )
Get D Entity1 by id

<p>Get D Entity1 by id.</p><p>Description of this entity, can be markdown text.</p><p><strong>Property List:</strong></p><ul><li>name: description of <strong>name</strong> field, can be markdown text.</li><li>number: description of <strong>number</strong> field, can be markdown text.</li></ul><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;D-Entity1&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20017Dentity1List**](inline_response_200_17_dentity1_list.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityUpdateDEntity1Gh29Uk**
> InlineResponse20017Dentity1List DynamicEntityUpdateDEntity1Gh29Uk(ctx, body)
Update D Entity1

<p>Update D Entity1.</p><p>Description of this entity, can be markdown text.</p><p><strong>Property List:</strong></p><ul><li>name: description of <strong>name</strong> field, can be markdown text.</li><li>number: description of <strong>number</strong> field, can be markdown text.</li></ul><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;D-Entity1&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body19**](Body19.md)| JObject object that needs to be added. | 

### Return type

[**InlineResponse20017Dentity1List**](inline_response_200_17_dentity1_list.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

