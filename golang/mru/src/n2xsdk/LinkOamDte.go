package n2xsdk

type LinkOamDte struct {
 Handler string
}

func(np *LinkOamDte) EnableAutomaticDiscovery () error {
 //parameters: SessionPoolHandle
 //AgtLinkOamDte EnableAutomaticDiscovery, m.Object, m.Name);
 return nil
}

func(np *LinkOamDte) DisableAutomaticDiscovery () error {
 //parameters: SessionPoolHandle
 //AgtLinkOamDte DisableAutomaticDiscovery, m.Object, m.Name);
 return nil
}

func(np *LinkOamDte) IsAutomaticDiscoveryEnabled () error {
 //parameters: SessionPoolHandle
 //AgtLinkOamDte IsAutomaticDiscoveryEnabled, m.Object, m.Name);
 return nil
}

func(np *LinkOamDte) EnableOrganizationSpecificInformation () error {
 //parameters: SessionPoolHandle
 //AgtLinkOamDte EnableOrganizationSpecificInformation, m.Object, m.Name);
 return nil
}

func(np *LinkOamDte) DisableOrganizationSpecificInformation () error {
 //parameters: SessionPoolHandle
 //AgtLinkOamDte DisableOrganizationSpecificInformation, m.Object, m.Name);
 return nil
}

func(np *LinkOamDte) IsOrganizationSpecificInformationEnabled () error {
 //parameters: SessionPoolHandle
 //AgtLinkOamDte IsOrganizationSpecificInformationEnabled, m.Object, m.Name);
 return nil
}

func(np *LinkOamDte) EnableCriticalLinkEvents () error {
 //parameters: SessionPoolHandle CriticalLinkEventType
 //AgtLinkOamDte EnableCriticalLinkEvents, m.Object, m.Name);
 return nil
}

func(np *LinkOamDte) DisableCriticalLinkEvents () error {
 //parameters: SessionPoolHandle CriticalLinkEventType
 //AgtLinkOamDte DisableCriticalLinkEvents, m.Object, m.Name);
 return nil
}

func(np *LinkOamDte) IsCriticalLinkEventsEnabled () error {
 //parameters: SessionPoolHandle CriticalLinkEventType
 //AgtLinkOamDte IsCriticalLinkEventsEnabled, m.Object, m.Name);
 return nil
}

func(np *LinkOamDte) EnableLinkEventNotification () error {
 //parameters: SessionPoolHandle LinkNotificationType
 //AgtLinkOamDte EnableLinkEventNotification, m.Object, m.Name);
 return nil
}

func(np *LinkOamDte) DisableLinkEventNotification () error {
 //parameters: SessionPoolHandle LinkNotificationType
 //AgtLinkOamDte DisableLinkEventNotification, m.Object, m.Name);
 return nil
}

func(np *LinkOamDte) IsLinkEventNotificationEnabled () error {
 //parameters: SessionPoolHandle LinkNotificationType
 //AgtLinkOamDte IsLinkEventNotificationEnabled, m.Object, m.Name);
 return nil
}

func(np *LinkOamDte) EnableOrganizationSpecificEvent () error {
 //parameters: SessionPoolHandle
 //AgtLinkOamDte EnableOrganizationSpecificEvent, m.Object, m.Name);
 return nil
}

func(np *LinkOamDte) DisableOrganizationSpecificEvent () error {
 //parameters: SessionPoolHandle
 //AgtLinkOamDte DisableOrganizationSpecificEvent, m.Object, m.Name);
 return nil
}

func(np *LinkOamDte) IsOrganizationSpecificEventEnabled () error {
 //parameters: SessionPoolHandle
 //AgtLinkOamDte IsOrganizationSpecificEventEnabled, m.Object, m.Name);
 return nil
}

func(np *LinkOamDte) EnableDiscoveryCapability () error {
 //parameters: SessionPoolHandle DiscoveryCapabilityType
 //AgtLinkOamDte EnableDiscoveryCapability, m.Object, m.Name);
 return nil
}

func(np *LinkOamDte) DisableDiscoveryCapability () error {
 //parameters: SessionPoolHandle DiscoveryCapabilityType
 //AgtLinkOamDte DisableDiscoveryCapability, m.Object, m.Name);
 return nil
}

func(np *LinkOamDte) IsDiscoveryCapabilityEnabled () error {
 //parameters: SessionPoolHandle DiscoveryCapabilityType
 //AgtLinkOamDte IsDiscoveryCapabilityEnabled, m.Object, m.Name);
 return nil
}

func(np *LinkOamDte) SetOamMode () error {
 //parameters: SessionPoolHandle OamMode
 //AgtLinkOamDte SetOamMode, m.Object, m.Name);
 return nil
}

func(np *LinkOamDte) GetOamMode ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLinkOamDte GetOamMode
 return "", nil
}

func(np *LinkOamDte) SetDestinationMacType () error {
 //parameters: SessionPoolHandle DestinationMacType
 //AgtLinkOamDte SetDestinationMacType, m.Object, m.Name);
 return nil
}

func(np *LinkOamDte) GetDestinationMacType ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLinkOamDte GetDestinationMacType
 return "", nil
}

func(np *LinkOamDte) SetUnicastMac () error {
 //parameters: SessionPoolHandle UnicastMac
 //AgtLinkOamDte SetUnicastMac, m.Object, m.Name);
 return nil
}

func(np *LinkOamDte) GetUnicastMac ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLinkOamDte GetUnicastMac
 return "", nil
}

func(np *LinkOamDte) GetOui ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLinkOamDte GetOui
 return "", nil
}

func(np *LinkOamDte) SetOui () error {
 //parameters: SessionPoolHandle Oui
 //AgtLinkOamDte SetOui, m.Object, m.Name);
 return nil
}

func(np *LinkOamDte) GetVendorSpecificInformation ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLinkOamDte GetVendorSpecificInformation
 return "", nil
}

func(np *LinkOamDte) SetVendorSpecificInformation () error {
 //parameters: SessionPoolHandle VendorSpecificInformation
 //AgtLinkOamDte SetVendorSpecificInformation, m.Object, m.Name);
 return nil
}

func(np *LinkOamDte) GetMaxPduLength ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLinkOamDte GetMaxPduLength
 return "", nil
}

func(np *LinkOamDte) SetMaxPduLength () error {
 //parameters: SessionPoolHandle MaxOamPduLength
 //AgtLinkOamDte SetMaxPduLength, m.Object, m.Name);
 return nil
}

func(np *LinkOamDte) GetNumberOfPdus ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLinkOamDte GetNumberOfPdus
 return "", nil
}

func(np *LinkOamDte) SetNumberOfPdus () error {
 //parameters: SessionPoolHandle PduCount
 //AgtLinkOamDte SetNumberOfPdus, m.Object, m.Name);
 return nil
}

func(np *LinkOamDte) SetLinkEventNotificationFlowControl () error {
 //parameters: SessionPoolHandle EventsPerInterval Interval NumberOfIntervals
 //AgtLinkOamDte SetLinkEventNotificationFlowControl, m.Object, m.Name);
 return nil
}

func(np *LinkOamDte) GetLinkEventNotificationFlowControl ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLinkOamDte GetLinkEventNotificationFlowControl
 return "", nil
}

func(np *LinkOamDte) GetOrganizationSpecificEventValue ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLinkOamDte GetOrganizationSpecificEventValue
 return "", nil
}

func(np *LinkOamDte) SetOrganizationSpecificEventValue () error {
 //parameters: SessionPoolHandle OrganizationSpecificEventValue
 //AgtLinkOamDte SetOrganizationSpecificEventValue, m.Object, m.Name);
 return nil
}

func(np *LinkOamDte) GetOrganizationSpecificInformationValue ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLinkOamDte GetOrganizationSpecificInformationValue
 return "", nil
}

func(np *LinkOamDte) SetOrganizationSpecificInformationValue () error {
 //parameters: SessionPoolHandle OrganizationSpecificInformationValue
 //AgtLinkOamDte SetOrganizationSpecificInformationValue, m.Object, m.Name);
 return nil
}

func(np *LinkOamDte) StartDiscovery () error {
 //parameters: SessionPoolHandle
 //AgtLinkOamDte StartDiscovery, m.Object, m.Name);
 return nil
}

func(np *LinkOamDte) SendEnableRemoteLoopback () error {
 //parameters: SessionPoolHandle
 //AgtLinkOamDte SendEnableRemoteLoopback, m.Object, m.Name);
 return nil
}

func(np *LinkOamDte) SendDisableRemoteLoopback () error {
 //parameters: SessionPoolHandle
 //AgtLinkOamDte SendDisableRemoteLoopback, m.Object, m.Name);
 return nil
}

func(np *LinkOamDte) StartLinkEventNotification () error {
 //parameters: SessionPoolHandle
 //AgtLinkOamDte StartLinkEventNotification, m.Object, m.Name);
 return nil
}

func(np *LinkOamDte) StopLinkEventNotification () error {
 //parameters: SessionPoolHandle
 //AgtLinkOamDte StopLinkEventNotification, m.Object, m.Name);
 return nil
}

