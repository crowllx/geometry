package geometry

type BB struct {
	L, T, R, B float64
}

// BB implements Shape.
func (bb *BB) BB() BB {
	return *bb
}

// Translate implements Shape.
func (bb *BB) Translate(v Vector) {
	bb.L += v.X
	bb.R += v.X
	bb.T += v.Y
	bb.B += v.Y
}

var _ Shape = &BB{}

func (bb BB) Collides(s Shape) bool {
	other := s.BB()
	if bb.R < other.L || bb.L > other.R {
		return false
	}
	if bb.B < other.T || bb.T > other.B {
		return false
	}
	return true
}

func (bb BB) Contains(other BB) bool {
	if bb.R < other.L || bb.L > other.R {
		return false
	}
	if bb.B < other.T || bb.T > other.B {
		return false
	}
	return true
}
