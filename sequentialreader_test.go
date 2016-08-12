package shp

import (
	"os"
	"testing"
)

func openFile(name string, t *testing.T) *os.File {
	f, err := os.Open(name)
	if err != nil {
		t.Fatalf("Failed to open %s: %v", name, err)
	}
	return f
}

func getShapesSequentially(prefix string, t *testing.T) (shapes []Shape) {
	shp := openFile(prefix+".shp", t)
	dbf := openFile(prefix+".dbf", t)

	sr := SequentialReaderFromExt(shp, dbf)
	if err := sr.Err(); err != nil {
		t.Fatalf("Error when iterating over the shapefile header: %v", err)
	}
	for sr.Next() {
		_, shape := sr.Shape()
		shapes = append(shapes, shape)
	}
	if err := sr.Err(); err != nil {
		t.Fatalf("Error when iterating over the shapes: %v", err)
	}

	if err := sr.Close(); err != nil {
		t.Fatalf("Could not close sequential reader: %v", err)
	}
	return shapes
}

func TestSequentialReader(t *testing.T) {
	for prefix, _ := range dataForReadTests {
		t.Logf("Testing sequential read for %s", prefix)
		test_shapeIdentity(t, prefix, getShapesSequentially)
	}
}
