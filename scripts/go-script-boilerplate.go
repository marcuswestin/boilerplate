//usr/local/bin/go run $0 $@ ; exit
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("go script boilerplate")
	os.Chdir(os.Getenv("PROJECTNAME_ROOT"))
}
