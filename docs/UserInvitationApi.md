# \UserInvitationApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserInvitation**](UserInvitationApi.md#CreateUserInvitation) | **Post** /obp/v5.1.0/banks/{BANK_ID}/user-invitation | Create User Invitation
[**GetUserInvitation**](UserInvitationApi.md#GetUserInvitation) | **Get** /obp/v5.1.0/banks/{BANK_ID}/user-invitations/SECRET_LINK | Get User Invitation
[**GetUserInvitationAnonymous**](UserInvitationApi.md#GetUserInvitationAnonymous) | **Post** /obp/v5.1.0/banks/{BANK_ID}/user-invitations | Get User Invitation Information
[**GetUserInvitations**](UserInvitationApi.md#GetUserInvitations) | **Get** /obp/v5.1.0/banks/{BANK_ID}/user-invitations | Get User Invitations


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

# **GetUserInvitation**
> UserInvitationJsonV400 GetUserInvitation(ctx, bANKID)
Get User Invitation

<p>Get User Invitation</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**UserInvitationJsonV400**](UserInvitationJsonV400.md)

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

# **GetUserInvitations**
> UserInvitationJsonV400 GetUserInvitations(ctx, bANKID)
Get User Invitations

<p>Get User Invitations</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**UserInvitationJsonV400**](UserInvitationJsonV400.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

