// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/user/addUser": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "新增一个用户",
                "parameters": [
                    {
                        "description": "ok",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.UserVo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "bool"
                        }
                    }
                }
            }
        },
        "/user/del/{id}": {
            "post": {
                "summary": "删除一个用户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "bool"
                        }
                    }
                }
            }
        },
        "/user/findById/{id}": {
            "get": {
                "summary": "查询一个用户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/dao.UserInfo"
                        }
                    }
                }
            }
        },
        "/user/findByIds": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "查询多个用户Ids",
                "parameters": [
                    {
                        "description": "ok",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.Ids"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"errno\":0,\"errmsg\": \"成功\",\"data\":userInfo}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/user/queryAll": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "查询所用用户",
                "responses": {
                    "200": {
                        "description": "{\"errno\":0,\"errmsg\": \"成功\",\"data\":userInfo}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/user/queryByUserParam": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "条件查询多个用户",
                "parameters": [
                    {
                        "description": "ok",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.UserVo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"errno\":0,\"errmsg\": \"成功\",\"data\":userInfo}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/user/queryPage": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "查询多个用户",
                "parameters": [
                    {
                        "description": "ok",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dao.Page"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"errno\":0,\"errmsg\": \"成功\",\"data\":userVo}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/user/updateUser": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "编辑一个用户",
                "parameters": [
                    {
                        "description": "UserVo",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.UserVo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "bool"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.Ids": {
            "type": "object",
            "properties": {
                "ids": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "api.UserVo": {
            "type": "object",
            "properties": {
                "age": {
                    "description": "每页数据量",
                    "type": "integer",
                    "example": 23
                },
                "email": {
                    "description": "排序依据",
                    "type": "string",
                    "example": "l1277314169@163.com"
                },
                "id": {
                    "description": "可以为空",
                    "type": "integer"
                },
                "name": {
                    "description": "页码",
                    "type": "string",
                    "example": "zhangsan"
                }
            }
        },
        "dao.Page": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer"
                },
                "size": {
                    "type": "integer"
                }
            }
        },
        "dao.UserInfo": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "createTime": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updateTime": {
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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