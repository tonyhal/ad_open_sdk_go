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

// AicMaterialPushV30ApiService AicMaterialPushV30Api service
type AicMaterialPushV30ApiService service

type ApiOpenApiV30AicMaterialPushPostRequest struct {
	ctx                       context.Context
	ApiService                *AicMaterialPushV30ApiService
	aicMaterialPushV30Request *AicMaterialPushV30Request
}

func (r *ApiOpenApiV30AicMaterialPushPostRequest) AicMaterialPushV30Request(aicMaterialPushV30Request AicMaterialPushV30Request) *ApiOpenApiV30AicMaterialPushPostRequest {
	r.aicMaterialPushV30Request = &aicMaterialPushV30Request
	return r
}

func (r *ApiOpenApiV30AicMaterialPushPostRequest) Execute() (*AicMaterialPushV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30AicMaterialPushPostRequest) AccessToken(accessToken string) *ApiOpenApiV30AicMaterialPushPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30AicMaterialPushPostRequest) WithLog(enable bool) *ApiOpenApiV30AicMaterialPushPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30AicMaterialPushPost Method for OpenApiV30AicMaterialPushPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30AicMaterialPushPostRequest
*/
func (a *AicMaterialPushV30ApiService) Post(ctx context.Context) *ApiOpenApiV30AicMaterialPushPostRequest {
	return &ApiOpenApiV30AicMaterialPushPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AicMaterialPushV30Response
func (a *AicMaterialPushV30ApiService) postExecute(r *ApiOpenApiV30AicMaterialPushPostRequest) (*AicMaterialPushV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *AicMaterialPushV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/aic/material/push/"

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
	localVarPostBody = r.aicMaterialPushV30Request
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
