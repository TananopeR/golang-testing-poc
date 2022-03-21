package testitySuite

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("main")
	run := m.Run()
	fmt.Println("main finished")
	os.Exit(run)
}

func Test1(t *testing.T) {
	t.Cleanup(func() { fmt.Println("test 1 clean up") })

	fmt.Println("test 1")
}
