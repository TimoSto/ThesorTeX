workspace(name = "com_github_tendermint_tendermint")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "56d8c5a5c91e1af73eca71a6fab2ced959b67c86d12ba37feedb0a2dfea441a6",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.37.0/rules_go-v0.37.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.37.0/rules_go-v0.37.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "448e37e0dbf61d6fa8f00aaa12d191745e14f07c31cabfa731f0c8e8a4f41b97",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.28.0/bazel-gazelle-v0.28.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.28.0/bazel-gazelle-v0.28.0.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(
    nogo = "@//:my_nogo",
    version = "1.18.9",
)

http_archive(
    name = "com_github_sluongng_nogo_analyzer",
    sha256 = "ab9ab7936b6d490ff92bb8e3e03bc3ace3406f0b4d1625cc0720d0e9e81a369a",
    strip_prefix = "nogo-analyzer-0.0.1",
    urls = [
        "https://github.com/sluongng/nogo-analyzer/archive/refs/tags/v0.0.1.tar.gz",
    ],
)

load("@com_github_sluongng_nogo_analyzer//staticcheck:deps.bzl", "staticcheck_deps")

staticcheck_deps()

gazelle_dependencies()

http_archive(
    name = "rules_pkg",
    sha256 = "62eeb544ff1ef41d786e329e1536c1d541bb9bcad27ae984d57f18f314018e66",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_pkg/releases/download/0.6.0/rules_pkg-0.6.0.tar.gz",
        "https://github.com/bazelbuild/rules_pkg/releases/download/0.6.0/rules_pkg-0.6.0.tar.gz",
    ],
)

load("@rules_pkg//:deps.bzl", "rules_pkg_dependencies", "rules_pkg_register_toolchains")

rules_pkg_dependencies()

rules_pkg_register_toolchains()

http_archive(
    name = "com_google_protobuf",
    sha256 = "930c2c3b5ecc6c9c12615cf5ad93f1cd6e12d0aba862b572e076259970ac3a53",
    strip_prefix = "protobuf-3.21.12",
    urls = [
        "https://github.com/protocolbuffers/protobuf/archive/v3.21.12.tar.gz",
    ],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

http_archive(
    name = "com_github_bazelbuild_buildtools",
    sha256 = "e3bb0dc8b0274ea1aca75f1f8c0c835adbe589708ea89bf698069d0790701ea3",
    strip_prefix = "buildtools-5.1.0",
    urls = [
        "https://github.com/bazelbuild/buildtools/archive/refs/tags/5.1.0.tar.gz",
    ],
)

##################
# rules_ts/_js setup #
##################
http_archive(
    name = "rules_nodejs",
    sha256 = "08337d4fffc78f7fe648a93be12ea2fc4e8eb9795a4e6aa48595b66b34555626",
    urls = ["https://github.com/bazelbuild/rules_nodejs/releases/download/5.8.0/rules_nodejs-core-5.8.0.tar.gz"],
)

# Fetch and register node, if you haven't already
load("@rules_nodejs//nodejs:repositories.bzl", "nodejs_register_toolchains")

nodejs_register_toolchains(
    name = "node",
    node_version = "16.17.0",
)

http_archive(
    name = "aspect_rules_js",
    sha256 = "9f51475dd2f99abb015939b1cf57ab5f15ef36ca6d2a67104450893fd0aa5c8b",
    strip_prefix = "rules_js-1.16.0",
    url = "https://github.com/aspect-build/rules_js/archive/refs/tags/v1.16.0.tar.gz",
)

load("@aspect_rules_js//js:repositories.bzl", "rules_js_dependencies")

rules_js_dependencies()

http_archive(
    name = "aspect_rules_ts",
    sha256 = "acb20a4e41295d07441fa940c8da9fd02f8637391fd74a14300586a3ee244d59",
    strip_prefix = "rules_ts-1.2.0",
    url = "https://github.com/aspect-build/rules_ts/archive/refs/tags/v1.2.0.tar.gz",
)

load("@aspect_rules_js//npm:npm_import.bzl", "npm_translate_lock")

npm_translate_lock(
    name = "pnpm",
    pnpm_lock = "//:pnpm-lock.yaml",
    verify_node_modules_ignored = "//:.bazelignore",
)

load("@pnpm//:repositories.bzl", "npm_repositories")

npm_repositories()

load("@aspect_rules_ts//ts:repositories.bzl", "rules_ts_dependencies")

rules_ts_dependencies(
    ts_version = "4.9.4",
)
