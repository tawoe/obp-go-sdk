/*
 * Open Bank Project API
 *
 * An Open Source API for Banks. (c) TESOBE GmbH. 2011 - 2024. Licensed under the AGPL and commercial licences.
 *
 * API version: v5.1.0
 * Contact: contact@tesobe.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package obp_golang

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type InsurancePremiumGh29UkApiService service

/*
InsurancePremiumGh29UkApiService Create new Insurance Premium
&lt;p&gt;Create new Insurance Premium.&lt;/p&gt;&lt;p&gt;Retrive the premium for the customer.&lt;/p&gt;&lt;p&gt;&lt;strong&gt;Property List:&lt;/strong&gt;&lt;/p&gt;&lt;ul&gt;&lt;li&gt;name: description of &lt;strong&gt;name&lt;/strong&gt; field, can be markdown text.&lt;/li&gt;&lt;li&gt;number: description of &lt;strong&gt;number&lt;/strong&gt; field, can be markdown text.&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;MethodRouting settings example:&lt;/p&gt;&lt;details&gt;&lt;pre&gt;&lt;code&gt;{  &amp;quot;is_bank_id_exact_match&amp;quot;:false,  &amp;quot;method_name&amp;quot;:&amp;quot;dynamicEntityProcess&amp;quot;,  &amp;quot;connector_name&amp;quot;:&amp;quot;rest_vMar2019&amp;quot;,  &amp;quot;bank_id_pattern&amp;quot;:&amp;quot;.*&amp;quot;,  &amp;quot;parameters&amp;quot;:[    {        &amp;quot;key&amp;quot;:&amp;quot;entityName&amp;quot;,        &amp;quot;value&amp;quot;:&amp;quot;InsurancePremium&amp;quot;    }    {        &amp;quot;key&amp;quot;:&amp;quot;url&amp;quot;,        &amp;quot;value&amp;quot;:&amp;quot;http://mydomain.com/xxx&amp;quot;    }  ]}&lt;/code&gt;&lt;/pre&gt;&lt;/details&gt;&lt;p&gt;Authentication is Mandatory&lt;/p&gt;
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body JObject object that needs to be added.

@return InlineResponse20019InsurancePremiumList
*/
func (a *InsurancePremiumGh29UkApiService) DynamicEntityCreateInsurancePremiumGh29Uk(ctx context.Context, body Body24) (InlineResponse20019InsurancePremiumList, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue InlineResponse20019InsurancePremiumList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/banks/gh.29.uk/InsurancePremium"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key

		}
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 201 {
			var v InlineResponse20019InsurancePremiumList
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 400 {
			var v ErrorUserNotLoggedIn
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
InsurancePremiumGh29UkApiService Delete Insurance Premium by id
&lt;p&gt;Delete Insurance Premium by id&lt;/p&gt;&lt;p&gt;MethodRouting settings example:&lt;/p&gt;&lt;details&gt;&lt;pre&gt;&lt;code&gt;{  &amp;quot;is_bank_id_exact_match&amp;quot;:false,  &amp;quot;method_name&amp;quot;:&amp;quot;dynamicEntityProcess&amp;quot;,  &amp;quot;connector_name&amp;quot;:&amp;quot;rest_vMar2019&amp;quot;,  &amp;quot;bank_id_pattern&amp;quot;:&amp;quot;.*&amp;quot;,  &amp;quot;parameters&amp;quot;:[    {        &amp;quot;key&amp;quot;:&amp;quot;entityName&amp;quot;,        &amp;quot;value&amp;quot;:&amp;quot;InsurancePremium&amp;quot;    }    {        &amp;quot;key&amp;quot;:&amp;quot;url&amp;quot;,        &amp;quot;value&amp;quot;:&amp;quot;http://mydomain.com/xxx&amp;quot;    }  ]}&lt;/code&gt;&lt;/pre&gt;&lt;/details&gt;&lt;p&gt;Authentication is Mandatory&lt;/p&gt;
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body JObject object that needs to be added.

@return InlineResponse20019InsurancePremiumList
*/
func (a *InsurancePremiumGh29UkApiService) DynamicEntityDeleteInsurancePremiumGh29Uk(ctx context.Context, body Body26) (InlineResponse20019InsurancePremiumList, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Delete")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue InlineResponse20019InsurancePremiumList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/banks/gh.29.uk/InsurancePremium/INSURANCE_PREMIUM_ID"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key

		}
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 204 {
			var v InlineResponse20019InsurancePremiumList
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 400 {
			var v ErrorUserNotLoggedIn
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
InsurancePremiumGh29UkApiService Get Insurance Premium List
&lt;p&gt;Get Insurance Premium List.&lt;/p&gt;&lt;p&gt;Retrive the premium for the customer.&lt;/p&gt;&lt;p&gt;&lt;strong&gt;Property List:&lt;/strong&gt;&lt;/p&gt;&lt;ul&gt;&lt;li&gt;name: description of &lt;strong&gt;name&lt;/strong&gt; field, can be markdown text.&lt;/li&gt;&lt;li&gt;number: description of &lt;strong&gt;number&lt;/strong&gt; field, can be markdown text.&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;MethodRouting settings example:&lt;/p&gt;&lt;details&gt;&lt;pre&gt;&lt;code&gt;{  &amp;quot;is_bank_id_exact_match&amp;quot;:false,  &amp;quot;method_name&amp;quot;:&amp;quot;dynamicEntityProcess&amp;quot;,  &amp;quot;connector_name&amp;quot;:&amp;quot;rest_vMar2019&amp;quot;,  &amp;quot;bank_id_pattern&amp;quot;:&amp;quot;.*&amp;quot;,  &amp;quot;parameters&amp;quot;:[    {        &amp;quot;key&amp;quot;:&amp;quot;entityName&amp;quot;,        &amp;quot;value&amp;quot;:&amp;quot;InsurancePremium&amp;quot;    }    {        &amp;quot;key&amp;quot;:&amp;quot;url&amp;quot;,        &amp;quot;value&amp;quot;:&amp;quot;http://mydomain.com/xxx&amp;quot;    }  ]}&lt;/code&gt;&lt;/pre&gt;&lt;/details&gt;&lt;p&gt;Authentication is Mandatory&lt;/p&gt;&lt;p&gt;Can do filter on the fields&lt;br /&gt;e.g: /InsurancePremium?name&#x3D;James%20Brown&amp;amp;number&#x3D;123.456&amp;amp;number&#x3D;11.11&lt;br /&gt;Will do filter by this rule: name &#x3D;&#x3D; &amp;quot;James Brown&amp;quot; &amp;amp;&amp;amp; (number&#x3D;&#x3D;123.456 || number&#x3D;11.11)&lt;/p&gt;
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return InlineResponse20019
*/
func (a *InsurancePremiumGh29UkApiService) DynamicEntityGetInsurancePremiumListGh29Uk(ctx context.Context) (InlineResponse20019, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue InlineResponse20019
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/banks/gh.29.uk/InsurancePremium"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key

		}
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse20019
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 400 {
			var v ErrorUserNotLoggedIn
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
InsurancePremiumGh29UkApiService Get Insurance Premium by id
&lt;p&gt;Get Insurance Premium by id.&lt;/p&gt;&lt;p&gt;Retrive the premium for the customer.&lt;/p&gt;&lt;p&gt;&lt;strong&gt;Property List:&lt;/strong&gt;&lt;/p&gt;&lt;ul&gt;&lt;li&gt;name: description of &lt;strong&gt;name&lt;/strong&gt; field, can be markdown text.&lt;/li&gt;&lt;li&gt;number: description of &lt;strong&gt;number&lt;/strong&gt; field, can be markdown text.&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;MethodRouting settings example:&lt;/p&gt;&lt;details&gt;&lt;pre&gt;&lt;code&gt;{  &amp;quot;is_bank_id_exact_match&amp;quot;:false,  &amp;quot;method_name&amp;quot;:&amp;quot;dynamicEntityProcess&amp;quot;,  &amp;quot;connector_name&amp;quot;:&amp;quot;rest_vMar2019&amp;quot;,  &amp;quot;bank_id_pattern&amp;quot;:&amp;quot;.*&amp;quot;,  &amp;quot;parameters&amp;quot;:[    {        &amp;quot;key&amp;quot;:&amp;quot;entityName&amp;quot;,        &amp;quot;value&amp;quot;:&amp;quot;InsurancePremium&amp;quot;    }    {        &amp;quot;key&amp;quot;:&amp;quot;url&amp;quot;,        &amp;quot;value&amp;quot;:&amp;quot;http://mydomain.com/xxx&amp;quot;    }  ]}&lt;/code&gt;&lt;/pre&gt;&lt;/details&gt;&lt;p&gt;Authentication is Mandatory&lt;/p&gt;
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return InlineResponse20019InsurancePremiumList
*/
func (a *InsurancePremiumGh29UkApiService) DynamicEntityGetSingleInsurancePremiumGh29Uk(ctx context.Context) (InlineResponse20019InsurancePremiumList, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue InlineResponse20019InsurancePremiumList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/banks/gh.29.uk/InsurancePremium/INSURANCE_PREMIUM_ID"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key

		}
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse20019InsurancePremiumList
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 400 {
			var v ErrorUserNotLoggedIn
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
InsurancePremiumGh29UkApiService Update Insurance Premium
&lt;p&gt;Update Insurance Premium.&lt;/p&gt;&lt;p&gt;Retrive the premium for the customer.&lt;/p&gt;&lt;p&gt;&lt;strong&gt;Property List:&lt;/strong&gt;&lt;/p&gt;&lt;ul&gt;&lt;li&gt;name: description of &lt;strong&gt;name&lt;/strong&gt; field, can be markdown text.&lt;/li&gt;&lt;li&gt;number: description of &lt;strong&gt;number&lt;/strong&gt; field, can be markdown text.&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;MethodRouting settings example:&lt;/p&gt;&lt;details&gt;&lt;pre&gt;&lt;code&gt;{  &amp;quot;is_bank_id_exact_match&amp;quot;:false,  &amp;quot;method_name&amp;quot;:&amp;quot;dynamicEntityProcess&amp;quot;,  &amp;quot;connector_name&amp;quot;:&amp;quot;rest_vMar2019&amp;quot;,  &amp;quot;bank_id_pattern&amp;quot;:&amp;quot;.*&amp;quot;,  &amp;quot;parameters&amp;quot;:[    {        &amp;quot;key&amp;quot;:&amp;quot;entityName&amp;quot;,        &amp;quot;value&amp;quot;:&amp;quot;InsurancePremium&amp;quot;    }    {        &amp;quot;key&amp;quot;:&amp;quot;url&amp;quot;,        &amp;quot;value&amp;quot;:&amp;quot;http://mydomain.com/xxx&amp;quot;    }  ]}&lt;/code&gt;&lt;/pre&gt;&lt;/details&gt;&lt;p&gt;Authentication is Mandatory&lt;/p&gt;
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body JObject object that needs to be added.

@return InlineResponse20019InsurancePremiumList
*/
func (a *InsurancePremiumGh29UkApiService) DynamicEntityUpdateInsurancePremiumGh29Uk(ctx context.Context, body Body25) (InlineResponse20019InsurancePremiumList, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Put")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue InlineResponse20019InsurancePremiumList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/banks/gh.29.uk/InsurancePremium/INSURANCE_PREMIUM_ID"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key

		}
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse20019InsurancePremiumList
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 400 {
			var v ErrorUserNotLoggedIn
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
