package authz

allow {
    input.url.path == ["/users"]
    input.method == "POST"
}

default deny
