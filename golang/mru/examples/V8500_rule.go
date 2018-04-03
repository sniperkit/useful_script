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
	Key  []TLV
	Data string
	Mask string
}

const (
	FP1 = iota
	FP2
	FP3
	FP4
	FIXED
)

type RuleEntry struct {
	Key         map[int][]TLV
	Index       int64
	FP1         string
	FP1_MASK    string
	FP2         string
	FP2_MASK    string
	FP3         string
	FP3_MASK    string
	FP4         string
	FP4_MASK    string
	FIXED       string
	FIXED_MASK  string
	PairedEntry *RuleEntry
}

func (re *RuleEntry) String() string {
	var res string
	if DB.EntryToSliceMap[re.Index]%2 == 0 {
		if DB.PFS[0].SliceFieldSelectors[DB.EntryToSliceMap[re.Index]+1].PARING_ODD_SLICE != 0 {
			res += fmt.Sprintf("[%05d] + [%05d]:\n", re.Index, re.PairedEntry.Index)
			res += fmt.Sprintf("   Key: \n")
			res += fmt.Sprintf("       FP1: %+v\n", re.Key[FP1])
			res += fmt.Sprintf("       FP2: %+v\n", re.Key[FP2])
			res += fmt.Sprintf("       FP3: %+v\n", re.Key[FP3])
			res += fmt.Sprintf("       FP4: %+v\n", re.Key[FP4])
			res += fmt.Sprintf("       FIXED: %+v\n", re.Key[FIXED])
			res += fmt.Sprintf("   Field: \n")
			res += fmt.Sprintf("       FP1: %40s:  FP1_MASK: %40s\n", re.FP1, re.FP1_MASK)
			res += fmt.Sprintf("       FP2: %40s:  FP2_MASK: %40s\n", re.FP2, re.FP2_MASK)
			res += fmt.Sprintf("       FP3: %40s:  FP3_MASK: %40s\n", re.FP3, re.FP3_MASK)
			res += fmt.Sprintf("       FP4: %40s:  FP4_MASK: %40s\n", re.FP4, re.FP4_MASK)
			res += fmt.Sprintf("       FIXED: %38s:  FIXED_MASK: %38s\n", re.FIXED, re.FIXED_MASK)
			res += fmt.Sprintf("   PairedEntryKey: \n")
			res += fmt.Sprintf("       FP1: %+v\n", re.PairedEntry.Key[FP1])
			res += fmt.Sprintf("       FP2: %+v\n", re.PairedEntry.Key[FP2])
			res += fmt.Sprintf("       FP3: %+v\n", re.PairedEntry.Key[FP3])
			res += fmt.Sprintf("       FP4: %+v\n", re.PairedEntry.Key[FP4])
			res += fmt.Sprintf("       FIXED: %+v\n", re.PairedEntry.Key[FIXED])
			res += fmt.Sprintf("   PairedEntryField: \n")
			res += fmt.Sprintf("       FP1: %40s:  FP1_MASK: %40s\n", re.PairedEntry.FP1, re.PairedEntry.FP1_MASK)
			res += fmt.Sprintf("       FP2: %40s:  FP2_MASK: %40s\n", re.PairedEntry.FP2, re.PairedEntry.FP2_MASK)
			res += fmt.Sprintf("       FP3: %40s:  FP3_MASK: %40s\n", re.PairedEntry.FP3, re.PairedEntry.FP3_MASK)
			res += fmt.Sprintf("       FP4: %40s:  FP4_MASK: %40s\n", re.PairedEntry.FP4, re.PairedEntry.FP4_MASK)
			res += fmt.Sprintf("       FIXED: %38s:  FIXED_MASK: %38s\n", re.PairedEntry.FIXED, re.PairedEntry.FIXED_MASK)
		} else {
			res += fmt.Sprintf("[%05d] \n", re.Index)
			res += fmt.Sprintf("   Key: \n")
			res += fmt.Sprintf("       FP1: %+v\n", re.Key[FP1])
			res += fmt.Sprintf("       FP2: %+v\n", re.Key[FP2])
			res += fmt.Sprintf("       FP3: %+v\n", re.Key[FP3])
			res += fmt.Sprintf("       FP4: %+v\n", re.Key[FP4])
			res += fmt.Sprintf("       FIXED: %+v\n", re.Key[FIXED])
			res += fmt.Sprintf("   Field: \n")
			res += fmt.Sprintf("       FP1: %40s:  FP1_MASK: %40s\n", re.FP1, re.FP1_MASK)
			res += fmt.Sprintf("       FP2: %40s:  FP2_MASK: %40s\n", re.FP2, re.FP2_MASK)
			res += fmt.Sprintf("       FP3: %40s:  FP3_MASK: %40s\n", re.FP3, re.FP3_MASK)
			res += fmt.Sprintf("       FP4: %40s:  FP4_MASK: %40s\n", re.FP4, re.FP4_MASK)
			res += fmt.Sprintf("       FIXED: %38s:  FIXED_MASK: %38s\n", re.FIXED, re.FIXED_MASK)
		}
	} else {
		if DB.PFS[0].SliceFieldSelectors[DB.EntryToSliceMap[re.Index]].PARING_ODD_SLICE == 0 {
			res += fmt.Sprintf("[%05d] \n", re.Index)
			res += fmt.Sprintf("   Key: \n")
			res += fmt.Sprintf("       FP1: %+v\n", re.Key[FP1])
			res += fmt.Sprintf("       FP2: %+v\n", re.Key[FP2])
			res += fmt.Sprintf("       FP3: %+v\n", re.Key[FP3])
			res += fmt.Sprintf("       FP4: %+v\n", re.Key[FP4])
			res += fmt.Sprintf("       FIXED: %+v\n", re.Key[FIXED])
			res += fmt.Sprintf("   Field: \n")
			res += fmt.Sprintf("       FP1: %40s:  FP1_MASK: %40s\n", re.FP1, re.FP1_MASK)
			res += fmt.Sprintf("       FP2: %40s:  FP2_MASK: %40s\n", re.FP2, re.FP2_MASK)
			res += fmt.Sprintf("       FP3: %40s:  FP3_MASK: %40s\n", re.FP3, re.FP3_MASK)
			res += fmt.Sprintf("       FP4: %40s:  FP4_MASK: %40s\n", re.FP4, re.FP4_MASK)
			res += fmt.Sprintf("       FIXED: %38s:  FIXED_MASK: %38s\n", re.FIXED, re.FIXED_MASK)
		}
	}

	return res
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
	EntryCount         int64
	PFS                map[int64]*PortFieldSelector
	RuleEntries        map[int64]*RuleEntry
}

var DB RuleDB = RuleDB{
	VirtualSliceMap:    make(map[int64]int64, 1),
	SliceGroupMap:      make(map[int64]int64, 1),
	SliceEntryCountMap: make(map[int64]int64, 1),
	SliceStartIndexMap: make(map[int64]int64, 1),
	SliceEndIndexMap:   make(map[int64]int64, 1),
	EntryToSliceMap:    make(map[int64]int64, 1),
	PFS:                make(map[int64]*PortFieldSelector, 1),
	RuleEntries:        make(map[int64]*RuleEntry, 1),
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
			TLV{Name: "IPV6_DIP", Size: 128, Offset: 0},
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
			TLV{Name: "IPV6_DIP", Size: 128, Offset: 0},
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
			TLV{Name: "IPV6_DIP", Size: 128, Offset: 0},
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

func AddRulePort(dev *rut.RUT, name string, flow string, action string, port string) error {
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
		&command.Command{Mode: "config", CMD: " interface " + port},
		&command.Command{Mode: "config-if", CMD: " service-policy input " + name},
		&command.Command{Mode: "config-if", CMD: " exit"},
	})

	return err
}

func AddRulePortPriority(dev *rut.RUT, name, flow, action, port, priority string) error {
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
		&command.Command{Mode: "config-policy", CMD: " priority " + priority},
		&command.Command{Mode: "config-policy", CMD: " apply"},
		&command.Command{Mode: "config-policy", CMD: " exit"},
		&command.Command{Mode: "config", CMD: " interface " + port},
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

func DelRulePort(dev *rut.RUT, name, port string) error {
	_, err := dev.RunCommands(CTX, []*command.Command{
		&command.Command{Mode: "config", CMD: " interface " + port},
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
}

type SliceFieldSelector struct {
	raw                string
	S_TYPE_SEL         int64
	PAIRING_IPBM       int64
	PAIRING_FIXED      int64
	NORMALIZE_MAC_ADDR int64
	NORMALIZE_IP_ADDR  int64
	FIELDS             string
	FP3                int64
	FP4                int64
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
}

var FPTCAMIndexReg = regexp.MustCompile(`FP_TCAM\.\*\[(?P<index>[0]*[xX]*[0-9a-fA-F]+)\]:`)

func FPTCAMEntryParse(file string) {
	table, _ := ioutil.ReadFile(file)
	lines := strings.Split(string(table), "\r\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "FP_TCAM") && strings.Contains(line, "VALID=3") {
			FPTCAMParseLine(line)
		}
	}
	fmt.Printf("Total entry count: (%d)\n", len(DB.RuleEntries))
}

var FPTCAMEntryF1Reg = regexp.MustCompile("F1_MASK=(?P<f1m>[0]*[xX]*[0-9a-fA-F])+,F1=(?P<f1>[0]*[xX]*[0-9a-fA-F]+)")
var FPTCAMEntryF2Reg = regexp.MustCompile("F2_MASK=(?P<f2m>[0]*[xX]*[0-9a-fA-F]+),F2=(?P<f2>[0]*[xX]*[0-9a-fA-F]+)")
var FPTCAMEntryF3Reg = regexp.MustCompile("F3_MASK=(?P<f3m>[0]*[xX]*[0-9a-fA-F]+),F3=(?P<f3>[0]*[xX]*[0-9a-fA-F]+)")
var FPTCAMEntryF4Reg = regexp.MustCompile("F4_MASK=(?P<f4m>[0]*[xX]*[0-9a-fA-F]+),F4=(?P<f4>[0]*[xX]*[0-9a-fA-F]+)")
var FPTCAMEntryFIXEDReg = regexp.MustCompile("FIXED_MASK=(?P<fim>[0]*[xX]*[0-9a-fA-F]+),FIXED=(?P<fi>[0]*[xX]*[0-9a-fA-F]+)")

func FPTCAMParseLine(line string) {
	var rule RuleEntry
	match := FPTCAMIndexReg.FindStringSubmatch(string(line))
	index, _ := strconv.ParseInt(match[1], 0, 32)
	rule.Index = index
	rule.Key = make(map[int][]TLV, 1)
	if DB.EntryToSliceMap[index]%2 != 0 {
		rule.Key[FP1] = BCM56850ICAPFieldSelector_single.FP1[int(DB.PFS[0].SliceFieldSelectors[DB.EntryToSliceMap[index]].FP1)]
		rule.Key[FP2] = BCM56850ICAPFieldSelector_single.FP2[int(DB.PFS[0].SliceFieldSelectors[DB.EntryToSliceMap[index]].FP2)]
		rule.Key[FP3] = BCM56850ICAPFieldSelector_single.FP3[int(DB.PFS[0].SliceFieldSelectors[DB.EntryToSliceMap[index]].FP3)]
		rule.Key[FP4] = BCM56850ICAPFieldSelector_single.FP4[int(DB.PFS[0].SliceFieldSelectors[DB.EntryToSliceMap[index]].FP4)]
		rule.Key[FIXED] = BCM56850ICAPFieldSelector_single.FIXED[int(DB.PFS[0].SliceFieldSelectors[DB.EntryToSliceMap[index]].PAIRING_FIXED)]
	} else {
		rule.Key[FP1] = BCM56850ICAPFieldSelector_Double.FP1[int(DB.PFS[0].SliceFieldSelectors[DB.EntryToSliceMap[index]].FP1)]
		rule.Key[FP2] = BCM56850ICAPFieldSelector_Double.FP2[int(DB.PFS[0].SliceFieldSelectors[DB.EntryToSliceMap[index]].FP2)]
		rule.Key[FP3] = BCM56850ICAPFieldSelector_Double.FP3[int(DB.PFS[0].SliceFieldSelectors[DB.EntryToSliceMap[index]].FP3)]
		rule.Key[FP4] = BCM56850ICAPFieldSelector_Double.FP4[int(DB.PFS[0].SliceFieldSelectors[DB.EntryToSliceMap[index]].FP4)]
		rule.Key[FIXED] = BCM56850ICAPFieldSelector_Double.FIXED[int(DB.PFS[0].SliceFieldSelectors[DB.EntryToSliceMap[index]].PAIRING_FIXED)]
	}

	f1match := FPTCAMEntryF1Reg.FindStringSubmatch(line)
	if len(f1match) > 1 {
		rule.FP1 = f1match[2]
		rule.FP1_MASK = f1match[1]
	}
	f2match := FPTCAMEntryF2Reg.FindStringSubmatch(line)
	if len(f2match) > 1 {
		rule.FP2 = f2match[2]
		rule.FP2_MASK = f2match[1]
	}

	f3match := FPTCAMEntryF3Reg.FindStringSubmatch(line)
	if len(f3match) > 1 {
		rule.FP3 = f3match[2]
		rule.FP3_MASK = f3match[1]
	}
	f4match := FPTCAMEntryF4Reg.FindStringSubmatch(line)
	if len(f4match) > 1 {
		rule.FP4 = f4match[2]
		rule.FP4_MASK = f4match[1]
	}

	fixedmatch := FPTCAMEntryFIXEDReg.FindStringSubmatch(line)
	if len(fixedmatch) > 1 {
		rule.FIXED = fixedmatch[2]
		rule.FIXED_MASK = fixedmatch[1]
	}
	DB.RuleEntries[index] = &rule

	for id, rule := range DB.RuleEntries {
		slice := DB.EntryToSliceMap[id]
		if slice%2 == 0 {
			pindex := id + DB.SliceEntryCountMap[slice]
			rule.PairedEntry = DB.RuleEntries[pindex]
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

func RulePortBindingTest(dev *rut.RUT) {
	DelRulePort(dev, "giga8_1", "gigabitethernet 8/1")
	DumpTable(dev, "giga8_1_b")
	AddRulePort(dev, "giga8_1", "ip any any", "copy-to-cpu", "gigabitethernet 8/1")
	DumpTable(dev, "giga8_1_a")
	util.DiffFile(FP_GLOBAL_MASK_TCAM_FILE("giga8_1_b"), FP_GLOBAL_MASK_TCAM_FILE("giga8_1_a"))
	util.DiffFile(FP_TCAM_FILE("giga8_1_b"), FP_TCAM_FILE("giga8_1_a"))

	DelRulePort(dev, "giga8_2", "gigabitethernet 8/2")
	DumpTable(dev, "giga8_2_b")
	AddRulePort(dev, "giga8_2", "ip any any", "copy-to-cpu", "gigabitethernet 8/2")
	DumpTable(dev, "giga8_2_a")
	util.DiffFile(FP_GLOBAL_MASK_TCAM_FILE("giga8_2_b"), FP_GLOBAL_MASK_TCAM_FILE("giga8_2_a"))
	util.DiffFile(FP_TCAM_FILE("giga8_2_b"), FP_TCAM_FILE("giga8_2_a"))

	DelRulePort(dev, "giga8_3", "gigabitethernet 8/3")
	DumpTable(dev, "giga8_3_b")
	AddRulePort(dev, "giga8_3", "ip any any", "copy-to-cpu", "gigabitethernet 8/3")
	DumpTable(dev, "giga8_3_a")
	util.DiffFile(FP_GLOBAL_MASK_TCAM_FILE("giga8_3_b"), FP_GLOBAL_MASK_TCAM_FILE("giga8_3_a"))
	util.DiffFile(FP_TCAM_FILE("giga8_3_b"), FP_TCAM_FILE("giga8_3_a"))

	DelRulePort(dev, "giga8_4", "gigabitethernet 8/4")
	DumpTable(dev, "giga8_4_b")
	AddRulePort(dev, "giga8_4", "ip any any", "copy-to-cpu", "gigabitethernet 8/4")
	DumpTable(dev, "giga8_4_a")
	util.DiffFile(FP_GLOBAL_MASK_TCAM_FILE("giga8_4_b"), FP_GLOBAL_MASK_TCAM_FILE("giga8_4_a"))
	util.DiffFile(FP_TCAM_FILE("giga8_4_b"), FP_TCAM_FILE("giga8_4_a"))

	DelRulePort(dev, "giga8_5", "gigabitethernet 8/5")
	DumpTable(dev, "giga8_5_b")
	AddRulePort(dev, "giga8_5", "ip any any", "copy-to-cpu", "gigabitethernet 8/5")
	DumpTable(dev, "giga8_5_a")
	util.DiffFile(FP_GLOBAL_MASK_TCAM_FILE("giga8_5_b"), FP_GLOBAL_MASK_TCAM_FILE("giga8_5_a"))
	util.DiffFile(FP_TCAM_FILE("giga8_5_b"), FP_TCAM_FILE("giga8_5_a"))

	DelRulePort(dev, "giga8_6", "gigabitethernet 8/6")
	DumpTable(dev, "giga8_6_b")
	AddRulePort(dev, "giga8_6", "ip any any", "copy-to-cpu", "gigabitethernet 8/6")
	DumpTable(dev, "giga8_6_a")
	util.DiffFile(FP_GLOBAL_MASK_TCAM_FILE("giga8_6_b"), FP_GLOBAL_MASK_TCAM_FILE("giga8_6_a"))
	util.DiffFile(FP_TCAM_FILE("giga8_6_b"), FP_TCAM_FILE("giga8_6_a"))

	DelRulePort(dev, "giga8_1", "gigabitethernet 8/1")
	DelRulePort(dev, "giga8_2", "gigabitethernet 8/2")
	DelRulePort(dev, "giga8_3", "gigabitethernet 8/3")
	DelRulePort(dev, "giga8_4", "gigabitethernet 8/4")
	DelRulePort(dev, "giga8_5", "gigabitethernet 8/5")
	DelRulePort(dev, "giga8_6", "gigabitethernet 8/6")
}

const (
	PriorityLow = iota
	PriorityMiddle
	PriorityHigh
	PriorityHighest
)

var RulePriorityMap = map[int]string{
	PriorityLow:     "low",
	PriorityMiddle:  "medium",
	PriorityHigh:    "high",
	PriorityHighest: "highest",
}

func RulePriorityTest(dev *rut.RUT) {
	for _, p := range RulePriorityMap {
		DelRulePort(dev, p, "gigabitethernet 8/6")
		DumpTable(dev, p+"_b")
		AddRulePortPriority(dev, p, "ip any any", "copy-to-cpu", "gigabitethernet 8/6", p)
		DumpTable(dev, p+"_a")
		util.DiffFile(FP_GLOBAL_MASK_TCAM_FILE(p+"_b"), FP_GLOBAL_MASK_TCAM_FILE(p+"_a"))
		util.DiffFile(FP_TCAM_FILE(p+"_b"), FP_TCAM_FILE(p+"_a"))
	}
}

func BasicRuleFunctionTest(dev *rut.RUT) {
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

	DelRule(dev, "ip_tcp_100_200")
	DumpTable(dev, "ip_tcp_100_200_b")
	AddRule(dev, "ip_tcp_100_200", "ip 120.120.120.120 6.6.6.6 tcp 100 200", "deny")
	DumpTable(dev, "ip_tcp_100_200_a")

	util.DiffFile(FP_TCAM_FILE("ip_tcp_100_200_b"), FP_TCAM_FILE("ip_tcp_100_200_a"))
	util.DiffFileToHtml(FP_TCAM_FILE("ip_tcp_100_200_b"), FP_TCAM_FILE("ip_tcp_100_200_a"))
	util.DiffFile(FP_POLICY_TABLE_FILE("ip_tcp_100_200_b"), FP_POLICY_TABLE_FILE("ip_tcp_100_200_a"))
	util.DiffFile(FP_PORT_FIELD_SEL_FILE("ip_tcp_100_200_b"), FP_PORT_FIELD_SEL_FILE("ip_tcp_100_200_a"))
	util.DiffFile(FP_GLOBAL_MASK_TCAM_FILE("ip_tcp_100_200_b"), FP_GLOBAL_MASK_TCAM_FILE("ip_tcp_100_200_a"))

	DelRule(dev, "ipv6_tcp_300_400")
	DumpTable(dev, "ipv6_tcp_300_400_b")
	AddRule(dev, "ipv6_tcp_300_400", "ipv6 2001:db8::2001 2001:db8::2002 tcp 300 400", "deny")
	DumpTable(dev, "ipv6_tcp_300_400_a")

	util.DiffFile(FP_TCAM_FILE("ipv6_tcp_300_400_b"), FP_TCAM_FILE("ipv6_tcp_300_400_a"))
	util.DiffFileToHtml(FP_TCAM_FILE("ipv6_tcp_300_400_b"), FP_TCAM_FILE("ipv6_tcp_300_400_a"))
	util.DiffFile(FP_POLICY_TABLE_FILE("ipv6_tcp_300_400_b"), FP_POLICY_TABLE_FILE("ipv6_tcp_300_400_a"))
	util.DiffFile(FP_PORT_FIELD_SEL_FILE("ipv6_tcp_300_400_b"), FP_PORT_FIELD_SEL_FILE("ipv6_tcp_300_400_a"))
	util.DiffFile(FP_GLOBAL_MASK_TCAM_FILE("ipv6_tcp_300_400_b"), FP_GLOBAL_MASK_TCAM_FILE("ipv6_tcp_300_400_a"))

	DelRule(dev, "ether_type_8899")
	DumpTable(dev, "ether_type_8899_b")
	AddRule(dev, "ether_type_8899", "ethtype 0x8899", "deny")
	DumpTable(dev, "ether_type_8899_a")

	util.DiffFile(FP_TCAM_FILE("ether_type_8899_b"), FP_TCAM_FILE("ether_type_8899_a"))
	util.DiffFileToHtml(FP_TCAM_FILE("ether_type_8899_b"), FP_TCAM_FILE("ether_type_8899_a"))
	util.DiffFile(FP_POLICY_TABLE_FILE("ether_type_8899_b"), FP_POLICY_TABLE_FILE("ether_type_8899_a"))
	util.DiffFile(FP_PORT_FIELD_SEL_FILE("ether_type_8899_b"), FP_PORT_FIELD_SEL_FILE("ether_type_8899_a"))
	util.DiffFile(FP_GLOBAL_MASK_TCAM_FILE("ether_type_8899_b"), FP_GLOBAL_MASK_TCAM_FILE("ether_type_8899_a"))
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

	DumpTable(dev, "raw")
	DumpSliceInfo(dev)
	DumpPortFieldSelector(dev)
	BasicRuleFunctionTest(dev)
	//RulePortBindingTest(dev)
	//RulePriorityTest(dev)
	FPTCAMEntryParse(FP_TCAM_FILE("ip10.10.10.10_a"))
	for _, r := range DB.RuleEntries {
		fmt.Printf("%+v\n", r)
	}
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
