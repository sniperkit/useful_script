package n2xsdk

type Bgp4UpdateMessageCaptureFilterAfiSafi struct {
 Handler string
}

func(np *Bgp4UpdateMessageCaptureFilterAfiSafi) SetAfiSafi () error {
 //parameters: AfiSafiHandle AFI SAFI
 //AgtBgp4UpdateMessageCaptureFilterAfiSafi SetAfiSafi, m.Object, m.Name);
 return nil
}

func(np *Bgp4UpdateMessageCaptureFilterAfiSafi) GetAfiSafi ()(string, error) {
 //parameters: AfiSafiHandle
 //AgtBgp4UpdateMessageCaptureFilterAfiSafi GetAfiSafi
 return "", nil
}

func(np *Bgp4UpdateMessageCaptureFilterAfiSafi) QueryAfiSafiCapture () error {
 //parameters: AfiSafiHandle
 //AgtBgp4UpdateMessageCaptureFilterAfiSafi QueryAfiSafiCapture, m.Object, m.Name);
 return nil
}

func(np *Bgp4UpdateMessageCaptureFilterAfiSafi) GetNextCapturedMessage ()(string, error) {
 //parameters: AfiSafiHandle
 //AgtBgp4UpdateMessageCaptureFilterAfiSafi GetNextCapturedMessage
 return "", nil
}

func(np *Bgp4UpdateMessageCaptureFilterAfiSafi) SetNextHopFilter () error {
 //parameters: AfiSafiHandle NextHopAddressType NextHopAddress
 //AgtBgp4UpdateMessageCaptureFilterAfiSafi SetNextHopFilter, m.Object, m.Name);
 return nil
}

func(np *Bgp4UpdateMessageCaptureFilterAfiSafi) GetNextHopFilter ()(string, error) {
 //parameters: AfiSafiHandle
 //AgtBgp4UpdateMessageCaptureFilterAfiSafi GetNextHopFilter
 return "", nil
}

func(np *Bgp4UpdateMessageCaptureFilterAfiSafi) SetNextHopFilterState () error {
 //parameters: AfiSafiHandle NextHopFiltering
 //AgtBgp4UpdateMessageCaptureFilterAfiSafi SetNextHopFilterState, m.Object, m.Name);
 return nil
}

func(np *Bgp4UpdateMessageCaptureFilterAfiSafi) IsNextHopFilterEnabled () error {
 //parameters: AfiSafiHandle
 //AgtBgp4UpdateMessageCaptureFilterAfiSafi IsNextHopFilterEnabled, m.Object, m.Name);
 return nil
}

func(np *Bgp4UpdateMessageCaptureFilterAfiSafi) Enable () error {
 //parameters: UpdateMessageCaptureFilterHandle
 //AgtBgp4UpdateMessageCaptureFilterAfiSafi Enable, m.Object, m.Name);
 return nil
}

func(np *Bgp4UpdateMessageCaptureFilterAfiSafi) Disable () error {
 //parameters: UpdateMessageCaptureFilterHandle
 //AgtBgp4UpdateMessageCaptureFilterAfiSafi Disable, m.Object, m.Name);
 return nil
}

func(np *Bgp4UpdateMessageCaptureFilterAfiSafi) Reset () error {
 //parameters: UpdateMessageCaptureFilterHandle
 //AgtBgp4UpdateMessageCaptureFilterAfiSafi Reset, m.Object, m.Name);
 return nil
}

