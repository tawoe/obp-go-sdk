# \MetricApi

All URIs are relative to *https://apisandbox.openbankproject.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ElasticSearchMetrics**](MetricApi.md#ElasticSearchMetrics) | **Get** /obp/v5.1.0/search/metrics | Search API Metrics via Elasticsearch
[**GetAggregateMetrics**](MetricApi.md#GetAggregateMetrics) | **Get** /obp/v5.1.0/management/aggregate-metrics | Get Aggregate Metrics
[**GetConnectorMetrics**](MetricApi.md#GetConnectorMetrics) | **Get** /obp/v5.1.0/management/connector/metrics | Get Connector Metrics
[**GetMetrics**](MetricApi.md#GetMetrics) | **Get** /obp/v5.1.0/management/metrics | Get Metrics
[**GetMetricsAtBank**](MetricApi.md#GetMetricsAtBank) | **Get** /obp/v5.1.0/management/metrics/banks/{BANK_ID} | Get Metrics at Bank
[**GetMetricsTopConsumers**](MetricApi.md#GetMetricsTopConsumers) | **Get** /obp/v5.1.0/management/metrics/top-consumers | Get Top Consumers
[**GetTopAPIs**](MetricApi.md#GetTopAPIs) | **Get** /obp/v5.1.0/management/metrics/top-apis | Get Top APIs


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

# **GetAggregateMetrics**
> AggregateMetricJson GetAggregateMetrics(ctx, )
Get Aggregate Metrics

<p>Returns aggregate metrics on api usage eg. total count, response time (in ms), etc.</p><p>Should be able to filter on the following fields</p><p>eg: /management/aggregate-metrics?from_date=1100-01-01T01:01:01.000Z&amp;to_date=1100-01-01T01:01:01.000Z&amp;consumer_id=5<br />&amp;user_id=66214b8e-259e-44ad-8868-3eb47be70646&amp;implemented_by_partial_function=getTransactionsForBankAccount<br />&amp;implemented_in_version=v3.0.0&amp;url=/obp/v3.0.0/banks/gh.29.uk/accounts/8ca8a7e4-6d02-48e3-a029-0b2bf89de9f0/owner/transactions<br />&amp;verb=GET&amp;anon=false&amp;app_name=MapperPostman<br />&amp;exclude_app_names=API-EXPLORER,API-Manager,SOFI,null</p><p>1 from_date (defaults to the day before the current date): eg:from_date=1100-01-01T01:01:01.000Z</p><p>2 to_date (defaults to the current date) eg:to_date=1100-01-01T01:01:01.000Z</p><p>3 consumer_id  (if null ignore)</p><p>4 user_id (if null ignore)</p><p>5 anon (if null ignore) only support two value : true (return where user_id is null.) or false (return where user_id is not null.)</p><p>6 url (if null ignore), note: can not contain '&amp;'.</p><p>7 app_name (if null ignore)</p><p>8 implemented_by_partial_function (if null ignore),</p><p>9 implemented_in_version (if null ignore)</p><p>10 verb (if null ignore)</p><p>11 correlation_id (if null ignore)</p><p>12 include_app_names (if null ignore).eg: &amp;include_app_names=API-EXPLORER,API-Manager,SOFI,null</p><p>13 include_url_patterns (if null ignore).you can design you own SQL LIKE pattern. eg: &amp;include_url_patterns=%management/metrics%,%management/aggregate-metrics%</p><p>14 include_implemented_by_partial_functions (if null ignore).eg: &amp;include_implemented_by_partial_functions=getMetrics,getConnectorMetrics,getAggregateMetrics</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AggregateMetricJson**](AggregateMetricJSON.md)

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

# **GetMetricsTopConsumers**
> TopConsumersJson GetMetricsTopConsumers(ctx, )
Get Top Consumers

<p>Get metrics about the top consumers of the API usage e.g. total count, consumer_id and app_name.</p><p>Should be able to filter on the following fields</p><p>e.g.: /management/metrics/top-consumers?from_date=1970-01-01T00:00:00.000Z&amp;to_date=2024-02-05T14:15:55.281Z&amp;consumer_id=5<br />&amp;user_id=66214b8e-259e-44ad-8868-3eb47be70646&amp;implemented_by_partial_function=getTransactionsForBankAccount<br />&amp;implemented_in_version=v3.0.0&amp;url=/obp/v3.0.0/banks/gh.29.uk/accounts/8ca8a7e4-6d02-48e3-a029-0b2bf89de9f0/owner/transactions<br />&amp;verb=GET&amp;anon=false&amp;app_name=MapperPostman<br />&amp;exclude_app_names=API-EXPLORER,API-Manager,SOFI,null<br />&amp;limit=100</p><p>1 from_date (defaults to the one year ago): eg:from_date=1970-01-01T00:00:00.000Z</p><p>2 to_date (defaults to the current date) eg:to_date=2024-02-05T14:15:55.282Z</p><p>3 consumer_id  (if null ignore)</p><p>4 user_id (if null ignore)</p><p>5 anon (if null ignore) only support two value : true (return where user_id is null.) or false (return where user_id is not null.)</p><p>6 url (if null ignore), note: can not contain '&amp;'.</p><p>7 app_name (if null ignore)</p><p>8 implemented_by_partial_function (if null ignore),</p><p>9 implemented_in_version (if null ignore)</p><p>10 verb (if null ignore)</p><p>11 correlation_id (if null ignore)</p><p>12 duration (if null ignore) non digit chars will be silently omitted</p><p>13 exclude_app_names (if null ignore).eg: &amp;exclude_app_names=API-EXPLORER,API-Manager,SOFI,null</p><p>14 exclude_url_patterns (if null ignore).you can design you own SQL NOT LIKE pattern. eg: &amp;exclude_url_patterns=%management/metrics%,%management/aggregate-metrics%</p><p>15 exclude_implemented_by_partial_functions (if null ignore).eg: &amp;exclude_implemented_by_partial_functions=getMetrics,getConnectorMetrics,getAggregateMetrics</p><p>16 limit (for pagination: defaults to 50)  eg:limit=200</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TopConsumersJson**](TopConsumersJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTopAPIs**
> TopApisJson GetTopAPIs(ctx, )
Get Top APIs

<p>Get metrics about the most popular APIs. e.g.: total count, response time (in ms), etc.</p><p>Should be able to filter on the following fields</p><p>eg: /management/metrics/top-apis?from_date=1970-01-01T00:00:00.000Z&amp;to_date=2024-02-05T14:15:55.280Z&amp;consumer_id=5<br />&amp;user_id=66214b8e-259e-44ad-8868-3eb47be70646&amp;implemented_by_partial_function=getTransactionsForBankAccount<br />&amp;implemented_in_version=v3.0.0&amp;url=/obp/v3.0.0/banks/gh.29.uk/accounts/8ca8a7e4-6d02-48e3-a029-0b2bf89de9f0/owner/transactions<br />&amp;verb=GET&amp;anon=false&amp;app_name=MapperPostman<br />&amp;exclude_app_names=API-EXPLORER,API-Manager,SOFI,null</p><p>1 from_date (defaults to the one year ago): eg:from_date=1970-01-01T00:00:00.000Z</p><p>2 to_date (defaults to the current date) eg:to_date=2024-02-05T14:15:55.280Z</p><p>3 consumer_id  (if null ignore)</p><p>4 user_id (if null ignore)</p><p>5 anon (if null ignore) only support two value : true (return where user_id is null.) or false (return where user_id is not null.)</p><p>6 url (if null ignore), note: can not contain '&amp;'.</p><p>7 app_name (if null ignore)</p><p>8 implemented_by_partial_function (if null ignore),</p><p>9 implemented_in_version (if null ignore)</p><p>10 verb (if null ignore)</p><p>11 correlation_id (if null ignore)</p><p>12 duration (if null ignore) non digit chars will be silently omitted</p><p>13 exclude_app_names (if null ignore).eg: &amp;exclude_app_names=API-EXPLORER,API-Manager,SOFI,null</p><p>14 exclude_url_patterns (if null ignore).you can design you own SQL NOT LIKE pattern. eg: &amp;exclude_url_patterns=%management/metrics%,%management/aggregate-metrics%</p><p>15 exclude_implemented_by_partial_functions (if null ignore).eg: &amp;exclude_implemented_by_partial_functions=getMetrics,getConnectorMetrics,getAggregateMetrics</p><p>Authentication is Mandatory</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TopApisJson**](TopApisJson.md)

### Authorization

[directLogin](../README.md#directLogin), [gatewayLogin](../README.md#gatewayLogin)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

