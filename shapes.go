package geometry

import "math"

type Shape interface {
	BB() BB
	Translate(Vector)
    Collides(Shape) bool
}

func CircleRectCollision(c *Circle, r *Rect) bool {
	closestX := math.Max(r.Min.X, math.Min(c.center.X, r.Max.X))
	closestY := math.Max(r.Min.Y, math.Min(c.center.Y, r.Max.Y))
	dx := closestX - c.center.X
	dy := closestY - c.center.Y
	distance_squared := dx*dx + dy*dy

	if distance_squared <= c.r*c.r {
		return true
	}
	return false
}

func CircleBBCollision(c *Circle, bb *BB) bool {
	closestX := math.Max(bb.L, math.Min(c.center.X, bb.R))
	closestY := math.Max(bb.T, math.Min(c.center.Y, bb.B))
	dx := closestX - c.center.X
	dy := closestY - c.center.Y
	distanceSquared := dx*dx + dy*dy

	if distanceSquared <= c.r*c.r {
		return true
	}
	return false
}

func CircleCircleCollision(cl *Circle, cr *Circle) bool {
	dx := cr.center.X - cl.center.X
	dy := cr.center.Y - cl.center.Y
	distanceSquared := dx*dx + dy*dy
	radiiSum := cl.r + cr.r
	radiiSumSquared := radiiSum * radiiSum

	if distanceSquared <= radiiSumSquared {
		return true
	}
	return false
}
