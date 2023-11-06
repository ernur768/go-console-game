package term

import (
	"bufio"
	"consoleTest/gamePhysics"
	"fmt"
	"os"
)

var screen = bufio.NewWriter(os.Stdout)

func HideCursor() {
	_, _ = fmt.Fprint(screen, "\033[?25l")
}

func ShowCursor() {
	_, _ = fmt.Fprint(screen, "\033[?25h")
}

func Clear() {
	_, _ = fmt.Fprint(screen, "\033[2J")
}

func ClearLine() {
	fmt.Print("\033[2K")
}

func MoveCursor(pos gamePhysics.Position) {
	_, _ = fmt.Fprintf(screen, "\033[%d;%dH", pos.Y, pos.X)
}

func Draw(str string) {
	_, _ = fmt.Fprint(screen, str)
}

func MoveCursorAndDraw(pos gamePhysics.Position, str string) {
	MoveCursor(pos)
	Draw(str)
}

func Render() {
	screen.Flush()
}
