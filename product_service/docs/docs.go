// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/products": {
            "post": {
                "description": "create new product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "create new product",
                "parameters": [
                    {
                        "description": "Create Product Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.CreateProductRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.CreateProductRequest": {
            "type": "object",
            "required": [
                "category_id",
                "condition",
                "description",
                "name",
                "price",
                "selling_by"
            ],
            "properties": {
                "category_id": {
                    "description": "required: true",
                    "type": "integer"
                },
                "condition": {
                    "description": "required: true",
                    "type": "string"
                },
                "description": {
                    "description": "required: true",
                    "type": "string"
                },
                "name": {
                    "description": "required: true",
                    "type": "string"
                },
                "price": {
                    "description": "required: true",
                    "type": "number"
                },
                "selling_by": {
                    "description": "required: true",
                    "type": "integer"
                },
                "thumbnail": {
                    "description": "required: false",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Product Service API Doc",
	Description:      "API Documentation",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
