package n2xsdk

type TestTopologyParameters struct {
 Handler string
}

func(np *TestTopologyParameters) GetLinkLayers ()(string, error) {
 //parameters: SessionHandle
 //AgtTestTopologyParameters GetLinkLayers
 return "", nil
}

func(np *TestTopologyParameters) SetLinkLayers () error {
 //parameters: SessionHandle EAgtTopologyParameterCategories
 //AgtTestTopologyParameters SetLinkLayers, m.Object, m.Name);
 return nil
}

func(np *TestTopologyParameters) SelectParameter () error {
 //parameters: SessionHandle Category Parameter
 //AgtTestTopologyParameters SelectParameter, m.Object, m.Name);
 return nil
}

func(np *TestTopologyParameters) DeselectParameter () error {
 //parameters: SessionHandle Category Parameter
 //AgtTestTopologyParameters DeselectParameter, m.Object, m.Name);
 return nil
}

func(np *TestTopologyParameters) IsParameterSelected () error {
 //parameters: SessionHandle Category Parameter
 //AgtTestTopologyParameters IsParameterSelected, m.Object, m.Name);
 return nil
}

func(np *TestTopologyParameters) SetParameterValue () error {
 //parameters: SessionHandle Category Parameter Value
 //AgtTestTopologyParameters SetParameterValue, m.Object, m.Name);
 return nil
}

func(np *TestTopologyParameters) GetParameterValue ()(string, error) {
 //parameters: SessionHandle Category Parameter
 //AgtTestTopologyParameters GetParameterValue
 return "", nil
}

func(np *TestTopologyParameters) SetParameterValueFixed () error {
 //parameters: SessionHandle Category Parameter Value
 //AgtTestTopologyParameters SetParameterValueFixed, m.Object, m.Name);
 return nil
}

func(np *TestTopologyParameters) SetIncrementingParameterValue () error {
 //parameters: SessionHandle Category Parameter StartValue Repeat Increment
 //AgtTestTopologyParameters SetIncrementingParameterValue, m.Object, m.Name);
 return nil
}

func(np *TestTopologyParameters) GetIncrementingParameterValue ()(string, error) {
 //parameters: SessionHandle Category Parameter
 //AgtTestTopologyParameters GetIncrementingParameterValue
 return "", nil
}

func(np *TestTopologyParameters) SetIncrementingParameterValueAndCount () error {
 //parameters: SessionHandle Category Parameter StartValue Repeat Increment Count
 //AgtTestTopologyParameters SetIncrementingParameterValueAndCount, m.Object, m.Name);
 return nil
}

func(np *TestTopologyParameters) GetIncrementingParameterValueAndCount ()(string, error) {
 //parameters: SessionHandle Category Parameter
 //AgtTestTopologyParameters GetIncrementingParameterValueAndCount
 return "", nil
}

func(np *TestTopologyParameters) SetIncrementingParameterValueAndOffset () error {
 //parameters: SessionHandle Category Parameter StartValue Repeat Increment Offset
 //AgtTestTopologyParameters SetIncrementingParameterValueAndOffset, m.Object, m.Name);
 return nil
}

func(np *TestTopologyParameters) GetIncrementingParameterValueAndOffset ()(string, error) {
 //parameters: SessionHandle Category Parameter
 //AgtTestTopologyParameters GetIncrementingParameterValueAndOffset
 return "", nil
}

