package n2xsdk

type PimGroup struct {
 Handler string
}

func(np *PimGroup) GetState ()(string, error) {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtPimGroup GetState
 return "", nil
}

func(np *PimGroup) EnableSourceSpecificJoinPrune () error {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtPimGroup EnableSourceSpecificJoinPrune, m.Object, m.Name);
 return nil
}

func(np *PimGroup) DisableSourceSpecificJoinPrune () error {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtPimGroup DisableSourceSpecificJoinPrune, m.Object, m.Name);
 return nil
}

func(np *PimGroup) EnableAllSourceSpecificJoinPruneInSession () error {
 //parameters: SessionHandle
 //AgtPimGroup EnableAllSourceSpecificJoinPruneInSession, m.Object, m.Name);
 return nil
}

func(np *PimGroup) DisableAllSourceSpecificJoinPruneInSession () error {
 //parameters: SessionHandle
 //AgtPimGroup DisableAllSourceSpecificJoinPruneInSession, m.Object, m.Name);
 return nil
}

func(np *PimGroup) EnableAllSourceSpecificJoinPrune () error {
 //parameters: 
 //AgtPimGroup EnableAllSourceSpecificJoinPrune, m.Object, m.Name);
 return nil
}

func(np *PimGroup) DisableAllSourceSpecificJoinPrune () error {
 //parameters: 
 //AgtPimGroup DisableAllSourceSpecificJoinPrune, m.Object, m.Name);
 return nil
}

func(np *PimGroup) IsSourceSpecificJoinPruneEnabled () error {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtPimGroup IsSourceSpecificJoinPruneEnabled, m.Object, m.Name);
 return nil
}

func(np *PimGroup) EnableWildcardGroup () error {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtPimGroup EnableWildcardGroup, m.Object, m.Name);
 return nil
}

func(np *PimGroup) DisableWildcardGroup () error {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtPimGroup DisableWildcardGroup, m.Object, m.Name);
 return nil
}

func(np *PimGroup) IsWildcardGroupEnabled () error {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtPimGroup IsWildcardGroupEnabled, m.Object, m.Name);
 return nil
}

func(np *PimGroup) EnableSGRpt () error {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtPimGroup EnableSGRpt, m.Object, m.Name);
 return nil
}

func(np *PimGroup) DisableSGRpt () error {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtPimGroup DisableSGRpt, m.Object, m.Name);
 return nil
}

func(np *PimGroup) IsSGRptEnabled () error {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtPimGroup IsSGRptEnabled, m.Object, m.Name);
 return nil
}

func(np *PimGroup) EnableJoinPrunePduAggregation () error {
 //parameters: SessionHandle MembershipPoolHandle AggrFactor
 //AgtPimGroup EnableJoinPrunePduAggregation, m.Object, m.Name);
 return nil
}

func(np *PimGroup) DisableJoinPrunePduAggregation () error {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtPimGroup DisableJoinPrunePduAggregation, m.Object, m.Name);
 return nil
}

func(np *PimGroup) EnableAllJoinPrunePduAggregationInSession () error {
 //parameters: SessionHandle AggrFactor
 //AgtPimGroup EnableAllJoinPrunePduAggregationInSession, m.Object, m.Name);
 return nil
}

func(np *PimGroup) DisableAllJoinPrunePduAggregationInSession () error {
 //parameters: SessionHandle
 //AgtPimGroup DisableAllJoinPrunePduAggregationInSession, m.Object, m.Name);
 return nil
}

func(np *PimGroup) EnableAllJoinPrunePduAggregation () error {
 //parameters: AggrFactor
 //AgtPimGroup EnableAllJoinPrunePduAggregation, m.Object, m.Name);
 return nil
}

func(np *PimGroup) DisableAllJoinPrunePduAggregation () error {
 //parameters: 
 //AgtPimGroup DisableAllJoinPrunePduAggregation, m.Object, m.Name);
 return nil
}

func(np *PimGroup) IsJoinPrunePduAggregationEnabled () error {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtPimGroup IsJoinPrunePduAggregationEnabled, m.Object, m.Name);
 return nil
}

func(np *PimGroup) GetJoinPrunePduAggregationFactor ()(string, error) {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtPimGroup GetJoinPrunePduAggregationFactor
 return "", nil
}

func(np *PimGroup) SetJoinPruneTransmissionRate () error {
 //parameters: SessionHandle MembershipPoolHandle NumPdusPerInterval Interval
 //AgtPimGroup SetJoinPruneTransmissionRate, m.Object, m.Name);
 return nil
}

func(np *PimGroup) GetJoinPruneTransmissionRate ()(string, error) {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtPimGroup GetJoinPruneTransmissionRate
 return "", nil
}

func(np *PimGroup) Join () error {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtPimGroup Join, m.Object, m.Name);
 return nil
}

func(np *PimGroup) Prune () error {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtPimGroup Prune, m.Object, m.Name);
 return nil
}

func(np *PimGroup) JoinAllGroupsInSession () error {
 //parameters: SessionHandle
 //AgtPimGroup JoinAllGroupsInSession, m.Object, m.Name);
 return nil
}

func(np *PimGroup) PruneAllGroupsInSession () error {
 //parameters: SessionHandle
 //AgtPimGroup PruneAllGroupsInSession, m.Object, m.Name);
 return nil
}

func(np *PimGroup) JoinAllGroups () error {
 //parameters: 
 //AgtPimGroup JoinAllGroups, m.Object, m.Name);
 return nil
}

func(np *PimGroup) PruneAllGroups () error {
 //parameters: 
 //AgtPimGroup PruneAllGroups, m.Object, m.Name);
 return nil
}

func(np *PimGroup) Cancel () error {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtPimGroup Cancel, m.Object, m.Name);
 return nil
}

func(np *PimGroup) CancelAllInSession () error {
 //parameters: SessionHandle
 //AgtPimGroup CancelAllInSession, m.Object, m.Name);
 return nil
}

func(np *PimGroup) CancelAll () error {
 //parameters: 
 //AgtPimGroup CancelAll, m.Object, m.Name);
 return nil
}

func(np *PimGroup) IsPoolAvailableForAction () error {
 //parameters: SessionHandle MembershipPoolHandle Action SourceSpecific
 //AgtPimGroup IsPoolAvailableForAction, m.Object, m.Name);
 return nil
}

func(np *PimGroup) IsPoolAvailableForCancel () error {
 //parameters: SessionHandle MembershipPoolHandle SourceSpecific
 //AgtPimGroup IsPoolAvailableForCancel, m.Object, m.Name);
 return nil
}

func(np *PimGroup) GetReceiverDrGroupState ()(string, error) {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtPimGroup GetReceiverDrGroupState
 return "", nil
}

func(np *PimGroup) GetSourceDrGroupState ()(string, error) {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtPimGroup GetSourceDrGroupState
 return "", nil
}

func(np *PimGroup) SetGroupActionAtSessionEnable () error {
 //parameters: SessionHandle MembershipPoolHandle GroupAction
 //AgtPimGroup SetGroupActionAtSessionEnable, m.Object, m.Name);
 return nil
}

func(np *PimGroup) GetGroupActionAtSessionEnable ()(string, error) {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtPimGroup GetGroupActionAtSessionEnable
 return "", nil
}

func(np *PimGroup) SetSGJoinMode () error {
 //parameters: SessionHandle MembershipPoolHandle SGJoinMode
 //AgtPimGroup SetSGJoinMode, m.Object, m.Name);
 return nil
}

func(np *PimGroup) GetSGJoinMode ()(string, error) {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtPimGroup GetSGJoinMode
 return "", nil
}

