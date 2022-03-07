package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func Line() {
	white := color.New(color.FgWhite, color.Bold)
	white.Println(strings.Repeat("-", 42))
}

func Input(name string) string {
	green := color.New(color.FgGreen, color.Bold)
	green.Print(name)

	color.Set(color.FgYellow, color.Bold)
	defer color.Unset()

	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	return input.Text()
}

func Encryption(data string) string {
	// sha256(data) result
	s256Hash := sha256.Sum256([]byte(data))
	s256Text := hex.EncodeToString(s256Hash[:])
	// md5(sha256(data)) result
	wandHash := md5.Sum([]byte(s256Text))
	wandText := hex.EncodeToString(wandHash[:])

	return wandText
}

func main() {

	Execute, _ := exec.Command("sh", "-c", "clear").Output()
	fmt.Print(string(Execute))

	green := color.New(color.FgGreen, color.Bold).SprintFunc()
	yellow := color.New(color.FgYellow, color.Bold).SprintFunc()

	Line()

	color.Set(color.FgYellow, color.Bold)
	defer color.Unset()

	δ := Input("δpq: ")
	Δ := Input("Δpq: ")

	Line()

	for i := 0; i < 64; i++ {
		δ = Encryption(δ)
		fmt.Printf("[%v]:\t%v\n", green(i), yellow(δ[:0x11], Δ))
	}

	Line()
}
