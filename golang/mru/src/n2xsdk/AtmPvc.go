package n2xsdk

type AtmPvc struct {
 Handler string
}

func(np *mPvc) GetVpiVci ()(string, error) {
 //parameters: PvcHandle
 //AgtAtmPvc GetVpiVci
 return "", nil
}

func(np *mPvc) SetVpiVci () error {
 //parameters: PvcHandle Vpi Vci
 //AgtAtmPvc SetVpiVci, m.Object, m.Name);
 return nil
}

func(np *mPvc) GetEncapsulation ()(string, error) {
 //parameters: PvcHandle
 //AgtAtmPvc GetEncapsulation
 return "", nil
}

func(np *mPvc) SetEncapsulation () error {
 //parameters: PvcHandle Encapsulation
 //AgtAtmPvc SetEncapsulation, m.Object, m.Name);
 return nil
}

func(np *mPvc) GetIpAddresses ()(string, error) {
 //parameters: PvcHandle
 //AgtAtmPvc GetIpAddresses
 return "", nil
}

func(np *mPvc) SetIpAddresses () error {
 //parameters: PvcHandle TesterIpAddress SutIpAddress PrefixLength
 //AgtAtmPvc SetIpAddresses, m.Object, m.Name);
 return nil
}

func(np *mPvc) GetIpv6Addresses ()(string, error) {
 //parameters: PvcHandle
 //AgtAtmPvc GetIpv6Addresses
 return "", nil
}

func(np *mPvc) SetIpv6Addresses () error {
 //parameters: PvcHandle TesterIpAddress SutIpAddress PrefixLength
 //AgtAtmPvc SetIpv6Addresses, m.Object, m.Name);
 return nil
}

func(np *mPvc) EnableTraffic () error {
 //parameters: PvcHandle
 //AgtAtmPvc EnableTraffic, m.Object, m.Name);
 return nil
}

func(np *mPvc) DisableTraffic () error {
 //parameters: PvcHandle
 //AgtAtmPvc DisableTraffic, m.Object, m.Name);
 return nil
}

func(np *mPvc) IsTrafficEnabled () error {
 //parameters: PvcHandle
 //AgtAtmPvc IsTrafficEnabled, m.Object, m.Name);
 return nil
}

func(np *mPvc) EnableIpv6 () error {
 //parameters: PvcHandle
 //AgtAtmPvc EnableIpv6, m.Object, m.Name);
 return nil
}

func(np *mPvc) DisableIpv6 () error {
 //parameters: PvcHandle
 //AgtAtmPvc DisableIpv6, m.Object, m.Name);
 return nil
}

func(np *mPvc) IsIpv6Enabled () error {
 //parameters: PvcHandle
 //AgtAtmPvc IsIpv6Enabled, m.Object, m.Name);
 return nil
}

