terraform {
  required_providers {
    amenesik = {
      source  = "hashicorp.com/edu/amenesik"
    }
  }
  required_version = ">= 1.1.0"
}

provider "amenesik" {
  apikey   = var.ace_api_key
  account  = "amenesik"
  host     = "phoenix.amenesik.com"
}

data "amenesik_beam" "edu" { }

variable "ace_api_key" {
	description = "ACE secret API KEY for user account"
	type = string
	sensitive = true
}

output "edu_beam" {
  value = data.amenesik_beam.edu
}

