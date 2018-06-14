package n2xsdk

type Bgp4Capabilities struct {
 Handler string
}

func(np *Bgp4Capabilities) SetSessionCapabilities () error {
 //parameters: PeerPoolHandle EmulationPduHandle
 //AgtBgp4Capabilities SetSessionCapabilities, m.Object, m.Name);
 return nil
}

func(np *Bgp4Capabilities) GetSessionCapabilities ()(string, error) {
 //parameters: PeerPoolHandle
 //AgtBgp4Capabilities GetSessionCapabilities
 return "", nil
}

func(np *Bgp4Capabilities) ClearSessionCapabilities () error {
 //parameters: PeerPoolHandle
 //AgtBgp4Capabilities ClearSessionCapabilities, m.Object, m.Name);
 return nil
}

