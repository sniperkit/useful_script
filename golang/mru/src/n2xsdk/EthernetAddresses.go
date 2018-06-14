package n2xsdk

type EthernetAddresses struct {
 Handler string
}

func(np *EthernetAddresses) IsVlanSupported () error {
 //parameters: PortHandle
 //AgtEthernetAddresses IsVlanSupported, m.Object, m.Name);
 return nil
}

func(np *EthernetAddresses) IsMacAddressAccessSupported () error {
 //parameters: PortHandle
 //AgtEthernetAddresses IsMacAddressAccessSupported, m.Object, m.Name);
 return nil
}

func(np *EthernetAddresses) AddSutIpAddresses () error {
 //parameters: PortHandle NumOfSutIpAddresses FirstSutIpAddress PrefixLength IpAddressIncrement
 //AgtEthernetAddresses AddSutIpAddresses, m.Object, m.Name);
 return nil
}

func(np *EthernetAddresses) AddSutIpAddress () error {
 //parameters: PortHandle IpAddress
 //AgtEthernetAddresses AddSutIpAddress, m.Object, m.Name);
 return nil
}

func(np *EthernetAddresses) ModifySutIpAddress () error {
 //parameters: PortHandle OldIpAddress NewIpAddress
 //AgtEthernetAddresses ModifySutIpAddress, m.Object, m.Name);
 return nil
}

func(np *EthernetAddresses) RemoveSutIpAddress () error {
 //parameters: PortHandle IpAddress
 //AgtEthernetAddresses RemoveSutIpAddress, m.Object, m.Name);
 return nil
}

func(np *EthernetAddresses) RemoveSutIpAddresses () error {
 //parameters: PortHandle NumAddresses pIpAddresses
 //AgtEthernetAddresses RemoveSutIpAddresses, m.Object, m.Name);
 return nil
}

func(np *EthernetAddresses) GetSutIpAddressCount ()(string, error) {
 //parameters: PortHandle
 //AgtEthernetAddresses GetSutIpAddressCount
 return "", nil
}

func(np *EthernetAddresses) ListSutIpAddresses ()(string, error) {
 //parameters: PortHandle
 //AgtEthernetAddresses ListSutIpAddresses
 return "", nil
}

func(np *EthernetAddresses) ListSutIpAddressesByPool ()(string, error) {
 //parameters: PortHandle hTesterAddressPool
 //AgtEthernetAddresses ListSutIpAddressesByPool
 return "", nil
}

func(np *EthernetAddresses) GetSutMacAddress ()(string, error) {
 //parameters: PortHandle IpAddress
 //AgtEthernetAddresses GetSutMacAddress
 return "", nil
}

func(np *EthernetAddresses) SetSutMacAddress () error {
 //parameters: PortHandle IpAddress MacAddress
 //AgtEthernetAddresses SetSutMacAddress, m.Object, m.Name);
 return nil
}

func(np *EthernetAddresses) AddAddressPools () error {
 //parameters: PortHandle NumOfAddressPools FirstTesterIpAddress PrefixLength NumAddressesPerPool InPoolAddressModifier IpAddressAcrossPoolIncrement VlanEnableFlag VlanNum psaFirstVlanIds IncrementsNum psaVlanIdIncrements
 //AgtEthernetAddresses AddAddressPools, m.Object, m.Name);
 return nil
}

func(np *EthernetAddresses) AddAddressPool () error {
 //parameters: PortHandle FirstTesterIpAddress PrefixLength NumAddresses AddressModifier
 //AgtEthernetAddresses AddAddressPool, m.Object, m.Name);
 return nil
}

func(np *EthernetAddresses) RemoveAddressPool () error {
 //parameters: PortHandle AddressPoolHandle
 //AgtEthernetAddresses RemoveAddressPool, m.Object, m.Name);
 return nil
}

func(np *EthernetAddresses) RemoveAddressPools () error {
 //parameters: PortHandle NumAddressPools pAddressPoolHandles
 //AgtEthernetAddresses RemoveAddressPools, m.Object, m.Name);
 return nil
}

func(np *EthernetAddresses) GetAddressPoolCount ()(string, error) {
 //parameters: PortHandle
 //AgtEthernetAddresses GetAddressPoolCount
 return "", nil
}

func(np *EthernetAddresses) GetAddressPoolLimit ()(string, error) {
 //parameters: PortHandle
 //AgtEthernetAddresses GetAddressPoolLimit
 return "", nil
}

func(np *EthernetAddresses) ListAddressPools ()(string, error) {
 //parameters: PortHandle
 //AgtEthernetAddresses ListAddressPools
 return "", nil
}

func(np *EthernetAddresses) GetTesterIpAddressCount ()(string, error) {
 //parameters: PortHandle
 //AgtEthernetAddresses GetTesterIpAddressCount
 return "", nil
}

func(np *EthernetAddresses) GetVlanCount ()(string, error) {
 //parameters: PortHandle
 //AgtEthernetAddresses GetVlanCount
 return "", nil
}

func(np *EthernetAddresses) ListVlans ()(string, error) {
 //parameters: PortHandle
 //AgtEthernetAddresses ListVlans
 return "", nil
}

func(np *EthernetAddresses) GetAddressPoolCountByVlan ()(string, error) {
 //parameters: PortHandle VlanNum psaVlanIds
 //AgtEthernetAddresses GetAddressPoolCountByVlan
 return "", nil
}

func(np *EthernetAddresses) ListAddressPoolsByVlan ()(string, error) {
 //parameters: PortHandle VlanNum psaVlanIds
 //AgtEthernetAddresses ListAddressPoolsByVlan
 return "", nil
}

func(np *EthernetAddresses) GetSutIpAddressCountByVlan ()(string, error) {
 //parameters: PortHandle VlanNum psaVlanIds
 //AgtEthernetAddresses GetSutIpAddressCountByVlan
 return "", nil
}

func(np *EthernetAddresses) ListSutIpAddressesByVlan ()(string, error) {
 //parameters: PortHandle VlanNum psaVlanIds
 //AgtEthernetAddresses ListSutIpAddressesByVlan
 return "", nil
}

func(np *EthernetAddresses) GetTesterIpAddressCountByVlan ()(string, error) {
 //parameters: PortHandle VlanNum psaVlanIds
 //AgtEthernetAddresses GetTesterIpAddressCountByVlan
 return "", nil
}

func(np *EthernetAddresses) SetVlanEtherType () error {
 //parameters: PortHandle VlanIndex EtherType
 //AgtEthernetAddresses SetVlanEtherType, m.Object, m.Name);
 return nil
}

func(np *EthernetAddresses) GetVlanEtherType ()(string, error) {
 //parameters: PortHandle VlanIndex
 //AgtEthernetAddresses GetVlanEtherType
 return "", nil
}

func(np *EthernetAddresses) ListStackedVlans ()(string, error) {
 //parameters: PortHandle
 //AgtEthernetAddresses ListStackedVlans
 return "", nil
}

