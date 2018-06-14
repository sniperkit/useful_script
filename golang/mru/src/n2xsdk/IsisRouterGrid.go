package n2xsdk

type IsisRouterGrid struct {
 Handler string
}

func(np *IsisRouterGrid) SetGridSize () error {
 //parameters: GridHandle nRows nColumns
 //AgtIsisRouterGrid SetGridSize, m.Object, m.Name);
 return nil
}

func(np *IsisRouterGrid) GetGridSize ()(string, error) {
 //parameters: GridHandle
 //AgtIsisRouterGrid GetGridSize
 return "", nil
}

func(np *IsisRouterGrid) SetStartingSystemId () error {
 //parameters: GridHandle StartingSystemId
 //AgtIsisRouterGrid SetStartingSystemId, m.Object, m.Name);
 return nil
}

func(np *IsisRouterGrid) GetStartingSystemId ()(string, error) {
 //parameters: GridHandle
 //AgtIsisRouterGrid GetStartingSystemId
 return "", nil
}

func(np *IsisRouterGrid) SetStartingRouterId () error {
 //parameters: GridHandle StartingRouterId
 //AgtIsisRouterGrid SetStartingRouterId, m.Object, m.Name);
 return nil
}

func(np *IsisRouterGrid) GetStartingRouterId ()(string, error) {
 //parameters: GridHandle
 //AgtIsisRouterGrid GetStartingRouterId
 return "", nil
}

func(np *IsisRouterGrid) SetGridConnection () error {
 //parameters: GridHandle SessionHandle Row Col
 //AgtIsisRouterGrid SetGridConnection, m.Object, m.Name);
 return nil
}

func(np *IsisRouterGrid) GetGridConnection ()(string, error) {
 //parameters: GridHandle SessionHandle
 //AgtIsisRouterGrid GetGridConnection
 return "", nil
}

func(np *IsisRouterGrid) DisconnectGrid () error {
 //parameters: GridHandle SessionHandle
 //AgtIsisRouterGrid DisconnectGrid, m.Object, m.Name);
 return nil
}

func(np *IsisRouterGrid) GetRouter ()(string, error) {
 //parameters: GridHandle nRow nCol
 //AgtIsisRouterGrid GetRouter
 return "", nil
}

func(np *IsisRouterGrid) EnableTe () error {
 //parameters: GridHandle
 //AgtIsisRouterGrid EnableTe, m.Object, m.Name);
 return nil
}

func(np *IsisRouterGrid) DisableTe () error {
 //parameters: GridHandle
 //AgtIsisRouterGrid DisableTe, m.Object, m.Name);
 return nil
}

func(np *IsisRouterGrid) IsTeEnabled () error {
 //parameters: GridHandle
 //AgtIsisRouterGrid IsTeEnabled, m.Object, m.Name);
 return nil
}

func(np *IsisRouterGrid) EnableMultiTopologies () error {
 //parameters: GridHandle
 //AgtIsisRouterGrid EnableMultiTopologies, m.Object, m.Name);
 return nil
}

func(np *IsisRouterGrid) DisableMultiTopologies () error {
 //parameters: GridHandle
 //AgtIsisRouterGrid DisableMultiTopologies, m.Object, m.Name);
 return nil
}

func(np *IsisRouterGrid) IsMultiTopologiesEnabled () error {
 //parameters: GridHandle
 //AgtIsisRouterGrid IsMultiTopologiesEnabled, m.Object, m.Name);
 return nil
}

func(np *IsisRouterGrid) AdvertiseRouter () error {
 //parameters: GridHandle nRow nCol
 //AgtIsisRouterGrid AdvertiseRouter, m.Object, m.Name);
 return nil
}

func(np *IsisRouterGrid) WithdrawRouter () error {
 //parameters: GridHandle nRow nCol
 //AgtIsisRouterGrid WithdrawRouter, m.Object, m.Name);
 return nil
}

func(np *IsisRouterGrid) Advertise () error {
 //parameters: GridHandle
 //AgtIsisRouterGrid Advertise, m.Object, m.Name);
 return nil
}

func(np *IsisRouterGrid) Withdraw () error {
 //parameters: GridHandle
 //AgtIsisRouterGrid Withdraw, m.Object, m.Name);
 return nil
}

func(np *IsisRouterGrid) GetRouterPosition ()(string, error) {
 //parameters: GridHandle RouterHandle
 //AgtIsisRouterGrid GetRouterPosition
 return "", nil
}

func(np *IsisRouterGrid) SetIpMode () error {
 //parameters: GridHandle IpMode
 //AgtIsisRouterGrid SetIpMode, m.Object, m.Name);
 return nil
}

func(np *IsisRouterGrid) GetIpMode ()(string, error) {
 //parameters: GridHandle
 //AgtIsisRouterGrid GetIpMode
 return "", nil
}

func(np *IsisRouterGrid) SetRoutingLevel () error {
 //parameters: GridHandle RoutingLevel
 //AgtIsisRouterGrid SetRoutingLevel, m.Object, m.Name);
 return nil
}

func(np *IsisRouterGrid) GetRoutingLevel ()(string, error) {
 //parameters: GridHandle
 //AgtIsisRouterGrid GetRoutingLevel
 return "", nil
}

