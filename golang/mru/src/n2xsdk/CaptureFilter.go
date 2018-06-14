package n2xsdk

type CaptureFilter struct {
 Handler string
}

func(np *CaptureFilter) ClearAllFilters () error {
 //parameters: PortHandle
 //AgtCaptureFilter ClearAllFilters, m.Object, m.Name);
 return nil
}

func(np *CaptureFilter) AddStatusFilter () error {
 //parameters: PortHandle StatusFilter FilterAction
 //AgtCaptureFilter AddStatusFilter, m.Object, m.Name);
 return nil
}

func(np *CaptureFilter) RemoveStatusFilter () error {
 //parameters: PortHandle StatusFilter
 //AgtCaptureFilter RemoveStatusFilter, m.Object, m.Name);
 return nil
}

func(np *CaptureFilter) ListStatusFilters ()(string, error) {
 //parameters: PortHandle
 //AgtCaptureFilter ListStatusFilters
 return "", nil
}

func(np *CaptureFilter) AddFrameMatcherFilters () error {
 //parameters: PortHandle FrameMatcherHandleCount psaFrameMatcherHandles FilterAction
 //AgtCaptureFilter AddFrameMatcherFilters, m.Object, m.Name);
 return nil
}

func(np *CaptureFilter) RemoveFrameMatcherFilters () error {
 //parameters: PortHandle FrameMatcherHandleCount psaFrameMatcherHandles
 //AgtCaptureFilter RemoveFrameMatcherFilters, m.Object, m.Name);
 return nil
}

func(np *CaptureFilter) ListFrameMatcherFilters ()(string, error) {
 //parameters: PortHandle
 //AgtCaptureFilter ListFrameMatcherFilters
 return "", nil
}

