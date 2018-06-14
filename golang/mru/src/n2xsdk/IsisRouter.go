package n2xsdk

type IsisRouter struct {
 Handler string
}

func(np *IsisRouter) SetAreaId () error {
 //parameters: RouterHandle AreaId
 //AgtIsisRouter SetAreaId, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) GetAreaId ()(string, error) {
 //parameters: RouterHandle
 //AgtIsisRouter GetAreaId
 return "", nil
}

func(np *IsisRouter) SetHostname () error {
 //parameters: RouterHandle Hostname
 //AgtIsisRouter SetHostname, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) GetHostname ()(string, error) {
 //parameters: RouterHandle
 //AgtIsisRouter GetHostname
 return "", nil
}

func(np *IsisRouter) SetAreaList ()(string, error) {
 //parameters: RouterHandle Count AreaList
 //AgtIsisRouter SetAreaList
 return "", nil
}

func(np *IsisRouter) GetAreaList ()(string, error) {
 //parameters: RouterHandle
 //AgtIsisRouter GetAreaList
 return "", nil
}

func(np *IsisRouter) SetSystemId () error {
 //parameters: RouterHandle SystemId
 //AgtIsisRouter SetSystemId, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) GetSystemId ()(string, error) {
 //parameters: RouterHandle
 //AgtIsisRouter GetSystemId
 return "", nil
}

func(np *IsisRouter) SetPseudonodeNumber () error {
 //parameters: RouterHandle PseudonodeNumber
 //AgtIsisRouter SetPseudonodeNumber, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) GetPseudonodeNumber ()(string, error) {
 //parameters: RouterHandle
 //AgtIsisRouter GetPseudonodeNumber
 return "", nil
}

func(np *IsisRouter) SetTeRouterId () error {
 //parameters: RouterHandle TeRouterId
 //AgtIsisRouter SetTeRouterId, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) GetTeRouterId ()(string, error) {
 //parameters: RouterHandle
 //AgtIsisRouter GetTeRouterId
 return "", nil
}

func(np *IsisRouter) EnableTe () error {
 //parameters: RouterHandle
 //AgtIsisRouter EnableTe, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) DisableTe () error {
 //parameters: RouterHandle
 //AgtIsisRouter DisableTe, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) IsTeEnabled () error {
 //parameters: RouterHandle
 //AgtIsisRouter IsTeEnabled, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) SetRoutingLevel () error {
 //parameters: RouterHandle RoutingLevel
 //AgtIsisRouter SetRoutingLevel, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) GetRoutingLevel ()(string, error) {
 //parameters: RouterHandle
 //AgtIsisRouter GetRoutingLevel
 return "", nil
}

func(np *IsisRouter) SetOverloadBit () error {
 //parameters: RouterHandle OverloadBit
 //AgtIsisRouter SetOverloadBit, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) GetOverloadBit ()(string, error) {
 //parameters: RouterHandle
 //AgtIsisRouter GetOverloadBit
 return "", nil
}

func(np *IsisRouter) SetAttachedBit () error {
 //parameters: RouterHandle AttachedBit
 //AgtIsisRouter SetAttachedBit, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) GetAttachedBit ()(string, error) {
 //parameters: RouterHandle
 //AgtIsisRouter GetAttachedBit
 return "", nil
}

func(np *IsisRouter) Advertise () error {
 //parameters: RouterHandle
 //AgtIsisRouter Advertise, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) Withdraw () error {
 //parameters: RouterHandle
 //AgtIsisRouter Withdraw, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) GetAdvertiseFlag ()(string, error) {
 //parameters: RouterHandle
 //AgtIsisRouter GetAdvertiseFlag
 return "", nil
}

func(np *IsisRouter) GetRouterLsp ()(string, error) {
 //parameters: RouterHandle
 //AgtIsisRouter GetRouterLsp
 return "", nil
}

func(np *IsisRouter) ConnectObjects () error {
 //parameters: RouterHandle OtherObjectHandle
 //AgtIsisRouter ConnectObjects, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) DisconnectObjects () error {
 //parameters: RouterHandle OtherObjectHandle
 //AgtIsisRouter DisconnectObjects, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) ListObjectConnections ()(string, error) {
 //parameters: RouterHandle
 //AgtIsisRouter ListObjectConnections
 return "", nil
}

func(np *IsisRouter) GetLinkNarrowMetric ()(string, error) {
 //parameters: RouterHandle OtherObjectHandle
 //AgtIsisRouter GetLinkNarrowMetric
 return "", nil
}

func(np *IsisRouter) SetLinkNarrowMetric () error {
 //parameters: RouterHandle OtherObjectHandle NarrowMetric
 //AgtIsisRouter SetLinkNarrowMetric, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) GetLinkWideMetric ()(string, error) {
 //parameters: RouterHandle OtherObjectHandle
 //AgtIsisRouter GetLinkWideMetric
 return "", nil
}

func(np *IsisRouter) SetLinkWideMetric () error {
 //parameters: RouterHandle OtherObjectHandle WideMetric
 //AgtIsisRouter SetLinkWideMetric, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) IsRouterLink () error {
 //parameters: RouterHandle OtherObjectHandle
 //AgtIsisRouter IsRouterLink, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) SetLinkIpAddress () error {
 //parameters: RouterHandle OtherObjectHandle IpAddress
 //AgtIsisRouter SetLinkIpAddress, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) GetLinkIpAddress ()(string, error) {
 //parameters: RouterHandle OtherObjectHandle
 //AgtIsisRouter GetLinkIpAddress
 return "", nil
}

func(np *IsisRouter) SetLinkIpv4Address () error {
 //parameters: RouterHandle OtherObjectHandle Ipv4Address
 //AgtIsisRouter SetLinkIpv4Address, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) GetLinkIpv4Address ()(string, error) {
 //parameters: RouterHandle OtherObjectHandle
 //AgtIsisRouter GetLinkIpv4Address
 return "", nil
}

func(np *IsisRouter) SetLinkIpv6Address () error {
 //parameters: RouterHandle OtherObjectHandle Ipv6Address
 //AgtIsisRouter SetLinkIpv6Address, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) GetLinkIpv6Address ()(string, error) {
 //parameters: RouterHandle OtherObjectHandle
 //AgtIsisRouter GetLinkIpv6Address
 return "", nil
}

func(np *IsisRouter) SetLinkPrefixLength () error {
 //parameters: RouterHandle OtherObjectHandle PrefixLength
 //AgtIsisRouter SetLinkPrefixLength, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) GetLinkPrefixLength ()(string, error) {
 //parameters: RouterHandle OtherObjectHandle
 //AgtIsisRouter GetLinkPrefixLength
 return "", nil
}

func(np *IsisRouter) SetLinkIpv4PrefixLength () error {
 //parameters: RouterHandle OtherObjectHandle PrefixLength
 //AgtIsisRouter SetLinkIpv4PrefixLength, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) GetLinkIpv4PrefixLength ()(string, error) {
 //parameters: RouterHandle OtherObjectHandle
 //AgtIsisRouter GetLinkIpv4PrefixLength
 return "", nil
}

func(np *IsisRouter) SetLinkIpv6PrefixLength () error {
 //parameters: RouterHandle OtherObjectHandle PrefixLength
 //AgtIsisRouter SetLinkIpv6PrefixLength, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) GetLinkIpv6PrefixLength ()(string, error) {
 //parameters: RouterHandle OtherObjectHandle
 //AgtIsisRouter GetLinkIpv6PrefixLength
 return "", nil
}

func(np *IsisRouter) SetRemoteLinkIpAddress () error {
 //parameters: RouterHandle OtherRouterHandle IpAddress
 //AgtIsisRouter SetRemoteLinkIpAddress, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) GetRemoteLinkIpAddress ()(string, error) {
 //parameters: RouterHandle OtherRouterHandle
 //AgtIsisRouter GetRemoteLinkIpAddress
 return "", nil
}

func(np *IsisRouter) SetRemoteLinkIpv4Address () error {
 //parameters: RouterHandle OtherRouterHandle Ipv4Address
 //AgtIsisRouter SetRemoteLinkIpv4Address, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) GetRemoteLinkIpv4Address ()(string, error) {
 //parameters: RouterHandle OtherRouterHandle
 //AgtIsisRouter GetRemoteLinkIpv4Address
 return "", nil
}

func(np *IsisRouter) SetRemoteLinkIpv6Address () error {
 //parameters: RouterHandle OtherRouterHandle Ipv6Address
 //AgtIsisRouter SetRemoteLinkIpv6Address, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) GetRemoteLinkIpv6Address ()(string, error) {
 //parameters: RouterHandle OtherRouterHandle
 //AgtIsisRouter GetRemoteLinkIpv6Address
 return "", nil
}

func(np *IsisRouter) SetMaxLinkBandwidth () error {
 //parameters: RouterHandle OtherRouterHandle MaxLinkBw
 //AgtIsisRouter SetMaxLinkBandwidth, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) GetMaxLinkBandwidth ()(string, error) {
 //parameters: RouterHandle OtherRouterHandle
 //AgtIsisRouter GetMaxLinkBandwidth
 return "", nil
}

func(np *IsisRouter) SetMaxReservableBandwidth () error {
 //parameters: RouterHandle OtherRouterHandle ReservableBw
 //AgtIsisRouter SetMaxReservableBandwidth, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) GetMaxReservableBandwidth ()(string, error) {
 //parameters: RouterHandle OtherRouterHandle
 //AgtIsisRouter GetMaxReservableBandwidth
 return "", nil
}

func(np *IsisRouter) SetUnreservedBandwidth () error {
 //parameters: RouterHandle OtherRouterHandle Count pUnreservedBw
 //AgtIsisRouter SetUnreservedBandwidth, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) GetUnreservedBandwidth ()(string, error) {
 //parameters: RouterHandle OtherRouterHandle
 //AgtIsisRouter GetUnreservedBandwidth
 return "", nil
}

func(np *IsisRouter) SetAdministrativeGroups () error {
 //parameters: RouterHandle OtherRouterHandle AdministrativeGroups
 //AgtIsisRouter SetAdministrativeGroups, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) GetAdministrativeGroups ()(string, error) {
 //parameters: RouterHandle OtherRouterHandle
 //AgtIsisRouter GetAdministrativeGroups
 return "", nil
}

func(np *IsisRouter) EnableLink () error {
 //parameters: RouterHandle OtherObjectHandle
 //AgtIsisRouter EnableLink, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) DisableLink () error {
 //parameters: RouterHandle OtherObjectHandle
 //AgtIsisRouter DisableLink, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) IsLinkEnabled () error {
 //parameters: RouterHandle OtherObjectHandle
 //AgtIsisRouter IsLinkEnabled, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) GetExternalFlag ()(string, error) {
 //parameters: RouterHandle OtherObjectHandle
 //AgtIsisRouter GetExternalFlag
 return "", nil
}

func(np *IsisRouter) SetExternalFlag () error {
 //parameters: RouterHandle OtherObjectHandle IsExternal
 //AgtIsisRouter SetExternalFlag, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) GetUpDownBit ()(string, error) {
 //parameters: RouterHandle OtherObjectHandle
 //AgtIsisRouter GetUpDownBit
 return "", nil
}

func(np *IsisRouter) SetUpDownBit () error {
 //parameters: RouterHandle OtherObjectHandle UpDownBit
 //AgtIsisRouter SetUpDownBit, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) EnableLinkTe () error {
 //parameters: RouterHandle OtherRouterHandle
 //AgtIsisRouter EnableLinkTe, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) DisableLinkTe () error {
 //parameters: RouterHandle OtherRouterHandle
 //AgtIsisRouter DisableLinkTe, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) IsLinkTeEnabled () error {
 //parameters: RouterHandle OtherRouterHandle
 //AgtIsisRouter IsLinkTeEnabled, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) SetIpMode () error {
 //parameters: RouterHandle IpMode
 //AgtIsisRouter SetIpMode, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) GetIpMode ()(string, error) {
 //parameters: RouterHandle
 //AgtIsisRouter GetIpMode
 return "", nil
}

func(np *IsisRouter) IsSessionRouter () error {
 //parameters: RouterHandle
 //AgtIsisRouter IsSessionRouter, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) IsNeighborRouter () error {
 //parameters: RouterHandle
 //AgtIsisRouter IsNeighborRouter, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) IsIpv6AddressAvailable () error {
 //parameters: RouterHandle
 //AgtIsisRouter IsIpv6AddressAvailable, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) EnableMultiTopologies () error {
 //parameters: RouterHandle
 //AgtIsisRouter EnableMultiTopologies, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) DisableMultiTopologies () error {
 //parameters: RouterHandle
 //AgtIsisRouter DisableMultiTopologies, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) IsMultiTopologiesEnabled () error {
 //parameters: RouterHandle
 //AgtIsisRouter IsMultiTopologiesEnabled, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) EnableForwardingAdjacency () error {
 //parameters: RouterHandle
 //AgtIsisRouter EnableForwardingAdjacency, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) DisableForwardingAdjacency () error {
 //parameters: RouterHandle
 //AgtIsisRouter DisableForwardingAdjacency, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) IsForwardingAdjacencyEnabled () error {
 //parameters: RouterHandle
 //AgtIsisRouter IsForwardingAdjacencyEnabled, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) SetFAPeerSystemID () error {
 //parameters: RouterHandle FAPeerSystemId
 //AgtIsisRouter SetFAPeerSystemID, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) GetFAPeerSystemID ()(string, error) {
 //parameters: RouterHandle
 //AgtIsisRouter GetFAPeerSystemID
 return "", nil
}

func(np *IsisRouter) SetFAPeerMetric () error {
 //parameters: RouterHandle FAPeerMetric
 //AgtIsisRouter SetFAPeerMetric, m.Object, m.Name);
 return nil
}

func(np *IsisRouter) GetFAPeerMetric ()(string, error) {
 //parameters: RouterHandle
 //AgtIsisRouter GetFAPeerMetric
 return "", nil
}

