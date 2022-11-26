package input 

import (
    "os"
    "flag"
    "log"
    "read/output"

)
type Input struct{
	ReadFile string	
	OutputToJSON string


}

func ReadArgs() Input {
    helpCommand := flag.NewFlagSet("help",flag.ExitOnError)
    ReadFileCommand := flag.NewFlagSet("file",flag.ExitOnError)
    OutputCommand := flag.NewFlagSet("out",flag.ExitOnError)

    /* Sub commands*/
    ReadFilePtr := ReadFileCommand.String("f", "", "file to read")
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
	    err := ReadFileCommand.Parse(os.Args[2:])
	    if err != nil {
                    log.Fatal(err)
            }
    case "out":
	    err := OutputCommand.Parse(os.Args[2:])


    }
    
    if ReadFileCommand.Parsed() {
	if *ReadFilePtr != "" {
		*ReadFilePtr = output.ReadFile() // build ReadFile() in /output/read.go
	}
    }


       if OutputCommand.Parsed() {
	       if *OutputToJSONPtr != "" {       //build AppplyExtension in /output/file.go
		       *OutputToJSONPtr = output.ApplyExtension(*OutputToJSONPtr, "json")  
	       }
       } 
}

