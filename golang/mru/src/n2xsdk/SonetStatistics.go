package n2xsdk

type SonetStatistics struct {
 Handler string
}

func(np *SonetStatistics) GetType ()(string, error) {
 //parameters: Handle
 //AgtSonetStatistics GetType
 return "", nil
}

func(np *SonetStatistics) GetName ()(string, error) {
 //parameters: Handle
 //AgtSonetStatistics GetName
 return "", nil
}

func(np *SonetStatistics) GetLockCount ()(string, error) {
 //parameters: Handle
 //AgtSonetStatistics GetLockCount
 return "", nil
}

func(np *SonetStatistics) ListSelectedPorts ()(string, error) {
 //parameters: StatisticsHandle
 //AgtSonetStatistics ListSelectedPorts
 return "", nil
}

func(np *SonetStatistics) SelectPorts () error {
 //parameters: StatisticsHandle NumPorts PortHandles
 //AgtSonetStatistics SelectPorts, m.Object, m.Name);
 return nil
}

func(np *SonetStatistics) GetAccumulatedValues ()(string, error) {
 //parameters: StatisticsHandle
 //AgtSonetStatistics GetAccumulatedValues
 return "", nil
}

func(np *SonetStatistics) LockItem () error {
 //parameters: StatisticsHandle
 //AgtSonetStatistics LockItem, m.Object, m.Name);
 return nil
}

func(np *SonetStatistics) UnlockItem () error {
 //parameters: StatisticsHandle
 //AgtSonetStatistics UnlockItem, m.Object, m.Name);
 return nil
}

func(np *SonetStatistics) ListAvailableStatistics ()(string, error) {
 //parameters: 
 //AgtSonetStatistics ListAvailableStatistics
 return "", nil
}

func(np *SonetStatistics) ListSelectedStatistics ()(string, error) {
 //parameters: Handle
 //AgtSonetStatistics ListSelectedStatistics
 return "", nil
}

func(np *SonetStatistics) SelectStatistics () error {
 //parameters: Handle NumStatistics Statistics
 //AgtSonetStatistics SelectStatistics, m.Object, m.Name);
 return nil
}

