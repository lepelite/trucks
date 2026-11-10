package calc_test

import (
	"testing"
	"trucks/internal/calc"
	"trucks/internal/models"
)

func TestDistanceZero(t *testing.T) {
	p1 := models.NewPoint(0, 0)
	if calc.Distance(p1, p1) != 0.0 {
		t.Fatalf("Distance not zero between same point")
	}
	p2 := models.NewPoint(5, 5)
	if calc.Distance(p2, p2) != 0.0 {
		t.Fatalf("Distance not zero between same point")
	}

}

func TestDistance345(t *testing.T) {
	d := calc.Distance(models.NewPoint(0, 0), models.NewPoint(3, 4))
	if d != 5.0 {
		t.Fatalf("Failed to make 3 4 5 triangle")
	}
}

func TestRouteSingleLoad(t *testing.T) {
	loads := []models.Load{{Id: 42, Start: models.NewPoint(3, 4), End: models.NewPoint(6, 8)}}
	routes := calc.Route(loads, 0)
	if len(routes) != 1 || len(routes[0]) != 1 || routes[0][0] != 42 {
		t.Fatalf("expected one route containing load 42, got %v", routes)
	}
	if distance := calc.RouteDistance(routes[0], loads); distance != 20 {
		t.Fatalf("expected round-trip distance 20, got %v", distance)
	}
}
