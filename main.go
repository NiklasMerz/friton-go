package main

import (
	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

func main() {
	// Enable BLE interface.
	must("enable BLE stack", adapter.Enable())

	// Start scanning.
	println("scanning...")
	err := adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
		println("found device:", device.Address.String(), device.RSSI, device.LocalName())
		if device.Address.String() == "FC:DA:D9:DE:CC:70" {
			connect(device.Address)
			adapter.StopScan()
		}
	})
	must("start scan", err)
}

func must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ": " + err.Error())
	}
}

func connect(a bluetooth.Addresser) {
	d, err := adapter.Connect(a, bluetooth.ConnectionParams{})
	if err != nil {
		panic(err)
	}
	s, err := d.DiscoverServices(nil)
	if err != nil {
		panic(err)
	}

	for _, service := range s {
		println(service.UUID().Get16Bit())
		c, err := service.DiscoverCharacteristics(nil)
		if err != nil {
			panic(err)
		}
		for _, cr := range c {
			println("Characteristic")
			println(cr.String())
		}
	}
}
