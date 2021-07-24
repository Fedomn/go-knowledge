package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

// http://vearne.cc/archives/543
// int32都占4个字节，不论这个数据是100、1000、还是1000,000。vint的理念是采用可变长的字节，来表示一个整数。数值较大的数，使用较多的字节来表示，数值较少的数，使用较少的字节来表示
// 每个字节仅使用第1至第7位(共7bits)，第8位作为标识，表明(为了得到这个整型)是否需要读取下一个字节。

// 如下encode64 来说就是 0xffff ffff ffff ffff，一个f代表4bit位，所以未经过压缩，uint64它需要8个byte
func TestVarintEncode(t *testing.T) {
	res := EncodeVarint(nil, 1)
	fmt.Println(res)
	require.Len(t, res, 1)
}

func TestVarintDecode(t *testing.T) {
	encodedRes := EncodeVarint(nil, 1)
	_, v, _ := DecodeVarint(encodedRes)
	fmt.Println(v)
	require.EqualValues(t, 1, v)
}

// EncodeVarint appends the encoded value to slice b and returns the appended slice.
// Note that the encoded result is not memcomparable.
func EncodeVarint(b []byte, v int64) []byte {
	var data [binary.MaxVarintLen64]byte
	n := binary.PutVarint(data[:], v)
	return append(b, data[:n]...)
}

// DecodeVarint decodes value encoded by EncodeVarint before.
// It returns the leftover un-decoded slice, decoded value if no error.
func DecodeVarint(b []byte) ([]byte, int64, error) {
	v, n := binary.Varint(b)
	if n > 0 {
		return b[n:], v, nil
	}
	if n < 0 {
		return nil, 0, errors.New("value larger than 64 bits")
	}
	return nil, 0, errors.New("insufficient bytes to decode value")
}
