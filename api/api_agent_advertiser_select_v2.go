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

// AgentAdvertiserSelectV2ApiService AgentAdvertiserSelectV2Api service
type AgentAdvertiserSelectV2ApiService service

type ApiOpenApi2AgentAdvertiserSelectGetRequest struct {
	ctx          context.Context
	ApiService   *AgentAdvertiserSelectV2ApiService
	advertiserId *int64
	companyIds   *[]int64
	count        *int64
	cursor       *int64
	filtering    *AgentAdvertiserSelectV2Filtering
	page         *int64
	pageSize     *int64
}

func (r *ApiOpenApi2AgentAdvertiserSelectGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2AgentAdvertiserSelectGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2AgentAdvertiserSelectGetRequest) CompanyIds(companyIds []int64) *ApiOpenApi2AgentAdvertiserSelectGetRequest {
	r.companyIds = &companyIds
	return r
}

func (r *ApiOpenApi2AgentAdvertiserSelectGetRequest) Count(count int64) *ApiOpenApi2AgentAdvertiserSelectGetRequest {
	r.count = &count
	return r
}

func (r *ApiOpenApi2AgentAdvertiserSelectGetRequest) Cursor(cursor int64) *ApiOpenApi2AgentAdvertiserSelectGetRequest {
	r.cursor = &cursor
	return r
}

func (r *ApiOpenApi2AgentAdvertiserSelectGetRequest) Filtering(filtering AgentAdvertiserSelectV2Filtering) *ApiOpenApi2AgentAdvertiserSelectGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2AgentAdvertiserSelectGetRequest) Page(page int64) *ApiOpenApi2AgentAdvertiserSelectGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2AgentAdvertiserSelectGetRequest) PageSize(pageSize int64) *ApiOpenApi2AgentAdvertiserSelectGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2AgentAdvertiserSelectGetRequest) Execute() (*AgentAdvertiserSelectV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2AgentAdvertiserSelectGetRequest) AccessToken(accessToken string) *ApiOpenApi2AgentAdvertiserSelectGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2AgentAdvertiserSelectGetRequest) WithLog(enable bool) *ApiOpenApi2AgentAdvertiserSelectGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2AgentAdvertiserSelectGet Method for OpenApi2AgentAdvertiserSelectGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2AgentAdvertiserSelectGetRequest
*/
func (a *AgentAdvertiserSelectV2ApiService) Get(ctx context.Context) *ApiOpenApi2AgentAdvertiserSelectGetRequest {
	return &ApiOpenApi2AgentAdvertiserSelectGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AgentAdvertiserSelectV2Response
func (a *AgentAdvertiserSelectV2ApiService) getExecute(r *ApiOpenApi2AgentAdvertiserSelectGetRequest) (*AgentAdvertiserSelectV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *AgentAdvertiserSelectV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/agent/advertiser/select/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.companyIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "company_ids", r.companyIds)
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count)
	}
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor)
	}
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
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
