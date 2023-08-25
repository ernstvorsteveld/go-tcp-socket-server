package arguments

import "fmt"

type Arguments struct {
	Server string
	Port   string
}

func HandleArguments(args []string) Arguments {
	return Arguments{
		Server: getArgument(args, 0, "server name"),
		Port:   getArgument(args, 1, "port number"),
	}
}

func getArgument(args []string, p int, attr string) string {
	if len(args) < p+1 {
		message := fmt.Sprintf("No argument found for %s. Please provide a valid value.", attr)
		fmt.Println(message)
		panic("Invalid argument list.\n" + message)
	}
	return args[p]
}
