
variable "region" {
  description = "AWS region"
  type        = string
  default     = "us-east-1"
}

variable "access_key" {
  description = "AWS access_key"
  type = string
  validation {
   condition     = length(var.access_key) != 0
   error_message = "Please provide a valid aws access key"
 }
 sensitive = true
}

variable "secret_key" {
  description = "AWS secret_key"
  type = string
  validation {
   condition     = length(var.secret_key) != 0
   error_message = "Please provide a valid aws secret key"
 }
 sensitive = true
}