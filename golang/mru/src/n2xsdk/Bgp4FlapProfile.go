package n2xsdk

type Bgp4FlapProfile struct {
 Handler string
}

func(np *Bgp4FlapProfile) SetRoutesPerUpdate () error {
 //parameters: PeerPoolHandle Routes
 //AgtBgp4FlapProfile SetRoutesPerUpdate, m.Object, m.Name);
 return nil
}

func(np *Bgp4FlapProfile) GetRoutesPerUpdate ()(string, error) {
 //parameters: PeerPoolHandle
 //AgtBgp4FlapProfile GetRoutesPerUpdate
 return "", nil
}

func(np *Bgp4FlapProfile) GetInterUpdateDelay ()(string, error) {
 //parameters: PeerPoolHandle
 //AgtBgp4FlapProfile GetInterUpdateDelay
 return "", nil
}

func(np *Bgp4FlapProfile) SetInterUpdateDelay () error {
 //parameters: PeerPoolHandle Delay
 //AgtBgp4FlapProfile SetInterUpdateDelay, m.Object, m.Name);
 return nil
}

func(np *Bgp4FlapProfile) GetAdvertiseToWithdrawDelay ()(string, error) {
 //parameters: PeerPoolHandle
 //AgtBgp4FlapProfile GetAdvertiseToWithdrawDelay
 return "", nil
}

func(np *Bgp4FlapProfile) SetAdvertiseToWithdrawDelay () error {
 //parameters: PeerPoolHandle Delay
 //AgtBgp4FlapProfile SetAdvertiseToWithdrawDelay, m.Object, m.Name);
 return nil
}

func(np *Bgp4FlapProfile) GetWithdrawToAdvertiseDelay ()(string, error) {
 //parameters: PeerPoolHandle
 //AgtBgp4FlapProfile GetWithdrawToAdvertiseDelay
 return "", nil
}

func(np *Bgp4FlapProfile) SetWithdrawToAdvertiseDelay () error {
 //parameters: PeerPoolHandle Delay
 //AgtBgp4FlapProfile SetWithdrawToAdvertiseDelay, m.Object, m.Name);
 return nil
}

func(np *Bgp4FlapProfile) SetRoutePools () error {
 //parameters: PeerPoolHandle RoutePoolHandles
 //AgtBgp4FlapProfile SetRoutePools, m.Object, m.Name);
 return nil
}

func(np *Bgp4FlapProfile) GetRoutePools ()(string, error) {
 //parameters: PeerPoolHandle
 //AgtBgp4FlapProfile GetRoutePools
 return "", nil
}

func(np *Bgp4FlapProfile) StartFlap () error {
 //parameters: PeerPoolHandle
 //AgtBgp4FlapProfile StartFlap, m.Object, m.Name);
 return nil
}

func(np *Bgp4FlapProfile) StopFlap () error {
 //parameters: PeerPoolHandle
 //AgtBgp4FlapProfile StopFlap, m.Object, m.Name);
 return nil
}

func(np *Bgp4FlapProfile) GetFlapState ()(string, error) {
 //parameters: PeerPoolHandle
 //AgtBgp4FlapProfile GetFlapState
 return "", nil
}

