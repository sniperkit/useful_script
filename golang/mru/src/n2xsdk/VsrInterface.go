package n2xsdk

type VsrInterface struct {
 Handler string
}

func(np *VsrInterface) SetCodeViolationErrorMode () error {
 //parameters: PortHandle ErrorMode
 //AgtVsrInterface SetCodeViolationErrorMode, m.Object, m.Name);
 return nil
}

func(np *VsrInterface) GetCodeViolationErrorMode ()(string, error) {
 //parameters: PortHandle
 //AgtVsrInterface GetCodeViolationErrorMode
 return "", nil
}

func(np *VsrInterface) EnableCrcCorrection () error {
 //parameters: PortHandle
 //AgtVsrInterface EnableCrcCorrection, m.Object, m.Name);
 return nil
}

func(np *VsrInterface) DisableCrcCorrection () error {
 //parameters: PortHandle
 //AgtVsrInterface DisableCrcCorrection, m.Object, m.Name);
 return nil
}

func(np *VsrInterface) IsCrcCorrectionEnabled () error {
 //parameters: PortHandle
 //AgtVsrInterface IsCrcCorrectionEnabled, m.Object, m.Name);
 return nil
}

func(np *VsrInterface) EnableChannelProtection () error {
 //parameters: PortHandle
 //AgtVsrInterface EnableChannelProtection, m.Object, m.Name);
 return nil
}

func(np *VsrInterface) DisableChannelProtection () error {
 //parameters: PortHandle
 //AgtVsrInterface DisableChannelProtection, m.Object, m.Name);
 return nil
}

func(np *VsrInterface) IsChannelProtectionEnabled () error {
 //parameters: PortHandle
 //AgtVsrInterface IsChannelProtectionEnabled, m.Object, m.Name);
 return nil
}

func(np *VsrInterface) EnableTxLoopback () error {
 //parameters: PortHandle
 //AgtVsrInterface EnableTxLoopback, m.Object, m.Name);
 return nil
}

func(np *VsrInterface) DisableTxLoopback () error {
 //parameters: PortHandle
 //AgtVsrInterface DisableTxLoopback, m.Object, m.Name);
 return nil
}

func(np *VsrInterface) IsTxLoopbackEnabled () error {
 //parameters: PortHandle
 //AgtVsrInterface IsTxLoopbackEnabled, m.Object, m.Name);
 return nil
}

func(np *VsrInterface) IsVsrSupported () error {
 //parameters: PortHandle
 //AgtVsrInterface IsVsrSupported, m.Object, m.Name);
 return nil
}

func(np *VsrInterface) SetErrorChannels () error {
 //parameters: PortHandle ErrorType Length psaChannels
 //AgtVsrInterface SetErrorChannels, m.Object, m.Name);
 return nil
}

func(np *VsrInterface) GetErrorChannels ()(string, error) {
 //parameters: PortHandle ErrorType
 //AgtVsrInterface GetErrorChannels
 return "", nil
}

func(np *VsrInterface) PatchcordReversalOn () error {
 //parameters: PortHandle
 //AgtVsrInterface PatchcordReversalOn, m.Object, m.Name);
 return nil
}

func(np *VsrInterface) PatchcordReversalOff () error {
 //parameters: PortHandle
 //AgtVsrInterface PatchcordReversalOff, m.Object, m.Name);
 return nil
}

func(np *VsrInterface) IsPatchcordReversalOn () error {
 //parameters: PortHandle
 //AgtVsrInterface IsPatchcordReversalOn, m.Object, m.Name);
 return nil
}

