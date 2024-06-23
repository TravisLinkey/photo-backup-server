
variable "bucket_name" {
  description = "The name of the bucket"
  type = string
}

variable "acl" {
  description = "The ACL to apply"
  type = string
  default = "private"
}
