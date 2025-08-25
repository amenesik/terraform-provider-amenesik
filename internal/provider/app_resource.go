package provider

import (
    "context"
    "fmt"
    "time"
    "errors"
    "github.com/hashicorp/terraform-plugin-framework/resource"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema"
    "github.com/hashicorp/terraform-plugin-framework/types"
    "github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
    _ resource.Resource              = &appResource{}
    _ resource.ResourceWithConfigure = &appResource{}
)

// NewAppResource is a helper function to simplify the provider implementation.
func NewAppResource() resource.Resource {
    return &appResource{}
}

// app Resource is the resource implementation.
type appResource struct{
	client *Client
}

// appResourceModel maps the resource schema data.
type appResourceModel struct {
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

func (r *appResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_app"
}

// Schema defines the schema for the resource.
func (r *appResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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

// ----------------------------------------------
// WAIT FOR STATUS
// ----------------------------------------------
// Accompany the operation that is underway until
// the expected current status transitions to the 
// required status. Failure if other status than
// that expected occurs.
// ----------------------------------------------
func WaitForStatus(r *appResource, ctx context.Context, br *BeamResponse, t string, p string, d string, waiting string, waited string ) (*BeamResponse, error) {
    if ( br.status == "200" ) {
        br.status = waiting
    }
    for br.status == waiting {
        var err  error
        var rr *BeamResponse
	time.Sleep( 3 * time.Second )
	// inspect the BEAM instance status
	rr, err = r.client.StatusBeamInstance(ctx, t, p, d )
	// signal but tolerate errors
	if err != nil {
	    tflog.Info(ctx,"AMENESIK:APP ERROR: STATUS BEAM INSTANCE: "+err.Error());
        } else {
	    br.status = rr.status
        }
    }
    // ensure required state reached
    if br.status == waited {
        return br, nil
    } else {
	return nil, errors.New("500")
    }
}

// ------------------------------------------------
// CREATE APP RESOURCE
// ------------------------------------------------
// Create the APP RESOURCE as described by the plan
// comprising template, program, domain, region, 
// category and parameter information.
// This is a multi phase operation requiring the
// following actions to be successfully performed:
//
// - CLONE  BEAM MODEL from TEMPLATE to PROGRAM
// - CREATE BEAM INSTANCE at REGION and CATEGORY
// - START  BEAM INSTANCE in DOMAIN with PARAM
// - LOCK   BEAM INSTANCE against unwanted actions
//
// Each transition will be accompanied to ensure
// the correct completion of APP RESOURCE CREATE.
// ------------------------------------------------
func (r *appResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
    var err  error
    var plan appResourceModel
    tflog.Info(ctx,"AMENESIK:APP ENTER:CREATE: Get Plan");
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
    param    := plan.Param.String()

    var br *BeamResponse

    // CLONE a BEAM model with specific provisioning characteristics
    br, err = r.client.CloneBeamModel(ctx,template, program, domain, region, category )
    if err != nil {
	tflog.Info(ctx,"AMENESIK:APP ERROR: CLONE BEAM MODEL: "+err.Error());
        return
    }

    // CREATE the BEAM instance for the domain and application specific parameters
    br, err = r.client.CreateBeamInstance(ctx,template,program,domain,param )
    if err != nil {
	tflog.Info(ctx,"AMENESIK:APP ERROR: CREATE BEAM INSTANCE: "+err.Error());
        return
    }
    // accompany the instance creation operation from creating to created or error
    br, err = WaitForStatus(r,ctx,br,template, program, domain, "creating", "created")
    if  err != nil {
	    return
    }
    // START the BEAM instance now
    br, err = r.client.StartBeamInstance(ctx, template, program )
    if err != nil {
        tflog.Info(ctx,"AMENESIK:APP ERROR: START BEAM INSTANCE: "+err.Error());
        return
    }
    // accompany the instance start operation from creating to starting to started
    br, err = WaitForStatus(r,ctx,br,template, program, domain, "starting", "started")
    if  err != nil {
	    return
    }
    // LOCK the BEAM instance to protect against undesired state change
    br, err = r.client.LockBeamInstance(ctx, template, program )
    if err != nil {
	tflog.Info(ctx,"AMENESIK:APP ERROR: LOCK BEAM INSTANCE: "+err.Error());
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
    tflog.Info(ctx,"AMENESIK:APP LEAVE:CREATE: SUCCESS");
    return
}

// Read refreshes the Terraform state with the latest data.
func (r *appResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *appResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *appResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
    var err  error
    var state appResourceModel
    tflog.Info(ctx,"AMENESIK:APP ENTER:DELETE: Get Plan");
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    // prepare the resource description parameters
    template := state.Template.String()
    program  := state.Program.String()
    domain   := state.Domain.String()

    var br *BeamResponse

    // UNLOCK the BEAM instance allowing sequence of required state change
    br, err = r.client.UnLockBeamInstance(ctx, template, program )
    if err != nil {
	tflog.Info(ctx,"AMENESIK:APP ERROR: UNLOCK BEAM INSTANCE: "+err.Error());
        return
    }
    if br.status == "200" {
	    br.status = "started"
    }

    // STOP the BEAM instance now
    br, err = r.client.StopBeamInstance(ctx, template, program )
    if err != nil {
	tflog.Info(ctx,"AMENESIK:APP ERROR: STOP BEAM INSTANCE: "+err.Error());
        return
    }
    // accompany the instance stop operation from stopping to idle
    br, err = WaitForStatus(r,ctx,br,template, program, domain, "stopping", "created")
    if  err != nil {
	    return
    }
    // DROP the BEAM instance now
    br, err = r.client.DropBeamInstance(ctx, template, program )
    if err != nil {
	tflog.Info(ctx,"AMENESIK:APP ERROR: DELETE BEAM INSTANCE: "+err.Error());
        return
    }
    // accompany the instance stop operation from stopping to idle
    br, err = WaitForStatus(r,ctx,br,template, program, domain, "deleting", "none")
    if  err != nil {
	    return
    }
    // DELETE the BEAM model now
    br, err = r.client.DeleteBeamModel(ctx, template, program )
    if err != nil {
	tflog.Info(ctx,"AMENESIK:APP ERROR: DELETE BEAM MODEL: "+err.Error());
        return
    }
    br.status = ""
    tflog.Info(ctx,"AMENESIK:APP LEAVE:DELETE: SUCCESS");
    return
}

// Configure adds the provider configured client to the resource.
func (r *appResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

