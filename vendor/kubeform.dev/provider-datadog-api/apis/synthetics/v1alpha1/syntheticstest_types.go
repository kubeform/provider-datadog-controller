/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Syntheticstest struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SyntheticstestSpec   `json:"spec,omitempty"`
	Status            SyntheticstestStatus `json:"status,omitempty"`
}

type SyntheticstestSpecApiStepAssertionTargetjsonpath struct {
	// The JSON path to assert.
	Jsonpath *string `json:"jsonpath" tf:"jsonpath"`
	// The specific operator to use on the path.
	Operator *string `json:"operator" tf:"operator"`
	// Expected matching value.
	Targetvalue *string `json:"targetvalue" tf:"targetvalue"`
}

type SyntheticstestSpecApiStepAssertion struct {
	// Assertion operator. **Note** Only some combinations of `type` and `operator` are valid (please refer to [Datadog documentation](https://docs.datadoghq.com/api/latest/synthetics/#create-a-test)).
	Operator *string `json:"operator" tf:"operator"`
	// If assertion type is `header`, this is the header name.
	// +optional
	Property *string `json:"property,omitempty" tf:"property"`
	// Expected value. Depends on the assertion type, refer to [Datadog documentation](https://docs.datadoghq.com/api/latest/synthetics/#create-a-test) for details.
	// +optional
	Target *string `json:"target,omitempty" tf:"target"`
	// Expected structure if `operator` is `validatesJSONPath`. Exactly one nested block is allowed with the structure below.
	// +optional
	Targetjsonpath *SyntheticstestSpecApiStepAssertionTargetjsonpath `json:"targetjsonpath,omitempty" tf:"targetjsonpath"`
	// Type of assertion. **Note** Only some combinations of `type` and `operator` are valid (please refer to [Datadog documentation](https://docs.datadoghq.com/api/latest/synthetics/#create-a-test)).
	Type *string `json:"type" tf:"type"`
}

type SyntheticstestSpecApiStepExtractedValueParser struct {
	// Type of parser for a Synthetics global variable from a synthetics test.
	Type *string `json:"type" tf:"type"`
	// Regex or JSON path used for the parser. Not used with type `raw`.
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type SyntheticstestSpecApiStepExtractedValue struct {
	// When type is `http_header`, name of the header to use to extract the value.
	// +optional
	Field  *string                                        `json:"field,omitempty" tf:"field"`
	Name   *string                                        `json:"name" tf:"name"`
	Parser *SyntheticstestSpecApiStepExtractedValueParser `json:"parser" tf:"parser"`
	// Property of the Synthetics Test Response to use for the variable.
	Type *string `json:"type" tf:"type"`
}

type SyntheticstestSpecApiStepRequestBasicauth struct {
	// Password for authentication.
	Password *string `json:"-" sensitive:"true" tf:"password"`
	// Username for authentication.
	Username *string `json:"username" tf:"username"`
}

type SyntheticstestSpecApiStepRequestClientCertificateCert struct {
	// Content of the certificate.
	Content *string `json:"-" sensitive:"true" tf:"content"`
	// File name for the certificate.
	// +optional
	Filename *string `json:"filename,omitempty" tf:"filename"`
}

type SyntheticstestSpecApiStepRequestClientCertificateKey struct {
	// Content of the certificate.
	Content *string `json:"-" sensitive:"true" tf:"content"`
	// File name for the certificate.
	// +optional
	Filename *string `json:"filename,omitempty" tf:"filename"`
}

type SyntheticstestSpecApiStepRequestClientCertificate struct {
	Cert *SyntheticstestSpecApiStepRequestClientCertificateCert `json:"cert" tf:"cert"`
	Key  *SyntheticstestSpecApiStepRequestClientCertificateKey  `json:"key" tf:"key"`
}

type SyntheticstestSpecApiStepRequestDefinition struct {
	// Allows loading insecure content for an HTTP test.
	// +optional
	AllowInsecure *bool `json:"allowInsecure,omitempty" tf:"allow_insecure"`
	// The request body.
	// +optional
	Body *string `json:"body,omitempty" tf:"body"`
	// DNS server to use for DNS tests (`subtype = "dns"`).
	// +optional
	DnsServer *string `json:"dnsServer,omitempty" tf:"dns_server"`
	// DNS server port to use for DNS tests.
	// +optional
	DnsServerPort *int64 `json:"dnsServerPort,omitempty" tf:"dns_server_port"`
	// Determines whether or not the API HTTP test should follow redirects.
	// +optional
	FollowRedirects *bool `json:"followRedirects,omitempty" tf:"follow_redirects"`
	// Host name to perform the test with.
	// +optional
	Host *string `json:"host,omitempty" tf:"host"`
	// For UDP and websocket tests, message to send with the request.
	// +optional
	Message *string `json:"message,omitempty" tf:"message"`
	// The HTTP method.
	// +optional
	Method *string `json:"method,omitempty" tf:"method"`
	// Determines whether or not to save the response body.
	// +optional
	NoSavingResponseBody *bool `json:"noSavingResponseBody,omitempty" tf:"no_saving_response_body"`
	// Number of pings to use per test for ICMP tests (`subtype = "icmp"`) between 0 and 10.
	// +optional
	NumberOfPackets *int64 `json:"numberOfPackets,omitempty" tf:"number_of_packets"`
	// Port to use when performing the test.
	// +optional
	Port *int64 `json:"port,omitempty" tf:"port"`
	// For SSL tests, it specifies on which server you want to initiate the TLS handshake, allowing the server to present one of multiple possible certificates on the same IP address and TCP port number.
	// +optional
	Servername *string `json:"servername,omitempty" tf:"servername"`
	// This will turn on a traceroute probe to discover all gateways along the path to the host destination. For ICMP tests (`subtype = "icmp"`).
	// +optional
	ShouldTrackHops *bool `json:"shouldTrackHops,omitempty" tf:"should_track_hops"`
	// Timeout in seconds for the test. Defaults to `60`.
	// +optional
	Timeout *int64 `json:"timeout,omitempty" tf:"timeout"`
	// The URL to send the request to.
	// +optional
	Url *string `json:"url,omitempty" tf:"url"`
}

type SyntheticstestSpecApiStepRetry struct {
	// Number of retries needed to consider a location as failed before sending a notification alert.
	// +optional
	Count *int64 `json:"count,omitempty" tf:"count"`
	// Interval between a failed test and the next retry in milliseconds.
	// +optional
	Interval *int64 `json:"interval,omitempty" tf:"interval"`
}

type SyntheticstestSpecApiStep struct {
	// Determines whether or not to continue with test if this step fails.
	// +optional
	AllowFailure *bool `json:"allowFailure,omitempty" tf:"allow_failure"`
	// Assertions used for the test. Multiple `assertion` blocks are allowed with the structure below.
	// +optional
	Assertion []SyntheticstestSpecApiStepAssertion `json:"assertion,omitempty" tf:"assertion"`
	// Values to parse and save as variables from the response.
	// +optional
	ExtractedValue []SyntheticstestSpecApiStepExtractedValue `json:"extractedValue,omitempty" tf:"extracted_value"`
	// Determines whether or not to consider the entire test as failed if this step fails. Can be used only if `allow_failure` is `true`.
	// +optional
	IsCritical *bool `json:"isCritical,omitempty" tf:"is_critical"`
	// The name of the step.
	Name *string `json:"name" tf:"name"`
	// The HTTP basic authentication credentials. Exactly one nested block is allowed with the structure below.
	// +optional
	RequestBasicauth *SyntheticstestSpecApiStepRequestBasicauth `json:"requestBasicauth,omitempty" tf:"request_basicauth"`
	// Client certificate to use when performing the test request. Exactly one nested block is allowed with the structure below.
	// +optional
	RequestClientCertificate *SyntheticstestSpecApiStepRequestClientCertificate `json:"requestClientCertificate,omitempty" tf:"request_client_certificate"`
	// The request for the api step.
	// +optional
	RequestDefinition *SyntheticstestSpecApiStepRequestDefinition `json:"requestDefinition,omitempty" tf:"request_definition"`
	// Header name and value map.
	// +optional
	RequestHeaders map[string]string `json:"requestHeaders,omitempty" tf:"request_headers"`
	// Query arguments name and value map.
	// +optional
	RequestQuery map[string]string `json:"requestQuery,omitempty" tf:"request_query"`
	// +optional
	Retry *SyntheticstestSpecApiStepRetry `json:"retry,omitempty" tf:"retry"`
	// The subtype of the Synthetic multistep API test step.
	// +optional
	Subtype *string `json:"subtype,omitempty" tf:"subtype"`
}

type SyntheticstestSpecAssertionTargetjsonpath struct {
	// The JSON path to assert.
	Jsonpath *string `json:"jsonpath" tf:"jsonpath"`
	// The specific operator to use on the path.
	Operator *string `json:"operator" tf:"operator"`
	// Expected matching value.
	Targetvalue *string `json:"targetvalue" tf:"targetvalue"`
}

type SyntheticstestSpecAssertion struct {
	// Assertion operator. **Note** Only some combinations of `type` and `operator` are valid (please refer to [Datadog documentation](https://docs.datadoghq.com/api/latest/synthetics/#create-a-test)).
	Operator *string `json:"operator" tf:"operator"`
	// If assertion type is `header`, this is the header name.
	// +optional
	Property *string `json:"property,omitempty" tf:"property"`
	// Expected value. Depends on the assertion type, refer to [Datadog documentation](https://docs.datadoghq.com/api/latest/synthetics/#create-a-test) for details.
	// +optional
	Target *string `json:"target,omitempty" tf:"target"`
	// Expected structure if `operator` is `validatesJSONPath`. Exactly one nested block is allowed with the structure below.
	// +optional
	Targetjsonpath *SyntheticstestSpecAssertionTargetjsonpath `json:"targetjsonpath,omitempty" tf:"targetjsonpath"`
	// Type of assertion. **Note** Only some combinations of `type` and `operator` are valid (please refer to [Datadog documentation](https://docs.datadoghq.com/api/latest/synthetics/#create-a-test)).
	Type *string `json:"type" tf:"type"`
}

type SyntheticstestSpecBrowserStepParamsVariable struct {
	// Example of the extracted variable.
	// +optional
	Example *string `json:"example,omitempty" tf:"example"`
	// Name of the extracted variable.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
}

type SyntheticstestSpecBrowserStepParams struct {
	// Name of the attribute to use for an "assert attribute" step.
	// +optional
	Attribute *string `json:"attribute,omitempty" tf:"attribute"`
	// Check type to use for an assertion step.
	// +optional
	Check *string `json:"check,omitempty" tf:"check"`
	// Type of click to use for a "click" step.
	// +optional
	ClickType *string `json:"clickType,omitempty" tf:"click_type"`
	// Javascript code to use for the step.
	// +optional
	Code *string `json:"code,omitempty" tf:"code"`
	// Delay between each key stroke for a "type test" step.
	// +optional
	Delay *int64 `json:"delay,omitempty" tf:"delay"`
	// Element to use for the step, json encoded string.
	// +optional
	Element *string `json:"element,omitempty" tf:"element"`
	// Details of the email for an "assert email" step.
	// +optional
	Email *string `json:"email,omitempty" tf:"email"`
	// For an "assert download" step.
	// +optional
	File *string `json:"file,omitempty" tf:"file"`
	// Details of the files for an "upload files" step, json encoded string.
	// +optional
	Files *string `json:"files,omitempty" tf:"files"`
	// Modifier to use for a "press key" step.
	// +optional
	Modifiers []string `json:"modifiers,omitempty" tf:"modifiers"`
	// ID of the tab to play the subtest.
	// +optional
	PlayingTabID *string `json:"playingTabID,omitempty" tf:"playing_tab_id"`
	// Request for an API step.
	// +optional
	Request *string `json:"request,omitempty" tf:"request"`
	// ID of the Synthetics test to use as subtest.
	// +optional
	SubtestPublicID *string `json:"subtestPublicID,omitempty" tf:"subtest_public_id"`
	// Value of the step.
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
	// Details of the variable to extract.
	// +optional
	Variable *SyntheticstestSpecBrowserStepParamsVariable `json:"variable,omitempty" tf:"variable"`
	// For "file upload" steps.
	// +optional
	WithClick *bool `json:"withClick,omitempty" tf:"with_click"`
	// X coordinates for a "scroll step".
	// +optional
	X *int64 `json:"x,omitempty" tf:"x"`
	// Y coordinates for a "scroll step".
	// +optional
	Y *int64 `json:"y,omitempty" tf:"y"`
}

type SyntheticstestSpecBrowserStep struct {
	// Determines if the step should be allowed to fail.
	// +optional
	AllowFailure *bool `json:"allowFailure,omitempty" tf:"allow_failure"`
	// Force update of the "element" parameter for the step
	// +optional
	ForceElementUpdate *bool `json:"forceElementUpdate,omitempty" tf:"force_element_update"`
	// Name of the step.
	Name *string `json:"name" tf:"name"`
	// Parameters for the step.
	Params *SyntheticstestSpecBrowserStepParams `json:"params" tf:"params"`
	// Used to override the default timeout of a step.
	// +optional
	Timeout *int64 `json:"timeout,omitempty" tf:"timeout"`
	// Type of the step.
	Type *string `json:"type" tf:"type"`
}

type SyntheticstestSpecBrowserVariable struct {
	// Example for the variable.
	// +optional
	Example *string `json:"example,omitempty" tf:"example"`
	// ID of the global variable to use. This is actually only used (and required) in the case of using a variable of type `global`.
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// Name of the variable.
	Name *string `json:"name" tf:"name"`
	// Pattern of the variable.
	// +optional
	Pattern *string `json:"pattern,omitempty" tf:"pattern"`
	// Type of browser test variable.
	Type *string `json:"type" tf:"type"`
}

type SyntheticstestSpecConfigVariable struct {
	// Example for the variable.
	// +optional
	Example *string `json:"example,omitempty" tf:"example"`
	// When type = `global`, ID of the global variable to use.
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// Name of the variable.
	Name *string `json:"name" tf:"name"`
	// Pattern of the variable.
	// +optional
	Pattern *string `json:"pattern,omitempty" tf:"pattern"`
	// Type of test configuration variable.
	Type *string `json:"type" tf:"type"`
}

type SyntheticstestSpecOptionsListMonitorOptions struct {
	// Specify a renotification frequency.
	// +optional
	RenotifyInterval *int64 `json:"renotifyInterval,omitempty" tf:"renotify_interval"`
}

type SyntheticstestSpecOptionsListRetry struct {
	// Number of retries needed to consider a location as failed before sending a notification alert.
	// +optional
	Count *int64 `json:"count,omitempty" tf:"count"`
	// Interval between a failed test and the next retry in milliseconds.
	// +optional
	Interval *int64 `json:"interval,omitempty" tf:"interval"`
}

type SyntheticstestSpecOptionsList struct {
	// For SSL test, whether or not the test should allow self signed certificates.
	// +optional
	AcceptSelfSigned *bool `json:"acceptSelfSigned,omitempty" tf:"accept_self_signed"`
	// Allows loading insecure content for an HTTP test.
	// +optional
	AllowInsecure *bool `json:"allowInsecure,omitempty" tf:"allow_insecure"`
	// Determines whether or not the API HTTP test should follow redirects.
	// +optional
	FollowRedirects *bool `json:"followRedirects,omitempty" tf:"follow_redirects"`
	// Minimum amount of time in failure required to trigger an alert. Default is `0`.
	// +optional
	MinFailureDuration *int64 `json:"minFailureDuration,omitempty" tf:"min_failure_duration"`
	// Minimum number of locations in failure required to trigger an alert. Default is `1`.
	// +optional
	MinLocationFailed *int64 `json:"minLocationFailed,omitempty" tf:"min_location_failed"`
	// The monitor name is used for the alert title as well as for all monitor dashboard widgets and SLOs.
	// +optional
	MonitorName *string `json:"monitorName,omitempty" tf:"monitor_name"`
	// +optional
	MonitorOptions *SyntheticstestSpecOptionsListMonitorOptions `json:"monitorOptions,omitempty" tf:"monitor_options"`
	// +optional
	MonitorPriority *int64 `json:"monitorPriority,omitempty" tf:"monitor_priority"`
	// Prevents saving screenshots of the steps.
	// +optional
	NoScreenshot *bool `json:"noScreenshot,omitempty" tf:"no_screenshot"`
	// +optional
	Retry *SyntheticstestSpecOptionsListRetry `json:"retry,omitempty" tf:"retry"`
	// How often the test should run (in seconds).
	TickEvery *int64 `json:"tickEvery" tf:"tick_every"`
}

type SyntheticstestSpecRequestBasicauth struct {
	// Password for authentication.
	Password *string `json:"-" sensitive:"true" tf:"password"`
	// Username for authentication.
	Username *string `json:"username" tf:"username"`
}

type SyntheticstestSpecRequestClientCertificateCert struct {
	// Content of the certificate.
	Content *string `json:"-" sensitive:"true" tf:"content"`
	// File name for the certificate.
	// +optional
	Filename *string `json:"filename,omitempty" tf:"filename"`
}

type SyntheticstestSpecRequestClientCertificateKey struct {
	// Content of the certificate.
	Content *string `json:"-" sensitive:"true" tf:"content"`
	// File name for the certificate.
	// +optional
	Filename *string `json:"filename,omitempty" tf:"filename"`
}

type SyntheticstestSpecRequestClientCertificate struct {
	Cert *SyntheticstestSpecRequestClientCertificateCert `json:"cert" tf:"cert"`
	Key  *SyntheticstestSpecRequestClientCertificateKey  `json:"key" tf:"key"`
}

type SyntheticstestSpecRequestDefinition struct {
	// The request body.
	// +optional
	Body *string `json:"body,omitempty" tf:"body"`
	// DNS server to use for DNS tests (`subtype = "dns"`).
	// +optional
	DnsServer *string `json:"dnsServer,omitempty" tf:"dns_server"`
	// DNS server port to use for DNS tests.
	// +optional
	DnsServerPort *int64 `json:"dnsServerPort,omitempty" tf:"dns_server_port"`
	// Host name to perform the test with.
	// +optional
	Host *string `json:"host,omitempty" tf:"host"`
	// For UDP and websocket tests, message to send with the request.
	// +optional
	Message *string `json:"message,omitempty" tf:"message"`
	// The HTTP method.
	// +optional
	Method *string `json:"method,omitempty" tf:"method"`
	// Determines whether or not to save the response body.
	// +optional
	NoSavingResponseBody *bool `json:"noSavingResponseBody,omitempty" tf:"no_saving_response_body"`
	// Number of pings to use per test for ICMP tests (`subtype = "icmp"`) between 0 and 10.
	// +optional
	NumberOfPackets *int64 `json:"numberOfPackets,omitempty" tf:"number_of_packets"`
	// Port to use when performing the test.
	// +optional
	Port *int64 `json:"port,omitempty" tf:"port"`
	// For SSL tests, it specifies on which server you want to initiate the TLS handshake, allowing the server to present one of multiple possible certificates on the same IP address and TCP port number.
	// +optional
	Servername *string `json:"servername,omitempty" tf:"servername"`
	// This will turn on a traceroute probe to discover all gateways along the path to the host destination. For ICMP tests (`subtype = "icmp"`).
	// +optional
	ShouldTrackHops *bool `json:"shouldTrackHops,omitempty" tf:"should_track_hops"`
	// Timeout in seconds for the test. Defaults to `60`.
	// +optional
	Timeout *int64 `json:"timeout,omitempty" tf:"timeout"`
	// The URL to send the request to.
	// +optional
	Url *string `json:"url,omitempty" tf:"url"`
}

type SyntheticstestSpec struct {
	State *SyntheticstestSpecResource `json:"state,omitempty" tf:"-"`

	Resource SyntheticstestSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type SyntheticstestSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Steps for multistep api tests
	// +optional
	ApiStep []SyntheticstestSpecApiStep `json:"apiStep,omitempty" tf:"api_step"`
	// Assertions used for the test. Multiple `assertion` blocks are allowed with the structure below.
	// +optional
	Assertion []SyntheticstestSpecAssertion `json:"assertion,omitempty" tf:"assertion"`
	// Steps for browser tests.
	// +optional
	BrowserStep []SyntheticstestSpecBrowserStep `json:"browserStep,omitempty" tf:"browser_step"`
	// Variables used for a browser test steps. Multiple `variable` blocks are allowed with the structure below.
	// +optional
	BrowserVariable []SyntheticstestSpecBrowserVariable `json:"browserVariable,omitempty" tf:"browser_variable"`
	// Variables used for the test configuration. Multiple `config_variable` blocks are allowed with the structure below.
	// +optional
	ConfigVariable []SyntheticstestSpecConfigVariable `json:"configVariable,omitempty" tf:"config_variable"`
	// Required if `type = "browser"`. Array with the different device IDs used to run the test.
	// +optional
	DeviceIDS []string `json:"deviceIDS,omitempty" tf:"device_ids"`
	// Array of locations used to run the test. Refer to [Datadog documentation](https://docs.datadoghq.com/synthetics/api_test/#request) for available locations (e.g. `aws:eu-central-1`).
	Locations []string `json:"locations" tf:"locations"`
	// A message to include with notifications for this synthetics test. Email notifications can be sent to specific users by using the same `@username` notation as events.
	// +optional
	Message *string `json:"message,omitempty" tf:"message"`
	// ID of the monitor associated with the Datadog synthetics test.
	// +optional
	MonitorID *int64 `json:"monitorID,omitempty" tf:"monitor_id"`
	// Name of Datadog synthetics test.
	Name *string `json:"name" tf:"name"`
	// +optional
	OptionsList *SyntheticstestSpecOptionsList `json:"optionsList,omitempty" tf:"options_list"`
	// The HTTP basic authentication credentials. Exactly one nested block is allowed with the structure below.
	// +optional
	RequestBasicauth *SyntheticstestSpecRequestBasicauth `json:"requestBasicauth,omitempty" tf:"request_basicauth"`
	// Client certificate to use when performing the test request. Exactly one nested block is allowed with the structure below.
	// +optional
	RequestClientCertificate *SyntheticstestSpecRequestClientCertificate `json:"requestClientCertificate,omitempty" tf:"request_client_certificate"`
	// Required if `type = "api"`. The synthetics test request.
	// +optional
	RequestDefinition *SyntheticstestSpecRequestDefinition `json:"requestDefinition,omitempty" tf:"request_definition"`
	// Header name and value map.
	// +optional
	RequestHeaders map[string]string `json:"requestHeaders,omitempty" tf:"request_headers"`
	// Query arguments name and value map.
	// +optional
	RequestQuery map[string]string `json:"requestQuery,omitempty" tf:"request_query"`
	// Cookies to be used for a browser test request, using the [Set-Cookie](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie) syntax.
	// +optional
	SetCookie *string `json:"setCookie,omitempty" tf:"set_cookie"`
	// Define whether you want to start (`live`) or pause (`paused`) a Synthetic test.
	Status *string `json:"status" tf:"status"`
	// The subtype of the Synthetic API test. Defaults to `http`.
	// +optional
	Subtype *string `json:"subtype,omitempty" tf:"subtype"`
	// A list of tags to associate with your synthetics test. This can help you categorize and filter tests in the manage synthetics page of the UI. Default is an empty list (`[]`).
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
	// Synthetics test type.
	Type *string `json:"type" tf:"type"`
}

type SyntheticstestStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SyntheticstestList is a list of Syntheticstests
type SyntheticstestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Syntheticstest CRD objects
	Items []Syntheticstest `json:"items,omitempty"`
}
