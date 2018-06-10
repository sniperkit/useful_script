package n2x

import (
	"fmt"
	"strings"
)

type OSPF struct {
	Name          string
	Version       string
	Handler       string
	RID           string
	SutRID        string
	AreaID        string
	RouterHandler string
	Neighbour     *OSPFNeighbour
	HelloInterval string
	DeadInterval  string
	PollInterval  string
	TransmitDelay string
	Priority      string
	Cost          string
	MTU           string
	*Port
}

type OSPFNeighbour struct {
	Handler string
	AreaID  string
	RID     string
	State   string
	*Port
}

func (on *OSPFNeighbour) Sync() error {
	if on.Handler == "" {
		return fmt.Errorf("You must set the handler for this neighbor")
	}

	if on.Port == nil {
		return fmt.Errorf("You must set the port for this neighbor")
	}

	_, err := on.GetRouterID()
	if err != nil {
		return fmt.Errorf("Cannot sync neighbor with: %s", err)
	}

	_, err = on.GetIPAddress()
	if err != nil {
		return fmt.Errorf("Cannot sync neighbor with: %s", err)
	}

	return nil
}

func (on *OSPFNeighbour) GetRouterID() (string, error) {
	cmd := fmt.Sprintf("AgtOspfNeighbor GetRouterId %s", on.Handler)
	res, err := on.Invoke(cmd)
	if err != nil {
		return "", fmt.Errorf("Cannot get ospf neighbor rid %s : %s", on.Handler, err.Error())
	}

	on.RID = strings.TrimSpace(res)

	return on.RID, nil
}

func (on *OSPFNeighbour) SetRouterID(rid string) error {
	//AgtOspfNeighbor SetRouterId
	cmd := fmt.Sprintf("AgtOspfNeighbor SetRouterId %s %s", on.Handler, rid)
	res, err := on.Invoke(cmd)
	if err != nil {
		return fmt.Errorf("Cannot set ospf neighbor rid %s : %s", on.Handler, err.Error())
	}

	fmt.Println(res)

	return nil
}

func (on *OSPFNeighbour) GetIPAddress() (string, error) {
	cmd := fmt.Sprintf("AgtOspfNeighbor GetIpAddress %s", on.Handler)
	res, err := on.Invoke(cmd)
	if err != nil {
		return "", fmt.Errorf("Cannot get ospf neighbor ip %s : %s", on.Handler, err.Error())
	}

	return strings.TrimSpace(res), nil
}

func (on *OSPFNeighbour) SetIPAddress(ip string) error {
	cmd := fmt.Sprintf("AgtOspfNeighbor SetIpAddress %s %s", on.Handler, ip)
	_, err := on.Invoke(cmd)
	if err != nil {
		return fmt.Errorf("Cannot set ospf neighbor ip %s : %s", on.Handler, err.Error())
	}

	return nil
}

func (o *OSPF) Sync() error {
	if o.Handler == "" {
		return fmt.Errorf("You must set the OSPF session handler first")
	}

	_, err := o.GetRouterHandler()
	if err != nil {
		return fmt.Errorf("Cannot sync ospf with: %s", err)
	}

	name, err := o.GetName()
	if err != nil {
		return fmt.Errorf("Cannot sync ospf with: %s", err)
	}

	area, err := o.GetAreaID()
	if err != nil {
		return fmt.Errorf("Cannot sync ospf with: %s", err)
	}

	rid, err := o.GetRouterID()
	if err != nil {
		return fmt.Errorf("Cannot sync ospf with: %s", err)
	}

	n, err := o.GetNeighbour()
	if err != nil {
		return fmt.Errorf("Cannot sync ospf with: %s", err)
	}

	o.Name = name
	o.AreaID = area
	o.RID = rid
	o.Neighbour = n

	return nil
}

func (o *OSPF) SetName(name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return fmt.Errorf("OSPF Session name must not be empty")
	}

	cmd := fmt.Sprintf("AgtTestTopology SetSessionName %s %s", o.Handler, name)
	_, err := o.Invoke(cmd)
	if err != nil {
		return fmt.Errorf("Cannot set ospf session name for %s : %s", o.Handler, err.Error())
	}

	o.Name = name

	return nil
}

func (o *OSPF) GetName() (string, error) {
	cmd := fmt.Sprintf("AgtTestTopology GetSessionName %s", o.Handler)
	res, err := o.Invoke(cmd)
	if err != nil {
		return "", fmt.Errorf("Cannot get ospf session name for %s : %s", o.Handler, err.Error())
	}

	o.Name = strings.TrimSpace(res)

	return o.Name, nil
}

func (o *OSPF) Enable() error {
	cmd := fmt.Sprintf("AgtTestTopology EnableSession %s", o.Handler)
	_, err := o.Invoke(cmd)
	if err != nil {
		return fmt.Errorf("Cannot enable ospf session for %s : %s", o.Handler, err.Error())
	}

	return nil
}

func (o *OSPF) Disable() error {
	cmd := fmt.Sprintf("AgtTestTopology DisableSession %s", o.Handler)
	_, err := o.Invoke(cmd)
	if err != nil {
		return fmt.Errorf("Cannot disable ospf session for %s : %s", o.Handler, err.Error())
	}

	return nil
}

//AgtOspfSession GetSessionRouter
func (o *OSPF) GetRouterHandler() (string, error) {
	cmd := fmt.Sprintf("AgtOspfSession GetSessionRouter %s", o.Handler)
	res, err := o.Invoke(cmd)
	if err != nil {
		return "", fmt.Errorf("Cannot get ospf session router for %s : %s", o.Handler, err.Error())
	}

	o.RouterHandler = strings.TrimSpace(res)

	return o.RouterHandler, nil
}

func (o *OSPF) SetRouterID(id string) error {
	cmd := fmt.Sprintf("AgtOspfRouter SetRouterId %s %s", o.RouterHandler, id)
	_, err := o.Invoke(cmd)
	if err != nil {
		return fmt.Errorf("Cannot set ospf router id for %s(%s) : %s", o.Handler, o.RouterHandler, err.Error())
	}

	o.RID = id

	return nil
}

func (o *OSPF) GetRouterID() (string, error) {
	cmd := fmt.Sprintf("AgtOspfRouter GetRouterId %s", o.RouterHandler)
	res, err := o.Invoke(cmd)
	if err != nil {
		return "", fmt.Errorf("Cannot get ospf router id for %s : %s", o.Handler, err.Error())
	}

	o.RID = strings.TrimSpace(res)
	fmt.Println(res)

	return o.RID, nil
}

func (o *OSPF) SetAreaID(id string) error {
	cmd := fmt.Sprintf("AgtOspfSession SetAreaId %s %s", o.Handler, id)
	_, err := o.Invoke(cmd)
	if err != nil {
		return fmt.Errorf("Cannot set ospf area id for %s : %s", o.Handler, err.Error())
	}

	o.AreaID = id

	return nil
}

func (o *OSPF) GetAreaID() (string, error) {
	cmd := fmt.Sprintf("AgtOspfSession GetAreaId %s", o.Handler)
	res, err := o.Invoke(cmd)
	if err != nil {
		return "", fmt.Errorf("Cannot get ospf area id for %s : %s", o.Handler, err.Error())
	}

	o.AreaID = strings.TrimSpace(res)
	fmt.Println(res)

	return o.RID, nil
}

func (o *OSPF) GetNeighbour() (*OSPFNeighbour, error) {
	cmd := fmt.Sprintf("AgtOspfSession ListNeighbors %s", o.Handler)
	res, err := o.Invoke(cmd)
	if err != nil {
		return nil, fmt.Errorf("Cannot get ospf neighbors for %s : %s", o.Handler, err.Error())
	}

	res = strings.Replace(res, "}", "", -1)
	res = strings.Replace(res, "{", "", -1)
	res = strings.TrimSpace(res)

	o.Neighbour = nil

	fields := strings.Split(res, " ")
	for i, field := range fields {
		if i == 0 {
			if strings.TrimSpace(field) == "" {
				continue
			}

			n := &OSPFNeighbour{
				Handler: strings.TrimSpace(field),
				Port:    o.Port,
			}

			err := n.Sync()
			if err != nil {
				return nil, fmt.Errorf("Cannot get neighbor with: %s", err)
			}

			//Currently we only use the first one
			o.Neighbour = n
			break
		}
	}

	if o.Neighbour == nil {
		return nil, fmt.Errorf("Cannot find any ospf neighbor on: %s", o.Handler)
	}

	return o.Neighbour, nil
}

//AgtOspfSession AddNeighbor
func (o *OSPF) AddNeighbor() (*OSPFNeighbour, error) {
	cmd := fmt.Sprintf("AgtOspfSession AddNeighbor %s", o.Handler)
	res, err := o.Invoke(cmd)
	if err != nil {
		return nil, fmt.Errorf("Cannot get add neighbor for port %s : %s", o.Handler, err.Error())
	}

	o.Neighbour.Handler = strings.TrimSpace(res)

	return o.Neighbour, nil
}

func (o *OSPF) SetSutIPAddress(ip string) error {
	if o.Neighbour == nil {
		return fmt.Errorf("No neighbor exist for: %s", o.Handler)
	}

	return o.Neighbour.SetIPAddress(ip)
}

func (o *OSPF) SetSutRouterID(rid string) error {
	if o.Neighbour == nil {
		return fmt.Errorf("No neighbor exist for: %s", o.Handler)
	}

	fmt.Println("Setting: ", rid)
	return o.Neighbour.SetRouterID(rid)
}

//AGT_OSPF_ROUTER_PRIORITY,
//AGT_OSPF_HELLO_INTERVAL,
//AGT_OSPF_DEAD_INTERVAL,
//AGT_OSPF_POLL_INTERVAL,
//AGT_OSPF_LSA_RETRANSMIT_DELAY,
//AGT_OSPF_INTERFACE_COST,
//AGT_OSPF_TRANSIT_DELAY,
//AGT_OSPF_OPTIONS,
//AGT_OSPF_MTU,
//AGT_OSPF_INSTANCE_ID,
//AGT_OSPF_PASS_EVTS_TO_CLIENT,
//AGT_OSPF_GRACE_PERIOD,
//AGT_OSPF_RESTART_REASON,
//AGT_OSPF_RESTART_INTERVAL,
//AGT_OSPF_RESTART_MODE

func (o *OSPF) SetPriority(pri string) error {
	cmd := fmt.Sprintf("AgtOspfSession SetInterfaceParameter %s AGT_OSPF_ROUTER_PRIORITY %s", o.Handler, pri)
	_, err := o.Invoke(cmd)
	if err != nil {
		return fmt.Errorf("Cannot set priorty for %s : %s", o.Handler, err.Error())
	}

	o.Priority = strings.TrimSpace(pri)

	return nil
}

func (o *OSPF) SetHelloInterval(inter string) error {
	cmd := fmt.Sprintf("AgtOspfSession SetInterfaceParameter %s AGT_OSPF_HELLO_INTERVAL %s", o.Handler, inter)
	_, err := o.Invoke(cmd)
	if err != nil {
		return fmt.Errorf("Cannot set Hello interval for %s : %s", o.Handler, err.Error())
	}

	o.HelloInterval = strings.TrimSpace(inter)

	return nil
}

func (o *OSPF) SetDeadInterval(inter string) error {
	cmd := fmt.Sprintf("AgtOspfSession SetInterfaceParameter %s AGT_OSPF_DEAD_INTERVAL %s", o.Handler, inter)
	_, err := o.Invoke(cmd)
	if err != nil {
		return fmt.Errorf("Cannot set dead interval for %s : %s", o.Handler, err.Error())
	}

	o.DeadInterval = strings.TrimSpace(inter)

	return nil
}

func (o *OSPF) SetPollInterval(inter string) error {
	cmd := fmt.Sprintf("AgtOspfSession SetInterfaceParameter %s AGT_OSPF_POLL_INTERVAL %s", o.Handler, inter)
	_, err := o.Invoke(cmd)
	if err != nil {
		return fmt.Errorf("Cannot set poll interval for %s : %s", o.Handler, err.Error())
	}

	o.PollInterval = strings.TrimSpace(inter)

	return nil
}

func (o *OSPF) SetTransmitDelay(delay string) error {
	cmd := fmt.Sprintf("AgtOspfSession SetInterfaceParameter %s AGT_OSPF_TRANSIT_DELAY %s", o.Handler, delay)
	_, err := o.Invoke(cmd)
	if err != nil {
		return fmt.Errorf("Cannot set transmit delay for %s : %s", o.Handler, err.Error())
	}

	o.TransmitDelay = strings.TrimSpace(delay)

	return nil
}

func (o *OSPF) SetMTU(mtu string) error {
	cmd := fmt.Sprintf("AgtOspfSession SetInterfaceParameter %s AGT_OSPF_MTU %s", o.Handler, mtu)
	_, err := o.Invoke(cmd)
	if err != nil {
		return fmt.Errorf("Cannot set mtu for %s : %s", o.Handler, err.Error())
	}

	o.MTU = strings.TrimSpace(mtu)

	return nil
}

func (o *OSPF) SetInterfaceCost(cost string) error {
	cmd := fmt.Sprintf("AgtOspfSession SetInterfaceParameter %s AGT_OSPF_INTERFACE_COST %s", o.Handler, cost)
	_, err := o.Invoke(cmd)
	if err != nil {
		return fmt.Errorf("Cannot set cost for %s : %s", o.Handler, err.Error())
	}

	o.Cost = strings.TrimSpace(cost)

	return nil
}

func (o *OSPF) GetPriority() (string, error) {
	cmd := fmt.Sprintf("AgtOspfSession GetInterfaceParameter %s AGT_OSPF_ROUTER_PRIORITY", o.Handler)
	res, err := o.Invoke(cmd)
	if err != nil {
		return "", fmt.Errorf("Cannot get priorty for %s : %s", o.Handler, err.Error())
	}

	o.Priority = strings.TrimSpace(res)

	return o.Priority, nil
}

func (o *OSPF) GetHelloInterval() (string, error) {
	cmd := fmt.Sprintf("AgtOspfSession GetInterfaceParameter %s AGT_OSPF_HELLO_INTERVAL", o.Handler)
	res, err := o.Invoke(cmd)
	if err != nil {
		return "", fmt.Errorf("Cannot get Hello interval for %s : %s", o.Handler, err.Error())
	}

	o.HelloInterval = strings.TrimSpace(res)

	return o.HelloInterval, nil
}

func (o *OSPF) GetDeadInterval() (string, error) {
	cmd := fmt.Sprintf("AgtOspfSession GetInterfaceParameter %s AGT_OSPF_DEAD_INTERVAL", o.Handler)
	res, err := o.Invoke(cmd)
	if err != nil {
		return "", fmt.Errorf("Cannot get dead interval for %s : %s", o.Handler, err.Error())
	}

	o.DeadInterval = strings.TrimSpace(res)

	return o.DeadInterval, nil
}

func (o *OSPF) GetPollInterval() (string, error) {
	cmd := fmt.Sprintf("AgtOspfSession GetInterfaceParameter %s AGT_OSPF_POLL_INTERVAL", o.Handler)
	res, err := o.Invoke(cmd)
	if err != nil {
		return "", fmt.Errorf("Cannot get poll interval for %s : %s", o.Handler, err.Error())
	}

	o.PollInterval = strings.TrimSpace(res)

	return o.PollInterval, nil
}

func (o *OSPF) GetTransmitDelay() (string, error) {
	cmd := fmt.Sprintf("AgtOspfSession GetInterfaceParameter %s AGT_OSPF_TRANSIT_DELAY", o.Handler)
	res, err := o.Invoke(cmd)
	if err != nil {
		return "", fmt.Errorf("Cannot get transmit delay for %s : %s", o.Handler, err.Error())
	}

	o.TransmitDelay = strings.TrimSpace(res)

	return o.TransmitDelay, nil
}

func (o *OSPF) GetMTU() (string, error) {
	cmd := fmt.Sprintf("AgtOspfSession GetInterfaceParameter %s AGT_OSPF_MTU", o.Handler)
	res, err := o.Invoke(cmd)
	if err != nil {
		return "", fmt.Errorf("Cannot get mtu for %s : %s", o.Handler, err.Error())
	}

	o.MTU = strings.TrimSpace(res)

	return o.MTU, nil
}

func (o *OSPF) GetInterfaceCost() (string, error) {
	cmd := fmt.Sprintf("AgtOspfSession GetInterfaceParameter %s AGT_OSPF_INTERFACE_COST", o.Handler)
	res, err := o.Invoke(cmd)
	if err != nil {
		return "", fmt.Errorf("Cannot get cost for %s : %s", o.Handler, err.Error())
	}

	o.Cost = strings.TrimSpace(res)

	return o.Cost, nil
}
