package main

import (
	"ball"
	"rect"
	"fmt"
)

func main() {
	var iRect rect.IRect = ball.CreateBall(1, 2, 2, 2)
	var iBall ball.IBall = ball.CreateBall(1, 2, 2, 2)

	fmt.Println(iRect.GetMaxX())
	fmt.Println(iBall.GetRadius())
}
