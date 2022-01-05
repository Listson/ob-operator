//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) 2021 OceanBase
ob-operator is licensed under Mulan PSL v2.
You can use this software according to the terms and conditions of the Mulan PSL v2.
You may obtain a copy of Mulan PSL v2 at:
         http://license.coscl.org.cn/MulanPSL2
THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
See the Mulan PSL v2 for more details.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = make([]Subset, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterOBZoneStatus) DeepCopyInto(out *ClusterOBZoneStatus) {
	*out = *in
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = make([]OBZoneInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterOBZoneStatus.
func (in *ClusterOBZoneStatus) DeepCopy() *ClusterOBZoneStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterOBZoneStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRootServiceStatus) DeepCopyInto(out *ClusterRootServiceStatus) {
	*out = *in
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = make([]ZoneRootServiceStatus, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRootServiceStatus.
func (in *ClusterRootServiceStatus) DeepCopy() *ClusterRootServiceStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterRootServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = make([]ZoneStatus, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBCluster) DeepCopyInto(out *OBCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBCluster.
func (in *OBCluster) DeepCopy() *OBCluster {
	if in == nil {
		return nil
	}
	out := new(OBCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OBCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBClusterList) DeepCopyInto(out *OBClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OBCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBClusterList.
func (in *OBClusterList) DeepCopy() *OBClusterList {
	if in == nil {
		return nil
	}
	out := new(OBClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OBClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBClusterSpec) DeepCopyInto(out *OBClusterSpec) {
	*out = *in
	if in.Topology != nil {
		in, out := &in.Topology, &out.Topology
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBClusterSpec.
func (in *OBClusterSpec) DeepCopy() *OBClusterSpec {
	if in == nil {
		return nil
	}
	out := new(OBClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBClusterStatus) DeepCopyInto(out *OBClusterStatus) {
	*out = *in
	if in.Topology != nil {
		in, out := &in.Topology, &out.Topology
		*out = make([]ClusterStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBClusterStatus.
func (in *OBClusterStatus) DeepCopy() *OBClusterStatus {
	if in == nil {
		return nil
	}
	out := new(OBClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBNode) DeepCopyInto(out *OBNode) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBNode.
func (in *OBNode) DeepCopy() *OBNode {
	if in == nil {
		return nil
	}
	out := new(OBNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBZone) DeepCopyInto(out *OBZone) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBZone.
func (in *OBZone) DeepCopy() *OBZone {
	if in == nil {
		return nil
	}
	out := new(OBZone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OBZone) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBZoneInfo) DeepCopyInto(out *OBZoneInfo) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]OBNode, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBZoneInfo.
func (in *OBZoneInfo) DeepCopy() *OBZoneInfo {
	if in == nil {
		return nil
	}
	out := new(OBZoneInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBZoneList) DeepCopyInto(out *OBZoneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OBZone, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBZoneList.
func (in *OBZoneList) DeepCopy() *OBZoneList {
	if in == nil {
		return nil
	}
	out := new(OBZoneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OBZoneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBZoneSpec) DeepCopyInto(out *OBZoneSpec) {
	*out = *in
	if in.Topology != nil {
		in, out := &in.Topology, &out.Topology
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBZoneSpec.
func (in *OBZoneSpec) DeepCopy() *OBZoneSpec {
	if in == nil {
		return nil
	}
	out := new(OBZoneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBZoneStatus) DeepCopyInto(out *OBZoneStatus) {
	*out = *in
	if in.Topology != nil {
		in, out := &in.Topology, &out.Topology
		*out = make([]ClusterOBZoneStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBZoneStatus.
func (in *OBZoneStatus) DeepCopy() *OBZoneStatus {
	if in == nil {
		return nil
	}
	out := new(OBZoneStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PVCStatus) DeepCopyInto(out *PVCStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PVCStatus.
func (in *PVCStatus) DeepCopy() *PVCStatus {
	if in == nil {
		return nil
	}
	out := new(PVCStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodStatus) DeepCopyInto(out *PodStatus) {
	*out = *in
	if in.PVCs != nil {
		in, out := &in.PVCs, &out.PVCs
		*out = make([]PVCStatus, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodStatus.
func (in *PodStatus) DeepCopy() *PodStatus {
	if in == nil {
		return nil
	}
	out := new(PodStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcesSpec) DeepCopyInto(out *ResourcesSpec) {
	*out = *in
	out.CPU = in.CPU.DeepCopy()
	out.Memory = in.Memory.DeepCopy()
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = make([]StorageSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcesSpec.
func (in *ResourcesSpec) DeepCopy() *ResourcesSpec {
	if in == nil {
		return nil
	}
	out := new(ResourcesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RootService) DeepCopyInto(out *RootService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RootService.
func (in *RootService) DeepCopy() *RootService {
	if in == nil {
		return nil
	}
	out := new(RootService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RootService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RootServiceList) DeepCopyInto(out *RootServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RootService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RootServiceList.
func (in *RootServiceList) DeepCopy() *RootServiceList {
	if in == nil {
		return nil
	}
	out := new(RootServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RootServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RootServiceSpec) DeepCopyInto(out *RootServiceSpec) {
	*out = *in
	if in.Topology != nil {
		in, out := &in.Topology, &out.Topology
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RootServiceSpec.
func (in *RootServiceSpec) DeepCopy() *RootServiceSpec {
	if in == nil {
		return nil
	}
	out := new(RootServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RootServiceStatus) DeepCopyInto(out *RootServiceStatus) {
	*out = *in
	if in.Topology != nil {
		in, out := &in.Topology, &out.Topology
		*out = make([]ClusterRootServiceStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RootServiceStatus.
func (in *RootServiceStatus) DeepCopy() *RootServiceStatus {
	if in == nil {
		return nil
	}
	out := new(RootServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulApp) DeepCopyInto(out *StatefulApp) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulApp.
func (in *StatefulApp) DeepCopy() *StatefulApp {
	if in == nil {
		return nil
	}
	out := new(StatefulApp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StatefulApp) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulAppList) DeepCopyInto(out *StatefulAppList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StatefulApp, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulAppList.
func (in *StatefulAppList) DeepCopy() *StatefulAppList {
	if in == nil {
		return nil
	}
	out := new(StatefulAppList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StatefulAppList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulAppSpec) DeepCopyInto(out *StatefulAppSpec) {
	*out = *in
	if in.Subsets != nil {
		in, out := &in.Subsets, &out.Subsets
		*out = make([]Subset, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.PodTemplate.DeepCopyInto(&out.PodTemplate)
	if in.StorageTemplates != nil {
		in, out := &in.StorageTemplates, &out.StorageTemplates
		*out = make([]StorageTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulAppSpec.
func (in *StatefulAppSpec) DeepCopy() *StatefulAppSpec {
	if in == nil {
		return nil
	}
	out := new(StatefulAppSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulAppStatus) DeepCopyInto(out *StatefulAppStatus) {
	*out = *in
	if in.Subsets != nil {
		in, out := &in.Subsets, &out.Subsets
		*out = make([]SubsetStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulAppStatus.
func (in *StatefulAppStatus) DeepCopy() *StatefulAppStatus {
	if in == nil {
		return nil
	}
	out := new(StatefulAppStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageSpec) DeepCopyInto(out *StorageSpec) {
	*out = *in
	out.Size = in.Size.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageSpec.
func (in *StorageSpec) DeepCopy() *StorageSpec {
	if in == nil {
		return nil
	}
	out := new(StorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageTemplate) DeepCopyInto(out *StorageTemplate) {
	*out = *in
	in.PVC.DeepCopyInto(&out.PVC)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageTemplate.
func (in *StorageTemplate) DeepCopy() *StorageTemplate {
	if in == nil {
		return nil
	}
	out := new(StorageTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subset) DeepCopyInto(out *Subset) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subset.
func (in *Subset) DeepCopy() *Subset {
	if in == nil {
		return nil
	}
	out := new(Subset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubsetStatus) DeepCopyInto(out *SubsetStatus) {
	*out = *in
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]PodStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubsetStatus.
func (in *SubsetStatus) DeepCopy() *SubsetStatus {
	if in == nil {
		return nil
	}
	out := new(SubsetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneRootServiceStatus) DeepCopyInto(out *ZoneRootServiceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneRootServiceStatus.
func (in *ZoneRootServiceStatus) DeepCopy() *ZoneRootServiceStatus {
	if in == nil {
		return nil
	}
	out := new(ZoneRootServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneStatus) DeepCopyInto(out *ZoneStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneStatus.
func (in *ZoneStatus) DeepCopy() *ZoneStatus {
	if in == nil {
		return nil
	}
	out := new(ZoneStatus)
	in.DeepCopyInto(out)
	return out
}