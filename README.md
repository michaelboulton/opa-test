# opa-test

Regenerate:

```
bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies -prune=true
bazel run //:gazelle -- fix
```
