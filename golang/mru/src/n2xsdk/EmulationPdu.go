package n2xsdk

type EmulationPdu struct {
 Handler string
}

func(np *EmulationPdu) TransmitPdu () error {
 //parameters: DeviceHandle PduHandle
 //AgtEmulationPdu TransmitPdu, m.Object, m.Name);
 return nil
}

func(np *EmulationPdu) IsTransmitPduSupported () error {
 //parameters: SessionHandle
 //AgtEmulationPdu IsTransmitPduSupported, m.Object, m.Name);
 return nil
}

