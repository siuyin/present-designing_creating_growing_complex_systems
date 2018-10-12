package engineer

import "fmt"

func BuildBridge(length int) string {
	return fmt.Sprintf("Built %d m bridge.", length)
}

func PaveRoad(length, width int) string {
	return fmt.Sprintf("Paved %d x %d = %d square metres of road surface.",
		length, width, length*width)
}
