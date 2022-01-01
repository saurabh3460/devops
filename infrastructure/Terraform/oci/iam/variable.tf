variable "compartment_ocid" {
  type = string
  default = ""
}

variable "policy_description" {
  type = string
  default = "just for testing"
}

variable "policy_name" {
  type = string
  default = "test"
}

variable "policy_statements" {
  type = list(string)
  default = [
    "Allow group admin to manage all-resources in compartment sandbox",
    ]
}



