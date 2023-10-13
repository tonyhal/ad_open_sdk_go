/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsBidSuggestV2ArticleCategory
type ToolsBidSuggestV2ArticleCategory string

// List of tools_bid_suggest_v2_article_category
const (
	PETS_ToolsBidSuggestV2ArticleCategory          ToolsBidSuggestV2ArticleCategory = "PETS"
	SPORTS_ToolsBidSuggestV2ArticleCategory        ToolsBidSuggestV2ArticleCategory = "SPORTS"
	EMOTION_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "EMOTION"
	DIGITAL_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "DIGITAL"
	ESTATE_ToolsBidSuggestV2ArticleCategory        ToolsBidSuggestV2ArticleCategory = "ESTATE"
	BUSINESS_ToolsBidSuggestV2ArticleCategory      ToolsBidSuggestV2ArticleCategory = "BUSINESS"
	CULTURE_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "CULTURE"
	PHOTOGRAPHY_ToolsBidSuggestV2ArticleCategory   ToolsBidSuggestV2ArticleCategory = "PHOTOGRAPHY"
	ESSAY_ToolsBidSuggestV2ArticleCategory         ToolsBidSuggestV2ArticleCategory = "ESSAY"
	RUMOR_CRACKER_ToolsBidSuggestV2ArticleCategory ToolsBidSuggestV2ArticleCategory = "RUMOR_CRACKER"
	MILITARY_ToolsBidSuggestV2ArticleCategory      ToolsBidSuggestV2ArticleCategory = "MILITARY"
	ENTERTAINMENT_ToolsBidSuggestV2ArticleCategory ToolsBidSuggestV2ArticleCategory = "ENTERTAINMENT"
	HEALTH_ToolsBidSuggestV2ArticleCategory        ToolsBidSuggestV2ArticleCategory = "HEALTH"
	LOCAL_ToolsBidSuggestV2ArticleCategory         ToolsBidSuggestV2ArticleCategory = "LOCAL"
	FINANCE_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "FINANCE"
	GOVERNMENT_ToolsBidSuggestV2ArticleCategory    ToolsBidSuggestV2ArticleCategory = "GOVERNMENT"
	SCIENCE_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "SCIENCE"
	WEIGHT_LOSING_ToolsBidSuggestV2ArticleCategory ToolsBidSuggestV2ArticleCategory = "WEIGHT_LOSING"
	STORIES_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "STORIES"
	MOVIE_ToolsBidSuggestV2ArticleCategory         ToolsBidSuggestV2ArticleCategory = "MOVIE"
	FREAK_ToolsBidSuggestV2ArticleCategory         ToolsBidSuggestV2ArticleCategory = "FREAK"
	PARENTING_ToolsBidSuggestV2ArticleCategory     ToolsBidSuggestV2ArticleCategory = "PARENTING"
	REGIMEN_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "REGIMEN"
	AMUSEMENT_ToolsBidSuggestV2ArticleCategory     ToolsBidSuggestV2ArticleCategory = "AMUSEMENT"
	COLLECTION_ToolsBidSuggestV2ArticleCategory    ToolsBidSuggestV2ArticleCategory = "COLLECTION"
	CONSTELLATION_ToolsBidSuggestV2ArticleCategory ToolsBidSuggestV2ArticleCategory = "CONSTELLATION"
	INTERNATIONAL_ToolsBidSuggestV2ArticleCategory ToolsBidSuggestV2ArticleCategory = "INTERNATIONAL"
	FASHION_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "FASHION"
	CARS_ToolsBidSuggestV2ArticleCategory          ToolsBidSuggestV2ArticleCategory = "CARS"
	GRADUATES_ToolsBidSuggestV2ArticleCategory     ToolsBidSuggestV2ArticleCategory = "GRADUATES"
	DESIGN_ToolsBidSuggestV2ArticleCategory        ToolsBidSuggestV2ArticleCategory = "DESIGN"
	GOURMET_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "GOURMET"
	ANIMATION_ToolsBidSuggestV2ArticleCategory     ToolsBidSuggestV2ArticleCategory = "ANIMATION"
	HOME_ToolsBidSuggestV2ArticleCategory          ToolsBidSuggestV2ArticleCategory = "HOME"
	COMICS_ToolsBidSuggestV2ArticleCategory        ToolsBidSuggestV2ArticleCategory = "COMICS"
	GAMES_ToolsBidSuggestV2ArticleCategory         ToolsBidSuggestV2ArticleCategory = "GAMES"
	EXPLORE_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "EXPLORE"
	TIPS_ToolsBidSuggestV2ArticleCategory          ToolsBidSuggestV2ArticleCategory = "TIPS"
	TECHNOLOGY_ToolsBidSuggestV2ArticleCategory    ToolsBidSuggestV2ArticleCategory = "TECHNOLOGY"
	TRAVEL_ToolsBidSuggestV2ArticleCategory        ToolsBidSuggestV2ArticleCategory = "TRAVEL"
	WORKPLACE_ToolsBidSuggestV2ArticleCategory     ToolsBidSuggestV2ArticleCategory = "WORKPLACE"
	HISTORY_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "HISTORY"
	EDUCATION_ToolsBidSuggestV2ArticleCategory     ToolsBidSuggestV2ArticleCategory = "EDUCATION"
	LAWS_ToolsBidSuggestV2ArticleCategory          ToolsBidSuggestV2ArticleCategory = "LAWS"
	SOCIETY_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "SOCIETY"
)

// All allowed values of ToolsBidSuggestV2ArticleCategory enum
var AllowedToolsBidSuggestV2ArticleCategoryEnumValues = []ToolsBidSuggestV2ArticleCategory{
	"PETS",
	"SPORTS",
	"EMOTION",
	"DIGITAL",
	"ESTATE",
	"BUSINESS",
	"CULTURE",
	"PHOTOGRAPHY",
	"ESSAY",
	"RUMOR_CRACKER",
	"MILITARY",
	"ENTERTAINMENT",
	"HEALTH",
	"LOCAL",
	"FINANCE",
	"GOVERNMENT",
	"SCIENCE",
	"WEIGHT_LOSING",
	"STORIES",
	"MOVIE",
	"FREAK",
	"PARENTING",
	"REGIMEN",
	"AMUSEMENT",
	"COLLECTION",
	"CONSTELLATION",
	"INTERNATIONAL",
	"FASHION",
	"CARS",
	"GRADUATES",
	"DESIGN",
	"GOURMET",
	"ANIMATION",
	"HOME",
	"COMICS",
	"GAMES",
	"EXPLORE",
	"TIPS",
	"TECHNOLOGY",
	"TRAVEL",
	"WORKPLACE",
	"HISTORY",
	"EDUCATION",
	"LAWS",
	"SOCIETY",
}

// NewToolsBidSuggestV2ArticleCategoryFromValue returns a pointer to a valid ToolsBidSuggestV2ArticleCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2ArticleCategoryFromValue(v string) (*ToolsBidSuggestV2ArticleCategory, error) {
	ev := ToolsBidSuggestV2ArticleCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2ArticleCategory: valid values are %v", v, AllowedToolsBidSuggestV2ArticleCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2ArticleCategory) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2ArticleCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_article_category value
func (v ToolsBidSuggestV2ArticleCategory) Ptr() *ToolsBidSuggestV2ArticleCategory {
	return &v
}
