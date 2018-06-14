package n2xsdk

type FrameMatcherList struct {
 Handler string
}

func(np *FrameMatcherLis) AddFrameMatcher () error {
 //parameters: PortHandle
 //AgtFrameMatcherList AddFrameMatcher, m.Object, m.Name);
 return nil
}

func(np *FrameMatcherLis) RemoveFrameMatcher () error {
 //parameters: PortHandle pFrameMatcherHandle
 //AgtFrameMatcherList RemoveFrameMatcher, m.Object, m.Name);
 return nil
}

func(np *FrameMatcherLis) ListFrameMatchers ()(string, error) {
 //parameters: PortHandle
 //AgtFrameMatcherList ListFrameMatchers
 return "", nil
}

func(np *FrameMatcherLis) GetFreeFrameMatcherCount ()(string, error) {
 //parameters: PortHandle
 //AgtFrameMatcherList GetFreeFrameMatcherCount
 return "", nil
}

func(np *FrameMatcherLis) SetName () error {
 //parameters: FrameMatcherHandle Name
 //AgtFrameMatcherList SetName, m.Object, m.Name);
 return nil
}

func(np *FrameMatcherLis) GetName ()(string, error) {
 //parameters: FrameMatcherHandle
 //AgtFrameMatcherList GetName
 return "", nil
}

