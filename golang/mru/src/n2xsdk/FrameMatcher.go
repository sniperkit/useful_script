package n2xsdk

type FrameMatcher struct {
 Handler string
}

func(np *FrameMatcher) AddFrameFlags () error {
 //parameters: FrameMatcherHandle FlagCount psaFrameFlags
 //AgtFrameMatcher AddFrameFlags, m.Object, m.Name);
 return nil
}

func(np *FrameMatcher) RemoveFrameFlags () error {
 //parameters: FrameMatcherHandle FlagCount psaFrameFlags
 //AgtFrameMatcher RemoveFrameFlags, m.Object, m.Name);
 return nil
}

func(np *FrameMatcher) ListFrameFlags ()(string, error) {
 //parameters: FrameMatcherHandle
 //AgtFrameMatcher ListFrameFlags
 return "", nil
}

func(np *FrameMatcher) AddPatterns () error {
 //parameters: FrameMatcherHandle PatternCount psaPatternHandles
 //AgtFrameMatcher AddPatterns, m.Object, m.Name);
 return nil
}

func(np *FrameMatcher) RemovePatterns () error {
 //parameters: FrameMatcherHandle PatternCount psaPatternHandles
 //AgtFrameMatcher RemovePatterns, m.Object, m.Name);
 return nil
}

func(np *FrameMatcher) ListPatterns ()(string, error) {
 //parameters: FrameMatcherHandle
 //AgtFrameMatcher ListPatterns
 return "", nil
}

func(np *FrameMatcher) GetFreePatterns ()(string, error) {
 //parameters: FrameMatcherHandle
 //AgtFrameMatcher GetFreePatterns
 return "", nil
}

func(np *FrameMatcher) AddStreamMatch () error {
 //parameters: FrameMatcherHandle StreamGroupHandle StreamIndex
 //AgtFrameMatcher AddStreamMatch, m.Object, m.Name);
 return nil
}

func(np *FrameMatcher) RemoveStreamMatch () error {
 //parameters: FrameMatcherHandle StreamGroupHandle StreamIndex
 //AgtFrameMatcher RemoveStreamMatch, m.Object, m.Name);
 return nil
}

func(np *FrameMatcher) ListStreamMatches ()(string, error) {
 //parameters: FrameMatcherHandle
 //AgtFrameMatcher ListStreamMatches
 return "", nil
}

