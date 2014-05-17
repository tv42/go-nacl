package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"syscall"
)

var prog = filepath.Base(os.Args[0])

var loaders = map[string]string{}

func main() {
	log.SetFlags(0)
	log.SetPrefix(prog + ": ")

	if err := os.Setenv("GOOS", "nacl"); err != nil {
		log.Fatal(err)
	}

	arch := runtime.GOARCH
	if userArch := os.Getenv("GOARCH"); userArch != "" {
		// explicit in environment
		arch = userArch
	}

	// translate to nacl-specific architecture
	switch arch {
	case "amd64":
		arch = "amd64p32"
	}

	if err := os.Setenv("GOARCH", arch); err != nil {
		log.Fatal(err)
	}

	var bin string
	var err error
	if bin, err = exec.LookPath("go"); err != nil {
		log.Fatal(err)
	}
	args := os.Args
	args[0] = bin

	if len(args) >= 2 {
		switch args[1] {
		case "test", "run":
			var loader string
			switch arch {
			case "amd64p32":
				loader = "sel_ldr_x86_64"
			case "386":
				loader = "sel_ldr_x86_32"
			default:
				log.Fatalf("no NaCl loader for architecture: %v", arch)
			}
			args = append(args, "")
			copy(args[2+1:], args[2:])
			args[2] = "-exec=" + loader
		}
	}

	if err := syscall.Exec(bin, args, os.Environ()); err != nil {
		log.Fatal(err)
	}
	panic("still running after successful exec")
}
