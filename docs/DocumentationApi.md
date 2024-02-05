# \DocumentationApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiGlossary**](DocumentationApi.md#GetApiGlossary) | **Get** /obp/v5.1.0/api/glossary | Get Glossary of the API
[**GetBankLevelDynamicResourceDocsObp**](DocumentationApi.md#GetBankLevelDynamicResourceDocsObp) | **Get** /obp/v5.1.0/banks/{BANK_ID}/resource-docs/{API_VERSION}/obp | Get Bank Level Dynamic Resource Docs.
[**GetMessageDocs**](DocumentationApi.md#GetMessageDocs) | **Get** /obp/v5.1.0/message-docs/CONNECTOR | Get Message Docs
[**GetMessageDocsSwagger**](DocumentationApi.md#GetMessageDocsSwagger) | **Get** /obp/v5.1.0/message-docs/CONNECTOR/swagger2.0 | Get Message Docs Swagger
[**GetResourceDocsObpV400**](DocumentationApi.md#GetResourceDocsObpV400) | **Get** /obp/v5.1.0/resource-docs/{API_VERSION}/obp | Get Resource Docs
[**GetResourceDocsSwagger**](DocumentationApi.md#GetResourceDocsSwagger) | **Get** /obp/v5.1.0/resource-docs/{API_VERSION}/swagger | Get Swagger documentation
[**GetScannedApiVersions**](DocumentationApi.md#GetScannedApiVersions) | **Get** /obp/v5.1.0/api/versions | Get scanned API Versions


# **GetApiGlossary**
> GlossaryItemsJsonV300 GetApiGlossary(ctx, )
Get Glossary of the API

<p>Get API Glossary</p><p>Returns the glossary of the API</p><p>Authentication is Optional</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GlossaryItemsJsonV300**](GlossaryItemsJsonV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBankLevelDynamicResourceDocsObp**
> ResourceDocsJson GetBankLevelDynamicResourceDocsObp(ctx, body, aPIVERSION, bANKID)
Get Bank Level Dynamic Resource Docs.

<p>Get documentation about the RESTful resources on this server including example bodies for POST and PUT requests.</p><p>This is the native data format used to document OBP endpoints. Each endpoint has a Resource Doc (a Scala case class) defined in the source code.</p><p>This endpoint is used by OBP API Explorer to display and work with the API documentation.</p><p>Most (but not all) fields are also available in swagger format. (The Swagger endpoint is built from Resource Docs.)</p><p>API_VERSION is the version you want documentation about e.g. v3.0.0</p><p>You may filter this endpoint with tags parameter e.g. ?tags=Account,Bank</p><p>You may filter this endpoint with functions parameter e.g. ?functions=enableDisableConsumers,getConnectorMetrics</p><p>For possible function values, see implemented_by.function in the JSON returned by this endpoint or the OBP source code or the footer of the API Explorer which produces a comma separated list of functions that reflect the server or filtering by API Explorer based on tags etc.</p><p>You may filter this endpoint using the 'content' url parameter, e.g. ?content=dynamic<br />if set content=dynamic, only show dynamic endpoints, if content=static, only show the static endpoints. if omit this parameter, we will show all the endpoints.</p><p>You may need some other language resource docs, now we support en_GB and es_ES at the moment.</p><p>You can filter with api-collection-id, but api-collection-id can not be used with others together. If api-collection-id is used in URL, it will ignore all other parameters.</p><p>See the Resource Doc endpoint for more information.</p><p>Note: Dynamic Resource Docs are cached, TTL is 3600 seconds<br />Static Resource Docs are cached, TTL is 3600 seconds</p><p>Following are more examples:<br /><a href=\"https://apisandbox.openbankproject.com/obp/v4.0.0/banks/BANK_ID/resource-docs/v4.0.0/obp\">https://apisandbox.openbankproject.com/obp/v4.0.0/banks/BANK_ID/resource-docs/v4.0.0/obp</a><br /><a href=\"https://apisandbox.openbankproject.com/obp/v4.0.0/banks/BANK_ID/resource-docs/v4.0.0/obp?tags=Account,Bank\">https://apisandbox.openbankproject.com/obp/v4.0.0/banks/BANK_ID/resource-docs/v4.0.0/obp?tags=Account,Bank</a><br /><a href=\"https://apisandbox.openbankproject.com/obp/v4.0.0/banks/BANK_ID/resource-docs/v4.0.0/obp?functions=getBanks,bankById\">https://apisandbox.openbankproject.com/obp/v4.0.0/banks/BANK_ID/resource-docs/v4.0.0/obp?functions=getBanks,bankById</a><br /><a href=\"https://apisandbox.openbankproject.com/obp/v4.0.0/banks/BANK_ID/resource-docs/v4.0.0/obp?locale=es_ES\">https://apisandbox.openbankproject.com/obp/v4.0.0/banks/BANK_ID/resource-docs/v4.0.0/obp?locale=es_ES</a><br /><a href=\"https://apisandbox.openbankproject.com/obp/v4.0.0/banks/BANK_ID/resource-docs/v4.0.0/obp?content=static,dynamic,all\">https://apisandbox.openbankproject.com/obp/v4.0.0/banks/BANK_ID/resource-docs/v4.0.0/obp?content=static,dynamic,all</a><br /><a href=\"https://apisandbox.openbankproject.com/obp/v4.0.0/banks/BANK_ID/resource-docs/v4.0.0/obp?api-collection-id=4e866c86-60c3-4268-a221-cb0bbf1ad221\">https://apisandbox.openbankproject.com/obp/v4.0.0/banks/BANK_ID/resource-docs/v4.0.0/obp?api-collection-id=4e866c86-60c3-4268-a221-cb0bbf1ad221</a></p><ul><li> operation_id is concatenation of \"v\", version and function and should be unique (used for DOM element IDs etc. maybe used to link to source code) </li><li> version references the version that the API call is defined in.</li><li> function is the (scala) partial function that implements this endpoint. It is unique per version of the API.</li><li> request_url is empty for the root call, else the path. It contains the standard prefix (e.g. /obp) and the implemented version (the version where this endpoint was defined) e.g. /obp/v1.2.0/resource</li><li> specified_url (recommended to use) is empty for the root call, else the path. It contains the standard prefix (e.g. /obp) and the version specified in the call e.g. /obp/v3.1.0/resource. In OBP, endpoints are first made available at the request_url, but the same resource (function call) is often made available under later versions (specified_url). To access the latest version of all endpoints use the latest version available on your OBP instance e.g. /obp/v3.1.0 - To get the original version use the request_url. We recommend to use the specified_url since non semantic improvements are more likely to be applied to later implementations of the call.</li><li> summary is a short description inline with the swagger terminology. </li><li> description may contain html markup (generated from markdown on the server).</li></ul><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **aPIVERSION** | **string**| eg:v2.2.0, v3.0.0 | 
  **bANKID** | **string**| The bank id | 

### Return type

[**ResourceDocsJson**](ResourceDocsJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMessageDocs**
> MessageDocsJson GetMessageDocs(ctx, body)
Get Message Docs

<p>These message docs provide example messages sent by OBP to the (Kafka) message queue for processing by the Core Banking / Payment system Adapter - together with an example expected response and possible error codes.<br />Integrators can use these messages to build Adapters that provide core banking services to OBP.</p><p>Note: API Explorer provides a Message Docs page where these messages are displayed.</p><p><code>CONNECTOR</code>: kafka_vSept2018, stored_procedure_vDec2019 ...</p><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 

### Return type

[**MessageDocsJson**](MessageDocsJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMessageDocsSwagger**
> MessageDocsJson GetMessageDocsSwagger(ctx, )
Get Message Docs Swagger

<p>This endpoint provides example message docs in swagger format.<br />It is only relavent for REST Connectors.</p><p>This endpoint can be used by the developer building a REST Adapter that connects to the Core Banking System (CBS).<br />That is, the Adapter developer can use the Swagger surfaced here to build the REST APIs that the OBP REST connector will call to consume CBS services.</p><p>i.e.:</p><p>OBP API (Core OBP API code) -&gt; OBP REST Connector (OBP REST Connector code) -&gt; OBP REST Adapter (Adapter developer code) -&gt; CBS (Main Frame)</p><p>Authentication is Optional</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MessageDocsJson**](MessageDocsJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResourceDocsObpV400**
> ResourceDocsJson GetResourceDocsObpV400(ctx, body, aPIVERSION)
Get Resource Docs

<p>Get documentation about the RESTful resources on this server including example bodies for POST and PUT requests.</p><p>This is the native data format used to document OBP endpoints. Each endpoint has a Resource Doc (a Scala case class) defined in the source code.</p><p>This endpoint is used by OBP API Explorer to display and work with the API documentation.</p><p>Most (but not all) fields are also available in swagger format. (The Swagger endpoint is built from Resource Docs.)</p><p>API_VERSION is the version you want documentation about e.g. v3.0.0</p><p>You may filter this endpoint with tags parameter e.g. ?tags=Account,Bank</p><p>You may filter this endpoint with functions parameter e.g. ?functions=enableDisableConsumers,getConnectorMetrics</p><p>For possible function values, see implemented_by.function in the JSON returned by this endpoint or the OBP source code or the footer of the API Explorer which produces a comma separated list of functions that reflect the server or filtering by API Explorer based on tags etc.</p><p>You may filter this endpoint using the 'content' url parameter, e.g. ?content=dynamic<br />if set content=dynamic, only show dynamic endpoints, if content=static, only show the static endpoints. if omit this parameter, we will show all the endpoints.</p><p>You may need some other language resource docs, now we support en_GB and es_ES at the moment.</p><p>You can filter with api-collection-id, but api-collection-id can not be used with others together. If api-collection-id is used in URL, it will ignore all other parameters.</p><p>See the Resource Doc endpoint for more information.</p><p>Note: Dynamic Resource Docs are cached, TTL is 3600 seconds<br />Static Resource Docs are cached, TTL is 3600 seconds</p><p>Following are more examples:<br /><a href=\"https://apisandbox.openbankproject.com/obp/v4.0.0/resource-docs/v4.0.0/obp\">https://apisandbox.openbankproject.com/obp/v4.0.0/resource-docs/v4.0.0/obp</a><br /><a href=\"https://apisandbox.openbankproject.com/obp/v4.0.0/resource-docs/v4.0.0/obp?tags=Account,Bank\">https://apisandbox.openbankproject.com/obp/v4.0.0/resource-docs/v4.0.0/obp?tags=Account,Bank</a><br /><a href=\"https://apisandbox.openbankproject.com/obp/v4.0.0/resource-docs/v4.0.0/obp?functions=getBanks,bankById\">https://apisandbox.openbankproject.com/obp/v4.0.0/resource-docs/v4.0.0/obp?functions=getBanks,bankById</a><br /><a href=\"https://apisandbox.openbankproject.com/obp/v4.0.0/resource-docs/v4.0.0/obp?locale=es_ES\">https://apisandbox.openbankproject.com/obp/v4.0.0/resource-docs/v4.0.0/obp?locale=es_ES</a><br /><a href=\"https://apisandbox.openbankproject.com/obp/v4.0.0/resource-docs/v4.0.0/obp?content=static,dynamic,all\">https://apisandbox.openbankproject.com/obp/v4.0.0/resource-docs/v4.0.0/obp?content=static,dynamic,all</a><br /><a href=\"https://apisandbox.openbankproject.com/obp/v4.0.0/resource-docs/v4.0.0/obp?api-collection-id=4e866c86-60c3-4268-a221-cb0bbf1ad221\">https://apisandbox.openbankproject.com/obp/v4.0.0/resource-docs/v4.0.0/obp?api-collection-id=4e866c86-60c3-4268-a221-cb0bbf1ad221</a></p><ul><li> operation_id is concatenation of \"v\", version and function and should be unique (used for DOM element IDs etc. maybe used to link to source code) </li><li> version references the version that the API call is defined in.</li><li> function is the (scala) partial function that implements this endpoint. It is unique per version of the API.</li><li> request_url is empty for the root call, else the path. It contains the standard prefix (e.g. /obp) and the implemented version (the version where this endpoint was defined) e.g. /obp/v1.2.0/resource</li><li> specified_url (recommended to use) is empty for the root call, else the path. It contains the standard prefix (e.g. /obp) and the version specified in the call e.g. /obp/v3.1.0/resource. In OBP, endpoints are first made available at the request_url, but the same resource (function call) is often made available under later versions (specified_url). To access the latest version of all endpoints use the latest version available on your OBP instance e.g. /obp/v3.1.0 - To get the original version use the request_url. We recommend to use the specified_url since non semantic improvements are more likely to be applied to later implementations of the call.</li><li> summary is a short description inline with the swagger terminology. </li><li> description may contain html markup (generated from markdown on the server).</li></ul><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **aPIVERSION** | **string**| eg:v2.2.0, v3.0.0 | 

### Return type

[**ResourceDocsJson**](ResourceDocsJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResourceDocsSwagger**
> EmptyClassJson GetResourceDocsSwagger(ctx, body, aPIVERSION)
Get Swagger documentation

<p>Returns documentation about the RESTful resources on this server in Swagger format.</p><p>API_VERSION is the version you want documentation about e.g. v3.0.0</p><p>You may filter this endpoint using the 'tags' url parameter e.g. ?tags=Account,Bank</p><p>(All endpoints are given one or more tags which for used in grouping)</p><p>You may filter this endpoint using the 'functions' url parameter e.g. ?functions=getBanks,bankById</p><p>(Each endpoint is implemented in the OBP Scala code by a 'function')</p><p>See the Resource Doc endpoint for more information.</p><p>Note: Resource Docs are cached, TTL is 3600 seconds</p><p>Following are more examples:<br /><a href=\"https://apisandbox.openbankproject.com/obp/v3.1.0/resource-docs/v3.1.0/swagger\">https://apisandbox.openbankproject.com/obp/v3.1.0/resource-docs/v3.1.0/swagger</a><br /><a href=\"https://apisandbox.openbankproject.com/obp/v3.1.0/resource-docs/v3.1.0/swagger?tags=Account,Bank\">https://apisandbox.openbankproject.com/obp/v3.1.0/resource-docs/v3.1.0/swagger?tags=Account,Bank</a><br /><a href=\"https://apisandbox.openbankproject.com/obp/v3.1.0/resource-docs/v3.1.0/swagger?functions=getBanks,bankById\">https://apisandbox.openbankproject.com/obp/v3.1.0/resource-docs/v3.1.0/swagger?functions=getBanks,bankById</a><br /><a href=\"https://apisandbox.openbankproject.com/obp/v3.1.0/resource-docs/v3.1.0/swagger?tags=Account,Bank,PSD2&amp;functions=getBanks,bankById\">https://apisandbox.openbankproject.com/obp/v3.1.0/resource-docs/v3.1.0/swagger?tags=Account,Bank,PSD2&amp;functions=getBanks,bankById</a></p><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **aPIVERSION** | **string**| eg:v2.2.0, v3.0.0 | 

### Return type

[**EmptyClassJson**](EmptyClassJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetScannedApiVersions**
> InlineResponse200 GetScannedApiVersions(ctx, )
Get scanned API Versions

<p>Get all the scanned API Versions.</p><p>Authentication is Optional</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

