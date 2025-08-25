package provider

import (
    "context"
    "fmt"
    "time"
    "github.com/hashicorp/terraform-plugin-framework/resource"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema"
    "github.com/hashicorp/terraform-plugin-framework/types"
    "github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
    _ resource.Resource              = &beamResource{}
    _ resource.ResourceWithConfigure = &beamResource{}
)

// NewBeamResource is a helper function to simplify the provider implementation.
func NewBeamResource() resource.Resource {
    return &beamResource{}
}

// beam Resource is the resource implementation.
type beamResource struct{
	client *Client
}

// beamResourceModel maps the resource schema data.
type beamResourceModel struct {
    ID          types.String     `tfsdk:"id"`
    Template    types.String     `tfsdk:"template"`
    Program	types.String     `tfsdk:"program"`
    Domain	types.String     `tfsdk:"domain"`
    Region	types.String     `tfsdk:"region"`
    Category	types.String     `tfsdk:"category"`
    Param	types.String     `tfsdk:"param"`
    State       types.String     `tfsdk:"state"`
    LastUpdated types.String     `tfsdk:"last_updated"`
}

func (r *beamResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_beam"
}

// Schema defines the schema for the resource.
func (r *beamResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        Attributes: map[string]schema.Attribute{
            "id": schema.StringAttribute{
                Computed: true,
            },
            "state": schema.StringAttribute{
                Computed: true,
            },
            "last_updated": schema.StringAttribute{
                Computed: true,
            },
            "template": &schema.StringAttribute{
                Computed: false,
		Required: true,
            },
            "program": &schema.StringAttribute{
                Computed: false,
		Required: true,
            },
            "domain": &schema.StringAttribute{
                Computed: false,
		Required: true,
            },
            "region": &schema.StringAttribute{
                Computed: false,
		Required: true,
            },
            "category": &schema.StringAttribute{
                Computed: false,
		Required: true,
            },
            "param": &schema.StringAttribute{
                Computed: false,
		Required: true,
            },
        },
    }
}

// -------------------------------------------------
// CREATE BEAM RESOURCE
// -------------------------------------------------
// Create the BEAM RESOURCE as described by the plan
// comprising template, program, domain, region, 
// category and parameter information.
// This is a multi phase operation requiring the
// following actions to be successfully performed:
//
// - CREATE BEAM MODEL from TEMPLATE to PROGRAM
//
// Each transition will be accompanied to ensure
// the correct completion of BEAM RESOURCE CREATE.
// -------------------------------------------------
func (r *beamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
    var err  error
    var plan beamResourceModel
    tflog.Info(ctx,"AMENESIK:BEAM ENTER:CREATE: Get Plan");
    diags := req.Plan.Get(ctx, &plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    // prepare the resource description parameters
    template := plan.Template.String()
    program  := plan.Program.String()
    domain   := plan.Domain.String()
    region   := plan.Region.String()
    category := plan.Category.String()

    var br *BeamResponse

    // CLONE a BEAM model with specific provisioning characteristics
    br, err = r.client.CloneBeamModel(ctx,template, program, domain, region, category )
    if err != nil {
	tflog.Info(ctx,"AMENESIK:BEAM ERROR: CLONE BEAM MODEL: "+err.Error());
        return
    }
    // prepare the final state description
    if br.status == "200" {
	    br.status = "locked"
    }
    plan.ID = types.StringValue(time.Now().Format(time.RFC3339))
    plan.State = types.StringValue(br.status)
    plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))
    diags = resp.State.Set(ctx, plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
    tflog.Info(ctx,"AMENESIK:BEAM LEAVE:CREATE: SUCCESS");
    return
}

// Read refreshes the Terraform state with the latest data.
func (r *beamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *beamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *beamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
    var err  error
    var state beamResourceModel
    tflog.Info(ctx,"AMENESIK:BEAM ENTER:DELETE: Get Plan");
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    // prepare the resource description parameters
    template := state.Template.String()
    program  := state.Program.String()

    var br *BeamResponse

    // DELETE the BEAM model now
    br, err = r.client.DeleteBeamModel(ctx, template, program )
    if err != nil {
	tflog.Info(ctx,"AMENESIK:BEAM ERROR: DELETE BEAM MODEL: "+err.Error());
        return
    }
    br.status = ""
    tflog.Info(ctx,"AMENESIK:BEAM LEAVE:DELETE: SUCCESS");
    return
}

// Configure adds the provider configured client to the resource.
func (r *beamResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
    // Add a nil check when handling ProviderData because Terraform
    // sets that data after it calls the ConfigureProvider RPC.
    if req.ProviderData == nil {
        return
    }

    client, ok := req.ProviderData.(*Client)

    if !ok {
        resp.Diagnostics.AddError(
            "Unexpected Data Source Configure Type",
            fmt.Sprintf("Expected *Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
        )

        return
    }

    r.client = client
}

