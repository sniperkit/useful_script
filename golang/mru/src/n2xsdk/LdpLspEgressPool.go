package n2xsdk

type LdpLspEgressPool struct {
 Handler string
}

func(np *LdpLspEgressPool) SetNumberOfLsps () error {
 //parameters: PoolHandle NumLsps Modifier
 //AgtLdpLspEgressPool SetNumberOfLsps, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) GetNumberOfLsps ()(string, error) {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool GetNumberOfLsps
 return "", nil
}

func(np *LdpLspEgressPool) GetFecType ()(string, error) {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool GetFecType
 return "", nil
}

func(np *LdpLspEgressPool) SetPrefixFec () error {
 //parameters: PoolHandle FirstAddress Incr PfxLength IncrMask
 //AgtLdpLspEgressPool SetPrefixFec, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) GetPrefixFec ()(string, error) {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool GetPrefixFec
 return "", nil
}

func(np *LdpLspEgressPool) SetHostFec () error {
 //parameters: PoolHandle IpAddress Incr IncrMask
 //AgtLdpLspEgressPool SetHostFec, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) GetHostFec ()(string, error) {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool GetHostFec
 return "", nil
}

func(np *LdpLspEgressPool) SetCrLdpFec () error {
 //parameters: PoolHandle ActionFlag LspId IpAddress PfxLength IncrLspId
 //AgtLdpLspEgressPool SetCrLdpFec, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) GetCrLdpFec ()(string, error) {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool GetCrLdpFec
 return "", nil
}

func(np *LdpLspEgressPool) SetL2MplsFec () error {
 //parameters: PoolHandle CBit VcType GroupId VcId IncrVcId
 //AgtLdpLspEgressPool SetL2MplsFec, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) GetL2MplsFec ()(string, error) {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool GetL2MplsFec
 return "", nil
}

func(np *LdpLspEgressPool) SetPathVectorTLV () error {
 //parameters: PoolHandle PathCount psaContentList
 //AgtLdpLspEgressPool SetPathVectorTLV, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) GetPathVectorTLV ()(string, error) {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool GetPathVectorTLV
 return "", nil
}

func(np *LdpLspEgressPool) SetHopCountTLV () error {
 //parameters: PoolHandle HopCount
 //AgtLdpLspEgressPool SetHopCountTLV, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) GetHopCountTLV ()(string, error) {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool GetHopCountTLV
 return "", nil
}

func(np *LdpLspEgressPool) EnableDiffServXBit () error {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool EnableDiffServXBit, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) DisableDiffServXBit () error {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool DisableDiffServXBit, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) IsDiffServXBitEnabled () error {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool IsDiffServXBitEnabled, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) SetDiffServPhbEncodingStyle () error {
 //parameters: PoolHandle EncodingStyle
 //AgtLdpLspEgressPool SetDiffServPhbEncodingStyle, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) GetDiffServPhbEncodingStyle ()(string, error) {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool GetDiffServPhbEncodingStyle
 return "", nil
}

func(np *LdpLspEgressPool) EnableTLV () error {
 //parameters: PoolHandle TlvType
 //AgtLdpLspEgressPool EnableTLV, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) DisableTLV () error {
 //parameters: PoolHandle TlvType
 //AgtLdpLspEgressPool DisableTLV, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) IsTlvEnabled () error {
 //parameters: PoolHandle TlvType
 //AgtLdpLspEgressPool IsTlvEnabled, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) Open () error {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool Open, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) Close () error {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool Close, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) IsEnabled () error {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool IsEnabled, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) Cancel () error {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool Cancel, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) GetPoolState ()(string, error) {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool GetPoolState
 return "", nil
}

func(np *LdpLspEgressPool) GetNumberOfOpenedLsps ()(string, error) {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool GetNumberOfOpenedLsps
 return "", nil
}

func(np *LdpLspEgressPool) GetLspInfo ()(string, error) {
 //parameters: PoolHandle LspCount
 //AgtLdpLspEgressPool GetLspInfo
 return "", nil
}

func(np *LdpLspEgressPool) SetChunkCount () error {
 //parameters: PoolHandle ChunkCount
 //AgtLdpLspEgressPool SetChunkCount, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) GetChunkCount ()(string, error) {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool GetChunkCount
 return "", nil
}

func(np *LdpLspEgressPool) SetChunkDelay () error {
 //parameters: PoolHandle ChunkDelay
 //AgtLdpLspEgressPool SetChunkDelay, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) GetChunkDelay ()(string, error) {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool GetChunkDelay
 return "", nil
}

func(np *LdpLspEgressPool) SetFilterFecType () error {
 //parameters: PoolHandle FecType
 //AgtLdpLspEgressPool SetFilterFecType, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) SetFilterState () error {
 //parameters: PoolHandle State
 //AgtLdpLspEgressPool SetFilterState, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) SetFilterLabel () error {
 //parameters: PoolHandle Label
 //AgtLdpLspEgressPool SetFilterLabel, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) SetFilterDiffServType () error {
 //parameters: PoolHandle DiffServType
 //AgtLdpLspEgressPool SetFilterDiffServType, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) SetFilterLspId () error {
 //parameters: PoolHandle LspId
 //AgtLdpLspEgressPool SetFilterLspId, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) SetFilterSrcAddress () error {
 //parameters: PoolHandle SrcAddress
 //AgtLdpLspEgressPool SetFilterSrcAddress, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) SetFilterLen () error {
 //parameters: PoolHandle Len
 //AgtLdpLspEgressPool SetFilterLen, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) SetFilterRouterId () error {
 //parameters: PoolHandle RouterId
 //AgtLdpLspEgressPool SetFilterRouterId, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) SetFilterActFlag () error {
 //parameters: PoolHandle ActFlag
 //AgtLdpLspEgressPool SetFilterActFlag, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) SetFilterCbit () error {
 //parameters: PoolHandle CBit
 //AgtLdpLspEgressPool SetFilterCbit, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) SetFilterVcType () error {
 //parameters: PoolHandle VcType
 //AgtLdpLspEgressPool SetFilterVcType, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) SetFilterGroupId () error {
 //parameters: PoolHandle GroupId
 //AgtLdpLspEgressPool SetFilterGroupId, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) SetFilterVcId () error {
 //parameters: PoolHandle VcId
 //AgtLdpLspEgressPool SetFilterVcId, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) SetFilterL2MplsMtu () error {
 //parameters: PoolHandle MtuValue
 //AgtLdpLspEgressPool SetFilterL2MplsMtu, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) SetFilterL2MplsConcatAtmCell () error {
 //parameters: PoolHandle ConcatAtmCell
 //AgtLdpLspEgressPool SetFilterL2MplsConcatAtmCell, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) SetFilterL2MplsPayload () error {
 //parameters: PoolHandle CemPayload
 //AgtLdpLspEgressPool SetFilterL2MplsPayload, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) SetFilterL2MplsInterfaceString () error {
 //parameters: PoolHandle InterfaceString
 //AgtLdpLspEgressPool SetFilterL2MplsInterfaceString, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) SetFilterL2MplsOptions () error {
 //parameters: PoolHandle CemOptions
 //AgtLdpLspEgressPool SetFilterL2MplsOptions, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) SetL2MplsGenPwIdFec () error {
 //parameters: PoolHandle CBit VcType
 //AgtLdpLspEgressPool SetL2MplsGenPwIdFec, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) GetL2MplsGenPwIdFec ()(string, error) {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool GetL2MplsGenPwIdFec
 return "", nil
}

func(np *LdpLspEgressPool) SetNullLabelType () error {
 //parameters: PoolHandle NullLabelType
 //AgtLdpLspEgressPool SetNullLabelType, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) GetNullLabelType ()(string, error) {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool GetNullLabelType
 return "", nil
}

func(np *LdpLspEgressPool) SetEgressDiffServEtypeTLV () error {
 //parameters: PoolHandle ExpCount psaExp PhbldCount psaPhbId
 //AgtLdpLspEgressPool SetEgressDiffServEtypeTLV, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) GetEgressDiffServEtypeTLV ()(string, error) {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool GetEgressDiffServEtypeTLV
 return "", nil
}

func(np *LdpLspEgressPool) SetEgressDiffServLtypeTLV () error {
 //parameters: PoolHandle Psc
 //AgtLdpLspEgressPool SetEgressDiffServLtypeTLV, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) GetEgressDiffServLtypeTLV ()(string, error) {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool GetEgressDiffServLtypeTLV
 return "", nil
}

func(np *LdpLspEgressPool) SetL2MplsMtu () error {
 //parameters: PoolHandle MtuValue
 //AgtLdpLspEgressPool SetL2MplsMtu, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) GetL2MplsMtu ()(string, error) {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool GetL2MplsMtu
 return "", nil
}

func(np *LdpLspEgressPool) SetL2MplsConcatAtmCell () error {
 //parameters: PoolHandle ConcatAtmCell
 //AgtLdpLspEgressPool SetL2MplsConcatAtmCell, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) GetL2MplsConcatAtmCell ()(string, error) {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool GetL2MplsConcatAtmCell
 return "", nil
}

func(np *LdpLspEgressPool) SetL2MplsPayload () error {
 //parameters: PoolHandle CemPayload
 //AgtLdpLspEgressPool SetL2MplsPayload, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) GetL2MplsPayload ()(string, error) {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool GetL2MplsPayload
 return "", nil
}

func(np *LdpLspEgressPool) SetL2MplsInterfaceString () error {
 //parameters: PoolHandle InterfaceString
 //AgtLdpLspEgressPool SetL2MplsInterfaceString, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) GetL2MplsInterfaceString ()(string, error) {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool GetL2MplsInterfaceString
 return "", nil
}

func(np *LdpLspEgressPool) SetL2MplsOptions () error {
 //parameters: PoolHandle CemOptions
 //AgtLdpLspEgressPool SetL2MplsOptions, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) GetL2MplsOptions ()(string, error) {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool GetL2MplsOptions
 return "", nil
}

func(np *LdpLspEgressPool) EnableL2MplsInterfaceParameter () error {
 //parameters: PoolHandle InterfaceType
 //AgtLdpLspEgressPool EnableL2MplsInterfaceParameter, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) DisableL2MplsInterfaceParameter () error {
 //parameters: PoolHandle InterfaceType
 //AgtLdpLspEgressPool DisableL2MplsInterfaceParameter, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) IsL2MplsInterfaceParameterEnabled () error {
 //parameters: PoolHandle InterfaceType
 //AgtLdpLspEgressPool IsL2MplsInterfaceParameterEnabled, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) SetRequestedVlanId () error {
 //parameters: PoolHandle RequestedVlanId
 //AgtLdpLspEgressPool SetRequestedVlanId, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) GetRequestedVlanId ()(string, error) {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool GetRequestedVlanId
 return "", nil
}

func(np *LdpLspEgressPool) SetCemBitRate () error {
 //parameters: PoolHandle CemBitRate
 //AgtLdpLspEgressPool SetCemBitRate, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) GetCemBitRate ()(string, error) {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool GetCemBitRate
 return "", nil
}

func(np *LdpLspEgressPool) SetFrDlciLength () error {
 //parameters: PoolHandle FrDlciLength
 //AgtLdpLspEgressPool SetFrDlciLength, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) GetFrDlciLength ()(string, error) {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool GetFrDlciLength
 return "", nil
}

func(np *LdpLspEgressPool) SetFcsRetentionIndicator () error {
 //parameters: PoolHandle FcsRetentionIndicator
 //AgtLdpLspEgressPool SetFcsRetentionIndicator, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) GetFcsRetentionIndicator ()(string, error) {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool GetFcsRetentionIndicator
 return "", nil
}

func(np *LdpLspEgressPool) SetFragmentationIndicator () error {
 //parameters: PoolHandle FragmentationIndicator
 //AgtLdpLspEgressPool SetFragmentationIndicator, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) GetFragmentationIndicator ()(string, error) {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool GetFragmentationIndicator
 return "", nil
}

func(np *LdpLspEgressPool) SetTdmOption () error {
 //parameters: PoolHandle TdmOption
 //AgtLdpLspEgressPool SetTdmOption, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) GetTdmOption ()(string, error) {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool GetTdmOption
 return "", nil
}

func(np *LdpLspEgressPool) SetVccvParameter () error {
 //parameters: PoolHandle VccvParameter
 //AgtLdpLspEgressPool SetVccvParameter, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) GetVccvParameter ()(string, error) {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool GetVccvParameter
 return "", nil
}

func(np *LdpLspEgressPool) EnableL2MplsPwStatus () error {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool EnableL2MplsPwStatus, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) DisableL2MplsPwStatus () error {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool DisableL2MplsPwStatus, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) IsL2MplsPwStatusEnabled () error {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool IsL2MplsPwStatusEnabled, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) SetPwStatus () error {
 //parameters: PoolHandle PWStatusCode
 //AgtLdpLspEgressPool SetPwStatus, m.Object, m.Name);
 return nil
}

func(np *LdpLspEgressPool) GetPwStatus ()(string, error) {
 //parameters: PoolHandle
 //AgtLdpLspEgressPool GetPwStatus
 return "", nil
}

