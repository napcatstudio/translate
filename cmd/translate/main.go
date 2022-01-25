// translate updates meaning ordered words files.
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
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
	add mainLang newLang
	  Add a new meaning ordered words file for newLang based on mainLang to
	  wordsDir.
	check
	  Quick wordsDir check.  Does not check translation accuracy just
	  consistency.  Does not call the Google Translate API.
	supported displayLang
	  Show the current Google supported languages in displayLang.
	update mainLang
	  Updates all meaning ordered words files in wordsDir.  Effectively,
	  calls add on each existing non-mainLang language.

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
		if len(args) != 3 {
			fatal_usage(fmt.Errorf("bad mainLang newLang"))
		}
		err = add(*wordsDir, *credentialsJson, args[1], args[2])
	case "check":
		err = check(*wordsDir)
	case "supported":
		if len(args) != 2 {
			fatal_usage(fmt.Errorf("bad displayLang"))
		}
		err = supported(*credentialsJson, args[1])
	case "update":
		if len(args) != 2 {
			fatal_usage(fmt.Errorf("bad mainLang"))
		}
		err = update(*wordsDir, *credentialsJson, args[1])
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(2)
	}
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

func add(wordsDir, credentialsJson, mainLang, newLang string) error {
	words, err := xlns.WordsGetWords(wordsDir, mainLang)
	if err != nil {
		return err
	}
	ctx := context.Background()
	option := option.WithCredentialsFile(credentialsJson)
	client, err := tr.NewTranslationClient(ctx, option)
	if err != nil {
		return fmt.Errorf("new client got %v", err)
	}
	defer client.Close()
	parent, err := parent(credentialsJson)
	if err != nil {
		return err
	}
	req := &trpb.TranslateTextRequest{
		Parent:             parent,
		SourceLanguageCode: mainLang,
		TargetLanguageCode: newLang,
		MimeType:           "text/plain", // Mime type plain or html
		Contents:           words,
	}
	resp, err := client.TranslateText(ctx, req)
	if err != nil {
		return fmt.Errorf("translate text got  %v", err)
	}
	translated := make([]string, len(words))
	for i, translation := range resp.GetTranslations() {
		translated[i] = translation.GetTranslatedText()
	}
	err = xlns.WordsWriteWords(wordsDir, newLang, translated)
	return nil
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

func supported(credentialsJson, lang string) error {
	ctx := context.Background()
	option := option.WithCredentialsFile(credentialsJson)
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

func update(wordsDir, credentialsJson, mainLang string) error {
	words, err := xlns.WordsGetWords(wordsDir, mainLang)
	if err != nil {
		return err
	}
	ctx := context.Background()
	option := option.WithCredentialsFile(credentialsJson)
	client, err := tr.NewTranslationClient(ctx, option)
	if err != nil {
		return fmt.Errorf("new client got %v", err)
	}
	defer client.Close()
	parent, err := parent(credentialsJson)
	if err != nil {
		return err
	}
	langs, err := xlns.WordsLanguages(wordsDir)
	if err != nil {
		return err
	}
	for _, lang := range langs {
		if lang == mainLang {
			continue
		}
		req := &trpb.TranslateTextRequest{
			Parent:             parent,
			SourceLanguageCode: mainLang,
			TargetLanguageCode: lang,
			MimeType:           "text/plain", // Mime type plain or html
			Contents:           words,
		}
		resp, err := client.TranslateText(ctx, req)
		if err != nil {
			return fmt.Errorf("translate text got  %v", err)
		}
		translated := make([]string, len(words))
		for i, translation := range resp.GetTranslations() {
			translated[i] = translation.GetTranslatedText()
		}
		err = xlns.WordsWriteWords(wordsDir, lang, translated)
		if err != nil {
			return err
		}
	}
	return nil
}
