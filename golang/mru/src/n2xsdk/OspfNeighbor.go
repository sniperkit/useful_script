package n2xsdk

type OspfNeighbor struct {
 Handler string
}

func(np *OspfNeighbor) SetRouterId () error {
 //parameters: NeighborHandle RouterId
 //AgtOspfNeighbor SetRouterId, m.Object, m.Name);
 return nil
}

func(np *OspfNeighbor) GetRouterId ()(string, error) {
 //parameters: NeighborHandle
 //AgtOspfNeighbor GetRouterId
 return "", nil
}

func(np *OspfNeighbor) SetIpAddress () error {
 //parameters: NeighborHandle IpAddress
 //AgtOspfNeighbor SetIpAddress, m.Object, m.Name);
 return nil
}

func(np *OspfNeighbor) GetIpAddress ()(string, error) {
 //parameters: NeighborHandle
 //AgtOspfNeighbor GetIpAddress
 return "", nil
}

func(np *OspfNeighbor) EnableDrEligibility () error {
 //parameters: NeighborHandle
 //AgtOspfNeighbor EnableDrEligibility, m.Object, m.Name);
 return nil
}

func(np *OspfNeighbor) DisableDrEligibility () error {
 //parameters: NeighborHandle
 //AgtOspfNeighbor DisableDrEligibility, m.Object, m.Name);
 return nil
}

func(np *OspfNeighbor) IsDrEligibilityEnabled () error {
 //parameters: NeighborHandle
 //AgtOspfNeighbor IsDrEligibilityEnabled, m.Object, m.Name);
 return nil
}

func(np *OspfNeighbor) SetMaxLsasPerPacket () error {
 //parameters: NeighborHandle MaxLsas
 //AgtOspfNeighbor SetMaxLsasPerPacket, m.Object, m.Name);
 return nil
}

func(np *OspfNeighbor) GetMaxLsasPerPacket ()(string, error) {
 //parameters: NeighborHandle
 //AgtOspfNeighbor GetMaxLsasPerPacket
 return "", nil
}

func(np *OspfNeighbor) GetNeighborType ()(string, error) {
 //parameters: NeighborHandle
 //AgtOspfNeighbor GetNeighborType
 return "", nil
}

func(np *OspfNeighbor) GetRestartState ()(string, error) {
 //parameters: NeighborHandle
 //AgtOspfNeighbor GetRestartState
 return "", nil
}

