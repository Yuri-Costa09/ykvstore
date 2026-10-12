package repl

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Yuri-Costa09/ykvstore/store"
)

var ReplPrompt = "yKvstore"

func StartRepl(store *store.Store) {

	commands := map[string]func(){
		".help":  displayHelp,
		".clear": clearScreen,
	}

	reader := bufio.NewScanner(os.Stdin)
	printPrompt()

	for reader.Scan() {
		text := reader.Text()

		if command, exists := commands[text]; exists {
			command()
		}

		if strings.EqualFold(text, ".exit") {
			exitApp()
		}

		handleCmd(store, text)
	}

}

func handleCmd(s *store.Store, text string) error {
	text = strings.ToLower(text)
	parts := strings.Fields(text)

	switch parts[0] {
	case "put":
		s.Put(parts[1], parts[2])

	case "get":
		value, err := s.Get(parts[1])
		if strings.EqualFold(value, "") {
			fmt.Println(parts[1], ": Not found")
		}
		fmt.Println(value)
		if err != nil {
			return err
		}
		return err

	case "remove":
		err := s.Remove(parts[1])
		if err != nil {
			return err
		}
	}

	return nil
}

func printUnknown(command string) {
	fmt.Println(command, ": Command not found.")
}

func printPrompt() {
	fmt.Println(ReplPrompt)
}

func displayHelp() {
	fmt.Printf(
		"Welcome to yKvstore! These are the available commands: \n",
	)
	fmt.Println(".help    - Show available commands")
	fmt.Println(".clear   - Clear the terminal screen")
	fmt.Println(".exit    - Closes your connection to yKvstore")
}

// clearScreen clears the terminal screen
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func exitApp() {
	cmd := exec.Command("exit")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// handleInvalidCmd attempts to recover from a bad command
func handleInvalidCmd(text string) {
	defer printUnknown(text)
}

// cleanInput preprocesses input to the db repl
func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}
