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
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user": {
            "post": {
                "description": "Create available user from db",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sample"
                ],
                "summary": "Create Available User",
                "operationId": "create-user",
                "parameters": [
                    {
                        "description": "user params",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/objects.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responder.ResponseBody"
                        }
                    },
                    "400": {}
                }
            }
        },
        "/user/{id}": {
            "get": {
                "description": "Get available user from db",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sample"
                ],
                "summary": "Get Available User",
                "operationId": "get-user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responder.ResponseBody"
                        }
                    }
                }
            },
            "put": {
                "description": "Update available user from db",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sample"
                ],
                "summary": "Update Available User",
                "operationId": "update-user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "user params",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/objects.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responder.ResponseBody"
                        }
                    },
                    "400": {}
                }
            },
            "delete": {
                "description": "Delete available user from db",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sample"
                ],
                "summary": "Delete Available User",
                "operationId": "delete-user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responder.ResponseBody"
                        }
                    },
                    "400": {}
                }
            }
        }
    },
    "definitions": {
        "objects.User": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string",
                    "example": "服务端会自动生成"
                },
                "description": {
                    "type": "string"
                },
                "dob": {
                    "type": "string"
                },
                "id": {
                    "type": "string",
                    "example": "服务端会自动生成"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "responder.ResponseBody": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "success": {
                    "type": "boolean"
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
	Host:        "localhost:8080",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "User Routing Service",
	Description: "API REST in Golang with Gin Framework",
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