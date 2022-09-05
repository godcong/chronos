package utils

import (
	"encoding/binary"
	"time"
)

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

func BytesToTime(buf []byte) time.Time {
	return time.Unix(int64(binary.BigEndian.Uint64(buf)), 0).UTC()
}
