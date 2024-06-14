/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
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

// QianchuanAdvertiserTypeGetV10ApiService QianchuanAdvertiserTypeGetV10Api service
type QianchuanAdvertiserTypeGetV10ApiService service

type ApiOpenApiV10QianchuanAdvertiserTypeGetGetRequest struct {
	ctx           context.Context
	ApiService    *QianchuanAdvertiserTypeGetV10ApiService
	advertiserIds *[]int64
}

func (r *ApiOpenApiV10QianchuanAdvertiserTypeGetGetRequest) AdvertiserIds(advertiserIds []int64) *ApiOpenApiV10QianchuanAdvertiserTypeGetGetRequest {
	r.advertiserIds = &advertiserIds
	return r
}

func (r *ApiOpenApiV10QianchuanAdvertiserTypeGetGetRequest) Execute() (*QianchuanAdvertiserTypeGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanAdvertiserTypeGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanAdvertiserTypeGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanAdvertiserTypeGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanAdvertiserTypeGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanAdvertiserTypeGetGet Method for OpenApiV10QianchuanAdvertiserTypeGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanAdvertiserTypeGetGetRequest
*/
func (a *QianchuanAdvertiserTypeGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanAdvertiserTypeGetGetRequest {
	return &ApiOpenApiV10QianchuanAdvertiserTypeGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanAdvertiserTypeGetV10Response
func (a *QianchuanAdvertiserTypeGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanAdvertiserTypeGetGetRequest) (*QianchuanAdvertiserTypeGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanAdvertiserTypeGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/advertiser/type/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserIds == nil {
		return localVarReturnValue, nil, ReportError("advertiserIds is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_ids", r.advertiserIds)
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
