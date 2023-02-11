# GQL(Gen) Armor 🛡

GQL(gen)Armor is a port of [GQL Armor](https://github.com/Escape-Technologies/graphql-armor) security middleware libary for the popular Golang [GQLGen](https://gqlgen.com/) project.

## Contents

- [Contents](#contents)
- [Usage](#usage)
  - [Installation](#installation)
  - [Supported Features](#supported-features)
  - [Examples](#examples)
- [Contributing](#contributing)


## Installation

```bash
go get github.com/kilianstallz/gqlgen-armor
```

## Supported Features

- [x] [Alias Limit]
- [x] [Character Limit]
- [x] [Field Suggestions Filter]
- [x] [Max Complexity Limit] (via [gqlgen extension])
- [] [Max Depth Limit]
- [] [Cost Limit]
- [] [Max Directives]
- [] [Max Tokens]

## Examples

### Default Configuration

```go
package graphql

import (
  "github.com/99designs/gqlgen/graphql/handler"
  "github.com/kilianstallz/gql-armor"
  "github.com/99designs/gqlgen/graphql/handler/extension"
)

func NewGQLServer() *handler.Server {
    srv := handler.NewDefaultServer(resolvers.NewSchema(client, controller))
    srv.SetErrorPresenter(armor.BlockFieldSuggestionPresenter())
    srv.Use(extension.FixedComplexityLimit(30))
    srv.Use(armor.FixedAliasLimit(5))
    srv.Use(armor.FixedCharacterLimit(armor.DefaultCharacterLimit))

	...
	
    return srv
}
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
