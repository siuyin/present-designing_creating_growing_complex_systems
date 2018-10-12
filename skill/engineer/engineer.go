package engineer

import "fmt"

func BuildBridge(length int) string {
	return fmt.Sprintf("Built %d m bridge.\n", length)
}

func PaveRoad(length, width int) string {
	return fmt.Sprintf("Paved %d x %d = %d square metres of road surface\n",
		length, width, length*width)
}
