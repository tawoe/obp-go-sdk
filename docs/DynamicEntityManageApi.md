# \DynamicEntityManageApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBankLevelDynamicEntity**](DynamicEntityManageApi.md#CreateBankLevelDynamicEntity) | **Post** /obp/v5.1.0/management/banks/{BANK_ID}/dynamic-entities | Create Bank Level Dynamic Entity
[**CreateSystemDynamicEntity**](DynamicEntityManageApi.md#CreateSystemDynamicEntity) | **Post** /obp/v5.1.0/management/system-dynamic-entities | Create System Level Dynamic Entity
[**DeleteBankLevelDynamicEntity**](DynamicEntityManageApi.md#DeleteBankLevelDynamicEntity) | **Delete** /obp/v5.1.0/management/banks/{BANK_ID}/dynamic-entities/{DYNAMIC_ENTITY_ID} | Delete Bank Level Dynamic Entity
[**DeleteMyDynamicEntity**](DynamicEntityManageApi.md#DeleteMyDynamicEntity) | **Delete** /obp/v5.1.0/my/dynamic-entities/{DYNAMIC_ENTITY_ID} | Delete My Dynamic Entity
[**DeleteSystemDynamicEntity**](DynamicEntityManageApi.md#DeleteSystemDynamicEntity) | **Delete** /obp/v5.1.0/management/system-dynamic-entities/{DYNAMIC_ENTITY_ID} | Delete System Level Dynamic Entity
[**GetBankLevelDynamicEntities**](DynamicEntityManageApi.md#GetBankLevelDynamicEntities) | **Get** /obp/v5.1.0/management/banks/{BANK_ID}/dynamic-entities | Get Bank Level Dynamic Entities
[**GetMyDynamicEntities**](DynamicEntityManageApi.md#GetMyDynamicEntities) | **Get** /obp/v5.1.0/my/dynamic-entities | Get My Dynamic Entities
[**GetSystemDynamicEntities**](DynamicEntityManageApi.md#GetSystemDynamicEntities) | **Get** /obp/v5.1.0/management/system-dynamic-entities | Get System Dynamic Entities
[**UpdateBankLevelDynamicEntity**](DynamicEntityManageApi.md#UpdateBankLevelDynamicEntity) | **Put** /obp/v5.1.0/management/banks/{BANK_ID}/dynamic-entities/{DYNAMIC_ENTITY_ID} | Update Bank Level Dynamic Entity
[**UpdateMyDynamicEntity**](DynamicEntityManageApi.md#UpdateMyDynamicEntity) | **Put** /obp/v5.1.0/my/dynamic-entities/{DYNAMIC_ENTITY_ID} | Update My Dynamic Entity
[**UpdateSystemDynamicEntity**](DynamicEntityManageApi.md#UpdateSystemDynamicEntity) | **Put** /obp/v5.1.0/management/system-dynamic-entities/{DYNAMIC_ENTITY_ID} | Update System Level Dynamic Entity


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

