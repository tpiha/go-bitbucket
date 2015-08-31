# go-bitbucket

go-bitbucket is a Go client library for accessing the Bitbucket API 2.0

### Authentication ###

The go-bitbucket library does not directly handle authentication.  Instead, when
creating a new client, pass an `http.Client` that can handle authentication for
you.  The easiest and recommended way to do this is using the [oauth2](https://github.com/golang/oauth2)
library, but you can always use any other library that provides an
`http.Client`.  If you have an OAuth2 access token, you can use it with oauth2 using:

```go
func main() {
  ts := oauth2.StaticTokenSource(
    &oauth2.Token{AccessToken: "... your access token ..."},
  )
  tc := oauth2.NewClient(oauth2.NoContext, ts)

  client := bitbucket.NewClient(tc)

  // list all repositories for the authenticated user
  repos, _, err := client.Repositories.List("", nil)
}
```

See the [oauth2 docs](https://godoc.org/golang.org/x/oauth2) for complete instructions on using that library.
