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
        "/admin/subject": {
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
                "summary": "Список предметов",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/subject.Model"
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
                "summary": "Создать предмет",
                "parameters": [
                    {
                        "description": " ",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateSubjectDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/subject.Model"
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
        "/admin/subject/{id}": {
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
                "summary": "Получить предмет",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Индефикатор предмета",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/subject.Model"
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
                "summary": "Удалить предмет",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Индефикатор преподавателя",
                        "name": "id",
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
                "summary": "Обновить предмет",
                "parameters": [
                    {
                        "description": " ",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateSubjectDTO"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Индефикатор предмета",
                        "name": "id",
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
            }
        },
        "/admin/teachers": {
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
                "summary": "Список преподователей",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/teacher.Model"
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
                "summary": "Создать преподователя",
                "parameters": [
                    {
                        "description": " ",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateTeacherDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/teacher.Model"
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
        "/admin/teachers/{id}": {
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
                "summary": "Получить преподователя",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Индефикатор преподавателя",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/teacher.Model"
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
                "summary": "Удалить преподавателя",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Индефикатор преподавателя",
                        "name": "id",
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
                "summary": "Обновить прподавателя",
                "parameters": [
                    {
                        "description": " ",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateTeacherDTO"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Индефикатор преподавателя",
                        "name": "id",
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
        "/auth/user-info": {
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
                    "Аунтификация"
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
                            "$ref": "#/definitions/api.ResponseErrors"
                        }
                    }
                }
            }
        },
        "/parent": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Родитель"
                ],
                "summary": "Создать",
                "parameters": [
                    {
                        "description": " ",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateParentDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/parent.Model"
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
        "/subjects": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Предметы"
                ],
                "summary": "Список предметов",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/subject.Model"
                            }
                        }
                    }
                }
            }
        },
        "/subjects/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Предметы"
                ],
                "summary": "Получить предмет",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Индефикатор предмета",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/subject.Model"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
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
        "dto.CreateParentDTO": {
            "type": "object",
            "properties": {
                "guardian": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/dto.CreateUserDTO"
                }
            }
        },
        "dto.CreateSubjectDTO": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "title": {
                    "type": "string",
                    "maxLength": 40
                }
            }
        },
        "dto.CreateTeacherDTO": {
            "type": "object",
            "properties": {
                "subject": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/subject.Model"
                    }
                },
                "user": {
                    "$ref": "#/definitions/dto.CreateUserDTO"
                }
            }
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
                    "maxLength": 30,
                    "minLength": 8
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
            "required": [
                "token"
            ],
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "dto.SignInDTO": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 100
                },
                "password": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 8
                },
                "remember": {
                    "type": "boolean",
                    "default": false
                }
            }
        },
        "dto.UpdateSubjectDTO": {
            "type": "object",
            "properties": {
                "title": {
                    "type": "string",
                    "maxLength": 40
                }
            }
        },
        "dto.UpdateTeacherDTO": {
            "type": "object",
            "properties": {
                "subject": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/subject.Model"
                    }
                },
                "user": {
                    "$ref": "#/definitions/dto.UpdateUserDTO"
                },
                "userId": {
                    "type": "string"
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
                    "maxLength": 30,
                    "minLength": 8
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
        "parent.Model": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "guardian": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/user.Model"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "subject.Model": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "teacher.Model": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "subject": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/subject.Model"
                    }
                },
                "updatedAt": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/user.Model"
                },
                "userId": {
                    "type": "string"
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
                    "$ref": "#/definitions/constants.Role"
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
