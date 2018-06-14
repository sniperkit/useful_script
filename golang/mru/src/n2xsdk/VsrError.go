package n2xsdk

type VsrError struct {
 Handler string
}

func(np *VsrError) VsrErrorOn () error {
 //parameters: PortHandle VsrError
 //AgtVsrError VsrErrorOn, m.Object, m.Name);
 return nil
}

func(np *VsrError) VsrErrorOff () error {
 //parameters: PortHandle VsrError
 //AgtVsrError VsrErrorOff, m.Object, m.Name);
 return nil
}

func(np *VsrError) IsVsrErrorOn () error {
 //parameters: PortHandle VsrError
 //AgtVsrError IsVsrErrorOn, m.Object, m.Name);
 return nil
}

