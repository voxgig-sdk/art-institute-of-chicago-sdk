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
			"ratelimit": map[string]any{
				"options": map[string]any{
					"active": false,
					"burst": 5,
					"rate": 5,
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"retry": map[string]any{
				"options": map[string]any{
					"active": false,
					"factor": 2,
					"maxDelay": 2000,
					"minDelay": 50,
					"retries": 2,
					"statuses": []any{
						408,
						425,
						429,
						500,
						502,
						503,
						504,
					},
				},
				"optspec": map[string]any{
					"jitter": "`$BOOLEAN`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"optspec": map[string]any{
					"entity": "`$MAP`",
					"net": "`$MAP`",
				},
				"strict": false,
				"transport": "base",
			},
			"timeout": map[string]any{
				"options": map[string]any{
					"active": false,
					"ms": 30000,
				},
				"optspec": map[string]any{
					"clearTimer": "`$FUNCTION`",
					"setTimer": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
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
						"title": "Alt Titles",
						"type": "`$ANY`",
						"short": "Alternate names for this agent",
					},
					map[string]any{
						"name": "api_link",
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "birth_date",
						"title": "Birth Date",
						"type": "`$ANY`",
						"short": "The year this agent was born",
					},
					map[string]any{
						"name": "death_date",
						"title": "Death Date",
						"type": "`$ANY`",
						"short": "The year this agent died",
					},
					map[string]any{
						"name": "description",
						"title": "Description",
						"type": "`$STRING`",
						"short": "A biographical description of the agent",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "is_artist",
						"title": "Is Artist",
						"type": "`$BOOLEAN`",
						"short": "Whether the agent is an artist.",
					},
					map[string]any{
						"name": "sort_title",
						"title": "Sort Title",
						"type": "`$ANY`",
						"short": "Sortable name for this agent, typically with last name first.",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "ulan_id",
						"title": "Ulan Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this agent in Getty's ULAN",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
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
								"kind": "http",
								"method": "GET",
								"orig": "/agents",
								"segments": []any{
									map[string]any{
										"lit": "agents",
									},
								},
								"parts": []any{
									"agents",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/artists",
								"segments": []any{
									map[string]any{
										"lit": "artists",
									},
								},
								"parts": []any{
									"artists",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"agents",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
							},
							map[string]any{
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
								"parts": []any{
									"artists",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
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
								"kind": "http",
								"method": "GET",
								"orig": "/agent-roles",
								"segments": []any{
									map[string]any{
										"lit": "agent-roles",
									},
								},
								"parts": []any{
									"agent-roles",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"agent-roles",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
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
								"kind": "http",
								"method": "GET",
								"orig": "/agent-types",
								"segments": []any{
									map[string]any{
										"lit": "agent-types",
									},
								},
								"parts": []any{
									"agent-types",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"agent-types",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "copy",
						"title": "Copy",
						"type": "`$ANY`",
						"short": "The text of the article",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
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
								"kind": "http",
								"method": "GET",
								"orig": "/articles",
								"segments": []any{
									map[string]any{
										"lit": "articles",
									},
								},
								"parts": []any{
									"articles",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"articles",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Alt Artist Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of the non-preferred artists/cultures associated with this work",
					},
					map[string]any{
						"name": "alt_classification_ids",
						"title": "Alt Classification Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of all other non-preferred classification terms for this work",
					},
					map[string]any{
						"name": "alt_image_ids",
						"title": "Alt Image Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of all non-preferred images of this work.",
					},
					map[string]any{
						"name": "alt_material_ids",
						"title": "Alt Material Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of all other non-preferred material terms for this work",
					},
					map[string]any{
						"name": "alt_style_ids",
						"title": "Alt Style Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of all other non-preferred style terms for this work",
					},
					map[string]any{
						"name": "alt_subject_ids",
						"title": "Alt Subject Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of all other non-preferred subject terms for this work",
					},
					map[string]any{
						"name": "alt_technique_ids",
						"title": "Alt Technique Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of all other non-preferred technique terms for this work",
					},
					map[string]any{
						"name": "alt_titles",
						"title": "Alt Titles",
						"type": "`$ANY`",
						"short": "Alternate names for this work",
					},
					map[string]any{
						"name": "api_link",
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "artist_display",
						"title": "Artist Display",
						"type": "`$ANY`",
						"short": "Readable description of the creator of this work.",
					},
					map[string]any{
						"name": "artist_id",
						"title": "Artist Id",
						"type": "`$STRING`",
						"short": "Unique identifier of the preferred artist/culture associated with this work",
					},
					map[string]any{
						"name": "artist_ids",
						"title": "Artist Ids",
						"type": "`$ANY`",
						"short": "Unique identifier of all artist/cultures associated with this work",
					},
					map[string]any{
						"name": "artist_title",
						"title": "Artist Title",
						"type": "`$ANY`",
						"short": "Name of the preferred artist/culture associated with this work",
					},
					map[string]any{
						"name": "artist_titles",
						"title": "Artist Titles",
						"type": "`$ANY`",
						"short": "Names of all artist/cultures associated with this work",
					},
					map[string]any{
						"name": "artwork_type_id",
						"title": "Artwork Type Id",
						"type": "`$STRING`",
						"short": "Unique identifier of the kind of object or work",
					},
					map[string]any{
						"name": "artwork_type_title",
						"title": "Artwork Type Title",
						"type": "`$ANY`",
						"short": "The kind of object or work (e.g.",
					},
					map[string]any{
						"name": "boost_rank",
						"title": "Boost Rank",
						"type": "`$ANY`",
						"short": "Manual indication of what rank this artwork should take in search results.",
					},
					map[string]any{
						"name": "catalog_based_search_keyword_titles",
						"title": "Catalog Based Search Keyword Titles",
						"type": "`$ANY`",
						"short": "The keyword search values that would be catalog-based searches on this record",
					},
					map[string]any{
						"name": "catalogue_display",
						"title": "Catalogue Display",
						"type": "`$ANY`",
						"short": "Brief text listing all the catalogues raisonnés which include this work.",
					},
					map[string]any{
						"name": "category_ids",
						"title": "Category Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of the categories this work is a part of",
					},
					map[string]any{
						"name": "category_titles",
						"title": "Category Titles",
						"type": "`$ANY`",
						"short": "Names of the categories this artwork is a part of",
					},
					map[string]any{
						"name": "classification_id",
						"title": "Classification Id",
						"type": "`$STRING`",
						"short": "Unique identifier of the preferred classification term for this work",
					},
					map[string]any{
						"name": "classification_ids",
						"title": "Classification Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of all classification terms for this work",
					},
					map[string]any{
						"name": "classification_title",
						"title": "Classification Title",
						"type": "`$ANY`",
						"short": "The name of the preferred classification term for this work",
					},
					map[string]any{
						"name": "classification_titles",
						"title": "Classification Titles",
						"type": "`$ANY`",
						"short": "The names of all classification terms related to this artwork",
					},
					map[string]any{
						"name": "color",
						"title": "Color",
						"type": "`$ANY`",
						"short": "Dominant color of this artwork in HSL",
					},
					map[string]any{
						"name": "colorfulness",
						"title": "Colorfulness",
						"type": "`$ANY`",
						"short": "Unbounded positive float representing an abstract measure of colorfulness.",
					},
					map[string]any{
						"name": "copyright_notice",
						"title": "Copyright Notice",
						"type": "`$ANY`",
						"short": "Statement notifying how the work is protected by copyright.",
					},
					map[string]any{
						"name": "credit_line",
						"title": "Credit Line",
						"type": "`$ANY`",
						"short": "Brief statement indicating how the work came into the collection",
					},
					map[string]any{
						"name": "date_display",
						"title": "Date Display",
						"type": "`$ANY`",
						"short": "Readable, free-text description of the period of time associated with the creation of this work.",
					},
					map[string]any{
						"name": "date_end",
						"title": "Date End",
						"type": "`$ANY`",
						"short": "The year of the period of time associated with the creation of this work",
					},
					map[string]any{
						"name": "date_qualifier_id",
						"title": "Date Qualifier Id",
						"type": "`$STRING`",
						"short": "Unique identifier of the qualifer to the dates provided for this record.",
					},
					map[string]any{
						"name": "date_qualifier_title",
						"title": "Date Qualifier Title",
						"type": "`$ANY`",
						"short": "Readable, text qualifer to the dates provided for this record.",
					},
					map[string]any{
						"name": "date_start",
						"title": "Date Start",
						"type": "`$ANY`",
						"short": "The year of the period of time associated with the creation of this work",
					},
					map[string]any{
						"name": "department_id",
						"title": "Department Id",
						"type": "`$STRING`",
						"short": "Unique identifier of the curatorial department that this work belongs to",
					},
					map[string]any{
						"name": "department_title",
						"title": "Department Title",
						"type": "`$ANY`",
						"short": "Name of the curatorial department that this work belongs to",
					},
					map[string]any{
						"name": "description",
						"title": "Description",
						"type": "`$STRING`",
						"short": "Longer explanation describing the work",
					},
					map[string]any{
						"name": "dimensions",
						"title": "Dimensions",
						"type": "`$ANY`",
						"short": "The size, shape, scale, and dimensions of the work.",
					},
					map[string]any{
						"name": "dimensions_detail",
						"title": "Dimensions Detail",
						"type": "`$ANY`",
						"short": "The height, width, depth, and/or diameter of each section of the work in centimeters",
					},
					map[string]any{
						"name": "document_ids",
						"title": "Document Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of assets that serve as documentation for this artwork",
					},
					map[string]any{
						"name": "edition",
						"title": "Edition",
						"type": "`$ANY`",
						"short": "Edition number if the work is one of many",
					},
					map[string]any{
						"name": "exhibition_history",
						"title": "Exhibition History",
						"type": "`$ANY`",
						"short": "List of all the places this work has been exhibited",
					},
					map[string]any{
						"name": "fiscal_year",
						"title": "Fiscal Year",
						"type": "`$ANY`",
						"short": "The fiscal year in which the work was acquired.",
					},
					map[string]any{
						"name": "fiscal_year_deaccession",
						"title": "Fiscal Year Deaccession",
						"type": "`$ANY`",
						"short": "The fiscal year in which the work was deaccessioned.",
					},
					map[string]any{
						"name": "gallery_id",
						"title": "Gallery Id",
						"type": "`$STRING`",
						"short": "Unique identifier of the location of this work in our museum",
					},
					map[string]any{
						"name": "gallery_title",
						"title": "Gallery Title",
						"type": "`$ANY`",
						"short": "The location of this work in our museum",
					},
					map[string]any{
						"name": "has_advanced_imaging",
						"title": "Has Advanced Imaging",
						"type": "`$BOOLEAN`",
						"short": "Whether this artwork is enhanced with 3D models, 360 image sequences, Mirador views, etc.",
					},
					map[string]any{
						"name": "has_educational_resources",
						"title": "Has Educational Resources",
						"type": "`$BOOLEAN`",
						"short": "Whether this artwork has any documents tagged as educational",
					},
					map[string]any{
						"name": "has_multimedia_resources",
						"title": "Has Multimedia Resources",
						"type": "`$BOOLEAN`",
						"short": "Whether this artwork has any associated microsites, digital publications, or documents tagged as multimedia",
					},
					map[string]any{
						"name": "has_not_been_viewed_much",
						"title": "Has Not Been Viewed Much",
						"type": "`$BOOLEAN`",
						"short": "Whether the artwork hasn't been visited on our website very much",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "image_embedding",
						"title": "Image Embedding",
						"type": "`$ANY`",
						"short": "The generated embeddings describing the artwork image",
					},
					map[string]any{
						"name": "image_id",
						"title": "Image Id",
						"type": "`$STRING`",
						"short": "Unique identifier of the preferred image to use to represent this work",
					},
					map[string]any{
						"name": "inscriptions",
						"title": "Inscriptions",
						"type": "`$ANY`",
						"short": "A description of distinguishing or identifying physical markings that are on the work",
					},
					map[string]any{
						"name": "internal_department_id",
						"title": "Internal Department Id",
						"type": "`$STRING`",
						"short": "An internal department id we use for analytics.",
					},
					map[string]any{
						"name": "is_boosted",
						"title": "Is Boosted",
						"type": "`$BOOLEAN`",
						"short": "Whether this document should be boosted in search",
					},
					map[string]any{
						"name": "is_on_view",
						"title": "Is On View",
						"type": "`$BOOLEAN`",
						"short": "Whether the work is on display",
					},
					map[string]any{
						"name": "is_public_domain",
						"title": "Is Public Domain",
						"type": "`$BOOLEAN`",
						"short": "Whether the work is in the public domain, meaning it was created before copyrights existed or has left the copyright term",
					},
					map[string]any{
						"name": "is_zoomable",
						"title": "Is Zoomable",
						"type": "`$BOOLEAN`",
						"short": "Whether images of the work are allowed to be displayed in a zoomable interface.",
					},
					map[string]any{
						"name": "latitude",
						"title": "Latitude",
						"type": "`$NUMBER`",
						"short": "Latitude coordinate of the location of this work in our galleries",
					},
					map[string]any{
						"name": "latlon",
						"title": "Latlon",
						"type": "`$ANY`",
						"short": "Latitude and longitude coordinates of the location of this work in our galleries",
					},
					map[string]any{
						"name": "longitude",
						"title": "Longitude",
						"type": "`$NUMBER`",
						"short": "Longitude coordinate of the location of this work in our galleries",
					},
					map[string]any{
						"name": "main_reference_number",
						"title": "Main Reference Number",
						"type": "`$INTEGER`",
						"short": "Unique identifier assigned to the artwork upon acquisition",
					},
					map[string]any{
						"name": "material_id",
						"title": "Material Id",
						"type": "`$STRING`",
						"short": "Unique identifier of the preferred material term for this work",
					},
					map[string]any{
						"name": "material_ids",
						"title": "Material Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of all material terms for this work",
					},
					map[string]any{
						"name": "material_titles",
						"title": "Material Titles",
						"type": "`$ANY`",
						"short": "The names of all material terms related to this artwork",
					},
					map[string]any{
						"name": "max_zoom_window_size",
						"title": "Max Zoom Window Size",
						"type": "`$ANY`",
						"short": "The maximum size of the window the image is allowed to be viewed in, in pixels.",
					},
					map[string]any{
						"name": "medium_display",
						"title": "Medium Display",
						"type": "`$ANY`",
						"short": "The substances or materials used in the creation of a work",
					},
					map[string]any{
						"name": "nomisma_id",
						"title": "Nomisma Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this work in the nomisma coin database",
					},
					map[string]any{
						"name": "on_loan_display",
						"title": "On Loan Display",
						"type": "`$ANY`",
						"short": "If an artwork is on loan, this contains details about the loan",
					},
					map[string]any{
						"name": "pageviews",
						"title": "Pageviews",
						"type": "`$ANY`",
						"short": "Approx.",
					},
					map[string]any{
						"name": "pageviews_recent",
						"title": "Pageviews Recent",
						"type": "`$ANY`",
						"short": "Approx.",
					},
					map[string]any{
						"name": "place_of_origin",
						"title": "Place Of Origin",
						"type": "`$ANY`",
						"short": "The location where the creation, design, or production of the work took place, or the original location of the work",
					},
					map[string]any{
						"name": "provenance_text",
						"title": "Provenance Text",
						"type": "`$ANY`",
						"short": "Ownership/collecting history of the work.",
					},
					map[string]any{
						"name": "publication_history",
						"title": "Publication History",
						"type": "`$ANY`",
						"short": "Bibliographic list of all the places this work has been published",
					},
					map[string]any{
						"name": "publishing_verification_level",
						"title": "Publishing Verification Level",
						"type": "`$ANY`",
						"short": "Indicator of how much metadata on the work in published.",
					},
					map[string]any{
						"name": "section_ids",
						"title": "Section Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of the digital publication chapters this work in included in",
					},
					map[string]any{
						"name": "section_titles",
						"title": "Section Titles",
						"type": "`$ANY`",
						"short": "Names of the digital publication chapters this work is included in",
					},
					map[string]any{
						"name": "short_description",
						"title": "Short Description",
						"type": "`$ANY`",
						"short": "Short explanation describing the work",
					},
					map[string]any{
						"name": "site_ids",
						"title": "Site Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of the microsites this work is a part of",
					},
					map[string]any{
						"name": "sound_ids",
						"title": "Sound Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of the audio about this work",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "style_id",
						"title": "Style Id",
						"type": "`$STRING`",
						"short": "Unique identifier of the preferred style term for this work",
					},
					map[string]any{
						"name": "style_ids",
						"title": "Style Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of all style terms for this work",
					},
					map[string]any{
						"name": "style_title",
						"title": "Style Title",
						"type": "`$ANY`",
						"short": "The name of the preferred style term for this work",
					},
					map[string]any{
						"name": "style_titles",
						"title": "Style Titles",
						"type": "`$ANY`",
						"short": "The names of all style terms related to this artwork",
					},
					map[string]any{
						"name": "subject_id",
						"title": "Subject Id",
						"type": "`$STRING`",
						"short": "Unique identifier of the preferred subject term for this work",
					},
					map[string]any{
						"name": "subject_ids",
						"title": "Subject Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of all subject terms for this work",
					},
					map[string]any{
						"name": "subject_titles",
						"title": "Subject Titles",
						"type": "`$ANY`",
						"short": "The names of all subject terms related to this artwork",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "technique_id",
						"title": "Technique Id",
						"type": "`$STRING`",
						"short": "Unique identifier of the preferred technique term for this work",
					},
					map[string]any{
						"name": "technique_ids",
						"title": "Technique Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of all technique terms for this work",
					},
					map[string]any{
						"name": "technique_titles",
						"title": "Technique Titles",
						"type": "`$ANY`",
						"short": "The names of all technique terms related to this artwork",
					},
					map[string]any{
						"name": "term_titles",
						"title": "Term Titles",
						"type": "`$ANY`",
						"short": "The names of the taxonomy tags for this work",
					},
					map[string]any{
						"name": "text_embedding",
						"title": "Text Embedding",
						"type": "`$ANY`",
						"short": "The generated embeddings of artwork text",
					},
					map[string]any{
						"name": "text_ids",
						"title": "Text Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of the texts about this work",
					},
					map[string]any{
						"name": "theme_titles",
						"title": "Theme Titles",
						"type": "`$ANY`",
						"short": "The names of all thematic publish categories related to this artwork",
					},
					map[string]any{
						"name": "thumbnail",
						"title": "Thumbnail",
						"type": "`$ANY`",
						"short": "Metadata about the image referenced by `image_id`.",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
					},
					map[string]any{
						"name": "video_ids",
						"title": "Video Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of the videos about this work",
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
								"kind": "http",
								"method": "GET",
								"orig": "/artworks",
								"segments": []any{
									map[string]any{
										"lit": "artworks",
									},
								},
								"parts": []any{
									"artworks",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"artworks",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
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
								"kind": "http",
								"method": "GET",
								"orig": "/artwork-date-qualifiers",
								"segments": []any{
									map[string]any{
										"lit": "artwork-date-qualifiers",
									},
								},
								"parts": []any{
									"artwork-date-qualifiers",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"artwork-date-qualifiers",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
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
								"kind": "http",
								"method": "GET",
								"orig": "/artwork-place-qualifiers",
								"segments": []any{
									map[string]any{
										"lit": "artwork-place-qualifiers",
									},
								},
								"parts": []any{
									"artwork-place-qualifiers",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"artwork-place-qualifiers",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Aat Id",
						"type": "`$STRING`",
						"short": "Identifier of reconciled (most similar) term in the Getty's Art and Architecture Thesaurus (AAT)",
					},
					map[string]any{
						"name": "api_link",
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
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
								"kind": "http",
								"method": "GET",
								"orig": "/artwork-types",
								"segments": []any{
									map[string]any{
										"lit": "artwork-types",
									},
								},
								"parts": []any{
									"artwork-types",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"artwork-types",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Aat Id",
						"type": "`$STRING`",
						"short": "Identifier of reconciled (most similar) term in the Getty's Art and Architecture Thesaurus (AAT)",
					},
					map[string]any{
						"name": "api_link",
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "parent_id",
						"title": "Parent Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this category's parent",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "subtype",
						"title": "Subtype",
						"type": "`$ANY`",
						"short": "Takes one of the following values: classification, material, technique, style, subject, department, theme",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
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
								"kind": "http",
								"method": "GET",
								"orig": "/category-terms",
								"segments": []any{
									map[string]any{
										"lit": "category-terms",
									},
								},
								"parts": []any{
									"category-terms",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"category-terms",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "copy",
						"title": "Copy",
						"type": "`$ANY`",
						"short": "The text of the page",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
					},
					map[string]any{
						"name": "web_url",
						"title": "Web Url",
						"type": "`$ANY`",
						"short": "The URL to this page on our website",
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
								"kind": "http",
								"method": "GET",
								"orig": "/digital-publications",
								"segments": []any{
									map[string]any{
										"lit": "digital-publications",
									},
								},
								"parts": []any{
									"digital-publications",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"digital-publications",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "author_display",
						"title": "Author Display",
						"type": "`$ANY`",
						"short": "A display-friendly text of the authors of this article",
					},
					map[string]any{
						"name": "copy",
						"title": "Copy",
						"type": "`$ANY`",
						"short": "The text of the article",
					},
					map[string]any{
						"name": "digital_publication_id",
						"title": "Digital Publication Id",
						"type": "`$STRING`",
						"short": "Unique identifier of the digital publication this article belongs to",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
					},
					map[string]any{
						"name": "web_url",
						"title": "Web Url",
						"type": "`$ANY`",
						"short": "The URL to this article on our website",
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
								"kind": "http",
								"method": "GET",
								"orig": "/digital-publication-articles",
								"segments": []any{
									map[string]any{
										"lit": "digital-publication-articles",
									},
								},
								"parts": []any{
									"digital-publication-articles",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"digital-publication-articles",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "copy",
						"title": "Copy",
						"type": "`$ANY`",
						"short": "The text of the page",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
					},
					map[string]any{
						"name": "web_url",
						"title": "Web Url",
						"type": "`$ANY`",
						"short": "The URL to this page on our website",
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
								"kind": "http",
								"method": "GET",
								"orig": "/educator-resources",
								"segments": []any{
									map[string]any{
										"lit": "educator-resources",
									},
								},
								"parts": []any{
									"educator-resources",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"educator-resources",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Alt Audience Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers indicating the alternate audiences for this event",
					},
					map[string]any{
						"name": "alt_event_type_ids",
						"title": "Alt Event Type Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers indicating the alternate types of this event",
					},
					map[string]any{
						"name": "api_link",
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "audience_id",
						"title": "Audience Id",
						"type": "`$STRING`",
						"short": "Unique identifier indicating the preferred audience for this event",
					},
					map[string]any{
						"name": "buy_button_caption",
						"title": "Buy Button Caption",
						"type": "`$ANY`",
						"short": "Additional text below the ticket/registration button",
					},
					map[string]any{
						"name": "buy_button_text",
						"title": "Buy Button Text",
						"type": "`$ANY`",
						"short": "The text used on the ticket/registration button",
					},
					map[string]any{
						"name": "date_display",
						"title": "Date Display",
						"type": "`$ANY`",
						"short": "A readable display of the event dates",
					},
					map[string]any{
						"name": "description",
						"title": "Description",
						"type": "`$STRING`",
						"short": "All copytext of the event",
					},
					map[string]any{
						"name": "door_time",
						"title": "Door Time",
						"type": "`$ANY`",
						"short": "The time the doors open for this event",
					},
					map[string]any{
						"name": "end_date",
						"title": "End Date",
						"type": "`$ANY`",
						"short": "The date the event ends",
					},
					map[string]any{
						"name": "end_time",
						"title": "End Time",
						"type": "`$ANY`",
						"short": "The time the event ends",
					},
					map[string]any{
						"name": "entrance",
						"title": "Entrance",
						"type": "`$ANY`",
						"short": "Which entrance to use for this event",
					},
					map[string]any{
						"name": "event_host_id",
						"title": "Event Host Id",
						"type": "`$STRING`",
						"short": "Unique identifier of the host (cf.",
					},
					map[string]any{
						"name": "event_host_title",
						"title": "Event Host Title",
						"type": "`$ANY`",
						"short": "Unique identifier of the host (cf.",
					},
					map[string]any{
						"name": "event_type_id",
						"title": "Event Type Id",
						"type": "`$STRING`",
						"short": "Unique identifier indicating the preferred type of this event",
					},
					map[string]any{
						"name": "header_description",
						"title": "Header Description",
						"type": "`$ANY`",
						"short": "Brief description of the event displayed below the title",
					},
					map[string]any{
						"name": "hero_caption",
						"title": "Hero Caption",
						"type": "`$ANY`",
						"short": "Text displayed with the hero image on the event",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "image_url",
						"title": "Image Url",
						"type": "`$ANY`",
						"short": "The URL of an image representing this page",
					},
					map[string]any{
						"name": "is_admission_required",
						"title": "Is Admission Required",
						"type": "`$BOOLEAN`",
						"short": "Whether admission to the museum is required to attend this event",
					},
					map[string]any{
						"name": "is_after_hours",
						"title": "Is After Hours",
						"type": "`$BOOLEAN`",
						"short": "Whether the event is to be held after the museum closes",
					},
					map[string]any{
						"name": "is_free",
						"title": "Is Free",
						"type": "`$BOOLEAN`",
						"short": "Whether the event is free",
					},
					map[string]any{
						"name": "is_member_exclusive",
						"title": "Is Member Exclusive",
						"type": "`$BOOLEAN`",
						"short": "Whether the event is exclusive to members of the museum",
					},
					map[string]any{
						"name": "is_private",
						"title": "Is Private",
						"type": "`$BOOLEAN`",
						"short": "Whether the event is private",
					},
					map[string]any{
						"name": "is_registration_required",
						"title": "Is Registration Required",
						"type": "`$BOOLEAN`",
						"short": "Whether registration is required to attend the event",
					},
					map[string]any{
						"name": "is_sales_button_hidden",
						"title": "Is Sales Button Hidden",
						"type": "`$BOOLEAN`",
						"short": "Whether the buy tickets button is hidden on the website event page",
					},
					map[string]any{
						"name": "is_sold_out",
						"title": "Is Sold Out",
						"type": "`$BOOLEAN`",
						"short": "Whether the event is sold out",
					},
					map[string]any{
						"name": "is_ticketed",
						"title": "Is Ticketed",
						"type": "`$BOOLEAN`",
						"short": "Whether a ticket is required to attend the event",
					},
					map[string]any{
						"name": "is_virtual_event",
						"title": "Is Virtual Event",
						"type": "`$BOOLEAN`",
						"short": "Whether the event is being held virtually",
					},
					map[string]any{
						"name": "join_url",
						"title": "Join Url",
						"type": "`$ANY`",
						"short": "URL to the membership signup page via this event",
					},
					map[string]any{
						"name": "layout_type",
						"title": "Layout Type",
						"type": "`$ANY`",
						"short": "Number indicating the type of layout this event page uses",
					},
					map[string]any{
						"name": "list_description",
						"title": "List Description",
						"type": "`$ANY`",
						"short": "One-sentence description of the event displayed in listings",
					},
					map[string]any{
						"name": "location",
						"title": "Location",
						"type": "`$ANY`",
						"short": "Where the event takes place",
					},
					map[string]any{
						"name": "program_ids",
						"title": "Program Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers indicating the programs this event is a part of",
					},
					map[string]any{
						"name": "program_titles",
						"title": "Program Titles",
						"type": "`$ANY`",
						"short": "Titles of the programs this event is a part of",
					},
					map[string]any{
						"name": "rsvp_link",
						"title": "Rsvp Link",
						"type": "`$ANY`",
						"short": "The URL to the sales site for this event",
					},
					map[string]any{
						"name": "search_tags",
						"title": "Search Tags",
						"type": "`$ANY`",
						"short": "Editor-specified list of tags to aid in internal search",
					},
					map[string]any{
						"name": "short_description",
						"title": "Short Description",
						"type": "`$ANY`",
						"short": "Brief description of the event",
					},
					map[string]any{
						"name": "slug",
						"title": "Slug",
						"type": "`$STRING`",
						"short": "A string used in the URL for this event",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "start_date",
						"title": "Start Date",
						"type": "`$ANY`",
						"short": "The date the event begins",
					},
					map[string]any{
						"name": "start_time",
						"title": "Start Time",
						"type": "`$ANY`",
						"short": "The time the event starts",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "survey_url",
						"title": "Survey Url",
						"type": "`$ANY`",
						"short": "URL to the survey associated with this event",
					},
					map[string]any{
						"name": "ticketed_event_id",
						"title": "Ticketed Event Id",
						"type": "`$STRING`",
						"short": "Unique identifier of the event in the ticketing system this website event is tied to",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "title_display",
						"title": "Title Display",
						"type": "`$ANY`",
						"short": "The name of this event formatted with HTML (optional)",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
					},
					map[string]any{
						"name": "virtual_event_passcode",
						"title": "Virtual Event Passcode",
						"type": "`$ANY`",
						"short": "Passcode to access the virtual event",
					},
					map[string]any{
						"name": "virtual_event_url",
						"title": "Virtual Event Url",
						"type": "`$ANY`",
						"short": "URL to the virtual event",
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
								"kind": "http",
								"method": "GET",
								"orig": "/events",
								"segments": []any{
									map[string]any{
										"lit": "events",
									},
								},
								"parts": []any{
									"events",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"events",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "button_caption",
						"title": "Button Caption",
						"type": "`$ANY`",
						"short": "Additional text below the ticket/registration button",
					},
					map[string]any{
						"name": "button_text",
						"title": "Button Text",
						"type": "`$ANY`",
						"short": "The text used on the ticket/registration button",
					},
					map[string]any{
						"name": "button_url",
						"title": "Button Url",
						"type": "`$ANY`",
						"short": "The URL to the sales site or an RSVP link for this event",
					},
					map[string]any{
						"name": "description",
						"title": "Description",
						"type": "`$STRING`",
						"short": "Description of the event",
					},
					map[string]any{
						"name": "end_at",
						"title": "End At",
						"type": "`$ANY`",
						"short": "The date the event occurrence ends",
					},
					map[string]any{
						"name": "event_id",
						"title": "Event Id",
						"type": "`$STRING`",
						"short": "Identifier of the master event of which this is an occurrence",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "image_url",
						"title": "Image Url",
						"type": "`$ANY`",
						"short": "The URL of an image representing this page",
					},
					map[string]any{
						"name": "is_private",
						"title": "Is Private",
						"type": "`$BOOLEAN`",
						"short": "Whether the event is private.",
					},
					map[string]any{
						"name": "is_sales_button_hidden",
						"title": "Is Sales Button Hidden",
						"type": "`$BOOLEAN`",
						"short": "Whether the buy tickets button is hidden on the website event page",
					},
					map[string]any{
						"name": "is_ticketed",
						"title": "Is Ticketed",
						"type": "`$BOOLEAN`",
						"short": "Whether a ticket is required to attend the event",
					},
					map[string]any{
						"name": "location",
						"title": "Location",
						"type": "`$ANY`",
						"short": "Where the event takes place",
					},
					map[string]any{
						"name": "off_sale_at",
						"title": "Off Sale At",
						"type": "`$ANY`",
						"short": "Date and time the event goes off sale",
					},
					map[string]any{
						"name": "on_sale_at",
						"title": "On Sale At",
						"type": "`$ANY`",
						"short": "Date and time the event goes on sale",
					},
					map[string]any{
						"name": "short_description",
						"title": "Short Description",
						"type": "`$ANY`",
						"short": "Brief description of the event",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "start_at",
						"title": "Start At",
						"type": "`$ANY`",
						"short": "The date the event occurrence begins",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "title_display",
						"title": "Title Display",
						"type": "`$ANY`",
						"short": "The name of this event formatted with HTML (optional)",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
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
								"kind": "http",
								"method": "GET",
								"orig": "/event-occurrences",
								"segments": []any{
									map[string]any{
										"lit": "event-occurrences",
									},
								},
								"parts": []any{
									"event-occurrences",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"event-occurrences",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "is_affiliate_group",
						"title": "Is Affiliate Group",
						"type": "`$BOOLEAN`",
						"short": "Whether this program represents an affiliate group",
					},
					map[string]any{
						"name": "is_event_host",
						"title": "Is Event Host",
						"type": "`$BOOLEAN`",
						"short": "Whether this program represents an event host",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
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
								"kind": "http",
								"method": "GET",
								"orig": "/event-programs",
								"segments": []any{
									map[string]any{
										"lit": "event-programs",
									},
								},
								"parts": []any{
									"event-programs",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"event-programs",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Aic End At",
						"type": "`$ANY`",
						"short": "Date the exhibition closed at the Art Institute of Chicago",
					},
					map[string]any{
						"name": "aic_start_at",
						"title": "Aic Start At",
						"type": "`$ANY`",
						"short": "Date the exhibition opened at the Art Institute of Chicago",
					},
					map[string]any{
						"name": "alt_image_ids",
						"title": "Alt Image Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of all non-preferred images of this exhibition.",
					},
					map[string]any{
						"name": "api_link",
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "artist_ids",
						"title": "Artist Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of the artist agent records representing who was shown in the exhibition",
					},
					map[string]any{
						"name": "artwork_ids",
						"title": "Artwork Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of the artworks that were part of the exhibition",
					},
					map[string]any{
						"name": "artwork_titles",
						"title": "Artwork Titles",
						"type": "`$ANY`",
						"short": "Names of the artworks that were part of the exhibition",
					},
					map[string]any{
						"name": "document_ids",
						"title": "Document Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of assets that serve as documentation for this exhibition",
					},
					map[string]any{
						"name": "gallery_id",
						"title": "Gallery Id",
						"type": "`$STRING`",
						"short": "Unique identifier of the gallery that mainly housed the exhibition",
					},
					map[string]any{
						"name": "gallery_title",
						"title": "Gallery Title",
						"type": "`$ANY`",
						"short": "The name of the gallery that mainly housed the exhibition",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "image_id",
						"title": "Image Id",
						"type": "`$STRING`",
						"short": "Unique identifier of the preferred image to use to represent this exhibition",
					},
					map[string]any{
						"name": "image_url",
						"title": "Image Url",
						"type": "`$ANY`",
						"short": "URL to the hero image from the website",
					},
					map[string]any{
						"name": "is_featured",
						"title": "Is Featured",
						"type": "`$BOOLEAN`",
						"short": "Is this exhibition currently featured on our website?",
					},
					map[string]any{
						"name": "is_published",
						"title": "Is Published",
						"type": "`$BOOLEAN`",
						"short": "Is this exhibition currently published on our website?",
					},
					map[string]any{
						"name": "position",
						"title": "Position",
						"type": "`$ANY`",
						"short": "Numering position represnting the order in which this exhibition is featured on the website",
					},
					map[string]any{
						"name": "short_description",
						"title": "Short Description",
						"type": "`$ANY`",
						"short": "Brief explanation of what this exhibition is",
					},
					map[string]any{
						"name": "site_ids",
						"title": "Site Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of the microsites this exhibition is a part of",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "status",
						"title": "Status",
						"type": "`$ANY`",
						"short": "Whether the exhibition is open or closed",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
					},
					map[string]any{
						"name": "web_url",
						"title": "Web Url",
						"type": "`$ANY`",
						"short": "URL to this exhibition on our website",
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
								"kind": "http",
								"method": "GET",
								"orig": "/exhibitions",
								"segments": []any{
									map[string]any{
										"lit": "exhibitions",
									},
								},
								"parts": []any{
									"exhibitions",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"exhibitions",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "floor",
						"title": "Floor",
						"type": "`$ANY`",
						"short": "The level the gallery is on, e.g., 1, 2, 3, or LL",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "is_closed",
						"title": "Is Closed",
						"type": "`$BOOLEAN`",
						"short": "Whether the gallery is currently closed",
					},
					map[string]any{
						"name": "latitude",
						"title": "Latitude",
						"type": "`$NUMBER`",
						"short": "Latitude coordinate of the center of the room",
					},
					map[string]any{
						"name": "latlon",
						"title": "Latlon",
						"type": "`$ANY`",
						"short": "Latitude and longitude coordinates of the center of the room",
					},
					map[string]any{
						"name": "longitude",
						"title": "Longitude",
						"type": "`$NUMBER`",
						"short": "Longitude coordinate of the center of the room",
					},
					map[string]any{
						"name": "number",
						"title": "Number",
						"type": "`$ANY`",
						"short": "The gallery's room number.",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "tgn_id",
						"title": "Tgn Id",
						"type": "`$STRING`",
						"short": "Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN)",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
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
								"kind": "http",
								"method": "GET",
								"orig": "/galleries",
								"segments": []any{
									map[string]any{
										"lit": "galleries",
									},
								},
								"parts": []any{
									"galleries",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"galleries",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "copy",
						"title": "Copy",
						"type": "`$ANY`",
						"short": "The text of the page",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "search_tags",
						"title": "Search Tags",
						"type": "`$ANY`",
						"short": "Editor-specified list of tags to aid in internal search",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
					},
					map[string]any{
						"name": "web_url",
						"title": "Web Url",
						"type": "`$ANY`",
						"short": "The URL to this page on our website",
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
								"kind": "http",
								"method": "GET",
								"orig": "/generic-pages",
								"segments": []any{
									map[string]any{
										"lit": "generic-pages",
									},
								},
								"parts": []any{
									"generic-pages",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"generic-pages",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "copy",
						"title": "Copy",
						"type": "`$ANY`",
						"short": "The text of the highlight description",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
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
								"kind": "http",
								"method": "GET",
								"orig": "/highlights",
								"segments": []any{
									map[string]any{
										"lit": "highlights",
									},
								},
								"parts": []any{
									"highlights",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"highlights",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Additional Text",
						"type": "`$ANY`",
						"short": "Additional information about the hours",
					},
					map[string]any{
						"name": "api_link",
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "friday_is_closed",
						"title": "Friday Is Closed",
						"type": "`$ANY`",
						"short": "Whether the museum is closed on Fridays",
					},
					map[string]any{
						"name": "friday_member_close",
						"title": "Friday Member Close",
						"type": "`$ANY`",
						"short": "The time member hours ends on Fridays",
					},
					map[string]any{
						"name": "friday_member_open",
						"title": "Friday Member Open",
						"type": "`$ANY`",
						"short": "The time member hours starts on Fridays",
					},
					map[string]any{
						"name": "friday_public_close",
						"title": "Friday Public Close",
						"type": "`$ANY`",
						"short": "The time public hours ends on Fridays",
					},
					map[string]any{
						"name": "friday_public_open",
						"title": "Friday Public Open",
						"type": "`$ANY`",
						"short": "The time public hours starts on Fridays",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "monday_is_closed",
						"title": "Monday Is Closed",
						"type": "`$ANY`",
						"short": "Whether the museum is closed on Mondays",
					},
					map[string]any{
						"name": "monday_member_close",
						"title": "Monday Member Close",
						"type": "`$ANY`",
						"short": "The time member hours ends on Mondays",
					},
					map[string]any{
						"name": "monday_member_open",
						"title": "Monday Member Open",
						"type": "`$ANY`",
						"short": "The time member hours starts on Mondays",
					},
					map[string]any{
						"name": "monday_public_close",
						"title": "Monday Public Close",
						"type": "`$ANY`",
						"short": "The time public hours ends on Mondays",
					},
					map[string]any{
						"name": "monday_public_open",
						"title": "Monday Public Open",
						"type": "`$ANY`",
						"short": "The time public hours starts on Mondays",
					},
					map[string]any{
						"name": "saturday_is_closed",
						"title": "Saturday Is Closed",
						"type": "`$ANY`",
						"short": "Whether the museum is closed on Saturdays",
					},
					map[string]any{
						"name": "saturday_member_close",
						"title": "Saturday Member Close",
						"type": "`$ANY`",
						"short": "The time member hours ends on Saturdays",
					},
					map[string]any{
						"name": "saturday_member_open",
						"title": "Saturday Member Open",
						"type": "`$ANY`",
						"short": "The time member hours starts on Saturdays",
					},
					map[string]any{
						"name": "saturday_public_close",
						"title": "Saturday Public Close",
						"type": "`$ANY`",
						"short": "The time public hours ends on Saturdays",
					},
					map[string]any{
						"name": "saturday_public_open",
						"title": "Saturday Public Open",
						"type": "`$ANY`",
						"short": "The time public hours starts on Saturdays",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "summary",
						"title": "Summary",
						"type": "`$ANY`",
						"short": "Readable summary of the hours",
					},
					map[string]any{
						"name": "sunday_is_closed",
						"title": "Sunday Is Closed",
						"type": "`$ANY`",
						"short": "Whether the museum is closed on Sundays",
					},
					map[string]any{
						"name": "sunday_member_close",
						"title": "Sunday Member Close",
						"type": "`$ANY`",
						"short": "The time member hours ends on Sundays",
					},
					map[string]any{
						"name": "sunday_member_open",
						"title": "Sunday Member Open",
						"type": "`$ANY`",
						"short": "The time member hours starts on Sundays",
					},
					map[string]any{
						"name": "sunday_public_close",
						"title": "Sunday Public Close",
						"type": "`$ANY`",
						"short": "The time public hours ends on Sundays",
					},
					map[string]any{
						"name": "sunday_public_open",
						"title": "Sunday Public Open",
						"type": "`$ANY`",
						"short": "The time public hours starts on Sundays",
					},
					map[string]any{
						"name": "thursday_is_closed",
						"title": "Thursday Is Closed",
						"type": "`$ANY`",
						"short": "Whether the museum is closed on Thursdays",
					},
					map[string]any{
						"name": "thursday_member_close",
						"title": "Thursday Member Close",
						"type": "`$ANY`",
						"short": "The time member hours ends on Thursdays",
					},
					map[string]any{
						"name": "thursday_member_open",
						"title": "Thursday Member Open",
						"type": "`$ANY`",
						"short": "The time member hours starts on Thursdays",
					},
					map[string]any{
						"name": "thursday_public_close",
						"title": "Thursday Public Close",
						"type": "`$ANY`",
						"short": "The time public hours ends on Thursdays",
					},
					map[string]any{
						"name": "thursday_public_open",
						"title": "Thursday Public Open",
						"type": "`$ANY`",
						"short": "The time public hours starts on Thursdays",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "tuesday_is_closed",
						"title": "Tuesday Is Closed",
						"type": "`$ANY`",
						"short": "Whether the museum is closed on Tuesdays",
					},
					map[string]any{
						"name": "tuesday_member_close",
						"title": "Tuesday Member Close",
						"type": "`$ANY`",
						"short": "The time member hours ends on Tuesdays",
					},
					map[string]any{
						"name": "tuesday_member_open",
						"title": "Tuesday Member Open",
						"type": "`$ANY`",
						"short": "The time member hours starts on Tuesdays",
					},
					map[string]any{
						"name": "tuesday_public_close",
						"title": "Tuesday Public Close",
						"type": "`$ANY`",
						"short": "The time public hours ends on Tuesdays",
					},
					map[string]any{
						"name": "tuesday_public_open",
						"title": "Tuesday Public Open",
						"type": "`$ANY`",
						"short": "The time public hours starts on Tuesdays",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
					},
					map[string]any{
						"name": "wednesday_is_closed",
						"title": "Wednesday Is Closed",
						"type": "`$ANY`",
						"short": "Whether the museum is closed on Wednesdays",
					},
					map[string]any{
						"name": "wednesday_member_close",
						"title": "Wednesday Member Close",
						"type": "`$ANY`",
						"short": "The time member hours ends on Wednesdays",
					},
					map[string]any{
						"name": "wednesday_member_open",
						"title": "Wednesday Member Open",
						"type": "`$ANY`",
						"short": "The time member hours starts on Wednesdays",
					},
					map[string]any{
						"name": "wednesday_public_close",
						"title": "Wednesday Public Close",
						"type": "`$ANY`",
						"short": "The time public hours ends on Wednesdays",
					},
					map[string]any{
						"name": "wednesday_public_open",
						"title": "Wednesday Public Open",
						"type": "`$ANY`",
						"short": "The time public hours starts on Wednesdays",
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
								"kind": "http",
								"method": "GET",
								"orig": "/hours",
								"segments": []any{
									map[string]any{
										"lit": "hours",
									},
								},
								"parts": []any{
									"hours",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"hours",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Ahash",
						"type": "`$ANY`",
						"short": "Image hash generated using ahash algorithm with 64 boolean subfields",
					},
					map[string]any{
						"name": "alt_text",
						"title": "Alt Text",
						"type": "`$ANY`",
						"short": "Alternative text for the asset to describe it to people with low or no vision",
					},
					map[string]any{
						"name": "api_link",
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "artwork_ids",
						"title": "Artwork Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of the artworks associated with this asset",
					},
					map[string]any{
						"name": "artwork_titles",
						"title": "Artwork Titles",
						"type": "`$ANY`",
						"short": "Names of the artworks associated with this asset",
					},
					map[string]any{
						"name": "color",
						"title": "Color",
						"type": "`$ANY`",
						"short": "Dominant color of this image in HSL",
					},
					map[string]any{
						"name": "colorfulness",
						"title": "Colorfulness",
						"type": "`$ANY`",
						"short": "Unbounded positive float representing an abstract measure of colorfulness.",
					},
					map[string]any{
						"name": "content",
						"title": "Content",
						"type": "`$ANY`",
						"short": "Text of or URL to the contents of this asset",
					},
					map[string]any{
						"name": "content_e_tag",
						"title": "Content E Tag",
						"type": "`$ANY`",
						"short": "Arbitrary unique identifier that changes when the binary file gets updated",
					},
					map[string]any{
						"name": "credit_line",
						"title": "Credit Line",
						"type": "`$ANY`",
						"short": "Asset-specific copyright information",
					},
					map[string]any{
						"name": "fingerprint",
						"title": "Fingerprint",
						"type": "`$ANY`",
						"short": "Image hashes: aHash, dHash, pHash, wHash",
					},
					map[string]any{
						"name": "height",
						"title": "Height",
						"type": "`$NUMBER`",
						"short": "Native height of the image",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "iiif_url",
						"title": "Iiif Url",
						"type": "`$ANY`",
						"short": "IIIF URL of this image",
					},
					map[string]any{
						"name": "is_educational_resource",
						"title": "Is Educational Resource",
						"type": "`$BOOLEAN`",
						"short": "Whether this resource is considered to be educational",
					},
					map[string]any{
						"name": "is_multimedia_resource",
						"title": "Is Multimedia Resource",
						"type": "`$BOOLEAN`",
						"short": "Whether this resource is considered to be multimedia",
					},
					map[string]any{
						"name": "is_teacher_resource",
						"title": "Is Teacher Resource",
						"type": "`$BOOLEAN`",
						"short": "Whether this resource is considered to be educational",
					},
					map[string]any{
						"name": "lake_guid",
						"title": "Lake Guid",
						"type": "`$ANY`",
						"short": "Unique UUID of this resource in LAKE, our DAMS.",
					},
					map[string]any{
						"name": "lqip",
						"title": "Lqip",
						"type": "`$ANY`",
						"short": "Low-quality image placeholder (LQIP).",
					},
					map[string]any{
						"name": "phash",
						"title": "Phash",
						"type": "`$ANY`",
						"short": "Image hash generated using phash algorithm with 64 boolean subfields",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "type",
						"title": "Type",
						"type": "`$ANY`",
						"short": "Type always takes one of the following values: image, sound, text, video",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
					},
					map[string]any{
						"name": "width",
						"title": "Width",
						"type": "`$NUMBER`",
						"short": "Native width of the image",
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
								"kind": "http",
								"method": "GET",
								"orig": "/images",
								"segments": []any{
									map[string]any{
										"lit": "images",
									},
								},
								"parts": []any{
									"images",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"images",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "copy",
						"title": "Copy",
						"type": "`$ANY`",
						"short": "The text of the page",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "search_tags",
						"title": "Search Tags",
						"type": "`$ANY`",
						"short": "Editor-specified list of tags to aid in internal search",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
					},
					map[string]any{
						"name": "web_url",
						"title": "Web Url",
						"type": "`$ANY`",
						"short": "The URL to this page on our website",
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
								"kind": "http",
								"method": "GET",
								"orig": "/landing-pages",
								"segments": []any{
									map[string]any{
										"lit": "landing-pages",
									},
								},
								"parts": []any{
									"landing-pages",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"landing-pages",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "latitude",
						"title": "Latitude",
						"type": "`$NUMBER`",
						"short": "Latitude coordinate of the center of the room",
					},
					map[string]any{
						"name": "longitude",
						"title": "Longitude",
						"type": "`$NUMBER`",
						"short": "Longitude coordinate of the center of the room",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "tgn_id",
						"title": "Tgn Id",
						"type": "`$STRING`",
						"short": "Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN)",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
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
								"kind": "http",
								"method": "GET",
								"orig": "/places",
								"segments": []any{
									map[string]any{
										"lit": "places",
									},
								},
								"parts": []any{
									"places",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"places",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "copy",
						"title": "Copy",
						"type": "`$ANY`",
						"short": "The text of the page",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
					},
					map[string]any{
						"name": "web_url",
						"title": "Web Url",
						"type": "`$ANY`",
						"short": "The URL to this page on our website",
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
								"kind": "http",
								"method": "GET",
								"orig": "/press-releases",
								"segments": []any{
									map[string]any{
										"lit": "press-releases",
									},
								},
								"parts": []any{
									"press-releases",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"press-releases",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "copy",
						"title": "Copy",
						"type": "`$ANY`",
						"short": "The text of the page",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
					},
					map[string]any{
						"name": "web_url",
						"title": "Web Url",
						"type": "`$ANY`",
						"short": "The URL to this page on our website",
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
								"kind": "http",
								"method": "GET",
								"orig": "/printed-publications",
								"segments": []any{
									map[string]any{
										"lit": "printed-publications",
									},
								},
								"parts": []any{
									"printed-publications",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"printed-publications",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "artist_ids",
						"title": "Artist Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of the artists associated with this product",
					},
					map[string]any{
						"name": "artwork_ids",
						"title": "Artwork Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of the artworks associated with this product",
					},
					map[string]any{
						"name": "description",
						"title": "Description",
						"type": "`$STRING`",
						"short": "Explanation of what this product is",
					},
					map[string]any{
						"name": "exhibition_ids",
						"title": "Exhibition Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of the exhibitions associated with this product",
					},
					map[string]any{
						"name": "external_sku",
						"title": "External Sku",
						"type": "`$ANY`",
						"short": "Numeric product identification code of a machine-readable barcode, when the customer sku differs from our internal one",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "image_url",
						"title": "Image Url",
						"type": "`$ANY`",
						"short": "URL of an image for this product",
					},
					map[string]any{
						"name": "max_compare_at_price",
						"title": "Max Compare At Price",
						"type": "`$ANY`",
						"short": "Number indicating how much the most expensive variant of a product cost before a sale",
					},
					map[string]any{
						"name": "max_current_price",
						"title": "Max Current Price",
						"type": "`$ANY`",
						"short": "Number indicating how much the most expensive variant of a product costs right now",
					},
					map[string]any{
						"name": "min_compare_at_price",
						"title": "Min Compare At Price",
						"type": "`$ANY`",
						"short": "Number indicating how much the least expensive variant of a product cost before a sale",
					},
					map[string]any{
						"name": "min_current_price",
						"title": "Min Current Price",
						"type": "`$ANY`",
						"short": "Number indicating how much the least expensive variant of a product costs right now",
					},
					map[string]any{
						"name": "price_display",
						"title": "Price Display",
						"type": "`$ANY`",
						"short": "Explanation of what this product is",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
					},
					map[string]any{
						"name": "web_url",
						"title": "Web Url",
						"type": "`$ANY`",
						"short": "URL of this product in the shop",
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
								"kind": "http",
								"method": "GET",
								"orig": "/products",
								"segments": []any{
									map[string]any{
										"lit": "products",
									},
								},
								"parts": []any{
									"products",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"products",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "section_ids",
						"title": "Section Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of the sections of this publication",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
					},
					map[string]any{
						"name": "web_url",
						"title": "Web Url",
						"type": "`$ANY`",
						"short": "URL to the publication",
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
								"kind": "http",
								"method": "GET",
								"orig": "/publications",
								"segments": []any{
									map[string]any{
										"lit": "publications",
									},
								},
								"parts": []any{
									"publications",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"publications",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Api Id",
						"type": "`$STRING`",
						"short": "API unique identifier",
					},
					map[string]any{
						"name": "api_link",
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "URL to this recource in the API",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "Name of the model the resource represents",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier within the search index",
					},
					map[string]any{
						"name": "is_boosted",
						"title": "Is Boosted",
						"type": "`$BOOLEAN`",
						"short": "Whether this record has been flagged to be boosted",
					},
					map[string]any{
						"name": "score",
						"title": "Score",
						"type": "`$NUMBER`",
						"short": "Search index ranking of the result",
					},
					map[string]any{
						"name": "thumbnail",
						"title": "Thumbnail",
						"type": "`$ANY`",
						"short": "Metadata on the image representing this record",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date this record was last updated in the API",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
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
								"parts": []any{
									"agents",
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
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
							},
							map[string]any{
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
								"parts": []any{
									"articles",
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
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
							},
							map[string]any{
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
								"parts": []any{
									"artworks",
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
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
							},
							map[string]any{
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
								"parts": []any{
									"category-terms",
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
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
							},
							map[string]any{
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
								"parts": []any{
									"digital-publication-articles",
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
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
							},
							map[string]any{
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
								"parts": []any{
									"digital-publications",
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
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
							},
							map[string]any{
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
								"parts": []any{
									"educator-resources",
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
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
							},
							map[string]any{
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
								"parts": []any{
									"event-occurrences",
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
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
							},
							map[string]any{
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
								"parts": []any{
									"event-programs",
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
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
							},
							map[string]any{
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
								"parts": []any{
									"events",
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
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
							},
							map[string]any{
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
								"parts": []any{
									"exhibitions",
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
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
							},
							map[string]any{
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
								"parts": []any{
									"galleries",
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
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
							},
							map[string]any{
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
								"parts": []any{
									"generic-pages",
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
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
							},
							map[string]any{
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
								"parts": []any{
									"highlights",
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
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
							},
							map[string]any{
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
								"parts": []any{
									"hours",
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
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
							},
							map[string]any{
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
								"parts": []any{
									"images",
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
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
							},
							map[string]any{
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
								"parts": []any{
									"landing-pages",
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
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
							},
							map[string]any{
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
								"parts": []any{
									"mobile-sounds",
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
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
							},
							map[string]any{
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
								"parts": []any{
									"places",
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
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
							},
							map[string]any{
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
								"parts": []any{
									"press-releases",
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
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
							},
							map[string]any{
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
								"parts": []any{
									"printed-publications",
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
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
							},
							map[string]any{
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
								"parts": []any{
									"products",
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
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
							},
							map[string]any{
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
								"parts": []any{
									"publications",
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
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
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/search",
								"segments": []any{
									map[string]any{
										"lit": "search",
									},
								},
								"parts": []any{
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
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
							},
							map[string]any{
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
								"parts": []any{
									"sections",
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
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
							},
							map[string]any{
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
								"parts": []any{
									"sites",
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
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
							},
							map[string]any{
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
								"parts": []any{
									"sounds",
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
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
							},
							map[string]any{
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
								"parts": []any{
									"static-pages",
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
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
							},
							map[string]any{
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
								"parts": []any{
									"texts",
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
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
							},
							map[string]any{
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
								"parts": []any{
									"tours",
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
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
							},
							map[string]any{
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
								"parts": []any{
									"videos",
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
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
						"title": "Accession",
						"type": "`$ANY`",
						"short": "An accession number parsed from the title or tombstone",
					},
					map[string]any{
						"name": "api_link",
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "artwork_id",
						"title": "Artwork Id",
						"type": "`$STRING`",
						"short": "Unique identifier of the artwork with which this section is associated",
					},
					map[string]any{
						"name": "content",
						"title": "Content",
						"type": "`$ANY`",
						"short": "Content of this section in plaintext",
					},
					map[string]any{
						"name": "generic_page_id",
						"title": "Generic Page Id",
						"type": "`$STRING`",
						"short": "Unique identifier of the page on the website that represents the publication this section belongs to",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "publication_id",
						"title": "Publication Id",
						"type": "`$STRING`",
						"short": "Unique identifier of the publication this section belongs to",
					},
					map[string]any{
						"name": "publication_title",
						"title": "Publication Title",
						"type": "`$ANY`",
						"short": "Name of the publication this section belongs to",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
					},
					map[string]any{
						"name": "web_url",
						"title": "Web Url",
						"type": "`$ANY`",
						"short": "URL to the section",
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
								"kind": "http",
								"method": "GET",
								"orig": "/sections",
								"segments": []any{
									map[string]any{
										"lit": "sections",
									},
								},
								"parts": []any{
									"sections",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"sections",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "artwork_ids",
						"title": "Artwork Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of the artworks this site is associated with",
					},
					map[string]any{
						"name": "artwork_titles",
						"title": "Artwork Titles",
						"type": "`$ANY`",
						"short": "Names of the artworks this site is associated with",
					},
					map[string]any{
						"name": "description",
						"title": "Description",
						"type": "`$STRING`",
						"short": "Explanation of what this site is",
					},
					map[string]any{
						"name": "exhibition_ids",
						"title": "Exhibition Ids",
						"type": "`$ANY`",
						"short": "Unique identifier of the exhibitions this site is associated with",
					},
					map[string]any{
						"name": "exhibition_titles",
						"title": "Exhibition Titles",
						"type": "`$ANY`",
						"short": "Names of the exhibitions this site is associated with",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
					},
					map[string]any{
						"name": "web_url",
						"title": "Web Url",
						"type": "`$ANY`",
						"short": "URL to this site",
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
								"kind": "http",
								"method": "GET",
								"orig": "/sites",
								"segments": []any{
									map[string]any{
										"lit": "sites",
									},
								},
								"parts": []any{
									"sites",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"sites",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Alt Text",
						"type": "`$ANY`",
						"short": "Alternative text for the asset to describe it to people with low or no vision",
					},
					map[string]any{
						"name": "api_link",
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "artwork_ids",
						"title": "Artwork Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of the artworks associated with this asset",
					},
					map[string]any{
						"name": "artwork_titles",
						"title": "Artwork Titles",
						"type": "`$ANY`",
						"short": "Names of the artworks associated with this asset",
					},
					map[string]any{
						"name": "content",
						"title": "Content",
						"type": "`$ANY`",
						"short": "Text of or URL to the contents of this asset",
					},
					map[string]any{
						"name": "content_e_tag",
						"title": "Content E Tag",
						"type": "`$ANY`",
						"short": "Arbitrary unique identifier that changes when the binary file gets updated",
					},
					map[string]any{
						"name": "credit_line",
						"title": "Credit Line",
						"type": "`$ANY`",
						"short": "Asset-specific copyright information",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "is_educational_resource",
						"title": "Is Educational Resource",
						"type": "`$BOOLEAN`",
						"short": "Whether this resource is considered to be educational",
					},
					map[string]any{
						"name": "is_multimedia_resource",
						"title": "Is Multimedia Resource",
						"type": "`$BOOLEAN`",
						"short": "Whether this resource is considered to be multimedia",
					},
					map[string]any{
						"name": "is_teacher_resource",
						"title": "Is Teacher Resource",
						"type": "`$BOOLEAN`",
						"short": "Whether this resource is considered to be educational",
					},
					map[string]any{
						"name": "lake_guid",
						"title": "Lake Guid",
						"type": "`$ANY`",
						"short": "Unique UUID of this resource in LAKE, our DAMS.",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "Name of this mobile audio file – derived from the artwork and tour titles",
					},
					map[string]any{
						"name": "transcript",
						"title": "Transcript",
						"type": "`$ANY`",
						"short": "Text transcription of the audio file",
					},
					map[string]any{
						"name": "type",
						"title": "Type",
						"type": "`$ANY`",
						"short": "Type always takes one of the following values: image, sound, text, video",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
					},
					map[string]any{
						"name": "web_url",
						"title": "Web Url",
						"type": "`$ANY`",
						"short": "URL to the audio file",
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
								"kind": "http",
								"method": "GET",
								"orig": "/mobile-sounds",
								"segments": []any{
									map[string]any{
										"lit": "mobile-sounds",
									},
								},
								"parts": []any{
									"mobile-sounds",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/sounds",
								"segments": []any{
									map[string]any{
										"lit": "sounds",
									},
								},
								"parts": []any{
									"sounds",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"mobile-sounds",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
							},
							map[string]any{
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
								"parts": []any{
									"sounds",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
					},
					map[string]any{
						"name": "web_url",
						"title": "Web Url",
						"type": "`$ANY`",
						"short": "The URL to this page on our website",
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
								"kind": "http",
								"method": "GET",
								"orig": "/static-pages",
								"segments": []any{
									map[string]any{
										"lit": "static-pages",
									},
								},
								"parts": []any{
									"static-pages",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"static-pages",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Alt Text",
						"type": "`$ANY`",
						"short": "Alternative text for the asset to describe it to people with low or no vision",
					},
					map[string]any{
						"name": "api_link",
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "artwork_ids",
						"title": "Artwork Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of the artworks associated with this asset",
					},
					map[string]any{
						"name": "artwork_titles",
						"title": "Artwork Titles",
						"type": "`$ANY`",
						"short": "Names of the artworks associated with this asset",
					},
					map[string]any{
						"name": "content",
						"title": "Content",
						"type": "`$ANY`",
						"short": "Text of or URL to the contents of this asset",
					},
					map[string]any{
						"name": "content_e_tag",
						"title": "Content E Tag",
						"type": "`$ANY`",
						"short": "Arbitrary unique identifier that changes when the binary file gets updated",
					},
					map[string]any{
						"name": "credit_line",
						"title": "Credit Line",
						"type": "`$ANY`",
						"short": "Asset-specific copyright information",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "is_educational_resource",
						"title": "Is Educational Resource",
						"type": "`$BOOLEAN`",
						"short": "Whether this resource is considered to be educational",
					},
					map[string]any{
						"name": "is_multimedia_resource",
						"title": "Is Multimedia Resource",
						"type": "`$BOOLEAN`",
						"short": "Whether this resource is considered to be multimedia",
					},
					map[string]any{
						"name": "is_teacher_resource",
						"title": "Is Teacher Resource",
						"type": "`$BOOLEAN`",
						"short": "Whether this resource is considered to be educational",
					},
					map[string]any{
						"name": "lake_guid",
						"title": "Lake Guid",
						"type": "`$ANY`",
						"short": "Unique UUID of this resource in LAKE, our DAMS.",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "type",
						"title": "Type",
						"type": "`$ANY`",
						"short": "Type always takes one of the following values: image, sound, text, video",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
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
								"kind": "http",
								"method": "GET",
								"orig": "/texts",
								"segments": []any{
									map[string]any{
										"lit": "texts",
									},
								},
								"parts": []any{
									"texts",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"texts",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "artist_titles",
						"title": "Artist Titles",
						"type": "`$ANY`",
						"short": "Names of the artists of the artworks featured in this tour's tour stops",
					},
					map[string]any{
						"name": "artwork_titles",
						"title": "Artwork Titles",
						"type": "`$ANY`",
						"short": "Names of the artworks featured in this tour's tour stops",
					},
					map[string]any{
						"name": "description",
						"title": "Description",
						"type": "`$STRING`",
						"short": "Explanation of what the tour is",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "image",
						"title": "Image",
						"type": "`$ANY`",
						"short": "The main image for the tour",
					},
					map[string]any{
						"name": "intro",
						"title": "Intro",
						"type": "`$ANY`",
						"short": "Text introducing the tour",
					},
					map[string]any{
						"name": "intro_link",
						"title": "Intro Link",
						"type": "`$ANY`",
						"short": "Link to the audio file of the introduction",
					},
					map[string]any{
						"name": "intro_transcript",
						"title": "Intro Transcript",
						"type": "`$ANY`",
						"short": "Transcript of the introduction audio to the tour",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
					},
					map[string]any{
						"name": "weight",
						"title": "Weight",
						"type": "`$NUMBER`",
						"short": "Number representing this tour's sort order",
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
								"kind": "http",
								"method": "GET",
								"orig": "/tours",
								"segments": []any{
									map[string]any{
										"lit": "tours",
									},
								},
								"parts": []any{
									"tours",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"tours",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Alt Text",
						"type": "`$ANY`",
						"short": "Alternative text for the asset to describe it to people with low or no vision",
					},
					map[string]any{
						"name": "api_link",
						"title": "Api Link",
						"type": "`$ANY`",
						"short": "REST API link for this resource",
					},
					map[string]any{
						"name": "api_model",
						"title": "Api Model",
						"type": "`$ANY`",
						"short": "REST API resource type or endpoint",
					},
					map[string]any{
						"name": "artwork_ids",
						"title": "Artwork Ids",
						"type": "`$ANY`",
						"short": "Unique identifiers of the artworks associated with this asset",
					},
					map[string]any{
						"name": "artwork_titles",
						"title": "Artwork Titles",
						"type": "`$ANY`",
						"short": "Names of the artworks associated with this asset",
					},
					map[string]any{
						"name": "content",
						"title": "Content",
						"type": "`$ANY`",
						"short": "Text of or URL to the contents of this asset",
					},
					map[string]any{
						"name": "content_e_tag",
						"title": "Content E Tag",
						"type": "`$ANY`",
						"short": "Arbitrary unique identifier that changes when the binary file gets updated",
					},
					map[string]any{
						"name": "credit_line",
						"title": "Credit Line",
						"type": "`$ANY`",
						"short": "Asset-specific copyright information",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier of this resource.",
					},
					map[string]any{
						"name": "is_educational_resource",
						"title": "Is Educational Resource",
						"type": "`$BOOLEAN`",
						"short": "Whether this resource is considered to be educational",
					},
					map[string]any{
						"name": "is_multimedia_resource",
						"title": "Is Multimedia Resource",
						"type": "`$BOOLEAN`",
						"short": "Whether this resource is considered to be multimedia",
					},
					map[string]any{
						"name": "is_teacher_resource",
						"title": "Is Teacher Resource",
						"type": "`$BOOLEAN`",
						"short": "Whether this resource is considered to be educational",
					},
					map[string]any{
						"name": "lake_guid",
						"title": "Lake Guid",
						"type": "`$ANY`",
						"short": "Unique UUID of this resource in LAKE, our DAMS.",
					},
					map[string]any{
						"name": "source_updated_at",
						"title": "Source Updated At",
						"type": "`$ANY`",
						"short": "Date and time the resource was updated in the source system",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"title": "Suggest Autocomplete All",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autosuggest` endpoint.",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"title": "Suggest Autocomplete Boosted",
						"type": "`$ANY`",
						"short": "Internal field to power the `/autocomplete` endpoint.",
					},
					map[string]any{
						"name": "timestamp",
						"title": "Timestamp",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator search index",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "The name of this resource",
					},
					map[string]any{
						"name": "type",
						"title": "Type",
						"type": "`$ANY`",
						"short": "Type always takes one of the following values: image, sound, text, video",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$ANY`",
						"short": "Date and time the record was updated in the aggregator database",
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
								"kind": "http",
								"method": "GET",
								"orig": "/videos",
								"segments": []any{
									map[string]any{
										"lit": "videos",
									},
								},
								"parts": []any{
									"videos",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"videos",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
	case "ratelimit":
		if NewRatelimitFeatureFunc != nil {
			return NewRatelimitFeatureFunc()
		}
	case "retry":
		if NewRetryFeatureFunc != nil {
			return NewRetryFeatureFunc()
		}
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	case "timeout":
		if NewTimeoutFeatureFunc != nil {
			return NewTimeoutFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
