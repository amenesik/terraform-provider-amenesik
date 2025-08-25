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

The variable ace_api_key allows the sensitive string value of the amenesik provider API KEY to be defined throiugh the Terraform variable management mechanisms, including environment vcariables, terraform command line switches and prompted user input values.

The resource section provides the values for the required parameters of an amenesik provider resource:
- template: this property indicates the name of the BEAM document describing the details of the application configuration.
- program: this property indicates the name of the application instance and will be used in conjunction with the template name for the preparation of the instance specific derivation of the template document, and in conjunction with the value of the domain property in the composition of the fully quallified endpoint domain name.
- domain: this property provides the domain name used in conjunction with the preceding program property for the composition of the fully quallified endpoint domain name.
- category: this property indicates the name of the Amenesik Enterprise Cloud service provisioning type.  The property may be either a single quoted value, such as "amazonec2", or a quoted, comma separated, square braced list of alternative provisioning categories, such as "[amazonec2,googlecompute,windowsazure]". In the first case a single application state will result. In the second instance three alternative application states will be created, one for each of the specified provisioning types, but only the first will actually be started. The other two, initially idle states,  will provide alternative application fail-over states, that will be deployed and or released, as required, in response to the eventual reception, by the application controller, of "change" and "revert" action events, issued as a result of critical failure detection by the life cycle management of the Amenesik Cloud Engine.  
- region: this property indicates the name of the region and will be used in conjunction with the category property value for cloud provider region selection. As for the preceding "category" property, this property may also be either a single quoted value, such as "france", or a quoted, comma separated, square braced list of alternative provisioning regions, such as "[france,germany,italy]". In the first case a single application state will result. In the second instance, three alternative application states will be created, one for each of the specified provisioning regions, with identical subsequent behaviour as described for multiple provisioning categories. The category and region properties may both specify multiple values, in which case the corresponding ordered combinations will be used. For example, with category set to "[a,b]" and region set = "[c,d]" then two application states would be created, one for category a in region c, and one for category b in region d.
- param: this property allows optional, supplimentary, application specific, parameters to be passed to the application instance during its startup.

Management of the deployment of this APP instance would be performed using the standard terraform command:

    $ terraform plan
    ...
    $ terraform apply
    ...
    $ teraform show
    ...
    $ terraform destroy 
    ...

from the folder containing the application configuration file.

## Beam
This resource type will be used to manage the BEAM description documents of your complex, multi-cloud, business application.

The BEAM document format, being one of the outcomes of the H2020 European project known as BASMATI, is an accronym for Basmati Enhanced Application Model.

BEAM is a derivation of the standard TOSCA document format with the addition of standardised TAG values for the specification of information relating to the Service Level Agreement terms that describe the conditions and guarantees required for the deployment and life-cycle management of the application service.

A standard BEAM document comprises a collection of Node Type definitions and the application's Service Template.

The Service Template comprises the collection of BEAM Tag values and the application Topology Template.

The Topology Template comprises the collection of Node Templates describing the nature, composition and requirements of the service nodes required by the application.

Node Templates may be of the following types:
- Hardware nodes describing machines
- Software nodes describing software layers and configurations to be applied to Hardware nodes.
- Service nodes performing service oriented operations on behalf of any of the other node types.

The amenesik terraform provider BEAM resource subsequently type allows the definition, creation, management and destruction of BEAM documents.








  

