
import { BaseFeature } from './feature/base/BaseFeature'
import { RatelimitFeature } from './feature/ratelimit/RatelimitFeature'
import { RetryFeature } from './feature/retry/RetryFeature'
import { TestFeature } from './feature/test/TestFeature'
import { TimeoutFeature } from './feature/timeout/TimeoutFeature'



const FEATURE_CLASS: Record<string, typeof BaseFeature> = {
   ratelimit: RatelimitFeature,
 retry: RetryFeature,
 test: TestFeature,
 timeout: TimeoutFeature,

}


const FEATURE_PLUGINS: Record<string, any[]> = {
  
}


class Config {

  makeFeature(this: any, fn: string) {
    const fc = FEATURE_CLASS[fn]
    const fi = new fc()
    return fi
  }

  // False for a feature added at runtime via options.extend (station's
  // adopt path) - the constructor uses this to skip makeFeature for names
  // no generated class backs.
  hasFeature(this: any, fn: string) {
    return null != FEATURE_CLASS[fn]
  }


  main = {
    name: 'ArtInstituteOfChicago',
        slug: "art-institute-of-chicago",
    version: "0.0.1",
    target: "ts",

  }


  feature = {
     ratelimit:     {
      "options": {
        "active": false,
        "burst": 5,
        "rate": 5
      },
      "optspec": {
        "now": "`$FUNCTION`",
        "sleep": "`$FUNCTION`"
      },
      "strict": false,
      "transport": "wrap"
    },
 retry:     {
      "options": {
        "active": false,
        "factor": 2,
        "maxDelay": 2000,
        "minDelay": 50,
        "retries": 2,
        "statuses": [
          408,
          425,
          429,
          500,
          502,
          503,
          504
        ]
      },
      "optspec": {
        "jitter": "`$BOOLEAN`",
        "sleep": "`$FUNCTION`"
      },
      "strict": false,
      "transport": "wrap"
    },
 test:     {
      "options": {
        "active": false
      },
      "optspec": {
        "entity": "`$MAP`",
        "net": "`$MAP`"
      },
      "strict": false,
      "transport": "base"
    },
 timeout:     {
      "options": {
        "active": false,
        "ms": 30000
      },
      "optspec": {
        "clearTimer": "`$FUNCTION`",
        "now": "`$FUNCTION`",
        "setTimer": "`$FUNCTION`"
      },
      "strict": false,
      "transport": "wrap"
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
          "title": "Alt Titles",
          "type": "`$ANY`",
          "short": "Alternate names for this agent"
        },
        {
          "name": "api_link",
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "birth_date",
          "title": "Birth Date",
          "type": "`$ANY`",
          "short": "The year this agent was born"
        },
        {
          "name": "death_date",
          "title": "Death Date",
          "type": "`$ANY`",
          "short": "The year this agent died"
        },
        {
          "name": "description",
          "title": "Description",
          "type": "`$STRING`",
          "short": "A biographical description of the agent"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "is_artist",
          "title": "Is Artist",
          "type": "`$BOOLEAN`",
          "short": "Whether the agent is an artist."
        },
        {
          "name": "sort_title",
          "title": "Sort Title",
          "type": "`$ANY`",
          "short": "Sortable name for this agent, typically with last name first."
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "ulan_id",
          "title": "Ulan Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this agent in Getty's ULAN"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "agent",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/agents",
              "segments": [
                {
                  "lit": "agents"
                }
              ],
              "parts": [
                "agents"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/artists",
              "segments": [
                {
                  "lit": "artists"
                }
              ],
              "parts": [
                "artists"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/agents/{id}",
              "segments": [
                {
                  "lit": "agents"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "agents",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
              }
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/artists/{id}",
              "segments": [
                {
                  "lit": "artists"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "artists",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "agent_role",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/agent-roles",
              "segments": [
                {
                  "lit": "agent-roles"
                }
              ],
              "parts": [
                "agent-roles"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/agent-roles/{id}",
              "segments": [
                {
                  "lit": "agent-roles"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "agent-roles",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "agent_type",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/agent-types",
              "segments": [
                {
                  "lit": "agent-types"
                }
              ],
              "parts": [
                "agent-types"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/agent-types/{id}",
              "segments": [
                {
                  "lit": "agent-types"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "agent-types",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "copy",
          "title": "Copy",
          "type": "`$ANY`",
          "short": "The text of the article"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "article",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/articles",
              "segments": [
                {
                  "lit": "articles"
                }
              ],
              "parts": [
                "articles"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/articles/{id}",
              "segments": [
                {
                  "lit": "articles"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "articles",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Alt Artist Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of the non-preferred artists/cultures associated with this work"
        },
        {
          "name": "alt_classification_ids",
          "title": "Alt Classification Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of all other non-preferred classification terms for this work"
        },
        {
          "name": "alt_image_ids",
          "title": "Alt Image Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of all non-preferred images of this work."
        },
        {
          "name": "alt_material_ids",
          "title": "Alt Material Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of all other non-preferred material terms for this work"
        },
        {
          "name": "alt_style_ids",
          "title": "Alt Style Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of all other non-preferred style terms for this work"
        },
        {
          "name": "alt_subject_ids",
          "title": "Alt Subject Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of all other non-preferred subject terms for this work"
        },
        {
          "name": "alt_technique_ids",
          "title": "Alt Technique Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of all other non-preferred technique terms for this work"
        },
        {
          "name": "alt_titles",
          "title": "Alt Titles",
          "type": "`$ANY`",
          "short": "Alternate names for this work"
        },
        {
          "name": "api_link",
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "artist_display",
          "title": "Artist Display",
          "type": "`$ANY`",
          "short": "Readable description of the creator of this work."
        },
        {
          "name": "artist_id",
          "title": "Artist Id",
          "type": "`$STRING`",
          "short": "Unique identifier of the preferred artist/culture associated with this work"
        },
        {
          "name": "artist_ids",
          "title": "Artist Ids",
          "type": "`$ANY`",
          "short": "Unique identifier of all artist/cultures associated with this work"
        },
        {
          "name": "artist_title",
          "title": "Artist Title",
          "type": "`$ANY`",
          "short": "Name of the preferred artist/culture associated with this work"
        },
        {
          "name": "artist_titles",
          "title": "Artist Titles",
          "type": "`$ANY`",
          "short": "Names of all artist/cultures associated with this work"
        },
        {
          "name": "artwork_type_id",
          "title": "Artwork Type Id",
          "type": "`$STRING`",
          "short": "Unique identifier of the kind of object or work"
        },
        {
          "name": "artwork_type_title",
          "title": "Artwork Type Title",
          "type": "`$ANY`",
          "short": "The kind of object or work (e.g."
        },
        {
          "name": "boost_rank",
          "title": "Boost Rank",
          "type": "`$ANY`",
          "short": "Manual indication of what rank this artwork should take in search results."
        },
        {
          "name": "catalog_based_search_keyword_titles",
          "title": "Catalog Based Search Keyword Titles",
          "type": "`$ANY`",
          "short": "The keyword search values that would be catalog-based searches on this record"
        },
        {
          "name": "catalogue_display",
          "title": "Catalogue Display",
          "type": "`$ANY`",
          "short": "Brief text listing all the catalogues raisonnés which include this work."
        },
        {
          "name": "category_ids",
          "title": "Category Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of the categories this work is a part of"
        },
        {
          "name": "category_titles",
          "title": "Category Titles",
          "type": "`$ANY`",
          "short": "Names of the categories this artwork is a part of"
        },
        {
          "name": "classification_id",
          "title": "Classification Id",
          "type": "`$STRING`",
          "short": "Unique identifier of the preferred classification term for this work"
        },
        {
          "name": "classification_ids",
          "title": "Classification Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of all classification terms for this work"
        },
        {
          "name": "classification_title",
          "title": "Classification Title",
          "type": "`$ANY`",
          "short": "The name of the preferred classification term for this work"
        },
        {
          "name": "classification_titles",
          "title": "Classification Titles",
          "type": "`$ANY`",
          "short": "The names of all classification terms related to this artwork"
        },
        {
          "name": "color",
          "title": "Color",
          "type": "`$ANY`",
          "short": "Dominant color of this artwork in HSL"
        },
        {
          "name": "colorfulness",
          "title": "Colorfulness",
          "type": "`$ANY`",
          "short": "Unbounded positive float representing an abstract measure of colorfulness."
        },
        {
          "name": "copyright_notice",
          "title": "Copyright Notice",
          "type": "`$ANY`",
          "short": "Statement notifying how the work is protected by copyright."
        },
        {
          "name": "credit_line",
          "title": "Credit Line",
          "type": "`$ANY`",
          "short": "Brief statement indicating how the work came into the collection"
        },
        {
          "name": "date_display",
          "title": "Date Display",
          "type": "`$ANY`",
          "short": "Readable, free-text description of the period of time associated with the creation of this work."
        },
        {
          "name": "date_end",
          "title": "Date End",
          "type": "`$ANY`",
          "short": "The year of the period of time associated with the creation of this work"
        },
        {
          "name": "date_qualifier_id",
          "title": "Date Qualifier Id",
          "type": "`$STRING`",
          "short": "Unique identifier of the qualifer to the dates provided for this record."
        },
        {
          "name": "date_qualifier_title",
          "title": "Date Qualifier Title",
          "type": "`$ANY`",
          "short": "Readable, text qualifer to the dates provided for this record."
        },
        {
          "name": "date_start",
          "title": "Date Start",
          "type": "`$ANY`",
          "short": "The year of the period of time associated with the creation of this work"
        },
        {
          "name": "department_id",
          "title": "Department Id",
          "type": "`$STRING`",
          "short": "Unique identifier of the curatorial department that this work belongs to"
        },
        {
          "name": "department_title",
          "title": "Department Title",
          "type": "`$ANY`",
          "short": "Name of the curatorial department that this work belongs to"
        },
        {
          "name": "description",
          "title": "Description",
          "type": "`$STRING`",
          "short": "Longer explanation describing the work"
        },
        {
          "name": "dimensions",
          "title": "Dimensions",
          "type": "`$ANY`",
          "short": "The size, shape, scale, and dimensions of the work."
        },
        {
          "name": "dimensions_detail",
          "title": "Dimensions Detail",
          "type": "`$ANY`",
          "short": "The height, width, depth, and/or diameter of each section of the work in centimeters"
        },
        {
          "name": "document_ids",
          "title": "Document Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of assets that serve as documentation for this artwork"
        },
        {
          "name": "edition",
          "title": "Edition",
          "type": "`$ANY`",
          "short": "Edition number if the work is one of many"
        },
        {
          "name": "exhibition_history",
          "title": "Exhibition History",
          "type": "`$ANY`",
          "short": "List of all the places this work has been exhibited"
        },
        {
          "name": "fiscal_year",
          "title": "Fiscal Year",
          "type": "`$ANY`",
          "short": "The fiscal year in which the work was acquired."
        },
        {
          "name": "fiscal_year_deaccession",
          "title": "Fiscal Year Deaccession",
          "type": "`$ANY`",
          "short": "The fiscal year in which the work was deaccessioned."
        },
        {
          "name": "gallery_id",
          "title": "Gallery Id",
          "type": "`$STRING`",
          "short": "Unique identifier of the location of this work in our museum"
        },
        {
          "name": "gallery_title",
          "title": "Gallery Title",
          "type": "`$ANY`",
          "short": "The location of this work in our museum"
        },
        {
          "name": "has_advanced_imaging",
          "title": "Has Advanced Imaging",
          "type": "`$BOOLEAN`",
          "short": "Whether this artwork is enhanced with 3D models, 360 image sequences, Mirador views, etc."
        },
        {
          "name": "has_educational_resources",
          "title": "Has Educational Resources",
          "type": "`$BOOLEAN`",
          "short": "Whether this artwork has any documents tagged as educational"
        },
        {
          "name": "has_multimedia_resources",
          "title": "Has Multimedia Resources",
          "type": "`$BOOLEAN`",
          "short": "Whether this artwork has any associated microsites, digital publications, or documents tagged as multimedia"
        },
        {
          "name": "has_not_been_viewed_much",
          "title": "Has Not Been Viewed Much",
          "type": "`$BOOLEAN`",
          "short": "Whether the artwork hasn't been visited on our website very much"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "image_embedding",
          "title": "Image Embedding",
          "type": "`$ANY`",
          "short": "The generated embeddings describing the artwork image"
        },
        {
          "name": "image_id",
          "title": "Image Id",
          "type": "`$STRING`",
          "short": "Unique identifier of the preferred image to use to represent this work"
        },
        {
          "name": "inscriptions",
          "title": "Inscriptions",
          "type": "`$ANY`",
          "short": "A description of distinguishing or identifying physical markings that are on the work"
        },
        {
          "name": "internal_department_id",
          "title": "Internal Department Id",
          "type": "`$STRING`",
          "short": "An internal department id we use for analytics."
        },
        {
          "name": "is_boosted",
          "title": "Is Boosted",
          "type": "`$BOOLEAN`",
          "short": "Whether this document should be boosted in search"
        },
        {
          "name": "is_on_view",
          "title": "Is On View",
          "type": "`$BOOLEAN`",
          "short": "Whether the work is on display"
        },
        {
          "name": "is_public_domain",
          "title": "Is Public Domain",
          "type": "`$BOOLEAN`",
          "short": "Whether the work is in the public domain, meaning it was created before copyrights existed or has left the copyright term"
        },
        {
          "name": "is_zoomable",
          "title": "Is Zoomable",
          "type": "`$BOOLEAN`",
          "short": "Whether images of the work are allowed to be displayed in a zoomable interface."
        },
        {
          "name": "latitude",
          "title": "Latitude",
          "type": "`$NUMBER`",
          "short": "Latitude coordinate of the location of this work in our galleries"
        },
        {
          "name": "latlon",
          "title": "Latlon",
          "type": "`$ANY`",
          "short": "Latitude and longitude coordinates of the location of this work in our galleries"
        },
        {
          "name": "longitude",
          "title": "Longitude",
          "type": "`$NUMBER`",
          "short": "Longitude coordinate of the location of this work in our galleries"
        },
        {
          "name": "main_reference_number",
          "title": "Main Reference Number",
          "type": "`$INTEGER`",
          "short": "Unique identifier assigned to the artwork upon acquisition"
        },
        {
          "name": "material_id",
          "title": "Material Id",
          "type": "`$STRING`",
          "short": "Unique identifier of the preferred material term for this work"
        },
        {
          "name": "material_ids",
          "title": "Material Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of all material terms for this work"
        },
        {
          "name": "material_titles",
          "title": "Material Titles",
          "type": "`$ANY`",
          "short": "The names of all material terms related to this artwork"
        },
        {
          "name": "max_zoom_window_size",
          "title": "Max Zoom Window Size",
          "type": "`$ANY`",
          "short": "The maximum size of the window the image is allowed to be viewed in, in pixels."
        },
        {
          "name": "medium_display",
          "title": "Medium Display",
          "type": "`$ANY`",
          "short": "The substances or materials used in the creation of a work"
        },
        {
          "name": "nomisma_id",
          "title": "Nomisma Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this work in the nomisma coin database"
        },
        {
          "name": "on_loan_display",
          "title": "On Loan Display",
          "type": "`$ANY`",
          "short": "If an artwork is on loan, this contains details about the loan"
        },
        {
          "name": "pageviews",
          "title": "Pageviews",
          "type": "`$ANY`",
          "short": "Approx."
        },
        {
          "name": "pageviews_recent",
          "title": "Pageviews Recent",
          "type": "`$ANY`",
          "short": "Approx."
        },
        {
          "name": "place_of_origin",
          "title": "Place Of Origin",
          "type": "`$ANY`",
          "short": "The location where the creation, design, or production of the work took place, or the original location of the work"
        },
        {
          "name": "provenance_text",
          "title": "Provenance Text",
          "type": "`$ANY`",
          "short": "Ownership/collecting history of the work."
        },
        {
          "name": "publication_history",
          "title": "Publication History",
          "type": "`$ANY`",
          "short": "Bibliographic list of all the places this work has been published"
        },
        {
          "name": "publishing_verification_level",
          "title": "Publishing Verification Level",
          "type": "`$ANY`",
          "short": "Indicator of how much metadata on the work in published."
        },
        {
          "name": "section_ids",
          "title": "Section Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of the digital publication chapters this work in included in"
        },
        {
          "name": "section_titles",
          "title": "Section Titles",
          "type": "`$ANY`",
          "short": "Names of the digital publication chapters this work is included in"
        },
        {
          "name": "short_description",
          "title": "Short Description",
          "type": "`$ANY`",
          "short": "Short explanation describing the work"
        },
        {
          "name": "site_ids",
          "title": "Site Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of the microsites this work is a part of"
        },
        {
          "name": "sound_ids",
          "title": "Sound Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of the audio about this work"
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "style_id",
          "title": "Style Id",
          "type": "`$STRING`",
          "short": "Unique identifier of the preferred style term for this work"
        },
        {
          "name": "style_ids",
          "title": "Style Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of all style terms for this work"
        },
        {
          "name": "style_title",
          "title": "Style Title",
          "type": "`$ANY`",
          "short": "The name of the preferred style term for this work"
        },
        {
          "name": "style_titles",
          "title": "Style Titles",
          "type": "`$ANY`",
          "short": "The names of all style terms related to this artwork"
        },
        {
          "name": "subject_id",
          "title": "Subject Id",
          "type": "`$STRING`",
          "short": "Unique identifier of the preferred subject term for this work"
        },
        {
          "name": "subject_ids",
          "title": "Subject Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of all subject terms for this work"
        },
        {
          "name": "subject_titles",
          "title": "Subject Titles",
          "type": "`$ANY`",
          "short": "The names of all subject terms related to this artwork"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "technique_id",
          "title": "Technique Id",
          "type": "`$STRING`",
          "short": "Unique identifier of the preferred technique term for this work"
        },
        {
          "name": "technique_ids",
          "title": "Technique Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of all technique terms for this work"
        },
        {
          "name": "technique_titles",
          "title": "Technique Titles",
          "type": "`$ANY`",
          "short": "The names of all technique terms related to this artwork"
        },
        {
          "name": "term_titles",
          "title": "Term Titles",
          "type": "`$ANY`",
          "short": "The names of the taxonomy tags for this work"
        },
        {
          "name": "text_embedding",
          "title": "Text Embedding",
          "type": "`$ANY`",
          "short": "The generated embeddings of artwork text"
        },
        {
          "name": "text_ids",
          "title": "Text Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of the texts about this work"
        },
        {
          "name": "theme_titles",
          "title": "Theme Titles",
          "type": "`$ANY`",
          "short": "The names of all thematic publish categories related to this artwork"
        },
        {
          "name": "thumbnail",
          "title": "Thumbnail",
          "type": "`$ANY`",
          "short": "Metadata about the image referenced by `image_id`."
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        },
        {
          "name": "video_ids",
          "title": "Video Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of the videos about this work"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "artwork",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/artworks",
              "segments": [
                {
                  "lit": "artworks"
                }
              ],
              "parts": [
                "artworks"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/artworks/{id}",
              "segments": [
                {
                  "lit": "artworks"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "artworks",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "artwork_date_qualifier",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/artwork-date-qualifiers",
              "segments": [
                {
                  "lit": "artwork-date-qualifiers"
                }
              ],
              "parts": [
                "artwork-date-qualifiers"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/artwork-date-qualifiers/{id}",
              "segments": [
                {
                  "lit": "artwork-date-qualifiers"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "artwork-date-qualifiers",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "artwork_place_qualifier",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/artwork-place-qualifiers",
              "segments": [
                {
                  "lit": "artwork-place-qualifiers"
                }
              ],
              "parts": [
                "artwork-place-qualifiers"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/artwork-place-qualifiers/{id}",
              "segments": [
                {
                  "lit": "artwork-place-qualifiers"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "artwork-place-qualifiers",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Aat Id",
          "type": "`$STRING`",
          "short": "Identifier of reconciled (most similar) term in the Getty's Art and Architecture Thesaurus (AAT)"
        },
        {
          "name": "api_link",
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "artwork_type",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/artwork-types",
              "segments": [
                {
                  "lit": "artwork-types"
                }
              ],
              "parts": [
                "artwork-types"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/artwork-types/{id}",
              "segments": [
                {
                  "lit": "artwork-types"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "artwork-types",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Aat Id",
          "type": "`$STRING`",
          "short": "Identifier of reconciled (most similar) term in the Getty's Art and Architecture Thesaurus (AAT)"
        },
        {
          "name": "api_link",
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "parent_id",
          "title": "Parent Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this category's parent"
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "subtype",
          "title": "Subtype",
          "type": "`$ANY`",
          "short": "Takes one of the following values: classification, material, technique, style, subject, department, theme"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "category_term",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/category-terms",
              "segments": [
                {
                  "lit": "category-terms"
                }
              ],
              "parts": [
                "category-terms"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/category-terms/{id}",
              "segments": [
                {
                  "lit": "category-terms"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "category-terms",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "copy",
          "title": "Copy",
          "type": "`$ANY`",
          "short": "The text of the page"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        },
        {
          "name": "web_url",
          "title": "Web Url",
          "type": "`$ANY`",
          "short": "The URL to this page on our website"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "digital_publication",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/digital-publications",
              "segments": [
                {
                  "lit": "digital-publications"
                }
              ],
              "parts": [
                "digital-publications"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/digital-publications/{id}",
              "segments": [
                {
                  "lit": "digital-publications"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "digital-publications",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "author_display",
          "title": "Author Display",
          "type": "`$ANY`",
          "short": "A display-friendly text of the authors of this article"
        },
        {
          "name": "copy",
          "title": "Copy",
          "type": "`$ANY`",
          "short": "The text of the article"
        },
        {
          "name": "digital_publication_id",
          "title": "Digital Publication Id",
          "type": "`$STRING`",
          "short": "Unique identifier of the digital publication this article belongs to"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        },
        {
          "name": "web_url",
          "title": "Web Url",
          "type": "`$ANY`",
          "short": "The URL to this article on our website"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "digital_publication_article",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/digital-publication-articles",
              "segments": [
                {
                  "lit": "digital-publication-articles"
                }
              ],
              "parts": [
                "digital-publication-articles"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/digital-publication-articles/{id}",
              "segments": [
                {
                  "lit": "digital-publication-articles"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "digital-publication-articles",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "copy",
          "title": "Copy",
          "type": "`$ANY`",
          "short": "The text of the page"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        },
        {
          "name": "web_url",
          "title": "Web Url",
          "type": "`$ANY`",
          "short": "The URL to this page on our website"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "educator_resource",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/educator-resources",
              "segments": [
                {
                  "lit": "educator-resources"
                }
              ],
              "parts": [
                "educator-resources"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/educator-resources/{id}",
              "segments": [
                {
                  "lit": "educator-resources"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "educator-resources",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Alt Audience Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers indicating the alternate audiences for this event"
        },
        {
          "name": "alt_event_type_ids",
          "title": "Alt Event Type Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers indicating the alternate types of this event"
        },
        {
          "name": "api_link",
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "audience_id",
          "title": "Audience Id",
          "type": "`$STRING`",
          "short": "Unique identifier indicating the preferred audience for this event"
        },
        {
          "name": "buy_button_caption",
          "title": "Buy Button Caption",
          "type": "`$ANY`",
          "short": "Additional text below the ticket/registration button"
        },
        {
          "name": "buy_button_text",
          "title": "Buy Button Text",
          "type": "`$ANY`",
          "short": "The text used on the ticket/registration button"
        },
        {
          "name": "date_display",
          "title": "Date Display",
          "type": "`$ANY`",
          "short": "A readable display of the event dates"
        },
        {
          "name": "description",
          "title": "Description",
          "type": "`$STRING`",
          "short": "All copytext of the event"
        },
        {
          "name": "door_time",
          "title": "Door Time",
          "type": "`$ANY`",
          "short": "The time the doors open for this event"
        },
        {
          "name": "end_date",
          "title": "End Date",
          "type": "`$ANY`",
          "short": "The date the event ends"
        },
        {
          "name": "end_time",
          "title": "End Time",
          "type": "`$ANY`",
          "short": "The time the event ends"
        },
        {
          "name": "entrance",
          "title": "Entrance",
          "type": "`$ANY`",
          "short": "Which entrance to use for this event"
        },
        {
          "name": "event_host_id",
          "title": "Event Host Id",
          "type": "`$STRING`",
          "short": "Unique identifier of the host (cf."
        },
        {
          "name": "event_host_title",
          "title": "Event Host Title",
          "type": "`$ANY`",
          "short": "Unique identifier of the host (cf."
        },
        {
          "name": "event_type_id",
          "title": "Event Type Id",
          "type": "`$STRING`",
          "short": "Unique identifier indicating the preferred type of this event"
        },
        {
          "name": "header_description",
          "title": "Header Description",
          "type": "`$ANY`",
          "short": "Brief description of the event displayed below the title"
        },
        {
          "name": "hero_caption",
          "title": "Hero Caption",
          "type": "`$ANY`",
          "short": "Text displayed with the hero image on the event"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "image_url",
          "title": "Image Url",
          "type": "`$ANY`",
          "short": "The URL of an image representing this page"
        },
        {
          "name": "is_admission_required",
          "title": "Is Admission Required",
          "type": "`$BOOLEAN`",
          "short": "Whether admission to the museum is required to attend this event"
        },
        {
          "name": "is_after_hours",
          "title": "Is After Hours",
          "type": "`$BOOLEAN`",
          "short": "Whether the event is to be held after the museum closes"
        },
        {
          "name": "is_free",
          "title": "Is Free",
          "type": "`$BOOLEAN`",
          "short": "Whether the event is free"
        },
        {
          "name": "is_member_exclusive",
          "title": "Is Member Exclusive",
          "type": "`$BOOLEAN`",
          "short": "Whether the event is exclusive to members of the museum"
        },
        {
          "name": "is_private",
          "title": "Is Private",
          "type": "`$BOOLEAN`",
          "short": "Whether the event is private"
        },
        {
          "name": "is_registration_required",
          "title": "Is Registration Required",
          "type": "`$BOOLEAN`",
          "short": "Whether registration is required to attend the event"
        },
        {
          "name": "is_sales_button_hidden",
          "title": "Is Sales Button Hidden",
          "type": "`$BOOLEAN`",
          "short": "Whether the buy tickets button is hidden on the website event page"
        },
        {
          "name": "is_sold_out",
          "title": "Is Sold Out",
          "type": "`$BOOLEAN`",
          "short": "Whether the event is sold out"
        },
        {
          "name": "is_ticketed",
          "title": "Is Ticketed",
          "type": "`$BOOLEAN`",
          "short": "Whether a ticket is required to attend the event"
        },
        {
          "name": "is_virtual_event",
          "title": "Is Virtual Event",
          "type": "`$BOOLEAN`",
          "short": "Whether the event is being held virtually"
        },
        {
          "name": "join_url",
          "title": "Join Url",
          "type": "`$ANY`",
          "short": "URL to the membership signup page via this event"
        },
        {
          "name": "layout_type",
          "title": "Layout Type",
          "type": "`$ANY`",
          "short": "Number indicating the type of layout this event page uses"
        },
        {
          "name": "list_description",
          "title": "List Description",
          "type": "`$ANY`",
          "short": "One-sentence description of the event displayed in listings"
        },
        {
          "name": "location",
          "title": "Location",
          "type": "`$ANY`",
          "short": "Where the event takes place"
        },
        {
          "name": "program_ids",
          "title": "Program Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers indicating the programs this event is a part of"
        },
        {
          "name": "program_titles",
          "title": "Program Titles",
          "type": "`$ANY`",
          "short": "Titles of the programs this event is a part of"
        },
        {
          "name": "rsvp_link",
          "title": "Rsvp Link",
          "type": "`$ANY`",
          "short": "The URL to the sales site for this event"
        },
        {
          "name": "search_tags",
          "title": "Search Tags",
          "type": "`$ANY`",
          "short": "Editor-specified list of tags to aid in internal search"
        },
        {
          "name": "short_description",
          "title": "Short Description",
          "type": "`$ANY`",
          "short": "Brief description of the event"
        },
        {
          "name": "slug",
          "title": "Slug",
          "type": "`$STRING`",
          "short": "A string used in the URL for this event"
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "start_date",
          "title": "Start Date",
          "type": "`$ANY`",
          "short": "The date the event begins"
        },
        {
          "name": "start_time",
          "title": "Start Time",
          "type": "`$ANY`",
          "short": "The time the event starts"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "survey_url",
          "title": "Survey Url",
          "type": "`$ANY`",
          "short": "URL to the survey associated with this event"
        },
        {
          "name": "ticketed_event_id",
          "title": "Ticketed Event Id",
          "type": "`$STRING`",
          "short": "Unique identifier of the event in the ticketing system this website event is tied to"
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "title_display",
          "title": "Title Display",
          "type": "`$ANY`",
          "short": "The name of this event formatted with HTML (optional)"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        },
        {
          "name": "virtual_event_passcode",
          "title": "Virtual Event Passcode",
          "type": "`$ANY`",
          "short": "Passcode to access the virtual event"
        },
        {
          "name": "virtual_event_url",
          "title": "Virtual Event Url",
          "type": "`$ANY`",
          "short": "URL to the virtual event"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "event",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/events",
              "segments": [
                {
                  "lit": "events"
                }
              ],
              "parts": [
                "events"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/events/{id}",
              "segments": [
                {
                  "lit": "events"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "events",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "button_caption",
          "title": "Button Caption",
          "type": "`$ANY`",
          "short": "Additional text below the ticket/registration button"
        },
        {
          "name": "button_text",
          "title": "Button Text",
          "type": "`$ANY`",
          "short": "The text used on the ticket/registration button"
        },
        {
          "name": "button_url",
          "title": "Button Url",
          "type": "`$ANY`",
          "short": "The URL to the sales site or an RSVP link for this event"
        },
        {
          "name": "description",
          "title": "Description",
          "type": "`$STRING`",
          "short": "Description of the event"
        },
        {
          "name": "end_at",
          "title": "End At",
          "type": "`$ANY`",
          "short": "The date the event occurrence ends"
        },
        {
          "name": "event_id",
          "title": "Event Id",
          "type": "`$STRING`",
          "short": "Identifier of the master event of which this is an occurrence"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "image_url",
          "title": "Image Url",
          "type": "`$ANY`",
          "short": "The URL of an image representing this page"
        },
        {
          "name": "is_private",
          "title": "Is Private",
          "type": "`$BOOLEAN`",
          "short": "Whether the event is private."
        },
        {
          "name": "is_sales_button_hidden",
          "title": "Is Sales Button Hidden",
          "type": "`$BOOLEAN`",
          "short": "Whether the buy tickets button is hidden on the website event page"
        },
        {
          "name": "is_ticketed",
          "title": "Is Ticketed",
          "type": "`$BOOLEAN`",
          "short": "Whether a ticket is required to attend the event"
        },
        {
          "name": "location",
          "title": "Location",
          "type": "`$ANY`",
          "short": "Where the event takes place"
        },
        {
          "name": "off_sale_at",
          "title": "Off Sale At",
          "type": "`$ANY`",
          "short": "Date and time the event goes off sale"
        },
        {
          "name": "on_sale_at",
          "title": "On Sale At",
          "type": "`$ANY`",
          "short": "Date and time the event goes on sale"
        },
        {
          "name": "short_description",
          "title": "Short Description",
          "type": "`$ANY`",
          "short": "Brief description of the event"
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "start_at",
          "title": "Start At",
          "type": "`$ANY`",
          "short": "The date the event occurrence begins"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "title_display",
          "title": "Title Display",
          "type": "`$ANY`",
          "short": "The name of this event formatted with HTML (optional)"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "event_occurrence",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/event-occurrences",
              "segments": [
                {
                  "lit": "event-occurrences"
                }
              ],
              "parts": [
                "event-occurrences"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/event-occurrences/{id}",
              "segments": [
                {
                  "lit": "event-occurrences"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "event-occurrences",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "is_affiliate_group",
          "title": "Is Affiliate Group",
          "type": "`$BOOLEAN`",
          "short": "Whether this program represents an affiliate group"
        },
        {
          "name": "is_event_host",
          "title": "Is Event Host",
          "type": "`$BOOLEAN`",
          "short": "Whether this program represents an event host"
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "event_program",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/event-programs",
              "segments": [
                {
                  "lit": "event-programs"
                }
              ],
              "parts": [
                "event-programs"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/event-programs/{id}",
              "segments": [
                {
                  "lit": "event-programs"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "event-programs",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Aic End At",
          "type": "`$ANY`",
          "short": "Date the exhibition closed at the Art Institute of Chicago"
        },
        {
          "name": "aic_start_at",
          "title": "Aic Start At",
          "type": "`$ANY`",
          "short": "Date the exhibition opened at the Art Institute of Chicago"
        },
        {
          "name": "alt_image_ids",
          "title": "Alt Image Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of all non-preferred images of this exhibition."
        },
        {
          "name": "api_link",
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "artist_ids",
          "title": "Artist Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of the artist agent records representing who was shown in the exhibition"
        },
        {
          "name": "artwork_ids",
          "title": "Artwork Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of the artworks that were part of the exhibition"
        },
        {
          "name": "artwork_titles",
          "title": "Artwork Titles",
          "type": "`$ANY`",
          "short": "Names of the artworks that were part of the exhibition"
        },
        {
          "name": "document_ids",
          "title": "Document Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of assets that serve as documentation for this exhibition"
        },
        {
          "name": "gallery_id",
          "title": "Gallery Id",
          "type": "`$STRING`",
          "short": "Unique identifier of the gallery that mainly housed the exhibition"
        },
        {
          "name": "gallery_title",
          "title": "Gallery Title",
          "type": "`$ANY`",
          "short": "The name of the gallery that mainly housed the exhibition"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "image_id",
          "title": "Image Id",
          "type": "`$STRING`",
          "short": "Unique identifier of the preferred image to use to represent this exhibition"
        },
        {
          "name": "image_url",
          "title": "Image Url",
          "type": "`$ANY`",
          "short": "URL to the hero image from the website"
        },
        {
          "name": "is_featured",
          "title": "Is Featured",
          "type": "`$BOOLEAN`",
          "short": "Is this exhibition currently featured on our website?"
        },
        {
          "name": "is_published",
          "title": "Is Published",
          "type": "`$BOOLEAN`",
          "short": "Is this exhibition currently published on our website?"
        },
        {
          "name": "position",
          "title": "Position",
          "type": "`$ANY`",
          "short": "Numering position represnting the order in which this exhibition is featured on the website"
        },
        {
          "name": "short_description",
          "title": "Short Description",
          "type": "`$ANY`",
          "short": "Brief explanation of what this exhibition is"
        },
        {
          "name": "site_ids",
          "title": "Site Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of the microsites this exhibition is a part of"
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "status",
          "title": "Status",
          "type": "`$ANY`",
          "short": "Whether the exhibition is open or closed"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        },
        {
          "name": "web_url",
          "title": "Web Url",
          "type": "`$ANY`",
          "short": "URL to this exhibition on our website"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "exhibition",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/exhibitions",
              "segments": [
                {
                  "lit": "exhibitions"
                }
              ],
              "parts": [
                "exhibitions"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/exhibitions/{id}",
              "segments": [
                {
                  "lit": "exhibitions"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "exhibitions",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "floor",
          "title": "Floor",
          "type": "`$ANY`",
          "short": "The level the gallery is on, e.g., 1, 2, 3, or LL"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "is_closed",
          "title": "Is Closed",
          "type": "`$BOOLEAN`",
          "short": "Whether the gallery is currently closed"
        },
        {
          "name": "latitude",
          "title": "Latitude",
          "type": "`$NUMBER`",
          "short": "Latitude coordinate of the center of the room"
        },
        {
          "name": "latlon",
          "title": "Latlon",
          "type": "`$ANY`",
          "short": "Latitude and longitude coordinates of the center of the room"
        },
        {
          "name": "longitude",
          "title": "Longitude",
          "type": "`$NUMBER`",
          "short": "Longitude coordinate of the center of the room"
        },
        {
          "name": "number",
          "title": "Number",
          "type": "`$ANY`",
          "short": "The gallery's room number."
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "tgn_id",
          "title": "Tgn Id",
          "type": "`$STRING`",
          "short": "Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN)"
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "gallery",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/galleries",
              "segments": [
                {
                  "lit": "galleries"
                }
              ],
              "parts": [
                "galleries"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/galleries/{id}",
              "segments": [
                {
                  "lit": "galleries"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "galleries",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "copy",
          "title": "Copy",
          "type": "`$ANY`",
          "short": "The text of the page"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "search_tags",
          "title": "Search Tags",
          "type": "`$ANY`",
          "short": "Editor-specified list of tags to aid in internal search"
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        },
        {
          "name": "web_url",
          "title": "Web Url",
          "type": "`$ANY`",
          "short": "The URL to this page on our website"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "generic_page",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/generic-pages",
              "segments": [
                {
                  "lit": "generic-pages"
                }
              ],
              "parts": [
                "generic-pages"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/generic-pages/{id}",
              "segments": [
                {
                  "lit": "generic-pages"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "generic-pages",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "copy",
          "title": "Copy",
          "type": "`$ANY`",
          "short": "The text of the highlight description"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "highlight",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/highlights",
              "segments": [
                {
                  "lit": "highlights"
                }
              ],
              "parts": [
                "highlights"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/highlights/{id}",
              "segments": [
                {
                  "lit": "highlights"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "highlights",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Additional Text",
          "type": "`$ANY`",
          "short": "Additional information about the hours"
        },
        {
          "name": "api_link",
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "friday_is_closed",
          "title": "Friday Is Closed",
          "type": "`$ANY`",
          "short": "Whether the museum is closed on Fridays"
        },
        {
          "name": "friday_member_close",
          "title": "Friday Member Close",
          "type": "`$ANY`",
          "short": "The time member hours ends on Fridays"
        },
        {
          "name": "friday_member_open",
          "title": "Friday Member Open",
          "type": "`$ANY`",
          "short": "The time member hours starts on Fridays"
        },
        {
          "name": "friday_public_close",
          "title": "Friday Public Close",
          "type": "`$ANY`",
          "short": "The time public hours ends on Fridays"
        },
        {
          "name": "friday_public_open",
          "title": "Friday Public Open",
          "type": "`$ANY`",
          "short": "The time public hours starts on Fridays"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "monday_is_closed",
          "title": "Monday Is Closed",
          "type": "`$ANY`",
          "short": "Whether the museum is closed on Mondays"
        },
        {
          "name": "monday_member_close",
          "title": "Monday Member Close",
          "type": "`$ANY`",
          "short": "The time member hours ends on Mondays"
        },
        {
          "name": "monday_member_open",
          "title": "Monday Member Open",
          "type": "`$ANY`",
          "short": "The time member hours starts on Mondays"
        },
        {
          "name": "monday_public_close",
          "title": "Monday Public Close",
          "type": "`$ANY`",
          "short": "The time public hours ends on Mondays"
        },
        {
          "name": "monday_public_open",
          "title": "Monday Public Open",
          "type": "`$ANY`",
          "short": "The time public hours starts on Mondays"
        },
        {
          "name": "saturday_is_closed",
          "title": "Saturday Is Closed",
          "type": "`$ANY`",
          "short": "Whether the museum is closed on Saturdays"
        },
        {
          "name": "saturday_member_close",
          "title": "Saturday Member Close",
          "type": "`$ANY`",
          "short": "The time member hours ends on Saturdays"
        },
        {
          "name": "saturday_member_open",
          "title": "Saturday Member Open",
          "type": "`$ANY`",
          "short": "The time member hours starts on Saturdays"
        },
        {
          "name": "saturday_public_close",
          "title": "Saturday Public Close",
          "type": "`$ANY`",
          "short": "The time public hours ends on Saturdays"
        },
        {
          "name": "saturday_public_open",
          "title": "Saturday Public Open",
          "type": "`$ANY`",
          "short": "The time public hours starts on Saturdays"
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "summary",
          "title": "Summary",
          "type": "`$ANY`",
          "short": "Readable summary of the hours"
        },
        {
          "name": "sunday_is_closed",
          "title": "Sunday Is Closed",
          "type": "`$ANY`",
          "short": "Whether the museum is closed on Sundays"
        },
        {
          "name": "sunday_member_close",
          "title": "Sunday Member Close",
          "type": "`$ANY`",
          "short": "The time member hours ends on Sundays"
        },
        {
          "name": "sunday_member_open",
          "title": "Sunday Member Open",
          "type": "`$ANY`",
          "short": "The time member hours starts on Sundays"
        },
        {
          "name": "sunday_public_close",
          "title": "Sunday Public Close",
          "type": "`$ANY`",
          "short": "The time public hours ends on Sundays"
        },
        {
          "name": "sunday_public_open",
          "title": "Sunday Public Open",
          "type": "`$ANY`",
          "short": "The time public hours starts on Sundays"
        },
        {
          "name": "thursday_is_closed",
          "title": "Thursday Is Closed",
          "type": "`$ANY`",
          "short": "Whether the museum is closed on Thursdays"
        },
        {
          "name": "thursday_member_close",
          "title": "Thursday Member Close",
          "type": "`$ANY`",
          "short": "The time member hours ends on Thursdays"
        },
        {
          "name": "thursday_member_open",
          "title": "Thursday Member Open",
          "type": "`$ANY`",
          "short": "The time member hours starts on Thursdays"
        },
        {
          "name": "thursday_public_close",
          "title": "Thursday Public Close",
          "type": "`$ANY`",
          "short": "The time public hours ends on Thursdays"
        },
        {
          "name": "thursday_public_open",
          "title": "Thursday Public Open",
          "type": "`$ANY`",
          "short": "The time public hours starts on Thursdays"
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "tuesday_is_closed",
          "title": "Tuesday Is Closed",
          "type": "`$ANY`",
          "short": "Whether the museum is closed on Tuesdays"
        },
        {
          "name": "tuesday_member_close",
          "title": "Tuesday Member Close",
          "type": "`$ANY`",
          "short": "The time member hours ends on Tuesdays"
        },
        {
          "name": "tuesday_member_open",
          "title": "Tuesday Member Open",
          "type": "`$ANY`",
          "short": "The time member hours starts on Tuesdays"
        },
        {
          "name": "tuesday_public_close",
          "title": "Tuesday Public Close",
          "type": "`$ANY`",
          "short": "The time public hours ends on Tuesdays"
        },
        {
          "name": "tuesday_public_open",
          "title": "Tuesday Public Open",
          "type": "`$ANY`",
          "short": "The time public hours starts on Tuesdays"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        },
        {
          "name": "wednesday_is_closed",
          "title": "Wednesday Is Closed",
          "type": "`$ANY`",
          "short": "Whether the museum is closed on Wednesdays"
        },
        {
          "name": "wednesday_member_close",
          "title": "Wednesday Member Close",
          "type": "`$ANY`",
          "short": "The time member hours ends on Wednesdays"
        },
        {
          "name": "wednesday_member_open",
          "title": "Wednesday Member Open",
          "type": "`$ANY`",
          "short": "The time member hours starts on Wednesdays"
        },
        {
          "name": "wednesday_public_close",
          "title": "Wednesday Public Close",
          "type": "`$ANY`",
          "short": "The time public hours ends on Wednesdays"
        },
        {
          "name": "wednesday_public_open",
          "title": "Wednesday Public Open",
          "type": "`$ANY`",
          "short": "The time public hours starts on Wednesdays"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "hour",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/hours",
              "segments": [
                {
                  "lit": "hours"
                }
              ],
              "parts": [
                "hours"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/hours/{id}",
              "segments": [
                {
                  "lit": "hours"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "hours",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Ahash",
          "type": "`$ANY`",
          "short": "Image hash generated using ahash algorithm with 64 boolean subfields"
        },
        {
          "name": "alt_text",
          "title": "Alt Text",
          "type": "`$ANY`",
          "short": "Alternative text for the asset to describe it to people with low or no vision"
        },
        {
          "name": "api_link",
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "artwork_ids",
          "title": "Artwork Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of the artworks associated with this asset"
        },
        {
          "name": "artwork_titles",
          "title": "Artwork Titles",
          "type": "`$ANY`",
          "short": "Names of the artworks associated with this asset"
        },
        {
          "name": "color",
          "title": "Color",
          "type": "`$ANY`",
          "short": "Dominant color of this image in HSL"
        },
        {
          "name": "colorfulness",
          "title": "Colorfulness",
          "type": "`$ANY`",
          "short": "Unbounded positive float representing an abstract measure of colorfulness."
        },
        {
          "name": "content",
          "title": "Content",
          "type": "`$ANY`",
          "short": "Text of or URL to the contents of this asset"
        },
        {
          "name": "content_e_tag",
          "title": "Content E Tag",
          "type": "`$ANY`",
          "short": "Arbitrary unique identifier that changes when the binary file gets updated"
        },
        {
          "name": "credit_line",
          "title": "Credit Line",
          "type": "`$ANY`",
          "short": "Asset-specific copyright information"
        },
        {
          "name": "fingerprint",
          "title": "Fingerprint",
          "type": "`$ANY`",
          "short": "Image hashes: aHash, dHash, pHash, wHash"
        },
        {
          "name": "height",
          "title": "Height",
          "type": "`$NUMBER`",
          "short": "Native height of the image"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "iiif_url",
          "title": "Iiif Url",
          "type": "`$ANY`",
          "short": "IIIF URL of this image"
        },
        {
          "name": "is_educational_resource",
          "title": "Is Educational Resource",
          "type": "`$BOOLEAN`",
          "short": "Whether this resource is considered to be educational"
        },
        {
          "name": "is_multimedia_resource",
          "title": "Is Multimedia Resource",
          "type": "`$BOOLEAN`",
          "short": "Whether this resource is considered to be multimedia"
        },
        {
          "name": "is_teacher_resource",
          "title": "Is Teacher Resource",
          "type": "`$BOOLEAN`",
          "short": "Whether this resource is considered to be educational"
        },
        {
          "name": "lake_guid",
          "title": "Lake Guid",
          "type": "`$ANY`",
          "short": "Unique UUID of this resource in LAKE, our DAMS."
        },
        {
          "name": "lqip",
          "title": "Lqip",
          "type": "`$ANY`",
          "short": "Low-quality image placeholder (LQIP)."
        },
        {
          "name": "phash",
          "title": "Phash",
          "type": "`$ANY`",
          "short": "Image hash generated using phash algorithm with 64 boolean subfields"
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "type",
          "title": "Type",
          "type": "`$ANY`",
          "short": "Type always takes one of the following values: image, sound, text, video"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        },
        {
          "name": "width",
          "title": "Width",
          "type": "`$NUMBER`",
          "short": "Native width of the image"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "image",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/images",
              "segments": [
                {
                  "lit": "images"
                }
              ],
              "parts": [
                "images"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/images/{id}",
              "segments": [
                {
                  "lit": "images"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "images",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "copy",
          "title": "Copy",
          "type": "`$ANY`",
          "short": "The text of the page"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "search_tags",
          "title": "Search Tags",
          "type": "`$ANY`",
          "short": "Editor-specified list of tags to aid in internal search"
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        },
        {
          "name": "web_url",
          "title": "Web Url",
          "type": "`$ANY`",
          "short": "The URL to this page on our website"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "landing_page",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/landing-pages",
              "segments": [
                {
                  "lit": "landing-pages"
                }
              ],
              "parts": [
                "landing-pages"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/landing-pages/{id}",
              "segments": [
                {
                  "lit": "landing-pages"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "landing-pages",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "latitude",
          "title": "Latitude",
          "type": "`$NUMBER`",
          "short": "Latitude coordinate of the center of the room"
        },
        {
          "name": "longitude",
          "title": "Longitude",
          "type": "`$NUMBER`",
          "short": "Longitude coordinate of the center of the room"
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "tgn_id",
          "title": "Tgn Id",
          "type": "`$STRING`",
          "short": "Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN)"
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "place",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/places",
              "segments": [
                {
                  "lit": "places"
                }
              ],
              "parts": [
                "places"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/places/{id}",
              "segments": [
                {
                  "lit": "places"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "places",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "copy",
          "title": "Copy",
          "type": "`$ANY`",
          "short": "The text of the page"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        },
        {
          "name": "web_url",
          "title": "Web Url",
          "type": "`$ANY`",
          "short": "The URL to this page on our website"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "press_release",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/press-releases",
              "segments": [
                {
                  "lit": "press-releases"
                }
              ],
              "parts": [
                "press-releases"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/press-releases/{id}",
              "segments": [
                {
                  "lit": "press-releases"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "press-releases",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "copy",
          "title": "Copy",
          "type": "`$ANY`",
          "short": "The text of the page"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        },
        {
          "name": "web_url",
          "title": "Web Url",
          "type": "`$ANY`",
          "short": "The URL to this page on our website"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "printed_publication",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/printed-publications",
              "segments": [
                {
                  "lit": "printed-publications"
                }
              ],
              "parts": [
                "printed-publications"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/printed-publications/{id}",
              "segments": [
                {
                  "lit": "printed-publications"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "printed-publications",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "artist_ids",
          "title": "Artist Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of the artists associated with this product"
        },
        {
          "name": "artwork_ids",
          "title": "Artwork Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of the artworks associated with this product"
        },
        {
          "name": "description",
          "title": "Description",
          "type": "`$STRING`",
          "short": "Explanation of what this product is"
        },
        {
          "name": "exhibition_ids",
          "title": "Exhibition Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of the exhibitions associated with this product"
        },
        {
          "name": "external_sku",
          "title": "External Sku",
          "type": "`$ANY`",
          "short": "Numeric product identification code of a machine-readable barcode, when the customer sku differs from our internal one"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "image_url",
          "title": "Image Url",
          "type": "`$ANY`",
          "short": "URL of an image for this product"
        },
        {
          "name": "max_compare_at_price",
          "title": "Max Compare At Price",
          "type": "`$ANY`",
          "short": "Number indicating how much the most expensive variant of a product cost before a sale"
        },
        {
          "name": "max_current_price",
          "title": "Max Current Price",
          "type": "`$ANY`",
          "short": "Number indicating how much the most expensive variant of a product costs right now"
        },
        {
          "name": "min_compare_at_price",
          "title": "Min Compare At Price",
          "type": "`$ANY`",
          "short": "Number indicating how much the least expensive variant of a product cost before a sale"
        },
        {
          "name": "min_current_price",
          "title": "Min Current Price",
          "type": "`$ANY`",
          "short": "Number indicating how much the least expensive variant of a product costs right now"
        },
        {
          "name": "price_display",
          "title": "Price Display",
          "type": "`$ANY`",
          "short": "Explanation of what this product is"
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        },
        {
          "name": "web_url",
          "title": "Web Url",
          "type": "`$ANY`",
          "short": "URL of this product in the shop"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "product",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/products",
              "segments": [
                {
                  "lit": "products"
                }
              ],
              "parts": [
                "products"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/products/{id}",
              "segments": [
                {
                  "lit": "products"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "products",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "section_ids",
          "title": "Section Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of the sections of this publication"
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        },
        {
          "name": "web_url",
          "title": "Web Url",
          "type": "`$ANY`",
          "short": "URL to the publication"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "publication",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/publications",
              "segments": [
                {
                  "lit": "publications"
                }
              ],
              "parts": [
                "publications"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/publications/{id}",
              "segments": [
                {
                  "lit": "publications"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "publications",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Api Id",
          "type": "`$STRING`",
          "short": "API unique identifier"
        },
        {
          "name": "api_link",
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "URL to this recource in the API"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "Name of the model the resource represents"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier within the search index"
        },
        {
          "name": "is_boosted",
          "title": "Is Boosted",
          "type": "`$BOOLEAN`",
          "short": "Whether this record has been flagged to be boosted"
        },
        {
          "name": "score",
          "title": "Score",
          "type": "`$NUMBER`",
          "short": "Search index ranking of the result"
        },
        {
          "name": "thumbnail",
          "title": "Thumbnail",
          "type": "`$ANY`",
          "short": "Metadata on the image representing this record"
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date this record was last updated in the API"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "search",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/agents/search",
              "segments": [
                {
                  "lit": "agents"
                },
                {
                  "lit": "search"
                }
              ],
              "parts": [
                "agents",
                "search"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "facet",
                    "orig": "facets",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
              },
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/articles/search",
              "segments": [
                {
                  "lit": "articles"
                },
                {
                  "lit": "search"
                }
              ],
              "parts": [
                "articles",
                "search"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "facet",
                    "orig": "facets",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
              },
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/artworks/search",
              "segments": [
                {
                  "lit": "artworks"
                },
                {
                  "lit": "search"
                }
              ],
              "parts": [
                "artworks",
                "search"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "facet",
                    "orig": "facets",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
              },
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/category-terms/search",
              "segments": [
                {
                  "lit": "category-terms"
                },
                {
                  "lit": "search"
                }
              ],
              "parts": [
                "category-terms",
                "search"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "facet",
                    "orig": "facets",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
              },
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/digital-publication-articles/search",
              "segments": [
                {
                  "lit": "digital-publication-articles"
                },
                {
                  "lit": "search"
                }
              ],
              "parts": [
                "digital-publication-articles",
                "search"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "facet",
                    "orig": "facets",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
              },
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/digital-publications/search",
              "segments": [
                {
                  "lit": "digital-publications"
                },
                {
                  "lit": "search"
                }
              ],
              "parts": [
                "digital-publications",
                "search"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "facet",
                    "orig": "facets",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
              },
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/educator-resources/search",
              "segments": [
                {
                  "lit": "educator-resources"
                },
                {
                  "lit": "search"
                }
              ],
              "parts": [
                "educator-resources",
                "search"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "facet",
                    "orig": "facets",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
              },
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/event-occurrences/search",
              "segments": [
                {
                  "lit": "event-occurrences"
                },
                {
                  "lit": "search"
                }
              ],
              "parts": [
                "event-occurrences",
                "search"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "facet",
                    "orig": "facets",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
              },
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/event-programs/search",
              "segments": [
                {
                  "lit": "event-programs"
                },
                {
                  "lit": "search"
                }
              ],
              "parts": [
                "event-programs",
                "search"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "facet",
                    "orig": "facets",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
              },
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/events/search",
              "segments": [
                {
                  "lit": "events"
                },
                {
                  "lit": "search"
                }
              ],
              "parts": [
                "events",
                "search"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "facet",
                    "orig": "facets",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
              },
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/exhibitions/search",
              "segments": [
                {
                  "lit": "exhibitions"
                },
                {
                  "lit": "search"
                }
              ],
              "parts": [
                "exhibitions",
                "search"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "facet",
                    "orig": "facets",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
              },
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/galleries/search",
              "segments": [
                {
                  "lit": "galleries"
                },
                {
                  "lit": "search"
                }
              ],
              "parts": [
                "galleries",
                "search"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "facet",
                    "orig": "facets",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
              },
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/generic-pages/search",
              "segments": [
                {
                  "lit": "generic-pages"
                },
                {
                  "lit": "search"
                }
              ],
              "parts": [
                "generic-pages",
                "search"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "facet",
                    "orig": "facets",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
              },
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/highlights/search",
              "segments": [
                {
                  "lit": "highlights"
                },
                {
                  "lit": "search"
                }
              ],
              "parts": [
                "highlights",
                "search"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "facet",
                    "orig": "facets",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
              },
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/hours/search",
              "segments": [
                {
                  "lit": "hours"
                },
                {
                  "lit": "search"
                }
              ],
              "parts": [
                "hours",
                "search"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "facet",
                    "orig": "facets",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
              },
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/images/search",
              "segments": [
                {
                  "lit": "images"
                },
                {
                  "lit": "search"
                }
              ],
              "parts": [
                "images",
                "search"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "facet",
                    "orig": "facets",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
              },
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/landing-pages/search",
              "segments": [
                {
                  "lit": "landing-pages"
                },
                {
                  "lit": "search"
                }
              ],
              "parts": [
                "landing-pages",
                "search"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "facet",
                    "orig": "facets",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
              },
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/mobile-sounds/search",
              "segments": [
                {
                  "lit": "mobile-sounds"
                },
                {
                  "lit": "search"
                }
              ],
              "parts": [
                "mobile-sounds",
                "search"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "facet",
                    "orig": "facets",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
              },
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/places/search",
              "segments": [
                {
                  "lit": "places"
                },
                {
                  "lit": "search"
                }
              ],
              "parts": [
                "places",
                "search"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "facet",
                    "orig": "facets",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
              },
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/press-releases/search",
              "segments": [
                {
                  "lit": "press-releases"
                },
                {
                  "lit": "search"
                }
              ],
              "parts": [
                "press-releases",
                "search"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "facet",
                    "orig": "facets",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
              },
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/printed-publications/search",
              "segments": [
                {
                  "lit": "printed-publications"
                },
                {
                  "lit": "search"
                }
              ],
              "parts": [
                "printed-publications",
                "search"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "facet",
                    "orig": "facets",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
              },
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/products/search",
              "segments": [
                {
                  "lit": "products"
                },
                {
                  "lit": "search"
                }
              ],
              "parts": [
                "products",
                "search"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "facet",
                    "orig": "facets",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
              },
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/publications/search",
              "segments": [
                {
                  "lit": "publications"
                },
                {
                  "lit": "search"
                }
              ],
              "parts": [
                "publications",
                "search"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "facet",
                    "orig": "facets",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
              },
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/search",
              "segments": [
                {
                  "lit": "search"
                }
              ],
              "parts": [
                "search"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "facet",
                    "orig": "facets",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
              },
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/sections/search",
              "segments": [
                {
                  "lit": "sections"
                },
                {
                  "lit": "search"
                }
              ],
              "parts": [
                "sections",
                "search"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "facet",
                    "orig": "facets",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
              },
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/sites/search",
              "segments": [
                {
                  "lit": "sites"
                },
                {
                  "lit": "search"
                }
              ],
              "parts": [
                "sites",
                "search"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "facet",
                    "orig": "facets",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
              },
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/sounds/search",
              "segments": [
                {
                  "lit": "sounds"
                },
                {
                  "lit": "search"
                }
              ],
              "parts": [
                "sounds",
                "search"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "facet",
                    "orig": "facets",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
              },
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/static-pages/search",
              "segments": [
                {
                  "lit": "static-pages"
                },
                {
                  "lit": "search"
                }
              ],
              "parts": [
                "static-pages",
                "search"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "facet",
                    "orig": "facets",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
              },
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/texts/search",
              "segments": [
                {
                  "lit": "texts"
                },
                {
                  "lit": "search"
                }
              ],
              "parts": [
                "texts",
                "search"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "facet",
                    "orig": "facets",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
              },
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/tours/search",
              "segments": [
                {
                  "lit": "tours"
                },
                {
                  "lit": "search"
                }
              ],
              "parts": [
                "tours",
                "search"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "facet",
                    "orig": "facets",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
              },
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/videos/search",
              "segments": [
                {
                  "lit": "videos"
                },
                {
                  "lit": "search"
                }
              ],
              "parts": [
                "videos",
                "search"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "facet",
                    "orig": "facets",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "from",
                    "orig": "from",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "query",
                    "orig": "query",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "size",
                    "orig": "size",
                    "type": "`$STRING`",
                    "kind": "query"
                  },
                  {
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`",
                    "kind": "query"
                  }
                ]
              },
              "select": {}
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
          "title": "Accession",
          "type": "`$ANY`",
          "short": "An accession number parsed from the title or tombstone"
        },
        {
          "name": "api_link",
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "artwork_id",
          "title": "Artwork Id",
          "type": "`$STRING`",
          "short": "Unique identifier of the artwork with which this section is associated"
        },
        {
          "name": "content",
          "title": "Content",
          "type": "`$ANY`",
          "short": "Content of this section in plaintext"
        },
        {
          "name": "generic_page_id",
          "title": "Generic Page Id",
          "type": "`$STRING`",
          "short": "Unique identifier of the page on the website that represents the publication this section belongs to"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "publication_id",
          "title": "Publication Id",
          "type": "`$STRING`",
          "short": "Unique identifier of the publication this section belongs to"
        },
        {
          "name": "publication_title",
          "title": "Publication Title",
          "type": "`$ANY`",
          "short": "Name of the publication this section belongs to"
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        },
        {
          "name": "web_url",
          "title": "Web Url",
          "type": "`$ANY`",
          "short": "URL to the section"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "section",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/sections",
              "segments": [
                {
                  "lit": "sections"
                }
              ],
              "parts": [
                "sections"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/sections/{id}",
              "segments": [
                {
                  "lit": "sections"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "sections",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "artwork_ids",
          "title": "Artwork Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of the artworks this site is associated with"
        },
        {
          "name": "artwork_titles",
          "title": "Artwork Titles",
          "type": "`$ANY`",
          "short": "Names of the artworks this site is associated with"
        },
        {
          "name": "description",
          "title": "Description",
          "type": "`$STRING`",
          "short": "Explanation of what this site is"
        },
        {
          "name": "exhibition_ids",
          "title": "Exhibition Ids",
          "type": "`$ANY`",
          "short": "Unique identifier of the exhibitions this site is associated with"
        },
        {
          "name": "exhibition_titles",
          "title": "Exhibition Titles",
          "type": "`$ANY`",
          "short": "Names of the exhibitions this site is associated with"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        },
        {
          "name": "web_url",
          "title": "Web Url",
          "type": "`$ANY`",
          "short": "URL to this site"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "site",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/sites",
              "segments": [
                {
                  "lit": "sites"
                }
              ],
              "parts": [
                "sites"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/sites/{id}",
              "segments": [
                {
                  "lit": "sites"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "sites",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Alt Text",
          "type": "`$ANY`",
          "short": "Alternative text for the asset to describe it to people with low or no vision"
        },
        {
          "name": "api_link",
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "artwork_ids",
          "title": "Artwork Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of the artworks associated with this asset"
        },
        {
          "name": "artwork_titles",
          "title": "Artwork Titles",
          "type": "`$ANY`",
          "short": "Names of the artworks associated with this asset"
        },
        {
          "name": "content",
          "title": "Content",
          "type": "`$ANY`",
          "short": "Text of or URL to the contents of this asset"
        },
        {
          "name": "content_e_tag",
          "title": "Content E Tag",
          "type": "`$ANY`",
          "short": "Arbitrary unique identifier that changes when the binary file gets updated"
        },
        {
          "name": "credit_line",
          "title": "Credit Line",
          "type": "`$ANY`",
          "short": "Asset-specific copyright information"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "is_educational_resource",
          "title": "Is Educational Resource",
          "type": "`$BOOLEAN`",
          "short": "Whether this resource is considered to be educational"
        },
        {
          "name": "is_multimedia_resource",
          "title": "Is Multimedia Resource",
          "type": "`$BOOLEAN`",
          "short": "Whether this resource is considered to be multimedia"
        },
        {
          "name": "is_teacher_resource",
          "title": "Is Teacher Resource",
          "type": "`$BOOLEAN`",
          "short": "Whether this resource is considered to be educational"
        },
        {
          "name": "lake_guid",
          "title": "Lake Guid",
          "type": "`$ANY`",
          "short": "Unique UUID of this resource in LAKE, our DAMS."
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "Name of this mobile audio file – derived from the artwork and tour titles"
        },
        {
          "name": "transcript",
          "title": "Transcript",
          "type": "`$ANY`",
          "short": "Text transcription of the audio file"
        },
        {
          "name": "type",
          "title": "Type",
          "type": "`$ANY`",
          "short": "Type always takes one of the following values: image, sound, text, video"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        },
        {
          "name": "web_url",
          "title": "Web Url",
          "type": "`$ANY`",
          "short": "URL to the audio file"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "sound",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/mobile-sounds",
              "segments": [
                {
                  "lit": "mobile-sounds"
                }
              ],
              "parts": [
                "mobile-sounds"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/sounds",
              "segments": [
                {
                  "lit": "sounds"
                }
              ],
              "parts": [
                "sounds"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/mobile-sounds/{id}",
              "segments": [
                {
                  "lit": "mobile-sounds"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "mobile-sounds",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
              }
            },
            {
              "kind": "http",
              "method": "GET",
              "orig": "/sounds/{id}",
              "segments": [
                {
                  "lit": "sounds"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "sounds",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        },
        {
          "name": "web_url",
          "title": "Web Url",
          "type": "`$ANY`",
          "short": "The URL to this page on our website"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "static_page",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/static-pages",
              "segments": [
                {
                  "lit": "static-pages"
                }
              ],
              "parts": [
                "static-pages"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/static-pages/{id}",
              "segments": [
                {
                  "lit": "static-pages"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "static-pages",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Alt Text",
          "type": "`$ANY`",
          "short": "Alternative text for the asset to describe it to people with low or no vision"
        },
        {
          "name": "api_link",
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "artwork_ids",
          "title": "Artwork Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of the artworks associated with this asset"
        },
        {
          "name": "artwork_titles",
          "title": "Artwork Titles",
          "type": "`$ANY`",
          "short": "Names of the artworks associated with this asset"
        },
        {
          "name": "content",
          "title": "Content",
          "type": "`$ANY`",
          "short": "Text of or URL to the contents of this asset"
        },
        {
          "name": "content_e_tag",
          "title": "Content E Tag",
          "type": "`$ANY`",
          "short": "Arbitrary unique identifier that changes when the binary file gets updated"
        },
        {
          "name": "credit_line",
          "title": "Credit Line",
          "type": "`$ANY`",
          "short": "Asset-specific copyright information"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "is_educational_resource",
          "title": "Is Educational Resource",
          "type": "`$BOOLEAN`",
          "short": "Whether this resource is considered to be educational"
        },
        {
          "name": "is_multimedia_resource",
          "title": "Is Multimedia Resource",
          "type": "`$BOOLEAN`",
          "short": "Whether this resource is considered to be multimedia"
        },
        {
          "name": "is_teacher_resource",
          "title": "Is Teacher Resource",
          "type": "`$BOOLEAN`",
          "short": "Whether this resource is considered to be educational"
        },
        {
          "name": "lake_guid",
          "title": "Lake Guid",
          "type": "`$ANY`",
          "short": "Unique UUID of this resource in LAKE, our DAMS."
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "type",
          "title": "Type",
          "type": "`$ANY`",
          "short": "Type always takes one of the following values: image, sound, text, video"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "text",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/texts",
              "segments": [
                {
                  "lit": "texts"
                }
              ],
              "parts": [
                "texts"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/texts/{id}",
              "segments": [
                {
                  "lit": "texts"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "texts",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "artist_titles",
          "title": "Artist Titles",
          "type": "`$ANY`",
          "short": "Names of the artists of the artworks featured in this tour's tour stops"
        },
        {
          "name": "artwork_titles",
          "title": "Artwork Titles",
          "type": "`$ANY`",
          "short": "Names of the artworks featured in this tour's tour stops"
        },
        {
          "name": "description",
          "title": "Description",
          "type": "`$STRING`",
          "short": "Explanation of what the tour is"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "image",
          "title": "Image",
          "type": "`$ANY`",
          "short": "The main image for the tour"
        },
        {
          "name": "intro",
          "title": "Intro",
          "type": "`$ANY`",
          "short": "Text introducing the tour"
        },
        {
          "name": "intro_link",
          "title": "Intro Link",
          "type": "`$ANY`",
          "short": "Link to the audio file of the introduction"
        },
        {
          "name": "intro_transcript",
          "title": "Intro Transcript",
          "type": "`$ANY`",
          "short": "Transcript of the introduction audio to the tour"
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        },
        {
          "name": "weight",
          "title": "Weight",
          "type": "`$NUMBER`",
          "short": "Number representing this tour's sort order"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "tour",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/tours",
              "segments": [
                {
                  "lit": "tours"
                }
              ],
              "parts": [
                "tours"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/tours/{id}",
              "segments": [
                {
                  "lit": "tours"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "tours",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
          "title": "Alt Text",
          "type": "`$ANY`",
          "short": "Alternative text for the asset to describe it to people with low or no vision"
        },
        {
          "name": "api_link",
          "title": "Api Link",
          "type": "`$ANY`",
          "short": "REST API link for this resource"
        },
        {
          "name": "api_model",
          "title": "Api Model",
          "type": "`$ANY`",
          "short": "REST API resource type or endpoint"
        },
        {
          "name": "artwork_ids",
          "title": "Artwork Ids",
          "type": "`$ANY`",
          "short": "Unique identifiers of the artworks associated with this asset"
        },
        {
          "name": "artwork_titles",
          "title": "Artwork Titles",
          "type": "`$ANY`",
          "short": "Names of the artworks associated with this asset"
        },
        {
          "name": "content",
          "title": "Content",
          "type": "`$ANY`",
          "short": "Text of or URL to the contents of this asset"
        },
        {
          "name": "content_e_tag",
          "title": "Content E Tag",
          "type": "`$ANY`",
          "short": "Arbitrary unique identifier that changes when the binary file gets updated"
        },
        {
          "name": "credit_line",
          "title": "Credit Line",
          "type": "`$ANY`",
          "short": "Asset-specific copyright information"
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`",
          "short": "Unique identifier of this resource."
        },
        {
          "name": "is_educational_resource",
          "title": "Is Educational Resource",
          "type": "`$BOOLEAN`",
          "short": "Whether this resource is considered to be educational"
        },
        {
          "name": "is_multimedia_resource",
          "title": "Is Multimedia Resource",
          "type": "`$BOOLEAN`",
          "short": "Whether this resource is considered to be multimedia"
        },
        {
          "name": "is_teacher_resource",
          "title": "Is Teacher Resource",
          "type": "`$BOOLEAN`",
          "short": "Whether this resource is considered to be educational"
        },
        {
          "name": "lake_guid",
          "title": "Lake Guid",
          "type": "`$ANY`",
          "short": "Unique UUID of this resource in LAKE, our DAMS."
        },
        {
          "name": "source_updated_at",
          "title": "Source Updated At",
          "type": "`$ANY`",
          "short": "Date and time the resource was updated in the source system"
        },
        {
          "name": "suggest_autocomplete_all",
          "title": "Suggest Autocomplete All",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autosuggest` endpoint."
        },
        {
          "name": "suggest_autocomplete_boosted",
          "title": "Suggest Autocomplete Boosted",
          "type": "`$ANY`",
          "short": "Internal field to power the `/autocomplete` endpoint."
        },
        {
          "name": "timestamp",
          "title": "Timestamp",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator search index"
        },
        {
          "name": "title",
          "title": "Title",
          "type": "`$STRING`",
          "short": "The name of this resource"
        },
        {
          "name": "type",
          "title": "Type",
          "type": "`$ANY`",
          "short": "Type always takes one of the following values: image, sound, text, video"
        },
        {
          "name": "updated_at",
          "title": "Updated At",
          "type": "`$ANY`",
          "short": "Date and time the record was updated in the aggregator database"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "video",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/videos",
              "segments": [
                {
                  "lit": "videos"
                }
              ],
              "parts": [
                "videos"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {},
              "select": {}
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/videos/{id}",
              "segments": [
                {
                  "lit": "videos"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "videos",
                "{id}"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "id"
                ]
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
  config,
  FEATURE_PLUGINS,
}

