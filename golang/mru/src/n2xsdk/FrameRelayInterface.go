package n2xsdk

type FrameRelayInterface struct {
 Handler string
}

func(np *FrameRelayInterface) ScramblerOn () error {
 //parameters: PortHandle
 //AgtFrameRelayInterface ScramblerOn, m.Object, m.Name);
 return nil
}

func(np *FrameRelayInterface) ScramblerOff () error {
 //parameters: PortHandle
 //AgtFrameRelayInterface ScramblerOff, m.Object, m.Name);
 return nil
}

func(np *FrameRelayInterface) IsScramblerOn () error {
 //parameters: PortHandle
 //AgtFrameRelayInterface IsScramblerOn, m.Object, m.Name);
 return nil
}

func(np *FrameRelayInterface) DescramblerOn () error {
 //parameters: PortHandle
 //AgtFrameRelayInterface DescramblerOn, m.Object, m.Name);
 return nil
}

func(np *FrameRelayInterface) DescramblerOff () error {
 //parameters: PortHandle
 //AgtFrameRelayInterface DescramblerOff, m.Object, m.Name);
 return nil
}

func(np *FrameRelayInterface) IsDescramblerOn () error {
 //parameters: PortHandle
 //AgtFrameRelayInterface IsDescramblerOn, m.Object, m.Name);
 return nil
}

func(np *FrameRelayInterface) SetFcsLength () error {
 //parameters: PortHandle FcsLength
 //AgtFrameRelayInterface SetFcsLength, m.Object, m.Name);
 return nil
}

func(np *FrameRelayInterface) GetFcsLength ()(string, error) {
 //parameters: PortHandle
 //AgtFrameRelayInterface GetFcsLength
 return "", nil
}

func(np *FrameRelayInterface) SetMinimumInterFrameGap () error {
 //parameters: PortHandle InterFrameGap
 //AgtFrameRelayInterface SetMinimumInterFrameGap, m.Object, m.Name);
 return nil
}

func(np *FrameRelayInterface) GetMinimumInterFrameGap ()(string, error) {
 //parameters: PortHandle
 //AgtFrameRelayInterface GetMinimumInterFrameGap
 return "", nil
}

func(np *FrameRelayInterface) GetSupportedMinimumInterFrameGaps ()(string, error) {
 //parameters: PortHandle
 //AgtFrameRelayInterface GetSupportedMinimumInterFrameGaps
 return "", nil
}

func(np *FrameRelayInterface) SetMtuSize () error {
 //parameters: PortHandle MtuSize
 //AgtFrameRelayInterface SetMtuSize, m.Object, m.Name);
 return nil
}

func(np *FrameRelayInterface) GetMtuSize ()(string, error) {
 //parameters: PortHandle
 //AgtFrameRelayInterface GetMtuSize
 return "", nil
}

func(np *FrameRelayInterface) SetMplsEncapsulation () error {
 //parameters: PortHandle Encapsulation
 //AgtFrameRelayInterface SetMplsEncapsulation, m.Object, m.Name);
 return nil
}

func(np *FrameRelayInterface) GetMplsEncapsulation ()(string, error) {
 //parameters: PortHandle
 //AgtFrameRelayInterface GetMplsEncapsulation
 return "", nil
}

func(np *FrameRelayInterface) IsFrameRelaySupported () error {
 //parameters: PortHandle
 //AgtFrameRelayInterface IsFrameRelaySupported, m.Object, m.Name);
 return nil
}

