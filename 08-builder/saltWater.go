package main

type SaltWater struct {
	Salt  float64
	Water float64
}

func NewSaltWater(salt float64, water float64) *SaltWater {
	return &SaltWater{
		Salt:  salt,
		Water: water,
	}
}
