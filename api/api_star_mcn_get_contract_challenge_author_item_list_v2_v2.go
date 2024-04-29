/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
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

// StarMcnGetContractChallengeAuthorItemListV2V2ApiService StarMcnGetContractChallengeAuthorItemListV2V2Api service
type StarMcnGetContractChallengeAuthorItemListV2V2ApiService service

type ApiOpenApi2StarMcnGetContractChallengeAuthorItemListV2GetRequest struct {
	ctx         context.Context
	ApiService  *StarMcnGetContractChallengeAuthorItemListV2V2ApiService
	starId      *int64
	demandId    *int64
	page        *int32
	pageSize    *int32
	developerId *int64
}

func (r *ApiOpenApi2StarMcnGetContractChallengeAuthorItemListV2GetRequest) StarId(starId int64) *ApiOpenApi2StarMcnGetContractChallengeAuthorItemListV2GetRequest {
	r.starId = &starId
	return r
}

func (r *ApiOpenApi2StarMcnGetContractChallengeAuthorItemListV2GetRequest) DemandId(demandId int64) *ApiOpenApi2StarMcnGetContractChallengeAuthorItemListV2GetRequest {
	r.demandId = &demandId
	return r
}

func (r *ApiOpenApi2StarMcnGetContractChallengeAuthorItemListV2GetRequest) Page(page int32) *ApiOpenApi2StarMcnGetContractChallengeAuthorItemListV2GetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2StarMcnGetContractChallengeAuthorItemListV2GetRequest) PageSize(pageSize int32) *ApiOpenApi2StarMcnGetContractChallengeAuthorItemListV2GetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2StarMcnGetContractChallengeAuthorItemListV2GetRequest) DeveloperId(developerId int64) *ApiOpenApi2StarMcnGetContractChallengeAuthorItemListV2GetRequest {
	r.developerId = &developerId
	return r
}

func (r *ApiOpenApi2StarMcnGetContractChallengeAuthorItemListV2GetRequest) Execute() (*StarMcnGetContractChallengeAuthorItemListV2V2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarMcnGetContractChallengeAuthorItemListV2GetRequest) AccessToken(accessToken string) *ApiOpenApi2StarMcnGetContractChallengeAuthorItemListV2GetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarMcnGetContractChallengeAuthorItemListV2GetRequest) WithLog(enable bool) *ApiOpenApi2StarMcnGetContractChallengeAuthorItemListV2GetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarMcnGetContractChallengeAuthorItemListV2Get Method for OpenApi2StarMcnGetContractChallengeAuthorItemListV2Get

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarMcnGetContractChallengeAuthorItemListV2GetRequest
*/
func (a *StarMcnGetContractChallengeAuthorItemListV2V2ApiService) Get(ctx context.Context) *ApiOpenApi2StarMcnGetContractChallengeAuthorItemListV2GetRequest {
	return &ApiOpenApi2StarMcnGetContractChallengeAuthorItemListV2GetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarMcnGetContractChallengeAuthorItemListV2V2Response
func (a *StarMcnGetContractChallengeAuthorItemListV2V2ApiService) getExecute(r *ApiOpenApi2StarMcnGetContractChallengeAuthorItemListV2GetRequest) (*StarMcnGetContractChallengeAuthorItemListV2V2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarMcnGetContractChallengeAuthorItemListV2V2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/mcn/get_contract_challenge_author_item_list_v2/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.starId == nil {
		return localVarReturnValue, nil, ReportError("starId is required and must be specified")
	}
	if r.demandId == nil {
		return localVarReturnValue, nil, ReportError("demandId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "demand_id", r.demandId)
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	if r.developerId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "developer_id", r.developerId)
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