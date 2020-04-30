// +build !ignore_autogenerated

/*

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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceProperty) DeepCopyInto(out *DeviceProperty) {
	*out = *in
	out.Visitor = in.Visitor
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceProperty.
func (in *DeviceProperty) DeepCopy() *DeviceProperty {
	if in == nil {
		return nil
	}
	out := new(DeviceProperty)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OPCUADevice) DeepCopyInto(out *OPCUADevice) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OPCUADevice.
func (in *OPCUADevice) DeepCopy() *OPCUADevice {
	if in == nil {
		return nil
	}
	out := new(OPCUADevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OPCUADevice) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OPCUADeviceList) DeepCopyInto(out *OPCUADeviceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OPCUADevice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OPCUADeviceList.
func (in *OPCUADeviceList) DeepCopy() *OPCUADeviceList {
	if in == nil {
		return nil
	}
	out := new(OPCUADeviceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OPCUADeviceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OPCUADeviceSpec) DeepCopyInto(out *OPCUADeviceSpec) {
	*out = *in
	if in.ProtocolConfig != nil {
		in, out := &in.ProtocolConfig, &out.ProtocolConfig
		*out = new(OPCUAProtocolConfig)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make([]DeviceProperty, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OPCUADeviceSpec.
func (in *OPCUADeviceSpec) DeepCopy() *OPCUADeviceSpec {
	if in == nil {
		return nil
	}
	out := new(OPCUADeviceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OPCUADeviceStatus) DeepCopyInto(out *OPCUADeviceStatus) {
	*out = *in
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make([]StatusProperties, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OPCUADeviceStatus.
func (in *OPCUADeviceStatus) DeepCopy() *OPCUADeviceStatus {
	if in == nil {
		return nil
	}
	out := new(OPCUADeviceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OPCUAProtocolConfig) DeepCopyInto(out *OPCUAProtocolConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OPCUAProtocolConfig.
func (in *OPCUAProtocolConfig) DeepCopy() *OPCUAProtocolConfig {
	if in == nil {
		return nil
	}
	out := new(OPCUAProtocolConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PropertyVisitor) DeepCopyInto(out *PropertyVisitor) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PropertyVisitor.
func (in *PropertyVisitor) DeepCopy() *PropertyVisitor {
	if in == nil {
		return nil
	}
	out := new(PropertyVisitor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusProperties) DeepCopyInto(out *StatusProperties) {
	*out = *in
	in.UpdatedAt.DeepCopyInto(&out.UpdatedAt)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusProperties.
func (in *StatusProperties) DeepCopy() *StatusProperties {
	if in == nil {
		return nil
	}
	out := new(StatusProperties)
	in.DeepCopyInto(out)
	return out
}