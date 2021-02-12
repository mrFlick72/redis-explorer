provider "aws" {}

resource "aws_dynamodb_table" "redis-connections-dynamodb-table" {
  name = "RedisExplorerConnections"
  billing_mode = "PAY_PER_REQUEST"
  hash_key = "ConnectionId"

  attribute {
    name = "ConnectionId"
    type = "S"
  }

  tags = {
    Name = "RedisExplorerConnections"
    environment = var.tags_environment
    scope = var.tags_scope
    role = var.tags_role
    project = var.tags_project
  }
}