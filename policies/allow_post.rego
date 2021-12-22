package authz

allow {
    input.path == ["users"]
    input.method == "POST"
}
