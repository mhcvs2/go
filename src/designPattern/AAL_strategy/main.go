package main

import "designPattern/AAL_strategy/strategy"

func main() {
	var s strategy.IStrategy
	var c *strategy.Context
	s = new(strategy.BackDoor)
	c = strategy.NewContext(s)
	c.Operate()
	s = new(strategy.GivenGreenLight)
	c = strategy.NewContext(s)
	c.Operate()
	s = new(strategy.BlockEnemy)
	c = strategy.NewContext(s)
	c.Operate()
}
//back door...
//given green light
//block enemy
