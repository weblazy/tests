package main

import (
	"fmt"
	"net"
)

func main() {
	LookupIP()
	LookupCNAME()
	LookupNS()
	LookupMX()
	LookupSRV()
	LookupTXT()
}
func LookupIP() {
	ipRecords, _ := net.LookupIP("google.com")
	for _, ip := range ipRecords {
		fmt.Println(ip)
	}
}

func LookupCNAME() {
	cname, _ := net.LookupCNAME("m.baidu.com")
	fmt.Println(cname)
}

func LookupNS() {
	nameServer, _ := net.LookupNS("baidu.com")
	for _, ns := range nameServer {
		fmt.Println(ns)
	}
}

func LookupMX() {
	mxRecords, _ := net.LookupMX("baidu.com")
	for _, mx := range mxRecords {
		fmt.Println(mx.Host, mx.Pref)
	}
}

func LookupSRV() {
	cname, srvs, err := net.LookupSRV("xmpp-server", "tcp", "baidu.com")
	if err != nil {
		panic(err)
	}

	fmt.Printf("\ncname: %s \n\n", cname)

	for _, srv := range srvs {
		fmt.Printf("%v:%v:%d:%d\n", srv.Target, srv.Port, srv.Priority, srv.Weight)
	}
}

func LookupTXT() {
	txtRecords, _ := net.LookupTXT("baidu.com")

	for _, txt := range txtRecords {
		fmt.Println(txt)
	}
}
