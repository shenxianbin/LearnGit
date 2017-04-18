package utils

import (
	"bytes"
	"encoding/binary"
	"math"
)

type BinaryBuffer struct {
	bytes.Buffer
}

func NewBinaryBuffer(default_size int) *BinaryBuffer {
	buf := new(BinaryBuffer)
	buf.Grow(default_size)
	return buf
}

func (this *BinaryBuffer) ReadByte() byte {
	b, _ := this.Buffer.ReadByte()
	return b
}

func (this *BinaryBuffer) WriteByte(v byte) {
	this.Buffer.WriteByte(v)
}

func (this *BinaryBuffer) ReadBool() bool {
	return this.ReadByte() != 0
}

func (this *BinaryBuffer) WriteBool(v bool) {
	if v {
		this.WriteByte(1)
	} else {
		this.WriteByte(0)
	}
}

func (this *BinaryBuffer) ReadInt8() int8 {
	return int8(this.ReadByte())
}

func (this *BinaryBuffer) WriteInt8(v int8) {
	this.WriteByte(byte(v))
}

func (this *BinaryBuffer) ReadUint8() uint8 {
	return uint8(this.ReadByte())
}

func (this *BinaryBuffer) WriteUint8(v uint8) {
	this.WriteByte(byte(v))
}

func (this *BinaryBuffer) ReadInt16() int16 {
	b := make([]byte, 2)
	if _, err := this.Buffer.Read(b); err != nil {
		return 0
	}
	return int16(binary.BigEndian.Uint16(b))
}

func (this *BinaryBuffer) WriteInt16(v int16) {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(v))
	this.Buffer.Write(b)
}

func (this *BinaryBuffer) ReadUint16() uint16 {
	b := make([]byte, 2)
	if _, err := this.Buffer.Read(b); err != nil {
		return 0
	}
	return binary.BigEndian.Uint16(b)
}

func (this *BinaryBuffer) WriteUint16(v uint16) {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, v)
	this.Buffer.Write(b)
}

func (this *BinaryBuffer) ReadInt32() int32 {
	b := make([]byte, 4)
	if _, err := this.Buffer.Read(b); err != nil {
		return 0
	}
	return int32(binary.BigEndian.Uint32(b))
}

func (this *BinaryBuffer) WriteInt32(v int32) {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(v))
	this.Buffer.Write(b)
}

func (this *BinaryBuffer) ReadUint32() uint32 {
	b := make([]byte, 4)
	if _, err := this.Buffer.Read(b); err != nil {
		return 0
	}
	return binary.BigEndian.Uint32(b)
}

func (this *BinaryBuffer) WriteUint32(v uint32) {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, v)
	this.Buffer.Write(b)
}

func (this *BinaryBuffer) ReadInt64() int64 {
	b := make([]byte, 8)
	if _, err := this.Buffer.Read(b); err != nil {
		return 0
	}
	return int64(binary.BigEndian.Uint64(b))
}

func (this *BinaryBuffer) WriteInt64(v int64) {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	this.Buffer.Write(b)
}

func (this *BinaryBuffer) ReadUint64() uint64 {
	b := make([]byte, 8)
	if _, err := this.Buffer.Read(b); err != nil {
		return 0
	}
	return binary.BigEndian.Uint64(b)
}

func (this *BinaryBuffer) WriteUint64(v uint64) {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	this.Buffer.Write(b)
}

func (this *BinaryBuffer) ReadFloat32() float32 {
	b := make([]byte, 4)
	if _, err := this.Buffer.Read(b); err != nil {
		return 0.0
	}
	return math.Float32frombits(binary.BigEndian.Uint32(b))
}

func (this *BinaryBuffer) WriteFloat32(v float32) {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, math.Float32bits(v))
	this.Buffer.Write(b)
}

func (this *BinaryBuffer) ReadFloat64() float64 {
	b := make([]byte, 8)
	if _, err := this.Buffer.Read(b); err != nil {
		return 0
	}
	return math.Float64frombits(binary.BigEndian.Uint64(b))
}

func (this *BinaryBuffer) WriteFloat64(v float64) {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, math.Float64bits(v))
	this.Buffer.Write(b)
}

func (this *BinaryBuffer) ReadString() string {
	size := this.ReadUint16()
	b := make([]byte, size)
	var offset uint16 = 0
	for offset < size {
		b[offset] = this.ReadByte()
		offset++
	}
	return string(b)
}

func (this *BinaryBuffer) WriteString(v string) {
	var size uint16 = uint16(len(v))
	this.WriteUint16(size)
	b := []byte(v)
	var offset uint16 = 0
	for offset < size {
		this.WriteByte(b[offset])
		offset++
	}
}
