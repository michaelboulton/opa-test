package authz

default allow=false

allow {
    input.path == "/users"
    input.method == "POST"
}
