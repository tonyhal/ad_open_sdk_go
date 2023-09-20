/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsOrangeSiteGetV30DataListSiteType
type ToolsOrangeSiteGetV30DataListSiteType string

// List of tools_orange_site_get_v3.0_data_list_site_type
const (
	ADL_AB_TEST_ToolsOrangeSiteGetV30DataListSiteType               ToolsOrangeSiteGetV30DataListSiteType = "ADL_AB_TEST"
	ADL_BIZ_ToolsOrangeSiteGetV30DataListSiteType                   ToolsOrangeSiteGetV30DataListSiteType = "ADL_BIZ"
	ADL_DCAR_ToolsOrangeSiteGetV30DataListSiteType                  ToolsOrangeSiteGetV30DataListSiteType = "ADL_DCAR"
	ADL_EDU_ToolsOrangeSiteGetV30DataListSiteType                   ToolsOrangeSiteGetV30DataListSiteType = "ADL_EDU"
	ADL_FORM_GROUP_ToolsOrangeSiteGetV30DataListSiteType            ToolsOrangeSiteGetV30DataListSiteType = "ADL_FORM_GROUP"
	ADL_GAME_ToolsOrangeSiteGetV30DataListSiteType                  ToolsOrangeSiteGetV30DataListSiteType = "ADL_GAME"
	ADL_HOME_ToolsOrangeSiteGetV30DataListSiteType                  ToolsOrangeSiteGetV30DataListSiteType = "ADL_HOME"
	ADL_LIFE_ToolsOrangeSiteGetV30DataListSiteType                  ToolsOrangeSiteGetV30DataListSiteType = "ADL_LIFE"
	ADL_NOVEL_ToolsOrangeSiteGetV30DataListSiteType                 ToolsOrangeSiteGetV30DataListSiteType = "ADL_NOVEL"
	AD_APP_ToolsOrangeSiteGetV30DataListSiteType                    ToolsOrangeSiteGetV30DataListSiteType = "AD_APP"
	AUTOGEN_CAR_ToolsOrangeSiteGetV30DataListSiteType               ToolsOrangeSiteGetV30DataListSiteType = "AUTOGEN_CAR"
	AUTOGEN_GAME_ToolsOrangeSiteGetV30DataListSiteType              ToolsOrangeSiteGetV30DataListSiteType = "AUTOGEN_GAME"
	AUTOGEN_NOVEL_ToolsOrangeSiteGetV30DataListSiteType             ToolsOrangeSiteGetV30DataListSiteType = "AUTOGEN_NOVEL"
	BUSINESS_FORM_COMMENT_ToolsOrangeSiteGetV30DataListSiteType     ToolsOrangeSiteGetV30DataListSiteType = "BUSINESS_FORM_COMMENT"
	BUSINESS_FORM_INTERACTIVE_ToolsOrangeSiteGetV30DataListSiteType ToolsOrangeSiteGetV30DataListSiteType = "BUSINESS_FORM_INTERACTIVE"
	BUSINESS_FORM_SEARCH_ToolsOrangeSiteGetV30DataListSiteType      ToolsOrangeSiteGetV30DataListSiteType = "BUSINESS_FORM_SEARCH"
	BUSINESS_FORM_V4_ToolsOrangeSiteGetV30DataListSiteType          ToolsOrangeSiteGetV30DataListSiteType = "BUSINESS_FORM_V4"
	CONSULT_PAGE_ToolsOrangeSiteGetV30DataListSiteType              ToolsOrangeSiteGetV30DataListSiteType = "CONSULT_PAGE"
	CREATIVE_COMP_ToolsOrangeSiteGetV30DataListSiteType             ToolsOrangeSiteGetV30DataListSiteType = "CREATIVE_COMP"
	CREATIVE_FORM_ToolsOrangeSiteGetV30DataListSiteType             ToolsOrangeSiteGetV30DataListSiteType = "CREATIVE_FORM"
	CREATIVE_FORM_OLD_ToolsOrangeSiteGetV30DataListSiteType         ToolsOrangeSiteGetV30DataListSiteType = "CREATIVE_FORM_OLD"
	DPA_ToolsOrangeSiteGetV30DataListSiteType                       ToolsOrangeSiteGetV30DataListSiteType = "DPA"
	ENTERPRISE_ToolsOrangeSiteGetV30DataListSiteType                ToolsOrangeSiteGetV30DataListSiteType = "ENTERPRISE"
	ENTERPRISE_CLOUD_CAR_SHOP_ToolsOrangeSiteGetV30DataListSiteType ToolsOrangeSiteGetV30DataListSiteType = "ENTERPRISE_CLOUD_CAR_SHOP"
	ENTERPRISE_CLUE_ToolsOrangeSiteGetV30DataListSiteType           ToolsOrangeSiteGetV30DataListSiteType = "ENTERPRISE_CLUE"
	ENTERPRISE_CONTACT_ToolsOrangeSiteGetV30DataListSiteType        ToolsOrangeSiteGetV30DataListSiteType = "ENTERPRISE_CONTACT"
	ENTERPRISE_GROUP_BUY_ToolsOrangeSiteGetV30DataListSiteType      ToolsOrangeSiteGetV30DataListSiteType = "ENTERPRISE_GROUP_BUY"
	ENTERPRISE_HOMEPAGE_ToolsOrangeSiteGetV30DataListSiteType       ToolsOrangeSiteGetV30DataListSiteType = "ENTERPRISE_HOMEPAGE"
	ENTERPRISE_RESERVE_ToolsOrangeSiteGetV30DataListSiteType        ToolsOrangeSiteGetV30DataListSiteType = "ENTERPRISE_RESERVE"
	ENTERPRISE_TOOLS_ToolsOrangeSiteGetV30DataListSiteType          ToolsOrangeSiteGetV30DataListSiteType = "ENTERPRISE_TOOLS"
	FORM_ToolsOrangeSiteGetV30DataListSiteType                      ToolsOrangeSiteGetV30DataListSiteType = "FORM"
	HEALTH_ToolsOrangeSiteGetV30DataListSiteType                    ToolsOrangeSiteGetV30DataListSiteType = "HEALTH"
	INNER_EXP_ToolsOrangeSiteGetV30DataListSiteType                 ToolsOrangeSiteGetV30DataListSiteType = "INNER_EXP"
	LOCAL_INDUSTRY_ToolsOrangeSiteGetV30DataListSiteType            ToolsOrangeSiteGetV30DataListSiteType = "LOCAL_INDUSTRY"
	MICRO_APP_ToolsOrangeSiteGetV30DataListSiteType                 ToolsOrangeSiteGetV30DataListSiteType = "MICRO_APP"
	MICRO_APP_SHELL_ToolsOrangeSiteGetV30DataListSiteType           ToolsOrangeSiteGetV30DataListSiteType = "MICRO_APP_SHELL"
	MICRO_GAME_ToolsOrangeSiteGetV30DataListSiteType                ToolsOrangeSiteGetV30DataListSiteType = "MICRO_GAME"
	MINIAPP_ToolsOrangeSiteGetV30DataListSiteType                   ToolsOrangeSiteGetV30DataListSiteType = "MINIAPP"
	MULTI_GOODS_ToolsOrangeSiteGetV30DataListSiteType               ToolsOrangeSiteGetV30DataListSiteType = "MULTI_GOODS"
	NATIVE_ToolsOrangeSiteGetV30DataListSiteType                    ToolsOrangeSiteGetV30DataListSiteType = "NATIVE"
	NATIVE_FORM_ToolsOrangeSiteGetV30DataListSiteType               ToolsOrangeSiteGetV30DataListSiteType = "NATIVE_FORM"
	NO_PUBLISH_SITE_ToolsOrangeSiteGetV30DataListSiteType           ToolsOrangeSiteGetV30DataListSiteType = "NO_PUBLISH_SITE"
	OCEAN_ENGINE_APP_ToolsOrangeSiteGetV30DataListSiteType          ToolsOrangeSiteGetV30DataListSiteType = "OCEAN_ENGINE_APP"
	POLL_ToolsOrangeSiteGetV30DataListSiteType                      ToolsOrangeSiteGetV30DataListSiteType = "POLL"
	PRIVACY_POLICY_ToolsOrangeSiteGetV30DataListSiteType            ToolsOrangeSiteGetV30DataListSiteType = "PRIVACY_POLICY"
	PROGRAM_ToolsOrangeSiteGetV30DataListSiteType                   ToolsOrangeSiteGetV30DataListSiteType = "PROGRAM"
	SELF_ToolsOrangeSiteGetV30DataListSiteType                      ToolsOrangeSiteGetV30DataListSiteType = "SELF"
	SELF_CREATIVE_ToolsOrangeSiteGetV30DataListSiteType             ToolsOrangeSiteGetV30DataListSiteType = "SELF_CREATIVE"
	SLIDE_ToolsOrangeSiteGetV30DataListSiteType                     ToolsOrangeSiteGetV30DataListSiteType = "SLIDE"
	STAR_ToolsOrangeSiteGetV30DataListSiteType                      ToolsOrangeSiteGetV30DataListSiteType = "STAR"
	STAR_API_ToolsOrangeSiteGetV30DataListSiteType                  ToolsOrangeSiteGetV30DataListSiteType = "STAR_API"
	STORE_SITE_ToolsOrangeSiteGetV30DataListSiteType                ToolsOrangeSiteGetV30DataListSiteType = "STORE_SITE"
	STREAMING_ToolsOrangeSiteGetV30DataListSiteType                 ToolsOrangeSiteGetV30DataListSiteType = "STREAMING"
	STRUCTURED_LANDING_PAGE_ToolsOrangeSiteGetV30DataListSiteType   ToolsOrangeSiteGetV30DataListSiteType = "STRUCTURED_LANDING_PAGE"
	SUBCHAIN_ToolsOrangeSiteGetV30DataListSiteType                  ToolsOrangeSiteGetV30DataListSiteType = "SUBCHAIN"
	THIRD_EXP_ToolsOrangeSiteGetV30DataListSiteType                 ToolsOrangeSiteGetV30DataListSiteType = "THIRD_EXP"
	THIRD_SITE_ToolsOrangeSiteGetV30DataListSiteType                ToolsOrangeSiteGetV30DataListSiteType = "THIRD_SITE"
	ULTRA_ToolsOrangeSiteGetV30DataListSiteType                     ToolsOrangeSiteGetV30DataListSiteType = "ULTRA"
	UNIVERSAL_ToolsOrangeSiteGetV30DataListSiteType                 ToolsOrangeSiteGetV30DataListSiteType = "UNIVERSAL"
	WECHAT_APPLET_ToolsOrangeSiteGetV30DataListSiteType             ToolsOrangeSiteGetV30DataListSiteType = "WECHAT_APPLET"
	WECHAT_GAME_ToolsOrangeSiteGetV30DataListSiteType               ToolsOrangeSiteGetV30DataListSiteType = "WECHAT_GAME"
)

// All allowed values of ToolsOrangeSiteGetV30DataListSiteType enum
var AllowedToolsOrangeSiteGetV30DataListSiteTypeEnumValues = []ToolsOrangeSiteGetV30DataListSiteType{
	"ADL_AB_TEST",
	"ADL_BIZ",
	"ADL_DCAR",
	"ADL_EDU",
	"ADL_FORM_GROUP",
	"ADL_GAME",
	"ADL_HOME",
	"ADL_LIFE",
	"ADL_NOVEL",
	"AD_APP",
	"AUTOGEN_CAR",
	"AUTOGEN_GAME",
	"AUTOGEN_NOVEL",
	"BUSINESS_FORM_COMMENT",
	"BUSINESS_FORM_INTERACTIVE",
	"BUSINESS_FORM_SEARCH",
	"BUSINESS_FORM_V4",
	"CONSULT_PAGE",
	"CREATIVE_COMP",
	"CREATIVE_FORM",
	"CREATIVE_FORM_OLD",
	"DPA",
	"ENTERPRISE",
	"ENTERPRISE_CLOUD_CAR_SHOP",
	"ENTERPRISE_CLUE",
	"ENTERPRISE_CONTACT",
	"ENTERPRISE_GROUP_BUY",
	"ENTERPRISE_HOMEPAGE",
	"ENTERPRISE_RESERVE",
	"ENTERPRISE_TOOLS",
	"FORM",
	"HEALTH",
	"INNER_EXP",
	"LOCAL_INDUSTRY",
	"MICRO_APP",
	"MICRO_APP_SHELL",
	"MICRO_GAME",
	"MINIAPP",
	"MULTI_GOODS",
	"NATIVE",
	"NATIVE_FORM",
	"NO_PUBLISH_SITE",
	"OCEAN_ENGINE_APP",
	"POLL",
	"PRIVACY_POLICY",
	"PROGRAM",
	"SELF",
	"SELF_CREATIVE",
	"SLIDE",
	"STAR",
	"STAR_API",
	"STORE_SITE",
	"STREAMING",
	"STRUCTURED_LANDING_PAGE",
	"SUBCHAIN",
	"THIRD_EXP",
	"THIRD_SITE",
	"ULTRA",
	"UNIVERSAL",
	"WECHAT_APPLET",
	"WECHAT_GAME",
}

// NewToolsOrangeSiteGetV30DataListSiteTypeFromValue returns a pointer to a valid ToolsOrangeSiteGetV30DataListSiteType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsOrangeSiteGetV30DataListSiteTypeFromValue(v string) (*ToolsOrangeSiteGetV30DataListSiteType, error) {
	ev := ToolsOrangeSiteGetV30DataListSiteType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsOrangeSiteGetV30DataListSiteType: valid values are %v", v, AllowedToolsOrangeSiteGetV30DataListSiteTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsOrangeSiteGetV30DataListSiteType) IsValid() bool {
	for _, existing := range AllowedToolsOrangeSiteGetV30DataListSiteTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_orange_site_get_v3.0_data_list_site_type value
func (v ToolsOrangeSiteGetV30DataListSiteType) Ptr() *ToolsOrangeSiteGetV30DataListSiteType {
	return &v
}
