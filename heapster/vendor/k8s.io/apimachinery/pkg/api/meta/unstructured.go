/*
Copyright 2016 The Kubernetes Authors.

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

package meta

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// InterfacesForUnstructuredConversion returns VersionInterfaces suitable for
// dealing with unstructured.Unstructured objects and supports conversion
// from typed objects (provided by parent) to untyped objects.
func InterfacesForUnstructuredConversion(parent VersionInterfacesFunc) VersionInterfacesFunc {
	return func(version schema.GroupVersion) (*VersionInterfaces, error) {
		if i, err := parent(version); err == nil {
			return &VersionInterfaces{
				ObjectConvertor:  i.ObjectConvertor,
				MetadataAccessor: NewAccessor(),
			}, nil
		}
		return InterfacesForUnstructured(version)
	}
}

// InterfacesForUnstructured returns VersionInterfaces suitable for
// dealing with unstructured.Unstructured objects. It will return errors for
// other conversions.
func InterfacesForUnstructured(schema.GroupVersion) (*VersionInterfaces, error) {
	return &VersionInterfaces{
		ObjectConvertor:  &unstructured.UnstructuredObjectConverter{},
		MetadataAccessor: NewAccessor(),
	}, nil
}
