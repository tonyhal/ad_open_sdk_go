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

// CgTransferWalletTransferCanTransferBalanceV30ApiService CgTransferWalletTransferCanTransferBalanceV30Api service
type CgTransferWalletTransferCanTransferBalanceV30ApiService service

type ApiOpenApiV30CgTransferWalletTransferCanTransferBalanceGetRequest struct {
	ctx               context.Context
	ApiService        *CgTransferWalletTransferCanTransferBalanceV30ApiService
	accountId         *int64
	accountType       *CgTransferWalletTransferCanTransferBalanceV30AccountType
	bizRequestNo      *string
	mainWalletId      *int64
	subWalletList     *[]int64
	transferDirection *CgTransferWalletTransferCanTransferBalanceV30TransferDirection
}

// 鉴权账户
func (r *ApiOpenApiV30CgTransferWalletTransferCanTransferBalanceGetRequest) AccountId(accountId int64) *ApiOpenApiV30CgTransferWalletTransferCanTransferBalanceGetRequest {
	r.accountId = &accountId
	return r
}

// 鉴权账户类型
func (r *ApiOpenApiV30CgTransferWalletTransferCanTransferBalanceGetRequest) AccountType(accountType CgTransferWalletTransferCanTransferBalanceV30AccountType) *ApiOpenApiV30CgTransferWalletTransferCanTransferBalanceGetRequest {
	r.accountType = &accountType
	return r
}

// 请求id，推荐uuid，方便请求链路对齐
func (r *ApiOpenApiV30CgTransferWalletTransferCanTransferBalanceGetRequest) BizRequestNo(bizRequestNo string) *ApiOpenApiV30CgTransferWalletTransferCanTransferBalanceGetRequest {
	r.bizRequestNo = &bizRequestNo
	return r
}

// 大钱包id
func (r *ApiOpenApiV30CgTransferWalletTransferCanTransferBalanceGetRequest) MainWalletId(mainWalletId int64) *ApiOpenApiV30CgTransferWalletTransferCanTransferBalanceGetRequest {
	r.mainWalletId = &mainWalletId
	return r
}

// 小钱包id列表
func (r *ApiOpenApiV30CgTransferWalletTransferCanTransferBalanceGetRequest) SubWalletList(subWalletList []int64) *ApiOpenApiV30CgTransferWalletTransferCanTransferBalanceGetRequest {
	r.subWalletList = &subWalletList
	return r
}

// 转账方向，以小钱包视角确定
func (r *ApiOpenApiV30CgTransferWalletTransferCanTransferBalanceGetRequest) TransferDirection(transferDirection CgTransferWalletTransferCanTransferBalanceV30TransferDirection) *ApiOpenApiV30CgTransferWalletTransferCanTransferBalanceGetRequest {
	r.transferDirection = &transferDirection
	return r
}

func (r *ApiOpenApiV30CgTransferWalletTransferCanTransferBalanceGetRequest) Execute() (*CgTransferWalletTransferCanTransferBalanceV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30CgTransferWalletTransferCanTransferBalanceGetRequest) AccessToken(accessToken string) *ApiOpenApiV30CgTransferWalletTransferCanTransferBalanceGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30CgTransferWalletTransferCanTransferBalanceGetRequest) WithLog(enable bool) *ApiOpenApiV30CgTransferWalletTransferCanTransferBalanceGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30CgTransferWalletTransferCanTransferBalanceGet Method for OpenApiV30CgTransferWalletTransferCanTransferBalanceGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30CgTransferWalletTransferCanTransferBalanceGetRequest
*/
func (a *CgTransferWalletTransferCanTransferBalanceV30ApiService) Get(ctx context.Context) *ApiOpenApiV30CgTransferWalletTransferCanTransferBalanceGetRequest {
	return &ApiOpenApiV30CgTransferWalletTransferCanTransferBalanceGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CgTransferWalletTransferCanTransferBalanceV30Response
func (a *CgTransferWalletTransferCanTransferBalanceV30ApiService) getExecute(r *ApiOpenApiV30CgTransferWalletTransferCanTransferBalanceGetRequest) (*CgTransferWalletTransferCanTransferBalanceV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *CgTransferWalletTransferCanTransferBalanceV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/cg_transfer/wallet/transfer/can_transfer_balance/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.accountId == nil {
		return localVarReturnValue, nil, ReportError("accountId is required and must be specified")
	}
	if r.accountType == nil {
		return localVarReturnValue, nil, ReportError("accountType is required and must be specified")
	}
	if r.bizRequestNo == nil {
		return localVarReturnValue, nil, ReportError("bizRequestNo is required and must be specified")
	}
	if r.mainWalletId == nil {
		return localVarReturnValue, nil, ReportError("mainWalletId is required and must be specified")
	}
	if r.subWalletList == nil {
		return localVarReturnValue, nil, ReportError("subWalletList is required and must be specified")
	}
	if r.transferDirection == nil {
		return localVarReturnValue, nil, ReportError("transferDirection is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "account_id", r.accountId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "account_type", r.accountType)
	parameterAddToHeaderOrQuery(localVarQueryParams, "biz_request_no", r.bizRequestNo)
	parameterAddToHeaderOrQuery(localVarQueryParams, "main_wallet_id", r.mainWalletId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "sub_wallet_list", r.subWalletList)
	parameterAddToHeaderOrQuery(localVarQueryParams, "transfer_direction", r.transferDirection)
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
