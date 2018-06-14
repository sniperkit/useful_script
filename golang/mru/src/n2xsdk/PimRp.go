package n2xsdk

type PimRp struct {
 Handler string
}

func(np *PimRp) SetRpEntityType () error {
 //parameters: SessionHandle RpEntityType
 //AgtPimRp SetRpEntityType, m.Object, m.Name);
 return nil
}

func(np *PimRp) GetRpEntityType ()(string, error) {
 //parameters: SessionHandle
 //AgtPimRp GetRpEntityType
 return "", nil
}

func(np *PimRp) GetRemoteRpAddress ()(string, error) {
 //parameters: SessionHandle
 //AgtPimRp GetRemoteRpAddress
 return "", nil
}

func(np *PimRp) GetCRpDetails ()(string, error) {
 //parameters: SessionHandle
 //AgtPimRp GetCRpDetails
 return "", nil
}

func(np *PimRp) ListCRpGroupRanges ()(string, error) {
 //parameters: SessionHandle
 //AgtPimRp ListCRpGroupRanges
 return "", nil
}

func(np *PimRp) GetCRpGroupRange ()(string, error) {
 //parameters: SessionHandle GroupRangeHandle
 //AgtPimRp GetCRpGroupRange
 return "", nil
}

func(np *PimRp) GetCRpGroupAdminScope ()(string, error) {
 //parameters: SessionHandle GroupRangeHandle
 //AgtPimRp GetCRpGroupAdminScope
 return "", nil
}

func(np *PimRp) GetCRpGroupBiDirBit ()(string, error) {
 //parameters: SessionHandle GroupRangeHandle
 //AgtPimRp GetCRpGroupBiDirBit
 return "", nil
}

func(np *PimRp) GetCRpFsmState ()(string, error) {
 //parameters: SessionHandle
 //AgtPimRp GetCRpFsmState
 return "", nil
}

func(np *PimRp) GetCBsrDetails ()(string, error) {
 //parameters: SessionHandle
 //AgtPimRp GetCBsrDetails
 return "", nil
}

func(np *PimRp) ListRpSets ()(string, error) {
 //parameters: SessionHandle
 //AgtPimRp ListRpSets
 return "", nil
}

func(np *PimRp) GetRpSetGroupRange ()(string, error) {
 //parameters: SessionHandle RpSetHandle
 //AgtPimRp GetRpSetGroupRange
 return "", nil
}

func(np *PimRp) GetRpSetGroupAdminScope ()(string, error) {
 //parameters: SessionHandle RpSetHandle
 //AgtPimRp GetRpSetGroupAdminScope
 return "", nil
}

func(np *PimRp) GetRpSetGroupBiDirBit ()(string, error) {
 //parameters: SessionHandle RpSetHandle
 //AgtPimRp GetRpSetGroupBiDirBit
 return "", nil
}

func(np *PimRp) ListRpSetRps ()(string, error) {
 //parameters: SessionHandle RpSetHandle
 //AgtPimRp ListRpSetRps
 return "", nil
}

func(np *PimRp) GetRpSetRpDetails ()(string, error) {
 //parameters: SessionHandle RpSetHandle RpHandle
 //AgtPimRp GetRpSetRpDetails
 return "", nil
}

func(np *PimRp) StopCBsrMessages () error {
 //parameters: SessionHandle
 //AgtPimRp StopCBsrMessages, m.Object, m.Name);
 return nil
}

func(np *PimRp) StartCBsrMessages () error {
 //parameters: SessionHandle
 //AgtPimRp StartCBsrMessages, m.Object, m.Name);
 return nil
}

func(np *PimRp) GetCBsrFsmState ()(string, error) {
 //parameters: SessionHandle
 //AgtPimRp GetCBsrFsmState
 return "", nil
}

func(np *PimRp) SetRemoteRpAddress () error {
 //parameters: 
 //AgtPimRp SetRemoteRpAddress, m.Object, m.Name);
 return nil
}

func(np *PimRp) SetCRpDetails () error {
 //parameters: 
 //AgtPimRp SetCRpDetails, m.Object, m.Name);
 return nil
}

func(np *PimRp) AddCRpGroupRanges () error {
 //parameters: 
 //AgtPimRp AddCRpGroupRanges, m.Object, m.Name);
 return nil
}

func(np *PimRp) RemoveCRpGroupRanges () error {
 //parameters: 
 //AgtPimRp RemoveCRpGroupRanges, m.Object, m.Name);
 return nil
}

func(np *PimRp) RemoveAllCRpGroupRanges () error {
 //parameters: 
 //AgtPimRp RemoveAllCRpGroupRanges, m.Object, m.Name);
 return nil
}

func(np *PimRp) SetCRpGroupAdminScope () error {
 //parameters: 
 //AgtPimRp SetCRpGroupAdminScope, m.Object, m.Name);
 return nil
}

func(np *PimRp) SetCRpGroupBiDirBit () error {
 //parameters: 
 //AgtPimRp SetCRpGroupBiDirBit, m.Object, m.Name);
 return nil
}

func(np *PimRp) StopCRpAdvertisements () error {
 //parameters: 
 //AgtPimRp StopCRpAdvertisements, m.Object, m.Name);
 return nil
}

func(np *PimRp) StartCRpAdvertisements () error {
 //parameters: 
 //AgtPimRp StartCRpAdvertisements, m.Object, m.Name);
 return nil
}

func(np *PimRp) SetCBsrDetails () error {
 //parameters: 
 //AgtPimRp SetCBsrDetails, m.Object, m.Name);
 return nil
}

func(np *PimRp) AddRpSet () error {
 //parameters: 
 //AgtPimRp AddRpSet, m.Object, m.Name);
 return nil
}

func(np *PimRp) RemoveRpSets () error {
 //parameters: 
 //AgtPimRp RemoveRpSets, m.Object, m.Name);
 return nil
}

func(np *PimRp) RemoveAllRpSets () error {
 //parameters: 
 //AgtPimRp RemoveAllRpSets, m.Object, m.Name);
 return nil
}

func(np *PimRp) SetRpSetGroupRange () error {
 //parameters: 
 //AgtPimRp SetRpSetGroupRange, m.Object, m.Name);
 return nil
}

func(np *PimRp) SetRpSetGroupAdminScope () error {
 //parameters: 
 //AgtPimRp SetRpSetGroupAdminScope, m.Object, m.Name);
 return nil
}

func(np *PimRp) SetRpSetGroupBiDirBit () error {
 //parameters: 
 //AgtPimRp SetRpSetGroupBiDirBit, m.Object, m.Name);
 return nil
}

func(np *PimRp) AddRpSetRps () error {
 //parameters: 
 //AgtPimRp AddRpSetRps, m.Object, m.Name);
 return nil
}

func(np *PimRp) RemoveRpSetRps () error {
 //parameters: 
 //AgtPimRp RemoveRpSetRps, m.Object, m.Name);
 return nil
}

func(np *PimRp) RemoveAllRpSetRps () error {
 //parameters: 
 //AgtPimRp RemoveAllRpSetRps, m.Object, m.Name);
 return nil
}

