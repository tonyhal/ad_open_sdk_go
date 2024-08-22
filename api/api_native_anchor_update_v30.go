/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"context"
	"net/http"
	"net/url"

	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// NativeAnchorUpdateV30ApiService NativeAnchorUpdateV30Api service
type NativeAnchorUpdateV30ApiService service

type ApiOpenApiV30NativeAnchorUpdatePostRequest struct {
	ctx                          context.Context
	ApiService                   *NativeAnchorUpdateV30ApiService
	nativeAnchorUpdateV30Request *NativeAnchorUpdateV30Request
}

func (r *ApiOpenApiV30NativeAnchorUpdatePostRequest) NativeAnchorUpdateV30Request(nativeAnchorUpdateV30Request NativeAnchorUpdateV30Request) *ApiOpenApiV30NativeAnchorUpdatePostRequest {
	r.nativeAnchorUpdateV30Request = &nativeAnchorUpdateV30Request
	return r
}

func (r *ApiOpenApiV30NativeAnchorUpdatePostRequest) Execute() (*NativeAnchorUpdateV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30NativeAnchorUpdatePostRequest) AccessToken(accessToken string) *ApiOpenApiV30NativeAnchorUpdatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30NativeAnchorUpdatePostRequest) WithLog(enable bool) *ApiOpenApiV30NativeAnchorUpdatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30NativeAnchorUpdatePost Method for OpenApiV30NativeAnchorUpdatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30NativeAnchorUpdatePostRequest
*/
func (a *NativeAnchorUpdateV30ApiService) Post(ctx context.Context) *ApiOpenApiV30NativeAnchorUpdatePostRequest {
	return &ApiOpenApiV30NativeAnchorUpdatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return NativeAnchorUpdateV30Response
func (a *NativeAnchorUpdateV30ApiService) postExecute(r *ApiOpenApiV30NativeAnchorUpdatePostRequest) (*NativeAnchorUpdateV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *NativeAnchorUpdateV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/native_anchor/update/"

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

	// body params
	localVarPostBody = r.nativeAnchorUpdateV30Request
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.call(r.ctx, req, &localVarReturnValue)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}
