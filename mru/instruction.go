package dsl
import ( 
 "command"
 )
type Instruction struct {
Name string
Switch Switch
}
func PortSlotTypeEnable(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["Port"]; !ok {
      return nil 
   }
if _, ok := in["Slot"]; !ok {
      return nil 
   }
if _, ok := in["Type"]; !ok {
      return nil 
   }
if _, ok := in["Enable"]; !ok {
      return nil 
   }
      return sw.PortSlotTypeEnable(in["Port"], in["Slot"], in["Type"], in["Enable"])
}

func PortSlotTypeDisable(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["Port"]; !ok {
      return nil 
   }
if _, ok := in["Slot"]; !ok {
      return nil 
   }
if _, ok := in["Type"]; !ok {
      return nil 
   }
if _, ok := in["Disable"]; !ok {
      return nil 
   }
      return sw.PortSlotTypeDisable(in["Port"], in["Slot"], in["Type"], in["Disable"])
}

func PortSlotTypeSpeed(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["Port"]; !ok {
      return nil 
   }
if _, ok := in["Slot"]; !ok {
      return nil 
   }
if _, ok := in["Type"]; !ok {
      return nil 
   }
if _, ok := in["Speed"]; !ok {
      return nil 
   }
      return sw.PortSlotTypeSpeed(in["Port"], in["Slot"], in["Type"], in["Speed"])
}

func PortSlotTypePvid(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["Port"]; !ok {
      return nil 
   }
if _, ok := in["Slot"]; !ok {
      return nil 
   }
if _, ok := in["Type"]; !ok {
      return nil 
   }
if _, ok := in["Pvid"]; !ok {
      return nil 
   }
      return sw.PortSlotTypePvid(in["Port"], in["Slot"], in["Type"], in["Pvid"])
}

func VLAN(sw Switch, in map[string]string) []*command.Command{
if len(in) != 1 {
      return nil 
   }
if _, ok := in["VLAN"]; !ok {
      return nil 
   }
      return sw.VLAN(in["VLAN"])
}

func NoVLAN(sw Switch, in map[string]string) []*command.Command{
if len(in) != 2 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["VLAN"]; !ok {
      return nil 
   }
      return sw.NoVLAN(in["VLAN"])
}

func VLANAddTypeSlotPort(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["VLAN"]; !ok {
      return nil 
   }
if _, ok := in["Add"]; !ok {
      return nil 
   }
if _, ok := in["Type"]; !ok {
      return nil 
   }
if _, ok := in["Slot"]; !ok {
      return nil 
   }
if _, ok := in["Port"]; !ok {
      return nil 
   }
      return sw.VLANAddTypeSlotPort(in["VLAN"], in["Add"], in["Type"], in["Slot"], in["Port"])
}

func VLANAddTTypeSlotPort(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["VLAN"]; !ok {
      return nil 
   }
if _, ok := in["AddT"]; !ok {
      return nil 
   }
if _, ok := in["Type"]; !ok {
      return nil 
   }
if _, ok := in["Slot"]; !ok {
      return nil 
   }
if _, ok := in["Port"]; !ok {
      return nil 
   }
      return sw.VLANAddTTypeSlotPort(in["VLAN"], in["AddT"], in["Type"], in["Slot"], in["Port"])
}

func VLANDelTypeSlotPort(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["VLAN"]; !ok {
      return nil 
   }
if _, ok := in["Del"]; !ok {
      return nil 
   }
if _, ok := in["Type"]; !ok {
      return nil 
   }
if _, ok := in["Slot"]; !ok {
      return nil 
   }
if _, ok := in["Port"]; !ok {
      return nil 
   }
      return sw.VLANDelTypeSlotPort(in["VLAN"], in["Del"], in["Type"], in["Slot"], in["Port"])
}

func VLANDelTTypeSlotPort(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["VLAN"]; !ok {
      return nil 
   }
if _, ok := in["DelT"]; !ok {
      return nil 
   }
if _, ok := in["Type"]; !ok {
      return nil 
   }
if _, ok := in["Slot"]; !ok {
      return nil 
   }
if _, ok := in["Port"]; !ok {
      return nil 
   }
      return sw.VLANDelTTypeSlotPort(in["VLAN"], in["DelT"], in["Type"], in["Slot"], in["Port"])
}

func VLANShutdown(sw Switch, in map[string]string) []*command.Command{
if len(in) != 2 {
      return nil 
   }
if _, ok := in["VLAN"]; !ok {
      return nil 
   }
if _, ok := in["Shutdown"]; !ok {
      return nil 
   }
      return sw.VLANShutdown(in["VLAN"], in["Shutdown"])
}

func VLANNoShutdown(sw Switch, in map[string]string) []*command.Command{
if len(in) != 2 {
      return nil 
   }
if _, ok := in["VLAN"]; !ok {
      return nil 
   }
if _, ok := in["NoShutdown"]; !ok {
      return nil 
   }
      return sw.VLANNoShutdown(in["VLAN"], in["NoShutdown"])
}

func VLANIP(sw Switch, in map[string]string) []*command.Command{
if len(in) != 2 {
      return nil 
   }
if _, ok := in["VLAN"]; !ok {
      return nil 
   }
if _, ok := in["IP"]; !ok {
      return nil 
   }
      return sw.VLANIP(in["VLAN"], in["IP"])
}

func NoVLANIP(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["VLAN"]; !ok {
      return nil 
   }
if _, ok := in["IP"]; !ok {
      return nil 
   }
      return sw.NoVLANIP(in["VLAN"], in["IP"])
}

func VLANIP2(sw Switch, in map[string]string) []*command.Command{
if len(in) != 2 {
      return nil 
   }
if _, ok := in["VLAN"]; !ok {
      return nil 
   }
if _, ok := in["IP2"]; !ok {
      return nil 
   }
      return sw.VLANIP2(in["VLAN"], in["IP2"])
}

func NoVLANIP2(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["VLAN"]; !ok {
      return nil 
   }
if _, ok := in["IP2"]; !ok {
      return nil 
   }
      return sw.NoVLANIP2(in["VLAN"], in["IP2"])
}

func NoInterfaceTypeIfname(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["Interface"]; !ok {
      return nil 
   }
if _, ok := in["Type"]; !ok {
      return nil 
   }
if _, ok := in["Ifname"]; !ok {
      return nil 
   }
      return sw.NoInterfaceTypeIfname(in["Interface"], in["Type"], in["Ifname"])
}

func VLANAddTypeSlotPortIP(sw Switch, in map[string]string) []*command.Command{
if len(in) != 6 {
      return nil 
   }
if _, ok := in["VLAN"]; !ok {
      return nil 
   }
if _, ok := in["Add"]; !ok {
      return nil 
   }
if _, ok := in["Type"]; !ok {
      return nil 
   }
if _, ok := in["Slot"]; !ok {
      return nil 
   }
if _, ok := in["Port"]; !ok {
      return nil 
   }
if _, ok := in["IP"]; !ok {
      return nil 
   }
      return sw.VLANAddTypeSlotPortIP(in["VLAN"], in["Add"], in["Type"], in["Slot"], in["Port"], in["IP"])
}

func VLANAddTypeSlotPortIP2(sw Switch, in map[string]string) []*command.Command{
if len(in) != 6 {
      return nil 
   }
if _, ok := in["VLAN"]; !ok {
      return nil 
   }
if _, ok := in["Add"]; !ok {
      return nil 
   }
if _, ok := in["Type"]; !ok {
      return nil 
   }
if _, ok := in["Slot"]; !ok {
      return nil 
   }
if _, ok := in["Port"]; !ok {
      return nil 
   }
if _, ok := in["IP2"]; !ok {
      return nil 
   }
      return sw.VLANAddTypeSlotPortIP2(in["VLAN"], in["Add"], in["Type"], in["Slot"], in["Port"], in["IP2"])
}

func VLANAddTypeSlotPortIPNoShutdown(sw Switch, in map[string]string) []*command.Command{
if len(in) != 7 {
      return nil 
   }
if _, ok := in["VLAN"]; !ok {
      return nil 
   }
if _, ok := in["Add"]; !ok {
      return nil 
   }
if _, ok := in["Type"]; !ok {
      return nil 
   }
if _, ok := in["Slot"]; !ok {
      return nil 
   }
if _, ok := in["Port"]; !ok {
      return nil 
   }
if _, ok := in["IP"]; !ok {
      return nil 
   }
if _, ok := in["NoShutdown"]; !ok {
      return nil 
   }
      return sw.VLANAddTypeSlotPortIPNoShutdown(in["VLAN"], in["Add"], in["Type"], in["Slot"], in["Port"], in["IP"], in["NoShutdown"])
}

func VLANAddTypeSlotPortIP2NoShutdown(sw Switch, in map[string]string) []*command.Command{
if len(in) != 7 {
      return nil 
   }
if _, ok := in["VLAN"]; !ok {
      return nil 
   }
if _, ok := in["Add"]; !ok {
      return nil 
   }
if _, ok := in["Type"]; !ok {
      return nil 
   }
if _, ok := in["Slot"]; !ok {
      return nil 
   }
if _, ok := in["Port"]; !ok {
      return nil 
   }
if _, ok := in["IP2"]; !ok {
      return nil 
   }
if _, ok := in["NoShutdown"]; !ok {
      return nil 
   }
      return sw.VLANAddTypeSlotPortIP2NoShutdown(in["VLAN"], in["Add"], in["Type"], in["Slot"], in["Port"], in["IP2"], in["NoShutdown"])
}

func VLANAddTTypeSlotPortIP(sw Switch, in map[string]string) []*command.Command{
if len(in) != 6 {
      return nil 
   }
if _, ok := in["VLAN"]; !ok {
      return nil 
   }
if _, ok := in["AddT"]; !ok {
      return nil 
   }
if _, ok := in["Type"]; !ok {
      return nil 
   }
if _, ok := in["Slot"]; !ok {
      return nil 
   }
if _, ok := in["Port"]; !ok {
      return nil 
   }
if _, ok := in["IP"]; !ok {
      return nil 
   }
      return sw.VLANAddTTypeSlotPortIP(in["VLAN"], in["AddT"], in["Type"], in["Slot"], in["Port"], in["IP"])
}

func VLANAddTTypeSlotPortIP2(sw Switch, in map[string]string) []*command.Command{
if len(in) != 6 {
      return nil 
   }
if _, ok := in["VLAN"]; !ok {
      return nil 
   }
if _, ok := in["AddT"]; !ok {
      return nil 
   }
if _, ok := in["Type"]; !ok {
      return nil 
   }
if _, ok := in["Slot"]; !ok {
      return nil 
   }
if _, ok := in["Port"]; !ok {
      return nil 
   }
if _, ok := in["IP2"]; !ok {
      return nil 
   }
      return sw.VLANAddTTypeSlotPortIP2(in["VLAN"], in["AddT"], in["Type"], in["Slot"], in["Port"], in["IP2"])
}

func VLANAddTTypeSlotPortIPNoShutdown(sw Switch, in map[string]string) []*command.Command{
if len(in) != 7 {
      return nil 
   }
if _, ok := in["VLAN"]; !ok {
      return nil 
   }
if _, ok := in["AddT"]; !ok {
      return nil 
   }
if _, ok := in["Type"]; !ok {
      return nil 
   }
if _, ok := in["Slot"]; !ok {
      return nil 
   }
if _, ok := in["Port"]; !ok {
      return nil 
   }
if _, ok := in["IP"]; !ok {
      return nil 
   }
if _, ok := in["NoShutdown"]; !ok {
      return nil 
   }
      return sw.VLANAddTTypeSlotPortIPNoShutdown(in["VLAN"], in["AddT"], in["Type"], in["Slot"], in["Port"], in["IP"], in["NoShutdown"])
}

func VLANAddTTypeSlotPortIP2NoShutdown(sw Switch, in map[string]string) []*command.Command{
if len(in) != 7 {
      return nil 
   }
if _, ok := in["VLAN"]; !ok {
      return nil 
   }
if _, ok := in["AddT"]; !ok {
      return nil 
   }
if _, ok := in["Type"]; !ok {
      return nil 
   }
if _, ok := in["Slot"]; !ok {
      return nil 
   }
if _, ok := in["Port"]; !ok {
      return nil 
   }
if _, ok := in["IP2"]; !ok {
      return nil 
   }
if _, ok := in["NoShutdown"]; !ok {
      return nil 
   }
      return sw.VLANAddTTypeSlotPortIP2NoShutdown(in["VLAN"], in["AddT"], in["Type"], in["Slot"], in["Port"], in["IP2"], in["NoShutdown"])
}

func VLANIP6Enable(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["VLAN"]; !ok {
      return nil 
   }
if _, ok := in["IP6"]; !ok {
      return nil 
   }
if _, ok := in["Enable"]; !ok {
      return nil 
   }
      return sw.VLANIP6Enable(in["VLAN"], in["IP6"], in["Enable"])
}

func NoVLANIP6Enable(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["VLAN"]; !ok {
      return nil 
   }
if _, ok := in["IP6"]; !ok {
      return nil 
   }
if _, ok := in["Enable"]; !ok {
      return nil 
   }
      return sw.NoVLANIP6Enable(in["VLAN"], in["IP6"], in["Enable"])
}

func VLANIP6(sw Switch, in map[string]string) []*command.Command{
if len(in) != 2 {
      return nil 
   }
if _, ok := in["VLAN"]; !ok {
      return nil 
   }
if _, ok := in["IP6"]; !ok {
      return nil 
   }
      return sw.VLANIP6(in["VLAN"], in["IP6"])
}

func NoVLANIP6(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["VLAN"]; !ok {
      return nil 
   }
if _, ok := in["IP6"]; !ok {
      return nil 
   }
      return sw.NoVLANIP6(in["VLAN"], in["IP6"])
}

func VLANIP6LL(sw Switch, in map[string]string) []*command.Command{
if len(in) != 2 {
      return nil 
   }
if _, ok := in["VLAN"]; !ok {
      return nil 
   }
if _, ok := in["IP6LL"]; !ok {
      return nil 
   }
      return sw.VLANIP6LL(in["VLAN"], in["IP6LL"])
}

func VLANIP6LLIP6(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["VLAN"]; !ok {
      return nil 
   }
if _, ok := in["IP6LL"]; !ok {
      return nil 
   }
if _, ok := in["IP6"]; !ok {
      return nil 
   }
      return sw.VLANIP6LLIP6(in["VLAN"], in["IP6LL"], in["IP6"])
}

func NoVLANIP6LL(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["VLAN"]; !ok {
      return nil 
   }
if _, ok := in["IP6LL"]; !ok {
      return nil 
   }
      return sw.NoVLANIP6LL(in["VLAN"], in["IP6LL"])
}

func OSPF6(sw Switch, in map[string]string) []*command.Command{
if len(in) != 1 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
      return sw.OSPF6(in["OSPF6"])
}

func NoOSPF6(sw Switch, in map[string]string) []*command.Command{
if len(in) != 2 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
      return sw.NoOSPF6(in["OSPF6"])
}

func OSPF6Rid(sw Switch, in map[string]string) []*command.Command{
if len(in) != 2 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Rid"]; !ok {
      return nil 
   }
      return sw.OSPF6Rid(in["OSPF6"], in["Rid"])
}

func NoOSPF6Rid(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Rid"]; !ok {
      return nil 
   }
      return sw.NoOSPF6Rid(in["OSPF6"], in["Rid"])
}

func OSPF6InterfaceArea(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Interface"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
      return sw.OSPF6InterfaceArea(in["OSPF6"], in["Interface"], in["Area"])
}

func NoOSPF6InterfaceArea(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Interface"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
      return sw.NoOSPF6InterfaceArea(in["OSPF6"], in["Interface"], in["Area"])
}

func OSPF6InterfaceCost(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Interface"]; !ok {
      return nil 
   }
if _, ok := in["Cost"]; !ok {
      return nil 
   }
      return sw.OSPF6InterfaceCost(in["OSPF6"], in["Interface"], in["Cost"])
}

func OSPF6InterfaceDeadInterval(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Interface"]; !ok {
      return nil 
   }
if _, ok := in["DeadInterval"]; !ok {
      return nil 
   }
      return sw.OSPF6InterfaceDeadInterval(in["OSPF6"], in["Interface"], in["DeadInterval"])
}

func OSPF6InterfaceHelloInterval(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Interface"]; !ok {
      return nil 
   }
if _, ok := in["HelloInterval"]; !ok {
      return nil 
   }
      return sw.OSPF6InterfaceHelloInterval(in["OSPF6"], in["Interface"], in["HelloInterval"])
}

func OSPF6InterfaceRetransmitInterval(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Interface"]; !ok {
      return nil 
   }
if _, ok := in["RetransmitInterval"]; !ok {
      return nil 
   }
      return sw.OSPF6InterfaceRetransmitInterval(in["OSPF6"], in["Interface"], in["RetransmitInterval"])
}

func OSPF6InterfaceTransmitDelay(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Interface"]; !ok {
      return nil 
   }
if _, ok := in["TransmitDelay"]; !ok {
      return nil 
   }
      return sw.OSPF6InterfaceTransmitDelay(in["OSPF6"], in["Interface"], in["TransmitDelay"])
}

func OSPF6InterfacePriority(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Interface"]; !ok {
      return nil 
   }
if _, ok := in["Priority"]; !ok {
      return nil 
   }
      return sw.OSPF6InterfacePriority(in["OSPF6"], in["Interface"], in["Priority"])
}

func OSPF6InterfaceNetworktype(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Interface"]; !ok {
      return nil 
   }
if _, ok := in["Networktype"]; !ok {
      return nil 
   }
      return sw.OSPF6InterfaceNetworktype(in["OSPF6"], in["Interface"], in["Networktype"])
}

func NoOSPF6InterfaceCost(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Interface"]; !ok {
      return nil 
   }
if _, ok := in["Cost"]; !ok {
      return nil 
   }
      return sw.NoOSPF6InterfaceCost(in["OSPF6"], in["Interface"], in["Cost"])
}

func NoOSPF6InterfaceDeadInterval(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Interface"]; !ok {
      return nil 
   }
if _, ok := in["DeadInterval"]; !ok {
      return nil 
   }
      return sw.NoOSPF6InterfaceDeadInterval(in["OSPF6"], in["Interface"], in["DeadInterval"])
}

func NoOSPF6InterfaceHelloInterval(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Interface"]; !ok {
      return nil 
   }
if _, ok := in["HelloInterval"]; !ok {
      return nil 
   }
      return sw.NoOSPF6InterfaceHelloInterval(in["OSPF6"], in["Interface"], in["HelloInterval"])
}

func NoOSPF6InterfaceRetransmitInterval(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Interface"]; !ok {
      return nil 
   }
if _, ok := in["RetransmitInterval"]; !ok {
      return nil 
   }
      return sw.NoOSPF6InterfaceRetransmitInterval(in["OSPF6"], in["Interface"], in["RetransmitInterval"])
}

func NoOSPF6InterfaceTransmitDelay(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Interface"]; !ok {
      return nil 
   }
if _, ok := in["TransmitDelay"]; !ok {
      return nil 
   }
      return sw.NoOSPF6InterfaceTransmitDelay(in["OSPF6"], in["Interface"], in["TransmitDelay"])
}

func NoOSPF6InterfacePriority(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Interface"]; !ok {
      return nil 
   }
if _, ok := in["Priority"]; !ok {
      return nil 
   }
      return sw.NoOSPF6InterfacePriority(in["OSPF6"], in["Interface"], in["Priority"])
}

func NoOSPF6InterfaceNetworktype(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Interface"]; !ok {
      return nil 
   }
if _, ok := in["Networktype"]; !ok {
      return nil 
   }
      return sw.NoOSPF6InterfaceNetworktype(in["OSPF6"], in["Interface"], in["Networktype"])
}

func OSPF6ReferenceBandwidth(sw Switch, in map[string]string) []*command.Command{
if len(in) != 2 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["ReferenceBandwidth"]; !ok {
      return nil 
   }
      return sw.OSPF6ReferenceBandwidth(in["OSPF6"], in["ReferenceBandwidth"])
}

func NoOSPF6ReferenceBandwidth(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["ReferenceBandwidth"]; !ok {
      return nil 
   }
      return sw.NoOSPF6ReferenceBandwidth(in["OSPF6"], in["ReferenceBandwidth"])
}

func OSPF6DefaultOriginate(sw Switch, in map[string]string) []*command.Command{
if len(in) != 2 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
      return sw.OSPF6DefaultOriginate(in["OSPF6"], in["DefaultOriginate"])
}

func OSPF6DefaultOriginateRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.OSPF6DefaultOriginateRoutemap(in["OSPF6"], in["DefaultOriginate"], in["Routemap"])
}

func OSPF6DefaultOriginateMetric(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
      return sw.OSPF6DefaultOriginateMetric(in["OSPF6"], in["DefaultOriginate"], in["Metric"])
}

func OSPF6DefaultOriginateMetrictype(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
      return sw.OSPF6DefaultOriginateMetrictype(in["OSPF6"], in["DefaultOriginate"], in["Metrictype"])
}

func OSPF6DefaultOriginateMetricMetrictype(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
      return sw.OSPF6DefaultOriginateMetricMetrictype(in["OSPF6"], in["DefaultOriginate"], in["Metric"], in["Metrictype"])
}

func OSPF6DefaultOriginateMetricRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.OSPF6DefaultOriginateMetricRoutemap(in["OSPF6"], in["DefaultOriginate"], in["Metric"], in["Routemap"])
}

func OSPF6DefaultOriginateMetrictypeRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.OSPF6DefaultOriginateMetrictypeRoutemap(in["OSPF6"], in["DefaultOriginate"], in["Metrictype"], in["Routemap"])
}

func OSPF6DefaultOriginateMetricMetrictypeRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.OSPF6DefaultOriginateMetricMetrictypeRoutemap(in["OSPF6"], in["DefaultOriginate"], in["Metric"], in["Metrictype"], in["Routemap"])
}

func OSPF6DefaultOriginateAlways(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Always"]; !ok {
      return nil 
   }
      return sw.OSPF6DefaultOriginateAlways(in["OSPF6"], in["DefaultOriginate"], in["Always"])
}

func OSPF6DefaultOriginateAlwaysRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Always"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.OSPF6DefaultOriginateAlwaysRoutemap(in["OSPF6"], in["DefaultOriginate"], in["Always"], in["Routemap"])
}

func OSPF6DefaultOriginateAlwaysMetric(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Always"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
      return sw.OSPF6DefaultOriginateAlwaysMetric(in["OSPF6"], in["DefaultOriginate"], in["Always"], in["Metric"])
}

func OSPF6DefaultOriginateAlwaysMetrictype(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Always"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
      return sw.OSPF6DefaultOriginateAlwaysMetrictype(in["OSPF6"], in["DefaultOriginate"], in["Always"], in["Metrictype"])
}

func OSPF6DefaultOriginateAlwaysMetricMetrictype(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Always"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
      return sw.OSPF6DefaultOriginateAlwaysMetricMetrictype(in["OSPF6"], in["DefaultOriginate"], in["Always"], in["Metric"], in["Metrictype"])
}

func OSPF6DefaultOriginateAlwaysMetricRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Always"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.OSPF6DefaultOriginateAlwaysMetricRoutemap(in["OSPF6"], in["DefaultOriginate"], in["Always"], in["Metric"], in["Routemap"])
}

func OSPF6DefaultOriginateAlwaysMetrictypeRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Always"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.OSPF6DefaultOriginateAlwaysMetrictypeRoutemap(in["OSPF6"], in["DefaultOriginate"], in["Always"], in["Metrictype"], in["Routemap"])
}

func OSPF6DefaultOriginateAlwaysMetricMetrictypeRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 6 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Always"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.OSPF6DefaultOriginateAlwaysMetricMetrictypeRoutemap(in["OSPF6"], in["DefaultOriginate"], in["Always"], in["Metric"], in["Metrictype"], in["Routemap"])
}

func NoOSPF6DefaultOriginate(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
      return sw.NoOSPF6DefaultOriginate(in["OSPF6"], in["DefaultOriginate"])
}

func NoOSPF6DefaultOriginateRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.NoOSPF6DefaultOriginateRoutemap(in["OSPF6"], in["DefaultOriginate"], in["Routemap"])
}

func NoOSPF6DefaultOriginateMetric(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
      return sw.NoOSPF6DefaultOriginateMetric(in["OSPF6"], in["DefaultOriginate"], in["Metric"])
}

func NoOSPF6DefaultOriginateMetrictype(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
      return sw.NoOSPF6DefaultOriginateMetrictype(in["OSPF6"], in["DefaultOriginate"], in["Metrictype"])
}

func NoOSPF6DefaultOriginateMetricMetrictype(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
      return sw.NoOSPF6DefaultOriginateMetricMetrictype(in["OSPF6"], in["DefaultOriginate"], in["Metric"], in["Metrictype"])
}

func NoOSPF6DefaultOriginateMetricRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.NoOSPF6DefaultOriginateMetricRoutemap(in["OSPF6"], in["DefaultOriginate"], in["Metric"], in["Routemap"])
}

func NoOSPF6DefaultOriginateMetrictypeRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.NoOSPF6DefaultOriginateMetrictypeRoutemap(in["OSPF6"], in["DefaultOriginate"], in["Metrictype"], in["Routemap"])
}

func NoOSPF6DefaultOriginateMetricMetrictypeRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 6 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.NoOSPF6DefaultOriginateMetricMetrictypeRoutemap(in["OSPF6"], in["DefaultOriginate"], in["Metric"], in["Metrictype"], in["Routemap"])
}

func NoOSPF6DefaultOriginateAlways(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Always"]; !ok {
      return nil 
   }
      return sw.NoOSPF6DefaultOriginateAlways(in["OSPF6"], in["DefaultOriginate"], in["Always"])
}

func NoOSPF6DefaultOriginateAlwaysRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Always"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.NoOSPF6DefaultOriginateAlwaysRoutemap(in["OSPF6"], in["DefaultOriginate"], in["Always"], in["Routemap"])
}

func NoOSPF6DefaultOriginateAlwaysMetric(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Always"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
      return sw.NoOSPF6DefaultOriginateAlwaysMetric(in["OSPF6"], in["DefaultOriginate"], in["Always"], in["Metric"])
}

func NoOSPF6DefaultOriginateAlwaysMetrictype(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Always"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
      return sw.NoOSPF6DefaultOriginateAlwaysMetrictype(in["OSPF6"], in["DefaultOriginate"], in["Always"], in["Metrictype"])
}

func NoOSPF6DefaultOriginateAlwaysMetricMetrictype(sw Switch, in map[string]string) []*command.Command{
if len(in) != 6 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Always"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
      return sw.NoOSPF6DefaultOriginateAlwaysMetricMetrictype(in["OSPF6"], in["DefaultOriginate"], in["Always"], in["Metric"], in["Metrictype"])
}

func NoOSPF6DefaultOriginateAlwaysMetricRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 6 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Always"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.NoOSPF6DefaultOriginateAlwaysMetricRoutemap(in["OSPF6"], in["DefaultOriginate"], in["Always"], in["Metric"], in["Routemap"])
}

func NoOSPF6DefaultOriginateAlwaysMetrictypeRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 6 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Always"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.NoOSPF6DefaultOriginateAlwaysMetrictypeRoutemap(in["OSPF6"], in["DefaultOriginate"], in["Always"], in["Metrictype"], in["Routemap"])
}

func NoOSPF6DefaultOriginateAlwaysMetricMetrictypeRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 7 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Always"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.NoOSPF6DefaultOriginateAlwaysMetricMetrictypeRoutemap(in["OSPF6"], in["DefaultOriginate"], in["Always"], in["Metric"], in["Metrictype"], in["Routemap"])
}

func OSPF6Redistribute(sw Switch, in map[string]string) []*command.Command{
if len(in) != 2 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Redistribute"]; !ok {
      return nil 
   }
      return sw.OSPF6Redistribute(in["OSPF6"], in["Redistribute"])
}

func OSPF6RedistributeMetric(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Redistribute"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
      return sw.OSPF6RedistributeMetric(in["OSPF6"], in["Redistribute"], in["Metric"])
}

func OSPF6RedistributeMetrictype(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Redistribute"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
      return sw.OSPF6RedistributeMetrictype(in["OSPF6"], in["Redistribute"], in["Metrictype"])
}

func OSPF6RedistributeRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Redistribute"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.OSPF6RedistributeRoutemap(in["OSPF6"], in["Redistribute"], in["Routemap"])
}

func OSPF6RedistributeMetricMetrictype(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Redistribute"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
      return sw.OSPF6RedistributeMetricMetrictype(in["OSPF6"], in["Redistribute"], in["Metric"], in["Metrictype"])
}

func OSPF6RedistributeMetricRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Redistribute"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.OSPF6RedistributeMetricRoutemap(in["OSPF6"], in["Redistribute"], in["Metric"], in["Routemap"])
}

func OSPF6RedistributeMetricMetrictypeRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Redistribute"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.OSPF6RedistributeMetricMetrictypeRoutemap(in["OSPF6"], in["Redistribute"], in["Metric"], in["Metrictype"], in["Routemap"])
}

func NoOSPF6Redistribute(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Redistribute"]; !ok {
      return nil 
   }
      return sw.NoOSPF6Redistribute(in["OSPF6"], in["Redistribute"])
}

func NoOSPF6RedistributeMetric(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Redistribute"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
      return sw.NoOSPF6RedistributeMetric(in["OSPF6"], in["Redistribute"], in["Metric"])
}

func NoOSPF6RedistributeMetrictype(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Redistribute"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
      return sw.NoOSPF6RedistributeMetrictype(in["OSPF6"], in["Redistribute"], in["Metrictype"])
}

func NoOSPF6RedistributeRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Redistribute"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.NoOSPF6RedistributeRoutemap(in["OSPF6"], in["Redistribute"], in["Routemap"])
}

func NoOSPF6RedistributeMetricMetrictype(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Redistribute"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
      return sw.NoOSPF6RedistributeMetricMetrictype(in["OSPF6"], in["Redistribute"], in["Metric"], in["Metrictype"])
}

func NoOSPF6RedistributeMetricRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Redistribute"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.NoOSPF6RedistributeMetricRoutemap(in["OSPF6"], in["Redistribute"], in["Metric"], in["Routemap"])
}

func NoOSPF6RedistributeMetricMetrictypeRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 6 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Redistribute"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.NoOSPF6RedistributeMetricMetrictypeRoutemap(in["OSPF6"], in["Redistribute"], in["Metric"], in["Metrictype"], in["Routemap"])
}

func OSPF6Summary(sw Switch, in map[string]string) []*command.Command{
if len(in) != 2 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Summary"]; !ok {
      return nil 
   }
      return sw.OSPF6Summary(in["OSPF6"], in["Summary"])
}

func OSPF6SummaryNoAdvertise(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Summary"]; !ok {
      return nil 
   }
if _, ok := in["NoAdvertise"]; !ok {
      return nil 
   }
      return sw.OSPF6SummaryNoAdvertise(in["OSPF6"], in["Summary"], in["NoAdvertise"])
}

func NoOSPF6Summary(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Summary"]; !ok {
      return nil 
   }
      return sw.NoOSPF6Summary(in["OSPF6"], in["Summary"])
}

func NoOSPF6SummaryNoAdvertise(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Summary"]; !ok {
      return nil 
   }
if _, ok := in["NoAdvertise"]; !ok {
      return nil 
   }
      return sw.NoOSPF6SummaryNoAdvertise(in["OSPF6"], in["Summary"], in["NoAdvertise"])
}

func OSPF6DefaultMetric(sw Switch, in map[string]string) []*command.Command{
if len(in) != 2 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultMetric"]; !ok {
      return nil 
   }
      return sw.OSPF6DefaultMetric(in["OSPF6"], in["DefaultMetric"])
}

func NoOSPF6DefaultMetric(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["DefaultMetric"]; !ok {
      return nil 
   }
      return sw.NoOSPF6DefaultMetric(in["OSPF6"], in["DefaultMetric"])
}

func OSPF6Passive(sw Switch, in map[string]string) []*command.Command{
if len(in) != 2 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Passive"]; !ok {
      return nil 
   }
      return sw.OSPF6Passive(in["OSPF6"], in["Passive"])
}

func NoOSPF6Passive(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Passive"]; !ok {
      return nil 
   }
      return sw.NoOSPF6Passive(in["OSPF6"], in["Passive"])
}

func OSPF6AdminDistance(sw Switch, in map[string]string) []*command.Command{
if len(in) != 2 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["AdminDistance"]; !ok {
      return nil 
   }
      return sw.OSPF6AdminDistance(in["OSPF6"], in["AdminDistance"])
}

func NoOSPF6AdminDistance(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["AdminDistance"]; !ok {
      return nil 
   }
      return sw.NoOSPF6AdminDistance(in["OSPF6"], in["AdminDistance"])
}

func OSPF6DistanceExternal(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Distance"]; !ok {
      return nil 
   }
if _, ok := in["External"]; !ok {
      return nil 
   }
      return sw.OSPF6DistanceExternal(in["OSPF6"], in["Distance"], in["External"])
}

func OSPF6DistanceInter(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Distance"]; !ok {
      return nil 
   }
if _, ok := in["Inter"]; !ok {
      return nil 
   }
      return sw.OSPF6DistanceInter(in["OSPF6"], in["Distance"], in["Inter"])
}

func OSPF6DistanceIntra(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Distance"]; !ok {
      return nil 
   }
if _, ok := in["Intra"]; !ok {
      return nil 
   }
      return sw.OSPF6DistanceIntra(in["OSPF6"], in["Distance"], in["Intra"])
}

func OSPF6DistanceInterIntra(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Distance"]; !ok {
      return nil 
   }
if _, ok := in["Inter"]; !ok {
      return nil 
   }
if _, ok := in["Intra"]; !ok {
      return nil 
   }
      return sw.OSPF6DistanceInterIntra(in["OSPF6"], in["Distance"], in["Inter"], in["Intra"])
}

func OSPF6DistanceInterExternal(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Distance"]; !ok {
      return nil 
   }
if _, ok := in["Inter"]; !ok {
      return nil 
   }
if _, ok := in["External"]; !ok {
      return nil 
   }
      return sw.OSPF6DistanceInterExternal(in["OSPF6"], in["Distance"], in["Inter"], in["External"])
}

func OSPF6DistanceInterIntraExternal(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Distance"]; !ok {
      return nil 
   }
if _, ok := in["Inter"]; !ok {
      return nil 
   }
if _, ok := in["Intra"]; !ok {
      return nil 
   }
if _, ok := in["External"]; !ok {
      return nil 
   }
      return sw.OSPF6DistanceInterIntraExternal(in["OSPF6"], in["Distance"], in["Inter"], in["Intra"], in["External"])
}

func NoOSPF6DistanceExternal(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Distance"]; !ok {
      return nil 
   }
if _, ok := in["External"]; !ok {
      return nil 
   }
      return sw.NoOSPF6DistanceExternal(in["OSPF6"], in["Distance"], in["External"])
}

func NoOSPF6DistanceInter(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Distance"]; !ok {
      return nil 
   }
if _, ok := in["Inter"]; !ok {
      return nil 
   }
      return sw.NoOSPF6DistanceInter(in["OSPF6"], in["Distance"], in["Inter"])
}

func NoOSPF6DistanceIntra(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Distance"]; !ok {
      return nil 
   }
if _, ok := in["Intra"]; !ok {
      return nil 
   }
      return sw.NoOSPF6DistanceIntra(in["OSPF6"], in["Distance"], in["Intra"])
}

func NoOSPF6DistanceInterIntra(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Distance"]; !ok {
      return nil 
   }
if _, ok := in["Inter"]; !ok {
      return nil 
   }
if _, ok := in["Intra"]; !ok {
      return nil 
   }
      return sw.NoOSPF6DistanceInterIntra(in["OSPF6"], in["Distance"], in["Inter"], in["Intra"])
}

func NoOSPF6DistanceInterExternal(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Distance"]; !ok {
      return nil 
   }
if _, ok := in["Inter"]; !ok {
      return nil 
   }
if _, ok := in["External"]; !ok {
      return nil 
   }
      return sw.NoOSPF6DistanceInterExternal(in["OSPF6"], in["Distance"], in["Inter"], in["External"])
}

func NoOSPF6DistanceInterIntraExternal(sw Switch, in map[string]string) []*command.Command{
if len(in) != 6 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Distance"]; !ok {
      return nil 
   }
if _, ok := in["Inter"]; !ok {
      return nil 
   }
if _, ok := in["Intra"]; !ok {
      return nil 
   }
if _, ok := in["External"]; !ok {
      return nil 
   }
      return sw.NoOSPF6DistanceInterIntraExternal(in["OSPF6"], in["Distance"], in["Inter"], in["Intra"], in["External"])
}

func OSPF6DistributelistIN(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Distributelist"]; !ok {
      return nil 
   }
if _, ok := in["IN"]; !ok {
      return nil 
   }
      return sw.OSPF6DistributelistIN(in["OSPF6"], in["Distributelist"], in["IN"])
}

func OSPF6DistributelistOUT(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Distributelist"]; !ok {
      return nil 
   }
if _, ok := in["OUT"]; !ok {
      return nil 
   }
      return sw.OSPF6DistributelistOUT(in["OSPF6"], in["Distributelist"], in["OUT"])
}

func NoOSPF6DistributelistIN(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Distributelist"]; !ok {
      return nil 
   }
if _, ok := in["IN"]; !ok {
      return nil 
   }
      return sw.NoOSPF6DistributelistIN(in["OSPF6"], in["Distributelist"], in["IN"])
}

func NoOSPF6DistributelistOUT(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Distributelist"]; !ok {
      return nil 
   }
if _, ok := in["OUT"]; !ok {
      return nil 
   }
      return sw.NoOSPF6DistributelistOUT(in["OSPF6"], in["Distributelist"], in["OUT"])
}

func OSPF6AreaDefaultCost(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["DefaultCost"]; !ok {
      return nil 
   }
      return sw.OSPF6AreaDefaultCost(in["OSPF6"], in["Area"], in["DefaultCost"])
}

func NoOSPF6AreaDefaultCost(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["DefaultCost"]; !ok {
      return nil 
   }
      return sw.NoOSPF6AreaDefaultCost(in["OSPF6"], in["Area"], in["DefaultCost"])
}

func OSPF6AreaNSSA(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["NSSA"]; !ok {
      return nil 
   }
      return sw.OSPF6AreaNSSA(in["OSPF6"], in["Area"], in["NSSA"])
}

func OSPF6AreaNSSADefaultOriginate(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["NSSA"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
      return sw.OSPF6AreaNSSADefaultOriginate(in["OSPF6"], in["Area"], in["NSSA"], in["DefaultOriginate"])
}

func OSPF6AreaNSSANoRedistribution(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["NSSA"]; !ok {
      return nil 
   }
if _, ok := in["NoRedistribution"]; !ok {
      return nil 
   }
      return sw.OSPF6AreaNSSANoRedistribution(in["OSPF6"], in["Area"], in["NSSA"], in["NoRedistribution"])
}

func OSPF6AreaNSSANoSummary(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["NSSA"]; !ok {
      return nil 
   }
if _, ok := in["NoSummary"]; !ok {
      return nil 
   }
      return sw.OSPF6AreaNSSANoSummary(in["OSPF6"], in["Area"], in["NSSA"], in["NoSummary"])
}

func OSPF6AreaNSSAStabilityInterval(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["NSSA"]; !ok {
      return nil 
   }
if _, ok := in["StabilityInterval"]; !ok {
      return nil 
   }
      return sw.OSPF6AreaNSSAStabilityInterval(in["OSPF6"], in["Area"], in["NSSA"], in["StabilityInterval"])
}

func OSPF6AreaTranslatorrole(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Translatorrole"]; !ok {
      return nil 
   }
      return sw.OSPF6AreaTranslatorrole(in["OSPF6"], in["Area"], in["Translatorrole"])
}

func NoOSPF6AreaNSSA(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["NSSA"]; !ok {
      return nil 
   }
      return sw.NoOSPF6AreaNSSA(in["OSPF6"], in["Area"], in["NSSA"])
}

func NoOSPF6AreaNSSADefaultOriginate(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["NSSA"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
      return sw.NoOSPF6AreaNSSADefaultOriginate(in["OSPF6"], in["Area"], in["NSSA"], in["DefaultOriginate"])
}

func NoOSPF6AreaNSSANoRedistribution(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["NSSA"]; !ok {
      return nil 
   }
if _, ok := in["NoRedistribution"]; !ok {
      return nil 
   }
      return sw.NoOSPF6AreaNSSANoRedistribution(in["OSPF6"], in["Area"], in["NSSA"], in["NoRedistribution"])
}

func NoOSPF6AreaNSSANoSummary(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["NSSA"]; !ok {
      return nil 
   }
if _, ok := in["NoSummary"]; !ok {
      return nil 
   }
      return sw.NoOSPF6AreaNSSANoSummary(in["OSPF6"], in["Area"], in["NSSA"], in["NoSummary"])
}

func NoOSPF6AreaNSSAStabilityInterval(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["NSSA"]; !ok {
      return nil 
   }
if _, ok := in["StabilityInterval"]; !ok {
      return nil 
   }
      return sw.NoOSPF6AreaNSSAStabilityInterval(in["OSPF6"], in["Area"], in["NSSA"], in["StabilityInterval"])
}

func NoOSPF6AreaTranslatorrole(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Translatorrole"]; !ok {
      return nil 
   }
      return sw.NoOSPF6AreaTranslatorrole(in["OSPF6"], in["Area"], in["Translatorrole"])
}

func OSPF6AreaStub(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Stub"]; !ok {
      return nil 
   }
      return sw.OSPF6AreaStub(in["OSPF6"], in["Area"], in["Stub"])
}

func OSPF6AreaStubNoSummary(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Stub"]; !ok {
      return nil 
   }
if _, ok := in["NoSummary"]; !ok {
      return nil 
   }
      return sw.OSPF6AreaStubNoSummary(in["OSPF6"], in["Area"], in["Stub"], in["NoSummary"])
}

func NoOSPF6AreaStub(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Stub"]; !ok {
      return nil 
   }
      return sw.NoOSPF6AreaStub(in["OSPF6"], in["Area"], in["Stub"])
}

func NoOSPF6AreaStubNoSummary(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Stub"]; !ok {
      return nil 
   }
if _, ok := in["NoSummary"]; !ok {
      return nil 
   }
      return sw.NoOSPF6AreaStubNoSummary(in["OSPF6"], in["Area"], in["Stub"], in["NoSummary"])
}

func OSPF6AreaRange(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Range"]; !ok {
      return nil 
   }
      return sw.OSPF6AreaRange(in["OSPF6"], in["Area"], in["Range"])
}

func OSPF6AreaRangeAdvertise(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Range"]; !ok {
      return nil 
   }
if _, ok := in["Advertise"]; !ok {
      return nil 
   }
      return sw.OSPF6AreaRangeAdvertise(in["OSPF6"], in["Area"], in["Range"], in["Advertise"])
}

func OSPF6AreaRangeNoAdvertise(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Range"]; !ok {
      return nil 
   }
if _, ok := in["NoAdvertise"]; !ok {
      return nil 
   }
      return sw.OSPF6AreaRangeNoAdvertise(in["OSPF6"], in["Area"], in["Range"], in["NoAdvertise"])
}

func NoOSPF6AreaRange(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Range"]; !ok {
      return nil 
   }
      return sw.NoOSPF6AreaRange(in["OSPF6"], in["Area"], in["Range"])
}

func NoOSPF6AreaRangeAdvertise(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Range"]; !ok {
      return nil 
   }
if _, ok := in["Advertise"]; !ok {
      return nil 
   }
      return sw.NoOSPF6AreaRangeAdvertise(in["OSPF6"], in["Area"], in["Range"], in["Advertise"])
}

func NoOSPF6AreaRangeNoAdvertise(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Range"]; !ok {
      return nil 
   }
if _, ok := in["NoAdvertise"]; !ok {
      return nil 
   }
      return sw.NoOSPF6AreaRangeNoAdvertise(in["OSPF6"], in["Area"], in["Range"], in["NoAdvertise"])
}

func OSPF6AreaVirtuallink(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Virtuallink"]; !ok {
      return nil 
   }
      return sw.OSPF6AreaVirtuallink(in["OSPF6"], in["Area"], in["Virtuallink"])
}

func OSPF6AreaVirtuallinkDeadInterval(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Virtuallink"]; !ok {
      return nil 
   }
if _, ok := in["DeadInterval"]; !ok {
      return nil 
   }
      return sw.OSPF6AreaVirtuallinkDeadInterval(in["OSPF6"], in["Area"], in["Virtuallink"], in["DeadInterval"])
}

func OSPF6AreaVirtuallinkHelloInterval(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Virtuallink"]; !ok {
      return nil 
   }
if _, ok := in["HelloInterval"]; !ok {
      return nil 
   }
      return sw.OSPF6AreaVirtuallinkHelloInterval(in["OSPF6"], in["Area"], in["Virtuallink"], in["HelloInterval"])
}

func OSPF6AreaVirtuallinkInstanceid(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Virtuallink"]; !ok {
      return nil 
   }
if _, ok := in["Instanceid"]; !ok {
      return nil 
   }
      return sw.OSPF6AreaVirtuallinkInstanceid(in["OSPF6"], in["Area"], in["Virtuallink"], in["Instanceid"])
}

func OSPF6AreaVirtuallinkRetransmitInterval(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Virtuallink"]; !ok {
      return nil 
   }
if _, ok := in["RetransmitInterval"]; !ok {
      return nil 
   }
      return sw.OSPF6AreaVirtuallinkRetransmitInterval(in["OSPF6"], in["Area"], in["Virtuallink"], in["RetransmitInterval"])
}

func OSPF6AreaVirtuallinkTransmitDelay(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Virtuallink"]; !ok {
      return nil 
   }
if _, ok := in["TransmitDelay"]; !ok {
      return nil 
   }
      return sw.OSPF6AreaVirtuallinkTransmitDelay(in["OSPF6"], in["Area"], in["Virtuallink"], in["TransmitDelay"])
}

func NoOSPF6AreaVirtuallink(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Virtuallink"]; !ok {
      return nil 
   }
      return sw.NoOSPF6AreaVirtuallink(in["OSPF6"], in["Area"], in["Virtuallink"])
}

func NoOSPF6AreaVirtuallinkDeadInterval(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Virtuallink"]; !ok {
      return nil 
   }
if _, ok := in["DeadInterval"]; !ok {
      return nil 
   }
      return sw.NoOSPF6AreaVirtuallinkDeadInterval(in["OSPF6"], in["Area"], in["Virtuallink"], in["DeadInterval"])
}

func NoOSPF6AreaVirtuallinkHelloInterval(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Virtuallink"]; !ok {
      return nil 
   }
if _, ok := in["HelloInterval"]; !ok {
      return nil 
   }
      return sw.NoOSPF6AreaVirtuallinkHelloInterval(in["OSPF6"], in["Area"], in["Virtuallink"], in["HelloInterval"])
}

func NoOSPF6AreaVirtuallinkInstanceid(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Virtuallink"]; !ok {
      return nil 
   }
if _, ok := in["Instanceid"]; !ok {
      return nil 
   }
      return sw.NoOSPF6AreaVirtuallinkInstanceid(in["OSPF6"], in["Area"], in["Virtuallink"], in["Instanceid"])
}

func NoOSPF6AreaVirtuallinkRetransmitInterval(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Virtuallink"]; !ok {
      return nil 
   }
if _, ok := in["RetransmitInterval"]; !ok {
      return nil 
   }
      return sw.NoOSPF6AreaVirtuallinkRetransmitInterval(in["OSPF6"], in["Area"], in["Virtuallink"], in["RetransmitInterval"])
}

func NoOSPF6AreaVirtuallinkTransmitDelay(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF6"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Virtuallink"]; !ok {
      return nil 
   }
if _, ok := in["TransmitDelay"]; !ok {
      return nil 
   }
      return sw.NoOSPF6AreaVirtuallinkTransmitDelay(in["OSPF6"], in["Area"], in["Virtuallink"], in["TransmitDelay"])
}

func OSPF(sw Switch, in map[string]string) []*command.Command{
if len(in) != 1 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
      return sw.OSPF(in["OSPF"])
}

func NoOSPF(sw Switch, in map[string]string) []*command.Command{
if len(in) != 2 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
      return sw.NoOSPF(in["OSPF"])
}

func OSPFRid(sw Switch, in map[string]string) []*command.Command{
if len(in) != 2 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Rid"]; !ok {
      return nil 
   }
      return sw.OSPFRid(in["OSPF"], in["Rid"])
}

func NoOSPFRid(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Rid"]; !ok {
      return nil 
   }
      return sw.NoOSPFRid(in["OSPF"], in["Rid"])
}

func OSPFNetworkArea(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Network"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
      return sw.OSPFNetworkArea(in["OSPF"], in["Network"], in["Area"])
}

func NoOSPFNetworkArea(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Network"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
      return sw.NoOSPFNetworkArea(in["OSPF"], in["Network"], in["Area"])
}

func OSPFInterfaceCost(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Interface"]; !ok {
      return nil 
   }
if _, ok := in["Cost"]; !ok {
      return nil 
   }
      return sw.OSPFInterfaceCost(in["OSPF"], in["Interface"], in["Cost"])
}

func OSPFInterfaceDeadInterval(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Interface"]; !ok {
      return nil 
   }
if _, ok := in["DeadInterval"]; !ok {
      return nil 
   }
      return sw.OSPFInterfaceDeadInterval(in["OSPF"], in["Interface"], in["DeadInterval"])
}

func OSPFInterfaceHelloInterval(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Interface"]; !ok {
      return nil 
   }
if _, ok := in["HelloInterval"]; !ok {
      return nil 
   }
      return sw.OSPFInterfaceHelloInterval(in["OSPF"], in["Interface"], in["HelloInterval"])
}

func OSPFInterfaceRetransmitInterval(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Interface"]; !ok {
      return nil 
   }
if _, ok := in["RetransmitInterval"]; !ok {
      return nil 
   }
      return sw.OSPFInterfaceRetransmitInterval(in["OSPF"], in["Interface"], in["RetransmitInterval"])
}

func OSPFInterfaceTransmitDelay(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Interface"]; !ok {
      return nil 
   }
if _, ok := in["TransmitDelay"]; !ok {
      return nil 
   }
      return sw.OSPFInterfaceTransmitDelay(in["OSPF"], in["Interface"], in["TransmitDelay"])
}

func OSPFInterfacePriority(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Interface"]; !ok {
      return nil 
   }
if _, ok := in["Priority"]; !ok {
      return nil 
   }
      return sw.OSPFInterfacePriority(in["OSPF"], in["Interface"], in["Priority"])
}

func OSPFInterfaceNetworktype(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Interface"]; !ok {
      return nil 
   }
if _, ok := in["Networktype"]; !ok {
      return nil 
   }
      return sw.OSPFInterfaceNetworktype(in["OSPF"], in["Interface"], in["Networktype"])
}

func NoOSPFInterfaceCost(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Interface"]; !ok {
      return nil 
   }
if _, ok := in["Cost"]; !ok {
      return nil 
   }
      return sw.NoOSPFInterfaceCost(in["OSPF"], in["Interface"], in["Cost"])
}

func NoOSPFInterfaceDeadInterval(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Interface"]; !ok {
      return nil 
   }
if _, ok := in["DeadInterval"]; !ok {
      return nil 
   }
      return sw.NoOSPFInterfaceDeadInterval(in["OSPF"], in["Interface"], in["DeadInterval"])
}

func NoOSPFInterfaceHelloInterval(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Interface"]; !ok {
      return nil 
   }
if _, ok := in["HelloInterval"]; !ok {
      return nil 
   }
      return sw.NoOSPFInterfaceHelloInterval(in["OSPF"], in["Interface"], in["HelloInterval"])
}

func NoOSPFInterfaceRetransmitInterval(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Interface"]; !ok {
      return nil 
   }
if _, ok := in["RetransmitInterval"]; !ok {
      return nil 
   }
      return sw.NoOSPFInterfaceRetransmitInterval(in["OSPF"], in["Interface"], in["RetransmitInterval"])
}

func NoOSPFInterfaceTransmitDelay(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Interface"]; !ok {
      return nil 
   }
if _, ok := in["TransmitDelay"]; !ok {
      return nil 
   }
      return sw.NoOSPFInterfaceTransmitDelay(in["OSPF"], in["Interface"], in["TransmitDelay"])
}

func NoOSPFInterfacePriority(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Interface"]; !ok {
      return nil 
   }
if _, ok := in["Priority"]; !ok {
      return nil 
   }
      return sw.NoOSPFInterfacePriority(in["OSPF"], in["Interface"], in["Priority"])
}

func NoOSPFInterfaceNetworktype(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Interface"]; !ok {
      return nil 
   }
if _, ok := in["Networktype"]; !ok {
      return nil 
   }
      return sw.NoOSPFInterfaceNetworktype(in["OSPF"], in["Interface"], in["Networktype"])
}

func OSPFReferenceBandwidth(sw Switch, in map[string]string) []*command.Command{
if len(in) != 2 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["ReferenceBandwidth"]; !ok {
      return nil 
   }
      return sw.OSPFReferenceBandwidth(in["OSPF"], in["ReferenceBandwidth"])
}

func NoOSPFReferenceBandwidth(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["ReferenceBandwidth"]; !ok {
      return nil 
   }
      return sw.NoOSPFReferenceBandwidth(in["OSPF"], in["ReferenceBandwidth"])
}

func OSPFDefaultOriginate(sw Switch, in map[string]string) []*command.Command{
if len(in) != 2 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
      return sw.OSPFDefaultOriginate(in["OSPF"], in["DefaultOriginate"])
}

func OSPFDefaultOriginateRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.OSPFDefaultOriginateRoutemap(in["OSPF"], in["DefaultOriginate"], in["Routemap"])
}

func OSPFDefaultOriginateMetric(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
      return sw.OSPFDefaultOriginateMetric(in["OSPF"], in["DefaultOriginate"], in["Metric"])
}

func OSPFDefaultOriginateMetrictype(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
      return sw.OSPFDefaultOriginateMetrictype(in["OSPF"], in["DefaultOriginate"], in["Metrictype"])
}

func OSPFDefaultOriginateMetricMetrictype(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
      return sw.OSPFDefaultOriginateMetricMetrictype(in["OSPF"], in["DefaultOriginate"], in["Metric"], in["Metrictype"])
}

func OSPFDefaultOriginateMetricRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.OSPFDefaultOriginateMetricRoutemap(in["OSPF"], in["DefaultOriginate"], in["Metric"], in["Routemap"])
}

func OSPFDefaultOriginateMetrictypeRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.OSPFDefaultOriginateMetrictypeRoutemap(in["OSPF"], in["DefaultOriginate"], in["Metrictype"], in["Routemap"])
}

func OSPFDefaultOriginateMetricMetrictypeRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.OSPFDefaultOriginateMetricMetrictypeRoutemap(in["OSPF"], in["DefaultOriginate"], in["Metric"], in["Metrictype"], in["Routemap"])
}

func OSPFDefaultOriginateAlways(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Always"]; !ok {
      return nil 
   }
      return sw.OSPFDefaultOriginateAlways(in["OSPF"], in["DefaultOriginate"], in["Always"])
}

func OSPFDefaultOriginateAlwaysRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Always"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.OSPFDefaultOriginateAlwaysRoutemap(in["OSPF"], in["DefaultOriginate"], in["Always"], in["Routemap"])
}

func OSPFDefaultOriginateAlwaysMetric(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Always"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
      return sw.OSPFDefaultOriginateAlwaysMetric(in["OSPF"], in["DefaultOriginate"], in["Always"], in["Metric"])
}

func OSPFDefaultOriginateAlwaysMetrictype(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Always"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
      return sw.OSPFDefaultOriginateAlwaysMetrictype(in["OSPF"], in["DefaultOriginate"], in["Always"], in["Metrictype"])
}

func OSPFDefaultOriginateAlwaysMetricMetrictype(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Always"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
      return sw.OSPFDefaultOriginateAlwaysMetricMetrictype(in["OSPF"], in["DefaultOriginate"], in["Always"], in["Metric"], in["Metrictype"])
}

func OSPFDefaultOriginateAlwaysMetricRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Always"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.OSPFDefaultOriginateAlwaysMetricRoutemap(in["OSPF"], in["DefaultOriginate"], in["Always"], in["Metric"], in["Routemap"])
}

func OSPFDefaultOriginateAlwaysMetrictypeRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Always"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.OSPFDefaultOriginateAlwaysMetrictypeRoutemap(in["OSPF"], in["DefaultOriginate"], in["Always"], in["Metrictype"], in["Routemap"])
}

func OSPFDefaultOriginateAlwaysMetricMetrictypeRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 6 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Always"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.OSPFDefaultOriginateAlwaysMetricMetrictypeRoutemap(in["OSPF"], in["DefaultOriginate"], in["Always"], in["Metric"], in["Metrictype"], in["Routemap"])
}

func NoOSPFDefaultOriginate(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
      return sw.NoOSPFDefaultOriginate(in["OSPF"], in["DefaultOriginate"])
}

func NoOSPFDefaultOriginateRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.NoOSPFDefaultOriginateRoutemap(in["OSPF"], in["DefaultOriginate"], in["Routemap"])
}

func NoOSPFDefaultOriginateMetric(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
      return sw.NoOSPFDefaultOriginateMetric(in["OSPF"], in["DefaultOriginate"], in["Metric"])
}

func NoOSPFDefaultOriginateMetrictype(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
      return sw.NoOSPFDefaultOriginateMetrictype(in["OSPF"], in["DefaultOriginate"], in["Metrictype"])
}

func NoOSPFDefaultOriginateMetricMetrictype(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
      return sw.NoOSPFDefaultOriginateMetricMetrictype(in["OSPF"], in["DefaultOriginate"], in["Metric"], in["Metrictype"])
}

func NoOSPFDefaultOriginateMetricRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.NoOSPFDefaultOriginateMetricRoutemap(in["OSPF"], in["DefaultOriginate"], in["Metric"], in["Routemap"])
}

func NoOSPFDefaultOriginateMetrictypeRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.NoOSPFDefaultOriginateMetrictypeRoutemap(in["OSPF"], in["DefaultOriginate"], in["Metrictype"], in["Routemap"])
}

func NoOSPFDefaultOriginateMetricMetrictypeRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 6 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.NoOSPFDefaultOriginateMetricMetrictypeRoutemap(in["OSPF"], in["DefaultOriginate"], in["Metric"], in["Metrictype"], in["Routemap"])
}

func NoOSPFDefaultOriginateAlways(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Always"]; !ok {
      return nil 
   }
      return sw.NoOSPFDefaultOriginateAlways(in["OSPF"], in["DefaultOriginate"], in["Always"])
}

func NoOSPFDefaultOriginateAlwaysRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Always"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.NoOSPFDefaultOriginateAlwaysRoutemap(in["OSPF"], in["DefaultOriginate"], in["Always"], in["Routemap"])
}

func NoOSPFDefaultOriginateAlwaysMetric(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Always"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
      return sw.NoOSPFDefaultOriginateAlwaysMetric(in["OSPF"], in["DefaultOriginate"], in["Always"], in["Metric"])
}

func NoOSPFDefaultOriginateAlwaysMetrictype(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Always"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
      return sw.NoOSPFDefaultOriginateAlwaysMetrictype(in["OSPF"], in["DefaultOriginate"], in["Always"], in["Metrictype"])
}

func NoOSPFDefaultOriginateAlwaysMetricMetrictype(sw Switch, in map[string]string) []*command.Command{
if len(in) != 6 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Always"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
      return sw.NoOSPFDefaultOriginateAlwaysMetricMetrictype(in["OSPF"], in["DefaultOriginate"], in["Always"], in["Metric"], in["Metrictype"])
}

func NoOSPFDefaultOriginateAlwaysMetricRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 6 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Always"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.NoOSPFDefaultOriginateAlwaysMetricRoutemap(in["OSPF"], in["DefaultOriginate"], in["Always"], in["Metric"], in["Routemap"])
}

func NoOSPFDefaultOriginateAlwaysMetrictypeRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 6 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Always"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.NoOSPFDefaultOriginateAlwaysMetrictypeRoutemap(in["OSPF"], in["DefaultOriginate"], in["Always"], in["Metrictype"], in["Routemap"])
}

func NoOSPFDefaultOriginateAlwaysMetricMetrictypeRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 7 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
if _, ok := in["Always"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.NoOSPFDefaultOriginateAlwaysMetricMetrictypeRoutemap(in["OSPF"], in["DefaultOriginate"], in["Always"], in["Metric"], in["Metrictype"], in["Routemap"])
}

func OSPFRedistribute(sw Switch, in map[string]string) []*command.Command{
if len(in) != 2 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Redistribute"]; !ok {
      return nil 
   }
      return sw.OSPFRedistribute(in["OSPF"], in["Redistribute"])
}

func OSPFRedistributeMetric(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Redistribute"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
      return sw.OSPFRedistributeMetric(in["OSPF"], in["Redistribute"], in["Metric"])
}

func OSPFRedistributeMetrictype(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Redistribute"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
      return sw.OSPFRedistributeMetrictype(in["OSPF"], in["Redistribute"], in["Metrictype"])
}

func OSPFRedistributeRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Redistribute"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.OSPFRedistributeRoutemap(in["OSPF"], in["Redistribute"], in["Routemap"])
}

func OSPFRedistributeMetricMetrictype(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Redistribute"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
      return sw.OSPFRedistributeMetricMetrictype(in["OSPF"], in["Redistribute"], in["Metric"], in["Metrictype"])
}

func OSPFRedistributeMetricRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Redistribute"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.OSPFRedistributeMetricRoutemap(in["OSPF"], in["Redistribute"], in["Metric"], in["Routemap"])
}

func OSPFRedistributeMetricMetrictypeRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Redistribute"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.OSPFRedistributeMetricMetrictypeRoutemap(in["OSPF"], in["Redistribute"], in["Metric"], in["Metrictype"], in["Routemap"])
}

func NoOSPFRedistribute(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Redistribute"]; !ok {
      return nil 
   }
      return sw.NoOSPFRedistribute(in["OSPF"], in["Redistribute"])
}

func NoOSPFRedistributeMetric(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Redistribute"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
      return sw.NoOSPFRedistributeMetric(in["OSPF"], in["Redistribute"], in["Metric"])
}

func NoOSPFRedistributeMetrictype(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Redistribute"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
      return sw.NoOSPFRedistributeMetrictype(in["OSPF"], in["Redistribute"], in["Metrictype"])
}

func NoOSPFRedistributeRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Redistribute"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.NoOSPFRedistributeRoutemap(in["OSPF"], in["Redistribute"], in["Routemap"])
}

func NoOSPFRedistributeMetricMetrictype(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Redistribute"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
      return sw.NoOSPFRedistributeMetricMetrictype(in["OSPF"], in["Redistribute"], in["Metric"], in["Metrictype"])
}

func NoOSPFRedistributeMetricRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Redistribute"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.NoOSPFRedistributeMetricRoutemap(in["OSPF"], in["Redistribute"], in["Metric"], in["Routemap"])
}

func NoOSPFRedistributeMetricMetrictypeRoutemap(sw Switch, in map[string]string) []*command.Command{
if len(in) != 6 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Redistribute"]; !ok {
      return nil 
   }
if _, ok := in["Metric"]; !ok {
      return nil 
   }
if _, ok := in["Metrictype"]; !ok {
      return nil 
   }
if _, ok := in["Routemap"]; !ok {
      return nil 
   }
      return sw.NoOSPFRedistributeMetricMetrictypeRoutemap(in["OSPF"], in["Redistribute"], in["Metric"], in["Metrictype"], in["Routemap"])
}

func OSPFSummary(sw Switch, in map[string]string) []*command.Command{
if len(in) != 2 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Summary"]; !ok {
      return nil 
   }
      return sw.OSPFSummary(in["OSPF"], in["Summary"])
}

func OSPFSummaryNoAdvertise(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Summary"]; !ok {
      return nil 
   }
if _, ok := in["NoAdvertise"]; !ok {
      return nil 
   }
      return sw.OSPFSummaryNoAdvertise(in["OSPF"], in["Summary"], in["NoAdvertise"])
}

func NoOSPFSummary(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Summary"]; !ok {
      return nil 
   }
      return sw.NoOSPFSummary(in["OSPF"], in["Summary"])
}

func NoOSPFSummaryNoAdvertise(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Summary"]; !ok {
      return nil 
   }
if _, ok := in["NoAdvertise"]; !ok {
      return nil 
   }
      return sw.NoOSPFSummaryNoAdvertise(in["OSPF"], in["Summary"], in["NoAdvertise"])
}

func OSPFDefaultMetric(sw Switch, in map[string]string) []*command.Command{
if len(in) != 2 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultMetric"]; !ok {
      return nil 
   }
      return sw.OSPFDefaultMetric(in["OSPF"], in["DefaultMetric"])
}

func NoOSPFDefaultMetric(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["DefaultMetric"]; !ok {
      return nil 
   }
      return sw.NoOSPFDefaultMetric(in["OSPF"], in["DefaultMetric"])
}

func OSPFPassive(sw Switch, in map[string]string) []*command.Command{
if len(in) != 2 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Passive"]; !ok {
      return nil 
   }
      return sw.OSPFPassive(in["OSPF"], in["Passive"])
}

func NoOSPFPassive(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Passive"]; !ok {
      return nil 
   }
      return sw.NoOSPFPassive(in["OSPF"], in["Passive"])
}

func OSPFAdminDistance(sw Switch, in map[string]string) []*command.Command{
if len(in) != 2 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["AdminDistance"]; !ok {
      return nil 
   }
      return sw.OSPFAdminDistance(in["OSPF"], in["AdminDistance"])
}

func NoOSPFAdminDistance(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["AdminDistance"]; !ok {
      return nil 
   }
      return sw.NoOSPFAdminDistance(in["OSPF"], in["AdminDistance"])
}

func OSPFDistanceExternal(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Distance"]; !ok {
      return nil 
   }
if _, ok := in["External"]; !ok {
      return nil 
   }
      return sw.OSPFDistanceExternal(in["OSPF"], in["Distance"], in["External"])
}

func OSPFDistanceInter(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Distance"]; !ok {
      return nil 
   }
if _, ok := in["Inter"]; !ok {
      return nil 
   }
      return sw.OSPFDistanceInter(in["OSPF"], in["Distance"], in["Inter"])
}

func OSPFDistanceIntra(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Distance"]; !ok {
      return nil 
   }
if _, ok := in["Intra"]; !ok {
      return nil 
   }
      return sw.OSPFDistanceIntra(in["OSPF"], in["Distance"], in["Intra"])
}

func OSPFDistanceInterIntra(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Distance"]; !ok {
      return nil 
   }
if _, ok := in["Inter"]; !ok {
      return nil 
   }
if _, ok := in["Intra"]; !ok {
      return nil 
   }
      return sw.OSPFDistanceInterIntra(in["OSPF"], in["Distance"], in["Inter"], in["Intra"])
}

func OSPFDistanceInterExternal(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Distance"]; !ok {
      return nil 
   }
if _, ok := in["Inter"]; !ok {
      return nil 
   }
if _, ok := in["External"]; !ok {
      return nil 
   }
      return sw.OSPFDistanceInterExternal(in["OSPF"], in["Distance"], in["Inter"], in["External"])
}

func OSPFDistanceInterIntraExternal(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Distance"]; !ok {
      return nil 
   }
if _, ok := in["Inter"]; !ok {
      return nil 
   }
if _, ok := in["Intra"]; !ok {
      return nil 
   }
if _, ok := in["External"]; !ok {
      return nil 
   }
      return sw.OSPFDistanceInterIntraExternal(in["OSPF"], in["Distance"], in["Inter"], in["Intra"], in["External"])
}

func NoOSPFDistanceExternal(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Distance"]; !ok {
      return nil 
   }
if _, ok := in["External"]; !ok {
      return nil 
   }
      return sw.NoOSPFDistanceExternal(in["OSPF"], in["Distance"], in["External"])
}

func NoOSPFDistanceInter(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Distance"]; !ok {
      return nil 
   }
if _, ok := in["Inter"]; !ok {
      return nil 
   }
      return sw.NoOSPFDistanceInter(in["OSPF"], in["Distance"], in["Inter"])
}

func NoOSPFDistanceIntra(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Distance"]; !ok {
      return nil 
   }
if _, ok := in["Intra"]; !ok {
      return nil 
   }
      return sw.NoOSPFDistanceIntra(in["OSPF"], in["Distance"], in["Intra"])
}

func NoOSPFDistanceInterIntra(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Distance"]; !ok {
      return nil 
   }
if _, ok := in["Inter"]; !ok {
      return nil 
   }
if _, ok := in["Intra"]; !ok {
      return nil 
   }
      return sw.NoOSPFDistanceInterIntra(in["OSPF"], in["Distance"], in["Inter"], in["Intra"])
}

func NoOSPFDistanceInterExternal(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Distance"]; !ok {
      return nil 
   }
if _, ok := in["Inter"]; !ok {
      return nil 
   }
if _, ok := in["External"]; !ok {
      return nil 
   }
      return sw.NoOSPFDistanceInterExternal(in["OSPF"], in["Distance"], in["Inter"], in["External"])
}

func NoOSPFDistanceInterIntraExternal(sw Switch, in map[string]string) []*command.Command{
if len(in) != 6 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Distance"]; !ok {
      return nil 
   }
if _, ok := in["Inter"]; !ok {
      return nil 
   }
if _, ok := in["Intra"]; !ok {
      return nil 
   }
if _, ok := in["External"]; !ok {
      return nil 
   }
      return sw.NoOSPFDistanceInterIntraExternal(in["OSPF"], in["Distance"], in["Inter"], in["Intra"], in["External"])
}

func OSPFDistributelistIN(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Distributelist"]; !ok {
      return nil 
   }
if _, ok := in["IN"]; !ok {
      return nil 
   }
      return sw.OSPFDistributelistIN(in["OSPF"], in["Distributelist"], in["IN"])
}

func OSPFDistributelistOUT(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Distributelist"]; !ok {
      return nil 
   }
if _, ok := in["OUT"]; !ok {
      return nil 
   }
      return sw.OSPFDistributelistOUT(in["OSPF"], in["Distributelist"], in["OUT"])
}

func NoOSPFDistributelistIN(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Distributelist"]; !ok {
      return nil 
   }
if _, ok := in["IN"]; !ok {
      return nil 
   }
      return sw.NoOSPFDistributelistIN(in["OSPF"], in["Distributelist"], in["IN"])
}

func NoOSPFDistributelistOUT(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Distributelist"]; !ok {
      return nil 
   }
if _, ok := in["OUT"]; !ok {
      return nil 
   }
      return sw.NoOSPFDistributelistOUT(in["OSPF"], in["Distributelist"], in["OUT"])
}

func OSPFAreaDefaultCost(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["DefaultCost"]; !ok {
      return nil 
   }
      return sw.OSPFAreaDefaultCost(in["OSPF"], in["Area"], in["DefaultCost"])
}

func NoOSPFAreaDefaultCost(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["DefaultCost"]; !ok {
      return nil 
   }
      return sw.NoOSPFAreaDefaultCost(in["OSPF"], in["Area"], in["DefaultCost"])
}

func OSPFAreaNSSA(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["NSSA"]; !ok {
      return nil 
   }
      return sw.OSPFAreaNSSA(in["OSPF"], in["Area"], in["NSSA"])
}

func OSPFAreaNSSADefaultOriginate(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["NSSA"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
      return sw.OSPFAreaNSSADefaultOriginate(in["OSPF"], in["Area"], in["NSSA"], in["DefaultOriginate"])
}

func OSPFAreaNSSANoRedistribution(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["NSSA"]; !ok {
      return nil 
   }
if _, ok := in["NoRedistribution"]; !ok {
      return nil 
   }
      return sw.OSPFAreaNSSANoRedistribution(in["OSPF"], in["Area"], in["NSSA"], in["NoRedistribution"])
}

func OSPFAreaNSSANoSummary(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["NSSA"]; !ok {
      return nil 
   }
if _, ok := in["NoSummary"]; !ok {
      return nil 
   }
      return sw.OSPFAreaNSSANoSummary(in["OSPF"], in["Area"], in["NSSA"], in["NoSummary"])
}

func OSPFAreaNSSAStabilityInterval(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["NSSA"]; !ok {
      return nil 
   }
if _, ok := in["StabilityInterval"]; !ok {
      return nil 
   }
      return sw.OSPFAreaNSSAStabilityInterval(in["OSPF"], in["Area"], in["NSSA"], in["StabilityInterval"])
}

func OSPFAreaTranslatorrole(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Translatorrole"]; !ok {
      return nil 
   }
      return sw.OSPFAreaTranslatorrole(in["OSPF"], in["Area"], in["Translatorrole"])
}

func NoOSPFAreaNSSA(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["NSSA"]; !ok {
      return nil 
   }
      return sw.NoOSPFAreaNSSA(in["OSPF"], in["Area"], in["NSSA"])
}

func NoOSPFAreaNSSADefaultOriginate(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["NSSA"]; !ok {
      return nil 
   }
if _, ok := in["DefaultOriginate"]; !ok {
      return nil 
   }
      return sw.NoOSPFAreaNSSADefaultOriginate(in["OSPF"], in["Area"], in["NSSA"], in["DefaultOriginate"])
}

func NoOSPFAreaNSSANoRedistribution(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["NSSA"]; !ok {
      return nil 
   }
if _, ok := in["NoRedistribution"]; !ok {
      return nil 
   }
      return sw.NoOSPFAreaNSSANoRedistribution(in["OSPF"], in["Area"], in["NSSA"], in["NoRedistribution"])
}

func NoOSPFAreaNSSANoSummary(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["NSSA"]; !ok {
      return nil 
   }
if _, ok := in["NoSummary"]; !ok {
      return nil 
   }
      return sw.NoOSPFAreaNSSANoSummary(in["OSPF"], in["Area"], in["NSSA"], in["NoSummary"])
}

func NoOSPFAreaNSSAStabilityInterval(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["NSSA"]; !ok {
      return nil 
   }
if _, ok := in["StabilityInterval"]; !ok {
      return nil 
   }
      return sw.NoOSPFAreaNSSAStabilityInterval(in["OSPF"], in["Area"], in["NSSA"], in["StabilityInterval"])
}

func NoOSPFAreaTranslatorrole(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Translatorrole"]; !ok {
      return nil 
   }
      return sw.NoOSPFAreaTranslatorrole(in["OSPF"], in["Area"], in["Translatorrole"])
}

func OSPFAreaStub(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Stub"]; !ok {
      return nil 
   }
      return sw.OSPFAreaStub(in["OSPF"], in["Area"], in["Stub"])
}

func OSPFAreaStubNoSummary(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Stub"]; !ok {
      return nil 
   }
if _, ok := in["NoSummary"]; !ok {
      return nil 
   }
      return sw.OSPFAreaStubNoSummary(in["OSPF"], in["Area"], in["Stub"], in["NoSummary"])
}

func NoOSPFAreaStub(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Stub"]; !ok {
      return nil 
   }
      return sw.NoOSPFAreaStub(in["OSPF"], in["Area"], in["Stub"])
}

func NoOSPFAreaStubNoSummary(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Stub"]; !ok {
      return nil 
   }
if _, ok := in["NoSummary"]; !ok {
      return nil 
   }
      return sw.NoOSPFAreaStubNoSummary(in["OSPF"], in["Area"], in["Stub"], in["NoSummary"])
}

func OSPFAreaRange(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Range"]; !ok {
      return nil 
   }
      return sw.OSPFAreaRange(in["OSPF"], in["Area"], in["Range"])
}

func OSPFAreaRangeAdvertise(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Range"]; !ok {
      return nil 
   }
if _, ok := in["Advertise"]; !ok {
      return nil 
   }
      return sw.OSPFAreaRangeAdvertise(in["OSPF"], in["Area"], in["Range"], in["Advertise"])
}

func OSPFAreaRangeNoAdvertise(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Range"]; !ok {
      return nil 
   }
if _, ok := in["NoAdvertise"]; !ok {
      return nil 
   }
      return sw.OSPFAreaRangeNoAdvertise(in["OSPF"], in["Area"], in["Range"], in["NoAdvertise"])
}

func NoOSPFAreaRange(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Range"]; !ok {
      return nil 
   }
      return sw.NoOSPFAreaRange(in["OSPF"], in["Area"], in["Range"])
}

func NoOSPFAreaRangeAdvertise(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Range"]; !ok {
      return nil 
   }
if _, ok := in["Advertise"]; !ok {
      return nil 
   }
      return sw.NoOSPFAreaRangeAdvertise(in["OSPF"], in["Area"], in["Range"], in["Advertise"])
}

func NoOSPFAreaRangeNoAdvertise(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Range"]; !ok {
      return nil 
   }
if _, ok := in["NoAdvertise"]; !ok {
      return nil 
   }
      return sw.NoOSPFAreaRangeNoAdvertise(in["OSPF"], in["Area"], in["Range"], in["NoAdvertise"])
}

func OSPFAreaVirtuallink(sw Switch, in map[string]string) []*command.Command{
if len(in) != 3 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Virtuallink"]; !ok {
      return nil 
   }
      return sw.OSPFAreaVirtuallink(in["OSPF"], in["Area"], in["Virtuallink"])
}

func OSPFAreaVirtuallinkDeadInterval(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Virtuallink"]; !ok {
      return nil 
   }
if _, ok := in["DeadInterval"]; !ok {
      return nil 
   }
      return sw.OSPFAreaVirtuallinkDeadInterval(in["OSPF"], in["Area"], in["Virtuallink"], in["DeadInterval"])
}

func OSPFAreaVirtuallinkHelloInterval(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Virtuallink"]; !ok {
      return nil 
   }
if _, ok := in["HelloInterval"]; !ok {
      return nil 
   }
      return sw.OSPFAreaVirtuallinkHelloInterval(in["OSPF"], in["Area"], in["Virtuallink"], in["HelloInterval"])
}

func OSPFAreaVirtuallinkInstanceid(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Virtuallink"]; !ok {
      return nil 
   }
if _, ok := in["Instanceid"]; !ok {
      return nil 
   }
      return sw.OSPFAreaVirtuallinkInstanceid(in["OSPF"], in["Area"], in["Virtuallink"], in["Instanceid"])
}

func OSPFAreaVirtuallinkRetransmitInterval(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Virtuallink"]; !ok {
      return nil 
   }
if _, ok := in["RetransmitInterval"]; !ok {
      return nil 
   }
      return sw.OSPFAreaVirtuallinkRetransmitInterval(in["OSPF"], in["Area"], in["Virtuallink"], in["RetransmitInterval"])
}

func OSPFAreaVirtuallinkTransmitDelay(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Virtuallink"]; !ok {
      return nil 
   }
if _, ok := in["TransmitDelay"]; !ok {
      return nil 
   }
      return sw.OSPFAreaVirtuallinkTransmitDelay(in["OSPF"], in["Area"], in["Virtuallink"], in["TransmitDelay"])
}

func NoOSPFAreaVirtuallink(sw Switch, in map[string]string) []*command.Command{
if len(in) != 4 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Virtuallink"]; !ok {
      return nil 
   }
      return sw.NoOSPFAreaVirtuallink(in["OSPF"], in["Area"], in["Virtuallink"])
}

func NoOSPFAreaVirtuallinkDeadInterval(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Virtuallink"]; !ok {
      return nil 
   }
if _, ok := in["DeadInterval"]; !ok {
      return nil 
   }
      return sw.NoOSPFAreaVirtuallinkDeadInterval(in["OSPF"], in["Area"], in["Virtuallink"], in["DeadInterval"])
}

func NoOSPFAreaVirtuallinkHelloInterval(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Virtuallink"]; !ok {
      return nil 
   }
if _, ok := in["HelloInterval"]; !ok {
      return nil 
   }
      return sw.NoOSPFAreaVirtuallinkHelloInterval(in["OSPF"], in["Area"], in["Virtuallink"], in["HelloInterval"])
}

func NoOSPFAreaVirtuallinkInstanceid(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Virtuallink"]; !ok {
      return nil 
   }
if _, ok := in["Instanceid"]; !ok {
      return nil 
   }
      return sw.NoOSPFAreaVirtuallinkInstanceid(in["OSPF"], in["Area"], in["Virtuallink"], in["Instanceid"])
}

func NoOSPFAreaVirtuallinkRetransmitInterval(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Virtuallink"]; !ok {
      return nil 
   }
if _, ok := in["RetransmitInterval"]; !ok {
      return nil 
   }
      return sw.NoOSPFAreaVirtuallinkRetransmitInterval(in["OSPF"], in["Area"], in["Virtuallink"], in["RetransmitInterval"])
}

func NoOSPFAreaVirtuallinkTransmitDelay(sw Switch, in map[string]string) []*command.Command{
if len(in) != 5 {
      return nil 
   }
if _, ok := in["No"]; !ok {
      return nil 
   }
if _, ok := in["OSPF"]; !ok {
      return nil 
   }
if _, ok := in["Area"]; !ok {
      return nil 
   }
if _, ok := in["Virtuallink"]; !ok {
      return nil 
   }
if _, ok := in["TransmitDelay"]; !ok {
      return nil 
   }
      return sw.NoOSPFAreaVirtuallinkTransmitDelay(in["OSPF"], in["Area"], in["Virtuallink"], in["TransmitDelay"])
}

