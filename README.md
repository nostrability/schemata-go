# schemata-go

Go data package for [Nostr](https://nostr.com/) protocol JSON schemas. Embeds compiled JSON schema files from [nostrability/schemata](https://github.com/nostrability/schemata) and provides a registry lookup.

## Usage

```go
import schemata "github.com/nostrability/schemata-go"

schema, ok := schemata.Get("kind1Schema")
if ok {
    // schema is a json.RawMessage
}

keys := schemata.Keys() // all available schema keys, sorted
```

## API

| Function | Description |
|----------|-------------|
| `Get(key) (json.RawMessage, bool)` | Look up a schema by registry key |
| `Keys() []string` | List all available schema keys |

## License

GPL-3.0-or-later
