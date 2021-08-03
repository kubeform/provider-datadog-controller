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

type LevelObjective struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LevelObjectiveSpec   `json:"spec,omitempty"`
	Status            LevelObjectiveStatus `json:"status,omitempty"`
}

type LevelObjectiveSpecQuery struct {
	// The sum of the `total` events.
	Denominator *string `json:"denominator" tf:"denominator"`
	// The sum of all the `good` events.
	Numerator *string `json:"numerator" tf:"numerator"`
}

type LevelObjectiveSpecThresholds struct {
	// The objective's target in`[0,100]`.
	Target *float64 `json:"target" tf:"target"`
	// A string representation of the target that indicates its precision. It uses trailing zeros to show significant decimal places (e.g. `98.00`).
	// +optional
	TargetDisplay *string `json:"targetDisplay,omitempty" tf:"target_display"`
	// The time frame for the objective. The mapping from these types to the types found in the Datadog Web UI can be found in the Datadog API documentation page.
	Timeframe *string `json:"timeframe" tf:"timeframe"`
	// The objective's warning value in `[0,100]`. This must be greater than the target value.
	// +optional
	Warning *float64 `json:"warning,omitempty" tf:"warning"`
	// A string representation of the warning target (see the description of the target_display field for details).
	// +optional
	WarningDisplay *string `json:"warningDisplay,omitempty" tf:"warning_display"`
}

type LevelObjectiveSpec struct {
	State *LevelObjectiveSpecResource `json:"state,omitempty" tf:"-"`

	Resource LevelObjectiveSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type LevelObjectiveSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// A description of this service level objective.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// A boolean indicating whether this monitor can be deleted even if it’s referenced by other resources (e.g. dashboards).
	// +optional
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete"`
	// A static set of groups to filter monitor-based SLOs
	// +optional
	Groups []string `json:"groups,omitempty" tf:"groups"`
	// A static set of monitor IDs to use as part of the SLO
	// +optional
	MonitorIDS []int64 `json:"monitorIDS,omitempty" tf:"monitor_ids"`
	// Name of Datadog service level objective
	Name *string `json:"name" tf:"name"`
	// The metric query of good / total events
	// +optional
	Query *LevelObjectiveSpecQuery `json:"query,omitempty" tf:"query"`
	// A list of tags to associate with your service level objective. This can help you categorize and filter service level objectives in the service level objectives page of the UI. Note: it's not currently possible to filter by these tags when querying via the API
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
	// A list of thresholds and targets that define the service level objectives from the provided SLIs.
	Thresholds []LevelObjectiveSpecThresholds `json:"thresholds" tf:"thresholds"`
	// The type of the service level objective. The mapping from these types to the types found in the Datadog Web UI can be found in the Datadog API [documentation page](https://docs.datadoghq.com/api/v1/service-level-objectives/#create-a-slo-object).
	Type *string `json:"type" tf:"type"`
	// Whether or not to validate the SLO.
	// +optional
	Validate *bool `json:"validate,omitempty" tf:"validate"`
}

type LevelObjectiveStatus struct {
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

// LevelObjectiveList is a list of LevelObjectives
type LevelObjectiveList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LevelObjective CRD objects
	Items []LevelObjective `json:"items,omitempty"`
}