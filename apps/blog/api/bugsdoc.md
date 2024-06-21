# this doc is for documenting the bugs

## didn't ack the message after consume it. the server kepting inserting the data after restarting it

```
c.Ack(false)
```

## use the middleware authorizer to audit the blog. but the non-auditor user could audit the blog even without permision. this was because the context didn't get aborted after the role verification

```
func Failed(..){
    defer c.Abort()
    ...
}
```
