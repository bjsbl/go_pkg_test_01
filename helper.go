package go_pkg_test_01

import "fmt"

const version = "1.0.1"

func cliVersion() {
	fmt.Println(" pkg version " + version)
}

func ShowVersion() {
	fmt.Println(" pkg version " + version)
}
