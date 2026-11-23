package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	subject_command := flag.NewFlagSet("example", flag.ExitOnError)
	subject_enable := subject_command.Bool("enable", false, "enable")
	subject_name := subject_command.String("name", "", "name")

	bar_command := flag.NewFlagSet("bar", flag.ExitOnError)
	bar_level := bar_command.Int("level", 0, "level")

	if len(os.Args) < 2 {
		fmt.Println("Expected cpmmands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "example":
		subject_command.Parse(os.Args[2:])
		fmt.Println("Enable:", *subject_enable)
		fmt.Println("Name:", *subject_name)
		fmt.Println("Tail:", subject_command.Args())
	case "bar":
		bar_command.Parse(os.Args[2:])
		fmt.Println("Subcommand 'example")
		fmt.Println("Level:", bar_level)
		fmt.Println("Tail:", bar_command.Args())
	default:
		fmt.Println("Exepcted 'example' or 'bar' subcommands00")
		os.Exit(1)
	}
}
