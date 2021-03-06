// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Kiswono Prayogo"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/product_manager/product/by-id/{productId}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get product by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "product id",
                        "name": "productId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/cProductManager.Product_ById_Response"
                        }
                    }
                }
            }
        },
        "/product_manager/product/list/{limit}/{offset}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get product list",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page, default: 0",
                        "name": "offset",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "page size, default: 10",
                        "name": "limit",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/cProductManager.Product_List_Response"
                        }
                    }
                }
            }
        },
        "/product_manager/product/modify": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "update product information",
                "parameters": [
                    {
                        "description": "id and fields of product that want to be updated/deleted, inserted=0",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/cProductManager.Product_Upsert_Request"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/cProductManager.Product_Upsert_Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "cProductManager.Product_ById_Response": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "row": {
                    "type": "object",
                    "$ref": "#/definitions/model.Product"
                }
            }
        },
        "cProductManager.Product_List_Response": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "error": {
                    "type": "string"
                },
                "rows": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Product"
                    }
                }
            }
        },
        "cProductManager.Product_Upsert_Request": {
            "type": "object",
            "properties": {
                "action": {
                    "type": "string"
                },
                "product": {
                    "type": "object",
                    "$ref": "#/definitions/model.Product"
                }
            }
        },
        "cProductManager.Product_Upsert_Response": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "product": {
                    "type": "object",
                    "$ref": "#/definitions/model.Product"
                }
            }
        },
        "model.Product": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "127.0.0.1:9090",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "SF1",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
