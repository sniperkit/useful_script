package n2xsdk

type PimGroupFlap struct {
 Handler string
}

func(np *PimGroupFlap) SetFlapTimers () error {
 //parameters: SessionHandle MembershipPoolHandle JoinToLeaveDelay LeaveToJoinDelay
 //AgtPimGroupFlap SetFlapTimers, m.Object, m.Name);
 return nil
}

func(np *PimGroupFlap) GetFlapTimers ()(string, error) {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtPimGroupFlap GetFlapTimers
 return "", nil
}

func(np *PimGroupFlap) StartFlap () error {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtPimGroupFlap StartFlap, m.Object, m.Name);
 return nil
}

func(np *PimGroupFlap) StopFlap () error {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtPimGroupFlap StopFlap, m.Object, m.Name);
 return nil
}

func(np *PimGroupFlap) StartAllFlapsInSession () error {
 //parameters: SessionHandle
 //AgtPimGroupFlap StartAllFlapsInSession, m.Object, m.Name);
 return nil
}

func(np *PimGroupFlap) StopAllFlapsInSession () error {
 //parameters: SessionHandle
 //AgtPimGroupFlap StopAllFlapsInSession, m.Object, m.Name);
 return nil
}

func(np *PimGroupFlap) StartAllFlaps () error {
 //parameters: 
 //AgtPimGroupFlap StartAllFlaps, m.Object, m.Name);
 return nil
}

func(np *PimGroupFlap) StopAllFlaps () error {
 //parameters: 
 //AgtPimGroupFlap StopAllFlaps, m.Object, m.Name);
 return nil
}

func(np *PimGroupFlap) Cancel () error {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtPimGroupFlap Cancel, m.Object, m.Name);
 return nil
}

func(np *PimGroupFlap) CancelAllInSession () error {
 //parameters: SessionHandle
 //AgtPimGroupFlap CancelAllInSession, m.Object, m.Name);
 return nil
}

func(np *PimGroupFlap) CancelAll () error {
 //parameters: 
 //AgtPimGroupFlap CancelAll, m.Object, m.Name);
 return nil
}

func(np *PimGroupFlap) IsPoolAvailableForAction () error {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtPimGroupFlap IsPoolAvailableForAction, m.Object, m.Name);
 return nil
}

func(np *PimGroupFlap) IsPoolAvailableForCancel () error {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtPimGroupFlap IsPoolAvailableForCancel, m.Object, m.Name);
 return nil
}

func(np *PimGroupFlap) EnableFlapFlag () error {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtPimGroupFlap EnableFlapFlag, m.Object, m.Name);
 return nil
}

func(np *PimGroupFlap) DisableFlapFlag () error {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtPimGroupFlap DisableFlapFlag, m.Object, m.Name);
 return nil
}

func(np *PimGroupFlap) GetFlapEnableFlag ()(string, error) {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtPimGroupFlap GetFlapEnableFlag
 return "", nil
}

