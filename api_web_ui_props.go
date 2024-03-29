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
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type WebUiPropsApiService service

/*
WebUiPropsApiService Create WebUiProps
&lt;p&gt;Create a WebUiProps.&lt;/p&gt;&lt;p&gt;Authentication is Mandatory&lt;/p&gt;&lt;p&gt;Explaination of Fields:&lt;/p&gt;&lt;ul&gt;&lt;li&gt;name is required String value&lt;/li&gt;&lt;li&gt;value is required String value&lt;/li&gt;&lt;/ul&gt;&lt;p&gt;The line break and double quotations should do escape, example:&lt;/p&gt;&lt;pre&gt;&lt;code&gt;{&amp;quot;name&amp;quot;: &amp;quot;webui_some&amp;quot;, &amp;quot;value&amp;quot;: &amp;quot;this valuehave &amp;quot;line break&amp;quot; and double quotations.&amp;quot;}&lt;/code&gt;&lt;/pre&gt;&lt;p&gt;should do escape like this:&lt;/p&gt;&lt;pre&gt;&lt;code&gt;{&amp;quot;name&amp;quot;: &amp;quot;webui_some&amp;quot;, &amp;quot;value&amp;quot;: &amp;quot;this value\\nhave \\&amp;quot;line break\\&amp;quot; and double quotations.&amp;quot;}&lt;/code&gt;&lt;/pre&gt;&lt;p&gt;Insert image examples:&lt;/p&gt;&lt;pre&gt;&lt;code&gt;// set width&#x3D;100 and height&#x3D;50{&amp;quot;name&amp;quot;: &amp;quot;webui_some_pic&amp;quot;, &amp;quot;value&amp;quot;: &amp;quot;here is a picture &amp;lt;img alt&#x3D;&amp;quot;hello&amp;quot; src&#x3D;&amp;quot;http://somedomain.com/images/pic.png&amp;quot; width&#x3D;&amp;quot;100&amp;quot; height&#x3D;&amp;quot;50&amp;quot; /&amp;gt;&amp;quot;}// only set height&#x3D;50{&amp;quot;name&amp;quot;: &amp;quot;webui_some_pic&amp;quot;, &amp;quot;value&amp;quot;: &amp;quot;here is a picture &amp;lt;img alt&#x3D;&amp;quot;hello&amp;quot; src&#x3D;&amp;quot;http://somedomain.com/images/pic.png&amp;quot; width&#x3D;&amp;quot;&amp;quot; height&#x3D;&amp;quot;50&amp;quot; /&amp;gt;&amp;quot;}// only width&#x3D;20%{&amp;quot;name&amp;quot;: &amp;quot;webui_some_pic&amp;quot;, &amp;quot;value&amp;quot;: &amp;quot;here is a picture &amp;lt;img alt&#x3D;&amp;quot;hello&amp;quot; src&#x3D;&amp;quot;http://somedomain.com/images/pic.png&amp;quot; width&#x3D;&amp;quot;20%&amp;quot; height&#x3D;&amp;quot;&amp;quot; /&amp;gt;&amp;quot;}&lt;/code&gt;&lt;/pre&gt;
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body WebUiPropsCommons object that needs to be added.

@return WebUiPropsCommons
*/
func (a *WebUiPropsApiService) CreateWebUiProps(ctx context.Context, body WebUiPropsCommons) (WebUiPropsCommons, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue WebUiPropsCommons
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/obp/v5.1.0/management/webui_props"

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
			var v WebUiPropsCommons
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
WebUiPropsApiService Delete WebUiProps
&lt;p&gt;Delete a WebUiProps specified by WEB_UI_PROPS_ID.&lt;/p&gt;&lt;p&gt;Authentication is Mandatory&lt;/p&gt;
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param wEBUIPROPSID the web ui props id
*/
func (a *WebUiPropsApiService) DeleteWebUiProps(ctx context.Context, wEBUIPROPSID string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/obp/v5.1.0/management/webui_props/{WEB_UI_PROPS_ID}"
	localVarPath = strings.Replace(localVarPath, "{"+"WEB_UI_PROPS_ID"+"}", fmt.Sprintf("%v", wEBUIPROPSID), -1)

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
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 400 {
			var v ErrorUserNotLoggedIn
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarHttpResponse, newErr
		}

		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
WebUiPropsApiService Get WebUiProps
&lt;p&gt;Get the all WebUiProps key values, those props key with &amp;quot;webui_&amp;quot; can be stored in DB, this endpoint get all from DB.&lt;/p&gt;&lt;p&gt;url query parameter:&lt;br /&gt;active: It must be a boolean string. and If active &#x3D; true, it will show&lt;br /&gt;combination of explicit (inserted) + implicit (default)  method_routings.&lt;/p&gt;&lt;p&gt;eg:&lt;br /&gt;&lt;a href&#x3D;\&quot;https://apisandbox.openbankproject.com/obp/v3.1.0/management/webui_props\&quot;&gt;https://apisandbox.openbankproject.com/obp/v3.1.0/management/webui_props&lt;/a&gt;&lt;br /&gt;&lt;a href&#x3D;\&quot;https://apisandbox.openbankproject.com/obp/v3.1.0/management/webui_props?active&#x3D;true\&quot;&gt;https://apisandbox.openbankproject.com/obp/v3.1.0/management/webui_props?active&#x3D;true&lt;/a&gt;&lt;/p&gt;&lt;p&gt;Authentication is Mandatory&lt;/p&gt;
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return InlineResponse20011
*/
func (a *WebUiPropsApiService) GetWebUiProps(ctx context.Context) (InlineResponse20011, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue InlineResponse20011
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/obp/v5.1.0/management/webui_props"

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
			var v InlineResponse20011
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
