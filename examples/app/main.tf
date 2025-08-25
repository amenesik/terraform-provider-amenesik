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

resource "amenesik_app" "terrabal" {
  template  = "abal64-u2004-mysql-small-template"
  program   = "terrabal"
  domain    = "openabal.com"
  region    = "france"
  category  = "amazonec2"
  param     = "4:8:16"
}

variable "ace_api_key" {
	description = "ACE secret API KEY for user account"
	type = string
	sensitive = true
}

output "one" {
  value = amenesik_app.terrabal
}

