package n2xsdk

type EthernetInterface struct {
 Handler string
}

func(np *EthernetInterface) SetFramingMode () error {
 //parameters: PortHandle FramingMode
 //AgtEthernetInterface SetFramingMode, m.Object, m.Name);
 return nil
}

func(np *EthernetInterface) GetFramingMode ()(string, error) {
 //parameters: PortHandle
 //AgtEthernetInterface GetFramingMode
 return "", nil
}

func(np *EthernetInterface) SetRj45Mode () error {
 //parameters: PortHandle CrossoverMode
 //AgtEthernetInterface SetRj45Mode, m.Object, m.Name);
 return nil
}

func(np *EthernetInterface) GetRj45Mode ()(string, error) {
 //parameters: PortHandle
 //AgtEthernetInterface GetRj45Mode
 return "", nil
}

func(np *EthernetInterface) IsRj45AutoCrossoverSupported () error {
 //parameters: PortHandle
 //AgtEthernetInterface IsRj45AutoCrossoverSupported, m.Object, m.Name);
 return nil
}

func(np *EthernetInterface) EnableFramePadding () error {
 //parameters: PortHandle
 //AgtEthernetInterface EnableFramePadding, m.Object, m.Name);
 return nil
}

func(np *EthernetInterface) DisableFramePadding () error {
 //parameters: PortHandle
 //AgtEthernetInterface DisableFramePadding, m.Object, m.Name);
 return nil
}

func(np *EthernetInterface) IsFramePaddingEnabled () error {
 //parameters: PortHandle
 //AgtEthernetInterface IsFramePaddingEnabled, m.Object, m.Name);
 return nil
}

func(np *EthernetInterface) EnableJumboFrames () error {
 //parameters: PortHandle
 //AgtEthernetInterface EnableJumboFrames, m.Object, m.Name);
 return nil
}

func(np *EthernetInterface) DisableJumboFrames () error {
 //parameters: PortHandle
 //AgtEthernetInterface DisableJumboFrames, m.Object, m.Name);
 return nil
}

func(np *EthernetInterface) IsJumboFramesEnabled () error {
 //parameters: PortHandle
 //AgtEthernetInterface IsJumboFramesEnabled, m.Object, m.Name);
 return nil
}

func(np *EthernetInterface) SetJumboFrameMtuSize () error {
 //parameters: PortHandle JumboFrameMtuSizeBytes
 //AgtEthernetInterface SetJumboFrameMtuSize, m.Object, m.Name);
 return nil
}

func(np *EthernetInterface) GetJumboFrameMtuSize ()(string, error) {
 //parameters: PortHandle
 //AgtEthernetInterface GetJumboFrameMtuSize
 return "", nil
}

func(np *EthernetInterface) GetMaxJumboFrameMtuSize ()(string, error) {
 //parameters: PortHandle
 //AgtEthernetInterface GetMaxJumboFrameMtuSize
 return "", nil
}

func(np *EthernetInterface) IsEthernetSupported () error {
 //parameters: PortHandle
 //AgtEthernetInterface IsEthernetSupported, m.Object, m.Name);
 return nil
}

func(np *EthernetInterface) IsEthernetPcsSupported () error {
 //parameters: PortHandle
 //AgtEthernetInterface IsEthernetPcsSupported, m.Object, m.Name);
 return nil
}

func(np *EthernetInterface) IsEthernet10100Supported () error {
 //parameters: PortHandle
 //AgtEthernetInterface IsEthernet10100Supported, m.Object, m.Name);
 return nil
}

func(np *EthernetInterface) GetLinkState ()(string, error) {
 //parameters: PortHandle
 //AgtEthernetInterface GetLinkState
 return "", nil
}

func(np *EthernetInterface) SetCollisionBackoff () error {
 //parameters: PortHandle Value
 //AgtEthernetInterface SetCollisionBackoff, m.Object, m.Name);
 return nil
}

func(np *EthernetInterface) GetCollisionBackoff ()(string, error) {
 //parameters: PortHandle
 //AgtEthernetInterface GetCollisionBackoff
 return "", nil
}

func(np *EthernetInterface) EnableCollisionGeneration () error {
 //parameters: PortHandle
 //AgtEthernetInterface EnableCollisionGeneration, m.Object, m.Name);
 return nil
}

func(np *EthernetInterface) DisableCollisionGeneration () error {
 //parameters: PortHandle
 //AgtEthernetInterface DisableCollisionGeneration, m.Object, m.Name);
 return nil
}

func(np *EthernetInterface) IsCollisionGenerationEnabled () error {
 //parameters: PortHandle
 //AgtEthernetInterface IsCollisionGenerationEnabled, m.Object, m.Name);
 return nil
}

func(np *EthernetInterface) EnablePauseFrameResponse () error {
 //parameters: PortHandle
 //AgtEthernetInterface EnablePauseFrameResponse, m.Object, m.Name);
 return nil
}

func(np *EthernetInterface) DisablePauseFrameResponse () error {
 //parameters: PortHandle
 //AgtEthernetInterface DisablePauseFrameResponse, m.Object, m.Name);
 return nil
}

func(np *EthernetInterface) IsPauseFrameResponseEnabled () error {
 //parameters: PortHandle
 //AgtEthernetInterface IsPauseFrameResponseEnabled, m.Object, m.Name);
 return nil
}

func(np *EthernetInterface) SetMinimumInterFrameGap () error {
 //parameters: PortHandle Value
 //AgtEthernetInterface SetMinimumInterFrameGap, m.Object, m.Name);
 return nil
}

func(np *EthernetInterface) GetMinimumInterFrameGap ()(string, error) {
 //parameters: PortHandle
 //AgtEthernetInterface GetMinimumInterFrameGap
 return "", nil
}

func(np *EthernetInterface) GetMinimumInterFrameGapRange ()(string, error) {
 //parameters: PortHandle
 //AgtEthernetInterface GetMinimumInterFrameGapRange
 return "", nil
}

func(np *EthernetInterface) SetConfigRegister () error {
 //parameters: PortHandle Register RegisterValue
 //AgtEthernetInterface SetConfigRegister, m.Object, m.Name);
 return nil
}

func(np *EthernetInterface) GetConfigRegister ()(string, error) {
 //parameters: PortHandle Register
 //AgtEthernetInterface GetConfigRegister
 return "", nil
}

func(np *EthernetInterface) SetConfigRegisterField () error {
 //parameters: PortHandle Register RegisterField FieldValue
 //AgtEthernetInterface SetConfigRegisterField, m.Object, m.Name);
 return nil
}

func(np *EthernetInterface) GetConfigRegisterField ()(string, error) {
 //parameters: PortHandle Register RegisterField
 //AgtEthernetInterface GetConfigRegisterField
 return "", nil
}

func(np *EthernetInterface) EnableAutoNegotiation () error {
 //parameters: PortHandle
 //AgtEthernetInterface EnableAutoNegotiation, m.Object, m.Name);
 return nil
}

func(np *EthernetInterface) DisableAutoNegotiation () error {
 //parameters: PortHandle
 //AgtEthernetInterface DisableAutoNegotiation, m.Object, m.Name);
 return nil
}

func(np *EthernetInterface) IsAutoNegotiationEnabled () error {
 //parameters: PortHandle
 //AgtEthernetInterface IsAutoNegotiationEnabled, m.Object, m.Name);
 return nil
}

func(np *EthernetInterface) RestartAutoNegotiation () error {
 //parameters: PortHandle
 //AgtEthernetInterface RestartAutoNegotiation, m.Object, m.Name);
 return nil
}

func(np *EthernetInterface) SetLinkSpeed () error {
 //parameters: PortHandle LinkSpeed
 //AgtEthernetInterface SetLinkSpeed, m.Object, m.Name);
 return nil
}

func(np *EthernetInterface) GetLinkSpeed ()(string, error) {
 //parameters: PortHandle
 //AgtEthernetInterface GetLinkSpeed
 return "", nil
}

func(np *EthernetInterface) SetDuplexMode () error {
 //parameters: PortHandle DuplexMode
 //AgtEthernetInterface SetDuplexMode, m.Object, m.Name);
 return nil
}

func(np *EthernetInterface) GetDuplexMode ()(string, error) {
 //parameters: PortHandle
 //AgtEthernetInterface GetDuplexMode
 return "", nil
}

func(np *EthernetInterface) SetPonPortMode () error {
 //parameters: PortHandle PortMode
 //AgtEthernetInterface SetPonPortMode, m.Object, m.Name);
 return nil
}

func(np *EthernetInterface) GetPonPortMode ()(string, error) {
 //parameters: PortHandle
 //AgtEthernetInterface GetPonPortMode
 return "", nil
}

