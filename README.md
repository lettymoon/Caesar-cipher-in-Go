# Caesar Chiper in Golang

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)

Caesar cipher is one of the simplest and most widely known encryption techniques. It is a type of substitution cipher in which each letter in the plaintext is replaced by a letter some fixed number of positions down the alphabet.

### How to use

```markdown
go run caesar-cipher.go [options]
```

### Options

```markdown
-e (encrypt) encrypt the message
-d (decrypt) decrypt the message
-k (key) key that will be used to encrypt and decrypt the message, if not selected, key 13 will be used as default
-m (message) message that will be encrypted or decrypted
```

### Example

```go run caesar-cipher.go -e -k 5 -m potato```

```
Key = 5

Message original = potato

Message encrypt = utyfyt
```
