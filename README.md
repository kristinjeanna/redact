# github.com/kristinjeanna/redact

[![GitHub license](https://img.shields.io/github/license/kristinjeanna/redact.svg?style=flat&label=License)](https://github.com/kristinjeanna/redact/blob/main/LICENSE) ![Last commit](https://img.shields.io/github/last-commit/kristinjeanna/redact?style=flat&label=Last%20commit) ![Build and test](https://github.com/kristinjeanna/redact/actions/workflows/build.yml/badge.svg?branch=main) ![Latest tag](https://img.shields.io/github/v/tag/kristinjeanna/redact?label=Latest%20tag) [![Go Report Card](https://goreportcard.com/badge/github.com/kristinjeanna/redact)](https://goreportcard.com/report/github.com/kristinjeanna/redact) [![codecov](https://codecov.io/gh/kristinjeanna/redact/branch/main/graph/badge.svg?token=mHRY7hXtrB)](https://codecov.io/gh/kristinjeanna/redact) [![Go Reference](https://pkg.go.dev/badge/github.com/kristinjeanna/redact.svg)](https://pkg.go.dev/github.com/kristinjeanna/redact)

Package `redact` provides a variety of string redactor implementations. The available redactors include: `simple`, `substring`, `blackout`, `middle`, `regex`, `url`, and `chain`.

<details open="open">
<summary>Table of Contents</summary>

- [Install](#install)
- [Overview of redactors](#overview-of-redactors)
  - [`simple`](#simple)
  - [`substring`](#substring)
  - [`blackout`](#blackout)
  - [`middle`](#middle)
  - [`regex`](#regex)
  - [`url`](#url)
  - [`chain`](#chain)

</details>

## Install

```shell
go get -u github.com/kristinjeanna/redact
```

## Overview of redactors

Each redactor implements the `redact.Redactor` interface:

```go
// Redactor is the common interface implemented by all redactors.
type Redactor interface {

    // Redact redacts the input string and returns the result.
    Redact(s string) (string, error)
}
```

### `simple`

The `simple` redactor is a redactor that simply replaces an entire
string with a specified replacement string.

``` go
package main

import (
    "fmt"
    "log"

    "github.com/kristinjeanna/redact/simple"
)

func main() {
    redactor := simple.New("[redacted]")
    input := "this string contains sensitive information"

    result, err := redactor.Redact(input)
    if err != nil {
        log.Fatalf("an error occurred while redacting: %s", err)
    }

    fmt.Println(result)
    // Output: [redacted]
}
```

### `substring`

The `substring` redactor is a redactor that replaces all occurrences of the
substring in the specified string.

```go
package main

import (
    "fmt"
    "log"

    "github.com/kristinjeanna/redact/substring"
)

func main() {
    redactor := substring.New("contains sensitive", "XXXXX")
    input := "this string contains sensitive information"

    result, err := redactor.Redact(input)
    if err != nil {
        log.Fatalf("an error occurred while redacting: %s", err)
    }

    fmt.Println(result)
    // Output: this string XXXXX information
}

```

### `blackout`

The `blackout` redactor strikes out each non-whitespace character of an input
string with a specified replacement string.

The following example uses a single Unicode full block character (U+2588) to
redact the input string:

```go
package main

import (
    "fmt"
    "log"

    "github.com/kristinjeanna/redact/blackout"
)

func main() {
    redactor := blackout.New("█")
    input := "this string contains sensitive information"

    result, err := redactor.Redact(input)
    if err != nil {
        log.Fatalf("an error occurred while redacting: %s", err)
    }

    fmt.Println(result)
    // Output: ████ ██████ ████████ █████████ ███████████
}
```

### `middle`

The `middle` redactor is a redactor that replaces the middle contents of a string with a replacement string, leaving a prefix of unredacted characters and a suffix of unredacted characters if the input string is long enough. For shorter input strings, the redactor uses only a prefix or suffix or just the replacement string itself.

The following example redacts the middle portion of the string:

```go
package main

import (
    "fmt"
    "log"

    "github.com/kristinjeanna/redact/middle"
)

func main() {
    redactor := middle.New()
    input := "this string contains sensitive information"

    result, err := redactor.Redact(input)
    if err != nil {
        log.Fatalf("an error occurred while redacting: %s", err)
    }

    fmt.Println(result)
    // Output: thi[redacted]ion
}
```

The prefix and suffix default to 3 characters but this is configurable when
creating a redactor via the `middle.NewFromOptions` function. Additionally,
a mode option can enable only the prefix or suffix of the input string to
remain unredacted.

```go
package main

import (
    "fmt"
    "log"

    m "github.com/kristinjeanna/redact/middle"
)

func main() {
    redactor, err := m.NewFromOptions(
        m.WithMode(PrefixOnlyMode),
        m.WithReplacementText("XXXXX"),
        m.WithPrefixLength(8),
    )
    if err != nil {
        log.Fatalf("an error occurred while creating redactor: %s", err)
    }
    input := "this string contains sensitive information"

    result, err := redactor.Redact(input)
    if err != nil {
        log.Fatalf("an error occurred while redacting: %s", err)
    }

    fmt.Println(result)
    // Output: this strXXXXX
}
```

### `regex`

The `regex` reactor is a redactor that replaces substrings matching a given
regular expression with a specified replacement string. Multiple pairs
of replacement strings and regular expressions can be specified to chain
the behavior.

The following example redacts the letters "i" and "s" from the input string:

```go
package main

import (
    "fmt"
    "log"

    "github.com/kristinjeanna/redact/regex"
)

func main() {
    // a regex.Pair holds the regex and the replacement string for matches
    pair, err := regex.NewPairUsingSimple("X", "[is]")
    if err != nil {
        log.Fatalf("an error occurred while create regex replacement pair: %s", err)
    }

    redactor, err := regex.New([]Pair{*pair})
    if err != nil {
        log.Fatalf("an error occurred while creating redactor: %s", err)
    }
    input = "this string contains sensitive information"

    result, err := redactor.Redact(input)
    if err != nil {
        log.Fatalf("an error occurred while redacting: %s", err)
    }

    fmt.Println(result)
    // Output: thXX XtrXng contaXnX XenXXtXve XnformatXon
}
```

### `url`

The `url` redactor enables redacting a password and, optionally a
username, from a URL string.

The following example redacts the password from a MySQL connection string:

```go
package main

import (
    "fmt"
    "log"

    "github.com/kristinjeanna/redact/url"
)

func main() {
    redactor := url.New("REDACTED", nil)
    input := "mysql://user:password@localhost:3306"

    result, err := redactor.Redact(input)
    if err != nil {
        log.Fatalf("an error occurred while redacting: %s", err)
    }

    fmt.Println(result)
    // Output: mysql://user:REDACTED@localhost:3306
}

```

### `chain`

The `chain` redactor consists of a slice of redactors that each redact an
input string in the order that they appear in the slice.

The following example chains a [`substring`](#substring) redactor with a
[`regex`](#regex) redactor:

```go
package main

import (
    "fmt"
    "log"

    "github.com/kristinjeanna/redact"
    "github.com/kristinjeanna/redact/chain"
    "github.com/kristinjeanna/redact/regex"
    "github.com/kristinjeanna/redact/substring"
)

func main() {
    substringRedactor := substring.New("contains", "HIDES")

    regexPair, err := regex.NewPair(" [redacted] ", `(?i)\ss\w*\s`)
    if err != nil {
        log.Fatalf("an error occurred while creating regex pair: %s", err)
    }
    regexRedactor, err := regex.New([]regex.Pair{*regexPair})
    if err != nil {
        log.Fatalf("an error occurred while creating regex redactor: %s", err)
    }

    chainRedactor := chain.New([]redact.Redactor{substringRedactor, regexRedactor})
    input := "this string contains sensitive information"

    result, err := chainRedactor.Redact(input)
    if err != nil {
        log.Fatalf("an error occurred while redacting: %s", err)
    }

    fmt.Println(result)
    // Output: this [redacted] HIDES [redacted] information
}
```
