# Leet code in Golang

## cli

```
echo -e $'alias lc_new_case="cd cmd/make/ && go run main.go"' >> ~/.zshrc && source ~/.zshrc
echo -e $'alias lc_cover="go test -v -coverprofile cover.out ./... && go tool cover -html cover.out -o cover.html"' >> ~/.zshrc && source ~/.zshrc
```

## Coverage

Run lc_cover or access [this link](https://yagobatista.github.io/leet-code-in-go/cover.html) to see the current coverage
