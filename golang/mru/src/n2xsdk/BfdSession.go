package n2xsdk

type BfdSession struct {
 Handler string
}

func(np *BfdSession) Shutdown () error {
 //parameters: SessionList
 //AgtBfdSession Shutdown, m.Object, m.Name);
 return nil
}

func(np *BfdSession) NoShutdown () error {
 //parameters: SessionList
 //AgtBfdSession NoShutdown, m.Object, m.Name);
 return nil
}

func(np *BfdSession) Suspend () error {
 //parameters: SessionList
 //AgtBfdSession Suspend, m.Object, m.Name);
 return nil
}

func(np *BfdSession) Resume () error {
 //parameters: SessionList
 //AgtBfdSession Resume, m.Object, m.Name);
 return nil
}

func(np *BfdSession) Reset () error {
 //parameters: SessionList
 //AgtBfdSession Reset, m.Object, m.Name);
 return nil
}

func(np *BfdSession) ResetMeasurements () error {
 //parameters: SessionList
 //AgtBfdSession ResetMeasurements, m.Object, m.Name);
 return nil
}

func(np *BfdSession) SetParameter () error {
 //parameters: SessionPoolHandle Parameter Value
 //AgtBfdSession SetParameter, m.Object, m.Name);
 return nil
}

func(np *BfdSession) GetParameter ()(string, error) {
 //parameters: SessionPoolHandle Parameter
 //AgtBfdSession GetParameter
 return "", nil
}

func(np *BfdSession) SetIncrementingParameter () error {
 //parameters: SessionPoolHandle Parameter Value Repeat Increment
 //AgtBfdSession SetIncrementingParameter, m.Object, m.Name);
 return nil
}

func(np *BfdSession) GetIncrementingParameter ()(string, error) {
 //parameters: SessionPoolHandle Parameter
 //AgtBfdSession GetIncrementingParameter
 return "", nil
}

func(np *BfdSession) IsEchoResponderEnabled () error {
 //parameters: SessionPoolHandle
 //AgtBfdSession IsEchoResponderEnabled, m.Object, m.Name);
 return nil
}

func(np *BfdSession) EnableEchoResponder () error {
 //parameters: SessionPoolHandle
 //AgtBfdSession EnableEchoResponder, m.Object, m.Name);
 return nil
}

func(np *BfdSession) DisableEchoResponder () error {
 //parameters: SessionPoolHandle
 //AgtBfdSession DisableEchoResponder, m.Object, m.Name);
 return nil
}

func(np *BfdSession) IsEchoInitiatorEnabled () error {
 //parameters: SessionPoolHandle
 //AgtBfdSession IsEchoInitiatorEnabled, m.Object, m.Name);
 return nil
}

func(np *BfdSession) EnableEchoInitiator () error {
 //parameters: SessionPoolHandle
 //AgtBfdSession EnableEchoInitiator, m.Object, m.Name);
 return nil
}

func(np *BfdSession) DisableEchoInitiator () error {
 //parameters: SessionPoolHandle
 //AgtBfdSession DisableEchoInitiator, m.Object, m.Name);
 return nil
}

