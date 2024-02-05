# \KYCApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddKycCheck**](KYCApi.md#AddKycCheck) | **Put** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/kyc_check/{KYC_CHECK_ID} | Add KYC Check
[**AddKycDocument**](KYCApi.md#AddKycDocument) | **Put** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/kyc_documents/{KYC_DOCUMENT_ID} | Add KYC Document
[**AddKycMedia**](KYCApi.md#AddKycMedia) | **Put** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/kyc_media/{KYC_MEDIA_ID} | Add KYC Media
[**AddKycStatus**](KYCApi.md#AddKycStatus) | **Put** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/kyc_statuses | Add KYC Status
[**CreateTaxResidence**](KYCApi.md#CreateTaxResidence) | **Post** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/tax-residence | Create Tax Residence
[**CreateUserInvitation**](KYCApi.md#CreateUserInvitation) | **Post** /obp/v5.1.0/banks/{BANK_ID}/user-invitation | Create User Invitation
[**DeleteCustomerAddress**](KYCApi.md#DeleteCustomerAddress) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/addresses/{CUSTOMER_ADDRESS_ID} | Delete Customer Address
[**DeleteTaxResidence**](KYCApi.md#DeleteTaxResidence) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/tax_residencies/{TAX_RESIDENCE_ID} | Delete Tax Residence
[**GetCustomerAddresses**](KYCApi.md#GetCustomerAddresses) | **Get** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/addresses | Get Customer Addresses
[**GetCustomerByCustomerNumber**](KYCApi.md#GetCustomerByCustomerNumber) | **Post** /obp/v5.1.0/banks/{BANK_ID}/customers/customer-number | Get Customer by CUSTOMER_NUMBER
[**GetCustomerOverview**](KYCApi.md#GetCustomerOverview) | **Post** /obp/v5.1.0/banks/{BANK_ID}/customers/customer-number-query/overview | Get Customer Overview
[**GetCustomerOverviewFlat**](KYCApi.md#GetCustomerOverviewFlat) | **Post** /obp/v5.1.0/banks/{BANK_ID}/customers/customer-number-query/overview-flat | Get Customer Overview Flat
[**GetCustomersByCustomerPhoneNumber**](KYCApi.md#GetCustomersByCustomerPhoneNumber) | **Post** /obp/v5.1.0/banks/{BANK_ID}/search/customers/mobile-phone-number | Get Customers by MOBILE_PHONE_NUMBER
[**GetKycChecks**](KYCApi.md#GetKycChecks) | **Get** /obp/v5.1.0/customers/{CUSTOMER_ID}/kyc_checks | Get Customer KYC Checks
[**GetKycDocuments**](KYCApi.md#GetKycDocuments) | **Get** /obp/v5.1.0/customers/{CUSTOMER_ID}/kyc_documents | Get Customer KYC Documents
[**GetKycMedia**](KYCApi.md#GetKycMedia) | **Get** /obp/v5.1.0/customers/{CUSTOMER_ID}/kyc_media | Get KYC Media for a customer
[**GetKycStatuses**](KYCApi.md#GetKycStatuses) | **Get** /obp/v5.1.0/customers/{CUSTOMER_ID}/kyc_statuses | Get Customer KYC statuses
[**GetTaxResidence**](KYCApi.md#GetTaxResidence) | **Get** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/tax-residences | Get Tax Residences of Customer
[**GetUserInvitationAnonymous**](KYCApi.md#GetUserInvitationAnonymous) | **Post** /obp/v5.1.0/banks/{BANK_ID}/user-invitations | Get User Invitation Information


# **AddKycCheck**
> KycCheckJson AddKycCheck(ctx, body, kYCCHECKID, cUSTOMERID, bANKID)
Add KYC Check

<p>Add a KYC check for the customer specified by CUSTOMER_ID. KYC Checks store details of checks on a customer made by the KYC team, their comments and a satisfied status</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostKycCheckJson**](PostKycCheckJson.md)| PostKycCheckJSON object that needs to be added. | 
  **kYCCHECKID** | **string**| The kyc check id | 
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**KycCheckJson**](KycCheckJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddKycDocument**
> KycDocumentJson AddKycDocument(ctx, body, kYCDOCUMENTID, cUSTOMERID, bANKID)
Add KYC Document

<p>Add a KYC document for the customer specified by CUSTOMER_ID. KYC Documents contain the document type (e.g. passport), place of issue, expiry etc.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostKycDocumentJson**](PostKycDocumentJson.md)| PostKycDocumentJSON object that needs to be added. | 
  **kYCDOCUMENTID** | **string**| The kyc document id | 
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**KycDocumentJson**](KycDocumentJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddKycMedia**
> KycMediaJson AddKycMedia(ctx, body, kYCMEDIAID, cUSTOMERID, bANKID)
Add KYC Media

<p>Add some KYC media for the customer specified by CUSTOMER_ID. KYC Media resources relate to KYC Documents and KYC Checks and contain media urls for scans of passports, utility bills etc</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostKycMediaJson**](PostKycMediaJson.md)| PostKycMediaJSON object that needs to be added. | 
  **kYCMEDIAID** | **string**| The kyc media id | 
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**KycMediaJson**](KycMediaJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddKycStatus**
> KycStatusJson AddKycStatus(ctx, body, cUSTOMERID, bANKID)
Add KYC Status

<p>Add a kyc_status for the customer specified by CUSTOMER_ID. KYC Status is a timeline of the KYC status of the customer</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostKycStatusJson**](PostKycStatusJson.md)| PostKycStatusJSON object that needs to be added. | 
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**KycStatusJson**](KycStatusJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTaxResidence**
> TaxResidenceV310 CreateTaxResidence(ctx, body, cUSTOMERID, bANKID)
Create Tax Residence

<p>Create a Tax Residence for a Customer specified by CUSTOMER_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostTaxResidenceJsonV310**](PostTaxResidenceJsonV310.md)| PostTaxResidenceJsonV310 object that needs to be added. | 
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**TaxResidenceV310**](TaxResidenceV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUserInvitation**
> UserInvitationJsonV400 CreateUserInvitation(ctx, body, bANKID)
Create User Invitation

<p>Create User Invitation.</p><p>This endpoint will send an invitation email to the developers, then they can use the link to create the obp user.</p><p>purpose filed only support:List(DEVELOPER, CUSTOMER).</p><p>You can customise the email details use the following webui props:</p><p>when purpose == DEVELOPER<br />webui_developer_user_invitation_email_subject<br />webui_developer_user_invitation_email_from<br />webui_developer_user_invitation_email_text<br />webui_developer_user_invitation_email_html_text</p><p>when purpose = == CUSTOMER<br />webui_customer_user_invitation_email_subject<br />webui_customer_user_invitation_email_from<br />webui_customer_user_invitation_email_text<br />webui_customer_user_invitation_email_html_text</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostUserInvitationJsonV400**](PostUserInvitationJsonV400.md)| PostUserInvitationJsonV400 object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**UserInvitationJsonV400**](UserInvitationJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCustomerAddress**
> DeleteCustomerAddress(ctx, cUSTOMERADDRESSID, cUSTOMERID, bANKID)
Delete Customer Address

<p>Delete an Address of the Customer specified by CUSTOMER_ADDRESS_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cUSTOMERADDRESSID** | **string**| the customer address id | 
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTaxResidence**
> DeleteTaxResidence(ctx, tAXRESIDENCEID, cUSTOMERID, bANKID)
Delete Tax Residence

<p>Delete a Tax Residence of the Customer specified by TAX_RESIDENCE_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tAXRESIDENCEID** | **string**| the tax residence id | 
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomerAddresses**
> CustomerAddressesJsonV310 GetCustomerAddresses(ctx, cUSTOMERID, bANKID)
Get Customer Addresses

<p>Get the Addresses of the Customer specified by CUSTOMER_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerAddressesJsonV310**](CustomerAddressesJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomerByCustomerNumber**
> CustomerWithAttributesJsonV310 GetCustomerByCustomerNumber(ctx, body, bANKID)
Get Customer by CUSTOMER_NUMBER

<p>Gets the Customer specified by CUSTOMER_NUMBER.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostCustomerNumberJsonV310**](PostCustomerNumberJsonV310.md)| PostCustomerNumberJsonV310 object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerWithAttributesJsonV310**](CustomerWithAttributesJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomerOverview**
> CustomerOverviewJsonV500 GetCustomerOverview(ctx, body, bANKID)
Get Customer Overview

<p>Gets the Customer Overview specified by customer_number and bank_code.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostCustomerOverviewJsonV500**](PostCustomerOverviewJsonV500.md)| PostCustomerOverviewJsonV500 object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerOverviewJsonV500**](CustomerOverviewJsonV500.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomerOverviewFlat**
> CustomerOverviewFlatJsonV500 GetCustomerOverviewFlat(ctx, body, bANKID)
Get Customer Overview Flat

<p>Gets the Customer Overview Flat specified by customer_number and bank_code.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostCustomerOverviewJsonV500**](PostCustomerOverviewJsonV500.md)| PostCustomerOverviewJsonV500 object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerOverviewFlatJsonV500**](CustomerOverviewFlatJsonV500.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomersByCustomerPhoneNumber**
> CustomerJsonV310 GetCustomersByCustomerPhoneNumber(ctx, body, bANKID)
Get Customers by MOBILE_PHONE_NUMBER

<p>Gets the Customers specified by MOBILE_PHONE_NUMBER.</p><p>There are two wildcards often used in conjunction with the LIKE operator:<br />% - The percent sign represents zero, one, or multiple characters<br />_ - The underscore represents a single character<br />For example {&quot;customer_phone_number&quot;:&quot;%381%&quot;} lists all numbers which contain 381 sequence</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostCustomerPhoneNumberJsonV400**](PostCustomerPhoneNumberJsonV400.md)| PostCustomerPhoneNumberJsonV400 object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerJsonV310**](CustomerJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKycChecks**
> KycChecksJson GetKycChecks(ctx, body, cUSTOMERID)
Get Customer KYC Checks

<p>Get KYC checks for the Customer specified by CUSTOMER_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **cUSTOMERID** | **string**| The customer id | 

### Return type

[**KycChecksJson**](KycChecksJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKycDocuments**
> KycDocumentsJson GetKycDocuments(ctx, body, cUSTOMERID)
Get Customer KYC Documents

<p>Get KYC (know your customer) documents for a customer specified by CUSTOMER_ID<br />Get a list of documents that affirm the identity of the customer<br />Passport, driving licence etc.<br />Authentication is Optional</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **cUSTOMERID** | **string**| The customer id | 

### Return type

[**KycDocumentsJson**](KycDocumentsJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKycMedia**
> KycMediasJson GetKycMedia(ctx, body, cUSTOMERID)
Get KYC Media for a customer

<p>Get KYC media (scans, pictures, videos) that affirms the identity of the customer.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **cUSTOMERID** | **string**| The customer id | 

### Return type

[**KycMediasJson**](KycMediasJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKycStatuses**
> KycStatusesJson GetKycStatuses(ctx, body, cUSTOMERID)
Get Customer KYC statuses

<p>Get the KYC statuses for a customer specified by CUSTOMER_ID over time.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **cUSTOMERID** | **string**| The customer id | 

### Return type

[**KycStatusesJson**](KycStatusesJSON.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTaxResidence**
> TaxResidenceJsonV310 GetTaxResidence(ctx, cUSTOMERID, bANKID)
Get Tax Residences of Customer

<p>Get the Tax Residences of the Customer specified by CUSTOMER_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**TaxResidenceJsonV310**](TaxResidenceJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserInvitationAnonymous**
> UserInvitationJsonV400 GetUserInvitationAnonymous(ctx, body, bANKID)
Get User Invitation Information

<p>Create User Invitation Information.</p><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostUserInvitationAnonymousJsonV400**](PostUserInvitationAnonymousJsonV400.md)| PostUserInvitationAnonymousJsonV400 object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**UserInvitationJsonV400**](UserInvitationJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

