# Go XPath Utilities

This project provides utility functions for working with XML documents using XPath expressions in Go.

## Usage

Drop an element by XPath:

```go
// Remove the first <bar> element from foo.xml
err := xpath.RemoveByXPathFromFile("foo.xml", "//bar")
```

Drop an element attribute by XPath:

```go
// Remove the buzz attribute from the first <fizz> element in foo.xml
err := xpath.RemoveByXPathFromFile("foo.xml", "//fizz/@buzz")
```

Update an element value by XPath:

```go
// Replace the value of the first <bar> element in foo.xml with "baz"
err := xpath.UpdateByXPathFromFile("foo.xml", "//bar", "baz")
```

Update an element attribute by XPath:

```go
// Replace the value of the buzz attribute in the first <fizz> element
// in foo.xml with "fizbuzz"
err := xpath.UpdateByXPathFromFile("foo.xml", "//fizz/@buzz", "fizbuzz")
```

## License

This project is licensed under the MIT License.
