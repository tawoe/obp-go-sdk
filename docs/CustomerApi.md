# \CustomerApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCustomerMessage**](CustomerApi.md#AddCustomerMessage) | **Post** /obp/v5.1.0/banks/{BANK_ID}/customer/{CUSTOMER_ID}/messages | Create Customer Message
[**AddKycCheck**](CustomerApi.md#AddKycCheck) | **Put** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/kyc_check/{KYC_CHECK_ID} | Add KYC Check
[**AddKycDocument**](CustomerApi.md#AddKycDocument) | **Put** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/kyc_documents/{KYC_DOCUMENT_ID} | Add KYC Document
[**AddKycMedia**](CustomerApi.md#AddKycMedia) | **Put** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/kyc_media/{KYC_MEDIA_ID} | Add KYC Media
[**AddKycStatus**](CustomerApi.md#AddKycStatus) | **Put** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/kyc_statuses | Add KYC Status
[**AddSocialMediaHandle**](CustomerApi.md#AddSocialMediaHandle) | **Post** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/social_media_handles | Create Customer Social Media Handle
[**CreateCustomer**](CustomerApi.md#CreateCustomer) | **Post** /obp/v5.1.0/banks/{BANK_ID}/customers | Create Customer
[**CreateCustomerAccountLink**](CustomerApi.md#CreateCustomerAccountLink) | **Post** /obp/v5.1.0/banks/{BANK_ID}/customer-account-links | Create Customer Account Link
[**CreateCustomerAddress**](CustomerApi.md#CreateCustomerAddress) | **Post** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/address | Create Address
[**CreateCustomerAttribute**](CustomerApi.md#CreateCustomerAttribute) | **Post** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/attribute | Create Customer Attribute
[**CreateCustomerMessage**](CustomerApi.md#CreateCustomerMessage) | **Post** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/messages | Create Customer Message
[**CreateMeeting**](CustomerApi.md#CreateMeeting) | **Post** /obp/v5.1.0/banks/{BANK_ID}/meetings | Create Meeting (video conference/call)
[**CreateOrUpdateCustomerAttributeAttributeDefinition**](CustomerApi.md#CreateOrUpdateCustomerAttributeAttributeDefinition) | **Put** /obp/v5.1.0/banks/{BANK_ID}/attribute-definitions/customer | Create or Update Customer Attribute Definition
[**CreateTaxResidence**](CustomerApi.md#CreateTaxResidence) | **Post** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/tax-residence | Create Tax Residence
[**CreateUserCustomerLinks**](CustomerApi.md#CreateUserCustomerLinks) | **Post** /obp/v5.1.0/banks/{BANK_ID}/user_customer_links | Create User Customer Link
[**DeleteCustomerAccountLinkById**](CustomerApi.md#DeleteCustomerAccountLinkById) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/customer-account-links/CUSTOMER_ACCOUNT_LINK_ID | Delete Customer Account Link
[**DeleteCustomerAddress**](CustomerApi.md#DeleteCustomerAddress) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/addresses/{CUSTOMER_ADDRESS_ID} | Delete Customer Address
[**DeleteCustomerAttribute**](CustomerApi.md#DeleteCustomerAttribute) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/{CUSTOMER_ID}/attributes/CUSTOMER_ATTRIBUTE_ID | Delete Customer Attribute
[**DeleteCustomerAttributeDefinition**](CustomerApi.md#DeleteCustomerAttributeDefinition) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/attribute-definitions/ATTRIBUTE_DEFINITION_ID/customer | Delete Customer Attribute Definition
[**DeleteCustomerCascade**](CustomerApi.md#DeleteCustomerCascade) | **Delete** /obp/v5.1.0/management/cascading/banks/{BANK_ID}/customers/{CUSTOMER_ID} | Delete Customer Cascade
[**DeleteTaxResidence**](CustomerApi.md#DeleteTaxResidence) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/tax_residencies/{TAX_RESIDENCE_ID} | Delete Tax Residence
[**DeleteUserCustomerLink**](CustomerApi.md#DeleteUserCustomerLink) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/user_customer_links/USER_CUSTOMER_LINK_ID | Delete User Customer Link
[**GetCorrelatedUsersInfoByCustomerId**](CustomerApi.md#GetCorrelatedUsersInfoByCustomerId) | **Get** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/correlated-users | Get Correlated User Info by Customer
[**GetCrmEvents**](CustomerApi.md#GetCrmEvents) | **Get** /obp/v5.1.0/banks/{BANK_ID}/crm-events | Get CRM Events
[**GetCustomerAccountLinkById**](CustomerApi.md#GetCustomerAccountLinkById) | **Get** /obp/v5.1.0/banks/{BANK_ID}/customer-account-links/CUSTOMER_ACCOUNT_LINK_ID | Get Customer Account Link by Id
[**GetCustomerAccountLinksByBankIdAccountId**](CustomerApi.md#GetCustomerAccountLinksByBankIdAccountId) | **Get** /obp/v5.1.0/banks/{BANK_ID}/accounts/{ACCOUNT_ID}/customer-account-links | Get Customer Account Links by ACCOUNT_ID
[**GetCustomerAccountLinksByCustomerId**](CustomerApi.md#GetCustomerAccountLinksByCustomerId) | **Get** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/customer-account-links | Get Customer Account Links by CUSTOMER_ID
[**GetCustomerAddresses**](CustomerApi.md#GetCustomerAddresses) | **Get** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/addresses | Get Customer Addresses
[**GetCustomerAttributeById**](CustomerApi.md#GetCustomerAttributeById) | **Get** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/attributes/ATTRIBUTE_ID | Get Customer Attribute By Id
[**GetCustomerAttributeDefinition**](CustomerApi.md#GetCustomerAttributeDefinition) | **Get** /obp/v5.1.0/banks/{BANK_ID}/attribute-definitions/customer | Get Customer Attribute Definition
[**GetCustomerAttributes**](CustomerApi.md#GetCustomerAttributes) | **Get** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/attributes | Get Customer Attributes
[**GetCustomerByCustomerId**](CustomerApi.md#GetCustomerByCustomerId) | **Get** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID} | Get Customer by CUSTOMER_ID
[**GetCustomerByCustomerNumber**](CustomerApi.md#GetCustomerByCustomerNumber) | **Post** /obp/v5.1.0/banks/{BANK_ID}/customers/customer-number | Get Customer by CUSTOMER_NUMBER
[**GetCustomerMessages**](CustomerApi.md#GetCustomerMessages) | **Get** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/messages | Get Customer Messages for a Customer
[**GetCustomerOverview**](CustomerApi.md#GetCustomerOverview) | **Post** /obp/v5.1.0/banks/{BANK_ID}/customers/customer-number-query/overview | Get Customer Overview
[**GetCustomerOverviewFlat**](CustomerApi.md#GetCustomerOverviewFlat) | **Post** /obp/v5.1.0/banks/{BANK_ID}/customers/customer-number-query/overview-flat | Get Customer Overview Flat
[**GetCustomersAtAnyBank**](CustomerApi.md#GetCustomersAtAnyBank) | **Get** /obp/v5.1.0/customers | Get Customers at Any Bank
[**GetCustomersAtOneBank**](CustomerApi.md#GetCustomersAtOneBank) | **Get** /obp/v5.1.0/banks/{BANK_ID}/customers | Get Customers at Bank
[**GetCustomersByCustomerPhoneNumber**](CustomerApi.md#GetCustomersByCustomerPhoneNumber) | **Post** /obp/v5.1.0/banks/{BANK_ID}/search/customers/mobile-phone-number | Get Customers by MOBILE_PHONE_NUMBER
[**GetCustomersForUser**](CustomerApi.md#GetCustomersForUser) | **Get** /obp/v5.1.0/users/current/customers | Get Customers for Current User
[**GetCustomersForUserIdsOnly**](CustomerApi.md#GetCustomersForUserIdsOnly) | **Get** /obp/v5.1.0/users/current/customers/customer_ids | Get Customers for Current User (IDs only)
[**GetCustomersMessages**](CustomerApi.md#GetCustomersMessages) | **Get** /obp/v5.1.0/banks/{BANK_ID}/customer/messages | Get Customer Messages for all Customers
[**GetCustomersMinimalAtAnyBank**](CustomerApi.md#GetCustomersMinimalAtAnyBank) | **Get** /obp/v5.1.0/customers-minimal | Get Customers Minimal at Any Bank
[**GetCustomersMinimalAtOneBank**](CustomerApi.md#GetCustomersMinimalAtOneBank) | **Get** /obp/v5.1.0/banks/{BANK_ID}/customers-minimal | Get Customers Minimal at Bank
[**GetFirehoseCustomers**](CustomerApi.md#GetFirehoseCustomers) | **Get** /obp/v5.1.0/banks/{BANK_ID}/firehose/customers | Get Firehose Customers
[**GetKycChecks**](CustomerApi.md#GetKycChecks) | **Get** /obp/v5.1.0/customers/{CUSTOMER_ID}/kyc_checks | Get Customer KYC Checks
[**GetKycDocuments**](CustomerApi.md#GetKycDocuments) | **Get** /obp/v5.1.0/customers/{CUSTOMER_ID}/kyc_documents | Get Customer KYC Documents
[**GetKycMedia**](CustomerApi.md#GetKycMedia) | **Get** /obp/v5.1.0/customers/{CUSTOMER_ID}/kyc_media | Get KYC Media for a customer
[**GetKycStatuses**](CustomerApi.md#GetKycStatuses) | **Get** /obp/v5.1.0/customers/{CUSTOMER_ID}/kyc_statuses | Get Customer KYC statuses
[**GetMeeting**](CustomerApi.md#GetMeeting) | **Get** /obp/v5.1.0/banks/{BANK_ID}/meetings/{MEETING_ID} | Get Meeting
[**GetMeetings**](CustomerApi.md#GetMeetings) | **Get** /obp/v5.1.0/banks/{BANK_ID}/meetings | Get Meetings
[**GetMyCorrelatedEntities**](CustomerApi.md#GetMyCorrelatedEntities) | **Get** /obp/v5.1.0/my/correlated-entities | Get Correlated Entities for the current User
[**GetMyCustomersAtAnyBank**](CustomerApi.md#GetMyCustomersAtAnyBank) | **Get** /obp/v5.1.0/my/customers | Get My Customers
[**GetMyCustomersAtBank**](CustomerApi.md#GetMyCustomersAtBank) | **Get** /obp/v5.1.0/banks/{BANK_ID}/my/customers | Get My Customers at Bank
[**GetSocialMediaHandles**](CustomerApi.md#GetSocialMediaHandles) | **Get** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/social_media_handles | Get Customer Social Media Handles
[**GetTaxResidence**](CustomerApi.md#GetTaxResidence) | **Get** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/tax-residences | Get Tax Residences of Customer
[**GetUserCustomerLinksByCustomerId**](CustomerApi.md#GetUserCustomerLinksByCustomerId) | **Get** /obp/v5.1.0/banks/{BANK_ID}/user_customer_links/customers/{CUSTOMER_ID} | Get User Customer Links by Customer
[**GetUserCustomerLinksByUserId**](CustomerApi.md#GetUserCustomerLinksByUserId) | **Get** /obp/v5.1.0/banks/{BANK_ID}/user_customer_links/users/{USER_ID} | Get User Customer Links by User
[**UpdateCustomerAccountLinkById**](CustomerApi.md#UpdateCustomerAccountLinkById) | **Put** /obp/v5.1.0/banks/{BANK_ID}/customer-account-links/CUSTOMER_ACCOUNT_LINK_ID | Update Customer Account Link by Id
[**UpdateCustomerAddress**](CustomerApi.md#UpdateCustomerAddress) | **Put** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/addresses/{CUSTOMER_ADDRESS_ID} | Update the Address of a Customer
[**UpdateCustomerAttribute**](CustomerApi.md#UpdateCustomerAttribute) | **Put** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/attributes/CUSTOMER_ATTRIBUTE_ID | Update Customer Attribute
[**UpdateCustomerBranch**](CustomerApi.md#UpdateCustomerBranch) | **Put** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/branch | Update the Branch of a Customer
[**UpdateCustomerCreditLimit**](CustomerApi.md#UpdateCustomerCreditLimit) | **Put** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/credit-limit | Update the credit limit of a Customer
[**UpdateCustomerCreditRatingAndSource**](CustomerApi.md#UpdateCustomerCreditRatingAndSource) | **Put** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/credit-rating-and-source | Update the credit rating and source of a Customer
[**UpdateCustomerData**](CustomerApi.md#UpdateCustomerData) | **Put** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/data | Update the other data of a Customer
[**UpdateCustomerEmail**](CustomerApi.md#UpdateCustomerEmail) | **Put** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/email | Update the email of a Customer
[**UpdateCustomerIdentity**](CustomerApi.md#UpdateCustomerIdentity) | **Put** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/identity | Update the identity data of a Customer
[**UpdateCustomerMobileNumber**](CustomerApi.md#UpdateCustomerMobileNumber) | **Put** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/mobile-number | Update the mobile number of a Customer
[**UpdateCustomerNumber**](CustomerApi.md#UpdateCustomerNumber) | **Put** /obp/v5.1.0/banks/{BANK_ID}/customers/{CUSTOMER_ID}/number | Update the number of a Customer


# **AddCustomerMessage**
> SuccessMessage AddCustomerMessage(ctx, body, cUSTOMERID, bANKID)
Create Customer Message

<p>Create a message for the customer specified by CUSTOMER_ID</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddCustomerMessageJson**](AddCustomerMessageJson.md)| AddCustomerMessageJson object that needs to be added. | 
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

# **AddSocialMediaHandle**
> SuccessMessage AddSocialMediaHandle(ctx, body, cUSTOMERID, bANKID)
Create Customer Social Media Handle

<p>Create a customer social media handle for the customer specified by CUSTOMER_ID</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SocialMediaJson**](SocialMediaJson.md)| SocialMediaJSON object that needs to be added. | 
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCustomer**
> CustomerJsonV310 CreateCustomer(ctx, body, bANKID)
Create Customer

<p>The Customer resource stores the customer number (which is set by the backend), legal name, email, phone number, their date of birth, relationship status, education attained, a url for a profile image, KYC status etc.<br />Dates need to be in the format 2013-01-21T23:08:00Z</p><p>Note: If you need to set a specific customer number, use the Update Customer Number endpoint after this call.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostCustomerJsonV500**](PostCustomerJsonV500.md)| PostCustomerJsonV500 object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerJsonV310**](CustomerJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCustomerAccountLink**
> CustomerAccountLinkJson CreateCustomerAccountLink(ctx, body, bANKID)
Create Customer Account Link

<p>Link a Customer to a Account</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateCustomerAccountLinkJson**](CreateCustomerAccountLinkJson.md)| CreateCustomerAccountLinkJson object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerAccountLinkJson**](CustomerAccountLinkJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCustomerAddress**
> CustomerAddressJsonV310 CreateCustomerAddress(ctx, body, cUSTOMERID, bANKID)
Create Address

<p>Create an Address for a Customer specified by CUSTOMER_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostCustomerAddressJsonV310**](PostCustomerAddressJsonV310.md)| PostCustomerAddressJsonV310 object that needs to be added. | 
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerAddressJsonV310**](CustomerAddressJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCustomerAttribute**
> CustomerAttributeResponseJsonV300 CreateCustomerAttribute(ctx, body, cUSTOMERID, bANKID)
Create Customer Attribute

<p>Create Customer Attribute</p><p>The type field must be one of &quot;STRING&quot;, &quot;INTEGER&quot;, &quot;DOUBLE&quot; or DATE_WITH_DAY&quot;</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CustomerAttributeJsonV400**](CustomerAttributeJsonV400.md)| CustomerAttributeJsonV400 object that needs to be added. | 
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerAttributeResponseJsonV300**](CustomerAttributeResponseJsonV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCustomerMessage**
> SuccessMessage CreateCustomerMessage(ctx, body, cUSTOMERID, bANKID)
Create Customer Message

<p>Create a message for the customer specified by CUSTOMER_ID<br />Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateMessageJsonV400**](CreateMessageJsonV400.md)| CreateMessageJsonV400 object that needs to be added. | 
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMeeting**
> MeetingJsonV310 CreateMeeting(ctx, body, bANKID)
Create Meeting (video conference/call)

<p>Create Meeting: Initiate a video conference/call with the bank.</p><p>The Meetings resource contains meta data about video/other conference sessions</p><p>provider_id determines the provider of the meeting / video chat service. MUST be url friendly (no spaces).</p><p>purpose_id explains the purpose of the chat. onboarding | mortgage | complaint etc. MUST be url friendly (no spaces).</p><p>Login is required.</p><p>This call is <strong>experimental</strong>. Currently staff_user_id is not set. Further calls will be needed to correctly set this.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateMeetingJsonV310**](CreateMeetingJsonV310.md)| CreateMeetingJsonV310 object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**MeetingJsonV310**](MeetingJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateCustomerAttributeAttributeDefinition**
> AttributeDefinitionResponseJsonV400 CreateOrUpdateCustomerAttributeAttributeDefinition(ctx, body, bANKID)
Create or Update Customer Attribute Definition

<p>Create or Update Customer Attribute Definition</p><p>The category field must be one of: Customer</p><p>The type field must be one of; DOUBLE, STRING, INTEGER and DATE_WITH_DAY</p><p>Authentication is Mandatory</p>

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

# **CreateUserCustomerLinks**
> UserCustomerLinkJson CreateUserCustomerLinks(ctx, body, bANKID)
Create User Customer Link

<p>Link a User to a Customer</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateUserCustomerLinkJson**](CreateUserCustomerLinkJson.md)| CreateUserCustomerLinkJson object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**UserCustomerLinkJson**](UserCustomerLinkJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCustomerAccountLinkById**
> DeleteCustomerAccountLinkById(ctx, bANKID)
Delete Customer Account Link

<p>Delete Customer Account Link by CUSTOMER_ACCOUNT_LINK_ID</p><p>Authentication is Mandatory</p>

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

# **DeleteCustomerAttribute**
> DeleteCustomerAttribute(ctx, cUSTOMERID, bANKID)
Delete Customer Attribute

<p>Delete Customer Attribute</p><p>CustomerAttributes are used to enhance the OBP Customer object with Bank specific entities.</p><p>Delete a Customer Attribute by its id.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

# **DeleteCustomerAttributeDefinition**
> DeleteCustomerAttributeDefinition(ctx, bANKID)
Delete Customer Attribute Definition

<p>Delete Customer Attribute Definition by ATTRIBUTE_DEFINITION_ID</p><p>Authentication is Mandatory</p>

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

# **DeleteCustomerCascade**
> DeleteCustomerCascade(ctx, cUSTOMERID, bANKID)
Delete Customer Cascade

<p>Delete a Customer Cascade specified by CUSTOMER_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

# **DeleteUserCustomerLink**
> DeleteUserCustomerLink(ctx, bANKID)
Delete User Customer Link

<p>Delete User Customer Link by USER_CUSTOMER_LINK_ID</p><p>Authentication is Mandatory</p>

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

# **GetCorrelatedUsersInfoByCustomerId**
> CustomerAndUsersWithAttributesResponseJson GetCorrelatedUsersInfoByCustomerId(ctx, cUSTOMERID, bANKID)
Get Correlated User Info by Customer

<p>Get Correlated User Info by CUSTOMER_ID</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerAndUsersWithAttributesResponseJson**](CustomerAndUsersWithAttributesResponseJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCrmEvents**
> CrmEventsJson GetCrmEvents(ctx, body, bANKID)
Get CRM Events

<p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CrmEventsJson**](CrmEventsJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomerAccountLinkById**
> CustomerAccountLinkJson GetCustomerAccountLinkById(ctx, bANKID)
Get Customer Account Link by Id

<p>Get Customer Account Link by CUSTOMER_ACCOUNT_LINK_ID</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerAccountLinkJson**](CustomerAccountLinkJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomerAccountLinksByBankIdAccountId**
> CustomerAccountLinksJson GetCustomerAccountLinksByBankIdAccountId(ctx, aCCOUNTID, bANKID)
Get Customer Account Links by ACCOUNT_ID

<p>Get Customer Account Links by ACCOUNT_ID</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aCCOUNTID** | **string**| The account id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerAccountLinksJson**](CustomerAccountLinksJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomerAccountLinksByCustomerId**
> CustomerAccountLinksJson GetCustomerAccountLinksByCustomerId(ctx, cUSTOMERID, bANKID)
Get Customer Account Links by CUSTOMER_ID

<p>Get Customer Account Links by CUSTOMER_ID</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerAccountLinksJson**](CustomerAccountLinksJson.md)

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

# **GetCustomerAttributeById**
> CustomerAttributeResponseJsonV300 GetCustomerAttributeById(ctx, cUSTOMERID, bANKID)
Get Customer Attribute By Id

<p>Get Customer Attribute By Id</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerAttributeResponseJsonV300**](CustomerAttributeResponseJsonV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomerAttributeDefinition**
> AttributeDefinitionsResponseJsonV400 GetCustomerAttributeDefinition(ctx, bANKID)
Get Customer Attribute Definition

<p>Get Customer Attribute Definition</p><p>Authentication is Mandatory</p>

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

# **GetCustomerAttributes**
> CustomerAttributesResponseJson GetCustomerAttributes(ctx, cUSTOMERID, bANKID)
Get Customer Attributes

<p>Get Customer Attributes</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerAttributesResponseJson**](CustomerAttributesResponseJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomerByCustomerId**
> CustomerWithAttributesJsonV310 GetCustomerByCustomerId(ctx, cUSTOMERID, bANKID)
Get Customer by CUSTOMER_ID

<p>Gets the Customer specified by CUSTOMER_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerWithAttributesJsonV310**](CustomerWithAttributesJsonV310.md)

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

# **GetCustomerMessages**
> CustomerMessagesJsonV400 GetCustomerMessages(ctx, cUSTOMERID, bANKID)
Get Customer Messages for a Customer

<p>Get messages for the customer specified by CUSTOMER_ID<br />Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerMessagesJsonV400**](CustomerMessagesJsonV400.md)

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

# **GetCustomersAtAnyBank**
> CustomerJsonsV300 GetCustomersAtAnyBank(ctx, )
Get Customers at Any Bank

<p>Get Customers at Any Bank.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CustomerJsonsV300**](CustomerJSONsV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomersAtOneBank**
> CustomerJsonsV300 GetCustomersAtOneBank(ctx, bANKID)
Get Customers at Bank

<p>Get Customers at Bank.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerJsonsV300**](CustomerJSONsV300.md)

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

# **GetCustomersForUser**
> CustomersWithAttributesJsonV300 GetCustomersForUser(ctx, )
Get Customers for Current User

<p>Gets all Customers that are linked to a User.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CustomersWithAttributesJsonV300**](CustomersWithAttributesJsonV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomersForUserIdsOnly**
> CustomersWithAttributesJsonV300 GetCustomersForUserIdsOnly(ctx, )
Get Customers for Current User (IDs only)

<p>Gets all Customers Ids that are linked to a User.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CustomersWithAttributesJsonV300**](CustomersWithAttributesJsonV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomersMessages**
> CustomerMessagesJson GetCustomersMessages(ctx, body, bANKID)
Get Customer Messages for all Customers

<p>Get messages for the logged in customer<br />Messages sent to the currently authenticated user.</p><p>Authentication via OAuth is required.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerMessagesJson**](CustomerMessagesJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomersMinimalAtAnyBank**
> CustomersMinimalJsonV400 GetCustomersMinimalAtAnyBank(ctx, )
Get Customers Minimal at Any Bank

<p>Get Customers Minimal at Any Bank.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CustomersMinimalJsonV400**](CustomersMinimalJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomersMinimalAtOneBank**
> CustomersMinimalJsonV400 GetCustomersMinimalAtOneBank(ctx, bANKID)
Get Customers Minimal at Bank

<p>Get Customers Minimal at Bank.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomersMinimalJsonV400**](CustomersMinimalJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFirehoseCustomers**
> CustomerJsons GetFirehoseCustomers(ctx, bANKID)
Get Firehose Customers

<p>Get Customers that has a firehose View.</p><p>Allows bulk access to customers.<br />User must have the CanUseFirehoseAtAnyBank Role</p><p>Possible custom url parameters for pagination:</p><ul><li>limit=NUMBER ==&gt; default value: 500</li><li>offset=NUMBER ==&gt; default value: 0</li></ul><p>eg1:?limit=100&amp;offset=0</p><ul><li>sort_direction=ASC/DESC ==&gt; default value: DESC.</li></ul><p>eg2:?limit=100&amp;offset=0&amp;sort_direction=ASC</p><ul><li>from_date=DATE =&gt; example value: 1970-01-01T00:00:00.000Z. NOTE! The default value is one year ago (1970-01-01T00:00:00.000Z).</li><li>to_date=DATE =&gt; example value: 2024-02-05T14:15:55.283Z. NOTE! The default value is now (2024-02-05T14:15:55.283Z).</li></ul><p>Date format parameter: yyyy-MM-dd'T'HH:mm:ss.SSS'Z'(1100-01-01T01:01:01.000Z) ==&gt; time zone is UTC.</p><p>eg3:?sort_direction=ASC&amp;limit=100&amp;offset=0&amp;from_date=1100-01-01T01:01:01.000Z&amp;to_date=1100-01-01T01:01:01.000Z</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerJsons**](CustomerJSONs.md)

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

# **GetMeeting**
> MeetingJsonV310 GetMeeting(ctx, mEETINGID, bANKID)
Get Meeting

<p>Get Meeting specified by BANK_ID / MEETING_ID<br />Meetings contain meta data about, and are used to facilitate, video conferences / chats etc.</p><p>The actual conference/chats are handled by external services.</p><p>Login is required.</p><p>This call is <strong>experimental</strong> and will require further authorisation in the future.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mEETINGID** | **string**| the meeting id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**MeetingJsonV310**](MeetingJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMeetings**
> MeetingsJsonV310 GetMeetings(ctx, bANKID)
Get Meetings

<p>Meetings contain meta data about, and are used to facilitate, video conferences / chats etc.</p><p>The actual conference/chats are handled by external services.</p><p>Login is required.</p><p>This call is <strong>experimental</strong> and will require further authorisation in the future.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**MeetingsJsonV310**](MeetingsJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMyCorrelatedEntities**
> CorrelatedEntities GetMyCorrelatedEntities(ctx, )
Get Correlated Entities for the current User

<p>Correlated Entities are users and customers linked to the currently authenticated user via User-Customer-Links</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CorrelatedEntities**](CorrelatedEntities.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMyCustomersAtAnyBank**
> CustomerJsonV210 GetMyCustomersAtAnyBank(ctx, )
Get My Customers

<p>Gets all Customers that are linked to me.</p><p>Authentication via OAuth is required.</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CustomerJsonV210**](CustomerJsonV210.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMyCustomersAtBank**
> CustomerJsons GetMyCustomersAtBank(ctx, bANKID)
Get My Customers at Bank

<p>Returns a list of Customers at the Bank that are linked to the currently authenticated User.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerJsons**](CustomerJSONs.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSocialMediaHandles**
> SocialMediasJson GetSocialMediaHandles(ctx, body, cUSTOMERID, bANKID)
Get Customer Social Media Handles

<p>Get social media handles for a customer specified by CUSTOMER_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmptyClassJson**](EmptyClassJson.md)| EmptyClassJson object that needs to be added. | 
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**SocialMediasJson**](SocialMediasJSON.md)

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

# **GetUserCustomerLinksByCustomerId**
> UserCustomerLinksJson GetUserCustomerLinksByCustomerId(ctx, cUSTOMERID, bANKID)
Get User Customer Links by Customer

<p>Get User Customer Links by CUSTOMER_ID</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**UserCustomerLinksJson**](UserCustomerLinksJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserCustomerLinksByUserId**
> UserCustomerLinksJson GetUserCustomerLinksByUserId(ctx, uSERID, bANKID)
Get User Customer Links by User

<p>Get User Customer Links by USER_ID</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uSERID** | **string**| The user id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**UserCustomerLinksJson**](UserCustomerLinksJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomerAccountLinkById**
> CustomerAccountLinkJson UpdateCustomerAccountLinkById(ctx, body, bANKID)
Update Customer Account Link by Id

<p>Update Customer Account Link by CUSTOMER_ACCOUNT_LINK_ID</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateCustomerAccountLinkJson**](UpdateCustomerAccountLinkJson.md)| UpdateCustomerAccountLinkJson object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerAccountLinkJson**](CustomerAccountLinkJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomerAddress**
> CustomerAddressJsonV310 UpdateCustomerAddress(ctx, body, cUSTOMERADDRESSID, cUSTOMERID, bANKID)
Update the Address of a Customer

<p>Update an Address of the Customer specified by CUSTOMER_ADDRESS_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostCustomerAddressJsonV310**](PostCustomerAddressJsonV310.md)| PostCustomerAddressJsonV310 object that needs to be added. | 
  **cUSTOMERADDRESSID** | **string**| the customer address id | 
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerAddressJsonV310**](CustomerAddressJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomerAttribute**
> CustomerAttributeResponseJsonV300 UpdateCustomerAttribute(ctx, body, cUSTOMERID, bANKID)
Update Customer Attribute

<p>Update Customer Attribute</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CustomerAttributeJsonV400**](CustomerAttributeJsonV400.md)| CustomerAttributeJsonV400 object that needs to be added. | 
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerAttributeResponseJsonV300**](CustomerAttributeResponseJsonV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomerBranch**
> CustomerJsonV310 UpdateCustomerBranch(ctx, body, cUSTOMERID, bANKID)
Update the Branch of a Customer

<p>Update the Branch of the Customer specified by CUSTOMER_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PutUpdateCustomerBranchJsonV310**](PutUpdateCustomerBranchJsonV310.md)| PutUpdateCustomerBranchJsonV310 object that needs to be added. | 
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerJsonV310**](CustomerJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomerCreditLimit**
> CustomerJsonV310 UpdateCustomerCreditLimit(ctx, body, cUSTOMERID, bANKID)
Update the credit limit of a Customer

<p>Update the credit limit of the Customer specified by CUSTOMER_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PutUpdateCustomerCreditLimitJsonV310**](PutUpdateCustomerCreditLimitJsonV310.md)| PutUpdateCustomerCreditLimitJsonV310 object that needs to be added. | 
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerJsonV310**](CustomerJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomerCreditRatingAndSource**
> CustomerJsonV310 UpdateCustomerCreditRatingAndSource(ctx, body, cUSTOMERID, bANKID)
Update the credit rating and source of a Customer

<p>Update the credit rating and source of the Customer specified by CUSTOMER_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PutUpdateCustomerCreditRatingAndSourceJsonV310**](PutUpdateCustomerCreditRatingAndSourceJsonV310.md)| PutUpdateCustomerCreditRatingAndSourceJsonV310 object that needs to be added. | 
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerJsonV310**](CustomerJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomerData**
> CustomerJsonV310 UpdateCustomerData(ctx, body, cUSTOMERID, bANKID)
Update the other data of a Customer

<p>Update the other data of the Customer specified by CUSTOMER_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PutUpdateCustomerDataJsonV310**](PutUpdateCustomerDataJsonV310.md)| PutUpdateCustomerDataJsonV310 object that needs to be added. | 
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerJsonV310**](CustomerJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomerEmail**
> CustomerJsonV310 UpdateCustomerEmail(ctx, body, cUSTOMERID, bANKID)
Update the email of a Customer

<p>Update an email of the Customer specified by CUSTOMER_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PutUpdateCustomerEmailJsonV310**](PutUpdateCustomerEmailJsonV310.md)| PutUpdateCustomerEmailJsonV310 object that needs to be added. | 
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerJsonV310**](CustomerJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomerIdentity**
> CustomerJsonV310 UpdateCustomerIdentity(ctx, body, cUSTOMERID, bANKID)
Update the identity data of a Customer

<p>Update the identity data of the Customer specified by CUSTOMER_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PutUpdateCustomerIdentityJsonV310**](PutUpdateCustomerIdentityJsonV310.md)| PutUpdateCustomerIdentityJsonV310 object that needs to be added. | 
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerJsonV310**](CustomerJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomerMobileNumber**
> CustomerJsonV310 UpdateCustomerMobileNumber(ctx, body, cUSTOMERID, bANKID)
Update the mobile number of a Customer

<p>Update the mobile number of the Customer specified by CUSTOMER_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PutUpdateCustomerMobilePhoneNumberJsonV310**](PutUpdateCustomerMobilePhoneNumberJsonV310.md)| PutUpdateCustomerMobilePhoneNumberJsonV310 object that needs to be added. | 
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerJsonV310**](CustomerJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomerNumber**
> CustomerJsonV310 UpdateCustomerNumber(ctx, body, cUSTOMERID, bANKID)
Update the number of a Customer

<p>Update the number of the Customer specified by CUSTOMER_ID.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PutUpdateCustomerNumberJsonV310**](PutUpdateCustomerNumberJsonV310.md)| PutUpdateCustomerNumberJsonV310 object that needs to be added. | 
  **cUSTOMERID** | **string**| The customer id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**CustomerJsonV310**](CustomerJsonV310.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

