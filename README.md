# translate

Tools to maintain meaning ordered word files.

## translate binary

translate is a tool for managing meaning ordered words files.

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
	recreate  [-credentials credentialsJson] mainLang
	  Recreate wordsDir meaning ordered words files.
	supported displayLang
	  Show the current Google supported languages in displayLang.
	update  [-credentials credentialsJson] mainLang
	  Updates all meaning ordered words files in wordsDir.

## More on meaning ordered word files

A meaning ordered word file is a just a list of words and phrases.  The file
names take the form *ISO639.words*.  Where *ISO639* is an ISO-639 two letter language code.

A directory of these files can be maintained with these tools so that for different languages the meanings correspond by line.

## Reference

[Google Translate API V3](https://pkg.go.dev/google.golang.org/api/translate/v3)

## golang

github.com/napcatstudio/translate