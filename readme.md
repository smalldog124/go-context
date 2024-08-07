# Go context
Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.


```
      handler
      /     \
     /       \
useCaseA    useCaseB
   |           |
   |           |
 repo        service
   |           |
   |           |
  DB         API other
```

# Ref
https://www.somkiat.cc/golang-using-context-package/

https://go.dev/blog/context

https://pkg.go.dev/context