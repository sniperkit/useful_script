package n2xsdk

type PosInterface struct {
 Handler string
}

func(np *PosInterface) SetFramingMode () error {
 //parameters: PortHandle Framing
 //AgtPosInterface SetFramingMode, m.Object, m.Name);
 return nil
}

func(np *PosInterface) GetFramingMode ()(string, error) {
 //parameters: PortHandle
 //AgtPosInterface GetFramingMode
 return "", nil
}

func(np *PosInterface) GetSupportedFramingModes ()(string, error) {
 //parameters: PortHandle
 //AgtPosInterface GetSupportedFramingModes
 return "", nil
}

func(np *PosInterface) ScramblerOn () error {
 //parameters: PortHandle
 //AgtPosInterface ScramblerOn, m.Object, m.Name);
 return nil
}

func(np *PosInterface) ScramblerOff () error {
 //parameters: PortHandle
 //AgtPosInterface ScramblerOff, m.Object, m.Name);
 return nil
}

func(np *PosInterface) IsScramblerOn () error {
 //parameters: PortHandle
 //AgtPosInterface IsScramblerOn, m.Object, m.Name);
 return nil
}

func(np *PosInterface) DescramblerOn () error {
 //parameters: PortHandle
 //AgtPosInterface DescramblerOn, m.Object, m.Name);
 return nil
}

func(np *PosInterface) DescramblerOff () error {
 //parameters: PortHandle
 //AgtPosInterface DescramblerOff, m.Object, m.Name);
 return nil
}

func(np *PosInterface) IsDescramblerOn () error {
 //parameters: PortHandle
 //AgtPosInterface IsDescramblerOn, m.Object, m.Name);
 return nil
}

func(np *PosInterface) IsPosSupported () error {
 //parameters: PortHandle
 //AgtPosInterface IsPosSupported, m.Object, m.Name);
 return nil
}

func(np *PosInterface) SetMinimumInterFrameGap () error {
 //parameters: PortHandle InterFrameGap
 //AgtPosInterface SetMinimumInterFrameGap, m.Object, m.Name);
 return nil
}

func(np *PosInterface) GetMinimumInterFrameGap ()(string, error) {
 //parameters: PortHandle
 //AgtPosInterface GetMinimumInterFrameGap
 return "", nil
}

func(np *PosInterface) GetSupportedMinimumInterFrameGaps ()(string, error) {
 //parameters: PortHandle
 //AgtPosInterface GetSupportedMinimumInterFrameGaps
 return "", nil
}

