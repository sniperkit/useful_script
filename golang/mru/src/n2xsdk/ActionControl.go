package n2xsdk

type ActionControl struct {
 Handler string
}

func(np *ctionControl) StartConstantRate () error {
 //parameters: InstanceCount Interval Object Method Parameters
 //AgtActionControl StartConstantRate, m.Object, m.Name);
 return nil
}

func(np *ctionControl) GetState ()(string, error) {
 //parameters: Handle
 //AgtActionControl GetState
 return "", nil
}

func(np *ctionControl) Cancel () error {
 //parameters: Handle
 //AgtActionControl Cancel, m.Object, m.Name);
 return nil
}

func(np *ctionControl) IsSupported () error {
 //parameters: Object Method
 //AgtActionControl IsSupported, m.Object, m.Name);
 return nil
}

