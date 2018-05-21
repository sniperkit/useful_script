!
hostname H_JOAM_V8500-1_TPS
!
service password-encryption
!
!
time-zone GMT+9
!
!
ntp 180.225.21.146 180.225.21.147 
ntp bind-address 172.17.48.243
!
syslog index epon 1-2 priority crit
syslog index epon 3 priority err
syslog index loop-detect 1-9 priority crit
syslog output info local volatile
syslog output info local non-volatile
syslog output warning remote 58.78.0.247
syslog output info remote 203.248.244.172
syslog local-code 7
syslog bind-address 172.17.48.243
!
debug packet log 500 50 60 10
!
slot planning siu 1 iu-gepon8 
slot planning siu 2 iu-gepon8 
slot planning siu 3 iu-gepon8 
slot planning siu 4 iu-gepon8 
slot planning siu 5 iu-gepon8 
slot planning siu 6 iu-gepon8 
slot planning siu 7 iu-gepon8 
slot planning siu 8 iu-gepon8 
slot planning siu 9 iu-gepon8 
slot planning siu 11 iu-10ge8 
slot planning siu 12 iu-10ge8 
slot unlock siu 1 
slot unlock siu 2 
slot unlock siu 3 
slot unlock siu 4 
slot unlock siu 5 
slot unlock siu 6 
slot unlock siu 7 
slot unlock siu 8 
slot unlock siu 9 
slot lock siu 10 
slot unlock siu 11 
slot unlock siu 12 
!
!
ip multicast-routing
!
bgp extended-asn-cap
!
fault-monitor memory threshold action switchover
!
ip unknown-multicast block
!
no ip multicast mrouter-pass-through
!
ip ecmp-hash sip-dip
!
service dhcp
!
fan operation full
!
cpu-storm arp rate 1024 
!
ip dhcp snooping
ip dhcp snooping arp-inspection start 21600
!
ip dhcp option82
 trust default permit
 system-circuit-id port-type physical
 system-remote-id option format remote-default
 dhcp-opt254-append enable
 exit
!
vlan database
 vlan 11  name EG1
 vlan 12  name EG2
 vlan 101  name 140779818_joamli_189_area
 vlan 104  name 140782157_joamli_483_area
 vlan 106  name 140782834_dokjeong_119_area
 vlan 107  name 140782830_eoeunli_736_area
 vlan 108  name 140783630_joamli_366_area
 vlan 181  name 1100295441_jayeon_A_dreamCS
 vlan 182  name 1100295434_jayeon_B_dreamCS
 vlan 202  name F/140784248/Joamri270-65area
 vlan 203  name F/140784878/sekpoli_395_area
 vlan 204  name F/140784879/eoeunli_15_area
 vlan 302  name F/140790715/joamri_311_1_area
 vlan 2001  name MSE1
 vlan 2002  name MSE2
 vlan 11-12,101,104,106-108,181-182,202-204,302,2001-2002,3000 state enable
 vlan description 101 140779818_joamli_189_area
 vlan description 104 140782157_joamli_483_area
 vlan description 106 140782834_dokjeong_119_area
 vlan description 107 140782830_eoeunli_736_area
 vlan description 108 F/140784729/Joamri_460
 vlan description 181 1100295441_jayeon_A_dreamCS
 vlan description 182 1100295434_jayeon_B_dreamCS
 vlan description 203 F/140784878/sekpoli_395_area
 vlan description 204 F/140784879/eoeunli_15_area
 vlan description 302 F/140790715/joamri_311_1_area
 exit
!
flow SQL_Over_UDP_1433_Dest create 
  ip any any udp any 1433
  apply
exit
flow SQL_Over_UDP_1434_Dest create 
  ip any any udp any 1434
  apply
exit
flow Blaster_Worm_TCP_4444_Dest create 
  ip any any tcp any 4444
  apply
exit
flow WMP create 
  ip any any tcp any 2869
  apply
exit
flow Win32_Welchia_worm create 
  ip any any tcp any 707
  apply
exit
flow Netbios_TCP_135_Dest create 
  ip any any tcp any 135
  apply
exit
flow Netbios_UDP_135_Dest create 
  ip any any udp any 135
  apply
exit
flow Worm_TCP_137_Dest create 
  ip any any tcp any 137
  apply
exit
flow Worm_UDP_137_Dest create 
  ip any any udp any 137
  apply
exit
flow Worm_UDP_138_Dest create 
  ip any any udp any 138
  apply
exit
flow Worm_TCP_139_Dest create 
  ip any any tcp any 139
  apply
exit
flow Worm_TCP_445_Dest create 
  ip any any tcp any 445
  apply
exit
flow Worm_UDP_445_Dest create 
  ip any any udp any 445
  apply
exit
flow UPnP create 
  ip any any udp any 1900
  apply
exit
flow Worm_traffic create 
  ip any any udp any 57321
  apply
exit
flow src_0 create 
  mac 00:00:00:00:00:00 any
  ip any any
  apply
exit
flow src_broadcast create 
  mac FF:FF:FF:FF:FF:FF any
  ip any any
  apply
exit
flow Netbios_UDP_3702 create 
  ip any any udp any 3702
  apply
exit
flow 88D9 create 
  ethtype 0x88D9
  apply
exit
flow 86DD create 
  ipv6 any any
  apply
exit
flow dst_67 create 
  ip any any udp any 67
  apply
exit
flow Mcast_deny create 
  ip any 224.0.0.0/4 udp any any
  apply
exit
flow voipqos create 
  ip any any
    dscp 46
  apply
exit
flow VOD create 
  ip any any
    dscp 34
  apply
exit
flow BGP create 
  ip any any tcp any 179
  apply
exit
flow User_Telnet_Permit1 create 
  ip 58.78.0.0/24 any tcp any 23
  apply
exit
flow User_Telnet_Permit2 create 
  ip 211.63.37.0/24 any tcp any 23
  apply
exit
flow User_Telnet_Permit3 create 
  ip 211.63.46.0/26 any tcp any 23
  apply
exit
flow User_Telnet_Permit4 create 
  ip 210.182.142.0/24 any tcp any 23
  apply
exit
flow User_Telnet_Permit5 create 
  ip 210.124.6.216/29 any tcp any 23
  apply
exit
flow User_Telnet_Permit6 create 
  ip 61.33.188.192/27 any tcp any 23
  apply
exit
flow User_Telnet_Permit7 create 
  ip 203.248.244.0/24 any tcp any 23
  apply
exit
flow User_Telnet_Permit8 create 
  ip 58.184.88.0/24 any tcp any 23
  apply
exit
flow User_Telnet_Permit9 create 
  ip 211.241.120.0/24 any tcp any 23
  apply
exit
flow User_Telnet_Permit10 create 
  ip 115.141.255.0/24 any tcp any 23
  apply
exit
flow User_Telnet_Permit11 create 
  ip 164.124.250.0/24 any tcp any 23
  apply
exit
flow User_Telnet_Permit12 create 
  ip 164.124.248.0/24 any tcp any 23
  apply
exit
flow User_Telnet_Permit13 create 
  ip 122.35.0.0/24 any tcp any 23
  apply
exit
flow User_Telnet_Permit14 create 
  ip 115.141.2.0/24 any tcp any 23
  apply
exit
flow User_Telnet_Permit15 create 
  ip 203.252.0.0/24 any tcp any 23
  apply
exit
flow User_Telnet_Permit16 create 
  ip 115.141.234.224/27 any tcp any 23
  apply
exit
flow User_Telnet_Permit17 create 
  ip 180.225.21.0/24 any tcp any 23
  apply
exit
flow User_Telnet_Permit18 create 
  ip 123.140.16.0/23 any tcp any 23
  apply
exit
flow User_Telnet_Deny create 
  ip any any tcp any 23
  apply
exit
flow Netbios_TCP_138_Dest create 
  ip any any tcp any 138
  apply
exit
flow Netbios_UDP_139_Dest create 
  ip any any udp any 139
  apply
exit
flow admin telnet_permit1 create
  ip 58.78.0.0/24 any tcp any 23
  apply
exit
flow admin telnet_permit2 create
  ip 203.248.244.0/24 any tcp any 23
  apply
exit
flow admin telnet_permit3 create
  ip 211.63.37.224/27 any tcp any 23
  apply
exit
flow admin telnet_permit4 create
  ip 211.63.37.200/29 any tcp any 23
  apply
exit
flow admin telnet_permit5 create
  ip 211.63.46.0/26 any tcp any 23
  apply
exit
flow admin telnet_permit6 create
  ip 203.248.240.0/24 any tcp any 23
  apply
exit
flow admin telnet_permit7 create
  ip 115.141.255.0/24 any tcp any 23
  apply
exit
flow admin telnet_permit8 create
  ip 164.124.250.0/24 any tcp any 23
  apply
exit
flow admin telnet_permit9 create
  ip 164.124.248.0/24 any tcp any 23
  apply
exit
flow admin telnet_permit10 create
  ip 210.124.6.216/29 any tcp any 23
  apply
exit
flow admin telnet_permit11 create
  ip 61.33.188.192/27 any tcp any 23
  apply
exit
flow admin telnet_permit12 create
  ip 210.182.142.0/24 any tcp any 23
  apply
exit
flow admin telnet_permit13 create
  ip 122.35.0.0/24 any tcp any 23
  apply
exit
flow admin telnet_permit14 create
  ip 115.141.2.0/24 any tcp any 23
  apply
exit
flow admin snmp_permit1 create
  ip 58.78.0.0/24 any udp any 161
  apply
exit
flow admin snmp_permit2 create
  ip 211.63.37.224/27 any udp any 161
  apply
exit
flow admin snmp_permit3 create
  ip 211.63.37.200/29 any udp any 161
  apply
exit
flow admin snmp_permit4 create
  ip 211.63.46.0/26 any udp any 161
  apply
exit
flow admin snmp_permit5 create
  ip 203.248.244.0/24 any udp any 161
  apply
exit
flow admin snmp_permit6 create
  ip 203.248.240.0/24 any udp any 161
  apply
exit
flow admin snmp_permit7 create
  ip 164.124.250.0/24 any udp any 161
  apply
exit
flow admin snmp_permit8 create
  ip 164.124.248.0/24 any udp any 161
  apply
exit
flow admin snmp_permit9 create
  ip 122.35.0.0/24 any udp any 161
  apply
exit
flow admin snmp_permit10 create
  ip 115.141.2.0/24 any udp any 161
  apply
exit
flow admin snmp_permit11 create
  ip 115.141.234.224/27 any udp any 161
  apply
exit
flow admin snmp_deny create
  ip any any udp any 161
  apply
exit
flow admin telnet_deny create
  ip any any tcp any 23
  apply
exit
flow admin ntp_permit1 create
  ip 180.225.21.146/32 any
  apply
exit
flow admin ntp_permit2 create
  ip 180.225.21.147/32 any
  apply
exit
flow admin ntp_deny create
  ip any any udp any 123
  apply
exit
flow admin ftp_deny create
  ip any any tcp any 20
  apply
exit
flow admin tftp_deny create
  ip any any udp any 69
  apply
exit
flow admin ftp_data_deny create
  ip any any tcp any 21
  apply
exit
flow admin ssh_deny create
  ip any any tcp any 22
  apply
exit
!
class deny_medium flow SQL_Over_UDP_1433_Dest SQL_Over_UDP_1434_Dest
class deny_medium flow Blaster_Worm_TCP_4444_Dest WMP Win32_Welchia_worm
class deny_medium flow Netbios_TCP_135_Dest Netbios_UDP_135_Dest
class deny_medium flow Worm_TCP_137_Dest Worm_UDP_137_Dest Worm_UDP_138_Dest
class deny_medium flow Worm_TCP_139_Dest Worm_TCP_445_Dest Worm_UDP_445_Dest
class deny_medium flow UPnP Worm_traffic src_0 src_broadcast
class deny_high flow Netbios_UDP_3702 88D9
class deny_highest flow Mcast_deny
class User_Telnet_Permit flow User_Telnet_Permit1 User_Telnet_Permit2
class User_Telnet_Permit flow User_Telnet_Permit3 User_Telnet_Permit4
class User_Telnet_Permit flow User_Telnet_Permit5 User_Telnet_Permit6
class User_Telnet_Permit flow User_Telnet_Permit7 User_Telnet_Permit8
class User_Telnet_Permit flow User_Telnet_Permit9 User_Telnet_Permit10
class User_Telnet_Permit flow User_Telnet_Permit11 User_Telnet_Permit12
class User_Telnet_Permit flow User_Telnet_Permit13 User_Telnet_Permit14
class User_Telnet_Permit flow User_Telnet_Permit15 User_Telnet_Permit16
class User_Telnet_Permit flow User_Telnet_Permit17 User_Telnet_Permit18
class uplink_deny_medium flow Netbios_TCP_135_Dest Netbios_UDP_135_Dest
class uplink_deny_medium flow Netbios_TCP_138_Dest Netbios_UDP_139_Dest
class admin telnet_permit flow telnet_permit1 telnet_permit2 telnet_permit3 telnet_permit4 telnet_permit5 telnet_permit6 telnet_permit7 telnet_permit8 telnet_permit9 telnet_permit10 telnet_permit11 telnet_permit12
class admin telnet_permit flow telnet_permit13 telnet_permit14
class admin snmp_permit flow snmp_permit1 snmp_permit2 snmp_permit3 snmp_permit4 snmp_permit5 snmp_permit6 snmp_permit7 snmp_permit8 snmp_permit9 snmp_permit10 snmp_permit11
class admin snmp_deny flow snmp_deny
!
policy deny_medium create 
  include-class deny_medium
  priority medium
  action match deny
  apply
exit
policy deny_highest create 
  include-class deny_highest
  priority highest
  action match deny
  apply
exit
policy 86DD create 
  include-flow 86DD
  priority high-middle
  action match deny
  apply
exit
policy voipqos create 
  include-flow voipqos
  priority high
  action match permit
  action match cos 4 overwrite
  apply
exit
policy VOD create 
  include-flow VOD
  priority high
  action match permit
  action match cos 3 overwrite
  apply
exit
policy BGP create 
  include-flow BGP
  priority high
  action match permit
  action match cos 7
  apply
exit
policy dst_67 create 
  include-flow dst_67
  priority highest
  action match deny
  action match copy-to-cpu
  apply
exit
policy User_Telnet_Permit create 
  include-class User_Telnet_Permit
  priority medium
  action match permit
  apply
exit
policy User_Telnet_Deny create 
  include-flow User_Telnet_Deny
  priority low
  action match deny
  apply
exit
policy deny_high create 
  include-class deny_high
  priority high
  action match deny
  apply
exit
policy uplink_deny_medium create 
  include-class uplink_deny_medium
  priority medium
  action match deny
  apply
exit
policy admin snmp_permit create
  include-class snmp_permit
  priority medium
  action match permit
  apply
exit
policy admin telnet_permit create
  include-class telnet_permit
  priority medium
  action match permit
  apply
exit
policy admin snmp_deny create
  include-class snmp_deny
  priority low
  action match deny
  apply
exit
policy admin ntp_permit1 create
  include-flow ntp_permit1
  priority medium
  action match permit
  apply
exit
policy admin ntp_permit2 create
  include-flow ntp_permit2
  priority medium
  action match permit
  apply
exit
policy admin ntp_deny create
  include-flow ntp_deny
  priority low
  action match deny
  apply
exit
policy admin ftp_deny create
  include-flow ftp_deny
  priority low
  action match deny
  apply
exit
policy admin tftp_deny create
  include-flow tftp_deny
  priority low
  action match deny
  apply
exit
policy admin ftp_data_deny create
  include-flow ftp_data_deny
  priority low
  action match deny
  apply
exit
policy admin ssh_deny create
  include-flow ssh_deny
  priority low
  action match deny
  apply
exit
!
epon olt redundancy standby olt rxpower alarm -40.0000 -30.0000
epon onu auto-upgrade enable
epon onu auto-upgrade reboot-time 4 to 5
!
eonu-rule flow 
 flow 1 dst-ip 224.0.0.1 
 flow 2 dst-ip 239.255.255.250 
 exit
!
eonu-rule action 
 action 1 deny 
 exit
!
eonu-rule all_deny create
 policy 1 upstream priority high flow 1 action 1 
 policy 2 upstream priority high flow 2 action 1 
 apply
exit
!
eonu-dba-profile FTTH_105M create
 link-sla max 105000 256 2 4
 queue-sla 0 upstream max 105000 256 2 16
 queue-sla 0 downstream max 105000 256 2 16
 apply
exit
!
eonu-dba-profile FTTH_1G create
 link-sla max 1000000 256 2 16
 queue-sla 0 upstream max 1000000 256 2 16
 queue-sla 0 downstream max 1000000 256 2 16
 apply
exit
!
eonu-dba-profile FTTH_525M create
 link-sla max 525000 256 2 4
 queue-sla 0 upstream max 525000 256 2 16
 queue-sla 0 downstream max 525000 256 2 16
 apply
exit
!
eonu-dba-profile HYBRIDONU2B create
 sr-report-interval high 
 link-sla min 10000 256 1 16
 link-sla max 1000000 256 2 16
 queue-sla 0 upstream max 1000000 256 2 16
 queue-sla 0 downstream max 1000000 256 2 16
 apply
exit
!
eonu-profile FTTH create
 dba-profile FTTH_105M
 max-host port 1-4 2
 vlan pvid uplink 1
 vlan pvid 1-4 1
 apply
exit
!
eonu-profile FTTH-1G create
 dba-profile FTTH_1G
 max-host port 1-4 3
 vlan pvid uplink 1
 vlan pvid 1-4 1
 apply
exit
!
eonu-profile FTTH-500 create
 dba-profile FTTH_525M
 max-host port 1-4 3
 vlan pvid uplink 1
 vlan pvid 1-4 1
 apply
exit
!
eonu-profile HYBRIDONU2A-default create
 dba-profile FTTH_1G
 vlan pvid uplink 1
 vlan pvid 1-4 1
 apply
exit
!
eonu-profile HYBRIDONU2B-default create
 dba-profile HYBRIDONU2B
 igmp-snooping disable
 vlan pvid uplink 1
 vlan pvid 1-4 1
 apply
exit
!
eonu-profile P6301-default create
 dba-profile FTTH_1G
 vlan pvid uplink 1
 vlan pvid 1-4 1
 apply
exit
!
eonu-profile P6302-default create
 dba-profile FTTH_1G
 vlan pvid uplink 1
 vlan pvid 1-4 1
 apply
exit
!
epon olt service-profile default FTTH
epon olt service-profile model-name HYBRIDONU2A HYBRIDONU2A-default
epon olt service-profile model-name HYBRIDONU2B HYBRIDONU2B-default
epon olt service-profile model-name P6301 P6301-default
epon olt service-profile model-name P6302 P6302-default
!
interface lo
 no shutdown
 ip address 172.17.48.243/32
 exit
!
!
interface epon 1/1
 pvid 101
 no shutdown
 switchport mode access
 switchport access vlan 101
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 onu add 1 00:d0:cb:d7:51:d9 
 onu add 2 00:d0:cb:9b:d0:60 
 onu add 3 00:d0:cb:fe:fd:d7 
 eonu-profile 1-2 FTTH
 eonu-profile 3 FTTH-500
 module ddm enable
 exit
!
interface epon 1/2
 pvid 101
 no shutdown
 switchport mode access
 switchport access vlan 101
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 onu add 1 00:d0:cb:fa:2b:c6 
 onu add 2 00:d0:cb:d7:53:51 
 eonu-profile 1-2 FTTH
 module ddm enable
 exit
!
interface epon 1/3
 pvid 101
 no shutdown
 switchport mode access
 switchport access vlan 101
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 onu add 1 00:d0:cb:f6:0d:b4 
 onu add 3 00:d0:cb:9d:71:0a 
 eonu-profile 3 FTTH
 eonu-profile 1 FTTH-500
 module ddm enable
 exit
!
interface epon 1/4
 pvid 104
 no shutdown
 switchport mode access
 switchport access vlan 104
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 module ddm enable
 exit
!
interface epon 1/5
 pvid 104
 no shutdown
 switchport mode access
 switchport access vlan 104
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 onu add 1 00:d0:cb:d7:5a:8d 
 onu add 2 00:d0:cb:fb:5b:85 
 onu add 3 00:d0:cb:f6:1a:0d 
 onu add 4 00:d0:cb:bb:f6:d3 
 onu add 5 00:d0:cb:fd:9f:1a 
 eonu-profile 2,4-5 FTTH
 eonu-profile 1,3 FTTH-500
 module ddm enable
 exit
!
interface epon 1/6
 pvid 106
 no shutdown
 switchport mode access
 switchport access vlan 106
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 onu add 1 00:d0:cb:fe:b8:37 
 eonu-profile 1 FTTH-500
 module ddm enable
 exit
!
interface epon 1/7
 pvid 107
 no shutdown
 switchport mode access
 switchport access vlan 107
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 module ddm enable
 exit
!
interface epon 1/8
 pvid 108
 no shutdown
 switchport mode trunk
 switchport trunk allowed vlan remove 1
 switchport trunk native vlan 108
 switchport trunk allowed vlan add 108,181-182,3000
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 onu add 1 00:d0:cb:99:e1:58 
 onu add 2 00:d0:cb:99:e1:57 
 onu add 3 00:d0:cb:fd:39:41 
 onu add 4 00:d0:cb:bb:f9:c0 
 eonu-profile 3-4 FTTH-500
 eonu-profile 1-2 P6302-default
 exit
!
interface epon 2/1
 pvid 108
 no shutdown
 switchport mode trunk
 switchport trunk allowed vlan remove 1
 switchport trunk native vlan 108
 switchport trunk allowed vlan add 108
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 module ddm enable
 exit
!
interface epon 2/2
 pvid 202
 no shutdown
 switchport mode access
 switchport access vlan 202
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 exit
!
interface epon 2/3
 pvid 203
 no shutdown
 switchport mode access
 switchport access vlan 203
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 exit
!
interface epon 2/4
 pvid 204
 no shutdown
 switchport mode access
 switchport access vlan 204
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 onu add 1 00:d0:cb:9f:9c:b1 
 onu add 2 00:d0:cb:bb:d7:80 
 eonu-profile 1 FTTH
 eonu-profile 2 FTTH-500
 exit
!
interface epon 2/5
 pvid 204
 no shutdown
 switchport mode access
 switchport access vlan 204
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 onu add 1 00:d0:cb:9f:aa:96 
 onu add 2 00:d0:cb:d7:61:ee 
 eonu-profile 1-2 FTTH
 exit
!
interface epon 2/6
 pvid 204
 no shutdown
 switchport mode access
 switchport access vlan 204
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 onu add 1 00:d0:cb:dd:f6:37 
 onu add 2 00:d0:cb:9f:96:10 
 onu add 3 00:d0:cb:d7:59:ac 
 onu add 4 00:d0:cb:d7:59:f0 
 eonu-profile 2,4 FTTH
 eonu-profile 1,3 FTTH-500
 exit
!
interface epon 2/7
 pvid 204
 no shutdown
 switchport mode access
 switchport access vlan 204
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 onu add 2 00:d0:cb:d7:59:f1 
 onu add 3 00:d0:cb:d7:62:02 
 eonu-profile 3 FTTH
 eonu-profile 2 FTTH-500
 exit
!
interface epon 2/8
 pvid 204
 no shutdown
 switchport mode access
 switchport access vlan 204
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 exit
!
interface epon 3/1
 pvid 204
 no shutdown
 switchport mode access
 switchport access vlan 204
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 exit
!
interface epon 3/2
 pvid 302
 no shutdown
 switchport mode trunk
 switchport trunk allowed vlan remove 1
 switchport trunk native vlan 302
 switchport trunk allowed vlan add 302,3000
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 exit
!
interface epon 3/3
 pvid 302
 no shutdown
 switchport mode trunk
 switchport trunk allowed vlan remove 1
 switchport trunk native vlan 302
 switchport trunk allowed vlan add 302,3000
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 exit
!
interface epon 3/4
 pvid 302
 no shutdown
 switchport mode trunk
 switchport trunk allowed vlan remove 1
 switchport trunk native vlan 302
 switchport trunk allowed vlan add 302,3000
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 exit
!
interface epon 3/5
 pvid 302
 no shutdown
 switchport mode trunk
 switchport trunk allowed vlan remove 1
 switchport trunk native vlan 302
 switchport trunk allowed vlan add 302,3000
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 exit
!
interface epon 3/6
 pvid 302
 no shutdown
 switchport mode trunk
 switchport trunk allowed vlan remove 1
 switchport trunk native vlan 302
 switchport trunk allowed vlan add 302,3000
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 exit
!
interface epon 3/7
 pvid 302
 no shutdown
 switchport mode trunk
 switchport trunk allowed vlan remove 1
 switchport trunk native vlan 302
 switchport trunk allowed vlan add 302,3000
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 exit
!
interface epon 3/8
 pvid 302
 no shutdown
 switchport mode trunk
 switchport trunk allowed vlan remove 1
 switchport trunk native vlan 302
 switchport trunk allowed vlan add 302,3000
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 exit
!
interface epon 4/1
 pvid 302
 no shutdown
 switchport mode trunk
 switchport trunk allowed vlan remove 1
 switchport trunk native vlan 302
 switchport trunk allowed vlan add 302,3000
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 exit
!
interface epon 4/2
 pvid 302
 no shutdown
 switchport mode trunk
 switchport trunk allowed vlan remove 1
 switchport trunk native vlan 302
 switchport trunk allowed vlan add 302,3000
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 exit
!
interface epon 4/3
 pvid 302
 no shutdown
 switchport mode trunk
 switchport trunk allowed vlan remove 1
 switchport trunk native vlan 302
 switchport trunk allowed vlan add 302,3000
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 exit
!
interface epon 4/4
 pvid 302
 no shutdown
 switchport mode trunk
 switchport trunk allowed vlan remove 1
 switchport trunk native vlan 302
 switchport trunk allowed vlan add 302,3000
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 exit
!
interface epon 4/5
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 exit
!
interface epon 4/6
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 exit
!
interface epon 4/7
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 exit
!
interface epon 4/8
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 exit
!
interface epon 5/1
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 exit
!
interface epon 5/2
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 exit
!
interface epon 5/3
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 exit
!
interface epon 5/4
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 exit
!
interface epon 5/5
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 exit
!
interface epon 5/6
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 exit
!
interface epon 5/7
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 exit
!
interface epon 5/8
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 exit
!
interface epon 6/1
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 exit
!
interface epon 6/2
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 exit
!
interface epon 6/3
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
 exit
!
interface epon 6/4
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
exit
!
interface epon 6/5
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
exit
!
interface epon 6/6
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
exit
!
interface epon 6/7
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
exit
!
interface epon 6/8
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
exit
!
interface epon 7/1
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
exit
!
interface epon 7/2
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
exit
!
interface epon 7/3
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
exit
!
interface epon 7/4
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
exit
!
interface epon 7/5
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
exit
!
interface epon 7/6
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
exit
!
interface epon 7/7
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
exit
!
interface epon 7/8
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
exit
!
interface epon 8/1
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
exit
!
interface epon 8/2
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
exit
!
interface epon 8/3
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
exit
!
interface epon 8/4
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
exit
!
interface epon 8/5
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
exit
!
interface epon 8/6
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
exit
!
interface epon 8/7
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
exit
!
interface epon 8/8
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
exit
!
interface epon 9/1
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
exit
!
interface epon 9/2
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
exit
!
interface epon 9/3
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
exit
!
interface epon 9/4
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
exit
!
interface epon 9/5
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
exit
!
interface epon 9/6
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
exit
!
interface epon 9/7
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
exit
!
interface epon 9/8
 no shutdown
 switchport mode access
 storm-control broadcast level 2040
 storm-control multicast level 2040
 storm-control dlf level 2040
 mac-flood-guard 100
 service-policy input deny_medium
 service-policy input deny_highest
 service-policy input 86DD
 service-policy input voipqos
 service-policy input VOD
 service-policy input dst_67
 service-policy input deny_high
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 dhcp-server-filter on
 system-circuit-id option format circuit-default
 olt register-mode auto-to-manual
 olt dhcp option254 enable
 onu inactive aging-time 15
 olt signal-check enable auto-onu-block
 alarm loss-of-signal disable
exit
!
interface tengigabitethernet 11/1
 description [TRK] JOAM_E7508-EG1 Te1/2/2
 pvid 11
 no shutdown
 switchport mode access
 switchport access vlan 11
 service-policy input voipqos
 service-policy input VOD
 service-policy input BGP
 service-policy input User_Telnet_Permit
 service-policy input User_Telnet_Deny
 service-policy input uplink_deny_medium
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 no system-circuit-id option format
exit 
!
interface tengigabitethernet 11/2
 description [MSE-T] JOAM_V6744-MSE1 Te37
 no shutdown
 switchport mode access
 channel-group lacp 1 mode active
 channel-group timeout short
 service-policy input voipqos
 service-policy input VOD
 service-policy input BGP
 service-policy input User_Telnet_Permit
 service-policy input User_Telnet_Deny
 service-policy input uplink_deny_medium
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 no system-circuit-id option format
exit 
!
interface tengigabitethernet 11/3
 no shutdown
 switchport mode access
 service-policy input uplink_deny_medium
 netbios-filter on 
 no system-circuit-id option format
exit 
!
interface tengigabitethernet 11/4
 no shutdown
 switchport mode access
 service-policy input uplink_deny_medium
 netbios-filter on 
 no system-circuit-id option format
exit 
!
interface tengigabitethernet 11/5
 no shutdown
 switchport mode access
 service-policy input uplink_deny_medium
 netbios-filter on 
 no system-circuit-id option format
exit 
!
interface tengigabitethernet 11/6
 no shutdown
 switchport mode access
 service-policy input uplink_deny_medium
 netbios-filter on 
 no system-circuit-id option format
exit 
!
interface tengigabitethernet 11/7
 no shutdown
 switchport mode access
 service-policy input uplink_deny_medium
 netbios-filter on 
 no system-circuit-id option format
exit 
!
interface tengigabitethernet 11/8
 no shutdown
 switchport mode access
 service-policy input uplink_deny_medium
 netbios-filter on 
 no system-circuit-id option format
exit 
!
interface tengigabitethernet 12/1
 description [TRK] JOAM_E7508-EG2 Te1/2/2
 pvid 12
 no shutdown
 switchport mode access
 switchport access vlan 12
 service-policy input voipqos
 service-policy input VOD
 service-policy input BGP
 service-policy input User_Telnet_Permit
 service-policy input User_Telnet_Deny
 service-policy input uplink_deny_medium
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 no system-circuit-id option format
exit 
!
interface tengigabitethernet 12/2
 description [MSE-T] JOAM_V6744-MSE2 Te37
 no shutdown
 switchport mode access
 channel-group lacp 2 mode active
 channel-group timeout short
 service-policy input voipqos
 service-policy input VOD
 service-policy input BGP
 service-policy input User_Telnet_Permit
 service-policy input User_Telnet_Deny
 service-policy input uplink_deny_medium
 qos scheduling-mode sp 
 netbios-filter on 
 rmon-simple crc-align-error 600 10 10
 no system-circuit-id option format
exit 
!
interface tengigabitethernet 12/3
 no shutdown
 switchport mode access
 service-policy input uplink_deny_medium
 netbios-filter on 
 no system-circuit-id option format
exit 
!
interface tengigabitethernet 12/4
 no shutdown
 switchport mode access
 service-policy input uplink_deny_medium
 netbios-filter on 
 no system-circuit-id option format
exit 
!
interface tengigabitethernet 12/5
 no shutdown
 switchport mode access
 service-policy input uplink_deny_medium
 netbios-filter on 
 no system-circuit-id option format
exit 
!
interface tengigabitethernet 12/6
 no shutdown
 switchport mode access
 service-policy input uplink_deny_medium
 netbios-filter on 
 no system-circuit-id option format
exit 
!
interface tengigabitethernet 12/7
 no shutdown
 switchport mode access
 service-policy input uplink_deny_medium
 netbios-filter on 
 no system-circuit-id option format
exit 
!
interface tengigabitethernet 12/8
 no shutdown
 switchport mode access
 service-policy input uplink_deny_medium
 netbios-filter on 
 no system-circuit-id option format
exit 
!
interface channelgroup 1
 description [MSE-T] JOAM_V6744-MSE1 Te37
 pvid 2001
 no shutdown
 switchport mode access
 switchport access vlan 2001
 channel-group distmode srcdstip
 no system-circuit-id option format
exit 
!
interface channelgroup 2
 description [MSE-T] JOAM_V6744-MSE2 Te37
 pvid 2002
 no shutdown
 switchport mode access
 switchport access vlan 2002
 channel-group distmode srcdstip
 system-circuit-id option format circuit-default
 exit
!
interface vlan 11
 no shutdown
 description [TRK] JOAM_E7508-EG1 Te1/2/2
 ip address 10.203.237.74/30
 exit
!
interface vlan 12
 no shutdown
 description [TRK] JOAM_E7508-EG2 Te1/2/2
 ip address 10.203.238.74/30
 exit
!
interface vlan 101
 no shutdown
 description 140779818_joamli_189_area
 ip address 112.155.138.65/26
 ip martian-filter
 ip igmp mroute-proxy vlan 2001
 ip igmp mroute-proxy vlan 2002
 ip dhcp helper-address 10.1.6.60 10.1.7.60
 ip dhcp snooping
 exit
!
interface vlan 104
 no shutdown
 description 140782157_joamli_483_area
 ip address 112.156.219.97/27
 ip martian-filter
 ip igmp mroute-proxy vlan 2001
 ip igmp mroute-proxy vlan 2002
 ip dhcp helper-address 10.1.6.60 10.1.7.60
 ip dhcp snooping
 exit
!
interface vlan 106
 no shutdown
 description 140782834_dokjeong_119_area
 ip address 182.211.194.33/27
 ip martian-filter
 ip igmp mroute-proxy vlan 2001
 ip igmp mroute-proxy vlan 2002
 ip dhcp helper-address 10.1.6.60 10.1.7.60
 ip dhcp snooping
 exit
!
interface vlan 107
 no shutdown
 description 140782830_eoeunli_736_area
 ip address 49.171.103.193/27
 ip martian-filter
 ip igmp mroute-proxy vlan 2001
 ip igmp mroute-proxy vlan 2002
 ip dhcp helper-address 10.1.6.60 10.1.7.60
 ip dhcp snooping
 exit
!
interface vlan 108
 no shutdown
 description F/140784729/Joamri_460
 ip address 112.156.216.129/27
 ip martian-filter
 ip igmp mroute-proxy vlan 2001
 ip igmp mroute-proxy vlan 2002
 ip dhcp helper-address 10.1.6.60 10.1.7.60
 ip dhcp snooping
 exit
!
interface vlan 181
 no shutdown
 description 1100295441_jayeon_A_dreamCS
 ip address 112.158.93.193/27
 ip address 112.159.139.61/30 secondary
 ip martian-filter
 ip igmp mroute-proxy vlan 2001
 ip igmp mroute-proxy vlan 2002
 ip dhcp helper-address 10.1.6.60 10.1.7.60
 ip dhcp snooping
 exit
!
interface vlan 182
 no shutdown
 description 1100295434_jayeon_B_dreamCS
 ip address 112.158.93.225/27
 ip address 112.159.139.65/30 secondary
 ip martian-filter
 ip igmp mroute-proxy vlan 2001
 ip igmp mroute-proxy vlan 2002
 ip dhcp helper-address 10.1.6.60 10.1.7.60
 ip dhcp snooping
 exit
!
interface vlan 202
 no shutdown
 description F/140784248/Joamri270-65area
 ip address 58.77.114.129/27
 ip martian-filter
 ip igmp mroute-proxy vlan 2001
 ip igmp mroute-proxy vlan 2002
 ip dhcp helper-address 10.1.6.60 10.1.7.60
 ip dhcp snooping
 exit
!
interface vlan 203
 no shutdown
 description F/140784878/sekpoli_395_area
 ip address 112.156.10.225/27
 ip martian-filter
 ip igmp mroute-proxy vlan 2001
 ip igmp mroute-proxy vlan 2002
 ip dhcp helper-address 10.1.6.60 10.1.7.60
 ip dhcp snooping
 exit
!
interface vlan 204
 no shutdown
 description F/140784879/eoeunli_15_area
 ip address 112.155.250.193/26
 ip martian-filter
 ip igmp mroute-proxy vlan 2001
 ip igmp mroute-proxy vlan 2002
 ip dhcp helper-address 10.1.6.60 10.1.7.60
 ip dhcp snooping
 exit
!
interface vlan 302
 no shutdown
 description 302  F/140790715/joamri_311_1_A
 ip address 180.231.234.1/25
 ip martian-filter
 ip igmp mroute-proxy vlan 2001
 ip igmp mroute-proxy vlan 2002
 ip dhcp helper-address 10.1.6.60 10.1.7.60
 ip dhcp snooping
 exit
!
interface vlan 2001
 no shutdown
 description [MSE-T] JOAM_V6744-MSE1 Te37
 ip address 10.204.242.74/30
 ip igmp proxy-service
 ip igmp proxy-service priority 100
 ip igmp querier-timeout 255
 ip igmp unsolicited-report-interval 2
 service-policy input voipqos
 exit
!
interface vlan 2002
 no shutdown
 description [MSE-T] JOAM_V6744-MSE2 Te37
 ip address 10.204.243.74/30
 ip igmp proxy-service
 ip igmp proxy-service priority 100
 ip igmp querier-timeout 255
 ip igmp unsolicited-report-interval 2
 service-policy input voipqos
 exit
!
interface vlan 3000
 no shutdown
 ip address 10.30.0.1/16
 ip igmp mroute-proxy vlan 2001
 ip igmp mroute-proxy vlan 2002
 exit
!
!
threshold memory 80 600
!
arp patrol 2 5 1
arp access-list DAI
 permit ip 112.159.139.60/30 mac any
 permit ip 112.159.139.64/30 mac any
 permit dhcp-snoop-inspection
 no check port dhcp-snoop-inspection
 exit
!
ip arp inspection validate src-mac dst-mac ip 
ip arp inspection vlan 101,104,106-108,181-182,202-204,302
ip arp inspection filter DAI vlan 101,104,106-108,181-182,202-204,302
ip arp inspection log-buffer entries 1024
ip arp inspection log-buffer logs 0 interval 86400
!
ip prefix-list ANNOUNCE seq 5 deny 0.0.0.0/0 le 20
ip prefix-list ANNOUNCE seq 10 permit 0.0.0.0/0 le 32
ip prefix-list DACOM seq 5 permit 127.0.0.1/32
ip prefix-list P_XPEED seq 5 permit 127.0.0.1/32
ip prefix-list P_XPEED seq 10 permit 112.155.138.64/26
ip prefix-list P_XPEED seq 15 permit 112.156.219.96/27
ip prefix-list P_XPEED seq 20 permit 182.211.194.32/27
ip prefix-list P_XPEED seq 25 permit 49.171.103.192/27
ip prefix-list P_XPEED seq 30 permit 112.158.93.192/27
ip prefix-list P_XPEED seq 35 permit 112.159.139.60/30
ip prefix-list P_XPEED seq 40 permit 112.158.93.224/27
ip prefix-list P_XPEED seq 45 permit 112.159.139.64/30
ip prefix-list P_XPEED seq 50 permit 58.77.114.128/27
ip prefix-list P_XPEED seq 55 permit 112.156.216.128/27
ip prefix-list P_XPEED seq 60 permit 112.156.10.224/27
ip prefix-list P_XPEED seq 65 permit 112.155.250.192/26
ip prefix-list P_XPEED seq 70 permit 180.231.234.0/25
!
route-map RED_STATIC_ROUTE permit 10
 match ip address prefix-list P_XPEED
 set community 17858:110
 exit
!
route-map RED_STATIC_ROUTE permit 20
 match ip address prefix-list DACOM
 set community 3786:73
 exit
!
router bgp 17858
 bgp router-id 172.17.48.243
 bgp log-neighbor-changes
 neighbor 172.16.2.153 remote-as 17858
 neighbor 172.16.2.153 description JOAM-EG1  
 neighbor 172.16.2.153 update-source lo
 neighbor 172.16.2.153 timers 30 90
 neighbor 172.16.2.153 dont-capability-negotiate
 neighbor 172.16.2.163 remote-as 17858
 neighbor 172.16.2.163 description JOAM-EG2  
 neighbor 172.16.2.163 update-source lo
 neighbor 172.16.2.163 timers 30 90
 neighbor 172.16.2.163 dont-capability-negotiate
 maximum-paths ibgp 2
 !
 address-family ipv4
 bgp network import-check
 redistribute connected
 redistribute static
 neighbor 172.16.2.153 activate
 neighbor 172.16.2.153 next-hop-self
 neighbor 172.16.2.153 soft-reconfiguration inbound
 neighbor 172.16.2.153 prefix-list ANNOUNCE out
 neighbor 172.16.2.153 route-map RED_STATIC_ROUTE out
 neighbor 172.16.2.163 activate
 neighbor 172.16.2.163 next-hop-self
 neighbor 172.16.2.163 soft-reconfiguration inbound
 neighbor 172.16.2.163 prefix-list ANNOUNCE out
 neighbor 172.16.2.163 route-map RED_STATIC_ROUTE out
 exit-address-family
 !
 address-family ipv6
 bgp network import-check
 exit-address-family
 exit
!
ip route 172.16.2.153/32 10.203.237.73
ip route 172.16.2.153/32 10.203.238.73 100
ip route 172.16.2.163/32 10.203.237.73 100
ip route 172.16.2.163/32 10.203.238.73
!
ip tcp ignore rst-unknown 
ip tcp syncookies
!
ip igmp proxy-service multipath grpip
ip igmp if flap discredit-unit 1
ip igmp if flap recover-unit 30
ip igmp if flap recover-interval 150
!
ip igmp snooping
ip igmp snooping immediate-leave
ip igmp snooping vlan 2001 mrouter port CG1
ip igmp snooping vlan 2002 mrouter port CG2
!
ip igmp profile 1 
 permit 
 range 233.14.202.1 233.14.202.255
 range 233.18.145.1 233.18.145.255
 range 233.38.222.1 233.38.222.255
 range 233.69.194.1 233.69.194.255
 exit
!
ip igmp filter port 1/1-10/8 profile 1 
ip igmp filter port 1/1-10/8 packet-type reportv1
ip igmp filter port 11/1,12/1,CG1-CG2 packet-type reportv2
ip igmp filter port 1/1-10/8 packet-type reportv3
ip igmp filter port 1/1-10/8 packet-type query
!
snmp community rw powernms
no snmp trap pps-control
!
!
end

