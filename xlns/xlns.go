// xlns.go
// Google Cloud API translate functions for words files.
package xlns

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	tr "cloud.google.com/go/translate/apiv3"
	"google.golang.org/api/option"
	trpb "google.golang.org/genproto/googleapis/cloud/translate/v3"
)

const (
	PROJECT_ID = "project_id"
)

// XlnsAdd adds new languages to a meaning ordered words directory.
func XlnsAdd(wordsDir, credentialsJson, mainLang string, newLangs []string) error {
	words, err := WordsGetWords(wordsDir, mainLang)
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
	for _, newLang := range newLangs {
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
		err = WordsWriteWords(wordsDir, newLang, translated)
		if err != nil {
			return fmt.Errorf("writing words for %s", newLang)
		}
	}
	return nil
}

// XlnsSupported outputs the list of Google Cloud Translate API supported
// languages.
func XlnsSupported(credentialsJson, lang string) error {
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

// parent returns the "parent" needed for some API calls.
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

// XlnsUpdate updates all meaning ordered words files based by translating
// from mainLang.  It retranslates the whole file.
func XlnsUpdate(wordsDir, credentialsJson, mainLang string) error {
	words, err := WordsGetWords(wordsDir, mainLang)
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
	langs, err := WordsLanguages(wordsDir)
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
		err = WordsWriteWords(wordsDir, lang, translated)
		if err != nil {
			return err
		}
	}
	return nil
}
