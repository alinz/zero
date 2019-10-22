# ZERO

Zero is an encoding package which uses zero width chars to hide information in any files. After encoding any files, the encoded file should show an empty file if opens with any text editors.

### Installation

```
go get github.com/alinz/zero
```

### Usage

Zero, implements both `io.Reader` and `io.Writer` which makes it very easy to plug into any streams in Go.

### Example

for a simple example, please refer to `cmd/cli` which show case how to use the library

for encoding a text:

```bash
echo "Hello World!" | go run cmd/cli/main.go > encoded.txt
```

for decoding:

```bash
cat encoded.txt | go run cmd/cli/main.go --decode
```
