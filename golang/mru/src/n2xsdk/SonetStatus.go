package n2xsdk

type SonetStatus struct {
 Handler string
}

func(np *SonetStatus) GetStatus ()(string, error) {
 //parameters: PortHandle
 //AgtSonetStatus GetStatus
 return "", nil
}

func(np *SonetStatus) GetStatusDescription ()(string, error) {
 //parameters: PortHandle
 //AgtSonetStatus GetStatusDescription
 return "", nil
}

func(np *SonetStatus) SaveStatus () error {
 //parameters: PortHandle LogFile
 //AgtSonetStatus SaveStatus, m.Object, m.Name);
 return nil
}

func(np *SonetStatus) GetStatusSummary ()(string, error) {
 //parameters: 
 //AgtSonetStatus GetStatusSummary
 return "", nil
}

func(np *SonetStatus) GetStatusSummaryDescription ()(string, error) {
 //parameters: 
 //AgtSonetStatus GetStatusSummaryDescription
 return "", nil
}

func(np *SonetStatus) SaveStatusSummary () error {
 //parameters: LogFile
 //AgtSonetStatus SaveStatusSummary, m.Object, m.Name);
 return nil
}

func(np *SonetStatus) GetStatusHistory ()(string, error) {
 //parameters: 
 //AgtSonetStatus GetStatusHistory
 return "", nil
}

func(np *SonetStatus) GetStatusHistoryDescription ()(string, error) {
 //parameters: 
 //AgtSonetStatus GetStatusHistoryDescription
 return "", nil
}

func(np *SonetStatus) SaveStatusHistory () error {
 //parameters: LogFile
 //AgtSonetStatus SaveStatusHistory, m.Object, m.Name);
 return nil
}

func(np *SonetStatus) ClearHistory () error {
 //parameters: 
 //AgtSonetStatus ClearHistory, m.Object, m.Name);
 return nil
}

func(np *SonetStatus) SetHistorySize () error {
 //parameters: HistorySize
 //AgtSonetStatus SetHistorySize, m.Object, m.Name);
 return nil
}

func(np *SonetStatus) GetHistorySize ()(string, error) {
 //parameters: 
 //AgtSonetStatus GetHistorySize
 return "", nil
}

func(np *SonetStatus) GetMaximumHistorySize ()(string, error) {
 //parameters: 
 //AgtSonetStatus GetMaximumHistorySize
 return "", nil
}

