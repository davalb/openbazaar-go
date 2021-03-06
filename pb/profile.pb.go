// Code generated by protoc-gen-go.
// source: profile.proto
// DO NOT EDIT!

/*
Package profile is a generated protocol buffer package.

It is generated from these files:
	profile.proto

It has these top-level messages:
	Profile
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Profile struct {
	Handle           string                   `protobuf:"bytes,1,opt,name=handle" json:"handle,omitempty"`
	Name             string                   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Location         string                   `protobuf:"bytes,3,opt,name=location" json:"location,omitempty"`
	About            string                   `protobuf:"bytes,4,opt,name=about" json:"about,omitempty"`
	ShortDescription string                   `protobuf:"bytes,5,opt,name=shortDescription" json:"shortDescription,omitempty"`
	Website          string                   `protobuf:"bytes,6,opt,name=website" json:"website,omitempty"`
	Email            string                   `protobuf:"bytes,7,opt,name=email" json:"email,omitempty"`
	PhoneNumber      string                   `protobuf:"bytes,8,opt,name=phoneNumber" json:"phoneNumber,omitempty"`
	Social           []*Profile_SocialAccount `protobuf:"bytes,9,rep,name=social" json:"social,omitempty"`
	Nsfw             bool                     `protobuf:"varint,10,opt,name=nsfw" json:"nsfw,omitempty"`
	Vendor           bool                     `protobuf:"varint,11,opt,name=vendor" json:"vendor,omitempty"`
	Moderator        bool                     `protobuf:"varint,12,opt,name=moderator" json:"moderator,omitempty"`
	PrimaryColor     string                   `protobuf:"bytes,13,opt,name=primaryColor" json:"primaryColor,omitempty"`
	SecondaryColor   string                   `protobuf:"bytes,14,opt,name=secondaryColor" json:"secondaryColor,omitempty"`
	TextColor        string                   `protobuf:"bytes,15,opt,name=textColor" json:"textColor,omitempty"`
	FollowerCount    uint32                   `protobuf:"varint,16,opt,name=followerCount" json:"followerCount,omitempty"`
	FollowingCount   uint32                   `protobuf:"varint,17,opt,name=followingCount" json:"followingCount,omitempty"`
	ListingCount     uint32                   `protobuf:"varint,18,opt,name=listingCount" json:"listingCount,omitempty"`
	LastModified     uint32                   `protobuf:"varint,19,opt,name=lastModified" json:"lastModified,omitempty"`
}

func (m *Profile) Reset()                    { *m = Profile{} }
func (m *Profile) String() string            { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()               {}
func (*Profile) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Profile) GetSocial() []*Profile_SocialAccount {
	if m != nil {
		return m.Social
	}
	return nil
}

type Profile_SocialAccount struct {
	Type     string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Proof    string `protobuf:"bytes,3,opt,name=proof" json:"proof,omitempty"`
}

func (m *Profile_SocialAccount) Reset()                    { *m = Profile_SocialAccount{} }
func (m *Profile_SocialAccount) String() string            { return proto.CompactTextString(m) }
func (*Profile_SocialAccount) ProtoMessage()               {}
func (*Profile_SocialAccount) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

func init() {
	proto.RegisterType((*Profile)(nil), "Profile")
	proto.RegisterType((*Profile_SocialAccount)(nil), "Profile.SocialAccount")
}

var fileDescriptor3 = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x92, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0x87, 0x35, 0xb6, 0xa6, 0xed, 0xe9, 0x32, 0x86, 0x41, 0x93, 0x35, 0x71, 0x51, 0x4d, 0x08,
	0x4d, 0x5c, 0xe4, 0x02, 0x9e, 0x00, 0x8d, 0x5b, 0x10, 0x0a, 0xe2, 0x01, 0x9c, 0xe4, 0x84, 0x5a,
	0x72, 0x7c, 0x22, 0xdb, 0xa1, 0xec, 0xc5, 0x78, 0x3e, 0xe2, 0x93, 0x34, 0x6d, 0xca, 0x55, 0xfd,
	0xfb, 0xce, 0x57, 0xff, 0x39, 0x39, 0x90, 0xb6, 0x8e, 0x6a, 0x6d, 0x30, 0xeb, 0x7f, 0x03, 0x3d,
	0xfc, 0x5d, 0xc0, 0xf2, 0xfb, 0x40, 0xc4, 0x1d, 0x24, 0x3b, 0x65, 0x2b, 0x83, 0xf2, 0x62, 0x7b,
	0xf1, 0xb8, 0xce, 0xc7, 0x24, 0x04, 0x5c, 0x59, 0xd5, 0xa0, 0x7c, 0xc1, 0x94, 0xd7, 0xe2, 0x1e,
	0x56, 0x86, 0x4a, 0x15, 0x34, 0x59, 0x79, 0xc9, 0x7c, 0xca, 0xe2, 0x0d, 0x2c, 0x54, 0x41, 0x5d,
	0x90, 0x57, 0x5c, 0x18, 0x82, 0xf8, 0x00, 0xb7, 0x7e, 0x47, 0x2e, 0x7c, 0x41, 0x5f, 0x3a, 0xdd,
	0xf2, 0x3f, 0x17, 0x2c, 0xfc, 0xc7, 0x85, 0x84, 0xe5, 0x1e, 0x0b, 0xaf, 0x03, 0xca, 0x84, 0x95,
	0x43, 0x8c, 0x7b, 0x63, 0xa3, 0xb4, 0x91, 0xcb, 0x61, 0x6f, 0x0e, 0x62, 0x0b, 0x9b, 0x76, 0x47,
	0x16, 0xbf, 0x75, 0x4d, 0x81, 0x4e, 0xae, 0xb8, 0x76, 0x8a, 0x44, 0x06, 0x89, 0xa7, 0x52, 0x2b,
	0x23, 0xd7, 0xdb, 0xcb, 0xc7, 0xcd, 0xc7, 0xbb, 0x6c, 0x7c, 0x75, 0xf6, 0x83, 0xf1, 0xe7, 0xb2,
	0xa4, 0xce, 0x86, 0x7c, 0xb4, 0xf8, 0xcd, 0xbe, 0xde, 0x4b, 0xe8, 0xb7, 0x5a, 0xe5, 0xbc, 0x8e,
	0xfd, 0xf9, 0x8d, 0xb6, 0x22, 0x27, 0x37, 0x4c, 0xc7, 0x24, 0xde, 0xc2, 0xba, 0xa1, 0x0a, 0x9d,
	0x0a, 0x7d, 0xe9, 0x9a, 0x4b, 0x47, 0x20, 0x1e, 0xe0, 0xba, 0x75, 0xba, 0x51, 0xee, 0xf9, 0x89,
	0x4c, 0x2f, 0xa4, 0x7c, 0xb9, 0x19, 0x13, 0xef, 0xe1, 0xc6, 0x63, 0x49, 0xb6, 0x9a, 0xac, 0x1b,
	0xb6, 0xce, 0x68, 0x3c, 0x29, 0xe0, 0x9f, 0x30, 0x28, 0x2f, 0x59, 0x39, 0x02, 0xf1, 0x0e, 0xd2,
	0x9a, 0x8c, 0xa1, 0x3d, 0xba, 0xa7, 0xf8, 0x18, 0x79, 0xdb, 0x1b, 0x69, 0x3e, 0x87, 0xf1, 0xac,
	0x01, 0x68, 0xfb, 0x6b, 0xd0, 0x5e, 0xb1, 0x76, 0x46, 0xe3, 0xbd, 0x8d, 0xf6, 0x61, 0xb2, 0x04,
	0x5b, 0x33, 0xc6, 0x8e, 0xf2, 0xe1, 0x2b, 0x55, 0xba, 0xd6, 0x58, 0xc9, 0xd7, 0xa3, 0x73, 0xc2,
	0xee, 0x7f, 0x42, 0x3a, 0x6b, 0x71, 0x6c, 0x6d, 0x78, 0x6e, 0x0f, 0x43, 0xc6, 0xeb, 0x38, 0x4e,
	0x9d, 0x47, 0x77, 0x32, 0x66, 0x53, 0x8e, 0x9f, 0xbc, 0x9f, 0x55, 0xaa, 0xc7, 0x39, 0x1b, 0x42,
	0x91, 0xf0, 0xfc, 0x7e, 0xfa, 0x17, 0x00, 0x00, 0xff, 0xff, 0x84, 0xeb, 0xbf, 0x85, 0xd0, 0x02,
	0x00, 0x00,
}
