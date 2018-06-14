package n2xsdk

type RipFlapProfile struct {
 Handler string
}

func(np *RipFlapProfile) GetAdvertiseToWithdrawDelay ()(string, error) {
 //parameters: SessionHandle
 //AgtRipFlapProfile GetAdvertiseToWithdrawDelay
 return "", nil
}

func(np *RipFlapProfile) SetAdvertiseToWithdrawDelay () error {
 //parameters: SessionHandle WithdrawDelay
 //AgtRipFlapProfile SetAdvertiseToWithdrawDelay, m.Object, m.Name);
 return nil
}

func(np *RipFlapProfile) GetWithdrawToAdvertiseDelay ()(string, error) {
 //parameters: SessionHandle
 //AgtRipFlapProfile GetWithdrawToAdvertiseDelay
 return "", nil
}

func(np *RipFlapProfile) SetWithdrawToAdvertiseDelay () error {
 //parameters: SessionHandle AdvertiseDelay
 //AgtRipFlapProfile SetWithdrawToAdvertiseDelay, m.Object, m.Name);
 return nil
}

func(np *RipFlapProfile) StartFlap () error {
 //parameters: SessionHandle
 //AgtRipFlapProfile StartFlap, m.Object, m.Name);
 return nil
}

func(np *RipFlapProfile) StopFlap () error {
 //parameters: SessionHandle
 //AgtRipFlapProfile StopFlap, m.Object, m.Name);
 return nil
}

func(np *RipFlapProfile) GetFlapState ()(string, error) {
 //parameters: SessionHandle
 //AgtRipFlapProfile GetFlapState
 return "", nil
}

func(np *RipFlapProfile) GetRoutePools ()(string, error) {
 //parameters: SessionHandle
 //AgtRipFlapProfile GetRoutePools
 return "", nil
}

func(np *RipFlapProfile) SetRoutePools () error {
 //parameters: SessionHandle Count pRoutePoolHandles
 //AgtRipFlapProfile SetRoutePools, m.Object, m.Name);
 return nil
}

