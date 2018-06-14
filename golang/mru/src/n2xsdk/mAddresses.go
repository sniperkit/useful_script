package n2xsdk

type mAddresses struct {
 Handler string
}

func(np *mAddresses) AddPvcPool () error {
 //parameters: PortHandle NumPvcs FirstVpi VpiIncrement FirstVci VciIncrement Encapsulation
 //AgtAtmAddresses AddPvcPool, m.Object, m.Name);
 return nil
}

func(np *mAddresses) RemovePvcPools () error {
 //parameters: PortHandle Count psaPvcPoolHandles
 //AgtAtmAddresses RemovePvcPools, m.Object, m.Name);
 return nil
}

func(np *mAddresses) ListPvcPools ()(string, error) {
 //parameters: PortHandle
 //AgtAtmAddresses ListPvcPools
 return "", nil
}

func(np *mAddresses) AddPvc () error {
 //parameters: PortHandle Vpi Vci Encapsulation TesterIpAddress SutIpAddress PrefixLength TrafficFlag
 //AgtAtmAddresses AddPvc, m.Object, m.Name);
 return nil
}

func(np *mAddresses) AddIpv6Pvc () error {
 //parameters: PortHandle Vpi Vci Encapsulation TesterIpv6Address SutIpv6Address Ipv6PrefixLength
 //AgtAtmAddresses AddIpv6Pvc, m.Object, m.Name);
 return nil
}

func(np *mAddresses) RemovePvc () error {
 //parameters: PortHandle PvcHandle
 //AgtAtmAddresses RemovePvc, m.Object, m.Name);
 return nil
}

func(np *mAddresses) ListPvcs ()(string, error) {
 //parameters: PortHandle
 //AgtAtmAddresses ListPvcs
 return "", nil
}

func(np *mAddresses) GetPvcCount ()(string, error) {
 //parameters: PortHandle
 //AgtAtmAddresses GetPvcCount
 return "", nil
}

func(np *mAddresses) SetVpiVciRangeForMplsNull () error {
 //parameters: PortHandle MinVpi MaxVpi MinVci MaxVci
 //AgtAtmAddresses SetVpiVciRangeForMplsNull, m.Object, m.Name);
 return nil
}

func(np *mAddresses) GetVpiVciRangeForMplsNull ()(string, error) {
 //parameters: PortHandle
 //AgtAtmAddresses GetVpiVciRangeForMplsNull
 return "", nil
}

