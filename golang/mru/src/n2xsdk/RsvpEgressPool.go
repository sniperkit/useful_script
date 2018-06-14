package n2xsdk

type RsvpEgressPool struct {
 Handler string
}

func(np *RsvpEgressPool) IsEgress () error {
 //parameters: PoolHandle
 //AgtRsvpEgressPool IsEgress, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) ConfigureAsSingleEndPoint () error {
 //parameters: PoolHandle IpSourceAddress IpDestinationAddress NumberTunnels NumberLsps
 //AgtRsvpEgressPool ConfigureAsSingleEndPoint, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) GetSingleEndPointDetails ()(string, error) {
 //parameters: PoolHandle
 //AgtRsvpEgressPool GetSingleEndPointDetails
 return "", nil
}

func(np *RsvpEgressPool) ConfigureAsMultipleEndPoint () error {
 //parameters: PoolHandle SourceIpAddressFirst SourceIpAddressIncrement SourceIpAddressNumber DestinationIpAddressFirst DestinationIpAddressIncrement DestinationIpAddressNumber PartialMesh
 //AgtRsvpEgressPool ConfigureAsMultipleEndPoint, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) GetMultipleEndPointDetails ()(string, error) {
 //parameters: PoolHandle
 //AgtRsvpEgressPool GetMultipleEndPointDetails
 return "", nil
}

func(np *RsvpEgressPool) ConfigureAsGridEndPoint () error {
 //parameters: PoolHandle SourceIpAddressFirst SourceIpAddressNumberOfRows SourceIpAddressNumberOfColumns DestinationIpAddressFirst DestinationIpAddressNumberOfRows DestinationIpAddressNumberOfColumns
 //AgtRsvpEgressPool ConfigureAsGridEndPoint, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) GetGridEndPointDetails ()(string, error) {
 //parameters: PoolHandle
 //AgtRsvpEgressPool GetGridEndPointDetails
 return "", nil
}

func(np *RsvpEgressPool) GetEndPointType ()(string, error) {
 //parameters: PoolHandle
 //AgtRsvpEgressPool GetEndPointType
 return "", nil
}

func(np *RsvpEgressPool) SetFirstTunnelId () error {
 //parameters: PoolHandle FirstTunnelId
 //AgtRsvpEgressPool SetFirstTunnelId, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) GetFirstTunnelId ()(string, error) {
 //parameters: PoolHandle
 //AgtRsvpEgressPool GetFirstTunnelId
 return "", nil
}

func(np *RsvpEgressPool) SetFirstLspId () error {
 //parameters: PoolHandle FirstLspId
 //AgtRsvpEgressPool SetFirstLspId, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) GetFirstLspId ()(string, error) {
 //parameters: PoolHandle
 //AgtRsvpEgressPool GetFirstLspId
 return "", nil
}

func(np *RsvpEgressPool) SetBurstInterval () error {
 //parameters: PoolHandle BurstInterval
 //AgtRsvpEgressPool SetBurstInterval, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) GetBurstInterval ()(string, error) {
 //parameters: PoolHandle
 //AgtRsvpEgressPool GetBurstInterval
 return "", nil
}

func(np *RsvpEgressPool) SetBurstSize () error {
 //parameters: PoolHandle BurstSize
 //AgtRsvpEgressPool SetBurstSize, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) GetBurstSize ()(string, error) {
 //parameters: PoolHandle
 //AgtRsvpEgressPool GetBurstSize
 return "", nil
}

func(np *RsvpEgressPool) Close () error {
 //parameters: PoolHandle
 //AgtRsvpEgressPool Close, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) CloseGracefully () error {
 //parameters: PoolHandle
 //AgtRsvpEgressPool CloseGracefully, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) CloseRemotely () error {
 //parameters: PoolHandle
 //AgtRsvpEgressPool CloseRemotely, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) CloseLsp () error {
 //parameters: PoolHandle IpSourceAddress IpDestinationAddress TunnelId LspId
 //AgtRsvpEgressPool CloseLsp, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) CloseLspGracefully () error {
 //parameters: PoolHandle IpSourceAddress IpDestinationAddress TunnelId LspId
 //AgtRsvpEgressPool CloseLspGracefully, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) CloseLspRemotely () error {
 //parameters: PoolHandle IpSourceAddress IpDestinationAddress TunnelId LspId
 //AgtRsvpEgressPool CloseLspRemotely, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) CloseTunnel () error {
 //parameters: PoolHandle IpSourceAddress IpDestinationAddress TunnelId
 //AgtRsvpEgressPool CloseTunnel, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) CloseTunnelGracefully () error {
 //parameters: PoolHandle IpSourceAddress IpDestinationAddress TunnelId
 //AgtRsvpEgressPool CloseTunnelGracefully, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) CloseTunnelRemotely () error {
 //parameters: PoolHandle IpSourceAddress IpDestinationAddress TunnelId
 //AgtRsvpEgressPool CloseTunnelRemotely, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) GetState ()(string, error) {
 //parameters: PoolHandle
 //AgtRsvpEgressPool GetState
 return "", nil
}

func(np *RsvpEgressPool) UploadTunnelDetails () error {
 //parameters: PoolHandle
 //AgtRsvpEgressPool UploadTunnelDetails, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) CancelUpload () error {
 //parameters: PoolHandle
 //AgtRsvpEgressPool CancelUpload, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) IsUploadComplete () error {
 //parameters: PoolHandle
 //AgtRsvpEgressPool IsUploadComplete, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) SetLspsPerUploadPage () error {
 //parameters: PoolHandle LspsPerPage
 //AgtRsvpEgressPool SetLspsPerUploadPage, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) GetLspsPerUploadPage ()(string, error) {
 //parameters: PoolHandle
 //AgtRsvpEgressPool GetLspsPerUploadPage
 return "", nil
}

func(np *RsvpEgressPool) GetNumberOfUploadedPages ()(string, error) {
 //parameters: PoolHandle
 //AgtRsvpEgressPool GetNumberOfUploadedPages
 return "", nil
}

func(np *RsvpEgressPool) GetTunnelLabels ()(string, error) {
 //parameters: PoolHandle PageNumber
 //AgtRsvpEgressPool GetTunnelLabels
 return "", nil
}

func(np *RsvpEgressPool) GetRequiredTunnelLabels ()(string, error) {
 //parameters: PoolHandle IpSourceAddressSize pRequiredSources IpDestinationAddressSize pRequiredDestinations TunnelIdSize pRequiredTunnelIds
 //AgtRsvpEgressPool GetRequiredTunnelLabels
 return "", nil
}

func(np *RsvpEgressPool) GetRequiredLspLabels ()(string, error) {
 //parameters: PoolHandle RequiredSourceAddress RequiredDestinationAddress RequiredTunnelId RequiredLspId
 //AgtRsvpEgressPool GetRequiredLspLabels
 return "", nil
}

func(np *RsvpEgressPool) GetTunnelLabelInformation ()(string, error) {
 //parameters: PoolHandle PageNumber
 //AgtRsvpEgressPool GetTunnelLabelInformation
 return "", nil
}

func(np *RsvpEgressPool) GetRequiredTunnelLabelInformation ()(string, error) {
 //parameters: PoolHandle IpSourceAddressSize pRequiredSources IpDestinationAddressSize pRequiredDestinations TunnelIdSize pRequiredTunnelIds
 //AgtRsvpEgressPool GetRequiredTunnelLabelInformation
 return "", nil
}

func(np *RsvpEgressPool) GetRequiredLspLabelInformation ()(string, error) {
 //parameters: PoolHandle RequiredSourceAddress RequiredDestinationAddress RequiredTunnelId RequiredLspId
 //AgtRsvpEgressPool GetRequiredLspLabelInformation
 return "", nil
}

func(np *RsvpEgressPool) SetBandwidthReservation () error {
 //parameters: PoolHandle Bandwidth
 //AgtRsvpEgressPool SetBandwidthReservation, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) GetBandwidthReservation ()(string, error) {
 //parameters: PoolHandle
 //AgtRsvpEgressPool GetBandwidthReservation
 return "", nil
}

func(np *RsvpEgressPool) SetPriority () error {
 //parameters: PoolHandle SetupPriority HoldPriority
 //AgtRsvpEgressPool SetPriority, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) GetPriority ()(string, error) {
 //parameters: PoolHandle
 //AgtRsvpEgressPool GetPriority
 return "", nil
}

func(np *RsvpEgressPool) SetAttributesFlag () error {
 //parameters: PoolHandle Flag
 //AgtRsvpEgressPool SetAttributesFlag, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) GetAttributesFlag ()(string, error) {
 //parameters: PoolHandle
 //AgtRsvpEgressPool GetAttributesFlag
 return "", nil
}

func(np *RsvpEgressPool) SetResourceAffinities () error {
 //parameters: PoolHandle ExcludeAny IncludeAny IncludeAll
 //AgtRsvpEgressPool SetResourceAffinities, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) GetResourceAffinities ()(string, error) {
 //parameters: PoolHandle
 //AgtRsvpEgressPool GetResourceAffinities
 return "", nil
}

func(np *RsvpEgressPool) EnableResourceAffinities () error {
 //parameters: PoolHandle
 //AgtRsvpEgressPool EnableResourceAffinities, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) DisableResourceAffinities () error {
 //parameters: PoolHandle
 //AgtRsvpEgressPool DisableResourceAffinities, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) IsResourceAffinitiesEnabled () error {
 //parameters: PoolHandle
 //AgtRsvpEgressPool IsResourceAffinitiesEnabled, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) SetSessionName () error {
 //parameters: PoolHandle Name
 //AgtRsvpEgressPool SetSessionName, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) GetSessionName ()(string, error) {
 //parameters: PoolHandle
 //AgtRsvpEgressPool GetSessionName
 return "", nil
}

func(np *RsvpEgressPool) EnableMakeBeforeBreak () error {
 //parameters: PoolHandle
 //AgtRsvpEgressPool EnableMakeBeforeBreak, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) DisableMakeBeforeBreak () error {
 //parameters: PoolHandle
 //AgtRsvpEgressPool DisableMakeBeforeBreak, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) IsMakeBeforeBreakEnabled () error {
 //parameters: PoolHandle
 //AgtRsvpEgressPool IsMakeBeforeBreakEnabled, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) EnableAutoMakeBeforeBreak () error {
 //parameters: PoolHandle
 //AgtRsvpEgressPool EnableAutoMakeBeforeBreak, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) DisableAutoMakeBeforeBreak () error {
 //parameters: PoolHandle
 //AgtRsvpEgressPool DisableAutoMakeBeforeBreak, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) IsAutoMakeBeforeBreakEnabled () error {
 //parameters: PoolHandle
 //AgtRsvpEgressPool IsAutoMakeBeforeBreakEnabled, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) LspPoolMakeBeforeBreak () error {
 //parameters: PoolHandle LspType
 //AgtRsvpEgressPool LspPoolMakeBeforeBreak, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) LspMakeBeforeBreak () error {
 //parameters: PoolHandle AddressPoolIndex LspType
 //AgtRsvpEgressPool LspMakeBeforeBreak, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) RemoveRsvpEroSubobject () error {
 //parameters: PoolHandle LspType Index
 //AgtRsvpEgressPool RemoveRsvpEroSubobject, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) RemoveAllRsvpEroSubobjects () error {
 //parameters: PoolHandle LspType
 //AgtRsvpEgressPool RemoveAllRsvpEroSubobjects, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) SetRsvpEroSubobject () error {
 //parameters: PoolHandle LspType Index EroSubobjectType IsLoose Count psaContents
 //AgtRsvpEgressPool SetRsvpEroSubobject, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) SetRsvpEroIpv4Subobject () error {
 //parameters: PoolHandle LspType Index IsLoose Address Prefix
 //AgtRsvpEgressPool SetRsvpEroIpv4Subobject, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) SetRsvpEroLabelSubobject () error {
 //parameters: PoolHandle LspType Index Upstream CType IsLoose Label
 //AgtRsvpEgressPool SetRsvpEroLabelSubobject, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) SetRsvpEroAsNumberSubobject () error {
 //parameters: PoolHandle LspType Index IsLoose AsNumber
 //AgtRsvpEgressPool SetRsvpEroAsNumberSubobject, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) SetRsvpEroUnnumberedSubobject () error {
 //parameters: PoolHandle LspType Index IsLoose RouterId InterfaceId
 //AgtRsvpEgressPool SetRsvpEroUnnumberedSubobject, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) GetRsvpEroSubobject ()(string, error) {
 //parameters: PoolHandle LspType Index
 //AgtRsvpEgressPool GetRsvpEroSubobject
 return "", nil
}

func(np *RsvpEgressPool) GetRsvpEroIpv4Subobject ()(string, error) {
 //parameters: PoolHandle LspType Index
 //AgtRsvpEgressPool GetRsvpEroIpv4Subobject
 return "", nil
}

func(np *RsvpEgressPool) GetRsvpEroLabelSubobject ()(string, error) {
 //parameters: PoolHandle LspType Index
 //AgtRsvpEgressPool GetRsvpEroLabelSubobject
 return "", nil
}

func(np *RsvpEgressPool) GetRsvpEroAsNumberSubobject ()(string, error) {
 //parameters: PoolHandle LspType Index
 //AgtRsvpEgressPool GetRsvpEroAsNumberSubobject
 return "", nil
}

func(np *RsvpEgressPool) GetRsvpEroUnnumberedSubobject ()(string, error) {
 //parameters: PoolHandle LspType Index
 //AgtRsvpEgressPool GetRsvpEroUnnumberedSubobject
 return "", nil
}

func(np *RsvpEgressPool) ListRsvpEroSubobjects ()(string, error) {
 //parameters: PoolHandle LspType
 //AgtRsvpEgressPool ListRsvpEroSubobjects
 return "", nil
}

func(np *RsvpEgressPool) GetNumberOfRsvpEroSubobjects ()(string, error) {
 //parameters: PoolHandle LspType
 //AgtRsvpEgressPool GetNumberOfRsvpEroSubobjects
 return "", nil
}

func(np *RsvpEgressPool) GetRsvpEroSubobjectType ()(string, error) {
 //parameters: PoolHandle LspType Index
 //AgtRsvpEgressPool GetRsvpEroSubobjectType
 return "", nil
}

func(np *RsvpEgressPool) SetTSpecIeeeParameter () error {
 //parameters: PoolHandle TSpecParameter FloatValue
 //AgtRsvpEgressPool SetTSpecIeeeParameter, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) GetTSpecIeeeParameter ()(string, error) {
 //parameters: PoolHandle TSpecParameter
 //AgtRsvpEgressPool GetTSpecIeeeParameter
 return "", nil
}

func(np *RsvpEgressPool) SetTSpecParameter () error {
 //parameters: PoolHandle TSpecParameter Value
 //AgtRsvpEgressPool SetTSpecParameter, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) GetTSpecParameter ()(string, error) {
 //parameters: PoolHandle TSpecParameter
 //AgtRsvpEgressPool GetTSpecParameter
 return "", nil
}

func(np *RsvpEgressPool) SetAdSpecParameter () error {
 //parameters: PoolHandle AdSpecParameter Value
 //AgtRsvpEgressPool SetAdSpecParameter, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) GetAdSpecParameter ()(string, error) {
 //parameters: PoolHandle AdSpecParameter
 //AgtRsvpEgressPool GetAdSpecParameter
 return "", nil
}

func(np *RsvpEgressPool) SetProtectionParameter () error {
 //parameters: PoolHandle ProtectionParameter Value
 //AgtRsvpEgressPool SetProtectionParameter, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) GetProtectionParameter ()(string, error) {
 //parameters: PoolHandle ProtectionParameter
 //AgtRsvpEgressPool GetProtectionParameter
 return "", nil
}

func(np *RsvpEgressPool) SetGeneralizedUniParameter () error {
 //parameters: PoolHandle GeneralizedUniParameter Value
 //AgtRsvpEgressPool SetGeneralizedUniParameter, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) GetGeneralizedUniParameter ()(string, error) {
 //parameters: PoolHandle GeneralizedUniParameter
 //AgtRsvpEgressPool GetGeneralizedUniParameter
 return "", nil
}

func(np *RsvpEgressPool) EnableProtectionObject () error {
 //parameters: PoolHandle
 //AgtRsvpEgressPool EnableProtectionObject, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) DisableProtectionObject () error {
 //parameters: PoolHandle
 //AgtRsvpEgressPool DisableProtectionObject, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) IsProtectionObjectEnabled () error {
 //parameters: PoolHandle
 //AgtRsvpEgressPool IsProtectionObjectEnabled, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) SelectAdSpecParameter () error {
 //parameters: PoolHandle AdSpecParameter
 //AgtRsvpEgressPool SelectAdSpecParameter, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) DeselectAdSpecParameter () error {
 //parameters: PoolHandle AdSpecParameter
 //AgtRsvpEgressPool DeselectAdSpecParameter, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) IsAdSpecParameterSelected () error {
 //parameters: PoolHandle AdSpecParameter
 //AgtRsvpEgressPool IsAdSpecParameterSelected, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) SelectGeneralizedUniParameter () error {
 //parameters: PoolHandle GeneralizedUniParameter
 //AgtRsvpEgressPool SelectGeneralizedUniParameter, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) DeselectGeneralizedUniParameter () error {
 //parameters: PoolHandle GeneralizedUniParameter
 //AgtRsvpEgressPool DeselectGeneralizedUniParameter, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) IsGeneralizedUniParameterSelected () error {
 //parameters: PoolHandle GeneralizedUniParameter
 //AgtRsvpEgressPool IsGeneralizedUniParameterSelected, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) SetDiffServObjectType () error {
 //parameters: PoolHandle DiffServObjectType
 //AgtRsvpEgressPool SetDiffServObjectType, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) GetDiffServObjectType ()(string, error) {
 //parameters: PoolHandle
 //AgtRsvpEgressPool GetDiffServObjectType
 return "", nil
}

func(np *RsvpEgressPool) EnableDiffServXBit () error {
 //parameters: PoolHandle
 //AgtRsvpEgressPool EnableDiffServXBit, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) DisableDiffServXBit () error {
 //parameters: PoolHandle
 //AgtRsvpEgressPool DisableDiffServXBit, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) IsDiffServXBitEnabled () error {
 //parameters: PoolHandle
 //AgtRsvpEgressPool IsDiffServXBitEnabled, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) SetDiffServPhbEncodingStyle () error {
 //parameters: PoolHandle EncodingStyle
 //AgtRsvpEgressPool SetDiffServPhbEncodingStyle, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) GetDiffServPhbEncodingStyle ()(string, error) {
 //parameters: PoolHandle
 //AgtRsvpEgressPool GetDiffServPhbEncodingStyle
 return "", nil
}

func(np *RsvpEgressPool) AddExpPhbIdMap () error {
 //parameters: PoolHandle Exp PhbId
 //AgtRsvpEgressPool AddExpPhbIdMap, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) RemoveExpPhbIdMap () error {
 //parameters: PoolHandle Exp PhbId
 //AgtRsvpEgressPool RemoveExpPhbIdMap, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) ListExpPhbIdMap ()(string, error) {
 //parameters: PoolHandle
 //AgtRsvpEgressPool ListExpPhbIdMap
 return "", nil
}

func(np *RsvpEgressPool) SetPhbSchedulingClass () error {
 //parameters: PoolHandle Psc
 //AgtRsvpEgressPool SetPhbSchedulingClass, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) GetPhbSchedulingClass ()(string, error) {
 //parameters: PoolHandle
 //AgtRsvpEgressPool GetPhbSchedulingClass
 return "", nil
}

func(np *RsvpEgressPool) SetPathObjectOrder () error {
 //parameters: PoolHandle Count psaObjects
 //AgtRsvpEgressPool SetPathObjectOrder, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) GetPathObjectOrder ()(string, error) {
 //parameters: PoolHandle
 //AgtRsvpEgressPool GetPathObjectOrder
 return "", nil
}

func(np *RsvpEgressPool) SetDefaultPathObjectOrder () error {
 //parameters: PoolHandle
 //AgtRsvpEgressPool SetDefaultPathObjectOrder, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) IsDefaultPathObjectOrder () error {
 //parameters: PoolHandle
 //AgtRsvpEgressPool IsDefaultPathObjectOrder, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) EnableUpstreamLabelOverride () error {
 //parameters: PoolHandle
 //AgtRsvpEgressPool EnableUpstreamLabelOverride, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) DisableUpstreamLabelOverride () error {
 //parameters: PoolHandle
 //AgtRsvpEgressPool DisableUpstreamLabelOverride, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) IsUpstreamLabelOverrideEnabled () error {
 //parameters: PoolHandle
 //AgtRsvpEgressPool IsUpstreamLabelOverrideEnabled, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) SetUnnumberedUpstreamLabel () error {
 //parameters: PoolHandle UpstreamLabel LocalInterfaceId
 //AgtRsvpEgressPool SetUnnumberedUpstreamLabel, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) GetUnnumberedUpstreamLabel ()(string, error) {
 //parameters: PoolHandle
 //AgtRsvpEgressPool GetUnnumberedUpstreamLabel
 return "", nil
}

func(np *RsvpEgressPool) SetNumberedUpstreamLabel () error {
 //parameters: PoolHandle UpstreamLabel LocalInterfaceId
 //AgtRsvpEgressPool SetNumberedUpstreamLabel, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) GetNumberedUpstreamLabel ()(string, error) {
 //parameters: PoolHandle
 //AgtRsvpEgressPool GetNumberedUpstreamLabel
 return "", nil
}

func(np *RsvpEgressPool) GetPortHandle ()(string, error) {
 //parameters: PoolHandle
 //AgtRsvpEgressPool GetPortHandle
 return "", nil
}

func(np *RsvpEgressPool) GetSessionHandle ()(string, error) {
 //parameters: PoolHandle
 //AgtRsvpEgressPool GetSessionHandle
 return "", nil
}

func(np *RsvpEgressPool) Enable () error {
 //parameters: PoolHandle
 //AgtRsvpEgressPool Enable, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) Disable () error {
 //parameters: PoolHandle
 //AgtRsvpEgressPool Disable, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) IsEnabled () error {
 //parameters: PoolHandle
 //AgtRsvpEgressPool IsEnabled, m.Object, m.Name);
 return nil
}

func(np *RsvpEgressPool) IsDefault () error {
 //parameters: PoolHandle
 //AgtRsvpEgressPool IsDefault, m.Object, m.Name);
 return nil
}

