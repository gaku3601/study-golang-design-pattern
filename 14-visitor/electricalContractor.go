package main

import "fmt"

type ElectricalContractor struct{}

func NewElectricalContractor() *ElectricalContractor {
	return &ElectricalContractor{}
}
func (self *ElectricalContractor) visit() {
	fmt.Println("電気工事屋で〜す。工事させてもらいに来ました〜。")
}
