package voxgigartinstituteofchicagosdk

import (
	"github.com/voxgig-sdk/art-institute-of-chicago-sdk/go/core"
	"github.com/voxgig-sdk/art-institute-of-chicago-sdk/go/entity"
	"github.com/voxgig-sdk/art-institute-of-chicago-sdk/go/feature"
	_ "github.com/voxgig-sdk/art-institute-of-chicago-sdk/go/utility"
)

// Type aliases preserve external API.
type ArtInstituteOfChicagoSDK = core.ArtInstituteOfChicagoSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type ArtInstituteOfChicagoEntity = core.ArtInstituteOfChicagoEntity
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type ArtInstituteOfChicagoError = core.ArtInstituteOfChicagoError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewAgentEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewAgentEntity(client, entopts)
	}
	core.NewAgentRoleEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewAgentRoleEntity(client, entopts)
	}
	core.NewAgentTypeEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewAgentTypeEntity(client, entopts)
	}
	core.NewArticleEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewArticleEntity(client, entopts)
	}
	core.NewArtworkEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewArtworkEntity(client, entopts)
	}
	core.NewArtworkDateQualifierEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewArtworkDateQualifierEntity(client, entopts)
	}
	core.NewArtworkPlaceQualifierEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewArtworkPlaceQualifierEntity(client, entopts)
	}
	core.NewArtworkTypeEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewArtworkTypeEntity(client, entopts)
	}
	core.NewCategoryTermEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewCategoryTermEntity(client, entopts)
	}
	core.NewDigitalPublicationEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewDigitalPublicationEntity(client, entopts)
	}
	core.NewDigitalPublicationArticleEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewDigitalPublicationArticleEntity(client, entopts)
	}
	core.NewEducatorResourceEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewEducatorResourceEntity(client, entopts)
	}
	core.NewEventEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewEventEntity(client, entopts)
	}
	core.NewEventOccurrenceEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewEventOccurrenceEntity(client, entopts)
	}
	core.NewEventProgramEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewEventProgramEntity(client, entopts)
	}
	core.NewExhibitionEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewExhibitionEntity(client, entopts)
	}
	core.NewGalleryEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewGalleryEntity(client, entopts)
	}
	core.NewGenericPageEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewGenericPageEntity(client, entopts)
	}
	core.NewHighlightEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewHighlightEntity(client, entopts)
	}
	core.NewHourEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewHourEntity(client, entopts)
	}
	core.NewImageEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewImageEntity(client, entopts)
	}
	core.NewLandingPageEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewLandingPageEntity(client, entopts)
	}
	core.NewPlaceEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewPlaceEntity(client, entopts)
	}
	core.NewPressReleaseEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewPressReleaseEntity(client, entopts)
	}
	core.NewPrintedPublicationEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewPrintedPublicationEntity(client, entopts)
	}
	core.NewProductEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewProductEntity(client, entopts)
	}
	core.NewPublicationEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewPublicationEntity(client, entopts)
	}
	core.NewSearchEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewSearchEntity(client, entopts)
	}
	core.NewSectionEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewSectionEntity(client, entopts)
	}
	core.NewSiteEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewSiteEntity(client, entopts)
	}
	core.NewSoundEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewSoundEntity(client, entopts)
	}
	core.NewStaticPageEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewStaticPageEntity(client, entopts)
	}
	core.NewTextEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewTextEntity(client, entopts)
	}
	core.NewTourEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewTourEntity(client, entopts)
	}
	core.NewVideoEntityFunc = func(client *core.ArtInstituteOfChicagoSDK, entopts map[string]any) core.ArtInstituteOfChicagoEntity {
		return entity.NewVideoEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewArtInstituteOfChicagoSDK = core.NewArtInstituteOfChicagoSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig
var SharedConfig = core.SharedConfig

// No-arg convenience constructors. Go has no default-argument syntax,
// so these aliases let callers write `sdk.New()` / `sdk.Test()`
// instead of `sdk.NewArtInstituteOfChicagoSDK(nil)` / `sdk.TestSDK(nil, nil)`
// for the common no-options case.
func New() *ArtInstituteOfChicagoSDK  { return NewArtInstituteOfChicagoSDK(nil) }
func Test() *ArtInstituteOfChicagoSDK { return TestSDK(nil, nil) }
var NewBaseFeature = feature.NewBaseFeature
var NewTestFeature = feature.NewTestFeature
