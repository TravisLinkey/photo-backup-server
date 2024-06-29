
variable "bucket_name" {
  description = "The name of the bucket"
  type = string
}

variable "acl" {
  description = "The ACL to apply"
  type = string
  default = "private"
}

variable "routes" {
  description = "A list of routes"
  type = list(object({
    path = string
    authorization = string
  }))
}
