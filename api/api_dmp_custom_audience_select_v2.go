/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
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

// DmpCustomAudienceSelectV2ApiService DmpCustomAudienceSelectV2Api service
type DmpCustomAudienceSelectV2ApiService service

type ApiOpenApi2DmpCustomAudienceSelectGetRequest struct {
	ctx          context.Context
	ApiService   *DmpCustomAudienceSelectV2ApiService
	advertiserId *int64
	limit        *int64
	offset       *int64
	selectType   *DmpCustomAudienceSelectV2SelectType
}

func (r *ApiOpenApi2DmpCustomAudienceSelectGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2DmpCustomAudienceSelectGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2DmpCustomAudienceSelectGetRequest) Limit(limit int64) *ApiOpenApi2DmpCustomAudienceSelectGetRequest {
	r.limit = &limit
	return r
}

func (r *ApiOpenApi2DmpCustomAudienceSelectGetRequest) Offset(offset int64) *ApiOpenApi2DmpCustomAudienceSelectGetRequest {
	r.offset = &offset
	return r
}

func (r *ApiOpenApi2DmpCustomAudienceSelectGetRequest) SelectType(selectType DmpCustomAudienceSelectV2SelectType) *ApiOpenApi2DmpCustomAudienceSelectGetRequest {
	r.selectType = &selectType
	return r
}

func (r *ApiOpenApi2DmpCustomAudienceSelectGetRequest) Execute() (*DmpCustomAudienceSelectV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2DmpCustomAudienceSelectGetRequest) AccessToken(accessToken string) *ApiOpenApi2DmpCustomAudienceSelectGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2DmpCustomAudienceSelectGetRequest) WithLog(enable bool) *ApiOpenApi2DmpCustomAudienceSelectGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2DmpCustomAudienceSelectGet Method for OpenApi2DmpCustomAudienceSelectGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2DmpCustomAudienceSelectGetRequest
*/
func (a *DmpCustomAudienceSelectV2ApiService) Get(ctx context.Context) *ApiOpenApi2DmpCustomAudienceSelectGetRequest {
	return &ApiOpenApi2DmpCustomAudienceSelectGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DmpCustomAudienceSelectV2Response
func (a *DmpCustomAudienceSelectV2ApiService) getExecute(r *ApiOpenApi2DmpCustomAudienceSelectGetRequest) (*DmpCustomAudienceSelectV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *DmpCustomAudienceSelectV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/dmp/custom_audience/select/"

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
	if r.selectType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "select_type", r.selectType)
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
