package main

import (
	"os"
)

const FILE_NAME = "storage.json"

type Store struct {
	File os.File	
}

