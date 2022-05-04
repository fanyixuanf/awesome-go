/**
 * @Author: Yuxue.Fan<yuxue.fan@yidatec.com>
 * @Description:
 * @File: md5.go
 * @Version: 1.0.0
 * @Date: 2022/2/18 14:45
 */

package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5V(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}
