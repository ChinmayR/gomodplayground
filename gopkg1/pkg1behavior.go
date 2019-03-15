package gopkg1

import (
	"fmt"

	"github.com/jessevdk/go-flags"
)

type options struct {
	Message string `long:"msg" short:"m"`
}

func main() {
	PrintPkg1Behavior()
	return
}

func PrintPkg1Behavior() {
	var opts options
	parser := flags.NewParser(&opts, flags.Default)
	if _, err := parser.Parse(); err != nil {
		return
	}

	fmt.Println("Hello from pkg1 got msgs: " + opts.Message)
}
