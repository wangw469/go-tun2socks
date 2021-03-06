package route

import (
	vnet "v2ray.com/core/common/net"
)

const (
	IPVERSION_4 = 4
	IPVERSION_6 = 6

	PROTOCOL_ICMP = 1
	PROTOCOL_TCP  = 6
	PROTOCOL_UDP  = 17
)

func PeekIPVersion(data []byte) uint8 {
	return uint8((data[0] & 0xf0) >> 4)
}

func PeekProtocol(data []byte) string {
	switch uint8(data[9]) {
	case PROTOCOL_ICMP:
		return "icmp"
	case PROTOCOL_TCP:
		return "tcp"
	case PROTOCOL_UDP:
		return "udp"
	default:
		return "unknown"
	}
}

func PeekDestinationAddress(data []byte) vnet.Address {
	return vnet.IPAddress(data[16:20])
}

func PeekDestinationPort(data []byte) vnet.Port {
	ihl := uint8(data[0] & 0x0f)
	return vnet.PortFromBytes(data[ihl*4+2 : ihl*4+4])
}

func IsSYNSegment(data []byte) bool {
	ihl := uint8(data[0] & 0x0f)
	if uint8(data[ihl*4+13]&(1<<1)) == 0 {
		return false
	} else {
		return true
	}
}
