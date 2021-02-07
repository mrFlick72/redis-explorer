provider "aws" {}

resource "aws_dynamodb_table" "redis-connections-dynamodb-table" {
  name           = "RedisExplorerConnections"
  hash_key       = "ConnectionName"

  attribute {
    name = "Username"
    type = "S"
  }

  attribute {
    name = "Password"
    type = "S"
  }

  tags = {
    Name        = "RedisExplorerConnections"
    environment = var.tags_environment
    scope = var.tags_scope
    role = var.tags_role
    project = var.tags_project
  }
}