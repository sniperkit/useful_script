package n2xsdk

type FrameTransmitter struct {
 Handler string
}

func(np *FrameTransmitter) AddPdu () error {
 //parameters: Count ProtocolList PduLength
 //AgtFrameTransmitter AddPdu, m.Object, m.Name);
 return nil
}

func(np *FrameTransmitter) AddPduWithPacketType () error {
 //parameters: OuterProtocol PacketType PduLength
 //AgtFrameTransmitter AddPduWithPacketType, m.Object, m.Name);
 return nil
}

func(np *FrameTransmitter) RemovePdu () error {
 //parameters: PduHandle
 //AgtFrameTransmitter RemovePdu, m.Object, m.Name);
 return nil
}

func(np *FrameTransmitter) GetPduLength ()(string, error) {
 //parameters: PduHandle
 //AgtFrameTransmitter GetPduLength
 return "", nil
}

func(np *FrameTransmitter) SetPduLength () error {
 //parameters: PduHandle PduLength
 //AgtFrameTransmitter SetPduLength, m.Object, m.Name);
 return nil
}

func(np *FrameTransmitter) TransmitPdu () error {
 //parameters: PortHandle PduHandle TransmitOptions
 //AgtFrameTransmitter TransmitPdu, m.Object, m.Name);
 return nil
}

func(np *FrameTransmitter) TransmitFrame () error {
 //parameters: PortHandle FrameLength psaFrameBytes TransmitOptions
 //AgtFrameTransmitter TransmitFrame, m.Object, m.Name);
 return nil
}

func(np *FrameTransmitter) GetLastFrameDetails ()(string, error) {
 //parameters: PortHandle
 //AgtFrameTransmitter GetLastFrameDetails
 return "", nil
}

