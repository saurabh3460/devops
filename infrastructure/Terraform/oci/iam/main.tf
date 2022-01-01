terraform {
  required_providers {
    oci = {
      source = "hashicorp/oci"
    }
  }
}

resource "oci_identity_policy" "test" {
    compartment_id = var.compartment_ocid
    description = var.policy_description
    name = var.policy_name
    statements = var.policy_statements

}