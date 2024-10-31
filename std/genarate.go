package std

import (
	cryRand "crypto/rand"
	"encoding/binary"
	"io"
	"log/slog"
	"math/rand"
	"time"
)

// RandStr 生成固定位数对随机值
func RandStr(n int) string {
	var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// SecureRandomInt 使用crypto/rand生成安全的随机数
// crypto/rand生成的随机数不容易被预测，
// 但相比math/rand，可能在性能上稍微慢一些，因为使用了操作系统的随机数生成器。
func SecureRandomInt(min, max int) int {
	// 8字节足以生成一个int64
	randomBytes := make([]byte, 8)
	// 使用crypto/rand的Read函数，从加密安全的源读取数据，填充随机字节
	if _, err := io.ReadFull(cryRand.Reader, randomBytes); err != nil {
		// 尽管生成失败这种情况很少见
		slog.Debug("SecureRandomInt err:", err)
		// 使用math/rand生成随机整数
		rand.New(rand.NewSource(time.Now().UnixNano()))
		return rand.Intn(max-min+1) + min
	}
	// 转换为int64
	randomUint64 := uint64(binary.BigEndian.Uint64(randomBytes))
	// 调整随机数到所需的范围内
	secureRandomInt := min + int(randomUint64%uint64(max-min+1))
	return secureRandomInt
}
