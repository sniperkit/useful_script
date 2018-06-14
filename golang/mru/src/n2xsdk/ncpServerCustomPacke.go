package n2xsdk

type ncpServerCustomPacke struct {
 Handler string
}

func(np *ncpServerCustomPacke) AddPdu () error {
 //parameters: PeerPoolHandle PduHandleList
 //AgtAncpServerCustomPacket AddPdu, m.Object, m.Name);
 return nil
}

func(np *ncpServerCustomPacke) RemovePdu () error {
 //parameters: PeerPoolHandle
 //AgtAncpServerCustomPacket RemovePdu, m.Object, m.Name);
 return nil
}

func(np *ncpServerCustomPacke) SendCustomPacket () error {
 //parameters: PeerPoolHandle
 //AgtAncpServerCustomPacket SendCustomPacket, m.Object, m.Name);
 return nil
}

