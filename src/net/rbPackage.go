package net

import (
	"encoding/binary"
	"unsafe"
)

const BIG_ENDIAN = false
const HEADER_LEN  = 40
const DELETE_3TH  = true

type Header struct {
	packLen uint16  // 整个包的长度
	packFlag uint8
	cmd uint32
	seqId uint32
	svrId uint32
	uid uint64
	reserver1 uint32
	reserver2 uint32
	transactionId uint32
	packExFlag uint32
}

type RbPackage struct {
	header Header // 只放头
	data []byte // 只放数据
	dataLen uint32 // 数据长度
	isToH bool // 头中的长度是否转为大端
}

func NewRbPackage() RbPackage {
	var pack RbPackage
	pack.init()
	return pack
}

func NewRbPackageFormBytes(bytes []byte) RbPackage {
	var pack RbPackage
	pack.initFromBytes(bytes)
	return pack
}

func (rbp *RbPackage) GetBytes() []byte {
	packLen := rbp.dataLen + HEADER_LEN
	data := make([]byte, packLen)

	pHeader := (*[HEADER_LEN]byte)(unsafe.Pointer(&rbp.header))
	// 修正包长度
	if BIG_ENDIAN{
		if !rbp.isToH {
			copy((*pHeader)[0:], htos(uint16(packLen)))
			rbp.isToH = true
		}
	}else{
		rbp.header.packLen = uint16(packLen)
	}
	copyLen := copy(data[0:], (*pHeader)[0:])
	copyLen = copy(data[copyLen:],rbp.data[0:])

	// 去掉第四个字节
	if DELETE_3TH {
		data = delete3thChar(data)
	}
	return data
}

func (rbp *RbPackage) GetData() []byte {
	return rbp.data
}

func (rbp *RbPackage) initFromBytes(inBytes []byte) []byte {
	if len(inBytes) <= 2 {
		return inBytes
	}
	if DELETE_3TH{
		inBytes = add3thChar(inBytes)
	}
	packLen := uint16(len(inBytes))
	if packLen < HEADER_LEN{
		return rbp.data
	}

	pHeader := (*[HEADER_LEN]byte)(unsafe.Pointer(&rbp.header))
	copy(pHeader[0:], inBytes[0:HEADER_LEN])
	// 大端读
	if BIG_ENDIAN {
		packLen = uint16(inBytes[0] << 8) + uint16(inBytes[1])
		rbp.header.packLen = packLen
	}

	rbp.data = make([]byte, rbp.header.packLen - HEADER_LEN)
	copy(rbp.data, inBytes[HEADER_LEN:])

	return rbp.data
}

func (rbp *RbPackage) init() {
	rbp.dataLen = 0
	rbp.data = make([]byte, 1024)
	rbp.isToH = false
}

func (rbp *RbPackage) AddData(data []byte) {
	copyLen := copy(rbp.data[rbp.dataLen:], data[0:])
	rbp.dataLen += uint32(copyLen)
}

func (rbp *RbPackage) SetCmd(cmd uint32) {
	if BIG_ENDIAN{
		pCmd := (*[4]byte)(unsafe.Pointer(&rbp.header.cmd))
		copy(pCmd[0:], htol(cmd))
	}else {
		rbp.header.cmd = cmd
	}
}

func (rbp *RbPackage) SetRoleId(roleId uint64) {
	if BIG_ENDIAN{
		pRoleId := (*[4]byte)(unsafe.Pointer(&rbp.header.uid))
		copy(pRoleId[0:], htoll(roleId))
	}else{
		rbp.header.uid = roleId
	}
}

func (rbp *RbPackage) SetRegionId(regionId uint32) {
	if BIG_ENDIAN {
		pRegionId := (*[4]byte)(unsafe.Pointer(&rbp.header.svrId))
		copy(pRegionId[0:], htol(regionId))
	}else {
		rbp.header.svrId = regionId
	}
}

func htos( number uint16) []byte {
	bytes := make([]byte, 2)
	binary.BigEndian.PutUint16(bytes, number)
	return bytes
}

func htol( number uint32) []byte {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, number)
	return bytes
}

func htoll( number uint64) []byte {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, number)
	return bytes
}

func delete3thChar(input []byte) []byte {
	result :=make([]byte, len(input)-1)
	copy(result[0:2], input[0:])
	copy(result[3:],input[4:])
	result[0]--
	return result
}

func add3thChar(input []byte) []byte {
	result :=make([]byte, len(input)+1)
	copy(result[0:2], input[0:])
	copy(result[4:],input[3:])
	result[0]++
	return result
}
