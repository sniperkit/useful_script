package n2xsdk

type EthernetPcsAlarms struct {
 Handler string
}

func(np *EthernetPcsAlarms) SetTransmitAlarm () error {
 //parameters: PortHandle Alarm
 //AgtEthernetPcsAlarms SetTransmitAlarm, m.Object, m.Name);
 return nil
}

func(np *EthernetPcsAlarms) GetTransmitAlarm ()(string, error) {
 //parameters: PortHandle
 //AgtEthernetPcsAlarms GetTransmitAlarm
 return "", nil
}

func(np *EthernetPcsAlarms) IsEthernetPcsAlarmsSupported () error {
 //parameters: PortHandle
 //AgtEthernetPcsAlarms IsEthernetPcsAlarmsSupported, m.Object, m.Name);
 return nil
}

