package models

import (
	"fmt"
	"html"
	"strconv"
)

const (
	Food   = "0x0001F35C"
	Finish = "0x0001F44C"
)

type Philosopher struct {
	ID                            int
	Name                          string
	LeftChopStick, RightChopStick *ChopStick
}

func (p Philosopher) Eat() {

	p.LeftChopStick.Lock()
	p.RightChopStick.Lock()

	fmt.Println(p.Name + " is eating " + GetEmoticon(Food))

	p.LeftChopStick.Unlock()
	p.RightChopStick.Unlock()
}

func GetEmoticon(value string) string {
	i, _ := strconv.ParseInt(value, 0, 64)
	foodEmoticon := html.UnescapeString(string(i))
	return foodEmoticon
}
