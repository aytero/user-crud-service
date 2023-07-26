package docs

import "github.com/swaggo/swag"

const docTemplate = `
{
  "openapi": "3.0.1",
  "info": {
    "title": "User CRUD",
    "version": "1.0"
  },
  "paths": {
    "/users/list": {
      "get": {
        "tags": [
          "user-controller"
        ],
        "operationId": "getUsersList",
        "responses": {
          "200": {
            "description": "ok",
            "content": {
              "application/json": {
                "schema": {
                  "type": "array",
                  "items": {
                    "$ref": "#/components/schemas/UserPublicInfoDto"
                  }
                }
              }
            }
          },
          "400": {
            "description": "bad request",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ErrorTextResponse"
                }
              }
            }
          }
        }
      }
    },
    "/add/users": {
      "post": {
        "tags": [
          "user-controller"
        ],
        "operationId": "addUsers",
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "type": "array",
                "items": {
                  "$ref": "#/components/schemas/CreateUserDto"
                }
              }
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "ok",
            "content": {
              "application/json": {
                "schema": {
                  "type": "array",
                  "items": {
                    "$ref": "#/components/schemas/UserDto"
                  }
                }
              }
            }
          },
          "400": {
            "description": "bad request",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ErrorTextResponse"
                }
              }
            }
          }
        }
      }
    },
    "/user/{id}": {
      "get": {
        "tags": [
          "user-controller"
        ],
        "operationId": "getUser",
        "parameters": [
          {
            "name": "user_id",
            "in": "path",
            "description": "User identifier",
            "required": true,
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/UserDto"
                }
              }
            }
          },
          "400": {
            "description": "bad request",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ErrorTextResponse"
                }
              }
            }
          },
          "404": {
            "description": "not found",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ErrorTextResponse"
                }
              }
            }
          }
        }
      },
      "patch": {
        "tags": [
          "user-controller"
        ],
        "operationId": "updateUser",
        "parameters": [
          {
            "name": "user_id",
            "in": "path",
            "description": "User identifier",
            "required": true,
            "schema": {
              "type": "string"
            }
          }
        ],
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/UpdateUserDto"
              }
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "ok",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/UserDto"
                }
              }
            }
          },
          "400": {
            "description": "bad request",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ErrorTextResponse"
                }
              }
            }
          },
          "404": {
            "description": "not found",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ErrorTextResponse"
                }
              }
            }
          }
        }
      }
    },

    "/login": {
      "post": {
        "tags": [
          "user-controller"
        ],
        "operationId": "loginUser",
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/LoginUserRequest"
              }
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "ok",
            "content": {
              "application/json": {
                "schema": {
                  "type": "array",
                  "items": {
                    "$ref": "#/components/schemas/LoginUserResponse"
                  }
                }
              }
            }
          },
          "400": {
            "description": "bad request",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ErrorTextResponse"
                }
              }
            }
          },
          "401": {
            "description": "unauthorized",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ErrorTextResponse"
                }
              }
            }
          },
          "501": {
            "description": "internal server error",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ErrorTextResponse"
                }
              }
            }
          }
        }
      }
    },

    "/delete/user/{id}": {
      "delete": {
        "tags": [
          "user-controller"
        ],
        "operationId": "deleteUser",
        "parameters": [
          {
            "name": "user_id",
            "in": "path",
            "description": "User identifier",
            "required": true,
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "content": {
              "application/json": {
                "schema": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "description": "bad request",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ErrorTextResponse"
                }
              }
            }
          },
          "404": {
            "description": "not found",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ErrorTextResponse"
                }
              }
            }
          }
        }
      }
    }
  },
  "components": {
    "schemas": {
      "ErrorTextResponse": {
        "type": "object",
        "properties": {
          "error": {
            "type": "string"
          }
        }
      },
      "UserDto": {
        "required": [
          "id",
          "password",
          "isActive",
          "balance",
          "age",
          "name",
          "gender",
          "company",
          "email",
          "phone",
          "address",
          "about",
          "registered",
          "latitude",
          "longitude",
          "tags",
          "friends",
          "data"
        ],
        "type": "object",
        "properties": {
          "id": {
            "type": "string"
          },
          "password": {
            "type": "string"
          },
          "isActive": {
            "type": "string"
          },
          "balance": {
            "type": "string"
          },
          "age": {
            "type": "string",
            "format": "int32"
          },
          "name": {
            "type": "string"
          },
          "gender": {
            "type": "string"
          },
          "company": {
            "type": "string"
          },
          "email": {
            "type": "string"
          },
          "phone": {
            "type": "string"
          },
          "address": {
            "type": "string"
          },
          "about": {
            "type": "string"
          },
          "registered": {
            "type": "string",
            "format": "date-time"
          },
          "latitude": {
            "type": "string",
            "format": "float64"
          },
          "longitude": {
            "type": "string",
            "format": "float64"
          },
          "tags": {
            "type": "array",
            "items": {
              "type": "string"
            }
          },
          "friends": {
            "type": "array",
            "items": {
              "$ref": "#/components/schemas/FriendsDto"
            }
          },
          "data": {
            "type": "string"
          }
        }
      },
      "FriendsDto": {
        "required": [
          "id",
          "name"
        ],
        "type": "object",
        "properties": {
          "id": {
            "type": "string",
            "format": "int64"
          },
          "name": {
            "type": "string"
          }
        }
      },
      "UserPublicInfoDto": {
        "required": [
          "id",
          "isActive",
          "balance",
          "age",
          "name",
          "gender",
          "company",
          "email",
          "phone",
          "address",
          "about",
          "registered",
          "latitude",
          "longitude",
          "tags",
          "friends",
          "data"
          ],
         "type": "object",
         "properties": {
          "id": {
            "type": "string"
          },
          "isActive": {
            "type": "string"
          },
          "balance": {
            "type": "string"
          },
          "age": {
            "type": "string",
            "format": "int32"
          },
          "name": {
            "type": "string"
          },
          "gender": {
            "type": "string"
          },
          "company": {
            "type": "string"
          },
          "email": {
            "type": "string"
          },
          "phone": {
            "type": "string"
          },
          "address": {
            "type": "string"
          },
          "about": {
            "type": "string"
          },
          "registered": {
            "type": "string",
            "format": "date-time"
          },
          "latitude": {
            "type": "string",
            "format": "float64"
          },
          "longitude": {
            "type": "string",
            "format": "float64"
          },
          "tags": {
            "type": "array",
            "items": {
              "type": "string"
            }
          },
          "friends": {
            "type": "array",
            "items": {
              "$ref": "#/components/schemas/FriendsDto"
            }
          },
          "data": {
            "type": "string"
          }
          }
      },
      "CreateUserRequest": {
        "required": [
          "users"
        ],
        "type": "object",
        "properties": {
          "users": {
            "type": "array",
            "items": {
              "$ref": "#/components/schemas/UserDto"
            }
          }
        }
      },
      "CreateUserResponse": {
        "required": [
          "users"
        ],
        "type": "object",
        "properties": {
          "users": {
            "type": "array",
            "items": {
              "$ref": "#/components/schemas/UserDto"
            }
          }
        }
      },

      "UpdateUserDto": {
        "required": [
          "id"
        ],
        "type": "object",
        "properties": {
          "id": {
            "type": "string"
          },
          "password": {
            "type": "string"
          },
          "isActive": {
            "type": "string"
          },
          "balance": {
            "type": "string"
          },
          "age": {
            "type": "string",
            "format": "int32"
          },
          "name": {
            "type": "string"
          },
          "gender": {
            "type": "string"
          },
          "company": {
            "type": "string"
          },
          "email": {
            "type": "string"
          },
          "phone": {
            "type": "string"
          },
          "address": {
            "type": "string"
          },
          "about": {
            "type": "string"
          },
          "registered": {
            "type": "string",
            "format": "date-time"
          },
          "latitude": {
            "type": "string",
            "format": "float64"
          },
          "longitude": {
            "type": "string",
            "format": "float64"
          },
          "tags": {
            "type": "array",
            "items": {
              "type": "string"
            }
          },
          "friends": {
            "type": "array",
            "items": {
              "$ref": "#/components/schemas/FriendsDto"
            }
          },
          "data": {
            "type": "string"
          }
        }
      },
      "LoginUserRequest": {
        "required": [
          "id",
          "password"
        ],
        "type": "object",
        "properties": {
          "id": {
            "type": "string"
          },
          "password": {
            "type": "string"
          }
        }
      },
      "LoginUserResponse": {
        "required": [
          "data"
        ],
        "type": "object",
        "properties": {
          "data": {
            "type": "string"
          }
        }
      },
      "BadRequestResponse": {
        "type": "object"
      },
      "NotFoundResponse": {
        "type": "object"
      }
    }
  }
}
`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Balance API",
	Description:      "Service for interactions with user's money accounts",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
