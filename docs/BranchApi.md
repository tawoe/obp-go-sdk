# \BranchApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBranch**](BranchApi.md#CreateBranch) | **Post** /obp/v5.1.0/banks/{BANK_ID}/branches | Create Branch
[**DeleteBranch**](BranchApi.md#DeleteBranch) | **Delete** /obp/v5.1.0/banks/{BANK_ID}/branches/{BRANCH_ID} | Delete Branch
[**GetBranch**](BranchApi.md#GetBranch) | **Get** /obp/v5.1.0/banks/{BANK_ID}/branches/{BRANCH_ID} | Get Branch
[**GetBranches**](BranchApi.md#GetBranches) | **Get** /obp/v5.1.0/banks/{BANK_ID}/branches | Get Branches for a Bank


# **CreateBranch**
> BranchJsonV300 CreateBranch(ctx, body, bANKID)
Create Branch

<p>Create Branch for the Bank.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BranchJsonV300**](BranchJsonV300.md)| BranchJsonV300 object that needs to be added. | 
  **bANKID** | **string**| The bank id | 

### Return type

[**BranchJsonV300**](BranchJsonV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBranch**
> DeleteBranch(ctx, bRANCHID, bANKID)
Delete Branch

<p>Delete Branch from given Bank.</p><p>Authentication is Mandatory</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bRANCHID** | **string**| The branch id | 
  **bANKID** | **string**| The bank id | 

### Return type

 (empty response body)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBranch**
> BranchJsonV300 GetBranch(ctx, bRANCHID, bANKID)
Get Branch

<p>Returns information about a single Branch specified by BANK_ID and BRANCH_ID including:</p><ul><li>Name</li><li>Address</li><li>Geo Location</li><li>License the data under this endpoint is released under.</li></ul><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bRANCHID** | **string**| The branch id | 
  **bANKID** | **string**| The bank id | 

### Return type

[**BranchJsonV300**](BranchJsonV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBranches**
> BranchesJsonV300 GetBranches(ctx, bANKID)
Get Branches for a Bank

<p>Returns information about branches for a single bank specified by BANK_ID including:</p><ul><li>Name</li><li>Address</li><li>Geo Location</li><li>License the data under this endpoint is released under</li><li>Structured opening hours</li><li>Accessible flag</li><li>Branch Type</li><li>More Info</li></ul><p>Pagination:</p><p>By default, 50 records are returned.</p><p>You can use the url query parameters <em>limit</em> and <em>offset</em> for pagination<br />You can also use the follow url query parameters:</p><ul><li><p>city - string, find Branches those in this city, optional</p></li><li><p>withinMetersOf - number, find Branches within given meters distance, optional</p></li><li>nearLatitude - number, a position of latitude value, cooperate with withMetersOf do query filter, optional</li><li>nearLongitude - number, a position of longitude value, cooperate with withMetersOf do query filter, optional</li></ul><p>note: withinMetersOf, nearLatitude and nearLongitude either all empty or all have value.</p><p>Authentication is Optional</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bANKID** | **string**| The bank id | 

### Return type

[**BranchesJsonV300**](BranchesJsonV300.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

