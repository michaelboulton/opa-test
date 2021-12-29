package authz

test_post_allowed {
    allow with input as {"path": "/users", "method": "POST"}
}

test_post_allowed_2 {
    not allow with input as {"path": "/users2", "method": "POST"}
}
