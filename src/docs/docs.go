// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-04-02 23:33:53.0740488 +0800 CST m=+0.046020501

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
        "termsOfService": "https://:).moe",
        "contact": {
            "name": "Kasper",
            "url": "https://kasper.moe",
            "email": "me@Kasper.moe"
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
        "/query": {
            "post": {
                "description": "选取特定的行人",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "选取特定的行人",
                "operationId": "Query",
                "parameters": [
                    {
                        "type": "file",
                        "description": "行人图片",
                        "name": "files",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\": \"uploaded successfully\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"error\": {}}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "415": {
                        "description": "{\"error\": {}}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "{\"error\": {}}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/search": {
            "post": {
                "description": "查找特定的行人",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "查找特定的行人",
                "operationId": "Search",
                "parameters": [
                    {
                        "type": "file",
                        "description": "待查找视频",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\": \"searched successfully\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"error\": {}}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "415": {
                        "description": "{\"error\": {}}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "{\"error\": {}}",
                        "schema": {
                            "type": "string"
                        }
                    }
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
	Host:        "127.0.0.1:8080",
	BasePath:    "/api",
	Schemes:     []string{},
	Title:       "ReID API",
	Description: "This is a ReID's server API.",
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
