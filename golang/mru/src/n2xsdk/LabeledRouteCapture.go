package n2xsdk

type LabeledRouteCapture struct {
 Handler string
}

func(np *LabeledRouteCapture) Start () error {
 //parameters: SessionHandle Type
 //AgtLabeledRouteCapture Start, m.Object, m.Name);
 return nil
}

func(np *LabeledRouteCapture) Stop () error {
 //parameters: SessionHandle Type
 //AgtLabeledRouteCapture Stop, m.Object, m.Name);
 return nil
}

func(np *LabeledRouteCapture) Clear () error {
 //parameters: SessionHandle Type
 //AgtLabeledRouteCapture Clear, m.Object, m.Name);
 return nil
}

func(np *LabeledRouteCapture) IsCaptureEnabled () error {
 //parameters: SessionHandle Type
 //AgtLabeledRouteCapture IsCaptureEnabled, m.Object, m.Name);
 return nil
}

func(np *LabeledRouteCapture) GetNumCapturedLabeledRoutes ()(string, error) {
 //parameters: Count psaSessionId Type
 //AgtLabeledRouteCapture GetNumCapturedLabeledRoutes
 return "", nil
}

func(np *LabeledRouteCapture) GetNumReceivedLabeledRoutesAdvertisements ()(string, error) {
 //parameters: Count psaSessionId Type
 //AgtLabeledRouteCapture GetNumReceivedLabeledRoutesAdvertisements
 return "", nil
}

func(np *LabeledRouteCapture) GetNumReceivedLabeledRoutesWithdraws ()(string, error) {
 //parameters: Count psaSessionId Type
 //AgtLabeledRouteCapture GetNumReceivedLabeledRoutesWithdraws
 return "", nil
}

func(np *LabeledRouteCapture) SetCaptureMode () error {
 //parameters: SessionHandle Type Mode
 //AgtLabeledRouteCapture SetCaptureMode, m.Object, m.Name);
 return nil
}

func(np *LabeledRouteCapture) GetCaptureMode ()(string, error) {
 //parameters: SessionHandle Type
 //AgtLabeledRouteCapture GetCaptureMode
 return "", nil
}

func(np *LabeledRouteCapture) UploadAllLabeledRoutes () error {
 //parameters: SessionIdCount psaSessionId Type
 //AgtLabeledRouteCapture UploadAllLabeledRoutes, m.Object, m.Name);
 return nil
}

func(np *LabeledRouteCapture) UploadSelectedLabeledRoutes () error {
 //parameters: SessionIdCount psaSessionId Type FirstRoute PrefixLength NumRoutes Modifier
 //AgtLabeledRouteCapture UploadSelectedLabeledRoutes, m.Object, m.Name);
 return nil
}

