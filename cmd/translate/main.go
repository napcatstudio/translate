// translate updates meaning ordered words files.
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	tr "cloud.google.com/go/translate/apiv3"
	"github.com/napcatstudio/translate/xlns"
	"google.golang.org/api/option"
	trpb "google.golang.org/genproto/googleapis/cloud/translate/v3"
)

const (
	PROJECT_ID = "project_id"
	USAGE      = `translate is a tool for managing meaning ordered words files.

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
	add [-credentials credentialsJson] mainLang newLang
	  Add a new meaning ordered words file for newLang based on mainLang to
	  wordsDir.
	check
	  Quick wordsDir check.  Does not check translation accuracy just
	  consistency.  Does not call the Google Translate API.
	recreate  [-credentials credentialsJson] mainLang
	  Recreate wordsDir meaning ordered words files.
	supported displayLang
	  Show the current Google supported languages in displayLang.
	update  [-credentials credentialsJson] mainLang
	  Updates all meaning ordered words files in wordsDir.

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
		err = add(*wordsDir)
	case "check":
		err = check(*wordsDir)
	case "recreate":
		err = recreate(*wordsDir)
	case "supported":
		if len(args) != 2 {
			fatal_usage(fmt.Errorf("bad displayLang"))
		}
		err = supported(*credentialsJson, args[1])
	case "update":
		err = update(*wordsDir)
	}
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println("it worked?")
}

func fatal_usage(err error) {
	fmt.Fprintf(os.Stderr, "error: %v", err)
	flag.Usage()
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

func supported(credentialsJson string, lang string) error {
	ctx := context.Background()
	option := option.WithCredentialsFile(credentialsJson)
	//fmt.Printf("%#v\n", option)
	client, err := tr.NewTranslationClient(ctx, option)
	if err != nil {
		return fmt.Errorf("new client got %v", err)
	}
	defer client.Close()
	parent, err := parent(credentialsJson)
	if err != nil {
		return err
	}
	req := &trpb.GetSupportedLanguagesRequest{
		Parent:              parent,
		DisplayLanguageCode: lang}
	langs, err := client.GetSupportedLanguages(ctx, req)
	if err != nil {
		return fmt.Errorf("supported languages got %v", err)
	}
	for _, lang := range langs.Languages {
		fmt.Printf("\t%s %s\n", lang.LanguageCode, lang.DisplayName)
	}
	return nil
}

func parent(credentialsJson string) (string, error) {
	credentials, err := ioutil.ReadFile(credentialsJson)
	if err != nil {
		return "", fmt.Errorf("reading credentials got %v", err)
	}
	data := make(map[string]string)
	err = json.Unmarshal(credentials, &data)
	if err != nil {
		return "", fmt.Errorf("unpacking credentials got %v", err)
	}
	id, ok := data[PROJECT_ID]
	if !ok {
		return "", fmt.Errorf("no %s in credentials", PROJECT_ID)
	}
	return fmt.Sprintf("projects/%s/locations/global", id), nil
}

func update(wordsDir string) error {
	fmt.Println("update")
	return fmt.Errorf("update not implemented")
}
