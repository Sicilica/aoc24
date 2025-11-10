package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Sicilica/aoc24/lib"
)

func main() {
	lib.Assert(len(os.Args) == 2)
	day := lib.Atoi(os.Args[1])
	runDay(day)
}

func runDay(day int) {
	fmt.Println("========")
	fmt.Println("Day", day)
	fmt.Println("========")
	cmd := exec.Command("go", "run", fmt.Sprintf("github.com/Sicilica/aoc24/days/%d", day))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	lib.NoErr(cmd.Run())
}
