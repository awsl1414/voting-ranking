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
        "/api/activity/add": {
            "post": {
                "description": "创建活动接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Activity"
                ],
                "summary": "创建活动接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ActivityDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/player/add": {
            "post": {
                "description": "更新选手信息接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Player"
                ],
                "summary": "更新选手信息接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AddPlayerDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/player/info": {
            "post": {
                "description": "获取选手详情接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Player"
                ],
                "summary": "获取选手详情接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/player/list": {
            "post": {
                "description": "获取选手列表接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Player"
                ],
                "summary": "获取选手列表接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.PlayerListDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/player/rank": {
            "post": {
                "description": "获取排行傍列表接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Player"
                ],
                "summary": "获取排行傍列表接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.PlayerListDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/user/login": {
            "post": {
                "description": "用户登录接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "用户登录接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserLoginDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/user/register": {
            "post": {
                "description": "用户注册接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "用户注册接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserRegisterDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/vote/add": {
            "post": {
                "description": "投票接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Vote"
                ],
                "summary": "投票接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.VoteDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.ActivityDto": {
            "type": "object",
            "properties": {
                "name": {
                    "description": "活动名",
                    "type": "string"
                }
            }
        },
        "dto.AddPlayerDto": {
            "type": "object",
            "required": [
                "aid",
                "avatar",
                "declaration",
                "nickname",
                "note",
                "phone",
                "ref",
                "score"
            ],
            "properties": {
                "aid": {
                    "description": "活动id",
                    "type": "integer"
                },
                "avatar": {
                    "description": "头像",
                    "type": "string"
                },
                "declaration": {
                    "description": "宣言",
                    "type": "string"
                },
                "nickname": {
                    "description": "昵称",
                    "type": "string"
                },
                "note": {
                    "description": "备注",
                    "type": "string"
                },
                "phone": {
                    "description": "电话",
                    "type": "string"
                },
                "ref": {
                    "description": "编号",
                    "type": "string"
                },
                "score": {
                    "description": "分数",
                    "type": "integer"
                }
            }
        },
        "dto.PlayerListDto": {
            "type": "object",
            "required": [
                "aid"
            ],
            "properties": {
                "aid": {
                    "description": "活动id",
                    "type": "integer"
                }
            }
        },
        "dto.UserLoginDto": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "dto.UserRegisterDto": {
            "type": "object",
            "required": [
                "confirmPassword",
                "password",
                "username"
            ],
            "properties": {
                "confirmPassword": {
                    "description": "密码",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "dto.VoteDto": {
            "type": "object",
            "required": [
                "playerId",
                "userId"
            ],
            "properties": {
                "playerId": {
                    "description": "用户名",
                    "type": "integer"
                },
                "userId": {
                    "description": "密码",
                    "type": "integer"
                }
            }
        },
        "result.Result": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
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
	BasePath:         "",
	Schemes:          []string{},
	Title:            "活动投票系统API接口文档",
	Description:      "活动投票系统API接口文档",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
