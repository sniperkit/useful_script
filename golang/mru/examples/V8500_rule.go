package main

import (
	//"cline"
	"command"
	//"configuration"
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"rut"
	"strconv"
	"strings"
	"util"
)

type TLV struct {
	Name   string
	Size   int
	Offset int
	Value  string
}

type RuleField struct {
	Key  []*TLV
	Data string
	Mask string
}

type RuleEntry struct {
	Mode  int
	Index int
	Parts []*RuleEntry
	Role  int
}

type RuleDB struct {
	Device             string
	SliceCount         int
	VirtualSliceMap    map[int64]int64 /* key physicla slice, value virtual slice */
	GroupCount         int
	SliceGroupMap      map[int64]int64 /* key slice, value group */
	SliceEntryCountMap map[int64]int64
	SliceStartIndexMap map[int64]int64
	SliceEndIndexMap   map[int64]int64
	EntryToSliceMap    map[int64]int64
	Entries            []*RuleEntry
	EntryCount         int64
	PFS                map[int64]*PortFieldSelector
}

var DB RuleDB = RuleDB{
	VirtualSliceMap:    make(map[int64]int64, 1),
	SliceGroupMap:      make(map[int64]int64, 1),
	SliceEntryCountMap: make(map[int64]int64, 1),
	SliceStartIndexMap: make(map[int64]int64, 1),
	SliceEndIndexMap:   make(map[int64]int64, 1),
	EntryToSliceMap:    make(map[int64]int64, 1),
	PFS:                make(map[int64]*PortFieldSelector, 1),
}

//Refer to AG201
//FP_SLICE_INDEX_CONTROL: Select Index for FP Port Field Select Table(UDF/Source Port)

type ICAPFieldSelector struct {
	FP0             map[int][]TLV
	FP1             map[int][]TLV
	FP2             map[int][]TLV
	FP3             map[int][]TLV
	FP4             map[int][]TLV
	FIXED           map[int][]TLV
	PAIRING_IPBM_F0 map[int][]TLV
	PAIRING_FIXED   map[int][]TLV
}

var BCM56540ICAPFieldSelector_single = ICAPFieldSelector{
	FP1: map[int][]TLV{
		0: []TLV{
			TLV{Name: "SVP_L3_IIF", Size: 16, Offset: 24},
			TLV{Name: "FORWARDING_FIELD", Size: 14, Offset: 10},
			TLV{Name: "SRC_DEST_CLASSID", Size: 10, Offset: 0},
		},
		1: []TLV{
			TLV{Name: "MH_OPCODE", Size: 3, Offset: 38},
			TLV{Name: "SVP_VALID", Size: 1, Offset: 37},
			TLV{Name: "S_FIELD", Size: 16, Offset: 21},
			TLV{Name: "D_TYPE", Size: 3, Offset: 18},
			TLV{Name: "D_FIELD", Size: 18, Offset: 0},
		},
		2: []TLV{
			TLV{Name: "SRC_CLASSID", Size: 10, Offset: 34},
			TLV{Name: "DEST_CLASSID", Size: 10, Offset: 24},
			TLV{Name: "MH_OPCODE", Size: 3, Offset: 21},
			TLV{Name: "D_TYPE", Size: 3, Offset: 18},
			TLV{Name: "D_FIELD", Size: 18, Offset: 0},
		},
		3: []TLV{
			TLV{Name: "OUTER_TPID_ENCODE", Size: 2, Offset: 32},
			TLV{Name: "ITAG", Size: 16, Offset: 16},
			TLV{Name: "OTAG", Size: 16, Offset: 0},
		},
		4: []TLV{
			TLV{Name: "FC_HDR_ENCODE_2", Size: 3, Offset: 35},
			TLV{Name: "FC_HDR_ENCODE_1", Size: 3, Offset: 32},
			TLV{Name: "ETHERTYPE", Size: 16, Offset: 16},
			TLV{Name: "OTAG", Size: 16, Offset: 0},
		},
		5: []TLV{
			TLV{Name: "INNER_TPID_ENCODE", Size: 2, Offset: 42},
			TLV{Name: "ITAG", Size: 16, Offset: 26},
			TLV{Name: "PKT_RES", Size: 6, Offset: 20},
			TLV{Name: "LOOKUP_STATUS", Size: 20, Offset: 0},
		},
		6: []TLV{
			TLV{Name: "FC_HDR_ENCODE_2", Size: 3, Offset: 37},
			TLV{Name: "FC_HDR_ENCODE_1", Size: 3, Offset: 34},
			TLV{Name: "IP_INFO", Size: 3, Offset: 31},
			TLV{Name: "PKT_RES", Size: 6, Offset: 25},
			TLV{Name: "MH_OPCODE", Size: 3, Offset: 22},
			TLV{Name: "TAG_STATUS", Size: 2, Offset: 20},
			TLV{Name: "PKT_FORMAT", Size: 4, Offset: 16},
			TLV{Name: "OTAG", Size: 16, Offset: 0},
		},
		7: []TLV{
			TLV{Name: "RAL_GAL", Size: 2, Offset: 40},
			TLV{Name: "FORWARDING_FIELD", Size: 14, Offset: 26},
			TLV{Name: "SRC_CLASSID", Size: 10, Offset: 16},
			TLV{Name: "OTAG", Size: 16, Offset: 0},
		},
		8: []TLV{
			TLV{Name: "MACSA_MACDA_NORMALIZED", Size: 1, Offset: 42},
			TLV{Name: "SIP_DIP_NORMALIZED", Size: 1, Offset: 41},
			TLV{Name: "DEST_IS_LOCAL", Size: 1, Offset: 40},
			TLV{Name: "FORWARDING_FIELD", Size: 14, Offset: 26},
			TLV{Name: "SRC_CLASSID", Size: 10, Offset: 16},
			TLV{Name: "TOS_FN", Size: 8, Offset: 8},
			TLV{Name: "IP_PROTOCOL/LAST_NH", Size: 8, Offset: 0},
		},
		9: []TLV{
			TLV{Name: "UDF1_VALID", Size: 2, Offset: 34},
			TLV{Name: "RAL_GAL", Size: 2, Offset: 32},
			TLV{Name: "UDF1", Size: 32, Offset: 0},
		},
		10: []TLV{
			TLV{Name: "DGLP", Size: 16, Offset: 21},
			TLV{Name: "D_TYPE", Size: 3, Offset: 18},
			TLV{Name: "D_FIELD", Size: 18, Offset: 0},
		},
		11: []TLV{
			TLV{Name: "CNG", Size: 2, Offset: 37},
			TLV{Name: "INT_PRI", Size: 4, Offset: 33},
			TLV{Name: "SVP_VALID", Size: 1, Offset: 32},
			TLV{Name: "SVP", Size: 16, Offset: 16},
			TLV{Name: "SGLP", Size: 16, Offset: 0},
		},
		12: []TLV{
			TLV{Name: "DEST_CLASSID", Size: 10, Offset: 38},
			TLV{Name: "SVP", Size: 16, Offset: 22},
			TLV{Name: "SRC_CLASSID", Size: 10, Offset: 12},
			TLV{Name: "OVID", Size: 12, Offset: 0},
		},
		13: []TLV{
			TLV{Name: "SW_VALID", Size: 1, Offset: 36},
			TLV{Name: "LABEL_ACTION", Size: 3, Offset: 33},
			TLV{Name: "AUX_TAG_VALID_1", Size: 1, Offset: 32},
			TLV{Name: "AUX_TAG_1", Size: 32, Offset: 0},
		},
	},
	FP2: map[int][]TLV{
		0: []TLV{
			TLV{Name: "SIP", Size: 32, Offset: 96},
			TLV{Name: "DIP", Size: 32, Offset: 64},
			TLV{Name: "IP_PROTOCOL/LAST_NH", Size: 8, Offset: 56},
			TLV{Name: "L4_SRC", Size: 16, Offset: 40},
			TLV{Name: "L4_DEST", Size: 16, Offset: 24},
			TLV{Name: "TOS_FN", Size: 8, Offset: 16},
			TLV{Name: "IP_FRAG_INFO", Size: 2, Offset: 14},
			TLV{Name: "TCP_FN", Size: 6, Offset: 8},
			TLV{Name: "TTL_FN", Size: 8, Offset: 0},
		},
		1: []TLV{
			TLV{Name: "SIP", Size: 32, Offset: 96},
			TLV{Name: "DIP", Size: 32, Offset: 64},
			TLV{Name: "IP_PROTOCOL/LAST_NH", Size: 8, Offset: 56},
			TLV{Name: "L4_SRC", Size: 16, Offset: 40},
			TLV{Name: "L4_DEST", Size: 16, Offset: 24},
			TLV{Name: "TOS_FN", Size: 8, Offset: 16},
			TLV{Name: "IP_FRAG_INFO", Size: 2, Offset: 14},
			TLV{Name: "TCP_FN", Size: 6, Offset: 8},
			TLV{Name: "TTL_FN", Size: 8, Offset: 0},
		},
		2: []TLV{
			TLV{Name: "IPV6_SIP", Size: 128, Offset: 0},
		},
		3: []TLV{
			TLV{Name: "IPV6_SIP", Size: 128, Offset: 0},
		},
		4: []TLV{
			TLV{Name: "IPV6_DIP_UPPER64", Size: 64, Offset: 64},
			TLV{Name: "IP_PROTOCOL/LAST_NH", Size: 8, Offset: 42},
			TLV{Name: "TOS_FN", Size: 8, Offset: 34},
			TLV{Name: "FL", Size: 20, Offset: 14},
			TLV{Name: "TCP_FN", Size: 6, Offset: 8},
			TLV{Name: "TTL_FN", Size: 8, Offset: 0},
		},
		5: []TLV{
			TLV{Name: "MACDA", Size: 48, Offset: 80},
			TLV{Name: "MACSA", Size: 48, Offset: 32},
			TLV{Name: "ETHERTYPE", Size: 16, Offset: 16},
			TLV{Name: "OTAG", Size: 16, Offset: 0},
		},
		6: []TLV{
			TLV{Name: "SIP", Size: 32, Offset: 80},
			TLV{Name: "MACSA", Size: 48, Offset: 32},
			TLV{Name: "ETHERTYPE", Size: 16, Offset: 16},
			TLV{Name: "OTAG", Size: 16, Offset: 0},
		},
		7: []TLV{
			TLV{Name: "MACDA", Size: 48, Offset: 80},
			TLV{Name: "DIP", Size: 32, Offset: 48},
			TLV{Name: "IP_PROTOCOL/LAST_NH", Size: 8, Offset: 40},
			TLV{Name: "TTL_FN", Size: 8, Offset: 32},
			TLV{Name: "ETHERTYPE", Size: 16, Offset: 16},
			TLV{Name: "OTAG", Size: 16, Offset: 0},
		},

		8: []TLV{
			TLV{Name: "UDF1", Size: 128, Offset: 0},
		},

		9: []TLV{
			TLV{Name: "UDF2", Size: 128, Offset: 0},
		},

		10: []TLV{
			TLV{Name: "IPV6_DIP_UPPER64", Size: 64, Offset: 64},
			TLV{Name: "IPV6_SIP_UPPER64", Size: 64, Offset: 0},
		},

		11: []TLV{
			TLV{Name: "MACDA", Size: 48, Offset: 80},
			TLV{Name: "MACSA", Size: 48, Offset: 32},
			TLV{Name: "DIPV6_DIP_UPPER32", Size: 32, Offset: 0},
		},
	},
	FP3: map[int][]TLV{
		0: []TLV{
			TLV{Name: "L3_IIF", Size: 16, Offset: 24},
			TLV{Name: "FORWARDING_FIELD", Size: 14, Offset: 10},
			TLV{Name: "SRC_DEST_CLASSID", Size: 10, Offset: 0},
		},
		1: []TLV{
			TLV{Name: "DEST_CLASSID", Size: 10, Offset: 36},
			TLV{Name: "OVID", Size: 12, Offset: 24},
			TLV{Name: "FORWARDING_FIELD", Size: 14, Offset: 10},
			TLV{Name: "SRC_CLASSID", Size: 10, Offset: 0},
		},
		2: []TLV{
			TLV{Name: "MH_OPCODE", Size: 3, Offset: 38},
			TLV{Name: "SVP_VALID", Size: 1, Offset: 37},
			TLV{Name: "S_FIELD", Size: 16, Offset: 21},
			TLV{Name: "D_TYPE", Size: 3, Offset: 18},
			TLV{Name: "D_FIELD", Size: 18, Offset: 0},
		},
		3: []TLV{
			TLV{Name: "FC_HDR_ENCODE_2", Size: 3, Offset: 37},
			TLV{Name: "FC_HDR_ENCODE_1", Size: 3, Offset: 34},
			TLV{Name: "IP_INFO", Size: 3, Offset: 31},
			TLV{Name: "PKT_RES", Size: 6, Offset: 25},
			TLV{Name: "MH_OPCODE", Size: 3, Offset: 22},
			TLV{Name: "TAG_STATUS", Size: 2, Offset: 20},
			TLV{Name: "PKT_FORMAT", Size: 4, Offset: 16},
			TLV{Name: "OTAG", Size: 16, Offset: 0},
		},

		4: []TLV{
			TLV{Name: "DEST_IS_LOCAL", Size: 1, Offset: 38},
			TLV{Name: "CNG", Size: 2, Offset: 36},
			TLV{Name: "INT_PRI", Size: 4, Offset: 32},
			TLV{Name: "ITAG", Size: 16, Offset: 16},
			TLV{Name: "OTAG", Size: 16, Offset: 0},
		},

		5: []TLV{
			TLV{Name: "FC_HDR_ENCODE_2", Size: 3, Offset: 35},
			TLV{Name: "FC_HDR_ENCODE_1", Size: 3, Offset: 32},
			TLV{Name: "EHTERTYPE", Size: 16, Offset: 16},
			TLV{Name: "OVID", Size: 16, Offset: 0},
		},

		6: []TLV{
			TLV{Name: "ITAG", Size: 16, Offset: 26},
			TLV{Name: "PKT_RES", Size: 6, Offset: 20},
			TLV{Name: "LOOKUP_STATUS", Size: 20, Offset: 0},
		},

		7: []TLV{
			TLV{Name: "INTERFACE_CLASSID", Size: 12, Offset: 24},
			TLV{Name: "RANGE_CHECK_RESULT", Size: 24, Offset: 0},
		},

		8: []TLV{
			TLV{Name: "OUTER_TPID_ENCODE", Size: 2, Offset: 28},
			TLV{Name: "INNER_TPID_ENCODE", Size: 2, Offset: 26},
			TLV{Name: "TAG_STATUS", Size: 2, Offset: 24},
			TLV{Name: "PKT_FORMAT", Size: 4, Offset: 20},
			TLV{Name: "IPV6_FL", Size: 20, Offset: 0},
		},

		9: []TLV{
			TLV{Name: "UDF_CHUNK_VALID_5_4", Size: 2, Offset: 34},
			TLV{Name: "RAL_GAL", Size: 2, Offset: 32},
			TLV{Name: "UDF1_95_64", Size: 32, Offset: 0},
		},

		10: []TLV{
			TLV{Name: "DGLP", Size: 16, Offset: 21},
			TLV{Name: "D_TYPE", Size: 3, Offset: 18},
			TLV{Name: "D_FIELD", Size: 18, Offset: 0},
		},

		11: []TLV{
			TLV{Name: "SVP_VALID", Size: 1, Offset: 32},
			TLV{Name: "SVP", Size: 16, Offset: 16},
			TLV{Name: "SGLP", Size: 18, Offset: 0},
		},

		12: []TLV{
			TLV{Name: "CW_VALID", Size: 1, Offset: 36},
			TLV{Name: "LABEL_ACTION", Size: 3, Offset: 33},
			TLV{Name: "AUX_TAG_VALID_2", Size: 1, Offset: 32},
			TLV{Name: "AUX_TAG_2", Size: 32, Offset: 0},
		},

		13: []TLV{
			TLV{Name: "MACSA_MACDA_NORMALIZED", Size: 1, Offset: 34},
			TLV{Name: "SIP_DIP_NORMALIZED", Size: 1, Offset: 33},
			TLV{Name: "DEST_IS_LOCAL", Size: 1, Offset: 32},
			TLV{Name: "IP_FIRST_PROTOCOL", Size: 8, Offset: 24},
			TLV{Name: "IPV6_FIRST_SUB_CODE", Size: 8, Offset: 16},
			TLV{Name: "IPV6_SECOND_NH", Size: 8, Offset: 8},
			TLV{Name: "TOS_FN", Size: 8, Offset: 0},
		},
	},
	FIXED: map[int][]TLV{
		0: []TLV{
			TLV{Name: "MY_STATION_HIT", Size: 1, Offset: 20},
			TLV{Name: "MIRROR_ONLY", Size: 1, Offset: 19},
			TLV{Name: "DROP", Size: 2, Offset: 17},
			TLV{Name: "TUNNEL_TYPE", Size: 5, Offset: 12},
			TLV{Name: "L3_ROUTABLE", Size: 1, Offset: 11},
			TLV{Name: "L4_VALID", Size: 1, Offset: 10},
			TLV{Name: "L3_TYPES", Size: 4, Offset: 5},
			TLV{Name: "SVP_OR_L3IIF", Size: 1, Offset: 4},
			TLV{Name: "FORWARDING_TYPE", Size: 3, Offset: 1},
			TLV{Name: "HIGIG", Size: 1, Offset: 0},
		},
	},

	FP4: map[int][]TLV{
		0: []TLV{
			TLV{Name: "PORT_FIELD_SEL_TABLE.INDEX", Size: 7, Offset: 0},
		},
	},
}

var BCM56540ICAPFieldSelector_Double = ICAPFieldSelector{
	FP0: map[int][]TLV{
		0: []TLV{
			TLV{Name: "IP_FRAG_INFO", Size: 2, Offset: 35},
			TLV{Name: "CNG", Size: 2, Offset: 33},
			TLV{Name: "SVP_VALID", Size: 1, Offset: 32},
			TLV{Name: "S_FIELD", Size: 16, Offset: 16},
			TLV{Name: "DGLP", Size: 16, Offset: 0},
		},
	},

	FP1: map[int][]TLV{
		0: []TLV{
			TLV{Name: "INT_PRI", Size: 4, Offset: 61},
			TLV{Name: "D_TYPE", Size: 3, Offset: 58},
			TLV{Name: "D_FIELD", Size: 18, Offset: 40},
			TLV{Name: "L4_SRC", Size: 16, Offset: 24},
			TLV{Name: "L4_DST", Size: 16, Offset: 8},
			TLV{Name: "TTL_FN1", Size: 8, Offset: 0},
		},
		1: []TLV{
			TLV{Name: "INT_PRI", Size: 4, Offset: 56},
			TLV{Name: "CW_VALID", Size: 1, Offset: 55},
			TLV{Name: "CW", Size: 30, Offset: 23},
			TLV{Name: "LABEL_ACTION", Size: 3, Offset: 20},
			TLV{Name: "LABEL_ID", Size: 20, Offset: 0},
		},
	},

	FP2: map[int][]TLV{
		0: []TLV{
			TLV{Name: "IPV4_SIP", Size: 32, Offset: 96},
			TLV{Name: "IPV4_DIP", Size: 32, Offset: 64},
			TLV{Name: "INTERFACE_CLASSID", Size: 12, Offset: 52},
			TLV{Name: "RANGE_CHECK_RESULT", Size: 24, Offset: 28},
			TLV{Name: "PKT_RES", Size: 6, Offset: 22},
			TLV{Name: "LOOKUP_STATUS", Size: 20, Offset: 2},
			TLV{Name: "ZEROS", Size: 2, Offset: 0},
		},
		1: []TLV{
			TLV{Name: "IPV6_SIP", Size: 128, Offset: 0},
		},

		2: []TLV{
			TLV{Name: "IPV6_DIP", Size: 128, Offset: 0},
		},

		3: []TLV{
			TLV{Name: "UDF2", Size: 128, Offset: 0},
		},
	},

	FP3: map[int][]TLV{
		0: []TLV{
			TLV{Name: "FIRST_NH", Size: 8, Offset: 24},
			TLV{Name: "FIRST_SUB_CODE", Size: 8, Offset: 16},
			TLV{Name: "IP_PROTOCOL_LAST_NH", Size: 8, Offset: 8},
			TLV{Name: "TOS_FN", Size: 8, Offset: 0},
		},
	},
	FP4: map[int][]TLV{
		0: []TLV{
			TLV{Name: "PORT_FIELD_SEL_TABLE.INDEX", Size: 7, Offset: 0},
		},
		1: []TLV{
			TLV{Name: "TCP_FN", Size: 6, Offset: 0},
		},
	},
}

//Even Slic Key
var BCM56850ICAPFieldSelector_single = ICAPFieldSelector{
	FP1: map[int][]TLV{
		0: []TLV{
			TLV{Name: "SVP", Size: 16, Offset: 33},
			TLV{Name: "FORWARDING_FIELD", Size: 13, Offset: 20},
			TLV{Name: "SRC_CLASSID", Size: 10, Offset: 10},
			TLV{Name: "DEST_CLASSID", Size: 10, Offset: 10},
		},
		1: []TLV{
			TLV{Name: "MH_OPCODE", Size: 3, Offset: 37},
			TLV{Name: "S_FIELD", Size: 16, Offset: 21},
			TLV{Name: "D_TYPE", Size: 3, Offset: 18},
			TLV{Name: "D_FIELD", Size: 18, Offset: 0},
		},
		2: []TLV{
			TLV{Name: "SRC_CLASSID", Size: 10, Offset: 34},
			TLV{Name: "DEST_CLASSID", Size: 10, Offset: 24},
			TLV{Name: "MH_OPCODE", Size: 3, Offset: 21},
			TLV{Name: "D_TYPE", Size: 3, Offset: 18},
			TLV{Name: "D_FIELD", Size: 18, Offset: 0},
		},
		3: []TLV{
			TLV{Name: "OUTER_TPID_ENCODE", Size: 2, Offset: 32},
			TLV{Name: "ITAG", Size: 16, Offset: 16},
			TLV{Name: "OTAG", Size: 16, Offset: 0},
		},
		4: []TLV{
			TLV{Name: "FC_HDR_ENCODE_2", Size: 3, Offset: 35},
			TLV{Name: "FC_HDR_ENCODE_1", Size: 3, Offset: 32},
			TLV{Name: "ETHERTYPE", Size: 16, Offset: 16},
			TLV{Name: "OTAG", Size: 16, Offset: 0},
		},
		5: []TLV{
			TLV{Name: "INNER_TPID_ENCODE", Size: 2, Offset: 42},
			TLV{Name: "ITAG", Size: 16, Offset: 26},
			TLV{Name: "PKT_RES", Size: 6, Offset: 20},
			TLV{Name: "LOOKUP_STATUS", Size: 20, Offset: 0},
		},
		6: []TLV{
			TLV{Name: "ETHERTYPE", Size: 6, Offset: 33},
			TLV{Name: "IP_FRAG_INFO", Size: 2, Offset: 31},
			TLV{Name: "PKT_RES", Size: 6, Offset: 25},
			TLV{Name: "MH_OPCODE", Size: 3, Offset: 22},
			TLV{Name: "TAG_STATUS", Size: 2, Offset: 20},
			TLV{Name: "PKT_FORMAT", Size: 4, Offset: 16},
			TLV{Name: "OTAG", Size: 16, Offset: 0},
		},
		7: []TLV{
			TLV{Name: "RAL_GAL", Size: 3, Offset: 39},
			TLV{Name: "FORWARDING_FIELD", Size: 13, Offset: 26},
			TLV{Name: "SRC_CLASSID", Size: 10, Offset: 16},
			TLV{Name: "OTAG", Size: 16, Offset: 0},
		},
		8: []TLV{
			TLV{Name: "DEST_IS_LOCAL", Size: 1, Offset: 47},
			TLV{Name: "ICMP_ERROR", Size: 1, Offset: 46},
			TLV{Name: "NAT_NEEDED", Size: 1, Offset: 45},
			TLV{Name: "NAT_DST_REALM_ID", Size: 2, Offset: 43},
			TLV{Name: "NAT_SRC_REALM_ID", Size: 2, Offset: 41},
			TLV{Name: "MACSA_MACDA_NORMALIZED", Size: 1, Offset: 40},
			TLV{Name: "SIP_DIP_NORMALIZED", Size: 1, Offset: 39},
			TLV{Name: "FORWARDING_FIELD", Size: 14, Offset: 26},
			TLV{Name: "SRC_CLASSID", Size: 10, Offset: 16},
			TLV{Name: "TOS_FN", Size: 8, Offset: 8},
			TLV{Name: "IP_PROTOCOL/LAST_NH", Size: 8, Offset: 0},
		},
		9: []TLV{
			TLV{Name: "L3_IIF", Size: 13, Offset: 36},
			TLV{Name: "UDF_CHUNK_VALID_1_0", Size: 2, Offset: 34},
			TLV{Name: "RAL_GAL", Size: 2, Offset: 32},
			TLV{Name: "UDF1_31_0", Size: 32, Offset: 0},
		},
		10: []TLV{
			TLV{Name: "DGLP", Size: 16, Offset: 21},
			TLV{Name: "D_TYPE", Size: 3, Offset: 18},
			TLV{Name: "D_FIELD", Size: 18, Offset: 0},
		},
		11: []TLV{
			TLV{Name: "CNG", Size: 2, Offset: 36},
			TLV{Name: "INT_PRI", Size: 4, Offset: 32},
			TLV{Name: "SVP", Size: 16, Offset: 16},
			TLV{Name: "SGLP", Size: 16, Offset: 0},
		},
		12: []TLV{
			TLV{Name: "DEST_CLASSID", Size: 10, Offset: 38},
			TLV{Name: "SVP", Size: 16, Offset: 22},
			TLV{Name: "SRC_CLASSID", Size: 10, Offset: 12},
			TLV{Name: "OVID", Size: 12, Offset: 0},
		},
		13: []TLV{
			TLV{Name: "CNG", Size: 2, Offset: 41},
			TLV{Name: "INT_PRI", Size: 4, Offset: 37},
			TLV{Name: "CW_VALID", Size: 1, Offset: 36},
			TLV{Name: "LABEL_ACTION", Size: 3, Offset: 33},
			TLV{Name: "AUX_TAG_VALID_1", Size: 1, Offset: 32},
			TLV{Name: "AUX_TAG_1", Size: 32, Offset: 0},
			TLV{Name: "LABLE_ID", Size: 20, Offset: 12},
			TLV{Name: "LABLE_EXP", Size: 3, Offset: 9},
			TLV{Name: "LABLE_BOS", Size: 1, Offset: 8},
			TLV{Name: "LABLE_TTL", Size: 9, Offset: 0},
		},

		14: []TLV{
			TLV{Name: "MACSA_MACDA_NORMALIZED", Size: 1, Offset: 48},
			TLV{Name: "MACSA", Size: 48, Offset: 0},
		},

		15: []TLV{
			TLV{Name: "IP_GRAG_INFO", Size: 2, Offset: 47},
			TLV{Name: "TCP_FN", Size: 6, Offset: 41},
			TLV{Name: "L4_NORMALIZED", Size: 1, Offset: 40},
			TLV{Name: "L4_SRC", Size: 16, Offset: 24},
			TLV{Name: "L4_DST", Size: 16, Offset: 8},
			TLV{Name: "IP_PROTOCOL/LAST_NH", Size: 8, Offset: 0},
		},

		16: []TLV{
			TLV{Name: "FCOE_ZONE_CHECK_STATUS", Size: 2, Offset: 47},
			TLV{Name: "FCOE_SRC_BIND_CHECK_STATUS", Size: 1, Offset: 46},
			TLV{Name: "FCOE_SRC_FPMA_PREFIX_CHECK_STATUS", Size: 1, Offset: 45},
			TLV{Name: "IFP_VSAN_PRI", Size: 3, Offset: 42},
			TLV{Name: "IFP_VSAN_ID", Size: 12, Offset: 30},
			TLV{Name: "FCOE_VFT_HOP_COUNT_FN", Size: 5, Offset: 25},
			TLV{Name: "FCOE_VFT_VERSION", Size: 2, Offset: 23},
			TLV{Name: "FCOE_STD_R_CTL", Size: 8, Offset: 15},
			TLV{Name: "FC_HDR_ENCODE_1", Size: 3, Offset: 12},
			TLV{Name: "FC_HDR_ENCODE_2", Size: 3, Offset: 9},
			TLV{Name: "FCOE_SOF", Size: 8, Offset: 1},
			TLV{Name: "FCOE_VERSION_IS_ZERO", Size: 1, Offset: 0},
		},
	},
	FP2: map[int][]TLV{
		0: []TLV{
			TLV{Name: "SIP", Size: 32, Offset: 96},
			TLV{Name: "DIP", Size: 32, Offset: 64},
			TLV{Name: "IP_PROTOCOL/LAST_NH", Size: 8, Offset: 56},
			TLV{Name: "L4_SRC", Size: 16, Offset: 40},
			TLV{Name: "L4_DEST", Size: 16, Offset: 24},
			TLV{Name: "TOS_FN", Size: 8, Offset: 16},
			TLV{Name: "IPFLAG", Size: 2, Offset: 14},
			TLV{Name: "TCP_FN", Size: 6, Offset: 8},
			TLV{Name: "TTL_FN", Size: 8, Offset: 0},
		},
		1: []TLV{
			TLV{Name: "SIP", Size: 32, Offset: 96},
			TLV{Name: "DIP", Size: 32, Offset: 64},
			TLV{Name: "IP_PROTOCOL/LAST_NH", Size: 8, Offset: 56},
			TLV{Name: "L4_SRC", Size: 16, Offset: 40},
			TLV{Name: "L4_DEST", Size: 16, Offset: 24},
			TLV{Name: "TOS_FN", Size: 8, Offset: 16},
			TLV{Name: "IP_FRAG_INFO", Size: 2, Offset: 14},
			TLV{Name: "TCP_FN", Size: 6, Offset: 8},
			TLV{Name: "TTL_FN", Size: 8, Offset: 0},
		},
		2: []TLV{
			TLV{Name: "IPV6_SIP", Size: 128, Offset: 0},
		},
		3: []TLV{
			TLV{Name: "IPV6_SIP", Size: 128, Offset: 0},
		},
		4: []TLV{
			TLV{Name: "IPV6_DIP_UPPER64", Size: 64, Offset: 64},
			TLV{Name: "RSVD", Size: 1, Offset: 63},
			TLV{Name: "L3_IIF", Size: 13, Offset: 50},
			TLV{Name: "IP_PROTOCOL/LAST_NH", Size: 8, Offset: 42},
			TLV{Name: "TOS_FN", Size: 8, Offset: 34},
			TLV{Name: "IPV6_FL", Size: 20, Offset: 14},
			TLV{Name: "TCP_FN", Size: 6, Offset: 8},
			TLV{Name: "TTL_FN", Size: 8, Offset: 0},
		},
		5: []TLV{
			TLV{Name: "MACDA", Size: 48, Offset: 80},
			TLV{Name: "MACSA", Size: 48, Offset: 32},
			TLV{Name: "ETHERTYPE", Size: 16, Offset: 16},
			TLV{Name: "OTAG", Size: 16, Offset: 0},
		},
		6: []TLV{
			TLV{Name: "SIP", Size: 32, Offset: 80},
			TLV{Name: "MACSA", Size: 48, Offset: 32},
			TLV{Name: "ETHERTYPE", Size: 16, Offset: 16},
			TLV{Name: "OTAG", Size: 16, Offset: 0},
		},
		7: []TLV{
			TLV{Name: "MACDA", Size: 48, Offset: 80},
			TLV{Name: "DIP", Size: 32, Offset: 48},
			TLV{Name: "IP_PROTOCOL/LAST_NH", Size: 8, Offset: 40},
			TLV{Name: "TTL_FN", Size: 8, Offset: 32},
			TLV{Name: "ETHERTYPE", Size: 16, Offset: 16},
			TLV{Name: "OTAG", Size: 16, Offset: 0},
		},

		8: []TLV{
			TLV{Name: "UDF1", Size: 128, Offset: 0},
		},

		9: []TLV{
			TLV{Name: "UDF2", Size: 128, Offset: 0},
		},

		10: []TLV{
			TLV{Name: "IPV6_DIP_UPPER64", Size: 64, Offset: 64},
			TLV{Name: "IPV6_SIP_UPPER64", Size: 64, Offset: 0},
		},

		11: []TLV{
			TLV{Name: "MACDA", Size: 48, Offset: 80},
			TLV{Name: "MACSA", Size: 48, Offset: 32},
			TLV{Name: "DIPV6_DIP_UPPER32", Size: 32, Offset: 0},
		},

		12: []TLV{
			TLV{Name: "FCOE_STD_S_ID", Size: 24, Offset: 104},
			TLV{Name: "FCOE_STD_D_ID", Size: 24, Offset: 80},
			TLV{Name: "FCOE_STD_TYPE", Size: 6, Offset: 72},
			TLV{Name: "FCOE_STD_F_CTL", Size: 24, Offset: 48},
			TLV{Name: "FCOE_STD_CS_CTL", Size: 8, Offset: 40},
			TLV{Name: "FCOE_STD_CS_CTL", Size: 8, Offset: 40},
			TLV{Name: "FCOE_STD_OX_ID", Size: 16, Offset: 16},
			TLV{Name: "FCOE_STD_RX_ID", Size: 16, Offset: 0},
		},
	},
	FP3: map[int][]TLV{
		0: []TLV{
			TLV{Name: "SVP", Size: 16, Offset: 33},
			TLV{Name: "FORWARDING_FIELD", Size: 13, Offset: 20},
			TLV{Name: "SRC_CLASSID", Size: 10, Offset: 10},
			TLV{Name: "DEST_CLASSID", Size: 10, Offset: 0},
		},
		1: []TLV{
			TLV{Name: "TOS_FN_LOW", Size: 4, Offset: 45},
			TLV{Name: "DEST_CLASSID", Size: 10, Offset: 35},
			TLV{Name: "OVID", Size: 12, Offset: 23},
			TLV{Name: "FORWARDING_FIELD", Size: 13, Offset: 10},
			TLV{Name: "SRC_CLASSID", Size: 10, Offset: 0},
		},
		2: []TLV{
			TLV{Name: "MH_OPCODE", Size: 3, Offset: 37},
			TLV{Name: "S_FIELD", Size: 16, Offset: 21},
			TLV{Name: "D_TYPE", Size: 3, Offset: 18},
			TLV{Name: "D_FIELD", Size: 18, Offset: 0},
		},
		3: []TLV{
			TLV{Name: "ETHERTYPE", Size: 16, Offset: 33},
			TLV{Name: "IP_FRAG_INFO", Size: 2, Offset: 31},
			TLV{Name: "PKT_RES", Size: 6, Offset: 25},
			TLV{Name: "MH_OPCODE", Size: 3, Offset: 22},
			TLV{Name: "TAG_STATUS", Size: 2, Offset: 20},
			TLV{Name: "PKT_FORMAT", Size: 4, Offset: 16},
			TLV{Name: "OTAG", Size: 16, Offset: 0},
		},

		4: []TLV{
			TLV{Name: "DEST_IS_LOCAL", Size: 1, Offset: 38},
			TLV{Name: "CNG", Size: 2, Offset: 36},
			TLV{Name: "INT_PRI", Size: 4, Offset: 32},
			TLV{Name: "ITAG", Size: 16, Offset: 16},
			TLV{Name: "OTAG", Size: 16, Offset: 0},
		},

		5: []TLV{
			TLV{Name: "FC_HDR_ENCODE_2", Size: 3, Offset: 35},
			TLV{Name: "FC_HDR_ENCODE_1", Size: 3, Offset: 32},
			TLV{Name: "EHTERTYPE", Size: 16, Offset: 16},
			TLV{Name: "OVID", Size: 16, Offset: 0},
		},

		6: []TLV{
			TLV{Name: "ITAG", Size: 16, Offset: 26},
			TLV{Name: "PKT_RES", Size: 6, Offset: 20},
			TLV{Name: "LOOKUP_STATUS", Size: 20, Offset: 0},
		},

		7: []TLV{
			TLV{Name: "INTERFACE_CLASSID", Size: 12, Offset: 24},
			TLV{Name: "RANGE_CHECK_RESULT", Size: 24, Offset: 0},
		},

		8: []TLV{
			TLV{Name: "OUTER_TPID_ENCODE", Size: 2, Offset: 28},
			TLV{Name: "INNER_TPID_ENCODE", Size: 2, Offset: 26},
			TLV{Name: "TAG_STATUS", Size: 2, Offset: 24},
			TLV{Name: "PKT_FORMAT", Size: 4, Offset: 20},
			TLV{Name: "IPV6_FL", Size: 20, Offset: 0},
		},

		9: []TLV{
			TLV{Name: "UDF_CHUNK_VALID_5_4", Size: 2, Offset: 34},
			TLV{Name: "RAL_GAL", Size: 2, Offset: 32},
			TLV{Name: "UDF1_95_64", Size: 32, Offset: 0},
		},

		10: []TLV{
			TLV{Name: "DGLP", Size: 16, Offset: 21},
			TLV{Name: "D_TYPE", Size: 3, Offset: 18},
			TLV{Name: "D_FIELD", Size: 18, Offset: 0},
		},

		11: []TLV{
			TLV{Name: "SVP", Size: 16, Offset: 16},
			TLV{Name: "SGLP", Size: 16, Offset: 0},
		},

		12: []TLV{
			TLV{Name: "CNG", Size: 2, Offset: 41},
			TLV{Name: "INT_PRI", Size: 4, Offset: 37},
			TLV{Name: "CW_VALID", Size: 1, Offset: 36},
			TLV{Name: "LABEL_ACTION", Size: 3, Offset: 33},
			TLV{Name: "AUX_TAG_VALID_2", Size: 1, Offset: 32},
			TLV{Name: "AUX_TAG_2", Size: 32, Offset: 0},
			TLV{Name: "LABEL_ID", Size: 20, Offset: 12},
			TLV{Name: "LABEL_EXP", Size: 3, Offset: 9},
			TLV{Name: "LABEL_BOS", Size: 1, Offset: 8},
			TLV{Name: "LABEL_TTL", Size: 8, Offset: 0},
		},

		13: []TLV{
			TLV{Name: "L3_IIF", Size: 13, Offset: 35},
			TLV{Name: "MACSA_MACDA_NORMALIZED", Size: 1, Offset: 34},
			TLV{Name: "SIP_DIP_NORMALIZED", Size: 1, Offset: 33},
			TLV{Name: "DEST_IS_LOCAL", Size: 1, Offset: 32},
			TLV{Name: "IP_FIRST_PROTOCOL", Size: 8, Offset: 24},
			TLV{Name: "IPV6_FIRST_SUB_CODE", Size: 8, Offset: 16},
			TLV{Name: "IPV6_SECOND_NH", Size: 8, Offset: 8},
			TLV{Name: "TOS_FN", Size: 8, Offset: 0},
		},

		14: []TLV{
			TLV{Name: "MACSA_MACDA_NORMALIZED", Size: 1, Offset: 48},
			TLV{Name: "MACDA", Size: 48, Offset: 0},
		},

		15: []TLV{
			TLV{Name: "OVID", Size: 12, Offset: 37},
			TLV{Name: "S_FIELD", Size: 16, Offset: 21},
			TLV{Name: "D_TYPE", Size: 3, Offset: 18},
			TLV{Name: "D_FIELD", Size: 18, Offset: 0},
		},

		16: []TLV{
			TLV{Name: "FCOE_ZONE_CHECK_STATUS", Size: 2, Offset: 47},
			TLV{Name: "FCOE_SRC_BIND_CHECK_STATUS", Size: 1, Offset: 46},
			TLV{Name: "FCOE_SRC_FPMA_PREFIX_CHECK_STATUS", Size: 1, Offset: 45},
			TLV{Name: "IFP_VSAN_PRI", Size: 3, Offset: 42},
			TLV{Name: "IFP_VSAN_ID", Size: 12, Offset: 30},
			TLV{Name: "FCOE_VFT_HOP_COUNT_FN", Size: 5, Offset: 25},
			TLV{Name: "FCOE_VFT_VERSION", Size: 2, Offset: 23},
			TLV{Name: "FCOE_STD_R_CTRL", Size: 8, Offset: 15},
			TLV{Name: "FCOE_HDR_ENCODE_1", Size: 3, Offset: 12},
			TLV{Name: "FCOE_HDR_ENCODE_2", Size: 3, Offset: 9},
			TLV{Name: "FCOE_SOF", Size: 8, Offset: 1},
			TLV{Name: "FCOE_VERSION_IS_ZERO", Size: 1, Offset: 0},
		},
	},
	FIXED: map[int][]TLV{
		0: []TLV{
			TLV{Name: "DROP", Size: 1, Offset: 21},
			TLV{Name: "IPV4_CHECKSUM_OK", Size: 1, Offset: 20},
			TLV{Name: "REP_COPY", Size: 1, Offset: 19},
			TLV{Name: "MIRROR_ONLY", Size: 1, Offset: 18},
			TLV{Name: "TUNNEL_TYPE", Size: 5, Offset: 13},
			TLV{Name: "L3_ROUTABLE", Size: 1, Offset: 12},
			TLV{Name: "L4_VALID", Size: 1, Offset: 11},
			TLV{Name: "L3_TYPE", Size: 5, Offset: 6},
			TLV{Name: "SVP_VALID", Size: 1, Offset: 5},
			TLV{Name: "FORWARDING_TYPE", Size: 3, Offset: 2},
			TLV{Name: "HIGIG", Size: 1, Offset: 1},
			TLV{Name: "MY_STATION_HIT", Size: 1, Offset: 0},
		},
	},

	FP4: map[int][]TLV{
		0: []TLV{
			TLV{Name: "PORT_FIELD_SEL_TABLE.INDEX", Size: 7, Offset: 0},
		},
	},
}

//Odd Slice key
var BCM56850ICAPFieldSelector_Double = ICAPFieldSelector{
	FP1: map[int][]TLV{
		0: []TLV{
			TLV{Name: "SVP", Size: 16, Offset: 33},
			TLV{Name: "FORWARDING_FIELD", Size: 13, Offset: 20},
			TLV{Name: "SRC_CLASSID", Size: 10, Offset: 10},
			TLV{Name: "DEST_CLASSID", Size: 10, Offset: 10},
		},
		1: []TLV{
			TLV{Name: "MH_OPCODE", Size: 3, Offset: 37},
			TLV{Name: "S_FIELD", Size: 16, Offset: 21},
			TLV{Name: "D_TYPE", Size: 3, Offset: 18},
			TLV{Name: "D_FIELD", Size: 18, Offset: 0},
		},
		2: []TLV{
			TLV{Name: "SRC_CLASSID", Size: 10, Offset: 34},
			TLV{Name: "DEST_CLASSID", Size: 10, Offset: 24},
			TLV{Name: "MH_OPCODE", Size: 3, Offset: 21},
			TLV{Name: "D_TYPE", Size: 3, Offset: 18},
			TLV{Name: "D_FIELD", Size: 18, Offset: 0},
		},
		3: []TLV{
			TLV{Name: "OUTER_TPID_ENCODE", Size: 2, Offset: 32},
			TLV{Name: "ITAG", Size: 16, Offset: 16},
			TLV{Name: "OTAG", Size: 16, Offset: 0},
		},
		4: []TLV{
			TLV{Name: "FC_HDR_ENCODE_2", Size: 3, Offset: 35},
			TLV{Name: "FC_HDR_ENCODE_1", Size: 3, Offset: 32},
			TLV{Name: "ETHERTYPE", Size: 16, Offset: 16},
			TLV{Name: "OTAG", Size: 16, Offset: 0},
		},
		5: []TLV{
			TLV{Name: "INNER_TPID_ENCODE", Size: 2, Offset: 42},
			TLV{Name: "ITAG", Size: 16, Offset: 26},
			TLV{Name: "PKT_RES", Size: 6, Offset: 20},
			TLV{Name: "LOOKUP_STATUS", Size: 20, Offset: 0},
		},
		6: []TLV{
			TLV{Name: "ETHERTYPE", Size: 6, Offset: 33},
			TLV{Name: "IP_FRAG_INFO", Size: 2, Offset: 31},
			TLV{Name: "PKT_RES", Size: 6, Offset: 25},
			TLV{Name: "MH_OPCODE", Size: 3, Offset: 22},
			TLV{Name: "TAG_STATUS", Size: 2, Offset: 20},
			TLV{Name: "PKT_FORMAT", Size: 4, Offset: 16},
			TLV{Name: "OTAG", Size: 16, Offset: 0},
		},
		7: []TLV{
			TLV{Name: "RAL_GAL", Size: 3, Offset: 39},
			TLV{Name: "FORWARDING_FIELD", Size: 13, Offset: 26},
			TLV{Name: "SRC_CLASSID", Size: 10, Offset: 16},
			TLV{Name: "OTAG", Size: 16, Offset: 0},
		},
		8: []TLV{
			TLV{Name: "DEST_IS_LOCAL", Size: 1, Offset: 47},
			TLV{Name: "ICMP_ERROR", Size: 1, Offset: 46},
			TLV{Name: "NAT_NEEDED", Size: 1, Offset: 45},
			TLV{Name: "NAT_DST_REALM_ID", Size: 2, Offset: 43},
			TLV{Name: "NAT_SRC_REALM_ID", Size: 2, Offset: 41},
			TLV{Name: "MACSA_MACDA_NORMALIZED", Size: 1, Offset: 40},
			TLV{Name: "SIP_DIP_NORMALIZED", Size: 1, Offset: 39},
			TLV{Name: "FORWARDING_FIELD", Size: 14, Offset: 26},
			TLV{Name: "SRC_CLASSID", Size: 10, Offset: 16},
			TLV{Name: "TOS_FN", Size: 8, Offset: 8},
			TLV{Name: "IP_PROTOCOL/LAST_NH", Size: 8, Offset: 0},
		},
		9: []TLV{
			TLV{Name: "L3_IIF", Size: 13, Offset: 36},
			TLV{Name: "UDF_CHUNK_VALID_1_0", Size: 2, Offset: 34},
			TLV{Name: "RAL_GAL", Size: 2, Offset: 32},
			TLV{Name: "UDF1_31_0", Size: 32, Offset: 0},
		},
		10: []TLV{
			TLV{Name: "DGLP", Size: 16, Offset: 21},
			TLV{Name: "D_TYPE", Size: 3, Offset: 18},
			TLV{Name: "D_FIELD", Size: 18, Offset: 0},
		},
		11: []TLV{
			TLV{Name: "CNG", Size: 2, Offset: 36},
			TLV{Name: "INT_PRI", Size: 4, Offset: 32},
			TLV{Name: "SVP", Size: 16, Offset: 16},
			TLV{Name: "SGLP", Size: 16, Offset: 0},
		},
		12: []TLV{
			TLV{Name: "DEST_CLASSID", Size: 10, Offset: 38},
			TLV{Name: "SVP", Size: 16, Offset: 22},
			TLV{Name: "SRC_CLASSID", Size: 10, Offset: 12},
			TLV{Name: "OVID", Size: 12, Offset: 0},
		},
		13: []TLV{
			TLV{Name: "CNG", Size: 2, Offset: 41},
			TLV{Name: "INT_PRI", Size: 4, Offset: 37},
			TLV{Name: "CW_VALID", Size: 1, Offset: 36},
			TLV{Name: "LABEL_ACTION", Size: 3, Offset: 33},
			TLV{Name: "AUX_TAG_VALID_1", Size: 1, Offset: 32},
			TLV{Name: "AUX_TAG_1", Size: 32, Offset: 0},
			TLV{Name: "LABLE_ID", Size: 20, Offset: 12},
			TLV{Name: "LABLE_EXP", Size: 3, Offset: 9},
			TLV{Name: "LABLE_BOS", Size: 1, Offset: 8},
			TLV{Name: "LABLE_TTL", Size: 9, Offset: 0},
		},

		14: []TLV{
			TLV{Name: "MACSA_MACDA_NORMALIZED", Size: 1, Offset: 48},
			TLV{Name: "MACSA", Size: 48, Offset: 0},
		},

		15: []TLV{
			TLV{Name: "IP_GRAG_INFO", Size: 2, Offset: 47},
			TLV{Name: "TCP_FN", Size: 6, Offset: 41},
			TLV{Name: "L4_NORMALIZED", Size: 1, Offset: 40},
			TLV{Name: "L4_SRC", Size: 16, Offset: 24},
			TLV{Name: "L4_DST", Size: 16, Offset: 8},
			TLV{Name: "IP_PROTOCOL/LAST_NH", Size: 8, Offset: 0},
		},

		16: []TLV{
			TLV{Name: "FCOE_ZONE_CHECK_STATUS", Size: 2, Offset: 47},
			TLV{Name: "FCOE_SRC_BIND_CHECK_STATUS", Size: 1, Offset: 46},
			TLV{Name: "FCOE_SRC_FPMA_PREFIX_CHECK_STATUS", Size: 1, Offset: 45},
			TLV{Name: "IFP_VSAN_PRI", Size: 3, Offset: 42},
			TLV{Name: "IFP_VSAN_ID", Size: 12, Offset: 30},
			TLV{Name: "FCOE_VFT_HOP_COUNT_FN", Size: 5, Offset: 25},
			TLV{Name: "FCOE_VFT_VERSION", Size: 2, Offset: 23},
			TLV{Name: "FCOE_STD_R_CTL", Size: 8, Offset: 15},
			TLV{Name: "FC_HDR_ENCODE_1", Size: 3, Offset: 12},
			TLV{Name: "FC_HDR_ENCODE_2", Size: 3, Offset: 9},
			TLV{Name: "FCOE_SOF", Size: 8, Offset: 1},
			TLV{Name: "FCOE_VERSION_IS_ZERO", Size: 1, Offset: 0},
		},
	},
	FP2: map[int][]TLV{
		0: []TLV{
			TLV{Name: "SIP", Size: 32, Offset: 96},
			TLV{Name: "DIP", Size: 32, Offset: 64},
			TLV{Name: "IP_PROTOCOL/LAST_NH", Size: 8, Offset: 56},
			TLV{Name: "L4_SRC", Size: 16, Offset: 40},
			TLV{Name: "L4_DEST", Size: 16, Offset: 24},
			TLV{Name: "TOS_FN", Size: 8, Offset: 16},
			TLV{Name: "IPFLAG", Size: 2, Offset: 14},
			TLV{Name: "TCP_FN", Size: 6, Offset: 8},
			TLV{Name: "TTL_FN", Size: 8, Offset: 0},
		},
		1: []TLV{
			TLV{Name: "SIP", Size: 32, Offset: 96},
			TLV{Name: "DIP", Size: 32, Offset: 64},
			TLV{Name: "IP_PROTOCOL/LAST_NH", Size: 8, Offset: 56},
			TLV{Name: "L4_SRC", Size: 16, Offset: 40},
			TLV{Name: "L4_DEST", Size: 16, Offset: 24},
			TLV{Name: "TOS_FN", Size: 8, Offset: 16},
			TLV{Name: "IP_FRAG_INFO", Size: 2, Offset: 14},
			TLV{Name: "TCP_FN", Size: 6, Offset: 8},
			TLV{Name: "TTL_FN", Size: 8, Offset: 0},
		},
		2: []TLV{
			TLV{Name: "IPV6_SIP", Size: 128, Offset: 0},
		},
		3: []TLV{
			TLV{Name: "IPV6_SIP", Size: 128, Offset: 0},
		},
		4: []TLV{
			TLV{Name: "IPV6_DIP_UPPER64", Size: 64, Offset: 64},
			TLV{Name: "RSVD", Size: 1, Offset: 63},
			TLV{Name: "L3_IIF", Size: 13, Offset: 50},
			TLV{Name: "IP_PROTOCOL/LAST_NH", Size: 8, Offset: 42},
			TLV{Name: "TOS_FN", Size: 8, Offset: 34},
			TLV{Name: "IPV6_FL", Size: 20, Offset: 14},
			TLV{Name: "TCP_FN", Size: 6, Offset: 8},
			TLV{Name: "TTL_FN", Size: 8, Offset: 0},
		},
		5: []TLV{
			TLV{Name: "MACDA", Size: 48, Offset: 80},
			TLV{Name: "MACSA", Size: 48, Offset: 32},
			TLV{Name: "ETHERTYPE", Size: 16, Offset: 16},
			TLV{Name: "OTAG", Size: 16, Offset: 0},
		},
		6: []TLV{
			TLV{Name: "SIP", Size: 32, Offset: 80},
			TLV{Name: "MACSA", Size: 48, Offset: 32},
			TLV{Name: "ETHERTYPE", Size: 16, Offset: 16},
			TLV{Name: "OTAG", Size: 16, Offset: 0},
		},
		7: []TLV{
			TLV{Name: "MACDA", Size: 48, Offset: 80},
			TLV{Name: "DIP", Size: 32, Offset: 48},
			TLV{Name: "IP_PROTOCOL/LAST_NH", Size: 8, Offset: 40},
			TLV{Name: "TTL_FN", Size: 8, Offset: 32},
			TLV{Name: "ETHERTYPE", Size: 16, Offset: 16},
			TLV{Name: "OTAG", Size: 16, Offset: 0},
		},

		8: []TLV{
			TLV{Name: "UDF1", Size: 128, Offset: 0},
		},

		9: []TLV{
			TLV{Name: "UDF2", Size: 128, Offset: 0},
		},

		10: []TLV{
			TLV{Name: "IPV6_DIP_UPPER64", Size: 64, Offset: 64},
			TLV{Name: "IPV6_SIP_UPPER64", Size: 64, Offset: 0},
		},

		11: []TLV{
			TLV{Name: "MACDA", Size: 48, Offset: 80},
			TLV{Name: "MACSA", Size: 48, Offset: 32},
			TLV{Name: "DIPV6_DIP_UPPER32", Size: 32, Offset: 0},
		},

		12: []TLV{
			TLV{Name: "FCOE_STD_S_ID", Size: 24, Offset: 104},
			TLV{Name: "FCOE_STD_D_ID", Size: 24, Offset: 80},
			TLV{Name: "FCOE_STD_TYPE", Size: 6, Offset: 72},
			TLV{Name: "FCOE_STD_F_CTL", Size: 24, Offset: 48},
			TLV{Name: "FCOE_STD_CS_CTL", Size: 8, Offset: 40},
			TLV{Name: "FCOE_STD_CS_CTL", Size: 8, Offset: 40},
			TLV{Name: "FCOE_STD_OX_ID", Size: 16, Offset: 16},
			TLV{Name: "FCOE_STD_RX_ID", Size: 16, Offset: 0},
		},
	},
	FP3: map[int][]TLV{
		0: []TLV{
			TLV{Name: "SVP", Size: 16, Offset: 33},
			TLV{Name: "FORWARDING_FIELD", Size: 13, Offset: 20},
			TLV{Name: "SRC_CLASSID", Size: 10, Offset: 10},
			TLV{Name: "DEST_CLASSID", Size: 10, Offset: 0},
		},
		1: []TLV{
			TLV{Name: "TOS_FN_LOW", Size: 4, Offset: 45},
			TLV{Name: "DEST_CLASSID", Size: 10, Offset: 35},
			TLV{Name: "OVID", Size: 12, Offset: 23},
			TLV{Name: "FORWARDING_FIELD", Size: 13, Offset: 10},
			TLV{Name: "SRC_CLASSID", Size: 10, Offset: 0},
		},
		2: []TLV{
			TLV{Name: "MH_OPCODE", Size: 3, Offset: 37},
			TLV{Name: "S_FIELD", Size: 16, Offset: 21},
			TLV{Name: "D_TYPE", Size: 3, Offset: 18},
			TLV{Name: "D_FIELD", Size: 18, Offset: 0},
		},
		3: []TLV{
			TLV{Name: "ETHERTYPE", Size: 16, Offset: 33},
			TLV{Name: "IP_FRAG_INFO", Size: 2, Offset: 31},
			TLV{Name: "PKT_RES", Size: 6, Offset: 25},
			TLV{Name: "MH_OPCODE", Size: 3, Offset: 22},
			TLV{Name: "TAG_STATUS", Size: 2, Offset: 20},
			TLV{Name: "PKT_FORMAT", Size: 4, Offset: 16},
			TLV{Name: "OTAG", Size: 16, Offset: 0},
		},

		4: []TLV{
			TLV{Name: "DEST_IS_LOCAL", Size: 1, Offset: 38},
			TLV{Name: "CNG", Size: 2, Offset: 36},
			TLV{Name: "INT_PRI", Size: 4, Offset: 32},
			TLV{Name: "ITAG", Size: 16, Offset: 16},
			TLV{Name: "OTAG", Size: 16, Offset: 0},
		},

		5: []TLV{
			TLV{Name: "FC_HDR_ENCODE_2", Size: 3, Offset: 35},
			TLV{Name: "FC_HDR_ENCODE_1", Size: 3, Offset: 32},
			TLV{Name: "EHTERTYPE", Size: 16, Offset: 16},
			TLV{Name: "OVID", Size: 16, Offset: 0},
		},

		6: []TLV{
			TLV{Name: "ITAG", Size: 16, Offset: 26},
			TLV{Name: "PKT_RES", Size: 6, Offset: 20},
			TLV{Name: "LOOKUP_STATUS", Size: 20, Offset: 0},
		},

		7: []TLV{
			TLV{Name: "INTERFACE_CLASSID", Size: 12, Offset: 24},
			TLV{Name: "RANGE_CHECK_RESULT", Size: 24, Offset: 0},
		},

		8: []TLV{
			TLV{Name: "OUTER_TPID_ENCODE", Size: 2, Offset: 28},
			TLV{Name: "INNER_TPID_ENCODE", Size: 2, Offset: 26},
			TLV{Name: "TAG_STATUS", Size: 2, Offset: 24},
			TLV{Name: "PKT_FORMAT", Size: 4, Offset: 20},
			TLV{Name: "IPV6_FL", Size: 20, Offset: 0},
		},

		9: []TLV{
			TLV{Name: "UDF_CHUNK_VALID_5_4", Size: 2, Offset: 34},
			TLV{Name: "RAL_GAL", Size: 2, Offset: 32},
			TLV{Name: "UDF1_95_64", Size: 32, Offset: 0},
		},

		10: []TLV{
			TLV{Name: "DGLP", Size: 16, Offset: 21},
			TLV{Name: "D_TYPE", Size: 3, Offset: 18},
			TLV{Name: "D_FIELD", Size: 18, Offset: 0},
		},

		11: []TLV{
			TLV{Name: "SVP", Size: 16, Offset: 16},
			TLV{Name: "SGLP", Size: 16, Offset: 0},
		},

		12: []TLV{
			TLV{Name: "CNG", Size: 2, Offset: 41},
			TLV{Name: "INT_PRI", Size: 4, Offset: 37},
			TLV{Name: "CW_VALID", Size: 1, Offset: 36},
			TLV{Name: "LABEL_ACTION", Size: 3, Offset: 33},
			TLV{Name: "AUX_TAG_VALID_2", Size: 1, Offset: 32},
			TLV{Name: "AUX_TAG_2", Size: 32, Offset: 0},
			TLV{Name: "LABEL_ID", Size: 20, Offset: 12},
			TLV{Name: "LABEL_EXP", Size: 3, Offset: 9},
			TLV{Name: "LABEL_BOS", Size: 1, Offset: 8},
			TLV{Name: "LABEL_TTL", Size: 8, Offset: 0},
		},

		13: []TLV{
			TLV{Name: "L3_IIF", Size: 13, Offset: 35},
			TLV{Name: "MACSA_MACDA_NORMALIZED", Size: 1, Offset: 34},
			TLV{Name: "SIP_DIP_NORMALIZED", Size: 1, Offset: 33},
			TLV{Name: "DEST_IS_LOCAL", Size: 1, Offset: 32},
			TLV{Name: "IP_FIRST_PROTOCOL", Size: 8, Offset: 24},
			TLV{Name: "IPV6_FIRST_SUB_CODE", Size: 8, Offset: 16},
			TLV{Name: "IPV6_SECOND_NH", Size: 8, Offset: 8},
			TLV{Name: "TOS_FN", Size: 8, Offset: 0},
		},

		14: []TLV{
			TLV{Name: "MACSA_MACDA_NORMALIZED", Size: 1, Offset: 48},
			TLV{Name: "MACDA", Size: 48, Offset: 0},
		},

		15: []TLV{
			TLV{Name: "OVID", Size: 12, Offset: 37},
			TLV{Name: "S_FIELD", Size: 16, Offset: 21},
			TLV{Name: "D_TYPE", Size: 3, Offset: 18},
			TLV{Name: "D_FIELD", Size: 18, Offset: 0},
		},

		16: []TLV{
			TLV{Name: "FCOE_ZONE_CHECK_STATUS", Size: 2, Offset: 47},
			TLV{Name: "FCOE_SRC_BIND_CHECK_STATUS", Size: 1, Offset: 46},
			TLV{Name: "FCOE_SRC_FPMA_PREFIX_CHECK_STATUS", Size: 1, Offset: 45},
			TLV{Name: "IFP_VSAN_PRI", Size: 3, Offset: 42},
			TLV{Name: "IFP_VSAN_ID", Size: 12, Offset: 30},
			TLV{Name: "FCOE_VFT_HOP_COUNT_FN", Size: 5, Offset: 25},
			TLV{Name: "FCOE_VFT_VERSION", Size: 2, Offset: 23},
			TLV{Name: "FCOE_STD_R_CTRL", Size: 8, Offset: 15},
			TLV{Name: "FCOE_HDR_ENCODE_1", Size: 3, Offset: 12},
			TLV{Name: "FCOE_HDR_ENCODE_2", Size: 3, Offset: 9},
			TLV{Name: "FCOE_SOF", Size: 8, Offset: 1},
			TLV{Name: "FCOE_VERSION_IS_ZERO", Size: 1, Offset: 0},
		},
	},
	FIXED: map[int][]TLV{
		0: []TLV{
			TLV{Name: "DROP", Size: 1, Offset: 21},
			TLV{Name: "IPV4_CHECKSUM_OK", Size: 1, Offset: 20},
			TLV{Name: "REP_COPY", Size: 1, Offset: 19},
			TLV{Name: "MIRROR_ONLY", Size: 1, Offset: 18},
			TLV{Name: "TUNNEL_TYPE", Size: 5, Offset: 13},
			TLV{Name: "L3_ROUTABLE", Size: 1, Offset: 12},
			TLV{Name: "L4_VALID", Size: 1, Offset: 11},
			TLV{Name: "L3_TYPE", Size: 5, Offset: 6},
			TLV{Name: "SVP_VALID", Size: 1, Offset: 5},
			TLV{Name: "FORWARDING_TYPE", Size: 3, Offset: 2},
			TLV{Name: "HIGIG", Size: 1, Offset: 1},
			TLV{Name: "MY_STATION_HIT", Size: 1, Offset: 0},
		},
	},

	FP4: map[int][]TLV{
		0: []TLV{
			TLV{Name: "PORT_FIELD_SEL_TABLE.INDEX", Size: 7, Offset: 0},
		},
	},

	PAIRING_IPBM_F0: map[int][]TLV{
		0: []TLV{
			TLV{Name: "L4_SRC", Size: 16, Offset: 22},
			TLV{Name: "L4_DST", Size: 16, Offset: 6},
			TLV{Name: "TCP_FN", Size: 6, Offset: 0},
		},
	},

	PAIRING_FIXED: map[int][]TLV{
		0: []TLV{
			TLV{Name: "TTL_FN1", Size: 8, Offset: 10},
			TLV{Name: "IP_PROTOCOL/LAST_NH", Size: 8, Offset: 2},
			TLV{Name: "IP_FRAG_INFO", Size: 2, Offset: 0},
		},
	},
}

var CTX = context.Background()

var IP = flag.String("ip", "10.71.20.191", "IP address of the remote device")
var Host = flag.String("hostname", "V8500_2", "Host name of the remote device")
var User = flag.String("username", "admin", "Username of the remote device")
var Password = flag.String("password", "", "Passwrod of the remote device")
var Phase = flag.String("p", "0", "rule stage(0/1)")
var SFU = flag.String("sfu", "A", "SFU (A/B)")

func AddRule(dev *rut.RUT, name string, flow string, action string) error {
	_, err := dev.RunCommands(CTX, []*command.Command{
		&command.Command{Mode: "config", CMD: " flow " + name + " create"},
		&command.Command{Mode: "config-flow", CMD: flow},
		&command.Command{Mode: "config-flow", CMD: " apply"},
		&command.Command{Mode: "config-flow", CMD: " exit"},
		&command.Command{Mode: "config", CMD: " policer " + name + " create"},
		&command.Command{Mode: "config-policer", CMD: " counter"},
		&command.Command{Mode: "config-policer", CMD: " apply"},
		&command.Command{Mode: "config-policer", CMD: " exit"},
		&command.Command{Mode: "config", CMD: " policy " + name + " create"},
		&command.Command{Mode: "config-policy", CMD: " include-flow " + name},
		&command.Command{Mode: "config-policy", CMD: " include-policer " + name},
		&command.Command{Mode: "config-policy", CMD: " action match " + action},
		&command.Command{Mode: "config-policy", CMD: " apply"},
		&command.Command{Mode: "config-policy", CMD: " exit"},
		&command.Command{Mode: "config", CMD: " interface gigabitethernet 8/2"},
		&command.Command{Mode: "config-if", CMD: " service-policy input " + name},
		&command.Command{Mode: "config-if", CMD: " exit"},
	})

	return err
}

func DelRule(dev *rut.RUT, name string) error {
	_, err := dev.RunCommands(CTX, []*command.Command{
		&command.Command{Mode: "config", CMD: " interface gigabitethernet 8/2"},
		&command.Command{Mode: "config-if", CMD: " no service-policy input " + name},
		&command.Command{Mode: "config-if", CMD: " exit"},
		&command.Command{Mode: "config", CMD: " no policy " + name},
		&command.Command{Mode: "config", CMD: " no policer " + name},
		&command.Command{Mode: "config", CMD: " no flow " + name},
	})

	return err
}

func dumpTableAndSaveToFile(dev *rut.RUT, name, start, end, file string) error {
	err := os.Remove(file)
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	data, err := dev.RunCommand(CTX, &command.Command{
		Mode: "config",
		CMD:  " do q sh -l",
	})
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("Cannot go to shell mode")
	}

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "shell",
		CMD:  " scontrol -f /proc/switch/ASIC/ctrl dump table 0 " + name + " " + start + " " + end,
	})
	if err != nil {
		fmt.Println(err)
	}

	util.SaveToFile(file, []byte(data))

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "shell",
		CMD:  " exit",
	})

	if err != nil {
		fmt.Println(err)
	}

	return nil
}

/*
VIRTUAL_SLICE_9_VIRTUAL_SLICE_GROUP_ENTRY_0=9,VIRTUAL_SLICE_9_PHYSICAL_SLICE_NUMBER_ENTRY_0=9
*/
var SliceGroupMapReg = regexp.MustCompile(`VIRTUAL_SLICE_(?P<sliceidx>0*[xX]*[0-9a-fA-F]+)_VIRTUAL_SLICE_GROUP_ENTRY_0=(?P<groupidx>0*[xX]*[0-9a-fA-F]+)`)
var VirtualSliceMapReg = regexp.MustCompile(`VIRTUAL_SLICE_(?P<sliceidx>0*[xX]*[0-9a-fA-F]+)_PHYSICAL_SLICE_NUMBER_ENTRY_0=(?P<physliceidx>0*[xX]*[0-9a-fA-F]+)`)

func DumpSliceInfo(dev *rut.RUT) {
	dumpTableAndSaveToFile(dev, "FP_SLICE_MAP", "0", "0", FP_SLICE_MAP_FILE("info"))
	table, _ := ioutil.ReadFile(FP_SLICE_MAP_FILE("info"))

	groups := SliceGroupMapReg.FindAllStringSubmatch(string(table), -1)
	for _, g := range groups {
		gidx, err := strconv.ParseInt(g[2], 0, 32)
		if err != nil {
			panic(err)
		}
		DB.GroupCount++
		sidx, err := strconv.ParseInt(g[1], 0, 32)
		if err != nil {
			panic(err)
		}
		DB.SliceGroupMap[sidx] = gidx
	}

	slices := VirtualSliceMapReg.FindAllStringSubmatch(string(table), -1)
	for _, s := range slices {
		pidx, err := strconv.ParseInt(s[2], 0, 32)
		if err != nil {
			panic(err)
		}
		DB.SliceCount++
		vidx, err := strconv.ParseInt(s[1], 0, 32)
		if err != nil {
			panic(err)
		}
		DB.SliceGroupMap[pidx] = vidx
	}

	for s := 0; s < DB.SliceCount; s++ {
		if s < 4 {
			DB.SliceEntryCountMap[int64(s)] = 512
			DB.EntryCount += 512
			if s == 0 {
				DB.SliceStartIndexMap[int64(s)] = 0
				DB.SliceEndIndexMap[int64(s)] = 512
			} else {
				DB.SliceStartIndexMap[int64(s)] = DB.SliceStartIndexMap[int64(s-1)] + DB.SliceEntryCountMap[int64(s-1)]
				DB.SliceEndIndexMap[int64(s)] = DB.SliceStartIndexMap[int64(s)] + DB.SliceEntryCountMap[int64(s)] - 1
			}
		} else {
			DB.SliceEntryCountMap[int64(s)] = 256
			DB.EntryCount += 256
			if s == 0 {
				DB.SliceStartIndexMap[int64(s)] = 0
				DB.SliceEndIndexMap[int64(s)] = 256
			} else {
				DB.SliceStartIndexMap[int64(s)] = DB.SliceStartIndexMap[int64(s-1)] + DB.SliceEntryCountMap[int64(s-1)]
				DB.SliceEndIndexMap[int64(s)] = DB.SliceStartIndexMap[int64(s)] + DB.SliceEntryCountMap[int64(s)] - 1
			}
		}
	}

	for e := 0; e < int(DB.EntryCount); e++ {
		for s := 0; s < DB.SliceCount; s++ {
			if e < int(DB.SliceStartIndexMap[int64(s)]+DB.SliceEntryCountMap[int64(s)])-1 {
				DB.EntryToSliceMap[int64(e)] = int64(s)
				break
			}
		}
	}
	fmt.Printf("%+v\n", DB)
}

/*
FP_PORT_FIELD_SEL.*[0]: <SLICE9_S_TYPE_SEL=1,SLICE9_PAIRING_IPBM_F0=0,SLICE9_PAIRING_FIXED=0,SLICE9_NORMALIZE_MAC_ADDR=0,SLICE9_NORMALIZE_IP_ADDR=0,SLICE9_FIELDS=0x401ea6,SLICE9_F3=0xf,SLICE9_F2=5,SLICE9_F1=6,SLICE9_D_TYPE_SEL=0,SLICE9_8_PAIRING=1,SLICE8_S_TYPE_SEL=0,SLICE8_PAIRING_IPBM_F0=0,SLICE8_PAIRING_FIXED=0,SLICE8_NORMALIZE_MAC_ADDR=0,SLICE8_NORMALIZE_IP_ADDR=0,SLICE8_FIELDS=0xe09,SLICE8_F3=7,SLICE8_F2=0,SLICE8_F1=9,SLICE8_D_TYPE_SEL=0,SLICE7_S_TYPE_SEL=0,SLICE7_PAIRING_IPBM_F0=0,SLICE7_PAIRING_FIXED=0,SLICE7_NORMALIZE_MAC_ADDR=0,SLICE7_NORMALIZE_IP_ADDR=0,SLICE7_FIELDS=0xe64,SLICE7_F3=7,SLICE7_F2=3,SLICE7_F1=4,SLICE7_D_TYPE_SEL=0,SLICE7_6_PAIRING=1,SLICE6_S_TYPE_SEL=1,SLICE6_PAIRING_IPBM_F0=0,SLICE6_PAIRING_FIXED=0,SLICE6_NORMALIZE_MAC_ADDR=0,SLICE6_NORMALIZE_IP_ADDR=0,SLICE6_FIELDS=0x401a41,SLICE6_F3=0xd,SLICE6_F2=2,SLICE6_F1=1,SLICE6_D_TYPE_SEL=0,SLICE5_S_TYPE_SEL=1,SLICE5_PAIRING_IPBM_F0=0,SLICE5_PAIRING_FIXED=0,SLICE5_NORMALIZE_MAC_ADDR=0,SLICE5_NORMALIZE_IP_ADDR=0,SLICE5_FIELDS=0x401ea6,SLICE5_F3=0xf,SLICE5_F2=5,SLICE5_F1=6,SLICE5_D_TYPE_SEL=0,SLICE5_4_PAIRING=1,SLICE4_S_TYPE_SEL=0,SLICE4_PAIRING_IPBM_F0=0,SLICE4_PAIRING_FIXED=0,SLICE4_NORMALIZE_MAC_ADDR=0,SLICE4_NORMALIZE_IP_ADDR=0,SLICE4_FIELDS=0xe09,SLICE4_F3=7,SLICE4_F2=0,SLICE4_F1=9,SLICE4_D_TYPE_SEL=0,SLICE3_S_TYPE_SEL=1,SLICE3_PAIRING_IPBM_F0=0,SLICE3_PAIRING_FIXED=0,SLICE3_NORMALIZE_MAC_ADDR=0,SLICE3_NORMALIZE_IP_ADDR=0,SLICE3_FIELDS=0x401ea6,SLICE3_F3=0xf,SLICE3_F2=5,SLICE3_F1=6,SLICE3_D_TYPE_SEL=0,SLICE3_2_PAIRING=1,SLICE2_S_TYPE_SEL=0,SLICE2_PAIRING_IPBM_F0=0,SLICE2_PAIRING_FIXED=0,SLICE2_NORMALIZE_MAC_ADDR=0,SLICE2_NORMALIZE_IP_ADDR=0,SLICE2_FIELDS=0xe09,SLICE2_F3=7,SLICE2_F2=0,SLICE2_F1=9,SLICE2_D_TYPE_SEL=0,SLICE1_S_TYPE_SEL=1,SLICE1_PAIRING_IPBM_F0=0,SLICE1_PAIRING_FIXED=0,SLICE1_NORMALIZE_MAC_ADDR=0,SLICE1_NORMALIZE_IP_ADDR=0,SLICE1_FIELDS=0x401ea6,SLICE1_F3=0xf,SLICE1_F2=5,SLICE1_F1=6,SLICE1_D_TYPE_SEL=0,SLICE1_0_PAIRING=1,SLICE11_S_TYPE_SEL=1,SLICE11_PAIRING_IPBM_F0=0,SLICE11_PAIRING_FIXED=0,SLICE11_NORMALIZE_MAC_ADDR=0,SLICE11_NORMALIZE_IP_ADDR=0,SLICE11_FIELDS=0x401ea6,SLICE11_F3=0xf,SLICE11_F2=5,SLICE11_F1=6,SLICE11_D_TYPE_SEL=0,SLICE11_10_PAIRING=1,SLICE10_S_TYPE_SEL=0,SLICE10_PAIRING_IPBM_F0=0,SLICE10_PAIRING_FIXED=0,SLICE10_NORMALIZE_MAC_ADDR=0,SLICE10_NORMALIZE_IP_ADDR=0,SLICE10_FIELDS=0xe09,SLICE10_F3=7,SLICE10_F2=0,SLICE10_F1=9,SLICE10_D_TYPE_SEL=0,SLICE0_S_TYPE_SEL=0,SLICE0_PAIRING_IPBM_F0=0,SLICE0_PAIRING_FIXED=0,SLICE0_NORMALIZE_MAC_ADDR=0,SLICE0_NORMALIZE_IP_ADDR=0,SLICE0_FIELDS=0xe09,SLICE0_F3=7,SLICE0_F2=0,SLICE0_F1=9,SLICE0_D_TYPE_SEL=0,EVEN_PARITY_1=0,EVEN_PARITY_0=0>
*/

type SliceFieldSelector struct {
	raw                string
	S_TYPE_SEL         int64
	PAIRING_IPBM       int64
	PAIRING_FIXED      int64
	NORMALIZE_MAC_ADDR int64
	NORMALIZE_IP_ADDR  int64
	FIELDS             string
	FP3                int64
	FP2                int64
	FP1                int64
	D_TYPE_SEL         int64
	PARING_ODD_SLICE   int64
	PAIRING_IPBM_F0    int64
}

func (sfs *SliceFieldSelector) String() string {
	return fmt.Sprintf("S_TYPE_SEL: %d, PAIRING_IPBB: %d, PAIRING_FIXED: %d, NORMALIZE_MAC_ADDR: %d, NORMALIZE_IP_ADDR: %d, FIELDS: %s, FP3: %d, FP2: %d, FP1: %d, D_TYPE_SEL: %d, PARING_ODD_SLICE: %d, PAIRING_IPBM_F0: %d", sfs.S_TYPE_SEL, sfs.PAIRING_IPBM, sfs.PAIRING_FIXED, sfs.NORMALIZE_MAC_ADDR, sfs.NORMALIZE_IP_ADDR, sfs.FIELDS, sfs.FP3, sfs.FP2, sfs.FP1, sfs.D_TYPE_SEL, sfs.PARING_ODD_SLICE, sfs.PAIRING_IPBM_F0)

}

var EvenSliceMatchFormat = "SLICE%d_S_TYPE_SEL=(?P<sv>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_PAIRING_IPBM_F0=(?P<pipbm>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_PAIRING_FIXED=(?P<pf>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_NORMALIZE_MAC_ADDR=(?P<nmac>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_NORMALIZE_IP_ADDR=(?P<nip>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_FIELDS=(?P<fields>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_F3=(?P<f3>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_F2=(?P<f2>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_F1=(?P<f1>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_D_TYPE_SEL=(?P<dts>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_%d_PAIRING=(?P<evp>[0]*[xX]*[0-9a-fA-F]+),"
var OddSliceMatchFormat = "SLICE%d_S_TYPE_SEL=(?P<sts>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_PAIRING_IPBM_F0=(?P<pipm>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_PAIRING_FIXED=(?P<pf>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_NORMALIZE_MAC_ADDR=(?P<nmac>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_NORMALIZE_IP_ADDR=(?P<nip>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_FIELDS=(?P<fields>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_F3=(?P<f3>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_F2=(?P<f2>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_F1=(?P<f1>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_D_TYPE_SEL=(?P<dts>[0]*[xX]*[0-9a-fA-F]+),"

type PortFieldSelector struct {
	Index               int64
	SliceFieldSelectors map[int64]*SliceFieldSelector
}

func (pfs *PortFieldSelector) String() string {
	res := fmt.Sprintf("Port: %d\n", pfs.Index)
	for i, sfs := range pfs.SliceFieldSelectors {
		res += fmt.Sprintf("     Slice: %d\n", i)
		res += fmt.Sprintf("            %s\n", sfs)
	}

	return res
}

var PFSIndexReg = regexp.MustCompile(`FP_PORT_FIELD_SEL\.\*\[(?P<index>[0]*[xX]*[0-9a-fA-F]+)\]:`)

func DumpPortFieldSelector(dev *rut.RUT) {
	dumpTableAndSaveToFile(dev, "FP_PORT_FIELD_SEL", "0", "127", FP_PORT_FIELD_SEL_FILE("info"))
	table, _ := ioutil.ReadFile(FP_PORT_FIELD_SEL_FILE("info"))

	lines := strings.Split(string(table), "\r\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "FP_PORT_FIELD_SEL") {
			index, _ := strconv.ParseInt(PFSIndexReg.FindStringSubmatch(line)[1], 0, 32)
			var pfs PortFieldSelector
			pfs.Index = index
			pfs.SliceFieldSelectors = make(map[int64]*SliceFieldSelector, 1)
			for i := 0; i < DB.SliceCount; i++ {
				var EvenSliceReg = regexp.MustCompile(fmt.Sprintf(EvenSliceMatchFormat, i, i, i, i, i, i, i, i, i, i, i, i-1))
				var OddSliceReg = regexp.MustCompile(fmt.Sprintf(OddSliceMatchFormat, i, i, i, i, i, i, i, i, i, i))
				var fs SliceFieldSelector

				if i%2 != 0 {
					match := EvenSliceReg.FindStringSubmatch(line)
					if len(match) != 0 {
						fs.raw = match[0]
						fs.S_TYPE_SEL, _ = strconv.ParseInt(match[1], 0, 32)
						fs.PAIRING_IPBM, _ = strconv.ParseInt(match[2], 0, 32)
						fs.PAIRING_FIXED, _ = strconv.ParseInt(match[3], 0, 32)
						fs.NORMALIZE_MAC_ADDR, _ = strconv.ParseInt(match[4], 0, 32)
						fs.NORMALIZE_IP_ADDR, _ = strconv.ParseInt(match[5], 0, 32)
						fs.FIELDS = match[6]
						fs.FP3, _ = strconv.ParseInt(match[7], 0, 32)
						fs.FP2, _ = strconv.ParseInt(match[8], 0, 32)
						fs.FP1, _ = strconv.ParseInt(match[9], 0, 32)
						fs.D_TYPE_SEL, _ = strconv.ParseInt(match[10], 0, 32)
						fs.PARING_ODD_SLICE, _ = strconv.ParseInt(match[11], 0, 32)
					}
				} else {
					match := OddSliceReg.FindStringSubmatch(line)
					if len(match) != 0 {
						fs.raw = match[0]
						fs.S_TYPE_SEL, _ = strconv.ParseInt(match[1], 0, 32)
						fs.PAIRING_IPBM_F0, _ = strconv.ParseInt(match[2], 0, 32)
						fs.PAIRING_FIXED, _ = strconv.ParseInt(match[3], 0, 32)
						fs.NORMALIZE_MAC_ADDR, _ = strconv.ParseInt(match[4], 0, 32)
						fs.NORMALIZE_IP_ADDR, _ = strconv.ParseInt(match[5], 0, 32)
						fs.FIELDS = match[6]
						fs.FP3, _ = strconv.ParseInt(match[7], 0, 32)
						fs.FP2, _ = strconv.ParseInt(match[8], 0, 32)
						fs.FP1, _ = strconv.ParseInt(match[9], 0, 32)
						fs.D_TYPE_SEL, _ = strconv.ParseInt(match[10], 0, 32)
					}
				}
				pfs.SliceFieldSelectors[int64(i)] = &fs
			}
			DB.PFS[pfs.Index] = &pfs
		}
	}
	for i := 0; i < DB.SliceCount; i++ {
		fmt.Printf("Slice: %d\n", i)
		fmt.Printf("       FP1: %+v\n", BCM56850ICAPFieldSelector_single.FP1[int(DB.PFS[0].SliceFieldSelectors[int64(i)].FP1)])
		fmt.Printf("       FP2: %+v\n", BCM56850ICAPFieldSelector_single.FP2[int(DB.PFS[0].SliceFieldSelectors[int64(i)].FP2)])
		fmt.Printf("       FP3: %+v\n", BCM56850ICAPFieldSelector_single.FP3[int(DB.PFS[0].SliceFieldSelectors[int64(i)].FP3)])
	}
}

var FPTCAMIndexReg = regexp.MustCompile(`FP_TCAM\.\*\[(?P<index>[0]*[xX]*[0-9a-fA-F]+)\]:`)

func FPTCAMEntryParse(file string) {
	table, _ := ioutil.ReadFile(file)
	match := FPTCAMIndexReg.FindAllStringSubmatch(string(table), -1)
	if len(match) > 0 {
		fmt.Printf("%+v\n", match)
		for _, m := range match {
			index, _ := strconv.ParseInt(m[1], 0, 32)
			fmt.Printf("Entry: %d is in Slice: %d\n", index, DB.EntryToSliceMap[index])
			fmt.Printf(" %+v\n", BCM56850ICAPFieldSelector_single.FP1[int(DB.PFS[0].SliceFieldSelectors[DB.EntryToSliceMap[index]].FP1)])
			fmt.Printf(" %+v\n", BCM56850ICAPFieldSelector_single.FP2[int(DB.PFS[0].SliceFieldSelectors[DB.EntryToSliceMap[index]].FP2)])
			fmt.Printf(" %+v\n", BCM56850ICAPFieldSelector_single.FP3[int(DB.PFS[0].SliceFieldSelectors[DB.EntryToSliceMap[index]].FP3)])
		}
	}
}

func DumpTable(dev *rut.RUT, version string) {
	err := os.Remove(FP_TCAM_FILE(version))
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}
	err = os.Remove(FP_POLICY_TABLE_FILE(version))
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	err = os.Remove(FP_GLOBAL_MASK_TCAM_FILE(version))
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	err = os.Remove(FP_PORT_FIELD_SEL_FILE(version))
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	err = os.Remove(FP_SLICE_KEY_CONTROL_FILE(version))
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	err = os.Remove(FP_SLICE_INDEX_CONTROL_FILE(version))
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	err = os.Remove(FP_SLICE_MAP_FILE(version))
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	data, err := dev.RunCommand(CTX, &command.Command{
		Mode: "config",
		CMD:  " do q sh -l",
	})
	if err != nil {
		fmt.Println(err)
	}

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "shell",
		CMD:  " scontrol -f /proc/switch/ASIC/ctrl dump table 0 FP_SLICE_MAP 1 1",
	})
	if err != nil {
		fmt.Println(err)
	}

	util.SaveToFile("FP_SLICE_MAP_"+version+".txt", []byte(data))

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "shell",
		CMD:  " scontrol -f /proc/switch/ASIC/ctrl dump table 0 FP_PORT_FIELD_SEL 0 127",
	})
	if err != nil {
		fmt.Println(err)
	}

	util.SaveToFile(FP_PORT_FIELD_SEL_FILE(version), []byte(data))

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "shell",
		CMD:  " scontrol -f /proc/switch/ASIC/ctrl dump table 0 FP_TCAM 0 4095",
	})
	if err != nil {
		fmt.Println(err)
	}

	util.SaveToFile(FP_TCAM_FILE(version), []byte(data))

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "shell",
		CMD:  " scontrol -f /proc/switch/ASIC/ctrl dump table 0 FP_GLOBAL_MASK_TCAM 0 4095",
	})
	if err != nil {
		fmt.Println(err)
	}

	util.SaveToFile(FP_GLOBAL_MASK_TCAM_FILE(version), []byte(data))

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "shell",
		CMD:  " scontrol -f /proc/switch/ASIC/ctrl dump table 0 FP_POLICY_TABLE 0 4095",
	})

	if err != nil {
		fmt.Println(err)
	}
	util.SaveToFile(FP_POLICY_TABLE_FILE(version), []byte(data))

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "shell",
		CMD:  " scontrol -f /proc/switch/ASIC/ctrl dump table 0 FP_SLICE_KEY_CONTROL 0 1",
	})

	if err != nil {
		fmt.Println(err)
	}
	util.SaveToFile(FP_SLICE_KEY_CONTROL_FILE(version), []byte(data))

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "shell",
		CMD:  " scontrol -f /proc/switch/ASIC/ctrl dump table 0 FP_SLICE_MAP 0 0",
	})

	if err != nil {
		fmt.Println(err)
	}
	util.SaveToFile(FP_SLICE_MAP_FILE(version), []byte(data))

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "shell",
		CMD:  " exit",
	})

	if err != nil {
		fmt.Println(err)
	}
}

func FP_TCAM_FILE(version string) string {
	return "FP_TCAM." + version + ".txt"
}

func FP_POLICY_TABLE_FILE(version string) string {
	return "FP_POLICY_TABLE." + version + ".txt"
}

func FP_PORT_FIELD_SEL_FILE(version string) string {
	return "FP_PORT_FIELD_SEL." + version + ".txt"
}

func FP_SLICE_KEY_CONTROL_FILE(version string) string {
	return "FP_SLICE_KEY_CONTROL." + version + ".txt"
}

func FP_SLICE_INDEX_CONTROL_FILE(version string) string {
	return "FP_SLICE_INDEX_CONTROL." + version + ".txt"
}

func FP_GLOBAL_MASK_TCAM_FILE(version string) string {
	return "FP_GLOBAL_MASK_TCAM_FILE." + version + ".txt"
}

func FP_SLICE_MAP_FILE(version string) string {
	return "FP_SLICE_MAP_FILE." + version + ".txt"
}

func main() {
	flag.Parse()

	ip := net.ParseIP(*IP)
	if ip == nil {
		fmt.Printf("Invalid IP address: %s\n", *IP)
		return
	}

	if *Host == "" {
		fmt.Println("Invalid Host name")
		return
	}

	if *User == "" {
		fmt.Println("Invalid username")
		return
	}

	if *SFU != "A" && *SFU != "B" {
		fmt.Printf("Invalid SFU: %s\n", *SFU)
		return
	}

	dev, err := rut.New(&rut.RUT{
		Name:     "V8500_SFU",
		Device:   "V8",
		IP:       *IP,
		Port:     "23",
		Username: *User,
		Hostname: *Host,
		Password: *Password,
		SFU:      *SFU,
	})

	if err != nil {
		panic(err)
	}

	err = dev.Init()
	if err != nil {
		panic(err)
	}

	data, err := dev.RunCommand(CTX, &command.Command{
		Mode: "normal",
		CMD:  " config terminal",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "config",
		CMD:  " show running-config",
	})
	if err != nil {
		fmt.Println(err)
	}

	DelRule(dev, "ipany")
	DumpTable(dev, "ipany_b")
	AddRule(dev, "ipany", "ip any any", "copy-to-cpu")
	DumpTable(dev, "ipany_a")

	util.DiffFile(FP_TCAM_FILE("ipany_b"), FP_TCAM_FILE("ipany_a"))
	util.DiffFileToHtml(FP_TCAM_FILE("ipany_b"), FP_TCAM_FILE("ipany_a"))
	util.DiffFile(FP_POLICY_TABLE_FILE("ipany_b"), FP_POLICY_TABLE_FILE("ipany_a"))
	util.DiffFile(FP_PORT_FIELD_SEL_FILE("ipany_b"), FP_PORT_FIELD_SEL_FILE("ipany_a"))
	util.DiffFile(FP_GLOBAL_MASK_TCAM_FILE("ipany_b"), FP_GLOBAL_MASK_TCAM_FILE("ipany_a"))

	DelRule(dev, "macany")
	DumpTable(dev, "macany_b")
	AddRule(dev, "macany", "mac any any", "deny")
	DumpTable(dev, "macany_a")

	util.DiffFile(FP_TCAM_FILE("macany_b"), FP_TCAM_FILE("macany_a"))
	util.DiffFileToHtml(FP_TCAM_FILE("macany_b"), FP_TCAM_FILE("macany_a"))
	util.DiffFile(FP_POLICY_TABLE_FILE("macany_b"), FP_POLICY_TABLE_FILE("macany_a"))
	util.DiffFile(FP_PORT_FIELD_SEL_FILE("macany_b"), FP_PORT_FIELD_SEL_FILE("macany_a"))
	util.DiffFile(FP_GLOBAL_MASK_TCAM_FILE("macany_b"), FP_GLOBAL_MASK_TCAM_FILE("macany_a"))

	DelRule(dev, "ip10.10.10.10")
	DumpTable(dev, "ip10.10.10.10_b")
	AddRule(dev, "ip10.10.10.10", "ip 10.10.10.10 8.8.8.8", "deny")
	DumpTable(dev, "ip10.10.10.10_a")

	util.DiffFile(FP_TCAM_FILE("ip10.10.10.10_b"), FP_TCAM_FILE("ip10.10.10.10_a"))
	util.DiffFileToHtml(FP_TCAM_FILE("ip10.10.10.10_b"), FP_TCAM_FILE("ip10.10.10.10_a"))
	util.DiffFile(FP_POLICY_TABLE_FILE("ip10.10.10.10_b"), FP_POLICY_TABLE_FILE("ip10.10.10.10_a"))
	util.DiffFile(FP_PORT_FIELD_SEL_FILE("ip10.10.10.10_b"), FP_PORT_FIELD_SEL_FILE("ip10.10.10.10_a"))
	util.DiffFile(FP_GLOBAL_MASK_TCAM_FILE("ip10.10.10.10_b"), FP_GLOBAL_MASK_TCAM_FILE("ip10.10.10.10_a"))

	DelRule(dev, "ip6_2001:db8")
	DumpTable(dev, "ip6_2002:db8_b")
	AddRule(dev, "ip6_2002:db8", "ipv6 2001:db8::10 2001:db8::20", "deny")
	DumpTable(dev, "ip6_2002:db8_a")

	util.DiffFile(FP_TCAM_FILE("ip6_2002:db8_b"), FP_TCAM_FILE("ip6_2002:db8_a"))
	util.DiffFileToHtml(FP_TCAM_FILE("ip6_2002:db8_b"), FP_TCAM_FILE("ip6_2002:db8_a"))
	util.DiffFile(FP_POLICY_TABLE_FILE("ip6_2002:db8_b"), FP_POLICY_TABLE_FILE("ip6_2002:db8_a"))
	util.DiffFile(FP_PORT_FIELD_SEL_FILE("ip6_2002:db8_b"), FP_PORT_FIELD_SEL_FILE("ip6_2002:db8_a"))
	util.DiffFile(FP_GLOBAL_MASK_TCAM_FILE("ip6_2002:db8_b"), FP_GLOBAL_MASK_TCAM_FILE("ip6_2002:db8_a"))

	DelRule(dev, "mac_00_d0_cb")
	DumpTable(dev, "mac_00_d0_cb_b")
	AddRule(dev, "mac_00_d0_cb", "mac 00:d0:cb:01:02:03 11:22:33:44:55:66", "deny")
	DumpTable(dev, "mac_00_d0_cb_a")

	util.DiffFile(FP_TCAM_FILE("mac_00_d0_cb_b"), FP_TCAM_FILE("mac_00_d0_cb_a"))
	util.DiffFileToHtml(FP_TCAM_FILE("mac_00_d0_cb_b"), FP_TCAM_FILE("mac_00_d0_cb_a"))
	util.DiffFile(FP_POLICY_TABLE_FILE("mac_00_d0_cb_b"), FP_POLICY_TABLE_FILE("mac_00_d0_cb_a"))
	util.DiffFile(FP_PORT_FIELD_SEL_FILE("mac_00_d0_cb_b"), FP_PORT_FIELD_SEL_FILE("mac_00_d0_cb_a"))
	util.DiffFile(FP_GLOBAL_MASK_TCAM_FILE("mac_00_d0_cb_b"), FP_GLOBAL_MASK_TCAM_FILE("mac_00_d0_cb_a"))

	DumpSliceInfo(dev)
	DumpPortFieldSelector(dev)
	FPTCAMEntryParse(FP_TCAM_FILE("ip10.10.10.10_a") + ".line.diff")
	FPTCAMEntryParse(FP_TCAM_FILE("mac_00_d0_cb_a") + ".line.diff")
	StartServer()
}

func StartServer() {

	mux := http.NewServeMux()
	//@liwei: This need more analysis.
	mux.HandleFunc("/", Login)
	mux.Handle("/asset/web/", http.FileServer(http.Dir(".")))

	//Add context support
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data, _ := ioutil.ReadFile("./" + r.URL.RequestURI())
		w.Write(data)
	}
}
