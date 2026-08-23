package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "ArtInstituteOfChicago",
			"slug": "art-institute-of-chicago",
			"version": "0.0.1",
			"target": "go",
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
			},
		},
		"options": map[string]any{
			"base": "https://api.artic.edu/api/v1",
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"agent": map[string]any{},
				"agent_role": map[string]any{},
				"agent_type": map[string]any{},
				"article": map[string]any{},
				"artwork": map[string]any{},
				"artwork_date_qualifier": map[string]any{},
				"artwork_place_qualifier": map[string]any{},
				"artwork_type": map[string]any{},
				"category_term": map[string]any{},
				"digital_publication": map[string]any{},
				"digital_publication_article": map[string]any{},
				"educator_resource": map[string]any{},
				"event": map[string]any{},
				"event_occurrence": map[string]any{},
				"event_program": map[string]any{},
				"exhibition": map[string]any{},
				"gallery": map[string]any{},
				"generic_page": map[string]any{},
				"highlight": map[string]any{},
				"hour": map[string]any{},
				"image": map[string]any{},
				"landing_page": map[string]any{},
				"place": map[string]any{},
				"press_release": map[string]any{},
				"printed_publication": map[string]any{},
				"product": map[string]any{},
				"publication": map[string]any{},
				"search": map[string]any{},
				"section": map[string]any{},
				"site": map[string]any{},
				"sound": map[string]any{},
				"static_page": map[string]any{},
				"text": map[string]any{},
				"tour": map[string]any{},
				"video": map[string]any{},
			},
		},
		"entity": map[string]any{
			"agent": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "alt_titles",
						"short": "Alternate names for this agent",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "birth_date",
						"short": "The year this agent was born",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "death_date",
						"short": "The year this agent died",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "description",
						"short": "A biographical description of the agent",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "is_artist",
						"short": "Whether the agent is an artist.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "sort_title",
						"short": "Sortable name for this agent, typically with last name first.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ulan_id",
						"short": "Unique identifier of this agent in Getty's ULAN",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
				},
				"name": "agent",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/agents",
								"parts": []any{
									"agents",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/artists",
								"parts": []any{
									"artists",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/agents/{id}",
								"parts": []any{
									"agents",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/artists/{id}",
								"parts": []any{
									"artists",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"agent_role": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
				},
				"name": "agent_role",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/agent-roles",
								"parts": []any{
									"agent-roles",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/agent-roles/{id}",
								"parts": []any{
									"agent-roles",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"agent_type": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
				},
				"name": "agent_type",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/agent-types",
								"parts": []any{
									"agent-types",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/agent-types/{id}",
								"parts": []any{
									"agent-types",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"article": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "copy",
						"short": "The text of the article",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
				},
				"name": "article",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/articles",
								"parts": []any{
									"articles",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/articles/{id}",
								"parts": []any{
									"articles",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"artwork": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "alt_artist_ids",
						"short": "Unique identifiers of the non-preferred artists/cultures associated with this work",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "alt_classification_ids",
						"short": "Unique identifiers of all other non-preferred classification terms for this work",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "alt_image_ids",
						"short": "Unique identifiers of all non-preferred images of this work.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "alt_material_ids",
						"short": "Unique identifiers of all other non-preferred material terms for this work",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "alt_style_ids",
						"short": "Unique identifiers of all other non-preferred style terms for this work",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "alt_subject_ids",
						"short": "Unique identifiers of all other non-preferred subject terms for this work",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "alt_technique_ids",
						"short": "Unique identifiers of all other non-preferred technique terms for this work",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "alt_titles",
						"short": "Alternate names for this work",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artist_display",
						"short": "Readable description of the creator of this work.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artist_id",
						"short": "Unique identifier of the preferred artist/culture associated with this work",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "artist_ids",
						"short": "Unique identifier of all artist/cultures associated with this work",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artist_title",
						"short": "Name of the preferred artist/culture associated with this work",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artist_titles",
						"short": "Names of all artist/cultures associated with this work",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artwork_type_id",
						"short": "Unique identifier of the kind of object or work",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "artwork_type_title",
						"short": "The kind of object or work (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "boost_rank",
						"short": "Manual indication of what rank this artwork should take in search results.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "catalog_based_search_keyword_titles",
						"short": "The keyword search values that would be catalog-based searches on this record",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "catalogue_display",
						"short": "Brief text listing all the catalogues raisonnés which include this work.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "category_ids",
						"short": "Unique identifiers of the categories this work is a part of",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "category_titles",
						"short": "Names of the categories this artwork is a part of",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "classification_id",
						"short": "Unique identifier of the preferred classification term for this work",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "classification_ids",
						"short": "Unique identifiers of all classification terms for this work",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "classification_title",
						"short": "The name of the preferred classification term for this work",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "classification_titles",
						"short": "The names of all classification terms related to this artwork",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "color",
						"short": "Dominant color of this artwork in HSL",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "colorfulness",
						"short": "Unbounded positive float representing an abstract measure of colorfulness.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "copyright_notice",
						"short": "Statement notifying how the work is protected by copyright.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "credit_line",
						"short": "Brief statement indicating how the work came into the collection",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "date_display",
						"short": "Readable, free-text description of the period of time associated with the creation of this work.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "date_end",
						"short": "The year of the period of time associated with the creation of this work",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "date_qualifier_id",
						"short": "Unique identifier of the qualifer to the dates provided for this record.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "date_qualifier_title",
						"short": "Readable, text qualifer to the dates provided for this record.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "date_start",
						"short": "The year of the period of time associated with the creation of this work",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "department_id",
						"short": "Unique identifier of the curatorial department that this work belongs to",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "department_title",
						"short": "Name of the curatorial department that this work belongs to",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "description",
						"short": "Longer explanation describing the work",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "dimensions",
						"short": "The size, shape, scale, and dimensions of the work.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "dimensions_detail",
						"short": "The height, width, depth, and/or diameter of each section of the work in centimeters",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "document_ids",
						"short": "Unique identifiers of assets that serve as documentation for this artwork",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "edition",
						"short": "Edition number if the work is one of many",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "exhibition_history",
						"short": "List of all the places this work has been exhibited",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "fiscal_year",
						"short": "The fiscal year in which the work was acquired.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "fiscal_year_deaccession",
						"short": "The fiscal year in which the work was deaccessioned.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "gallery_id",
						"short": "Unique identifier of the location of this work in our museum",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gallery_title",
						"short": "The location of this work in our museum",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "has_advanced_imaging",
						"short": "Whether this artwork is enhanced with 3D models, 360 image sequences, Mirador views, etc.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "has_educational_resources",
						"short": "Whether this artwork has any documents tagged as educational",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "has_multimedia_resources",
						"short": "Whether this artwork has any associated microsites, digital publications, or documents tagged as multimedia",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "has_not_been_viewed_much",
						"short": "Whether the artwork hasn't been visited on our website very much",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_embedding",
						"short": "The generated embeddings describing the artwork image",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "image_id",
						"short": "Unique identifier of the preferred image to use to represent this work",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inscriptions",
						"short": "A description of distinguishing or identifying physical markings that are on the work",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "internal_department_id",
						"short": "An internal department id we use for analytics.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "is_boosted",
						"short": "Whether this document should be boosted in search",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_on_view",
						"short": "Whether the work is on display",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_public_domain",
						"short": "Whether the work is in the public domain, meaning it was created before copyrights existed or has left the copyright term",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_zoomable",
						"short": "Whether images of the work are allowed to be displayed in a zoomable interface.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "latitude",
						"short": "Latitude coordinate of the location of this work in our galleries",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "latlon",
						"short": "Latitude and longitude coordinates of the location of this work in our galleries",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "longitude",
						"short": "Longitude coordinate of the location of this work in our galleries",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "main_reference_number",
						"short": "Unique identifier assigned to the artwork upon acquisition",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "material_id",
						"short": "Unique identifier of the preferred material term for this work",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "material_ids",
						"short": "Unique identifiers of all material terms for this work",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "material_titles",
						"short": "The names of all material terms related to this artwork",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "max_zoom_window_size",
						"short": "The maximum size of the window the image is allowed to be viewed in, in pixels.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "medium_display",
						"short": "The substances or materials used in the creation of a work",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "nomisma_id",
						"short": "Unique identifier of this work in the nomisma coin database",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "on_loan_display",
						"short": "If an artwork is on loan, this contains details about the loan",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "pageviews",
						"short": "Approx.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "pageviews_recent",
						"short": "Approx.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "place_of_origin",
						"short": "The location where the creation, design, or production of the work took place, or the original location of the work",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance_text",
						"short": "Ownership/collecting history of the work.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "publication_history",
						"short": "Bibliographic list of all the places this work has been published",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "publishing_verification_level",
						"short": "Indicator of how much metadata on the work in published.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "section_ids",
						"short": "Unique identifiers of the digital publication chapters this work in included in",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "section_titles",
						"short": "Names of the digital publication chapters this work is included in",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "short_description",
						"short": "Short explanation describing the work",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "site_ids",
						"short": "Unique identifiers of the microsites this work is a part of",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "sound_ids",
						"short": "Unique identifiers of the audio about this work",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "style_id",
						"short": "Unique identifier of the preferred style term for this work",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "style_ids",
						"short": "Unique identifiers of all style terms for this work",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "style_title",
						"short": "The name of the preferred style term for this work",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "style_titles",
						"short": "The names of all style terms related to this artwork",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "subject_id",
						"short": "Unique identifier of the preferred subject term for this work",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "subject_ids",
						"short": "Unique identifiers of all subject terms for this work",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "subject_titles",
						"short": "The names of all subject terms related to this artwork",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "technique_id",
						"short": "Unique identifier of the preferred technique term for this work",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "technique_ids",
						"short": "Unique identifiers of all technique terms for this work",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "technique_titles",
						"short": "The names of all technique terms related to this artwork",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "term_titles",
						"short": "The names of the taxonomy tags for this work",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "text_embedding",
						"short": "The generated embeddings of artwork text",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "text_ids",
						"short": "Unique identifiers of the texts about this work",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "theme_titles",
						"short": "The names of all thematic publish categories related to this artwork",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "thumbnail",
						"short": "Metadata about the image referenced by `image_id`.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "video_ids",
						"short": "Unique identifiers of the videos about this work",
						"type": "`$ANY`",
					},
				},
				"name": "artwork",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/artworks",
								"parts": []any{
									"artworks",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/artworks/{id}",
								"parts": []any{
									"artworks",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"artwork_date_qualifier": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
				},
				"name": "artwork_date_qualifier",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/artwork-date-qualifiers",
								"parts": []any{
									"artwork-date-qualifiers",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/artwork-date-qualifiers/{id}",
								"parts": []any{
									"artwork-date-qualifiers",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"artwork_place_qualifier": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
				},
				"name": "artwork_place_qualifier",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/artwork-place-qualifiers",
								"parts": []any{
									"artwork-place-qualifiers",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/artwork-place-qualifiers/{id}",
								"parts": []any{
									"artwork-place-qualifiers",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"artwork_type": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "aat_id",
						"short": "Identifier of reconciled (most similar) term in the Getty's Art and Architecture Thesaurus (AAT)",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
				},
				"name": "artwork_type",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/artwork-types",
								"parts": []any{
									"artwork-types",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/artwork-types/{id}",
								"parts": []any{
									"artwork-types",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"category_term": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "aat_id",
						"short": "Identifier of reconciled (most similar) term in the Getty's Art and Architecture Thesaurus (AAT)",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "parent_id",
						"short": "Unique identifier of this category's parent",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "subtype",
						"short": "Takes one of the following values: classification, material, technique, style, subject, department, theme",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
				},
				"name": "category_term",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/category-terms",
								"parts": []any{
									"category-terms",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/category-terms/{id}",
								"parts": []any{
									"category-terms",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"digital_publication": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "copy",
						"short": "The text of the page",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "web_url",
						"short": "The URL to this page on our website",
						"type": "`$ANY`",
					},
				},
				"name": "digital_publication",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/digital-publications",
								"parts": []any{
									"digital-publications",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/digital-publications/{id}",
								"parts": []any{
									"digital-publications",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"digital_publication_article": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "author_display",
						"short": "A display-friendly text of the authors of this article",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "copy",
						"short": "The text of the article",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "digital_publication_id",
						"short": "Unique identifier of the digital publication this article belongs to",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "web_url",
						"short": "The URL to this article on our website",
						"type": "`$ANY`",
					},
				},
				"name": "digital_publication_article",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/digital-publication-articles",
								"parts": []any{
									"digital-publication-articles",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/digital-publication-articles/{id}",
								"parts": []any{
									"digital-publication-articles",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"educator_resource": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "copy",
						"short": "The text of the page",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "web_url",
						"short": "The URL to this page on our website",
						"type": "`$ANY`",
					},
				},
				"name": "educator_resource",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/educator-resources",
								"parts": []any{
									"educator-resources",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/educator-resources/{id}",
								"parts": []any{
									"educator-resources",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"event": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "alt_audience_ids",
						"short": "Unique identifiers indicating the alternate audiences for this event",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "alt_event_type_ids",
						"short": "Unique identifiers indicating the alternate types of this event",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "audience_id",
						"short": "Unique identifier indicating the preferred audience for this event",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "buy_button_caption",
						"short": "Additional text below the ticket/registration button",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "buy_button_text",
						"short": "The text used on the ticket/registration button",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "date_display",
						"short": "A readable display of the event dates",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "description",
						"short": "All copytext of the event",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "door_time",
						"short": "The time the doors open for this event",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "end_date",
						"short": "The date the event ends",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "end_time",
						"short": "The time the event ends",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "entrance",
						"short": "Which entrance to use for this event",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "event_host_id",
						"short": "Unique identifier of the host (cf.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "event_host_title",
						"short": "Unique identifier of the host (cf.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "event_type_id",
						"short": "Unique identifier indicating the preferred type of this event",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "header_description",
						"short": "Brief description of the event displayed below the title",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "hero_caption",
						"short": "Text displayed with the hero image on the event",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_url",
						"short": "The URL of an image representing this page",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "is_admission_required",
						"short": "Whether admission to the museum is required to attend this event",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_after_hours",
						"short": "Whether the event is to be held after the museum closes",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_free",
						"short": "Whether the event is free",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_member_exclusive",
						"short": "Whether the event is exclusive to members of the museum",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_private",
						"short": "Whether the event is private",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_registration_required",
						"short": "Whether registration is required to attend the event",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_sales_button_hidden",
						"short": "Whether the buy tickets button is hidden on the website event page",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_sold_out",
						"short": "Whether the event is sold out",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_ticketed",
						"short": "Whether a ticket is required to attend the event",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_virtual_event",
						"short": "Whether the event is being held virtually",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "join_url",
						"short": "URL to the membership signup page via this event",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "layout_type",
						"short": "Number indicating the type of layout this event page uses",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "list_description",
						"short": "One-sentence description of the event displayed in listings",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "location",
						"short": "Where the event takes place",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "program_ids",
						"short": "Unique identifiers indicating the programs this event is a part of",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "program_titles",
						"short": "Titles of the programs this event is a part of",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "rsvp_link",
						"short": "The URL to the sales site for this event",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "search_tags",
						"short": "Editor-specified list of tags to aid in internal search",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "short_description",
						"short": "Brief description of the event",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "slug",
						"short": "A string used in the URL for this event",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "start_date",
						"short": "The date the event begins",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "start_time",
						"short": "The time the event starts",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "survey_url",
						"short": "URL to the survey associated with this event",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ticketed_event_id",
						"short": "Unique identifier of the event in the ticketing system this website event is tied to",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "title_display",
						"short": "The name of this event formatted with HTML (optional)",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "virtual_event_passcode",
						"short": "Passcode to access the virtual event",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "virtual_event_url",
						"short": "URL to the virtual event",
						"type": "`$ANY`",
					},
				},
				"name": "event",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/events",
								"parts": []any{
									"events",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/events/{id}",
								"parts": []any{
									"events",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"event_occurrence": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "button_caption",
						"short": "Additional text below the ticket/registration button",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "button_text",
						"short": "The text used on the ticket/registration button",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "button_url",
						"short": "The URL to the sales site or an RSVP link for this event",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "description",
						"short": "Description of the event",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "end_at",
						"short": "The date the event occurrence ends",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "event_id",
						"short": "Identifier of the master event of which this is an occurrence",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_url",
						"short": "The URL of an image representing this page",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "is_private",
						"short": "Whether the event is private.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_sales_button_hidden",
						"short": "Whether the buy tickets button is hidden on the website event page",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_ticketed",
						"short": "Whether a ticket is required to attend the event",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "location",
						"short": "Where the event takes place",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "off_sale_at",
						"short": "Date and time the event goes off sale",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "on_sale_at",
						"short": "Date and time the event goes on sale",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "short_description",
						"short": "Brief description of the event",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "start_at",
						"short": "The date the event occurrence begins",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "title_display",
						"short": "The name of this event formatted with HTML (optional)",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
				},
				"name": "event_occurrence",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/event-occurrences",
								"parts": []any{
									"event-occurrences",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/event-occurrences/{id}",
								"parts": []any{
									"event-occurrences",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"event_program": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "is_affiliate_group",
						"short": "Whether this program represents an affiliate group",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_event_host",
						"short": "Whether this program represents an event host",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
				},
				"name": "event_program",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/event-programs",
								"parts": []any{
									"event-programs",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/event-programs/{id}",
								"parts": []any{
									"event-programs",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"exhibition": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "aic_end_at",
						"short": "Date the exhibition closed at the Art Institute of Chicago",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "aic_start_at",
						"short": "Date the exhibition opened at the Art Institute of Chicago",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "alt_image_ids",
						"short": "Unique identifiers of all non-preferred images of this exhibition.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artist_ids",
						"short": "Unique identifiers of the artist agent records representing who was shown in the exhibition",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artwork_ids",
						"short": "Unique identifiers of the artworks that were part of the exhibition",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artwork_titles",
						"short": "Names of the artworks that were part of the exhibition",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "document_ids",
						"short": "Unique identifiers of assets that serve as documentation for this exhibition",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "gallery_id",
						"short": "Unique identifier of the gallery that mainly housed the exhibition",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gallery_title",
						"short": "The name of the gallery that mainly housed the exhibition",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_id",
						"short": "Unique identifier of the preferred image to use to represent this exhibition",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_url",
						"short": "URL to the hero image from the website",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "is_featured",
						"short": "Is this exhibition currently featured on our website?",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_published",
						"short": "Is this exhibition currently published on our website?",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "position",
						"short": "Numering position represnting the order in which this exhibition is featured on the website",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "short_description",
						"short": "Brief explanation of what this exhibition is",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "site_ids",
						"short": "Unique identifiers of the microsites this exhibition is a part of",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "status",
						"short": "Whether the exhibition is open or closed",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "web_url",
						"short": "URL to this exhibition on our website",
						"type": "`$ANY`",
					},
				},
				"name": "exhibition",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/exhibitions",
								"parts": []any{
									"exhibitions",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/exhibitions/{id}",
								"parts": []any{
									"exhibitions",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"gallery": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "floor",
						"short": "The level the gallery is on, e.g., 1, 2, 3, or LL",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "is_closed",
						"short": "Whether the gallery is currently closed",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "latitude",
						"short": "Latitude coordinate of the center of the room",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "latlon",
						"short": "Latitude and longitude coordinates of the center of the room",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "longitude",
						"short": "Longitude coordinate of the center of the room",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "number",
						"short": "The gallery's room number.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "tgn_id",
						"short": "Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN)",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
				},
				"name": "gallery",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/galleries",
								"parts": []any{
									"galleries",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/galleries/{id}",
								"parts": []any{
									"galleries",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"generic_page": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "copy",
						"short": "The text of the page",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "search_tags",
						"short": "Editor-specified list of tags to aid in internal search",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "web_url",
						"short": "The URL to this page on our website",
						"type": "`$ANY`",
					},
				},
				"name": "generic_page",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/generic-pages",
								"parts": []any{
									"generic-pages",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/generic-pages/{id}",
								"parts": []any{
									"generic-pages",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"highlight": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "copy",
						"short": "The text of the highlight description",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
				},
				"name": "highlight",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/highlights",
								"parts": []any{
									"highlights",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/highlights/{id}",
								"parts": []any{
									"highlights",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"hour": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "additional_text",
						"short": "Additional information about the hours",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "friday_is_closed",
						"short": "Whether the museum is closed on Fridays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "friday_member_close",
						"short": "The time member hours ends on Fridays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "friday_member_open",
						"short": "The time member hours starts on Fridays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "friday_public_close",
						"short": "The time public hours ends on Fridays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "friday_public_open",
						"short": "The time public hours starts on Fridays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "monday_is_closed",
						"short": "Whether the museum is closed on Mondays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "monday_member_close",
						"short": "The time member hours ends on Mondays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "monday_member_open",
						"short": "The time member hours starts on Mondays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "monday_public_close",
						"short": "The time public hours ends on Mondays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "monday_public_open",
						"short": "The time public hours starts on Mondays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "saturday_is_closed",
						"short": "Whether the museum is closed on Saturdays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "saturday_member_close",
						"short": "The time member hours ends on Saturdays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "saturday_member_open",
						"short": "The time member hours starts on Saturdays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "saturday_public_close",
						"short": "The time public hours ends on Saturdays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "saturday_public_open",
						"short": "The time public hours starts on Saturdays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "summary",
						"short": "Readable summary of the hours",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "sunday_is_closed",
						"short": "Whether the museum is closed on Sundays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "sunday_member_close",
						"short": "The time member hours ends on Sundays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "sunday_member_open",
						"short": "The time member hours starts on Sundays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "sunday_public_close",
						"short": "The time public hours ends on Sundays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "sunday_public_open",
						"short": "The time public hours starts on Sundays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "thursday_is_closed",
						"short": "Whether the museum is closed on Thursdays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "thursday_member_close",
						"short": "The time member hours ends on Thursdays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "thursday_member_open",
						"short": "The time member hours starts on Thursdays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "thursday_public_close",
						"short": "The time public hours ends on Thursdays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "thursday_public_open",
						"short": "The time public hours starts on Thursdays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tuesday_is_closed",
						"short": "Whether the museum is closed on Tuesdays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "tuesday_member_close",
						"short": "The time member hours ends on Tuesdays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "tuesday_member_open",
						"short": "The time member hours starts on Tuesdays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "tuesday_public_close",
						"short": "The time public hours ends on Tuesdays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "tuesday_public_open",
						"short": "The time public hours starts on Tuesdays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "wednesday_is_closed",
						"short": "Whether the museum is closed on Wednesdays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "wednesday_member_close",
						"short": "The time member hours ends on Wednesdays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "wednesday_member_open",
						"short": "The time member hours starts on Wednesdays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "wednesday_public_close",
						"short": "The time public hours ends on Wednesdays",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "wednesday_public_open",
						"short": "The time public hours starts on Wednesdays",
						"type": "`$ANY`",
					},
				},
				"name": "hour",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/hours",
								"parts": []any{
									"hours",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/hours/{id}",
								"parts": []any{
									"hours",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"image": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "ahash",
						"short": "Image hash generated using ahash algorithm with 64 boolean subfields",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "alt_text",
						"short": "Alternative text for the asset to describe it to people with low or no vision",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artwork_ids",
						"short": "Unique identifiers of the artworks associated with this asset",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artwork_titles",
						"short": "Names of the artworks associated with this asset",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "color",
						"short": "Dominant color of this image in HSL",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "colorfulness",
						"short": "Unbounded positive float representing an abstract measure of colorfulness.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "content",
						"short": "Text of or URL to the contents of this asset",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "content_e_tag",
						"short": "Arbitrary unique identifier that changes when the binary file gets updated",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "credit_line",
						"short": "Asset-specific copyright information",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "fingerprint",
						"short": "Image hashes: aHash, dHash, pHash, wHash",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "height",
						"short": "Native height of the image",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "iiif_url",
						"short": "IIIF URL of this image",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "is_educational_resource",
						"short": "Whether this resource is considered to be educational",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_multimedia_resource",
						"short": "Whether this resource is considered to be multimedia",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_teacher_resource",
						"short": "Whether this resource is considered to be educational",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "lake_guid",
						"short": "Unique UUID of this resource in LAKE, our DAMS.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "lqip",
						"short": "Low-quality image placeholder (LQIP).",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "phash",
						"short": "Image hash generated using phash algorithm with 64 boolean subfields",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "type",
						"short": "Type always takes one of the following values: image, sound, text, video",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "width",
						"short": "Native width of the image",
						"type": "`$NUMBER`",
					},
				},
				"name": "image",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/images",
								"parts": []any{
									"images",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/images/{id}",
								"parts": []any{
									"images",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"landing_page": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "copy",
						"short": "The text of the page",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "search_tags",
						"short": "Editor-specified list of tags to aid in internal search",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "web_url",
						"short": "The URL to this page on our website",
						"type": "`$ANY`",
					},
				},
				"name": "landing_page",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/landing-pages",
								"parts": []any{
									"landing-pages",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/landing-pages/{id}",
								"parts": []any{
									"landing-pages",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"place": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "latitude",
						"short": "Latitude coordinate of the center of the room",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "longitude",
						"short": "Longitude coordinate of the center of the room",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "tgn_id",
						"short": "Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN)",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
				},
				"name": "place",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/places",
								"parts": []any{
									"places",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/places/{id}",
								"parts": []any{
									"places",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"press_release": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "copy",
						"short": "The text of the page",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "web_url",
						"short": "The URL to this page on our website",
						"type": "`$ANY`",
					},
				},
				"name": "press_release",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/press-releases",
								"parts": []any{
									"press-releases",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/press-releases/{id}",
								"parts": []any{
									"press-releases",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"printed_publication": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "copy",
						"short": "The text of the page",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "web_url",
						"short": "The URL to this page on our website",
						"type": "`$ANY`",
					},
				},
				"name": "printed_publication",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/printed-publications",
								"parts": []any{
									"printed-publications",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/printed-publications/{id}",
								"parts": []any{
									"printed-publications",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"product": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artist_ids",
						"short": "Unique identifiers of the artists associated with this product",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artwork_ids",
						"short": "Unique identifiers of the artworks associated with this product",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "description",
						"short": "Explanation of what this product is",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "exhibition_ids",
						"short": "Unique identifiers of the exhibitions associated with this product",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "external_sku",
						"short": "Numeric product identification code of a machine-readable barcode, when the customer sku differs from our internal one",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_url",
						"short": "URL of an image for this product",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "max_compare_at_price",
						"short": "Number indicating how much the most expensive variant of a product cost before a sale",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "max_current_price",
						"short": "Number indicating how much the most expensive variant of a product costs right now",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "min_compare_at_price",
						"short": "Number indicating how much the least expensive variant of a product cost before a sale",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "min_current_price",
						"short": "Number indicating how much the least expensive variant of a product costs right now",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "price_display",
						"short": "Explanation of what this product is",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "web_url",
						"short": "URL of this product in the shop",
						"type": "`$ANY`",
					},
				},
				"name": "product",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/products",
								"parts": []any{
									"products",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/products/{id}",
								"parts": []any{
									"products",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"publication": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "section_ids",
						"short": "Unique identifiers of the sections of this publication",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "web_url",
						"short": "URL to the publication",
						"type": "`$ANY`",
					},
				},
				"name": "publication",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/publications",
								"parts": []any{
									"publications",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/publications/{id}",
								"parts": []any{
									"publications",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"search": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_id",
						"short": "API unique identifier",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "api_link",
						"short": "URL to this recource in the API",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "Name of the model the resource represents",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier within the search index",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "is_boosted",
						"short": "Whether this record has been flagged to be boosted",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "score",
						"short": "Search index ranking of the result",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "thumbnail",
						"short": "Metadata on the image representing this record",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date this record was last updated in the API",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
				},
				"name": "search",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/agents/search",
								"parts": []any{
									"agents",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/articles/search",
								"parts": []any{
									"articles",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/artworks/search",
								"parts": []any{
									"artworks",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/category-terms/search",
								"parts": []any{
									"category-terms",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/digital-publication-articles/search",
								"parts": []any{
									"digital-publication-articles",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/digital-publications/search",
								"parts": []any{
									"digital-publications",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/educator-resources/search",
								"parts": []any{
									"educator-resources",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/event-occurrences/search",
								"parts": []any{
									"event-occurrences",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/event-programs/search",
								"parts": []any{
									"event-programs",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/events/search",
								"parts": []any{
									"events",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/exhibitions/search",
								"parts": []any{
									"exhibitions",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/galleries/search",
								"parts": []any{
									"galleries",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/generic-pages/search",
								"parts": []any{
									"generic-pages",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/highlights/search",
								"parts": []any{
									"highlights",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/hours/search",
								"parts": []any{
									"hours",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/images/search",
								"parts": []any{
									"images",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/landing-pages/search",
								"parts": []any{
									"landing-pages",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/mobile-sounds/search",
								"parts": []any{
									"mobile-sounds",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/places/search",
								"parts": []any{
									"places",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/press-releases/search",
								"parts": []any{
									"press-releases",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/printed-publications/search",
								"parts": []any{
									"printed-publications",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/products/search",
								"parts": []any{
									"products",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/publications/search",
								"parts": []any{
									"publications",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/search",
								"parts": []any{
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/sections/search",
								"parts": []any{
									"sections",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/sites/search",
								"parts": []any{
									"sites",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/sounds/search",
								"parts": []any{
									"sounds",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/static-pages/search",
								"parts": []any{
									"static-pages",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/texts/search",
								"parts": []any{
									"texts",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/tours/search",
								"parts": []any{
									"tours",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/videos/search",
								"parts": []any{
									"videos",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"section": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "accession",
						"short": "An accession number parsed from the title or tombstone",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artwork_id",
						"short": "Unique identifier of the artwork with which this section is associated",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "content",
						"short": "Content of this section in plaintext",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "generic_page_id",
						"short": "Unique identifier of the page on the website that represents the publication this section belongs to",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "publication_id",
						"short": "Unique identifier of the publication this section belongs to",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "publication_title",
						"short": "Name of the publication this section belongs to",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "web_url",
						"short": "URL to the section",
						"type": "`$ANY`",
					},
				},
				"name": "section",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/sections",
								"parts": []any{
									"sections",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/sections/{id}",
								"parts": []any{
									"sections",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"site": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artwork_ids",
						"short": "Unique identifiers of the artworks this site is associated with",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artwork_titles",
						"short": "Names of the artworks this site is associated with",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "description",
						"short": "Explanation of what this site is",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "exhibition_ids",
						"short": "Unique identifier of the exhibitions this site is associated with",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "exhibition_titles",
						"short": "Names of the exhibitions this site is associated with",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "web_url",
						"short": "URL to this site",
						"type": "`$ANY`",
					},
				},
				"name": "site",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/sites",
								"parts": []any{
									"sites",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/sites/{id}",
								"parts": []any{
									"sites",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"sound": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "alt_text",
						"short": "Alternative text for the asset to describe it to people with low or no vision",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artwork_ids",
						"short": "Unique identifiers of the artworks associated with this asset",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artwork_titles",
						"short": "Names of the artworks associated with this asset",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "content",
						"short": "Text of or URL to the contents of this asset",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "content_e_tag",
						"short": "Arbitrary unique identifier that changes when the binary file gets updated",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "credit_line",
						"short": "Asset-specific copyright information",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "is_educational_resource",
						"short": "Whether this resource is considered to be educational",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_multimedia_resource",
						"short": "Whether this resource is considered to be multimedia",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_teacher_resource",
						"short": "Whether this resource is considered to be educational",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "lake_guid",
						"short": "Unique UUID of this resource in LAKE, our DAMS.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "Name of this mobile audio file – derived from the artwork and tour titles",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "transcript",
						"short": "Text transcription of the audio file",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "type",
						"short": "Type always takes one of the following values: image, sound, text, video",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "web_url",
						"short": "URL to the audio file",
						"type": "`$ANY`",
					},
				},
				"name": "sound",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/mobile-sounds",
								"parts": []any{
									"mobile-sounds",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/sounds",
								"parts": []any{
									"sounds",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/mobile-sounds/{id}",
								"parts": []any{
									"mobile-sounds",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/sounds/{id}",
								"parts": []any{
									"sounds",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"static_page": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "web_url",
						"short": "The URL to this page on our website",
						"type": "`$ANY`",
					},
				},
				"name": "static_page",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/static-pages",
								"parts": []any{
									"static-pages",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/static-pages/{id}",
								"parts": []any{
									"static-pages",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"text": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "alt_text",
						"short": "Alternative text for the asset to describe it to people with low or no vision",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artwork_ids",
						"short": "Unique identifiers of the artworks associated with this asset",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artwork_titles",
						"short": "Names of the artworks associated with this asset",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "content",
						"short": "Text of or URL to the contents of this asset",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "content_e_tag",
						"short": "Arbitrary unique identifier that changes when the binary file gets updated",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "credit_line",
						"short": "Asset-specific copyright information",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "is_educational_resource",
						"short": "Whether this resource is considered to be educational",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_multimedia_resource",
						"short": "Whether this resource is considered to be multimedia",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_teacher_resource",
						"short": "Whether this resource is considered to be educational",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "lake_guid",
						"short": "Unique UUID of this resource in LAKE, our DAMS.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "type",
						"short": "Type always takes one of the following values: image, sound, text, video",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
				},
				"name": "text",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/texts",
								"parts": []any{
									"texts",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/texts/{id}",
								"parts": []any{
									"texts",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"tour": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artist_titles",
						"short": "Names of the artists of the artworks featured in this tour's tour stops",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artwork_titles",
						"short": "Names of the artworks featured in this tour's tour stops",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "description",
						"short": "Explanation of what the tour is",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image",
						"short": "The main image for the tour",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "intro",
						"short": "Text introducing the tour",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "intro_link",
						"short": "Link to the audio file of the introduction",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "intro_transcript",
						"short": "Transcript of the introduction audio to the tour",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "weight",
						"short": "Number representing this tour's sort order",
						"type": "`$NUMBER`",
					},
				},
				"name": "tour",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/tours",
								"parts": []any{
									"tours",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/tours/{id}",
								"parts": []any{
									"tours",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"video": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "alt_text",
						"short": "Alternative text for the asset to describe it to people with low or no vision",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_link",
						"short": "REST API link for this resource",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"short": "REST API resource type or endpoint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artwork_ids",
						"short": "Unique identifiers of the artworks associated with this asset",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artwork_titles",
						"short": "Names of the artworks associated with this asset",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "content",
						"short": "Text of or URL to the contents of this asset",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "content_e_tag",
						"short": "Arbitrary unique identifier that changes when the binary file gets updated",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "credit_line",
						"short": "Asset-specific copyright information",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier of this resource.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "is_educational_resource",
						"short": "Whether this resource is considered to be educational",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_multimedia_resource",
						"short": "Whether this resource is considered to be multimedia",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_teacher_resource",
						"short": "Whether this resource is considered to be educational",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "lake_guid",
						"short": "Unique UUID of this resource in LAKE, our DAMS.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "source_updated_at",
						"short": "Date and time the resource was updated in the source system",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"short": "Internal field to power the `/autosuggest` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"short": "Internal field to power the `/autocomplete` endpoint.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"short": "Date and time the record was updated in the aggregator search index",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"short": "The name of this resource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "type",
						"short": "Type always takes one of the following values: image, sound, text, video",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Date and time the record was updated in the aggregator database",
						"type": "`$ANY`",
					},
				},
				"name": "video",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/videos",
								"parts": []any{
									"videos",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/videos/{id}",
								"parts": []any{
									"videos",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
