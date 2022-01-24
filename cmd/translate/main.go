// translate updates meaning ordered words files.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/napcatstudio/translate/xlns"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 3 {
		usage("incorrect number of arguments")
	}
	credentialsJson := args[0]
	//wordsDir := args[1]
	//mainLang := args[2]

	_, err := xlns.GetTranslateService(credentialsJson)
	if err != nil {
		log.Fatalf("error %v", err)
	}

	//editId, err := apta.EditsInsert(service, packageName)
	//if err != nil {
	//	log.Fatalf("error %v", err)
	//}

	// Details
	//appDetails, err := service.Edits.Details.Get(packageName, editId).Do()
	//if err != nil {
	//	log.Fatalf("getting %s details got %v", packageName, err)
	//}
	//fmt.Printf("%s %s\n%s\n%s\n",
	//	packageName, appDetails.DefaultLanguage,
	//	appDetails.ContactEmail,
	//	appDetails.ContactWebsite)
	//defLang := appDetails.DefaultLanguage

	fmt.Println("it worked?")
}

func usage(why string) {
	text := `credentialsJson wordsDir mainLang

    Uses the service information in credentialsJson to access the Google
    Translate API to update meaning ordered words files in wordsDir.

Where:
    credentialsJson is a JSON file with Google Service info and keys.
    wordsDir is a directory containing meaning ordered words files.
    mainLang is the ISO639 code for the words file to be used to update 
        the others.

Example:
    translate yourServiceKey.json words en
`
	log.Fatalf("ERROR: %s.\nUSAGE:\n    %s %s",
		why,
		os.Args[0],
		text)
}
