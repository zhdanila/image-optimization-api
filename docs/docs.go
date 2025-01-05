// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/api/image": {
            "post": {
                "description": "Uploads an image",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Images"
                ],
                "operationId": "upload-image",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Images to upload",
                        "name": "images",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/image-optimization-api_pkg_server.EmptyResponse"
                        }
                    }
                }
            }
        },
        "/api/image/list": {
            "get": {
                "description": "Retrieves a list of images",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Images"
                ],
                "operationId": "list-images",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/image-optimization-api_internal_service_image.ListImageResponse"
                        }
                    }
                }
            }
        },
        "/api/image/origin": {
            "get": {
                "description": "Retrieves a list of original images, excluding those with compression quality suffixes",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Images"
                ],
                "operationId": "list-origin-images",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/image-optimization-api_internal_service_image.ListOriginImageResponse"
                        }
                    }
                }
            }
        },
        "/api/image/{image_id}": {
            "get": {
                "description": "Retrieves an image by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Images"
                ],
                "operationId": "get-image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Image ID",
                        "name": "image_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Compression Quality (one of: 100, 75, 50, 25)",
                        "name": "quality",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/image-optimization-api_internal_service_image.GetImageResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "image-optimization-api_internal_service_image.GetImageResponse": {
            "type": "object",
            "properties": {
                "image": {
                    "$ref": "#/definitions/image-optimization-api_internal_service_image.Info"
                }
            }
        },
        "image-optimization-api_internal_service_image.Info": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "image-optimization-api_internal_service_image.ListImageResponse": {
            "type": "object",
            "properties": {
                "images": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/image-optimization-api_internal_service_image.Info"
                    }
                }
            }
        },
        "image-optimization-api_internal_service_image.ListOriginImageResponse": {
            "type": "object",
            "properties": {
                "keys": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "image-optimization-api_pkg_server.EmptyResponse": {
            "type": "object"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
