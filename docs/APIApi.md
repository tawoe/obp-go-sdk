# \APIApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Config**](APIApi.md#Config) | **Get** /obp/v5.1.0/config | Get API Configuration
[**CreateBankLevelDynamicEndpoint**](APIApi.md#CreateBankLevelDynamicEndpoint) | **Post** /obp/v5.1.0/management/banks/{BANK_ID}/dynamic-endpoints | Create Bank Level Dynamic Endpoint
[**CreateBankLevelDynamicEntity**](APIApi.md#CreateBankLevelDynamicEntity) | **Post** /obp/v5.1.0/management/banks/{BANK_ID}/dynamic-entities | Create Bank Level Dynamic Entity
[**CreateBankLevelEndpointTag**](APIApi.md#CreateBankLevelEndpointTag) | **Post** /obp/v5.1.0/management/banks/{BANK_ID}/endpoints/OPERATION_ID/tags | Create Bank Level Endpoint Tag
[**CreateDynamicEndpoint**](APIApi.md#CreateDynamicEndpoint) | **Post** /obp/v5.1.0/management/dynamic-endpoints | Create Dynamic Endpoint
[**CreateMethodRouting**](APIApi.md#CreateMethodRouting) | **Post** /obp/v5.1.0/management/method_routings | Create MethodRouting
[**CreateRegulatedEntity**](APIApi.md#CreateRegulatedEntity) | **Post** /obp/v5.1.0/regulated-entities | Create Regulated Entity
[**CreateSystemDynamicEntity**](APIApi.md#CreateSystemDynamicEntity) | **Post** /obp/v5.1.0/management/system-dynamic-entities | Create System Level Dynamic Entity
[**CreateSystemLevelEndpointTag**](APIApi.md#CreateSystemLevelEndpointTag) | **Post** /obp/v5.1.0/management/endpoints/OPERATION_ID/tags | Create System Level Endpoint Tag
[**DeleteBankLevelDynamicEndpoint**](APIApi.md#DeleteBankLevelDynamicEndpoint) | **Delete** /obp/v5.1.0/management/banks/{BANK_ID}/dynamic-endpoints/DYNAMIC_ENDPOINT_ID |  Delete Bank Level Dynamic Endpoint
[**DeleteBankLevelDynamicEntity**](APIApi.md#DeleteBankLevelDynamicEntity) | **Delete** /obp/v5.1.0/management/banks/{BANK_ID}/dynamic-entities/{DYNAMIC_ENTITY_ID} | Delete Bank Level Dynamic Entity
[**DeleteBankLevelEndpointTag**](APIApi.md#DeleteBankLevelEndpointTag) | **Delete** /obp/v5.1.0/management/banks/{BANK_ID}/endpoints/OPERATION_ID/tags/ENDPOINT_TAG_ID | Delete Bank Level Endpoint Tag
[**DeleteDynamicEndpoint**](APIApi.md#DeleteDynamicEndpoint) | **Delete** /obp/v5.1.0/management/dynamic-endpoints/DYNAMIC_ENDPOINT_ID |  Delete Dynamic Endpoint
[**DeleteMethodRouting**](APIApi.md#DeleteMethodRouting) | **Delete** /obp/v5.1.0/management/method_routings/{METHOD_ROUTING_ID} | Delete MethodRouting
[**DeleteMyDynamicEndpoint**](APIApi.md#DeleteMyDynamicEndpoint) | **Delete** /obp/v5.1.0/my/dynamic-endpoints/DYNAMIC_ENDPOINT_ID | Delete My Dynamic Endpoint
[**DeleteMyDynamicEntity**](APIApi.md#DeleteMyDynamicEntity) | **Delete** /obp/v5.1.0/my/dynamic-entities/{DYNAMIC_ENTITY_ID} | Delete My Dynamic Entity
[**DeleteRegulatedEntity**](APIApi.md#DeleteRegulatedEntity) | **Delete** /obp/v5.1.0/regulated-entities/REGULATED_ENTITY_ID | Delete Regulated Entity
[**DeleteSystemDynamicEntity**](APIApi.md#DeleteSystemDynamicEntity) | **Delete** /obp/v5.1.0/management/system-dynamic-entities/{DYNAMIC_ENTITY_ID} | Delete System Level Dynamic Entity
[**DeleteSystemLevelEndpointTag**](APIApi.md#DeleteSystemLevelEndpointTag) | **Delete** /obp/v5.1.0/management/endpoints/OPERATION_ID/tags/ENDPOINT_TAG_ID | Delete System Level Endpoint Tag
[**ElasticSearchMetrics**](APIApi.md#ElasticSearchMetrics) | **Get** /obp/v5.1.0/search/metrics | Search API Metrics via Elasticsearch
[**GetAdapterInfo**](APIApi.md#GetAdapterInfo) | **Get** /obp/v5.1.0/adapter | Get Adapter Info
[**GetAdapterInfoForBank**](APIApi.md#GetAdapterInfoForBank) | **Get** /obp/v5.1.0/banks/{BANK_ID}/adapter | Get Adapter Info for a bank
[**GetBankLevelDynamicEndpoint**](APIApi.md#GetBankLevelDynamicEndpoint) | **Get** /obp/v5.1.0/management/banks/{BANK_ID}/dynamic-endpoints/DYNAMIC_ENDPOINT_ID |  Get Bank Level Dynamic Endpoint
[**GetBankLevelDynamicEndpoints**](APIApi.md#GetBankLevelDynamicEndpoints) | **Get** /obp/v5.1.0/management/banks/{BANK_ID}/dynamic-endpoints | Get Bank Level Dynamic Endpoints
[**GetBankLevelDynamicEntities**](APIApi.md#GetBankLevelDynamicEntities) | **Get** /obp/v5.1.0/management/banks/{BANK_ID}/dynamic-entities | Get Bank Level Dynamic Entities
[**GetBankLevelDynamicResourceDocsObp**](APIApi.md#GetBankLevelDynamicResourceDocsObp) | **Get** /obp/v5.1.0/banks/{BANK_ID}/resource-docs/{API_VERSION}/obp | Get Bank Level Dynamic Resource Docs.
[**GetBankLevelEndpointTags**](APIApi.md#GetBankLevelEndpointTags) | **Get** /obp/v5.1.0/management/banks/{BANK_ID}/endpoints/OPERATION_ID/tags | Get Bank Level Endpoint Tags
[**GetCallContext**](APIApi.md#GetCallContext) | **Get** /obp/v5.1.0/development/call_context | Get the Call Context of a current call
[**GetConnectorMetrics**](APIApi.md#GetConnectorMetrics) | **Get** /obp/v5.1.0/management/connector/metrics | Get Connector Metrics
[**GetDynamicEndpoint**](APIApi.md#GetDynamicEndpoint) | **Get** /obp/v5.1.0/management/dynamic-endpoints/DYNAMIC_ENDPOINT_ID | Get Dynamic Endpoint
[**GetDynamicEndpoints**](APIApi.md#GetDynamicEndpoints) | **Get** /obp/v5.1.0/management/dynamic-endpoints |  Get Dynamic Endpoints
[**GetMapperDatabaseInfo**](APIApi.md#GetMapperDatabaseInfo) | **Get** /obp/v5.1.0/database/info | Get Mapper Database Info
[**GetMessageDocs**](APIApi.md#GetMessageDocs) | **Get** /obp/v5.1.0/message-docs/CONNECTOR | Get Message Docs
[**GetMessageDocsSwagger**](APIApi.md#GetMessageDocsSwagger) | **Get** /obp/v5.1.0/message-docs/CONNECTOR/swagger2.0 | Get Message Docs Swagger
[**GetMethodRoutings**](APIApi.md#GetMethodRoutings) | **Get** /obp/v5.1.0/management/method_routings | Get MethodRoutings
[**GetMetrics**](APIApi.md#GetMetrics) | **Get** /obp/v5.1.0/management/metrics | Get Metrics
[**GetMetricsAtBank**](APIApi.md#GetMetricsAtBank) | **Get** /obp/v5.1.0/management/metrics/banks/{BANK_ID} | Get Metrics at Bank
[**GetMyDynamicEndpoints**](APIApi.md#GetMyDynamicEndpoints) | **Get** /obp/v5.1.0/my/dynamic-endpoints | Get My Dynamic Endpoints
[**GetMyDynamicEntities**](APIApi.md#GetMyDynamicEntities) | **Get** /obp/v5.1.0/my/dynamic-entities | Get My Dynamic Entities
[**GetOAuth2ServerJWKsURIs**](APIApi.md#GetOAuth2ServerJWKsURIs) | **Get** /obp/v5.1.0/jwks-uris | Get JSON Web Key (JWK) URIs
[**GetObpConnectorLoopback**](APIApi.md#GetObpConnectorLoopback) | **Get** /obp/v5.1.0/connector/loopback | Get Connector Status (Loopback)
[**GetRateLimitingInfo**](APIApi.md#GetRateLimitingInfo) | **Get** /obp/v5.1.0/rate-limiting | Get Rate Limiting Info
[**GetRegulatedEntityById**](APIApi.md#GetRegulatedEntityById) | **Get** /obp/v5.1.0/regulated-entities/REGULATED_ENTITY_ID | Get Regulated Entity
[**GetResourceDocsObpV400**](APIApi.md#GetResourceDocsObpV400) | **Get** /obp/v5.1.0/resource-docs/{API_VERSION}/obp | Get Resource Docs
[**GetResourceDocsSwagger**](APIApi.md#GetResourceDocsSwagger) | **Get** /obp/v5.1.0/resource-docs/{API_VERSION}/swagger | Get Swagger documentation
[**GetScannedApiVersions**](APIApi.md#GetScannedApiVersions) | **Get** /obp/v5.1.0/api/versions | Get scanned API Versions
[**GetServerJWK**](APIApi.md#GetServerJWK) | **Get** /obp/v5.1.0/certs | Get JSON Web Key (JWK)
[**GetSystemDynamicEntities**](APIApi.md#GetSystemDynamicEntities) | **Get** /obp/v5.1.0/management/system-dynamic-entities | Get System Dynamic Entities
[**GetSystemLevelEndpointTags**](APIApi.md#GetSystemLevelEndpointTags) | **Get** /obp/v5.1.0/management/endpoints/OPERATION_ID/tags | Get System Level Endpoint Tags
[**RegulatedEntities**](APIApi.md#RegulatedEntities) | **Get** /obp/v5.1.0/regulated-entities | Get Regulated Entities
[**Root**](APIApi.md#Root) | **Get** /obp/v5.1.0/root | Get API Info (root)
[**SuggestedSessionTimeout**](APIApi.md#SuggestedSessionTimeout) | **Get** /obp/v5.1.0/ui/suggested-session-timeout | Get Suggested Session Timeout
[**UpdateBankLevelDynamicEndpointHost**](APIApi.md#UpdateBankLevelDynamicEndpointHost) | **Put** /obp/v5.1.0/management/banks/{BANK_ID}/dynamic-endpoints/DYNAMIC_ENDPOINT_ID/host |  Update Bank Level Dynamic Endpoint Host
[**UpdateBankLevelDynamicEntity**](APIApi.md#UpdateBankLevelDynamicEntity) | **Put** /obp/v5.1.0/management/banks/{BANK_ID}/dynamic-entities/{DYNAMIC_ENTITY_ID} | Update Bank Level Dynamic Entity
[**UpdateBankLevelEndpointTag**](APIApi.md#UpdateBankLevelEndpointTag) | **Put** /obp/v5.1.0/management/banks/{BANK_ID}/endpoints/OPERATION_ID/tags/ENDPOINT_TAG_ID | Update Bank Level Endpoint Tag
[**UpdateDynamicEndpointHost**](APIApi.md#UpdateDynamicEndpointHost) | **Put** /obp/v5.1.0/management/dynamic-endpoints/DYNAMIC_ENDPOINT_ID/host |  Update Dynamic Endpoint Host
[**UpdateMethodRouting**](APIApi.md#UpdateMethodRouting) | **Put** /obp/v5.1.0/management/method_routings/{METHOD_ROUTING_ID} | Update MethodRouting
[**UpdateMyDynamicEntity**](APIApi.md#UpdateMyDynamicEntity) | **Put** /obp/v5.1.0/my/dynamic-entities/{DYNAMIC_ENTITY_ID} | Update My Dynamic Entity
[**UpdateSystemDynamicEntity**](APIApi.md#UpdateSystemDynamicEntity) | **Put** /obp/v5.1.0/management/system-dynamic-entities/{DYNAMIC_ENTITY_ID} | Update System Level Dynamic Entity
[**UpdateSystemLevelEndpointTag**](APIApi.md#UpdateSystemLevelEndpointTag) | **Put** /obp/v5.1.0/management/endpoints/OPERATION_ID/tags/ENDPOINT_TAG_ID | Update System Level Endpoint Tag
[**VerifyRequestSignResponse**](APIApi.md#VerifyRequestSignResponse) | **Get** /obp/v5.1.0/development/echo/jws-verified-request-jws-signed-response | Verify Request and Sign Response of a current call
[**WaitingForGodot**](APIApi.md#WaitingForGodot) | **Get** /obp/v5.1.0/waiting-for-godot | Waiting For Godot


# **Config**
> ConfigurationJson Config(ctx, )
Get API Configuration

<p>Returns information about:</p><ul><li>The default bank_id</li><li>Akka configuration</li><li>Elastic Search configuration</li><li>Cached functions</li></ul><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ConfigurationJson**](ConfigurationJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBankLevelDynamicEndpoint**
> InlineResponse201 CreateBankLevelDynamicEndpoint(ctx, body, bANKID)
Create Bank Level Dynamic Endpoint

<p>Create dynamic endpoints.</p><p>Create dynamic endpoints with one json format swagger content.</p><p>If the host of swagger is <code>dynamic_entity</code>, then you need link the swagger fields to the dynamic entity fields,<br />please check <code>Endpoint Mapping</code> endpoints.</p><p>If the host of swagger is <code>obp_mock</code>, every dynamic endpoint will return example response of swagger,</p><p>when create MethodRouting for given dynamic endpoint, it will be routed to given url.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body**](Body.md)| JObject object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBankLevelDynamicEntity**
> DynamicEntityFooBar CreateBankLevelDynamicEntity(ctx, body, bANKID)
Create Bank Level Dynamic Entity

<p>Create a Bank Level DynamicEntity.</p><p>Authentication is Mandatory</p><p>Create a DynamicEntity. If creation is successful, the corresponding POST, GET, PUT and DELETE (Create, Read, Update, Delete or CRUD for short) endpoints will be generated automatically</p><p>The following field types are as supported:<br />[number, integer, boolean, string, DATE_WITH_DAY, reference]</p><p>The DATE_WITH_DAY format is: yyyy-MM-dd</p><p>Reference types are like foreign keys and composite foreign keys are supported. The value you need to supply as the (composite) foreign key is a UUID (or several UUIDs in the case of a composite key) that match value in another Entity..<br />The following list shows all the possible reference types in the system with corresponding examples values so you can see how to construct each reference type value.</p><pre><code>&quot;someField0&quot;: {    &quot;type&quot;: &quot;reference:FishPort&quot;,    &quot;example&quot;: &quot;e6a6f108-930a-4df8-8c70-68387180b81f&quot;}&quot;someField1&quot;: {    &quot;type&quot;: &quot;reference:FooBar&quot;,    &quot;example&quot;: &quot;e6a6f108-930a-4df8-8c70-68387180b81f&quot;}&quot;someField2&quot;: {    &quot;type&quot;: &quot;reference:sustrans&quot;,    &quot;example&quot;: &quot;e6a6f108-930a-4df8-8c70-68387180b81f&quot;}&quot;someField3&quot;: {    &quot;type&quot;: &quot;reference:SimonCovid&quot;,    &quot;example&quot;: &quot;e6a6f108-930a-4df8-8c70-68387180b81f&quot;}&quot;someField4&quot;: {    &quot;type&quot;: &quot;reference:CovidAPIDays&quot;,    &quot;example&quot;: &quot;e6a6f108-930a-4df8-8c70-68387180b81f&quot;}&quot;someField5&quot;: {    &quot;type&quot;: &quot;reference:customer_cars&quot;,    &quot;example&quot;: &quot;e6a6f108-930a-4df8-8c70-68387180b81f&quot;}&quot;someField6&quot;: {    &quot;type&quot;: &quot;reference:MarchHare&quot;,    &quot;example&quot;: &quot;e6a6f108-930a-4df8-8c70-68387180b81f&quot;}&quot;someField7&quot;: {    &quot;type&quot;: &quot;reference:InsurancePolicy&quot;,    &quot;example&quot;: &quot;e6a6f108-930a-4df8-8c70-68387180b81f&quot;}&quot;someField8&quot;: {    &quot;type&quot;: &quot;reference:Odometer&quot;,    &quot;example&quot;: &quot;e6a6f108-930a-4df8-8c70-68387180b81f&quot;}&quot;someField9&quot;: {    &quot;type&quot;: &quot;reference:InsurancePremium&quot;,    &quot;example&quot;: &quot;e6a6f108-930a-4df8-8c70-68387180b81f&quot;}&quot;someField10&quot;: {    &quot;type&quot;: &quot;reference:ObpActivity&quot;,    &quot;example&quot;: &quot;e6a6f108-930a-4df8-8c70-68387180b81f&quot;}&quot;someField11&quot;: {    &quot;type&quot;: &quot;reference:test1&quot;,    &quot;example&quot;: &quot;e6a6f108-930a-4df8-8c70-68387180b81f&quot;}&quot;someField12&quot;: {    &quot;type&quot;: &quot;reference:D-Entity1&quot;,    &quot;example&quot;: &quot;e6a6f108-930a-4df8-8c70-68387180b81f&quot;}&quot;someField13&quot;: {    &quot;type&quot;: &quot;reference:test_daniel707&quot;,    &quot;example&quot;: &quot;e6a6f108-930a-4df8-8c70-68387180b81f&quot;}&quot;someField14&quot;: {    &quot;type&quot;: &quot;reference:Bank&quot;,    &quot;example&quot;: &quot;e6a6f108-930a-4df8-8c70-68387180b81f&quot;}&quot;someField15&quot;: {    &quot;type&quot;: &quot;reference:Consumer&quot;,    &quot;example&quot;: &quot;e6a6f108-930a-4df8-8c70-68387180b81f&quot;}&quot;someField16&quot;: {    &quot;type&quot;: &quot;reference:Customer&quot;,    &quot;example&quot;: &quot;e6a6f108-930a-4df8-8c70-68387180b81f&quot;}&quot;someField17&quot;: {    &quot;type&quot;: &quot;reference:MethodRouting&quot;,    &quot;example&quot;: &quot;e6a6f108-930a-4df8-8c70-68387180b81f&quot;}&quot;someField18&quot;: {    &quot;type&quot;: &quot;reference:DynamicEntity&quot;,    &quot;example&quot;: &quot;e6a6f108-930a-4df8-8c70-68387180b81f&quot;}&quot;someField19&quot;: {    &quot;type&quot;: &quot;reference:TransactionRequest&quot;,    &quot;example&quot;: &quot;e6a6f108-930a-4df8-8c70-68387180b81f&quot;}&quot;someField20&quot;: {    &quot;type&quot;: &quot;reference:ProductAttribute&quot;,    &quot;example&quot;: &quot;e6a6f108-930a-4df8-8c70-68387180b81f&quot;}&quot;someField21&quot;: {    &quot;type&quot;: &quot;reference:AccountAttribute&quot;,    &quot;example&quot;: &quot;e6a6f108-930a-4df8-8c70-68387180b81f&quot;}&quot;someField22&quot;: {    &quot;type&quot;: &quot;reference:TransactionAttribute&quot;,    &quot;example&quot;: &quot;e6a6f108-930a-4df8-8c70-68387180b81f&quot;}&quot;someField23&quot;: {    &quot;type&quot;: &quot;reference:CustomerAttribute&quot;,    &quot;example&quot;: &quot;e6a6f108-930a-4df8-8c70-68387180b81f&quot;}&quot;someField24&quot;: {    &quot;type&quot;: &quot;reference:AccountApplication&quot;,    &quot;example&quot;: &quot;e6a6f108-930a-4df8-8c70-68387180b81f&quot;}&quot;someField25&quot;: {    &quot;type&quot;: &quot;reference:CardAttribute&quot;,    &quot;example&quot;: &quot;e6a6f108-930a-4df8-8c70-68387180b81f&quot;}&quot;someField26&quot;: {    &quot;type&quot;: &quot;reference:Counterparty&quot;,    &quot;example&quot;: &quot;e6a6f108-930a-4df8-8c70-68387180b81f&quot;}&quot;someField27&quot;: {    &quot;type&quot;: &quot;reference:Branch:bankId&amp;branchId&quot;,    &quot;example&quot;: &quot;bankId=e6a6f108-930a-4df8-8c70-68387180b81f&amp;branchId=49d57c01-59e8-40b2-b2bd-7acc1b0c3645&quot;}&quot;someField28&quot;: {    &quot;type&quot;: &quot;reference:Atm:bankId&amp;atmId&quot;,    &quot;example&quot;: &quot;bankId=e6a6f108-930a-4df8-8c70-68387180b81f&amp;atmId=49d57c01-59e8-40b2-b2bd-7acc1b0c3645&quot;}&quot;someField29&quot;: {    &quot;type&quot;: &quot;reference:BankAccount:bankId&amp;accountId&quot;,    &quot;example&quot;: &quot;bankId=e6a6f108-930a-4df8-8c70-68387180b81f&amp;accountId=49d57c01-59e8-40b2-b2bd-7acc1b0c3645&quot;}&quot;someField30&quot;: {    &quot;type&quot;: &quot;reference:Product:bankId&amp;productCode&quot;,    &quot;example&quot;: &quot;bankId=e6a6f108-930a-4df8-8c70-68387180b81f&amp;productCode=49d57c01-59e8-40b2-b2bd-7acc1b0c3645&quot;}&quot;someField31&quot;: {    &quot;type&quot;: &quot;reference:PhysicalCard:bankId&amp;cardId&quot;,    &quot;example&quot;: &quot;bankId=e6a6f108-930a-4df8-8c70-68387180b81f&amp;cardId=49d57c01-59e8-40b2-b2bd-7acc1b0c3645&quot;}&quot;someField32&quot;: {    &quot;type&quot;: &quot;reference:Transaction:bankId&amp;accountId&amp;transactionId&quot;,    &quot;example&quot;: &quot;bankId=e6a6f108-930a-4df8-8c70-68387180b81f&amp;accountId=49d57c01-59e8-40b2-b2bd-7acc1b0c3645&amp;transactionId=ae582fc1-41e9-49ac-bb84-eace98062292&quot;}&quot;someField33&quot;: {    &quot;type&quot;: &quot;reference:Counterparty:bankId&amp;accountId&amp;counterpartyId&quot;,    &quot;example&quot;: &quot;bankId=e6a6f108-930a-4df8-8c70-68387180b81f&amp;accountId=49d57c01-59e8-40b2-b2bd-7acc1b0c3645&amp;counterpartyId=ae582fc1-41e9-49ac-bb84-eace98062292&quot;}</code></pre><p>Note: if you set <code>hasPersonalEntity</code> = false, then OBP will not generate the CRUD my FooBar endpoints.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DynamicEntityFooBar**](DynamicEntityFooBar.md)| DynamicEntityFooBar object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**DynamicEntityFooBar**](DynamicEntityFooBar.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBankLevelEndpointTag**
> BankLevelEndpointTagResponseJson400 CreateBankLevelEndpointTag(ctx, body, bANKID)
Create Bank Level Endpoint Tag

<p>Create Bank Level Endpoint Tag</p><p>Note: Resource Docs are cached, TTL is 3600 seconds</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EndpointTagJson400**](EndpointTagJson400.md)| EndpointTagJson400 object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**BankLevelEndpointTagResponseJson400**](BankLevelEndpointTagResponseJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDynamicEndpoint**
> InlineResponse2011 CreateDynamicEndpoint(ctx, body)
Create Dynamic Endpoint

<p>Create dynamic endpoints.</p><p>Create dynamic endpoints with one json format swagger content.</p><p>If the host of swagger is <code>dynamic_entity</code>, then you need link the swagger fields to the dynamic entity fields,<br />please check <code>Endpoint Mapping</code> endpoints.</p><p>If the host of swagger is <code>obp_mock</code>, every dynamic endpoint will return example response of swagger,</p><p>when create MethodRouting for given dynamic endpoint, it will be routed to given url.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body3**](Body3.md)| JObject object that needs to be added. | 

### Return type

[**InlineResponse2011**](inline_response_201_1.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

# **CreateRegulatedEntity**
> RegulatedEntitiesJsonV510 CreateRegulatedEntity(ctx, body)
Create Regulated Entity

<p>Create Regulated Entity</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RegulatedEntityPostJsonV510**](RegulatedEntityPostJsonV510.md)| RegulatedEntityPostJsonV510 object that needs to be added. | 

### Return type

[**RegulatedEntitiesJsonV510**](RegulatedEntitiesJsonV510.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSystemDynamicEntity**
> DynamicEntityFooBar CreateSystemDynamicEntity(ctx, body)
Create System Level Dynamic Entity

<p>Create a system level Dynamic Entity.</p><p>Authentication is Mandatory</p><p>Create a DynamicEntity. If creation is successful, the corresponding POST, GET, PUT and DELETE (Create, Read, Update, Delete or CRUD for short) endpoints will be generated automatically</p><p>The following field types are as supported:<br />[number, integer, boolean, string, DATE_WITH_DAY, reference]</p><p>The DATE_WITH_DAY format is: yyyy-MM-dd</p><p>Reference types are like foreign keys and composite foreign keys are supported. The value you need to supply as the (composite) foreign key is a UUID (or several UUIDs in the case of a composite key) that match value in another Entity..<br />See the following list of currently available reference types and examples of how to construct key values correctly. Note: As more Dynamic Entities are created on this instance, this list will grow:</p><pre><code>&quot;someField0&quot;: {    &quot;type&quot;: &quot;reference:FishPort&quot;,    &quot;example&quot;: &quot;7b95a2c4-0663-49b6-a8a3-59f1221f2933&quot;}&quot;someField1&quot;: {    &quot;type&quot;: &quot;reference:FooBar&quot;,    &quot;example&quot;: &quot;7b95a2c4-0663-49b6-a8a3-59f1221f2933&quot;}&quot;someField2&quot;: {    &quot;type&quot;: &quot;reference:sustrans&quot;,    &quot;example&quot;: &quot;7b95a2c4-0663-49b6-a8a3-59f1221f2933&quot;}&quot;someField3&quot;: {    &quot;type&quot;: &quot;reference:SimonCovid&quot;,    &quot;example&quot;: &quot;7b95a2c4-0663-49b6-a8a3-59f1221f2933&quot;}&quot;someField4&quot;: {    &quot;type&quot;: &quot;reference:CovidAPIDays&quot;,    &quot;example&quot;: &quot;7b95a2c4-0663-49b6-a8a3-59f1221f2933&quot;}&quot;someField5&quot;: {    &quot;type&quot;: &quot;reference:customer_cars&quot;,    &quot;example&quot;: &quot;7b95a2c4-0663-49b6-a8a3-59f1221f2933&quot;}&quot;someField6&quot;: {    &quot;type&quot;: &quot;reference:MarchHare&quot;,    &quot;example&quot;: &quot;7b95a2c4-0663-49b6-a8a3-59f1221f2933&quot;}&quot;someField7&quot;: {    &quot;type&quot;: &quot;reference:InsurancePolicy&quot;,    &quot;example&quot;: &quot;7b95a2c4-0663-49b6-a8a3-59f1221f2933&quot;}&quot;someField8&quot;: {    &quot;type&quot;: &quot;reference:Odometer&quot;,    &quot;example&quot;: &quot;7b95a2c4-0663-49b6-a8a3-59f1221f2933&quot;}&quot;someField9&quot;: {    &quot;type&quot;: &quot;reference:InsurancePremium&quot;,    &quot;example&quot;: &quot;7b95a2c4-0663-49b6-a8a3-59f1221f2933&quot;}&quot;someField10&quot;: {    &quot;type&quot;: &quot;reference:ObpActivity&quot;,    &quot;example&quot;: &quot;7b95a2c4-0663-49b6-a8a3-59f1221f2933&quot;}&quot;someField11&quot;: {    &quot;type&quot;: &quot;reference:test1&quot;,    &quot;example&quot;: &quot;7b95a2c4-0663-49b6-a8a3-59f1221f2933&quot;}&quot;someField12&quot;: {    &quot;type&quot;: &quot;reference:D-Entity1&quot;,    &quot;example&quot;: &quot;7b95a2c4-0663-49b6-a8a3-59f1221f2933&quot;}&quot;someField13&quot;: {    &quot;type&quot;: &quot;reference:test_daniel707&quot;,    &quot;example&quot;: &quot;7b95a2c4-0663-49b6-a8a3-59f1221f2933&quot;}&quot;someField14&quot;: {    &quot;type&quot;: &quot;reference:Bank&quot;,    &quot;example&quot;: &quot;7b95a2c4-0663-49b6-a8a3-59f1221f2933&quot;}&quot;someField15&quot;: {    &quot;type&quot;: &quot;reference:Consumer&quot;,    &quot;example&quot;: &quot;7b95a2c4-0663-49b6-a8a3-59f1221f2933&quot;}&quot;someField16&quot;: {    &quot;type&quot;: &quot;reference:Customer&quot;,    &quot;example&quot;: &quot;7b95a2c4-0663-49b6-a8a3-59f1221f2933&quot;}&quot;someField17&quot;: {    &quot;type&quot;: &quot;reference:MethodRouting&quot;,    &quot;example&quot;: &quot;7b95a2c4-0663-49b6-a8a3-59f1221f2933&quot;}&quot;someField18&quot;: {    &quot;type&quot;: &quot;reference:DynamicEntity&quot;,    &quot;example&quot;: &quot;7b95a2c4-0663-49b6-a8a3-59f1221f2933&quot;}&quot;someField19&quot;: {    &quot;type&quot;: &quot;reference:TransactionRequest&quot;,    &quot;example&quot;: &quot;7b95a2c4-0663-49b6-a8a3-59f1221f2933&quot;}&quot;someField20&quot;: {    &quot;type&quot;: &quot;reference:ProductAttribute&quot;,    &quot;example&quot;: &quot;7b95a2c4-0663-49b6-a8a3-59f1221f2933&quot;}&quot;someField21&quot;: {    &quot;type&quot;: &quot;reference:AccountAttribute&quot;,    &quot;example&quot;: &quot;7b95a2c4-0663-49b6-a8a3-59f1221f2933&quot;}&quot;someField22&quot;: {    &quot;type&quot;: &quot;reference:TransactionAttribute&quot;,    &quot;example&quot;: &quot;7b95a2c4-0663-49b6-a8a3-59f1221f2933&quot;}&quot;someField23&quot;: {    &quot;type&quot;: &quot;reference:CustomerAttribute&quot;,    &quot;example&quot;: &quot;7b95a2c4-0663-49b6-a8a3-59f1221f2933&quot;}&quot;someField24&quot;: {    &quot;type&quot;: &quot;reference:AccountApplication&quot;,    &quot;example&quot;: &quot;7b95a2c4-0663-49b6-a8a3-59f1221f2933&quot;}&quot;someField25&quot;: {    &quot;type&quot;: &quot;reference:CardAttribute&quot;,    &quot;example&quot;: &quot;7b95a2c4-0663-49b6-a8a3-59f1221f2933&quot;}&quot;someField26&quot;: {    &quot;type&quot;: &quot;reference:Counterparty&quot;,    &quot;example&quot;: &quot;7b95a2c4-0663-49b6-a8a3-59f1221f2933&quot;}&quot;someField27&quot;: {    &quot;type&quot;: &quot;reference:Branch:bankId&amp;branchId&quot;,    &quot;example&quot;: &quot;bankId=7b95a2c4-0663-49b6-a8a3-59f1221f2933&amp;branchId=44f26efa-35c9-45b2-8951-b78d5638714d&quot;}&quot;someField28&quot;: {    &quot;type&quot;: &quot;reference:Atm:bankId&amp;atmId&quot;,    &quot;example&quot;: &quot;bankId=7b95a2c4-0663-49b6-a8a3-59f1221f2933&amp;atmId=44f26efa-35c9-45b2-8951-b78d5638714d&quot;}&quot;someField29&quot;: {    &quot;type&quot;: &quot;reference:BankAccount:bankId&amp;accountId&quot;,    &quot;example&quot;: &quot;bankId=7b95a2c4-0663-49b6-a8a3-59f1221f2933&amp;accountId=44f26efa-35c9-45b2-8951-b78d5638714d&quot;}&quot;someField30&quot;: {    &quot;type&quot;: &quot;reference:Product:bankId&amp;productCode&quot;,    &quot;example&quot;: &quot;bankId=7b95a2c4-0663-49b6-a8a3-59f1221f2933&amp;productCode=44f26efa-35c9-45b2-8951-b78d5638714d&quot;}&quot;someField31&quot;: {    &quot;type&quot;: &quot;reference:PhysicalCard:bankId&amp;cardId&quot;,    &quot;example&quot;: &quot;bankId=7b95a2c4-0663-49b6-a8a3-59f1221f2933&amp;cardId=44f26efa-35c9-45b2-8951-b78d5638714d&quot;}&quot;someField32&quot;: {    &quot;type&quot;: &quot;reference:Transaction:bankId&amp;accountId&amp;transactionId&quot;,    &quot;example&quot;: &quot;bankId=7b95a2c4-0663-49b6-a8a3-59f1221f2933&amp;accountId=44f26efa-35c9-45b2-8951-b78d5638714d&amp;transactionId=11f9c7eb-5307-4536-8279-bc9c8f1ee94d&quot;}&quot;someField33&quot;: {    &quot;type&quot;: &quot;reference:Counterparty:bankId&amp;accountId&amp;counterpartyId&quot;,    &quot;example&quot;: &quot;bankId=7b95a2c4-0663-49b6-a8a3-59f1221f2933&amp;accountId=44f26efa-35c9-45b2-8951-b78d5638714d&amp;counterpartyId=11f9c7eb-5307-4536-8279-bc9c8f1ee94d&quot;}</code></pre><p>Note: if you set <code>hasPersonalEntity</code> = false, then OBP will not generate the CRUD my FooBar endpoints.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DynamicEntityFooBar**](DynamicEntityFooBar.md)| DynamicEntityFooBar object that needs to be added. | 

### Return type

[**DynamicEntityFooBar**](DynamicEntityFooBar.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSystemLevelEndpointTag**
> BankLevelEndpointTagResponseJson400 CreateSystemLevelEndpointTag(ctx, body)
Create System Level Endpoint Tag

<p>Create System Level Endpoint Tag</p><p>Note: Resource Docs are cached, TTL is 3600 seconds</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EndpointTagJson400**](EndpointTagJson400.md)| EndpointTagJson400 object that needs to be added. | 

### Return type

[**BankLevelEndpointTagResponseJson400**](BankLevelEndpointTagResponseJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBankLevelDynamicEndpoint**
> DeleteBankLevelDynamicEndpoint(ctx, bANKID)
 Delete Bank Level Dynamic Endpoint

<p>Delete a Bank Level DynamicEndpoint specified by DYNAMIC_ENDPOINT_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBankLevelDynamicEntity**
> DeleteBankLevelDynamicEntity(ctx, dYNAMICENTITYID, bANKID)
Delete Bank Level Dynamic Entity

<p>Delete a Bank Level DynamicEntity specified by DYNAMIC_ENTITY_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dYNAMICENTITYID** | **string**| the dynamic entity id  | 
  **bANKID** | **string**| The bank id | 

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBankLevelEndpointTag**
> Full DeleteBankLevelEndpointTag(ctx, bANKID)
Delete Bank Level Endpoint Tag

<p>Delete Bank Level Endpoint Tag.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**Full**](Full.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDynamicEndpoint**
> DeleteDynamicEndpoint(ctx, )
 Delete Dynamic Endpoint

<p>Delete a DynamicEndpoint specified by DYNAMIC_ENDPOINT_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

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

# **DeleteMyDynamicEndpoint**
> DeleteMyDynamicEndpoint(ctx, )
Delete My Dynamic Endpoint

<p>Delete a DynamicEndpoint specified by DYNAMIC_ENDPOINT_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMyDynamicEntity**
> DeleteMyDynamicEntity(ctx, dYNAMICENTITYID)
Delete My Dynamic Entity

<p>Delete my DynamicEntity specified by DYNAMIC_ENTITY_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dYNAMICENTITYID** | **string**| the dynamic entity id  | 

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRegulatedEntity**
> DeleteRegulatedEntity(ctx, )
Delete Regulated Entity

<p>Delete Regulated Entity specified by REGULATED_ENTITY_ID</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSystemDynamicEntity**
> DeleteSystemDynamicEntity(ctx, dYNAMICENTITYID)
Delete System Level Dynamic Entity

<p>Delete a DynamicEntity specified by DYNAMIC_ENTITY_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dYNAMICENTITYID** | **string**| the dynamic entity id  | 

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSystemLevelEndpointTag**
> Full DeleteSystemLevelEndpointTag(ctx, )
Delete System Level Endpoint Tag

<p>Delete System Level Endpoint Tag.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Full**](Full.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ElasticSearchMetrics**
> EmptyClassJson ElasticSearchMetrics(ctx, body)
Search API Metrics via Elasticsearch

<p>Search the API calls made to this API instance via Elastic Search.</p><p>Login is required.</p><p>CanSearchMetrics entitlement is required to search metrics data.</p><p>parameters:</p><p>esType  - elasticsearch type</p><p>simple query:</p><p>q       - plain_text_query</p><p>df      - default field to search</p><p>sort    - field to sort on</p><p>size    - number of hits returned, default 10</p><p>from    - show hits starting from</p><p>json query:</p><p>source  - JSON_query_(URL-escaped)</p><p>example usage:</p><p>/search/metrics/q=findThis</p><p>or:</p><p>/search/metrics/source={&quot;query&quot;:{&quot;query_string&quot;:{&quot;query&quot;:&quot;findThis&quot;}}}</p><p>Note!!</p><p>The whole JSON query string MUST be URL-encoded:</p><ul><li>For {  use %7B</li><li>For }  use %7D</li><li>For : use %3A</li><li>For &quot; use %22</li></ul><p>etc..</p><p>Only q, source and esType are passed to Elastic</p><p>Elastic simple query: <a href=\"https://www.elastic.co/guide/en/elasticsearch/reference/current/search-uri-request.html\">https://www.elastic.co/guide/en/elasticsearch/reference/current/search-uri-request.html</a></p><p>Elastic JSON query: <a href=\"https://www.elastic.co/guide/en/elasticsearch/reference/current/query-filter-context.html\">https://www.elastic.co/guide/en/elasticsearch/reference/current/query-filter-context.html</a></p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 

### Return type

[**EmptyClassJson**](EmptyClassJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAdapterInfo**
> AdapterInfoJsonV500 GetAdapterInfo(ctx, )
Get Adapter Info

<p>Get basic information about the Adapter.</p><p>Authentication is Optional</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AdapterInfoJsonV500**](AdapterInfoJsonV500.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAdapterInfoForBank**
> AdapterInfoJsonV300 GetAdapterInfoForBank(ctx, bANKID)
Get Adapter Info for a bank

<p>Get basic information about the Adapter listening on behalf of this bank.</p><p>Authentication is Optional</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**AdapterInfoJsonV300**](AdapterInfoJsonV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBankLevelDynamicEndpoint**
> InlineResponse201 GetBankLevelDynamicEndpoint(ctx, bANKID)
 Get Bank Level Dynamic Endpoint

<p>Get a Bank Level Dynamic Endpoint.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBankLevelDynamicEndpoints**
> InlineResponse2003 GetBankLevelDynamicEndpoints(ctx, bANKID)
Get Bank Level Dynamic Endpoints

<p>Get Bank Level Dynamic Endpoints.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBankLevelDynamicEntities**
> InlineResponse2004 GetBankLevelDynamicEntities(ctx, bANKID)
Get Bank Level Dynamic Entities

<p>Get all the bank level Dynamic Entities for one bank.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

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

# **GetBankLevelEndpointTags**
> NotSupportedYet GetBankLevelEndpointTags(ctx, bANKID)
Get Bank Level Endpoint Tags

<p>Get Bank Level Endpoint Tags.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**NotSupportedYet**](NotSupportedYet.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCallContext**
> GetCallContext(ctx, )
Get the Call Context of a current call

<p>Get the Call Context of the current call.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectorMetrics**
> ConnectorMetricsJson GetConnectorMetrics(ctx, body)
Get Connector Metrics

<p>Get the all metrics</p><p>require CanGetConnectorMetrics role</p><p>Filters Part 1.<em>filtering</em> (no wilde cards etc.) parameters to GET /management/connector/metrics</p><p>Should be able to filter on the following metrics fields</p><p>eg: /management/connector/metrics?from_date=1100-01-01T01:01:01.000Z&amp;to_date=1100-01-01T01:01:01.000Z&amp;limit=50&amp;offset=2</p><p>1 from_date (defaults to one week before current date): eg:from_date=1100-01-01T01:01:01.000Z</p><p>2 to_date (defaults to current date) eg:to_date=1100-01-01T01:01:01.000Z</p><p>3 limit (for pagination: defaults to 1000)  eg:limit=2000</p><p>4 offset (for pagination: zero index, defaults to 0) eg: offset=10</p><p>eg: /management/connector/metrics?from_date=1100-01-01T01:01:01.000Z&amp;to_date=1100-01-01T01:01:01.000Z&amp;limit=100&amp;offset=300</p><p>Other filters:</p><p>5 connector_name  (if null ignore)</p><p>6 function_name (if null ignore)</p><p>7 correlation_id (if null ignore)</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 

### Return type

[**ConnectorMetricsJson**](ConnectorMetricsJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDynamicEndpoint**
> InlineResponse2011 GetDynamicEndpoint(ctx, )
Get Dynamic Endpoint

<p>Get a Dynamic Endpoint.</p><p>Get one DynamicEndpoint,</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2011**](inline_response_201_1.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDynamicEndpoints**
> InlineResponse2009 GetDynamicEndpoints(ctx, )
 Get Dynamic Endpoints

<p>Get Dynamic Endpoints.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMapperDatabaseInfo**
> AdapterInfoJsonV300 GetMapperDatabaseInfo(ctx, )
Get Mapper Database Info

<p>Get basic information about the Mapper Database.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AdapterInfoJsonV300**](AdapterInfoJsonV300.md)

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

# **GetMetrics**
> MetricsJsonV510 GetMetrics(ctx, body)
Get Metrics

<p>Get API metrics rows. These are records of each REST API call.</p><p>require CanReadMetrics role</p><p>Filters Part 1.<em>filtering</em> (no wilde cards etc.) parameters to GET /management/metrics</p><p>You can filter by the following fields by applying url parameters</p><p>eg: /management/metrics?from_date=1100-01-01T01:01:01.000Z&amp;to_date=1100-01-01T01:01:01.000Z&amp;limit=50&amp;offset=2</p><p>1 from_date e.g.:from_date=1100-01-01T01:01:01.000Z Defaults to the Unix Epoch i.e. Thu Jan 01 00:00:00 UTC 1970</p><p>2 to_date e.g.:to_date=1100-01-01T01:01:01.000Z Defaults to a far future date i.e. Sat Jan 01 00:00:00 UTC 4000</p><p>Note: it is recommended you send a valid from_date (e.g. 5 seconds ago) and to_date (now + 1 second) if you want to get the latest records<br />Otherwise you may receive stale cached results.</p><p>3 limit (for pagination: defaults to 50)  eg:limit=200</p><p>4 offset (for pagination: zero index, defaults to 0) eg: offset=10</p><p>5 sort_by (defaults to date field) eg: sort_by=date<br />possible values:<br />&quot;url&quot;,<br />&quot;date&quot;,<br />&quot;user_name&quot;,<br />&quot;app_name&quot;,<br />&quot;developer_email&quot;,<br />&quot;implemented_by_partial_function&quot;,<br />&quot;implemented_in_version&quot;,<br />&quot;consumer_id&quot;,<br />&quot;verb&quot;</p><p>6 direction (defaults to date desc) eg: direction=desc</p><p>eg: /management/metrics?from_date=1100-01-01T01:01:01.000Z&amp;to_date=1100-01-01T01:01:<a href=\"&#x6d;a&#x69;&#x6c;&#116;&#x6f;&#x3a;&#x30;1&#x2e;&#x30;&#48;&#x30;Z&#38;&#108;&#105;&#109;&#x69;&#116;&#61;&#x31;&#x30;&#x30;0&#x30;&#38;&#111;&#102;&#102;&#x73;&#101;t&#x3d;&#x30;&amp;a&#x6e;&#x6f;n&#x3d;&#x66;a&#108;&#115;&#101;&#38;&#97;&#x70;&#112;&#95;&#110;&#97;&#109;&#x65;&#61;&#x54;&#101;&#97;&#116;A&#x70;&#x70;&#38;&#x69;&#x6d;&#112;&#108;&#x65;&#x6d;en&#x74;&#x65;d&#x5f;&#105;&#x6e;_&#118;&#101;&#114;&#x73;&#105;&#111;n&#x3d;&#x76;2&#46;&#49;.&#48;&amp;&#118;&#x65;&#114;&#x62;&#61;&#x50;&#79;&#83;&#84;&#x26;&#x75;&#115;e&#x72;&#95;&#x69;&#x64;&#x3d;&#99;&#x37;&#98;&#54;&#99;&#x62;&#52;7&#x2d;&#99;&#x62;&#x39;&#x36;&#x2d;4&#x34;&#52;1&#45;&#56;&#x38;&#48;1&#x2d;&#x33;&#53;&#x62;&#53;&#x37;45&#x36;&#55;&#x35;&#51;a&#x26;&#117;&#x73;&#x65;&#114;&#x5f;n&#97;&#109;e&#61;&#x73;u&#x73;&#97;n&#x2e;&#x75;&#107;&#46;&#x32;&#x39;&#x40;e&#120;&#97;&#109;p&#x6c;&#101;&#x2e;&#99;&#111;&#x6d;\">&#x30;&#49;&#46;&#x30;&#48;&#48;&#x5a;&#38;&#108;&#x69;m&#105;&#116;=&#x31;0&#48;&#48;&#x30;&#x26;&#111;&#102;&#x66;&#115;&#x65;&#x74;&#x3d;&#x30;&#38;&#97;&#x6e;&#x6f;n&#x3d;&#102;&#97;&#x6c;&#115;&#101;&amp;&#97;&#x70;&#112;&#x5f;&#110;&#x61;m&#x65;&#x3d;&#84;&#x65;&#x61;&#x74;&#65;&#x70;&#x70;&#38;&#x69;&#109;p&#108;e&#x6d;&#x65;&#x6e;&#116;&#101;&#100;&#95;&#105;&#110;_&#118;&#x65;&#114;&#x73;&#x69;o&#x6e;&#x3d;&#x76;&#50;.&#x31;&#x2e;&#48;&#x26;&#x76;&#101;&#114;&#x62;&#x3d;&#80;&#79;&#83;T&#x26;&#117;&#x73;e&#x72;&#x5f;&#105;&#100;&#61;&#99;&#x37;&#98;&#x36;&#99;&#x62;&#52;7&#x2d;&#99;&#98;9&#x36;&#x2d;&#x34;&#52;&#52;1&#45;&#x38;80&#49;&#x2d;&#51;&#x35;&#98;&#x35;7&#52;&#53;&#x36;&#x37;&#53;&#x33;&#97;&amp;&#117;&#x73;e&#114;&#x5f;&#x6e;a&#x6d;&#x65;=&#115;&#117;&#x73;&#97;&#x6e;&#x2e;&#117;&#x6b;&#46;2&#57;&#x40;&#101;&#120;&#x61;m&#x70;&#x6c;&#101;&#x2e;&#x63;&#x6f;&#x6d;</a>&amp;consumer_id=78</p><p>Other filters:</p><p>7 consumer_id  (if null ignore)</p><p>8 user_id (if null ignore)</p><p>9 anon (if null ignore) only support two value : true (return where user_id is null.) or false (return where user_id is not null.)</p><p>10 url (if null ignore), note: can not contain '&amp;'.</p><p>11 app_name (if null ignore)</p><p>12 implemented_by_partial_function (if null ignore),</p><p>13 implemented_in_version (if null ignore)</p><p>14 verb (if null ignore)</p><p>15 correlation_id (if null ignore)</p><p>16 duration (if null ignore) non digit chars will be silently omitted</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 

### Return type

[**MetricsJsonV510**](MetricsJsonV510.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMetricsAtBank**
> MetricsJson GetMetricsAtBank(ctx, bANKID)
Get Metrics at Bank

<p>Get the all metrics at the Bank specified by BANK_ID</p><p>require CanReadMetrics role</p><p>Filters Part 1.<em>filtering</em> (no wilde cards etc.) parameters to GET /management/metrics</p><p>Should be able to filter on the following metrics fields</p><p>eg: /management/metrics?from_date=1100-01-01T01:01:01.000Z&amp;to_date=1100-01-01T01:01:01.000Z&amp;limit=50&amp;offset=2</p><p>1 from_date (defaults to one week before current date): eg:from_date=1100-01-01T01:01:01.000Z</p><p>2 to_date (defaults to current date) eg:to_date=1100-01-01T01:01:01.000Z</p><p>3 limit (for pagination: defaults to 50)  eg:limit=200</p><p>4 offset (for pagination: zero index, defaults to 0) eg: offset=10</p><p>5 sort_by (defaults to date field) eg: sort_by=date<br />possible values:<br />&quot;url&quot;,<br />&quot;date&quot;,<br />&quot;user_name&quot;,<br />&quot;app_name&quot;,<br />&quot;developer_email&quot;,<br />&quot;implemented_by_partial_function&quot;,<br />&quot;implemented_in_version&quot;,<br />&quot;consumer_id&quot;,<br />&quot;verb&quot;</p><p>6 direction (defaults to date desc) eg: direction=desc</p><p>eg: /management/metrics?from_date=1100-01-01T01:01:01.000Z&amp;to_date=1100-01-01T01:01:<a href=\"&#109;&#x61;&#x69;lt&#111;&#58;&#x30;&#x31;&#46;&#48;0&#x30;&#90;&#38;&#108;&#x69;&#109;&#x69;&#x74;=&#x31;0&#x30;&#x30;&#x30;&amp;&#x6f;&#102;fset=&#48;&#x26;&#x61;&#x6e;&#111;&#x6e;&#x3d;&#102;a&#x6c;se&#38;&#97;&#x70;&#x70;&#95;&#110;&#x61;&#x6d;&#x65;=&#84;&#101;a&#x74;&#65;&#x70;&#x70;&#38;&#x69;&#109;pl&#101;&#x6d;&#x65;&#110;&#116;&#x65;&#100;&#x5f;&#x69;&#x6e;&#95;ve&#x72;s&#105;&#x6f;&#110;&#61;&#x76;&#50;&#46;&#x31;&#x2e;0&#x26;v&#101;&#x72;&#x62;&#x3d;&#x50;&#79;&#x53;&#x54;&amp;&#117;&#x73;&#101;&#114;&#x5f;&#x69;&#x64;&#61;&#x63;&#x37;b&#54;c&#98;&#x34;&#55;&#45;&#x63;&#x62;&#x39;&#54;&#45;&#52;&#x34;&#x34;&#x31;&#45;&#x38;&#56;0&#49;-35&#98;&#53;&#55;&#x34;&#x35;&#x36;7&#x35;&#51;&#97;&#x26;&#x75;&#115;e&#114;&#x5f;&#110;a&#x6d;&#101;&#x3d;&#x73;&#117;&#x73;&#97;&#110;.uk&#x2e;2&#x39;&#64;ex&#97;&#109;&#112;l&#101;&#x2e;&#x63;&#x6f;&#109;\">&#48;&#x31;&#x2e;&#48;&#x30;0&#90;&#x26;l&#x69;&#x6d;&#105;&#116;&#61;&#x31;&#48;&#x30;&#x30;&#x30;&#38;&#111;&#x66;&#102;s&#101;&#116;&#x3d;&#48;&#x26;&#x61;&#110;o&#110;&#61;fal&#x73;&#x65;&#x26;&#97;&#112;&#x70;&#95;&#x6e;&#97;&#x6d;&#101;=&#x54;&#x65;&#x61;&#116;&#65;&#x70;&#112;&#x26;&#x69;&#109;&#x70;&#x6c;&#101;&#x6d;&#x65;&#x6e;&#116;&#x65;&#100;&#95;&#105;&#x6e;&#x5f;&#118;&#101;&#x72;&#115;&#105;o&#x6e;=&#x76;2&#x2e;&#x31;.&#48;&#x26;&#118;e&#x72;&#98;&#x3d;&#x50;&#79;&#83;T&#x26;&#x75;&#115;&#101;&#x72;&#95;&#x69;&#100;=&#x63;7&#x62;&#x36;cb&#x34;7&#45;c&#98;&#57;&#54;&#x2d;&#x34;&#x34;&#52;&#49;-&#56;8&#x30;&#x31;&#x2d;3&#53;&#x62;57&#52;&#x35;&#54;&#x37;&#53;3&#97;&#38;&#117;&#x73;&#x65;r&#x5f;n&#97;m&#x65;&#61;&#x73;u&#x73;&#x61;&#110;&#46;&#x75;&#107;&#46;&#x32;9&#x40;&#101;x&#97;&#109;&#x70;&#x6c;&#101;&#46;c&#x6f;&#x6d;</a>&amp;consumer_id=78</p><p>Other filters:</p><p>7 consumer_id  (if null ignore)</p><p>8 user_id (if null ignore)</p><p>9 anon (if null ignore) only support two value : true (return where user_id is null.) or false (return where user_id is not null.)</p><p>10 url (if null ignore), note: can not contain '&amp;'.</p><p>11 app_name (if null ignore)</p><p>12 implemented_by_partial_function (if null ignore),</p><p>13 implemented_in_version (if null ignore)</p><p>14 verb (if null ignore)</p><p>15 correlation_id (if null ignore)</p><p>16 duration (if null ignore) non digit chars will be silently omitted</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**MetricsJson**](MetricsJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMyDynamicEndpoints**
> InlineResponse20012 GetMyDynamicEndpoints(ctx, )
Get My Dynamic Endpoints

<p>Get My Dynamic Endpoints.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMyDynamicEntities**
> InlineResponse2004 GetMyDynamicEntities(ctx, )
Get My Dynamic Entities

<p>Get all my Dynamic Entities.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOAuth2ServerJWKsURIs**
> OAuth2ServerJwksUrisJson GetOAuth2ServerJWKsURIs(ctx, )
Get JSON Web Key (JWK) URIs

<p>Get the OAuth2 server's public JSON Web Key (JWK) URIs.<br />It is required by client applications to validate ID tokens, self-contained access tokens and other issued objects.</p><p>Authentication is Optional</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**OAuth2ServerJwksUrisJson**](OAuth2ServerJwksUrisJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetObpConnectorLoopback**
> ObpApiLoopbackJson GetObpConnectorLoopback(ctx, )
Get Connector Status (Loopback)

<p>This endpoint makes a call to the Connector to check the backend transport (e.g. Kafka) is reachable.</p><p>Currently this is only implemented for Kafka based connectors.</p><p>For Kafka based connectors, this endpoint writes a message to Kafka and reads it again.</p><p>In the future, this endpoint may also return information about database connections etc.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ObpApiLoopbackJson**](ObpApiLoopbackJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRateLimitingInfo**
> RateLimitingInfoV310 GetRateLimitingInfo(ctx, )
Get Rate Limiting Info

<p>Get information about the Rate Limiting setup on this OBP Instance such as:</p><p>Is rate limiting enabled and active?<br />What backend is used to keep track of the API calls (e.g. REDIS).</p><p>Note: Rate limiting can be set at the Consumer level and also for anonymous calls.</p><p>See the consumer rate limits / call limits endpoints.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RateLimitingInfoV310**](RateLimitingInfoV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRegulatedEntityById**
> RegulatedEntitiesJsonV510 GetRegulatedEntityById(ctx, )
Get Regulated Entity

<p>Get Regulated Entity By REGULATED_ENTITY_ID</p><p>Authentication is Optional</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RegulatedEntitiesJsonV510**](RegulatedEntitiesJsonV510.md)

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

# **GetSystemDynamicEntities**
> InlineResponse2004 GetSystemDynamicEntities(ctx, )
Get System Dynamic Entities

<p>Get all System Dynamic Entities</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSystemLevelEndpointTags**
> NotSupportedYet GetSystemLevelEndpointTags(ctx, )
Get System Level Endpoint Tags

<p>Get System Level Endpoint Tags.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NotSupportedYet**](NotSupportedYet.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegulatedEntities**
> RegulatedEntitiesJsonV510 RegulatedEntities(ctx, )
Get Regulated Entities

<p>Returns information about:</p><ul><li>Regulated Entities</li></ul><p>Authentication is Optional</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RegulatedEntitiesJsonV510**](RegulatedEntitiesJsonV510.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Root**
> ApiInfoJson400 Root(ctx, )
Get API Info (root)

<p>Returns information about:</p><ul><li>API version</li><li>Hosted by information</li><li>Hosted at information</li><li>Energy source information</li><li>Git Commit</li></ul><p>Authentication is Optional</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiInfoJson400**](APIInfoJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SuggestedSessionTimeout**
> SuggestedSessionTimeoutV510 SuggestedSessionTimeout(ctx, )
Get Suggested Session Timeout

<p>Returns information about:</p><ul><li>Suggested session timeout in case of a user inactivity</li></ul><p>Authentication is Optional</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SuggestedSessionTimeoutV510**](SuggestedSessionTimeoutV510.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBankLevelDynamicEndpointHost**
> DynamicEndpointHostJson400 UpdateBankLevelDynamicEndpointHost(ctx, body, bANKID)
 Update Bank Level Dynamic Endpoint Host

<p>Update Bank Level  dynamic endpoint Host.<br />The value can be obp_mock, dynamic_entity, or some service url.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DynamicEndpointHostJson400**](DynamicEndpointHostJson400.md)| DynamicEndpointHostJson400 object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**DynamicEndpointHostJson400**](DynamicEndpointHostJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBankLevelDynamicEntity**
> DynamicEntityFooBar UpdateBankLevelDynamicEntity(ctx, body, dYNAMICENTITYID, bANKID)
Update Bank Level Dynamic Entity

<p>Update a Bank Level DynamicEntity.</p><p>Authentication is Mandatory</p><p>Update one DynamicEntity, after update finished, the corresponding CRUD endpoints will be changed.</p><p>The following field types are as supported:<br />[number, integer, boolean, string, DATE_WITH_DAY, reference]</p><p>DATE_WITH_DAY format: yyyy-MM-dd</p><p>Reference types are like foreign keys and composite foreign keys are supported. The value you need to supply as the (composite) foreign key is a UUID (or several UUIDs in the case of a composite key) that match value in another Entity..<br />The following list shows all the possible reference types in the system with corresponding examples values so you can see how to construct each reference type value.</p><pre><code>&quot;someField0&quot;: {    &quot;type&quot;: &quot;reference:FishPort&quot;,    &quot;example&quot;: &quot;d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&quot;}&quot;someField1&quot;: {    &quot;type&quot;: &quot;reference:FooBar&quot;,    &quot;example&quot;: &quot;d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&quot;}&quot;someField2&quot;: {    &quot;type&quot;: &quot;reference:sustrans&quot;,    &quot;example&quot;: &quot;d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&quot;}&quot;someField3&quot;: {    &quot;type&quot;: &quot;reference:SimonCovid&quot;,    &quot;example&quot;: &quot;d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&quot;}&quot;someField4&quot;: {    &quot;type&quot;: &quot;reference:CovidAPIDays&quot;,    &quot;example&quot;: &quot;d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&quot;}&quot;someField5&quot;: {    &quot;type&quot;: &quot;reference:customer_cars&quot;,    &quot;example&quot;: &quot;d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&quot;}&quot;someField6&quot;: {    &quot;type&quot;: &quot;reference:MarchHare&quot;,    &quot;example&quot;: &quot;d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&quot;}&quot;someField7&quot;: {    &quot;type&quot;: &quot;reference:InsurancePolicy&quot;,    &quot;example&quot;: &quot;d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&quot;}&quot;someField8&quot;: {    &quot;type&quot;: &quot;reference:Odometer&quot;,    &quot;example&quot;: &quot;d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&quot;}&quot;someField9&quot;: {    &quot;type&quot;: &quot;reference:InsurancePremium&quot;,    &quot;example&quot;: &quot;d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&quot;}&quot;someField10&quot;: {    &quot;type&quot;: &quot;reference:ObpActivity&quot;,    &quot;example&quot;: &quot;d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&quot;}&quot;someField11&quot;: {    &quot;type&quot;: &quot;reference:test1&quot;,    &quot;example&quot;: &quot;d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&quot;}&quot;someField12&quot;: {    &quot;type&quot;: &quot;reference:D-Entity1&quot;,    &quot;example&quot;: &quot;d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&quot;}&quot;someField13&quot;: {    &quot;type&quot;: &quot;reference:test_daniel707&quot;,    &quot;example&quot;: &quot;d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&quot;}&quot;someField14&quot;: {    &quot;type&quot;: &quot;reference:Bank&quot;,    &quot;example&quot;: &quot;d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&quot;}&quot;someField15&quot;: {    &quot;type&quot;: &quot;reference:Consumer&quot;,    &quot;example&quot;: &quot;d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&quot;}&quot;someField16&quot;: {    &quot;type&quot;: &quot;reference:Customer&quot;,    &quot;example&quot;: &quot;d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&quot;}&quot;someField17&quot;: {    &quot;type&quot;: &quot;reference:MethodRouting&quot;,    &quot;example&quot;: &quot;d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&quot;}&quot;someField18&quot;: {    &quot;type&quot;: &quot;reference:DynamicEntity&quot;,    &quot;example&quot;: &quot;d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&quot;}&quot;someField19&quot;: {    &quot;type&quot;: &quot;reference:TransactionRequest&quot;,    &quot;example&quot;: &quot;d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&quot;}&quot;someField20&quot;: {    &quot;type&quot;: &quot;reference:ProductAttribute&quot;,    &quot;example&quot;: &quot;d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&quot;}&quot;someField21&quot;: {    &quot;type&quot;: &quot;reference:AccountAttribute&quot;,    &quot;example&quot;: &quot;d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&quot;}&quot;someField22&quot;: {    &quot;type&quot;: &quot;reference:TransactionAttribute&quot;,    &quot;example&quot;: &quot;d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&quot;}&quot;someField23&quot;: {    &quot;type&quot;: &quot;reference:CustomerAttribute&quot;,    &quot;example&quot;: &quot;d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&quot;}&quot;someField24&quot;: {    &quot;type&quot;: &quot;reference:AccountApplication&quot;,    &quot;example&quot;: &quot;d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&quot;}&quot;someField25&quot;: {    &quot;type&quot;: &quot;reference:CardAttribute&quot;,    &quot;example&quot;: &quot;d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&quot;}&quot;someField26&quot;: {    &quot;type&quot;: &quot;reference:Counterparty&quot;,    &quot;example&quot;: &quot;d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&quot;}&quot;someField27&quot;: {    &quot;type&quot;: &quot;reference:Branch:bankId&amp;branchId&quot;,    &quot;example&quot;: &quot;bankId=d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&amp;branchId=3487733b-59c5-467a-a823-23285c5b88ed&quot;}&quot;someField28&quot;: {    &quot;type&quot;: &quot;reference:Atm:bankId&amp;atmId&quot;,    &quot;example&quot;: &quot;bankId=d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&amp;atmId=3487733b-59c5-467a-a823-23285c5b88ed&quot;}&quot;someField29&quot;: {    &quot;type&quot;: &quot;reference:BankAccount:bankId&amp;accountId&quot;,    &quot;example&quot;: &quot;bankId=d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&amp;accountId=3487733b-59c5-467a-a823-23285c5b88ed&quot;}&quot;someField30&quot;: {    &quot;type&quot;: &quot;reference:Product:bankId&amp;productCode&quot;,    &quot;example&quot;: &quot;bankId=d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&amp;productCode=3487733b-59c5-467a-a823-23285c5b88ed&quot;}&quot;someField31&quot;: {    &quot;type&quot;: &quot;reference:PhysicalCard:bankId&amp;cardId&quot;,    &quot;example&quot;: &quot;bankId=d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&amp;cardId=3487733b-59c5-467a-a823-23285c5b88ed&quot;}&quot;someField32&quot;: {    &quot;type&quot;: &quot;reference:Transaction:bankId&amp;accountId&amp;transactionId&quot;,    &quot;example&quot;: &quot;bankId=d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&amp;accountId=3487733b-59c5-467a-a823-23285c5b88ed&amp;transactionId=c870df43-7938-49d1-8b13-d880105b1c4b&quot;}&quot;someField33&quot;: {    &quot;type&quot;: &quot;reference:Counterparty:bankId&amp;accountId&amp;counterpartyId&quot;,    &quot;example&quot;: &quot;bankId=d21c7c48-8ffc-4c3f-8d4d-f4c20432d8bf&amp;accountId=3487733b-59c5-467a-a823-23285c5b88ed&amp;counterpartyId=c870df43-7938-49d1-8b13-d880105b1c4b&quot;}</code></pre>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DynamicEntityFooBar**](DynamicEntityFooBar.md)| DynamicEntityFooBar object that needs to be added. | 
  **dYNAMICENTITYID** | **string**| the dynamic entity id  | 
  **bANKID** | **string**| The bank id | 

### Return type

[**DynamicEntityFooBar**](DynamicEntityFooBar.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBankLevelEndpointTag**
> BankLevelEndpointTagResponseJson400 UpdateBankLevelEndpointTag(ctx, body, bANKID)
Update Bank Level Endpoint Tag

<p>Update Endpoint Tag, you can only update the tag_name here, operation_id can not be updated.</p><p>Note: Resource Docs are cached, TTL is 3600 seconds</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EndpointTagJson400**](EndpointTagJson400.md)| EndpointTagJson400 object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**BankLevelEndpointTagResponseJson400**](BankLevelEndpointTagResponseJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDynamicEndpointHost**
> DynamicEndpointHostJson400 UpdateDynamicEndpointHost(ctx, body)
 Update Dynamic Endpoint Host

<p>Update dynamic endpoint Host.<br />The value can be obp_mock, dynamic_entity, or some service url.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DynamicEndpointHostJson400**](DynamicEndpointHostJson400.md)| DynamicEndpointHostJson400 object that needs to be added. | 

### Return type

[**DynamicEndpointHostJson400**](DynamicEndpointHostJson400.md)

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

# **UpdateMyDynamicEntity**
> DynamicEntityFooBar UpdateMyDynamicEntity(ctx, body, dYNAMICENTITYID)
Update My Dynamic Entity

<p>Update my DynamicEntity.</p><p>Authentication is Mandatory</p><p>Update one of my DynamicEntity, after update finished, the corresponding CRUD endpoints will be changed.</p><p>Current support filed types as follow:<br />[number, integer, boolean, string, DATE_WITH_DAY, reference]</p><p>DATE_WITH_DAY format: yyyy-MM-dd</p><p>Reference types are like foreign keys and composite foreign keys are supported. The value you need to supply as the (composite) foreign key is a UUID (or several UUIDs in the case of a composite key) that match value in another Entity..<br />The following list shows all the possible reference types in the system with corresponding examples values so you can see how to construct each reference type value.</p><pre><code>&quot;someField0&quot;: {    &quot;type&quot;: &quot;reference:FishPort&quot;,    &quot;example&quot;: &quot;c18a77b4-5d30-4fcb-84a0-dca287871a4f&quot;}&quot;someField1&quot;: {    &quot;type&quot;: &quot;reference:FooBar&quot;,    &quot;example&quot;: &quot;c18a77b4-5d30-4fcb-84a0-dca287871a4f&quot;}&quot;someField2&quot;: {    &quot;type&quot;: &quot;reference:sustrans&quot;,    &quot;example&quot;: &quot;c18a77b4-5d30-4fcb-84a0-dca287871a4f&quot;}&quot;someField3&quot;: {    &quot;type&quot;: &quot;reference:SimonCovid&quot;,    &quot;example&quot;: &quot;c18a77b4-5d30-4fcb-84a0-dca287871a4f&quot;}&quot;someField4&quot;: {    &quot;type&quot;: &quot;reference:CovidAPIDays&quot;,    &quot;example&quot;: &quot;c18a77b4-5d30-4fcb-84a0-dca287871a4f&quot;}&quot;someField5&quot;: {    &quot;type&quot;: &quot;reference:customer_cars&quot;,    &quot;example&quot;: &quot;c18a77b4-5d30-4fcb-84a0-dca287871a4f&quot;}&quot;someField6&quot;: {    &quot;type&quot;: &quot;reference:MarchHare&quot;,    &quot;example&quot;: &quot;c18a77b4-5d30-4fcb-84a0-dca287871a4f&quot;}&quot;someField7&quot;: {    &quot;type&quot;: &quot;reference:InsurancePolicy&quot;,    &quot;example&quot;: &quot;c18a77b4-5d30-4fcb-84a0-dca287871a4f&quot;}&quot;someField8&quot;: {    &quot;type&quot;: &quot;reference:Odometer&quot;,    &quot;example&quot;: &quot;c18a77b4-5d30-4fcb-84a0-dca287871a4f&quot;}&quot;someField9&quot;: {    &quot;type&quot;: &quot;reference:InsurancePremium&quot;,    &quot;example&quot;: &quot;c18a77b4-5d30-4fcb-84a0-dca287871a4f&quot;}&quot;someField10&quot;: {    &quot;type&quot;: &quot;reference:ObpActivity&quot;,    &quot;example&quot;: &quot;c18a77b4-5d30-4fcb-84a0-dca287871a4f&quot;}&quot;someField11&quot;: {    &quot;type&quot;: &quot;reference:test1&quot;,    &quot;example&quot;: &quot;c18a77b4-5d30-4fcb-84a0-dca287871a4f&quot;}&quot;someField12&quot;: {    &quot;type&quot;: &quot;reference:D-Entity1&quot;,    &quot;example&quot;: &quot;c18a77b4-5d30-4fcb-84a0-dca287871a4f&quot;}&quot;someField13&quot;: {    &quot;type&quot;: &quot;reference:test_daniel707&quot;,    &quot;example&quot;: &quot;c18a77b4-5d30-4fcb-84a0-dca287871a4f&quot;}&quot;someField14&quot;: {    &quot;type&quot;: &quot;reference:Bank&quot;,    &quot;example&quot;: &quot;c18a77b4-5d30-4fcb-84a0-dca287871a4f&quot;}&quot;someField15&quot;: {    &quot;type&quot;: &quot;reference:Consumer&quot;,    &quot;example&quot;: &quot;c18a77b4-5d30-4fcb-84a0-dca287871a4f&quot;}&quot;someField16&quot;: {    &quot;type&quot;: &quot;reference:Customer&quot;,    &quot;example&quot;: &quot;c18a77b4-5d30-4fcb-84a0-dca287871a4f&quot;}&quot;someField17&quot;: {    &quot;type&quot;: &quot;reference:MethodRouting&quot;,    &quot;example&quot;: &quot;c18a77b4-5d30-4fcb-84a0-dca287871a4f&quot;}&quot;someField18&quot;: {    &quot;type&quot;: &quot;reference:DynamicEntity&quot;,    &quot;example&quot;: &quot;c18a77b4-5d30-4fcb-84a0-dca287871a4f&quot;}&quot;someField19&quot;: {    &quot;type&quot;: &quot;reference:TransactionRequest&quot;,    &quot;example&quot;: &quot;c18a77b4-5d30-4fcb-84a0-dca287871a4f&quot;}&quot;someField20&quot;: {    &quot;type&quot;: &quot;reference:ProductAttribute&quot;,    &quot;example&quot;: &quot;c18a77b4-5d30-4fcb-84a0-dca287871a4f&quot;}&quot;someField21&quot;: {    &quot;type&quot;: &quot;reference:AccountAttribute&quot;,    &quot;example&quot;: &quot;c18a77b4-5d30-4fcb-84a0-dca287871a4f&quot;}&quot;someField22&quot;: {    &quot;type&quot;: &quot;reference:TransactionAttribute&quot;,    &quot;example&quot;: &quot;c18a77b4-5d30-4fcb-84a0-dca287871a4f&quot;}&quot;someField23&quot;: {    &quot;type&quot;: &quot;reference:CustomerAttribute&quot;,    &quot;example&quot;: &quot;c18a77b4-5d30-4fcb-84a0-dca287871a4f&quot;}&quot;someField24&quot;: {    &quot;type&quot;: &quot;reference:AccountApplication&quot;,    &quot;example&quot;: &quot;c18a77b4-5d30-4fcb-84a0-dca287871a4f&quot;}&quot;someField25&quot;: {    &quot;type&quot;: &quot;reference:CardAttribute&quot;,    &quot;example&quot;: &quot;c18a77b4-5d30-4fcb-84a0-dca287871a4f&quot;}&quot;someField26&quot;: {    &quot;type&quot;: &quot;reference:Counterparty&quot;,    &quot;example&quot;: &quot;c18a77b4-5d30-4fcb-84a0-dca287871a4f&quot;}&quot;someField27&quot;: {    &quot;type&quot;: &quot;reference:Branch:bankId&amp;branchId&quot;,    &quot;example&quot;: &quot;bankId=c18a77b4-5d30-4fcb-84a0-dca287871a4f&amp;branchId=bb989f3d-6fb8-45b7-b60d-c20e6ae9f21f&quot;}&quot;someField28&quot;: {    &quot;type&quot;: &quot;reference:Atm:bankId&amp;atmId&quot;,    &quot;example&quot;: &quot;bankId=c18a77b4-5d30-4fcb-84a0-dca287871a4f&amp;atmId=bb989f3d-6fb8-45b7-b60d-c20e6ae9f21f&quot;}&quot;someField29&quot;: {    &quot;type&quot;: &quot;reference:BankAccount:bankId&amp;accountId&quot;,    &quot;example&quot;: &quot;bankId=c18a77b4-5d30-4fcb-84a0-dca287871a4f&amp;accountId=bb989f3d-6fb8-45b7-b60d-c20e6ae9f21f&quot;}&quot;someField30&quot;: {    &quot;type&quot;: &quot;reference:Product:bankId&amp;productCode&quot;,    &quot;example&quot;: &quot;bankId=c18a77b4-5d30-4fcb-84a0-dca287871a4f&amp;productCode=bb989f3d-6fb8-45b7-b60d-c20e6ae9f21f&quot;}&quot;someField31&quot;: {    &quot;type&quot;: &quot;reference:PhysicalCard:bankId&amp;cardId&quot;,    &quot;example&quot;: &quot;bankId=c18a77b4-5d30-4fcb-84a0-dca287871a4f&amp;cardId=bb989f3d-6fb8-45b7-b60d-c20e6ae9f21f&quot;}&quot;someField32&quot;: {    &quot;type&quot;: &quot;reference:Transaction:bankId&amp;accountId&amp;transactionId&quot;,    &quot;example&quot;: &quot;bankId=c18a77b4-5d30-4fcb-84a0-dca287871a4f&amp;accountId=bb989f3d-6fb8-45b7-b60d-c20e6ae9f21f&amp;transactionId=48e0d625-eae0-4c74-89e6-b99e67e2ab55&quot;}&quot;someField33&quot;: {    &quot;type&quot;: &quot;reference:Counterparty:bankId&amp;accountId&amp;counterpartyId&quot;,    &quot;example&quot;: &quot;bankId=c18a77b4-5d30-4fcb-84a0-dca287871a4f&amp;accountId=bb989f3d-6fb8-45b7-b60d-c20e6ae9f21f&amp;counterpartyId=48e0d625-eae0-4c74-89e6-b99e67e2ab55&quot;}</code></pre>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DynamicEntityFooBar**](DynamicEntityFooBar.md)| DynamicEntityFooBar object that needs to be added. | 
  **dYNAMICENTITYID** | **string**| the dynamic entity id  | 

### Return type

[**DynamicEntityFooBar**](DynamicEntityFooBar.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSystemDynamicEntity**
> DynamicEntityFooBar UpdateSystemDynamicEntity(ctx, body, dYNAMICENTITYID)
Update System Level Dynamic Entity

<p>Update a System Level Dynamic Entity.</p><p>Authentication is Mandatory</p><p>Update one DynamicEntity, after update finished, the corresponding CRUD endpoints will be changed.</p><p>The following field types are as supported:<br />[number, integer, boolean, string, DATE_WITH_DAY, reference]</p><p>DATE_WITH_DAY format: yyyy-MM-dd</p><p>Reference types are like foreign keys and composite foreign keys are supported. The value you need to supply as the (composite) foreign key is a UUID (or several UUIDs in the case of a composite key) that match value in another Entity..<br />The following list shows all the possible reference types in the system with corresponding examples values so you can see how to construct each reference type value.</p><pre><code>&quot;someField0&quot;: {    &quot;type&quot;: &quot;reference:FishPort&quot;,    &quot;example&quot;: &quot;05e6621b-e445-41fa-b140-699e106aa72a&quot;}&quot;someField1&quot;: {    &quot;type&quot;: &quot;reference:FooBar&quot;,    &quot;example&quot;: &quot;05e6621b-e445-41fa-b140-699e106aa72a&quot;}&quot;someField2&quot;: {    &quot;type&quot;: &quot;reference:sustrans&quot;,    &quot;example&quot;: &quot;05e6621b-e445-41fa-b140-699e106aa72a&quot;}&quot;someField3&quot;: {    &quot;type&quot;: &quot;reference:SimonCovid&quot;,    &quot;example&quot;: &quot;05e6621b-e445-41fa-b140-699e106aa72a&quot;}&quot;someField4&quot;: {    &quot;type&quot;: &quot;reference:CovidAPIDays&quot;,    &quot;example&quot;: &quot;05e6621b-e445-41fa-b140-699e106aa72a&quot;}&quot;someField5&quot;: {    &quot;type&quot;: &quot;reference:customer_cars&quot;,    &quot;example&quot;: &quot;05e6621b-e445-41fa-b140-699e106aa72a&quot;}&quot;someField6&quot;: {    &quot;type&quot;: &quot;reference:MarchHare&quot;,    &quot;example&quot;: &quot;05e6621b-e445-41fa-b140-699e106aa72a&quot;}&quot;someField7&quot;: {    &quot;type&quot;: &quot;reference:InsurancePolicy&quot;,    &quot;example&quot;: &quot;05e6621b-e445-41fa-b140-699e106aa72a&quot;}&quot;someField8&quot;: {    &quot;type&quot;: &quot;reference:Odometer&quot;,    &quot;example&quot;: &quot;05e6621b-e445-41fa-b140-699e106aa72a&quot;}&quot;someField9&quot;: {    &quot;type&quot;: &quot;reference:InsurancePremium&quot;,    &quot;example&quot;: &quot;05e6621b-e445-41fa-b140-699e106aa72a&quot;}&quot;someField10&quot;: {    &quot;type&quot;: &quot;reference:ObpActivity&quot;,    &quot;example&quot;: &quot;05e6621b-e445-41fa-b140-699e106aa72a&quot;}&quot;someField11&quot;: {    &quot;type&quot;: &quot;reference:test1&quot;,    &quot;example&quot;: &quot;05e6621b-e445-41fa-b140-699e106aa72a&quot;}&quot;someField12&quot;: {    &quot;type&quot;: &quot;reference:D-Entity1&quot;,    &quot;example&quot;: &quot;05e6621b-e445-41fa-b140-699e106aa72a&quot;}&quot;someField13&quot;: {    &quot;type&quot;: &quot;reference:test_daniel707&quot;,    &quot;example&quot;: &quot;05e6621b-e445-41fa-b140-699e106aa72a&quot;}&quot;someField14&quot;: {    &quot;type&quot;: &quot;reference:Bank&quot;,    &quot;example&quot;: &quot;05e6621b-e445-41fa-b140-699e106aa72a&quot;}&quot;someField15&quot;: {    &quot;type&quot;: &quot;reference:Consumer&quot;,    &quot;example&quot;: &quot;05e6621b-e445-41fa-b140-699e106aa72a&quot;}&quot;someField16&quot;: {    &quot;type&quot;: &quot;reference:Customer&quot;,    &quot;example&quot;: &quot;05e6621b-e445-41fa-b140-699e106aa72a&quot;}&quot;someField17&quot;: {    &quot;type&quot;: &quot;reference:MethodRouting&quot;,    &quot;example&quot;: &quot;05e6621b-e445-41fa-b140-699e106aa72a&quot;}&quot;someField18&quot;: {    &quot;type&quot;: &quot;reference:DynamicEntity&quot;,    &quot;example&quot;: &quot;05e6621b-e445-41fa-b140-699e106aa72a&quot;}&quot;someField19&quot;: {    &quot;type&quot;: &quot;reference:TransactionRequest&quot;,    &quot;example&quot;: &quot;05e6621b-e445-41fa-b140-699e106aa72a&quot;}&quot;someField20&quot;: {    &quot;type&quot;: &quot;reference:ProductAttribute&quot;,    &quot;example&quot;: &quot;05e6621b-e445-41fa-b140-699e106aa72a&quot;}&quot;someField21&quot;: {    &quot;type&quot;: &quot;reference:AccountAttribute&quot;,    &quot;example&quot;: &quot;05e6621b-e445-41fa-b140-699e106aa72a&quot;}&quot;someField22&quot;: {    &quot;type&quot;: &quot;reference:TransactionAttribute&quot;,    &quot;example&quot;: &quot;05e6621b-e445-41fa-b140-699e106aa72a&quot;}&quot;someField23&quot;: {    &quot;type&quot;: &quot;reference:CustomerAttribute&quot;,    &quot;example&quot;: &quot;05e6621b-e445-41fa-b140-699e106aa72a&quot;}&quot;someField24&quot;: {    &quot;type&quot;: &quot;reference:AccountApplication&quot;,    &quot;example&quot;: &quot;05e6621b-e445-41fa-b140-699e106aa72a&quot;}&quot;someField25&quot;: {    &quot;type&quot;: &quot;reference:CardAttribute&quot;,    &quot;example&quot;: &quot;05e6621b-e445-41fa-b140-699e106aa72a&quot;}&quot;someField26&quot;: {    &quot;type&quot;: &quot;reference:Counterparty&quot;,    &quot;example&quot;: &quot;05e6621b-e445-41fa-b140-699e106aa72a&quot;}&quot;someField27&quot;: {    &quot;type&quot;: &quot;reference:Branch:bankId&amp;branchId&quot;,    &quot;example&quot;: &quot;bankId=05e6621b-e445-41fa-b140-699e106aa72a&amp;branchId=0bcccc1a-68de-4a4c-9082-396f28a65142&quot;}&quot;someField28&quot;: {    &quot;type&quot;: &quot;reference:Atm:bankId&amp;atmId&quot;,    &quot;example&quot;: &quot;bankId=05e6621b-e445-41fa-b140-699e106aa72a&amp;atmId=0bcccc1a-68de-4a4c-9082-396f28a65142&quot;}&quot;someField29&quot;: {    &quot;type&quot;: &quot;reference:BankAccount:bankId&amp;accountId&quot;,    &quot;example&quot;: &quot;bankId=05e6621b-e445-41fa-b140-699e106aa72a&amp;accountId=0bcccc1a-68de-4a4c-9082-396f28a65142&quot;}&quot;someField30&quot;: {    &quot;type&quot;: &quot;reference:Product:bankId&amp;productCode&quot;,    &quot;example&quot;: &quot;bankId=05e6621b-e445-41fa-b140-699e106aa72a&amp;productCode=0bcccc1a-68de-4a4c-9082-396f28a65142&quot;}&quot;someField31&quot;: {    &quot;type&quot;: &quot;reference:PhysicalCard:bankId&amp;cardId&quot;,    &quot;example&quot;: &quot;bankId=05e6621b-e445-41fa-b140-699e106aa72a&amp;cardId=0bcccc1a-68de-4a4c-9082-396f28a65142&quot;}&quot;someField32&quot;: {    &quot;type&quot;: &quot;reference:Transaction:bankId&amp;accountId&amp;transactionId&quot;,    &quot;example&quot;: &quot;bankId=05e6621b-e445-41fa-b140-699e106aa72a&amp;accountId=0bcccc1a-68de-4a4c-9082-396f28a65142&amp;transactionId=f83b7548-74c0-4bc7-9f2c-bb195290bf78&quot;}&quot;someField33&quot;: {    &quot;type&quot;: &quot;reference:Counterparty:bankId&amp;accountId&amp;counterpartyId&quot;,    &quot;example&quot;: &quot;bankId=05e6621b-e445-41fa-b140-699e106aa72a&amp;accountId=0bcccc1a-68de-4a4c-9082-396f28a65142&amp;counterpartyId=f83b7548-74c0-4bc7-9f2c-bb195290bf78&quot;}</code></pre>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DynamicEntityFooBar**](DynamicEntityFooBar.md)| DynamicEntityFooBar object that needs to be added. | 
  **dYNAMICENTITYID** | **string**| the dynamic entity id  | 

### Return type

[**DynamicEntityFooBar**](DynamicEntityFooBar.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSystemLevelEndpointTag**
> BankLevelEndpointTagResponseJson400 UpdateSystemLevelEndpointTag(ctx, body)
Update System Level Endpoint Tag

<p>Update System Level Endpoint Tag, you can only update the tag_name here, operation_id can not be updated.</p><p>Note: Resource Docs are cached, TTL is 3600 seconds</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EndpointTagJson400**](EndpointTagJson400.md)| EndpointTagJson400 object that needs to be added. | 

### Return type

[**BankLevelEndpointTagResponseJson400**](BankLevelEndpointTagResponseJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyRequestSignResponse**
> VerifyRequestSignResponse(ctx, )
Verify Request and Sign Response of a current call

<p>Verify Request and Sign Response of a current call.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WaitingForGodot**
> WaitingForGodotJsonV510 WaitingForGodot(ctx, )
Waiting For Godot

<p>Waiting For Godot</p><p>Uses query parameter &quot;sleep&quot; in milliseconds.<br />For instance: .../waiting-for-godot?sleep=50 means postpone response in 50 milliseconds.</p><p>Authentication is Optional</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**WaitingForGodotJsonV510**](WaitingForGodotJsonV510.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

