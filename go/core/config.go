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
				"transport": "base",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "agents",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"agents",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/artists",
								"segments": []any{
									map[string]any{
										"lit": "artists",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"artists",
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
								"segments": []any{
									map[string]any{
										"lit": "agents",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"agents",
									"{id}",
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
								"segments": []any{
									map[string]any{
										"lit": "artists",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"artists",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "agent-roles",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"agent-roles",
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
								"segments": []any{
									map[string]any{
										"lit": "agent-roles",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"agent-roles",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "agent-types",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"agent-types",
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
								"segments": []any{
									map[string]any{
										"lit": "agent-types",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"agent-types",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "articles",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"articles",
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
								"segments": []any{
									map[string]any{
										"lit": "articles",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"articles",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "artworks",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"artworks",
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
								"segments": []any{
									map[string]any{
										"lit": "artworks",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"artworks",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "artwork-date-qualifiers",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"artwork-date-qualifiers",
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
								"segments": []any{
									map[string]any{
										"lit": "artwork-date-qualifiers",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"artwork-date-qualifiers",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "artwork-place-qualifiers",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"artwork-place-qualifiers",
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
								"segments": []any{
									map[string]any{
										"lit": "artwork-place-qualifiers",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"artwork-place-qualifiers",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "artwork-types",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"artwork-types",
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
								"segments": []any{
									map[string]any{
										"lit": "artwork-types",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"artwork-types",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "category-terms",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"category-terms",
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
								"segments": []any{
									map[string]any{
										"lit": "category-terms",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"category-terms",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "digital-publications",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"digital-publications",
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
								"segments": []any{
									map[string]any{
										"lit": "digital-publications",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"digital-publications",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "digital-publication-articles",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"digital-publication-articles",
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
								"segments": []any{
									map[string]any{
										"lit": "digital-publication-articles",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"digital-publication-articles",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "educator-resources",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"educator-resources",
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
								"segments": []any{
									map[string]any{
										"lit": "educator-resources",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"educator-resources",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "events",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"events",
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
								"segments": []any{
									map[string]any{
										"lit": "events",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"events",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "event-occurrences",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"event-occurrences",
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
								"segments": []any{
									map[string]any{
										"lit": "event-occurrences",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"event-occurrences",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "event-programs",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"event-programs",
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
								"segments": []any{
									map[string]any{
										"lit": "event-programs",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"event-programs",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "exhibitions",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"exhibitions",
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
								"segments": []any{
									map[string]any{
										"lit": "exhibitions",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"exhibitions",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "galleries",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"galleries",
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
								"segments": []any{
									map[string]any{
										"lit": "galleries",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"galleries",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "generic-pages",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"generic-pages",
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
								"segments": []any{
									map[string]any{
										"lit": "generic-pages",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"generic-pages",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "highlights",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"highlights",
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
								"segments": []any{
									map[string]any{
										"lit": "highlights",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"highlights",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "hours",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"hours",
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
								"segments": []any{
									map[string]any{
										"lit": "hours",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"hours",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "images",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"images",
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
								"segments": []any{
									map[string]any{
										"lit": "images",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"images",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "landing-pages",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"landing-pages",
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
								"segments": []any{
									map[string]any{
										"lit": "landing-pages",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"landing-pages",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "places",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"places",
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
								"segments": []any{
									map[string]any{
										"lit": "places",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"places",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "press-releases",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"press-releases",
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
								"segments": []any{
									map[string]any{
										"lit": "press-releases",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"press-releases",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "printed-publications",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"printed-publications",
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
								"segments": []any{
									map[string]any{
										"lit": "printed-publications",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"printed-publications",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "products",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"products",
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
								"segments": []any{
									map[string]any{
										"lit": "products",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"products",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "publications",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"publications",
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
								"segments": []any{
									map[string]any{
										"lit": "publications",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"publications",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "agents",
									},
									map[string]any{
										"lit": "search",
									},
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
								"parts": []any{
									"agents",
									"search",
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
								"segments": []any{
									map[string]any{
										"lit": "articles",
									},
									map[string]any{
										"lit": "search",
									},
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
								"parts": []any{
									"articles",
									"search",
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
								"segments": []any{
									map[string]any{
										"lit": "artworks",
									},
									map[string]any{
										"lit": "search",
									},
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
								"parts": []any{
									"artworks",
									"search",
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
								"segments": []any{
									map[string]any{
										"lit": "category-terms",
									},
									map[string]any{
										"lit": "search",
									},
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
								"parts": []any{
									"category-terms",
									"search",
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
								"segments": []any{
									map[string]any{
										"lit": "digital-publication-articles",
									},
									map[string]any{
										"lit": "search",
									},
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
								"parts": []any{
									"digital-publication-articles",
									"search",
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
								"segments": []any{
									map[string]any{
										"lit": "digital-publications",
									},
									map[string]any{
										"lit": "search",
									},
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
								"parts": []any{
									"digital-publications",
									"search",
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
								"segments": []any{
									map[string]any{
										"lit": "educator-resources",
									},
									map[string]any{
										"lit": "search",
									},
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
								"parts": []any{
									"educator-resources",
									"search",
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
								"segments": []any{
									map[string]any{
										"lit": "event-occurrences",
									},
									map[string]any{
										"lit": "search",
									},
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
								"parts": []any{
									"event-occurrences",
									"search",
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
								"segments": []any{
									map[string]any{
										"lit": "event-programs",
									},
									map[string]any{
										"lit": "search",
									},
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
								"parts": []any{
									"event-programs",
									"search",
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
								"segments": []any{
									map[string]any{
										"lit": "events",
									},
									map[string]any{
										"lit": "search",
									},
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
								"parts": []any{
									"events",
									"search",
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
								"segments": []any{
									map[string]any{
										"lit": "exhibitions",
									},
									map[string]any{
										"lit": "search",
									},
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
								"parts": []any{
									"exhibitions",
									"search",
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
								"segments": []any{
									map[string]any{
										"lit": "galleries",
									},
									map[string]any{
										"lit": "search",
									},
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
								"parts": []any{
									"galleries",
									"search",
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
								"segments": []any{
									map[string]any{
										"lit": "generic-pages",
									},
									map[string]any{
										"lit": "search",
									},
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
								"parts": []any{
									"generic-pages",
									"search",
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
								"segments": []any{
									map[string]any{
										"lit": "highlights",
									},
									map[string]any{
										"lit": "search",
									},
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
								"parts": []any{
									"highlights",
									"search",
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
								"segments": []any{
									map[string]any{
										"lit": "hours",
									},
									map[string]any{
										"lit": "search",
									},
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
								"parts": []any{
									"hours",
									"search",
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
								"segments": []any{
									map[string]any{
										"lit": "images",
									},
									map[string]any{
										"lit": "search",
									},
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
								"parts": []any{
									"images",
									"search",
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
								"segments": []any{
									map[string]any{
										"lit": "landing-pages",
									},
									map[string]any{
										"lit": "search",
									},
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
								"parts": []any{
									"landing-pages",
									"search",
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
								"segments": []any{
									map[string]any{
										"lit": "mobile-sounds",
									},
									map[string]any{
										"lit": "search",
									},
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
								"parts": []any{
									"mobile-sounds",
									"search",
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
								"segments": []any{
									map[string]any{
										"lit": "places",
									},
									map[string]any{
										"lit": "search",
									},
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
								"parts": []any{
									"places",
									"search",
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
								"segments": []any{
									map[string]any{
										"lit": "press-releases",
									},
									map[string]any{
										"lit": "search",
									},
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
								"parts": []any{
									"press-releases",
									"search",
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
								"segments": []any{
									map[string]any{
										"lit": "printed-publications",
									},
									map[string]any{
										"lit": "search",
									},
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
								"parts": []any{
									"printed-publications",
									"search",
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
								"segments": []any{
									map[string]any{
										"lit": "products",
									},
									map[string]any{
										"lit": "search",
									},
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
								"parts": []any{
									"products",
									"search",
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
								"segments": []any{
									map[string]any{
										"lit": "publications",
									},
									map[string]any{
										"lit": "search",
									},
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
								"parts": []any{
									"publications",
									"search",
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
								"segments": []any{
									map[string]any{
										"lit": "search",
									},
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
								"parts": []any{
									"search",
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
								"segments": []any{
									map[string]any{
										"lit": "sections",
									},
									map[string]any{
										"lit": "search",
									},
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
								"parts": []any{
									"sections",
									"search",
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
								"segments": []any{
									map[string]any{
										"lit": "sites",
									},
									map[string]any{
										"lit": "search",
									},
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
								"parts": []any{
									"sites",
									"search",
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
								"segments": []any{
									map[string]any{
										"lit": "sounds",
									},
									map[string]any{
										"lit": "search",
									},
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
								"parts": []any{
									"sounds",
									"search",
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
								"segments": []any{
									map[string]any{
										"lit": "static-pages",
									},
									map[string]any{
										"lit": "search",
									},
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
								"parts": []any{
									"static-pages",
									"search",
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
								"segments": []any{
									map[string]any{
										"lit": "texts",
									},
									map[string]any{
										"lit": "search",
									},
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
								"parts": []any{
									"texts",
									"search",
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
								"segments": []any{
									map[string]any{
										"lit": "tours",
									},
									map[string]any{
										"lit": "search",
									},
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
								"parts": []any{
									"tours",
									"search",
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
								"segments": []any{
									map[string]any{
										"lit": "videos",
									},
									map[string]any{
										"lit": "search",
									},
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
								"parts": []any{
									"videos",
									"search",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "sections",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"sections",
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
								"segments": []any{
									map[string]any{
										"lit": "sections",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"sections",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "sites",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"sites",
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
								"segments": []any{
									map[string]any{
										"lit": "sites",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"sites",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "mobile-sounds",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"mobile-sounds",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/sounds",
								"segments": []any{
									map[string]any{
										"lit": "sounds",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"sounds",
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
								"segments": []any{
									map[string]any{
										"lit": "mobile-sounds",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"mobile-sounds",
									"{id}",
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
								"segments": []any{
									map[string]any{
										"lit": "sounds",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"sounds",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "static-pages",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"static-pages",
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
								"segments": []any{
									map[string]any{
										"lit": "static-pages",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"static-pages",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "texts",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"texts",
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
								"segments": []any{
									map[string]any{
										"lit": "texts",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"texts",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "tours",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"tours",
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
								"segments": []any{
									map[string]any{
										"lit": "tours",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"tours",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "videos",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"videos",
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
								"segments": []any{
									map[string]any{
										"lit": "videos",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"videos",
									"{id}",
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

// The plugin definitions the model selected per feature, as []any so a
// feature package can consume them without core naming its types. Empty
// when no active feature declares active plugin groups for this target.
var featurePlugins = map[string][]any{
}

// FeaturePlugins is the definitions list for one feature's chain.
func FeaturePlugins(name string) []any {
	return featurePlugins[name]
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
