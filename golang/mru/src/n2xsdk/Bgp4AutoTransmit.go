package n2xsdk

type Bgp4AutoTransmit struct {
 Handler string
}

func(np *Bgp4AutoTransmi) AddPdus () error {
 //parameters: PeerPoolHandle PduHandleList
 //AgtBgp4AutoTransmit AddPdus, m.Object, m.Name);
 return nil
}

func(np *Bgp4AutoTransmi) RemovePdus () error {
 //parameters: PeerPoolHandle PduHandleList
 //AgtBgp4AutoTransmit RemovePdus, m.Object, m.Name);
 return nil
}

func(np *Bgp4AutoTransmi) ListPdus ()(string, error) {
 //parameters: PeerPoolHandle
 //AgtBgp4AutoTransmit ListPdus
 return "", nil
}

