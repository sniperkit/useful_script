package n2xsdk

type HistogramClass struct {
 Handler string
}

func(np *HistogramClass) SetLabel () error {
 //parameters: ClassHandle Label
 //AgtHistogramClass SetLabel, m.Object, m.Name);
 return nil
}

func(np *HistogramClass) GetLabel ()(string, error) {
 //parameters: ClassHandle
 //AgtHistogramClass GetLabel
 return "", nil
}

func(np *HistogramClass) SetRange () error {
 //parameters: ClassHandle LowerBound UpperBound
 //AgtHistogramClass SetRange, m.Object, m.Name);
 return nil
}

func(np *HistogramClass) GetRange ()(string, error) {
 //parameters: ClassHandle
 //AgtHistogramClass GetRange
 return "", nil
}

