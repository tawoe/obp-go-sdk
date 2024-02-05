# \CounterpartyApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCounterpartyCorporateLocation**](CounterpartyApi.md#AddCounterpartyCorporateLocation) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/other_accounts/{OTHER_ACCOUNT_ID}/metadata/corporate_location | Add Corporate Location to Counterparty
[**AddCounterpartyImageUrl**](CounterpartyApi.md#AddCounterpartyImageUrl) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/other_accounts/{OTHER_ACCOUNT_ID}/metadata/image_url | Add image url to other bank account
[**AddCounterpartyMoreInfo**](CounterpartyApi.md#AddCounterpartyMoreInfo) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/other_accounts/{OTHER_ACCOUNT_ID}/metadata/more_info | Add Counterparty More Info
[**AddCounterpartyOpenCorporatesUrl**](CounterpartyApi.md#AddCounterpartyOpenCorporatesUrl) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/other_accounts/{OTHER_ACCOUNT_ID}/metadata/open_corporates_url | Add Open Corporates URL to Counterparty
[**AddCounterpartyPhysicalLocation**](CounterpartyApi.md#AddCounterpartyPhysicalLocation) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/other_accounts/{OTHER_ACCOUNT_ID}/metadata/physical_location | Add physical location to other bank account
[**AddCounterpartyPublicAlias**](CounterpartyApi.md#AddCounterpartyPublicAlias) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/other_accounts/{OTHER_ACCOUNT_ID}/public_alias | Add public alias to other bank account
[**AddCounterpartyUrl**](CounterpartyApi.md#AddCounterpartyUrl) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/other_accounts/{OTHER_ACCOUNT_ID}/metadata/url | Add url to other bank account
[**AddOtherAccountPrivateAlias**](CounterpartyApi.md#AddOtherAccountPrivateAlias) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/other_accounts/{OTHER_ACCOUNT_ID}/private_alias | Create Other Account Private Alias
[**CreateCounterparty**](CounterpartyApi.md#CreateCounterparty) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/counterparties | Create Counterparty (Explicit)
[**CreateCounterpartyForAnyAccount**](CounterpartyApi.md#CreateCounterpartyForAnyAccount) | **Post** /obp/v5.1.0/management/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/counterparties | Create Counterparty for any account (Explicit)
[**DeleteCounterpartyCorporateLocation**](CounterpartyApi.md#DeleteCounterpartyCorporateLocation) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/other_accounts/{OTHER_ACCOUNT_ID}/metadata/corporate_location | Delete Counterparty Corporate Location
[**DeleteCounterpartyForAnyAccount**](CounterpartyApi.md#DeleteCounterpartyForAnyAccount) | **Delete** /obp/v5.1.0/management/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/counterparties/{COUNTERPARTY_ID} | Delete Counterparty for any account (Explicit)
[**DeleteCounterpartyImageUrl**](CounterpartyApi.md#DeleteCounterpartyImageUrl) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/other_accounts/{OTHER_ACCOUNT_ID}/metadata/image_url | Delete Counterparty Image URL
[**DeleteCounterpartyMoreInfo**](CounterpartyApi.md#DeleteCounterpartyMoreInfo) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/other_accounts/{OTHER_ACCOUNT_ID}/metadata/more_info | Delete more info of other bank account
[**DeleteCounterpartyOpenCorporatesUrl**](CounterpartyApi.md#DeleteCounterpartyOpenCorporatesUrl) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/other_accounts/{OTHER_ACCOUNT_ID}/metadata/open_corporates_url | Delete Counterparty Open Corporates URL
[**DeleteCounterpartyPhysicalLocation**](CounterpartyApi.md#DeleteCounterpartyPhysicalLocation) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/other_accounts/{OTHER_ACCOUNT_ID}/metadata/physical_location | Delete Counterparty Physical Location
[**DeleteCounterpartyPrivateAlias**](CounterpartyApi.md#DeleteCounterpartyPrivateAlias) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/other_accounts/{OTHER_ACCOUNT_ID}/private_alias | Delete Counterparty Private Alias
[**DeleteCounterpartyPublicAlias**](CounterpartyApi.md#DeleteCounterpartyPublicAlias) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/other_accounts/{OTHER_ACCOUNT_ID}/public_alias | Delete Counterparty Public Alias
[**DeleteCounterpartyUrl**](CounterpartyApi.md#DeleteCounterpartyUrl) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/other_accounts/{OTHER_ACCOUNT_ID}/metadata/url | Delete url of other bank account
[**DeleteExplicitCounterparty**](CounterpartyApi.md#DeleteExplicitCounterparty) | **Post** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/counterparties/{COUNTERPARTY_ID} | Delete Counterparty (Explicit)
[**GetCounterpartiesForAnyAccount**](CounterpartyApi.md#GetCounterpartiesForAnyAccount) | **Get** /obp/v5.1.0/management/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/counterparties | Get Counterparties for any account (Explicit)
[**GetCounterpartyByIdForAnyAccount**](CounterpartyApi.md#GetCounterpartyByIdForAnyAccount) | **Get** /obp/v5.1.0/management/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/counterparties/{COUNTERPARTY_ID} | Get Counterparty by Id for any account (Explicit) 
[**GetCounterpartyByNameForAnyAccount**](CounterpartyApi.md#GetCounterpartyByNameForAnyAccount) | **Get** /obp/v5.1.0/management/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/counterparty-names/{COUNTERPARTY_NAME} | Get Counterparty by name for any account (Explicit) 
[**GetCounterpartyPublicAlias**](CounterpartyApi.md#GetCounterpartyPublicAlias) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/other_accounts/{OTHER_ACCOUNT_ID}/public_alias | Get public alias of other bank account
[**GetExplictCounterpartiesForAccount**](CounterpartyApi.md#GetExplictCounterpartiesForAccount) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/counterparties | Get Counterparties (Explicit)
[**GetExplictCounterpartyById**](CounterpartyApi.md#GetExplictCounterpartyById) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/counterparties/{COUNTERPARTY_ID} | Get Counterparty by Id (Explicit)
[**GetOtherAccountByIdForBankAccount**](CounterpartyApi.md#GetOtherAccountByIdForBankAccount) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/other_accounts/{OTHER_ACCOUNT_ID} | Get Other Account by Id
[**GetOtherAccountForTransaction**](CounterpartyApi.md#GetOtherAccountForTransaction) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/transactions/{TRANSACTION_ID}/other_account | Get Other Account of Transaction
[**GetOtherAccountMetadata**](CounterpartyApi.md#GetOtherAccountMetadata) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/other_accounts/{OTHER_ACCOUNT_ID}/metadata | Get Other Account Metadata
[**GetOtherAccountPrivateAlias**](CounterpartyApi.md#GetOtherAccountPrivateAlias) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/other_accounts/{OTHER_ACCOUNT_ID}/private_alias | Get Other Account Private Alias
[**GetOtherAccountsForBankAccount**](CounterpartyApi.md#GetOtherAccountsForBankAccount) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/other_accounts | Get Other Accounts of one Account
[**UpdateCounterpartyCorporateLocation**](CounterpartyApi.md#UpdateCounterpartyCorporateLocation) | **Put** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/other_accounts/{OTHER_ACCOUNT_ID}/metadata/corporate_location | Update Counterparty Corporate Location
[**UpdateCounterpartyImageUrl**](CounterpartyApi.md#UpdateCounterpartyImageUrl) | **Put** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/other_accounts/{OTHER_ACCOUNT_ID}/metadata/image_url | Update Counterparty Image Url
[**UpdateCounterpartyMoreInfo**](CounterpartyApi.md#UpdateCounterpartyMoreInfo) | **Put** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/other_accounts/{OTHER_ACCOUNT_ID}/metadata/more_info | Update Counterparty More Info
[**UpdateCounterpartyOpenCorporatesUrl**](CounterpartyApi.md#UpdateCounterpartyOpenCorporatesUrl) | **Put** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/other_accounts/{OTHER_ACCOUNT_ID}/metadata/open_corporates_url | Update Open Corporates Url of Counterparty
[**UpdateCounterpartyPhysicalLocation**](CounterpartyApi.md#UpdateCounterpartyPhysicalLocation) | **Put** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/other_accounts/{OTHER_ACCOUNT_ID}/metadata/physical_location | Update Counterparty Physical Location
[**UpdateCounterpartyPrivateAlias**](CounterpartyApi.md#UpdateCounterpartyPrivateAlias) | **Put** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/other_accounts/{OTHER_ACCOUNT_ID}/private_alias | Update Counterparty Private Alias
[**UpdateCounterpartyPublicAlias**](CounterpartyApi.md#UpdateCounterpartyPublicAlias) | **Put** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/other_accounts/{OTHER_ACCOUNT_ID}/public_alias | Update public alias of other bank account
[**UpdateCounterpartyUrl**](CounterpartyApi.md#UpdateCounterpartyUrl) | **Put** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/{VIEW_ID}/other_accounts/{OTHER_ACCOUNT_ID}/metadata/url | Update url of other bank account


# **AddCounterpartyCorporateLocation**
> SuccessMessage AddCounterpartyCorporateLocation(ctx, body, oTHERACCOUNTID, vIEWID, aCCOUNTID, bANKID)
Add Corporate Location to Counterparty

<p>Add the geolocation of the counterparty's registered address</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CorporateLocationJson**](CorporateLocationJson.md)| CorporateLocationJSON object that needs to be added. | 
  **oTHERACCOUNTID** | **string**| The other account id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddCounterpartyImageUrl**
> SuccessMessage AddCounterpartyImageUrl(ctx, body, oTHERACCOUNTID, vIEWID, aCCOUNTID, bANKID)
Add image url to other bank account

<p>Add a url that points to the logo of the counterparty</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ImageUrlJson**](ImageUrlJson.md)| ImageUrlJSON object that needs to be added. | 
  **oTHERACCOUNTID** | **string**| The other account id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddCounterpartyMoreInfo**
> SuccessMessage AddCounterpartyMoreInfo(ctx, body, oTHERACCOUNTID, vIEWID, aCCOUNTID, bANKID)
Add Counterparty More Info

<p>Add a description of the counter party from the perpestive of the account e.g. My dentist</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MoreInfoJson**](MoreInfoJson.md)| MoreInfoJSON object that needs to be added. | 
  **oTHERACCOUNTID** | **string**| The other account id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddCounterpartyOpenCorporatesUrl**
> SuccessMessage AddCounterpartyOpenCorporatesUrl(ctx, body, oTHERACCOUNTID, vIEWID, aCCOUNTID, bANKID)
Add Open Corporates URL to Counterparty

<p>Add open corporates url to other bank account</p><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OpenCorporateUrlJson**](OpenCorporateUrlJson.md)| OpenCorporateUrlJSON object that needs to be added. | 
  **oTHERACCOUNTID** | **string**| The other account id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddCounterpartyPhysicalLocation**
> SuccessMessage AddCounterpartyPhysicalLocation(ctx, body, oTHERACCOUNTID, vIEWID, aCCOUNTID, bANKID)
Add physical location to other bank account

<p>Add geocoordinates of the counterparty's main location</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PhysicalLocationJson**](PhysicalLocationJson.md)| PhysicalLocationJSON object that needs to be added. | 
  **oTHERACCOUNTID** | **string**| The other account id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddCounterpartyPublicAlias**
> SuccessMessage AddCounterpartyPublicAlias(ctx, body, oTHERACCOUNTID, vIEWID, aCCOUNTID, bANKID)
Add public alias to other bank account

<p>Creates the public alias for the other account OTHER_ACCOUNT_ID.</p><p>Authentication is Optional<br />Authentication is required if the view is not public.</p><p>Note: Public aliases are automatically generated for new 'other accounts / counterparties', so this call should only be used if<br />the public alias was deleted.</p><p>The VIEW_ID parameter should be a view the caller is permitted to access to and that has permission to create public aliases.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AliasJson**](AliasJson.md)| AliasJSON object that needs to be added. | 
  **oTHERACCOUNTID** | **string**| The other account id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddCounterpartyUrl**
> SuccessMessage AddCounterpartyUrl(ctx, body, oTHERACCOUNTID, vIEWID, aCCOUNTID, bANKID)
Add url to other bank account

<p>A url which represents the counterparty (home page url etc.)</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UrlJson**](UrlJson.md)| UrlJSON object that needs to be added. | 
  **oTHERACCOUNTID** | **string**| The other account id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddOtherAccountPrivateAlias**
> SuccessMessage AddOtherAccountPrivateAlias(ctx, body, oTHERACCOUNTID, vIEWID, aCCOUNTID, bANKID)
Create Other Account Private Alias

<p>Creates a private alias for the other account OTHER_ACCOUNT_ID.</p><p>Authentication is Optional<br />Authentication is required if the view is not public.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AliasJson**](AliasJson.md)| AliasJSON object that needs to be added. | 
  **oTHERACCOUNTID** | **string**| The other account id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCounterparty**
> CounterpartyWithMetadataJson400 CreateCounterparty(ctx, body, vIEWID, aCCOUNTID, bANKID)
Create Counterparty (Explicit)

<p>Create Counterparty (Explicit) for an Account.</p><p>In OBP, there are two types of Counterparty.</p><ul><li><p>Explicit Counterparties (those here) which we create explicitly and are used in COUNTERPARTY Transaction Requests</p></li><li><p>Implicit Counterparties (AKA Other Accounts) which are generated automatically from the other sides of Transactions.</p></li></ul><p>Explicit Counterparties are created for the account / view<br />They are how the user of the view (e.g. account owner) refers to the other side of the transaction</p><p>name : the human readable name (e.g. Piano teacher, Miss Nipa)</p><p>description : the human readable name (e.g. Piano teacher, Miss Nipa)</p><p>currency : counterparty account currency (e.g. EUR, GBP, USD, ...)</p><p>bank_routing_scheme : eg: bankId or bankCode or any other strings</p><p>bank_routing_address : eg: <code>gh.29.uk</code>, must be valid sandbox bankIds</p><p>account_routing_scheme : eg: AccountId or AccountNumber or any other strings</p><p>account_routing_address : eg: <code>1d65db7c-a7b2-4839-af41-95</code>, must be valid accountIds</p><p>other_account_secondary_routing_scheme : eg: IBAN or any other strings</p><p>other_account_secondary_routing_address : if it is an IBAN, it should be unique for each counterparty.</p><p>other_branch_routing_scheme : eg: branchId or any other strings or you can leave it empty, not useful in sandbox mode.</p><p>other_branch_routing_address : eg: <code>branch-id-123</code> or you can leave it empty, not useful in sandbox mode.</p><p>is_beneficiary : must be set to <code>true</code> in order to send payments to this counterparty</p><p>bespoke: It supports a list of key-value, you can add it to the counterparty.</p><p>bespoke.key : any info-key you want to add to this counterparty</p><p>bespoke.value : any info-value you want to add to this counterparty</p><p>The view specified by VIEW_ID must have the canAddCounterparty permission</p><p>A minimal example for TransactionRequestType == COUNTERPARTY<br />{<br />&quot;name&quot;: &quot;Tesobe1&quot;,<br />&quot;description&quot;: &quot;Good Company&quot;,<br />&quot;currency&quot;: &quot;EUR&quot;,<br />&quot;other_bank_routing_scheme&quot;: &quot;OBP&quot;,<br />&quot;other_bank_routing_address&quot;: &quot;gh.29.uk&quot;,<br />&quot;other_account_routing_scheme&quot;: &quot;OBP&quot;,<br />&quot;other_account_routing_address&quot;: &quot;8ca8a7e4-6d02-48e3-a029-0b2bf89de9f0&quot;,<br />&quot;is_beneficiary&quot;: true,<br />&quot;other_account_secondary_routing_scheme&quot;: &quot;&quot;,<br />&quot;other_account_secondary_routing_address&quot;: &quot;&quot;,<br />&quot;other_branch_routing_scheme&quot;: &quot;&quot;,<br />&quot;other_branch_routing_address&quot;: &quot;&quot;,<br />&quot;bespoke&quot;: []<br />}</p><p>A minimal example for TransactionRequestType == SEPA</p><p>{<br />&quot;name&quot;: &quot;Tesobe2&quot;,<br />&quot;description&quot;: &quot;Good Company&quot;,<br />&quot;currency&quot;: &quot;EUR&quot;,<br />&quot;other_bank_routing_scheme&quot;: &quot;OBP&quot;,<br />&quot;other_bank_routing_address&quot;: &quot;gh.29.uk&quot;,<br />&quot;other_account_routing_scheme&quot;: &quot;OBP&quot;,<br />&quot;other_account_routing_address&quot;: &quot;8ca8a7e4-6d02-48e3-a029-0b2bf89de9f0&quot;,<br />&quot;other_account_secondary_routing_scheme&quot;: &quot;IBAN&quot;,<br />&quot;other_account_secondary_routing_address&quot;: &quot;DE89 3704 0044 0532 0130 00&quot;,<br />&quot;is_beneficiary&quot;: true,<br />&quot;other_branch_routing_scheme&quot;: &quot;&quot;,<br />&quot;other_branch_routing_address&quot;: &quot;&quot;,<br />&quot;bespoke&quot;: []<br />}</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostCounterpartyJson400**](PostCounterpartyJson400.md)| PostCounterpartyJson400 object that needs to be added. | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CounterpartyWithMetadataJson400**](CounterpartyWithMetadataJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCounterpartyForAnyAccount**
> CounterpartyWithMetadataJson400 CreateCounterpartyForAnyAccount(ctx, body, vIEWID, aCCOUNTID, bANKID)
Create Counterparty for any account (Explicit)

<p>Create Counterparty for any Account. (Explicit)</p><p>In OBP, there are two types of Counterparty.</p><ul><li><p>Explicit Counterparties (those here) which we create explicitly and are used in COUNTERPARTY Transaction Requests</p></li><li><p>Implicit Counterparties (AKA Other Accounts) which are generated automatically from the other sides of Transactions.</p></li></ul><p>Explicit Counterparties are created for the account / view<br />They are how the user of the view (e.g. account owner) refers to the other side of the transaction</p><p>name : the human readable name (e.g. Piano teacher, Miss Nipa)</p><p>description : the human readable name (e.g. Piano teacher, Miss Nipa)</p><p>currency : counterparty account currency (e.g. EUR, GBP, USD, ...)</p><p>bank_routing_scheme : eg: bankId or bankCode or any other strings</p><p>bank_routing_address : eg: <code>gh.29.uk</code>, must be valid sandbox bankIds</p><p>account_routing_scheme : eg: AccountId or AccountNumber or any other strings</p><p>account_routing_address : eg: <code>1d65db7c-a7b2-4839-af41-95</code>, must be valid accountIds</p><p>other_account_secondary_routing_scheme : eg: IBAN or any other strings</p><p>other_account_secondary_routing_address : if it is an IBAN, it should be unique for each counterparty.</p><p>other_branch_routing_scheme : eg: branchId or any other strings or you can leave it empty, not useful in sandbox mode.</p><p>other_branch_routing_address : eg: <code>branch-id-123</code> or you can leave it empty, not useful in sandbox mode.</p><p>is_beneficiary : must be set to <code>true</code> in order to send payments to this counterparty</p><p>bespoke: It supports a list of key-value, you can add it to the counterparty.</p><p>bespoke.key : any info-key you want to add to this counterparty</p><p>bespoke.value : any info-value you want to add to this counterparty</p><p>The view specified by VIEW_ID must have the canAddCounterparty permission</p><p>A minimal example for TransactionRequestType == COUNTERPARTY<br />{<br />&quot;name&quot;: &quot;Tesobe1&quot;,<br />&quot;description&quot;: &quot;Good Company&quot;,<br />&quot;currency&quot;: &quot;EUR&quot;,<br />&quot;other_bank_routing_scheme&quot;: &quot;OBP&quot;,<br />&quot;other_bank_routing_address&quot;: &quot;gh.29.uk&quot;,<br />&quot;other_account_routing_scheme&quot;: &quot;OBP&quot;,<br />&quot;other_account_routing_address&quot;: &quot;8ca8a7e4-6d02-48e3-a029-0b2bf89de9f0&quot;,<br />&quot;is_beneficiary&quot;: true,<br />&quot;other_account_secondary_routing_scheme&quot;: &quot;&quot;,<br />&quot;other_account_secondary_routing_address&quot;: &quot;&quot;,<br />&quot;other_branch_routing_scheme&quot;: &quot;&quot;,<br />&quot;other_branch_routing_address&quot;: &quot;&quot;,<br />&quot;bespoke&quot;: []<br />}</p><p>A minimal example for TransactionRequestType == SEPA</p><p>{<br />&quot;name&quot;: &quot;Tesobe2&quot;,<br />&quot;description&quot;: &quot;Good Company&quot;,<br />&quot;currency&quot;: &quot;EUR&quot;,<br />&quot;other_bank_routing_scheme&quot;: &quot;OBP&quot;,<br />&quot;other_bank_routing_address&quot;: &quot;gh.29.uk&quot;,<br />&quot;other_account_routing_scheme&quot;: &quot;OBP&quot;,<br />&quot;other_account_routing_address&quot;: &quot;8ca8a7e4-6d02-48e3-a029-0b2bf89de9f0&quot;,<br />&quot;other_account_secondary_routing_scheme&quot;: &quot;IBAN&quot;,<br />&quot;other_account_secondary_routing_address&quot;: &quot;DE89 3704 0044 0532 0130 00&quot;,<br />&quot;is_beneficiary&quot;: true,<br />&quot;other_branch_routing_scheme&quot;: &quot;&quot;,<br />&quot;other_branch_routing_address&quot;: &quot;&quot;,<br />&quot;bespoke&quot;: []<br />}</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostCounterpartyJson400**](PostCounterpartyJson400.md)| PostCounterpartyJson400 object that needs to be added. | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CounterpartyWithMetadataJson400**](CounterpartyWithMetadataJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCounterpartyCorporateLocation**
> EmptyClassJson DeleteCounterpartyCorporateLocation(ctx, body, oTHERACCOUNTID, vIEWID, aCCOUNTID, bANKID)
Delete Counterparty Corporate Location

<p>Delete corporate location of other bank account. Delete the geolocation of the counterparty's registered address</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **oTHERACCOUNTID** | **string**| The other account id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**EmptyClassJson**](EmptyClassJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCounterpartyForAnyAccount**
> DeleteCounterpartyForAnyAccount(ctx, cOUNTERPARTYID, vIEWID, aCCOUNTID, bANKID)
Delete Counterparty for any account (Explicit)

<p>Delete Counterparty (Explicit) for any account<br />and also delete the Metadata for its counterparty.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cOUNTERPARTYID** | **string**| the counterparty id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCounterpartyImageUrl**
> EmptyClassJson DeleteCounterpartyImageUrl(ctx, body, oTHERACCOUNTID, vIEWID, aCCOUNTID, bANKID)
Delete Counterparty Image URL

<p>Delete image url of other bank account</p><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **oTHERACCOUNTID** | **string**| The other account id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**EmptyClassJson**](EmptyClassJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCounterpartyMoreInfo**
> EmptyClassJson DeleteCounterpartyMoreInfo(ctx, body, oTHERACCOUNTID, vIEWID, aCCOUNTID, bANKID)
Delete more info of other bank account

<p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **oTHERACCOUNTID** | **string**| The other account id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**EmptyClassJson**](EmptyClassJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCounterpartyOpenCorporatesUrl**
> EmptyClassJson DeleteCounterpartyOpenCorporatesUrl(ctx, body, oTHERACCOUNTID, vIEWID, aCCOUNTID, bANKID)
Delete Counterparty Open Corporates URL

<p>Delete open corporate url of other bank account</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **oTHERACCOUNTID** | **string**| The other account id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**EmptyClassJson**](EmptyClassJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCounterpartyPhysicalLocation**
> EmptyClassJson DeleteCounterpartyPhysicalLocation(ctx, body, oTHERACCOUNTID, vIEWID, aCCOUNTID, bANKID)
Delete Counterparty Physical Location

<p>Delete physical location of other bank account</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **oTHERACCOUNTID** | **string**| The other account id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**EmptyClassJson**](EmptyClassJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCounterpartyPrivateAlias**
> EmptyClassJson DeleteCounterpartyPrivateAlias(ctx, body, oTHERACCOUNTID, vIEWID, aCCOUNTID, bANKID)
Delete Counterparty Private Alias

<p>Deletes the private alias of the other account OTHER_ACCOUNT_ID.</p><p>Authentication is Optional<br />Authentication is required if the view is not public.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **oTHERACCOUNTID** | **string**| The other account id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**EmptyClassJson**](EmptyClassJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCounterpartyPublicAlias**
> EmptyClassJson DeleteCounterpartyPublicAlias(ctx, body, oTHERACCOUNTID, vIEWID, aCCOUNTID, bANKID)
Delete Counterparty Public Alias

<p>Deletes the public alias of the other account OTHER_ACCOUNT_ID.</p><p>Authentication is Optional<br />Authentication is required if the view is not public.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **oTHERACCOUNTID** | **string**| The other account id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**EmptyClassJson**](EmptyClassJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCounterpartyUrl**
> EmptyClassJson DeleteCounterpartyUrl(ctx, body, oTHERACCOUNTID, vIEWID, aCCOUNTID, bANKID)
Delete url of other bank account

<p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **oTHERACCOUNTID** | **string**| The other account id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**EmptyClassJson**](EmptyClassJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteExplicitCounterparty**
> DeleteExplicitCounterparty(ctx, cOUNTERPARTYID, vIEWID, aCCOUNTID, bANKID)
Delete Counterparty (Explicit)

<p>Delete Counterparty (Explicit) for an Account.<br />and also delete the Metadata for its counterparty.</p><p>need the view permission <code>can_delete_counterparty</code><br />Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cOUNTERPARTYID** | **string**| the counterparty id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCounterpartiesForAnyAccount**
> CounterpartiesJson400 GetCounterpartiesForAnyAccount(ctx, vIEWID, aCCOUNTID, bANKID)
Get Counterparties for any account (Explicit)

<p>Get the Counterparties (Explicit) for any account .</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CounterpartiesJson400**](CounterpartiesJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCounterpartyByIdForAnyAccount**
> CounterpartyWithMetadataJson400 GetCounterpartyByIdForAnyAccount(ctx, cOUNTERPARTYID, vIEWID, aCCOUNTID, bANKID)
Get Counterparty by Id for any account (Explicit) 

<p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cOUNTERPARTYID** | **string**| the counterparty id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CounterpartyWithMetadataJson400**](CounterpartyWithMetadataJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCounterpartyByNameForAnyAccount**
> CounterpartyWithMetadataJson400 GetCounterpartyByNameForAnyAccount(ctx, cOUNTERPARTYNAME, vIEWID, aCCOUNTID, bANKID)
Get Counterparty by name for any account (Explicit) 

<p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cOUNTERPARTYNAME** | **string**| the counterparty name | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CounterpartyWithMetadataJson400**](CounterpartyWithMetadataJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCounterpartyPublicAlias**
> AliasJson GetCounterpartyPublicAlias(ctx, body, oTHERACCOUNTID, vIEWID, aCCOUNTID, bANKID)
Get public alias of other bank account

<p>Returns the public alias of the other account OTHER_ACCOUNT_ID.<br />Authentication is Optional<br />Authentication is Mandatory if the view is not public.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **oTHERACCOUNTID** | **string**| The other account id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**AliasJson**](AliasJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExplictCounterpartiesForAccount**
> CounterpartiesJson400 GetExplictCounterpartiesForAccount(ctx, vIEWID, aCCOUNTID, bANKID)
Get Counterparties (Explicit)

<p>Get the Counterparties (Explicit) for the account / view.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CounterpartiesJson400**](CounterpartiesJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExplictCounterpartyById**
> CounterpartyWithMetadataJson400 GetExplictCounterpartyById(ctx, cOUNTERPARTYID, vIEWID, aCCOUNTID, bANKID)
Get Counterparty by Id (Explicit)

<p>Information returned about the Counterparty specified by COUNTERPARTY_ID:</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cOUNTERPARTYID** | **string**| the counterparty id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CounterpartyWithMetadataJson400**](CounterpartyWithMetadataJson400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOtherAccountByIdForBankAccount**
> OtherAccountJsonV300 GetOtherAccountByIdForBankAccount(ctx, oTHERACCOUNTID, vIEWID, aCCOUNTID, bANKID)
Get Other Account by Id

<p>Returns data about the Other Account that has shared at least one transaction with ACCOUNT_ID at BANK_ID.<br />Authentication is Optional</p><p>Authentication is required if the view is not public.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oTHERACCOUNTID** | **string**| The other account id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**OtherAccountJsonV300**](OtherAccountJsonV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOtherAccountForTransaction**
> OtherAccountJson GetOtherAccountForTransaction(ctx, body, tRANSACTIONID, vIEWID, aCCOUNTID, bANKID)
Get Other Account of Transaction

<p>Get other account of a transaction.<br />Returns details of the other party involved in the transaction, moderated by the <a href=\"#1_2_1-getViewsForBankAccount\">view</a> (VIEW_ID).<br />Authentication via OAuth is required if the view is not public.</p><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **tRANSACTIONID** | **string**| The transaction id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**OtherAccountJson**](OtherAccountJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOtherAccountMetadata**
> OtherAccountMetadataJson GetOtherAccountMetadata(ctx, body, oTHERACCOUNTID, vIEWID, aCCOUNTID, bANKID)
Get Other Account Metadata

<p>Get metadata of one other account.<br />Returns only the metadata about one other bank account (OTHER_ACCOUNT_ID) that had shared at least one transaction with ACCOUNT_ID at BANK_ID.</p><p>Authentication via OAuth is required if the view is not public.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **oTHERACCOUNTID** | **string**| The other account id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**OtherAccountMetadataJson**](OtherAccountMetadataJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOtherAccountPrivateAlias**
> AliasJson GetOtherAccountPrivateAlias(ctx, body, oTHERACCOUNTID, vIEWID, aCCOUNTID, bANKID)
Get Other Account Private Alias

<p>Returns the private alias of the other account OTHER_ACCOUNT_ID.</p><p>Authentication is Optional<br />Authentication is required if the view is not public.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **oTHERACCOUNTID** | **string**| The other account id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**AliasJson**](AliasJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOtherAccountsForBankAccount**
> OtherAccountsJsonV300 GetOtherAccountsForBankAccount(ctx, vIEWID, aCCOUNTID, bANKID)
Get Other Accounts of one Account

<p>Returns data about all the other accounts that have shared at least one transaction with the ACCOUNT_ID at BANK_ID.<br />Authentication is Optional</p><p>Authentication is required if the view VIEW_ID is not public.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**OtherAccountsJsonV300**](OtherAccountsJsonV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCounterpartyCorporateLocation**
> SuccessMessage UpdateCounterpartyCorporateLocation(ctx, body, oTHERACCOUNTID, vIEWID, aCCOUNTID, bANKID)
Update Counterparty Corporate Location

<p>Update the geolocation of the counterparty's registered address</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CorporateLocationJson**](CorporateLocationJson.md)| CorporateLocationJSON object that needs to be added. | 
  **oTHERACCOUNTID** | **string**| The other account id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCounterpartyImageUrl**
> SuccessMessage UpdateCounterpartyImageUrl(ctx, body, oTHERACCOUNTID, vIEWID, aCCOUNTID, bANKID)
Update Counterparty Image Url

<p>Update the url that points to the logo of the counterparty</p><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ImageUrlJson**](ImageUrlJson.md)| ImageUrlJSON object that needs to be added. | 
  **oTHERACCOUNTID** | **string**| The other account id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCounterpartyMoreInfo**
> SuccessMessage UpdateCounterpartyMoreInfo(ctx, body, oTHERACCOUNTID, vIEWID, aCCOUNTID, bANKID)
Update Counterparty More Info

<p>Update the more info description of the counter party from the perpestive of the account e.g. My dentist</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MoreInfoJson**](MoreInfoJson.md)| MoreInfoJSON object that needs to be added. | 
  **oTHERACCOUNTID** | **string**| The other account id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCounterpartyOpenCorporatesUrl**
> SuccessMessage UpdateCounterpartyOpenCorporatesUrl(ctx, body, oTHERACCOUNTID, vIEWID, aCCOUNTID, bANKID)
Update Open Corporates Url of Counterparty

<p>Update open corporate url of other bank account</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OpenCorporateUrlJson**](OpenCorporateUrlJson.md)| OpenCorporateUrlJSON object that needs to be added. | 
  **oTHERACCOUNTID** | **string**| The other account id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCounterpartyPhysicalLocation**
> SuccessMessage UpdateCounterpartyPhysicalLocation(ctx, body, oTHERACCOUNTID, vIEWID, aCCOUNTID, bANKID)
Update Counterparty Physical Location

<p>Update geocoordinates of the counterparty's main location</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PhysicalLocationJson**](PhysicalLocationJson.md)| PhysicalLocationJSON object that needs to be added. | 
  **oTHERACCOUNTID** | **string**| The other account id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCounterpartyPrivateAlias**
> SuccessMessage UpdateCounterpartyPrivateAlias(ctx, body, oTHERACCOUNTID, vIEWID, aCCOUNTID, bANKID)
Update Counterparty Private Alias

<p>Updates the private alias of the counterparty (AKA other account) OTHER_ACCOUNT_ID.</p><p>Authentication is Optional<br />Authentication is required if the view is not public.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AliasJson**](AliasJson.md)| AliasJSON object that needs to be added. | 
  **oTHERACCOUNTID** | **string**| The other account id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCounterpartyPublicAlias**
> SuccessMessage UpdateCounterpartyPublicAlias(ctx, body, oTHERACCOUNTID, vIEWID, aCCOUNTID, bANKID)
Update public alias of other bank account

<p>Updates the public alias of the other account / counterparty OTHER_ACCOUNT_ID.</p><p>Authentication is Optional<br />Authentication is required if the view is not public.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AliasJson**](AliasJson.md)| AliasJSON object that needs to be added. | 
  **oTHERACCOUNTID** | **string**| The other account id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCounterpartyUrl**
> SuccessMessage UpdateCounterpartyUrl(ctx, body, oTHERACCOUNTID, vIEWID, aCCOUNTID, bANKID)
Update url of other bank account

<p>A url which represents the counterparty (home page url etc.)</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UrlJson**](UrlJson.md)| UrlJSON object that needs to be added. | 
  **oTHERACCOUNTID** | **string**| The other account id | 
  **vIEWID** | **string**| The view id | 
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

