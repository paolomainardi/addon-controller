//go:build !ignore_autogenerated

/*
Copyright 2024. projectsveltos.io. All rights reserved.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	apiv1beta1 "github.com/projectsveltos/libsveltos/api/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Chart) DeepCopyInto(out *Chart) {
	*out = *in
	if in.LastAppliedTime != nil {
		in, out := &in.LastAppliedTime, &out.LastAppliedTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Chart.
func (in *Chart) DeepCopy() *Chart {
	if in == nil {
		return nil
	}
	out := new(Chart)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfiguration) DeepCopyInto(out *ClusterConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfiguration.
func (in *ClusterConfiguration) DeepCopy() *ClusterConfiguration {
	if in == nil {
		return nil
	}
	out := new(ClusterConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfigurationList) DeepCopyInto(out *ClusterConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfigurationList.
func (in *ClusterConfigurationList) DeepCopy() *ClusterConfigurationList {
	if in == nil {
		return nil
	}
	out := new(ClusterConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfigurationStatus) DeepCopyInto(out *ClusterConfigurationStatus) {
	*out = *in
	if in.ClusterProfileResources != nil {
		in, out := &in.ClusterProfileResources, &out.ClusterProfileResources
		*out = make([]ClusterProfileResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProfileResources != nil {
		in, out := &in.ProfileResources, &out.ProfileResources
		*out = make([]ProfileResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfigurationStatus.
func (in *ClusterConfigurationStatus) DeepCopy() *ClusterConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterProfile) DeepCopyInto(out *ClusterProfile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterProfile.
func (in *ClusterProfile) DeepCopy() *ClusterProfile {
	if in == nil {
		return nil
	}
	out := new(ClusterProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterProfile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterProfileList) DeepCopyInto(out *ClusterProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterProfileList.
func (in *ClusterProfileList) DeepCopy() *ClusterProfileList {
	if in == nil {
		return nil
	}
	out := new(ClusterProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterProfileResource) DeepCopyInto(out *ClusterProfileResource) {
	*out = *in
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = make([]Feature, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterProfileResource.
func (in *ClusterProfileResource) DeepCopy() *ClusterProfileResource {
	if in == nil {
		return nil
	}
	out := new(ClusterProfileResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterReport) DeepCopyInto(out *ClusterReport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterReport.
func (in *ClusterReport) DeepCopy() *ClusterReport {
	if in == nil {
		return nil
	}
	out := new(ClusterReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterReport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterReportList) DeepCopyInto(out *ClusterReportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterReport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterReportList.
func (in *ClusterReportList) DeepCopy() *ClusterReportList {
	if in == nil {
		return nil
	}
	out := new(ClusterReportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterReportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterReportSpec) DeepCopyInto(out *ClusterReportSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterReportSpec.
func (in *ClusterReportSpec) DeepCopy() *ClusterReportSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterReportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterReportStatus) DeepCopyInto(out *ClusterReportStatus) {
	*out = *in
	if in.ReleaseReports != nil {
		in, out := &in.ReleaseReports, &out.ReleaseReports
		*out = make([]ReleaseReport, len(*in))
		copy(*out, *in)
	}
	if in.ResourceReports != nil {
		in, out := &in.ResourceReports, &out.ResourceReports
		*out = make([]ResourceReport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.KustomizeResourceReports != nil {
		in, out := &in.KustomizeResourceReports, &out.KustomizeResourceReports
		*out = make([]ResourceReport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterReportStatus.
func (in *ClusterReportStatus) DeepCopy() *ClusterReportStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterReportStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSummary) DeepCopyInto(out *ClusterSummary) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSummary.
func (in *ClusterSummary) DeepCopy() *ClusterSummary {
	if in == nil {
		return nil
	}
	out := new(ClusterSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterSummary) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSummaryList) DeepCopyInto(out *ClusterSummaryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterSummary, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSummaryList.
func (in *ClusterSummaryList) DeepCopy() *ClusterSummaryList {
	if in == nil {
		return nil
	}
	out := new(ClusterSummaryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterSummaryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSummarySpec) DeepCopyInto(out *ClusterSummarySpec) {
	*out = *in
	in.ClusterProfileSpec.DeepCopyInto(&out.ClusterProfileSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSummarySpec.
func (in *ClusterSummarySpec) DeepCopy() *ClusterSummarySpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSummarySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSummaryStatus) DeepCopyInto(out *ClusterSummaryStatus) {
	*out = *in
	if in.Dependencies != nil {
		in, out := &in.Dependencies, &out.Dependencies
		*out = new(string)
		**out = **in
	}
	if in.FeatureSummaries != nil {
		in, out := &in.FeatureSummaries, &out.FeatureSummaries
		*out = make([]FeatureSummary, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DeployedGVKs != nil {
		in, out := &in.DeployedGVKs, &out.DeployedGVKs
		*out = make([]FeatureDeploymentInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HelmReleaseSummaries != nil {
		in, out := &in.HelmReleaseSummaries, &out.HelmReleaseSummaries
		*out = make([]HelmChartSummary, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSummaryStatus.
func (in *ClusterSummaryStatus) DeepCopy() *ClusterSummaryStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterSummaryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Clusters) DeepCopyInto(out *Clusters) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]corev1.ObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Clusters.
func (in *Clusters) DeepCopy() *Clusters {
	if in == nil {
		return nil
	}
	out := new(Clusters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DriftExclusion) DeepCopyInto(out *DriftExclusion) {
	*out = *in
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(apiv1beta1.PatchSelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DriftExclusion.
func (in *DriftExclusion) DeepCopy() *DriftExclusion {
	if in == nil {
		return nil
	}
	out := new(DriftExclusion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DryRunReconciliationError) DeepCopyInto(out *DryRunReconciliationError) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DryRunReconciliationError.
func (in *DryRunReconciliationError) DeepCopy() *DryRunReconciliationError {
	if in == nil {
		return nil
	}
	out := new(DryRunReconciliationError)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Feature) DeepCopyInto(out *Feature) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]Resource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Charts != nil {
		in, out := &in.Charts, &out.Charts
		*out = make([]Chart, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Feature.
func (in *Feature) DeepCopy() *Feature {
	if in == nil {
		return nil
	}
	out := new(Feature)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureDeploymentInfo) DeepCopyInto(out *FeatureDeploymentInfo) {
	*out = *in
	if in.DeployedGroupVersionKind != nil {
		in, out := &in.DeployedGroupVersionKind, &out.DeployedGroupVersionKind
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureDeploymentInfo.
func (in *FeatureDeploymentInfo) DeepCopy() *FeatureDeploymentInfo {
	if in == nil {
		return nil
	}
	out := new(FeatureDeploymentInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureSummary) DeepCopyInto(out *FeatureSummary) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.FailureReason != nil {
		in, out := &in.FailureReason, &out.FailureReason
		*out = new(string)
		**out = **in
	}
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
	if in.DeployedGroupVersionKind != nil {
		in, out := &in.DeployedGroupVersionKind, &out.DeployedGroupVersionKind
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LastAppliedTime != nil {
		in, out := &in.LastAppliedTime, &out.LastAppliedTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureSummary.
func (in *FeatureSummary) DeepCopy() *FeatureSummary {
	if in == nil {
		return nil
	}
	out := new(FeatureSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmChart) DeepCopyInto(out *HelmChart) {
	*out = *in
	if in.ValuesFrom != nil {
		in, out := &in.ValuesFrom, &out.ValuesFrom
		*out = make([]ValueFrom, len(*in))
		copy(*out, *in)
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = new(HelmOptions)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmChart.
func (in *HelmChart) DeepCopy() *HelmChart {
	if in == nil {
		return nil
	}
	out := new(HelmChart)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmChartSummary) DeepCopyInto(out *HelmChartSummary) {
	*out = *in
	if in.ValuesHash != nil {
		in, out := &in.ValuesHash, &out.ValuesHash
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmChartSummary.
func (in *HelmChartSummary) DeepCopy() *HelmChartSummary {
	if in == nil {
		return nil
	}
	out := new(HelmChartSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmInstallOptions) DeepCopyInto(out *HelmInstallOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmInstallOptions.
func (in *HelmInstallOptions) DeepCopy() *HelmInstallOptions {
	if in == nil {
		return nil
	}
	out := new(HelmInstallOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmOptions) DeepCopyInto(out *HelmOptions) {
	*out = *in
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.InstallOptions = in.InstallOptions
	out.UpgradeOptions = in.UpgradeOptions
	out.UninstallOptions = in.UninstallOptions
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmOptions.
func (in *HelmOptions) DeepCopy() *HelmOptions {
	if in == nil {
		return nil
	}
	out := new(HelmOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmUninstallOptions) DeepCopyInto(out *HelmUninstallOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmUninstallOptions.
func (in *HelmUninstallOptions) DeepCopy() *HelmUninstallOptions {
	if in == nil {
		return nil
	}
	out := new(HelmUninstallOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmUpgradeOptions) DeepCopyInto(out *HelmUpgradeOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmUpgradeOptions.
func (in *HelmUpgradeOptions) DeepCopy() *HelmUpgradeOptions {
	if in == nil {
		return nil
	}
	out := new(HelmUpgradeOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KustomizationRef) DeepCopyInto(out *KustomizationRef) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ValuesFrom != nil {
		in, out := &in.ValuesFrom, &out.ValuesFrom
		*out = make([]ValueFrom, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KustomizationRef.
func (in *KustomizationRef) DeepCopy() *KustomizationRef {
	if in == nil {
		return nil
	}
	out := new(KustomizationRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyRef) DeepCopyInto(out *PolicyRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyRef.
func (in *PolicyRef) DeepCopy() *PolicyRef {
	if in == nil {
		return nil
	}
	out := new(PolicyRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Profile) DeepCopyInto(out *Profile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Profile.
func (in *Profile) DeepCopy() *Profile {
	if in == nil {
		return nil
	}
	out := new(Profile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Profile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProfileList) DeepCopyInto(out *ProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Profile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProfileList.
func (in *ProfileList) DeepCopy() *ProfileList {
	if in == nil {
		return nil
	}
	out := new(ProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProfileResource) DeepCopyInto(out *ProfileResource) {
	*out = *in
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = make([]Feature, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProfileResource.
func (in *ProfileResource) DeepCopy() *ProfileResource {
	if in == nil {
		return nil
	}
	out := new(ProfileResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseReport) DeepCopyInto(out *ReleaseReport) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseReport.
func (in *ReleaseReport) DeepCopy() *ReleaseReport {
	if in == nil {
		return nil
	}
	out := new(ReleaseReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resource) DeepCopyInto(out *Resource) {
	*out = *in
	if in.LastAppliedTime != nil {
		in, out := &in.LastAppliedTime, &out.LastAppliedTime
		*out = (*in).DeepCopy()
	}
	out.Owner = in.Owner
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resource.
func (in *Resource) DeepCopy() *Resource {
	if in == nil {
		return nil
	}
	out := new(Resource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceReport) DeepCopyInto(out *ResourceReport) {
	*out = *in
	in.Resource.DeepCopyInto(&out.Resource)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceReport.
func (in *ResourceReport) DeepCopy() *ResourceReport {
	if in == nil {
		return nil
	}
	out := new(ResourceReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Spec) DeepCopyInto(out *Spec) {
	*out = *in
	in.ClusterSelector.DeepCopyInto(&out.ClusterSelector)
	if in.ClusterRefs != nil {
		in, out := &in.ClusterRefs, &out.ClusterRefs
		*out = make([]corev1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.SetRefs != nil {
		in, out := &in.SetRefs, &out.SetRefs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MaxUpdate != nil {
		in, out := &in.MaxUpdate, &out.MaxUpdate
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.TemplateResourceRefs != nil {
		in, out := &in.TemplateResourceRefs, &out.TemplateResourceRefs
		*out = make([]TemplateResourceRef, len(*in))
		copy(*out, *in)
	}
	if in.DependsOn != nil {
		in, out := &in.DependsOn, &out.DependsOn
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PolicyRefs != nil {
		in, out := &in.PolicyRefs, &out.PolicyRefs
		*out = make([]PolicyRef, len(*in))
		copy(*out, *in)
	}
	if in.HelmCharts != nil {
		in, out := &in.HelmCharts, &out.HelmCharts
		*out = make([]HelmChart, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.KustomizationRefs != nil {
		in, out := &in.KustomizationRefs, &out.KustomizationRefs
		*out = make([]KustomizationRef, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ValidateHealths != nil {
		in, out := &in.ValidateHealths, &out.ValidateHealths
		*out = make([]ValidateHealth, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Patches != nil {
		in, out := &in.Patches, &out.Patches
		*out = make([]apiv1beta1.Patch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DriftExclusions != nil {
		in, out := &in.DriftExclusions, &out.DriftExclusions
		*out = make([]DriftExclusion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExtraLabels != nil {
		in, out := &in.ExtraLabels, &out.ExtraLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ExtraAnnotations != nil {
		in, out := &in.ExtraAnnotations, &out.ExtraAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Spec.
func (in *Spec) DeepCopy() *Spec {
	if in == nil {
		return nil
	}
	out := new(Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Status) DeepCopyInto(out *Status) {
	*out = *in
	if in.MatchingClusterRefs != nil {
		in, out := &in.MatchingClusterRefs, &out.MatchingClusterRefs
		*out = make([]corev1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	in.UpdatingClusters.DeepCopyInto(&out.UpdatingClusters)
	in.UpdatedClusters.DeepCopyInto(&out.UpdatedClusters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Status.
func (in *Status) DeepCopy() *Status {
	if in == nil {
		return nil
	}
	out := new(Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateResourceRef) DeepCopyInto(out *TemplateResourceRef) {
	*out = *in
	out.Resource = in.Resource
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateResourceRef.
func (in *TemplateResourceRef) DeepCopy() *TemplateResourceRef {
	if in == nil {
		return nil
	}
	out := new(TemplateResourceRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidateHealth) DeepCopyInto(out *ValidateHealth) {
	*out = *in
	if in.LabelFilters != nil {
		in, out := &in.LabelFilters, &out.LabelFilters
		*out = make([]apiv1beta1.LabelFilter, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidateHealth.
func (in *ValidateHealth) DeepCopy() *ValidateHealth {
	if in == nil {
		return nil
	}
	out := new(ValidateHealth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValueFrom) DeepCopyInto(out *ValueFrom) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValueFrom.
func (in *ValueFrom) DeepCopy() *ValueFrom {
	if in == nil {
		return nil
	}
	out := new(ValueFrom)
	in.DeepCopyInto(out)
	return out
}
