// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/neo-hu/docker_el6/containerd/linux/runctypes/runc.proto

/*
	Package runctypes is a generated protocol buffer package.

	It is generated from these files:
		github.com/neo-hu/docker_el6/containerd/linux/runctypes/runc.proto

	It has these top-level messages:
		RuncOptions
		CreateOptions
		CheckpointOptions
		ProcessDetails
*/
package runctypes

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// skipping weak import gogoproto "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type RuncOptions struct {
	Runtime       string `protobuf:"bytes,1,opt,name=runtime,proto3" json:"runtime,omitempty"`
	RuntimeRoot   string `protobuf:"bytes,2,opt,name=runtime_root,json=runtimeRoot,proto3" json:"runtime_root,omitempty"`
	CriuPath      string `protobuf:"bytes,3,opt,name=criu_path,json=criuPath,proto3" json:"criu_path,omitempty"`
	SystemdCgroup bool   `protobuf:"varint,4,opt,name=systemd_cgroup,json=systemdCgroup,proto3" json:"systemd_cgroup,omitempty"`
}

func (m *RuncOptions) Reset()                    { *m = RuncOptions{} }
func (*RuncOptions) ProtoMessage()               {}
func (*RuncOptions) Descriptor() ([]byte, []int) { return fileDescriptorRunc, []int{0} }

type CreateOptions struct {
	NoPivotRoot         bool     `protobuf:"varint,1,opt,name=no_pivot_root,json=noPivotRoot,proto3" json:"no_pivot_root,omitempty"`
	OpenTcp             bool     `protobuf:"varint,2,opt,name=open_tcp,json=openTcp,proto3" json:"open_tcp,omitempty"`
	ExternalUnixSockets bool     `protobuf:"varint,3,opt,name=external_unix_sockets,json=externalUnixSockets,proto3" json:"external_unix_sockets,omitempty"`
	Terminal            bool     `protobuf:"varint,4,opt,name=terminal,proto3" json:"terminal,omitempty"`
	FileLocks           bool     `protobuf:"varint,5,opt,name=file_locks,json=fileLocks,proto3" json:"file_locks,omitempty"`
	EmptyNamespaces     []string `protobuf:"bytes,6,rep,name=empty_namespaces,json=emptyNamespaces" json:"empty_namespaces,omitempty"`
	CgroupsMode         string   `protobuf:"bytes,7,opt,name=cgroups_mode,json=cgroupsMode,proto3" json:"cgroups_mode,omitempty"`
	NoNewKeyring        bool     `protobuf:"varint,8,opt,name=no_new_keyring,json=noNewKeyring,proto3" json:"no_new_keyring,omitempty"`
	ShimCgroup          string   `protobuf:"bytes,9,opt,name=shim_cgroup,json=shimCgroup,proto3" json:"shim_cgroup,omitempty"`
	IoUid               uint32   `protobuf:"varint,10,opt,name=io_uid,json=ioUid,proto3" json:"io_uid,omitempty"`
	IoGid               uint32   `protobuf:"varint,11,opt,name=io_gid,json=ioGid,proto3" json:"io_gid,omitempty"`
}

func (m *CreateOptions) Reset()                    { *m = CreateOptions{} }
func (*CreateOptions) ProtoMessage()               {}
func (*CreateOptions) Descriptor() ([]byte, []int) { return fileDescriptorRunc, []int{1} }

type CheckpointOptions struct {
	Exit                bool     `protobuf:"varint,1,opt,name=exit,proto3" json:"exit,omitempty"`
	OpenTcp             bool     `protobuf:"varint,2,opt,name=open_tcp,json=openTcp,proto3" json:"open_tcp,omitempty"`
	ExternalUnixSockets bool     `protobuf:"varint,3,opt,name=external_unix_sockets,json=externalUnixSockets,proto3" json:"external_unix_sockets,omitempty"`
	Terminal            bool     `protobuf:"varint,4,opt,name=terminal,proto3" json:"terminal,omitempty"`
	FileLocks           bool     `protobuf:"varint,5,opt,name=file_locks,json=fileLocks,proto3" json:"file_locks,omitempty"`
	EmptyNamespaces     []string `protobuf:"bytes,6,rep,name=empty_namespaces,json=emptyNamespaces" json:"empty_namespaces,omitempty"`
	CgroupsMode         string   `protobuf:"bytes,7,opt,name=cgroups_mode,json=cgroupsMode,proto3" json:"cgroups_mode,omitempty"`
}

func (m *CheckpointOptions) Reset()                    { *m = CheckpointOptions{} }
func (*CheckpointOptions) ProtoMessage()               {}
func (*CheckpointOptions) Descriptor() ([]byte, []int) { return fileDescriptorRunc, []int{2} }

type ProcessDetails struct {
	ExecID string `protobuf:"bytes,1,opt,name=exec_id,json=execId,proto3" json:"exec_id,omitempty"`
}

func (m *ProcessDetails) Reset()                    { *m = ProcessDetails{} }
func (*ProcessDetails) ProtoMessage()               {}
func (*ProcessDetails) Descriptor() ([]byte, []int) { return fileDescriptorRunc, []int{3} }

func init() {
	proto.RegisterType((*RuncOptions)(nil), "containerd.linux.runc.RuncOptions")
	proto.RegisterType((*CreateOptions)(nil), "containerd.linux.runc.CreateOptions")
	proto.RegisterType((*CheckpointOptions)(nil), "containerd.linux.runc.CheckpointOptions")
	proto.RegisterType((*ProcessDetails)(nil), "containerd.linux.runc.ProcessDetails")
}
func (m *RuncOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RuncOptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Runtime) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRunc(dAtA, i, uint64(len(m.Runtime)))
		i += copy(dAtA[i:], m.Runtime)
	}
	if len(m.RuntimeRoot) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRunc(dAtA, i, uint64(len(m.RuntimeRoot)))
		i += copy(dAtA[i:], m.RuntimeRoot)
	}
	if len(m.CriuPath) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRunc(dAtA, i, uint64(len(m.CriuPath)))
		i += copy(dAtA[i:], m.CriuPath)
	}
	if m.SystemdCgroup {
		dAtA[i] = 0x20
		i++
		if m.SystemdCgroup {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *CreateOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateOptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.NoPivotRoot {
		dAtA[i] = 0x8
		i++
		if m.NoPivotRoot {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.OpenTcp {
		dAtA[i] = 0x10
		i++
		if m.OpenTcp {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.ExternalUnixSockets {
		dAtA[i] = 0x18
		i++
		if m.ExternalUnixSockets {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Terminal {
		dAtA[i] = 0x20
		i++
		if m.Terminal {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.FileLocks {
		dAtA[i] = 0x28
		i++
		if m.FileLocks {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.EmptyNamespaces) > 0 {
		for _, s := range m.EmptyNamespaces {
			dAtA[i] = 0x32
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.CgroupsMode) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintRunc(dAtA, i, uint64(len(m.CgroupsMode)))
		i += copy(dAtA[i:], m.CgroupsMode)
	}
	if m.NoNewKeyring {
		dAtA[i] = 0x40
		i++
		if m.NoNewKeyring {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.ShimCgroup) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintRunc(dAtA, i, uint64(len(m.ShimCgroup)))
		i += copy(dAtA[i:], m.ShimCgroup)
	}
	if m.IoUid != 0 {
		dAtA[i] = 0x50
		i++
		i = encodeVarintRunc(dAtA, i, uint64(m.IoUid))
	}
	if m.IoGid != 0 {
		dAtA[i] = 0x58
		i++
		i = encodeVarintRunc(dAtA, i, uint64(m.IoGid))
	}
	return i, nil
}

func (m *CheckpointOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CheckpointOptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Exit {
		dAtA[i] = 0x8
		i++
		if m.Exit {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.OpenTcp {
		dAtA[i] = 0x10
		i++
		if m.OpenTcp {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.ExternalUnixSockets {
		dAtA[i] = 0x18
		i++
		if m.ExternalUnixSockets {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Terminal {
		dAtA[i] = 0x20
		i++
		if m.Terminal {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.FileLocks {
		dAtA[i] = 0x28
		i++
		if m.FileLocks {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.EmptyNamespaces) > 0 {
		for _, s := range m.EmptyNamespaces {
			dAtA[i] = 0x32
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.CgroupsMode) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintRunc(dAtA, i, uint64(len(m.CgroupsMode)))
		i += copy(dAtA[i:], m.CgroupsMode)
	}
	return i, nil
}

func (m *ProcessDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProcessDetails) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ExecID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRunc(dAtA, i, uint64(len(m.ExecID)))
		i += copy(dAtA[i:], m.ExecID)
	}
	return i, nil
}

func encodeVarintRunc(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RuncOptions) Size() (n int) {
	var l int
	_ = l
	l = len(m.Runtime)
	if l > 0 {
		n += 1 + l + sovRunc(uint64(l))
	}
	l = len(m.RuntimeRoot)
	if l > 0 {
		n += 1 + l + sovRunc(uint64(l))
	}
	l = len(m.CriuPath)
	if l > 0 {
		n += 1 + l + sovRunc(uint64(l))
	}
	if m.SystemdCgroup {
		n += 2
	}
	return n
}

func (m *CreateOptions) Size() (n int) {
	var l int
	_ = l
	if m.NoPivotRoot {
		n += 2
	}
	if m.OpenTcp {
		n += 2
	}
	if m.ExternalUnixSockets {
		n += 2
	}
	if m.Terminal {
		n += 2
	}
	if m.FileLocks {
		n += 2
	}
	if len(m.EmptyNamespaces) > 0 {
		for _, s := range m.EmptyNamespaces {
			l = len(s)
			n += 1 + l + sovRunc(uint64(l))
		}
	}
	l = len(m.CgroupsMode)
	if l > 0 {
		n += 1 + l + sovRunc(uint64(l))
	}
	if m.NoNewKeyring {
		n += 2
	}
	l = len(m.ShimCgroup)
	if l > 0 {
		n += 1 + l + sovRunc(uint64(l))
	}
	if m.IoUid != 0 {
		n += 1 + sovRunc(uint64(m.IoUid))
	}
	if m.IoGid != 0 {
		n += 1 + sovRunc(uint64(m.IoGid))
	}
	return n
}

func (m *CheckpointOptions) Size() (n int) {
	var l int
	_ = l
	if m.Exit {
		n += 2
	}
	if m.OpenTcp {
		n += 2
	}
	if m.ExternalUnixSockets {
		n += 2
	}
	if m.Terminal {
		n += 2
	}
	if m.FileLocks {
		n += 2
	}
	if len(m.EmptyNamespaces) > 0 {
		for _, s := range m.EmptyNamespaces {
			l = len(s)
			n += 1 + l + sovRunc(uint64(l))
		}
	}
	l = len(m.CgroupsMode)
	if l > 0 {
		n += 1 + l + sovRunc(uint64(l))
	}
	return n
}

func (m *ProcessDetails) Size() (n int) {
	var l int
	_ = l
	l = len(m.ExecID)
	if l > 0 {
		n += 1 + l + sovRunc(uint64(l))
	}
	return n
}

func sovRunc(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRunc(x uint64) (n int) {
	return sovRunc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *RuncOptions) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RuncOptions{`,
		`Runtime:` + fmt.Sprintf("%v", this.Runtime) + `,`,
		`RuntimeRoot:` + fmt.Sprintf("%v", this.RuntimeRoot) + `,`,
		`CriuPath:` + fmt.Sprintf("%v", this.CriuPath) + `,`,
		`SystemdCgroup:` + fmt.Sprintf("%v", this.SystemdCgroup) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CreateOptions) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CreateOptions{`,
		`NoPivotRoot:` + fmt.Sprintf("%v", this.NoPivotRoot) + `,`,
		`OpenTcp:` + fmt.Sprintf("%v", this.OpenTcp) + `,`,
		`ExternalUnixSockets:` + fmt.Sprintf("%v", this.ExternalUnixSockets) + `,`,
		`Terminal:` + fmt.Sprintf("%v", this.Terminal) + `,`,
		`FileLocks:` + fmt.Sprintf("%v", this.FileLocks) + `,`,
		`EmptyNamespaces:` + fmt.Sprintf("%v", this.EmptyNamespaces) + `,`,
		`CgroupsMode:` + fmt.Sprintf("%v", this.CgroupsMode) + `,`,
		`NoNewKeyring:` + fmt.Sprintf("%v", this.NoNewKeyring) + `,`,
		`ShimCgroup:` + fmt.Sprintf("%v", this.ShimCgroup) + `,`,
		`IoUid:` + fmt.Sprintf("%v", this.IoUid) + `,`,
		`IoGid:` + fmt.Sprintf("%v", this.IoGid) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CheckpointOptions) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CheckpointOptions{`,
		`Exit:` + fmt.Sprintf("%v", this.Exit) + `,`,
		`OpenTcp:` + fmt.Sprintf("%v", this.OpenTcp) + `,`,
		`ExternalUnixSockets:` + fmt.Sprintf("%v", this.ExternalUnixSockets) + `,`,
		`Terminal:` + fmt.Sprintf("%v", this.Terminal) + `,`,
		`FileLocks:` + fmt.Sprintf("%v", this.FileLocks) + `,`,
		`EmptyNamespaces:` + fmt.Sprintf("%v", this.EmptyNamespaces) + `,`,
		`CgroupsMode:` + fmt.Sprintf("%v", this.CgroupsMode) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ProcessDetails) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ProcessDetails{`,
		`ExecID:` + fmt.Sprintf("%v", this.ExecID) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringRunc(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *RuncOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRunc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RuncOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RuncOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Runtime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRunc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Runtime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RuntimeRoot", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRunc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RuntimeRoot = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CriuPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRunc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CriuPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SystemdCgroup", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.SystemdCgroup = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRunc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRunc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CreateOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRunc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreateOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NoPivotRoot", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.NoPivotRoot = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OpenTcp", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OpenTcp = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalUnixSockets", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ExternalUnixSockets = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Terminal", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Terminal = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileLocks", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.FileLocks = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EmptyNamespaces", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRunc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EmptyNamespaces = append(m.EmptyNamespaces, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CgroupsMode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRunc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CgroupsMode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NoNewKeyring", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.NoNewKeyring = bool(v != 0)
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShimCgroup", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRunc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ShimCgroup = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IoUid", wireType)
			}
			m.IoUid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IoUid |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IoGid", wireType)
			}
			m.IoGid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IoGid |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRunc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRunc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CheckpointOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRunc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CheckpointOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CheckpointOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exit", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Exit = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OpenTcp", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OpenTcp = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalUnixSockets", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ExternalUnixSockets = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Terminal", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Terminal = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileLocks", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.FileLocks = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EmptyNamespaces", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRunc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EmptyNamespaces = append(m.EmptyNamespaces, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CgroupsMode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRunc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CgroupsMode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRunc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRunc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProcessDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRunc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProcessDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProcessDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRunc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExecID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRunc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRunc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRunc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRunc
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRunc
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRunc
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRunc
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRunc
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRunc(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRunc = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRunc   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/neo-hu/docker_el6/containerd/linux/runctypes/runc.proto", fileDescriptorRunc)
}

var fileDescriptorRunc = []byte{
	// 540 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x6b, 0xda, 0x26, 0xce, 0xa4, 0x29, 0xb0, 0x50, 0xc9, 0x14, 0x91, 0x86, 0x00, 0x52,
	0xb8, 0xa4, 0x12, 0x88, 0x0b, 0xbd, 0xb5, 0x45, 0xa8, 0x02, 0x4a, 0x65, 0x5a, 0x09, 0x71, 0x59,
	0xb9, 0xeb, 0x21, 0x59, 0x25, 0xde, 0x59, 0xed, 0xae, 0xa9, 0x73, 0xeb, 0x13, 0xf0, 0x5c, 0x3d,
	0x72, 0xe4, 0x84, 0x68, 0x5e, 0x04, 0xe4, 0x75, 0x1c, 0xb8, 0x72, 0xe5, 0xf6, 0xcf, 0xf7, 0x8f,
	0x3d, 0xa3, 0x7f, 0x35, 0xb0, 0x37, 0x92, 0x6e, 0x9c, 0x9f, 0x0f, 0x05, 0x65, 0xbb, 0x82, 0x94,
	0x4b, 0xa4, 0x42, 0x93, 0xfe, 0x2d, 0xa7, 0x52, 0xe5, 0xc5, 0xae, 0xc9, 0x95, 0x70, 0x33, 0x8d,
	0xd6, 0xab, 0xa1, 0x36, 0xe4, 0x88, 0x6d, 0xfd, 0x69, 0x1b, 0xfa, 0xb6, 0x61, 0x69, 0x6e, 0xdf,
	0x1d, 0xd1, 0x88, 0x7c, 0xc7, 0x6e, 0xa9, 0xaa, 0xe6, 0xfe, 0xd7, 0x00, 0xda, 0x71, 0xae, 0xc4,
	0x7b, 0xed, 0x24, 0x29, 0xcb, 0x22, 0x68, 0x9a, 0x5c, 0x39, 0x99, 0x61, 0x14, 0xf4, 0x82, 0x41,
	0x2b, 0xae, 0x4b, 0xf6, 0x10, 0x36, 0x16, 0x92, 0x1b, 0x22, 0x17, 0xdd, 0xf0, 0x76, 0x7b, 0xc1,
	0x62, 0x22, 0xc7, 0xee, 0x43, 0x4b, 0x18, 0x99, 0x73, 0x9d, 0xb8, 0x71, 0xb4, 0xea, 0xfd, 0xb0,
	0x04, 0x27, 0x89, 0x1b, 0xb3, 0x27, 0xb0, 0x69, 0x67, 0xd6, 0x61, 0x96, 0x72, 0x31, 0x32, 0x94,
	0xeb, 0x68, 0xad, 0x17, 0x0c, 0xc2, 0xb8, 0xb3, 0xa0, 0x07, 0x1e, 0xf6, 0x2f, 0x57, 0xa1, 0x73,
	0x60, 0x30, 0x71, 0x58, 0xaf, 0xd4, 0x87, 0x8e, 0x22, 0xae, 0xe5, 0x17, 0x72, 0xd5, 0xe4, 0xc0,
	0x7f, 0xd7, 0x56, 0x74, 0x52, 0x32, 0x3f, 0xf9, 0x1e, 0x84, 0xa4, 0x51, 0x71, 0x27, 0xb4, 0x5f,
	0x2c, 0x8c, 0x9b, 0x65, 0x7d, 0x2a, 0x34, 0x7b, 0x06, 0x5b, 0x58, 0x38, 0x34, 0x2a, 0x99, 0xf2,
	0x5c, 0xc9, 0x82, 0x5b, 0x12, 0x13, 0x74, 0xd6, 0x2f, 0x18, 0xc6, 0x77, 0x6a, 0xf3, 0x4c, 0xc9,
	0xe2, 0x43, 0x65, 0xb1, 0x6d, 0x08, 0x1d, 0x9a, 0x4c, 0xaa, 0x64, 0xba, 0xd8, 0x72, 0x59, 0xb3,
	0x07, 0x00, 0x9f, 0xe5, 0x14, 0xf9, 0x94, 0xc4, 0xc4, 0x46, 0xeb, 0xde, 0x6d, 0x95, 0xe4, 0x6d,
	0x09, 0xd8, 0x53, 0xb8, 0x85, 0x99, 0x76, 0x33, 0xae, 0x92, 0x0c, 0xad, 0x4e, 0x04, 0xda, 0xa8,
	0xd1, 0x5b, 0x1d, 0xb4, 0xe2, 0x9b, 0x9e, 0x1f, 0x2f, 0x71, 0x99, 0x68, 0x95, 0x84, 0xe5, 0x19,
	0xa5, 0x18, 0x35, 0xab, 0x44, 0x17, 0xec, 0x1d, 0xa5, 0xc8, 0x1e, 0xc3, 0xa6, 0x22, 0xae, 0xf0,
	0x82, 0x4f, 0x70, 0x66, 0xa4, 0x1a, 0x45, 0xa1, 0x1f, 0xb8, 0xa1, 0xe8, 0x18, 0x2f, 0xde, 0x54,
	0x8c, 0xed, 0x40, 0xdb, 0x8e, 0x65, 0x56, 0xe7, 0xda, 0xf2, 0xff, 0x81, 0x12, 0x55, 0xa1, 0xb2,
	0x2d, 0x68, 0x48, 0xe2, 0xb9, 0x4c, 0x23, 0xe8, 0x05, 0x83, 0x4e, 0xbc, 0x2e, 0xe9, 0x4c, 0xa6,
	0x0b, 0x3c, 0x92, 0x69, 0xd4, 0xae, 0xf1, 0x6b, 0x99, 0xf6, 0x7f, 0x05, 0x70, 0xfb, 0x60, 0x8c,
	0x62, 0xa2, 0x49, 0x2a, 0x57, 0x3f, 0x03, 0x83, 0x35, 0x2c, 0x64, 0x9d, 0xbe, 0xd7, 0xff, 0x6b,
	0xec, 0xfd, 0x17, 0xb0, 0x79, 0x62, 0x48, 0xa0, 0xb5, 0x87, 0xe8, 0x12, 0x39, 0xb5, 0xec, 0x11,
	0x34, 0xb1, 0x40, 0xc1, 0x65, 0x5a, 0xdd, 0xc5, 0x3e, 0xcc, 0x7f, 0xec, 0x34, 0x5e, 0x15, 0x28,
	0x8e, 0x0e, 0xe3, 0x46, 0x69, 0x1d, 0xa5, 0xfb, 0xa7, 0x57, 0xd7, 0xdd, 0x95, 0xef, 0xd7, 0xdd,
	0x95, 0xcb, 0x79, 0x37, 0xb8, 0x9a, 0x77, 0x83, 0x6f, 0xf3, 0x6e, 0xf0, 0x73, 0xde, 0x0d, 0x3e,
	0xbd, 0xfc, 0xd7, 0x83, 0xde, 0x5b, 0xaa, 0x8f, 0x2b, 0xe7, 0x0d, 0x7f, 0xab, 0xcf, 0x7f, 0x07,
	0x00, 0x00, 0xff, 0xff, 0xb1, 0xca, 0x85, 0x39, 0x17, 0x04, 0x00, 0x00,
}