resource "amenesik_beam" "redhat" { 
	template = "example-template"
	program  = "rhel9-template"
	domain   = "openabal.com"
	region   = "any"
	category = "any"
	param    = "none"
}

resource "amenesik_beam" "redhat_mysql" { 
	template = "example-rhel9-template"
	program  = "mysql-template"
	domain   = "openabal.com"
	region   = "any"
	category = "any"
	param    = "none"
}

resource "amenesik_beam" "redhat_mariadb" { 
	template = "example-rhel9-template"
	program  = "mariadb-template"
	domain   = "openabal.com"
	region   = "any"
	category = "any"
	param    = "none"
}

resource "amenesik_beam" "redhat_postgres" { 
	template = "example-rhel9-template"
	program  = "postgres-template"
	domain   = "openabal.com"
	region   = "any"
	category = "any"
	param    = "none"
}

