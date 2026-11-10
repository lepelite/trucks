package models_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"trucks/internal/models"
)

func TestNewLoadsSkipsBlankLines(t *testing.T) {
	path := filepath.Join(t.TempDir(), "manifest.txt")
	contents := "loadNumber pickup dropoff\n\n1 (3,4) (6,8)\n \t\n2 (-1,2) (0,0)\n"
	if err := os.WriteFile(path, []byte(contents), 0600); err != nil {
		t.Fatal(err)
	}
	loads := models.NewLoads(path)
	if len(loads) != 2 || loads[0].Id != 1 || loads[1].Id != 2 {
		t.Fatalf("expected loads 1 and 2, got %v", loads)
	}
	if loads[0].Start.X != 3 || loads[0].Start.Y != 4 || loads[1].Start.X != -1 {
		t.Fatalf("unexpected pickup coordinates: %v", loads)
	}
}

func TestNewLoadsRejectsTruncatedManifest(t *testing.T) {
	path := filepath.Join(t.TempDir(), "manifest.txt")
	contents := "loadNumber pickup dropoff\n1 (0,0) (1,1)\n" + strings.Repeat("x", 70*1024)
	if err := os.WriteFile(path, []byte(contents), 0600); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if recover() == nil {
			t.Fatal("expected a read error instead of a partial manifest")
		}
	}()
	models.NewLoads(path)
}

func TestParsePointRejectsNonFiniteCoordinates(t *testing.T) {
	for _, input := range []string{"(NaN,0)", "(0,NaN)", "(+Inf,0)", "(0,-Inf)"} {
		t.Run(input, func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Fatalf("expected rejection of %s", input)
				}
			}()
			models.ParsePoint(input)
		})
	}
}
