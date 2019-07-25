
Tricky Go
=========

A collection of runnable pieces of Go code which demonstrate tricky aspects of the language.

The approach is basically similar to the famous [Go by Example](https://gobyexample.com/), with the following
differencies:

1. Tricky Go is about tricky and less obvious aspects only.
2. All examples are runnable.
3. Setup requirements are minimal &mdash; all you need is Go.

## Table of contents

### Reflection

* [JSON decoder skips unexported fields](https://godoc.org/github.com/dadooda/trickygo#JsonDecoderSkipsUnexportedFields)
* [Reflection cheatsheet](https://godoc.org/github.com/dadooda/trickygo#ReflectionCheatsheet)

You can also browse [full project's documentation on Godoc](https://godoc.org/github.com/dadooda/trickygo).

## Local setup

1. Clone.
2. To run tests, do a: `./Test`.
3. To serve a local documentation server, do a: `./ServeDocs`.

## Contributing

If you've discovered other tricky parts of Go you'd like to share, I invite you to collaborate:

1. Fork, clone.
2. Following existing code and tests form, add your case to the tree (e.g. <a href="reflection.go">code</a> and <a href="reflection_test.go">tests</a>).
3. Run tests.
4. Push, send me a pull request.

## Cheers!

Feedback of any kind is highly appreciated.

&mdash; Alex Fortuna, &copy; 2019
