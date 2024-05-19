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
        "/connection_info": {
            "get": {
                "description": "Get a list of connection information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[On-premise] ConnectionInfo"
                ],
                "summary": "List ConnectionInfo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Page of the connection information list.",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Row of the connection information list.",
                        "name": "row",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "UUID of the connection information.",
                        "name": "uuid",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Migration group UUID.",
                        "name": "group_uuid",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "IP address of the connection information.",
                        "name": "ip_address",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "SSH port of the connection information.",
                        "name": "ssh_port",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "User of the connection information.",
                        "name": "user",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Type of the connection information.",
                        "name": "type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully get a list of connection information.",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_model.ConnectionInfo"
                            }
                        }
                    },
                    "400": {
                        "description": "Sent bad request.",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_common.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to get a list of connection information.",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_common.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create the connection information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[On-premise] ConnectionInfo"
                ],
                "summary": "Create ConnectionInfo",
                "parameters": [
                    {
                        "description": "Connection information of the node.",
                        "name": "ConnectionInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_model.ConnectionInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully register the connection information",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_model.ConnectionInfo"
                        }
                    },
                    "400": {
                        "description": "Sent bad request.",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_common.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to register the connection information",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_common.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/connection_info/{uuid}": {
            "get": {
                "description": "Get the connection information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[On-premise] ConnectionInfo"
                ],
                "summary": "Get ConnectionInfo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UUID of the connectionInfo",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully get the connection information",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_model.ConnectionInfo"
                        }
                    },
                    "400": {
                        "description": "Sent bad request.",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_common.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to get the connection information",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_common.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update the connection information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[On-premise] ConnectionInfo"
                ],
                "summary": "Update ConnectionInfo",
                "parameters": [
                    {
                        "description": "Connection information to modify.",
                        "name": "ConnectionInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_model.ConnectionInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully update the connection information",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_model.ConnectionInfo"
                        }
                    },
                    "400": {
                        "description": "Sent bad request.",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_common.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to update the connection information",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_common.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete the connection information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[On-premise] ConnectionInfo"
                ],
                "summary": "Delete ConnectionInfo",
                "responses": {
                    "200": {
                        "description": "Successfully delete the connection information",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_model.ConnectionInfo"
                        }
                    },
                    "400": {
                        "description": "Sent bad request.",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_common.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to delete the connection information",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_common.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/import/infra/{uuid}": {
            "get": {
                "description": "Import the infra information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[Import] ImportInfra"
                ],
                "summary": "Import Infra",
                "responses": {
                    "200": {
                        "description": "Successfully saved the infra information",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_model.SavedInfraInfo"
                        }
                    },
                    "400": {
                        "description": "Sent bad request.",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_common.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to save the infra information",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_common.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/import/software/{uuid}": {
            "get": {
                "description": "Import the software information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[Import] ImportSoftware"
                ],
                "summary": "Import software",
                "responses": {
                    "200": {
                        "description": "Successfully saved the software information",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_model.SavedInfraInfo"
                        }
                    },
                    "400": {
                        "description": "Sent bad request.",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_common.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to save the software information",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_common.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/migration_group": {
            "get": {
                "description": "Get a list of migration group.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[On-premise] MigrationGroup"
                ],
                "summary": "List MigrationGroup",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Page of the migration group list.",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Row of the migration group list.",
                        "name": "row",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "UUID of the migration group.",
                        "name": "uuid",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Migration group name.",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully get a list of migration group.",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_model.MigrationGroup"
                            }
                        }
                    },
                    "400": {
                        "description": "Sent bad request.",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_common.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to get a list of migration group.",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_common.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Register the migration group.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[On-premise] MigrationGroup"
                ],
                "summary": "Register MigrationGroup",
                "parameters": [
                    {
                        "description": "migration group of the node.",
                        "name": "MigrationGroup",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_model.MigrationGroup"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully register the migration group",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_model.MigrationGroup"
                        }
                    },
                    "400": {
                        "description": "Sent bad request.",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_common.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to register the migration group",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_common.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/migration_group/check/{uuid}": {
            "get": {
                "description": "Check if SSH connection is available for each connection info in migration group. Show each status by returning connection info list.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[On-premise] MigrationGroup"
                ],
                "summary": "Check Connection MigrationGroup",
                "parameters": [
                    {
                        "description": "migration group to check SSH connection for each connection info in migration group",
                        "name": "MigrationGroup",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_model.MigrationGroup"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully checked SSH connection for the migration group",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_model.ConnectionInfo"
                        }
                    },
                    "400": {
                        "description": "Sent bad request.",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_common.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to check SSH connection for the migration group",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_common.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/migration_group/{uuid}": {
            "get": {
                "description": "Get the migration group.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[On-premise] MigrationGroup"
                ],
                "summary": "Get MigrationGroup",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UUID of the MigrationGroup",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully get the migration group",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_model.MigrationGroup"
                        }
                    },
                    "400": {
                        "description": "Sent bad request.",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_common.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to get the migration group",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_common.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update the migration group.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[On-premise] MigrationGroup"
                ],
                "summary": "Update MigrationGroup",
                "parameters": [
                    {
                        "description": "migration group to modify.",
                        "name": "MigrationGroup",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_model.MigrationGroup"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully update the migration group",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_model.MigrationGroup"
                        }
                    },
                    "400": {
                        "description": "Sent bad request.",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_common.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to update the migration group",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_common.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete the migration group.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[On-premise] MigrationGroup"
                ],
                "summary": "Delete MigrationGroup",
                "responses": {
                    "200": {
                        "description": "Successfully delete the migration group",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_model.MigrationGroup"
                        }
                    },
                    "400": {
                        "description": "Sent bad request.",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_common.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to delete the migration group",
                        "schema": {
                            "$ref": "#/definitions/github_com_cloud-barista_cm-honeybee_pkg_api_rest_common.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_cloud-barista_cm-honeybee_pkg_api_rest_common.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "github_com_cloud-barista_cm-honeybee_pkg_api_rest_model.ConnectionInfo": {
            "type": "object",
            "required": [
                "group_uuid",
                "ip_address",
                "ssh_port",
                "user",
                "uuid"
            ],
            "properties": {
                "failed_message": {
                    "type": "string"
                },
                "group_uuid": {
                    "type": "string"
                },
                "ip_address": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "private_key": {
                    "type": "string"
                },
                "public_key": {
                    "type": "string"
                },
                "ssh_port": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "user": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "github_com_cloud-barista_cm-honeybee_pkg_api_rest_model.MigrationGroup": {
            "type": "object",
            "required": [
                "name",
                "uuid"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "github_com_cloud-barista_cm-honeybee_pkg_api_rest_model.SavedInfraInfo": {
            "type": "object",
            "required": [
                "connection_uuid",
                "infra_data"
            ],
            "properties": {
                "connection_uuid": {
                    "type": "string"
                },
                "infra_data": {
                    "type": "string"
                },
                "saved_time": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "pkg_api_rest_controller.SimpleMsg": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
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
