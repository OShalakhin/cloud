package main

import (
	"flag"
	"fmt"
)

const (
	// VERSION of package
	VERSION = "VERSION\n" +
		"\t 0.1.0\n\n"
	// DESCRIPTION what is package for
	DESCRIPTION = "DESCRIPTION\n" +
		"\tSync your data with cloud storages (like Amazon S3, Rackspace CloudFiles etc.)\n\n"
	// COMMANDS lists commands available
	COMMANDS = "COMMANDS\n" +
		"\tinit\tinitialize .cloudcore and .cloud files\n" +
		"\tsync\tsynchronize folder with the cloud\n" +
		"\tupdate\tupdate file or folder in container\n" +
		"\turl\tshow container url\n" +
		"\tignore\tignore particular file with .cloudignore\n" +
		"\tclear\tclear container\n" +
		"\thelp\tshow this message\n\n"
	// CONTRIBUTORS shows package author & contributors
	CONTRIBUTORS = "CONTRIBUTORS\n" +
		"\tOlexandr Shalakhin <olexandr@shalakhin.com>\n\n"
)

// usage is responsible to generate command help message
func usage() {
	fmt.Printf(DESCRIPTION + VERSION + COMMANDS + CONTRIBUTORS)
	flag.PrintDefaults()
}

// getContainerName if it is presented in args as a second parameter
func getContainerName(args []string) string {
	// default container is the first in .cloud
	container := ""
	if len(args) > 2 {
		// use defined
		container = args[1]
	}
	return container
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	switch {
	case args[0] == "init":
		initConfigs()
	case args[0] == "sync":
		Sync(getContainerName(args))
	// case args[0] == "update":
	// 	Update(, getContainerName(args))
	case args[0] == "url":
		GetContainerURL(getContainerName(args))
	case args[0] == "ignore":
		fmt.Println("Nothing here yet...")
	case args[0] == "clear":
		fmt.Println("Nothing here yet...")
	case args[0] == "help":
		usage()
	default:
		usage()
	}
}
