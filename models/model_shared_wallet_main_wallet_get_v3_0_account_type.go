/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SharedWalletMainWalletGetV30AccountType
type SharedWalletMainWalletGetV30AccountType string

// List of shared_wallet_main_wallet_get_v3.0_account_type
const (
	AD_SharedWalletMainWalletGetV30AccountType        SharedWalletMainWalletGetV30AccountType = "AD"
	AGENT_SharedWalletMainWalletGetV30AccountType     SharedWalletMainWalletGetV30AccountType = "AGENT"
	LOCAL_SharedWalletMainWalletGetV30AccountType     SharedWalletMainWalletGetV30AccountType = "LOCAL"
	QIANCHUAN_SharedWalletMainWalletGetV30AccountType SharedWalletMainWalletGetV30AccountType = "QIANCHUAN"
)

// Ptr returns reference to shared_wallet_main_wallet_get_v3.0_account_type value
func (v SharedWalletMainWalletGetV30AccountType) Ptr() *SharedWalletMainWalletGetV30AccountType {
	return &v
}
