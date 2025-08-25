resource "amenesik_beam" "ubuntu" { 
	template = "example-template"
	program  = "u2004-template"
	domain   = "openabal.com"
	region   = "any"
	category = "any"
	param    = "none"
}

resource "amenesik_beam" "ubuntu_mysql" { 
	template = "example-u2004-template"
	program  = "mysql-template"
	domain   = "openabal.com"
	region   = "any"
	category = "any"
	param    = "none"
}

resource "amenesik_beam" "ubuntu_mariadb" { 
	template = "example-u2004-template"
	program  = "mariadb-template"
	domain   = "openabal.com"
	region   = "any"
	category = "any"
	param    = "none"
}

resource "amenesik_beam" "ubuntu_postgres" { 
	template = "example-u2004-template"
	program  = "postgres-template"
	domain   = "openabal.com"
	region   = "any"
	category = "any"
	param    = "none"
}


