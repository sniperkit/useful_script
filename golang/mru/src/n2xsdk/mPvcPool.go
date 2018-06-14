package n2xsdk

type mPvcPool struct {
 Handler string
}

func(np *mPvcPool) SetNumberOfPvcs () error {
 //parameters: PvcPoolHandle NumPvcs
 //AgtAtmPvcPool SetNumberOfPvcs, m.Object, m.Name);
 return nil
}

func(np *mPvcPool) GetNumberOfPvcs ()(string, error) {
 //parameters: PvcPoolHandle
 //AgtAtmPvcPool GetNumberOfPvcs
 return "", nil
}

func(np *mPvcPool) SetVpiVcis () error {
 //parameters: PvcPoolHandle FirstVpi VpiIncrement FirstVci VciIncrement
 //AgtAtmPvcPool SetVpiVcis, m.Object, m.Name);
 return nil
}

func(np *mPvcPool) GetVpiVcis ()(string, error) {
 //parameters: PvcPoolHandle
 //AgtAtmPvcPool GetVpiVcis
 return "", nil
}

func(np *mPvcPool) SetEncapsulation () error {
 //parameters: PvcPoolHandle Encapsulation
 //AgtAtmPvcPool SetEncapsulation, m.Object, m.Name);
 return nil
}

func(np *mPvcPool) GetEncapsulation ()(string, error) {
 //parameters: PvcPoolHandle
 //AgtAtmPvcPool GetEncapsulation
 return "", nil
}

func(np *mPvcPool) SetIpAddresses () error {
 //parameters: PvcPoolHandle FirstTesterIpv4Address FirstSutIpv4Address Ipv4PrefixLength Ipv4Modifier
 //AgtAtmPvcPool SetIpAddresses, m.Object, m.Name);
 return nil
}

func(np *mPvcPool) GetIpAddresses ()(string, error) {
 //parameters: PvcPoolHandle
 //AgtAtmPvcPool GetIpAddresses
 return "", nil
}

func(np *mPvcPool) SetIpv6Addresses () error {
 //parameters: PvcPoolHandle FirstTesterIpv6Address FirstSutIpv6Address Ipv6PrefixLength Ipv6Modifier
 //AgtAtmPvcPool SetIpv6Addresses, m.Object, m.Name);
 return nil
}

func(np *mPvcPool) GetIpv6Addresses ()(string, error) {
 //parameters: PvcPoolHandle
 //AgtAtmPvcPool GetIpv6Addresses
 return "", nil
}

func(np *mPvcPool) SetMacAddresses () error {
 //parameters: PvcPoolHandle FirstTesterMacAddress SutMacAddress
 //AgtAtmPvcPool SetMacAddresses, m.Object, m.Name);
 return nil
}

func(np *mPvcPool) GetMacAddresses ()(string, error) {
 //parameters: PvcPoolHandle
 //AgtAtmPvcPool GetMacAddresses
 return "", nil
}

func(np *mPvcPool) EnableTrafficDestinations () error {
 //parameters: PvcPoolHandle
 //AgtAtmPvcPool EnableTrafficDestinations, m.Object, m.Name);
 return nil
}

func(np *mPvcPool) DisableTrafficDestinations () error {
 //parameters: PvcPoolHandle
 //AgtAtmPvcPool DisableTrafficDestinations, m.Object, m.Name);
 return nil
}

func(np *mPvcPool) IsTrafficDestinationEnabled () error {
 //parameters: PvcPoolHandle
 //AgtAtmPvcPool IsTrafficDestinationEnabled, m.Object, m.Name);
 return nil
}

