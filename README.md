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

From the above examples it should be noted that the data array of the BEAM resource describes the properties and their values of the required BEAM document.

BEAM documents comprise ordered collections of NODES, RELATIONS and PROBES (a specialisation of the node).

A Data Path must be defined with respect to one of these three document roots or arrays:

- node.<number>[.<capability>].<property>
- relation.node.<number>
- probe.<number>.property

The corresponding value will depend on the nature of the path.
- for nodes : the value will be the required value of the property.
- for relations : the value will be the required target node of the relation.
- for probes : the value will be the required value of the property.

The following property names exist for all node paths outside of capability extensions.

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

The following structure describes the properties of a typical combination of hardware and software nodes

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

