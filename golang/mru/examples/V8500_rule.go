package main

import (
	//"cline"
	"command"
	//"configuration"
	"context"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"rut"
	"sort"
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
	IPBM
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
	IPBM        string
	IPBM_MASK   string
	PairedEntry *RuleEntry
	Policy      *PolicyEntry
}

type RuleEntrySlice []*RuleEntry

func (res RuleEntrySlice) Len() int           { return len(res) }
func (res RuleEntrySlice) Swap(i, j int)      { res[i], res[j] = res[j], res[i] }
func (res RuleEntrySlice) Less(i, j int) bool { return res[i].Index < res[j].Index }

var PolicyEntryFields = []string{
	"Y_NEW_PKT_PRI",
	"Y_NEW_ECN",
	"Y_NEW_DSCP",
	"Y_DROP_PRECEDENCE",
	"Y_DROP",
	"Y_COS_INT_PRI",
	"Y_COPY_TO_CPU",
	"Y_CHANGE_PKT_PRI",
	"Y_CHANGE_ECN",
	"Y_CHANGE_DSCP",
	"Y_CHANGE_COS_OR_INT_PRI",
	"UNICAST_REDIRECT_CONTROL",
	"SUPPRESS_VXLT",
	"SFLOW_ING_SAMPLE",
	"SFLOW_EGR_SAMPLE",
	"R_NEW_PKT_PRI",
	"R_NEW_ECN",
	"R_NEW_DSCP",
	"R_DROP_PRECEDENCE",
	"R_DROP",
	"R_COS_INT_PRI",
	"R_COPY_TO_CPU",
	"R_CHANGE_PKT_PRI",
	"R_CHANGE_ECN",
	"R_CHANGE_DSCP",
	"R_CHANGE_COS_OR_INT_PRI",
	"RSVD_NEXT_HOP_INDEX",
	"RESERVED_EH_TM",
	"RESERVED_3",
	"RESERVED_2",
	"RESERVED_1",
	"RESERVED_0",
	"REDIRECT_USE_IFP_PBM",
	"REDIRECT_T",
	"REDIRECT_PORT_NUM",
	"REDIRECT_NHI",
	"REDIRECT_MODID",
	"REDIRECT_L3MC_INDEX",
	"REDIRECT_L2MC_INDEX",
	"REDIRECT_INDEX_TYPE",
	"REDIRECT_IFP_PROFILE_INDEX",
	"REDIRECT_ECMP_GROUP",
	"REDIRECT_DVP",
	"REDIRECT_DGLP",
	"REDIRECTION",
	"PPD3_CLASS_TAG",
	"OAM_UP_MEP",
	"OAM_TX",
	"OAM_TUNNEL_CONTROL",
	"OAM_TAG_STATUS_CHECK_CONTROL",
	"OAM_SERVICE_PRI_MAPPING_PTR",
	"OAM_LM_EN",
	"OAM_LM_BASE_PTR",
	"OAM_LMEP_MDL",
	"OAM_LMEP_EN",
	"OAM_ENABLE_LM_DM_SAMPLE",
	"OAM_DM_TYPE",
	"OAM_DM_EN",
	"NEXT_HOP_INDEX",
	"NAT_PACKET_EDIT_IDX",
	"NAT_PACKET_EDIT_ENTRY_SEL",
	"NAT_ENABLE",
	"MTP_INDEX3",
	"MTP_INDEX2",
	"MTP_INDEX1",
	"MTP_INDEX0",
	"MIRROR_OVERRIDE",
	"MIRROR",
	"METER_UPDATE_ODD",
	"METER_UPDATE_EVEN",
	"METER_TEST_ODD",
	"METER_TEST_EVEN",
	"METER_PAIR_MODE_MODIFIER",
	"METER_PAIR_MODE",
	"METER_PAIR_INDEX_ODD",
	"METER_PAIR_INDEX_EVEN",
	"MATCHED_RULE",
	"LAG_RH_DISABLE",
	"I2E_CLASSID_SEL",
	"I2E_CLASSID",
	"HG_CLASSID_SEL",
	"HGT_RH_DISABLE",
	"HGT_DLB_DISABLE",
	"G_PACKET_REDIRECTION",
	"G_NEW_PKT_PRI",
	"G_NEW_ECN",
	"G_NEW_DSCP_TOS",
	"G_L3SW_CHANGE_L2_FIELDS",
	"G_DROP_PRECEDENCE",
	"G_DROP",
	"G_COS_INT_PRI",
	"G_COPY_TO_CPU",
	"G_CHANGE_PKT_PRI",
	"G_CHANGE_ECN",
	"G_CHANGE_DSCP_TOS",
	"G_CHANGE_COS_OR_INT_PRI",
	"GREEN_TO_PID",
	"FCOE_ZONE_CHECK_ACTION",
	"FCOE_VSAN_PRI_VALID",
	"FCOE_VSAN_PRI",
	"FCOE_VSAN_ID",
	"EVEN_PARITY",
	"EH_TAG_TYPE",
	"EH_QUEUE_TAG",
	"ECMP_RH_DISABLE",
	"ECMP_PTR",
	"ECMP_NH_INFO",
	"ECMP_HASH_SEL",
	"ECMP",
	"DO_NOT_URPF",
	"DO_NOT_NAT",
	"DO_NOT_GENERATE_CNM",
	"DO_NOT_CUT_THROUGH",
	"DO_NOT_CHANGE_TTL",
	"CPU_COS",
	"COUNTER_MODE",
	"COUNTER_INDEX",
	"CHANGE_CPU_COS",
}

type PolicyEntry struct {
	Fields map[string]int64
}

type FieldSlice []map[string]int64

func (fs FieldSlice) Len() int      { return len(fs) }
func (fs FieldSlice) Swap(i, j int) { fs[i], fs[j] = fs[j], fs[i] }
func (fs FieldSlice) Less(i, j int) bool {
	var ik string
	var jk string
	for k, _ := range fs[i] {
		ik = k
		break
	}
	for k, _ := range fs[j] {
		jk = k
		break
	}

	return ik < jk
}

func (pe *PolicyEntry) String() string {
	var res string
	var fs FieldSlice = make([]map[string]int64, 0, 1)
	res += fmt.Sprintf("[")
	for k, field := range pe.Fields {
		if field != 0 {
			fs = append(fs, map[string]int64{k: field})
		}
	}

	sort.Sort(fs)

	for _, f := range fs {
		for k, field := range f {
			res += fmt.Sprintf(" %10s -> %2d ", k, field)
		}
	}
	res += fmt.Sprintf("]")

	return res
}

//Flow 配置的内容是由FP_GLOBAL_MASK_TCAM 和FP_TCAM两张表决定的, 所以Rule Entry的内容需要解析这两张表.
// F2, F4, FIXED 从FP_TCAM中获取，
// F1, F3, IPBM 从FP_GLOBAL_MASK_TCAM中获取
/* FP_GLOBAL_MASK_TCAM.*[1]: <VALID=3,SPARE_MASK=0,SPARE=0,MASK=0x807fffffffffffffffffffffffc000000000000000000000000,KEY=0x7fffffffffffffffffffffffc000000000000000000000000,IPBM_MASK=0x201ffffffffffffffffffffffff,IPBM=0x1ffffffffffffffffffffffff,F3_MASK=0,F3=0,F1_MASK=0,F1=0>
 */

func (re *RuleEntry) String() string {
	var res string
	var Yellow = color.New(color.FgYellow)
	var Cyan = color.New(color.FgCyan)
	var Green = color.New(color.FgGreen)
	if DB.EntryToSliceMap[re.Index]%2 == 0 {
		res += fmt.Sprintf("{")
		if DB.PFS[0].SliceFieldSelectors[DB.EntryToSliceMap[re.Index]+1].PAIRING_EVEN_SLICE != 0 {
			res += Yellow.Sprintf("[%05d(%d)] + [%05d(%d)]:\n", re.Index, DB.EntryToSliceMap[re.Index], re.PairedEntry.Index, DB.EntryToSliceMap[re.PairedEntry.Index])
			res += fmt.Sprintf("   Key: \n")
			res += fmt.Sprintf("       FP1: %+v\n", re.Key[FP1])
			res += fmt.Sprintf("       FP2: %+v\n", re.Key[FP2])
			res += fmt.Sprintf("       FP3: %+v\n", re.Key[FP3])
			res += fmt.Sprintf("       FP4: %+v\n", re.Key[FP4])
			res += fmt.Sprintf("       FIXED: %+v\n", re.Key[FIXED])
			res += Cyan.Sprintf("   Field: \n")
			res += fmt.Sprintf("       FP1: %40s:  FP1_MASK: %40s\n", re.FP1, re.FP1_MASK)
			res += fmt.Sprintf("       FP2: %40s:  FP2_MASK: %40s\n", re.FP2, re.FP2_MASK)
			res += fmt.Sprintf("       FP3: %40s:  FP3_MASK: %40s\n", re.FP3, re.FP3_MASK)
			res += fmt.Sprintf("       FP4: %40s:  FP4_MASK: %40s\n", re.FP4, re.FP4_MASK)
			res += fmt.Sprintf("       FIXED: %38s:  FIXED_MASK: %38s\n", re.FIXED, re.FIXED_MASK)
			res += fmt.Sprintf("       IPBM: %39s:  IPBM_MASK: %39s\n", re.IPBM, re.IPBM_MASK)
			res += Green.Sprintf("   Action: \n")
			res += fmt.Sprintf("       %s\n", re.Policy)
			res += fmt.Sprintf("   PairedEntryKey: \n")
			res += fmt.Sprintf("       FP1: %+v\n", re.PairedEntry.Key[FP1])
			res += fmt.Sprintf("       FP2: %+v\n", re.PairedEntry.Key[FP2])
			res += fmt.Sprintf("       FP3: %+v\n", re.PairedEntry.Key[FP3])
			res += fmt.Sprintf("       FP4: %+v\n", re.PairedEntry.Key[FP4])
			res += fmt.Sprintf("       FIXED: %+v\n", re.PairedEntry.Key[FIXED])
			res += fmt.Sprintf("       PAIRING_IPBM_F0: %+v\n", re.PairedEntry.Key[IPBM])
			res += Cyan.Sprintf("   PairedEntryField: \n")
			res += fmt.Sprintf("       FP1: %40s:  FP1_MASK: %40s\n", re.PairedEntry.FP1, re.PairedEntry.FP1_MASK)
			res += fmt.Sprintf("       FP2: %40s:  FP2_MASK: %40s\n", re.PairedEntry.FP2, re.PairedEntry.FP2_MASK)
			res += fmt.Sprintf("       FP3: %40s:  FP3_MASK: %40s\n", re.PairedEntry.FP3, re.PairedEntry.FP3_MASK)
			res += fmt.Sprintf("       FP4: %40s:  FP4_MASK: %40s\n", re.PairedEntry.FP4, re.PairedEntry.FP4_MASK)
			res += fmt.Sprintf("       FIXED: %38s:  FIXED_MASK: %38s\n", re.PairedEntry.FIXED, re.PairedEntry.FIXED_MASK)
			res += fmt.Sprintf("       PAIRING_IPBM_F0: %28s:  PAIRING_IPBM_F0_MASK: %28s\n", re.PairedEntry.IPBM, re.PairedEntry.IPBM_MASK)
			res += Green.Sprintf("   PairedEntryAction: \n")
			res += fmt.Sprintf("       %s\n", re.PairedEntry.Policy)
		} else {
			res += Yellow.Sprintf("[%05d(%d)] \n", re.Index, DB.EntryToSliceMap[re.Index])
			res += fmt.Sprintf("   Key: \n")
			res += fmt.Sprintf("       FP1: %+v\n", re.Key[FP1])
			res += fmt.Sprintf("       FP2: %+v\n", re.Key[FP2])
			res += fmt.Sprintf("       FP3: %+v\n", re.Key[FP3])
			res += fmt.Sprintf("       FP4: %+v\n", re.Key[FP4])
			res += fmt.Sprintf("       FIXED: %+v\n", re.Key[FIXED])
			res += Cyan.Sprintf("   Field: \n")
			res += fmt.Sprintf("       FP1: %40s:  FP1_MASK: %40s\n", re.FP1, re.FP1_MASK)
			res += fmt.Sprintf("       FP2: %40s:  FP2_MASK: %40s\n", re.FP2, re.FP2_MASK)
			res += fmt.Sprintf("       FP3: %40s:  FP3_MASK: %40s\n", re.FP3, re.FP3_MASK)
			res += fmt.Sprintf("       FP4: %40s:  FP4_MASK: %40s\n", re.FP4, re.FP4_MASK)
			res += fmt.Sprintf("       FIXED: %38s:  FIXED_MASK: %38s\n", re.FIXED, re.FIXED_MASK)
			res += fmt.Sprintf("       IPBM: %39s:  IPBM_MASK: %39s\n", re.IPBM, re.IPBM_MASK)
			res += Green.Sprintf("   Action: \n")
			res += fmt.Sprintf("       %s\n", re.Policy)
		}
		res += fmt.Sprintf("}")
	} else {
		if DB.PFS[0].SliceFieldSelectors[DB.EntryToSliceMap[re.Index]].PAIRING_EVEN_SLICE == 0 {
			res += fmt.Sprintf("{")
			res += Yellow.Sprintf("[%05d(%d)] \n", re.Index, DB.EntryToSliceMap[re.Index])
			res += fmt.Sprintf("   Key: \n")
			res += fmt.Sprintf("       FP1: %+v\n", re.Key[FP1])
			res += fmt.Sprintf("       FP2: %+v\n", re.Key[FP2])
			res += fmt.Sprintf("       FP3: %+v\n", re.Key[FP3])
			res += fmt.Sprintf("       FP4: %+v\n", re.Key[FP4])
			res += fmt.Sprintf("       FIXED: %+v\n", re.Key[FIXED])
			res += Cyan.Sprintf("   Field: \n")
			res += fmt.Sprintf("       FP1: %40s:  FP1_MASK: %40s\n", re.FP1, re.FP1_MASK)
			res += fmt.Sprintf("       FP2: %40s:  FP2_MASK: %40s\n", re.FP2, re.FP2_MASK)
			res += fmt.Sprintf("       FP3: %40s:  FP3_MASK: %40s\n", re.FP3, re.FP3_MASK)
			res += fmt.Sprintf("       FP4: %40s:  FP4_MASK: %40s\n", re.FP4, re.FP4_MASK)
			res += fmt.Sprintf("       FIXED: %38s:  FIXED_MASK: %38s\n", re.FIXED, re.FIXED_MASK)
			res += fmt.Sprintf("       IPBM: %39s:  IPBM_MASK: %39s\n", re.IPBM, re.IPBM_MASK)
			res += Green.Sprintf("   Action: \n")
			res += fmt.Sprintf("       %s\n", re.Policy)
			res += fmt.Sprintf("}")
		}
	}

	return res
}

type RuleRawEntry struct {
	Index               int64
	FP_TCAM             string
	FP_GLOBAL_MASK_TCAM string
	FP_POLICY_TABLE     string
}

type RuleDB struct {
	Device             *rut.RUT
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
	RuleEntriesOrdered RuleEntrySlice
	RawEntries         map[int64]*RuleRawEntry
	RawTables          map[string]string
}

func (rdb *RuleDB) Clear() {
	rdb.SliceCount = 0
	rdb.GroupCount = 0
	rdb.EntryCount = 0
	rdb.VirtualSliceMap = make(map[int64]int64, 1)
	rdb.SliceGroupMap = make(map[int64]int64, 1)
	rdb.SliceEntryCountMap = make(map[int64]int64, 1)
	rdb.SliceStartIndexMap = make(map[int64]int64, 1)
	rdb.SliceEndIndexMap = make(map[int64]int64, 1)
	rdb.EntryToSliceMap = make(map[int64]int64, 1)
	rdb.PFS = make(map[int64]*PortFieldSelector, 1)
	rdb.RuleEntries = make(map[int64]*RuleEntry, 1)
	rdb.RawEntries = make(map[int64]*RuleRawEntry, 1)
	rdb.RawTables = make(map[string]string, 1)
	rdb.RuleEntriesOrdered = make([]*RuleEntry, 1)
}

func (rdb *RuleDB) IsInitialized() bool {
	if _, ok := rdb.RawTables["FP_TCAM"]; ok {
		if _, ok := rdb.RawTables["FP_GLOBAL_MASK_TCAM"]; ok {
			if _, ok := rdb.RawTables["FP_POLICY_TABLE"]; ok {
			}
			return true
		}

	}

	return false
}

var FPTCAMIndexReg = regexp.MustCompile(`FP_TCAM\.\*\[(?P<index>[0]*[xX]*[0-9a-fA-F]+)\]:`)

var FPGlobalMaskTCAMEntryRegFmt = `FP_GLOBAL_MASK_TCAM\.\*\[%d\]:[[:space:]]*<[a-zA-Z0-9,=_[:space:]]+>`
var FPPolicyTableEntryRegFmt = `FP_POLICY_TABLE\.\*\[%d\]:[[:space:]]*<[a-zA-Z0-9,=_[:space:]+]+>`

func (rdb *RuleDB) GetRawEntries() error {
	if !rdb.IsInitialized() {
		panic("Cannot get rule entries, DB not initilaized")
	}

	//Clear the Raw DB before a new iteration
	rdb.RawEntries = make(map[int64]*RuleRawEntry, 1)

	ft := rdb.RawTables["FP_TCAM"]
	fgmt := rdb.RawTables["FP_GLOBAL_MASK_TCAM"]
	fpt := rdb.RawTables["FP_POLICY_TABLE"]

	lines := strings.Split(string(ft), "\r\n")
	for _, line := range lines {
		var rule RuleRawEntry
		if strings.HasPrefix(line, "FP_TCAM") && strings.Contains(line, "VALID=3") {
			rule.FP_TCAM = line
			match := FPTCAMIndexReg.FindStringSubmatch(string(line))
			rule.Index, _ = strconv.ParseInt(match[1], 0, 32)

			fgmtEntryReg := regexp.MustCompile(fmt.Sprintf(FPGlobalMaskTCAMEntryRegFmt, rule.Index))
			fgmtEntry := fgmtEntryReg.FindStringSubmatch(string(fgmt))
			rule.FP_GLOBAL_MASK_TCAM = fgmtEntry[0]

			fptEntryReg := regexp.MustCompile(fmt.Sprintf(FPPolicyTableEntryRegFmt, rule.Index))
			fptEntry := fptEntryReg.FindStringSubmatch(string(fpt))
			rule.FP_POLICY_TABLE = fptEntry[0]

			rdb.RawEntries[rule.Index] = &rule
		}
	}

	return nil
}

func (rdb *RuleDB) ParseRawEntries() {
	if !rdb.IsInitialized() {
		panic("Cannot parse rule entries, DB not initilaized")
	}

	rdb.RuleEntries = make(map[int64]*RuleEntry, 1)
	rdb.RuleEntriesOrdered = make([]*RuleEntry, 0, 1)
	for index, rr := range rdb.RawEntries {
		rdb.RuleEntries[index] = rdb.ParseRawEntry(rr)
		rdb.RuleEntriesOrdered = append(rdb.RuleEntriesOrdered, rdb.RuleEntries[index])
	}

	sort.Sort(rdb.RuleEntriesOrdered)

	for id, rule := range DB.RuleEntries {
		slice := DB.EntryToSliceMap[id]
		if slice%2 == 0 {
			pindex := id + DB.SliceEntryCountMap[slice]
			rule.PairedEntry = DB.RuleEntries[pindex]
		}
	}
}

func (rdb *RuleDB) ParseRawEntry(raw *RuleRawEntry) *RuleEntry {
	var rule RuleEntry
	rule.Index = raw.Index
	rule.Key = make(map[int][]TLV, 1)
	if rdb.EntryToSliceMap[rule.Index]%2 == 0 { /* Even Slice Key */
		rule.Key[FP1] = BCM56850ICAPFieldSelector_Even.FP1[int(rdb.PFS[0].SliceFieldSelectors[rdb.EntryToSliceMap[rule.Index]].FP1)]
		rule.Key[FP2] = BCM56850ICAPFieldSelector_Even.FP2[int(rdb.PFS[0].SliceFieldSelectors[rdb.EntryToSliceMap[rule.Index]].FP2)]
		rule.Key[FP3] = BCM56850ICAPFieldSelector_Even.FP3[int(rdb.PFS[0].SliceFieldSelectors[rdb.EntryToSliceMap[rule.Index]].FP3)]
		rule.Key[FP4] = BCM56850ICAPFieldSelector_Even.FP4[int(rdb.PFS[0].SliceFieldSelectors[rdb.EntryToSliceMap[rule.Index]].FP4)]
		rule.Key[FIXED] = BCM56850ICAPFieldSelector_Even.FIXED[int(rdb.PFS[0].SliceFieldSelectors[rdb.EntryToSliceMap[rule.Index]].PAIRING_FIXED)]
	} else { /* Odd Slice Key */
		rule.Key[FP1] = BCM56850ICAPFieldSelector_Odd.FP1[int(rdb.PFS[0].SliceFieldSelectors[rdb.EntryToSliceMap[rule.Index]].FP1)]
		rule.Key[FP2] = BCM56850ICAPFieldSelector_Odd.FP2[int(rdb.PFS[0].SliceFieldSelectors[rdb.EntryToSliceMap[rule.Index]].FP2)]
		rule.Key[FP3] = BCM56850ICAPFieldSelector_Odd.FP3[int(rdb.PFS[0].SliceFieldSelectors[rdb.EntryToSliceMap[rule.Index]].FP3)]
		rule.Key[FP4] = BCM56850ICAPFieldSelector_Odd.FP4[int(rdb.PFS[0].SliceFieldSelectors[rdb.EntryToSliceMap[rule.Index]].FP4)]
		rule.Key[FIXED] = BCM56850ICAPFieldSelector_Odd.PAIRING_FIXED[int(rdb.PFS[0].SliceFieldSelectors[rdb.EntryToSliceMap[rule.Index]].PAIRING_FIXED)]
		rule.Key[IPBM] = BCM56850ICAPFieldSelector_Odd.PAIRING_IPBM_F0[int(rdb.PFS[0].SliceFieldSelectors[rdb.EntryToSliceMap[rule.Index]].PAIRING_IPBM_F0)]
	}

	/* Get F2, F4, FIXED from FP_TCAM */
	f2match := FPTCAMEntryF2Reg.FindStringSubmatch(raw.FP_TCAM)
	if len(f2match) > 1 {
		rule.FP2 = f2match[2]
		rule.FP2_MASK = f2match[1]
	}

	f4match := FPTCAMEntryF4Reg.FindStringSubmatch(raw.FP_TCAM)
	if len(f4match) > 1 {
		rule.FP4 = f4match[2]
		rule.FP4_MASK = f4match[1]
	}

	fixedmatch := FPTCAMEntryFIXEDReg.FindStringSubmatch(raw.FP_TCAM)
	if len(fixedmatch) > 1 {
		rule.FIXED = fixedmatch[2]
		rule.FIXED_MASK = fixedmatch[1]
	}
	/* Get F1, F3, IPBM from FP_GLOBAL_MASK_TCAM */

	f1match := FPGlobalMaskTCAMF1Reg.FindStringSubmatch(raw.FP_GLOBAL_MASK_TCAM)
	if len(f1match) > 1 {
		rule.FP1 = f1match[2]
		rule.FP1_MASK = f1match[1]
	}

	f3match := FPGlobalMaskTCAMF3Reg.FindStringSubmatch(raw.FP_GLOBAL_MASK_TCAM)
	if len(f3match) > 1 {
		rule.FP3 = f3match[2]
		rule.FP3_MASK = f3match[1]
	}

	ipbmmatch := FPGlobalMaskTCAMIPBMReg.FindStringSubmatch(raw.FP_GLOBAL_MASK_TCAM)
	if len(ipbmmatch) > 1 {
		rule.IPBM = ipbmmatch[2]
		rule.IPBM_MASK = ipbmmatch[1]
	}

	rule.Policy = rdb.parsePolicyEntry(raw.FP_POLICY_TABLE)

	return &rule
}

var FieldValueRegFmt = `=(?P<value>0*[xX]*[0-9a-fA-F]+)`

func (rdb *RuleDB) parsePolicyEntry(line string) *PolicyEntry {
	var policy PolicyEntry
	policy.Fields = make(map[string]int64)
	for _, field := range PolicyEntryFields {
		matcher := regexp.MustCompile(field + FieldValueRegFmt)
		matches := matcher.FindStringSubmatch(line)
		policy.Fields[field], _ = strconv.ParseInt(matches[1], 0, 32)
	}

	return &policy
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
	RawEntries:         make(map[int64]*RuleRawEntry, 1),
	RawTables:          make(map[string]string, 1),
	RuleEntriesOrdered: make([]*RuleEntry, 0, 1),
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
var BCM56850ICAPFieldSelector_Even = ICAPFieldSelector{
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
var BCM56850ICAPFieldSelector_Odd = ICAPFieldSelector{
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
	PAIRING_EVEN_SLICE int64
	PAIRING_IPBM_F0    int64
}

func (sfs *SliceFieldSelector) String() string {
	return fmt.Sprintf("S_TYPE_SEL: %d, PAIRING_IPBB: %d, PAIRING_FIXED: %d, NORMALIZE_MAC_ADDR: %d, NORMALIZE_IP_ADDR: %d, FIELDS: %s, FP3: %d, FP2: %d, FP1: %d, D_TYPE_SEL: %d, PAIRING_EVEN_SLICE: %d, PAIRING_IPBM_F0: %d", sfs.S_TYPE_SEL, sfs.PAIRING_IPBM, sfs.PAIRING_FIXED, sfs.NORMALIZE_MAC_ADDR, sfs.NORMALIZE_IP_ADDR, sfs.FIELDS, sfs.FP3, sfs.FP2, sfs.FP1, sfs.D_TYPE_SEL, sfs.PAIRING_EVEN_SLICE, sfs.PAIRING_IPBM_F0)

}

var OddSliceMatchFormat = "SLICE%d_S_TYPE_SEL=(?P<sv>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_PAIRING_IPBM_F0=(?P<pipbm>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_PAIRING_FIXED=(?P<pf>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_NORMALIZE_MAC_ADDR=(?P<nmac>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_NORMALIZE_IP_ADDR=(?P<nip>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_FIELDS=(?P<fields>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_F3=(?P<f3>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_F2=(?P<f2>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_F1=(?P<f1>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_D_TYPE_SEL=(?P<dts>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_%d_PAIRING=(?P<evp>[0]*[xX]*[0-9a-fA-F]+),"
var EvenSliceMatchFormat = "SLICE%d_S_TYPE_SEL=(?P<sts>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_PAIRING_IPBM_F0=(?P<pipm>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_PAIRING_FIXED=(?P<pf>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_NORMALIZE_MAC_ADDR=(?P<nmac>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_NORMALIZE_IP_ADDR=(?P<nip>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_FIELDS=(?P<fields>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_F3=(?P<f3>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_F2=(?P<f2>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_F1=(?P<f1>[0]*[xX]*[0-9a-fA-F]+),SLICE%d_D_TYPE_SEL=(?P<dts>[0]*[xX]*[0-9a-fA-F]+),"

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

var FPTCAMEntryF2Reg = regexp.MustCompile("F2_MASK=(?P<f2m>[0]*[xX]*[0-9a-fA-F]+),F2=(?P<f2>[0]*[xX]*[0-9a-fA-F]+)")
var FPTCAMEntryF4Reg = regexp.MustCompile("F4_MASK=(?P<f4m>[0]*[xX]*[0-9a-fA-F]+),F4=(?P<f4>[0]*[xX]*[0-9a-fA-F]+)")
var FPTCAMEntryFIXEDReg = regexp.MustCompile("FIXED_MASK=(?P<fim>[0]*[xX]*[0-9a-fA-F]+),FIXED=(?P<fi>[0]*[xX]*[0-9a-fA-F]+)")
var FPGlobalMaskTCAMF1Reg = regexp.MustCompile("F1_MASK=(?P<f1m>[0]*[xX]*[0-9a-fA-F]+),F1=(?P<f1>[0]*[xX]*[0-9a-fA-F]+)")
var FPGlobalMaskTCAMF3Reg = regexp.MustCompile("F3_MASK=(?P<f3m>[0]*[xX]*[0-9a-fA-F]+),F3=(?P<f3>[0]*[xX]*[0-9a-fA-F]+)")
var FPGlobalMaskTCAMIPBMReg = regexp.MustCompile("IPBM_MASK=(?P<ipbmm>[0]*[xX]*[0-9a-fA-F]+),IPBM=(?P<ipbm>[0]*[xX]*[0-9a-fA-F]+)")

func (rdb *RuleDB) ParseSliceInfo() {
	if !rdb.IsInitialized() {
		panic("Cannot parse slice info: not initialized!")
	}
	table := rdb.RawTables["FP_SLICE_MAP"]
	rdb.SliceCount = 0
	rdb.GroupCount = 0
	rdb.EntryCount = 0
	rdb.SliceGroupMap = make(map[int64]int64, 1)
	rdb.SliceEntryCountMap = make(map[int64]int64, 1)
	rdb.SliceStartIndexMap = make(map[int64]int64, 1)
	rdb.SliceEndIndexMap = make(map[int64]int64, 1)
	rdb.EntryToSliceMap = make(map[int64]int64, 1)

	groups := SliceGroupMapReg.FindAllStringSubmatch(string(table), -1)
	for _, g := range groups {
		gidx, err := strconv.ParseInt(g[2], 0, 32)
		if err != nil {
			panic(err)
		}
		rdb.GroupCount++
		sidx, err := strconv.ParseInt(g[1], 0, 32)
		if err != nil {
			panic(err)
		}
		rdb.SliceGroupMap[sidx] = gidx
	}

	slices := VirtualSliceMapReg.FindAllStringSubmatch(string(table), -1)
	for _, s := range slices {
		pidx, err := strconv.ParseInt(s[2], 0, 32)
		if err != nil {
			panic(err)
		}
		rdb.SliceCount++
		vidx, err := strconv.ParseInt(s[1], 0, 32)
		if err != nil {
			panic(err)
		}
		rdb.SliceGroupMap[pidx] = vidx
	}

	for s := 0; s < rdb.SliceCount; s++ {
		if s < 8 {
			rdb.SliceEntryCountMap[int64(s)] = 256
			rdb.EntryCount += 256
			if s == 0 {
				rdb.SliceStartIndexMap[int64(s)] = 0
				rdb.SliceEndIndexMap[int64(s)] = 256
			} else {
				rdb.SliceStartIndexMap[int64(s)] = rdb.SliceStartIndexMap[int64(s-1)] + rdb.SliceEntryCountMap[int64(s-1)]
				rdb.SliceEndIndexMap[int64(s)] = rdb.SliceStartIndexMap[int64(s)] + rdb.SliceEntryCountMap[int64(s)] - 1
			}
		} else {
			rdb.SliceEntryCountMap[int64(s)] = 512
			rdb.EntryCount += 512
			if s == 0 {
				rdb.SliceStartIndexMap[int64(s)] = 512
				rdb.SliceEndIndexMap[int64(s)] = 512
			} else {
				rdb.SliceStartIndexMap[int64(s)] = rdb.SliceStartIndexMap[int64(s-1)] + rdb.SliceEntryCountMap[int64(s-1)]
				rdb.SliceEndIndexMap[int64(s)] = rdb.SliceStartIndexMap[int64(s)] + rdb.SliceEntryCountMap[int64(s)] - 1
			}
		}
	}

	for e := 0; e < int(rdb.EntryCount); e++ {
		for s := 0; s < rdb.SliceCount; s++ {
			if e < int(rdb.SliceStartIndexMap[int64(s)]+rdb.SliceEntryCountMap[int64(s)])-1 {
				rdb.EntryToSliceMap[int64(e)] = int64(s)
				break
			}
		}
	}
}

func (rdb *RuleDB) ParseKeys() {
	if !rdb.IsInitialized() {
		panic("Cannot parse slice info: not initialized!")
	}
	//Reset first
	rdb.PFS = make(map[int64]*PortFieldSelector, 1)

	table := rdb.RawTables["FP_PORT_FIELD_SEL"]
	lines := strings.Split(string(table), "\r\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "FP_PORT_FIELD_SEL") {
			index, _ := strconv.ParseInt(PFSIndexReg.FindStringSubmatch(line)[1], 0, 32)
			var pfs PortFieldSelector
			pfs.Index = index
			pfs.SliceFieldSelectors = make(map[int64]*SliceFieldSelector, 1)
			for i := 0; i < DB.SliceCount; i++ {
				var OddSliceReg = regexp.MustCompile(fmt.Sprintf(OddSliceMatchFormat, i, i, i, i, i, i, i, i, i, i, i, i-1))
				var EvenSliceReg = regexp.MustCompile(fmt.Sprintf(EvenSliceMatchFormat, i, i, i, i, i, i, i, i, i, i))
				var fs SliceFieldSelector

				if i%2 == 0 {
					match := EvenSliceReg.FindStringSubmatch(line)
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
				} else {
					match := OddSliceReg.FindStringSubmatch(line)
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
						fs.PAIRING_EVEN_SLICE, _ = strconv.ParseInt(match[11], 0, 32)
					}
				}
				pfs.SliceFieldSelectors[int64(i)] = &fs
			}
			rdb.PFS[pfs.Index] = &pfs
		}
	}
}

func (rdb *RuleDB) Prepare(dev *rut.RUT) {
	rdb.DumpTables(dev, "realtime")
}

func (rdb *RuleDB) Analysis() {
	rdb.ParseSliceInfo()
	rdb.ParseKeys()
	rdb.GetRawEntries()
	rdb.ParseRawEntries()
}

func (rdb *RuleDB) Dump(dev *rut.RUT, file string) {
	rdb.Prepare(dev)
	rdb.Analysis()

	var result string
	for _, r := range DB.RuleEntriesOrdered {
		fmt.Printf("%+v\n", r)
		result += fmt.Sprintf("%+v\n", r)
	}

	util.SaveToFile(file, []byte(result))
}

func (rdb *RuleDB) DumpTables(dev *rut.RUT, version string) {
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

	rdb.RawTables["FP_SLICE_MAP"] = string(data)

	util.SaveToFile("FP_SLICE_MAP_"+version+".txt", []byte(data))

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "shell",
		CMD:  " scontrol -f /proc/switch/ASIC/ctrl dump table 0 FP_PORT_FIELD_SEL 0 127",
	})
	if err != nil {
		fmt.Println(err)
	}

	rdb.RawTables["FP_PORT_FIELD_SEL"] = string(data)
	util.SaveToFile(FP_PORT_FIELD_SEL_FILE(version), []byte(data))

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "shell",
		CMD:  " scontrol -f /proc/switch/ASIC/ctrl dump table 0 FP_TCAM 0 4095",
	})
	if err != nil {
		fmt.Println(err)
	}

	rdb.RawTables["FP_TCAM"] = string(data)
	util.SaveToFile(FP_TCAM_FILE(version), []byte(data))

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "shell",
		CMD:  " scontrol -f /proc/switch/ASIC/ctrl dump table 0 FP_GLOBAL_MASK_TCAM 0 4095",
	})
	if err != nil {
		fmt.Println(err)
	}
	rdb.RawTables["FP_GLOBAL_MASK_TCAM"] = string(data)

	util.SaveToFile(FP_GLOBAL_MASK_TCAM_FILE(version), []byte(data))

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "shell",
		CMD:  " scontrol -f /proc/switch/ASIC/ctrl dump table 0 FP_POLICY_TABLE 0 4095",
	})

	if err != nil {
		fmt.Println(err)
	}
	rdb.RawTables["FP_POLICY_TABLE"] = string(data)
	util.SaveToFile(FP_POLICY_TABLE_FILE(version), []byte(data))

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "shell",
		CMD:  " scontrol -f /proc/switch/ASIC/ctrl dump table 0 FP_SLICE_KEY_CONTROL 0 0",
	})

	if err != nil {
		fmt.Println(err)
	}
	rdb.RawTables["FP_SLICE_KEY_CONTROL"] = string(data)
	util.SaveToFile(FP_SLICE_KEY_CONTROL_FILE(version), []byte(data))

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "shell",
		CMD:  " scontrol -f /proc/switch/ASIC/ctrl dump table 0 FP_SLICE_MAP 0 0",
	})

	if err != nil {
		fmt.Println(err)
	}
	rdb.RawTables["FP_SLICE_MAP"] = string(data)
	util.SaveToFile(FP_SLICE_MAP_FILE(version), []byte(data))

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "shell",
		CMD:  " scontrol -f /proc/switch/ASIC/ctrl dump table 0 UDF_TCAM 0 511",
	})

	if err != nil {
		fmt.Println(err)
	}
	rdb.RawTables["UDF_TCAM"] = string(data)
	util.SaveToFile(FP_UDF_TCAM_FILE(version), []byte(data))

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "shell",
		CMD:  " scontrol -f /proc/switch/ASIC/ctrl dump table 0 UDF_OFFSET 0 511",
	})

	if err != nil {
		fmt.Println(err)
	}
	rdb.RawTables["UDF_OFFSET"] = string(data)
	util.SaveToFile(FP_UDF_OFFSET_FILE(version), []byte(data))

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "shell",
		CMD:  " scontrol -f /proc/switch/ASIC/ctrl dump table 0 FP_RANGE_CHECK 0 31",
	})

	if err != nil {
		fmt.Println(err)
	}
	rdb.RawTables["FP_RANGE_CHECK"] = string(data)
	util.SaveToFile(FP_RANGE_CHECK_FILE(version), []byte(data))

	data, err = dev.RunCommand(CTX, &command.Command{
		Mode: "shell",
		CMD:  " exit",
	})

	if err != nil {
		fmt.Println(err)
	}
}

func (rdb *RuleDB) AnalysisRule(dev *rut.RUT, name, flow, action, port, priority string) {
	rdb.Device = dev
	//First Remove if already exist.
	rdb.RuleDel(dev, name, flow, action, port, priority)
	rdb.DumpTables(dev, "before."+name)
	rdb.RuleAdd(dev, name, flow, action, port, priority)
	rdb.DumpTables(dev, "after."+name)
	rdb.CompareTables("before."+name, "after."+name)
	rdb.Clear()
	rdb.Dump(dev, "after."+name+".txt")
	//Clear After analysis
	rdb.RuleDel(dev, name, flow, action, port, priority)
}

func (rdb *RuleDB) CompareTables(version1, version2 string) {
	util.DiffFile(FP_TCAM_FILE(version1), FP_TCAM_FILE(version2))
	util.DiffFile(FP_GLOBAL_MASK_TCAM_FILE(version1), FP_GLOBAL_MASK_TCAM_FILE(version2))
	util.DiffFile(FP_POLICY_TABLE_FILE(version1), FP_POLICY_TABLE_FILE(version2))
	util.DiffFile(FP_UDF_OFFSET_FILE(version1), FP_UDF_OFFSET_FILE(version2))
	util.DiffFile(FP_UDF_TCAM_FILE(version1), FP_UDF_TCAM_FILE(version2))
	util.DiffFile(FP_RANGE_CHECK_FILE(version1), FP_RANGE_CHECK_FILE(version2))
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

func FP_UDF_TCAM_FILE(version string) string {
	return "FP_UDF_TCAM." + version + ".txt"
}

func FP_UDF_OFFSET_FILE(version string) string {
	return "FP_UDF_OFFSET." + version + ".txt"
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

func FP_RANGE_CHECK_FILE(version string) string {
	return "FP_RANGE_CHECK_FILE." + version + ".txt"
}

func FP_SLICE_MAP_FILE(version string) string {
	return "FP_SLICE_MAP_FILE." + version + ".txt"
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

func (rdb *RuleDB) RuleAdd(dev *rut.RUT, name, flow, action, port, priority string) {
	dev.RunCommands(CTX, []*command.Command{
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
}

func (rdb *RuleDB) RuleDel(dev *rut.RUT, name, flow, action, port, priority string) {
	dev.RunCommands(CTX, []*command.Command{
		&command.Command{Mode: "config", CMD: " interface " + port},
		&command.Command{Mode: "config-if", CMD: " no service-policy input " + name},
		&command.Command{Mode: "config-if", CMD: " exit"},
		&command.Command{Mode: "config", CMD: " no policy " + name},
		&command.Command{Mode: "config", CMD: " no policer " + name},
		&command.Command{Mode: "config", CMD: " no flow " + name},
	})
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

	_, err = dev.RunCommand(CTX, &command.Command{
		Mode: "normal",
		CMD:  " config terminal",
	})
	if err != nil {
		fmt.Println(err)
	}

	_, err = dev.RunCommand(CTX, &command.Command{
		Mode: "config",
		CMD:  " show running-config",
	})
	if err != nil {
		fmt.Println(err)
	}

	DB.Dump(dev, "before.txt")
	/*
			DB.AnalysisRule(dev, "ether_any", "ethtype any", "copy-to-cpu", "gigabitethernet 8/2", "high")
			DB.AnalysisRule(dev, "ether_8765", "ethtype 0x8765", "copy-to-cpu", "gigabitethernet 8/2", "high")
			DB.AnalysisRule(dev, "length_1234", "length 1234", "copy-to-cpu", "gigabitethernet 8/2", "high")
			DB.AnalysisRule(dev, "inner_vid_2345", "inner-vid 1234", "copy-to-cpu", "gigabitethernet 8/2", "high")
			DB.AnalysisRule(dev, "outer_vid_2345", "outer-vid 1234", "copy-to-cpu", "gigabitethernet 8/2", "high")
			DB.AnalysisRule(dev, "cos_6", "cos 6", "copy-to-cpu", "gigabitethernet 8/2", "high")
			DB.AnalysisRule(dev, "mac_any_any", "mac any any", "copy-to-cpu", "gigabitethernet 8/2", "high")
			DB.AnalysisRule(dev, "mac_08_07", "mac 08:08:08:08:08:08 07:07:07:07:07:07", "copy-to-cpu", "gigabitethernet 8/2", "high")
			DB.AnalysisRule(dev, "ip_any_any", "ip any any", "copy-to-cpu", "gigabitethernet 8/2", "high")
			DB.AnalysisRule(dev, "ip_50_40", "ip 50.50.50.50 40.40.40.40", "deny", "gigabitethernet 8/2", "high")
		DB.AnalysisRule(dev, "ip_50_40_proto_123", "ip 50.50.50.50 40.40.40.40 123", "deny", "gigabitethernet 8/2", "high")
		DB.AnalysisRule(dev, "ip_50_40_tcp_100_200", "ip 50.50.50.50 40.40.40.40 tcp 100 200", "deny", "gigabitethernet 8/2", "high")
		DB.AnalysisRule(dev, "ip_50_40_udp_400_200", "ip 50.50.50.50 40.40.40.40 udp 400 200", "deny", "gigabitethernet 8/2", "high")
		DB.AnalysisRule(dev, "ip_50_40_udp_300_400", "ip 50.50.50.50 40.40.40.40 udp 300 400", "deny", "gigabitethernet 8/2", "high")
		DB.AnalysisRule(dev, "ip_50_40_udp_500_600", "ip 50.50.50.50 40.40.40.40 udp 500 600", "deny", "gigabitethernet 8/2", "high")
		DB.AnalysisRule(dev, "ip_50_40_udp_800_700", "ip 50.50.50.50 40.40.40.40 udp 800 700", "deny", "gigabitethernet 8/2", "high")
		DB.AnalysisRule(dev, "ip_50_40_udp_900_20", "ip 50.50.50.50 40.40.40.40 udp 900 20", "deny", "gigabitethernet 8/2", "high")
		DB.AnalysisRule(dev, "ip_50_40", "ip 50.50.50.50 40.40.40.40", "deny", "gigabitethernet 8/2", "high")
	*/
	/*
		DB.AnalysisRule(dev, "ipv6_1000_2000", "ipv6 2001:db8::1000 2001:db8::2000", "copy-to-cpu", "gigabitethernet 8/2", "high")
		DB.AnalysisRule(dev, "ipv6_1000_2000_nh_50", "ipv6 2001:db8::1000 2001:db8::2000 50", "copy-to-cpu", "gigabitethernet 8/2", "high")
		DB.AnalysisRule(dev, "ipv6_1000_2000_nh_51", "ipv6 2001:db8::1000 2001:db8::2000 51", "copy-to-cpu", "gigabitethernet 8/2", "high")
		DB.AnalysisRule(dev, "ipv6_1000_2000_nh_52", "ipv6 2001:db8::1000 2001:db8::2000 52", "copy-to-cpu", "gigabitethernet 8/2", "high")
		DB.AnalysisRule(dev, "ipv6_1000_2000_nh_53", "ipv6 2001:db8::1000 2001:db8::2000 53", "copy-to-cpu", "gigabitethernet 8/2", "high")
		DB.AnalysisRule(dev, "ipv6_1000_2000_nh_54", "ipv6 2001:db8::1000 2001:db8::2000 54", "copy-to-cpu", "gigabitethernet 8/2", "high")
		DB.AnalysisRule(dev, "ipv6_1000_2000_tcp_100_200", "ipv6 2001:db8::1000 2001:db8::2000 tcp 100 200", "copy-to-cpu", "gigabitethernet 8/2", "high")
		DB.AnalysisRule(dev, "ipv6_1000_2000_tcp_300_400", "ipv6 2001:db8::1000 2001:db8::2000 tcp 300 400", "copy-to-cpu", "gigabitethernet 8/2", "high")
	*/
	DB.AnalysisRule(dev, "ipv6_1000_2000_udp_100_200", "ipv6 2001:db8::1000 2001:db8::2000 udp 100-200 3004-400", "copy-to-cpu", "gigabitethernet 8/2", "high")
	DB.AnalysisRule(dev, "ipv6_1000_2000_udp_500_600", "ipv6 2001:db8::1000 2001:db8::2000 udp 500-700 600", "copy-to-cpu", "gigabitethernet 8/2", "high")
	DB.AnalysisRule(dev, "ipv6_1000_2000_tcp_500_600", "ipv6 2001:db8::1000 2001:db8::2000 tcp 500 600-800", "copy-to-cpu", "gigabitethernet 8/2", "high")
	DB.AnalysisRule(dev, "ipv6_1000_2000_tcp_500", "ipv6 2001:db8::1000 2001:db8::2000 tcp 500 any", "copy-to-cpu", "gigabitethernet 8/2", "high")
	DB.Dump(dev, "after.txt")
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
