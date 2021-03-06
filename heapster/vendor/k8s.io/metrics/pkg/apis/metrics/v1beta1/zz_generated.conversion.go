// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1beta1

import (
	v1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	metrics "k8s.io/metrics/pkg/apis/metrics"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1beta1_ContainerMetrics_To_metrics_ContainerMetrics,
		Convert_metrics_ContainerMetrics_To_v1beta1_ContainerMetrics,
		Convert_v1beta1_NodeMetrics_To_metrics_NodeMetrics,
		Convert_metrics_NodeMetrics_To_v1beta1_NodeMetrics,
		Convert_v1beta1_NodeMetricsList_To_metrics_NodeMetricsList,
		Convert_metrics_NodeMetricsList_To_v1beta1_NodeMetricsList,
		Convert_v1beta1_PodMetrics_To_metrics_PodMetrics,
		Convert_metrics_PodMetrics_To_v1beta1_PodMetrics,
		Convert_v1beta1_PodMetricsList_To_metrics_PodMetricsList,
		Convert_metrics_PodMetricsList_To_v1beta1_PodMetricsList,
	)
}

func autoConvert_v1beta1_ContainerMetrics_To_metrics_ContainerMetrics(in *ContainerMetrics, out *metrics.ContainerMetrics, s conversion.Scope) error {
	out.Name = in.Name
	out.Usage = *(*metrics.ResourceList)(unsafe.Pointer(&in.Usage))
	return nil
}

// Convert_v1beta1_ContainerMetrics_To_metrics_ContainerMetrics is an autogenerated conversion function.
func Convert_v1beta1_ContainerMetrics_To_metrics_ContainerMetrics(in *ContainerMetrics, out *metrics.ContainerMetrics, s conversion.Scope) error {
	return autoConvert_v1beta1_ContainerMetrics_To_metrics_ContainerMetrics(in, out, s)
}

func autoConvert_metrics_ContainerMetrics_To_v1beta1_ContainerMetrics(in *metrics.ContainerMetrics, out *ContainerMetrics, s conversion.Scope) error {
	out.Name = in.Name
	out.Usage = *(*v1.ResourceList)(unsafe.Pointer(&in.Usage))
	return nil
}

// Convert_metrics_ContainerMetrics_To_v1beta1_ContainerMetrics is an autogenerated conversion function.
func Convert_metrics_ContainerMetrics_To_v1beta1_ContainerMetrics(in *metrics.ContainerMetrics, out *ContainerMetrics, s conversion.Scope) error {
	return autoConvert_metrics_ContainerMetrics_To_v1beta1_ContainerMetrics(in, out, s)
}

func autoConvert_v1beta1_NodeMetrics_To_metrics_NodeMetrics(in *NodeMetrics, out *metrics.NodeMetrics, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Timestamp = in.Timestamp
	out.Window = in.Window
	out.Usage = *(*metrics.ResourceList)(unsafe.Pointer(&in.Usage))
	return nil
}

// Convert_v1beta1_NodeMetrics_To_metrics_NodeMetrics is an autogenerated conversion function.
func Convert_v1beta1_NodeMetrics_To_metrics_NodeMetrics(in *NodeMetrics, out *metrics.NodeMetrics, s conversion.Scope) error {
	return autoConvert_v1beta1_NodeMetrics_To_metrics_NodeMetrics(in, out, s)
}

func autoConvert_metrics_NodeMetrics_To_v1beta1_NodeMetrics(in *metrics.NodeMetrics, out *NodeMetrics, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Timestamp = in.Timestamp
	out.Window = in.Window
	out.Usage = *(*v1.ResourceList)(unsafe.Pointer(&in.Usage))
	return nil
}

// Convert_metrics_NodeMetrics_To_v1beta1_NodeMetrics is an autogenerated conversion function.
func Convert_metrics_NodeMetrics_To_v1beta1_NodeMetrics(in *metrics.NodeMetrics, out *NodeMetrics, s conversion.Scope) error {
	return autoConvert_metrics_NodeMetrics_To_v1beta1_NodeMetrics(in, out, s)
}

func autoConvert_v1beta1_NodeMetricsList_To_metrics_NodeMetricsList(in *NodeMetricsList, out *metrics.NodeMetricsList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]metrics.NodeMetrics)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_NodeMetricsList_To_metrics_NodeMetricsList is an autogenerated conversion function.
func Convert_v1beta1_NodeMetricsList_To_metrics_NodeMetricsList(in *NodeMetricsList, out *metrics.NodeMetricsList, s conversion.Scope) error {
	return autoConvert_v1beta1_NodeMetricsList_To_metrics_NodeMetricsList(in, out, s)
}

func autoConvert_metrics_NodeMetricsList_To_v1beta1_NodeMetricsList(in *metrics.NodeMetricsList, out *NodeMetricsList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]NodeMetrics)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_metrics_NodeMetricsList_To_v1beta1_NodeMetricsList is an autogenerated conversion function.
func Convert_metrics_NodeMetricsList_To_v1beta1_NodeMetricsList(in *metrics.NodeMetricsList, out *NodeMetricsList, s conversion.Scope) error {
	return autoConvert_metrics_NodeMetricsList_To_v1beta1_NodeMetricsList(in, out, s)
}

func autoConvert_v1beta1_PodMetrics_To_metrics_PodMetrics(in *PodMetrics, out *metrics.PodMetrics, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Timestamp = in.Timestamp
	out.Window = in.Window
	out.Containers = *(*[]metrics.ContainerMetrics)(unsafe.Pointer(&in.Containers))
	return nil
}

// Convert_v1beta1_PodMetrics_To_metrics_PodMetrics is an autogenerated conversion function.
func Convert_v1beta1_PodMetrics_To_metrics_PodMetrics(in *PodMetrics, out *metrics.PodMetrics, s conversion.Scope) error {
	return autoConvert_v1beta1_PodMetrics_To_metrics_PodMetrics(in, out, s)
}

func autoConvert_metrics_PodMetrics_To_v1beta1_PodMetrics(in *metrics.PodMetrics, out *PodMetrics, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Timestamp = in.Timestamp
	out.Window = in.Window
	out.Containers = *(*[]ContainerMetrics)(unsafe.Pointer(&in.Containers))
	return nil
}

// Convert_metrics_PodMetrics_To_v1beta1_PodMetrics is an autogenerated conversion function.
func Convert_metrics_PodMetrics_To_v1beta1_PodMetrics(in *metrics.PodMetrics, out *PodMetrics, s conversion.Scope) error {
	return autoConvert_metrics_PodMetrics_To_v1beta1_PodMetrics(in, out, s)
}

func autoConvert_v1beta1_PodMetricsList_To_metrics_PodMetricsList(in *PodMetricsList, out *metrics.PodMetricsList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]metrics.PodMetrics)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_PodMetricsList_To_metrics_PodMetricsList is an autogenerated conversion function.
func Convert_v1beta1_PodMetricsList_To_metrics_PodMetricsList(in *PodMetricsList, out *metrics.PodMetricsList, s conversion.Scope) error {
	return autoConvert_v1beta1_PodMetricsList_To_metrics_PodMetricsList(in, out, s)
}

func autoConvert_metrics_PodMetricsList_To_v1beta1_PodMetricsList(in *metrics.PodMetricsList, out *PodMetricsList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]PodMetrics)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_metrics_PodMetricsList_To_v1beta1_PodMetricsList is an autogenerated conversion function.
func Convert_metrics_PodMetricsList_To_v1beta1_PodMetricsList(in *metrics.PodMetricsList, out *PodMetricsList, s conversion.Scope) error {
	return autoConvert_metrics_PodMetricsList_To_v1beta1_PodMetricsList(in, out, s)
}
