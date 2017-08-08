//usr/local/bin/go run $0 $@ ; exit
package main

import (
	"fmt"
	"os"
)

func main() {
	os.Chdir(os.Getenv("PROJECTNAME_ROOT"))
	fmt.Println("go script boilerplate")
}
