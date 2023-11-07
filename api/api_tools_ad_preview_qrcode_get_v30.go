/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
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

// ToolsAdPreviewQrcodeGetV30ApiService ToolsAdPreviewQrcodeGetV30Api service
type ToolsAdPreviewQrcodeGetV30ApiService service

type ApiOpenApiV30ToolsAdPreviewQrcodeGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsAdPreviewQrcodeGetV30ApiService
	advertiserId *int64
	idType       *string
	promotionId  *int64
	materialId   *int64
}

// 广告主ID
func (r *ApiOpenApiV30ToolsAdPreviewQrcodeGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30ToolsAdPreviewQrcodeGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 查询条件，可选值:\&quot;ID_TYPE_PROJECT\&quot;,\&quot;ID_TYPE_PROMOTION\&quot;
func (r *ApiOpenApiV30ToolsAdPreviewQrcodeGetGetRequest) IdType(idType string) *ApiOpenApiV30ToolsAdPreviewQrcodeGetGetRequest {
	r.idType = &idType
	return r
}

// 广告ID
func (r *ApiOpenApiV30ToolsAdPreviewQrcodeGetGetRequest) PromotionId(promotionId int64) *ApiOpenApiV30ToolsAdPreviewQrcodeGetGetRequest {
	r.promotionId = &promotionId
	return r
}

// 素材ID，查询素材预览时使用
func (r *ApiOpenApiV30ToolsAdPreviewQrcodeGetGetRequest) MaterialId(materialId int64) *ApiOpenApiV30ToolsAdPreviewQrcodeGetGetRequest {
	r.materialId = &materialId
	return r
}

func (r *ApiOpenApiV30ToolsAdPreviewQrcodeGetGetRequest) Execute() (*ToolsAdPreviewQrcodeGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30ToolsAdPreviewQrcodeGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30ToolsAdPreviewQrcodeGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ToolsAdPreviewQrcodeGetGetRequest) WithLog(enable bool) *ApiOpenApiV30ToolsAdPreviewQrcodeGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ToolsAdPreviewQrcodeGetGet Method for OpenApiV30ToolsAdPreviewQrcodeGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ToolsAdPreviewQrcodeGetGetRequest
*/
func (a *ToolsAdPreviewQrcodeGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30ToolsAdPreviewQrcodeGetGetRequest {
	return &ApiOpenApiV30ToolsAdPreviewQrcodeGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAdPreviewQrcodeGetV30Response
func (a *ToolsAdPreviewQrcodeGetV30ApiService) getExecute(r *ApiOpenApiV30ToolsAdPreviewQrcodeGetGetRequest) (*ToolsAdPreviewQrcodeGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsAdPreviewQrcodeGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/tools/ad_preview/qrcode_get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.idType == nil {
		return localVarReturnValue, nil, ReportError("idType is required and must be specified")
	}
	if r.promotionId == nil {
		return localVarReturnValue, nil, ReportError("promotionId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "id_type", r.idType)
	parameterAddToHeaderOrQuery(localVarQueryParams, "promotion_id", r.promotionId)
	if r.materialId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "material_id", r.materialId)
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
