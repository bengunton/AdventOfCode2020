package main

import(
	"testing"
)

func TestAbs(t *testing.T) {
	res := abs(-1)

	if res != 1 {
		t.Errorf("You messed up a simple abs function")
	}
}

func checkDistance(expected, actual int64, t *testing.T) {
	if actual != expected {
		t.Errorf("Expected distance: %d, actual: %d", expected, actual)
	}
}

func TestRunSample(t *testing.T) {
	checkDistance(runFile("sample"), 286, t)
}

func TestRunRotate(t *testing.T) {
	checkDistance(runFile("rotate"), 51, t)
}

func checkLocation(expected1, expected2, actual1, actual2 int64, t *testing.T) {
	if actual1 != expected1 {
		t.Errorf("Expected: %d, actual: %d", expected1, actual1)
	}
	if actual2 != expected2 {
		t.Errorf("Expected: %d, actual: %d", expected2, actual2)
	}
}

func TestMoveToWaypointLeft(t *testing.T) {
	boat := Boat{North, 0, 0}
	wp := Waypoint{1, 1}

	wp = boat.moveWaypoint64("F1", wp)
	checkLocation(1, 1, boat.northOffset, boat.eastOffset, t)

	wp = boat.moveWaypoint64("L90", wp)
	wp = boat.moveWaypoint64("F1", wp)
	checkLocation(2, 0, boat.northOffset, boat.eastOffset, t)

	wp = boat.moveWaypoint64("N1", wp)
	wp = boat.moveWaypoint64("F1", wp)
	checkLocation(4, -1, boat.northOffset, boat.eastOffset, t)
}

func TestMoveToWaypointLeftTwice(t *testing.T) {
	boat := Boat{North, 0, 0}
	wp := Waypoint{1, 1}

	wp = boat.moveWaypoint64("F1", wp)
	checkLocation(1, 1, boat.northOffset, boat.eastOffset, t)

	wp = boat.moveWaypoint64("L90", wp)
	wp = boat.moveWaypoint64("L90", wp)
	wp = boat.moveWaypoint64("F1", wp)
	checkLocation(0, 0, boat.northOffset, boat.eastOffset, t)

	wp = boat.moveWaypoint64("N1", wp)
	wp = boat.moveWaypoint64("F1", wp)
	checkLocation(0, -1, boat.northOffset, boat.eastOffset, t)
}

func TestMoveToWaypointRight(t *testing.T) {
	boat := Boat{North, 0, 0}
	wp := Waypoint{1, 1}

	wp = boat.moveWaypoint64("F1", wp)
	checkLocation(1, 1, boat.northOffset, boat.eastOffset, t)

	wp = boat.moveWaypoint64("R90", wp)
	wp = boat.moveWaypoint64("R90", wp)
	wp = boat.moveWaypoint64("R90", wp)
	wp = boat.moveWaypoint64("F1", wp)
	checkLocation(2, 0, boat.northOffset, boat.eastOffset, t)

	wp = boat.moveWaypoint64("R90", wp)
	wp = boat.moveWaypoint64("N1", wp)
	wp = boat.moveWaypoint64("F1", wp)
	checkLocation(4, 1, boat.northOffset, boat.eastOffset, t)
}