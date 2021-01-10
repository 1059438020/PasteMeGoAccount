package util

import (
	"crypto/md5"
	"encoding/hex"
)

func EncryptionPassword(password string, salt string) string {
	m5 := md5.New()
	m5.Write([]byte(password))
	m5.Write([]byte(salt))
	encryptionPassword := hex.EncodeToString(m5.Sum(nil))
	return encryptionPassword
}
