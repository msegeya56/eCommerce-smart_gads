// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "after user login user will seen this page with user informations",
                "tags": [
                    "Home"
                ],
                "summary": "api for user home page",
                "operationId": "User Home",
                "responses": {
                    "200": {
                        "description": "Welcome to home !"
                    }
                }
            }
        },
        "/admin": {
            "get": {
                "tags": [
                    "Admin Home"
                ],
                "summary": "api admin home",
                "operationId": "AdminHome",
                "responses": {
                    "200": {
                        "description": "Welcome to Admin Home",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/admin/login": {
            "post": {
                "tags": [
                    "Admin Login"
                ],
                "summary": "api for admin to login",
                "operationId": "AdminLogin",
                "parameters": [
                    {
                        "description": "inputs",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.LoginData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successfully logged in",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Failed to login",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Generate JWT failure",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/admin/products": {
            "put": {
                "tags": [
                    "Admin Product"
                ],
                "summary": "api for admin to delete a product",
                "operationId": "UpdateProduct",
                "parameters": [
                    {
                        "description": "inputs",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.DeleteProductReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfuly deleted product",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Missing or invalid input",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Missing or invalid input",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "Admin Brand"
                ],
                "summary": "api for admin to add a parent brand",
                "operationId": "AddBrand",
                "parameters": [
                    {
                        "description": "inputs",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ReqProduct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfuly added a new brand in database",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Missing or invalid entry",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/admin/users": {
            "get": {
                "tags": [
                    "Admin User"
                ],
                "summary": "api for admin to list users",
                "operationId": "ListUsers",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page Number",
                        "name": "page_number",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Count Of Order",
                        "name": "count",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List user successful",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "failed to get all users",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/brands": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Product brands"
                ],
                "summary": "api for admin to list all brands",
                "operationId": "ListBrands-admin",
                "responses": {
                    "200": {
                        "description": "Successfuly listed all brands",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Failed to get brands",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "User Login"
                ],
                "summary": "api for user login",
                "operationId": "UserLogin",
                "parameters": [
                    {
                        "description": "Input Fields",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.LoginData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login successful",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Missing or invalid entry",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "failed to send OTP",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/logout": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "user can logout",
                "tags": [
                    "User Logout"
                ],
                "summary": "api for user to logout",
                "operationId": "UserLogout",
                "responses": {
                    "200": {
                        "description": "Log out successful"
                    }
                }
            }
        },
        "/otp-verify": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "User OTP verification"
                ],
                "summary": "api for user otp verification",
                "operationId": "UserOtpVerify",
                "parameters": [
                    {
                        "description": "Input Fields",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.OTPVerify"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login successful",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Missing or invalid entry",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "failed to send OTP",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/products": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "User Products"
                ],
                "summary": "api for user to list all products",
                "operationId": "ListProducts-User",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page Number",
                        "name": "page_number",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Count Of Order",
                        "name": "count",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Product listed successfuly",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "failed to get all products",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "User Signup"
                ],
                "summary": "api for register user",
                "operationId": "UserSignUp",
                "parameters": [
                    {
                        "description": "Input Fields",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.SignupUserData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Account created successfuly"
                    },
                    "400": {
                        "description": "invalid entry"
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Admin": {
            "type": "object",
            "required": [
                "password"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 5
                },
                "updated_at": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string",
                    "maxLength": 15,
                    "minLength": 4
                }
            }
        },
        "request.DeleteProductReq": {
            "type": "object",
            "properties": {
                "Prod_id": {
                    "type": "integer"
                }
            }
        },
        "request.LoginData": {
            "type": "object",
            "required": [
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 5
                },
                "user_name": {
                    "type": "string",
                    "maxLength": 15,
                    "minLength": 3
                }
            }
        },
        "request.OTPVerify": {
            "type": "object",
            "required": [
                "otp",
                "user_id"
            ],
            "properties": {
                "otp": {
                    "type": "string",
                    "maxLength": 8,
                    "minLength": 4
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "request.ReqProduct": {
            "type": "object",
            "required": [
                "brand_id",
                "description",
                "image",
                "price",
                "product_name"
            ],
            "properties": {
                "brand_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string",
                    "maxLength": 1000,
                    "minLength": 10
                },
                "image": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "product_name": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 3
                }
            }
        },
        "request.SignupUserData": {
            "type": "object",
            "required": [
                "age",
                "confirm_password",
                "email",
                "first_name",
                "last_name",
                "password",
                "phone",
                "user_name"
            ],
            "properties": {
                "age": {
                    "type": "integer"
                },
                "confirm_password": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                },
                "last_name": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 1
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string",
                    "maxLength": 10,
                    "minLength": 10
                },
                "user_name": {
                    "type": "string",
                    "maxLength": 15,
                    "minLength": 3
                }
            }
        },
        "request.UpdateProductReq": {
            "type": "object",
            "properties": {
                "brand_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "product_name": {
                    "type": "string"
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "errors": {},
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
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
