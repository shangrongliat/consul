// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package multiclusterv2

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using ComputedExportedServices within kubernetes types, where deepcopy-gen is used.
func (in *ComputedExportedServices) DeepCopyInto(out *ComputedExportedServices) {
	proto.Reset(out)
	proto.Merge(out, proto.Clone(in))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputedExportedServices. Required by controller-gen.
func (in *ComputedExportedServices) DeepCopy() *ComputedExportedServices {
	if in == nil {
		return nil
	}
	out := new(ComputedExportedServices)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ComputedExportedServices. Required by controller-gen.
func (in *ComputedExportedServices) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ComputedExportedService within kubernetes types, where deepcopy-gen is used.
func (in *ComputedExportedService) DeepCopyInto(out *ComputedExportedService) {
	proto.Reset(out)
	proto.Merge(out, proto.Clone(in))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputedExportedService. Required by controller-gen.
func (in *ComputedExportedService) DeepCopy() *ComputedExportedService {
	if in == nil {
		return nil
	}
	out := new(ComputedExportedService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ComputedExportedService. Required by controller-gen.
func (in *ComputedExportedService) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ComputedExportedServiceConsumer within kubernetes types, where deepcopy-gen is used.
func (in *ComputedExportedServiceConsumer) DeepCopyInto(out *ComputedExportedServiceConsumer) {
	proto.Reset(out)
	proto.Merge(out, proto.Clone(in))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputedExportedServiceConsumer. Required by controller-gen.
func (in *ComputedExportedServiceConsumer) DeepCopy() *ComputedExportedServiceConsumer {
	if in == nil {
		return nil
	}
	out := new(ComputedExportedServiceConsumer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ComputedExportedServiceConsumer. Required by controller-gen.
func (in *ComputedExportedServiceConsumer) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
