/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// ToolsWechatAppletUpdateV30ApiService ToolsWechatAppletUpdateV30Api service
type ToolsWechatAppletUpdateV30ApiService service

type ApiOpenApiV30ToolsWechatAppletUpdatePostRequest struct {
	ctx                               context.Context
	ApiService                        *ToolsWechatAppletUpdateV30ApiService
	toolsWechatAppletUpdateV30Request *ToolsWechatAppletUpdateV30Request
}

func (r *ApiOpenApiV30ToolsWechatAppletUpdatePostRequest) ToolsWechatAppletUpdateV30Request(toolsWechatAppletUpdateV30Request ToolsWechatAppletUpdateV30Request) *ApiOpenApiV30ToolsWechatAppletUpdatePostRequest {
	r.toolsWechatAppletUpdateV30Request = &toolsWechatAppletUpdateV30Request
	return r
}

func (r *ApiOpenApiV30ToolsWechatAppletUpdatePostRequest) Execute() (*ToolsWechatAppletUpdateV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30ToolsWechatAppletUpdatePostRequest) AccessToken(accessToken string) *ApiOpenApiV30ToolsWechatAppletUpdatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ToolsWechatAppletUpdatePostRequest) WithLog(enable bool) *ApiOpenApiV30ToolsWechatAppletUpdatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ToolsWechatAppletUpdatePost Method for OpenApiV30ToolsWechatAppletUpdatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ToolsWechatAppletUpdatePostRequest
*/
func (a *ToolsWechatAppletUpdateV30ApiService) Post(ctx context.Context) *ApiOpenApiV30ToolsWechatAppletUpdatePostRequest {
	return &ApiOpenApiV30ToolsWechatAppletUpdatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsWechatAppletUpdateV30Response
func (a *ToolsWechatAppletUpdateV30ApiService) postExecute(r *ApiOpenApiV30ToolsWechatAppletUpdatePostRequest) (*ToolsWechatAppletUpdateV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsWechatAppletUpdateV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/tools/wechat_applet/update/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.toolsWechatAppletUpdateV30Request
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
