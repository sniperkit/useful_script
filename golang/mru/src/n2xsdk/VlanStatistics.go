package n2xsdk

type VlanStatistics struct {
 Handler string
}

func(np *VlanStatistics) GetType ()(string, error) {
 //parameters: Handle
 //AgtVlanStatistics GetType
 return "", nil
}

func(np *VlanStatistics) GetName ()(string, error) {
 //parameters: Handle
 //AgtVlanStatistics GetName
 return "", nil
}

func(np *VlanStatistics) GetLockCount ()(string, error) {
 //parameters: Handle
 //AgtVlanStatistics GetLockCount
 return "", nil
}

func(np *VlanStatistics) ListSelectedPorts ()(string, error) {
 //parameters: StatisticsHandle
 //AgtVlanStatistics ListSelectedPorts
 return "", nil
}

func(np *VlanStatistics) SelectPorts () error {
 //parameters: StatisticsHandle NumPorts PortHandles
 //AgtVlanStatistics SelectPorts, m.Object, m.Name);
 return nil
}

func(np *VlanStatistics) GetAccumulatedValues ()(string, error) {
 //parameters: StatisticsHandle
 //AgtVlanStatistics GetAccumulatedValues
 return "", nil
}

func(np *VlanStatistics) LockItem () error {
 //parameters: StatisticsHandle
 //AgtVlanStatistics LockItem, m.Object, m.Name);
 return nil
}

func(np *VlanStatistics) UnlockItem () error {
 //parameters: StatisticsHandle
 //AgtVlanStatistics UnlockItem, m.Object, m.Name);
 return nil
}

func(np *VlanStatistics) ListAvailableStatistics ()(string, error) {
 //parameters: 
 //AgtVlanStatistics ListAvailableStatistics
 return "", nil
}

func(np *VlanStatistics) ListSelectedStatistics ()(string, error) {
 //parameters: Handle
 //AgtVlanStatistics ListSelectedStatistics
 return "", nil
}

func(np *VlanStatistics) SelectStatistics () error {
 //parameters: Handle NumStatistics Statistics
 //AgtVlanStatistics SelectStatistics, m.Object, m.Name);
 return nil
}

