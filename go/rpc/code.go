package rpc

import (
	"bytes"
	"encoding/gob"
)

func decode(bs []byte) (RpcData, error) {
	buf := bytes.NewBuffer(bs)
	d := gob.NewDecoder(buf)
	var data RpcData
	err := d.Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func encode(data RpcData) ([]byte, error) {
	var buf bytes.Buffer
	bufEncode := gob.NewEncoder(&buf)
	err := bufEncode.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
