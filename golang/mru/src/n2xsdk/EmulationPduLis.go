package n2xsdk

type EmulationPduLis struct {
 Handler string
}

func(np *EmulationPduLis) AddPdu () error {
 //parameters: Count ProtocolList
 //AgtEmulationPduList AddPdu, m.Object, m.Name);
 return nil
}

func(np *EmulationPduLis) AddPduWithPacketType () error {
 //parameters: OuterProtocol PacketType
 //AgtEmulationPduList AddPduWithPacketType, m.Object, m.Name);
 return nil
}

func(np *EmulationPduLis) RemovePdu () error {
 //parameters: EmulationPduHandle
 //AgtEmulationPduList RemovePdu, m.Object, m.Name);
 return nil
}

func(np *EmulationPduLis) SetName () error {
 //parameters: EmulationPduHandle Name
 //AgtEmulationPduList SetName, m.Object, m.Name);
 return nil
}

func(np *EmulationPduLis) GetName ()(string, error) {
 //parameters: EmulationPduHandle
 //AgtEmulationPduList GetName
 return "", nil
}

func(np *EmulationPduLis) GetHandle ()(string, error) {
 //parameters: Name
 //AgtEmulationPduList GetHandle
 return "", nil
}

func(np *EmulationPduLis) ListHandles ()(string, error) {
 //parameters: 
 //AgtEmulationPduList ListHandles
 return "", nil
}

func(np *EmulationPduLis) ListHandlesByOuterProtocol ()(string, error) {
 //parameters: OuterProtocol
 //AgtEmulationPduList ListHandlesByOuterProtocol
 return "", nil
}

