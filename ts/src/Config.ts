
import { BaseFeature } from './feature/base/BaseFeature'
import { TestFeature } from './feature/test/TestFeature'



const FEATURE_CLASS: Record<string, typeof BaseFeature> = {
   test: TestFeature,

}


class Config {

  makeFeature(this: any, fn: string) {
    const fc = FEATURE_CLASS[fn]
    const fi = new fc()
    // TODO: errors etc
    return fi
  }


  main = {
    name: 'ArtInstituteOfChicago',
  }


  feature = {
     test:     {
      "options": {
        "active": false
      }
    },

  }


  options = {
    base: "https://api.artic.edu/api/v1",

    headers: {
      "content-type": "application/json"
    },

    entity: {
      
      agent: {
      },

      agent_role: {
      },

      agent_type: {
      },

      article: {
      },

      artwork: {
      },

      artwork_date_qualifier: {
      },

      artwork_place_qualifier: {
      },

      artwork_type: {
      },

      category_term: {
      },

      digital_publication: {
      },

      digital_publication_article: {
      },

      educator_resource: {
      },

      event: {
      },

      event_occurrence: {
      },

      event_program: {
      },

      exhibition: {
      },

      gallery: {
      },

      generic_page: {
      },

      highlight: {
      },

      hour: {
      },

      image: {
      },

      landing_page: {
      },

      place: {
      },

      press_release: {
      },

      printed_publication: {
      },

      product: {
      },

      publication: {
      },

      search: {
      },

      section: {
      },

      site: {
      },

      sound: {
      },

      static_page: {
      },

      text: {
      },

      tour: {
      },

      video: {
      },

    }
  }


  entity = {
    "agent": {
      "fields": [
        {
          "name": "alt_titles",
          "type": "`$ANY`"
        },
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "birth_date",
          "type": "`$ANY`"
        },
        {
          "name": "death_date",
          "type": "`$ANY`"
        },
        {
          "name": "description",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "is_artist",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "sort_title",
          "type": "`$ANY`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "ulan_id",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        }
      ],
      "name": "agent",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/agents",
              "parts": [
                "agents"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/artists",
              "parts": [
                "artists"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/agents/{id}",
              "parts": [
                "agents",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/artists/{id}",
              "parts": [
                "artists",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "agent_role": {
      "fields": [
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        }
      ],
      "name": "agent_role",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/agent-roles",
              "parts": [
                "agent-roles"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/agent-roles/{id}",
              "parts": [
                "agent-roles",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "agent_type": {
      "fields": [
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        }
      ],
      "name": "agent_type",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/agent-types",
              "parts": [
                "agent-types"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/agent-types/{id}",
              "parts": [
                "agent-types",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "article": {
      "fields": [
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "copy",
          "type": "`$ANY`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        }
      ],
      "name": "article",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/articles",
              "parts": [
                "articles"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/articles/{id}",
              "parts": [
                "articles",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "artwork": {
      "fields": [
        {
          "name": "alt_artist_ids",
          "type": "`$ANY`"
        },
        {
          "name": "alt_classification_ids",
          "type": "`$ANY`"
        },
        {
          "name": "alt_image_ids",
          "type": "`$ANY`"
        },
        {
          "name": "alt_material_ids",
          "type": "`$ANY`"
        },
        {
          "name": "alt_style_ids",
          "type": "`$ANY`"
        },
        {
          "name": "alt_subject_ids",
          "type": "`$ANY`"
        },
        {
          "name": "alt_technique_ids",
          "type": "`$ANY`"
        },
        {
          "name": "alt_titles",
          "type": "`$ANY`"
        },
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "artist_display",
          "type": "`$ANY`"
        },
        {
          "name": "artist_id",
          "type": "`$STRING`"
        },
        {
          "name": "artist_ids",
          "type": "`$ANY`"
        },
        {
          "name": "artist_title",
          "type": "`$ANY`"
        },
        {
          "name": "artist_titles",
          "type": "`$ANY`"
        },
        {
          "name": "artwork_type_id",
          "type": "`$STRING`"
        },
        {
          "name": "artwork_type_title",
          "type": "`$ANY`"
        },
        {
          "name": "boost_rank",
          "type": "`$ANY`"
        },
        {
          "name": "catalog_based_search_keyword_titles",
          "type": "`$ANY`"
        },
        {
          "name": "catalogue_display",
          "type": "`$ANY`"
        },
        {
          "name": "category_ids",
          "type": "`$ANY`"
        },
        {
          "name": "category_titles",
          "type": "`$ANY`"
        },
        {
          "name": "classification_id",
          "type": "`$STRING`"
        },
        {
          "name": "classification_ids",
          "type": "`$ANY`"
        },
        {
          "name": "classification_title",
          "type": "`$ANY`"
        },
        {
          "name": "classification_titles",
          "type": "`$ANY`"
        },
        {
          "name": "color",
          "type": "`$ANY`"
        },
        {
          "name": "colorfulness",
          "type": "`$ANY`"
        },
        {
          "name": "copyright_notice",
          "type": "`$ANY`"
        },
        {
          "name": "credit_line",
          "type": "`$ANY`"
        },
        {
          "name": "date_display",
          "type": "`$ANY`"
        },
        {
          "name": "date_end",
          "type": "`$ANY`"
        },
        {
          "name": "date_qualifier_id",
          "type": "`$STRING`"
        },
        {
          "name": "date_qualifier_title",
          "type": "`$ANY`"
        },
        {
          "name": "date_start",
          "type": "`$ANY`"
        },
        {
          "name": "department_id",
          "type": "`$STRING`"
        },
        {
          "name": "department_title",
          "type": "`$ANY`"
        },
        {
          "name": "description",
          "type": "`$STRING`"
        },
        {
          "name": "dimensions",
          "type": "`$ANY`"
        },
        {
          "name": "dimensions_detail",
          "type": "`$ANY`"
        },
        {
          "name": "document_ids",
          "type": "`$ANY`"
        },
        {
          "name": "edition",
          "type": "`$ANY`"
        },
        {
          "name": "exhibition_history",
          "type": "`$ANY`"
        },
        {
          "name": "fiscal_year",
          "type": "`$ANY`"
        },
        {
          "name": "fiscal_year_deaccession",
          "type": "`$ANY`"
        },
        {
          "name": "gallery_id",
          "type": "`$STRING`"
        },
        {
          "name": "gallery_title",
          "type": "`$ANY`"
        },
        {
          "name": "has_advanced_imaging",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "has_educational_resources",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "has_multimedia_resources",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "has_not_been_viewed_much",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "image_embedding",
          "type": "`$ANY`"
        },
        {
          "name": "image_id",
          "type": "`$STRING`"
        },
        {
          "name": "inscriptions",
          "type": "`$ANY`"
        },
        {
          "name": "internal_department_id",
          "type": "`$STRING`"
        },
        {
          "name": "is_boosted",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "is_on_view",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "is_public_domain",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "is_zoomable",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "latitude",
          "type": "`$NUMBER`"
        },
        {
          "name": "latlon",
          "type": "`$ANY`"
        },
        {
          "name": "longitude",
          "type": "`$NUMBER`"
        },
        {
          "name": "main_reference_number",
          "type": "`$INTEGER`"
        },
        {
          "name": "material_id",
          "type": "`$STRING`"
        },
        {
          "name": "material_ids",
          "type": "`$ANY`"
        },
        {
          "name": "material_titles",
          "type": "`$ANY`"
        },
        {
          "name": "max_zoom_window_size",
          "type": "`$ANY`"
        },
        {
          "name": "medium_display",
          "type": "`$ANY`"
        },
        {
          "name": "nomisma_id",
          "type": "`$STRING`"
        },
        {
          "name": "on_loan_display",
          "type": "`$ANY`"
        },
        {
          "name": "pageviews",
          "type": "`$ANY`"
        },
        {
          "name": "pageviews_recent",
          "type": "`$ANY`"
        },
        {
          "name": "place_of_origin",
          "type": "`$ANY`"
        },
        {
          "name": "provenance_text",
          "type": "`$ANY`"
        },
        {
          "name": "publication_history",
          "type": "`$ANY`"
        },
        {
          "name": "publishing_verification_level",
          "type": "`$ANY`"
        },
        {
          "name": "section_ids",
          "type": "`$ANY`"
        },
        {
          "name": "section_titles",
          "type": "`$ANY`"
        },
        {
          "name": "short_description",
          "type": "`$ANY`"
        },
        {
          "name": "site_ids",
          "type": "`$ANY`"
        },
        {
          "name": "sound_ids",
          "type": "`$ANY`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "style_id",
          "type": "`$STRING`"
        },
        {
          "name": "style_ids",
          "type": "`$ANY`"
        },
        {
          "name": "style_title",
          "type": "`$ANY`"
        },
        {
          "name": "style_titles",
          "type": "`$ANY`"
        },
        {
          "name": "subject_id",
          "type": "`$STRING`"
        },
        {
          "name": "subject_ids",
          "type": "`$ANY`"
        },
        {
          "name": "subject_titles",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "technique_id",
          "type": "`$STRING`"
        },
        {
          "name": "technique_ids",
          "type": "`$ANY`"
        },
        {
          "name": "technique_titles",
          "type": "`$ANY`"
        },
        {
          "name": "term_titles",
          "type": "`$ANY`"
        },
        {
          "name": "text_embedding",
          "type": "`$ANY`"
        },
        {
          "name": "text_ids",
          "type": "`$ANY`"
        },
        {
          "name": "theme_titles",
          "type": "`$ANY`"
        },
        {
          "name": "thumbnail",
          "type": "`$ANY`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "video_ids",
          "type": "`$ANY`"
        }
      ],
      "name": "artwork",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/artworks",
              "parts": [
                "artworks"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/artworks/{id}",
              "parts": [
                "artworks",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "artwork_date_qualifier": {
      "fields": [
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        }
      ],
      "name": "artwork_date_qualifier",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/artwork-date-qualifiers",
              "parts": [
                "artwork-date-qualifiers"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/artwork-date-qualifiers/{id}",
              "parts": [
                "artwork-date-qualifiers",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "artwork_place_qualifier": {
      "fields": [
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        }
      ],
      "name": "artwork_place_qualifier",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/artwork-place-qualifiers",
              "parts": [
                "artwork-place-qualifiers"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/artwork-place-qualifiers/{id}",
              "parts": [
                "artwork-place-qualifiers",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "artwork_type": {
      "fields": [
        {
          "name": "aat_id",
          "type": "`$STRING`"
        },
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        }
      ],
      "name": "artwork_type",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/artwork-types",
              "parts": [
                "artwork-types"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/artwork-types/{id}",
              "parts": [
                "artwork-types",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "category_term": {
      "fields": [
        {
          "name": "aat_id",
          "type": "`$STRING`"
        },
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "parent_id",
          "type": "`$STRING`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "subtype",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        }
      ],
      "name": "category_term",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/category-terms",
              "parts": [
                "category-terms"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/category-terms/{id}",
              "parts": [
                "category-terms",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "digital_publication": {
      "fields": [
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "copy",
          "type": "`$ANY`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "web_url",
          "type": "`$ANY`"
        }
      ],
      "name": "digital_publication",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/digital-publications",
              "parts": [
                "digital-publications"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/digital-publications/{id}",
              "parts": [
                "digital-publications",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "digital_publication_article": {
      "fields": [
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "author_display",
          "type": "`$ANY`"
        },
        {
          "name": "copy",
          "type": "`$ANY`"
        },
        {
          "name": "digital_publication_id",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "web_url",
          "type": "`$ANY`"
        }
      ],
      "name": "digital_publication_article",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/digital-publication-articles",
              "parts": [
                "digital-publication-articles"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/digital-publication-articles/{id}",
              "parts": [
                "digital-publication-articles",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "educator_resource": {
      "fields": [
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "copy",
          "type": "`$ANY`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "web_url",
          "type": "`$ANY`"
        }
      ],
      "name": "educator_resource",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/educator-resources",
              "parts": [
                "educator-resources"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/educator-resources/{id}",
              "parts": [
                "educator-resources",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "event": {
      "fields": [
        {
          "name": "alt_audience_ids",
          "type": "`$ANY`"
        },
        {
          "name": "alt_event_type_ids",
          "type": "`$ANY`"
        },
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "audience_id",
          "type": "`$STRING`"
        },
        {
          "name": "buy_button_caption",
          "type": "`$ANY`"
        },
        {
          "name": "buy_button_text",
          "type": "`$ANY`"
        },
        {
          "name": "date_display",
          "type": "`$ANY`"
        },
        {
          "name": "description",
          "type": "`$STRING`"
        },
        {
          "name": "door_time",
          "type": "`$ANY`"
        },
        {
          "name": "end_date",
          "type": "`$ANY`"
        },
        {
          "name": "end_time",
          "type": "`$ANY`"
        },
        {
          "name": "entrance",
          "type": "`$ANY`"
        },
        {
          "name": "event_host_id",
          "type": "`$STRING`"
        },
        {
          "name": "event_host_title",
          "type": "`$ANY`"
        },
        {
          "name": "event_type_id",
          "type": "`$STRING`"
        },
        {
          "name": "header_description",
          "type": "`$ANY`"
        },
        {
          "name": "hero_caption",
          "type": "`$ANY`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "image_url",
          "type": "`$ANY`"
        },
        {
          "name": "is_admission_required",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "is_after_hours",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "is_free",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "is_member_exclusive",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "is_private",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "is_registration_required",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "is_sales_button_hidden",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "is_sold_out",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "is_ticketed",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "is_virtual_event",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "join_url",
          "type": "`$ANY`"
        },
        {
          "name": "layout_type",
          "type": "`$ANY`"
        },
        {
          "name": "list_description",
          "type": "`$ANY`"
        },
        {
          "name": "location",
          "type": "`$ANY`"
        },
        {
          "name": "program_ids",
          "type": "`$ANY`"
        },
        {
          "name": "program_titles",
          "type": "`$ANY`"
        },
        {
          "name": "rsvp_link",
          "type": "`$ANY`"
        },
        {
          "name": "search_tags",
          "type": "`$ANY`"
        },
        {
          "name": "short_description",
          "type": "`$ANY`"
        },
        {
          "name": "slug",
          "type": "`$STRING`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "start_date",
          "type": "`$ANY`"
        },
        {
          "name": "start_time",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "survey_url",
          "type": "`$ANY`"
        },
        {
          "name": "ticketed_event_id",
          "type": "`$STRING`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "title_display",
          "type": "`$ANY`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "virtual_event_passcode",
          "type": "`$ANY`"
        },
        {
          "name": "virtual_event_url",
          "type": "`$ANY`"
        }
      ],
      "name": "event",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/events",
              "parts": [
                "events"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/events/{id}",
              "parts": [
                "events",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "event_occurrence": {
      "fields": [
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "button_caption",
          "type": "`$ANY`"
        },
        {
          "name": "button_text",
          "type": "`$ANY`"
        },
        {
          "name": "button_url",
          "type": "`$ANY`"
        },
        {
          "name": "description",
          "type": "`$STRING`"
        },
        {
          "name": "end_at",
          "type": "`$ANY`"
        },
        {
          "name": "event_id",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "image_url",
          "type": "`$ANY`"
        },
        {
          "name": "is_private",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "is_sales_button_hidden",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "is_ticketed",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "location",
          "type": "`$ANY`"
        },
        {
          "name": "off_sale_at",
          "type": "`$ANY`"
        },
        {
          "name": "on_sale_at",
          "type": "`$ANY`"
        },
        {
          "name": "short_description",
          "type": "`$ANY`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "start_at",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "title_display",
          "type": "`$ANY`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        }
      ],
      "name": "event_occurrence",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/event-occurrences",
              "parts": [
                "event-occurrences"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/event-occurrences/{id}",
              "parts": [
                "event-occurrences",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "event_program": {
      "fields": [
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "is_affiliate_group",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "is_event_host",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        }
      ],
      "name": "event_program",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/event-programs",
              "parts": [
                "event-programs"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/event-programs/{id}",
              "parts": [
                "event-programs",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "exhibition": {
      "fields": [
        {
          "name": "aic_end_at",
          "type": "`$ANY`"
        },
        {
          "name": "aic_start_at",
          "type": "`$ANY`"
        },
        {
          "name": "alt_image_ids",
          "type": "`$ANY`"
        },
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "artist_ids",
          "type": "`$ANY`"
        },
        {
          "name": "artwork_ids",
          "type": "`$ANY`"
        },
        {
          "name": "artwork_titles",
          "type": "`$ANY`"
        },
        {
          "name": "document_ids",
          "type": "`$ANY`"
        },
        {
          "name": "gallery_id",
          "type": "`$STRING`"
        },
        {
          "name": "gallery_title",
          "type": "`$ANY`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "image_id",
          "type": "`$STRING`"
        },
        {
          "name": "image_url",
          "type": "`$ANY`"
        },
        {
          "name": "is_featured",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "is_published",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "position",
          "type": "`$ANY`"
        },
        {
          "name": "short_description",
          "type": "`$ANY`"
        },
        {
          "name": "site_ids",
          "type": "`$ANY`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "status",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "web_url",
          "type": "`$ANY`"
        }
      ],
      "name": "exhibition",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/exhibitions",
              "parts": [
                "exhibitions"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/exhibitions/{id}",
              "parts": [
                "exhibitions",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "gallery": {
      "fields": [
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "floor",
          "type": "`$ANY`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "is_closed",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "latitude",
          "type": "`$NUMBER`"
        },
        {
          "name": "latlon",
          "type": "`$ANY`"
        },
        {
          "name": "longitude",
          "type": "`$NUMBER`"
        },
        {
          "name": "number",
          "type": "`$ANY`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "tgn_id",
          "type": "`$STRING`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        }
      ],
      "name": "gallery",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/galleries",
              "parts": [
                "galleries"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/galleries/{id}",
              "parts": [
                "galleries",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "generic_page": {
      "fields": [
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "copy",
          "type": "`$ANY`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "search_tags",
          "type": "`$ANY`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "web_url",
          "type": "`$ANY`"
        }
      ],
      "name": "generic_page",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/generic-pages",
              "parts": [
                "generic-pages"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/generic-pages/{id}",
              "parts": [
                "generic-pages",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "highlight": {
      "fields": [
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "copy",
          "type": "`$ANY`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        }
      ],
      "name": "highlight",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/highlights",
              "parts": [
                "highlights"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/highlights/{id}",
              "parts": [
                "highlights",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "hour": {
      "fields": [
        {
          "name": "additional_text",
          "type": "`$ANY`"
        },
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "friday_is_closed",
          "type": "`$ANY`"
        },
        {
          "name": "friday_member_close",
          "type": "`$ANY`"
        },
        {
          "name": "friday_member_open",
          "type": "`$ANY`"
        },
        {
          "name": "friday_public_close",
          "type": "`$ANY`"
        },
        {
          "name": "friday_public_open",
          "type": "`$ANY`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "monday_is_closed",
          "type": "`$ANY`"
        },
        {
          "name": "monday_member_close",
          "type": "`$ANY`"
        },
        {
          "name": "monday_member_open",
          "type": "`$ANY`"
        },
        {
          "name": "monday_public_close",
          "type": "`$ANY`"
        },
        {
          "name": "monday_public_open",
          "type": "`$ANY`"
        },
        {
          "name": "saturday_is_closed",
          "type": "`$ANY`"
        },
        {
          "name": "saturday_member_close",
          "type": "`$ANY`"
        },
        {
          "name": "saturday_member_open",
          "type": "`$ANY`"
        },
        {
          "name": "saturday_public_close",
          "type": "`$ANY`"
        },
        {
          "name": "saturday_public_open",
          "type": "`$ANY`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "summary",
          "type": "`$ANY`"
        },
        {
          "name": "sunday_is_closed",
          "type": "`$ANY`"
        },
        {
          "name": "sunday_member_close",
          "type": "`$ANY`"
        },
        {
          "name": "sunday_member_open",
          "type": "`$ANY`"
        },
        {
          "name": "sunday_public_close",
          "type": "`$ANY`"
        },
        {
          "name": "sunday_public_open",
          "type": "`$ANY`"
        },
        {
          "name": "thursday_is_closed",
          "type": "`$ANY`"
        },
        {
          "name": "thursday_member_close",
          "type": "`$ANY`"
        },
        {
          "name": "thursday_member_open",
          "type": "`$ANY`"
        },
        {
          "name": "thursday_public_close",
          "type": "`$ANY`"
        },
        {
          "name": "thursday_public_open",
          "type": "`$ANY`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "tuesday_is_closed",
          "type": "`$ANY`"
        },
        {
          "name": "tuesday_member_close",
          "type": "`$ANY`"
        },
        {
          "name": "tuesday_member_open",
          "type": "`$ANY`"
        },
        {
          "name": "tuesday_public_close",
          "type": "`$ANY`"
        },
        {
          "name": "tuesday_public_open",
          "type": "`$ANY`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "wednesday_is_closed",
          "type": "`$ANY`"
        },
        {
          "name": "wednesday_member_close",
          "type": "`$ANY`"
        },
        {
          "name": "wednesday_member_open",
          "type": "`$ANY`"
        },
        {
          "name": "wednesday_public_close",
          "type": "`$ANY`"
        },
        {
          "name": "wednesday_public_open",
          "type": "`$ANY`"
        }
      ],
      "name": "hour",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/hours",
              "parts": [
                "hours"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/hours/{id}",
              "parts": [
                "hours",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "image": {
      "fields": [
        {
          "name": "ahash",
          "type": "`$ANY`"
        },
        {
          "name": "alt_text",
          "type": "`$ANY`"
        },
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "artwork_ids",
          "type": "`$ANY`"
        },
        {
          "name": "artwork_titles",
          "type": "`$ANY`"
        },
        {
          "name": "color",
          "type": "`$ANY`"
        },
        {
          "name": "colorfulness",
          "type": "`$ANY`"
        },
        {
          "name": "content",
          "type": "`$ANY`"
        },
        {
          "name": "content_e_tag",
          "type": "`$ANY`"
        },
        {
          "name": "credit_line",
          "type": "`$ANY`"
        },
        {
          "name": "fingerprint",
          "type": "`$ANY`"
        },
        {
          "name": "height",
          "type": "`$NUMBER`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "iiif_url",
          "type": "`$ANY`"
        },
        {
          "name": "is_educational_resource",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "is_multimedia_resource",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "is_teacher_resource",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "lake_guid",
          "type": "`$ANY`"
        },
        {
          "name": "lqip",
          "type": "`$ANY`"
        },
        {
          "name": "phash",
          "type": "`$ANY`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "type",
          "type": "`$ANY`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "width",
          "type": "`$NUMBER`"
        }
      ],
      "name": "image",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/images",
              "parts": [
                "images"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/images/{id}",
              "parts": [
                "images",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "landing_page": {
      "fields": [
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "copy",
          "type": "`$ANY`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "search_tags",
          "type": "`$ANY`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "web_url",
          "type": "`$ANY`"
        }
      ],
      "name": "landing_page",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/landing-pages",
              "parts": [
                "landing-pages"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/landing-pages/{id}",
              "parts": [
                "landing-pages",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "place": {
      "fields": [
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "latitude",
          "type": "`$NUMBER`"
        },
        {
          "name": "longitude",
          "type": "`$NUMBER`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "tgn_id",
          "type": "`$STRING`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        }
      ],
      "name": "place",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/places",
              "parts": [
                "places"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/places/{id}",
              "parts": [
                "places",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "press_release": {
      "fields": [
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "copy",
          "type": "`$ANY`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "web_url",
          "type": "`$ANY`"
        }
      ],
      "name": "press_release",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/press-releases",
              "parts": [
                "press-releases"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/press-releases/{id}",
              "parts": [
                "press-releases",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "printed_publication": {
      "fields": [
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "copy",
          "type": "`$ANY`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "web_url",
          "type": "`$ANY`"
        }
      ],
      "name": "printed_publication",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/printed-publications",
              "parts": [
                "printed-publications"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/printed-publications/{id}",
              "parts": [
                "printed-publications",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "product": {
      "fields": [
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "artist_ids",
          "type": "`$ANY`"
        },
        {
          "name": "artwork_ids",
          "type": "`$ANY`"
        },
        {
          "name": "description",
          "type": "`$STRING`"
        },
        {
          "name": "exhibition_ids",
          "type": "`$ANY`"
        },
        {
          "name": "external_sku",
          "type": "`$ANY`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "image_url",
          "type": "`$ANY`"
        },
        {
          "name": "max_compare_at_price",
          "type": "`$ANY`"
        },
        {
          "name": "max_current_price",
          "type": "`$ANY`"
        },
        {
          "name": "min_compare_at_price",
          "type": "`$ANY`"
        },
        {
          "name": "min_current_price",
          "type": "`$ANY`"
        },
        {
          "name": "price_display",
          "type": "`$ANY`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "web_url",
          "type": "`$ANY`"
        }
      ],
      "name": "product",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/products",
              "parts": [
                "products"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/products/{id}",
              "parts": [
                "products",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "publication": {
      "fields": [
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "section_ids",
          "type": "`$ANY`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "web_url",
          "type": "`$ANY`"
        }
      ],
      "name": "publication",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/publications",
              "parts": [
                "publications"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/publications/{id}",
              "parts": [
                "publications",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "search": {
      "fields": [
        {
          "name": "api_id",
          "type": "`$STRING`"
        },
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "is_boosted",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "score",
          "type": "`$NUMBER`"
        },
        {
          "name": "thumbnail",
          "type": "`$ANY`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        }
      ],
      "name": "search",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "facet",
                    "orig": "facet",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/agents/search",
              "parts": [
                "agents",
                "search"
              ],
              "select": {
                "exist": [
                  "facet",
                  "from",
                  "q",
                  "query",
                  "size",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "facet",
                    "orig": "facet",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/articles/search",
              "parts": [
                "articles",
                "search"
              ],
              "select": {
                "exist": [
                  "facet",
                  "from",
                  "q",
                  "query",
                  "size",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "facet",
                    "orig": "facet",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/artworks/search",
              "parts": [
                "artworks",
                "search"
              ],
              "select": {
                "exist": [
                  "facet",
                  "from",
                  "q",
                  "query",
                  "size",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "facet",
                    "orig": "facet",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/category-terms/search",
              "parts": [
                "category-terms",
                "search"
              ],
              "select": {
                "exist": [
                  "facet",
                  "from",
                  "q",
                  "query",
                  "size",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "facet",
                    "orig": "facet",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/digital-publication-articles/search",
              "parts": [
                "digital-publication-articles",
                "search"
              ],
              "select": {
                "exist": [
                  "facet",
                  "from",
                  "q",
                  "query",
                  "size",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "facet",
                    "orig": "facet",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/digital-publications/search",
              "parts": [
                "digital-publications",
                "search"
              ],
              "select": {
                "exist": [
                  "facet",
                  "from",
                  "q",
                  "query",
                  "size",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "facet",
                    "orig": "facet",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/educator-resources/search",
              "parts": [
                "educator-resources",
                "search"
              ],
              "select": {
                "exist": [
                  "facet",
                  "from",
                  "q",
                  "query",
                  "size",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "facet",
                    "orig": "facet",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/event-occurrences/search",
              "parts": [
                "event-occurrences",
                "search"
              ],
              "select": {
                "exist": [
                  "facet",
                  "from",
                  "q",
                  "query",
                  "size",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "facet",
                    "orig": "facet",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/event-programs/search",
              "parts": [
                "event-programs",
                "search"
              ],
              "select": {
                "exist": [
                  "facet",
                  "from",
                  "q",
                  "query",
                  "size",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "facet",
                    "orig": "facet",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/events/search",
              "parts": [
                "events",
                "search"
              ],
              "select": {
                "exist": [
                  "facet",
                  "from",
                  "q",
                  "query",
                  "size",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "facet",
                    "orig": "facet",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/exhibitions/search",
              "parts": [
                "exhibitions",
                "search"
              ],
              "select": {
                "exist": [
                  "facet",
                  "from",
                  "q",
                  "query",
                  "size",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "facet",
                    "orig": "facet",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/galleries/search",
              "parts": [
                "galleries",
                "search"
              ],
              "select": {
                "exist": [
                  "facet",
                  "from",
                  "q",
                  "query",
                  "size",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "facet",
                    "orig": "facet",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/generic-pages/search",
              "parts": [
                "generic-pages",
                "search"
              ],
              "select": {
                "exist": [
                  "facet",
                  "from",
                  "q",
                  "query",
                  "size",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "facet",
                    "orig": "facet",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/highlights/search",
              "parts": [
                "highlights",
                "search"
              ],
              "select": {
                "exist": [
                  "facet",
                  "from",
                  "q",
                  "query",
                  "size",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "facet",
                    "orig": "facet",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/hours/search",
              "parts": [
                "hours",
                "search"
              ],
              "select": {
                "exist": [
                  "facet",
                  "from",
                  "q",
                  "query",
                  "size",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "facet",
                    "orig": "facet",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/images/search",
              "parts": [
                "images",
                "search"
              ],
              "select": {
                "exist": [
                  "facet",
                  "from",
                  "q",
                  "query",
                  "size",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "facet",
                    "orig": "facet",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/landing-pages/search",
              "parts": [
                "landing-pages",
                "search"
              ],
              "select": {
                "exist": [
                  "facet",
                  "from",
                  "q",
                  "query",
                  "size",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "facet",
                    "orig": "facet",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/mobile-sounds/search",
              "parts": [
                "mobile-sounds",
                "search"
              ],
              "select": {
                "exist": [
                  "facet",
                  "from",
                  "q",
                  "query",
                  "size",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "facet",
                    "orig": "facet",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/places/search",
              "parts": [
                "places",
                "search"
              ],
              "select": {
                "exist": [
                  "facet",
                  "from",
                  "q",
                  "query",
                  "size",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "facet",
                    "orig": "facet",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/press-releases/search",
              "parts": [
                "press-releases",
                "search"
              ],
              "select": {
                "exist": [
                  "facet",
                  "from",
                  "q",
                  "query",
                  "size",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "facet",
                    "orig": "facet",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/printed-publications/search",
              "parts": [
                "printed-publications",
                "search"
              ],
              "select": {
                "exist": [
                  "facet",
                  "from",
                  "q",
                  "query",
                  "size",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "facet",
                    "orig": "facet",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/products/search",
              "parts": [
                "products",
                "search"
              ],
              "select": {
                "exist": [
                  "facet",
                  "from",
                  "q",
                  "query",
                  "size",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "facet",
                    "orig": "facet",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/publications/search",
              "parts": [
                "publications",
                "search"
              ],
              "select": {
                "exist": [
                  "facet",
                  "from",
                  "q",
                  "query",
                  "size",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "facet",
                    "orig": "facet",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/search",
              "parts": [
                "search"
              ],
              "select": {
                "exist": [
                  "facet",
                  "from",
                  "q",
                  "query",
                  "size",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "facet",
                    "orig": "facet",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/sections/search",
              "parts": [
                "sections",
                "search"
              ],
              "select": {
                "exist": [
                  "facet",
                  "from",
                  "q",
                  "query",
                  "size",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "facet",
                    "orig": "facet",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/sites/search",
              "parts": [
                "sites",
                "search"
              ],
              "select": {
                "exist": [
                  "facet",
                  "from",
                  "q",
                  "query",
                  "size",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "facet",
                    "orig": "facet",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/sounds/search",
              "parts": [
                "sounds",
                "search"
              ],
              "select": {
                "exist": [
                  "facet",
                  "from",
                  "q",
                  "query",
                  "size",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "facet",
                    "orig": "facet",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/static-pages/search",
              "parts": [
                "static-pages",
                "search"
              ],
              "select": {
                "exist": [
                  "facet",
                  "from",
                  "q",
                  "query",
                  "size",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "facet",
                    "orig": "facet",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/texts/search",
              "parts": [
                "texts",
                "search"
              ],
              "select": {
                "exist": [
                  "facet",
                  "from",
                  "q",
                  "query",
                  "size",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "facet",
                    "orig": "facet",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/tours/search",
              "parts": [
                "tours",
                "search"
              ],
              "select": {
                "exist": [
                  "facet",
                  "from",
                  "q",
                  "query",
                  "size",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "facet",
                    "orig": "facet",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/videos/search",
              "parts": [
                "videos",
                "search"
              ],
              "select": {
                "exist": [
                  "facet",
                  "from",
                  "q",
                  "query",
                  "size",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "section": {
      "fields": [
        {
          "name": "accession",
          "type": "`$ANY`"
        },
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "artwork_id",
          "type": "`$STRING`"
        },
        {
          "name": "content",
          "type": "`$ANY`"
        },
        {
          "name": "generic_page_id",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "publication_id",
          "type": "`$STRING`"
        },
        {
          "name": "publication_title",
          "type": "`$ANY`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "web_url",
          "type": "`$ANY`"
        }
      ],
      "name": "section",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/sections",
              "parts": [
                "sections"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/sections/{id}",
              "parts": [
                "sections",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "site": {
      "fields": [
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "artwork_ids",
          "type": "`$ANY`"
        },
        {
          "name": "artwork_titles",
          "type": "`$ANY`"
        },
        {
          "name": "description",
          "type": "`$STRING`"
        },
        {
          "name": "exhibition_ids",
          "type": "`$ANY`"
        },
        {
          "name": "exhibition_titles",
          "type": "`$ANY`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "web_url",
          "type": "`$ANY`"
        }
      ],
      "name": "site",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/sites",
              "parts": [
                "sites"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/sites/{id}",
              "parts": [
                "sites",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "sound": {
      "fields": [
        {
          "name": "alt_text",
          "type": "`$ANY`"
        },
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "artwork_ids",
          "type": "`$ANY`"
        },
        {
          "name": "artwork_titles",
          "type": "`$ANY`"
        },
        {
          "name": "content",
          "type": "`$ANY`"
        },
        {
          "name": "content_e_tag",
          "type": "`$ANY`"
        },
        {
          "name": "credit_line",
          "type": "`$ANY`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "is_educational_resource",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "is_multimedia_resource",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "is_teacher_resource",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "lake_guid",
          "type": "`$ANY`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "transcript",
          "type": "`$ANY`"
        },
        {
          "name": "type",
          "type": "`$ANY`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "web_url",
          "type": "`$ANY`"
        }
      ],
      "name": "sound",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/mobile-sounds",
              "parts": [
                "mobile-sounds"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/sounds",
              "parts": [
                "sounds"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/mobile-sounds/{id}",
              "parts": [
                "mobile-sounds",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/sounds/{id}",
              "parts": [
                "sounds",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "static_page": {
      "fields": [
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "web_url",
          "type": "`$ANY`"
        }
      ],
      "name": "static_page",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/static-pages",
              "parts": [
                "static-pages"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/static-pages/{id}",
              "parts": [
                "static-pages",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "text": {
      "fields": [
        {
          "name": "alt_text",
          "type": "`$ANY`"
        },
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "artwork_ids",
          "type": "`$ANY`"
        },
        {
          "name": "artwork_titles",
          "type": "`$ANY`"
        },
        {
          "name": "content",
          "type": "`$ANY`"
        },
        {
          "name": "content_e_tag",
          "type": "`$ANY`"
        },
        {
          "name": "credit_line",
          "type": "`$ANY`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "is_educational_resource",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "is_multimedia_resource",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "is_teacher_resource",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "lake_guid",
          "type": "`$ANY`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "type",
          "type": "`$ANY`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        }
      ],
      "name": "text",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/texts",
              "parts": [
                "texts"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/texts/{id}",
              "parts": [
                "texts",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "tour": {
      "fields": [
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "artist_titles",
          "type": "`$ANY`"
        },
        {
          "name": "artwork_titles",
          "type": "`$ANY`"
        },
        {
          "name": "description",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "image",
          "type": "`$ANY`"
        },
        {
          "name": "intro",
          "type": "`$ANY`"
        },
        {
          "name": "intro_link",
          "type": "`$ANY`"
        },
        {
          "name": "intro_transcript",
          "type": "`$ANY`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "weight",
          "type": "`$NUMBER`"
        }
      ],
      "name": "tour",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/tours",
              "parts": [
                "tours"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/tours/{id}",
              "parts": [
                "tours",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "video": {
      "fields": [
        {
          "name": "alt_text",
          "type": "`$ANY`"
        },
        {
          "name": "api_link",
          "type": "`$ANY`"
        },
        {
          "name": "api_model",
          "type": "`$ANY`"
        },
        {
          "name": "artwork_ids",
          "type": "`$ANY`"
        },
        {
          "name": "artwork_titles",
          "type": "`$ANY`"
        },
        {
          "name": "content",
          "type": "`$ANY`"
        },
        {
          "name": "content_e_tag",
          "type": "`$ANY`"
        },
        {
          "name": "credit_line",
          "type": "`$ANY`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "is_educational_resource",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "is_multimedia_resource",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "is_teacher_resource",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "lake_guid",
          "type": "`$ANY`"
        },
        {
          "name": "source_updated_at",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_all",
          "type": "`$ANY`"
        },
        {
          "name": "suggest_autocomplete_boosted",
          "type": "`$ANY`"
        },
        {
          "name": "timestamp",
          "type": "`$ANY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "type",
          "type": "`$ANY`"
        },
        {
          "name": "updated_at",
          "type": "`$ANY`"
        }
      ],
      "name": "video",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/videos",
              "parts": [
                "videos"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/videos/{id}",
              "parts": [
                "videos",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    }
  }
}


const config = new Config()

export {
  config
}

