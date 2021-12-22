known_versions = {
    "0.35.0": "a60742685950163f565e896891f300ab0cea13e1bb26bd10a3656287c1537894",
}

def _download_opa_impl(repository_ctx):
    opa_version = repository_ctx.attr.version
    if opa_version not in known_versions:
        if repository_ctx.attr.sha256 == None:
            fail("Need to specify sha for version {}".format(opa_version))
        else:
            sha = repository_ctx.attr.sha256
    else:
        sha = known_versions[opa_version]

    url = "https://github.com/open-policy-agent/opa/releases/download/v{version}/opa_linux_amd64".format(version = opa_version)
    repository_ctx.download(url, sha256 = sha, executable = True)

download_opa = repository_rule(
    implementation = _download_opa_impl,
    attrs = {
        "version": attr.string(default = "0.35.0"),
        "sha256": attr.string(),
    },
)
