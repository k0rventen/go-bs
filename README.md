# go-bs

A bullshit generator written in go.

Inspiration from [here](https://www.makebullshit.com) and [here](https://www.bullshitgenerator.com/)

## build

_src_
```
git clone https://github.com/k0rventen/go-bs
go get
go run 
```

_docker_

```
git clone https://github.com/k0rventen/go-bs
docker build -t go-bs .
docker run -p 8080:8080 go-bs
```

then just `curl localhost:8080`.
Voila.

## License

WTFPL.