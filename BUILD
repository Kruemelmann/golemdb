load("@bazel_gazelle//:def.bzl", "gazelle")
load("@io_bazel_rules_docker//container:container.bzl", "container_image", "container_push")

# gazelle : exclude build
# gazelle:prefix github.com/kruemelmann/golemdb/
gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)

container_image(
    name = "golemdb",
    base = "@alpine_linux_amd64//image",
    files = [
        "//cmd/golemdb:crosslinux_golemdb",
    ],
)

