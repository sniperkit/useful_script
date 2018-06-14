package n2xsdk

type HardwareFailureStatus struct {
 Handler string
}

func(np *HardwareFailureStatus) GetStatus ()(string, error) {
 //parameters: PortHandle
 //AgtHardwareFailureStatus GetStatus
 return "", nil
}

func(np *HardwareFailureStatus) GetStatusDescription ()(string, error) {
 //parameters: PortHandle
 //AgtHardwareFailureStatus GetStatusDescription
 return "", nil
}

func(np *HardwareFailureStatus) SaveStatus () error {
 //parameters: PortHandle LogFile
 //AgtHardwareFailureStatus SaveStatus, m.Object, m.Name);
 return nil
}

func(np *HardwareFailureStatus) GetStatusSummary ()(string, error) {
 //parameters: 
 //AgtHardwareFailureStatus GetStatusSummary
 return "", nil
}

func(np *HardwareFailureStatus) GetStatusSummaryDescription ()(string, error) {
 //parameters: 
 //AgtHardwareFailureStatus GetStatusSummaryDescription
 return "", nil
}

func(np *HardwareFailureStatus) SaveStatusSummary () error {
 //parameters: LogFile
 //AgtHardwareFailureStatus SaveStatusSummary, m.Object, m.Name);
 return nil
}

func(np *HardwareFailureStatus) GetStatusHistory ()(string, error) {
 //parameters: 
 //AgtHardwareFailureStatus GetStatusHistory
 return "", nil
}

func(np *HardwareFailureStatus) GetStatusHistoryDescription ()(string, error) {
 //parameters: 
 //AgtHardwareFailureStatus GetStatusHistoryDescription
 return "", nil
}

func(np *HardwareFailureStatus) SaveStatusHistory () error {
 //parameters: LogFile
 //AgtHardwareFailureStatus SaveStatusHistory, m.Object, m.Name);
 return nil
}

func(np *HardwareFailureStatus) ClearHistory () error {
 //parameters: 
 //AgtHardwareFailureStatus ClearHistory, m.Object, m.Name);
 return nil
}

func(np *HardwareFailureStatus) SetHistorySize () error {
 //parameters: HistorySize
 //AgtHardwareFailureStatus SetHistorySize, m.Object, m.Name);
 return nil
}

func(np *HardwareFailureStatus) GetHistorySize ()(string, error) {
 //parameters: 
 //AgtHardwareFailureStatus GetHistorySize
 return "", nil
}

func(np *HardwareFailureStatus) GetMaximumHistorySize ()(string, error) {
 //parameters: 
 //AgtHardwareFailureStatus GetMaximumHistorySize
 return "", nil
}

