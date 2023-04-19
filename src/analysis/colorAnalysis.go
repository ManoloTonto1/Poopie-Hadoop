package analysis

import (
	"fmt"
	"os"
	"strconv"
)

type Color struct {
	Name       string
	HexMinimum string
	HexMaximum string
}

var colorRanges = []Color{
	{Name: "Red", HexMinimum: "#FF0000", HexMaximum: "#FF9999"},
	{Name: "Orange", HexMinimum: "#FFA500", HexMaximum: "#FFC200"},
	{Name: "Yellow", HexMinimum: "#FFFF00", HexMaximum: "#FFFF99"},
	{Name: "Green", HexMinimum: "#008000", HexMaximum: "#00C957"},
	{Name: "Blue", HexMinimum: "#0000FF", HexMaximum: "#6699FF"},
	{Name: "Indigo", HexMinimum: "#4B0082", HexMaximum: "#6F00FF"},
	{Name: "Violet", HexMinimum: "#EE82EE", HexMaximum: "#D58BF6"},
	{Name: "White", HexMinimum: "#FFFFFF", HexMaximum: "#FFFFFF"},
	{Name: "Black", HexMinimum: "#000000", HexMaximum: "#333333"},
	{Name: "Gray", HexMinimum: "#808080", HexMaximum: "#C0C0C0"},
}

func GetColorInDecimal(hexColor string) uint64 {
	color, err := strconv.ParseUint(hexColor[1:], 16, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return color
}
func CheckRangesForColors(color string) string {
	for _, colorRange := range colorRanges {
		if (GetColorInDecimal(color) >= GetColorInDecimal(colorRange.HexMinimum)) && (GetColorInDecimal(color) <= GetColorInDecimal(colorRange.HexMaximum)) {
			return colorRange.Name
		}
	}
	return "None"
}
