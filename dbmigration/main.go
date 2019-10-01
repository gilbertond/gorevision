package main

import (
	"flag"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pressly/goose"
	"log"
	"os"
)

var (
	flags     = flag.NewFlagSet("gooseMig", flag.ExitOnError)
	directory = flags.String("dir", ".", "directory with migration files")
)

func main() /*(string, error) */{
	flags.Parse(os.Args[1:])
	args := flags.Args()

	if len(args) < 2 {
		flags.Usage() //called when an error occurs
		return
	}

	dbstring, command := args[1], args[2]
	db, err := goose.OpenDBWithDriver("sqlserver", dbstring)

	if err != nil {
		log.Fatalf("Failed to open connection to database: %#v\n", err)
	}
	arguments := []string{}
	if len(args) > 3 {
		arguments = append(arguments, args[3:]...)
	}
	if err := goose.Run(command, db, *directory, arguments...); err != nil {
		log.Fatalf("goose command \n%v\n failed. \nerror: %v\n", command, err)
	}
}

//Possibly use this and let it have an S3 trigger event
func gooseStart() {
	lambda.Start(main)
}

//jdbc:sqlserver://localhost;database=master
