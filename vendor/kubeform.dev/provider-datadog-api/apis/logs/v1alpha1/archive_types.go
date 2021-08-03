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

type Archive struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ArchiveSpec   `json:"spec,omitempty"`
	Status            ArchiveStatus `json:"status,omitempty"`
}

type ArchiveSpecAzureArchive struct {
	// Your client id.
	ClientID *string `json:"clientID" tf:"client_id"`
	// The container where the archive will be stored.
	Container *string `json:"container" tf:"container"`
	// The path where the archive will be stored.
	// +optional
	Path *string `json:"path,omitempty" tf:"path"`
	// The associated storage account.
	StorageAccount *string `json:"storageAccount" tf:"storage_account"`
	// Your tenant id.
	TenantID *string `json:"tenantID" tf:"tenant_id"`
}

type ArchiveSpecGcsArchive struct {
	// Name of your GCS bucket.
	Bucket *string `json:"bucket" tf:"bucket"`
	// Your client email.
	ClientEmail *string `json:"clientEmail" tf:"client_email"`
	// Path where the archive will be stored.
	Path *string `json:"path" tf:"path"`
	// Your project id.
	ProjectID *string `json:"projectID" tf:"project_id"`
}

type ArchiveSpecS3Archive struct {
	// Your AWS account id.
	AccountID *string `json:"accountID" tf:"account_id"`
	// Name of your s3 bucket.
	Bucket *string `json:"bucket" tf:"bucket"`
	// Path where the archive will be stored.
	Path *string `json:"path" tf:"path"`
	// Your AWS role name
	RoleName *string `json:"roleName" tf:"role_name"`
}

type ArchiveSpec struct {
	State *ArchiveSpecResource `json:"state,omitempty" tf:"-"`

	Resource ArchiveSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ArchiveSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Definition of an azure archive.
	// +optional
	AzureArchive *ArchiveSpecAzureArchive `json:"azureArchive,omitempty" tf:"azure_archive"`
	// Definition of a GCS archive.
	// +optional
	GcsArchive *ArchiveSpecGcsArchive `json:"gcsArchive,omitempty" tf:"gcs_archive"`
	// To store the tags in the archive, set the value `true`. If it is set to `false`, the tags will be dropped when the logs are sent to the archive.
	// +optional
	IncludeTags *bool `json:"includeTags,omitempty" tf:"include_tags"`
	// Your archive name.
	Name *string `json:"name" tf:"name"`
	// The archive query/filter. Logs matching this query are included in the archive.
	Query *string `json:"query" tf:"query"`
	// An array of tags to add to rehydrated logs from an archive.
	// +optional
	RehydrationTags []string `json:"rehydrationTags,omitempty" tf:"rehydration_tags"`
	// Definition of an s3 archive.
	// +optional
	S3Archive *ArchiveSpecS3Archive `json:"s3Archive,omitempty" tf:"s3_archive"`
}

type ArchiveStatus struct {
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

// ArchiveList is a list of Archives
type ArchiveList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Archive CRD objects
	Items []Archive `json:"items,omitempty"`
}