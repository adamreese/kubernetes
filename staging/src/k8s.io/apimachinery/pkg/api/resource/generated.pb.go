/*
Copyright 2017 The Kubernetes Authors.

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

// Code generated by protoc-gen-gogo.
// source: k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/api/resource/generated.proto
// DO NOT EDIT!

/*
	Package resource is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/api/resource/generated.proto

	It has these top-level messages:
		Quantity
*/
package resource

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

func (m *Quantity) Reset()                    { *m = Quantity{} }
func (*Quantity) ProtoMessage()               {}
func (*Quantity) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func init() {
	proto.RegisterType((*Quantity)(nil), "k8s.io.apimachinery.pkg.api.resource.Quantity")
}

var fileDescriptorGenerated = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x8f, 0xb1, 0x4a, 0x03, 0x41,
	0x10, 0x86, 0x77, 0x1b, 0x89, 0x57, 0x06, 0x11, 0x49, 0xb1, 0x17, 0xc4, 0x42, 0x04, 0x77, 0xb0,
	0x0b, 0x96, 0xf6, 0x16, 0x5a, 0xda, 0xdd, 0x5d, 0xc6, 0xcd, 0x72, 0x66, 0xf7, 0x98, 0x9d, 0x15,
	0xd2, 0xa5, 0xb4, 0x4c, 0x69, 0x99, 0x7b, 0x9b, 0x94, 0x29, 0x2d, 0x2c, 0xbc, 0xf3, 0x45, 0xe4,
	0x2e, 0x09, 0x88, 0x76, 0xf3, 0xff, 0xc3, 0x37, 0x7c, 0x93, 0xdc, 0x97, 0x93, 0xa0, 0xad, 0x87,
	0x32, 0xe6, 0x48, 0x0e, 0x19, 0x03, 0xbc, 0xa2, 0x9b, 0x7a, 0x82, 0xfd, 0x22, 0xab, 0xec, 0x3c,
	0x2b, 0x66, 0xd6, 0x21, 0x2d, 0xa0, 0x2a, 0x4d, 0x57, 0x00, 0x61, 0xf0, 0x91, 0x0a, 0x04, 0x83,
	0x0e, 0x29, 0x63, 0x9c, 0xea, 0x8a, 0x3c, 0xfb, 0xe1, 0xc5, 0x8e, 0xd2, 0xbf, 0x29, 0x5d, 0x95,
	0xa6, 0x2b, 0xf4, 0x81, 0x1a, 0x5d, 0x1b, 0xcb, 0xb3, 0x98, 0xeb, 0xc2, 0xcf, 0xc1, 0x78, 0xe3,
	0xa1, 0x87, 0xf3, 0xf8, 0xdc, 0xa7, 0x3e, 0xf4, 0xd3, 0xee, 0xe8, 0xe8, 0xe6, 0xbf, 0x63, 0x27,
	0x12, 0xd9, 0xbe, 0x80, 0x75, 0x1c, 0x98, 0xfe, 0x7a, 0x9c, 0x4f, 0x92, 0xc1, 0x43, 0xcc, 0x1c,
	0x5b, 0x5e, 0x0c, 0x4f, 0x93, 0xa3, 0xc0, 0x64, 0x9d, 0x39, 0x93, 0x63, 0x79, 0x79, 0xfc, 0xb8,
	0x4f, 0xb7, 0x27, 0xef, 0xeb, 0x54, 0xbc, 0xd5, 0xa9, 0x58, 0xd5, 0xa9, 0x58, 0xd7, 0xa9, 0x58,
	0x7e, 0x8e, 0xc5, 0xdd, 0xd5, 0xa6, 0x51, 0x62, 0xdb, 0x28, 0xf1, 0xd1, 0x28, 0xb1, 0x6c, 0x95,
	0xdc, 0xb4, 0x4a, 0x6e, 0x5b, 0x25, 0xbf, 0x5a, 0x25, 0x57, 0xdf, 0x4a, 0x3c, 0x0d, 0x0e, 0x7f,
	0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xaa, 0x3e, 0x00, 0x3d, 0x01, 0x00, 0x00,
}
