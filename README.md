# Go XPath Utilities

A utility library for manipulating XML documents with XPath in Go.  

Powered by [`xmlquery`](https://github.com/antchfx/xmlquery).

## Table of contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
  - [Get attribute by XPath](#get-attribute-by-xpath)
  - [Add attribute by XPath](#add-attribute-by-xpath)
  - [Update attribute by XPath](#update-attribute-by-xpath)
  - [Remove by XPath](#remove-by-xpath)
- [License](#license)

## Features

Supports the following XPath operations:

- Retrieve an attribute's value
- Add, update, or remove an attribute
- Remove a node

**Note:** Nodes cannot be retrieved, added, or updated via XPath. This simplifies the utility, as add and update operations require a string value.

## Installation

```sh
go get github.com/rmkane/go-xpath-utils
```

## Usage

### Get attribute by XPath

Get an attribute by XPath:

```go
// Get the value of the "buzz" attribute in the first <fizz> element in foo.xml
value, err := xpathutils.GetAttrByXPathFromFile("foo.xml", "", "//fizz/@buzz")
// value = "fizzbuzz"

// For XML strings, use:
value, err := xpathutils.GetAttrByXPathFromString(xmlStr, "//fizz/@buzz")
```

### Add attribute by XPath

Add an attribute by XPath:

```go
// Add a "buzz" attribute with value "fizzbuzz" to the first <fizz> element
err := xpathutils.AddAttrByXPathFromFile("foo.xml", "", "//fizz/@buzz", "fizzbuzz")

// For XML strings, use:
updatedXML, err := xpathutils.AddAttrByXPathFromString(xmlStr, "//fizz/@buzz", "fizzbuzz")
```

### Update attribute by XPath

Update an attribute by XPath:

```go
// Update the "buzz" attribute in the first <fizz> element in foo.xml to "xyz"
err := xpathutils.UpdateAttrByXPathFromFile("foo.xml", "", "//fizz/@buzz", "xyz")

// For XML strings, use:
updatedXML, err := xpathutils.UpdateAttrByXPathFromString(xmlStr, "//fizz/@buzz", "xyz")
```

### Remove by XPath

Remove an attribute by XPath:

```go
/// Remove the "buzz" attribute from the first <fizz> element in foo.xml
err := xpathutils.RemoveAttrByXPathFromFile("foo.xml", "", "//fizz/@buzz")

// For XML strings, use:
updatedXML, err := xpathutils.RemoveAttrByXPathFromString(xmlStr, "//fizz/@buzz")
```

Remove an element by XPath:

```go
// Remove the first <bar> element from foo.xml
err := xpathutils.RemoveNodeByXPathFromFile("foo.xml", "", "//bar")

// For XML strings, use:
updatedXML, err := xpathutils.RemoveNodeByXPathFromString(xmlStr, "//bar")
```

## License

This project is licensed under the MIT License.
