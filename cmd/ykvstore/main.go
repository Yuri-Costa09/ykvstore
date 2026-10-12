package main

import (
	"github.com/Yuri-Costa09/ykvstore/repl"
	"github.com/Yuri-Costa09/ykvstore/store"
)

func main() {
	// initialize in-memory key-value store

	store := store.NewMemoryStore()

	repl.StartRepl(store)
}
