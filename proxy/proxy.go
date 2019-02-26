package main

import "fmt"

type GiveGift interface {
	GiveFollower()
}

type Pursuit struct {
	lover string
}

func (this *Pursuit) GiveFollower() {
	fmt.Println(this.lover + "：给你花！")
}

type Proxy struct {
	Pursuit Pursuit
}

func (this *Proxy) GiveFollower() {
	this.Pursuit.GiveFollower()
}

func main() {
	me := Pursuit{lover: "小美"}
	var you GiveGift = &Proxy{Pursuit: me}
	you.GiveFollower()
}
