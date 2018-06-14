package n2xsdk

type FrameRelayPvcPool struct {
 Handler string
}

func(np *FrameRelayPvcPool) SetNumberOfPvcs () error {
 //parameters: PvcPoolHandle NumPvcs
 //AgtFrameRelayPvcPool SetNumberOfPvcs, m.Object, m.Name);
 return nil
}

func(np *FrameRelayPvcPool) GetNumberOfPvcs ()(string, error) {
 //parameters: PvcPoolHandle
 //AgtFrameRelayPvcPool GetNumberOfPvcs
 return "", nil
}

func(np *FrameRelayPvcPool) SetDlcis () error {
 //parameters: PvcPoolHandle FirstDlci DlciIncrement
 //AgtFrameRelayPvcPool SetDlcis, m.Object, m.Name);
 return nil
}

func(np *FrameRelayPvcPool) GetDlcis ()(string, error) {
 //parameters: PvcPoolHandle
 //AgtFrameRelayPvcPool GetDlcis
 return "", nil
}

func(np *FrameRelayPvcPool) SetEncapsulation () error {
 //parameters: PvcPoolHandle Encapsulation
 //AgtFrameRelayPvcPool SetEncapsulation, m.Object, m.Name);
 return nil
}

func(np *FrameRelayPvcPool) GetEncapsulation ()(string, error) {
 //parameters: PvcPoolHandle
 //AgtFrameRelayPvcPool GetEncapsulation
 return "", nil
}

func(np *FrameRelayPvcPool) SetIpAddresses () error {
 //parameters: PvcPoolHandle FirstTesterIpv4Address FirstSutIpv4Address Ipv4PrefixLength Ipv4Modifier
 //AgtFrameRelayPvcPool SetIpAddresses, m.Object, m.Name);
 return nil
}

func(np *FrameRelayPvcPool) GetIpAddresses ()(string, error) {
 //parameters: PvcPoolHandle
 //AgtFrameRelayPvcPool GetIpAddresses
 return "", nil
}

func(np *FrameRelayPvcPool) SetIpv6Addresses () error {
 //parameters: PvcPoolHandle FirstTesterIpv6Address FirstSutIpv6Address Ipv6PrefixLength Ipv6Modifier
 //AgtFrameRelayPvcPool SetIpv6Addresses, m.Object, m.Name);
 return nil
}

func(np *FrameRelayPvcPool) GetIpv6Addresses ()(string, error) {
 //parameters: PvcPoolHandle
 //AgtFrameRelayPvcPool GetIpv6Addresses
 return "", nil
}

func(np *FrameRelayPvcPool) EnableTrafficDestinations () error {
 //parameters: PvcPoolHandle
 //AgtFrameRelayPvcPool EnableTrafficDestinations, m.Object, m.Name);
 return nil
}

func(np *FrameRelayPvcPool) DisableTrafficDestinations () error {
 //parameters: PvcPoolHandle
 //AgtFrameRelayPvcPool DisableTrafficDestinations, m.Object, m.Name);
 return nil
}

func(np *FrameRelayPvcPool) IsTrafficDestinationEnabled () error {
 //parameters: PvcPoolHandle
 //AgtFrameRelayPvcPool IsTrafficDestinationEnabled, m.Object, m.Name);
 return nil
}

