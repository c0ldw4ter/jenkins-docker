package main

import (
	"fmt"
	"net/http"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)
func cpuHandler(w http.ResponseWriter, r *http.Request) {
    
    //CPU загрузка в %
    percent, err := cpu.Percent(0, false)
    if err != nil || len(percent) == 0 {
        http.Error(w, "Cannot get CPU usage", http.StatusInternalServerError)
        return
    }
    // Выдаем JSON
    fmt.Fprintf(w, `{"cpu": "%.2f"}`, percent[0]) 
}

//CPU частота в МГц
func cpuFreq(w http.ResponseWriter, r *http.Request){
    info, err := cpu.Info()
    if err != nil{
        http.Error(w, "Cannot get CPU info", http.StatusInternalServerError)
        return
    }
    fmt.Fprintf(w, `{"freq" : "%v MHz"}`, info[0].Mhz) 
}
//Load avarage
func loadAvgHandler(w http.ResponseWriter, r *http.Request){
    loadavg, err := load.Avg()
    if err != nil{
        http.Error(w, "Cannot get LoadAvg stat", http.StatusInternalServerError)
        return
    }
    fmt.Fprintf(w, `{"load_average": "%.2f, %.2f, %.2f"}`, loadavg.Load1, loadavg.Load5, loadavg.Load15)
}
// Использование оперативки
func memUsage(w http.ResponseWriter, r *http.Request){
    mem_u, err := mem.VirtualMemory()
    if err !=nil {
        http.Error(w, "Cannot get RAM usage", http.StatusInternalServerError)
        return
    }
    fmt.Fprintf(w, `{"RAM_Used" : "%d MB"}`, mem_u.Used / (1024*1024))
}


func main() {

  
    //Ручики для Golang
    http.HandleFunc("/frequency", cpuFreq) 
    http.HandleFunc("/loadavg", loadAvgHandler)
    http.HandleFunc("/memory", memUsage)
    http.HandleFunc("/cpu", cpuHandler)
    http.ListenAndServe(":8080", nil)
}
