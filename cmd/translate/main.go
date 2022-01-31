// translate updates meaning ordered words files.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/napcatstudio/translate/xlns"
)

const (
	USAGE = `translate is a tool for managing meaning ordered words files.

A meaning ordered words file is a file which has words, in one language,
based on another file in a different language.  The file name specifies the
language.  The filename must be of the form XX.words, where XX is a ISO-639
two letter language code.

For instance:

en.words
Easy to use.
Easy.

de.words
Einfach zu gebrauchen.
Einfach.

fi.words
Helppokäyttöinen.
Helppo.

It uses the Google Translate API V3 for translating.

Usage:
	translate [-words wordsDir] [-credentials credentialsJson] command [arguments]

The commands are:
	add mainLang newLang [newLang...]
	  Add a new meaning ordered words file for newLang based on mainLang to
	  wordsDir.
	check
	  Quick wordsDir check.  Does not check translation accuracy just
	  consistency.  Does not call the Google Translate API.
	merge [fromWordsDir]
	  Updates wordsDir with the words in fromWordsDir.
	supported displayLang
	  Show the current Google supported languages in displayLang.
	update mainLang
	  Updates all meaning ordered words files in wordsDir.  Effectively,
	  calls add on each existing non-mainLang language.

Example:
	translate add en es-419 pl

	Uses Google Translate service credentials in 'credentials.json' translates
	  'words/en.words' into ES-419 and PL, and adds them as
	  'words/es-419.words' and 'words/pl.words'.

`
)

func main() {
	wordsDir := flag.String(
		"words", "words", "meaning ordered words directory")
	credentialsJson := flag.String(
		"credentials", "credentials.json", "Google service account information")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, USAGE)
		flag.PrintDefaults()
	}
	flag.Parse()
	err := isDir(*wordsDir)
	if err != nil {
		fatal_usage(fmt.Errorf("bad wordsDir (%v)", err))
	}
	args := flag.Args()
	if len(flag.Args()) < 1 {
		fatal_usage(fmt.Errorf("no command"))
	}

	// Run command.
	switch args[0] {
	case "add":
		if len(args) < 3 {
			fatal_usage(fmt.Errorf("no newLang"))
		}
		err = xlns.XlnsAdd(*wordsDir, *credentialsJson, args[1], args[2:])
	case "check":
		err = xlns.WordsCheck(*wordsDir)
	case "merge":
		if len(args) != 2 {
			fatal_usage(fmt.Errorf("wrong number of arguments"))
		}
		err = isDir(args[1])
		if err != nil {
			fatal_usage(err)
		}
		err = xlns.WordsMerge(*wordsDir, args[1])
	case "supported":
		if len(args) != 2 {
			fatal_usage(fmt.Errorf("bad displayLang"))
		}
		err = xlns.XlnsSupported(*credentialsJson, args[1])
	case "update":
		if len(args) != 2 {
			fatal_usage(fmt.Errorf("bad mainLang"))
		}
		err = xlns.XlnsUpdate(*wordsDir, *credentialsJson, args[1])
	}
	if err != nil {
		fatal(err)
	}
}

func fatal_usage(err error) {
	fmt.Fprintf(os.Stderr, "error: %v", err)
	flag.Usage()
	os.Exit(2)
}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "error: %v", err)
	os.Exit(2)
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
