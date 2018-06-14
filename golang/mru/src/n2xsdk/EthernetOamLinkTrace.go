package n2xsdk

type EthernetOamLinkTrace struct {
 Handler string
}

func(np *EthernetOamLinkTrace) SetStandard () error {
 //parameters: SessionPoolHandle EthOamStandard
 //AgtEthernetOamLinkTrace SetStandard, m.Object, m.Name);
 return nil
}

func(np *EthernetOamLinkTrace) GetStandard ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtEthernetOamLinkTrace GetStandard
 return "", nil
}

func(np *EthernetOamLinkTrace) SetTargetMacAddress () error {
 //parameters: SessionPoolHandle MacAddress
 //AgtEthernetOamLinkTrace SetTargetMacAddress, m.Object, m.Name);
 return nil
}

func(np *EthernetOamLinkTrace) GetTargetMacAddress ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtEthernetOamLinkTrace GetTargetMacAddress
 return "", nil
}

func(np *EthernetOamLinkTrace) SetStartTtl () error {
 //parameters: SessionPoolHandle StartTtl
 //AgtEthernetOamLinkTrace SetStartTtl, m.Object, m.Name);
 return nil
}

func(np *EthernetOamLinkTrace) GetStartTtl ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtEthernetOamLinkTrace GetStartTtl
 return "", nil
}

func(np *EthernetOamLinkTrace) SetMdLevel () error {
 //parameters: SessionPoolHandle Level
 //AgtEthernetOamLinkTrace SetMdLevel, m.Object, m.Name);
 return nil
}

func(np *EthernetOamLinkTrace) GetMdLevel ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtEthernetOamLinkTrace GetMdLevel
 return "", nil
}

func(np *EthernetOamLinkTrace) SetMegLevel () error {
 //parameters: SessionPoolHandle Level
 //AgtEthernetOamLinkTrace SetMegLevel, m.Object, m.Name);
 return nil
}

func(np *EthernetOamLinkTrace) GetMegLevel ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtEthernetOamLinkTrace GetMegLevel
 return "", nil
}

func(np *EthernetOamLinkTrace) SetTransactionIdentifier () error {
 //parameters: SessionPoolHandle TransId
 //AgtEthernetOamLinkTrace SetTransactionIdentifier, m.Object, m.Name);
 return nil
}

func(np *EthernetOamLinkTrace) GetTransactionIdentifier ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtEthernetOamLinkTrace GetTransactionIdentifier
 return "", nil
}

func(np *EthernetOamLinkTrace) EnableSenderIdTlv () error {
 //parameters: SessionPoolHandle
 //AgtEthernetOamLinkTrace EnableSenderIdTlv, m.Object, m.Name);
 return nil
}

func(np *EthernetOamLinkTrace) DisableSenderIdTlv () error {
 //parameters: SessionPoolHandle
 //AgtEthernetOamLinkTrace DisableSenderIdTlv, m.Object, m.Name);
 return nil
}

func(np *EthernetOamLinkTrace) IsSenderIdTlvEnabled () error {
 //parameters: SessionPoolHandle
 //AgtEthernetOamLinkTrace IsSenderIdTlvEnabled, m.Object, m.Name);
 return nil
}

func(np *EthernetOamLinkTrace) SetSenderIdTlv () error {
 //parameters: SessionPoolHandle ChassisIdLength ChassisIdSubType ChassisID MgmtAddressDomain MgmtAddressDomainLength MgmtAddress MgmtAddresslength
 //AgtEthernetOamLinkTrace SetSenderIdTlv, m.Object, m.Name);
 return nil
}

func(np *EthernetOamLinkTrace) GetSenderIdTlv ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtEthernetOamLinkTrace GetSenderIdTlv
 return "", nil
}

func(np *EthernetOamLinkTrace) EnableOrganizationSpecificTlv () error {
 //parameters: SessionPoolHandle
 //AgtEthernetOamLinkTrace EnableOrganizationSpecificTlv, m.Object, m.Name);
 return nil
}

func(np *EthernetOamLinkTrace) DisableOrganizationSpecificTlv () error {
 //parameters: SessionPoolHandle
 //AgtEthernetOamLinkTrace DisableOrganizationSpecificTlv, m.Object, m.Name);
 return nil
}

func(np *EthernetOamLinkTrace) IsOrganizationSpecificTlvEnabled () error {
 //parameters: SessionPoolHandle
 //AgtEthernetOamLinkTrace IsOrganizationSpecificTlvEnabled, m.Object, m.Name);
 return nil
}

func(np *EthernetOamLinkTrace) SetOrganizationSpecificTlv () error {
 //parameters: SessionPoolHandle OrgSpecTlvLen OrgUniqId OrgSpecSubType OrgSpecTlvValue
 //AgtEthernetOamLinkTrace SetOrganizationSpecificTlv, m.Object, m.Name);
 return nil
}

func(np *EthernetOamLinkTrace) GetOrganizationSpecificTlv ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtEthernetOamLinkTrace GetOrganizationSpecificTlv
 return "", nil
}

func(np *EthernetOamLinkTrace) SetBurstSize () error {
 //parameters: SessionPoolHandle BurstSize
 //AgtEthernetOamLinkTrace SetBurstSize, m.Object, m.Name);
 return nil
}

func(np *EthernetOamLinkTrace) GetBurstSize ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtEthernetOamLinkTrace GetBurstSize
 return "", nil
}

func(np *EthernetOamLinkTrace) SetBurstDelay () error {
 //parameters: SessionPoolHandle BurstDelay
 //AgtEthernetOamLinkTrace SetBurstDelay, m.Object, m.Name);
 return nil
}

func(np *EthernetOamLinkTrace) GetBurstDelay ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtEthernetOamLinkTrace GetBurstDelay
 return "", nil
}

func(np *EthernetOamLinkTrace) EnableRenew () error {
 //parameters: SessionPoolHandle
 //AgtEthernetOamLinkTrace EnableRenew, m.Object, m.Name);
 return nil
}

func(np *EthernetOamLinkTrace) DisableRenew () error {
 //parameters: SessionPoolHandle
 //AgtEthernetOamLinkTrace DisableRenew, m.Object, m.Name);
 return nil
}

func(np *EthernetOamLinkTrace) IsRenewEnabled () error {
 //parameters: SessionPoolHandle
 //AgtEthernetOamLinkTrace IsRenewEnabled, m.Object, m.Name);
 return nil
}

func(np *EthernetOamLinkTrace) SetRenewTimeout () error {
 //parameters: SessionPoolHandle RenewTimeout
 //AgtEthernetOamLinkTrace SetRenewTimeout, m.Object, m.Name);
 return nil
}

func(np *EthernetOamLinkTrace) GetRenewTimeout ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtEthernetOamLinkTrace GetRenewTimeout
 return "", nil
}

func(np *EthernetOamLinkTrace) SetResponseTimeout () error {
 //parameters: SessionPoolHandle ResponseTimeout
 //AgtEthernetOamLinkTrace SetResponseTimeout, m.Object, m.Name);
 return nil
}

func(np *EthernetOamLinkTrace) GetResponseTimeout ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtEthernetOamLinkTrace GetResponseTimeout
 return "", nil
}

func(np *EthernetOamLinkTrace) Start () error {
 //parameters: SessionPoolHandle
 //AgtEthernetOamLinkTrace Start, m.Object, m.Name);
 return nil
}

func(np *EthernetOamLinkTrace) ForceRenew () error {
 //parameters: SessionPoolHandle
 //AgtEthernetOamLinkTrace ForceRenew, m.Object, m.Name);
 return nil
}

func(np *EthernetOamLinkTrace) SetIncrementingParameter () error {
 //parameters: SessionPoolHandle IncrementingParameter Value Repeat Increment
 //AgtEthernetOamLinkTrace SetIncrementingParameter, m.Object, m.Name);
 return nil
}

func(np *EthernetOamLinkTrace) GetIncrementingParameter ()(string, error) {
 //parameters: SessionPoolHandle IncrementingParameter
 //AgtEthernetOamLinkTrace GetIncrementingParameter
 return "", nil
}

func(np *EthernetOamLinkTrace) GetVlanPriority ()(string, error) {
 //parameters: DeviceHandle VlanTagIndex
 //AgtEthernetOamLinkTrace GetVlanPriority
 return "", nil
}

func(np *EthernetOamLinkTrace) SetVlanPriority () error {
 //parameters: DeviceHandle VlanTagIndex VlanPriority
 //AgtEthernetOamLinkTrace SetVlanPriority, m.Object, m.Name);
 return nil
}

