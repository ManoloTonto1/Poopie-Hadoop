package analysis

import (
	"os"
	"os/exec"
)

type ChartData struct {
	Labels []string
	Data   []int
}

func StartDashBoard() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cmd := exec.Command("cmd", "/c", "start", "", "C:/Users/manny/Desktop/pbi.lnk", dir+"/dashboard.pbix")
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
