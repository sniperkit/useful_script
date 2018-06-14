package n2xsdk

type Threshold struct {
 Handler string
}

func(np *Threshold) SetLatencyThreshold () error {
 //parameters: PortHandle LatencyThreshold ThresholdType
 //AgtThreshold SetLatencyThreshold, m.Object, m.Name);
 return nil
}

func(np *Threshold) GetLatencyThreshold ()(string, error) {
 //parameters: PortHandle
 //AgtThreshold GetLatencyThreshold
 return "", nil
}

func(np *Threshold) SetRateThreshold () error {
 //parameters: PortHandle ThresholdType Source FrameMatcherHandle RateThreshold MaximumBurstSize
 //AgtThreshold SetRateThreshold, m.Object, m.Name);
 return nil
}

func(np *Threshold) GetRateThreshold ()(string, error) {
 //parameters: PortHandle
 //AgtThreshold GetRateThreshold
 return "", nil
}

func(np *Threshold) ClearRateThreshold () error {
 //parameters: PortHandle
 //AgtThreshold ClearRateThreshold, m.Object, m.Name);
 return nil
}

func(np *Threshold) SetCountThreshold () error {
 //parameters: PortHandle Source FrameMatcherHandle CountThreshold
 //AgtThreshold SetCountThreshold, m.Object, m.Name);
 return nil
}

func(np *Threshold) GetCountThreshold ()(string, error) {
 //parameters: PortHandle
 //AgtThreshold GetCountThreshold
 return "", nil
}

func(np *Threshold) ClearCountThreshold () error {
 //parameters: PortHandle
 //AgtThreshold ClearCountThreshold, m.Object, m.Name);
 return nil
}

