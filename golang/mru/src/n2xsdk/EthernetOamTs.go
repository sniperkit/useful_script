package n2xsdk

type EthernetOamTs struct {
 Handler string
}

func(np *EthernetOamTs) SetDestinationMacAddress () error {
 //parameters: SessionPoolHandle MacAddress
 //AgtEthernetOamTst SetDestinationMacAddress, m.Object, m.Name);
 return nil
}

func(np *EthernetOamTs) GetDestinationMacAddress ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtEthernetOamTst GetDestinationMacAddress
 return "", nil
}

func(np *EthernetOamTs) SetMegLevel () error {
 //parameters: SessionPoolHandle Level
 //AgtEthernetOamTst SetMegLevel, m.Object, m.Name);
 return nil
}

func(np *EthernetOamTs) GetMegLevel ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtEthernetOamTst GetMegLevel
 return "", nil
}

func(np *EthernetOamTs) SetSeqNumber () error {
 //parameters: SessionPoolHandle SeqNumber
 //AgtEthernetOamTst SetSeqNumber, m.Object, m.Name);
 return nil
}

func(np *EthernetOamTs) GetSeqNumber ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtEthernetOamTst GetSeqNumber
 return "", nil
}

func(np *EthernetOamTs) SetTestTlv () error {
 //parameters: SessionPoolHandle TestTlvLength TestTlvPatternType
 //AgtEthernetOamTst SetTestTlv, m.Object, m.Name);
 return nil
}

func(np *EthernetOamTs) GetTestTlv ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtEthernetOamTst GetTestTlv
 return "", nil
}

func(np *EthernetOamTs) SetBurstSize () error {
 //parameters: SessionPoolHandle BurstSize
 //AgtEthernetOamTst SetBurstSize, m.Object, m.Name);
 return nil
}

func(np *EthernetOamTs) GetBurstSize ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtEthernetOamTst GetBurstSize
 return "", nil
}

func(np *EthernetOamTs) SetBurstDelay () error {
 //parameters: SessionPoolHandle BurstDelay
 //AgtEthernetOamTst SetBurstDelay, m.Object, m.Name);
 return nil
}

func(np *EthernetOamTs) GetBurstDelay ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtEthernetOamTst GetBurstDelay
 return "", nil
}

func(np *EthernetOamTs) EnableRenew () error {
 //parameters: SessionPoolHandle
 //AgtEthernetOamTst EnableRenew, m.Object, m.Name);
 return nil
}

func(np *EthernetOamTs) DisableRenew () error {
 //parameters: SessionPoolHandle
 //AgtEthernetOamTst DisableRenew, m.Object, m.Name);
 return nil
}

func(np *EthernetOamTs) IsRenewEnabled () error {
 //parameters: SessionPoolHandle
 //AgtEthernetOamTst IsRenewEnabled, m.Object, m.Name);
 return nil
}

func(np *EthernetOamTs) SetRenewTimeout () error {
 //parameters: SessionPoolHandle RenewTimeout
 //AgtEthernetOamTst SetRenewTimeout, m.Object, m.Name);
 return nil
}

func(np *EthernetOamTs) GetRenewTimeout ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtEthernetOamTst GetRenewTimeout
 return "", nil
}

func(np *EthernetOamTs) Start () error {
 //parameters: SessionPoolHandle
 //AgtEthernetOamTst Start, m.Object, m.Name);
 return nil
}

func(np *EthernetOamTs) SetIncrementingParameter () error {
 //parameters: SessionPoolHandle IncrementingParameter Value Repeat Increment
 //AgtEthernetOamTst SetIncrementingParameter, m.Object, m.Name);
 return nil
}

func(np *EthernetOamTs) GetIncrementingParameter ()(string, error) {
 //parameters: SessionPoolHandle IncrementingParameter
 //AgtEthernetOamTst GetIncrementingParameter
 return "", nil
}

func(np *EthernetOamTs) GetVlanPriority ()(string, error) {
 //parameters: DeviceHandle VlanTagIndex
 //AgtEthernetOamTst GetVlanPriority
 return "", nil
}

func(np *EthernetOamTs) SetVlanPriority () error {
 //parameters: DeviceHandle VlanTagIndex VlanPriority
 //AgtEthernetOamTst SetVlanPriority, m.Object, m.Name);
 return nil
}

