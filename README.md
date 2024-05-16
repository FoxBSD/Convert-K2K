# Convert-K2K

This is a small aplication to convert a `Kernel BSD Like` in a mount of files sorted by extension. 

Can be useless to coding but, we can study the languages and logic in most organized enviroment.

> [!WARNING]
> This code occupes a lot of storage because we prefer copy the file instead moving

> [!CAUTION]
> This code want work with some BSD OS code to work and needs the flag `-dir {BSD OS dir}` or the flag `-nbsd` to work in non bsd repository
> `nbsd` assumes than you don't want work with bsd kernel, this make the code walk by all repository

## License

This code is written under license [MIT](https://github.com/FoxBSD/Convert-K2K/blob/main/LICENSE)

## How run this code

Start instealing dependencies like sqlite3 and run: 

```bash
go mod tidy
```

Before this do you can build the code with:

```bash
go build cmd/ck2k/main.go
```

or run interactively

```bash
go run cmd/ck2k/main.go
```

