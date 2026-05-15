package core

import (
	"fmt"

	vs "github.com/voxgig/struct"
)

type ArtInstituteOfChicagoSDK struct {
	Mode     string
	options  map[string]any
	utility  *Utility
	Features []Feature
	rootctx  *Context
}

func NewArtInstituteOfChicagoSDK(options map[string]any) *ArtInstituteOfChicagoSDK {
	sdk := &ArtInstituteOfChicagoSDK{
		Mode:     "live",
		Features: []Feature{},
	}

	sdk.utility = NewUtility()

	config := MakeConfig()

	sdk.rootctx = sdk.utility.MakeContext(map[string]any{
		"client":  sdk,
		"utility": sdk.utility,
		"config":  config,
		"options": options,
		"shared":  map[string]any{},
	}, nil)

	sdk.options = sdk.utility.MakeOptions(sdk.rootctx)

	if vs.GetPath([]any{"feature", "test", "active"}, sdk.options) == true {
		sdk.Mode = "test"
	}

	sdk.rootctx.Options = sdk.options

	// Add features from config.
	featureOpts := ToMapAny(vs.GetProp(sdk.options, "feature"))
	if featureOpts != nil {
		for _, item := range vs.Items(featureOpts) {
			fname, _ := item[0].(string)
			fopts := ToMapAny(item[1])
			if fopts != nil {
				if active, ok := fopts["active"]; ok {
					if ab, ok := active.(bool); ok && ab {
						sdk.utility.FeatureAdd(sdk.rootctx, makeFeature(fname))
					}
				}
			}
		}
	}

	// Add extension features.
	if extend := vs.GetProp(sdk.options, "extend"); extend != nil {
		if extList, ok := extend.([]any); ok {
			for _, f := range extList {
				if feat, ok := f.(Feature); ok {
					sdk.utility.FeatureAdd(sdk.rootctx, feat)
				}
			}
		}
	}

	// Initialize features.
	for _, f := range sdk.Features {
		sdk.utility.FeatureInit(sdk.rootctx, f)
	}

	sdk.utility.FeatureHook(sdk.rootctx, "PostConstruct")

	return sdk
}

func (sdk *ArtInstituteOfChicagoSDK) OptionsMap() map[string]any {
	out := vs.Clone(sdk.options)
	if om, ok := out.(map[string]any); ok {
		return om
	}
	return map[string]any{}
}

func (sdk *ArtInstituteOfChicagoSDK) GetUtility() *Utility {
	return CopyUtility(sdk.utility)
}

func (sdk *ArtInstituteOfChicagoSDK) GetRootCtx() *Context {
	return sdk.rootctx
}

func (sdk *ArtInstituteOfChicagoSDK) Prepare(fetchargs map[string]any) (map[string]any, error) {
	utility := sdk.utility

	if fetchargs == nil {
		fetchargs = map[string]any{}
	}

	var ctrl map[string]any
	if c := vs.GetProp(fetchargs, "ctrl"); c != nil {
		if cm, ok := c.(map[string]any); ok {
			ctrl = cm
		}
	}
	if ctrl == nil {
		ctrl = map[string]any{}
	}

	ctx := utility.MakeContext(map[string]any{
		"opname": "prepare",
		"ctrl":   ctrl,
	}, sdk.rootctx)

	options := sdk.options

	path, _ := vs.GetProp(fetchargs, "path").(string)
	method, _ := vs.GetProp(fetchargs, "method").(string)
	if method == "" {
		method = "GET"
	}

	params := ToMapAny(vs.GetProp(fetchargs, "params"))
	if params == nil {
		params = map[string]any{}
	}
	query := ToMapAny(vs.GetProp(fetchargs, "query"))
	if query == nil {
		query = map[string]any{}
	}

	headers := utility.PrepareHeaders(ctx)

	base, _ := vs.GetProp(options, "base").(string)
	prefix, _ := vs.GetProp(options, "prefix").(string)
	suffix, _ := vs.GetProp(options, "suffix").(string)

	ctx.Spec = NewSpec(map[string]any{
		"base":    base,
		"prefix":  prefix,
		"suffix":  suffix,
		"path":    path,
		"method":  method,
		"params":  params,
		"query":   query,
		"headers": headers,
		"body":    vs.GetProp(fetchargs, "body"),
		"step":    "start",
	})

	// Merge user-provided headers.
	if uh := vs.GetProp(fetchargs, "headers"); uh != nil {
		if uhm, ok := uh.(map[string]any); ok {
			for k, v := range uhm {
				ctx.Spec.Headers[k] = v
			}
		}
	}

	_, err := utility.PrepareAuth(ctx)
	if err != nil {
		return nil, err
	}

	return utility.MakeFetchDef(ctx)
}

func (sdk *ArtInstituteOfChicagoSDK) Direct(fetchargs map[string]any) (map[string]any, error) {
	utility := sdk.utility

	fetchdef, err := sdk.Prepare(fetchargs)
	if err != nil {
		return map[string]any{"ok": false, "err": err}, nil
	}

	if fetchargs == nil {
		fetchargs = map[string]any{}
	}

	var ctrl map[string]any
	if c := vs.GetProp(fetchargs, "ctrl"); c != nil {
		if cm, ok := c.(map[string]any); ok {
			ctrl = cm
		}
	}
	if ctrl == nil {
		ctrl = map[string]any{}
	}

	ctx := utility.MakeContext(map[string]any{
		"opname": "direct",
		"ctrl":   ctrl,
	}, sdk.rootctx)

	url, _ := fetchdef["url"].(string)
	fetched, fetchErr := utility.Fetcher(ctx, url, fetchdef)

	if fetchErr != nil {
		return map[string]any{"ok": false, "err": fetchErr}, nil
	}

	if fetched == nil {
		return map[string]any{
			"ok":  false,
			"err": ctx.MakeError("direct_no_response", "response: undefined"),
		}, nil
	}

	if fm, ok := fetched.(map[string]any); ok {
		status := ToInt(vs.GetProp(fm, "status"))
		headers := vs.GetProp(fm, "headers")

		// No-body responses (204, 304) and explicit zero content-length
		// must skip JSON parsing — calling json() on an empty body errors.
		var contentLength string
		if hm, ok := headers.(map[string]any); ok {
			if cl, ok := hm["content-length"]; ok {
				contentLength = fmt.Sprintf("%v", cl)
			}
		}
		noBody := status == 204 || status == 304 || contentLength == "0"

		var jsonData any
		if !noBody {
			if jf := vs.GetProp(fm, "json"); jf != nil {
				if f, ok := jf.(func() any); ok {
					// f() returns nil on parse error in our fetcher.
					jsonData = f()
				}
			}
		}

		return map[string]any{
			"ok":      status >= 200 && status < 300,
			"status":  status,
			"headers": headers,
			"data":    jsonData,
		}, nil
	}

	return map[string]any{"ok": false, "err": ctx.MakeError("direct_invalid", "invalid response type")}, nil
}


func (sdk *ArtInstituteOfChicagoSDK) Agent(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewAgentEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) AgentRole(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewAgentRoleEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) AgentType(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewAgentTypeEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) Article(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewArticleEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) Artwork(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewArtworkEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) ArtworkDateQualifier(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewArtworkDateQualifierEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) ArtworkPlaceQualifier(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewArtworkPlaceQualifierEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) ArtworkType(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewArtworkTypeEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) CategoryTerm(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewCategoryTermEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) DigitalPublication(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewDigitalPublicationEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) DigitalPublicationArticle(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewDigitalPublicationArticleEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) EducatorResource(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewEducatorResourceEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) Event(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewEventEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) EventOccurrence(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewEventOccurrenceEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) EventProgram(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewEventProgramEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) Exhibition(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewExhibitionEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) Gallery(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewGalleryEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) GenericPage(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewGenericPageEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) Highlight(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewHighlightEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) Hour(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewHourEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) Image(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewImageEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) LandingPage(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewLandingPageEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) Place(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewPlaceEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) PressRelease(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewPressReleaseEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) PrintedPublication(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewPrintedPublicationEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) Product(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewProductEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) Publication(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewPublicationEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) Search(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewSearchEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) Section(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewSectionEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) Site(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewSiteEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) Sound(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewSoundEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) StaticPage(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewStaticPageEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) Text(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewTextEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) Tour(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewTourEntityFunc(sdk, data)
}


func (sdk *ArtInstituteOfChicagoSDK) Video(data map[string]any) ArtInstituteOfChicagoEntity {
	return NewVideoEntityFunc(sdk, data)
}



func TestSDK(testopts map[string]any, sdkopts map[string]any) *ArtInstituteOfChicagoSDK {
	if sdkopts == nil {
		sdkopts = map[string]any{}
	}
	sdkopts = vs.Clone(sdkopts).(map[string]any)

	if testopts == nil {
		testopts = map[string]any{}
	}
	testopts = vs.Clone(testopts).(map[string]any)
	testopts["active"] = true

	vs.SetPath(sdkopts, []any{"feature", "test"}, testopts)

	sdk := NewArtInstituteOfChicagoSDK(sdkopts)
	sdk.Mode = "test"

	return sdk
}
