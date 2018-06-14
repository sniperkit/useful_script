package n2xsdk

type NdpEmulation struct {
 Handler string
}

func(np *NdpEmulation) GetState ()(string, error) {
 //parameters: PortHandle
 //AgtNdpEmulation GetState
 return "", nil
}

func(np *NdpEmulation) SendNeighborSolicitationsByAddressPool () error {
 //parameters: PortHandle SutIpv6Address AddressPoolHandle
 //AgtNdpEmulation SendNeighborSolicitationsByAddressPool, m.Object, m.Name);
 return nil
}

func(np *NdpEmulation) SendAllNeighborSolicitationsByAddressPool () error {
 //parameters: PortHandle AddressPoolHandle
 //AgtNdpEmulation SendAllNeighborSolicitationsByAddressPool, m.Object, m.Name);
 return nil
}

func(np *NdpEmulation) EnableSendAllNdpRequestsAtTestStart () error {
 //parameters: 
 //AgtNdpEmulation EnableSendAllNdpRequestsAtTestStart, m.Object, m.Name);
 return nil
}

func(np *NdpEmulation) DisableSendAllNdpRequestsAtTestStart () error {
 //parameters: 
 //AgtNdpEmulation DisableSendAllNdpRequestsAtTestStart, m.Object, m.Name);
 return nil
}

func(np *NdpEmulation) IsSendAllNdpRequestsAtTestStartEnabled () error {
 //parameters: 
 //AgtNdpEmulation IsSendAllNdpRequestsAtTestStartEnabled, m.Object, m.Name);
 return nil
}

func(np *NdpEmulation) GetSelfPacedState ()(string, error) {
 //parameters: PortHandle
 //AgtNdpEmulation GetSelfPacedState
 return "", nil
}

func(np *NdpEmulation) SendNeighborSolicitations () error {
 //parameters: PortHandle VlanIds SutIpv6Address
 //AgtNdpEmulation SendNeighborSolicitations, m.Object, m.Name);
 return nil
}

func(np *NdpEmulation) SendAllNeighborSolicitations () error {
 //parameters: PortHandle
 //AgtNdpEmulation SendAllNeighborSolicitations, m.Object, m.Name);
 return nil
}

func(np *NdpEmulation) SendNeighborSolicitationsBySession () error {
 //parameters: PortHandle SutIpv6Address SessionPoolHandle
 //AgtNdpEmulation SendNeighborSolicitationsBySession, m.Object, m.Name);
 return nil
}

func(np *NdpEmulation) SendAllNeighborSolicitationsBySession () error {
 //parameters: PortHandle SessionPoolHandle
 //AgtNdpEmulation SendAllNeighborSolicitationsBySession, m.Object, m.Name);
 return nil
}

func(np *NdpEmulation) ClearMacTable () error {
 //parameters: PortHandle
 //AgtNdpEmulation ClearMacTable, m.Object, m.Name);
 return nil
}

func(np *NdpEmulation) Enable () error {
 //parameters: PortHandle
 //AgtNdpEmulation Enable, m.Object, m.Name);
 return nil
}

func(np *NdpEmulation) Disable () error {
 //parameters: PortHandle
 //AgtNdpEmulation Disable, m.Object, m.Name);
 return nil
}

func(np *NdpEmulation) EnableSelfPacedResolution () error {
 //parameters: PortHandle
 //AgtNdpEmulation EnableSelfPacedResolution, m.Object, m.Name);
 return nil
}

func(np *NdpEmulation) DisableSelfPacedResolution () error {
 //parameters: PortHandle
 //AgtNdpEmulation DisableSelfPacedResolution, m.Object, m.Name);
 return nil
}

