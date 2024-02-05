# \InsurancePolicyGh29UkApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DynamicEntityCreateInsurancePolicyGh29Uk**](InsurancePolicyGh29UkApi.md#DynamicEntityCreateInsurancePolicyGh29Uk) | **Post** /banks/gh.29.uk/InsurancePolicy | Create new Insurance Policy
[**DynamicEntityDeleteInsurancePolicyGh29Uk**](InsurancePolicyGh29UkApi.md#DynamicEntityDeleteInsurancePolicyGh29Uk) | **Delete** /banks/gh.29.uk/InsurancePolicy/INSURANCE_POLICY_ID | Delete Insurance Policy by id
[**DynamicEntityGetInsurancePolicyListGh29Uk**](InsurancePolicyGh29UkApi.md#DynamicEntityGetInsurancePolicyListGh29Uk) | **Get** /banks/gh.29.uk/InsurancePolicy | Get Insurance Policy List
[**DynamicEntityGetSingleInsurancePolicyGh29Uk**](InsurancePolicyGh29UkApi.md#DynamicEntityGetSingleInsurancePolicyGh29Uk) | **Get** /banks/gh.29.uk/InsurancePolicy/INSURANCE_POLICY_ID | Get Insurance Policy by id
[**DynamicEntityUpdateInsurancePolicyGh29Uk**](InsurancePolicyGh29UkApi.md#DynamicEntityUpdateInsurancePolicyGh29Uk) | **Put** /banks/gh.29.uk/InsurancePolicy/INSURANCE_POLICY_ID | Update Insurance Policy


# **DynamicEntityCreateInsurancePolicyGh29Uk**
> InlineResponse20018InsurancePolicyList DynamicEntityCreateInsurancePolicyGh29Uk(ctx, body)
Create new Insurance Policy

<p>Create new Insurance Policy.</p><p>Which insurance policies the customer already has.</p><p><strong>Property List:</strong></p><ul><li>name: description of <strong>name</strong> field, can be markdown text.</li><li>number: description of <strong>number</strong> field, can be markdown text.</li></ul><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;InsurancePolicy&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body21**](Body21.md)| JObject object that needs to be added. | 

### Return type

[**InlineResponse20018InsurancePolicyList**](inline_response_200_18_insurance_policy_list.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityDeleteInsurancePolicyGh29Uk**
> InlineResponse20018InsurancePolicyList DynamicEntityDeleteInsurancePolicyGh29Uk(ctx, body)
Delete Insurance Policy by id

<p>Delete Insurance Policy by id</p><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;InsurancePolicy&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body23**](Body23.md)| JObject object that needs to be added. | 

### Return type

[**InlineResponse20018InsurancePolicyList**](inline_response_200_18_insurance_policy_list.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityGetInsurancePolicyListGh29Uk**
> InlineResponse20018 DynamicEntityGetInsurancePolicyListGh29Uk(ctx, )
Get Insurance Policy List

<p>Get Insurance Policy List.</p><p>Which insurance policies the customer already has.</p><p><strong>Property List:</strong></p><ul><li>name: description of <strong>name</strong> field, can be markdown text.</li><li>number: description of <strong>number</strong> field, can be markdown text.</li></ul><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;InsurancePolicy&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p><p>Can do filter on the fields<br />e.g: /InsurancePolicy?name=James%20Brown&amp;number=123.456&amp;number=11.11<br />Will do filter by this rule: name == &quot;James Brown&quot; &amp;&amp; (number==123.456 || number=11.11)</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityGetSingleInsurancePolicyGh29Uk**
> InlineResponse20018InsurancePolicyList DynamicEntityGetSingleInsurancePolicyGh29Uk(ctx, )
Get Insurance Policy by id

<p>Get Insurance Policy by id.</p><p>Which insurance policies the customer already has.</p><p><strong>Property List:</strong></p><ul><li>name: description of <strong>name</strong> field, can be markdown text.</li><li>number: description of <strong>number</strong> field, can be markdown text.</li></ul><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;InsurancePolicy&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20018InsurancePolicyList**](inline_response_200_18_insurance_policy_list.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicEntityUpdateInsurancePolicyGh29Uk**
> InlineResponse20018InsurancePolicyList DynamicEntityUpdateInsurancePolicyGh29Uk(ctx, body)
Update Insurance Policy

<p>Update Insurance Policy.</p><p>Which insurance policies the customer already has.</p><p><strong>Property List:</strong></p><ul><li>name: description of <strong>name</strong> field, can be markdown text.</li><li>number: description of <strong>number</strong> field, can be markdown text.</li></ul><p>MethodRouting settings example:</p><details><pre><code>{  &quot;is_bank_id_exact_match&quot;:false,  &quot;method_name&quot;:&quot;dynamicEntityProcess&quot;,  &quot;connector_name&quot;:&quot;rest_vMar2019&quot;,  &quot;bank_id_pattern&quot;:&quot;.*&quot;,  &quot;parameters&quot;:[    {        &quot;key&quot;:&quot;entityName&quot;,        &quot;value&quot;:&quot;InsurancePolicy&quot;    }    {        &quot;key&quot;:&quot;url&quot;,        &quot;value&quot;:&quot;http://mydomain.com/xxx&quot;    }  ]}</code></pre></details><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body22**](Body22.md)| JObject object that needs to be added. | 

### Return type

[**InlineResponse20018InsurancePolicyList**](inline_response_200_18_insurance_policy_list.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

