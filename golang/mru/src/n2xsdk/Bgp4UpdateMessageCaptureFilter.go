package n2xsdk

type Bgp4UpdateMessageCaptureFilter struct {
 Handler string
}

func(np *Bgp4UpdateMessageCaptureFilter) Enable () error {
 //parameters: UpdateMessageCaptureFilterHandle
 //AgtBgp4UpdateMessageCaptureFilter Enable, m.Object, m.Name);
 return nil
}

func(np *Bgp4UpdateMessageCaptureFilter) Disable () error {
 //parameters: UpdateMessageCaptureFilterHandle
 //AgtBgp4UpdateMessageCaptureFilter Disable, m.Object, m.Name);
 return nil
}

func(np *Bgp4UpdateMessageCaptureFilter) Reset () error {
 //parameters: UpdateMessageCaptureFilterHandle
 //AgtBgp4UpdateMessageCaptureFilter Reset, m.Object, m.Name);
 return nil
}

