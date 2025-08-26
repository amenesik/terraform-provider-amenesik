// -------------------------------------------
// AMENESIK CLOUD ENGINE (ACE)
// BASMATI ENHANCED APPLICATION MODEL (BEAM)
// APPLICATION PROVISIONING PACKAGE   (APP)
// -------------------------------------------
// This TERRAFORM RESOURCE allows provisioning
// of business application workloads on any of
// the major public or private cloud providers
// in any providers data center or region
// -------------------------------------------
// Applications are described by their BEAM.
// An extension to the TOSCA document format
// describing instance configuration, life
// cycle management, fault tolerance and cost.
// -------------------------------------------

package provider

import(
  "context"
  "encoding/json"
  "fmt"
  "strings"
  "net/http"
  "bytes"
  "io/ioutil"
  "github.com/hashicorp/terraform-plugin-log/tflog"
)

// ---------------------
// The ACE / BEAM CLIENT 
// ---------------------
type Client struct {
    httpClient *http.Client
    baseURL string
    account string
    apikey  string
    token   string
}

// -------------------------
// The ACE / BEAM AUTH TOKEN
// -------------------------
type BeamToken struct {
    status  string
    auth    string
    account string
    user    string
    role    string
    expires string
}

// --------------------------
// The ACE / BEAM / APP STATE
// --------------------------
type BeamInstanceState struct {
    name   string
    status string
}

// ------------------------------
// The standard BEAM API RESPONSE
// ------------------------------
type BeamResponse struct {
    status string
    result BeamInstanceState
}

// --------------------------------------
// remove the quotes from around a string
// --------------------------------------
func UnQuote( v string ) string {
	if v[0:1] != "\"" {
	    return v
	}
	nn := v[1:]
	return nn[:len(nn)-1]
}

// parse a simple JSON formated string to find a named value
func BeamJsonParser(js string, jn string) string {
    one := js[1:]
    two := one[:len(one)-1]
    parts := strings.Split(two,",")
    items := len(parts)
    for items > 0 {
	items--;
	three := strings.Split(parts[items],":")
	nnn := three[0]
	nn := nnn[1:]
	name := nn[:len(nn)-1]
	if name == jn {
	    nnn := three[1]
	    nn := nnn[1:]
	    return nn[:len(nn)-1]
	}
    }
    return "none"
}

// ---------------------------------
// Creation of a new ACE/BEAM CLIENT
// ---------------------------------
func NewClient(ctx context.Context,baseURL string, account string, apikey string) (*Client, error) {
    tflog.Info(ctx,"AMENESIK:ACE: NEW CLIENT: "+baseURL)
    reqBody := map[string]string{"action": "login", "user": account, "secret": apikey }
    body, _ := json.Marshal(reqBody)
    url := "https://"+baseURL+"/aec/api.php"
    req, _ := http.NewRequest("POST", url, bytes.NewReader(body))
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    bodyBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    auth := BeamJsonParser( string(bodyBytes), "auth" )

    tflog.Info(ctx,"AMENESIK:ACE: NEW CLIENT: SUCCESS: AUTH: "+baseURL)

    return &Client{
        httpClient: &http.Client{},
        baseURL:    url,
	account:    account,
	apikey:     apikey,
	token:      auth,
    }, nil
}

// ----------------------------------------------------------------------
// CLONE BEAM MODEL ( template, program, domain region, category )
// ----------------------------------------------------------------------
// Clones a new BEMA template from the specified template using program
// to replace the "template", and setting the instance hostname as the
// concatenation of the values of the program and domain parameters. The
// Cloud Provider will be set to the value of "category" in the "region".
// ----------------------------------------------------------------------
func (c *Client) CloneBeamModel(ctx context.Context,template string, program string, domain string, region string, category string) (*BeamResponse, error) {
    tflog.Info(ctx,"AMENESIK:ACE: CLONE BEAM MODEL: "+template+"-"+program)
    reqBody := map[string]string{"auth":c.token,"action": "clone", "subject": "beam", "account": c.account, "template": UnQuote(template), "program": UnQuote(program), "domain": UnQuote(domain), "region": UnQuote(region), "provider": UnQuote(category) }
    body, _ := json.Marshal(reqBody)

    req, _ := http.NewRequest("POST", c.baseURL, bytes.NewReader(body))
    req.Header.Set("Authorization", "Bearer "+c.token)
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return nil, fmt.Errorf("failed to create model: %s", resp.Status)
    }

    var bi BeamResponse
    bi.status = "cloned"
    bi.result.name = program
    bi.result.status = "200"
    tflog.Info(ctx,"AMENESIK:ACE: CLONE BEAM MODEL: "+template+"-"+program+": SUCCESS")
    return &bi, nil
}

// ----------------------------------------------------------------------
// CHANGE BEAM MODEL ( template, program, domain region, category )
// ----------------------------------------------------------------------
// Clones a new BEAM template from the specified template using program
// to replace the "template", and setting the instance hostname as the
// concatenation of the values of the program and domain parameters. The
// Cloud Provider will be set to the value of "category" in the "region".
// ----------------------------------------------------------------------
func (c *Client) ChangeBeamModel(ctx context.Context,template string, program string, domain string, region string, category string, data string ) (*BeamResponse, error) {
    tflog.Info(ctx,"AMENESIK:ACE: CLONE BEAM MODEL: "+template+"-"+program)
    reqBody := map[string]string{"auth":c.token,"action": "change", "subject": "beam", "account": c.account, "template": UnQuote(template), "program": UnQuote(program), "domain": UnQuote(domain), "region": UnQuote(region), "provider": UnQuote(category), "data":UnQuote(data) }
    body, _ := json.Marshal(reqBody)

    req, _ := http.NewRequest("POST", c.baseURL, bytes.NewReader(body))
    req.Header.Set("Authorization", "Bearer "+c.token)
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return nil, fmt.Errorf("failed to create model: %s", resp.Status)
    }

    var bi BeamResponse
    bi.status = "cloned"
    bi.result.name = program
    bi.result.status = "200"
    tflog.Info(ctx,"AMENESIK:ACE: CLONE BEAM MODEL: "+template+"-"+program+": SUCCESS")
    return &bi, nil
}

// ----------------------------------------------------------------------
// CREATE BEAM INSTANCE ( template, program, domain, param )
// ----------------------------------------------------------------------
// Creates a BEAM Application Controller instance fromi the cloned BEAM
// as described by the template and program parameters, using the domain
// and param information as required by the actual BEAM model.
// ----------------------------------------------------------------------
func (c *Client) CreateBeamInstance(ctx context.Context,template string, program string, domain string, param string) (*BeamResponse, error) {
    tflog.Info(ctx,"AMENESIK:ACE: CREATE BEAM INSTANCE: "+template+"-"+program)
    reqBody := map[string]string{"auth":c.token,"action": "create", "subject": "beam", "account": c.account, "template": UnQuote(template), "program": UnQuote(program), "domain": UnQuote(domain), "param": UnQuote(param) }
    body, _ := json.Marshal(reqBody)

    req, _ := http.NewRequest("POST", c.baseURL, bytes.NewReader(body))
    req.Header.Set("Authorization", "Bearer "+c.token)
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return nil, fmt.Errorf("failed to create instance: %s", resp.Status)
    }

    bodyBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    tflog.Info(ctx,"AMENESIK:ACE: STATUS: "+string(bodyBytes));

    var bi BeamResponse
    bi.status = BeamJsonParser(string(bodyBytes),"status")
    bi.result.name = BeamJsonParser(string(bodyBytes),"id")
    bi.result.status = "200"

    tflog.Info(ctx,"AMENESIK:ACE: CREATE BEAM INSTANCE: "+template+"-"+program+": SUCCESS")
    return &bi, nil
}

// ----------------------------------------------------------------------
// START BEAM INSTANCE ( template, program )
// ----------------------------------------------------------------------
// Starts the BEAM Application Controller instance as described by the 
// cloned template described by the template and program parameters.
// ----------------------------------------------------------------------
func (c *Client) StartBeamInstance(ctx context.Context,template string, program string) (*BeamResponse, error) {
    tflog.Info(ctx,"AMENESIK:ACE: START BEAM INSTANCE: "+template+"-"+program)
    reqBody := map[string]string{"auth":c.token,"action": "start", "subject": "beam", "account": c.account, "template": UnQuote(template), "program": UnQuote(program) }
    body, _ := json.Marshal(reqBody)

    req, _ := http.NewRequest("POST", c.baseURL, bytes.NewReader(body))
    req.Header.Set("Authorization", "Bearer "+c.token)
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return nil, fmt.Errorf("failed to start instance: %s", resp.Status)
    }

    bodyBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    tflog.Info(ctx,"AMENESIK:ACE: STATUS: "+string(bodyBytes));

    var bi BeamResponse
    bi.status = BeamJsonParser(string(bodyBytes),"status")
    bi.result.name = BeamJsonParser(string(bodyBytes),"id")
    bi.result.status = "200"
    tflog.Info(ctx,"AMENESIK:ACE: START BEAM INSTANCE: "+template+"-"+program+": SUCCESS")
    return &bi, nil
}

// ----------------------------------------------------------------------
// LOCK BEAM INSTANCE ( template, program )
// ----------------------------------------------------------------------
// Locks the BEAM Application Controller instance as described by the 
// cloned template described by the template and program parameters.
// ----------------------------------------------------------------------
func (c *Client) LockBeamInstance(ctx context.Context,template string, program string) (*BeamResponse, error) {
    tflog.Info(ctx,"AMENESIK:ACE: LOCK BEAM INSTANCE: "+template+"-"+program)
    reqBody := map[string]string{"auth":c.token,"action": "lock", "subject": "beam", "account": c.account, "template": UnQuote(template), "program": UnQuote(program) }
    body, _ := json.Marshal(reqBody)

    req, _ := http.NewRequest("POST", c.baseURL, bytes.NewReader(body))
    req.Header.Set("Authorization", "Bearer "+c.token)
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return nil, fmt.Errorf("failed to start instance: %s", resp.Status)
    }

    bodyBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    tflog.Info(ctx,"AMENESIK:ACE: STATUS: "+string(bodyBytes));

    var bi BeamResponse
    bi.status = BeamJsonParser(string(bodyBytes),"status")
    bi.result.name = BeamJsonParser(string(bodyBytes),"id")
    bi.result.status = "200"
    tflog.Info(ctx,"AMENESIK:ACE: LOCK BEAM INSTANCE: "+template+"-"+program+": SUCCESS")
    return &bi, nil
}

// ----------------------------------------------------------------------
// UNLOCK BEAM INSTANCE ( template, program )
// ----------------------------------------------------------------------
// Unlocks the BEAM Application Controller instance as described by the 
// cloned template described by the template and program parameters.
// ----------------------------------------------------------------------
func (c *Client) UnLockBeamInstance(ctx context.Context,template string, program string) (*BeamResponse, error) {
    tflog.Info(ctx,"AMENESIK:ACE: UNLOCK BEAM INSTANCE: "+template+"-"+program)
    reqBody := map[string]string{"auth":c.token,"action": "unlock", "subject": "beam", "account": c.account, "template": UnQuote(template), "program": UnQuote(program) }
    body, _ := json.Marshal(reqBody)

    req, _ := http.NewRequest("POST", c.baseURL, bytes.NewReader(body))
    req.Header.Set("Authorization", "Bearer "+c.token)
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return nil, fmt.Errorf("failed to start instance: %s", resp.Status)
    }

    bodyBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    tflog.Info(ctx,"AMENESIK:ACE: STATUS: "+string(bodyBytes));

    var bi BeamResponse
    bi.status = BeamJsonParser(string(bodyBytes),"status")
    bi.result.name = BeamJsonParser(string(bodyBytes),"id")
    bi.result.status = "200"
    tflog.Info(ctx,"AMENESIK:ACE: UNLOCK BEAM INSTANCE: "+template+"-"+program+": SUCCESS")
    return &bi, nil
}

// ----------------------------------------------------------------------
// STATUS BEAM INSTANCE ( template, program, domain )
// ----------------------------------------------------------------------
// Checks the BEAM Application Controller instance as described by the 
// cloned template described by the template and program parameters.
// ----------------------------------------------------------------------
func (c *Client) StatusBeamInstance(ctx context.Context,template string, program string, domain string,) (*BeamResponse, error) {
    tflog.Info(ctx,"AMENESIK:ACE: STATUS BEAM INSTANCE: "+template+"-"+program)
    reqBody := map[string]string{"auth":c.token,"action": "status", "subject": "beam", "account": c.account, "template": UnQuote(template), "program": UnQuote(program), "domain": UnQuote(domain) }
    body, _ := json.Marshal(reqBody)

    req, _ := http.NewRequest("POST", c.baseURL, bytes.NewReader(body))
    req.Header.Set("Authorization", "Bearer "+c.token)
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return nil, fmt.Errorf("failed to retrieve status of instance: %s", resp.Status)
    }

    bodyBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    tflog.Info(ctx,"AMENESIK:ACE: STATUS: "+string(bodyBytes));

    var bi BeamResponse
    bi.status = BeamJsonParser(string(bodyBytes),"status")
    bi.result.name = BeamJsonParser(string(bodyBytes),"id")
    bi.result.status = "200"
    json.NewDecoder(resp.Body).Decode(&bi)
    return &bi, nil
}

// ----------------------------------------------------------------------
// STOP BEAM INSTANCE ( template, program )
// ----------------------------------------------------------------------
// Stops the BEAM Application Controller instance as described by the 
// cloned template described by the template and program parameters.
// ----------------------------------------------------------------------
func (c *Client) StopBeamInstance(ctx context.Context,template string, program string) (*BeamResponse, error) {
    tflog.Info(ctx,"AMENESIK:ACE: STOP BEAM INSTANCE: "+template+"-"+program)
    reqBody := map[string]string{"auth":c.token,"action": "stop", "subject": "beam", "account": c.account, "template": UnQuote(template), "program": UnQuote(program) }
    body, _ := json.Marshal(reqBody)

    req, _ := http.NewRequest("POST", c.baseURL, bytes.NewReader(body))
    req.Header.Set("Authorization", "Bearer "+c.token)
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return nil, fmt.Errorf("failed to stop instance: %s", resp.Status)
    }

    bodyBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    tflog.Info(ctx,"AMENESIK:ACE: STATUS: "+string(bodyBytes));

    var bi BeamResponse
    bi.status = BeamJsonParser(string(bodyBytes),"status")
    bi.result.name = BeamJsonParser(string(bodyBytes),"id")
    bi.result.status = "200"
    tflog.Info(ctx,"AMENESIK:ACE: STOP BEAM INSTANCE: "+template+"-"+program+": SUCCESS")
    return &bi, nil
}

// ----------------------------------------------------------------------
// SUSPEND BEAM INSTANCE ( template, program )
// ----------------------------------------------------------------------
// Suspends the BEAM Application Controller instance as described by the 
// cloned template described by the template and program parameters.
// ----------------------------------------------------------------------
func (c *Client) SuspendBeamInstance(ctx context.Context,template string, program string) (*BeamResponse, error) {
    tflog.Info(ctx,"AMENESIK:ACE: SUSPEND BEAM INSTANCE: "+template+"-"+program)
    reqBody := map[string]string{"auth":c.token,"action": "suspend", "subject": "beam", "account": c.account, "template": UnQuote(template), "program": UnQuote(program) }
    body, _ := json.Marshal(reqBody)

    req, _ := http.NewRequest("POST", c.baseURL, bytes.NewReader(body))
    req.Header.Set("Authorization", "Bearer "+c.token)
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return nil, fmt.Errorf("failed to suspend instance: %s", resp.Status)
    }

    bodyBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    tflog.Info(ctx,"AMENESIK:ACE: STATUS: "+string(bodyBytes));

    var bi BeamResponse
    bi.status = BeamJsonParser(string(bodyBytes),"status")
    bi.result.name = BeamJsonParser(string(bodyBytes),"id")
    bi.result.status = "200"
    tflog.Info(ctx,"AMENESIK:ACE: SUSPEND BEAM INSTANCE: "+template+"-"+program+": SUCCESS")
    return &bi, nil
}

// ----------------------------------------------------------------------
// RESUME BEAM INSTANCE ( template, program )
// ----------------------------------------------------------------------
// Resumes the BEAM Application Controller instance as described by the 
// cloned template described by the template and program parameters.
// ----------------------------------------------------------------------
func (c *Client) ResumeBeamInstance(ctx context.Context,template string, program string) (*BeamResponse, error) {
    tflog.Info(ctx,"AMENESIK:ACE: RESUME BEAM INSTANCE: "+template+"-"+program)
    reqBody := map[string]string{"auth":c.token,"action": "resume", "subject": "beam", "account": c.account, "template": UnQuote(template), "program": UnQuote(program) }
    body, _ := json.Marshal(reqBody)

    req, _ := http.NewRequest("POST", c.baseURL, bytes.NewReader(body))
    req.Header.Set("Authorization", "Bearer "+c.token)
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return nil, fmt.Errorf("failed to resume instance: %s", resp.Status)
    }

    bodyBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    tflog.Info(ctx,"AMENESIK:ACE: STATUS: "+string(bodyBytes));

    var bi BeamResponse
    bi.status = BeamJsonParser(string(bodyBytes),"status")
    bi.result.name = BeamJsonParser(string(bodyBytes),"id")
    bi.result.status = "200"
    tflog.Info(ctx,"AMENESIK:ACE: RESUME BEAM INSTANCE: "+template+"-"+program+": SUCCESS")
    return &bi, nil
}

// ----------------------------------------------------------------------
// DROP BEAM INSTANCE ( template, program )
// ----------------------------------------------------------------------
// Deletes the BEAM Application Controller instance as described by the 
// cloned template described by the template and program parameters.
// ----------------------------------------------------------------------
func (c *Client) DropBeamInstance(ctx context.Context,template string, program string) (*BeamResponse, error) {
    tflog.Info(ctx,"AMENESIK:ACE: DROP BEAM INSTANCE: "+template+"-"+program)
    reqBody := map[string]string{"auth":c.token,"action": "drop", "subject": "beam", "account": c.account, "template": UnQuote(template), "program": UnQuote(program) }
    body, _ := json.Marshal(reqBody)

    req, _ := http.NewRequest("POST", c.baseURL, bytes.NewReader(body))
    req.Header.Set("Authorization", "Bearer "+c.token)
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return nil, fmt.Errorf("failed to delete instance: %s", resp.Status)
    }

    bodyBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    tflog.Info(ctx,"AMENESIK:ACE: STATUS: "+string(bodyBytes));

    var bi BeamResponse
    bi.status = BeamJsonParser(string(bodyBytes),"status")
    bi.result.name = BeamJsonParser(string(bodyBytes),"id")
    bi.result.status = "200"
    tflog.Info(ctx,"AMENESIK:ACE: DROP BEAM INSTANCE: "+template+"-"+program+": SUCCESS")
    return &bi, nil
}

// ----------------------------------------------------------------------
// DELETE BEAM MODEL ( template, program )
// ----------------------------------------------------------------------
// Deletes the BEAM model described by template and program parameters.
// ----------------------------------------------------------------------
func (c *Client) DeleteBeamModel(ctx context.Context,template string, program string) (*BeamResponse, error) {
    tflog.Info(ctx,"AMENESIK:ACE: DELETE BEAM MODEL: "+template+"-"+program)
    reqBody := map[string]string{"auth":c.token,"action": "delete", "subject": "beam", "account": c.account, "template": UnQuote(template), "program": UnQuote(program) }
    body, _ := json.Marshal(reqBody)

    req, _ := http.NewRequest("POST", c.baseURL, bytes.NewReader(body))
    req.Header.Set("Authorization", "Bearer "+c.token)
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return nil, fmt.Errorf("failed to delete model instance: %s", resp.Status)
    }

    var bi BeamResponse
    bi.status = "deleted"
    bi.result.name = program
    bi.result.status = "200"
    tflog.Info(ctx,"AMENESIK:ACE: DELETE BEAM MODEL: "+template+"-"+program+": SUCCESS")
    return &bi, nil
}

