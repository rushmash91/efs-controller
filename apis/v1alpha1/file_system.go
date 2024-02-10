// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FileSystemSpec defines the desired state of FileSystem.
type FileSystemSpec struct {

	// Used to create a One Zone file system. It specifies the Amazon Web Services
	// Availability Zone in which to create the file system. Use the format us-east-1a
	// to specify the Availability Zone. For more information about One Zone file
	// systems, see Using EFS storage classes (https://docs.aws.amazon.com/efs/latest/ug/storage-classes.html)
	// in the Amazon EFS User Guide.
	//
	// One Zone file systems are not available in all Availability Zones in Amazon
	// Web Services Regions where Amazon EFS is available.
	AvailabilityZoneName *string `json:"availabilityZoneName,omitempty"`
	// Specifies whether automatic backups are enabled on the file system that you
	// are creating. Set the value to true to enable automatic backups. If you are
	// creating a One Zone file system, automatic backups are enabled by default.
	// For more information, see Automatic backups (https://docs.aws.amazon.com/efs/latest/ug/awsbackup.html#automatic-backups)
	// in the Amazon EFS User Guide.
	//
	// Default is false. However, if you specify an AvailabilityZoneName, the default
	// is true.
	//
	// Backup is not available in all Amazon Web Services Regions where Amazon EFS
	// is available.
	Backup *bool `json:"backup,omitempty"`
	// The backup policy included in the PutBackupPolicy request.
	BackupPolicy *BackupPolicy `json:"backupPolicy,omitempty"`
	// A Boolean value that, if true, creates an encrypted file system. When creating
	// an encrypted file system, you have the option of specifying an existing Key
	// Management Service key (KMS key). If you don't specify a KMS key, then the
	// default KMS key for Amazon EFS, /aws/elasticfilesystem, is used to protect
	// the encrypted file system.
	Encrypted            *bool                            `json:"encrypted,omitempty"`
	FileSystemProtection *UpdateFileSystemProtectionInput `json:"fileSystemProtection,omitempty"`
	// The ID of the KMS key that you want to use to protect the encrypted file
	// system. This parameter is required only if you want to use a non-default
	// KMS key. If this parameter is not specified, the default KMS key for Amazon
	// EFS is used. You can specify a KMS key ID using the following formats:
	//
	//   - Key ID - A unique identifier of the key, for example 1234abcd-12ab-34cd-56ef-1234567890ab.
	//
	//   - ARN - An Amazon Resource Name (ARN) for the key, for example arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab.
	//
	//   - Key alias - A previously created display name for a key, for example
	//     alias/projectKey1.
	//
	//   - Key alias ARN - An ARN for a key alias, for example arn:aws:kms:us-west-2:444455556666:alias/projectKey1.
	//
	// If you use KmsKeyId, you must set the CreateFileSystemRequest$Encrypted parameter
	// to true.
	//
	// EFS accepts only symmetric KMS keys. You cannot use asymmetric KMS keys with
	// Amazon EFS file systems.
	KMSKeyID  *string                                  `json:"kmsKeyID,omitempty"`
	KMSKeyRef *ackv1alpha1.AWSResourceReferenceWrapper `json:"kmsKeyRef,omitempty"`
	// An array of LifecyclePolicy objects that define the file system's LifecycleConfiguration
	// object. A LifecycleConfiguration object informs EFS Lifecycle management
	// of the following:
	//
	//   - TransitionToIA – When to move files in the file system from primary
	//     storage (Standard storage class) into the Infrequent Access (IA) storage.
	//
	//   - TransitionToArchive – When to move files in the file system from their
	//     current storage class (either IA or Standard storage) into the Archive
	//     storage. File systems cannot transition into Archive storage before transitioning
	//     into IA storage. Therefore, TransitionToArchive must either not be set
	//     or must be later than TransitionToIA. The Archive storage class is available
	//     only for file systems that use the Elastic Throughput mode and the General
	//     Purpose Performance mode.
	//
	//   - TransitionToPrimaryStorageClass – Whether to move files in the file
	//     system back to primary storage (Standard storage class) after they are
	//     accessed in IA or Archive storage.
	//
	// When using the put-lifecycle-configuration CLI command or the PutLifecycleConfiguration
	// API action, Amazon EFS requires that each LifecyclePolicy object have only
	// a single transition. This means that in a request body, LifecyclePolicies
	// must be structured as an array of LifecyclePolicy objects, one object for
	// each storage transition. See the example requests in the following section
	// for more information.
	LifecyclePolicies []*LifecyclePolicy `json:"lifecyclePolicies,omitempty"`
	// The Performance mode of the file system. We recommend generalPurpose performance
	// mode for all file systems. File systems using the maxIO performance mode
	// can scale to higher levels of aggregate throughput and operations per second
	// with a tradeoff of slightly higher latencies for most file operations. The
	// performance mode can't be changed after the file system has been created.
	// The maxIO mode is not supported on One Zone file systems.
	//
	// Due to the higher per-operation latencies with Max I/O, we recommend using
	// General Purpose performance mode for all file systems.
	//
	// Default is generalPurpose.
	PerformanceMode *string `json:"performanceMode,omitempty"`
	// The FileSystemPolicy that you're creating. Accepts a JSON formatted policy
	// definition. EFS file system policies have a 20,000 character limit. To find
	// out more about the elements that make up a file system policy, see EFS Resource-based
	// Policies (https://docs.aws.amazon.com/efs/latest/ug/access-control-overview.html#access-control-manage-access-intro-resource-policies).
	Policy *string `json:"policy,omitempty"`
	// The throughput, measured in mebibytes per second (MiBps), that you want to
	// provision for a file system that you're creating. Required if ThroughputMode
	// is set to provisioned. Valid values are 1-3414 MiBps, with the upper limit
	// depending on Region. To increase this limit, contact Amazon Web Services
	// Support. For more information, see Amazon EFS quotas that you can increase
	// (https://docs.aws.amazon.com/efs/latest/ug/limits.html#soft-limits) in the
	// Amazon EFS User Guide.
	ProvisionedThroughputInMiBps *float64 `json:"provisionedThroughputInMiBps,omitempty"`
	// Use to create one or more tags associated with the file system. Each tag
	// is a user-defined key-value pair. Name your file system on creation by including
	// a "Key":"Name","Value":"{value}" key-value pair. Each key must be unique.
	// For more information, see Tagging Amazon Web Services resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// in the Amazon Web Services General Reference Guide.
	Tags []*Tag `json:"tags,omitempty"`
	// Specifies the throughput mode for the file system. The mode can be bursting,
	// provisioned, or elastic. If you set ThroughputMode to provisioned, you must
	// also set a value for ProvisionedThroughputInMibps. After you create the file
	// system, you can decrease your file system's Provisioned throughput or change
	// between the throughput modes, with certain time restrictions. For more information,
	// see Specifying throughput with provisioned mode (https://docs.aws.amazon.com/efs/latest/ug/performance.html#provisioned-throughput)
	// in the Amazon EFS User Guide.
	//
	// Default is bursting.
	ThroughputMode *string `json:"throughputMode,omitempty"`
}

// FileSystemStatus defines the observed state of FileSystem
type FileSystemStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The unique and consistent identifier of the Availability Zone in which the
	// file system is located, and is valid only for One Zone file systems. For
	// example, use1-az1 is an Availability Zone ID for the us-east-1 Amazon Web
	// Services Region, and it has the same location in every Amazon Web Services
	// account.
	// +kubebuilder:validation:Optional
	AvailabilityZoneID *string `json:"availabilityZoneID,omitempty"`
	// The time that the file system was created, in seconds (since 1970-01-01T00:00:00Z).
	// +kubebuilder:validation:Optional
	CreationTime *metav1.Time `json:"creationTime,omitempty"`
	// The ID of the file system, assigned by Amazon EFS.
	// +kubebuilder:validation:Optional
	FileSystemID *string `json:"fileSystemID,omitempty"`
	// The lifecycle phase of the file system.
	// +kubebuilder:validation:Optional
	LifeCycleState *string `json:"lifeCycleState,omitempty"`
	// You can add tags to a file system, including a Name tag. For more information,
	// see CreateFileSystem. If the file system has a Name tag, Amazon EFS returns
	// the value in this field.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty"`
	// The current number of mount targets that the file system has. For more information,
	// see CreateMountTarget.
	// +kubebuilder:validation:Optional
	NumberOfMountTargets *int64 `json:"numberOfMountTargets,omitempty"`
	// The Amazon Web Services account that created the file system.
	// +kubebuilder:validation:Optional
	OwnerID *string `json:"ownerID,omitempty"`
	// The latest known metered size (in bytes) of data stored in the file system,
	// in its Value field, and the time at which that size was determined in its
	// Timestamp field. The Timestamp value is the integer number of seconds since
	// 1970-01-01T00:00:00Z. The SizeInBytes value doesn't represent the size of
	// a consistent snapshot of the file system, but it is eventually consistent
	// when there are no writes to the file system. That is, SizeInBytes represents
	// actual size only if the file system is not modified for a period longer than
	// a couple of hours. Otherwise, the value is not the exact size that the file
	// system was at any point in time.
	// +kubebuilder:validation:Optional
	SizeInBytes *FileSystemSize `json:"sizeInBytes,omitempty"`
}

// FileSystem is the Schema for the FileSystems API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="ID",type=string,priority=0,JSONPath=`.status.fileSystemID`
// +kubebuilder:printcolumn:name="ENCRYPTED",type=boolean,priority=0,JSONPath=`.spec.encrypted`
// +kubebuilder:printcolumn:name="PERFORMANCEMODE",type=string,priority=1,JSONPath=`.spec.performanceMode`
// +kubebuilder:printcolumn:name="THROUGHPUTMODE",type=string,priority=1,JSONPath=`.spec.throughputMode`
// +kubebuilder:printcolumn:name="PROVISIONEDTHROUGHPUT",type=string,priority=1,JSONPath=`.status.provisionedThroughputInMiBps`
// +kubebuilder:printcolumn:name="SIZE",type=integer,priority=0,JSONPath=`.status.sizeInBytes.value`
// +kubebuilder:printcolumn:name="MOUNTTARGETS",type=integer,priority=0,JSONPath=`.status.numberOfMountTargets`
// +kubebuilder:printcolumn:name="STATE",type=string,priority=0,JSONPath=`.status.lifeCycleState`
// +kubebuilder:printcolumn:name="Synced",type="string",priority=0,JSONPath=".status.conditions[?(@.type==\"ACK.ResourceSynced\")].status"
// +kubebuilder:printcolumn:name="Age",type="date",priority=0,JSONPath=".metadata.creationTimestamp"
type FileSystem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FileSystemSpec   `json:"spec,omitempty"`
	Status            FileSystemStatus `json:"status,omitempty"`
}

// FileSystemList contains a list of FileSystem
// +kubebuilder:object:root=true
type FileSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FileSystem `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FileSystem{}, &FileSystemList{})
}
