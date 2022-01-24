// translate updates meaning ordered words files.
package main

import (
	"flag"
	"fmt"
	"github.com/napcatstudio/translate/xlns"
	"log"
	"os"
	"path"
)

/*
    Uses the service information in credentialsJson to access the Google
    Translate API to update meaning ordered words files in wordsDir.

Where:
    credentialsJson is a JSON file with Google Service info and keys.
    wordsDir is a directory containing meaning ordered words files.
    mainLang is the ISO639 code for the words file to be used to update
        the others.

Example:
    translate yourServiceKey.json words en
*/

const (
	USAGE = `translate is a tool for managing meaning ordered words files.

It uses the Google Translate API for translating.

Usage:
	translate [-words wordsDir] command [arguments]

The commands are:
	add [-credentials credentialsJson] mainLang newLang
	  Add a new meaning ordered words file for newLang based on mainLang to
	  wordsDir.
	check
	  Quick wordsDir check.  Does not check translation accuracy just
	  consistency.
	recreate  [-credentials credentialsJson] mainLang
	  Recreate wordsDir meaning ordered words files.
	update  [-credentials credentialsJson] mainLang
	  Updates all meaning ordered words files in wordsDir.

`
)

func main() {
	wordsDir := flag.String("words", "words", "meaning ordered words directory")
	//addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	//addCreds :=
	//checkCmd := flag.NewFlagSet("check", flag.ExitOnError)
	//checkWordsDir := checkCmd.String("words", "words", "meaning ordered words directory")
	//recreateCmd := flag.NewFlagSet("recreate", flag.ExitOnError)
	//updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	//fooEnable := fooCmd.Bool("enable", false, "enable")
	//fooName := fooCmd.String("name", "", "name")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, USAGE)
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nadd mainLang newLang\n")
		fmt.Fprintf(os.Stderr, "\ncheck\n")
		//checkCmd.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nrecreate\n")
		fmt.Fprintf(os.Stderr, "\nupdate\n")
	}
	flag.Parse()
	err := isDir(*wordsDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "bad wordsDir (%v)", err)
		flag.Usage()
		log.Fatal(err)
	}
	if len(flag.Args()) < 1 {
		flag.Usage()
		log.Fatal("no command")
	}

	// Run command.
	switch flag.Arg(0) {
	case "add":
		err = add(*wordsDir)
	case "check":
		err = check(*wordsDir)
	case "recreate":
		err = recreate(*wordsDir)
	case "update":
		err = update(*wordsDir)
	}
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	/*
		args := flag.Args()
		if len(args) != 3 {
			usage("incorrect number of arguments")
		}
	*/
	/*
		credentialsJson := args[0]
		wordsDir := args[1]
		mainLang := args[2] // ISO-639 two letter code.
		if !isDir(wordsDir) {
			usage(fmt.Sprintf("%s is not a directory", wordsDir))
		}
		lang := xlns.LanguageForIso639(mainLang)
		if lang == "" {
			usage(fmt.Sprintf("%s is not an ISO-639 language", mainLang))
		}
		baseWordsFile := path.Join(wordsDir, mainLang+".words")
		if !isFile(baseWordsFile) {
			usage(fmt.Sprintf("%s not found", baseWordsFile))
		}

		_, err := xlns.GetTranslateService(credentialsJson)
		if err != nil {
			log.Fatalf("error %v", err)
		}
	*/

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

func isDir(dir string) error {
	fileInfo, err := os.Stat(dir)
	if err != nil {
		return fmt.Errorf("bad path %s", dir)
	}
	if !fileInfo.IsDir() {
		return fmt.Errorf("%s not directory", dir)
	}
	return nil
}

func isFile(file string) error {
	fileInfo, err := os.Stat(file)
	if err != nil {
		return fmt.Errorf("bad path %s", file)
	}
	if fileInfo.IsDir() {
		return fmt.Errorf("%s not file", file)
	}
	return nil
}

func add(wordsDir string) error {
	fmt.Println("add")
	return fmt.Errorf("add not implemented")
}

func check(wordsDir string) error {
	// Getting a list of languages will find any we don't know about.
	langs, err := xlns.WordsLanguages(wordsDir)
	if err != nil {
		return fmt.Errorf("check got %v", err)
	}
	// Checking each file to the last will make sure that the files have the
	// same number of words (not that the words are good translations).
	last := ""
	for _, lang := range langs {
		current := path.Join(wordsDir, fmt.Sprintf("%s.words", lang))
		fmt.Println(current)
		if last == "" {
			last = current
			continue
		}
		_, err = xlns.XlnsMapFromFiles(last, current)
		if err != nil {
			return fmt.Errorf("check maps got %v", err)
		}
		last = current
	}
	return nil
}

func recreate(wordsDir string) error {
	fmt.Println("recreate")
	return fmt.Errorf("recreate not implemented")
}

func update(wordsDir string) error {
	fmt.Println("update")
	return fmt.Errorf("update not implemented")
}

//func usage() {
//	text := `credentialsJson wordsDir mainLang
//
//    Uses the service information in credentialsJson to access the Google
//    Translate API to update meaning ordered words files in wordsDir.
//
//Where:
//    credentialsJson is a JSON file with Google Service info and keys.
//    wordsDir is a directory containing meaning ordered words files.
//    mainLang is the ISO639 code for the words file to be used to update
//        the others.
//
//Example:
//    translate yourServiceKey.json words en
//`
//	log.Fatalf("ERROR: %s.\nUSAGE:\n    %s %s",
//		why,
//		os.Args[0],
//		text)
//}
//
