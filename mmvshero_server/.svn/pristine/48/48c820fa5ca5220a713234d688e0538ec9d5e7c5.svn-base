package ip

import (
	"encoding/binary"
	"errors"
	. "galaxy"
	"io/ioutil"
	"net"
	"os"
	"strings"

	zh "golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

var ip_query *IpQuery

const (
	INDEX_LEN       = 7
	REDIRECT_MODE_1 = 0x01
	REDIRECT_MODE_2 = 0x02
)

type IpQuery struct {
	file       *os.File
	indexStart uint32
	indexEnd   uint32
}

func Init(filepath string) error {
	query, err := newIpQuery(filepath)
	if err != nil {
		return err
	}
	ip_query = query
	return nil
}

func Find(ip string) (string, error) {
	return ip_query.Find(ip)
}

func Close() {
	ip_query.Close()
}

func newIpQuery(filepath string) (*IpQuery, error) {
	iq := new(IpQuery)
	if err := iq.init(filepath); err != nil {
		return nil, err
	}
	return iq, nil
}

func (this *IpQuery) init(filepath string) error {
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0400)
	if err != nil {
		LogError(err)
		return err
	}
	this.file = file

	buf := make([]byte, 8)
	this.file.ReadAt(buf, 0)

	this.indexStart = binary.LittleEndian.Uint32(buf[0:4])
	this.indexEnd = binary.LittleEndian.Uint32(buf[4:8])
	LogInfo("Ip Index range: ", this.indexStart, " - ", this.indexEnd)
	return nil
}

func (this *IpQuery) Close() {
	this.file.Close()
}

func (this *IpQuery) Find(ip string) (string, error) {
	LogDebug("IpQuery.Find ip : ", ip)
	if ip == "" {
		return "", errors.New("Ip Null Error")
	}

	offset := this.searchIndex(binary.BigEndian.Uint32(net.ParseIP(ip).To4()))
	if offset <= 0 {
		return "", errors.New("Data Offset Error")
	}

	var country string
	var area string

	mode := this.readMode(offset + 4)
	if mode == REDIRECT_MODE_1 {
		countryOffset := this.readUInt24()
		mode = this.readMode(countryOffset)
		if mode == REDIRECT_MODE_2 {
			c := this.readUInt24()
			country = this.readString(c)
			countryOffset += 4
		} else {
			country = this.readString(countryOffset)
			countryOffset += uint32(len(country) + 1)
		}
		area = this.readArea(countryOffset)
	} else if mode == REDIRECT_MODE_2 {
		countryOffset := this.readUInt24()
		country = this.readString(countryOffset)
		area = this.readArea(offset + 8)
	} else {
		country = this.readString(offset + 4)
		area = this.readArea(offset + uint32(5+len(country)))
	}

	var tr *transform.Reader
	tr = transform.NewReader(strings.NewReader(country), zh.GBK.NewDecoder())
	if s, err := ioutil.ReadAll(tr); err == nil {
		country = string(s)
	}

	tr = transform.NewReader(strings.NewReader(area), zh.GBK.NewDecoder())
	if s, err := ioutil.ReadAll(tr); err == nil {
		area = string(s)
	}

	return country + area, nil
}

func (this *IpQuery) searchIndex(ip uint32) uint32 {
	start := this.indexStart
	end := this.indexEnd
	for {
		mid := this.getMiddleOffset(start, end)
		this.file.Seek(int64(mid), 0)
		buf := make([]byte, INDEX_LEN)
		this.file.Read(buf)
		_ip := binary.LittleEndian.Uint32(buf[:4])

		if end-start == INDEX_LEN {
			offset := byte3ToUInt32(buf[4:])
			this.file.Read(buf)
			if ip < binary.LittleEndian.Uint32(buf[:4]) {
				return offset
			} else {
				return 0
			}
		}

		if _ip > ip {
			end = mid
		} else if _ip < ip {
			start = mid
		} else if _ip == ip {
			return byte3ToUInt32(buf[4:])
		}

	}
	return 0
}

func (this *IpQuery) readMode(offset uint32) byte {
	this.file.Seek(int64(offset), 0)
	mode := make([]byte, 1)
	this.file.Read(mode)
	return mode[0]
}

func (this *IpQuery) readArea(offset uint32) string {
	mode := this.readMode(offset)
	if mode == REDIRECT_MODE_1 || mode == REDIRECT_MODE_2 {
		areaOffset := this.readUInt24()
		if areaOffset == 0 {
			return ""
		} else {
			return this.readString(areaOffset)
		}
	} else {
		return this.readString(offset)
	}
	return ""
}

func (this *IpQuery) readString(offset uint32) string {
	this.file.Seek(int64(offset), 0)
	data := make([]byte, 0, 30)
	buf := make([]byte, 1)
	for {
		this.file.Read(buf)
		if buf[0] == 0 {
			break
		}
		data = append(data, buf[0])
	}
	return string(data)
}

func (this *IpQuery) readUInt24() uint32 {
	buf := make([]byte, 3)
	this.file.Read(buf)
	return byte3ToUInt32(buf)
}

func (this *IpQuery) getMiddleOffset(start uint32, end uint32) uint32 {
	records := ((end - start) / INDEX_LEN) >> 1
	return start + records*INDEX_LEN
}

func byte3ToUInt32(data []byte) uint32 {
	i := uint32(data[0]) & 0xff
	i |= (uint32(data[1]) << 8) & 0xff00
	i |= (uint32(data[2]) << 16) & 0xff0000
	return i
}
