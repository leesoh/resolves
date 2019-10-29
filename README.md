# Resolves

Resolves checks whether the provided DNS resolver(s) actually resolve things and, optionally, whether they do it quickly.

## Installation

`go get -u github.com/leesoh/hacks/resolves`

## Usage

```sh
$ cat resolvers.txt
4.2.2.2
8.8.8.8
1.1.1.1
192.168.1.1
...

$ cat resolvers.txt | resolves
4.2.2.2
8.8.8.8
1.1.1.1
...
```

## Resources

- [public-dns.info](https://public-dns.info/nameservers.txt) - big list of public DNS resolvers.
