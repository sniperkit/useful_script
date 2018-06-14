package n2xsdk

type FrameRelayAddresses struct {
 Handler string
}

func(np *FrameRelayAddresses) AddPvcPool () error {
 //parameters: PortHandle NumPvcs FirstDlci DlciIncrement Encapsulation
 //AgtFrameRelayAddresses AddPvcPool, m.Object, m.Name);
 return nil
}

func(np *FrameRelayAddresses) RemovePvcPools () error {
 //parameters: PortHandle Count psaPvcPoolHandles
 //AgtFrameRelayAddresses RemovePvcPools, m.Object, m.Name);
 return nil
}

func(np *FrameRelayAddresses) ListPvcPools ()(string, error) {
 //parameters: PortHandle
 //AgtFrameRelayAddresses ListPvcPools
 return "", nil
}

func(np *FrameRelayAddresses) AddPvcs () error {
 //parameters: PortHandle NumPvcs FirstDlci DlciIncrement Encapsulation FirstTesterIpAddress FirstSutIpAddress PrefixLength IpAddressModifier
 //AgtFrameRelayAddresses AddPvcs, m.Object, m.Name);
 return nil
}

func(np *FrameRelayAddresses) RemovePvcs () error {
 //parameters: PortHandle Count psaPvcs
 //AgtFrameRelayAddresses RemovePvcs, m.Object, m.Name);
 return nil
}

func(np *FrameRelayAddresses) ListPvcs ()(string, error) {
 //parameters: PortHandle
 //AgtFrameRelayAddresses ListPvcs
 return "", nil
}

func(np *FrameRelayAddresses) GetPvcCount ()(string, error) {
 //parameters: PortHandle
 //AgtFrameRelayAddresses GetPvcCount
 return "", nil
}

func(np *FrameRelayAddresses) GetMaxPvcs ()(string, error) {
 //parameters: PortHandle
 //AgtFrameRelayAddresses GetMaxPvcs
 return "", nil
}

