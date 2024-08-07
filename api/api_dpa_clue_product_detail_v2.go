/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
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

// DpaClueProductDetailV2ApiService DpaClueProductDetailV2Api service
type DpaClueProductDetailV2ApiService service

type ApiOpenApi2DpaClueProductDetailGetRequest struct {
	ctx          context.Context
	ApiService   *DpaClueProductDetailV2ApiService
	advertiserId *int64
	productIds   *[]int64
}

func (r *ApiOpenApi2DpaClueProductDetailGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2DpaClueProductDetailGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2DpaClueProductDetailGetRequest) ProductIds(productIds []int64) *ApiOpenApi2DpaClueProductDetailGetRequest {
	r.productIds = &productIds
	return r
}

func (r *ApiOpenApi2DpaClueProductDetailGetRequest) Execute() (*DpaClueProductDetailV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2DpaClueProductDetailGetRequest) AccessToken(accessToken string) *ApiOpenApi2DpaClueProductDetailGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2DpaClueProductDetailGetRequest) WithLog(enable bool) *ApiOpenApi2DpaClueProductDetailGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2DpaClueProductDetailGet Method for OpenApi2DpaClueProductDetailGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2DpaClueProductDetailGetRequest
*/
func (a *DpaClueProductDetailV2ApiService) Get(ctx context.Context) *ApiOpenApi2DpaClueProductDetailGetRequest {
	return &ApiOpenApi2DpaClueProductDetailGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DpaClueProductDetailV2Response
func (a *DpaClueProductDetailV2ApiService) getExecute(r *ApiOpenApi2DpaClueProductDetailGetRequest) (*DpaClueProductDetailV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *DpaClueProductDetailV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/dpa/clue_product/detail/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.productIds == nil {
		return localVarReturnValue, nil, ReportError("productIds is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "product_ids", r.productIds)
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
