package n2xsdk

type HdlcStatus struct {
 Handler string
}

func(np *HdlcStatus) GetStatus ()(string, error) {
 //parameters: PortHandle
 //AgtHdlcStatus GetStatus
 return "", nil
}

func(np *HdlcStatus) GetStatusDescription ()(string, error) {
 //parameters: PortHandle
 //AgtHdlcStatus GetStatusDescription
 return "", nil
}

func(np *HdlcStatus) SaveStatus () error {
 //parameters: PortHandle LogFile
 //AgtHdlcStatus SaveStatus, m.Object, m.Name);
 return nil
}

func(np *HdlcStatus) GetStatusSummary ()(string, error) {
 //parameters: 
 //AgtHdlcStatus GetStatusSummary
 return "", nil
}

func(np *HdlcStatus) GetStatusSummaryDescription ()(string, error) {
 //parameters: 
 //AgtHdlcStatus GetStatusSummaryDescription
 return "", nil
}

func(np *HdlcStatus) SaveStatusSummary () error {
 //parameters: LogFile
 //AgtHdlcStatus SaveStatusSummary, m.Object, m.Name);
 return nil
}

func(np *HdlcStatus) GetStatusHistory ()(string, error) {
 //parameters: 
 //AgtHdlcStatus GetStatusHistory
 return "", nil
}

func(np *HdlcStatus) GetStatusHistoryDescription ()(string, error) {
 //parameters: 
 //AgtHdlcStatus GetStatusHistoryDescription
 return "", nil
}

func(np *HdlcStatus) SaveStatusHistory () error {
 //parameters: LogFile
 //AgtHdlcStatus SaveStatusHistory, m.Object, m.Name);
 return nil
}

func(np *HdlcStatus) ClearHistory () error {
 //parameters: 
 //AgtHdlcStatus ClearHistory, m.Object, m.Name);
 return nil
}

func(np *HdlcStatus) SetHistorySize () error {
 //parameters: HistorySize
 //AgtHdlcStatus SetHistorySize, m.Object, m.Name);
 return nil
}

func(np *HdlcStatus) GetHistorySize ()(string, error) {
 //parameters: 
 //AgtHdlcStatus GetHistorySize
 return "", nil
}

func(np *HdlcStatus) GetMaximumHistorySize ()(string, error) {
 //parameters: 
 //AgtHdlcStatus GetMaximumHistorySize
 return "", nil
}

