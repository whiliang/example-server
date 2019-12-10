// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-12-10 21:47:00.944611 +0800 CST m=+0.078238126

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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "login"
                ],
                "summary": "user login",
                "parameters": [
                    {
                        "description": "login param",
                        "name": "loginParam",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/apimodel.LoginParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\":\"ok\",\"code\":200,\"data\":{\"token\":\"9a22723c3589fd6922f65cbef2310b541e7721ca351f2cf4a17a9d84e6b9599e\"}}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "{\"status\":\"error\",\"code\":401,\"msg\":\"Wrong user name or password\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/logout": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "logout"
                ],
                "summary": "user logout",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\":\"ok\",\"code\":200,\"data\":\"admin: logged out\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "{\"status\":\"error\",\"code\":401,\"msg\":\"Wrong token or token expires\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "apimodel.LoginParam": {
            "type": "object",
            "required": [
                "password",
                "user"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "example": "123456"
                },
                "user": {
                    "type": "string",
                    "example": "admin"
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
	Host:        "localhost:8000",
	BasePath:    "/v1",
	Schemes:     []string{},
	Title:       "Swagger Example API",
	Description: "This is a sample server Petstore server.",
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
