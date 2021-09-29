package utils

import (
	"encoding/binary"
	"math"
	"math/big"
	"net"
	"runtime"
)

//PanicTrace panic捕获并且输入到日志中
func PanicTrace() []byte {
	kbSize := 1024
	stack := make([]byte, kbSize<<10) //4KB
	stack = stack[:runtime.Stack(stack, false)]
	return stack
}

func Ip2long(ipstr string) uint32 {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}

func Long2ip(ipLong uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, ipLong)
	ip := net.IP(ipByte)
	return ip.String()
}

func TryParseAmount(Reserve0, Reserve1 string, Decimals0, Decimals1 int, base bool) float64 {

	rs0, _ := new(big.Float).SetString(Reserve0)
	rs1, _ := new(big.Float).SetString(Reserve1)
	value0 := new(big.Float).Quo(rs0, big.NewFloat(math.Pow10(Decimals0)))
	value1 := new(big.Float).Quo(rs1, big.NewFloat(math.Pow10(Decimals1)))
	if base {
		price, _ := new(big.Float).Quo(value0, value1).Float64()
		return price
	}

	price, _ := new(big.Float).Quo(value1, value0).Float64()
	return price
}

func TryParsePrice(Reserve0, Reserve1 *big.Float, Decimals0, Decimals1 int, base bool) float64 {

	value0 := new(big.Float).Quo(Reserve0, big.NewFloat(math.Pow10(Decimals0)))
	value1 := new(big.Float).Quo(Reserve1, big.NewFloat(math.Pow10(Decimals1)))
	if base {
		price, _ := new(big.Float).Quo(value0, value1).Float64()
		return price
	}

	price, _ := new(big.Float).Quo(value1, value0).Float64()
	return price
}
