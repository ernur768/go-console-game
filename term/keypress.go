package term

import (
	"github.com/mattn/go-tty"
	"log"
	"sync"
)

const closeKey = 'p'

var keyInputs []KeyInput
var mu sync.Mutex

type KeyInput struct {
	id int
	C  chan rune
}

func newKeyInput() *KeyInput {
	return &KeyInput{
		id: len(keyInputs),
		C:  make(chan rune, 1),
	}
}

func (ki *KeyInput) Close() {
	mu.Lock()
	defer mu.Unlock()

	// Closing channel
	close(ki.C)

	// Removing from keyInputs slice
	freeIndex := ki.id
	lastIndex := len(keyInputs) - 1
	keyInputs[lastIndex].id = freeIndex

	keyInputs[freeIndex] = keyInputs[lastIndex]
	//keyInputs[lastIndex] = nil
	keyInputs = keyInputs[:lastIndex]
}

func GetKeyInput() KeyInput {
	mu.Lock()
	defer mu.Unlock()

	if keyInputs == nil {
		keyInputs = make([]KeyInput, 0)
		newKI := newKeyInput()
		keyInputs = append(keyInputs, *newKI)
		go handleKeyEvents()
		return *newKI
	}

	newKI := newKeyInput()
	keyInputs = append(keyInputs, *newKI)
	return *newKI
}

func handleKeyEvents() {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	for {
		key, err := tty.ReadRune()
		if err != nil {
			panic(err)
		}

		if key == 0 {
			continue
		}

		if key == closeKey {
			for _, ki := range keyInputs {
				close(ki.C)
			}
			break
		}

		for _, ki := range keyInputs {
			ki.C <- key
		}
	}
}
