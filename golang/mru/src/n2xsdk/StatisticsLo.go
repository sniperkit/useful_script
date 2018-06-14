package n2xsdk

type StatisticsLo struct {
 Handler string
}

func(np *StatisticsLo) EnableLogging () error {
 //parameters: 
 //AgtStatisticsLog EnableLogging, m.Object, m.Name);
 return nil
}

func(np *StatisticsLo) DisableLogging () error {
 //parameters: 
 //AgtStatisticsLog DisableLogging, m.Object, m.Name);
 return nil
}

func(np *StatisticsLo) IsLoggingEnabled () error {
 //parameters: 
 //AgtStatisticsLog IsLoggingEnabled, m.Object, m.Name);
 return nil
}

func(np *StatisticsLo) SetLogFile () error {
 //parameters: LogFile
 //AgtStatisticsLog SetLogFile, m.Object, m.Name);
 return nil
}

func(np *StatisticsLo) GetLogFile ()(string, error) {
 //parameters: 
 //AgtStatisticsLog GetLogFile
 return "", nil
}

func(np *StatisticsLo) SetLoggingInterval () error {
 //parameters: Multiple
 //AgtStatisticsLog SetLoggingInterval, m.Object, m.Name);
 return nil
}

func(np *StatisticsLo) GetLoggingInterval ()(string, error) {
 //parameters: 
 //AgtStatisticsLog GetLoggingInterval
 return "", nil
}

func(np *StatisticsLo) SelectPorts () error {
 //parameters: NumPorts psaPorts
 //AgtStatisticsLog SelectPorts, m.Object, m.Name);
 return nil
}

func(np *StatisticsLo) ListSelectedPorts ()(string, error) {
 //parameters: 
 //AgtStatisticsLog ListSelectedPorts
 return "", nil
}

func(np *StatisticsLo) ClearSelectedStreams () error {
 //parameters: 
 //AgtStatisticsLog ClearSelectedStreams, m.Object, m.Name);
 return nil
}

func(np *StatisticsLo) ListPortsWithStreamsSelected ()(string, error) {
 //parameters: 
 //AgtStatisticsLog ListPortsWithStreamsSelected
 return "", nil
}

func(np *StatisticsLo) SelectStreamGroups () error {
 //parameters: Count psaStreamGroupHandles
 //AgtStatisticsLog SelectStreamGroups, m.Object, m.Name);
 return nil
}

func(np *StatisticsLo) ListSelectedStreamGroups ()(string, error) {
 //parameters: 
 //AgtStatisticsLog ListSelectedStreamGroups
 return "", nil
}

func(np *StatisticsLo) SelectStreams () error {
 //parameters: NumStreamGroupHandles psaStreamGroupHandles NumStreamIndexes psaStreamIndexes NumDestinationPortHandles psaDestinationPorts
 //AgtStatisticsLog SelectStreams, m.Object, m.Name);
 return nil
}

func(np *StatisticsLo) ListSelectedStreams ()(string, error) {
 //parameters: 
 //AgtStatisticsLog ListSelectedStreams
 return "", nil
}

func(np *StatisticsLo) SelectFieldValues () error {
 //parameters: NumFieldInitialValues psaFieldInitialValues NumFieldValueCounts psaFieldValueCounts NumFieldValueIncrements psaFieldValueIncrements NumDestinationPortHandles psaDestinationPortHandles
 //AgtStatisticsLog SelectFieldValues, m.Object, m.Name);
 return nil
}

func(np *StatisticsLo) ListSelectedFieldValues ()(string, error) {
 //parameters: 
 //AgtStatisticsLog ListSelectedFieldValues
 return "", nil
}

func(np *StatisticsLo) ListAvailableStatistics ()(string, error) {
 //parameters: 
 //AgtStatisticsLog ListAvailableStatistics
 return "", nil
}

func(np *StatisticsLo) SelectStatistics () error {
 //parameters: NumStatistics psaStatistics
 //AgtStatisticsLog SelectStatistics, m.Object, m.Name);
 return nil
}

func(np *StatisticsLo) ListSelectedStatistics ()(string, error) {
 //parameters: 
 //AgtStatisticsLog ListSelectedStatistics
 return "", nil
}

