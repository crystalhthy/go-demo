package lib

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

func EncryptIntToStr(id int) (string, error) {

	hashBytes := []byte(fmt.Sprintf("hello_%d", id))
	return hex.EncodeToString(hashBytes), nil
}

func DecryptStrToInt(str string) (int, error) {

	var id int
	var sid []byte
	var err error
	sid, err = hex.DecodeString(str)
	if err != nil {
		return 0, err
	}

	id, err = strconv.Atoi(string(sid)[6:])
	if err != nil {
		return 0, err
	}
	return id, nil
}
