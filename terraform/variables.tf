variable "bucket_name" {
  description = "The name of the bucket"
  type = string
}

variable "thumbnail_server_arn" {
  description = "Server to handle thumbnail creation"
  type = string
}

variable "thumbnail_execution_role_arn" {
  description = "Server to handle thumbnail creation"
  type = string
}

variable "routes" {
  description = "A list of routes"
  type = list(object({
    path = string
    authorization = string
  }))
}
