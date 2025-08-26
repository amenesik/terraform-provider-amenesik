resource "amenesik_beam" "redhat" { 
	template = "openabal-template"
	program  = "rhel9-template"
	domain   = "openabal.com"
	region   = "any"
	category = "any"
	param    = "none"
	data     = [
		{
		path = "node.1.name"
		value = "oa-rhel-hw"
		},
		{
		path = "node.2.name"
		value = "oa-rhel-sw"
		},
		{
		path = "node.2.base"
		value = "oa-rhel-hw"
		},

	]
}

resource "amenesik_beam" "redhat_mysql" { 
	depends_on = [ amenesik_beam.redhat ]
	template = "openabal-rhel9-template"
	program  = "mysql-template"
	domain   = "openabal.com"
	region   = "any"
	category = "any"
	param    = "none"
	data     = [
		{
		path = "node.1.name"
		value = "oa-rhel-mysql-hw"
		},
		{
		path = "node.2.name"
		value = "oa-rhel-mysql-sw"
		},
		{
		path = "node.2.type"
		value = "Abal64Rhel9MySql"
		},
		{
		path = "node.2.base"
		value = "oa-rhel-mysql-hw"
		},
		{
		path = "node.1.os.distribution"
		value = "RHEL"
		},
                {
                path = "node.1.os.version"
                value = "9"
                },
		{
		path = "node.2.inxs.INXSHOST"
		value = "localhost:3306:0"
		},
		{
		path = "node.2.inxs.INXSTYPE"
		value = "MYSQL"
		},
		{
		path = "node.2.lts.LTSSQLTYPE"
		value = "MYSQL"
		},
		{
		path = "node.2.lts.LTSSQLPORT"
		value = "3306"
		}
	]
}

resource "amenesik_beam" "redhat_mysql_small" { 
	depends_on = [ amenesik_beam.redhat_mysql ]
	template = "openabal-rhel9-mysql-template"
	program  = "small-template"
	domain   = "openabal.com"
	region   = "any"
	category = "any"
	param    = "none"
	data     = [
		{
		path = "node.1.host.num_cpus"
		value = "1"
		},
		{
		path = "node.1.host.mem_size"
		value = "2G"
		},
		{
		path = "node.1.host.disk_size"
		value = "20G"
		}
	]
}

resource "amenesik_beam" "redhat_mysql_medium" { 
	depends_on = [ amenesik_beam.redhat_mysql ]
	template = "openabal-rhel9-mysql-template"
	program  = "medium-template"
	domain   = "openabal.com"
	region   = "any"
	category = "any"
	param    = "none"
	data     = [
		{
		path = "node.1.host.num_cpus"
		value = "2"
		},
		{
		path = "node.1.host.mem_size"
		value = "4G"
		},
		{
		path = "node.1.host.disk_size"
		value = "40G"
		}
	]
}

resource "amenesik_beam" "redhat_mysql_large" { 
	depends_on = [ amenesik_beam.redhat_mysql ]
	template = "openabal-rhel9-mysql-template"
	program  = "large-template"
	domain   = "openabal.com"
	region   = "any"
	category = "any"
	param    = "none"
	data     = [
		{
		path = "node.1.host.num_cpus"
		value = "4"
		},
		{
		path = "node.1.host.mem_size"
		value = "16G"
		},
		{
		path = "node.1.host.disk_size"
		value = "32G"
		}
	]
}

resource "amenesik_beam" "redhat_mysql_huge" { 
	depends_on = [ amenesik_beam.redhat_mysql ]
	template = "openabal-rhel9-mysql-template"
	program  = "huge-template"
	domain   = "openabal.com"
	region   = "any"
	category = "any"
	param    = "none"
	data     = [
		{
		path = "node.1.host.num_cpus"
		value = "8"
		},
		{
		path = "node.1.host.mem_size"
		value = "32G"
		},
		{
		path = "node.1.host.disk_size"
		value = "64G"
		}
	]
}

resource "amenesik_beam" "redhat_mariadb" { 
	depends_on = [ amenesik_beam.redhat ]
	template = "openabal-rhel9-template"
	program  = "mariadb-template"
	domain   = "openabal.com"
	region   = "any"
	category = "any"
	param    = "none"
	data     = [
		{
		path = "node.1.name"
		value = "oa-rhel-maria-hw"
		},
		{
		path = "node.2.name"
		value = "oa-rhel-maria-sw"
		},
		{
		path = "node.2.type"
		value = "Abal64Rhel9MariaDb"
		},
		{
		path = "node.2.base"
		value = "oa-rhel-maria-hw"
		},
		{
		path = "node.2.inxs.INXSHOST"
		value = "localhost:3306:0"
		},
		{
		path = "node.2.inxs.INXSTYPE"
		value = "MARIA"
		},
		{
		path = "node.2.lts.LTSSQLTYPE"
		value = "MARIA"
		},
		{
		path = "node.2.lts.LTSSQLPORT"
		value = "3306"
		}
	]
}

resource "amenesik_beam" "redhat_mariadb_small" { 
	depends_on = [ amenesik_beam.redhat_mariadb ]
	template = "openabal-rhel9-mariadb-template"
	program  = "small-template"
	domain   = "openabal.com"
	region   = "any"
	category = "any"
	param    = "none"
	data     = [
		{
		path = "node.1.host.num_cpus"
		value = "1"
		},
		{
		path = "node.1.host.mem_size"
		value = "2G"
		},
		{
		path = "node.1.host.disk_size"
		value = "20G"
		}
	]
}

resource "amenesik_beam" "redhat_mariadb_medium" { 
	depends_on = [ amenesik_beam.redhat_mariadb ]
	template = "openabal-rhel9-mariadb-template"
	program  = "medium-template"
	domain   = "openabal.com"
	region   = "any"
	category = "any"
	param    = "none"
	data     = [
		{
		path = "node.1.host.num_cpus"
		value = "2"
		},
		{
		path = "node.1.host.mem_size"
		value = "4G"
		},
		{
		path = "node.1.host.disk_size"
		value = "40G"
		}
	]
}

resource "amenesik_beam" "redhat_mariadb_large" { 
	depends_on = [ amenesik_beam.redhat_mariadb ]
	template = "openabal-rhel9-mariadb-template"
	program  = "large-template"
	domain   = "openabal.com"
	region   = "any"
	category = "any"
	param    = "none"
	data     = [
		{
		path = "node.1.host.num_cpus"
		value = "4"
		},
		{
		path = "node.1.host.mem_size"
		value = "16G"
		},
		{
		path = "node.1.host.disk_size"
		value = "32G"
		}
	]
}

resource "amenesik_beam" "redhat_mariadb_huge" { 
	depends_on = [ amenesik_beam.redhat_mariadb ]
	template = "openabal-rhel9-mariadb-template"
	program  = "huge-template"
	domain   = "openabal.com"
	region   = "any"
	category = "any"
	param    = "none"
	data     = [
		{
		path = "node.1.host.num_cpus"
		value = "8"
		},
		{
		path = "node.1.host.mem_size"
		value = "32G"
		},
		{
		path = "node.1.host.disk_size"
		value = "64G"
		}
	]
}

resource "amenesik_beam" "redhat_postgres" { 
	depends_on = [ amenesik_beam.redhat ]
	template = "openabal-rhel9-template"
	program  = "postgres-template"
	domain   = "openabal.com"
	region   = "any"
	category = "any"
	param    = "none"
	data     = [
		{
		path = "node.1.name"
		value = "oa-rhel-pgsql-hw"
		},
		{
		path = "node.2.name"
		value = "oa-rhel-pgsql-sw"
		},
		{
		path = "node.2.type"
		value = "Abal64Rhel9PostGres"
		},
		{
		path = "node.2.base"
		value = "oa-rhel-pgsql-hw"
		},
		{
		path = "node.2.inxs.INXSUSER"
		value = "postgres"
		},
		{
		path = "node.2.inxs.INXSHOST"
		value = "localhost:5432:0"
		},
		{
		path = "node.2.inxs.INXSTYPE"
		value = "PGSQL"
		},
		{
		path = "node.2.lts.LTSSQLTYPE"
		value = "PGSQL"
		},
		{
		path = "node.2.lts.LTSSQLPORT"
		value = "5432"
		}
	]
}

resource "amenesik_beam" "redhat_postgres_small" { 
	depends_on = [ amenesik_beam.redhat_postgres ]
	template = "openabal-rhel9-postgres-template"
	program  = "small-template"
	domain   = "openabal.com"
	region   = "any"
	category = "any"
	param    = "none"
	data     = [
		{
		path = "node.1.host.num_cpus"
		value = "1"
		},
		{
		path = "node.1.host.mem_size"
		value = "2G"
		},
		{
		path = "node.1.host.disk_size"
		value = "20G"
		}
	]
}

resource "amenesik_beam" "redhat_postgres_medium" { 
	depends_on = [ amenesik_beam.redhat_postgres ]
	template = "openabal-rhel9-postgres-template"
	program  = "medium-template"
	domain   = "openabal.com"
	region   = "any"
	category = "any"
	param    = "none"
	data     = [
		{
		path = "node.1.host.num_cpus"
		value = "2"
		},
		{
		path = "node.1.host.mem_size"
		value = "4G"
		},
		{
		path = "node.1.host.disk_size"
		value = "40G"
		}
	]
}

resource "amenesik_beam" "redhat_postgres_large" { 
	depends_on = [ amenesik_beam.redhat_postgres ]
	template = "openabal-rhel9-postgres-template"
	program  = "large-template"
	domain   = "openabal.com"
	region   = "any"
	category = "any"
	param    = "none"
	data     = [
		{
		path = "node.1.host.num_cpus"
		value = "4"
		},
		{
		path = "node.1.host.mem_size"
		value = "16G"
		},
		{
		path = "node.1.host.disk_size"
		value = "32G"
		}
	]
}

resource "amenesik_beam" "redhat_postgres_huge" { 
	depends_on = [ amenesik_beam.redhat_postgres ]
	template = "openabal-rhel9-postgres-template"
	program  = "huge-template"
	domain   = "openabal.com"
	region   = "any"
	category = "any"
	param    = "none"
	data     = [
		{
		path = "node.1.host.num_cpus"
		value = "8"
		},
		{
		path = "node.1.host.mem_size"
		value = "32G"
		},
		{
		path = "node.1.host.disk_size"
		value = "64G"
		}
	]
}

