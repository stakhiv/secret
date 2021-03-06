# Secret [![Build Status](https://travis-ci.org/stakhiv/secret.svg?branch=master)](https://travis-ci.org/stakhiv/secret)
Simple password manager utility. Uses RSA OAEP for password encryption and
SHA256 for key hashing. RSA key is generated on startup and encrypted with
AES256.
`.key` and `.storage` files are stored in `~/.secret` directory.

# Installation
Just `go get` it.

```
go get -u github.com/stakhiv/secret
```

# Usage

When storing password you have two options.

Provide password as an argument:

```
secret store your@mail.com password
```

Or enter it using password prompt:

```
secret store your@mail.com
```

To retreive stored password use:

```
secret get your@mail.com
```

Or you can automatically pipe it into clipboard.
Use `pbcopy`, `clip`, `xclip` on Mac, Windows, Linux respectively.

```
secret get your@mail.com | pbcopy
```

# Obligatory asciinema
[![asciicast](https://asciinema.org/a/bx8a0xm7xfjav90j6s2iba3vd.png)](https://asciinema.org/a/bx8a0xm7xfjav90j6s2iba3vd)
