package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"html/template"
	"net/http"
	"strconv"
)

const (
	Port = ":8080"
)

type Page struct {
	Title    string
	Hostname string
	OS       string
	MemTotal string
	MemUsed  string
	CpuModel string
}

func serveStatus(w http.ResponseWriter, r *http.Request) {
	v, _ := mem.VirtualMemory()
	c, _ := cpu.Info()
	h, _ := host.Info()
	thisPage := Page{
		Title:    "Status Page",
		Hostname: h.Hostname,
		OS:       h.OS,
		MemTotal: strconv.Itoa(int(v.Total)),
		MemUsed:  strconv.Itoa(int(v.Used)),
		CpuModel: c[0].ModelName,
	}
	t, _ := template.ParseFiles("status.html")
	t.Execute(w, thisPage)
}

func main() {
	fmt.Println("Starting Webserver")
	http.HandleFunc("/status", serveStatus)
	http.ListenAndServe(Port, nil)
}
