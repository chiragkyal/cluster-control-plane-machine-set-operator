// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/config/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClusterOperatorStatusConditionApplyConfiguration represents a declarative configuration of the ClusterOperatorStatusCondition type for use
// with apply.
type ClusterOperatorStatusConditionApplyConfiguration struct {
	Type               *v1.ClusterStatusConditionType `json:"type,omitempty"`
	Status             *v1.ConditionStatus            `json:"status,omitempty"`
	LastTransitionTime *metav1.Time                   `json:"lastTransitionTime,omitempty"`
	Reason             *string                        `json:"reason,omitempty"`
	Message            *string                        `json:"message,omitempty"`
}

// ClusterOperatorStatusConditionApplyConfiguration constructs a declarative configuration of the ClusterOperatorStatusCondition type for use with
// apply.
func ClusterOperatorStatusCondition() *ClusterOperatorStatusConditionApplyConfiguration {
	return &ClusterOperatorStatusConditionApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *ClusterOperatorStatusConditionApplyConfiguration) WithType(value v1.ClusterStatusConditionType) *ClusterOperatorStatusConditionApplyConfiguration {
	b.Type = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *ClusterOperatorStatusConditionApplyConfiguration) WithStatus(value v1.ConditionStatus) *ClusterOperatorStatusConditionApplyConfiguration {
	b.Status = &value
	return b
}

// WithLastTransitionTime sets the LastTransitionTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastTransitionTime field is set to the value of the last call.
func (b *ClusterOperatorStatusConditionApplyConfiguration) WithLastTransitionTime(value metav1.Time) *ClusterOperatorStatusConditionApplyConfiguration {
	b.LastTransitionTime = &value
	return b
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *ClusterOperatorStatusConditionApplyConfiguration) WithReason(value string) *ClusterOperatorStatusConditionApplyConfiguration {
	b.Reason = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *ClusterOperatorStatusConditionApplyConfiguration) WithMessage(value string) *ClusterOperatorStatusConditionApplyConfiguration {
	b.Message = &value
	return b
}
