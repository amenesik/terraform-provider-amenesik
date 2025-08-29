# Terraform Provider Amenesik

This provider plugin for the Terraform platform allows deployment and life-cycle management of complex business management application workloads through the Multi-Cloud Amenesik Enterprise Cloud Service Federation platform.

## Resources
The Amenesik Terraform Provider plugin describes the following resources:

- App : These resource types are used to manage complete application deployment instances
- Beam : These resource types are used to manage Basmati Enhanced Application Model descriptions that are used for the description of the deployment details for the preceding App resource.

## App
This resource type will be used to manage the deployment instances of your complex, multi-cloud business application configurations.

A simple example of the use of this type of resource can be seen in the sample Terraform configuration document.

  
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

The provider section defines the configuration values required for the Amenesik provider:
- Host : the domain name of the Amenesik Enterprise Cloud platform, usually "phoenix.amenesik.com" or "mycompany.amenesik.com"
- Account : the provisioning account name on the corresponding Amenesik Enterprise Cloud platform
- Apikey : The API KEY associated with the provisioning account. This is a sensitive value and should be not be written in plain text in configuration documents.

The variable ace_api_key allows the sensitive string value of the amenesik provider API KEY to be defined through the Terraform variable management mechanisms, including environment variables, terraform command line switches and prompted user input values.

The resource section provides the values for the required parameters of a single amenesik provider APP resource:

- Template: The value of this property indicates the name of the BEAM resource describing the details of the application configuration.
  
- Program: The value of this property indicates the name of the application instance. It will be used in conjunction with the template name for the preparation of the instance specific derivation of the template document. It will also be used in conjunction with the value of the domain property in the composition of the fully quallified endpoint domain name.
  
- Domain: The value of this property provides the domain name used in conjunction with the preceding program property for the composition of the fully quallified endpoint domain name.
  
- Category: The value of this property indicates the name of the Amenesik Enterprise Cloud service provisioning type.  The property may be either a single quoted value, such as "amazonec2", or a quoted, comma separated, square braced list of alternative provisioning categories, such as "[amazonec2,googlecompute,windowsazure]". In the first case a single application state will result. In the second instance three alternative application states will be created, one for each of the specified provisioning types, but only the first will actually be started. The other two, initially idle states,  will provide alternative application fail-over states. They will be deployed and or released, as required, in response to the eventual reception, by the application controller, of "change" and "revert" action events. These events are issued as a result of critical failure detection by the life cycle management of the Amenesik Cloud Engine.
    
- Region: The value of this property indicates the name of the region and will be used in conjunction with the category property value for cloud provider region selection. As for the preceding "category" property, this property may also be either a single quoted value, such as "france", or a quoted, comma separated, square braced list of alternative provisioning regions, such as "[france,germany,italy]". In the first case a single application state will result. In the second instance, three alternative application states will be created, one for each of the specified provisioning regions, with identical subsequent behaviour as described for multiple provisioning categories. The category and region properties may both specify multiple values, in which case the corresponding ordered combinations will be used. For example, with category set to "[a,b]" and region set = "[c,d]" then two application states would be created, one for category a in region c, and one for category b in region d.
  
- Param: the value of this property allows optional application specific parameters to be passed to the application instance during its startup.

Management of the deployment of a suitably defined APP instance would be performed using the standard terraform command, as can be seen below:

    $ terraform plan
    ...
    $ terraform apply
    ...
    $ teraform show
    ...
    $ terraform destroy 
    ...

These commands must be launched from the folder containing the application configuration file, unless the terraform directory change command line switch is specified.

## Beam
This resource type will be used to manage the BEAM description documents of complex, multi-cloud business applications.

The BEAM document format, being one of the outcomes of the H2020 European project known as BASMATI, is an accronym for Basmati Enhanced Application Model.

### Introduction
BEAM is a conforming derivation of the standard TOSCA document format, with the addition of standardised TAG values for the specification of information relating to the Service Level Agreement terms that describe the conditions and guarantees required for the deployment and life-cycle management of the application service.

A complete BEAM document comprises a collection of Node Type definitions, which may result from import statements, and the application's Service Template.

The Service Template comprises the collection of BEAM Tag values and the application's Topology Template.

The Topology Template comprises the collection of Node Templates describing the nature, composition and requirements of the service nodes needed by the application.

The Topology Template also comprises an optional collection of Relationship Templates describing the heirarchy and interconnection of the Nodes.

Node Templates may be of the following types:
- Hardware nodes describing machines
- Software nodes describing software layers and configurations to be applied to Hardware nodes
- Service nodes performing service oriented operations on behalf of any of the other node types

The amenesik terraform provider's BEAM resource type subsequently allows the definition, creation, management and destruction of BEAM documents.

A example of a complex Topology Template, described by a single BEAM document, as managed by the Amenesik Enterprise Cloud, is shown below. The links between the nodes represent the information provided by the collection of Relationship Templates. 

<img width="1920" height="1080" alt="image" src="https://github.com/user-attachments/assets/774c4b9a-b5a5-4e85-8f24-6f5538fe6d6e" />

This complex, yet concrete example, describes the deployment, configuration and interconnection of:
- twelve virtual machine nodes in one geographical region
- seven virtual machine nodes in a secondary geographical region
- accessible by both regional entry points
- and balanced by a global traffic manager service instance

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

The above example shows two BEAM (TOSCA) nodes, the first a hardware node of type Compute, the second a software node of type Database. The second uses the hardware Compute node as its base.

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

From the above examples it should be noted that the data array of the BEAM resource describes the properties and their values of the BEAM document.

### Syntax
Conceptually, BEAM documents comprise ordered collections of TAGS, TYPES, IMPORTS, NODES, RELATIONS and PROBES (a specialisation of the node).

A Data Path must be defined with respect to one of these document roots or arrays:

- node . < identifier > [ . < capability > ] . < property >
- relation . node . < number > . [ hostname | contract ]
- probe . < identifier > . < property >
- type . < name > . < property >
- tag . < name >
- import

In addtition, the data path also allows replication of preceding nodes and probes using the following operations:

- copy . node
- copy . probe

In the general syntax above, the term 'identifier' may be either a number or the name of the node or probe. After the use of the 'copy' operation the term 'last' may be used to address the most recently created node.

The corresponding value will depend on the nature of the path.
- For nodes : The value will provide the value required to be set for the property of the node.
- For relations : The value will be the identification of required target node of the relation
- For probes : The value will provide the value required to be set for the property of the probe.

#### Nodes
The following property names exist for all node paths, outside of any capability extensions.

- Name : the name of the node
- Type : the usage type of the node as Compute or other Software layer definitions.
  Description : a human readable text string describing the purpose of the node.
- Base : the parent node of a node layering collection, where the hardware node provides a base for a subsequent chain of software node layers. 

The optional capability of a node path expression may be one of the following.

- Host : the collection of host properties of the Compute node type
- Os : the collection of operating system properties of the Compute node type
- Other capability values may be used for the node type specific capability parameters of software node configurations.

The host capability, of the Compute node type, defines the following properties

- Num_cpus : the number of cores or virtual cpus for the Compute node
- Mem_size : the size of the memory suffixed by M or G
- Disk_size : the size of the disk, suffixed by M, G or T
- Volume : the name of an attached volume
- Entry : the entry point description
- Hostname : the fully qualified host and domain name
- Provider : the provisioning category, when specific to a node
- Region : the provisioning region when specific to a node
- Vlan : the vlan to which the node should be attached
- Tcp_port : a TCP port to be opened in the firewall or security group
- Udp_port : a UDP port to be opened in the firewall or security group
- Tcp_range : a dash or comma separated range of TCP ports to be opened in the firewall or security group
- Udp_range : a dash or comma separated range of UDP ports to be opened in the firewall or security group
- Protocol : the network protocol
- Cluster : the name of the cluster for a container compute node
- Namespace : the name of the namespace for a container compute node
 
The OS capability, of the Compute node type, defines the following properties

- Architecture : describes the hardware architecture such as x86_64
- Type : indicates the operating system family as linux or windows,
- Distribution : the name of the distribution as WINDOWS, UBUNTU or RHEL
- Version : the version of the specified distribution such as 20.04 or 9

The properties of the software node types are "type specific" and require consultation of the product capabilites on the amenesik web site.

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

In real world situations a large number of nodes would be defined each with their own specific collections of capabilities and their associated properties.

The complex example shown above, requires 22 hardware (Compute) nodes, 19 software nodes of four different classes (LDAP, TOMCAT, APACHE, HAPROXY) and roughly 50 relations. In addition, a variety of probes would be required for both hardware and software operation and failover monitoring.  The resulting BEAM document would naturally be correspondingly complex.

### Probes
The following properties are defined for the monitoring probes.

- Name : This property provides the logical name of the probe and should match the one of the Probe Tag values or a Node Probe name.
- Metric : The definition of the type of information, its collection frequency and its means of collection. These may be defined per account by the Amenesik Enterprise Cloud.
- Condition : the nature of the comparaison with the threshold value (eq, gr, ls, ge, le, ne)
- Threshold : the threshold value which when reached requires remediative "penalty" action to be engaged
- Type : the nature of the remediation action invocation (one of OCCISCRIPT, BASH, PYTHON)
- Nature : the purpose or nature of the action (one of penalty, reward, both). When "reward" or "both" then the actio will be engaged before threshold is reached.
- Behaviour : the name of the OCCI, BASH or PYTHON script describing the subsequent action

Probes are attached to hardware nodes either implicitly through the collection of probes defined in the Tag section, or explicitely by the collection of probes defined as requirements of a node, or a combination of both.

### Relations
A relation is required to be defined when a secondary (target) node construction (hardware and software elements) requires autotamtion of its connection to a primary (source) node construction (hardware and software elements) during the deployment of infrastructural APP resources.

The target node is said to receive connection information from the source node during the configuration (or construction) phase of its life cycle.

The following Data definition of a BEAM resource shows the properties required to describe and establish such a connection, always defined between the hardware nodes.

    resource "amenesik_beam" "small" { 
      ...
    	data     = [
        { path = "node.1.type", value = "Compute" },
        { path = "node.2.type", value = "Database" },
        { path = "node.3.type", value = "Compute" },
        { path = "node.4.type", value = "WebServer" },
    		{
    		path = "relation . node . 1 . hostname"
    		value = "node.3"
    		}      
      ]
    }

In this example the Compute node of the WebServer node of the configuration will receive the hostname of the Compute node of the Database node of the configuration.

In most cases, the subject of the relationship will be the hostname of the source. 

In certain cases, especially concerning automation of load balancing scenarios, it is necessary that the subject of the relationship be the contract identifier of the source, facilitating replication of the source, by the target, in accordance with encountered real time load.

### Tags 
Tags are used to define service deployment conditions and other metadata and are defined for a BEAM resource in terms of theier name and value.

The following tags are currently availabe for use in BEAM documents, and other than the Probe tag, should all be self-explanatory:

- Title : defines the title of the BEAM resource
- SubTitle : defines the sub title of the BEAM resource
- Author : defines the author of the BEAM resource
- Version : defines the version of the BEAM resource
- Date : defines the date of creation or modification of the document
- Account : defines the default, or global, provisioning account
- Provider : defines the default, or global, provisioning category
- Zone : defines the default, or global, provisioning zone
- Domain : defines the default, or global, domain name
- Probe : defines the name of an Amenesik Enterprise Cloud Probe definition

The following configuration document snippet shows an example of tag definitions.

    resource "amenesik_beam" "small" { 
      ...
    	data     = [
        ...
        { path = tag.Title", value = "My Document Title" },
        { path = "tag.Author", value = "The document Author name" },
        { path = "tag.Probe", value = "memory-free" },
        ...
      ]
    }

### Types
The Node types that are used to define the behaviour of software layer nodes may be defined locally for a  BEAM resource in terms of the following properties:

- Name : the terminal name portion of the node type ( will be prefixed by tosca-nodes- by the Amenesik Enterprise Cloud BEAM resource processor )
- Create : the public, web fetchable action script to be fetched and launched when a node is created
- Start : the public, web fetchable action script to be fetched and launched when a node is started
- Stop : the public, web fetchable action script, to be fetched and launched when a node is stopped
- Save : the public, web fetchable action script, to be fetched and launched when a node is saved
- Delete : the public, web fetchable action script, to be fetched and launched when a node is deleted
- Tcp_Port : allows the collection of TCP network ports of the node type to be described individually.
- Tcp_Range : allows the collection of TCP network ports of the node type to be described as a range.
- Udp_Port : allows the collection of UDP network ports of the node type to be described individually.
- Udp_Range : allows the collection of UDP network ports of the node type to be described as a range.

### Imports
Node types may be imported instead of being defined locally in BEAM documents. This encourages reusability.

The following configuration document snippet shows an example of node import definitions.

    resource "amenesik_beam" "small" { 
      ...
    	data     = [
        ...
        { path = import", value = "Database" },
        { path = "import", value = "WebServer" },
        ...
      ]
    }

The use of formal import statments is facultative since node types encountered during node type statements will be added automatically to the document as import statements unless the presence of a local node type definitions is detected.

## Complete Example
The following configuration document shows the creation of a complex multi layer load balenced web application scenario with 6 application servers and two database servers and three load balancers.

    resource "amenesik_beam" "mybeam" {
            template = "template"
            program  = "mybeam-template"
            domain   = "myhost.com"
            region   = "any"
            category = "any"
            param    = "none"
            data     = [
                    # beam document tags
                    { path = "tag.Title", value="My New Beam Document" },
                    { path = "tag.SubTitle", value="Generated by Terraform" },
                    { path = "tag.Author", value="Iain James Marshall" },
                    { path = "tag.Email", value="ijm@amenesik.com" ],
                    { path = "tag.Zone", value="any" },
                    { path = "tag.Provider", value="any" },
                    { path = "tag.Domain", value="myhost.com" },

                    # beam document default probe tags (probes to be added to all hardware nodes)
                    { path = "tag.Probe", value="memory-free" },
                    { path = "tag.Probe", value="disk-free" },
                    { path = "tag.Probe", value="load-average" },
                    { path = "tag.Probe", value="net-rx" },
                    { path = "tag.Probe", value="net-tx" },

                    # memory-free probe
                    { path = "probe.1.name", value = "memory-free" },
                    { path = "probe.1.metric", value = "memory:free" },
                    { path = "probe.1.condition", value = "gr" },
                    { path = "probe.1.threshold", value = "0" },
                    { path = "probe.1.type", value = "OCCISCRIPT" },
                    { path = "probe.1.value", value = "both" },
                    { path = "probe.1.behaviour", value = "activity" },
    
                    # disk-free probe
                    { path = "copy.probe", "value" = "1" },
                    { path = "probe.last.name", value = "disk-free" },
                    { path = "probe.last.metric", value = "disk:free" },
                    { path = "probe.last.threshold", value = "20" },
    
                    # load-average probe
                    { path = "copy.probe", "value" = "1" },
                    { path = "probe.last.name", value = "load-average" },
                    { path = "probe.last.metric", value = "load:average:1" },
                    { path = "probe.1.condition", value = "ge" },
    
                    # net-rx probe
                    { path = "copy.probe", "value" = "1" },
                    { path = "probe.last.name", value = "net-rx" },
                    { path = "probe.last.metric", value = "net:eth0:rx:rate" },
                    { path = "probe.1.condition", value = "ge" },
    
                    # net-tx probe
                    { path = "copy.probe", "value" = "1" },
                    { path = "probe.last.name", value = "net-tx" },
                    { path = "probe.last.metric", value = "net:eth0:tx:rate" },
                    { path = "probe.1.condition", value = "ge" },

                    # template specific type MyDatabase
                    { path = "type.1.name", value = "MyDatabase" },
                    { path = "type.1.create", value = "wget https://www.myhost.com/install-database.sh; bash ./install-database.sh" },
                    { path = "type.1.start", value = "bash ./install-database.sh" },
                    { path = "type.1.delete", value = "wget https://www.myhost.com/delete-database.sh; bash ./delete-database.sh" },
                    { path = "type.1.tcp_port", value = "3306" },

                    # template specific type MyWebServer
                    { path = "type.2.name", value = "MyWebServer" },
                    { path = "type.2.create", value = "wget https://www.myhost.com/install-web-server.sh; bash ./install-web-server.sh" },
                    { path = "type.2.start", value = "bash ./install-web-server.sh" },
                    { path = "type.2.delete", value = "wget https://www.myhost.com/delete-web-server.sh; bash ./delete-web-server.sh" },
                    { path = "type.2.tcp_port", value = "443"  },

                    # template specific type MyLoadBalancer
                    { path = "type.2.name", value = "MyLoadBalancer" },
                    { path = "type.2.create", value = "wget https://www.myhost.com/install-load-balancer.sh; bash ./install-load-balancer.sh" },
                    { path = "type.2.start", value = "bash ./install-load-balancer.sh" },
                    { path = "type.2.delete", value = "wget https://www.myhost.com/delete-load-balancer.sh; bash ./delete-load-balancer.sh" },
                    { path = "type.2.tcp_port", value = "443"  },
                    { path = "type.2.tcp_port", value = "3306" },
                    
                    # beam document node: database hardware
                    { path = "node.1.name", value = "dbhwa" },
                    { path = "node.1.type", value = "Compute" },
                    { path = "node.1.host.num_cpus", value = "8" },
                    { path = "node.1.host.mem_size", value = "32G" },
                    { path = "node.1.host.disk_size", value = "110G" },
                    { path = "node.1.os.distribution", value = "UBUNTU" },
                    { path = "node.1.os.version", value = "20.04" },
    
                    # beam document node: database software
                    { path = "node.2.name", value = "dbswa" },
                    { path = "node.2.base", value = "dbhwa" },
                    { path = "node.2.type", value = "MyDatabase" },
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
                    { path = "node.last.type", value = "MyWebServer" },
    
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
                    { path = "copy.node", "value" = "dbhwa" },
                    { path = "node.last.name", value = "lbhwa" },
    
                    # beam document node: load balancer software
                    { path = "copy.node", "value" = "dbswa" },
                    { path = "node.last.name", value = "lbswa" },
                    { path = "node.last.base", value = "lbhwa" },
                    { path = "node.last.type", value = "MyLoadBalancer" },
    
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
                    { path = "node.last.host.hostname", value = "www.myhost.com" },
                    { path = "node.last.host.entry", value = "443/https" },
    
                    # beam document node: load balancer software copy
                    { path = "copy.node", "value" = "lbswa" },
                    { path = "node.last.name", value = "lbswc" },
                    { path = "node.last.base", value = "lbhwc" },
    
                    # beam document relations

                    # connection of web servers to database a
                    { path = "relation.node.dbhwa.hostname", value="node.wshwa" },
                    { path = "relation.node.dbhwa.hostname", value="node.wshwb" },
                    { path = "relation.node.dbhwa.hostname", value="node.wshwc" },
                    { path = "relation.node.dbhwa.hostname", value="node.wshwd" },
                    { path = "relation.node.dbhwa.hostname", value="node.wshwe" },
                    { path = "relation.node.dbhwa.hostname", value="node.wshwf" },

                    #connection of web servers to database b
                    { path = "relation.node.dbhwb.hostname", value="node.wshwa" },
                    { path = "relation.node.dbhwb.hostname", value="node.wshwb" },
                    { path = "relation.node.dbhwb.hostname", value="node.wshwc" },
                    { path = "relation.node.dbhwb.hostname", value="node.wshwd" },
                    { path = "relation.node.dbhwb.hostname", value="node.wshwe" },
                    { path = "relation.node.dbhwb.hostname", value="node.wshwf" },

                    # connection of web servers to load balancer a
                    { path = "relation.node.wshwa.hostname", value="node.lbhwa" },
                    { path = "relation.node.wshwb.hostname", value="node.lbhwa" },
                    { path = "relation.node.wshwc.hostname", value="node.lbhwa" },
                    { path = "relation.node.wshwd.hostname", value="node.lbhwa" },
                    { path = "relation.node.wshwe.hostname", value="node.lbhwa" },
                    { path = "relation.node.wshwf.hostname", value="node.lbhwa" },

                    # connection of web servers to load balancer b
                    { path = "relation.node.wshwa.hostname", value="node.lbhwb" },
                    { path = "relation.node.wshwb.hostname", value="node.lbhwb" },
                    { path = "relation.node.wshwc.hostname", value="node.lbhwb" },
                    { path = "relation.node.wshwd.hostname", value="node.lbhwb" },
                    { path = "relation.node.wshwe.hostname", value="node.lbhwb" },
                    { path = "relation.node.wshwf.hostname", value="node.lbhwb" },

                    # connection of load balancers a and b to c
                    { path = "relation.node.lbhwa.hostname", value="node.lbhwc" },
                    { path = "relation.node.lbhwb.hostname", value="node.lbhwc" },
            ]
    }

The processing of this BEAM resource, using Terraform Apply, would result in the following BEAM Topology being created in the Amenesik Enterprise Cloud.

<img width="1920" height="1080" alt="image" src="https://github.com/user-attachments/assets/371fddb9-2b96-420c-bd9b-4d0909f2466e" />

Naturally the use of Terraform Destroy would delete the BEAM resource from the Amenesik Enterprise Cloud.



