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

// QianchuanAdBidUpdateV10ApiService QianchuanAdBidUpdateV10Api service
type QianchuanAdBidUpdateV10ApiService service

type ApiOpenApiV10QianchuanAdBidUpdatePostRequest struct {
	ctx                            context.Context
	ApiService                     *QianchuanAdBidUpdateV10ApiService
	qianchuanAdBidUpdateV10Request *QianchuanAdBidUpdateV10Request
}

func (r *ApiOpenApiV10QianchuanAdBidUpdatePostRequest) QianchuanAdBidUpdateV10Request(qianchuanAdBidUpdateV10Request QianchuanAdBidUpdateV10Request) *ApiOpenApiV10QianchuanAdBidUpdatePostRequest {
	r.qianchuanAdBidUpdateV10Request = &qianchuanAdBidUpdateV10Request
	return r
}

func (r *ApiOpenApiV10QianchuanAdBidUpdatePostRequest) Execute() (*QianchuanAdBidUpdateV10Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV10QianchuanAdBidUpdatePostRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanAdBidUpdatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanAdBidUpdatePostRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanAdBidUpdatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanAdBidUpdatePost Method for OpenApiV10QianchuanAdBidUpdatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanAdBidUpdatePostRequest
*/
func (a *QianchuanAdBidUpdateV10ApiService) Post(ctx context.Context) *ApiOpenApiV10QianchuanAdBidUpdatePostRequest {
	return &ApiOpenApiV10QianchuanAdBidUpdatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanAdBidUpdateV10Response
func (a *QianchuanAdBidUpdateV10ApiService) postExecute(r *ApiOpenApiV10QianchuanAdBidUpdatePostRequest) (*QianchuanAdBidUpdateV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanAdBidUpdateV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/ad/bid/update/"

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
	localVarPostBody = r.qianchuanAdBidUpdateV10Request
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
