package geometry

type ConvexPolygon struct {
	Vertices []Vector
}
type Edge struct {
	a, b Vector
}

func NewConvexPolygon(vertices ...Vector) *ConvexPolygon {
	return &ConvexPolygon{
		Vertices: vertices,
	}
}

func (cp *ConvexPolygon) BB() BB {
	minX := cp.Vertices[0].X
	maxX := cp.Vertices[0].X
	minY := cp.Vertices[0].Y
	maxY := cp.Vertices[0].Y

	for _, vec := range cp.Vertices[1:] {
		if vec.X < minX {
			minX = vec.X
		}

		if vec.X > maxX {
			maxX = vec.X
		}
		if vec.Y < minY {
			minY = vec.Y
		}

		if vec.Y > maxY {
			maxY = vec.Y
		}
	}

	return BB{minX, minY, maxX, maxY}
}

// apply translate vector to all vertices of polygon
func (cp *ConvexPolygon) Translate(t Vector) {
	for i, vec := range cp.Vertices {
		cp.Vertices[i] = vec.Add(t)
	}
}

// unimplemented
func (cp *ConvexPolygon) Collides(s Shape) bool {
	if cp.BB().Contains(s.BB()) {
		switch s.(type) {
		default:
			return true
		}
	}
	return false
}

func (cp ConvexPolygon) Edges() []Edge {
	var edges []Edge

	for i := range len(cp.Vertices) - 1 {
		edges = append(edges, Edge{cp.Vertices[i], cp.Vertices[i+1]})
	}
	edges = append(edges, Edge{cp.Vertices[len(cp.Vertices)-1], cp.Vertices[0]})

	return edges
}

func (cp *ConvexPolygon) Rotate(angle float32, origin Vector) {
	for i, vec := range cp.Vertices {
		cp.Vertices[i] = vec.Rotate(angle, origin)
	}
}

func (cp ConvexPolygon) Centroid() Vector {
	edges := cp.Edges()
	var midX, midY float32

	for _, edge := range edges {
		midX += (edge.a.X + edge.b.X) / 2
		midY += (edge.a.Y + edge.b.Y) / 2
	}
	midX = midX / float32(len(edges))
	midY = midY / float32(len(edges))

	return NewVector(midX, midY)
}

var _ Shape = &ConvexPolygon{}
