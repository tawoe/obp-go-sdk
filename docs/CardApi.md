# \CardApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCardForBank**](CardApi.md#AddCardForBank) | **Post** /obp/v5.1.0/management/banks/{BANK_ID}/cards | Create Card
[**CreateCardAttribute**](CardApi.md#CreateCardAttribute) | **Post** /obp/v5.1.0/management/banks/{BANK_ID}/cards/{CARD_ID}/attribute | Create Card Attribute
[**CreateOrUpdateCardAttributeDefinition**](CardApi.md#CreateOrUpdateCardAttributeDefinition) | **Put** /obp/v5.1.0/banks/{BANK_ID}/attribute-definitions/card | Create or Update Card Attribute Definition
[**DeleteCardAttributeDefinition**](CardApi.md#DeleteCardAttributeDefinition) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/attribute-definitions/ATTRIBUTE_DEFINITION_ID/card | Delete Card Attribute Definition
[**DeleteCardForBank**](CardApi.md#DeleteCardForBank) | **Delete** /obp/v5.1.0/management/banks/{BANK_ID}/cards/{CARD_ID} | Delete Card
[**GetCardAttributeDefinition**](CardApi.md#GetCardAttributeDefinition) | **Get** /obp/v5.1.0/banks/{BANK_ID}/attribute-definitions/card | Get Card Attribute Definition
[**GetCardForBank**](CardApi.md#GetCardForBank) | **Get** /obp/v5.1.0/management/banks/{BANK_ID}/cards/{CARD_ID} | Get Card By Id
[**GetCards**](CardApi.md#GetCards) | **Get** /obp/v5.1.0/cards | Get cards for the current user
[**GetCardsForBank**](CardApi.md#GetCardsForBank) | **Get** /obp/v5.1.0/management/banks/{BANK_ID}/cards | Get Cards for the specified bank
[**GetStatusOfCreditCardOrder**](CardApi.md#GetStatusOfCreditCardOrder) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/credit_cards/orders | Get status of Credit Card order 
[**UpdateCardAttribute**](CardApi.md#UpdateCardAttribute) | **Put** /obp/v5.1.0/management/banks/{BANK_ID}/cards/{CARD_ID}/attributes/{CARD_ATTRIBUTE_ID} | Update Card Attribute
[**UpdatedCardForBank**](CardApi.md#UpdatedCardForBank) | **Put** /obp/v5.1.0/management/banks/{BANK_ID}/cards/{CARD_ID} | Update Card


# **AddCardForBank**
> PhysicalCardJsonV500 AddCardForBank(ctx, body, bANKID)
Create Card

<p>Create Card at bank specified by BANK_ID .</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreatePhysicalCardJsonV500**](CreatePhysicalCardJsonV500.md)| CreatePhysicalCardJsonV500 object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**PhysicalCardJsonV500**](PhysicalCardJsonV500.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCardAttribute**
> CardAttributeCommons CreateCardAttribute(ctx, body, cARDID, bANKID)
Create Card Attribute

<p>Create Card Attribute</p><p>Card Attributes are used to describe a financial Product with a list of typed key value pairs.</p><p>Each Card Attribute is linked to its Card by CARD_ID</p><p>The type field must be one of &quot;STRING&quot;, &quot;INTEGER&quot;, &quot;DOUBLE&quot; or DATE_WITH_DAY&quot;</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CardAttributeJson**](CardAttributeJson.md)| CardAttributeJson object that needs to be added. | 
  **cARDID** | **string**| the card id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CardAttributeCommons**](CardAttributeCommons.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateCardAttributeDefinition**
> AttributeDefinitionResponseJsonV400 CreateOrUpdateCardAttributeDefinition(ctx, body, bANKID)
Create or Update Card Attribute Definition

<p>Create or Update Card Attribute Definition</p><p>The category field must be Card</p><p>The type field must be one of; DOUBLE, STRING, INTEGER and DATE_WITH_DAY</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AttributeDefinitionJsonV400**](AttributeDefinitionJsonV400.md)| AttributeDefinitionJsonV400 object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**AttributeDefinitionResponseJsonV400**](AttributeDefinitionResponseJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCardAttributeDefinition**
> DeleteCardAttributeDefinition(ctx, bANKID)
Delete Card Attribute Definition

<p>Delete Card Attribute Definition by ATTRIBUTE_DEFINITION_ID</p><p>Authentication is Mandatory</p>

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

# **DeleteCardForBank**
> DeleteCardForBank(ctx, cARDID, bANKID)
Delete Card

<p>Delete a Card at bank specified by CARD_ID .</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cARDID** | **string**| the card id | 
  **bANKID** | **string**| The bank id | 

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCardAttributeDefinition**
> AttributeDefinitionsResponseJsonV400 GetCardAttributeDefinition(ctx, bANKID)
Get Card Attribute Definition

<p>Get Card Attribute Definition</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**AttributeDefinitionsResponseJsonV400**](AttributeDefinitionsResponseJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCardForBank**
> PhysicalCardWithAttributesJsonV310 GetCardForBank(ctx, cARDID, bANKID)
Get Card By Id

<p>This will the datails of the card.<br />It shows the account infomation which linked the the card.<br />Also shows the card attributes of the card.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cARDID** | **string**| the card id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**PhysicalCardWithAttributesJsonV310**](PhysicalCardWithAttributesJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCards**
> PhysicalCardsJson GetCards(ctx, body)
Get cards for the current user

<p>Returns data about all the physical cards a user has been issued. These could be debit cards, credit cards, etc.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 

### Return type

[**PhysicalCardsJson**](PhysicalCardsJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCardsForBank**
> PhysicalCardsJsonV310 GetCardsForBank(ctx, bANKID)
Get Cards for the specified bank

<p>Should be able to filter on the following fields</p><p>eg:/management/banks/BANK_ID/cards?customer_id=66214b8e-259e-44ad-8868-3eb47be70646&amp;account_id=8ca8a7e4-6d02-48e3-a029-0b2bf89de9f0</p><p>1 customer_id should be valid customer_id, otherwise, it will return an empty card list.</p><p>2 account_id should be valid account_id , otherwise, it will return an empty card list.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**PhysicalCardsJsonV310**](PhysicalCardsJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStatusOfCreditCardOrder**
> CreditCardOrderStatusResponseJson GetStatusOfCreditCardOrder(ctx, vIEWID, aCCOUNTID, bANKID)
Get status of Credit Card order 

<pre><code>  Get status of Credit Card orders</code></pre><p>Get all orders</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CreditCardOrderStatusResponseJson**](CreditCardOrderStatusResponseJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCardAttribute**
> CardAttributeCommons UpdateCardAttribute(ctx, body, cARDATTRIBUTEID, cARDID, bANKID)
Update Card Attribute

<p>Update Card Attribute</p><p>Card Attributes are used to describe a financial Product with a list of typed key value pairs.</p><p>Each Card Attribute is linked to its Card by CARD_ID</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CardAttributeJson**](CardAttributeJson.md)| CardAttributeJson object that needs to be added. | 
  **cARDATTRIBUTEID** | **string**| the card attribute id | 
  **cARDID** | **string**| the card id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CardAttributeCommons**](CardAttributeCommons.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatedCardForBank**
> PhysicalCardJsonV310 UpdatedCardForBank(ctx, body, cARDID, bANKID)
Update Card

<p>Update Card at bank specified by CARD_ID .<br />Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdatePhysicalCardJsonV310**](UpdatePhysicalCardJsonV310.md)| UpdatePhysicalCardJsonV310 object that needs to be added. | 
  **cARDID** | **string**| the card id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**PhysicalCardJsonV310**](PhysicalCardJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

