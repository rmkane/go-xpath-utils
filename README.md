# Go XPath Utilities

This project provides utility functions for working with XML documents using XPath expressions in Go.  
It allows adding, updating, and removing elements or attributes in **XML files or XML strings**.

## Installation

```sh
go get github.com/rmkane/go-xpath-utils
```

## Usage

### Add

Add an element by XPath:

```go
// Add a <baz>zip</baz> element under the first <bar> element in foo.xml
err := xpath.AddByXPathFromFile("foo.xml", "//bar", "baz", "zip")
```

Add an attribute by XPath:

```go
// Add a "buzz" attribute with value "fizzbuzz" to the first <fizz> element
err := xpath.AddByXPathFromFile("foo.xml", "//fizz/@buzz", "buzz", "fizzbuzz")
```

### Update

Update an element value by XPath:

```go
// Update the value of the first <bar> element in foo.xml to "baz"
err := xpath.UpdateByXPathFromFile("foo.xml", "//bar", "baz")
```

Update an attribute by XPath:

```go
// Update the "buzz" attribute in the first <fizz> element in foo.xml to "fizzbuzz"
err := xpath.UpdateByXPathFromFile("foo.xml", "//fizz/@buzz", "fizzbuzz")
```

### Remove

Remove an element by XPath:

```go
// Remove the first <bar> element from foo.xml
err := xpath.RemoveByXPathFromFile("foo.xml", "//bar")
```

Remove an attribute by XPath:

```go
// Remove the "buzz" attribute from the first <fizz> element in foo.xml
err := xpath.RemoveByXPathFromFile("foo.xml", "//fizz/@buzz")
```

### Working with Strings

These functions also support modifying **XML strings** instead of files:

```go
updatedXML, err := xpath.AddByXPathFromString(xmlStr, "//bar", "baz", "zip")
updatedXML, err := xpath.UpdateByXPathFromString(xmlStr, "//bar", "baz")
updatedXML, err := xpath.RemoveByXPathFromString(xmlStr, "//bar")
```

## License

This project is licensed under the MIT License.
