package n2xsdk

type GmplsLabelManager struct {
 Handler string
}

func(np *GmplsLabelManager) GetLabelSessionType ()(string, error) {
 //parameters: PortHandle
 //AgtGmplsLabelManager GetLabelSessionType
 return "", nil
}

func(np *GmplsLabelManager) SetSwitchingCapability () error {
 //parameters: PortHandle SwitchingCapability
 //AgtGmplsLabelManager SetSwitchingCapability, m.Object, m.Name);
 return nil
}

func(np *GmplsLabelManager) GetSwitchingCapability ()(string, error) {
 //parameters: PortHandle
 //AgtGmplsLabelManager GetSwitchingCapability
 return "", nil
}

func(np *GmplsLabelManager) SetInterfaceNumbering () error {
 //parameters: PortHandle InterfaceNumbering
 //AgtGmplsLabelManager SetInterfaceNumbering, m.Object, m.Name);
 return nil
}

func(np *GmplsLabelManager) GetInterfaceNumbering ()(string, error) {
 //parameters: PortHandle
 //AgtGmplsLabelManager GetInterfaceNumbering
 return "", nil
}

func(np *GmplsLabelManager) AddUnnumberedInterfaceGroup () error {
 //parameters: PortHandle InterfaceIdStart InterfaceIdIncr NumberOfInterfaces
 //AgtGmplsLabelManager AddUnnumberedInterfaceGroup, m.Object, m.Name);
 return nil
}

func(np *GmplsLabelManager) AddNumberedInterfaceGroup () error {
 //parameters: PortHandle InterfaceIdStart InterfaceIdIncr NumberOfInterfaces RemoteIdStart RemoteIdIncr
 //AgtGmplsLabelManager AddNumberedInterfaceGroup, m.Object, m.Name);
 return nil
}

func(np *GmplsLabelManager) SetUnnumberedInterfaceGroup () error {
 //parameters: PortHandle GroupHandle InterfaceIdStart InterfaceIdIncr NumberOfInterfaces
 //AgtGmplsLabelManager SetUnnumberedInterfaceGroup, m.Object, m.Name);
 return nil
}

func(np *GmplsLabelManager) SetNumberedInterfaceGroup () error {
 //parameters: PortHandle GroupHandle InterfaceIdStart InterfaceIdIncr NumberOfInterfaces RemoteIdStart RemoteIdIncr
 //AgtGmplsLabelManager SetNumberedInterfaceGroup, m.Object, m.Name);
 return nil
}

func(np *GmplsLabelManager) GetUnnumberedInterfaceGroup ()(string, error) {
 //parameters: PortHandle GroupHandle
 //AgtGmplsLabelManager GetUnnumberedInterfaceGroup
 return "", nil
}

func(np *GmplsLabelManager) GetNumberedInterfaceGroup ()(string, error) {
 //parameters: PortHandle GroupHandle
 //AgtGmplsLabelManager GetNumberedInterfaceGroup
 return "", nil
}

func(np *GmplsLabelManager) SetInterfaceGroupTeLink () error {
 //parameters: PortHandle GroupHandle UseTeLink TeLinkAddr
 //AgtGmplsLabelManager SetInterfaceGroupTeLink, m.Object, m.Name);
 return nil
}

func(np *GmplsLabelManager) GetInterfaceGroupTeLink ()(string, error) {
 //parameters: PortHandle GroupHandle
 //AgtGmplsLabelManager GetInterfaceGroupTeLink
 return "", nil
}

func(np *GmplsLabelManager) SetUnnumberedInterfaceGroupRemoteId () error {
 //parameters: PortHandle GroupHandle RemoteIdStart RemoteIdIncr
 //AgtGmplsLabelManager SetUnnumberedInterfaceGroupRemoteId, m.Object, m.Name);
 return nil
}

func(np *GmplsLabelManager) GetUnnumberedInterfaceGroupRemoteId ()(string, error) {
 //parameters: PortHandle GroupHandle
 //AgtGmplsLabelManager GetUnnumberedInterfaceGroupRemoteId
 return "", nil
}

func(np *GmplsLabelManager) EnableChannelization () error {
 //parameters: PortHandle GroupHandle
 //AgtGmplsLabelManager EnableChannelization, m.Object, m.Name);
 return nil
}

func(np *GmplsLabelManager) DisableChannelization () error {
 //parameters: PortHandle GroupHandle
 //AgtGmplsLabelManager DisableChannelization, m.Object, m.Name);
 return nil
}

func(np *GmplsLabelManager) IsChannelizationEnabled () error {
 //parameters: PortHandle GroupHandle
 //AgtGmplsLabelManager IsChannelizationEnabled, m.Object, m.Name);
 return nil
}

func(np *GmplsLabelManager) SetTdmInterfaceGroupRate () error {
 //parameters: PortHandle GroupHandle InterfaceRate
 //AgtGmplsLabelManager SetTdmInterfaceGroupRate, m.Object, m.Name);
 return nil
}

func(np *GmplsLabelManager) GetTdmInterfaceGroupRate ()(string, error) {
 //parameters: PortHandle GroupHandle
 //AgtGmplsLabelManager GetTdmInterfaceGroupRate
 return "", nil
}

func(np *GmplsLabelManager) SetLscInterfaceGroupRate () error {
 //parameters: PortHandle GroupHandle LambdaIdStart LambdaIdIncr NumberOfLambdas InterfaceRate
 //AgtGmplsLabelManager SetLscInterfaceGroupRate, m.Object, m.Name);
 return nil
}

func(np *GmplsLabelManager) GetLscInterfaceGroupRate ()(string, error) {
 //parameters: PortHandle GroupHandle
 //AgtGmplsLabelManager GetLscInterfaceGroupRate
 return "", nil
}

func(np *GmplsLabelManager) SetFscInterfaceGroupRate () error {
 //parameters: PortHandle GroupHandle FscInterfaceStart FscInterfaceIncr InterfaceRate
 //AgtGmplsLabelManager SetFscInterfaceGroupRate, m.Object, m.Name);
 return nil
}

func(np *GmplsLabelManager) GetFscInterfaceGroupRate ()(string, error) {
 //parameters: PortHandle GroupHandle
 //AgtGmplsLabelManager GetFscInterfaceGroupRate
 return "", nil
}

func(np *GmplsLabelManager) RemoveInterfaceGroup () error {
 //parameters: PortHandle GroupHandle
 //AgtGmplsLabelManager RemoveInterfaceGroup, m.Object, m.Name);
 return nil
}

func(np *GmplsLabelManager) ListInterfaceGroups ()(string, error) {
 //parameters: PortHandle
 //AgtGmplsLabelManager ListInterfaceGroups
 return "", nil
}

