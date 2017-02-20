package main

import "fmt"

type SaltWaterBuilder struct {
	SaltWater *SaltWater
}

func NewSaltWaterBuilder() *SaltWaterBuilder {
	return &SaltWaterBuilder{
		SaltWater: NewSaltWater(0, 0),
	}
}
func (self *SaltWaterBuilder) addSolute(saltAmount float64) {
	self.SaltWater.Salt += saltAmount
}

func (self *SaltWaterBuilder) addSolvent(waterAmount float64) {
	self.SaltWater.Water += waterAmount
}
func (self *SaltWaterBuilder) abandonSolution(saltWaterAmount float64) {
	saltDelta := saltWaterAmount * (self.SaltWater.Salt / (self.SaltWater.Salt + self.SaltWater.Water))
	waterDelta := saltWaterAmount * (self.SaltWater.Water / (self.SaltWater.Salt + self.SaltWater.Water))
	self.SaltWater.Salt -= saltDelta
	self.SaltWater.Water -= waterDelta
}
func (self *SaltWaterBuilder) getResult() {
	fmt.Println(self.SaltWater)
}
