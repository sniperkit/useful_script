package n2xsdk

type ArpEmulation struct {
 Handler string
}

func(np *rpEmulation) GetState ()(string, error) {
 //parameters: PortHandle
 //AgtArpEmulation GetState
 return "", nil
}

func(np *rpEmulation) SendAllArpRequests () error {
 //parameters: PortHandle
 //AgtArpEmulation SendAllArpRequests, m.Object, m.Name);
 return nil
}

func(np *rpEmulation) SendArpRequestsByAddressPool () error {
 //parameters: PortHandle SutIpAddress AddressPoolHandle
 //AgtArpEmulation SendArpRequestsByAddressPool, m.Object, m.Name);
 return nil
}

func(np *rpEmulation) SendAllArpRequestsByAddressPool () error {
 //parameters: PortHandle AddressPoolHandle
 //AgtArpEmulation SendAllArpRequestsByAddressPool, m.Object, m.Name);
 return nil
}

func(np *rpEmulation) EnableSendAllArpRequestsAtTestStart () error {
 //parameters: 
 //AgtArpEmulation EnableSendAllArpRequestsAtTestStart, m.Object, m.Name);
 return nil
}

func(np *rpEmulation) DisableSendAllArpRequestsAtTestStart () error {
 //parameters: 
 //AgtArpEmulation DisableSendAllArpRequestsAtTestStart, m.Object, m.Name);
 return nil
}

func(np *rpEmulation) IsSendAllArpRequestsAtTestStartEnabled () error {
 //parameters: 
 //AgtArpEmulation IsSendAllArpRequestsAtTestStartEnabled, m.Object, m.Name);
 return nil
}

func(np *rpEmulation) GetSelfPacedState ()(string, error) {
 //parameters: PortHandle
 //AgtArpEmulation GetSelfPacedState
 return "", nil
}

func(np *rpEmulation) SendArpRequests () error {
 //parameters: PortHandle SutIpAddress
 //AgtArpEmulation SendArpRequests, m.Object, m.Name);
 return nil
}

func(np *rpEmulation) SendArpRequestsBySession () error {
 //parameters: PortHandle SutIpAddress SessionPoolHandle
 //AgtArpEmulation SendArpRequestsBySession, m.Object, m.Name);
 return nil
}

func(np *rpEmulation) SendAllArpRequestsBySession () error {
 //parameters: PortHandle SessionPoolHandle
 //AgtArpEmulation SendAllArpRequestsBySession, m.Object, m.Name);
 return nil
}

func(np *rpEmulation) ClearMacTable () error {
 //parameters: PortHandle
 //AgtArpEmulation ClearMacTable, m.Object, m.Name);
 return nil
}

func(np *rpEmulation) Enable () error {
 //parameters: PortHandle
 //AgtArpEmulation Enable, m.Object, m.Name);
 return nil
}

func(np *rpEmulation) Disable () error {
 //parameters: PortHandle
 //AgtArpEmulation Disable, m.Object, m.Name);
 return nil
}

func(np *rpEmulation) EnableSelfPacedResolution () error {
 //parameters: PortHandle
 //AgtArpEmulation EnableSelfPacedResolution, m.Object, m.Name);
 return nil
}

func(np *rpEmulation) DisableSelfPacedResolution () error {
 //parameters: PortHandle
 //AgtArpEmulation DisableSelfPacedResolution, m.Object, m.Name);
 return nil
}

