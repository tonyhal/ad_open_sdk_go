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

// ToolsPromotionCardRecommendGetV2ApiService ToolsPromotionCardRecommendGetV2Api service
type ToolsPromotionCardRecommendGetV2ApiService service

type ApiOpenApi2ToolsPromotionCardRecommendGetGetRequest struct {
	ctx                  context.Context
	ApiService           *ToolsPromotionCardRecommendGetV2ApiService
	adId                 *int64
	advancedCreativeType *ToolsPromotionCardRecommendGetV2AdvancedCreativeType
	advertiserId         *int64
	downloadType         *ToolsPromotionCardRecommendGetV2DownloadType
	recommendType        *ToolsPromotionCardRecommendGetV2RecommendType
	titleList            *[]*ToolsPromotionCardRecommendGetV2TitleListInner
}

func (r *ApiOpenApi2ToolsPromotionCardRecommendGetGetRequest) AdId(adId int64) *ApiOpenApi2ToolsPromotionCardRecommendGetGetRequest {
	r.adId = &adId
	return r
}

func (r *ApiOpenApi2ToolsPromotionCardRecommendGetGetRequest) AdvancedCreativeType(advancedCreativeType ToolsPromotionCardRecommendGetV2AdvancedCreativeType) *ApiOpenApi2ToolsPromotionCardRecommendGetGetRequest {
	r.advancedCreativeType = &advancedCreativeType
	return r
}

func (r *ApiOpenApi2ToolsPromotionCardRecommendGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsPromotionCardRecommendGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsPromotionCardRecommendGetGetRequest) DownloadType(downloadType ToolsPromotionCardRecommendGetV2DownloadType) *ApiOpenApi2ToolsPromotionCardRecommendGetGetRequest {
	r.downloadType = &downloadType
	return r
}

func (r *ApiOpenApi2ToolsPromotionCardRecommendGetGetRequest) RecommendType(recommendType ToolsPromotionCardRecommendGetV2RecommendType) *ApiOpenApi2ToolsPromotionCardRecommendGetGetRequest {
	r.recommendType = &recommendType
	return r
}

func (r *ApiOpenApi2ToolsPromotionCardRecommendGetGetRequest) TitleList(titleList []*ToolsPromotionCardRecommendGetV2TitleListInner) *ApiOpenApi2ToolsPromotionCardRecommendGetGetRequest {
	r.titleList = &titleList
	return r
}

func (r *ApiOpenApi2ToolsPromotionCardRecommendGetGetRequest) Execute() (*ToolsPromotionCardRecommendGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsPromotionCardRecommendGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsPromotionCardRecommendGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsPromotionCardRecommendGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsPromotionCardRecommendGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsPromotionCardRecommendGetGet Method for OpenApi2ToolsPromotionCardRecommendGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsPromotionCardRecommendGetGetRequest
*/
func (a *ToolsPromotionCardRecommendGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsPromotionCardRecommendGetGetRequest {
	return &ApiOpenApi2ToolsPromotionCardRecommendGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsPromotionCardRecommendGetV2Response
func (a *ToolsPromotionCardRecommendGetV2ApiService) getExecute(r *ApiOpenApi2ToolsPromotionCardRecommendGetGetRequest) (*ToolsPromotionCardRecommendGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsPromotionCardRecommendGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/promotion_card/recommend/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.adId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ad_id", r.adId)
	}
	if r.advancedCreativeType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advanced_creative_type", r.advancedCreativeType)
	}
	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.downloadType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "download_type", r.downloadType)
	}
	if r.recommendType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "recommend_type", r.recommendType)
	}
	if r.titleList != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "title_list", r.titleList)
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
