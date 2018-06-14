package n2xsdk

type IsisLsp struct {
 Handler string
}

func(np *IsisLsp) SetAreaId () error {
 //parameters: SessionHandle AreaId
 //AgtIsisLsp SetAreaId, m.Object, m.Name);
 return nil
}

func(np *IsisLsp) GetAreaId ()(string, error) {
 //parameters: SessionHandle
 //AgtIsisLsp GetAreaId
 return "", nil
}

func(np *IsisLsp) SetSystemId () error {
 //parameters: SessionHandle SystemId
 //AgtIsisLsp SetSystemId, m.Object, m.Name);
 return nil
}

func(np *IsisLsp) GetSystemId ()(string, error) {
 //parameters: SessionHandle
 //AgtIsisLsp GetSystemId
 return "", nil
}

func(np *IsisLsp) SetTeRouterId () error {
 //parameters: LspHandle TeRouterId
 //AgtIsisLsp SetTeRouterId, m.Object, m.Name);
 return nil
}

func(np *IsisLsp) GetTeRouterId ()(string, error) {
 //parameters: LspHandle
 //AgtIsisLsp GetTeRouterId
 return "", nil
}

