package n2xsdk

type ncpClientCustomPacke struct {
 Handler string
}

func(np *ncpClientCustomPacke) AddPdu () error {
 //parameters: PeerPoolHandle PduHandleList
 //AgtAncpClientCustomPacket AddPdu, m.Object, m.Name);
 return nil
}

func(np *ncpClientCustomPacke) RemovePdu () error {
 //parameters: PeerPoolHandle
 //AgtAncpClientCustomPacket RemovePdu, m.Object, m.Name);
 return nil
}

func(np *ncpClientCustomPacke) SendCustomPacket () error {
 //parameters: PeerPoolHandle
 //AgtAncpClientCustomPacket SendCustomPacket, m.Object, m.Name);
 return nil
}

