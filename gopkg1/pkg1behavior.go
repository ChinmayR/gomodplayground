package gopkg1

import (
	"fmt"

	"github.com/ChinmayR/gomodplayground/gopkg2"

	"github.com/davecgh/go-spew/spew"
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
	gopkg2.PrintPkg2Behavior()
	spew.Dump("Hello from pkg1 using spew")
}
