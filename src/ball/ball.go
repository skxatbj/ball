package ball

import (
	"rect"
)

type IBall interface {
	SetBallPos(cx, cy, radius float32)

	GetRadius() float32
	GetWeight() float32
}

type Ball struct {
	*rect.Rect

	radius float32
	weight float32
}

func (ball *Ball) SetBallPos(cx, cy, radius float32) {
	if (ball.radius != radius) {
		ball.radius = radius
		ball.weight = RadiusToWeight(radius)

		ball.SetW(radius * 2)
		ball.SetH(radius * 2)
	}

	ball.SetMidX(cx)
	ball.SetMidY(cy)
}

func (ball *Ball) GetRadius() float32 {
	return ball.radius
}

func (ball *Ball) GetWeight() float32 {
	return ball.weight
}

func CreateBall(bid int32, cx, cy, radius float32) *Ball {
	return &Ball {
		rect.CreateRect(bid, radius * 2, radius * 2, cx, cy),
		radius,
		RadiusToWeight(radius),
	}
}

/////////////////////////////////////////////////////////////////
func RadiusToWeight(radius float32) float32 {
	return radius * 3 * 3
}

func WeightToRadius(weight float32) float32 {
	return weight / 3 / 3
}
