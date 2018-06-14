package n2xsdk

type Layer2Learning struct {
 Handler string
}

func(np *Layer2Learnin) EnableAutoSend () error {
 //parameters: DevicePoolHandle
 //AgtLayer2Learning EnableAutoSend, m.Object, m.Name);
 return nil
}

func(np *Layer2Learnin) DisableAutoSend () error {
 //parameters: DevicePoolHandle
 //AgtLayer2Learning DisableAutoSend, m.Object, m.Name);
 return nil
}

func(np *Layer2Learnin) IsAutoSendEnabled () error {
 //parameters: DevicePoolHandle
 //AgtLayer2Learning IsAutoSendEnabled, m.Object, m.Name);
 return nil
}

func(np *Layer2Learnin) SetFrameLength () error {
 //parameters: DevicePoolHandle FrameLength
 //AgtLayer2Learning SetFrameLength, m.Object, m.Name);
 return nil
}

func(np *Layer2Learnin) GetFrameLength ()(string, error) {
 //parameters: DevicePoolHandle
 //AgtLayer2Learning GetFrameLength
 return "", nil
}

func(np *Layer2Learnin) GetCustomState ()(string, error) {
 //parameters: DevicePoolInstance
 //AgtLayer2Learning GetCustomState
 return "", nil
}

func(np *Layer2Learnin) Start () error {
 //parameters: DevicePoolHandle
 //AgtLayer2Learning Start, m.Object, m.Name);
 return nil
}

func(np *Layer2Learnin) Stop () error {
 //parameters: DevicePoolHandle
 //AgtLayer2Learning Stop, m.Object, m.Name);
 return nil
}

func(np *Layer2Learnin) StartByPort () error {
 //parameters: PortHandle
 //AgtLayer2Learning StartByPort, m.Object, m.Name);
 return nil
}

func(np *Layer2Learnin) StopByPort () error {
 //parameters: PortHandle
 //AgtLayer2Learning StopByPort, m.Object, m.Name);
 return nil
}

func(np *Layer2Learnin) StartAll () error {
 //parameters: 
 //AgtLayer2Learning StartAll, m.Object, m.Name);
 return nil
}

func(np *Layer2Learnin) StopAll () error {
 //parameters: 
 //AgtLayer2Learning StopAll, m.Object, m.Name);
 return nil
}

func(np *Layer2Learnin) SetRateControl () error {
 //parameters: PortHandle RateControlMode
 //AgtLayer2Learning SetRateControl, m.Object, m.Name);
 return nil
}

func(np *Layer2Learnin) GetRateControl ()(string, error) {
 //parameters: PortHandle
 //AgtLayer2Learning GetRateControl
 return "", nil
}

func(np *Layer2Learnin) SetFixedRateProfile () error {
 //parameters: PortHandle FixedRate MaxBurstSize
 //AgtLayer2Learning SetFixedRateProfile, m.Object, m.Name);
 return nil
}

func(np *Layer2Learnin) GetFixedRateProfile ()(string, error) {
 //parameters: PortHandle
 //AgtLayer2Learning GetFixedRateProfile
 return "", nil
}

