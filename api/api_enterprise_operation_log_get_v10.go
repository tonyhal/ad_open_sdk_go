/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
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

// EnterpriseOperationLogGetV10ApiService EnterpriseOperationLogGetV10Api service
type EnterpriseOperationLogGetV10ApiService service

type ApiOpenApiV10EnterpriseOperationLogGetGetRequest struct {
	ctx          context.Context
	ApiService   *EnterpriseOperationLogGetV10ApiService
	advertiserId *int64
	limit        *int64
	offset       *int64
	openId       *string
	startTime    **string
}

func (r *ApiOpenApiV10EnterpriseOperationLogGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10EnterpriseOperationLogGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10EnterpriseOperationLogGetGetRequest) Limit(limit int64) *ApiOpenApiV10EnterpriseOperationLogGetGetRequest {
	r.limit = &limit
	return r
}

func (r *ApiOpenApiV10EnterpriseOperationLogGetGetRequest) Offset(offset int64) *ApiOpenApiV10EnterpriseOperationLogGetGetRequest {
	r.offset = &offset
	return r
}

func (r *ApiOpenApiV10EnterpriseOperationLogGetGetRequest) OpenId(openId string) *ApiOpenApiV10EnterpriseOperationLogGetGetRequest {
	r.openId = &openId
	return r
}

func (r *ApiOpenApiV10EnterpriseOperationLogGetGetRequest) StartTime(startTime *string) *ApiOpenApiV10EnterpriseOperationLogGetGetRequest {
	r.startTime = &startTime
	return r
}

func (r *ApiOpenApiV10EnterpriseOperationLogGetGetRequest) Execute() (*EnterpriseOperationLogGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10EnterpriseOperationLogGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10EnterpriseOperationLogGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10EnterpriseOperationLogGetGetRequest) WithLog(enable bool) *ApiOpenApiV10EnterpriseOperationLogGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10EnterpriseOperationLogGetGet Method for OpenApiV10EnterpriseOperationLogGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10EnterpriseOperationLogGetGetRequest
*/
func (a *EnterpriseOperationLogGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10EnterpriseOperationLogGetGetRequest {
	return &ApiOpenApiV10EnterpriseOperationLogGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return EnterpriseOperationLogGetV10Response
func (a *EnterpriseOperationLogGetV10ApiService) getExecute(r *ApiOpenApiV10EnterpriseOperationLogGetGetRequest) (*EnterpriseOperationLogGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *EnterpriseOperationLogGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/enterprise/operation/log/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit)
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset)
	}
	if r.openId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "open_id", r.openId)
	}
	if r.startTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
	}
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
