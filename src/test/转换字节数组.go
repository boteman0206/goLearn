package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
)

func main() {

	num := []int{346, 210, 220, 345, 212, 237}

	for _, v := range num {
		toBytes := IntToBytes(v)
		fmt.Println(string(toBytes))
	}

	data := `
	{\"msg\":\"\346\210\220\345\212\237\",\"result\":{\"orgCode\":null,\"resultCode\":\"1009001\",\"outSkuId\":null,\"failedDetail\":\"\347\261\273\347\233\256\344\277\256\346\224\271\345\244\261\350\264\245\357\274\214\346\255\243\347\241\256\347\261\273\347\233\256Id\357\274\23224113\357\274\214\347\261\273\347\233\256\345\220\215\347\247\260\357\274\232\345\251\264\345\204\277\347\272\270\345\267\276/\345\251\264\345\204\277\346\271\277\345\267\276\",\"skuId\":null,\"upcCode\":null,\"resultMsg\":\"\345\244\261\350\264\245,\345\216\237\345\233\240\347\274\226\345\217\267[1007085]\"},\"code\":\"0\",\"success\":true,\"detail\":null}
`

	strings := make(map[string]string, 0)

	unmarshal := json.Unmarshal([]byte(data), &strings)
	fmt.Println(unmarshal, strings)

}

func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, x)
	return bytesBuffer.Bytes()
}
