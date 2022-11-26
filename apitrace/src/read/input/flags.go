package input 

import (
    "os"
    "flag" //https://pkg.go.dev/flag
    "log"
    "read/output"

)

/*
// could probably call this instead of running a cmp on each case
func check(err) {
	if err != nil {
		log.Fatal(err)
	}
	 
}
*/
type Input struct{
	ParseFile string	
	OutputToJSON string


}

func ReadArgs() Input {
    helpCommand := flag.NewFlagSet("help",flag.ExitOnError)
    ParseFileCommand := flag.NewFlagSet("file",flag.ExitOnError)
    OutputCommand := flag.NewFlagSet("out",flag.ExitOnError)

    /* Sub commands*/
    ParseFilePtr := ParseFileCommand.String("f", "", "file to read")
    OutputToJSONPtr := OutputCommand.String("o", "", "JSON output path")
   
    if len(os.Args) < 2 {
	    output.Help()
	    os.Exit(1)
    }
    
    switch os.Args[1] {
    case "help":
	    output.Help()
	    err := helpCommand.Parse(os.Args[2:])
	    if err != nil {
		    log.Fatal(err)
	    }
    case "file":
	    err := ParseFileCommand.Parse(os.Args[2:])
	    if err != nil {
                    log.Fatal(err)
            }
    case "out":
	    err := OutputCommand.Parse(os.Args[2:])


    }
    
    if ParseFileCommand.Parsed() {
	if *ParseFilePtr != "" {
		*ParseFilePtr = output.ParseFile() // build ParseFile() in /output/read.go
	}
    }


       if OutputCommand.Parsed() {
	       if *OutputToJSONPtr != "" {       //build AppplyExtension in /output/file.go
		       *OutputToJSONPtr = output.ApplyExtension(*OutputToJSONPtr, "json")  
	       }
       } 
}

