// Code generated by swaggo/swag. DO NOT EDIT
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
        "/admin/users": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Администратор"
                ],
                "summary": "Список пользователей",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/user.Model"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Администратор"
                ],
                "summary": "Создать пользователя",
                "parameters": [
                    {
                        "description": " ",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateUserDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.Model"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseError"
                        }
                    }
                }
            }
        },
        "/admin/users/{userId}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Администратор"
                ],
                "summary": "Удалить пользователя",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Индефикатор пользователя",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseError"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Администратор"
                ],
                "summary": "Обновить пользователя",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Индефикатор пользователя",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": " ",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateUserDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseError"
                        }
                    }
                }
            }
        },
        "/auth/refresh": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Аунтификация"
                ],
                "summary": "Обновление токена",
                "parameters": [
                    {
                        "description": "credentials",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RefreshTokenDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.Tokens"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseErrors"
                        }
                    }
                }
            }
        },
        "/auth/sign-in": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Аунтификация"
                ],
                "summary": "Авторизация",
                "parameters": [
                    {
                        "description": "credentials",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SignInDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.Model"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseErrors"
                        }
                    }
                }
            }
        },
        "/user/info": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Пользователь"
                ],
                "summary": "Информация о пользователи",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.Model"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.ParseErrorType": {
            "type": "object",
            "additionalProperties": {
                "type": "string"
            }
        },
        "api.Response": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "api.ResponseError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "api.ResponseErrors": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.ParseErrorType"
                    }
                }
            }
        },
        "auth.Tokens": {
            "type": "object",
            "properties": {
                "access": {
                    "type": "string"
                },
                "refresh": {
                    "type": "string"
                }
            }
        },
        "constants.Role": {
            "type": "string",
            "enum": [
                "teacher",
                "student",
                "parent"
            ],
            "x-enum-varnames": [
                "TeacherRole",
                "StudentRole",
                "ParentRole"
            ]
        },
        "dto.CreateUserDTO": {
            "type": "object",
            "required": [
                "address",
                "email",
                "name",
                "password",
                "patronymic",
                "phone",
                "role",
                "surname"
            ],
            "properties": {
                "address": {
                    "type": "string",
                    "maxLength": 255
                },
                "email": {
                    "type": "string",
                    "maxLength": 100
                },
                "isActive": {
                    "type": "boolean",
                    "default": false
                },
                "isSuperUser": {
                    "type": "boolean",
                    "default": false
                },
                "name": {
                    "type": "string",
                    "maxLength": 40
                },
                "password": {
                    "type": "string",
                    "maxLength": 30
                },
                "patronymic": {
                    "type": "string",
                    "maxLength": 40
                },
                "phone": {
                    "type": "string",
                    "maxLength": 30
                },
                "role": {
                    "enum": [
                        "student",
                        "teacher",
                        "parent"
                    ],
                    "allOf": [
                        {
                            "$ref": "#/definitions/constants.Role"
                        }
                    ]
                },
                "surname": {
                    "type": "string",
                    "maxLength": 40
                }
            }
        },
        "dto.RefreshTokenDTO": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "dto.SignInDTO": {
            "type": "object",
            "properties": {
                "email": {
                    "description": "@required Обязательное поле",
                    "type": "string"
                },
                "password": {
                    "description": "@required Обязательное поле",
                    "type": "string"
                },
                "remember": {
                    "type": "boolean"
                }
            }
        },
        "dto.UpdateUserDTO": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string",
                    "maxLength": 255
                },
                "email": {
                    "type": "string",
                    "maxLength": 100
                },
                "isActive": {
                    "type": "boolean",
                    "default": false
                },
                "isSuperUser": {
                    "type": "boolean",
                    "default": false
                },
                "name": {
                    "type": "string",
                    "maxLength": 40
                },
                "password": {
                    "type": "string",
                    "maxLength": 30
                },
                "patronymic": {
                    "type": "string",
                    "maxLength": 40
                },
                "phone": {
                    "type": "string",
                    "maxLength": 30
                },
                "role": {
                    "enum": [
                        "student",
                        "teacher",
                        "parent"
                    ],
                    "allOf": [
                        {
                            "$ref": "#/definitions/constants.Role"
                        }
                    ]
                },
                "surname": {
                    "type": "string",
                    "maxLength": 40
                }
            }
        },
        "user.Model": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "isActive": {
                    "type": "boolean"
                },
                "isSuperUser": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "patronymic": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "role": {
                    "enum": [
                        "student",
                        "teacher",
                        "parent"
                    ],
                    "allOf": [
                        {
                            "$ref": "#/definitions/constants.Role"
                        }
                    ]
                },
                "surname": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
