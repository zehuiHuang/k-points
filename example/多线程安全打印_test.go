package example

import "testing"

func TestExample1(t *testing.T) {
	example1()
}

func TestExample2(t *testing.T) {
	example2()
}

func TestBenchmarkSyncMap(t *testing.T) {
	benchmarkSyncMap()
}

func TestBenchmarkMap(t *testing.T) {
	benchmarkMap()
}
