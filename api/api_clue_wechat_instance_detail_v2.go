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

// ClueWechatInstanceDetailV2ApiService ClueWechatInstanceDetailV2Api service
type ClueWechatInstanceDetailV2ApiService service

type ApiOpenApi2ClueWechatInstanceDetailGetRequest struct {
	ctx          context.Context
	ApiService   *ClueWechatInstanceDetailV2ApiService
	advertiserId *int64
	instanceId   *int64
}

// 广告主ID
func (r *ApiOpenApi2ClueWechatInstanceDetailGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ClueWechatInstanceDetailGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 微信号码包ID
func (r *ApiOpenApi2ClueWechatInstanceDetailGetRequest) InstanceId(instanceId int64) *ApiOpenApi2ClueWechatInstanceDetailGetRequest {
	r.instanceId = &instanceId
	return r
}

func (r *ApiOpenApi2ClueWechatInstanceDetailGetRequest) Execute() (*ClueWechatInstanceDetailV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ClueWechatInstanceDetailGetRequest) AccessToken(accessToken string) *ApiOpenApi2ClueWechatInstanceDetailGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ClueWechatInstanceDetailGetRequest) WithLog(enable bool) *ApiOpenApi2ClueWechatInstanceDetailGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ClueWechatInstanceDetailGet Method for OpenApi2ClueWechatInstanceDetailGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ClueWechatInstanceDetailGetRequest
*/
func (a *ClueWechatInstanceDetailV2ApiService) Get(ctx context.Context) *ApiOpenApi2ClueWechatInstanceDetailGetRequest {
	return &ApiOpenApi2ClueWechatInstanceDetailGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ClueWechatInstanceDetailV2Response
func (a *ClueWechatInstanceDetailV2ApiService) getExecute(r *ApiOpenApi2ClueWechatInstanceDetailGetRequest) (*ClueWechatInstanceDetailV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ClueWechatInstanceDetailV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/clue/wechat_instance/detail/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.instanceId == nil {
		return localVarReturnValue, nil, ReportError("instanceId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "instance_id", r.instanceId)
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
