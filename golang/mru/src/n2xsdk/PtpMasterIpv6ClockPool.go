package n2xsdk

type PtpMasterIpv6ClockPool struct {
 Handler string
}

func(np *PtpMasterIpv6ClockPool) SetClockQuality () error {
 //parameters: SessionPoolHandle ClockClass ClockAccuracy OffsetScaledLogVariance Priority1 Priority2 AlternateMasterFlag StepsRemoved TimeSource
 //AgtPtpMasterIpv6ClockPool SetClockQuality, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) GetClockQuality ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool GetClockQuality
 return "", nil
}

func(np *PtpMasterIpv6ClockPool) EnableAnnounce () error {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool EnableAnnounce, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) DisableAnnounce () error {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool DisableAnnounce, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) IsEnabledAnnounce () error {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool IsEnabledAnnounce, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) SetAnnounceAttributes () error {
 //parameters: SessionPoolHandle LogMessageInterval ReceiptTimeout
 //AgtPtpMasterIpv6ClockPool SetAnnounceAttributes, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) GetAnnounceAttributes ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool GetAnnounceAttributes
 return "", nil
}

func(np *PtpMasterIpv6ClockPool) EnableSync () error {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool EnableSync, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) DisableSync () error {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool DisableSync, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) IsEnabledSync () error {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool IsEnabledSync, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) SetSyncAttributes () error {
 //parameters: SessionPoolHandle TwoStepFlagMode LogMessageInterval
 //AgtPtpMasterIpv6ClockPool SetSyncAttributes, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) GetSyncAttributes ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool GetSyncAttributes
 return "", nil
}

func(np *PtpMasterIpv6ClockPool) EnableFollowUp () error {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool EnableFollowUp, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) DisableFollowUp () error {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool DisableFollowUp, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) IsEnabledFollowUp () error {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool IsEnabledFollowUp, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) SetFollowUpAttributes () error {
 //parameters: SessionPoolHandle FollowUpDelay
 //AgtPtpMasterIpv6ClockPool SetFollowUpAttributes, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) GetFollowUpAttributes ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool GetFollowUpAttributes
 return "", nil
}

func(np *PtpMasterIpv6ClockPool) EnableDelayResponse () error {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool EnableDelayResponse, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) DisableDelayResponse () error {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool DisableDelayResponse, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) IsEnabledDelayResponse () error {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool IsEnabledDelayResponse, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) SetDelayResponseAttributes () error {
 //parameters: SessionPoolHandle LogMessageInterval
 //AgtPtpMasterIpv6ClockPool SetDelayResponseAttributes, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) GetDelayResponseAttributes ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool GetDelayResponseAttributes
 return "", nil
}

func(np *PtpMasterIpv6ClockPool) SetDomainNumber () error {
 //parameters: SessionPoolHandle DomainNumber
 //AgtPtpMasterIpv6ClockPool SetDomainNumber, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) GetDomainNumber ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool GetDomainNumber
 return "", nil
}

func(np *PtpMasterIpv6ClockPool) SetPortNumber () error {
 //parameters: SessionPoolHandle PortNumber
 //AgtPtpMasterIpv6ClockPool SetPortNumber, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) GetPortNumber ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool GetPortNumber
 return "", nil
}

func(np *PtpMasterIpv6ClockPool) SetClockId () error {
 //parameters: SessionPoolHandle ClockId
 //AgtPtpMasterIpv6ClockPool SetClockId, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) GetClockId ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool GetClockId
 return "", nil
}

func(np *PtpMasterIpv6ClockPool) SetClockIdIncrementingRange () error {
 //parameters: SessionPoolHandle FirstClockId Repeat Increment
 //AgtPtpMasterIpv6ClockPool SetClockIdIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) GetClockIdIncrementingRange ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool GetClockIdIncrementingRange
 return "", nil
}

func(np *PtpMasterIpv6ClockPool) SetClockIdList ()(string, error) {
 //parameters: SessionPoolHandle ClockIdList
 //AgtPtpMasterIpv6ClockPool SetClockIdList
 return "", nil
}

func(np *PtpMasterIpv6ClockPool) GetClockIdList ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool GetClockIdList
 return "", nil
}

func(np *PtpMasterIpv6ClockPool) SetPathDelayMode () error {
 //parameters: SessionPoolHandle PathDelayMode
 //AgtPtpMasterIpv6ClockPool SetPathDelayMode, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) GetPathDelayMode ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool GetPathDelayMode
 return "", nil
}

func(np *PtpMasterIpv6ClockPool) EnablePeerDelayRequest () error {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool EnablePeerDelayRequest, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) DisablePeerDelayRequest () error {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool DisablePeerDelayRequest, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) IsEnabledPeerDelayRequest () error {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool IsEnabledPeerDelayRequest, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) EnablePeerDelayResponse () error {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool EnablePeerDelayResponse, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) DisablePeerDelayResponse () error {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool DisablePeerDelayResponse, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) IsEnabledPeerDelayResponse () error {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool IsEnabledPeerDelayResponse, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) EnablePeerDelayResponseFollowUp () error {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool EnablePeerDelayResponseFollowUp, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) DisablePeerDelayResponseFollowUp () error {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool DisablePeerDelayResponseFollowUp, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) IsEnabledPeerDelayResponseFollowUp () error {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool IsEnabledPeerDelayResponseFollowUp, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) SetPeerToPeerPathDelayAttributes () error {
 //parameters: SessionPoolHandle LogMessageInterval TwoStepFlagMode FollowUpDelay
 //AgtPtpMasterIpv6ClockPool SetPeerToPeerPathDelayAttributes, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) GetPeerToPeerPathDelayAttributes ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool GetPeerToPeerPathDelayAttributes
 return "", nil
}

func(np *PtpMasterIpv6ClockPool) SetTimestampCalibration () error {
 //parameters: SessionPoolHandle TxOffset RxOffset
 //AgtPtpMasterIpv6ClockPool SetTimestampCalibration, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) GetTimestampCalibration ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool GetTimestampCalibration
 return "", nil
}

func(np *PtpMasterIpv6ClockPool) GetVlanPriority ()(string, error) {
 //parameters: DeviceHandle VlanTagIndex
 //AgtPtpMasterIpv6ClockPool GetVlanPriority
 return "", nil
}

func(np *PtpMasterIpv6ClockPool) SetVlanPriority () error {
 //parameters: DeviceHandle VlanTagIndex VlanPriority
 //AgtPtpMasterIpv6ClockPool SetVlanPriority, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) SetDestinationAddressType () error {
 //parameters: SessionPoolHandle AddressType
 //AgtPtpMasterIpv6ClockPool SetDestinationAddressType, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) GetDestinationAddressType ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool GetDestinationAddressType
 return "", nil
}

func(np *PtpMasterIpv6ClockPool) IsIpv6PriorityNoCodePointFieldSelected () error {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool IsIpv6PriorityNoCodePointFieldSelected, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) SelectIpv6PriorityNoCodePointField () error {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool SelectIpv6PriorityNoCodePointField, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) GetIpv6Priority ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool GetIpv6Priority
 return "", nil
}

func(np *PtpMasterIpv6ClockPool) SetIpv6Priority () error {
 //parameters: SessionPoolHandle Ipv6Priority
 //AgtPtpMasterIpv6ClockPool SetIpv6Priority, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) IsIpv6PriorityDiffServFieldSelected () error {
 //parameters: SessionPoolHandle DiffServPerHopBehaviorGroup
 //AgtPtpMasterIpv6ClockPool IsIpv6PriorityDiffServFieldSelected, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) SelectIpv6PriorityDiffServField () error {
 //parameters: SessionPoolHandle DiffServPerHopBehaviorGroup
 //AgtPtpMasterIpv6ClockPool SelectIpv6PriorityDiffServField, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) GetIpv6PriorityDiffServValue ()(string, error) {
 //parameters: SessionPoolHandle DiffServCodePointField
 //AgtPtpMasterIpv6ClockPool GetIpv6PriorityDiffServValue
 return "", nil
}

func(np *PtpMasterIpv6ClockPool) SetIpv6PriorityDiffServValue () error {
 //parameters: SessionPoolHandle DiffServCodePointConfigurableField Value
 //AgtPtpMasterIpv6ClockPool SetIpv6PriorityDiffServValue, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) SetDestinationAddress () error {
 //parameters: SessionPoolHandle FirstAddress
 //AgtPtpMasterIpv6ClockPool SetDestinationAddress, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) GetDestinationAddress ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool GetDestinationAddress
 return "", nil
}

func(np *PtpMasterIpv6ClockPool) SetDestinationAddressIncrementingRange () error {
 //parameters: SessionPoolHandle FirstAddress PrefixLength Increment Repeat PercentageOverlap
 //AgtPtpMasterIpv6ClockPool SetDestinationAddressIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) GetDestinationAddressIncrementingRange ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool GetDestinationAddressIncrementingRange
 return "", nil
}

func(np *PtpMasterIpv6ClockPool) SetDestinationAddressList ()(string, error) {
 //parameters: SessionPoolHandle IpAddressList
 //AgtPtpMasterIpv6ClockPool SetDestinationAddressList
 return "", nil
}

func(np *PtpMasterIpv6ClockPool) GetDestinationAddressList ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool GetDestinationAddressList
 return "", nil
}

func(np *PtpMasterIpv6ClockPool) SetPeerDestinationAddress () error {
 //parameters: SessionPoolHandle PeerAddress
 //AgtPtpMasterIpv6ClockPool SetPeerDestinationAddress, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) GetPeerDestinationAddress ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool GetPeerDestinationAddress
 return "", nil
}

func(np *PtpMasterIpv6ClockPool) EnableUnicastNegotiation () error {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool EnableUnicastNegotiation, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) DisableUnicastNegotiation () error {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool DisableUnicastNegotiation, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) IsEnabledUnicastNegotiation () error {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool IsEnabledUnicastNegotiation, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) SetUnicastGrantDuration () error {
 //parameters: SessionPoolHandle UnicastGrantDuration
 //AgtPtpMasterIpv6ClockPool SetUnicastGrantDuration, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) GetUnicastGrantDuration ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool GetUnicastGrantDuration
 return "", nil
}

func(np *PtpMasterIpv6ClockPool) SetNumberOfUnicastDestinations () error {
 //parameters: SessionPoolHandle UnicastDestinationCount
 //AgtPtpMasterIpv6ClockPool SetNumberOfUnicastDestinations, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) GetNumberOfUnicastDestinations ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool GetNumberOfUnicastDestinations
 return "", nil
}

func(np *PtpMasterIpv6ClockPool) SetPeerPortNumber () error {
 //parameters: SessionPoolHandle PortNumber
 //AgtPtpMasterIpv6ClockPool SetPeerPortNumber, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) GetPeerPortNumber ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool GetPeerPortNumber
 return "", nil
}

func(np *PtpMasterIpv6ClockPool) SetPeerClockId () error {
 //parameters: SessionPoolHandle ClockId
 //AgtPtpMasterIpv6ClockPool SetPeerClockId, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) GetPeerClockId ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool GetPeerClockId
 return "", nil
}

func(np *PtpMasterIpv6ClockPool) SetUdpSourcePort () error {
 //parameters: SessionPoolHandle FirstPort
 //AgtPtpMasterIpv6ClockPool SetUdpSourcePort, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) GetUdpSourcePort ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool GetUdpSourcePort
 return "", nil
}

func(np *PtpMasterIpv6ClockPool) SetUdpSourcePortIncrementingRange () error {
 //parameters: SessionPoolHandle FirstPort Increment Repeat
 //AgtPtpMasterIpv6ClockPool SetUdpSourcePortIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *PtpMasterIpv6ClockPool) GetUdpSourcePortIncrementingRange ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool GetUdpSourcePortIncrementingRange
 return "", nil
}

func(np *PtpMasterIpv6ClockPool) SetUdpSourcePortList ()(string, error) {
 //parameters: SessionPoolHandle UdpPortList
 //AgtPtpMasterIpv6ClockPool SetUdpSourcePortList
 return "", nil
}

func(np *PtpMasterIpv6ClockPool) GetUdpSourcePortList ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPtpMasterIpv6ClockPool GetUdpSourcePortList
 return "", nil
}

