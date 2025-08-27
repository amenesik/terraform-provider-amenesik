# Terraform Provider Amenesik

This provider plugin for the Terraform platform allows deployment and life-cycle management of complex business management applications workloads through the Multi-Cloud Amenesik Enterprise Cloud Service Federation platform.

## Resources
The Amenesik Terraform Provider plugin describes the following resources:

- App : These resource types are used to manage complete application deployment instances
- Beam : These resource types are used to manage Basmati Enhanced Application Model descriptions that are used for the description of deployment details for the preceding App resource.

## App
This resource type will be used to manage the deployment instances of your complex, multi-cloud, business application configurations.

An simple example of the use of this type of resource can be seen in the sample Terraform configuration document.

  
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
      account  = "myaccount"
      host     = "phoenix.amenesik.com"
    }
    
    resource "amenesik_app" "myapp" {
      template  = "abal64-u2004-mysql-small-template"
      program   = "myapp"
      domain    = "mydomain.com"
      category  = "amazonec2"
      region    = "france"
      param     = "4:8:16"
    }
  
    variable "ace_api_key" {
      description = "ACE secret API KEY for user account"
      type = string
      sensitive = true
    }
    
    output "myapp" {
      value = amenesik_app.myapp
    }

The terraform section indicates the required use of the involved providers, as is customary for Terraform configuration documents, in this case the amenesik provider.

The provider section defines the configuration values required for the amenesik provider:
- host : the domain name of the Amenesik Enterprise Cloud platform, usually "phoenix.amenesik.com" or "mycompany.amenesik.com"
- account : the provisioning account name on the corresponding Amenesik Enterprise Cloud platform.
- apikey : the API KEY associated with the provisioning account. This is a sensitive value and should be not be written in plain text on configuration documents.

The variable ace_api_key allows the sensitive string value of the amenesik provider API KEY to be defined through the Terraform variable management mechanisms, including environment variables, terraform command line switches and prompted user input values.

The resource section provides the values for the required parameters of an amenesik provider APP resource:
- template: The value of this property indicates the name of the BEAM document describing the details of the application configuration.
- program: The value of this property indicates the name of the application instance. It will be used in conjunction with the template name for the preparation of the instance specific derivation of the template document. It will also be used in conjunction with the value of the domain property in the composition of the fully quallified endpoint domain name.
- domain: The value of this property provides the domain name used in conjunction with the preceding program property for the composition of the fully quallified endpoint domain name.
- category: The value of this property indicates the name of the Amenesik Enterprise Cloud service provisioning type.  The property may be either a single quoted value, such as "amazonec2", or a quoted, comma separated, square braced list of alternative provisioning categories, such as "[amazonec2,googlecompute,windowsazure]". In the first case a single application state will result. In the second instance three alternative application states will be created, one for each of the specified provisioning types, but only the first will actually be started. The other two, initially idle states,  will provide alternative application fail-over states, that will be deployed and or released, as required, in response to the eventual reception, by the application controller, of "change" and "revert" action events, issued as a result of critical failure detection by the life cycle management of the Amenesik Cloud Engine.  
- region: The value of this property indicates the name of the region and will be used in conjunction with the category property value for cloud provider region selection. As for the preceding "category" property, this property may also be either a single quoted value, such as "france", or a quoted, comma separated, square braced list of alternative provisioning regions, such as "[france,germany,italy]". In the first case a single application state will result. In the second instance, three alternative application states will be created, one for each of the specified provisioning regions, with identical subsequent behaviour as described for multiple provisioning categories. The category and region properties may both specify multiple values, in which case the corresponding ordered combinations will be used. For example, with category set to "[a,b]" and region set = "[c,d]" then two application states would be created, one for category a in region c, and one for category b in region d.
- param: The value of this property allows optional application specific parameters to be passed to the application instance during its startup.

Management of the deployment of a suitably defined APP instance would be performed using the standard terraform command, as can be seen below:

    $ terraform plan
    ...
    $ terraform apply
    ...
    $ teraform show
    ...
    $ terraform destroy 
    ...

These command must be launched from the folder containing the application configuration file, unless the terraform directory change command line switch is specified.

## Beam
This resource type will be used to manage the BEAM description documents of complex, multi-cloud business applications.

The BEAM document format, being one of the outcomes of the H2020 European project known as BASMATI, is an accronym for Basmati Enhanced Application Model.

### Introduction
BEAM is a conforming derivation of the standard TOSCA document format, with the addition of standardised TAG values for the specification of information relating to the Service Level Agreement terms that describe the conditions and guarantees required for the deployment and life-cycle management of the application service.

A complete BEAM document comprises a collection of Node Type definitions, which may result from import statements, and the application's Service Template.

The Service Template comprises the collection of BEAM Tag values and the application's Topology Template.

The Topology Template comprises the collection of Node Templates describing the nature, composition and requirements of the service nodes needed by the application.

The Topology Template also comprises an optional collection of Relationship Templates describing the heirarchy of the Nodes.

Node Templates may be of the following types:
- Hardware nodes describing machines
- Software nodes describing software layers and configurations to be applied to Hardware nodes.
- Service nodes performing service oriented operations on behalf of any of the other node types.

The amenesik terraform provider BEAM resource type subsequently allows the definition, creation, management and destruction of BEAM documents.

A example of a complex Topology Template, described by a single BEAM document, as managed by the Amenesik Enterprise Cloud, is shown below. The links between the nodes represent the information provided by the collection of Relationship Templates. 

<img width="1920" height="1080" alt="image" src="https://github.com/user-attachments/assets/774c4b9a-b5a5-4e85-8f24-6f5538fe6d6e" />

This complex, yet concrete example, describes the deployment, configuration and interconnection of:
- twelve virtual machine nodes in one geographical region,
- seven virtual machine nodes in a secondary geographical region,
- accessible by both regional entry points,
- and balanced by a global traffic manager service instance.

### Example
The following Terraform configuration file shows a simple example of a BEAM resource.

    resource "amenesik_beam" "redhat" { 
    	template = "example-template"
    	program  = "rhel9-template"
    	domain   = "mydomain.com"
    	region   = "any"
    	category = "any"
    	param    = "none"
    	data     = [
    		{
    		path = "node.1.name"
    		value = "hardware-node"
    		},
      	{
    		path = "node.1.type"
    		value = "Compute"
    		},
    		{
    		path = "node.2.name"
    		value = software-node"
    		},
        {
    		path = "node.2.type"
    		value = "Database"
    		},
    		{
    		path = "node.2.base"
    		value = "hardware-node"
    		}
    	]
    }

The above example shows two BEAM (TOSCA) nodes, the first a hardware node of type Compute, and a software node, of type Database, that  uses the hardware Compute node as its base.

This can be extended further, to produce two derived templates of differing infrastructure dimensions.

Firstly a small single cpu machine with 2G of memory and 20G of disk.

    resource "amenesik_beam" "small" { 
    	template = "example-rhel9-template"
    	program  = "small-template"
    	domain   = "mydomain.com"
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

Secondly a large quad-cored cpu with 16G or memory and 100G of disk.

    resource "amenesik_beam" "small" { 
    	template = "example-rhel9-template"
    	program  = "small-template"
    	domain   = "mydomain.com"
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
    		value = "100G"
    		}      
      ]
    }

From the above examples it should be noted that the data array of the BEAM resource describes the properties and their values of the required BEAM document.

### Syntax
BEAM documents comprise ordered collections of NODES, RELATIONS and PROBES (a specialisation of the node).

A Data Path must be defined with respect to one of these three document roots or arrays:

- node . < identifier > [ . < capability > ] . < property >
- relation . node . < number > . [ hostname | contract ]
- probe . < identifier > . < property >
- type . < name > . < property >
- tag . < name >
- import
- copy . node

In the general syntax above, the term 'identifier' may be a number of the name of the node or probe. Afte the use of the 'copy' operation the term 'last' may be used to address the most recently created node.

The corresponding value will depend on the nature of the path.
- for nodes : the value will be the required value of the property.
- for relations : the value will be the required target node of the relation.
- for probes : the value will be the required value of the property.

#### Nodes
The following property names exist for all node paths outside of any capability extensions.

- name : the name of the node
- type : the usage type of the node as Compute or other Software layer definitions.
- base : the parent node of a node layering collection, where the hardware node provides a base for a subsequent chain of software node layers. 

The optional capability of a node path expression may be one of the following.

- host : the collection of host properties of the Compute node type.
- os : the collection of operating system properties of the Compute node type.
- other capability values may be used for node type specific capability parameters.

The host capability defines the following properties

- num_cpus : the number of cores or virtual cpus for the Compute node.
- mem_size : the size of the memory suffixed by M or G
- disk_size : the size of the disk, suffixed by M, G or T
- volume : the name of an attached volume
- entry : the entry point description
- hostname : the fully qualified host and domain name
- provider : the provisioning category, when specific to a node
- region : the provisioning region when specific to a node
- vlan : the vlan to which the node should be attached
- tcp_port : a TCP port to be opened in the firewall or security group
- udp_port : a UDP port to be opened in the firewall or security group
- tcp_range : a dash or comma separated range of TCP ports to be opened in the firewall or security group
- udp_range : a dash or comma separated range of UDP ports to be opened in the firewall or security group
- protocol : the network protocol
- cluster : the name of the cluster for a container compute node
- namespace : the name of the namespace for a container compute node
 
The os capability defines the following properties

- architecture : describes the hardware architecture such as x86_64
- type : indicates the operating system family as linux or windows
- distribution : the name of the distribution as WINDOWS, UBUNTU or RHEL
- version : the version of the specified distribution such as 20.04 or 9

The properties of the software node types are type specific and require consultation of the product capabilites on the amenesik web site.

The following structure describes the properties that would be required for a typical combination of hardware and software nodes

- node.1.name
- node.1.type
- node.1.host.num_cpus
- node.1.host.mem_size
- node.1.host.disk_size
- node.1.host.hostname
- node.2.name
- node.2.type
- node.2.base
- node.2.capability.property1
- node.2.capability.propertyN

In real world situations a a large number of nodes would be defined each with their own specific collections of capabilities and their associated properties.

The complex example shown above, requires 22 hardware (Compute) nodes, 19 software nodes of four different classes (LDAP, TOMCAT, APACHE, HAPROXY) and roughly 50 relations. In addition a variety of probes would be required for both hardware and software operation and failover monitoring.  The resulting BEAM document would naturally be correspondingly complex.

### Probes
The following properties are defined for the monitoring probes.

- metric : the definition of the type of information, its collection frequency and its means of collection. These may be defined per account by the Amenesik Enterprise Cloud.
- condition : the nature of the comparaison with the threshold value (eq, gr, ls, ge, le, ne)
- threshold : the threshold value which when reached requires remediative "penalty" action to be engaged.
- type : the nature of the remediation action invocation (one of OCCISCRIPT, BASH, PYTHON)
- nature : the purpose or nature of the action (one of penalty, reward, both). When "reward" or "both" then the actio will be engaged before threshold is reached.
- behaviour : the name of the OCCI, BASH or PYTHON script describing the subsequent action.

### Relations
A relation is required to be defined when a secondary (target) node construction (hardware and software elements) requires autotamtion of its connection to a primary (source) node construction (hardware and software elements) during the deployment of infrastructural APP resources.

The target node is said to receive connection information from the source node during the configuration (or construction) phase of its life cycle.

The following Data definition of a BEAM resource shows the properties required to describe and establish such a connection, always defined between the hardware nodes.

    resource "amenesik_beam" "small" { 
      ...
    	data     = [
        { path = "node.1.type" value = "Compute" }
        { path = "node.2.type" value = "Database" }
        { path = "node.3.type" value = "Compute" }
        { path = "node.4.type" value = "WebServer" }
    		{
    		path = "relation . node . 1 . hostname"
    		value = "node.3"
    		}      
      ]
    }

In this example the Compute node of the WebServer configuration will receive the hostname of the Compute node of the Database configuration.

In most cases, the subject of the relationship will be the hostname of the source. 

In certain cases, especially concerning automation of load balancing scenarios, it is necessary that the subject of the relationship be the contract identifier of the source, facilitating replication of the source, by the target, in accordance with percieved load.

### Tags 
Tags that are used to define service deployment conditions and other metadata, may be defined through the BEAM resource in terms of theier name and value.

The following tags are currently defined for BEAM documents, and other than the Probe tag, should all be self-explanatory:

- Title : defines the title of the BEAM resource
- SubTitle : defines the sub title of the BEAM resource
- Author : defines the author of the BEAM resource
- Version : defines the version of the BEAM resource
- Date : defines the date of creation or modification of the document
- Account : defines the default, or global, provisioning account
- Provider : defines the default, or global, provisioning category
- Zone : defines the default, or global, provisioning zone
- Domain : defines the default, or global, domain name
- Probe : defines the name of an Amenesik Enterprise Cloud Probe definition.

The following configuration document snippet shows an example of tag definitions.

    resource "amenesik_beam" "small" { 
      ...
    	data     = [
        ...
        { path = tag.Title" value = "My Document Title" }
        { path = "tag.Author" value = "The document Author name" }
        { path = "tag.Probe" value = "memory-free" }
        ...
      ]
    }

### Types
The Node types that are used to define the behaviour of software layer nodes may be defined through the BEAM resource in terms of the following properties:

- name : the terminal name portion of the node type.
- create : the public, web fetchable action script to be fetched and launched when a node is created.
- start : the public, web fetchable action script to be fetched and launched when a node is started.
- stop : the public, web fetchable action script, to be fetched and launched when a node is stopped.
- save : the public, web fetchable action script, to be fetched and launched when a node is saved.
- delete : the public, web fetchable action script, to be fetched and launched when a node is deleted.

### Imports
Node types may be imported instead of being defined in BEAM documents. This encourages reusability.

The following configuration document snippet shows an example of tag definitions.

    resource "amenesik_beam" "small" { 
      ...
    	data     = [
        ...
        { path = import" value = "Database" }
        { path = "import" value = "WebServer" }
        ...
      ]
    }

## Complete Example
The following configuration document shows the creation of a complex multi layer load balenced web application scenario with 6 application servers and two database servers.

    resource "amenesik_beam" "mybeam" {
            template = "template"
            program  = "mybeam-template"
            domain   = "openabal.com"
            region   = "any"
            category = "any"
            param    = "none"
            data     = [
                    # beam document tags
                    { path = "tag.Title", value="My New Beam Document" },
                    { path = "tag.SubTitle", value="Generated by Terraform" },
                    { path = "tag.Author", value="Iain James Marshall" },
                    { path = "tag.Probe", value="memory-free" },
                    { path = "tag.Probe", value="disk-free" },
    
                    # beam document node: database hardware
                    { path = "node.1.name", value = "dbhwa" },
                    { path = "node.1.type", value = "Compute" },
                    { path = "node.1.host.num_cpus", value = "8" },
                    { path = "node.1.host.mem_size", value = "32G" },
                    { path = "node.1.host.disk_size", value = "i110G" },
                    { path = "node.1.os.distribution", value = "UBUNTU" },
                    { path = "node.1.os.version", value = "20.04" },
    
                    # beam document node: database software
                    { path = "node.2.name", value = "dbswa" },
                    { path = "node.2.base", value = "dbhwa" },
                    { path = "node.2.type", value = "Database" },
                    { path = "node.2.db.HOST", value = "localhost:3306:1" },
                    { path = "node.2.db.USER", value = "myuser" },
                    { path = "node.2.db.PASS", value = "mypass" },
                    { path = "node.2.db.BASE", value = "mybase" },
    
                    # beam document node : duplicate database hardware
                    { path = "copy.node", "value" = "dbhwa" },
                    { path = "node.last.name", value = "dbhwb" },
    
                    # beam document node: duplicate database software
                    { path = "copy.node", "value" = "dbswa" },
                    { path = "node.last.name", value = "dbswb" },
                    { path = "node.last.base", value = "dbhwb" },
    
                    # beam document node: web server hardware
                    { path = "copy.node", "value" = "dbhwa" },
                    { path = "node.last.name", value = "wshwa" },
    
                    # beam document node: web server software
                    { path = "copy.node", "value" = "dbswa" },
                    { path = "node.last.name", value = "wsswa" },
                    { path = "node.last.base", value = "wshwa" },
                    { path = "node.last.type", value = "WebServer" },
    
                    # beam document node: web server hardware copy
                    { path = "copy.node", "value" = "wshwa" },
                    { path = "node.last.name", value = "wshwb" },
    
                    # beam document node: web server software copy
                    { path = "copy.node", "value" = "wsswa" },
                    { path = "node.last.name", value = "wsswb" },
                    { path = "node.last.base", value = "wshwb" },
    
                    # beam document node: web server hardware copy
                    { path = "copy.node", "value" = "wshwa" },
                    { path = "node.last.name", value = "wshwc" },
    
                    # beam document nodes 6a : web server software copy
                    { path = "copy.node", "value" = "wsswa" },
                    { path = "node.last.name", value = "wsswc" },
                    { path = "node.last.base", value = "wshwc" },
    
                    # beam document node: web server hardware copy
                    { path = "copy.node", "value" = "wshwa" },
                    { path = "node.last.name", value = "wshwd" },
    
                    # beam document node: web server software copy
                    { path = "copy.node", "value" = "wsswa" },
                    { path = "node.last.name", value = "wsswd" },
    
                    # beam document node: load balancer hardware
                    { path = "copy.node", "value" = "dbhw" },
                    { path = "node.last.name", value = "lbhwa" },
    
                    # beam document node: load balancer software
                    { path = "copy.node", "value" = "dbsw" },
                    { path = "node.last.name", value = "lbswa" },
                    { path = "node.last.base", value = "lbhwa" },
                    { path = "node.last.type", value = "LoadBalancer" },
    
                    # beam document node: web server hardware copy
                    { path = "copy.node", "value" = "wshwa" },
                    { path = "node.last.name", value = "wshwe" },
    
                    # beam document node: web server software copy
                    { path = "copy.node", "value" = "wsswa" },
                    { path = "node.last.name", value = "wsswe" },
                    { path = "node.last.base", value = "wshwe" },
    
                    # beam document node: web server hardware copy
                    { path = "copy.node", "value" = "wshwa" },
                    { path = "node.last.name", value = "wshwf" },
    
                    # beam document node: web server software copy
                    { path = "copy.node", "value" = "wsswa" },
                    { path = "node.last.name", value = "wsswf" },
                    { path = "node.last.base", value = "wshwf" },
    
                    # beam document node: load balancer hardware copy
                    { path = "copy.node", "value" = "lbhwa" },
                    { path = "node.last.name", value = "lbhwb" },
    
                    # beam document node: load balancer software copy
                    { path = "copy.node", "value" = "lbswa" },
                    { path = "node.last.name", value = "lbswb" },
                    { path = "node.last.base", value = "lbhwb" },
                    
                    # beam document node: load balancer hardware copy
                    { path = "copy.node", "value" = "lbhwa" },
                    { path = "node.last.name", value = "lbhwc" },
    
                    # beam document node: load balancer software copy
                    { path = "copy.node", "value" = "lbswa" },
                    { path = "node.last.name", value = "lbswc" },
                    { path = "node.last.base", value = "lbhwc" },
    
                    # beam document relations
                    { path = "relation.node.dbhwa.hostname", value="node.wshwa" },
                    { path = "relation.node.dbhwa.hostname", value="node.wshwb" },
                    { path = "relation.node.dbhwa.hostname", value="node.wshwc" },
                    { path = "relation.node.dbhwa.hostname", value="node.wshwd" },
                    { path = "relation.node.dbhwa.hostname", value="node.wshwe" },
                    { path = "relation.node.dbhwa.hostname", value="node.wshwf" },
                    { path = "relation.node.dbhwb.hostname", value="node.wshwa" },
                    { path = "relation.node.dbhwb.hostname", value="node.wshwb" },
                    { path = "relation.node.dbhwb.hostname", value="node.wshwc" },
                    { path = "relation.node.dbhwb.hostname", value="node.wshwd" },
                    { path = "relation.node.dbhwb.hostname", value="node.wshwe" },
                    { path = "relation.node.dbhwb.hostname", value="node.wshwf" },
                    { path = "relation.node.wshwa.hostname", value="node.lbhwa" },
                    { path = "relation.node.wshwb.hostname", value="node.lbhwa" },
                    { path = "relation.node.wshwc.hostname", value="node.lbhwa" },
                    { path = "relation.node.wshwd.hostname", value="node.lbhwa" },
                    { path = "relation.node.wshwe.hostname", value="node.lbhwa" },
                    { path = "relation.node.wshwf.hostname", value="node.lbhwa" },
                    { path = "relation.node.wshwa.hostname", value="node.lbhwb" },
                    { path = "relation.node.wshwb.hostname", value="node.lbhwb" },
                    { path = "relation.node.wshwc.hostname", value="node.lbhwb" },
                    { path = "relation.node.wshwd.hostname", value="node.lbhwb" },
                    { path = "relation.node.wshwe.hostname", value="node.lbhwb" },
                    { path = "relation.node.wshwf.hostname", value="node.lbhwb" },
                    { path = "relation.node.lbhwa.hostname", value="node.lbhwc" },
                    { path = "relation.node.lbhwb.hostname", value="node.lbhwc" },
            ]
    }

The processing of this BEAM resource using Terraform Apply would result in the following BEAM Topology.

<img width="1920" height="1080" alt="image" src="https://github.com/user-attachments/assets/371fddb9-2b96-420c-bd9b-4d0909f2466e" />





