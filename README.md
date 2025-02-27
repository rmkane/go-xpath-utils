# Go XPath Utilities

This project provides utility functions for working with XML documents using XPath expressions in Go.  
It allows adding, updating, and removing elements or attributes in **XML files or XML strings**.

## Installation

```sh
go get github.com/rmkane/go-xpath-utils
```

## Usage

### Get

Get an attribute by XPath:

```go
// Get the value of the "buzz" attribute in the first <fizz> element in foo.xml
value, err := xpathutils.GetAttrByXPathFromFile("foo.xml", "", "//fizz/@buzz")
// value = "fizzbuzz"
```

### Add

Add an attribute by XPath:

```go
// Add a "buzz" attribute with value "fizzbuzz" to the first <fizz> element
err := xpathutils.AddAttrByXPathFromFile("foo.xml", "", "//fizz/@buzz", "fizzbuzz")
```

### Update

Update an attribute by XPath:

```go
// Update the "buzz" attribute in the first <fizz> element in foo.xml to "fizzbuzz"
err := xpathutils.UpdateAttrByXPathFromFile("foo.xml", "", "//fizz/@buzz", "fizzbuzz")
```

### Remove

Remove an element by XPath:

```go
// Remove the first <bar> element from foo.xml
err := xpathutils.RemoveNodeByXPathFromFile("foo.xml", "", "//bar")
```

Remove an attribute by XPath:

```go
// Remove the "buzz" attribute from the first <fizz> element in foo.xml
err := xpathutils.RemoveAttrByXPathFromFile("foo.xml", "", "//fizz/@buzz")
```

### Working with Strings

These functions also support modifying **XML strings** instead of files:

```go
updatedXML, err := xpathutils.AddByXPathFromString(xmlStr, "//bar", "baz", "zip")
updatedXML, err := xpathutils.UpdateByXPathFromString(xmlStr, "//bar", "baz")
updatedXML, err := xpathutils.RemoveByXPathFromString(xmlStr, "//bar")
```

## License

This project is licensed under the MIT License.
