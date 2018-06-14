package n2xsdk

type Bgp4VpnIpv4RouteProfile struct {
 Handler string
}

func(np *Bgp4VpnIpv4RouteProfile) SetRoutesPerVpn () error {
 //parameters: RouteProfileHandle NumRoutesPerVpn
 //AgtBgp4VpnIpv4RouteProfile SetRoutesPerVpn, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetRoutesPerVpn ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetRoutesPerVpn
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) CreateMatchingVpnVrfPool () error {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile CreateMatchingVpnVrfPool, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) EnableConnector () error {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile EnableConnector, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) DisableConnector () error {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile DisableConnector, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) IsConnectorEnabled () error {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile IsConnectorEnabled, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) EnableAutoConnectorIpAddress () error {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile EnableAutoConnectorIpAddress, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) DisableAutoConnectorIpAddress () error {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile DisableAutoConnectorIpAddress, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) IsAutoConnectorIpAddressEnabled () error {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile IsAutoConnectorIpAddressEnabled, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetConnectorIpAddressAuto ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetConnectorIpAddressAuto
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetConnectorRouteDistinguisherIncrementingRange () error {
 //parameters: RouteProfileHandle FirstRouteDistinguisher RouteDistinguisherIncrement RouteDistinguisherRepeat PercentOverlap
 //AgtBgp4VpnIpv4RouteProfile SetConnectorRouteDistinguisherIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetConnectorRouteDistinguisherIncrementingRange ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetConnectorRouteDistinguisherIncrementingRange
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetConnectorIpAddressIncrementingRange () error {
 //parameters: RouteProfileHandle FirstIpAddress PrefixLength AddressIncrement AddressRepeat
 //AgtBgp4VpnIpv4RouteProfile SetConnectorIpAddressIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetConnectorIpAddressIncrementingRange ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetConnectorIpAddressIncrementingRange
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetConnectorRouteDistinguisherList ()(string, error) {
 //parameters: RouteProfileHandle RouteDistinguisherList
 //AgtBgp4VpnIpv4RouteProfile SetConnectorRouteDistinguisherList
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetConnectorRouteDistinguisherList ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetConnectorRouteDistinguisherList
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetConnectorIpAddressList ()(string, error) {
 //parameters: RouteProfileHandle IpAddressList
 //AgtBgp4VpnIpv4RouteProfile SetConnectorIpAddressList
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetConnectorIpAddressList ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetConnectorIpAddressList
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetConnectorRouteDistinguisher () error {
 //parameters: RouteProfileHandle RouteDistinguisher
 //AgtBgp4VpnIpv4RouteProfile SetConnectorRouteDistinguisher, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetConnectorRouteDistinguisher ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetConnectorRouteDistinguisher
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetConnectorIpAddress () error {
 //parameters: RouteProfileHandle IpAddress
 //AgtBgp4VpnIpv4RouteProfile SetConnectorIpAddress, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetConnectorIpAddress ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetConnectorIpAddress
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetExportRouteTargetType () error {
 //parameters: RouteProfileHandle RouteTargetTypeValue
 //AgtBgp4VpnIpv4RouteProfile SetExportRouteTargetType, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetExportRouteTargetType ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetExportRouteTargetType
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetRouteDistinguisherType () error {
 //parameters: RouteProfileHandle RouteDistinguisherTypeValue
 //AgtBgp4VpnIpv4RouteProfile SetRouteDistinguisherType, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetRouteDistinguisherType ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetRouteDistinguisherType
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) EnableAutoRouteDistinguisher () error {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile EnableAutoRouteDistinguisher, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) DisableAutoRouteDistinguisher () error {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile DisableAutoRouteDistinguisher, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) IsAutoRouteDistinguisherEnabled () error {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile IsAutoRouteDistinguisherEnabled, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetVpnsPerPeer () error {
 //parameters: RouteProfileHandle NumVpnsPerPeer
 //AgtBgp4VpnIpv4RouteProfile SetVpnsPerPeer, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetVpnsPerPeer ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetVpnsPerPeer
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetTotalNumberOfVpns ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetTotalNumberOfVpns
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetExportRouteTargetIncrementingRange () error {
 //parameters: RouteProfileHandle RouteTargetType FirstRouteTarget RouteTargetIncrement RouteTargetRepeat PercentOverlap
 //AgtBgp4VpnIpv4RouteProfile SetExportRouteTargetIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetExportRouteTargetIncrementingRange ()(string, error) {
 //parameters: RouteProfileHandle RouteTargetType
 //AgtBgp4VpnIpv4RouteProfile GetExportRouteTargetIncrementingRange
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetExportRouteTargetList ()(string, error) {
 //parameters: RouteProfileHandle RouteTargetType RouteTargetList
 //AgtBgp4VpnIpv4RouteProfile SetExportRouteTargetList
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetExportRouteTargetList ()(string, error) {
 //parameters: RouteProfileHandle RouteTargetType
 //AgtBgp4VpnIpv4RouteProfile GetExportRouteTargetList
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetExportRouteTarget () error {
 //parameters: RouteProfileHandle RouteTargetType RouteTarget
 //AgtBgp4VpnIpv4RouteProfile SetExportRouteTarget, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetExportRouteTarget ()(string, error) {
 //parameters: RouteProfileHandle RouteTargetType
 //AgtBgp4VpnIpv4RouteProfile GetExportRouteTarget
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetRouteDistinguisherIncrementingRange () error {
 //parameters: RouteProfileHandle RouteDistinguisherType FirstRouteDistinguisher RouteDistinguisherIncrement RouteDistinguisherRepeat PercentOverlap
 //AgtBgp4VpnIpv4RouteProfile SetRouteDistinguisherIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetRouteDistinguisherIncrementingRange ()(string, error) {
 //parameters: RouteProfileHandle RouteDistinguisherType
 //AgtBgp4VpnIpv4RouteProfile GetRouteDistinguisherIncrementingRange
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetRouteDistinguisherAuto ()(string, error) {
 //parameters: RouteProfileHandle RouteDistinguisherType
 //AgtBgp4VpnIpv4RouteProfile GetRouteDistinguisherAuto
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetRouteDistinguisherList ()(string, error) {
 //parameters: RouteProfileHandle RouteDistinguisherType RouteDistinguisherList
 //AgtBgp4VpnIpv4RouteProfile SetRouteDistinguisherList
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetRouteDistinguisherList ()(string, error) {
 //parameters: RouteProfileHandle RouteDistinguisherType
 //AgtBgp4VpnIpv4RouteProfile GetRouteDistinguisherList
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetRouteDistinguisher () error {
 //parameters: RouteProfileHandle RouteDistinguisherType RouteDistinguisher
 //AgtBgp4VpnIpv4RouteProfile SetRouteDistinguisher, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetRouteDistinguisher ()(string, error) {
 //parameters: RouteProfileHandle RouteDistinguisherType
 //AgtBgp4VpnIpv4RouteProfile GetRouteDistinguisher
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) EnableAutoPathAttribute () error {
 //parameters: RouteProfileHandle PathAttributeAuto
 //AgtBgp4VpnIpv4RouteProfile EnableAutoPathAttribute, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) DisableAutoPathAttribute () error {
 //parameters: RouteProfileHandle PathAttributeAuto
 //AgtBgp4VpnIpv4RouteProfile DisableAutoPathAttribute, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) IsAutoPathAttributeEnabled () error {
 //parameters: RouteProfileHandle PathAttributeAuto
 //AgtBgp4VpnIpv4RouteProfile IsAutoPathAttributeEnabled, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetNextHopIncrementingRange () error {
 //parameters: RouteProfileHandle FirstNextHop PrefixLength NextHopIncrement NextHopRepeat
 //AgtBgp4VpnIpv4RouteProfile SetNextHopIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetNextHopIncrementingRange ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetNextHopIncrementingRange
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetNextHopList ()(string, error) {
 //parameters: RouteProfileHandle NextHopList
 //AgtBgp4VpnIpv4RouteProfile SetNextHopList
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetNextHopList ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetNextHopList
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetNextHop () error {
 //parameters: RouteProfileHandle NextHop
 //AgtBgp4VpnIpv4RouteProfile SetNextHop, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetNextHop ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetNextHop
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetNextHopAuto ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetNextHopAuto
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetRoutesIncrementingRange () error {
 //parameters: RouteProfileHandle FirstRoute PrefixLengthIpv4 RouteIncrement PercentOverlap
 //AgtBgp4VpnIpv4RouteProfile SetRoutesIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetRoutesIncrementingRange ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetRoutesIncrementingRange
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetRoutesPrefixLength () error {
 //parameters: RouteProfileHandle PrefixLengthIpv4
 //AgtBgp4VpnIpv4RouteProfile SetRoutesPrefixLength, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetRoutesPrefixLength ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetRoutesPrefixLength
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetRoutesSubRangeOffsets () error {
 //parameters: RouteProfileHandle FirstRoute MsbOffsetList
 //AgtBgp4VpnIpv4RouteProfile SetRoutesSubRangeOffsets, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetRoutesSubRangeOffsets ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetRoutesSubRangeOffsets
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetRoutesSubRange () error {
 //parameters: RouteProfileHandle SubRangeInstance RouteIncrement Count Repeat
 //AgtBgp4VpnIpv4RouteProfile SetRoutesSubRange, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetRoutesSubRange ()(string, error) {
 //parameters: RouteProfileHandle SubRangeInstance
 //AgtBgp4VpnIpv4RouteProfile GetRoutesSubRange
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetRoutesPerPeer () error {
 //parameters: RouteProfileHandle RoutesPerPeer
 //AgtBgp4VpnIpv4RouteProfile SetRoutesPerPeer, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetRoutesPerPeer ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetRoutesPerPeer
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetTotalRoutesInProfile ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetTotalRoutesInProfile
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetNthRouteInProfile ()(string, error) {
 //parameters: RouteProfileHandle RouteProfileIndex
 //AgtBgp4VpnIpv4RouteProfile GetNthRouteInProfile
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetNumAdvertisedRoutes ()(string, error) {
 //parameters: RouteProfileInstance
 //AgtBgp4VpnIpv4RouteProfile GetNumAdvertisedRoutes
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetNumWithdrawnRoutes ()(string, error) {
 //parameters: RouteProfileInstance
 //AgtBgp4VpnIpv4RouteProfile GetNumWithdrawnRoutes
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) EnableLabeling () error {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile EnableLabeling, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) DisableLabeling () error {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile DisableLabeling, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) IsLabelingEnabled () error {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile IsLabelingEnabled, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetLabelingMode () error {
 //parameters: RouteProfileHandle LabelingMode
 //AgtBgp4VpnIpv4RouteProfile SetLabelingMode, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetLabelingMode ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetLabelingMode
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetUserLabelType () error {
 //parameters: RouteProfileHandle UserLabelType
 //AgtBgp4VpnIpv4RouteProfile SetUserLabelType, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetUserLabelType ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetUserLabelType
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetUserLabelIncrementingRange () error {
 //parameters: RouteProfileHandle StartLabel LabelIncrement LabelRepeat
 //AgtBgp4VpnIpv4RouteProfile SetUserLabelIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetUserLabelIncrementingRange ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetUserLabelIncrementingRange
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetUserLabelList ()(string, error) {
 //parameters: RouteProfileHandle LabelList
 //AgtBgp4VpnIpv4RouteProfile SetUserLabelList
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetUserLabelList ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetUserLabelList
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetUserLabel () error {
 //parameters: RouteProfileHandle Label
 //AgtBgp4VpnIpv4RouteProfile SetUserLabel, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetUserLabel ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetUserLabel
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetLabels ()(string, error) {
 //parameters: RouteProfileInstance
 //AgtBgp4VpnIpv4RouteProfile GetLabels
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetLabelPoolState ()(string, error) {
 //parameters: RouteProfileInstance
 //AgtBgp4VpnIpv4RouteProfile GetLabelPoolState
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetPeerPoolHandle ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetPeerPoolHandle
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetState ()(string, error) {
 //parameters: RouteProfileInstance
 //AgtBgp4VpnIpv4RouteProfile GetState
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) IsEnabled () error {
 //parameters: RouteProfileInstance
 //AgtBgp4VpnIpv4RouteProfile IsEnabled, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) EnableAsPath () error {
 //parameters: RouteProfileHandle AsPathAttributeType
 //AgtBgp4VpnIpv4RouteProfile EnableAsPath, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) Enable4ByteAsPath () error {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeType
 //AgtBgp4VpnIpv4RouteProfile Enable4ByteAsPath, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) DisableAsPath () error {
 //parameters: RouteProfileHandle AsPathAttributeType
 //AgtBgp4VpnIpv4RouteProfile DisableAsPath, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) Disable4ByteAsPath () error {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeType
 //AgtBgp4VpnIpv4RouteProfile Disable4ByteAsPath, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) IsAsPathEnabled () error {
 //parameters: RouteProfileHandle AsPathAttributeType
 //AgtBgp4VpnIpv4RouteProfile IsAsPathEnabled, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) Is4ByteAsPathEnabled () error {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeType
 //AgtBgp4VpnIpv4RouteProfile Is4ByteAsPathEnabled, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) EnableAutoAsPath () error {
 //parameters: RouteProfileHandle AsPathAttributeType
 //AgtBgp4VpnIpv4RouteProfile EnableAutoAsPath, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) EnableAuto4ByteAsPath () error {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeType
 //AgtBgp4VpnIpv4RouteProfile EnableAuto4ByteAsPath, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) DisableAutoAsPath () error {
 //parameters: RouteProfileHandle AsPathAttributeType
 //AgtBgp4VpnIpv4RouteProfile DisableAutoAsPath, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) DisableAuto4ByteAsPath () error {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeType
 //AgtBgp4VpnIpv4RouteProfile DisableAuto4ByteAsPath, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) IsAutoAsPathEnabled () error {
 //parameters: RouteProfileHandle AsPathAttributeType
 //AgtBgp4VpnIpv4RouteProfile IsAutoAsPathEnabled, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) IsAuto4ByteAsPathEnabled () error {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeType
 //AgtBgp4VpnIpv4RouteProfile IsAuto4ByteAsPathEnabled, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetAsPathIncrementingRange () error {
 //parameters: RouteProfileHandle AsPathAttributeType FirstAsNumber AsNumberIncrement AsNumberRepeat
 //AgtBgp4VpnIpv4RouteProfile SetAsPathIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) Set4ByteAsPathIncrementingRange () error {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeType First4ByteAsNumber AsNumberIncrement AsNumberRepeat
 //AgtBgp4VpnIpv4RouteProfile Set4ByteAsPathIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetAsPathIncrementingRange ()(string, error) {
 //parameters: RouteProfileHandle AsPathAttributeType
 //AgtBgp4VpnIpv4RouteProfile GetAsPathIncrementingRange
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) Get4ByteAsPathIncrementingRange ()(string, error) {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeType
 //AgtBgp4VpnIpv4RouteProfile Get4ByteAsPathIncrementingRange
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetAsPathList ()(string, error) {
 //parameters: RouteProfileHandle AsPathAttributeType AsNumberList
 //AgtBgp4VpnIpv4RouteProfile SetAsPathList
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) Set4ByteAsPathList ()(string, error) {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeType AsNumberList
 //AgtBgp4VpnIpv4RouteProfile Set4ByteAsPathList
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetAsPathList ()(string, error) {
 //parameters: RouteProfileHandle AsPathAttributeType
 //AgtBgp4VpnIpv4RouteProfile GetAsPathList
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) Get4ByteAsPathList ()(string, error) {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeType
 //AgtBgp4VpnIpv4RouteProfile Get4ByteAsPathList
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetAsPath () error {
 //parameters: RouteProfileHandle AsPathAttributeType AsNumber
 //AgtBgp4VpnIpv4RouteProfile SetAsPath, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) Set4ByteAsPath () error {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeType AsNumber
 //AgtBgp4VpnIpv4RouteProfile Set4ByteAsPath, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetAsPath ()(string, error) {
 //parameters: RouteProfileHandle AsPathAttributeType
 //AgtBgp4VpnIpv4RouteProfile GetAsPath
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) Get4ByteAsPath ()(string, error) {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeType
 //AgtBgp4VpnIpv4RouteProfile Get4ByteAsPath
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) AddAsPathToListOfLists ()(string, error) {
 //parameters: RouteProfileHandle AsPathAttributeType AsNumberList
 //AgtBgp4VpnIpv4RouteProfile AddAsPathToListOfLists
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) Add4ByteAsPathToListOfLists ()(string, error) {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeType AsNumberList
 //AgtBgp4VpnIpv4RouteProfile Add4ByteAsPathToListOfLists
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) RemoveAsPathFromListOfLists ()(string, error) {
 //parameters: RouteProfileHandle AsPathAttributeType AsPathListHandle
 //AgtBgp4VpnIpv4RouteProfile RemoveAsPathFromListOfLists
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) Remove4ByteAsPathFromListOfLists ()(string, error) {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeType AsPathListHandle
 //AgtBgp4VpnIpv4RouteProfile Remove4ByteAsPathFromListOfLists
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) ClearAsPathListOfLists ()(string, error) {
 //parameters: RouteProfileHandle AsPathAttributeType
 //AgtBgp4VpnIpv4RouteProfile ClearAsPathListOfLists
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) Clear4ByteAsPathListOfLists ()(string, error) {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeType
 //AgtBgp4VpnIpv4RouteProfile Clear4ByteAsPathListOfLists
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) ListAsPathListOfLists ()(string, error) {
 //parameters: RouteProfileHandle AsPathAttributeType AsPathListHandle
 //AgtBgp4VpnIpv4RouteProfile ListAsPathListOfLists
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) List4ByteAsPathListOfLists ()(string, error) {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeType AsPathListHandle
 //AgtBgp4VpnIpv4RouteProfile List4ByteAsPathListOfLists
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetAsPathAuto ()(string, error) {
 //parameters: RouteProfileHandle AsPathAttributeTypeAuto
 //AgtBgp4VpnIpv4RouteProfile GetAsPathAuto
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) Get4ByteAsPathAuto ()(string, error) {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeTypeAuto
 //AgtBgp4VpnIpv4RouteProfile Get4ByteAsPathAuto
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) Get4ByteAsPathAutoWithFormat ()(string, error) {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeTypeAuto FourByteFormatType
 //AgtBgp4VpnIpv4RouteProfile Get4ByteAsPathAutoWithFormat
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) EnablePathAttribute () error {
 //parameters: RouteProfileHandle PathAttribute
 //AgtBgp4VpnIpv4RouteProfile EnablePathAttribute, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) DisablePathAttribute () error {
 //parameters: RouteProfileHandle PathAttribute
 //AgtBgp4VpnIpv4RouteProfile DisablePathAttribute, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) IsPathAttributeEnabled () error {
 //parameters: RouteProfileHandle PathAttribute
 //AgtBgp4VpnIpv4RouteProfile IsPathAttributeEnabled, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetOrigin () error {
 //parameters: RouteProfileHandle Origin
 //AgtBgp4VpnIpv4RouteProfile SetOrigin, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetOrigin ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetOrigin
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetLocalPreference () error {
 //parameters: RouteProfileHandle LocalPreference
 //AgtBgp4VpnIpv4RouteProfile SetLocalPreference, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetLocalPreference ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetLocalPreference
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetMultiExitDiscriminator () error {
 //parameters: RouteProfileHandle MultiExitDiscriminator
 //AgtBgp4VpnIpv4RouteProfile SetMultiExitDiscriminator, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetMultiExitDiscriminator ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetMultiExitDiscriminator
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetAggregatorIncrementingRange () error {
 //parameters: RouteProfileHandle FirstAsNumber AsNumberIncrement AsNumberRepeat FirstIpAddress PrefixLength AddressIncrement AddressRepeat
 //AgtBgp4VpnIpv4RouteProfile SetAggregatorIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) Set4ByteAsAggregatorIncrementingRange () error {
 //parameters: RouteProfileHandle First4ByteAsNumber AsNumberIncrement AsNumberRepeat FirstIpAddress PrefixLength AddressIncrement AddressRepeat
 //AgtBgp4VpnIpv4RouteProfile Set4ByteAsAggregatorIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) Set4ByteAsAggregatorIncrementingRangeWithFormat () error {
 //parameters: RouteProfileHandle FourByteFormatType First4ByteAsNumber AsNumberIncrement AsNumberRepeat FirstIpAddress PrefixLength AddressIncrement AddressRepeat
 //AgtBgp4VpnIpv4RouteProfile Set4ByteAsAggregatorIncrementingRangeWithFormat, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetAggregatorIncrementingRange ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetAggregatorIncrementingRange
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) Get4ByteAsAggregatorIncrementingRange ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile Get4ByteAsAggregatorIncrementingRange
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) Get4ByteAsAggregatorIncrementingRangeWithFormat ()(string, error) {
 //parameters: RouteProfileHandle FourByteFormatType
 //AgtBgp4VpnIpv4RouteProfile Get4ByteAsAggregatorIncrementingRangeWithFormat
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetAggregatorList ()(string, error) {
 //parameters: RouteProfileHandle AsNumberList IpAddressList
 //AgtBgp4VpnIpv4RouteProfile SetAggregatorList
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) Set4ByteAsAggregatorList ()(string, error) {
 //parameters: RouteProfileHandle AsNumberList IpAddressList
 //AgtBgp4VpnIpv4RouteProfile Set4ByteAsAggregatorList
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) Set4ByteAsAggregatorListWithFormat ()(string, error) {
 //parameters: RouteProfileHandle FourByteFormatType AsNumberList IpAddressList
 //AgtBgp4VpnIpv4RouteProfile Set4ByteAsAggregatorListWithFormat
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetAggregatorList ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetAggregatorList
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) Get4ByteAsAggregatorList ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile Get4ByteAsAggregatorList
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) Get4ByteAsAggregatorListWithFormat ()(string, error) {
 //parameters: RouteProfileHandle FourByteFormatType
 //AgtBgp4VpnIpv4RouteProfile Get4ByteAsAggregatorListWithFormat
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetAggregator () error {
 //parameters: RouteProfileHandle AsNumber IpAddress
 //AgtBgp4VpnIpv4RouteProfile SetAggregator, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) Set4ByteAsAggregator () error {
 //parameters: RouteProfileHandle AsNumber IpAddress
 //AgtBgp4VpnIpv4RouteProfile Set4ByteAsAggregator, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) Set4ByteAsAggregatorWithFormat () error {
 //parameters: RouteProfileHandle FourByteFormatType AsNumber IpAddress
 //AgtBgp4VpnIpv4RouteProfile Set4ByteAsAggregatorWithFormat, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetAggregator ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetAggregator
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) Get4ByteAsAggregator ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile Get4ByteAsAggregator
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) Get4ByteAsAggregatorWithFormat ()(string, error) {
 //parameters: RouteProfileHandle FourByteFormatType
 //AgtBgp4VpnIpv4RouteProfile Get4ByteAsAggregatorWithFormat
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetAggregatorAuto ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetAggregatorAuto
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) Get4ByteAggregatorAuto ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile Get4ByteAggregatorAuto
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) Get4ByteAggregatorAutoWithFormat ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile Get4ByteAggregatorAutoWithFormat
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetOriginatorIdIncrementingRange () error {
 //parameters: RouteProfileHandle FirstOriginatorId PrefixLength AddressIncrement AddressRepeat
 //AgtBgp4VpnIpv4RouteProfile SetOriginatorIdIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetOriginatorIdIncrementingRange ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetOriginatorIdIncrementingRange
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetOriginatorIdList ()(string, error) {
 //parameters: RouteProfileHandle OriginatorIdList
 //AgtBgp4VpnIpv4RouteProfile SetOriginatorIdList
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetOriginatorIdList ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetOriginatorIdList
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetOriginatorId () error {
 //parameters: RouteProfileHandle OriginatorId
 //AgtBgp4VpnIpv4RouteProfile SetOriginatorId, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetOriginatorId ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetOriginatorId
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetOriginatorIdAuto ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetOriginatorIdAuto
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetClusterList ()(string, error) {
 //parameters: RouteProfileHandle ClusterList
 //AgtBgp4VpnIpv4RouteProfile SetClusterList
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetClusterList ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetClusterList
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetClusterListAuto ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetClusterListAuto
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetCommunities () error {
 //parameters: RouteProfileHandle CommunitiesList
 //AgtBgp4VpnIpv4RouteProfile SetCommunities, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetCommunities ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetCommunities
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetCommunitiesIncrementingRange () error {
 //parameters: RouteProfileHandle CommunitiesList CommunityIncrement CommunityRepeat
 //AgtBgp4VpnIpv4RouteProfile SetCommunitiesIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetCommunitiesIncrementingRange ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetCommunitiesIncrementingRange
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetCommunitiesList ()(string, error) {
 //parameters: RouteProfileHandle CommunitiesList
 //AgtBgp4VpnIpv4RouteProfile SetCommunitiesList
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetCommunitiesList ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetCommunitiesList
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) EnableTrafficDestinations () error {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile EnableTrafficDestinations, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) DisableTrafficDestinations () error {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile DisableTrafficDestinations, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) IsTrafficDestinationsEnabled () error {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile IsTrafficDestinationsEnabled, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) SetFlapDistribution () error {
 //parameters: RouteProfileHandle FlapPercentage FlapOffset
 //AgtBgp4VpnIpv4RouteProfile SetFlapDistribution, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) GetFlapDistribution ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile GetFlapDistribution
 return "", nil
}

func(np *Bgp4VpnIpv4RouteProfile) Advertise () error {
 //parameters: RouteProfileIdentifiers
 //AgtBgp4VpnIpv4RouteProfile Advertise, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) Withdraw () error {
 //parameters: RouteProfileIdentifiers
 //AgtBgp4VpnIpv4RouteProfile Withdraw, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) Enable () error {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile Enable, m.Object, m.Name);
 return nil
}

func(np *Bgp4VpnIpv4RouteProfile) Disable () error {
 //parameters: RouteProfileHandle
 //AgtBgp4VpnIpv4RouteProfile Disable, m.Object, m.Name);
 return nil
}

