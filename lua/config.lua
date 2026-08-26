-- ArtInstituteOfChicago SDK configuration

-- Build a fresh, fully materialised config table. Every call rebuilds the
-- whole structure, so prefer require("config_shared") unless you need a
-- private copy you intend to mutate.
local function make_config()
  return {
    main = {
      name = "ArtInstituteOfChicago",
      slug = "art-institute-of-chicago",
      version = "0.0.1",
      target = "lua",
    },
    feature = {
      ["test"] = {
        ["options"] = {
          ["active"] = false,
        },
        ["transport"] = "base",
      },
    },
    options = {
      base = "https://api.artic.edu/api/v1",
      headers = {
        ["content-type"] = "application/json",
      },
      entity = {
        ["agent"] = {},
        ["agent_role"] = {},
        ["agent_type"] = {},
        ["article"] = {},
        ["artwork"] = {},
        ["artwork_date_qualifier"] = {},
        ["artwork_place_qualifier"] = {},
        ["artwork_type"] = {},
        ["category_term"] = {},
        ["digital_publication"] = {},
        ["digital_publication_article"] = {},
        ["educator_resource"] = {},
        ["event"] = {},
        ["event_occurrence"] = {},
        ["event_program"] = {},
        ["exhibition"] = {},
        ["gallery"] = {},
        ["generic_page"] = {},
        ["highlight"] = {},
        ["hour"] = {},
        ["image"] = {},
        ["landing_page"] = {},
        ["place"] = {},
        ["press_release"] = {},
        ["printed_publication"] = {},
        ["product"] = {},
        ["publication"] = {},
        ["search"] = {},
        ["section"] = {},
        ["site"] = {},
        ["sound"] = {},
        ["static_page"] = {},
        ["text"] = {},
        ["tour"] = {},
        ["video"] = {},
      },
    },
    entity = {
      ["agent"] = {
        ["fields"] = {
          {
            ["name"] = "alt_titles",
            ["short"] = "Alternate names for this agent",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "birth_date",
            ["short"] = "The year this agent was born",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "death_date",
            ["short"] = "The year this agent died",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "description",
            ["short"] = "A biographical description of the agent",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "is_artist",
            ["short"] = "Whether the agent is an artist.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "sort_title",
            ["short"] = "Sortable name for this agent, typically with last name first.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "ulan_id",
            ["short"] = "Unique identifier of this agent in Getty's ULAN",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
        },
        ["name"] = "agent",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/agents",
                ["parts"] = {
                  "agents",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/artists",
                ["parts"] = {
                  "artists",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/agents/{id}",
                ["parts"] = {
                  "agents",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/artists/{id}",
                ["parts"] = {
                  "artists",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["agent_role"] = {
        ["fields"] = {
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
        },
        ["name"] = "agent_role",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/agent-roles",
                ["parts"] = {
                  "agent-roles",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/agent-roles/{id}",
                ["parts"] = {
                  "agent-roles",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["agent_type"] = {
        ["fields"] = {
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
        },
        ["name"] = "agent_type",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/agent-types",
                ["parts"] = {
                  "agent-types",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/agent-types/{id}",
                ["parts"] = {
                  "agent-types",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["article"] = {
        ["fields"] = {
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "copy",
            ["short"] = "The text of the article",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
        },
        ["name"] = "article",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/articles",
                ["parts"] = {
                  "articles",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/articles/{id}",
                ["parts"] = {
                  "articles",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["artwork"] = {
        ["fields"] = {
          {
            ["name"] = "alt_artist_ids",
            ["short"] = "Unique identifiers of the non-preferred artists/cultures associated with this work",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "alt_classification_ids",
            ["short"] = "Unique identifiers of all other non-preferred classification terms for this work",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "alt_image_ids",
            ["short"] = "Unique identifiers of all non-preferred images of this work.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "alt_material_ids",
            ["short"] = "Unique identifiers of all other non-preferred material terms for this work",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "alt_style_ids",
            ["short"] = "Unique identifiers of all other non-preferred style terms for this work",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "alt_subject_ids",
            ["short"] = "Unique identifiers of all other non-preferred subject terms for this work",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "alt_technique_ids",
            ["short"] = "Unique identifiers of all other non-preferred technique terms for this work",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "alt_titles",
            ["short"] = "Alternate names for this work",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "artist_display",
            ["short"] = "Readable description of the creator of this work.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "artist_id",
            ["short"] = "Unique identifier of the preferred artist/culture associated with this work",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "artist_ids",
            ["short"] = "Unique identifier of all artist/cultures associated with this work",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "artist_title",
            ["short"] = "Name of the preferred artist/culture associated with this work",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "artist_titles",
            ["short"] = "Names of all artist/cultures associated with this work",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "artwork_type_id",
            ["short"] = "Unique identifier of the kind of object or work",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "artwork_type_title",
            ["short"] = "The kind of object or work (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "boost_rank",
            ["short"] = "Manual indication of what rank this artwork should take in search results.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "catalog_based_search_keyword_titles",
            ["short"] = "The keyword search values that would be catalog-based searches on this record",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "catalogue_display",
            ["short"] = "Brief text listing all the catalogues raisonnés which include this work.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "category_ids",
            ["short"] = "Unique identifiers of the categories this work is a part of",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "category_titles",
            ["short"] = "Names of the categories this artwork is a part of",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "classification_id",
            ["short"] = "Unique identifier of the preferred classification term for this work",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "classification_ids",
            ["short"] = "Unique identifiers of all classification terms for this work",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "classification_title",
            ["short"] = "The name of the preferred classification term for this work",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "classification_titles",
            ["short"] = "The names of all classification terms related to this artwork",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "color",
            ["short"] = "Dominant color of this artwork in HSL",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "colorfulness",
            ["short"] = "Unbounded positive float representing an abstract measure of colorfulness.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "copyright_notice",
            ["short"] = "Statement notifying how the work is protected by copyright.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "credit_line",
            ["short"] = "Brief statement indicating how the work came into the collection",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "date_display",
            ["short"] = "Readable, free-text description of the period of time associated with the creation of this work.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "date_end",
            ["short"] = "The year of the period of time associated with the creation of this work",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "date_qualifier_id",
            ["short"] = "Unique identifier of the qualifer to the dates provided for this record.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "date_qualifier_title",
            ["short"] = "Readable, text qualifer to the dates provided for this record.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "date_start",
            ["short"] = "The year of the period of time associated with the creation of this work",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "department_id",
            ["short"] = "Unique identifier of the curatorial department that this work belongs to",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "department_title",
            ["short"] = "Name of the curatorial department that this work belongs to",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "description",
            ["short"] = "Longer explanation describing the work",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "dimensions",
            ["short"] = "The size, shape, scale, and dimensions of the work.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "dimensions_detail",
            ["short"] = "The height, width, depth, and/or diameter of each section of the work in centimeters",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "document_ids",
            ["short"] = "Unique identifiers of assets that serve as documentation for this artwork",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "edition",
            ["short"] = "Edition number if the work is one of many",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "exhibition_history",
            ["short"] = "List of all the places this work has been exhibited",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "fiscal_year",
            ["short"] = "The fiscal year in which the work was acquired.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "fiscal_year_deaccession",
            ["short"] = "The fiscal year in which the work was deaccessioned.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "gallery_id",
            ["short"] = "Unique identifier of the location of this work in our museum",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "gallery_title",
            ["short"] = "The location of this work in our museum",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "has_advanced_imaging",
            ["short"] = "Whether this artwork is enhanced with 3D models, 360 image sequences, Mirador views, etc.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "has_educational_resources",
            ["short"] = "Whether this artwork has any documents tagged as educational",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "has_multimedia_resources",
            ["short"] = "Whether this artwork has any associated microsites, digital publications, or documents tagged as multimedia",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "has_not_been_viewed_much",
            ["short"] = "Whether the artwork hasn't been visited on our website very much",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "image_embedding",
            ["short"] = "The generated embeddings describing the artwork image",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "image_id",
            ["short"] = "Unique identifier of the preferred image to use to represent this work",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "inscriptions",
            ["short"] = "A description of distinguishing or identifying physical markings that are on the work",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "internal_department_id",
            ["short"] = "An internal department id we use for analytics.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "is_boosted",
            ["short"] = "Whether this document should be boosted in search",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "is_on_view",
            ["short"] = "Whether the work is on display",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "is_public_domain",
            ["short"] = "Whether the work is in the public domain, meaning it was created before copyrights existed or has left the copyright term",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "is_zoomable",
            ["short"] = "Whether images of the work are allowed to be displayed in a zoomable interface.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "latitude",
            ["short"] = "Latitude coordinate of the location of this work in our galleries",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "latlon",
            ["short"] = "Latitude and longitude coordinates of the location of this work in our galleries",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "longitude",
            ["short"] = "Longitude coordinate of the location of this work in our galleries",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "main_reference_number",
            ["short"] = "Unique identifier assigned to the artwork upon acquisition",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "material_id",
            ["short"] = "Unique identifier of the preferred material term for this work",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "material_ids",
            ["short"] = "Unique identifiers of all material terms for this work",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "material_titles",
            ["short"] = "The names of all material terms related to this artwork",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "max_zoom_window_size",
            ["short"] = "The maximum size of the window the image is allowed to be viewed in, in pixels.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "medium_display",
            ["short"] = "The substances or materials used in the creation of a work",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "nomisma_id",
            ["short"] = "Unique identifier of this work in the nomisma coin database",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "on_loan_display",
            ["short"] = "If an artwork is on loan, this contains details about the loan",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "pageviews",
            ["short"] = "Approx.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "pageviews_recent",
            ["short"] = "Approx.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "place_of_origin",
            ["short"] = "The location where the creation, design, or production of the work took place, or the original location of the work",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance_text",
            ["short"] = "Ownership/collecting history of the work.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "publication_history",
            ["short"] = "Bibliographic list of all the places this work has been published",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "publishing_verification_level",
            ["short"] = "Indicator of how much metadata on the work in published.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "section_ids",
            ["short"] = "Unique identifiers of the digital publication chapters this work in included in",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "section_titles",
            ["short"] = "Names of the digital publication chapters this work is included in",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "short_description",
            ["short"] = "Short explanation describing the work",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "site_ids",
            ["short"] = "Unique identifiers of the microsites this work is a part of",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "sound_ids",
            ["short"] = "Unique identifiers of the audio about this work",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "style_id",
            ["short"] = "Unique identifier of the preferred style term for this work",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "style_ids",
            ["short"] = "Unique identifiers of all style terms for this work",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "style_title",
            ["short"] = "The name of the preferred style term for this work",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "style_titles",
            ["short"] = "The names of all style terms related to this artwork",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "subject_id",
            ["short"] = "Unique identifier of the preferred subject term for this work",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "subject_ids",
            ["short"] = "Unique identifiers of all subject terms for this work",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "subject_titles",
            ["short"] = "The names of all subject terms related to this artwork",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "technique_id",
            ["short"] = "Unique identifier of the preferred technique term for this work",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "technique_ids",
            ["short"] = "Unique identifiers of all technique terms for this work",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "technique_titles",
            ["short"] = "The names of all technique terms related to this artwork",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "term_titles",
            ["short"] = "The names of the taxonomy tags for this work",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "text_embedding",
            ["short"] = "The generated embeddings of artwork text",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "text_ids",
            ["short"] = "Unique identifiers of the texts about this work",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "theme_titles",
            ["short"] = "The names of all thematic publish categories related to this artwork",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "thumbnail",
            ["short"] = "Metadata about the image referenced by `image_id`.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "video_ids",
            ["short"] = "Unique identifiers of the videos about this work",
            ["type"] = "`$ANY`",
          },
        },
        ["name"] = "artwork",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/artworks",
                ["parts"] = {
                  "artworks",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/artworks/{id}",
                ["parts"] = {
                  "artworks",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["artwork_date_qualifier"] = {
        ["fields"] = {
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
        },
        ["name"] = "artwork_date_qualifier",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/artwork-date-qualifiers",
                ["parts"] = {
                  "artwork-date-qualifiers",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/artwork-date-qualifiers/{id}",
                ["parts"] = {
                  "artwork-date-qualifiers",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["artwork_place_qualifier"] = {
        ["fields"] = {
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
        },
        ["name"] = "artwork_place_qualifier",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/artwork-place-qualifiers",
                ["parts"] = {
                  "artwork-place-qualifiers",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/artwork-place-qualifiers/{id}",
                ["parts"] = {
                  "artwork-place-qualifiers",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["artwork_type"] = {
        ["fields"] = {
          {
            ["name"] = "aat_id",
            ["short"] = "Identifier of reconciled (most similar) term in the Getty's Art and Architecture Thesaurus (AAT)",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
        },
        ["name"] = "artwork_type",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/artwork-types",
                ["parts"] = {
                  "artwork-types",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/artwork-types/{id}",
                ["parts"] = {
                  "artwork-types",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["category_term"] = {
        ["fields"] = {
          {
            ["name"] = "aat_id",
            ["short"] = "Identifier of reconciled (most similar) term in the Getty's Art and Architecture Thesaurus (AAT)",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "parent_id",
            ["short"] = "Unique identifier of this category's parent",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "subtype",
            ["short"] = "Takes one of the following values: classification, material, technique, style, subject, department, theme",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
        },
        ["name"] = "category_term",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/category-terms",
                ["parts"] = {
                  "category-terms",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/category-terms/{id}",
                ["parts"] = {
                  "category-terms",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["digital_publication"] = {
        ["fields"] = {
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "copy",
            ["short"] = "The text of the page",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "web_url",
            ["short"] = "The URL to this page on our website",
            ["type"] = "`$ANY`",
          },
        },
        ["name"] = "digital_publication",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/digital-publications",
                ["parts"] = {
                  "digital-publications",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/digital-publications/{id}",
                ["parts"] = {
                  "digital-publications",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["digital_publication_article"] = {
        ["fields"] = {
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "author_display",
            ["short"] = "A display-friendly text of the authors of this article",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "copy",
            ["short"] = "The text of the article",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "digital_publication_id",
            ["short"] = "Unique identifier of the digital publication this article belongs to",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "web_url",
            ["short"] = "The URL to this article on our website",
            ["type"] = "`$ANY`",
          },
        },
        ["name"] = "digital_publication_article",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/digital-publication-articles",
                ["parts"] = {
                  "digital-publication-articles",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/digital-publication-articles/{id}",
                ["parts"] = {
                  "digital-publication-articles",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["educator_resource"] = {
        ["fields"] = {
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "copy",
            ["short"] = "The text of the page",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "web_url",
            ["short"] = "The URL to this page on our website",
            ["type"] = "`$ANY`",
          },
        },
        ["name"] = "educator_resource",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/educator-resources",
                ["parts"] = {
                  "educator-resources",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/educator-resources/{id}",
                ["parts"] = {
                  "educator-resources",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["event"] = {
        ["fields"] = {
          {
            ["name"] = "alt_audience_ids",
            ["short"] = "Unique identifiers indicating the alternate audiences for this event",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "alt_event_type_ids",
            ["short"] = "Unique identifiers indicating the alternate types of this event",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "audience_id",
            ["short"] = "Unique identifier indicating the preferred audience for this event",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "buy_button_caption",
            ["short"] = "Additional text below the ticket/registration button",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "buy_button_text",
            ["short"] = "The text used on the ticket/registration button",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "date_display",
            ["short"] = "A readable display of the event dates",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "description",
            ["short"] = "All copytext of the event",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "door_time",
            ["short"] = "The time the doors open for this event",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "end_date",
            ["short"] = "The date the event ends",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "end_time",
            ["short"] = "The time the event ends",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "entrance",
            ["short"] = "Which entrance to use for this event",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "event_host_id",
            ["short"] = "Unique identifier of the host (cf.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "event_host_title",
            ["short"] = "Unique identifier of the host (cf.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "event_type_id",
            ["short"] = "Unique identifier indicating the preferred type of this event",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "header_description",
            ["short"] = "Brief description of the event displayed below the title",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "hero_caption",
            ["short"] = "Text displayed with the hero image on the event",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "image_url",
            ["short"] = "The URL of an image representing this page",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "is_admission_required",
            ["short"] = "Whether admission to the museum is required to attend this event",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "is_after_hours",
            ["short"] = "Whether the event is to be held after the museum closes",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "is_free",
            ["short"] = "Whether the event is free",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "is_member_exclusive",
            ["short"] = "Whether the event is exclusive to members of the museum",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "is_private",
            ["short"] = "Whether the event is private",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "is_registration_required",
            ["short"] = "Whether registration is required to attend the event",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "is_sales_button_hidden",
            ["short"] = "Whether the buy tickets button is hidden on the website event page",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "is_sold_out",
            ["short"] = "Whether the event is sold out",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "is_ticketed",
            ["short"] = "Whether a ticket is required to attend the event",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "is_virtual_event",
            ["short"] = "Whether the event is being held virtually",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "join_url",
            ["short"] = "URL to the membership signup page via this event",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "layout_type",
            ["short"] = "Number indicating the type of layout this event page uses",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "list_description",
            ["short"] = "One-sentence description of the event displayed in listings",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "location",
            ["short"] = "Where the event takes place",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "program_ids",
            ["short"] = "Unique identifiers indicating the programs this event is a part of",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "program_titles",
            ["short"] = "Titles of the programs this event is a part of",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "rsvp_link",
            ["short"] = "The URL to the sales site for this event",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "search_tags",
            ["short"] = "Editor-specified list of tags to aid in internal search",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "short_description",
            ["short"] = "Brief description of the event",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "slug",
            ["short"] = "A string used in the URL for this event",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "start_date",
            ["short"] = "The date the event begins",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "start_time",
            ["short"] = "The time the event starts",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "survey_url",
            ["short"] = "URL to the survey associated with this event",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ticketed_event_id",
            ["short"] = "Unique identifier of the event in the ticketing system this website event is tied to",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "title_display",
            ["short"] = "The name of this event formatted with HTML (optional)",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "virtual_event_passcode",
            ["short"] = "Passcode to access the virtual event",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "virtual_event_url",
            ["short"] = "URL to the virtual event",
            ["type"] = "`$ANY`",
          },
        },
        ["name"] = "event",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/events",
                ["parts"] = {
                  "events",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/events/{id}",
                ["parts"] = {
                  "events",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["event_occurrence"] = {
        ["fields"] = {
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "button_caption",
            ["short"] = "Additional text below the ticket/registration button",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "button_text",
            ["short"] = "The text used on the ticket/registration button",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "button_url",
            ["short"] = "The URL to the sales site or an RSVP link for this event",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "description",
            ["short"] = "Description of the event",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "end_at",
            ["short"] = "The date the event occurrence ends",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "event_id",
            ["short"] = "Identifier of the master event of which this is an occurrence",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "image_url",
            ["short"] = "The URL of an image representing this page",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "is_private",
            ["short"] = "Whether the event is private.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "is_sales_button_hidden",
            ["short"] = "Whether the buy tickets button is hidden on the website event page",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "is_ticketed",
            ["short"] = "Whether a ticket is required to attend the event",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "location",
            ["short"] = "Where the event takes place",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "off_sale_at",
            ["short"] = "Date and time the event goes off sale",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "on_sale_at",
            ["short"] = "Date and time the event goes on sale",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "short_description",
            ["short"] = "Brief description of the event",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "start_at",
            ["short"] = "The date the event occurrence begins",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "title_display",
            ["short"] = "The name of this event formatted with HTML (optional)",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
        },
        ["name"] = "event_occurrence",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/event-occurrences",
                ["parts"] = {
                  "event-occurrences",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/event-occurrences/{id}",
                ["parts"] = {
                  "event-occurrences",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["event_program"] = {
        ["fields"] = {
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "is_affiliate_group",
            ["short"] = "Whether this program represents an affiliate group",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "is_event_host",
            ["short"] = "Whether this program represents an event host",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
        },
        ["name"] = "event_program",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/event-programs",
                ["parts"] = {
                  "event-programs",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/event-programs/{id}",
                ["parts"] = {
                  "event-programs",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["exhibition"] = {
        ["fields"] = {
          {
            ["name"] = "aic_end_at",
            ["short"] = "Date the exhibition closed at the Art Institute of Chicago",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "aic_start_at",
            ["short"] = "Date the exhibition opened at the Art Institute of Chicago",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "alt_image_ids",
            ["short"] = "Unique identifiers of all non-preferred images of this exhibition.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "artist_ids",
            ["short"] = "Unique identifiers of the artist agent records representing who was shown in the exhibition",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "artwork_ids",
            ["short"] = "Unique identifiers of the artworks that were part of the exhibition",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "artwork_titles",
            ["short"] = "Names of the artworks that were part of the exhibition",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "document_ids",
            ["short"] = "Unique identifiers of assets that serve as documentation for this exhibition",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "gallery_id",
            ["short"] = "Unique identifier of the gallery that mainly housed the exhibition",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "gallery_title",
            ["short"] = "The name of the gallery that mainly housed the exhibition",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "image_id",
            ["short"] = "Unique identifier of the preferred image to use to represent this exhibition",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "image_url",
            ["short"] = "URL to the hero image from the website",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "is_featured",
            ["short"] = "Is this exhibition currently featured on our website?",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "is_published",
            ["short"] = "Is this exhibition currently published on our website?",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "position",
            ["short"] = "Numering position represnting the order in which this exhibition is featured on the website",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "short_description",
            ["short"] = "Brief explanation of what this exhibition is",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "site_ids",
            ["short"] = "Unique identifiers of the microsites this exhibition is a part of",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "status",
            ["short"] = "Whether the exhibition is open or closed",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "web_url",
            ["short"] = "URL to this exhibition on our website",
            ["type"] = "`$ANY`",
          },
        },
        ["name"] = "exhibition",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/exhibitions",
                ["parts"] = {
                  "exhibitions",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/exhibitions/{id}",
                ["parts"] = {
                  "exhibitions",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["gallery"] = {
        ["fields"] = {
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "floor",
            ["short"] = "The level the gallery is on, e.g., 1, 2, 3, or LL",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "is_closed",
            ["short"] = "Whether the gallery is currently closed",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "latitude",
            ["short"] = "Latitude coordinate of the center of the room",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "latlon",
            ["short"] = "Latitude and longitude coordinates of the center of the room",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "longitude",
            ["short"] = "Longitude coordinate of the center of the room",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "number",
            ["short"] = "The gallery's room number.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "tgn_id",
            ["short"] = "Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN)",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
        },
        ["name"] = "gallery",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/galleries",
                ["parts"] = {
                  "galleries",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/galleries/{id}",
                ["parts"] = {
                  "galleries",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["generic_page"] = {
        ["fields"] = {
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "copy",
            ["short"] = "The text of the page",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "search_tags",
            ["short"] = "Editor-specified list of tags to aid in internal search",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "web_url",
            ["short"] = "The URL to this page on our website",
            ["type"] = "`$ANY`",
          },
        },
        ["name"] = "generic_page",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/generic-pages",
                ["parts"] = {
                  "generic-pages",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/generic-pages/{id}",
                ["parts"] = {
                  "generic-pages",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["highlight"] = {
        ["fields"] = {
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "copy",
            ["short"] = "The text of the highlight description",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
        },
        ["name"] = "highlight",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/highlights",
                ["parts"] = {
                  "highlights",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/highlights/{id}",
                ["parts"] = {
                  "highlights",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["hour"] = {
        ["fields"] = {
          {
            ["name"] = "additional_text",
            ["short"] = "Additional information about the hours",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "friday_is_closed",
            ["short"] = "Whether the museum is closed on Fridays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "friday_member_close",
            ["short"] = "The time member hours ends on Fridays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "friday_member_open",
            ["short"] = "The time member hours starts on Fridays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "friday_public_close",
            ["short"] = "The time public hours ends on Fridays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "friday_public_open",
            ["short"] = "The time public hours starts on Fridays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "monday_is_closed",
            ["short"] = "Whether the museum is closed on Mondays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "monday_member_close",
            ["short"] = "The time member hours ends on Mondays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "monday_member_open",
            ["short"] = "The time member hours starts on Mondays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "monday_public_close",
            ["short"] = "The time public hours ends on Mondays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "monday_public_open",
            ["short"] = "The time public hours starts on Mondays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "saturday_is_closed",
            ["short"] = "Whether the museum is closed on Saturdays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "saturday_member_close",
            ["short"] = "The time member hours ends on Saturdays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "saturday_member_open",
            ["short"] = "The time member hours starts on Saturdays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "saturday_public_close",
            ["short"] = "The time public hours ends on Saturdays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "saturday_public_open",
            ["short"] = "The time public hours starts on Saturdays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "summary",
            ["short"] = "Readable summary of the hours",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "sunday_is_closed",
            ["short"] = "Whether the museum is closed on Sundays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "sunday_member_close",
            ["short"] = "The time member hours ends on Sundays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "sunday_member_open",
            ["short"] = "The time member hours starts on Sundays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "sunday_public_close",
            ["short"] = "The time public hours ends on Sundays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "sunday_public_open",
            ["short"] = "The time public hours starts on Sundays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "thursday_is_closed",
            ["short"] = "Whether the museum is closed on Thursdays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "thursday_member_close",
            ["short"] = "The time member hours ends on Thursdays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "thursday_member_open",
            ["short"] = "The time member hours starts on Thursdays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "thursday_public_close",
            ["short"] = "The time public hours ends on Thursdays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "thursday_public_open",
            ["short"] = "The time public hours starts on Thursdays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tuesday_is_closed",
            ["short"] = "Whether the museum is closed on Tuesdays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "tuesday_member_close",
            ["short"] = "The time member hours ends on Tuesdays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "tuesday_member_open",
            ["short"] = "The time member hours starts on Tuesdays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "tuesday_public_close",
            ["short"] = "The time public hours ends on Tuesdays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "tuesday_public_open",
            ["short"] = "The time public hours starts on Tuesdays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "wednesday_is_closed",
            ["short"] = "Whether the museum is closed on Wednesdays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "wednesday_member_close",
            ["short"] = "The time member hours ends on Wednesdays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "wednesday_member_open",
            ["short"] = "The time member hours starts on Wednesdays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "wednesday_public_close",
            ["short"] = "The time public hours ends on Wednesdays",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "wednesday_public_open",
            ["short"] = "The time public hours starts on Wednesdays",
            ["type"] = "`$ANY`",
          },
        },
        ["name"] = "hour",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/hours",
                ["parts"] = {
                  "hours",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/hours/{id}",
                ["parts"] = {
                  "hours",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["image"] = {
        ["fields"] = {
          {
            ["name"] = "ahash",
            ["short"] = "Image hash generated using ahash algorithm with 64 boolean subfields",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "alt_text",
            ["short"] = "Alternative text for the asset to describe it to people with low or no vision",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "artwork_ids",
            ["short"] = "Unique identifiers of the artworks associated with this asset",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "artwork_titles",
            ["short"] = "Names of the artworks associated with this asset",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "color",
            ["short"] = "Dominant color of this image in HSL",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "colorfulness",
            ["short"] = "Unbounded positive float representing an abstract measure of colorfulness.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "content",
            ["short"] = "Text of or URL to the contents of this asset",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "content_e_tag",
            ["short"] = "Arbitrary unique identifier that changes when the binary file gets updated",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "credit_line",
            ["short"] = "Asset-specific copyright information",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "fingerprint",
            ["short"] = "Image hashes: aHash, dHash, pHash, wHash",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "height",
            ["short"] = "Native height of the image",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "iiif_url",
            ["short"] = "IIIF URL of this image",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "is_educational_resource",
            ["short"] = "Whether this resource is considered to be educational",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "is_multimedia_resource",
            ["short"] = "Whether this resource is considered to be multimedia",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "is_teacher_resource",
            ["short"] = "Whether this resource is considered to be educational",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "lake_guid",
            ["short"] = "Unique UUID of this resource in LAKE, our DAMS.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "lqip",
            ["short"] = "Low-quality image placeholder (LQIP).",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "phash",
            ["short"] = "Image hash generated using phash algorithm with 64 boolean subfields",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "type",
            ["short"] = "Type always takes one of the following values: image, sound, text, video",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "width",
            ["short"] = "Native width of the image",
            ["type"] = "`$NUMBER`",
          },
        },
        ["name"] = "image",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/images",
                ["parts"] = {
                  "images",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/images/{id}",
                ["parts"] = {
                  "images",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["landing_page"] = {
        ["fields"] = {
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "copy",
            ["short"] = "The text of the page",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "search_tags",
            ["short"] = "Editor-specified list of tags to aid in internal search",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "web_url",
            ["short"] = "The URL to this page on our website",
            ["type"] = "`$ANY`",
          },
        },
        ["name"] = "landing_page",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/landing-pages",
                ["parts"] = {
                  "landing-pages",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/landing-pages/{id}",
                ["parts"] = {
                  "landing-pages",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["place"] = {
        ["fields"] = {
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "latitude",
            ["short"] = "Latitude coordinate of the center of the room",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "longitude",
            ["short"] = "Longitude coordinate of the center of the room",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "tgn_id",
            ["short"] = "Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN)",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
        },
        ["name"] = "place",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/places",
                ["parts"] = {
                  "places",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/places/{id}",
                ["parts"] = {
                  "places",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["press_release"] = {
        ["fields"] = {
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "copy",
            ["short"] = "The text of the page",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "web_url",
            ["short"] = "The URL to this page on our website",
            ["type"] = "`$ANY`",
          },
        },
        ["name"] = "press_release",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/press-releases",
                ["parts"] = {
                  "press-releases",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/press-releases/{id}",
                ["parts"] = {
                  "press-releases",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["printed_publication"] = {
        ["fields"] = {
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "copy",
            ["short"] = "The text of the page",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "web_url",
            ["short"] = "The URL to this page on our website",
            ["type"] = "`$ANY`",
          },
        },
        ["name"] = "printed_publication",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/printed-publications",
                ["parts"] = {
                  "printed-publications",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/printed-publications/{id}",
                ["parts"] = {
                  "printed-publications",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["product"] = {
        ["fields"] = {
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "artist_ids",
            ["short"] = "Unique identifiers of the artists associated with this product",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "artwork_ids",
            ["short"] = "Unique identifiers of the artworks associated with this product",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "description",
            ["short"] = "Explanation of what this product is",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "exhibition_ids",
            ["short"] = "Unique identifiers of the exhibitions associated with this product",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "external_sku",
            ["short"] = "Numeric product identification code of a machine-readable barcode, when the customer sku differs from our internal one",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "image_url",
            ["short"] = "URL of an image for this product",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "max_compare_at_price",
            ["short"] = "Number indicating how much the most expensive variant of a product cost before a sale",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "max_current_price",
            ["short"] = "Number indicating how much the most expensive variant of a product costs right now",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "min_compare_at_price",
            ["short"] = "Number indicating how much the least expensive variant of a product cost before a sale",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "min_current_price",
            ["short"] = "Number indicating how much the least expensive variant of a product costs right now",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "price_display",
            ["short"] = "Explanation of what this product is",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "web_url",
            ["short"] = "URL of this product in the shop",
            ["type"] = "`$ANY`",
          },
        },
        ["name"] = "product",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/products",
                ["parts"] = {
                  "products",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/products/{id}",
                ["parts"] = {
                  "products",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["publication"] = {
        ["fields"] = {
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "section_ids",
            ["short"] = "Unique identifiers of the sections of this publication",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "web_url",
            ["short"] = "URL to the publication",
            ["type"] = "`$ANY`",
          },
        },
        ["name"] = "publication",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/publications",
                ["parts"] = {
                  "publications",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/publications/{id}",
                ["parts"] = {
                  "publications",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["search"] = {
        ["fields"] = {
          {
            ["name"] = "api_id",
            ["short"] = "API unique identifier",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "api_link",
            ["short"] = "URL to this recource in the API",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "Name of the model the resource represents",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier within the search index",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "is_boosted",
            ["short"] = "Whether this record has been flagged to be boosted",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "score",
            ["short"] = "Search index ranking of the result",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "thumbnail",
            ["short"] = "Metadata on the image representing this record",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date this record was last updated in the API",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "search",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "facet",
                      ["orig"] = "facet",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from",
                      ["orig"] = "from",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "q",
                      ["orig"] = "q",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "query",
                      ["orig"] = "query",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/agents/search",
                ["parts"] = {
                  "agents",
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "facet",
                    "from",
                    "q",
                    "query",
                    "size",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "facet",
                      ["orig"] = "facet",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from",
                      ["orig"] = "from",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "q",
                      ["orig"] = "q",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "query",
                      ["orig"] = "query",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/articles/search",
                ["parts"] = {
                  "articles",
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "facet",
                    "from",
                    "q",
                    "query",
                    "size",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "facet",
                      ["orig"] = "facet",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from",
                      ["orig"] = "from",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "q",
                      ["orig"] = "q",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "query",
                      ["orig"] = "query",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/artworks/search",
                ["parts"] = {
                  "artworks",
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "facet",
                    "from",
                    "q",
                    "query",
                    "size",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "facet",
                      ["orig"] = "facet",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from",
                      ["orig"] = "from",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "q",
                      ["orig"] = "q",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "query",
                      ["orig"] = "query",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/category-terms/search",
                ["parts"] = {
                  "category-terms",
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "facet",
                    "from",
                    "q",
                    "query",
                    "size",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "facet",
                      ["orig"] = "facet",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from",
                      ["orig"] = "from",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "q",
                      ["orig"] = "q",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "query",
                      ["orig"] = "query",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/digital-publication-articles/search",
                ["parts"] = {
                  "digital-publication-articles",
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "facet",
                    "from",
                    "q",
                    "query",
                    "size",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "facet",
                      ["orig"] = "facet",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from",
                      ["orig"] = "from",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "q",
                      ["orig"] = "q",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "query",
                      ["orig"] = "query",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/digital-publications/search",
                ["parts"] = {
                  "digital-publications",
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "facet",
                    "from",
                    "q",
                    "query",
                    "size",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "facet",
                      ["orig"] = "facet",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from",
                      ["orig"] = "from",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "q",
                      ["orig"] = "q",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "query",
                      ["orig"] = "query",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/educator-resources/search",
                ["parts"] = {
                  "educator-resources",
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "facet",
                    "from",
                    "q",
                    "query",
                    "size",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "facet",
                      ["orig"] = "facet",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from",
                      ["orig"] = "from",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "q",
                      ["orig"] = "q",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "query",
                      ["orig"] = "query",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/event-occurrences/search",
                ["parts"] = {
                  "event-occurrences",
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "facet",
                    "from",
                    "q",
                    "query",
                    "size",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "facet",
                      ["orig"] = "facet",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from",
                      ["orig"] = "from",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "q",
                      ["orig"] = "q",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "query",
                      ["orig"] = "query",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/event-programs/search",
                ["parts"] = {
                  "event-programs",
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "facet",
                    "from",
                    "q",
                    "query",
                    "size",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "facet",
                      ["orig"] = "facet",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from",
                      ["orig"] = "from",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "q",
                      ["orig"] = "q",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "query",
                      ["orig"] = "query",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/events/search",
                ["parts"] = {
                  "events",
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "facet",
                    "from",
                    "q",
                    "query",
                    "size",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "facet",
                      ["orig"] = "facet",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from",
                      ["orig"] = "from",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "q",
                      ["orig"] = "q",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "query",
                      ["orig"] = "query",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/exhibitions/search",
                ["parts"] = {
                  "exhibitions",
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "facet",
                    "from",
                    "q",
                    "query",
                    "size",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "facet",
                      ["orig"] = "facet",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from",
                      ["orig"] = "from",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "q",
                      ["orig"] = "q",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "query",
                      ["orig"] = "query",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/galleries/search",
                ["parts"] = {
                  "galleries",
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "facet",
                    "from",
                    "q",
                    "query",
                    "size",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "facet",
                      ["orig"] = "facet",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from",
                      ["orig"] = "from",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "q",
                      ["orig"] = "q",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "query",
                      ["orig"] = "query",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/generic-pages/search",
                ["parts"] = {
                  "generic-pages",
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "facet",
                    "from",
                    "q",
                    "query",
                    "size",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "facet",
                      ["orig"] = "facet",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from",
                      ["orig"] = "from",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "q",
                      ["orig"] = "q",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "query",
                      ["orig"] = "query",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/highlights/search",
                ["parts"] = {
                  "highlights",
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "facet",
                    "from",
                    "q",
                    "query",
                    "size",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "facet",
                      ["orig"] = "facet",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from",
                      ["orig"] = "from",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "q",
                      ["orig"] = "q",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "query",
                      ["orig"] = "query",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/hours/search",
                ["parts"] = {
                  "hours",
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "facet",
                    "from",
                    "q",
                    "query",
                    "size",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "facet",
                      ["orig"] = "facet",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from",
                      ["orig"] = "from",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "q",
                      ["orig"] = "q",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "query",
                      ["orig"] = "query",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/images/search",
                ["parts"] = {
                  "images",
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "facet",
                    "from",
                    "q",
                    "query",
                    "size",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "facet",
                      ["orig"] = "facet",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from",
                      ["orig"] = "from",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "q",
                      ["orig"] = "q",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "query",
                      ["orig"] = "query",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/landing-pages/search",
                ["parts"] = {
                  "landing-pages",
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "facet",
                    "from",
                    "q",
                    "query",
                    "size",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "facet",
                      ["orig"] = "facet",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from",
                      ["orig"] = "from",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "q",
                      ["orig"] = "q",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "query",
                      ["orig"] = "query",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/mobile-sounds/search",
                ["parts"] = {
                  "mobile-sounds",
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "facet",
                    "from",
                    "q",
                    "query",
                    "size",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "facet",
                      ["orig"] = "facet",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from",
                      ["orig"] = "from",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "q",
                      ["orig"] = "q",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "query",
                      ["orig"] = "query",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/places/search",
                ["parts"] = {
                  "places",
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "facet",
                    "from",
                    "q",
                    "query",
                    "size",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "facet",
                      ["orig"] = "facet",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from",
                      ["orig"] = "from",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "q",
                      ["orig"] = "q",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "query",
                      ["orig"] = "query",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/press-releases/search",
                ["parts"] = {
                  "press-releases",
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "facet",
                    "from",
                    "q",
                    "query",
                    "size",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "facet",
                      ["orig"] = "facet",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from",
                      ["orig"] = "from",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "q",
                      ["orig"] = "q",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "query",
                      ["orig"] = "query",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/printed-publications/search",
                ["parts"] = {
                  "printed-publications",
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "facet",
                    "from",
                    "q",
                    "query",
                    "size",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "facet",
                      ["orig"] = "facet",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from",
                      ["orig"] = "from",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "q",
                      ["orig"] = "q",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "query",
                      ["orig"] = "query",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/products/search",
                ["parts"] = {
                  "products",
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "facet",
                    "from",
                    "q",
                    "query",
                    "size",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "facet",
                      ["orig"] = "facet",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from",
                      ["orig"] = "from",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "q",
                      ["orig"] = "q",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "query",
                      ["orig"] = "query",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/publications/search",
                ["parts"] = {
                  "publications",
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "facet",
                    "from",
                    "q",
                    "query",
                    "size",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "facet",
                      ["orig"] = "facet",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from",
                      ["orig"] = "from",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "q",
                      ["orig"] = "q",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "query",
                      ["orig"] = "query",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/search",
                ["parts"] = {
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "facet",
                    "from",
                    "q",
                    "query",
                    "size",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "facet",
                      ["orig"] = "facet",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from",
                      ["orig"] = "from",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "q",
                      ["orig"] = "q",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "query",
                      ["orig"] = "query",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/sections/search",
                ["parts"] = {
                  "sections",
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "facet",
                    "from",
                    "q",
                    "query",
                    "size",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "facet",
                      ["orig"] = "facet",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from",
                      ["orig"] = "from",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "q",
                      ["orig"] = "q",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "query",
                      ["orig"] = "query",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/sites/search",
                ["parts"] = {
                  "sites",
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "facet",
                    "from",
                    "q",
                    "query",
                    "size",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "facet",
                      ["orig"] = "facet",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from",
                      ["orig"] = "from",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "q",
                      ["orig"] = "q",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "query",
                      ["orig"] = "query",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/sounds/search",
                ["parts"] = {
                  "sounds",
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "facet",
                    "from",
                    "q",
                    "query",
                    "size",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "facet",
                      ["orig"] = "facet",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from",
                      ["orig"] = "from",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "q",
                      ["orig"] = "q",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "query",
                      ["orig"] = "query",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/static-pages/search",
                ["parts"] = {
                  "static-pages",
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "facet",
                    "from",
                    "q",
                    "query",
                    "size",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "facet",
                      ["orig"] = "facet",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from",
                      ["orig"] = "from",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "q",
                      ["orig"] = "q",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "query",
                      ["orig"] = "query",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/texts/search",
                ["parts"] = {
                  "texts",
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "facet",
                    "from",
                    "q",
                    "query",
                    "size",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "facet",
                      ["orig"] = "facet",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from",
                      ["orig"] = "from",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "q",
                      ["orig"] = "q",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "query",
                      ["orig"] = "query",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/tours/search",
                ["parts"] = {
                  "tours",
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "facet",
                    "from",
                    "q",
                    "query",
                    "size",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "facet",
                      ["orig"] = "facet",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from",
                      ["orig"] = "from",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "q",
                      ["orig"] = "q",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "query",
                      ["orig"] = "query",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/videos/search",
                ["parts"] = {
                  "videos",
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "facet",
                    "from",
                    "q",
                    "query",
                    "size",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["section"] = {
        ["fields"] = {
          {
            ["name"] = "accession",
            ["short"] = "An accession number parsed from the title or tombstone",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "artwork_id",
            ["short"] = "Unique identifier of the artwork with which this section is associated",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "content",
            ["short"] = "Content of this section in plaintext",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "generic_page_id",
            ["short"] = "Unique identifier of the page on the website that represents the publication this section belongs to",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "publication_id",
            ["short"] = "Unique identifier of the publication this section belongs to",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "publication_title",
            ["short"] = "Name of the publication this section belongs to",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "web_url",
            ["short"] = "URL to the section",
            ["type"] = "`$ANY`",
          },
        },
        ["name"] = "section",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/sections",
                ["parts"] = {
                  "sections",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/sections/{id}",
                ["parts"] = {
                  "sections",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["site"] = {
        ["fields"] = {
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "artwork_ids",
            ["short"] = "Unique identifiers of the artworks this site is associated with",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "artwork_titles",
            ["short"] = "Names of the artworks this site is associated with",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "description",
            ["short"] = "Explanation of what this site is",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "exhibition_ids",
            ["short"] = "Unique identifier of the exhibitions this site is associated with",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "exhibition_titles",
            ["short"] = "Names of the exhibitions this site is associated with",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "web_url",
            ["short"] = "URL to this site",
            ["type"] = "`$ANY`",
          },
        },
        ["name"] = "site",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/sites",
                ["parts"] = {
                  "sites",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/sites/{id}",
                ["parts"] = {
                  "sites",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["sound"] = {
        ["fields"] = {
          {
            ["name"] = "alt_text",
            ["short"] = "Alternative text for the asset to describe it to people with low or no vision",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "artwork_ids",
            ["short"] = "Unique identifiers of the artworks associated with this asset",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "artwork_titles",
            ["short"] = "Names of the artworks associated with this asset",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "content",
            ["short"] = "Text of or URL to the contents of this asset",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "content_e_tag",
            ["short"] = "Arbitrary unique identifier that changes when the binary file gets updated",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "credit_line",
            ["short"] = "Asset-specific copyright information",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "is_educational_resource",
            ["short"] = "Whether this resource is considered to be educational",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "is_multimedia_resource",
            ["short"] = "Whether this resource is considered to be multimedia",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "is_teacher_resource",
            ["short"] = "Whether this resource is considered to be educational",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "lake_guid",
            ["short"] = "Unique UUID of this resource in LAKE, our DAMS.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "Name of this mobile audio file – derived from the artwork and tour titles",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "transcript",
            ["short"] = "Text transcription of the audio file",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "type",
            ["short"] = "Type always takes one of the following values: image, sound, text, video",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "web_url",
            ["short"] = "URL to the audio file",
            ["type"] = "`$ANY`",
          },
        },
        ["name"] = "sound",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/mobile-sounds",
                ["parts"] = {
                  "mobile-sounds",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/sounds",
                ["parts"] = {
                  "sounds",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/mobile-sounds/{id}",
                ["parts"] = {
                  "mobile-sounds",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/sounds/{id}",
                ["parts"] = {
                  "sounds",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["static_page"] = {
        ["fields"] = {
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "web_url",
            ["short"] = "The URL to this page on our website",
            ["type"] = "`$ANY`",
          },
        },
        ["name"] = "static_page",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/static-pages",
                ["parts"] = {
                  "static-pages",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/static-pages/{id}",
                ["parts"] = {
                  "static-pages",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["text"] = {
        ["fields"] = {
          {
            ["name"] = "alt_text",
            ["short"] = "Alternative text for the asset to describe it to people with low or no vision",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "artwork_ids",
            ["short"] = "Unique identifiers of the artworks associated with this asset",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "artwork_titles",
            ["short"] = "Names of the artworks associated with this asset",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "content",
            ["short"] = "Text of or URL to the contents of this asset",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "content_e_tag",
            ["short"] = "Arbitrary unique identifier that changes when the binary file gets updated",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "credit_line",
            ["short"] = "Asset-specific copyright information",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "is_educational_resource",
            ["short"] = "Whether this resource is considered to be educational",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "is_multimedia_resource",
            ["short"] = "Whether this resource is considered to be multimedia",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "is_teacher_resource",
            ["short"] = "Whether this resource is considered to be educational",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "lake_guid",
            ["short"] = "Unique UUID of this resource in LAKE, our DAMS.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "type",
            ["short"] = "Type always takes one of the following values: image, sound, text, video",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
        },
        ["name"] = "text",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/texts",
                ["parts"] = {
                  "texts",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/texts/{id}",
                ["parts"] = {
                  "texts",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["tour"] = {
        ["fields"] = {
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "artist_titles",
            ["short"] = "Names of the artists of the artworks featured in this tour's tour stops",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "artwork_titles",
            ["short"] = "Names of the artworks featured in this tour's tour stops",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "description",
            ["short"] = "Explanation of what the tour is",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "image",
            ["short"] = "The main image for the tour",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "intro",
            ["short"] = "Text introducing the tour",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "intro_link",
            ["short"] = "Link to the audio file of the introduction",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "intro_transcript",
            ["short"] = "Transcript of the introduction audio to the tour",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "weight",
            ["short"] = "Number representing this tour's sort order",
            ["type"] = "`$NUMBER`",
          },
        },
        ["name"] = "tour",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/tours",
                ["parts"] = {
                  "tours",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/tours/{id}",
                ["parts"] = {
                  "tours",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["video"] = {
        ["fields"] = {
          {
            ["name"] = "alt_text",
            ["short"] = "Alternative text for the asset to describe it to people with low or no vision",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_link",
            ["short"] = "REST API link for this resource",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "api_model",
            ["short"] = "REST API resource type or endpoint",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "artwork_ids",
            ["short"] = "Unique identifiers of the artworks associated with this asset",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "artwork_titles",
            ["short"] = "Names of the artworks associated with this asset",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "content",
            ["short"] = "Text of or URL to the contents of this asset",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "content_e_tag",
            ["short"] = "Arbitrary unique identifier that changes when the binary file gets updated",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "credit_line",
            ["short"] = "Asset-specific copyright information",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique identifier of this resource.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "is_educational_resource",
            ["short"] = "Whether this resource is considered to be educational",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "is_multimedia_resource",
            ["short"] = "Whether this resource is considered to be multimedia",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "is_teacher_resource",
            ["short"] = "Whether this resource is considered to be educational",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "lake_guid",
            ["short"] = "Unique UUID of this resource in LAKE, our DAMS.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "source_updated_at",
            ["short"] = "Date and time the resource was updated in the source system",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_all",
            ["short"] = "Internal field to power the `/autosuggest` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "suggest_autocomplete_boosted",
            ["short"] = "Internal field to power the `/autocomplete` endpoint.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "timestamp",
            ["short"] = "Date and time the record was updated in the aggregator search index",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "title",
            ["short"] = "The name of this resource",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "type",
            ["short"] = "Type always takes one of the following values: image, sound, text, video",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "updated_at",
            ["short"] = "Date and time the record was updated in the aggregator database",
            ["type"] = "`$ANY`",
          },
        },
        ["name"] = "video",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/videos",
                ["parts"] = {
                  "videos",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/videos/{id}",
                ["parts"] = {
                  "videos",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
    },
  }
end


local function make_feature(name)
  local features = require("features")
  local factory = features[name]
  if factory ~= nil then
    return factory()
  end
  return features.base()
end


-- Attach make_feature to the SDK class
local function setup_sdk(SDK)
  SDK._make_feature = make_feature
end


return make_config
