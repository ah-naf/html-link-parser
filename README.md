# HTML Link Parser

This project is an HTML link parser written in Go. It extracts all links (anchor tags) from an HTML file and outputs the `href` and associated text content.

## Features

- Parses an HTML file to extract links.
- Outputs each link's `href` and the corresponding text content.

## Installation

1. Clone the repository or download the source code.

```bash
git clone https://github.com/ah-naf/html-link-parser
cd html-link-parser
```

2. Build the project:

```bash
go build -o html-link-parser
```

## Usage

1. Prepare an HTML file (e.g., `index.html`) that contains the content you want to parse.

2. Run the program using the following command:

```bash
./html-link-parser -file <path-to-html-file>
```

For example:

```bash
./html-link-parser -file index.html
```

## Code Overview

### `main.go`

The main program:

- Accepts an HTML file path via the `-file` flag.
- Opens the file and parses it to extract links.
- Prints the `href` and text content of each link.

### Key Functions

#### `Parse(r io.Reader) ([]Link, error)`

Parses the HTML content from the provided `io.Reader` and returns a slice of `Link` structs containing the `href` and text content.

#### `getLink(n *html.Node) []Link`

Recursively traverses the HTML nodes to extract anchor (`<a>`) tags and their associated text content.

#### `extractText(n *html.Node) string`

Recursively extracts text content from a given HTML node.

### Example Output

Given the following HTML content in `index.html`:

```html
<!DOCTYPE html>
<html>
  <body>
    <a href="https://example.com">Example</a>
    <div>
      <a href="https://another.com">Another Example</a>
    </div>
  </body>
</html>
```

The program will output:

```plaintext
{https://example.com Example}
{https://another.com Another Example}
```

## Dependencies

This project uses the `golang.org/x/net/html` package for parsing HTML.

