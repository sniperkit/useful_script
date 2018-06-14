package n2xsdk

type TestController struct {
 Handler string
}

func(np *TestController) SetTestMode () error {
 //parameters: TestMode
 //AgtTestController SetTestMode, m.Object, m.Name);
 return nil
}

func(np *TestController) GetTestMode ()(string, error) {
 //parameters: 
 //AgtTestController GetTestMode
 return "", nil
}

func(np *TestController) SetTestDuration () error {
 //parameters: Duration
 //AgtTestController SetTestDuration, m.Object, m.Name);
 return nil
}

func(np *TestController) GetTestDuration ()(string, error) {
 //parameters: 
 //AgtTestController GetTestDuration
 return "", nil
}

func(np *TestController) GetTestDurationLimits ()(string, error) {
 //parameters: 
 //AgtTestController GetTestDurationLimits
 return "", nil
}

func(np *TestController) SetSamplingInterval () error {
 //parameters: SamplingInterval
 //AgtTestController SetSamplingInterval, m.Object, m.Name);
 return nil
}

func(np *TestController) GetSamplingInterval ()(string, error) {
 //parameters: 
 //AgtTestController GetSamplingInterval
 return "", nil
}

func(np *TestController) GetSamplingIntervalLimits ()(string, error) {
 //parameters: 
 //AgtTestController GetSamplingIntervalLimits
 return "", nil
}

func(np *TestController) SetTrickleTime () error {
 //parameters: TrickleTime
 //AgtTestController SetTrickleTime, m.Object, m.Name);
 return nil
}

func(np *TestController) GetTrickleTime ()(string, error) {
 //parameters: 
 //AgtTestController GetTrickleTime
 return "", nil
}

func(np *TestController) GetTrickleTimeLimits ()(string, error) {
 //parameters: 
 //AgtTestController GetTrickleTimeLimits
 return "", nil
}

func(np *TestController) EnableTrickleTime () error {
 //parameters: 
 //AgtTestController EnableTrickleTime, m.Object, m.Name);
 return nil
}

func(np *TestController) DisableTrickleTime () error {
 //parameters: 
 //AgtTestController DisableTrickleTime, m.Object, m.Name);
 return nil
}

func(np *TestController) IsTrickleTimeEnabled () error {
 //parameters: 
 //AgtTestController IsTrickleTimeEnabled, m.Object, m.Name);
 return nil
}

func(np *TestController) StartTest () error {
 //parameters: 
 //AgtTestController StartTest, m.Object, m.Name);
 return nil
}

func(np *TestController) StopTest () error {
 //parameters: 
 //AgtTestController StopTest, m.Object, m.Name);
 return nil
}

func(np *TestController) WaitUntilTestStops () error {
 //parameters: 
 //AgtTestController WaitUntilTestStops, m.Object, m.Name);
 return nil
}

func(np *TestController) GetTestState ()(string, error) {
 //parameters: 
 //AgtTestController GetTestState
 return "", nil
}

func(np *TestController) GetStartTime ()(string, error) {
 //parameters: 
 //AgtTestController GetStartTime
 return "", nil
}

func(np *TestController) GetElapsedTime ()(string, error) {
 //parameters: 
 //AgtTestController GetElapsedTime
 return "", nil
}

