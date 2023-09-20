/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
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

// QianchuanFileVideoAwemeGetV10ApiService QianchuanFileVideoAwemeGetV10Api service
type QianchuanFileVideoAwemeGetV10ApiService service

type ApiOpenApiV10QianchuanFileVideoAwemeGetGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanFileVideoAwemeGetV10ApiService
	advertiserId *int64
	awemeId      *int64
	filtering    *QianchuanFileVideoAwemeGetV10Filtering
	cursor       *int64
	count        *int64
}

func (r *ApiOpenApiV10QianchuanFileVideoAwemeGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanFileVideoAwemeGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanFileVideoAwemeGetGetRequest) AwemeId(awemeId int64) *ApiOpenApiV10QianchuanFileVideoAwemeGetGetRequest {
	r.awemeId = &awemeId
	return r
}

func (r *ApiOpenApiV10QianchuanFileVideoAwemeGetGetRequest) Filtering(filtering QianchuanFileVideoAwemeGetV10Filtering) *ApiOpenApiV10QianchuanFileVideoAwemeGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApiV10QianchuanFileVideoAwemeGetGetRequest) Cursor(cursor int64) *ApiOpenApiV10QianchuanFileVideoAwemeGetGetRequest {
	r.cursor = &cursor
	return r
}

func (r *ApiOpenApiV10QianchuanFileVideoAwemeGetGetRequest) Count(count int64) *ApiOpenApiV10QianchuanFileVideoAwemeGetGetRequest {
	r.count = &count
	return r
}

func (r *ApiOpenApiV10QianchuanFileVideoAwemeGetGetRequest) Execute() (*QianchuanFileVideoAwemeGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanFileVideoAwemeGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanFileVideoAwemeGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanFileVideoAwemeGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanFileVideoAwemeGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanFileVideoAwemeGetGet Method for OpenApiV10QianchuanFileVideoAwemeGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanFileVideoAwemeGetGetRequest
*/
func (a *QianchuanFileVideoAwemeGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanFileVideoAwemeGetGetRequest {
	return &ApiOpenApiV10QianchuanFileVideoAwemeGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanFileVideoAwemeGetV10Response
func (a *QianchuanFileVideoAwemeGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanFileVideoAwemeGetGetRequest) (*QianchuanFileVideoAwemeGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *QianchuanFileVideoAwemeGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/file/video/aweme/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.awemeId == nil {
		return localVarReturnValue, nil, ReportError("awemeId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "aweme_id", r.awemeId)
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	}
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor)
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count)
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