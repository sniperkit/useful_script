package n2xsdk

type OspfRouterGrid struct {
 Handler string
}

func(np *OspfRouterGrid) SetGridSize () error {
 //parameters: GridHandle nRows nColumns
 //AgtOspfRouterGrid SetGridSize, m.Object, m.Name);
 return nil
}

func(np *OspfRouterGrid) GetGridSize ()(string, error) {
 //parameters: GridHandle
 //AgtOspfRouterGrid GetGridSize
 return "", nil
}

func(np *OspfRouterGrid) SetStartingRouterId () error {
 //parameters: GridHandle StartingRouterId
 //AgtOspfRouterGrid SetStartingRouterId, m.Object, m.Name);
 return nil
}

func(np *OspfRouterGrid) GetStartingRouterId ()(string, error) {
 //parameters: GridHandle
 //AgtOspfRouterGrid GetStartingRouterId
 return "", nil
}

func(np *OspfRouterGrid) SetStartingTeInterfaceAddress () error {
 //parameters: GridHandle StartingTeInterfaceAddress
 //AgtOspfRouterGrid SetStartingTeInterfaceAddress, m.Object, m.Name);
 return nil
}

func(np *OspfRouterGrid) GetStartingTeInterfaceAddress ()(string, error) {
 //parameters: GridHandle
 //AgtOspfRouterGrid GetStartingTeInterfaceAddress
 return "", nil
}

func(np *OspfRouterGrid) SetStartingGmplsLinkId () error {
 //parameters: GridHandle StartingGmplsLinkId
 //AgtOspfRouterGrid SetStartingGmplsLinkId, m.Object, m.Name);
 return nil
}

func(np *OspfRouterGrid) GetStartingGmplsLinkId ()(string, error) {
 //parameters: GridHandle
 //AgtOspfRouterGrid GetStartingGmplsLinkId
 return "", nil
}

func(np *OspfRouterGrid) SetGridConnection () error {
 //parameters: GridHandle SessionHandle RowNum ColNum
 //AgtOspfRouterGrid SetGridConnection, m.Object, m.Name);
 return nil
}

func(np *OspfRouterGrid) GetGridConnection ()(string, error) {
 //parameters: GridHandle OspfSessionHandle
 //AgtOspfRouterGrid GetGridConnection
 return "", nil
}

func(np *OspfRouterGrid) SetGridDisconnection () error {
 //parameters: GridHandle SessionHandle
 //AgtOspfRouterGrid SetGridDisconnection, m.Object, m.Name);
 return nil
}

func(np *OspfRouterGrid) GetRouter ()(string, error) {
 //parameters: GridHandle RowNum ColNum
 //AgtOspfRouterGrid GetRouter
 return "", nil
}

func(np *OspfRouterGrid) EnableTe () error {
 //parameters: GridHandle
 //AgtOspfRouterGrid EnableTe, m.Object, m.Name);
 return nil
}

func(np *OspfRouterGrid) DisableTe () error {
 //parameters: GridHandle
 //AgtOspfRouterGrid DisableTe, m.Object, m.Name);
 return nil
}

func(np *OspfRouterGrid) IsTeEnabled () error {
 //parameters: GridHandle
 //AgtOspfRouterGrid IsTeEnabled, m.Object, m.Name);
 return nil
}

func(np *OspfRouterGrid) AdvertiseRouter () error {
 //parameters: GridHandle RowNum ColNum
 //AgtOspfRouterGrid AdvertiseRouter, m.Object, m.Name);
 return nil
}

func(np *OspfRouterGrid) WithdrawRouter () error {
 //parameters: GridHandle RowNum ColNum
 //AgtOspfRouterGrid WithdrawRouter, m.Object, m.Name);
 return nil
}

func(np *OspfRouterGrid) SetGridLinkType () error {
 //parameters: GridHandle GridLinkType
 //AgtOspfRouterGrid SetGridLinkType, m.Object, m.Name);
 return nil
}

func(np *OspfRouterGrid) GetGridLinkType ()(string, error) {
 //parameters: GridHandle
 //AgtOspfRouterGrid GetGridLinkType
 return "", nil
}

func(np *OspfRouterGrid) GetRouterPosition ()(string, error) {
 //parameters: GridHandle RouterHandle
 //AgtOspfRouterGrid GetRouterPosition
 return "", nil
}

