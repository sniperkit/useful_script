package n2xsdk

type AtmStatus struct {
 Handler string
}

func(np *mStatus) GetStatus ()(string, error) {
 //parameters: PortHandle
 //AgtAtmStatus GetStatus
 return "", nil
}

func(np *mStatus) GetStatusDescription ()(string, error) {
 //parameters: PortHandle
 //AgtAtmStatus GetStatusDescription
 return "", nil
}

func(np *mStatus) SaveStatus () error {
 //parameters: PortHandle LogFile
 //AgtAtmStatus SaveStatus, m.Object, m.Name);
 return nil
}

func(np *mStatus) GetStatusSummary ()(string, error) {
 //parameters: 
 //AgtAtmStatus GetStatusSummary
 return "", nil
}

func(np *mStatus) GetStatusSummaryDescription ()(string, error) {
 //parameters: 
 //AgtAtmStatus GetStatusSummaryDescription
 return "", nil
}

func(np *mStatus) SaveStatusSummary () error {
 //parameters: LogFile
 //AgtAtmStatus SaveStatusSummary, m.Object, m.Name);
 return nil
}

func(np *mStatus) GetStatusHistory ()(string, error) {
 //parameters: 
 //AgtAtmStatus GetStatusHistory
 return "", nil
}

func(np *mStatus) GetStatusHistoryDescription ()(string, error) {
 //parameters: 
 //AgtAtmStatus GetStatusHistoryDescription
 return "", nil
}

func(np *mStatus) SaveStatusHistory () error {
 //parameters: LogFile
 //AgtAtmStatus SaveStatusHistory, m.Object, m.Name);
 return nil
}

func(np *mStatus) ClearHistory () error {
 //parameters: 
 //AgtAtmStatus ClearHistory, m.Object, m.Name);
 return nil
}

func(np *mStatus) SetHistorySize () error {
 //parameters: HistorySize
 //AgtAtmStatus SetHistorySize, m.Object, m.Name);
 return nil
}

func(np *mStatus) GetHistorySize ()(string, error) {
 //parameters: 
 //AgtAtmStatus GetHistorySize
 return "", nil
}

func(np *mStatus) GetMaximumHistorySize ()(string, error) {
 //parameters: 
 //AgtAtmStatus GetMaximumHistorySize
 return "", nil
}

