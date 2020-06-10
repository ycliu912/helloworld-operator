// NOTE: Boilerplate only.  Ignore this file.

// Package v1alpha1 contains API Schema definitions for the example v1alpha1 API group
// +k8s:deepcopy-gen=package,register
// +groupName=example.ycliu912.com
package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: "example.ycliu912.com", Version: "v1alpha1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
)
