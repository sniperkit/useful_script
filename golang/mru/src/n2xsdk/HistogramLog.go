package n2xsdk

type HistogramLog struct {
 Handler string
}

func(np *HistogramLo) EnableLogging () error {
 //parameters: 
 //AgtHistogramLog EnableLogging, m.Object, m.Name);
 return nil
}

func(np *HistogramLo) DisableLogging () error {
 //parameters: 
 //AgtHistogramLog DisableLogging, m.Object, m.Name);
 return nil
}

func(np *HistogramLo) IsLoggingEnabled () error {
 //parameters: 
 //AgtHistogramLog IsLoggingEnabled, m.Object, m.Name);
 return nil
}

func(np *HistogramLo) SetLogFile () error {
 //parameters: LogFile
 //AgtHistogramLog SetLogFile, m.Object, m.Name);
 return nil
}

func(np *HistogramLo) GetLogFile ()(string, error) {
 //parameters: 
 //AgtHistogramLog GetLogFile
 return "", nil
}

func(np *HistogramLo) SetLoggingInterval () error {
 //parameters: Multiple
 //AgtHistogramLog SetLoggingInterval, m.Object, m.Name);
 return nil
}

func(np *HistogramLo) GetLoggingInterval ()(string, error) {
 //parameters: 
 //AgtHistogramLog GetLoggingInterval
 return "", nil
}

