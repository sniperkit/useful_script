package n2xsdk

type NdpAaRouterCustomPacke struct {
 Handler string
}

func(np *NdpAaRouterCustomPacke) AddPdu () error {
 //parameters: DeviceHandle PduHandleList
 //AgtNdpAaRouterCustomPacket AddPdu, m.Object, m.Name);
 return nil
}

func(np *NdpAaRouterCustomPacke) RemovePdu () error {
 //parameters: DeviceHandle
 //AgtNdpAaRouterCustomPacket RemovePdu, m.Object, m.Name);
 return nil
}

