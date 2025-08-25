
package provider

import (
  "context"
  "os"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/path"
  "github.com/hashicorp/terraform-plugin-framework/provider"
  "github.com/hashicorp/terraform-plugin-framework/provider/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
    _ provider.Provider = &amenesikProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
    return func() provider.Provider {
        return &amenesikProvider{
            version: version,
        }
    }
}

// amenesikProvider is the provider implementation.
type amenesikProvider struct {
    // version is set to the provider version on release, "dev" when the
    // provider is built and ran locally, and "test" when running acceptance
    // testing.
    version string
}

// amenesikProviderModel maps provider schema data to a Go type.
type amenesikProviderModel struct {
    Host    types.String `tfsdk:"host"`
    Account types.String `tfsdk:"account"`
    ApiKey  types.String `tfsdk:"apikey"`
}

// Metadata returns the provider type name.
func (p *amenesikProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
    resp.TypeName = "amenesik"
    resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *amenesikProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
    resp.Schema = schema.Schema{
        Attributes: map[string]schema.Attribute{
            "host": schema.StringAttribute{
                Optional: true,
            },
            "account": schema.StringAttribute{
                Optional: true,
            },
            "apikey": schema.StringAttribute{
                Optional: true,
            },
        },
    }
}

// Configure prepares a Amenesik API client for data sources and resources.
func (p *amenesikProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
    tflog.Info(ctx,"Configuring Amenesik client")
    // Retrieve provider data from configuration
    var config amenesikProviderModel
    diags := req.Config.Get(ctx, &config)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    // If practitioner provided a configuration value for any of the
    // attributes, it must be a known value.

    if config.Host.IsUnknown() {
        resp.Diagnostics.AddAttributeError(
            path.Root("host"),
            "Unknown Amenesik API Host",
            "The provider cannot create the Amenesik API client as there is an unknown configuration value for the Amenesik API host. "+
                "Either target apply the source of the value first, set the value statically in the configuration, or use the ACE_HOST environment variable.",
        )
    }

    if config.Account.IsUnknown() {
        resp.Diagnostics.AddAttributeError(
            path.Root("account"),
            "Unknown Amenesik API Account",
            "The provider cannot create the Amenesik API client as there is an unknown configuration value for the Amenesik API account. "+
                "Either target apply the source of the value first, set the value statically in the configuration, or use the ACE_ACCOUNT environment variable.",
        )
    }

    if config.ApiKey.IsUnknown() {
        resp.Diagnostics.AddAttributeError(
            path.Root("apikey"),
            "Unknown Amenesik API KEY",
            "The provider cannot create the Amenesik API client as there is an unknown configuration value for the Amenesik API KEY. "+
                "Either target apply the source of the value first, set the value statically in the configuration, or use the ACE_APIKEY environment variable.",
        )
    }

    if resp.Diagnostics.HasError() {
        return
    }

    // Default values to environment variables, but override
    // with Terraform configuration value if set.

    host := os.Getenv("ACE_HOST")

    if !config.Host.IsNull() {
        host = config.Host.ValueString()
    }

    account := os.Getenv("ACE_ACCOUNT")

    if !config.Account.IsNull() {
        account = config.Account.ValueString()
    }

    apikey := os.Getenv("ACE_APIKEY")

    if !config.ApiKey.IsNull() {
        apikey = config.ApiKey.ValueString()
    }

    // If any of the expected configurations are missing, return
    // errors with provider-specific guidance.

    if host == "" {
        resp.Diagnostics.AddAttributeError(
            path.Root("host"),
            "Missing Amenesik API Host",
            "The provider cannot create the Amenesik API client as there is a missing or empty value for the Amenesik API host. "+
                "Set the host value in the configuration or use the ACE_HOST environment variable. "+
                "If either is already set, ensure the value is not empty.",
        )
    }

    if account == "" {
        resp.Diagnostics.AddAttributeError(
            path.Root("account"),
            "Missing Amenesik API Account",
            "The provider cannot create the Amenesik API client as there is a missing or empty value for the Amenesik API account. "+
                "Set the account value in the configuration or use the ACE_ACCOUNT environment variable. "+
                "If either is already set, ensure the value is not empty.",
        )
    }

    if apikey == "" {
        resp.Diagnostics.AddAttributeError(
            path.Root("apikey"),
            "Missing Amenesik API KEY",
            "The provider cannot create the Amenesik API client as there is a missing or empty value for the Amenesik API KEY. "+
                "Set the API KEY in the configuration or better still use the ACE_APIKEY environment variable. "+
                "If either is already set, ensure the value is not empty.",
        )
    }

    if resp.Diagnostics.HasError() {
        return
    }

    ctx = tflog.SetField(ctx,"amenesik_host", host)
    ctx = tflog.SetField(ctx,"amenesik_account", account)
    ctx = tflog.SetField(ctx,"amenesik_apikey", apikey)

    tflog.Debug(ctx,"Creating Amenesik client")

    // Create a new Amenesik client using the configuration values
    client, err := NewClient(ctx,host,account,apikey)
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to Create Amenesik API Client",
            "An unexpected error occurred when creating the Amenesik API client. "+
                "If the error is not clear, please contact the provider developers.\n\n"+
                "Amenesik Client Error: "+err.Error(),
        )
        return
    }

    // Make the Amenesik client available during DataSource and Resource
    // type Configure methods.
    resp.DataSourceData = client
    resp.ResourceData = client
    tflog.Info(ctx,"Configured Amenesik client", map[string]any{"success":true})
}

// DataSources defines the data sources implemented in the provider.
func (p *amenesikProvider) DataSources(_ context.Context) []func() datasource.DataSource {
  return nil
}

// Resources defines the resources implemented in the provider.
func (p *amenesikProvider) Resources(_ context.Context) []func() resource.Resource {
    return []func() resource.Resource{
        NewAppResource,
	NewBeamResource,
    }
}


