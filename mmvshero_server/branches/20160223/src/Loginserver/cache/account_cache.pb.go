// Code generated by protoc-gen-go.
// source: account_cache.proto
// DO NOT EDIT!

/*
Package cache is a generated protocol buffer package.

It is generated from these files:
	account_cache.proto

It has these top-level messages:
	AccountCache
*/
package cache

import proto "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type AccountCache struct {
	RoleId         int64  `protobuf:"varint,1,opt,name=roleId" json:"roleId,omitempty"`
	Username       string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Password       string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	Salt           string `protobuf:"bytes,4,opt,name=salt" json:"salt,omitempty"`
	AccountType    int32  `protobuf:"varint,5,opt,name=accountType" json:"accountType,omitempty"`
	ThirdpartToken string `protobuf:"bytes,6,opt,name=thirdpartToken" json:"thirdpartToken,omitempty"`
	LoginTime      int64  `protobuf:"varint,7,opt,name=loginTime" json:"loginTime,omitempty"`
	CreateTime     int64  `protobuf:"varint,8,opt,name=createTime" json:"createTime,omitempty"`
}

func (m *AccountCache) Reset()         { *m = AccountCache{} }
func (m *AccountCache) String() string { return proto.CompactTextString(m) }
func (*AccountCache) ProtoMessage()    {}

func init() {
}
