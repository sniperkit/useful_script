package n2xsdk

type RstpSession struct {
 Handler string
}

func(np *RstpSession) ForceMigrationCheck () error {
 //parameters: SessionInstanceHandle
 //AgtRstpSession ForceMigrationCheck, m.Object, m.Name);
 return nil
}

func(np *RstpSession) EnablePvstPlus () error {
 //parameters: SessionPoolHandle
 //AgtRstpSession EnablePvstPlus, m.Object, m.Name);
 return nil
}

func(np *RstpSession) DisablePvstPlus () error {
 //parameters: SessionPoolHandle
 //AgtRstpSession DisablePvstPlus, m.Object, m.Name);
 return nil
}

func(np *RstpSession) IsPvstPlusEnabled () error {
 //parameters: SessionPoolHandle
 //AgtRstpSession IsPvstPlusEnabled, m.Object, m.Name);
 return nil
}

func(np *RstpSession) OpenActive () error {
 //parameters: SessionInstanceHandle
 //AgtRstpSession OpenActive, m.Object, m.Name);
 return nil
}

func(np *RstpSession) OpenPassive () error {
 //parameters: SessionInstanceHandle
 //AgtRstpSession OpenPassive, m.Object, m.Name);
 return nil
}

func(np *RstpSession) Close () error {
 //parameters: SessionInstanceHandle
 //AgtRstpSession Close, m.Object, m.Name);
 return nil
}

func(np *RstpSession) Reset () error {
 //parameters: SessionInstanceHandle
 //AgtRstpSession Reset, m.Object, m.Name);
 return nil
}

func(np *RstpSession) ResetMeasurements () error {
 //parameters: SessionInstanceHandle
 //AgtRstpSession ResetMeasurements, m.Object, m.Name);
 return nil
}

func(np *RstpSession) SignalTopologyChange () error {
 //parameters: SessionInstanceHandle
 //AgtRstpSession SignalTopologyChange, m.Object, m.Name);
 return nil
}

func(np *RstpSession) GetIdentifierEncoding ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtRstpSession GetIdentifierEncoding
 return "", nil
}

func(np *RstpSession) SetIdentifierEncoding () error {
 //parameters: SessionPoolHandle IdentifierEncoding
 //AgtRstpSession SetIdentifierEncoding, m.Object, m.Name);
 return nil
}

func(np *RstpSession) GetBridgeIdentifier ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtRstpSession GetBridgeIdentifier
 return "", nil
}

func(np *RstpSession) SetBridgeIdentifier () error {
 //parameters: SessionPoolHandle Priority MacAddress
 //AgtRstpSession SetBridgeIdentifier, m.Object, m.Name);
 return nil
}

func(np *RstpSession) GetPortIdentifier ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtRstpSession GetPortIdentifier
 return "", nil
}

func(np *RstpSession) SetPortIdentifier () error {
 //parameters: SessionPoolHandle Priority Number
 //AgtRstpSession SetPortIdentifier, m.Object, m.Name);
 return nil
}

func(np *RstpSession) GetPathCostMode ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtRstpSession GetPathCostMode
 return "", nil
}

func(np *RstpSession) SetPathCostMode () error {
 //parameters: SessionPoolHandle PathCostMode
 //AgtRstpSession SetPathCostMode, m.Object, m.Name);
 return nil
}

func(np *RstpSession) GetPathCostValue ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtRstpSession GetPathCostValue
 return "", nil
}

func(np *RstpSession) SetPathCostValue () error {
 //parameters: SessionPoolHandle PathCost
 //AgtRstpSession SetPathCostValue, m.Object, m.Name);
 return nil
}

func(np *RstpSession) GetRootIdentifier ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtRstpSession GetRootIdentifier
 return "", nil
}

func(np *RstpSession) SetRootIdentifier () error {
 //parameters: SessionPoolHandle Priority MacAddress
 //AgtRstpSession SetRootIdentifier, m.Object, m.Name);
 return nil
}

func(np *RstpSession) GetRootPathCost ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtRstpSession GetRootPathCost
 return "", nil
}

func(np *RstpSession) SetRootPathCost () error {
 //parameters: SessionPoolHandle RootPathCost
 //AgtRstpSession SetRootPathCost, m.Object, m.Name);
 return nil
}

func(np *RstpSession) GetTime ()(string, error) {
 //parameters: SessionPoolHandle TimeParameter
 //AgtRstpSession GetTime
 return "", nil
}

func(np *RstpSession) SetTime () error {
 //parameters: SessionPoolHandle TimeParameter Time
 //AgtRstpSession SetTime, m.Object, m.Name);
 return nil
}

func(np *RstpSession) GetOperationalModifier ()(string, error) {
 //parameters: SessionPoolHandle OperationalParameter
 //AgtRstpSession GetOperationalModifier
 return "", nil
}

func(np *RstpSession) SetOperationalModifier () error {
 //parameters: SessionPoolHandle OperationalParameter Value
 //AgtRstpSession SetOperationalModifier, m.Object, m.Name);
 return nil
}

func(np *RstpSession) SetIncrementingParameter () error {
 //parameters: SessionPoolHandle IncrementingParameter Value Repeat Increment
 //AgtRstpSession SetIncrementingParameter, m.Object, m.Name);
 return nil
}

func(np *RstpSession) GetIncrementingParameter ()(string, error) {
 //parameters: SessionPoolHandle IncrementingParameter
 //AgtRstpSession GetIncrementingParameter
 return "", nil
}

func(np *RstpSession) GetBridgeGroupAddressType ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtRstpSession GetBridgeGroupAddressType
 return "", nil
}

func(np *RstpSession) SetBridgeGroupAddressType () error {
 //parameters: SessionPoolHandle BridgeGroupAddressType
 //AgtRstpSession SetBridgeGroupAddressType, m.Object, m.Name);
 return nil
}

