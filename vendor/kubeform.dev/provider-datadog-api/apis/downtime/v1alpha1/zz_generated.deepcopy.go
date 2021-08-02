// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Downtime) DeepCopyInto(out *Downtime) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Downtime.
func (in *Downtime) DeepCopy() *Downtime {
	if in == nil {
		return nil
	}
	out := new(Downtime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Downtime) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DowntimeList) DeepCopyInto(out *DowntimeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Downtime, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DowntimeList.
func (in *DowntimeList) DeepCopy() *DowntimeList {
	if in == nil {
		return nil
	}
	out := new(DowntimeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DowntimeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DowntimeSpec) DeepCopyInto(out *DowntimeSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(DowntimeSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DowntimeSpec.
func (in *DowntimeSpec) DeepCopy() *DowntimeSpec {
	if in == nil {
		return nil
	}
	out := new(DowntimeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DowntimeSpecRecurrence) DeepCopyInto(out *DowntimeSpecRecurrence) {
	*out = *in
	if in.Period != nil {
		in, out := &in.Period, &out.Period
		*out = new(int64)
		**out = **in
	}
	if in.Rrule != nil {
		in, out := &in.Rrule, &out.Rrule
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.UntilDate != nil {
		in, out := &in.UntilDate, &out.UntilDate
		*out = new(int64)
		**out = **in
	}
	if in.UntilOccurrences != nil {
		in, out := &in.UntilOccurrences, &out.UntilOccurrences
		*out = new(int64)
		**out = **in
	}
	if in.WeekDays != nil {
		in, out := &in.WeekDays, &out.WeekDays
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DowntimeSpecRecurrence.
func (in *DowntimeSpecRecurrence) DeepCopy() *DowntimeSpecRecurrence {
	if in == nil {
		return nil
	}
	out := new(DowntimeSpecRecurrence)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DowntimeSpecResource) DeepCopyInto(out *DowntimeSpecResource) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = new(bool)
		**out = **in
	}
	if in.ActiveChildID != nil {
		in, out := &in.ActiveChildID, &out.ActiveChildID
		*out = new(int64)
		**out = **in
	}
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = new(bool)
		**out = **in
	}
	if in.End != nil {
		in, out := &in.End, &out.End
		*out = new(int64)
		**out = **in
	}
	if in.EndDate != nil {
		in, out := &in.EndDate, &out.EndDate
		*out = new(string)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.MonitorID != nil {
		in, out := &in.MonitorID, &out.MonitorID
		*out = new(int64)
		**out = **in
	}
	if in.MonitorTags != nil {
		in, out := &in.MonitorTags, &out.MonitorTags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Recurrence != nil {
		in, out := &in.Recurrence, &out.Recurrence
		*out = new(DowntimeSpecRecurrence)
		(*in).DeepCopyInto(*out)
	}
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Start != nil {
		in, out := &in.Start, &out.Start
		*out = new(int64)
		**out = **in
	}
	if in.StartDate != nil {
		in, out := &in.StartDate, &out.StartDate
		*out = new(string)
		**out = **in
	}
	if in.Timezone != nil {
		in, out := &in.Timezone, &out.Timezone
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DowntimeSpecResource.
func (in *DowntimeSpecResource) DeepCopy() *DowntimeSpecResource {
	if in == nil {
		return nil
	}
	out := new(DowntimeSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DowntimeStatus) DeepCopyInto(out *DowntimeStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DowntimeStatus.
func (in *DowntimeStatus) DeepCopy() *DowntimeStatus {
	if in == nil {
		return nil
	}
	out := new(DowntimeStatus)
	in.DeepCopyInto(out)
	return out
}
