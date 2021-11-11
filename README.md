### YS - MemoliDB - Golang Backend Developer Study Case

#### About
It is a simple in memory database, it can be configured to store snapshots of in-memory data.

#### Installation
First copy .env_example to .env and set your variables if you'd like.

You can run this project via Docker Compose  
```
    docker-compose up -d 
```

or since it doesn't have any other heavy dependencies **you can run with** 

```
go run cmd/app/main.go
```

#### Explore

You can explore rest API endpoints by visiting swagger endpoint

```
    localhost:8080/swagger/index.html
```

#### Test

You can run all unit test

```
   go test ./... 
```

#### Linters
This project developed by support of [golangci-lint](https://golangci-lint.run/usage/quick-start) linter, so you can setup on your machine and lint.

#### Contributions
Feel free to ask any questions or contribute to this project. 
Just fork it and create a PR, and we can check and discuss improvement

#### License
MIT License

Copyright (c) [2021] [Recep Can]

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

