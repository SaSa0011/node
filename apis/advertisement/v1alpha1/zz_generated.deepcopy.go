//go:build !ignore_autogenerated

// Copyright 2022-2023 FLUIDOS Project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"

	nodecorev1alpha1 "github.com/fluidos-project/node/apis/nodecore/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Discovery) DeepCopyInto(out *Discovery) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Discovery.
func (in *Discovery) DeepCopy() *Discovery {
	if in == nil {
		return nil
	}
	out := new(Discovery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Discovery) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryList) DeepCopyInto(out *DiscoveryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Discovery, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryList.
func (in *DiscoveryList) DeepCopy() *DiscoveryList {
	if in == nil {
		return nil
	}
	out := new(DiscoveryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DiscoveryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoverySpec) DeepCopyInto(out *DiscoverySpec) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(nodecorev1alpha1.FlavourSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoverySpec.
func (in *DiscoverySpec) DeepCopy() *DiscoverySpec {
	if in == nil {
		return nil
	}
	out := new(DiscoverySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryStatus) DeepCopyInto(out *DiscoveryStatus) {
	*out = *in
	out.Phase = in.Phase
	out.PeeringCandidate = in.PeeringCandidate
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryStatus.
func (in *DiscoveryStatus) DeepCopy() *DiscoveryStatus {
	if in == nil {
		return nil
	}
	out := new(DiscoveryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeeringCandidate) DeepCopyInto(out *PeeringCandidate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeeringCandidate.
func (in *PeeringCandidate) DeepCopy() *PeeringCandidate {
	if in == nil {
		return nil
	}
	out := new(PeeringCandidate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PeeringCandidate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeeringCandidateList) DeepCopyInto(out *PeeringCandidateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PeeringCandidate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeeringCandidateList.
func (in *PeeringCandidateList) DeepCopy() *PeeringCandidateList {
	if in == nil {
		return nil
	}
	out := new(PeeringCandidateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PeeringCandidateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeeringCandidateSpec) DeepCopyInto(out *PeeringCandidateSpec) {
	*out = *in
	in.Flavour.DeepCopyInto(&out.Flavour)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeeringCandidateSpec.
func (in *PeeringCandidateSpec) DeepCopy() *PeeringCandidateSpec {
	if in == nil {
		return nil
	}
	out := new(PeeringCandidateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeeringCandidateStatus) DeepCopyInto(out *PeeringCandidateStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeeringCandidateStatus.
func (in *PeeringCandidateStatus) DeepCopy() *PeeringCandidateStatus {
	if in == nil {
		return nil
	}
	out := new(PeeringCandidateStatus)
	in.DeepCopyInto(out)
	return out
}
