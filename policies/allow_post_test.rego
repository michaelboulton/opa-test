package authz

test_post_allowed {
    allow with input as {"path": "/users", "method": "POST"}
}
