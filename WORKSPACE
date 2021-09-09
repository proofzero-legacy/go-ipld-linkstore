# WORKSPACE

## Golang
#
# Configure Golang support by:
# - fetching the Bazel rules (rules_go) for building Golang software
# - downloading a recent supported Go toolchain
# - registering the toolchain for use
#
# We have the option of crafting BUILD files by hand, or using
# bazel_gazelle to inspect go.mod files and generating them
# instead. This tool is installed below to provide us that option.
#
# NB: specify a more recent commit of rules_go by using a git_repository().
#
# Further documentation is available at:
#
#   https://github.com/bazelbuild/rules_go/blob/master/proto/core.rst
#

workspace(name = "kubelt")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

# TOOL DEPENDENCIES
# ------------------------------------------------------------------------------

##
## Gazelle (golang BUILD file generator)
##

http_archive(
    name = "bazel_gazelle",
    sha256 = "62ca106be173579c0a167deb23358fdfe71ffa1e4cfdddf5582af26520f1c66f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
    ],
)

##
## Protocol Buffers
##

# fixes nogo bug
# https://github.com/bazelbuild/rules_go/issues/2479
# git_repository(
#     name = "bazel_gazelle",
#     commit = "0ac66c98675a24d58f89a614b84dcd920a7e1762",
#     remote = "https://github.com/bazelbuild/bazel-gazelle",
#     shallow_since = "1626107853 -0400",
# )

# rules_cc defines rules for generating C++ code from Protocol Buffers.
http_archive(
    name = "rules_cc",
    sha256 = "35f2fb4ea0b3e61ad64a369de284e4fbbdcdba71836a5555abb5e194cf119509",
    strip_prefix = "rules_cc-624b5d59dfb45672d4239422fa1e3de1822ee110",
    urls = [
        "https://github.com/bazelbuild/rules_cc/archive/624b5d59dfb45672d4239422fa1e3de1822ee110.tar.gz",
        "https://mirror.bazel.build/github.com/bazelbuild/rules_cc/archive/624b5d59dfb45672d4239422fa1e3de1822ee110.tar.gz",
    ],
)

http_archive(
    name = "com_google_protobuf",
    sha256 = "c6003e1d2e7fefa78a3039f19f383b4f3a61e81be8c19356f85b6461998ad3db",
    strip_prefix = "protobuf-3.17.3",
    urls = [
        "https://github.com/protocolbuffers/protobuf/archive/v3.17.3.tar.gz",
        "https://mirror.bazel.build/github.com/protocolbuffers/protobuf/archive/v3.17.3.tar.gz",
    ],
)

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "8e968b5fcea1d2d64071872b12737bbb5514524ee5f0a4f54f5920266c261acb",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.28.0/rules_go-v0.28.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.28.0/rules_go-v0.28.0.zip",
    ],
)

http_archive(
    name = "rules_proto",
    sha256 = "602e7161d9195e50246177e7c55b2f39950a9cf7366f74ed5f22fd45750cd208",
    strip_prefix = "rules_proto-97d8af4dc474595af3900dd85cb3a29ad28cc313",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_proto/archive/97d8af4dc474595af3900dd85cb3a29ad28cc313.tar.gz",
        "https://github.com/bazelbuild/rules_proto/archive/97d8af4dc474595af3900dd85cb3a29ad28cc313.tar.gz",
    ],
)

##
## Python
##

rules_python_version = "740825b7f74930c62f44af95c9a4c1bd428d2c53"  # Latest @ 2021-06-23

http_archive(
    name = "rules_python",
    sha256 = "09a3c4791c61b62c2cbc5b2cbea4ccc32487b38c7a2cc8f87a794d7a659cc742",
    strip_prefix = "rules_python-{}".format(rules_python_version),
    url = "https://github.com/bazelbuild/rules_python/archive/{}.zip".format(rules_python_version),
)

##
## zlib
##

http_archive(
    name = "zlib",
    build_file = "@com_google_protobuf//:third_party/zlib.BUILD",
    sha256 = "c3e5e9fdd5004dcb542feda5ee4f0ff0744628baf8ed2dd5d66f8ca1197cb1a1",
    strip_prefix = "zlib-1.2.11",
    urls = [
        "https://mirror.bazel.build/zlib.net/zlib-1.2.11.tar.gz",
        "https://zlib.net/zlib-1.2.11.tar.gz",
    ],
)

##
## nodejs / javascript
##

http_archive(
    name = "build_bazel_rules_nodejs",
    sha256 = "8f5f192ba02319254aaf2cdcca00ec12eaafeb979a80a1e946773c520ae0a2c9",
    urls = [
        "https://github.com/bazelbuild/rules_nodejs/releases/download/3.7.0/rules_nodejs-3.7.0.tar.gz",
    ],
)

# INITIALIZATION
# ------------------------------------------------------------------------------

load("@bazel_gazelle//:deps.bzl", "go_repository")
load("@build_bazel_rules_nodejs//:index.bzl", "node_repositories")
load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")
load("@io_bazel_rules_go//extras:embed_data_deps.bzl", "go_embed_data_dependencies")
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

# NB: we declare some dependencies before initalization in order to use
# newer versions than those referenced by the rules.

# go_repository(
#     name = "org_golang_x_tools",
#     importpath = "golang.org/x/tools",
#     sum = "h1:ouewzE6p+/VEB31YYnTbEJdi8pFqKp4P4n85vwo3DHA=",
#     version = "v0.1.5",
# )

##
## Toolchain
##

# Declare indirect dependencies and init toolchains.
go_rules_dependencies()

go_register_toolchains(version = "1.17")

go_embed_data_dependencies()

##
## Protocol Buffers
##

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

rules_proto_dependencies()

rules_proto_toolchains()

protobuf_deps()

##
## Gazelle
##

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

go_repository(
    name = "com_github_sahilm_fuzzy",
    importpath = "github.com/sahilm/fuzzy",
    sum = "h1:FzWGaw2Opqyu+794ZQ9SYifWv2EIXpwP4q8dY1kDAwI=",
    version = "v0.1.0",
)

gazelle_dependencies()

##
## NodeJS
##

node_repositories(
    node_version = "14.17.5",
    package_json = ["//ktrl/plugins/ktrl-web/web:package.json"],
    yarn_version = "1.22.11",
)

##
## Typescript
##

# load("@npm_bazel_typescript//:index.bzl", "ts_setup_workspace")

# ts_setup_workspace()

# GO DEPENDENCIES
# ------------------------------------------------------------------------------
# These are the dependencies required to build the projects in the
# repo. They are referenced from BUILD files by "name" to pull them
# in. As a result, changes here update the dependency version in all
# places where it is used throughout the repo. If it becomes necessary
# to use two separate versions of the dependency, use two different
# names to describe the distinct versions, and update any references to
# use the correct name.
#
# To update a dependency use the gazelle tool. To install it using Go:
#   $ go get github.com/bazelbuild/bazel-gazelle/cmd/gazelle
#
# For the latest available version:
#   $ gazelle update-repos example.com/foo/bar
#
# To require a specific version:
#   $ gazelle update-repos example.com/foo/bar@1.2.3
#
# To import from a go.mod:
#   $ gazelle update-repos -from_file=go.mod

go_repository(
    name = "com_github_whyrusleeping_tar_utils",
    importpath = "github.com/whyrusleeping/tar-utils",
    sum = "h1:wA3QeTsaAXybLL2kb2cKhCAQTHgYTMwuI8lBlJSv5V8=",
    version = "v0.0.0-20201201191210-20a61371de5b",
)

go_repository(
    name = "com_github_warpfork_go_testmark",
    importpath = "github.com/warpfork/go-testmark",
    sum = "h1:Q81c4u7hT+BR5kNfNQhEF0VT2pmL7+Kk0wD+ORYl7iA=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_ipfs_go_ipfs_api",
    importpath = "github.com/ipfs/go-ipfs-api",
    sum = "h1:BXRctUU8YOUOQT/jW1s56d9wLa85ntOqK6bptvCKb8c=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_atotto_clipboard",
    importpath = "github.com/atotto/clipboard",
    sum = "h1:EH0zSVneZPSuFR11BlR9YppQTVDbh5+16AmcJi4g1z4=",
    version = "v0.1.4",
)

go_repository(
    name = "com_github_charmbracelet_bubbles",
    importpath = "github.com/charmbracelet/bubbles",
    sum = "h1:uE0NRa15pVNdnF0EUIwK48ARdRdZkEZ4wLnP/GuoHkI=",
    version = "v0.8.1-0.20210823213054-d987ef84f266",
)

go_repository(
    name = "com_github_charmbracelet_bubbletea",
    importpath = "github.com/charmbracelet/bubbletea",
    sum = "h1:pD/bM5LBEH/nDo7nKcgNUgi4uRHQhpWTIHZbG5vuSlc=",
    version = "v0.14.1",
)

go_repository(
    name = "com_github_charmbracelet_lipgloss",
    importpath = "github.com/charmbracelet/lipgloss",
    sum = "h1:5MysOD6sHr4RP4jkZNWGVIul5GKoOsP12NgbgXPvAlA=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_containerd_console",
    importpath = "github.com/containerd/console",
    sum = "h1:u7SFAJyRqWcG6ogaMAx3KjSTy1e3hT9QxqX7Jco7dRc=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_google_goterm",
    importpath = "github.com/google/goterm",
    sum = "h1:5CjVwnuUcp5adK4gmY6i72gpVFVnZDP2h5TmPScB6u4=",
    version = "v0.0.0-20190703233501-fc88cf888a3f",
)

go_repository(
    name = "com_github_lucasb_eyer_go_colorful",
    importpath = "github.com/lucasb-eyer/go-colorful",
    sum = "h1:1nnpGOrhyZZuNyfu1QjKiUICQ74+3FNCN69Aj6K7nkY=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_muesli_reflow",
    importpath = "github.com/muesli/reflow",
    sum = "h1:IFsN6K9NfGtjeggFP+68I4chLZV2yIKsXJFNZ+eWh6s=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_muesli_termenv",
    importpath = "github.com/muesli/termenv",
    sum = "h1:wnbOaGz+LUR3jNT0zOzinPnyDaCZUQRZj9GxK8eRVl8=",
    version = "v0.9.0",
)

go_repository(
    name = "com_github_rivo_uniseg",
    importpath = "github.com/rivo/uniseg",
    sum = "h1:S1pD9weZBuJdFmowNwbpi7BJ8TNftyUImj/0WQi72jY=",
    version = "v0.2.0",
)

go_repository(
    name = "org_golang_x_crypto",
    importpath = "golang.org/x/crypto",
    sum = "h1:/UOmuWzQfxxo9UtlXMwuQU8CMgg1eZXqTRwkSQJWKOI=",
    version = "v0.0.0-20210711020723-a769d52b0f97",
)

go_repository(
    name = "com_github_azure_go_ansiterm",
    importpath = "github.com/Azure/go-ansiterm",
    sum = "h1:w+iIsaOQNcT7OZ575w+acHgRric5iCyQh+xv+KJ4HB8=",
    version = "v0.0.0-20170929234023-d6e3b3328b78",
)

go_repository(
    name = "com_github_badgerodon_peg",
    importpath = "github.com/badgerodon/peg",
    sum = "h1:77KAMse6RWRpPfVnIZcAtJ/5ZK/oRCeY94ZjIWSbe0g=",
    version = "v0.0.0-20130729175151-9e5f7f4d07ca",
)

go_repository(
    name = "com_github_boltdb_bolt",
    importpath = "github.com/boltdb/bolt",
    sum = "h1:JQmyP4ZBrce+ZQu0dY660FMfatumYDLun9hBCUVIkF4=",
    version = "v1.3.1",
)

go_repository(
    name = "com_github_cayleygraph_cayley",
    importpath = "github.com/cayleygraph/cayley",
    sum = "h1:z+7xkAbg6bKiXJOtOkEG3zCm2K084sr/aGwFV7xcQNs=",
    version = "v0.7.7",
)

go_repository(
    name = "com_github_cayleygraph_quad",
    importpath = "github.com/cayleygraph/quad",
    sum = "h1:+6u09WxA7zg9ILonK8DChwzWKLKsDkjyvX+CXXhI/mM=",
    version = "v1.2.4",
)

go_repository(
    name = "com_github_containerd_continuity",
    importpath = "github.com/containerd/continuity",
    sum = "h1:TP+534wVlf61smEIq1nwLLAjQVEK2EADoW3CX9AuT+8=",
    version = "v0.0.0-20190426062206-aaeac12a7ffc",
)

go_repository(
    name = "com_github_cznic_mathutil",
    importpath = "github.com/cznic/mathutil",
    sum = "h1:oad14P7M0/ZAPSMH1nl1vC8zdKVkA3kfHLO59z1l8Eg=",
    version = "v0.0.0-20170313102836-1447ad269d64",
)

go_repository(
    name = "com_github_d4l3k_messagediff",
    importpath = "github.com/d4l3k/messagediff",
    sum = "h1:ZcAIMYsUg0EAp9X+tt8/enBE/Q8Yd5kzPynLyKptt9U=",
    version = "v1.2.1",
)

go_repository(
    name = "com_github_dennwc_base",
    importpath = "github.com/dennwc/base",
    sum = "h1:xlBzvBNRvkQ1LFI/jom7rr0vZsvYDKtvMM6lIpjFb3M=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_dennwc_graphql",
    importpath = "github.com/dennwc/graphql",
    sum = "h1:QWlaiMNg63HE5qimJd4stjg9l1Ca4BKcgs+UNSWPJ+s=",
    version = "v0.0.0-20180603144102-12cfed44bc5d",
)

go_repository(
    name = "com_github_dlclark_regexp2",
    importpath = "github.com/dlclark/regexp2",
    sum = "h1:1udHhhGkIMplSrLeMJpPN7BHz1Iq2wVBUcb+3fxzhQM=",
    version = "v1.1.4",
)

go_repository(
    name = "com_github_docker_docker",
    importpath = "github.com/docker/docker",
    sum = "h1:zlYHASK/UP+Fs28TyQhgUlOubRKCQ+38o/aa6GgouTs=",
    version = "v0.7.3-0.20180412203414-a422774e593b",
)

go_repository(
    name = "com_github_docker_go_connections",
    importpath = "github.com/docker/go-connections",
    sum = "h1:El9xVISelRB7BuFusrZozjnkIM5YnzCViNKohAFqRJQ=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_docker_go_units",
    importpath = "github.com/docker/go-units",
    sum = "h1:3uh0PgVws3nIA0Q+MwDC8yjEPf9zjRfZZWXZYDct3Tw=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_dop251_goja",
    importpath = "github.com/dop251/goja",
    sum = "h1:cA2OMt2CQ2yq2WhQw16mHv6ej9YY07H4pzfR/z/y+1Q=",
    version = "v0.0.0-20190105122144-6d5bf35058fa",
)

go_repository(
    name = "com_github_flimzy_diff",
    importpath = "github.com/flimzy/diff",
    sum = "h1:ufTsTKcDtlaczpJTo3u1NeYqzuP6oRpy1VwQUIrgmBY=",
    version = "v0.1.6",
)

go_repository(
    name = "com_github_flimzy_kivik",
    importpath = "github.com/flimzy/kivik",
    sum = "h1:URl7e0OnfSvAu3ZHQ5BkvzRZlCmyYuDyWUCcPWIHlU0=",
    version = "v1.8.1",
)

go_repository(
    name = "com_github_flimzy_testy",
    importpath = "github.com/flimzy/testy",
    sum = "h1:nchF7XYCkfHJiZKMRhAVKQp8jzpXFPwJYnSrnFysqlI=",
    version = "v0.1.16",
)

go_repository(
    name = "com_github_fortytw2_leaktest",
    importpath = "github.com/fortytw2/leaktest",
    sum = "h1:u8491cBMTQ8ft8aeV+adlcytMZylmA5nnwwkRZjI8vw=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_fsouza_go_dockerclient",
    importpath = "github.com/fsouza/go-dockerclient",
    sum = "h1:rFDrkgZUIlruULXD2gRhT8JhqbjA6vHszAIStg/juEY=",
    version = "v1.2.2",
)

go_repository(
    name = "com_github_go_kivik_couchdb",
    importpath = "github.com/go-kivik/couchdb",
    sum = "h1:2yjmysS48JYpyWTkx2E3c7ASZP8Kh0eABWnkKlV8bbw=",
    version = "v1.8.1",
)

go_repository(
    name = "com_github_go_kivik_kivik",
    importpath = "github.com/go-kivik/kivik",
    sum = "h1:GScP1mS5wP2km2awszvKzPEjC21lYjQGr3GY+4a/o2U=",
    version = "v1.8.1",
)

go_repository(
    name = "com_github_go_kivik_kiviktest",
    importpath = "github.com/go-kivik/kiviktest",
    sum = "h1:mn93FbOY459cGsYw4xp2t7myGb2YtfvOIivOtMlLaPA=",
    version = "v1.1.2",
)

go_repository(
    name = "com_github_go_kivik_pouchdb",
    importpath = "github.com/go-kivik/pouchdb",
    sum = "h1:PFnvbZJzfTeutcfcozHc649Y44tnvqdbuCTU/ZEZQ6Q=",
    version = "v1.3.5",
)

go_repository(
    name = "com_github_go_sourcemap_sourcemap",
    importpath = "github.com/go-sourcemap/sourcemap",
    sum = "h1:0b/xya7BKGhXuqFESKM4oIiRo9WOt2ebz7KxfreD6ug=",
    version = "v2.1.2+incompatible",
)

go_repository(
    name = "com_github_gobuffalo_envy",
    importpath = "github.com/gobuffalo/envy",
    sum = "h1:OQl5ys5MBea7OGCdvPbBJWRgnhC/fGona6QKfvFeau8=",
    version = "v1.7.1",
)

go_repository(
    name = "com_github_gobuffalo_logger",
    importpath = "github.com/gobuffalo/logger",
    sum = "h1:ZEgyRGgAm4ZAhAO45YXMs5Fp+bzGLESFewzAVBMKuTg=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_gobuffalo_packd",
    importpath = "github.com/gobuffalo/packd",
    sum = "h1:eMwymTkA1uXsqxS0Tpoop3Lc0u3kTfiMBE6nKtQU4g4=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_gobuffalo_packr_v2",
    importpath = "github.com/gobuffalo/packr/v2",
    sum = "h1:n3CIW5T17T8v4GGK5sWXLVWJhCz7b5aNLSxW6gYim4o=",
    version = "v2.7.1",
)

go_repository(
    name = "com_github_gopherjs_jsbuiltin",
    importpath = "github.com/gopherjs/jsbuiltin",
    sum = "h1:atBEgNR1C5+LFkl8ipQtLee9RStheS8YeCSkiYqBhOg=",
    version = "v0.0.0-20180426082241-50091555e127",
)

go_repository(
    name = "com_github_gotestyourself_gotestyourself",
    importpath = "github.com/gotestyourself/gotestyourself",
    sum = "h1:AQwinXlbQR2HvPjQZOmDhRqsv5mZf+Jb1RnSLxcqZcI=",
    version = "v2.2.0+incompatible",
)

go_repository(
    name = "com_github_hidal_go_hidalgo",
    importpath = "github.com/hidal-go/hidalgo",
    sum = "h1:hBE4LGxApbZiV/3YoEPv7uYlUMWOogG1hwtkpiU87zQ=",
    version = "v0.0.0-20190814174001-42e03f3b5eaa",
)

go_repository(
    name = "com_github_jackc_fake",
    importpath = "github.com/jackc/fake",
    sum = "h1:vr3AYkKovP8uR8AvSGGUK1IDqRa5lAAvEkZG1LKaCRc=",
    version = "v0.0.0-20150926172116-812a484cc733",
)

go_repository(
    name = "com_github_jackc_pgx",
    importpath = "github.com/jackc/pgx",
    sum = "h1:Wa90/+qsITBAPkAZjiByeIGHFcj3Ztu+VzrrIpHjL90=",
    version = "v3.3.0+incompatible",
)

go_repository(
    name = "com_github_joho_godotenv",
    importpath = "github.com/joho/godotenv",
    sum = "h1:Zjp+RcGpHhGlrMbJzXTrZZPrWj+1vfm90La1wgB6Bhc=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_linkeddata_gojsonld",
    importpath = "github.com/linkeddata/gojsonld",
    sum = "h1:YP3lfXXYiQV5MKeUqVnxRP5uuMQTLPx+PGYm1UBoU98=",
    version = "v0.0.0-20170418210642-4f5db6791326",
)

go_repository(
    name = "com_github_microsoft_go_winio",
    importpath = "github.com/Microsoft/go-winio",
    sum = "h1:xAfWHN1IrQ0NJ9TBC0KBZoqLjzDTr1ML+4MywiUOryc=",
    version = "v0.4.12",
)

go_repository(
    name = "com_github_nvveen_gotty",
    importpath = "github.com/Nvveen/Gotty",
    sum = "h1:TngWCqHvy9oXAN6lEVMRuU21PR1EtLVZJmdB18Gu3Rw=",
    version = "v0.0.0-20120604004816-cd527374f1e5",
)

go_repository(
    name = "com_github_opencontainers_go_digest",
    importpath = "github.com/opencontainers/go-digest",
    sum = "h1:WzifXhOVOEOuFYOJAW6aQqW0TooG2iki3E3Ii+WN7gQ=",
    version = "v1.0.0-rc1",
)

go_repository(
    name = "com_github_opencontainers_image_spec",
    importpath = "github.com/opencontainers/image-spec",
    sum = "h1:JMemWkRwHx4Zj+fVxWoMCFm/8sYGGrUVojFA6h/TRcI=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_opencontainers_runc",
    importpath = "github.com/opencontainers/runc",
    sum = "h1:GlxAyO6x8rfZYN9Tt0Kti5a/cP41iuiO2yYT0IJGY8Y=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_opencontainers_selinux",
    importpath = "github.com/opencontainers/selinux",
    sum = "h1:AYFJmdZd1xjz5UIb8YpDHthdwAzlM5FVY6PzoNMgAMk=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_ory_dockertest",
    importpath = "github.com/ory/dockertest",
    sum = "h1:VrpM6Gqg7CrPm3bL4Wm1skO+zFWLbh7/Xb5kGEbJRh8=",
    version = "v3.3.4+incompatible",
)

go_repository(
    name = "com_github_peterh_liner",
    importpath = "github.com/peterh/liner",
    sum = "h1:8uaXtUkxiy+T/zdLWuxa/PG4so0TPZDZfafFNNSaptE=",
    version = "v0.0.0-20170317030525-88609521dc4b",
)

go_repository(
    name = "com_github_piprate_json_gold",
    importpath = "github.com/piprate/json-gold",
    sum = "h1:a1vHx7Q1jOO1pjCtKwTI/WCzwaQwRt9VM7apK2uy200=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_pquerna_cachecontrol",
    importpath = "github.com/pquerna/cachecontrol",
    sum = "h1:J9b7z+QKAmPf4YLrFg6oQUotqHQeUNWwkvo7jZp1GLU=",
    version = "v0.0.0-20180517163645-1555304b9b35",
)

go_repository(
    name = "com_github_satori_go_uuid",
    importpath = "github.com/satori/go.uuid",
    sum = "h1:0uYX9dsZ2yD7q2RtLRtPSdGDWzjeM3TbMJP9utgA0ww=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_shopspring_decimal",
    importpath = "github.com/shopspring/decimal",
    sum = "h1:pntxY8Ary0t43dCZ5dqY4YTJCObLY1kIXl0uzMv+7DE=",
    version = "v0.0.0-20180709203117-cd690d0c9e24",
)

go_repository(
    name = "com_github_smartystreets_go_aws_auth",
    importpath = "github.com/smartystreets/go-aws-auth",
    sum = "h1:hp2CYQUINdZMHdvTdXtPOY2ainKl4IoMcpAXEf2xj3Q=",
    version = "v0.0.0-20180515143844-0c1422d1fdb9",
)

go_repository(
    name = "com_github_tidwall_pretty",
    importpath = "github.com/tidwall/pretty",
    sum = "h1:HsD+QiTn7sK6flMKIvNmpqz1qrpP3Ps6jOKIKMooyg4=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_tylertreat_boomfilters",
    importpath = "github.com/tylertreat/BoomFilters",
    sum = "h1:7X4KYG3guI2mPQGxm/ZNNsiu4BjKnef0KG0TblMC+Z8=",
    version = "v0.0.0-20181028192813-611b3dbe80e8",
)

# Forked version of wasmer that includes Bazel fixes. Replaces this
# upstream commented out version:
#
# go_repository(
#     name = "com_github_wasmerio_wasmer_go",
#     importpath = "github.com/wasmerio/wasmer-go",
#     sum = "h1:MnqHoOGfiQ8MMq2RF6wyCeebKOe84G88h5yv+vmxJgs=",
#     version = "v1.0.4",
# )
#
# Cf. https://github.com/bazelbuild/rules_go/issues/2861
#
# TODO if these fixes don't land upstream by the time we need wasmer, we
# will likely need to fork wasmer ourselves.

go_repository(
    name = "com_github_joehattori_wasmer_go",
    importpath = "github.com/joehattori/wasmer-go",
    sum = "h1:XTAa9GhlG5zioyfeXq4Ody85LhvNxhW51uh+KWE4vyQ=",
    version = "v1.0.4",
)

go_repository(
    name = "com_github_xdg_scram",
    importpath = "github.com/xdg/scram",
    sum = "h1:u40Z8hqBAAQyv+vATcGgV0YCnDjqSL7/q/JyPhhJSPk=",
    version = "v0.0.0-20180814205039-7eeb5667e42c",
)

go_repository(
    name = "com_github_xdg_stringprep",
    importpath = "github.com/xdg/stringprep",
    sum = "h1:d9X0esnoa3dFsV0FG35rAT0RIhYFlPq7MiP+DW89La0=",
    version = "v1.0.0",
)

go_repository(
    name = "in_gopkg_olivere_elastic_v5",
    importpath = "gopkg.in/olivere/elastic.v5",
    sum = "h1:21Vu9RMT2qXVLqXINtiOhwVPYz/87+Omsxh/Re+gK4k=",
    version = "v5.0.81",
)

go_repository(
    name = "org_mongodb_go_mongo_driver",
    importpath = "go.mongodb.org/mongo-driver",
    sum = "h1:bHxbjH6iwh1uInchXadI6hQR107KEbgYsMzoblDONmQ=",
    version = "v1.0.4",
)

go_repository(
    name = "tools_gotest",
    importpath = "gotest.tools",
    sum = "h1:VsBPFP1AI068pPrMxtb/S8Zkgf9xEmTLJjfM+P5UIEo=",
    version = "v2.2.0+incompatible",
)

go_repository(
    name = "com_github_go_logr_logr",
    importpath = "github.com/go-logr/logr",
    sum = "h1:K7/B1jt6fIBQVd4Owv2MqGQClcgf0R266+7C/QjRcLc=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_ipld_go_car_v2",
    importpath = "github.com/ipld/go-car/v2",
    sum = "h1:R1oIAPwrGp26mEFzcGf5bfTZAAHDOkaVnZTEVebaWX4=",
    version = "v2.0.2",
)

go_repository(
    name = "com_github_petar_gollrb",
    importpath = "github.com/petar/GoLLRB",
    sum = "h1:1/WtZae0yGtPq+TI6+Tv1WTxkukpXeMlviSxvL7SRgk=",
    version = "v0.0.0-20210522233825-ae3b015fd3e9",
)

go_repository(
    name = "com_github_rs_xid",
    importpath = "github.com/rs/xid",
    sum = "h1:mhH9Nq+C1fY2l1XIpgxIiUOfNpRBYH1kKcr+qfKgjRc=",
    version = "v1.2.1",
)

go_repository(
    name = "com_github_rs_zerolog",
    importpath = "github.com/rs/zerolog",
    sum = "h1:Q3vdXlfLNT+OftyBHsU0Y445MD+8m8axjKgf2si0QcM=",
    version = "v1.21.0",
)

go_repository(
    name = "com_github_whyrusleeping_cbor",
    importpath = "github.com/whyrusleeping/cbor",
    sum = "h1:5HZfQkwe0mIfyDmc1Em5GqlNRzcdtlv4HTNmdpt7XH0=",
    version = "v0.0.0-20171005072247-63513f603b11",
)

go_repository(
    name = "io_opentelemetry_go_otel",
    importpath = "go.opentelemetry.io/otel",
    sum = "h1:eaP0Fqu7SXHwvjiqDq83zImeehOHX8doTvU9AwXON8g=",
    version = "v0.20.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_metric",
    importpath = "go.opentelemetry.io/otel/metric",
    sum = "h1:4kzhXFP+btKm4jwxpjIqjs41A7MakRFUS86bqLHTIw8=",
    version = "v0.20.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_oteltest",
    importpath = "go.opentelemetry.io/otel/oteltest",
    sum = "h1:HiITxCawalo5vQzdHfKeZurV8x7ljcqAgiWzF6Vaeaw=",
    version = "v0.20.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_sdk",
    importpath = "go.opentelemetry.io/otel/sdk",
    sum = "h1:JsxtGXd06J8jrnya7fdI/U/MR6yXA5DtbZy+qoHQlr8=",
    version = "v0.20.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_trace",
    importpath = "go.opentelemetry.io/otel/trace",
    sum = "h1:1DL6EXUdcg95gukhuRRvLDO/4X5THh/5dIV52lqtnbw=",
    version = "v0.20.0",
)

go_repository(
    name = "com_github_golang_mock",
    importpath = "github.com/golang/mock",
    sum = "h1:G5FRp8JnTd7RQH5kemVNlMeyXQAztQ3mOWV95KxsXH8=",
    version = "v1.1.1",
)

# NB: keeping this version because of an issue with 1.38.0 onwards
# around reflection.
go_repository(
    name = "org_golang_google_grpc",
    importpath = "google.golang.org/grpc",
    sum = "h1:2dTRdpdFEEhJYQD8EMLB61nnrzSCTbG38PhqdhvOltg=",
    version = "v1.26.0",
)

go_repository(
    name = "com_github_fatih_color",
    importpath = "github.com/fatih/color",
    sum = "h1:DkWD4oS2D8LGGgTQ6IvwJJXSL5Vp2ffcQg58nFV38Ys=",
    version = "v1.7.0",
)

go_repository(
    name = "org_golang_google_protobuf",
    importpath = "google.golang.org/protobuf",
    sum = "h1:SnqbnDw1V7RiZcXPx5MEeqPv2s79L9i7BJUlG/+RurQ=",
    version = "v1.27.1",
)

go_repository(
    name = "cc_mvdan_gofumpt",
    importpath = "mvdan.cc/gofumpt",
    sum = "h1:bi/1aS/5W00E2ny5q65w9SnKpWEF/UIOqDYBILpo9rA=",
    version = "v0.1.1",
)

go_repository(
    name = "cc_mvdan_interfacer",
    importpath = "mvdan.cc/interfacer",
    sum = "h1:WX1yoOaKQfddO/mLzdV4wptyWgoH/6hwLs7QHTixo0I=",
    version = "v0.0.0-20180901003855-c20040233aed",
)

go_repository(
    name = "cc_mvdan_lint",
    importpath = "mvdan.cc/lint",
    sum = "h1:DxJ5nJdkhDlLok9K6qO+5290kphDJbHOQO1DFFFTeBo=",
    version = "v0.0.0-20170908181259-adc824a0674b",
)

go_repository(
    name = "cc_mvdan_unparam",
    importpath = "mvdan.cc/unparam",
    sum = "h1:HT3e4Krq+IE44tiN36RvVEb6tvqeIdtsVSsxmNPqlFU=",
    version = "v0.0.0-20210104141923-aac4ce9116a7",
)

go_repository(
    name = "co_honnef_go_tools",
    importpath = "honnef.co/go/tools",
    sum = "h1:qTakTkI6ni6LFD5sBwwsdSO+AQqbSIxOauHTTQKZ/7o=",
    version = "v0.1.3",
)

go_repository(
    name = "com_4d63_gochecknoglobals",
    importpath = "4d63.com/gochecknoglobals",
    sum = "h1:wFEQiK85fRsEVF0CRrPAos5LoAryUsIX1kPW/WrIqFw=",
    version = "v0.0.0-20201008074935-acfc0b28355a",
)

go_repository(
    name = "com_github_adrg_xdg",
    importpath = "github.com/adrg/xdg",
    sum = "h1:s/tV7MdqQnzB1nKY8aqHvAMD+uCiuEDzVB5HLRY849U=",
    version = "v0.3.3",
)

go_repository(
    name = "com_github_alecthomas_template",
    importpath = "github.com/alecthomas/template",
    sum = "h1:JYp7IbQjafoB+tBA3gMyHYHrpOtNuDiK/uB5uXxq5wM=",
    version = "v0.0.0-20190718012654-fb15b899a751",
)

go_repository(
    name = "com_github_alecthomas_units",
    importpath = "github.com/alecthomas/units",
    sum = "h1:Hs82Z41s6SdL1CELW+XaDYmOH4hkBN4/N9og/AsOv7E=",
    version = "v0.0.0-20190717042225-c3de453c63f4",
)

go_repository(
    name = "com_github_alexkohler_prealloc",
    importpath = "github.com/alexkohler/prealloc",
    sum = "h1:Hbq0/3fJPQhNkN0dR95AVrr6R7tou91y0uHG5pOcUuw=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_andybalholm_brotli",
    importpath = "github.com/andybalholm/brotli",
    sum = "h1:7UCwP93aiSfvWpapti8g88vVVGp2qqtGyePsSuDafo4=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_antihax_optional",
    importpath = "github.com/antihax/optional",
    sum = "h1:xK2lYat7ZLaVVcIuj82J8kIro4V6kDe0AUDFboUCwcg=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_aokoli_goutils",
    importpath = "github.com/aokoli/goutils",
    sum = "h1:7fpzNGoJ3VA8qcrm++XEE1QUe0mIwNeLa02Nwq7RDkg=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_armon_circbuf",
    importpath = "github.com/armon/circbuf",
    sum = "h1:QEF07wC0T1rKkctt1RINW/+RMTVmiwxETico2l3gxJA=",
    version = "v0.0.0-20150827004946-bbbad097214e",
)

go_repository(
    name = "com_github_armon_consul_api",
    importpath = "github.com/armon/consul-api",
    sum = "h1:G1bPvciwNyF7IUmKXNt9Ak3m6u9DE1rF+RmtIkBpVdA=",
    version = "v0.0.0-20180202201655-eb2c6b5be1b6",
)

go_repository(
    name = "com_github_armon_go_metrics",
    importpath = "github.com/armon/go-metrics",
    sum = "h1:8GUt8eRujhVEGZFFEjBj46YV4rDjvGrNxb0KMWYkL2I=",
    version = "v0.0.0-20180917152333-f0300d1749da",
)

go_repository(
    name = "com_github_armon_go_radix",
    importpath = "github.com/armon/go-radix",
    sum = "h1:BUAU3CGlLvorLI26FmByPp2eC2qla6E1Tw+scpcg/to=",
    version = "v0.0.0-20180808171621-7fddfc383310",
)

go_repository(
    name = "com_github_ashanbrown_forbidigo",
    importpath = "github.com/ashanbrown/forbidigo",
    sum = "h1:SJOPJyqsrVL3CvR0veFZFmIM0fXS/Kvyikqvfphd0Z4=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_ashanbrown_makezero",
    importpath = "github.com/ashanbrown/makezero",
    sum = "h1:27owMIbvO33XL56BKWPy+SCU69I9wPwPXuMf5mAbVGU=",
    version = "v0.0.0-20210308000810-4155955488a0",
)

go_repository(
    name = "com_github_aws_aws_sdk_go",
    importpath = "github.com/aws/aws-sdk-go",
    sum = "h1:0xphMHGMLBrPMfxR2AmVjZKcMEESEgWF8Kru94BNByk=",
    version = "v1.27.0",
)

go_repository(
    name = "com_github_beorn7_perks",
    importpath = "github.com/beorn7/perks",
    sum = "h1:VlbKKnNfV8bJzeqoa4cOKqO6bYr3WgKZxO8Z16+hsOM=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_bgentry_speakeasy",
    importpath = "github.com/bgentry/speakeasy",
    sum = "h1:ByYyxL9InA1OWqxJqqp2A5pYHUrCiAL6K3J+LKSsQkY=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_bketelsen_crypt",
    importpath = "github.com/bketelsen/crypt",
    sum = "h1:w/jqZtC9YD4DS/Vp9GhWfWcCpuAL58oTnLoI8vE9YHU=",
    version = "v0.0.4",
)

go_repository(
    name = "com_github_bkielbasa_cyclop",
    importpath = "github.com/bkielbasa/cyclop",
    sum = "h1:7Jmnh0yL2DjKfw28p86YTd/B4lRGcNuu12sKE35sM7A=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_bombsimon_wsl_v3",
    importpath = "github.com/bombsimon/wsl/v3",
    sum = "h1:Mka/+kRLoQJq7g2rggtgQsjuI/K5Efd87WX96EWFxjM=",
    version = "v3.3.0",
)

go_repository(
    name = "com_github_burntsushi_toml",
    importpath = "github.com/BurntSushi/toml",
    sum = "h1:WXkYYl6Yr3qBf1K79EBnL4mak0OimBfB0XUf9Vl28OQ=",
    version = "v0.3.1",
)

go_repository(
    name = "com_github_burntsushi_xgb",
    importpath = "github.com/BurntSushi/xgb",
    sum = "h1:1BDTz0u9nC3//pOCMdNH+CiXJVYJh5UQNCOBG7jbELc=",
    version = "v0.0.0-20160522181843-27f122750802",
)

go_repository(
    name = "com_github_census_instrumentation_opencensus_proto",
    importpath = "github.com/census-instrumentation/opencensus-proto",
    sum = "h1:glEXhBS5PSLLv4IXzLA5yPRVX4bilULVyxxbrfOtDAk=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_cespare_xxhash",
    importpath = "github.com/cespare/xxhash",
    sum = "h1:a6HrQnmkObjyL+Gs60czilIUGqrzKutQD6XZog3p+ko=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_cespare_xxhash_v2",
    importpath = "github.com/cespare/xxhash/v2",
    sum = "h1:6MnRN8NT7+YBpUIWxHtefFZOKTAPgGjpQSxqLNn0+qY=",
    version = "v2.1.1",
)

go_repository(
    name = "com_github_charithe_durationcheck",
    importpath = "github.com/charithe/durationcheck",
    sum = "h1:Tsy7EppNow2pDC0jN7Hsmcb6mHd71ZbI1vFissRBtc0=",
    version = "v0.0.6",
)

go_repository(
    name = "com_github_chavacava_garif",
    importpath = "github.com/chavacava/garif",
    sum = "h1:StHNkfM8nXnNQnk5/0uYYhIqvvENd14hoHPnZsakTNo=",
    version = "v0.0.0-20210405163807-87a70f3d418b",
)

go_repository(
    name = "com_github_chzyer_logex",
    importpath = "github.com/chzyer/logex",
    sum = "h1:Swpa1K6QvQznwJRcfTfQJmTE72DqScAa40E+fbHEXEE=",
    version = "v1.1.10",
)

go_repository(
    name = "com_github_chzyer_readline",
    importpath = "github.com/chzyer/readline",
    sum = "h1:fY5BOSpyZCqRo5OhCuC+XN+r/bBCmeuuJtjz+bCNIf8=",
    version = "v0.0.0-20180603132655-2972be24d48e",
)

go_repository(
    name = "com_github_chzyer_test",
    importpath = "github.com/chzyer/test",
    sum = "h1:q763qf9huN11kDQavWsoZXJNW3xEE4JJyHa5Q25/sd8=",
    version = "v0.0.0-20180213035817-a1ea475d72b1",
)

go_repository(
    name = "com_github_client9_misspell",
    importpath = "github.com/client9/misspell",
    sum = "h1:ta993UF76GwbvJcIo3Y68y/M3WxlpEHPWIGDkJYwzJI=",
    version = "v0.3.4",
)

go_repository(
    name = "com_github_cncf_udpa_go",
    importpath = "github.com/cncf/udpa/go",
    sum = "h1:cqQfy1jclcSy/FwLjemeg3SR1yaINm74aQyupQ0Bl8M=",
    version = "v0.0.0-20201120205902-5459f2c99403",
)

go_repository(
    name = "com_github_cockroachdb_apd",
    importpath = "github.com/cockroachdb/apd",
    sum = "h1:3LFP3629v+1aKXU5Q37mxmRxX/pIu1nijXydLShEq5I=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_cockroachdb_apd_v2",
    importpath = "github.com/cockroachdb/apd/v2",
    sum = "h1:y1Rh3tEU89D+7Tgbw+lp52T6p/GJLpDmNvr10UWqLTE=",
    version = "v2.0.1",
)

go_repository(
    name = "com_github_cockroachdb_datadriven",
    importpath = "github.com/cockroachdb/datadriven",
    sum = "h1:OaNxuTZr7kxeODyLWsRMC+OD03aFUH+mW6r2d+MWa5Y=",
    version = "v0.0.0-20190809214429-80d97fb3cbaa",
)

go_repository(
    name = "com_github_coreos_bbolt",
    importpath = "github.com/coreos/bbolt",
    sum = "h1:wZwiHHUieZCquLkDL0B8UhzreNWsPHooDAG3q34zk0s=",
    version = "v1.3.2",
)

go_repository(
    name = "com_github_coreos_etcd",
    importpath = "github.com/coreos/etcd",
    sum = "h1:jFneRYjIvLMLhDLCzuTuU4rSJUjRplcJQ7pD7MnhC04=",
    version = "v3.3.10+incompatible",
)

go_repository(
    name = "com_github_coreos_go_etcd",
    importpath = "github.com/coreos/go-etcd",
    sum = "h1:bXhRBIXoTm9BYHS3gE0TtQuyNZyeEMux2sDi4oo5YOo=",
    version = "v2.0.0+incompatible",
)

go_repository(
    name = "com_github_coreos_go_semver",
    importpath = "github.com/coreos/go-semver",
    sum = "h1:wkHLiw0WNATZnSG7epLsujiMCgPAc9xhjJ4tgnAxmfM=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_coreos_go_systemd",
    importpath = "github.com/coreos/go-systemd",
    sum = "h1:Wf6HqHfScWJN9/ZjdUKyjop4mf3Qdd+1TvvltAvM3m8=",
    version = "v0.0.0-20190321100706-95778dfbb74e",
)

go_repository(
    name = "com_github_coreos_go_systemd_v22",
    importpath = "github.com/coreos/go-systemd/v22",
    sum = "h1:D9/bQk5vlXQFZ6Kwuu6zaiXJ9oTPe68++AzAJc1DzSI=",
    version = "v22.3.2",
)

go_repository(
    name = "com_github_coreos_pkg",
    importpath = "github.com/coreos/pkg",
    sum = "h1:CAKfRE2YtTUIjjh1bkBtyYFaUT/WmOqsJjgtihT0vMI=",
    version = "v0.0.0-20160727233714-3ac0863d7acf",
)

go_repository(
    name = "com_github_cpuguy83_go_md2man",
    importpath = "github.com/cpuguy83/go-md2man",
    sum = "h1:BSKMNlYxDvnunlTymqtgONjNnaRV1sTpcovwwjF22jk=",
    version = "v1.0.10",
)

go_repository(
    name = "com_github_cpuguy83_go_md2man_v2",
    importpath = "github.com/cpuguy83/go-md2man/v2",
    sum = "h1:U+s90UTSYgptZMwQh2aRr3LuazLJIa+Pg3Kc1ylSYVY=",
    version = "v2.0.0-20190314233015-f79a8a8ca69d",
)

go_repository(
    name = "com_github_creack_pty",
    importpath = "github.com/creack/pty",
    sum = "h1:6pwm8kMQKCmgUg0ZHTm5+/YvRK0s3THD/28+T6/kk4A=",
    version = "v1.1.7",
)

go_repository(
    name = "com_github_daixiang0_gci",
    importpath = "github.com/daixiang0/gci",
    sum = "h1:1mrIGMBQsBu0P7j7m1M8Lb+ZeZxsZL+jyGX4YoMJJpg=",
    version = "v0.2.8",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    importpath = "github.com/davecgh/go-spew",
    sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_denis_tingajkin_go_header",
    importpath = "github.com/denis-tingajkin/go-header",
    sum = "h1:jEeSF4sdv8/3cT/WY8AgDHUoItNSoEZ7qg9dX7pc218=",
    version = "v0.4.2",
)

go_repository(
    name = "com_github_dgrijalva_jwt_go",
    importpath = "github.com/dgrijalva/jwt-go",
    sum = "h1:7qlOGliEKZXTDg6OTjfoBKDXWrumCAMpl/TFQ4/5kLM=",
    version = "v3.2.0+incompatible",
)

go_repository(
    name = "com_github_dgryski_go_sip13",
    importpath = "github.com/dgryski/go-sip13",
    sum = "h1:RMLoZVzv4GliuWafOuPuQDKSm1SJph7uCRnnS61JAn4=",
    version = "v0.0.0-20181026042036-e10d5fee7954",
)

go_repository(
    name = "com_github_djarvur_go_err113",
    importpath = "github.com/Djarvur/go-err113",
    sum = "h1:sHglBQTwgx+rWPdisA5ynNEsoARbiCBOyGcJM4/OzsM=",
    version = "v0.0.0-20210108212216-aea10b59be24",
)

go_repository(
    name = "com_github_dustin_go_humanize",
    importpath = "github.com/dustin/go-humanize",
    sum = "h1:VSnTsYCnlFHaM2/igO1h6X3HA71jcobQuxemgkq4zYo=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_emicklei_proto",
    importpath = "github.com/emicklei/proto",
    sum = "h1:XbpwxmuOPrdES97FrSfpyy67SSCV/wBIKXqgJzh6hNw=",
    version = "v1.6.15",
)

go_repository(
    name = "com_github_envoyproxy_go_control_plane_a",
    importpath = "github.com/envoyproxy/go-control-plane",
    sum = "h1:4cmBvAEBNJaGARUEs3/suWRyfyBfhf7I60WBZq+bv2w=",
    version = "v0.9.1-0.20191026205805-5f8ba28d4473",
)

go_repository(
    name = "com_github_envoyproxy_protoc_gen_validate",
    importpath = "github.com/envoyproxy/protoc-gen-validate",
    sum = "h1:EQciDnbrYxy13PgWoY8AqoxGiPrpgBZ1R8UNe3ddc+A=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_esimonov_ifshort",
    importpath = "github.com/esimonov/ifshort",
    sum = "h1:K5s1W2fGfkoWXsFlxBNqT6J0ZCncPaKrGM5qe0bni68=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_ettle_strcase",
    importpath = "github.com/ettle/strcase",
    sum = "h1:htFueZyVeE1XNnMEfbqp5r67qAN/4r6ya1ysq8Q+Zcw=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_fatih_structtag",
    importpath = "github.com/fatih/structtag",
    sum = "h1:/OdNE99OxoI/PqaW/SuSK9uxxT3f/tcSZgon/ssNSx4=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_fsnotify_fsnotify",
    importpath = "github.com/fsnotify/fsnotify",
    sum = "h1:IXs+QLmnXW2CcXuY+8Mzv/fWEsPGWxqefPtCP5CnV9I=",
    version = "v1.4.7",
)

go_repository(
    name = "com_github_fullstorydev_grpcurl",
    importpath = "github.com/fullstorydev/grpcurl",
    sum = "h1:p8BB6VZF8O7w6MxGr3KJ9E6EVKaswCevSALK6FBtMzA=",
    version = "v1.6.0",
)

go_repository(
    name = "com_github_fzipp_gocyclo",
    importpath = "github.com/fzipp/gocyclo",
    sum = "h1:A9UeX3HJSXTBzvHzhqoYVuE0eAhe+aM8XBCCwsPMZOc=",
    version = "v0.3.1",
)

go_repository(
    name = "com_github_ghodss_yaml",
    importpath = "github.com/ghodss/yaml",
    sum = "h1:wQHKEahhL6wmXdzwWG11gIVCkOv05bNOh+Rxn0yngAk=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_go_critic_go_critic",
    importpath = "github.com/go-critic/go-critic",
    sum = "h1:siUR1+322iVikWXoV75I1YRfNaC/yaLzhdF9Zwd8Tus=",
    version = "v0.5.6",
)

go_repository(
    name = "com_github_go_gl_glfw",
    importpath = "github.com/go-gl/glfw",
    sum = "h1:QbL/5oDUmRBzO9/Z7Seo6zf912W/a6Sr4Eu0G/3Jho0=",
    version = "v0.0.0-20190409004039-e6da0acd62b1",
)

go_repository(
    name = "com_github_go_gl_glfw_v3_3_glfw",
    importpath = "github.com/go-gl/glfw/v3.3/glfw",
    sum = "h1:WtGNWLvXpe6ZudgnXrq0barxBImvnnJoMEhXAzcbM0I=",
    version = "v0.0.0-20200222043503-6f7a984d4dc4",
)

go_repository(
    name = "com_github_go_kit_kit",
    importpath = "github.com/go-kit/kit",
    sum = "h1:dXFJfIHVvUcpSgDOV+Ne6t7jXri8Tfv2uOLHUZ2XNuo=",
    version = "v0.10.0",
)

go_repository(
    name = "com_github_go_lintpack_lintpack",
    importpath = "github.com/go-lintpack/lintpack",
    sum = "h1:DI5mA3+eKdWeJ40nU4d6Wc26qmdG8RCi/btYq0TuRN0=",
    version = "v0.5.2",
)

go_repository(
    name = "com_github_go_logfmt_logfmt",
    importpath = "github.com/go-logfmt/logfmt",
    sum = "h1:TrB8swr/68K7m9CcGut2g3UOihhbcbiMAYiuTXdEih4=",
    version = "v0.5.0",
)

go_repository(
    name = "com_github_go_ole_go_ole",
    importpath = "github.com/go-ole/go-ole",
    sum = "h1:nNBDSCOigTSiarFpYE9J/KtEA1IOW4CNeqT9TQDqCxI=",
    version = "v1.2.4",
)

go_repository(
    name = "com_github_go_redis_redis",
    importpath = "github.com/go-redis/redis",
    sum = "h1:BKZuG6mCnRj5AOaWJXoCgf6rqTYnYJLe4en2hxT7r9o=",
    version = "v6.15.8+incompatible",
)

go_repository(
    name = "com_github_go_sql_driver_mysql",
    importpath = "github.com/go-sql-driver/mysql",
    sum = "h1:7LxgVwFb2hIQtMm87NdgAVfXjnt4OePseqT1tKx+opk=",
    version = "v1.4.0",
)

go_repository(
    name = "com_github_go_stack_stack",
    importpath = "github.com/go-stack/stack",
    sum = "h1:5SgMzNM5HxrEjV0ww2lTmX6E2Izsfxas4+YHWRs3Lsk=",
    version = "v1.8.0",
)

go_repository(
    name = "com_github_go_toolsmith_astcast",
    importpath = "github.com/go-toolsmith/astcast",
    sum = "h1:JojxlmI6STnFVG9yOImLeGREv8W2ocNUM+iOhR6jE7g=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_go_toolsmith_astcopy",
    importpath = "github.com/go-toolsmith/astcopy",
    sum = "h1:OMgl1b1MEpjFQ1m5ztEO06rz5CUd3oBv9RF7+DyvdG8=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_go_toolsmith_astequal",
    importpath = "github.com/go-toolsmith/astequal",
    sum = "h1:4zxD8j3JRFNyLN46lodQuqz3xdKSrur7U/sr0SDS/gQ=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_go_toolsmith_astfmt",
    importpath = "github.com/go-toolsmith/astfmt",
    sum = "h1:A0vDDXt+vsvLEdbMFJAUBI/uTbRw1ffOPnxsILnFL6k=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_go_toolsmith_astinfo",
    importpath = "github.com/go-toolsmith/astinfo",
    sum = "h1:wP6mXeB2V/d1P1K7bZ5vDUO3YqEzcvOREOxZPEu3gVI=",
    version = "v0.0.0-20180906194353-9809ff7efb21",
)

go_repository(
    name = "com_github_go_toolsmith_astp",
    importpath = "github.com/go-toolsmith/astp",
    sum = "h1:alXE75TXgcmupDsMK1fRAy0YUzLzqPVvBKoyWV+KPXg=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_go_toolsmith_pkgload",
    importpath = "github.com/go-toolsmith/pkgload",
    sum = "h1:4DFWWMXVfbcN5So1sBNW9+yeiMqLFGl1wFLTL5R0Tgg=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_go_toolsmith_strparse",
    importpath = "github.com/go-toolsmith/strparse",
    sum = "h1:Vcw78DnpCAKlM20kSbAyO4mPfJn/lyYA4BJUDxe2Jb4=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_go_toolsmith_typep",
    importpath = "github.com/go-toolsmith/typep",
    sum = "h1:8xdsa1+FSIH/RhEkgnD1j2CJOy5mNllW1Q9tRiYwvlk=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_go_xmlfmt_xmlfmt",
    importpath = "github.com/go-xmlfmt/xmlfmt",
    sum = "h1:khEcpUM4yFcxg4/FHQWkvVRmgijNXRfzkIDHh23ggEo=",
    version = "v0.0.0-20191208150333-d5b6f63a941b",
)

go_repository(
    name = "com_github_gobwas_glob",
    importpath = "github.com/gobwas/glob",
    sum = "h1:A4xDbljILXROh+kObIiy5kIaPYD8e96x1tgBhUI5J+Y=",
    version = "v0.2.3",
)

go_repository(
    name = "com_github_godbus_dbus_v5",
    importpath = "github.com/godbus/dbus/v5",
    sum = "h1:9349emZab16e7zQvpmsbtjc18ykshndd8y2PG3sgJbA=",
    version = "v5.0.4",
)

go_repository(
    name = "com_github_gofrs_flock",
    importpath = "github.com/gofrs/flock",
    sum = "h1:MSdYClljsF3PbENUUEx85nkWfJSGfzYI9yEBZOJz6CY=",
    version = "v0.8.0",
)

go_repository(
    name = "com_github_gogo_protobuf",
    importpath = "github.com/gogo/protobuf",
    sum = "h1:DqDEcV5aeaTmdFBePNpYsp3FlcVH/2ISVVM9Qf8PSls=",
    version = "v1.3.1",
)

go_repository(
    name = "com_github_golang_glog",
    importpath = "github.com/golang/glog",
    sum = "h1:VKtxabqXZkF25pY9ekfRL6a582T4P37/31XEstQ5p58=",
    version = "v0.0.0-20160126235308-23def4e6c14b",
)

go_repository(
    name = "com_github_golang_groupcache",
    importpath = "github.com/golang/groupcache",
    sum = "h1:ZgQEtGgCBiWRM39fZuwSd1LwSqqSW0hOdXCYYDX0R3I=",
    version = "v0.0.0-20190702054246-869f871628b6",
)

go_repository(
    name = "com_github_golang_protobuf",
    importpath = "github.com/golang/protobuf",
    sum = "h1:6nsPYzhq5kReh6QImI3k5qWzO4PEbvbIW2cwSfR/6xs=",
    version = "v1.3.2",
)

go_repository(
    name = "com_github_golangci_check",
    importpath = "github.com/golangci/check",
    sum = "h1:23T5iq8rbUYlhpt5DB4XJkc6BU31uODLD1o1gKvZmD0=",
    version = "v0.0.0-20180506172741-cfe4005ccda2",
)

go_repository(
    name = "com_github_golangci_dupl",
    importpath = "github.com/golangci/dupl",
    sum = "h1:w8hkcTqaFpzKqonE9uMCefW1WDie15eSP/4MssdenaM=",
    version = "v0.0.0-20180902072040-3e9179ac440a",
)

go_repository(
    name = "com_github_golangci_errcheck",
    importpath = "github.com/golangci/errcheck",
    sum = "h1:YYWNAGTKWhKpcLLt7aSj/odlKrSrelQwlovBpDuf19w=",
    version = "v0.0.0-20181223084120-ef45e06d44b6",
)

go_repository(
    name = "com_github_golangci_go_misc",
    importpath = "github.com/golangci/go-misc",
    sum = "h1:9kfjN3AdxcbsZBf8NjltjWihK2QfBBBZuv91cMFfDHw=",
    version = "v0.0.0-20180628070357-927a3d87b613",
)

go_repository(
    name = "com_github_golangci_go_tools",
    importpath = "github.com/golangci/go-tools",
    sum = "h1:/7detzz5stiXWPzkTlPTzkBEIIE4WGpppBJYjKqBiPI=",
    version = "v0.0.0-20190318055746-e32c54105b7c",
)

go_repository(
    name = "com_github_golangci_goconst",
    importpath = "github.com/golangci/goconst",
    sum = "h1:pe9JHs3cHHDQgOFXJJdYkK6fLz2PWyYtP4hthoCMvs8=",
    version = "v0.0.0-20180610141641-041c5f2b40f3",
)

go_repository(
    name = "com_github_golangci_gocyclo",
    importpath = "github.com/golangci/gocyclo",
    sum = "h1:J2XAy40+7yz70uaOiMbNnluTg7gyQhtGqLQncQh+4J8=",
    version = "v0.0.0-20180528134321-2becd97e67ee",
)

go_repository(
    name = "com_github_golangci_gofmt",
    importpath = "github.com/golangci/gofmt",
    sum = "h1:iR3fYXUjHCR97qWS8ch1y9zPNsgXThGwjKPrYfqMPks=",
    version = "v0.0.0-20190930125516-244bba706f1a",
)

go_repository(
    name = "com_github_golangci_golangci_lint",
    importpath = "github.com/golangci/golangci-lint",
    sum = "h1:pBrCqt9BgI9LfGCTKRTSe1DfMjR6BkOPERPaXJYXA6Q=",
    version = "v1.40.1",
)

go_repository(
    name = "com_github_golangci_gosec",
    importpath = "github.com/golangci/gosec",
    sum = "h1:fUdgm/BdKvwOHxg5AhNbkNRp2mSy8sxTXyBVs/laQHo=",
    version = "v0.0.0-20190211064107-66fb7fc33547",
)

go_repository(
    name = "com_github_golangci_ineffassign",
    importpath = "github.com/golangci/ineffassign",
    sum = "h1:gLLhTLMk2/SutryVJ6D4VZCU3CUqr8YloG7FPIBWFpI=",
    version = "v0.0.0-20190609212857-42439a7714cc",
)

go_repository(
    name = "com_github_golangci_lint_1",
    importpath = "github.com/golangci/lint-1",
    sum = "h1:MfyDlzVjl1hoaPzPD4Gpb/QgoRfSBR0jdhwGyAWwMSA=",
    version = "v0.0.0-20191013205115-297bf364a8e0",
)

go_repository(
    name = "com_github_golangci_maligned",
    importpath = "github.com/golangci/maligned",
    sum = "h1:kNY3/svz5T29MYHubXix4aDDuE3RWHkPvopM/EDv/MA=",
    version = "v0.0.0-20180506175553-b1d89398deca",
)

go_repository(
    name = "com_github_golangci_misspell",
    importpath = "github.com/golangci/misspell",
    sum = "h1:pLzmVdl3VxTOncgzHcvLOKirdvcx/TydsClUQXTehjo=",
    version = "v0.3.5",
)

go_repository(
    name = "com_github_golangci_prealloc",
    importpath = "github.com/golangci/prealloc",
    sum = "h1:leSNB7iYzLYSSx3J/s5sVf4Drkc68W2wm4Ixh/mr0us=",
    version = "v0.0.0-20180630174525-215b22d4de21",
)

go_repository(
    name = "com_github_golangci_revgrep",
    importpath = "github.com/golangci/revgrep",
    sum = "h1:c9Mqqrm/Clj5biNaG7rABrmwUq88nHh0uABo2b/WYmc=",
    version = "v0.0.0-20210208091834-cd28932614b5",
)

go_repository(
    name = "com_github_golangci_unconvert",
    importpath = "github.com/golangci/unconvert",
    sum = "h1:zwtduBRr5SSWhqsYNgcuWO2kFlpdOZbP0+yRjmvPGys=",
    version = "v0.0.0-20180507085042-28b1c447d1f4",
)

go_repository(
    name = "com_github_google_btree",
    importpath = "github.com/google/btree",
    sum = "h1:0udJVsspx3VBr5FwtLhQQtuAsVc79tTq0ocGIPAU6qo=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_google_certificate_transparency_go",
    importpath = "github.com/google/certificate-transparency-go",
    sum = "h1:6JHXZhXEvilMcTjR4MGZn5KV0IRkcFl4CJx5iHVhjFE=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_google_go_cmp",
    importpath = "github.com/google/go-cmp",
    sum = "h1:Khx7svrCpmxxtHBq5j2mp/xVjsi8hQMfNLvJFAlrGgU=",
    version = "v0.5.5",
)

go_repository(
    name = "com_github_google_gofuzz",
    importpath = "github.com/google/gofuzz",
    sum = "h1:A8PeW59pxE9IoFRqBp37U+mSNaQoZ46F1f0f863XSXw=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_google_martian",
    importpath = "github.com/google/martian",
    sum = "h1:/CP5g8u/VJHijgedC/Legn3BAbAaWPgecwXBIDzw5no=",
    version = "v2.1.0+incompatible",
)

go_repository(
    name = "com_github_google_martian_v3",
    importpath = "github.com/google/martian/v3",
    sum = "h1:wCKgOCHuUEVfsaQLpPSJb7VdYCdTVZQAuOdYm1yc/60=",
    version = "v3.1.0",
)

go_repository(
    name = "com_github_google_pprof",
    importpath = "github.com/google/pprof",
    sum = "h1:zIaiqGYDQwa4HVx5wGRTXbx38Pqxjemn4BP98wpzpXo=",
    version = "v0.0.0-20210226084205-cbba55b83ad5",
)

go_repository(
    name = "com_github_google_renameio",
    importpath = "github.com/google/renameio",
    sum = "h1:GOZbcHa3HfsPKPlmyPyN2KEohoMXOhdMbHrvbpl2QaA=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_google_trillian",
    importpath = "github.com/google/trillian",
    sum = "h1:pPzJPkK06mvXId1LHEAJxIegGgHzzp/FUnycPYfoCMI=",
    version = "v1.3.11",
)

go_repository(
    name = "com_github_google_uuid",
    importpath = "github.com/google/uuid",
    sum = "h1:Gkbcsh/GbpXz7lPftLA3P6TYMwjCLYm83jiFQZF/3gY=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_googleapis_gax_go_v2",
    importpath = "github.com/googleapis/gax-go/v2",
    sum = "h1:sjZBwGj9Jlw33ImPtvFviGYvseOtDM7hkSKB7+Tv3SM=",
    version = "v2.0.5",
)

go_repository(
    name = "com_github_gookit_color",
    importpath = "github.com/gookit/color",
    sum = "h1:w2WcSwaCa1ojRWO60Mm4GJUJomBNKR9G+x9DwaaCL1c=",
    version = "v1.3.8",
)

go_repository(
    name = "com_github_gopherjs_gopherjs",
    importpath = "github.com/gopherjs/gopherjs",
    sum = "h1:EGx4pi6eqNxGaHF6qqu48+N2wcFQ5qg5FXgOdqsJ5d8=",
    version = "v0.0.0-20181017120253-0766667cb4d1",
)

go_repository(
    name = "com_github_gordonklaus_ineffassign",
    importpath = "github.com/gordonklaus/ineffassign",
    sum = "h1:Nb2aRlC404yz7gQIfRZxX9/MLvQiqXyiBTJtgAy6yrI=",
    version = "v0.0.0-20210225214923-2e10b2664254",
)

go_repository(
    name = "com_github_gorhill_cronexpr",
    importpath = "github.com/gorhill/cronexpr",
    sum = "h1:f0n1xnMSmBLzVfsMMvriDyA75NB/oBgILX2GcHXIQzY=",
    version = "v0.0.0-20180427100037-88b0669f7d75",
)

go_repository(
    name = "com_github_gorilla_mux",
    importpath = "github.com/gorilla/mux",
    sum = "h1:gnP5JzjVOuiZD07fKKToCAOjS0yOpj/qPETTXCCS6hw=",
    version = "v1.7.3",
)

go_repository(
    name = "com_github_gorilla_websocket",
    importpath = "github.com/gorilla/websocket",
    sum = "h1:WDFjx/TMzVgy9VdMMQi2K2Emtwi2QcUQsztZ/zLaH/Q=",
    version = "v1.4.0",
)

go_repository(
    name = "com_github_gostaticanalysis_analysisutil",
    importpath = "github.com/gostaticanalysis/analysisutil",
    sum = "h1:/7clKqrVfiVwiBQLM0Uke4KvXnO6JcCTS7HwF2D6wG8=",
    version = "v0.4.1",
)

go_repository(
    name = "com_github_gostaticanalysis_comment",
    importpath = "github.com/gostaticanalysis/comment",
    sum = "h1:xHopR5L2lRz6OsjH4R2HG5wRhW9ySl3FsHIvi5pcXwc=",
    version = "v1.4.1",
)

go_repository(
    name = "com_github_gostaticanalysis_forcetypeassert",
    importpath = "github.com/gostaticanalysis/forcetypeassert",
    sum = "h1:rx8127mFPqXXsfPSo8BwnIU97MKFZc89WHAHt8PwDVY=",
    version = "v0.0.0-20200621232751-01d4955beaa5",
)

go_repository(
    name = "com_github_gostaticanalysis_nilerr",
    importpath = "github.com/gostaticanalysis/nilerr",
    sum = "h1:ThE+hJP0fEp4zWLkWHWcRyI2Od0p7DlgYG3Uqrmrcpk=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_gregjones_httpcache",
    importpath = "github.com/gregjones/httpcache",
    sum = "h1:pdN6V1QBWetyv/0+wjACpqVH+eVULgEjkurDLq3goeM=",
    version = "v0.0.0-20180305231024-9cad4c3443a7",
)

go_repository(
    name = "com_github_grpc_ecosystem_go_grpc_middleware",
    importpath = "github.com/grpc-ecosystem/go-grpc-middleware",
    sum = "h1:z53tR0945TRRQO/fLEVPI6SMv7ZflF0TEaTAoU7tOzg=",
    version = "v1.0.1-0.20190118093823-f849b5445de4",
)

go_repository(
    name = "com_github_grpc_ecosystem_go_grpc_prometheus",
    importpath = "github.com/grpc-ecosystem/go-grpc-prometheus",
    sum = "h1:Ovs26xHkKqVztRpIrF/92BcuyuQ/YW4NSIpoGtfXNho=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_grpc_ecosystem_grpc_gateway",
    importpath = "github.com/grpc-ecosystem/grpc-gateway",
    sum = "h1:UImYN5qQ8tuGpGE16ZmjvcTtTw24zw1QAp/SlnNrZhI=",
    version = "v1.9.5",
)

go_repository(
    name = "com_github_hashicorp_consul_api",
    importpath = "github.com/hashicorp/consul/api",
    sum = "h1:HXNYlRkkM/t+Y/Yhxtwcy02dlYwIaoxzvxPnS+cqy78=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_hashicorp_consul_sdk",
    importpath = "github.com/hashicorp/consul/sdk",
    sum = "h1:UOxjlb4xVNF93jak1mzzoBatyFju9nrkxpVwIp/QqxQ=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_hashicorp_errwrap",
    importpath = "github.com/hashicorp/errwrap",
    sum = "h1:hLrqtEDnRye3+sgx6z4qVLNuviH3MR5aQ0ykNJa/UYA=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_go_cleanhttp",
    importpath = "github.com/hashicorp/go-cleanhttp",
    sum = "h1:dH3aiDG9Jvb5r5+bYHsikaOUIpcM0xvgMXVoDkXMzJM=",
    version = "v0.5.1",
)

go_repository(
    name = "com_github_hashicorp_go_immutable_radix_a",
    importpath = "github.com/hashicorp/go-immutable-radix",
    sum = "h1:AKDB1HM5PWEA7i4nhcpwOrO2byshxBjXVn/J/3+z5/0=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_go_immutable_radix_b",
    importpath = "github.com/hashicorp/go-immutable-radix",
    sum = "h1:DKHmCUm2hRBK510BaiZlwvpD40f8bJFeZnpfm2KLowc=",
    version = "v1.3.1",
)

go_repository(
    name = "com_github_hashicorp_go_msgpack",
    importpath = "github.com/hashicorp/go-msgpack",
    sum = "h1:zKjpN5BK/P5lMYrLmBHdBULWbJ0XpYR+7NGzqkZzoD4=",
    version = "v0.5.3",
)

go_repository(
    name = "com_github_hashicorp_go_multierror",
    importpath = "github.com/hashicorp/go-multierror",
    sum = "h1:iVjPR7a6H0tWELX5NxNe7bYopibicUzc7uPribsnS6o=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_go_net",
    importpath = "github.com/hashicorp/go.net",
    sum = "h1:sNCoNyDEvN1xa+X0baata4RdcpKwcMS6DH+xwfqPgjw=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_hashicorp_go_rootcerts",
    importpath = "github.com/hashicorp/go-rootcerts",
    sum = "h1:Rqb66Oo1X/eSV1x66xbDccZjhJigjg0+e82kpwzSwCI=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_go_sockaddr",
    importpath = "github.com/hashicorp/go-sockaddr",
    sum = "h1:GeH6tui99pF4NJgfnhp+L6+FfobzVW3Ah46sLo0ICXs=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_go_syslog",
    importpath = "github.com/hashicorp/go-syslog",
    sum = "h1:KaodqZuhUoZereWVIYmpUgZysurB1kBLX2j0MwMrUAE=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_go_uuid",
    importpath = "github.com/hashicorp/go-uuid",
    sum = "h1:fv1ep09latC32wFoVwnqcnKJGnMSdBanPczbHAYm1BE=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_hashicorp_golang_lru_a",
    importpath = "github.com/hashicorp/golang-lru",
    sum = "h1:0hERBMJE1eitiLkihrMvRVBYAkpHzc/J3QdDN+dAcgU=",
    version = "v0.5.1",
)

go_repository(
    name = "com_github_hashicorp_golang_lru_b",
    importpath = "github.com/hashicorp/golang-lru",
    sum = "h1:YDjusn29QI/Das2iO9M0BHnIbxPeyuCHsjMW+lJfyTc=",
    version = "v0.5.4",
)

go_repository(
    name = "com_github_hashicorp_hcl",
    importpath = "github.com/hashicorp/hcl",
    sum = "h1:0Anlzjpi4vEasTeNFn2mLJgTSwt0+6sfsiTG8qcWGx4=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_logutils",
    importpath = "github.com/hashicorp/logutils",
    sum = "h1:dLEQVugN8vlakKOUE3ihGLTZJRB4j+M2cdTm/ORI65Y=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_mdns",
    importpath = "github.com/hashicorp/mdns",
    sum = "h1:WhIgCr5a7AaVH6jPUwjtRuuE7/RDufnUvzIr48smyxs=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_memberlist",
    importpath = "github.com/hashicorp/memberlist",
    sum = "h1:EmmoJme1matNzb+hMpDuR/0sbJSUisxyqBGG676r31M=",
    version = "v0.1.3",
)

go_repository(
    name = "com_github_hashicorp_serf",
    importpath = "github.com/hashicorp/serf",
    sum = "h1:YZ7UKsJv+hKjqGVUUbtE3HNj79Eln2oQ75tniF6iPt0=",
    version = "v0.8.2",
)

go_repository(
    name = "com_github_hpcloud_tail",
    importpath = "github.com/hpcloud/tail",
    sum = "h1:nfCOvKYfkgYP8hkirhJocXT2+zOD8yUNjXaWfTlyFKI=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_huandu_xstrings",
    importpath = "github.com/huandu/xstrings",
    sum = "h1:yPeWdRnmynF7p+lLYz0H2tthW9lqhMJrQV/U7yy4wX0=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_ianlancetaylor_demangle",
    importpath = "github.com/ianlancetaylor/demangle",
    sum = "h1:mV02weKRL81bEnm8A0HT1/CAelMQDBuQIfLw8n+d6xI=",
    version = "v0.0.0-20200824232613-28f6c0f3b639",
)

go_repository(
    name = "com_github_imdario_mergo",
    importpath = "github.com/imdario/mergo",
    sum = "h1:Y+UAYTZ7gDEuOfhxKWy+dvb5dRQ6rJjFSdX2HZY1/gI=",
    version = "v0.3.7",
)

go_repository(
    name = "com_github_inconshreveable_mousetrap",
    importpath = "github.com/inconshreveable/mousetrap",
    sum = "h1:Z8tu5sraLXCXIcARxBp/8cbvlwVa7Z1NHg9XEKhtSvM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_jgautheron_goconst",
    importpath = "github.com/jgautheron/goconst",
    sum = "h1:hp9XKUpe/MPyDamUbfsrGpe+3dnY2whNK4EtB86dvLM=",
    version = "v1.4.0",
)

go_repository(
    name = "com_github_jhump_protoreflect",
    importpath = "github.com/jhump/protoreflect",
    sum = "h1:4/2yi5LyDPP7nN+Hiird1SAJ6YoxUm13/oxHGRnbPd8=",
    version = "v1.6.1",
)

go_repository(
    name = "com_github_jingyugao_rowserrcheck",
    importpath = "github.com/jingyugao/rowserrcheck",
    sum = "h1:4Rlb26NqzNtbDH69CRpr0vZooj3jAlXTycWCX3xRYAY=",
    version = "v0.0.0-20210315055705-d907ca737bb1",
)

go_repository(
    name = "com_github_jirfag_go_printf_func_name",
    importpath = "github.com/jirfag/go-printf-func-name",
    sum = "h1:KA9BjwUk7KlCh6S9EAGWBt1oExIUv9WyNCiRz5amv48=",
    version = "v0.0.0-20200119135958-7558a9eaa5af",
)

go_repository(
    name = "com_github_jmespath_go_jmespath",
    importpath = "github.com/jmespath/go-jmespath",
    sum = "h1:pmfjZENx5imkbgOkpRUYLnmbU7UEFbjtDA2hxJ1ichM=",
    version = "v0.0.0-20180206201540-c2b33e8439af",
)

go_repository(
    name = "com_github_jmespath_go_jmespath_internal_testify",
    importpath = "github.com/jmespath/go-jmespath/internal/testify",
    sum = "h1:shLQSRRSCCPj3f2gpwzGwWFoC7ycTf1rcQZHOlsJ6N8=",
    version = "v1.5.1",
)

go_repository(
    name = "com_github_jmoiron_sqlx",
    importpath = "github.com/jmoiron/sqlx",
    sum = "h1:41Ip0zITnmWNR/vHV+S4m+VoUivnWY5E4OJfLZjCJMA=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_jonboulle_clockwork",
    importpath = "github.com/jonboulle/clockwork",
    sum = "h1:VKV+ZcuP6l3yW9doeqz6ziZGgcynBVQO+obU0+0hcPo=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_json_iterator_go",
    importpath = "github.com/json-iterator/go",
    sum = "h1:QiWkFLKq0T7mpzwOTu6BzNDbfTE8OLrYhVKYMLF46Ok=",
    version = "v1.1.8",
)

go_repository(
    name = "com_github_jstemmer_go_junit_report",
    importpath = "github.com/jstemmer/go-junit-report",
    sum = "h1:6QPYqodiu3GuPL+7mfx+NwDdp2eTkp9IfEUpgAwUN0o=",
    version = "v0.9.1",
)

go_repository(
    name = "com_github_jtolds_gls",
    importpath = "github.com/jtolds/gls",
    sum = "h1:xdiiI2gbIgH/gLH7ADydsJ1uDOEzR8yvV7C0MuV77Wo=",
    version = "v4.20.0+incompatible",
)

go_repository(
    name = "com_github_juju_ratelimit",
    importpath = "github.com/juju/ratelimit",
    sum = "h1:+7AIFJVQ0EQgq/K9+0Krm7m530Du7tIz0METWzN0RgY=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_julienschmidt_httprouter",
    importpath = "github.com/julienschmidt/httprouter",
    sum = "h1:TDTW5Yz1mjftljbcKqRcrYhd4XeOoI98t+9HbQbYf7g=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_julz_importas",
    importpath = "github.com/julz/importas",
    sum = "h1:XeSMXURZPtUffuWAaq90o6kLgZdgu+QA8wk4MPC8ikI=",
    version = "v0.0.0-20210419104244-841f0c0fe66d",
)

go_repository(
    name = "com_github_k0kubun_colorstring",
    importpath = "github.com/k0kubun/colorstring",
    sum = "h1:uC1QfSlInpQF+M0ao65imhwqKnz3Q2z/d8PWZRMQvDM=",
    version = "v0.0.0-20150214042306-9440f1994b88",
)

go_repository(
    name = "com_github_kisielk_errcheck",
    importpath = "github.com/kisielk/errcheck",
    sum = "h1:reN85Pxc5larApoH1keMBiu2GWtPqXQ1nc9gx+jOU+E=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_kisielk_gotool",
    importpath = "github.com/kisielk/gotool",
    sum = "h1:AV2c/EiW3KqPNT9ZKl07ehoAGi4C5/01Cfbblndcapg=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_klauspost_compress",
    importpath = "github.com/klauspost/compress",
    sum = "h1:BtAvtV1+h0YwSVwWoYXMREPpYu9VzTJ9QDI1TEg/iQQ=",
    version = "v1.13.3",
)

go_repository(
    name = "com_github_klauspost_cpuid",
    importpath = "github.com/klauspost/cpuid",
    sum = "h1:NMpwD2G9JSFOE1/TJjGSo5zG7Yb2bTe7eq1jH+irmeE=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_konsorten_go_windows_terminal_sequences",
    importpath = "github.com/konsorten/go-windows-terminal-sequences",
    sum = "h1:mweAR1A6xJ3oS2pRaGiHgQ4OO8tzTaLawm8vnODuwDk=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_kr_fs",
    importpath = "github.com/kr/fs",
    sum = "h1:Jskdu9ieNAYnjxsi0LbQp1ulIKZV1LAFgK1tWhpZgl8=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_kr_logfmt",
    importpath = "github.com/kr/logfmt",
    sum = "h1:T+h1c/A9Gawja4Y9mFVWj2vyii2bbUNDw3kt9VxK2EY=",
    version = "v0.0.0-20140226030751-b84e30acd515",
)

go_repository(
    name = "com_github_kr_pretty",
    importpath = "github.com/kr/pretty",
    sum = "h1:Fmg33tUaq4/8ym9TJN1x7sLJnHVwhP33CNkpYV/7rwI=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_kr_pty",
    importpath = "github.com/kr/pty",
    sum = "h1:VkoXIwSboBpnk99O/KFauAEILuNHv5DVFKZMBN/gUgw=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_kr_text",
    importpath = "github.com/kr/text",
    sum = "h1:45sCR5RtlFHMR4UwH9sdQ5TC8v0qDQCHnXt+kaKSTVE=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_kulti_thelper",
    importpath = "github.com/kulti/thelper",
    sum = "h1:2Nx7XbdbE/BYZeoip2mURKUdtHQRuy6Ug+wR7K9ywNM=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_kunwardeep_paralleltest",
    importpath = "github.com/kunwardeep/paralleltest",
    sum = "h1:/jJRv0TiqPoEy/Y8dQxCFJhD56uS/pnvtatgTZBHokU=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_kylelemons_godebug",
    importpath = "github.com/kylelemons/godebug",
    sum = "h1:RPNrshWIDI6G2gRW9EHilWtl7Z6Sb1BR0xunSBf0SNc=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_kyoh86_exportloopref",
    importpath = "github.com/kyoh86/exportloopref",
    sum = "h1:5Ry/at+eFdkX9Vsdw3qU4YkvGtzuVfzT4X7S77LoN/M=",
    version = "v0.1.8",
)

go_repository(
    name = "com_github_kyoh86_richgo",
    importpath = "github.com/kyoh86/richgo",
    sum = "h1:om0AqQ6LmNYvaL7pMK7s3gixTM6iWzH3ICkfS0qNs+k=",
    version = "v0.3.9",
)

go_repository(
    name = "com_github_kyoh86_xdg",
    importpath = "github.com/kyoh86/xdg",
    sum = "h1:CERuT/ShdTDj+A2UaX3hQ3mOV369+Sj+wyn2nIRIIkI=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_ldez_gomoddirectives",
    importpath = "github.com/ldez/gomoddirectives",
    sum = "h1:9pAcW9KRZW7HQjFwbozNvFMcNVwdCBufU7os5QUwLIY=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_ldez_tagliatelle",
    importpath = "github.com/ldez/tagliatelle",
    sum = "h1:693V8Bf1NdShJ8eu/s84QySA0J2VWBanVBa2WwXD/Wk=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_letsencrypt_pkcs11key_v4",
    importpath = "github.com/letsencrypt/pkcs11key/v4",
    sum = "h1:qLc/OznH7xMr5ARJgkZCCWk+EomQkiNTOoOF5LAgagc=",
    version = "v4.0.0",
)

go_repository(
    name = "com_github_lib_pq",
    importpath = "github.com/lib/pq",
    sum = "h1:X5PMW56eZitiTeO7tKzZxFCSpbFZJtkMMooicw2us9A=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_logrusorgru_aurora",
    importpath = "github.com/logrusorgru/aurora",
    sum = "h1:9MlwzLdW7QSDrhDjFlsEYmxpFyIoXmYRon3dt0io31k=",
    version = "v0.0.0-20181002194514-a7b3b318ed4e",
)

go_repository(
    name = "com_github_magiconair_properties",
    importpath = "github.com/magiconair/properties",
    sum = "h1:b6kJs+EmPFMYGkow9GiUyCyOvIwYetYJ3fSaWak/Gls=",
    version = "v1.8.5",
)

go_repository(
    name = "com_github_maratori_testpackage",
    importpath = "github.com/maratori/testpackage",
    sum = "h1:QtJ5ZjqapShm0w5DosRjg0PRlSdAdlx+W6cCKoALdbQ=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_masterminds_goutils",
    importpath = "github.com/Masterminds/goutils",
    sum = "h1:zukEsf/1JZwCMgHiK3GZftabmxiCw4apj3a28RPBiVg=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_masterminds_semver",
    importpath = "github.com/Masterminds/semver",
    sum = "h1:H65muMkzWKEuNDnfl9d70GUjFniHKHRbFPGBuZ3QEww=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_masterminds_sprig",
    importpath = "github.com/Masterminds/sprig",
    sum = "h1:z4yfnGrZ7netVz+0EDJ0Wi+5VZCSYp4Z0m2dk6cEM60=",
    version = "v2.22.0+incompatible",
)

go_repository(
    name = "com_github_matoous_godox",
    importpath = "github.com/matoous/godox",
    sum = "h1:pWxk9e//NbPwfxat7RXkts09K+dEBJWakUWwICVqYbA=",
    version = "v0.0.0-20210227103229-6504466cf951",
)

go_repository(
    name = "com_github_mattn_go_colorable_a",
    importpath = "github.com/mattn/go-colorable",
    sum = "h1:UVL0vNpWh04HeJXV0KLcaT7r06gOH2l4OW6ddYRUIY4=",
    version = "v0.0.9",
)

go_repository(
    name = "com_github_mattn_go_colorable_b",
    importpath = "github.com/mattn/go-colorable",
    sum = "h1:c1ghPdyEDarC70ftn0y+A/Ee++9zz8ljHG1b13eJ0s8=",
    version = "v0.1.8",
)

go_repository(
    name = "com_github_mattn_go_isatty_a",
    importpath = "github.com/mattn/go-isatty",
    sum = "h1:ns/ykhmWi7G9O+8a448SecJU3nSMBXJfqQkl0upE1jI=",
    version = "v0.0.3",
)

go_repository(
    name = "com_github_mattn_go_isatty_b",
    importpath = "github.com/mattn/go-isatty",
    sum = "h1:wuysRhFDzyxgEmMf5xjvJ2M9dZoWAXNNr5LSBS7uHXY=",
    version = "v0.0.12",
)

go_repository(
    name = "com_github_mattn_go_runewidth",
    importpath = "github.com/mattn/go-runewidth",
    sum = "h1:UnlwIPBGaTZfPQ6T1IGzPI0EkYAQmT9fAEJ/poFC63o=",
    version = "v0.0.2",
)

go_repository(
    name = "com_github_mattn_go_sqlite3",
    importpath = "github.com/mattn/go-sqlite3",
    sum = "h1:jbhqpg7tQe4SupckyijYiy0mJJ/pRyHvXf7JdWK860o=",
    version = "v1.10.0",
)

go_repository(
    name = "com_github_mattn_goveralls",
    importpath = "github.com/mattn/goveralls",
    sum = "h1:7eJB6EqsPhRVxvwEXGnqdO2sJI0PTsrWoTMXEk9/OQc=",
    version = "v0.0.2",
)

go_repository(
    name = "com_github_matttproud_golang_protobuf_extensions",
    importpath = "github.com/matttproud/golang_protobuf_extensions",
    sum = "h1:4hp9jkHxhMHkqkrB3Ix0jegS5sx/RkqARlsWZ6pIwiU=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_mbilski_exhaustivestruct",
    importpath = "github.com/mbilski/exhaustivestruct",
    sum = "h1:wCBmUnSYufAHO6J4AVWY6ff+oxWxsVFrwgOdMUQePUo=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_mgechev_dots",
    importpath = "github.com/mgechev/dots",
    sum = "h1:QASJXOGm2RZ5Ardbc86qNFvby9AqkLDibfChMtAg5QM=",
    version = "v0.0.0-20190921121421-c36f7dcfbb81",
)

go_repository(
    name = "com_github_mgechev_revive",
    importpath = "github.com/mgechev/revive",
    sum = "h1:MgRQ3ys2uQCyVjelaDhVs8oSvOPYInzGA/nNGMa+MNU=",
    version = "v1.0.6",
)

go_repository(
    name = "com_github_miekg_dns",
    importpath = "github.com/miekg/dns",
    sum = "h1:WMhc1ik4LNkTg8U9l3hI1LvxKmIL+f1+WV/SZtCbDDA=",
    version = "v1.1.12",
)

go_repository(
    name = "com_github_miekg_pkcs11",
    importpath = "github.com/miekg/pkcs11",
    sum = "h1:iMwmD7I5225wv84WxIG/bmxz9AXjWvTWIbM/TYHvWtw=",
    version = "v1.0.3",
)

go_repository(
    name = "com_github_mitchellh_cli",
    importpath = "github.com/mitchellh/cli",
    sum = "h1:iGBIsUe3+HZ/AD/Vd7DErOt5sU9fa8Uj7A2s1aggv1Y=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_mitchellh_copystructure",
    importpath = "github.com/mitchellh/copystructure",
    sum = "h1:Laisrj+bAB6b/yJwB5Bt3ITZhGJdqmxquMKeZ+mmkFQ=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_mitchellh_go_homedir",
    importpath = "github.com/mitchellh/go-homedir",
    sum = "h1:vKb8ShqSby24Yrqr/yDYkuFz8d0WUjys40rvnGC8aR0=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_mitchellh_go_ps",
    importpath = "github.com/mitchellh/go-ps",
    sum = "h1:i6ampVEEF4wQFF+bkYfwYgY+F/uYJDktmvLPf7qIgjc=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_mitchellh_go_testing_interface",
    importpath = "github.com/mitchellh/go-testing-interface",
    sum = "h1:fzU/JVNcaqHQEcVFAKeR41fkiLdIPrefOvVG1VZ96U0=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_mitchellh_gox",
    importpath = "github.com/mitchellh/gox",
    sum = "h1:lfGJxY7ToLJQjHHwi0EX6uYBdK78egf954SQl13PQJc=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_mitchellh_iochan",
    importpath = "github.com/mitchellh/iochan",
    sum = "h1:C+X3KsSTLFVBr/tK1eYN/vs4rJcvsiLU338UhYPJWeY=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_mitchellh_mapstructure",
    importpath = "github.com/mitchellh/mapstructure",
    sum = "h1:fmNYVwqnSfB9mZU6OS2O6GsXM+wcskZDuKQzvN1EDeE=",
    version = "v1.1.2",
)

go_repository(
    name = "com_github_mitchellh_reflectwalk",
    importpath = "github.com/mitchellh/reflectwalk",
    sum = "h1:FVzMWA5RllMAKIdUSC8mdWo3XtwoecrH79BY70sEEpE=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_modern_go_concurrent",
    importpath = "github.com/modern-go/concurrent",
    sum = "h1:TRLaZ9cD/w8PVh93nsPXa1VrQ6jlwL5oN8l14QlcNfg=",
    version = "v0.0.0-20180306012644-bacd9c7ef1dd",
)

go_repository(
    name = "com_github_modern_go_reflect2",
    importpath = "github.com/modern-go/reflect2",
    sum = "h1:9f412s+6RmYXLWZSEzVVgPGK7C2PphHj5RJrvfx9AWI=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_mohae_deepcopy",
    importpath = "github.com/mohae/deepcopy",
    sum = "h1:RWengNIwukTxcDr9M+97sNutRR1RKhG96O6jWumTTnw=",
    version = "v0.0.0-20170929034955-c48cc78d4826",
)

go_repository(
    name = "com_github_moricho_tparallel",
    importpath = "github.com/moricho/tparallel",
    sum = "h1:95FytivzT6rYzdJLdtfn6m1bfFJylOJK41+lgv/EHf4=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_morikuni_aec",
    importpath = "github.com/morikuni/aec",
    sum = "h1:nP9CBfwrvYnBRgY6qfDQkygYDmYwOilePFkwzv4dU8A=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_mozilla_scribe",
    importpath = "github.com/mozilla/scribe",
    sum = "h1:29NKShH4TWd3lxCDUhS4Xe16EWMA753dtIxYtwddklU=",
    version = "v0.0.0-20180711195314-fb71baf557c1",
)

go_repository(
    name = "com_github_mozilla_tls_observatory",
    importpath = "github.com/mozilla/tls-observatory",
    sum = "h1:DXaIt8v4XXkFoVZXkG/PjLS5Rz3I2yoflOQrnuGgJeA=",
    version = "v0.0.0-20210209181001-cf43108d6880",
)

go_repository(
    name = "com_github_mpvl_unique",
    importpath = "github.com/mpvl/unique",
    sum = "h1:D5x39vF5KCwKQaw+OC9ZPiLVHXz3UFw2+psEX+gYcto=",
    version = "v0.0.0-20150818121801-cbe035fff7de",
)

go_repository(
    name = "com_github_mwitkow_go_conntrack",
    importpath = "github.com/mwitkow/go-conntrack",
    sum = "h1:F9x/1yl3T2AeKLr2AMdilSD8+f9bvMnNN8VS5iDtovc=",
    version = "v0.0.0-20161129095857-cc309e4a2223",
)

go_repository(
    name = "com_github_mwitkow_go_proto_validators",
    importpath = "github.com/mwitkow/go-proto-validators",
    sum = "h1:F6LFfmgVnfULfaRsQWBbe7F7ocuHCr9+7m+GAeDzNbQ=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_nakabonne_nestif",
    importpath = "github.com/nakabonne/nestif",
    sum = "h1:+yOViDGhg8ygGrmII72nV9B/zGxY188TYpfolntsaPw=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_nbutton23_zxcvbn_go",
    importpath = "github.com/nbutton23/zxcvbn-go",
    sum = "h1:4kuARK6Y6FxaNu/BnU2OAaLF86eTVhP2hjTB6iMvItA=",
    version = "v0.0.0-20210217022336-fa2cb2858354",
)

go_repository(
    name = "com_github_niemeyer_pretty",
    importpath = "github.com/niemeyer/pretty",
    sum = "h1:fD57ERR4JtEqsWbfPhv4DMiApHyliiK5xCTNVSPiaAs=",
    version = "v0.0.0-20200227124842-a10e7caefd8e",
)

go_repository(
    name = "com_github_nishanths_exhaustive",
    importpath = "github.com/nishanths/exhaustive",
    sum = "h1:kVlMw8h2LHPMGUVqUj6230oQjjTMFjwcZrnkhXzFfl8=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_nishanths_predeclared",
    importpath = "github.com/nishanths/predeclared",
    sum = "h1:1TXtjmy4f3YCFjTxRd8zcFHOmoUir+gp0ESzjFzG2sw=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_nxadm_tail",
    importpath = "github.com/nxadm/tail",
    sum = "h1:DQuhQpB1tVlglWS2hLQ5OV6B5r8aGxSrPc5Qo6uTN78=",
    version = "v1.4.4",
)

go_repository(
    name = "com_github_oklog_ulid",
    importpath = "github.com/oklog/ulid",
    sum = "h1:EGfNDEx6MqHz8B3uNV6QAib1UR2Lm97sHi3ocA6ESJ4=",
    version = "v1.3.1",
)

go_repository(
    name = "com_github_olekukonko_tablewriter",
    importpath = "github.com/olekukonko/tablewriter",
    sum = "h1:58+kh9C6jJVXYjt8IE48G2eWl6BjwU5Gj0gqY84fy78=",
    version = "v0.0.0-20170122224234-a0225b3f23b5",
)

go_repository(
    name = "com_github_oneofone_xxhash",
    importpath = "github.com/OneOfOne/xxhash",
    sum = "h1:KMrpdQIwFcEqXDklaen+P1axHaj9BSKzvpUUfnHldSE=",
    version = "v1.2.2",
)

go_repository(
    name = "com_github_onsi_ginkgo",
    importpath = "github.com/onsi/ginkgo",
    sum = "h1:VkHVNpR4iVnU8XQR6DBm8BqYjN7CRzw+xKUbVVbbW9w=",
    version = "v1.8.0",
)

go_repository(
    name = "com_github_onsi_gomega",
    importpath = "github.com/onsi/gomega",
    sum = "h1:izbySO9zDPmjJ8rDjLvkA2zJHIo+HkYXHnf7eN7SSyo=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_openpeedeep_depguard",
    importpath = "github.com/OpenPeeDeeP/depguard",
    sum = "h1:VlW4R6jmBIv3/u1JNlawEvJMM4J+dPORPaZasQee8Us=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_opentracing_opentracing_go",
    importpath = "github.com/opentracing/opentracing-go",
    sum = "h1:pWlfV3Bxv7k65HYwkikxat0+s3pV4bsqf19k25Ur8rU=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_pascaldekloe_goe",
    importpath = "github.com/pascaldekloe/goe",
    sum = "h1:Lgl0gzECD8GnQ5QCWA8o6BtfL6mDH5rQgM4/fX3avOs=",
    version = "v0.0.0-20180627143212-57f6aae5913c",
)

go_repository(
    name = "com_github_pborman_uuid",
    importpath = "github.com/pborman/uuid",
    sum = "h1:J7Q5mO4ysT1dv8hyrUGHb9+ooztCXu1D8MY8DZYsu3g=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_pelletier_go_toml",
    importpath = "github.com/pelletier/go-toml",
    sum = "h1:zeC5b1GviRUyKYd6OJPvBU/mcVDVoL1OhT17FCt5dSQ=",
    version = "v1.9.3",
)

go_repository(
    name = "com_github_peterbourgon_diskv",
    importpath = "github.com/peterbourgon/diskv",
    sum = "h1:UBdAOUP5p4RWqPBg048CAvpKN+vxiaj6gdUUzhl4XmI=",
    version = "v2.0.1+incompatible",
)

go_repository(
    name = "com_github_phayes_checkstyle",
    importpath = "github.com/phayes/checkstyle",
    sum = "h1:CdDQnGF8Nq9ocOS/xlSptM1N3BbrA6/kmaep5ggwaIA=",
    version = "v0.0.0-20170904204023-bfd46e6a821d",
)

go_repository(
    name = "com_github_pkg_diff",
    importpath = "github.com/pkg/diff",
    sum = "h1:aoZm08cpOy4WuID//EZDgcC4zIxODThtZNPirFr42+A=",
    version = "v0.0.0-20210226163009-20ebb0f2a09e",
)

go_repository(
    name = "com_github_pkg_errors",
    importpath = "github.com/pkg/errors",
    sum = "h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=",
    version = "v0.9.1",
)

go_repository(
    name = "com_github_pkg_sftp",
    importpath = "github.com/pkg/sftp",
    sum = "h1:VasscCm72135zRysgrJDKsntdmPN+OuU3+nnHYA9wyc=",
    version = "v1.10.1",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    importpath = "github.com/pmezard/go-difflib",
    sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_polyfloyd_go_errorlint",
    importpath = "github.com/polyfloyd/go-errorlint",
    sum = "h1:uuOfAQo7em74dKh41UzjlQ6dXmE9wYxjvUcfg2EHTDw=",
    version = "v0.0.0-20210418123303-74da32850375",
)

go_repository(
    name = "com_github_posener_complete",
    importpath = "github.com/posener/complete",
    sum = "h1:ccV59UEOTzVDnDUEFdT95ZzHVZ+5+158q8+SJb2QV5w=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_prometheus_client_golang",
    importpath = "github.com/prometheus/client_golang",
    sum = "h1:miYCvYqFXtl/J9FIy8eNpBfYthAEFg+Ys0XyUVEcDsc=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_prometheus_client_model",
    importpath = "github.com/prometheus/client_model",
    sum = "h1:ElTg5tNp4DqfV7UQjDqv2+RJlNzsDtvNAWccbItceIE=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_prometheus_common",
    importpath = "github.com/prometheus/common",
    sum = "h1:L+1lyG48J1zAQXA3RBX/nG/B3gjlHq0zTt2tlbJLyCY=",
    version = "v0.7.0",
)

go_repository(
    name = "com_github_prometheus_procfs",
    importpath = "github.com/prometheus/procfs",
    sum = "h1:+fpWZdT24pJBiqJdAwYBjPSk+5YmQzYNPYzQsdzLkt8=",
    version = "v0.0.8",
)

go_repository(
    name = "com_github_prometheus_tsdb",
    importpath = "github.com/prometheus/tsdb",
    sum = "h1:YZcsG11NqnK4czYLrWd9mpEuAJIHVQLwdrleYfszMAA=",
    version = "v0.7.1",
)

go_repository(
    name = "com_github_protocolbuffers_txtpbfmt",
    importpath = "github.com/protocolbuffers/txtpbfmt",
    sum = "h1:gSVONBi2HWMFXCa9jFdYvYk7IwW/mTLxWOF7rXS4LO0=",
    version = "v0.0.0-20201118171849-f6a6b3f636fc",
    # build_directives = [
    #     "gazelle:proto disable",
    #     "gazelle:proto disable_global",
    # ],
)

go_repository(
    name = "com_github_pseudomuto_protoc_gen_doc",
    importpath = "github.com/pseudomuto/protoc-gen-doc",
    sum = "h1:61vWZuxYa8D7Rn4h+2dgoTNqnluBmJya2MgbqO32z6g=",
    version = "v1.3.2",
)

go_repository(
    name = "com_github_pseudomuto_protokit",
    importpath = "github.com/pseudomuto/protokit",
    sum = "h1:hlnBDcy3YEDXH7kc9gV+NLaN0cDzhDvD1s7Y6FZ8RpM=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_quasilyte_go_consistent",
    importpath = "github.com/quasilyte/go-consistent",
    sum = "h1:JoUA0uz9U0FVFq5p4LjEq4C0VgQ0El320s3Ms0V4eww=",
    version = "v0.0.0-20190521200055-c6f3937de18c",
)

go_repository(
    name = "com_github_quasilyte_go_ruleguard",
    importpath = "github.com/quasilyte/go-ruleguard",
    sum = "h1:F6l5p6+7WBcTKS7foNQ4wqA39zjn2+RbdbyzGxIq1B0=",
    version = "v0.3.4",
)

go_repository(
    name = "com_github_quasilyte_go_ruleguard_dsl",
    importpath = "github.com/quasilyte/go-ruleguard/dsl",
    sum = "h1:ULi3SLXvDUgb0u2IM5xU6er9KeWBSaUh1NlDjCgLHU8=",
    version = "v0.3.2",
)

go_repository(
    name = "com_github_quasilyte_go_ruleguard_rules",
    importpath = "github.com/quasilyte/go-ruleguard/rules",
    sum = "h1:PeTrJiH/dSeruL/Z9Db39NRMwI/yoA3oHCdCkg+Wh8A=",
    version = "v0.0.0-20210203162857-b223e0831f88",
)

go_repository(
    name = "com_github_quasilyte_regex_syntax",
    importpath = "github.com/quasilyte/regex/syntax",
    sum = "h1:L8QM9bvf68pVdQ3bCFZMDmnt9yqcMBro1pC7F+IPYMY=",
    version = "v0.0.0-20200407221936-30656e2c4a95",
)

go_repository(
    name = "com_github_rakyll_statik",
    importpath = "github.com/rakyll/statik",
    sum = "h1:OF3QCZUuyPxuGEP7B4ypUa7sB/iHtqOTDYZXGM8KOdQ=",
    version = "v0.1.7",
)

go_repository(
    name = "com_github_rogpeppe_fastuuid",
    importpath = "github.com/rogpeppe/fastuuid",
    sum = "h1:gu+uRPtBe88sKxUCEXRoeCvVG90TJmwhiqRpvdhQFng=",
    version = "v0.0.0-20150106093220-6724a57986af",
)

go_repository(
    name = "com_github_rogpeppe_go_internal",
    importpath = "github.com/rogpeppe/go-internal",
    sum = "h1:RR9dF3JtopPvtkroDZuVD7qquD0bnHlKSqaQhgwt8yk=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_rs_cors",
    importpath = "github.com/rs/cors",
    sum = "h1:+88SsELBHx5r+hZ8TCkggzSstaWNbDvThkVK8H6f9ik=",
    version = "v1.7.0",
)

go_repository(
    name = "com_github_russross_blackfriday",
    importpath = "github.com/russross/blackfriday",
    sum = "h1:HyvC0ARfnZBqnXwABFeSZHpKvJHJJfPz81GNueLj0oo=",
    version = "v1.5.2",
)

go_repository(
    name = "com_github_russross_blackfriday_v2",
    importpath = "github.com/russross/blackfriday/v2",
    sum = "h1:lPqVAte+HuHNfhJ/0LC98ESWRz8afy9tM/0RK8m9o+Q=",
    version = "v2.0.1",
)

go_repository(
    name = "com_github_ryancurrah_gomodguard",
    importpath = "github.com/ryancurrah/gomodguard",
    sum = "h1:YWfhGOrXwLGiqcC/u5EqG6YeS8nh+1fw0HEc85CVZro=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_ryanrolds_sqlclosecheck",
    importpath = "github.com/ryanrolds/sqlclosecheck",
    sum = "h1:AZx+Bixh8zdUBxUA1NxbxVAS78vTPq4rCb8OUZI9xFw=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_ryanuber_columnize",
    importpath = "github.com/ryanuber/columnize",
    sum = "h1:UFr9zpz4xgTnIE5yIMtWAMngCdZ9p/+q6lTbgelo80M=",
    version = "v0.0.0-20160712163229-9b3edd62028f",
)

go_repository(
    name = "com_github_ryanuber_go_glob",
    importpath = "github.com/ryanuber/go-glob",
    sum = "h1:7YvPJVmEeFHR1Tj9sZEYsmarJEQfMVYpd/Vyy/A8dqE=",
    version = "v0.0.0-20170128012129-256dc444b735",
)

go_repository(
    name = "com_github_sanposhiho_wastedassign",
    importpath = "github.com/sanposhiho/wastedassign",
    sum = "h1:dB+7OV0iJ5b0SpGwKjKlPCr8GDZJX6Ylm3YG+66xGpc=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_sean_seed",
    importpath = "github.com/sean-/seed",
    sum = "h1:nn5Wsu0esKSJiIVhscUtVbo7ada43DJhG55ua/hjS5I=",
    version = "v0.0.0-20170313163322-e2103e2c3529",
)

go_repository(
    name = "com_github_securego_gosec_v2",
    importpath = "github.com/securego/gosec/v2",
    sum = "h1:mOhJv5w6UyNLpSssQOQCc7eGkKLuicAxvf66Ey/X4xk=",
    version = "v2.7.0",
)

go_repository(
    name = "com_github_sergi_go_diff",
    importpath = "github.com/sergi/go-diff",
    sum = "h1:Kpca3qRNrduNnOQeazBd0ysaKrUJiIuISHxogkT9RPQ=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_shazow_go_diff",
    importpath = "github.com/shazow/go-diff",
    sum = "h1:W65qqJCIOVP4jpqPQ0YvHYKwcMEMVWIzWC5iNQQfBTU=",
    version = "v0.0.0-20160112020656-b6b7b6733b8c",
)

go_repository(
    name = "com_github_shirou_gopsutil",
    importpath = "github.com/shirou/gopsutil",
    sum = "h1:80VN+vGkqM773Br/uNNTSheo3KatTgV8IpjIKjvVLng=",
    version = "v0.0.0-20180427012116-c95755e4bcd7",
)

go_repository(
    name = "com_github_shirou_gopsutil_v3",
    importpath = "github.com/shirou/gopsutil/v3",
    sum = "h1:XB/+p+kVnyYLuPHCfa99lxz2aJyvVhnyd+FxZqH/k7M=",
    version = "v3.21.4",
)

go_repository(
    name = "com_github_shirou_w32",
    importpath = "github.com/shirou/w32",
    sum = "h1:udFKJ0aHUL60LboW/A+DfgoHVedieIzIXE8uylPue0U=",
    version = "v0.0.0-20160930032740-bb4de0191aa4",
)

go_repository(
    name = "com_github_shurcool_go",
    importpath = "github.com/shurcooL/go",
    sum = "h1:MZM7FHLqUHYI0Y/mQAt3d2aYa0SiNms/hFqC9qJYolM=",
    version = "v0.0.0-20180423040247-9e1955d9fb6e",
)

go_repository(
    name = "com_github_shurcool_go_goon",
    importpath = "github.com/shurcooL/go-goon",
    sum = "h1:llrF3Fs4018ePo4+G/HV/uQUqEI1HMDjCeOf2V6puPc=",
    version = "v0.0.0-20170922171312-37c2f522c041",
)

go_repository(
    name = "com_github_shurcool_sanitized_anchor_name",
    importpath = "github.com/shurcooL/sanitized_anchor_name",
    sum = "h1:PdmoCO6wvbs+7yrJyMORt4/BmY5IYyJwS/kOiWx8mHo=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_sirupsen_logrus",
    importpath = "github.com/sirupsen/logrus",
    sum = "h1:dJKuHgqk1NNQlqoA6BTlM1Wf9DOH3NBjQyu0h9+AZZE=",
    version = "v1.8.1",
)

go_repository(
    name = "com_github_smartystreets_assertions",
    importpath = "github.com/smartystreets/assertions",
    sum = "h1:zE9ykElWQ6/NYmHa3jpm/yHnI4xSofP+UP6SpjHcSeM=",
    version = "v0.0.0-20180927180507-b2de0cb4f26d",
)

go_repository(
    name = "com_github_smartystreets_goconvey",
    importpath = "github.com/smartystreets/goconvey",
    sum = "h1:fv0U8FUIMPNf1L9lnHLvLhgicrIVChEkdzIKYqbNC9s=",
    version = "v1.6.4",
)

go_repository(
    name = "com_github_soheilhy_cmux",
    importpath = "github.com/soheilhy/cmux",
    sum = "h1:0HKaf1o97UwFjHH9o5XsHUOF+tqmdA7KEzXLpiyaw0E=",
    version = "v0.1.4",
)

go_repository(
    name = "com_github_sonatard_noctx",
    importpath = "github.com/sonatard/noctx",
    sum = "h1:VC1Qhl6Oxx9vvWo3UDgrGXYCeKCe3Wbw7qAWL6FrmTY=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_sourcegraph_go_diff",
    importpath = "github.com/sourcegraph/go-diff",
    sum = "h1:hmA1LzxW0n1c3Q4YbrFgg4P99GSnebYa3x8gr0HZqLQ=",
    version = "v0.6.1",
)

go_repository(
    name = "com_github_spaolacci_murmur3",
    importpath = "github.com/spaolacci/murmur3",
    sum = "h1:7c1g84S4BPRrfL5Xrdp6fOJ206sU9y293DDHaoy0bLI=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_spf13_afero",
    importpath = "github.com/spf13/afero",
    sum = "h1:xoax2sJ2DT8S8xA2paPFjDCScCNeWsg75VG0DLRreiY=",
    version = "v1.6.0",
)

go_repository(
    name = "com_github_spf13_cast",
    importpath = "github.com/spf13/cast",
    sum = "h1:nFm6S0SMdyzrzcmThSipiEubIDy8WEXKNZ0UOgiRpng=",
    version = "v1.3.1",
)

go_repository(
    name = "com_github_spf13_cobra",
    importpath = "github.com/spf13/cobra",
    sum = "h1:ZlrZ4XsMRm04Fr5pSFxBgfND2EBVa1nLpiy1stUsX/8=",
    version = "v0.0.3",
)

go_repository(
    name = "com_github_spf13_jwalterweatherman",
    importpath = "github.com/spf13/jwalterweatherman",
    sum = "h1:ue6voC5bR5F8YxI5S67j9i582FU4Qvo2bmqnqMYADFk=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_spf13_pflag",
    importpath = "github.com/spf13/pflag",
    sum = "h1:aCvUg6QPl3ibpQUxyLkrEkCHtPqYJL4x9AuhqVqFis4=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_spf13_viper",
    importpath = "github.com/spf13/viper",
    sum = "h1:Kq1fyeebqsBfbjZj4EL7gj2IO0mMaiyjYUWcUsl2O44=",
    version = "v1.8.1",
)

go_repository(
    name = "com_github_ssgreg_nlreturn_v2",
    importpath = "github.com/ssgreg/nlreturn/v2",
    sum = "h1:6/s4Rc49L6Uo6RLjhWZGBpWWjfzk2yrf1nIW8m4wgVA=",
    version = "v2.1.0",
)

go_repository(
    name = "com_github_stackexchange_wmi",
    importpath = "github.com/StackExchange/wmi",
    sum = "h1:G0m3OIz70MZUWq3EgK3CesDbo8upS2Vm9/P3FtgI+Jk=",
    version = "v0.0.0-20190523213315-cbe66965904d",
)

go_repository(
    name = "com_github_stretchr_objx",
    importpath = "github.com/stretchr/objx",
    sum = "h1:2vfRuCMp5sSVIDSqO8oNnWJq7mPa6KVP3iPIwFBuy8A=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_stretchr_testify",
    importpath = "github.com/stretchr/testify",
    sum = "h1:nwc3DEeHmmLAfoZucVR881uASk0Mfjw8xYJ99tb5CcY=",
    version = "v1.7.0",
)

go_repository(
    name = "com_github_subosito_gotenv",
    importpath = "github.com/subosito/gotenv",
    sum = "h1:Slr1R9HxAlEKefgq5jn9U+DnETlIUa6HfgEzj0g5d7s=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_tdakkota_asciicheck",
    importpath = "github.com/tdakkota/asciicheck",
    sum = "h1:HxLVTlqcHhFAz3nWUcuvpH7WuOMv8LQoCWmruLfFH2U=",
    version = "v0.0.0-20200416200610-e657995f937b",
)

go_repository(
    name = "com_github_tetafro_godot",
    importpath = "github.com/tetafro/godot",
    sum = "h1:NCglcF0Ct5vVUeRJVsUz9TPKyxkE/lKv7QYJfjxRuvw=",
    version = "v1.4.6",
)

go_repository(
    name = "com_github_timakin_bodyclose",
    importpath = "github.com/timakin/bodyclose",
    sum = "h1:ig99OeTyDwQWhPe2iw9lwfQVF1KB3Q4fpP3X7/2VBG8=",
    version = "v0.0.0-20200424151742-cb6215831a94",
)

go_repository(
    name = "com_github_tklauser_go_sysconf",
    importpath = "github.com/tklauser/go-sysconf",
    sum = "h1:HT8SVixZd3IzLdfs/xlpq0jeSfTX57g1v6wB1EuzV7M=",
    version = "v0.3.4",
)

go_repository(
    name = "com_github_tklauser_numcpus",
    importpath = "github.com/tklauser/numcpus",
    sum = "h1:ct88eFm+Q7m2ZfXJdan1xYoXKlmwsfP+k88q05KvlZc=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_tmc_grpc_websocket_proxy",
    importpath = "github.com/tmc/grpc-websocket-proxy",
    sum = "h1:ndzgwNDnKIqyCvHTXaCqh9KlOWKvBry6nuXMJmonVsE=",
    version = "v0.0.0-20170815181823-89b8d40f7ca8",
)

go_repository(
    name = "com_github_tomarrell_wrapcheck_v2",
    importpath = "github.com/tomarrell/wrapcheck/v2",
    sum = "h1:LTzwrYlgBUwi9JldazhbJN84fN9nS2UNGrZIo2syqxE=",
    version = "v2.1.0",
)

go_repository(
    name = "com_github_tomasen_realip",
    importpath = "github.com/tomasen/realip",
    sum = "h1:fb190+cK2Xz/dvi9Hv8eCYJYvIGUTN2/KLq1pT6CjEc=",
    version = "v0.0.0-20180522021738-f0c99a92ddce",
)

go_repository(
    name = "com_github_tommy_muehle_go_mnd_v2",
    importpath = "github.com/tommy-muehle/go-mnd/v2",
    sum = "h1:SLkFtxVVkoypCu6eTERr5U2IC3Kce/zOhA4IyNesPV4=",
    version = "v2.3.2",
)

go_repository(
    name = "com_github_ugorji_go",
    importpath = "github.com/ugorji/go",
    sum = "h1:j4s+tAvLfL3bZyefP2SEWmhBzmuIlH/eqNuPdFPgngw=",
    version = "v1.1.4",
)

go_repository(
    name = "com_github_ugorji_go_codec",
    importpath = "github.com/ugorji/go/codec",
    sum = "h1:3SVOIvH7Ae1KRYyQWRjXWJEA9sS/c/pjvH++55Gr648=",
    version = "v0.0.0-20181204163529-d75b2dcb6bc8",
)

go_repository(
    name = "com_github_ultraware_funlen",
    importpath = "github.com/ultraware/funlen",
    sum = "h1:5ylVWm8wsNwH5aWo9438pwvsK0QiqVuUrt9bn7S/iLA=",
    version = "v0.0.3",
)

go_repository(
    name = "com_github_ultraware_whitespace",
    importpath = "github.com/ultraware/whitespace",
    sum = "h1:If7Va4cM03mpgrNH9k49/VOicWpGoG70XPBFFODYDsg=",
    version = "v0.0.4",
)

go_repository(
    name = "com_github_urfave_cli",
    importpath = "github.com/urfave/cli",
    sum = "h1:+mkCCcOFKPnCmVYVcURKps1Xe+3zP90gSYGNfRkjoIY=",
    version = "v1.22.1",
)

go_repository(
    name = "com_github_uudashr_gocognit",
    importpath = "github.com/uudashr/gocognit",
    sum = "h1:MoG2fZ0b/Eo7NXoIwCVFLG5JED3qgQz5/NEE+rOsjPs=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_valyala_bytebufferpool",
    importpath = "github.com/valyala/bytebufferpool",
    sum = "h1:GqA5TC/0021Y/b9FG4Oi9Mr3q7XYx6KllzawFIhcdPw=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_valyala_fasthttp",
    importpath = "github.com/valyala/fasthttp",
    sum = "h1:9zAqOYLl8Tuy3E5R6ckzGDJ1g8+pw15oQp2iL9Jl6gQ=",
    version = "v1.16.0",
)

go_repository(
    name = "com_github_valyala_quicktemplate",
    importpath = "github.com/valyala/quicktemplate",
    sum = "h1:O7EuMwuH7Q94U2CXD6sOX8AYHqQqWtmIk690IhmpkKA=",
    version = "v1.6.3",
)

go_repository(
    name = "com_github_valyala_tcplisten",
    importpath = "github.com/valyala/tcplisten",
    sum = "h1:0R4NLDRDZX6JcmhJgXi5E4b8Wg84ihbmUKp/GvSPEzc=",
    version = "v0.0.0-20161114210144-ceec8f93295a",
)

go_repository(
    name = "com_github_viki_org_dnscache",
    importpath = "github.com/viki-org/dnscache",
    sum = "h1:EVObHAr8DqpoJCVv6KYTle8FEImKhtkfcZetNqxDoJQ=",
    version = "v0.0.0-20130720023526-c70c1f23c5d8",
)

go_repository(
    name = "com_github_wacul_ptr",
    importpath = "github.com/wacul/ptr",
    sum = "h1:FIKu08Wx0YUIf9MNsfF62OCmBSmz5A1Tk65zWhOIL/I=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_xiang90_probing",
    importpath = "github.com/xiang90/probing",
    sum = "h1:eY9dn8+vbi4tKz5Qo6v2eYzo7kUS51QINcR5jNpbZS8=",
    version = "v0.0.0-20190116061207-43a291ad63a2",
)

go_repository(
    name = "com_github_xordataexchange_crypt",
    importpath = "github.com/xordataexchange/crypt",
    sum = "h1:ESFSdwYZvkeru3RtdrYueztKhOBCSAAzS4Gf+k0tEow=",
    version = "v0.0.3-0.20170626215501-b2862e3d0a77",
)

go_repository(
    name = "com_github_yeya24_promlinter",
    importpath = "github.com/yeya24/promlinter",
    sum = "h1:goWULN0jH5Yajmu/K+v1xCqIREeB+48OiJ2uu2ssc7U=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_yudai_gojsondiff",
    importpath = "github.com/yudai/gojsondiff",
    sum = "h1:27cbfqXLVEJ1o8I6v3y9lg8Ydm53EKqHXAOMxEGlCOA=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_yudai_golcs",
    importpath = "github.com/yudai/golcs",
    sum = "h1:BHyfKlQyqbsFN5p3IfnEUduWvb9is428/nNb5L3U01M=",
    version = "v0.0.0-20170316035057-ecda9a501e82",
)

go_repository(
    name = "com_github_yudai_pp",
    importpath = "github.com/yudai/pp",
    sum = "h1:Q4//iY4pNF6yPLZIigmvcl7k/bPgrcTPIFIcmawg5bI=",
    version = "v2.0.1+incompatible",
)

go_repository(
    name = "com_github_yuin_goldmark",
    importpath = "github.com/yuin/goldmark",
    sum = "h1:ruQGxdhGHe7FWOJPT0mKs5+pD2Xs1Bm/kdGlHO04FmM=",
    version = "v1.2.1",
)

go_repository(
    name = "com_google_cloud_go",
    importpath = "cloud.google.com/go",
    sum = "h1:eOI3/cP2VTU6uZLDYAoic+eyzzB9YyGmJ7eIjl8rOPg=",
    version = "v0.34.0",
)

go_repository(
    name = "com_google_cloud_go_bigquery",
    importpath = "cloud.google.com/go/bigquery",
    sum = "h1:PQcPefKFdaIzjQFbiyOgAqyx8q5djaE7x9Sqe712DPA=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_datastore",
    importpath = "cloud.google.com/go/datastore",
    sum = "h1:/May9ojXjRkPBNVrq+oWLqmWCkr4OU5uRY29bu0mRyQ=",
    version = "v1.1.0",
)

go_repository(
    name = "com_google_cloud_go_firestore",
    importpath = "cloud.google.com/go/firestore",
    sum = "h1:9x7Bx0A9R5/M9jibeJeZWqjeVEIxYW9fZYqB9a70/bY=",
    version = "v1.1.0",
)

go_repository(
    name = "com_google_cloud_go_pubsub",
    importpath = "cloud.google.com/go/pubsub",
    sum = "h1:ukjixP1wl0LpnZ6LWtZJ0mX5tBmjp1f8Sqer8Z2OMUU=",
    version = "v1.3.1",
)

go_repository(
    name = "com_google_cloud_go_spanner",
    importpath = "cloud.google.com/go/spanner",
    sum = "h1:mvgDB+f4CfKXebTIi36aqY02eVG4nU/Z3KV6stQaYCc=",
    version = "v1.7.0",
)

go_repository(
    name = "com_google_cloud_go_storage",
    importpath = "cloud.google.com/go/storage",
    sum = "h1:STgFzyU5/8miMl0//zKh2aQeTyeaUH3WN9bSUiJ09bA=",
    version = "v1.10.0",
)

go_repository(
    name = "com_shuralyov_dmitri_gpu_mtl",
    importpath = "dmitri.shuralyov.com/gpu/mtl",
    sum = "h1:+PdD6GLKejR9DizMAKT5DpSAkKswvZrurk1/eEt9+pw=",
    version = "v0.0.0-20201218220906-28db891af037",
)

go_repository(
    name = "com_sourcegraph_sqs_pbtypes",
    importpath = "sourcegraph.com/sqs/pbtypes",
    sum = "h1:JPJh2pk3+X4lXAkZIk2RuE/7/FoK9maXw+TNPJhVS/c=",
    version = "v0.0.0-20180604144634-d3ebe8f20ae4",
)

go_repository(
    name = "in_gopkg_airbrake_gobrake_v2",
    importpath = "gopkg.in/airbrake/gobrake.v2",
    sum = "h1:7z2uVWwn7oVeeugY1DtlPAy5H+KYgB1KeKTnqjNatLo=",
    version = "v2.0.9",
)

go_repository(
    name = "in_gopkg_alecthomas_kingpin_v2",
    importpath = "gopkg.in/alecthomas/kingpin.v2",
    sum = "h1:jMFz6MfLP0/4fUyZle81rXUoxOBFi19VUFKVDOQfozc=",
    version = "v2.2.6",
)

go_repository(
    name = "in_gopkg_check_v1",
    importpath = "gopkg.in/check.v1",
    sum = "h1:YR8cESwS4TdDjEe65xsg0ogRM/Nc3DYOhEAlW+xobZo=",
    version = "v1.0.0-20190902080502-41f04d3bba15",
)

go_repository(
    name = "in_gopkg_cheggaaa_pb_v1",
    importpath = "gopkg.in/cheggaaa/pb.v1",
    sum = "h1:Ev7yu1/f6+d+b3pi5vPdRPc6nNtP1umSfcWiEfRqv6I=",
    version = "v1.0.25",
)

go_repository(
    name = "in_gopkg_errgo_v2",
    importpath = "gopkg.in/errgo.v2",
    sum = "h1:0vLT13EuvQ0hNvakwLuFZ/jYrLp5F3kcWHXdRggjCE8=",
    version = "v2.1.0",
)

go_repository(
    name = "in_gopkg_fsnotify_v1",
    importpath = "gopkg.in/fsnotify.v1",
    sum = "h1:xOHLXZwVvI9hhs+cLKq5+I5onOuwQLhQwiu63xxlHs4=",
    version = "v1.4.7",
)

go_repository(
    name = "in_gopkg_gcfg_v1",
    importpath = "gopkg.in/gcfg.v1",
    sum = "h1:m8OOJ4ccYHnx2f4gQwpno8nAX5OGOh7RLaaz0pj3Ogs=",
    version = "v1.2.3",
)

go_repository(
    name = "in_gopkg_gemnasium_logrus_airbrake_hook_v2",
    importpath = "gopkg.in/gemnasium/logrus-airbrake-hook.v2",
    sum = "h1:OAj3g0cR6Dx/R07QgQe8wkA9RNjB2u4i700xBkIT4e0=",
    version = "v2.1.2",
)

go_repository(
    name = "in_gopkg_ini_v1",
    importpath = "gopkg.in/ini.v1",
    sum = "h1:duBzk771uxoUuOlyRLkHsygud9+5lrlGjdFBb4mSKDU=",
    version = "v1.62.0",
)

go_repository(
    name = "in_gopkg_resty_v1",
    importpath = "gopkg.in/resty.v1",
    sum = "h1:CuXP0Pjfw9rOuY6EP+UvtNvt5DSqHpIxILZKT/quCZI=",
    version = "v1.12.0",
)

go_repository(
    name = "in_gopkg_tomb_v1",
    importpath = "gopkg.in/tomb.v1",
    sum = "h1:uRGJdciOHaEIrze2W8Q3AKkepLTh2hOroT7a+7czfdQ=",
    version = "v1.0.0-20141024135613-dd632973f1e7",
)

go_repository(
    name = "in_gopkg_warnings_v0",
    importpath = "gopkg.in/warnings.v0",
    sum = "h1:wFXVbFY8DY5/xOe1ECiWdKCzZlxgshcYVNkBHstARME=",
    version = "v0.1.2",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    importpath = "gopkg.in/yaml.v2",
    sum = "h1:ZCJp+EgiOT7lHqUV2J862kp8Qj64Jo6az82+3Td9dZw=",
    version = "v2.2.2",
)

go_repository(
    name = "in_gopkg_yaml_v3",
    importpath = "gopkg.in/yaml.v3",
    sum = "h1:dUUwHk2QECo/6vqA44rthZ8ie2QXMNeKRTHCNY2nXvo=",
    version = "v3.0.0-20200313102051-9f266ea9e77c",
)

go_repository(
    name = "io_etcd_go_bbolt",
    importpath = "go.etcd.io/bbolt",
    sum = "h1:MUGmc65QhB3pIlaQ5bB4LwqSj6GIonVJXpZiaKNyaKk=",
    version = "v1.3.3",
)

go_repository(
    name = "io_etcd_go_etcd",
    importpath = "go.etcd.io/etcd",
    sum = "h1:VcrIfasaLFkyjk6KNlXQSzO+B0fZcnECiDrKJsfxka0=",
    version = "v0.0.0-20191023171146-3cf2f69b5738",
)

go_repository(
    name = "io_etcd_go_etcd_api_v3",
    importpath = "go.etcd.io/etcd/api/v3",
    sum = "h1:GsV3S+OfZEOCNXdtNkBSR7kgLobAa/SO6tCxRa0GAYw=",
    version = "v3.5.0",
)

go_repository(
    name = "io_etcd_go_etcd_client_pkg_v3",
    importpath = "go.etcd.io/etcd/client/pkg/v3",
    sum = "h1:2aQv6F436YnN7I4VbI8PPYrBhu+SmrTaADcf8Mi/6PU=",
    version = "v3.5.0",
)

go_repository(
    name = "io_etcd_go_etcd_client_v2",
    importpath = "go.etcd.io/etcd/client/v2",
    sum = "h1:ftQ0nOOHMcbMS3KIaDQ0g5Qcd6bhaBrQT6b89DfwLTs=",
    version = "v2.305.0",
)

go_repository(
    name = "io_k8s_sigs_yaml",
    importpath = "sigs.k8s.io/yaml",
    sum = "h1:4A07+ZFc2wgJwo8YNlQpr1rVlgUDlxXHhPJciaPY5gs=",
    version = "v1.1.0",
)

go_repository(
    name = "io_opencensus_go",
    importpath = "go.opencensus.io",
    sum = "h1:75k/FF0Q2YM8QYo07VPddOLBslDt1MZOdEslOHvmzAs=",
    version = "v0.22.2",
)

go_repository(
    name = "io_opencensus_go_contrib_exporter_stackdriver",
    importpath = "contrib.go.opencensus.io/exporter/stackdriver",
    sum = "h1:ksUxwH3OD5sxkjzEqGxNTl+Xjsmu3BnC/300MhSVTSc=",
    version = "v0.13.4",
)

go_repository(
    name = "io_rsc_binaryregexp",
    importpath = "rsc.io/binaryregexp",
    sum = "h1:HfqmD5MEmC0zvwBuF187nq9mdnXjXsSivRiXN7SmRkE=",
    version = "v0.2.0",
)

go_repository(
    name = "io_rsc_quote_v3",
    importpath = "rsc.io/quote/v3",
    sum = "h1:9JKUTTIUgS6kzR9mK1YuGKv6Nl+DijDNIc0ghT58FaY=",
    version = "v3.1.0",
)

go_repository(
    name = "io_rsc_sampler",
    importpath = "rsc.io/sampler",
    sum = "h1:7uVkIFmeBqHfdjD+gZwtXXI+RODJ2Wc4O7MPEh/QiW4=",
    version = "v1.3.0",
)

go_repository(
    name = "org_bitbucket_creachadair_shell",
    importpath = "bitbucket.org/creachadair/shell",
    sum = "h1:reJflDbKqnlnqb4Oo2pQ1/BqmY/eCWcNGHrIUO8qIzc=",
    version = "v0.0.6",
)

go_repository(
    name = "org_cuelang_go",
    importpath = "cuelang.org/go",
    sum = "h1:GLJblw6m2WGGCA3k1v6Wbk9gTOt2qto48ahO2MmSd6I=",
    version = "v0.4.0",
)

go_repository(
    name = "org_golang_google_api",
    importpath = "google.golang.org/api",
    sum = "h1:oJra/lMfmtm13/rgY/8i3MzjFWYXvQIAKjQ3HqofMk8=",
    version = "v0.3.1",
)

go_repository(
    name = "org_golang_google_appengine",
    importpath = "google.golang.org/appengine",
    sum = "h1:/wp5JvzpHIxhs/dumFmF7BXTf3Z+dd4uXta4kVyO508=",
    version = "v1.4.0",
)

go_repository(
    name = "org_golang_google_genproto",
    importpath = "google.golang.org/genproto",
    sum = "h1:gSJIx1SDwno+2ElGhA4+qG2zF97qiUzTM+rQ0klBOcE=",
    version = "v0.0.0-20190819201941-24fa4b261c55",
)

go_repository(
    name = "org_golang_x_exp",
    importpath = "golang.org/x/exp",
    sum = "h1:ddvpKwqE7dm58PoWjRCmaCiA3DANEW0zWGfNYQD212Y=",
    version = "v0.0.0-20210615023648-acb5c1269671",
)

go_repository(
    name = "org_golang_x_image",
    importpath = "golang.org/x/image",
    sum = "h1:+qEpEAPhDZ1o0x3tHzZTQDArnOixOzGD9HUJfcg0mb4=",
    version = "v0.0.0-20190802002840-cff245a6509b",
)

go_repository(
    name = "org_golang_x_lint",
    importpath = "golang.org/x/lint",
    sum = "h1:5hukYrvBGR8/eNkX5mdUezrA6JiaEZDtJb9Ei+1LlBs=",
    version = "v0.0.0-20190930215403-16217165b5de",
)

go_repository(
    name = "org_golang_x_mobile",
    importpath = "golang.org/x/mobile",
    sum = "h1:kgfVkAEEQXXQ0qc6dH7n6y37NAYmTFmz0YRwrRjgxKw=",
    version = "v0.0.0-20201217150744-e6ae53a27f4f",
)

go_repository(
    name = "org_golang_x_mod",
    importpath = "golang.org/x/mod",
    sum = "h1:Gz96sIWK3OalVv/I/qNygP42zyoKp3xptRVCWRFEBvo=",
    version = "v0.4.2",
)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    sum = "h1:qWPm9rbaAMKs8Bq/9LRpbMqxWRVUAQwMI9fVrssnTfw=",
    version = "v0.0.0-20210226172049-e18ecbb05110",
)

go_repository(
    name = "org_golang_x_oauth2",
    importpath = "golang.org/x/oauth2",
    sum = "h1:Wo7BWFiOk0QRFMLYMqJGFMd9CgUAcGx7V+qEg/h5IBI=",
    version = "v0.0.0-20190226205417-e64efc72b421",
)

go_repository(
    name = "org_golang_x_sync",
    importpath = "golang.org/x/sync",
    sum = "h1:SQFwaSi55rU7vdNs9Yr0Z324VNlrF+0wMqRXT4St8ck=",
    version = "v0.0.0-20201020160332-67f06af15bc9",
)

go_repository(
    name = "org_golang_x_sys_a",
    importpath = "golang.org/x/sys",
    sum = "h1:SrN+KX8Art/Sf4HNj6Zcz06G7VEz+7w9tdXTPOZ7+l4=",
    version = "v0.0.0-20210615035016-665e8c7367d1",
)

go_repository(
    name = "org_golang_x_sys_b",
    importpath = "golang.org/x/sys",
    sum = "h1:gG67DSER+11cZvqIMb8S8bt0vZtiN6xWYARwirrOSfE=",
    version = "v0.0.0-20210510120138-977fb7262007",
)

go_repository(
    name = "org_golang_x_term",
    importpath = "golang.org/x/term",
    sum = "h1:v+OssWQX+hTHEmOBgwxdZxK4zHq3yOs8F9J7mk0PY8E=",
    version = "v0.0.0-20201126162022-7de9c90e9dd1",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    sum = "h1:cokOdA+Jmi5PJGXLlLllQSgYigAEfHXJAERHVMaCc2k=",
    version = "v0.3.3",
)

go_repository(
    name = "org_golang_x_time",
    importpath = "golang.org/x/time",
    sum = "h1:/5xXl8Y5W96D+TtHSlonuFqGHIWVuyCkGJLwGh9JJFs=",
    version = "v0.0.0-20191024005414-555d28b269f0",
)

go_repository(
    name = "org_golang_x_tools",
    importpath = "golang.org/x/tools",
    sum = "h1:po9/4sTYwZU9lPhi1tOrb4hCv3qrhiQ77LZfGa2OjwY=",
    version = "v0.1.0",
)

go_repository(
    name = "org_golang_x_xerrors",
    importpath = "golang.org/x/xerrors",
    sum = "h1:go1bK/D/BFZV2I8cIQd1NKEZ+0owSTG1fDTci4IqFcE=",
    version = "v0.0.0-20200804184101-5ec99f83aff1",
)

go_repository(
    name = "org_mozilla_go_mozlog",
    importpath = "go.mozilla.org/mozlog",
    sum = "h1:rKyWXYDfrVOpMFBion4Pmx5sJbQreQNXycHvm4KwJSg=",
    version = "v0.0.0-20170222151521-4bb13139d403",
)

go_repository(
    name = "org_uber_go_atomic",
    importpath = "go.uber.org/atomic",
    sum = "h1:ADUqmZGgLDDfbSL9ZmPxKTybcoEYHgpYfELNoN+7hsw=",
    version = "v1.7.0",
)

go_repository(
    name = "org_uber_go_multierr",
    importpath = "go.uber.org/multierr",
    sum = "h1:y6IPFStTAIT5Ytl7/XYmHvzXQ7S3g/IeZW9hyZ5thw4=",
    version = "v1.6.0",
)

go_repository(
    name = "org_uber_go_tools",
    importpath = "go.uber.org/tools",
    sum = "h1:0mgffUl7nfd+FpvXMVz4IDEaUSmT1ysygQC7qYo7sG4=",
    version = "v0.0.0-20190618225709-2cfd321de3ee",
)

go_repository(
    name = "org_uber_go_zap",
    importpath = "go.uber.org/zap",
    sum = "h1:uFRZXykJGK9lLY4HtgSw44DnIcAM+kRBP7x5m+NpAOM=",
    version = "v1.16.0",
)

go_repository(
    name = "com_github_cncf_xds_go",
    importpath = "github.com/cncf/xds/go",
    sum = "h1:OZmjad4L3H8ncOIR8rnb5MREYqG8ixi5+WbeUsquF0c=",
    version = "v0.0.0-20210312221358-fbca930ec8ed",
)

go_repository(
    name = "io_opentelemetry_go_proto_otlp",
    importpath = "go.opentelemetry.io/proto/otlp",
    sum = "h1:rwOQPCuKAKmwGKq2aVNnYIibI6wnV7EvzgfTCzcdGg8=",
    version = "v0.7.0",
)

go_repository(
    name = "com_github_c_bata_go_prompt",
    importpath = "github.com/c-bata/go-prompt",
    sum = "h1:POP+nrHE+DfLYx370bedwNhsqmpCUynWPxuHi0C5vZI=",
    version = "v0.2.6",
)

go_repository(
    name = "com_github_juju_ansiterm",
    importpath = "github.com/juju/ansiterm",
    sum = "h1:FaWFmfWdAUKbSCtOU2QjDaorUexogfaMgbipgYATUMU=",
    version = "v0.0.0-20180109212912-720a0952cc2a",
)

go_repository(
    name = "com_github_lunixbochs_vtclean",
    importpath = "github.com/lunixbochs/vtclean",
    sum = "h1:xu2sLAri4lGiovBDQKxl5mrXyESr3gUr5m5SM5+LVb8=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_manifoldco_promptui",
    importpath = "github.com/manifoldco/promptui",
    sum = "h1:R95mMF+McvXZQ7j1g8ucVZE1gLP3Sv6j9vlF9kyRqQo=",
    version = "v0.8.0",
)

go_repository(
    name = "com_github_mattn_go_tty",
    importpath = "github.com/mattn/go-tty",
    sum = "h1:5OfyWorkyO7xP52Mq7tB36ajHDG5OHrmBGIS/DtakQI=",
    version = "v0.0.3",
)

go_repository(
    name = "com_github_pkg_term",
    importpath = "github.com/pkg/term",
    sum = "h1:L3y/h2jkuBVFdWiJvNfYfKmzcCnILw7mJWm2JQuMppw=",
    version = "v1.2.0-beta.2",
)

go_repository(
    name = "com_github_aead_siphash",
    importpath = "github.com/aead/siphash",
    sum = "h1:FwHfE/T45KPKYuuSAKyyvE+oPWcaQ+CUmFW0bPlM+kg=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_afex_hystrix_go",
    importpath = "github.com/afex/hystrix-go",
    sum = "h1:rFw4nCn9iMW+Vajsk51NtYIcwSTkXr+JGrMd36kTDJw=",
    version = "v0.0.0-20180502004556-fa1af6a1f4f5",
)

go_repository(
    name = "com_github_alexbrainman_goissue34681",
    importpath = "github.com/alexbrainman/goissue34681",
    sum = "h1:iW0a5ljuFxkLGPNem5Ui+KBjFJzKg4Fv2fnxe4dvzpM=",
    version = "v0.0.0-20191006012335-3fc7a47baff5",
)

go_repository(
    name = "com_github_andreasbriese_bbloom",
    importpath = "github.com/AndreasBriese/bbloom",
    sum = "h1:PqzgE6kAMi81xWQA2QIVxjWkFHptGgC547vchpUbtFo=",
    version = "v0.0.0-20180913140656-343706a395b7",
)

go_repository(
    name = "com_github_anmitsu_go_shlex",
    importpath = "github.com/anmitsu/go-shlex",
    sum = "h1:kFOfPq6dUM1hTo4JG6LR5AXSUEsOjtdm0kw0FtQtMJA=",
    version = "v0.0.0-20161002113705-648efa622239",
)

go_repository(
    name = "com_github_apache_thrift",
    importpath = "github.com/apache/thrift",
    sum = "h1:5hryIiq9gtn+MiLVn0wP37kb/uTeRZgN08WoCsAhIhI=",
    version = "v0.13.0",
)

go_repository(
    name = "com_github_aryann_difflib",
    importpath = "github.com/aryann/difflib",
    sum = "h1:pv34s756C4pEXnjgPfGYgdhg/ZdajGhyOvzx8k+23nw=",
    version = "v0.0.0-20170710044230-e206f873d14a",
)

go_repository(
    name = "com_github_aws_aws_lambda_go",
    importpath = "github.com/aws/aws-lambda-go",
    sum = "h1:SuCy7H3NLyp+1Mrfp+m80jcbi9KYWAs9/BXwppwRDzY=",
    version = "v1.13.3",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2",
    importpath = "github.com/aws/aws-sdk-go-v2",
    sum = "h1:qZ+woO4SamnH/eEbjM2IDLhRNwIwND/RQyVlBLp3Jqg=",
    version = "v0.18.0",
)

go_repository(
    name = "com_github_benbjohnson_clock",
    importpath = "github.com/benbjohnson/clock",
    sum = "h1:Q92kusRqC1XV2MjkWETPvjJVqKetz1OzxZB7mHJLju8=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_blang_semver_v4",
    importpath = "github.com/blang/semver/v4",
    sum = "h1:1PFHFE6yCCTv8C1TeyNNarDzntLi7wMI5i/pzqYIsAM=",
    version = "v4.0.0",
)

go_repository(
    name = "com_github_bradfitz_go_smtpd",
    importpath = "github.com/bradfitz/go-smtpd",
    sum = "h1:ckJgFhFWywOx+YLEMIJsTb+NV6NexWICk5+AMSuz3ss=",
    version = "v0.0.0-20170404230938-deb6d6237625",
)

go_repository(
    name = "com_github_btcsuite_btcd",
    importpath = "github.com/btcsuite/btcd",
    sum = "h1:aEbSeNALREWXk0G7UdNhR3ayBV7tZ4M2PNmnrCAph6Q=",
    version = "v0.0.0-20190523000118-16327141da8c",
)

go_repository(
    name = "com_github_btcsuite_btclog",
    importpath = "github.com/btcsuite/btclog",
    sum = "h1:bAs4lUbRJpnnkd9VhRV3jjAVU7DJVjMaK+IsvSeZvFo=",
    version = "v0.0.0-20170628155309-84c8d2346e9f",
)

go_repository(
    name = "com_github_btcsuite_btcutil",
    importpath = "github.com/btcsuite/btcutil",
    sum = "h1:yJzD/yFppdVCf6ApMkVy8cUxV0XrxdP9rVf6D87/Mng=",
    version = "v0.0.0-20190425235716-9e5f4b9a998d",
)

go_repository(
    name = "com_github_btcsuite_go_socks",
    importpath = "github.com/btcsuite/go-socks",
    sum = "h1:R/opQEbFEy9JGkIguV40SvRY1uliPX8ifOvi6ICsFCw=",
    version = "v0.0.0-20170105172521-4720035b7bfd",
)

go_repository(
    name = "com_github_btcsuite_goleveldb",
    importpath = "github.com/btcsuite/goleveldb",
    sum = "h1:qdGvebPBDuYDPGi1WCPjy1tGyMpmDK8IEapSsszn7HE=",
    version = "v0.0.0-20160330041536-7834afc9e8cd",
)

go_repository(
    name = "com_github_btcsuite_snappy_go",
    importpath = "github.com/btcsuite/snappy-go",
    sum = "h1:ZA/jbKoGcVAnER6pCHPEkGdZOV7U1oLUedErBHCUMs0=",
    version = "v0.0.0-20151229074030-0bdef8d06723",
)

go_repository(
    name = "com_github_btcsuite_websocket",
    importpath = "github.com/btcsuite/websocket",
    sum = "h1:R8vQdOQdZ9Y3SkEwmHoWBmX1DNXhXZqlTpq6s4tyJGc=",
    version = "v0.0.0-20150119174127-31079b680792",
)

go_repository(
    name = "com_github_btcsuite_winsvc",
    importpath = "github.com/btcsuite/winsvc",
    sum = "h1:J9B4L7e3oqhXOcm+2IuNApwzQec85lE+QaikUcCs+dk=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_buger_jsonparser",
    importpath = "github.com/buger/jsonparser",
    sum = "h1:D21IyuvjDCshj1/qq+pCNd3VZOAEI9jy6Bi131YlXgI=",
    version = "v0.0.0-20181115193947-bf1c66bbce23",
)

go_repository(
    name = "com_github_casbin_casbin_v2",
    importpath = "github.com/casbin/casbin/v2",
    sum = "h1:bTwon/ECRx9dwBy2ewRVr5OiqjeXSGiTUY74sDPQi/g=",
    version = "v2.1.2",
)

go_repository(
    name = "com_github_cenkalti_backoff",
    importpath = "github.com/cenkalti/backoff",
    sum = "h1:tNowT99t7UNflLxfYYSlKYsBpXdEet03Pg2g16Swow4=",
    version = "v2.2.1+incompatible",
)

go_repository(
    name = "com_github_cheekybits_genny",
    importpath = "github.com/cheekybits/genny",
    sum = "h1:uGGa4nei+j20rOSeDeP5Of12XVm7TGUd4dJA9RDitfE=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_cheggaaa_pb",
    importpath = "github.com/cheggaaa/pb",
    sum = "h1:FckUN5ngEk2LpvuG0fw1GEFx6LtyY2pWI/Z2QgCnEYo=",
    version = "v1.0.29",
)

go_repository(
    name = "com_github_clbanning_x2j",
    importpath = "github.com/clbanning/x2j",
    sum = "h1:EdRZT3IeKQmfCSrgo8SZ8V3MEnskuJP0wCYNpe+aiXo=",
    version = "v0.0.0-20191024224557-825249438eec",
)

go_repository(
    name = "com_github_codahale_hdrhistogram",
    importpath = "github.com/codahale/hdrhistogram",
    sum = "h1:qMd81Ts1T2OTKmB4acZcyKaMtRnY5Y44NuXGX2GFJ1w=",
    version = "v0.0.0-20161010025455-3a0bb77429bd",
)

go_repository(
    name = "com_github_crackcomm_go_gitignore",
    importpath = "github.com/crackcomm/go-gitignore",
    sum = "h1:HVTnpeuvF6Owjd5mniCL8DEXo7uYXdQEmOP4FJbV5tg=",
    version = "v0.0.0-20170627025303-887ab5e44cc3",
)

go_repository(
    name = "com_github_cskr_pubsub",
    importpath = "github.com/cskr/pubsub",
    sum = "h1:vlOzMhl6PFn60gRlTQQsIfVwaPB/B/8MziK8FhEPt/0=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_davidlazar_go_crypto",
    importpath = "github.com/davidlazar/go-crypto",
    sum = "h1:pFUpOrbxDR6AkioZ1ySsx5yxlDQZ8stG2b88gTPxgJU=",
    version = "v0.0.0-20200604182044-b73af7476f6c",
)

go_repository(
    name = "com_github_decred_dcrd_lru",
    importpath = "github.com/decred/dcrd/lru",
    sum = "h1:Kbsb1SFDsIlaupWPwsPp+dkxiBY1frcS07PCPgotKz8=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_dgraph_io_badger",
    importpath = "github.com/dgraph-io/badger",
    sum = "h1:6itBiEUtu+gOzXZWn46bM5/qm8LlV6/byR7Yflx/y6M=",
    version = "v1.5.5-0.20190226225317-8115aed38f8f",
)

go_repository(
    name = "com_github_dgraph_io_badger_v3",
    build_directives = [
        "gazelle:proto disable",
        "gazelle:proto disable_global",
    ],
    importpath = "github.com/dgraph-io/badger/v3",
    sum = "h1:zaX53IRg7ycxVlkd5pYdCeFp1FynD6qBGQoQql3R3Hk=",
    version = "v3.2103.1",
)

go_repository(
    name = "com_github_dgraph_io_ristretto",
    importpath = "github.com/dgraph-io/ristretto",
    sum = "h1:Jv3CGQHp9OjuMBSne1485aDpUkTKEcUqF+jm/LuerPI=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_dgryski_go_farm",
    importpath = "github.com/dgryski/go-farm",
    sum = "h1:dDxpBYafY/GYpcl+LS4Bn3ziLPuEdGRkRjYAbSlWxSA=",
    version = "v0.0.0-20190104051053-3adb47b1fb0f",
)

go_repository(
    name = "com_github_eapache_go_resiliency",
    importpath = "github.com/eapache/go-resiliency",
    sum = "h1:1NtRmCAqadE2FN4ZcN6g90TP3uk8cg9rn9eNK2197aU=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_eapache_go_xerial_snappy",
    importpath = "github.com/eapache/go-xerial-snappy",
    sum = "h1:YEetp8/yCZMuEPMUDHG0CW/brkkEp8mzqk2+ODEitlw=",
    version = "v0.0.0-20180814174437-776d5712da21",
)

go_repository(
    name = "com_github_eapache_queue",
    importpath = "github.com/eapache/queue",
    sum = "h1:YOEu7KNc61ntiQlcEeUIoDTJ2o8mQznoNvUhiigpIqc=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_edsrzf_mmap_go",
    importpath = "github.com/edsrzf/mmap-go",
    sum = "h1:CEBF7HpRnUCSJgGUb5h1Gm7e3VkmVDrR8lvWVLtrOFw=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_elgris_jsondiff",
    importpath = "github.com/elgris/jsondiff",
    sum = "h1:QV0ZrfBLpFc2KDk+a4LJefDczXnonRwrYrQJY/9L4dA=",
    version = "v0.0.0-20160530203242-765b5c24c302",
)

go_repository(
    name = "com_github_facebookgo_atomicfile",
    importpath = "github.com/facebookgo/atomicfile",
    sum = "h1:BBso6MBKW8ncyZLv37o+KNyy0HrrHgfnOaGQC2qvN+A=",
    version = "v0.0.0-20151019160806-2de1f203e7d5",
)

go_repository(
    name = "com_github_flynn_go_shlex",
    importpath = "github.com/flynn/go-shlex",
    sum = "h1:BHsljHzVlRcyQhjrss6TZTdY2VfCqZPbv5k3iBFa2ZQ=",
    version = "v0.0.0-20150515145356-3f9db97f8568",
)

go_repository(
    name = "com_github_flynn_noise",
    importpath = "github.com/flynn/noise",
    sum = "h1:DlTHqmzmvcEiKj+4RYo/imoswx/4r6iBlCMfVtrMXpQ=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_francoispqt_gojay",
    importpath = "github.com/francoispqt/gojay",
    sum = "h1:d2m3sFjloqoIUQU3TsHBgj6qg/BVGlTBeHDUmyJnXKk=",
    version = "v1.2.13",
)

go_repository(
    name = "com_github_franela_goblin",
    importpath = "github.com/franela/goblin",
    sum = "h1:gb2Z18BhTPJPpLQWj4T+rfKHYCHxRHCtRxhKKjRidVw=",
    version = "v0.0.0-20200105215937-c9ffbefa60db",
)

go_repository(
    name = "com_github_franela_goreq",
    importpath = "github.com/franela/goreq",
    sum = "h1:a9ENSRDFBUPkJ5lCgVZh26+ZbGyoVJG7yb5SSzF5H54=",
    version = "v0.0.0-20171204163338-bcd34c9993f8",
)

go_repository(
    name = "com_github_frankban_quicktest",
    importpath = "github.com/frankban/quicktest",
    sum = "h1:8sXhOn0uLys67V8EsXLc6eszDs8VXWxL3iRvebPhedY=",
    version = "v1.11.3",
)

go_repository(
    name = "com_github_gabriel_vasile_mimetype",
    importpath = "github.com/gabriel-vasile/mimetype",
    sum = "h1:gaPnPcNor5aZSVCJVSGipcpbgMWiAAj9z182ocSGbHU=",
    version = "v1.1.2",
)

go_repository(
    name = "com_github_gammazero_deque",
    importpath = "github.com/gammazero/deque",
    sum = "h1:f9LnNmq66VDeuAlSAapemq/U7hJ2jpIWa4c09q8Dlik=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_gliderlabs_ssh",
    importpath = "github.com/gliderlabs/ssh",
    sum = "h1:j3L6gSLQalDETeEg/Jg0mGY0/y/N6zI2xX1978P0Uqw=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_go_bindata_go_bindata_v3",
    importpath = "github.com/go-bindata/go-bindata/v3",
    sum = "h1:F0nVttLC3ws0ojc7p60veTurcOm//D4QBODNM7EGrCI=",
    version = "v3.1.3",
)

go_repository(
    name = "com_github_go_check_check",
    importpath = "github.com/go-check/check",
    sum = "h1:0gkP6mzaMqkmpcJYCFOLkIBwI7xFExG03bbkOkCvUPI=",
    version = "v0.0.0-20180628173108-788fd7840127",
)

go_repository(
    name = "com_github_go_errors_errors",
    importpath = "github.com/go-errors/errors",
    sum = "h1:LUHzmkK3GUKUrL/1gfBUxAHzcev3apQlezX/+O7ma6w=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_go_task_slim_sprig",
    importpath = "github.com/go-task/slim-sprig",
    sum = "h1:p104kn46Q8WdvHunIJ9dAyjPVtrBPhSr3KT2yUst43I=",
    version = "v0.0.0-20210107165309-348f09dbbbc0",
)

go_repository(
    name = "com_github_gogo_googleapis",
    importpath = "github.com/gogo/googleapis",
    sum = "h1:kFkMAZBNAn4j7K0GiZr8cRYzejq68VbheufiV3YuyFI=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_golang_lint",
    importpath = "github.com/golang/lint",
    sum = "h1:2hRPrmiwPrp3fQX967rNJIhQPtiGXdlQWAxKbKw3VHA=",
    version = "v0.0.0-20180702182130-06c8688daad7",
)

go_repository(
    name = "com_github_golang_snappy",
    importpath = "github.com/golang/snappy",
    sum = "h1:woRePGFeVFfLKN/pOkfl+p/TAqKOfFu+7KPlMVpok/w=",
    version = "v0.0.0-20180518054509-2e65f85255db",
)

go_repository(
    name = "com_github_google_flatbuffers",
    importpath = "github.com/google/flatbuffers",
    sum = "h1:dicJ2oXwypfwUGnB2/TYWYEKiuk9eYQlQO/AnOHl5mI=",
    version = "v2.0.0+incompatible",
)

go_repository(
    name = "com_github_google_go_github",
    importpath = "github.com/google/go-github",
    sum = "h1:N0LgJ1j65A7kfXrZnUDaYCs/Sf4rEjNlfyDHW9dolSY=",
    version = "v17.0.0+incompatible",
)

go_repository(
    name = "com_github_google_go_querystring",
    importpath = "github.com/google/go-querystring",
    sum = "h1:Xkwi/a1rcvNg1PPYe5vI8GbeBY/jrVuDX5ASuANWTrk=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_google_gopacket",
    importpath = "github.com/google/gopacket",
    sum = "h1:ves8RnFZPGiFnTS0uPQStjwru6uO6h+nlr9j6fL7kF8=",
    version = "v1.1.19",
)

go_repository(
    name = "com_github_googleapis_gax_go",
    importpath = "github.com/googleapis/gax-go",
    sum = "h1:j0GKcs05QVmm7yesiZq2+9cxHkNK9YM6zKx4D2qucQU=",
    version = "v2.0.0+incompatible",
)

go_repository(
    name = "com_github_gorilla_context",
    importpath = "github.com/gorilla/context",
    sum = "h1:AWwleXJkX/nhcU9bZSnZoi3h/qGYqQAGhq6zZe/aQW8=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_gxed_hashland_keccakpg",
    importpath = "github.com/gxed/hashland/keccakpg",
    sum = "h1:wrk3uMNaMxbXiHibbPO4S0ymqJMm41WiudyFSs7UnsU=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_gxed_hashland_murmur3",
    importpath = "github.com/gxed/hashland/murmur3",
    sum = "h1:SheiaIt0sda5K+8FLz952/1iWS9zrnKsEJaOJu4ZbSc=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_hannahhoward_cbor_gen_for",
    importpath = "github.com/hannahhoward/cbor-gen-for",
    sum = "h1:F9k+7wv5OIk1zcq23QpdiL0hfDuXPjuOmMNaC6fgQ0Q=",
    version = "v0.0.0-20200817222906-ea96cece81f1",
)

go_repository(
    name = "com_github_hannahhoward_go_pubsub",
    importpath = "github.com/hannahhoward/go-pubsub",
    sum = "h1:3YKHER4nmd7b5qy5t0GWDTwSn4OyRgfAXSmo6VnryBY=",
    version = "v0.0.0-20200423002714-8d62886cc36e",
)

go_repository(
    name = "com_github_hashicorp_go_memdb",
    importpath = "github.com/hashicorp/go-memdb",
    sum = "h1:RBKHOsnSszpU6vxq80LzC2BaQjuuvoyaQbkLTf7V7g8=",
    version = "v1.3.2",
)

go_repository(
    name = "com_github_hashicorp_go_version",
    importpath = "github.com/hashicorp/go-version",
    sum = "h1:3vNe/fWF5CBgRIguda1meWhsZHy3m8gCJ5wx+dIzX/E=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_hudl_fargo",
    importpath = "github.com/hudl/fargo",
    sum = "h1:0U6+BtN6LhaYuTnIJq4Wyq5cpn6O2kWrxAtcqBmYY6w=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_huin_goupnp",
    importpath = "github.com/huin/goupnp",
    sum = "h1:wg75sLpL6DZqwHQN6E1Cfk6mtfzS45z8OV+ic+DtHRo=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_huin_goutil",
    importpath = "github.com/huin/goutil",
    sum = "h1:vlNjIqmUZ9CMAWsbURYl3a6wZbw7q5RHVvlXTNS/Bs8=",
    version = "v0.0.0-20170803182201-1ca381bf3150",
)

go_repository(
    name = "com_github_influxdata_influxdb1_client",
    importpath = "github.com/influxdata/influxdb1-client",
    sum = "h1:/WZQPMZNsjZ7IlCpsLGdQBINg5bxKQ1K1sh6awxLtkA=",
    version = "v0.0.0-20191209144304-8bf82d3c094d",
)

go_repository(
    name = "com_github_ipfs_bbloom",
    importpath = "github.com/ipfs/bbloom",
    sum = "h1:Gi+8EGJ2y5qiD5FbsbpX/TMNcJw8gSqr7eyjHa4Fhvs=",
    version = "v0.0.4",
)

go_repository(
    name = "com_github_ipfs_go_bitswap",
    build_directives = [
        "gazelle:proto disable",
        "gazelle:proto disable_global",
    ],
    importpath = "github.com/ipfs/go-bitswap",
    sum = "h1:28YsHYw9ut6wootnImPXH0WpnU5Dbo3qm6cvQ6e6wYY=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_ipfs_go_block_format",
    importpath = "github.com/ipfs/go-block-format",
    sum = "h1:r8t66QstRp/pd/or4dpnbVfXT5Gt7lOqRvC+/dDTpMc=",
    version = "v0.0.3",
)

go_repository(
    name = "com_github_ipfs_go_blockservice",
    importpath = "github.com/ipfs/go-blockservice",
    sum = "h1:dh2i7xjMbCtf0ZSMyQAF2qpV/pEEmM7yVpQ00+gik6U=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_ipfs_go_cid",
    importpath = "github.com/ipfs/go-cid",
    sum = "h1:YN33LQulcRHjfom/i25yoOZR4Telp1Hr/2RU3d0PnC0=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_ipfs_go_cidutil",
    importpath = "github.com/ipfs/go-cidutil",
    sum = "h1:CNOboQf1t7Qp0nuNh8QMmhJs0+Q//bRL1axtCnIB1Yo=",
    version = "v0.0.2",
)

go_repository(
    name = "com_github_ipfs_go_datastore",
    importpath = "github.com/ipfs/go-datastore",
    sum = "h1:h8/n7WPzhp239kkLws+epN3Ic7YtcBPgcaXfEfdVDWM=",
    version = "v0.4.2",
)

go_repository(
    name = "com_github_ipfs_go_detect_race",
    importpath = "github.com/ipfs/go-detect-race",
    sum = "h1:qX/xay2W3E4Q1U7d9lNs1sU9nvguX0a7319XbyQ6cOk=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_ipfs_go_ds_badger",
    importpath = "github.com/ipfs/go-ds-badger",
    sum = "h1:7ToQt7QByBhOTuZF2USMv+PGlMcBC7FW7FdgQ4FCsoo=",
    version = "v0.0.2",
)

go_repository(
    name = "com_github_ipfs_go_ds_flatfs",
    importpath = "github.com/ipfs/go-ds-flatfs",
    sum = "h1:4QceuKEbH+HVZ2ZommstJMi3o3II+dWS3IhLaD7IGHs=",
    version = "v0.4.5",
)

go_repository(
    name = "com_github_ipfs_go_ds_leveldb",
    importpath = "github.com/ipfs/go-ds-leveldb",
    sum = "h1:Z0lsTFciec9qYsyngAw1f/czhRU35qBLR2vhavPFgqA=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_ipfs_go_ds_measure",
    importpath = "github.com/ipfs/go-ds-measure",
    sum = "h1:vE4TyY4aeLeVgnnPBC5QzKIjKrqzha0NCujTfgvVbVQ=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_ipfs_go_fetcher",
    importpath = "github.com/ipfs/go-fetcher",
    sum = "h1:f8uY8AlQ7uB9tzkbzCN0ziPb/LuNJXSVDn65NJMYOXM=",
    version = "v1.4.0",
)

go_repository(
    name = "com_github_ipfs_go_filestore",
    build_directives = [
        "gazelle:proto disable",
        "gazelle:proto disable_global",
    ],
    importpath = "github.com/ipfs/go-filestore",
    sum = "h1:f+Z1sF+BZ4YEKkrVx1oqsWO/94UQ38FCu2QIF02ttYM=",
    version = "v1.0.1-0.20210602165910-58393b4f5f66",
)

go_repository(
    name = "com_github_ipfs_go_fs_lock",
    importpath = "github.com/ipfs/go-fs-lock",
    sum = "h1:sn3TWwNVQqSeNjlWy6zQ1uUGAZrV3hPOyEA6y1/N2a0=",
    version = "v0.0.6",
)

go_repository(
    name = "com_github_ipfs_go_graphsync",
    build_directives = [
        "gazelle:proto disable",
        "gazelle:proto disable_global",
    ],
    importpath = "github.com/ipfs/go-graphsync",
    sum = "h1:Zhh6QdTqdipYHD71ncLO8eA6c8EGUTOoJ4Rqybw3K+o=",
    version = "v0.8.0",
)

go_repository(
    name = "com_github_ipfs_go_ipfs",
    importpath = "github.com/ipfs/go-ipfs",
    sum = "h1:fqilyDPFCEMr3AUFWG/UOgwleTrzt4Ou2UuY7+/Dvcg=",
    version = "v0.9.1",
)

go_repository(
    name = "com_github_ipfs_go_ipfs_blockstore",
    importpath = "github.com/ipfs/go-ipfs-blockstore",
    sum = "h1:RDhK6fdg5YsonkpMuMpdvk/pRtOQlrIRIybuQfkvB2M=",
    version = "v1.0.3",
)

go_repository(
    name = "com_github_ipfs_go_ipfs_blocksutil",
    importpath = "github.com/ipfs/go-ipfs-blocksutil",
    sum = "h1:Eh/H4pc1hsvhzsQoMEP3Bke/aW5P5rVM1IWFJMcGIPQ=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_ipfs_go_ipfs_chunker",
    importpath = "github.com/ipfs/go-ipfs-chunker",
    sum = "h1:ojCf7HV/m+uS2vhUGWcogIIxiO5ubl5O57Q7NapWLY8=",
    version = "v0.0.5",
)

go_repository(
    name = "com_github_ipfs_go_ipfs_cmds",
    importpath = "github.com/ipfs/go-ipfs-cmds",
    sum = "h1:yAxdowQZzoFKjcLI08sXVNnqVj3jnABbf9smrPQmBsw=",
    version = "v0.6.0",
)

go_repository(
    name = "com_github_ipfs_go_ipfs_config",
    importpath = "github.com/ipfs/go-ipfs-config",
    sum = "h1:CBtIYyp/iWIczCv83bmfge8EA2KqxOOfqmETs3tUnnU=",
    version = "v0.16.0",
)

go_repository(
    name = "com_github_ipfs_go_ipfs_delay",
    importpath = "github.com/ipfs/go-ipfs-delay",
    sum = "h1:r/UXYyRcddO6thwOnhiznIAiSvxMECGgtv35Xs1IeRQ=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_ipfs_go_ipfs_ds_help",
    importpath = "github.com/ipfs/go-ipfs-ds-help",
    sum = "h1:bEQ8hMGs80h0sR8O4tfDgV6B01aaF9qeTrujrTLYV3g=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_ipfs_go_ipfs_exchange_interface",
    importpath = "github.com/ipfs/go-ipfs-exchange-interface",
    sum = "h1:LJXIo9W7CAmugqI+uofioIpRb6rY30GUu7G6LUfpMvM=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_ipfs_go_ipfs_exchange_offline",
    importpath = "github.com/ipfs/go-ipfs-exchange-offline",
    sum = "h1:P56jYKZF7lDDOLx5SotVh5KFxoY6C81I1NSHW1FxGew=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_ipfs_go_ipfs_files",
    importpath = "github.com/ipfs/go-ipfs-files",
    sum = "h1:8o0oFJkJ8UkO/ABl8T6ac6tKF3+NIpj67aAB6ZpusRg=",
    version = "v0.0.8",
)

go_repository(
    name = "com_github_ipfs_go_ipfs_keystore",
    importpath = "github.com/ipfs/go-ipfs-keystore",
    sum = "h1:Fa9xg9IFD1VbiZtrNLzsD0GuELVHUFXCWF64kCPfEXU=",
    version = "v0.0.2",
)

go_repository(
    name = "com_github_ipfs_go_ipfs_pinner",
    importpath = "github.com/ipfs/go-ipfs-pinner",
    sum = "h1:iJd1gwILGQJSZhhI0jn6yFOLg34Ua7fdKcB6mXp6k/M=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_ipfs_go_ipfs_posinfo",
    importpath = "github.com/ipfs/go-ipfs-posinfo",
    sum = "h1:Esoxj+1JgSjX0+ylc0hUmJCOv6V2vFoZiETLR6OtpRs=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_ipfs_go_ipfs_pq",
    importpath = "github.com/ipfs/go-ipfs-pq",
    sum = "h1:zgUotX8dcAB/w/HidJh1zzc1yFq6Vm8J7T2F4itj/RU=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_ipfs_go_ipfs_provider",
    importpath = "github.com/ipfs/go-ipfs-provider",
    sum = "h1:kZj72jzWLtGcorlwnMvBL6y6KJk6klO2Kb8QSeqEB0o=",
    version = "v0.5.1",
)

go_repository(
    name = "com_github_ipfs_go_ipfs_routing",
    importpath = "github.com/ipfs/go-ipfs-routing",
    sum = "h1:gAJTT1cEeeLj6/DlLX6t+NxD9fQe2ymTO6qWRDI/HQQ=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_ipfs_go_ipfs_util",
    importpath = "github.com/ipfs/go-ipfs-util",
    sum = "h1:59Sswnk1MFaiq+VcaknX7aYEyGyGDAA73ilhEK2POp8=",
    version = "v0.0.2",
)

go_repository(
    name = "com_github_ipfs_go_ipld_cbor",
    importpath = "github.com/ipfs/go-ipld-cbor",
    sum = "h1:ovz4CHKogtG2KB/h1zUp5U0c/IzZrL435rCh5+K/5G8=",
    version = "v0.0.5",
)

go_repository(
    name = "com_github_ipfs_go_ipld_format",
    importpath = "github.com/ipfs/go-ipld-format",
    sum = "h1:xGlJKkArkmBvowr+GMCX0FEZtkro71K1AwiKnL37mwA=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_ipfs_go_ipld_git",
    importpath = "github.com/ipfs/go-ipld-git",
    sum = "h1:fQv2Alq72g6mH+heDWQ9Awu5FQYc3hcCUVtzuWj/Mno=",
    version = "v0.0.4",
)

go_repository(
    name = "com_github_ipfs_go_ipld_legacy",
    importpath = "github.com/ipfs/go-ipld-legacy",
    sum = "h1:wxkkc4k8cnvIGIjPO0waJCe7SHEyFgl+yQdafdjGrpA=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_ipfs_go_ipns",
    build_directives = [
        "gazelle:proto disable",
    ],
    importpath = "github.com/ipfs/go-ipns",
    sum = "h1:O/s/0ht+4Jl9+VoxoUo0zaHjnZUS+aBQIKTuzdZ/ucI=",
    version = "v0.1.2",
)

go_repository(
    name = "com_github_ipfs_go_log",
    importpath = "github.com/ipfs/go-log",
    sum = "h1:9XTUN/rW64BCG1YhPK9Hoy3q8nr4gOmHHBpgFdfw6Lc=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_ipfs_go_log_v2",
    importpath = "github.com/ipfs/go-log/v2",
    sum = "h1:31Re/cPqFHpsRHgyVwjWADPoF0otB1WrjTy8ZFYwEZU=",
    version = "v2.3.0",
)

go_repository(
    name = "com_github_ipfs_go_metrics_interface",
    importpath = "github.com/ipfs/go-metrics-interface",
    sum = "h1:j+cpbjYvu4R8zbleSs36gvB7jR+wsL2fGD6n0jO4kdg=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_ipfs_go_metrics_prometheus",
    importpath = "github.com/ipfs/go-metrics-prometheus",
    sum = "h1:9i2iljLg12S78OhC6UAiXi176xvQGiZaGVF1CUVdE+s=",
    version = "v0.0.2",
)

go_repository(
    name = "com_github_ipfs_go_mfs",
    importpath = "github.com/ipfs/go-mfs",
    sum = "h1:DlelNSmH+yz/Riy0RjPKlooPg0KML4lXGdLw7uZkfAg=",
    version = "v0.1.2",
)

go_repository(
    name = "com_github_ipfs_go_namesys",
    build_directives = [
        "gazelle:proto disable",
        "gazelle:proto disable_global",
    ],
    importpath = "github.com/ipfs/go-namesys",
    sum = "h1:6lytKWj1rG0Ot6J0nTHvFw+06q1a6n7DLA2CbSGmZco=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_ipfs_go_path",
    importpath = "github.com/ipfs/go-path",
    sum = "h1:BIi831cNED8YnIlIKo9y1SI3u+E+FwQQD+rIIw8PwFA=",
    version = "v0.0.9",
)

go_repository(
    name = "com_github_ipfs_go_peertaskqueue",
    importpath = "github.com/ipfs/go-peertaskqueue",
    sum = "h1:bpRbgv76eT4avutNPDFZuCPOQus6qTgurEYxfulgZW4=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_ipfs_go_pinning_service_http_client",
    importpath = "github.com/ipfs/go-pinning-service-http-client",
    sum = "h1:Au0P4NglL5JfzhNSZHlZ1qra+IcJyO3RWMd9EYCwqSY=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_ipfs_go_unixfs",
    importpath = "github.com/ipfs/go-unixfs",
    sum = "h1:VsZwK3Z6+rjFxha87GBrp3kZHDsztSIuKlsScr3Iw4s=",
    version = "v0.2.3",
)

go_repository(
    name = "com_github_ipfs_go_unixfsnode",
    importpath = "github.com/ipfs/go-unixfsnode",
    sum = "h1:aTsCdhwU0F4dMShMwYGroAj4v4EzSONLdoENebvTRb0=",
    version = "v1.1.2",
)

go_repository(
    name = "com_github_ipfs_go_verifcid",
    importpath = "github.com/ipfs/go-verifcid",
    sum = "h1:m2HI7zIuR5TFyQ1b79Da5N9dnnCP1vcu2QqawmWlK2E=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_ipfs_interface_go_ipfs_core",
    importpath = "github.com/ipfs/interface-go-ipfs-core",
    sum = "h1:+mUiamyHIwedqP8ZgbCIwpy40oX7QcXUbo4CZOeJVJg=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_ipfs_tar_utils",
    importpath = "github.com/ipfs/tar-utils",
    sum = "h1:8Na0KBD6GddGyXwU4rXNtVTE24iuZws8mENJQPLG7W4=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_ipld_go_car",
    importpath = "github.com/ipld/go-car",
    sum = "h1:WT+3cdmXlvmWOlGxk9webhj4auGO5QvgqC2vCCkFRXs=",
    version = "v0.3.1",
)

go_repository(
    name = "com_github_ipld_go_codec_dagpb",
    importpath = "github.com/ipld/go-codec-dagpb",
    sum = "h1:2umV7ud8HBMkRuJgd8gXw95cLhwmcYrihS3cQEy9zpI=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_ipld_go_ipld_prime",
    importpath = "github.com/ipld/go-ipld-prime",
    sum = "h1:JapyKWTsJgmhrPI7hfx4V798c/RClr85sXfBZnH1VIw=",
    version = "v0.12.0",
)

go_repository(
    name = "com_github_ipld_go_ipld_prime_v12",
    importpath = "github.com/ipld/go-ipld-prime",
    sum = "h1:JapyKWTsJgmhrPI7hfx4V798c/RClr85sXfBZnH1VIw=",
    version = "v0.12.0",
)

go_repository(
    name = "com_github_jackpal_gateway",
    importpath = "github.com/jackpal/gateway",
    sum = "h1:qzXWUJfuMdlLMtt0a3Dgt+xkWQiA5itDEITVJtuSwMc=",
    version = "v1.0.5",
)

go_repository(
    name = "com_github_jackpal_go_nat_pmp",
    importpath = "github.com/jackpal/go-nat-pmp",
    sum = "h1:i0LektDkO1QlrTm/cSuP+PyBCDnYvjPLGl4LdWEMiaA=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_jbenet_go_cienv",
    importpath = "github.com/jbenet/go-cienv",
    sum = "h1:Vc/s0QbQtoxX8MwwSLWWh+xNNZvM3Lw7NsTcHrvvhMc=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_jbenet_go_random",
    importpath = "github.com/jbenet/go-random",
    sum = "h1:uUx61FiAa1GI6ZmVd2wf2vULeQZIKG66eybjNXKYCz4=",
    version = "v0.0.0-20190219211222-123a90aedc0c",
)

go_repository(
    name = "com_github_jbenet_go_temp_err_catcher",
    importpath = "github.com/jbenet/go-temp-err-catcher",
    sum = "h1:vhC1OXXiT9R2pczegwz6moDvuRpggaroAXhPIseh57A=",
    version = "v0.0.0-20150120210811-aac704a3f4f2",
)

go_repository(
    name = "com_github_jbenet_goprocess",
    importpath = "github.com/jbenet/goprocess",
    sum = "h1:YKyIEECS/XvcfHtBzxtjBBbWK+MbvA6dG8ASiqwvr10=",
    version = "v0.1.3",
)

go_repository(
    name = "com_github_jellevandenhooff_dkim",
    importpath = "github.com/jellevandenhooff/dkim",
    sum = "h1:ujPKutqRlJtcfWk6toYVYagwra7HQHbXOaS171b4Tg8=",
    version = "v0.0.0-20150330215556-f50fe3d243e1",
)

go_repository(
    name = "com_github_jessevdk_go_flags",
    importpath = "github.com/jessevdk/go-flags",
    sum = "h1:12K8AlpT0/6QUXSfV0yi4Q0jkbq8NDtIKFtF61AoqV0=",
    version = "v0.0.0-20141203071132-1679536dcc89",
)

go_repository(
    name = "com_github_jpillora_backoff",
    importpath = "github.com/jpillora/backoff",
    sum = "h1:uvFg412JmmHBHw7iwprIxkPMI+sGQ4kzOWsMeHnm2EA=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_jrick_logrotate",
    importpath = "github.com/jrick/logrotate",
    sum = "h1:lQ1bL/n9mBNeIXoTUoYRlK4dHuNJVofX9oWqBtPnSzI=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_kami_zh_go_capturer",
    importpath = "github.com/kami-zh/go-capturer",
    sum = "h1:cVtBfNW5XTHiKQe7jDaDBSh/EVM4XLPutLAGboIXuM0=",
    version = "v0.0.0-20171211120116-e492ea43421d",
)

go_repository(
    name = "com_github_kkdai_bstream",
    importpath = "github.com/kkdai/bstream",
    sum = "h1:FOOIBWrEkLgmlgGfMuZT83xIwfPDxEI2OHu6xUmJMFE=",
    version = "v0.0.0-20161212061736-f391b8402d23",
)

go_repository(
    name = "com_github_klauspost_cpuid_v2_a",
    importpath = "github.com/klauspost/cpuid/v2",
    sum = "h1:g0I61F2K2DjRHz1cnxlkNSBIaePVoJIjjnHui8QHbiw=",
    version = "v2.0.4",
)

go_repository(
    name = "com_github_klauspost_cpuid_v2_b",
    importpath = "github.com/klauspost/cpuid/v2",
    sum = "h1:lgaqFMSdTdQYdZ04uHyN2d/eKdOMyi2YLSvlQIBFYa4=",
    version = "v2.0.9",
)

go_repository(
    name = "com_github_knetic_govaluate",
    importpath = "github.com/Knetic/govaluate",
    sum = "h1:1G1pk05UrOh0NlF1oeaaix1x8XzrfjIDK47TY0Zehcw=",
    version = "v3.0.1-0.20171022003610-9aa49832a739+incompatible",
)

go_repository(
    name = "com_github_koron_go_ssdp",
    importpath = "github.com/koron/go-ssdp",
    sum = "h1:wxtKgYHEncAU00muMD06dzLiahtGM1eouRNOzVV7tdQ=",
    version = "v0.0.0-20180514024734-4a0ed625a78b",
)

go_repository(
    name = "com_github_kubuxu_go_os_helper",
    importpath = "github.com/Kubuxu/go-os-helper",
    sum = "h1:EJiD2VUQyh5A9hWJLmc6iWg6yIcJ7jpBcwC8GMGXfDk=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_libp2p_go_addr_util",
    importpath = "github.com/libp2p/go-addr-util",
    sum = "h1:TpTQm9cXVRVSKsYbgQ7GKc3KbbHVTnbostgGaDEP+88=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_libp2p_go_buffer_pool",
    importpath = "github.com/libp2p/go-buffer-pool",
    sum = "h1:QNK2iAFa8gjAe1SPz6mHSMuCcjs+X1wlHzeOSqcmlfs=",
    version = "v0.0.2",
)

go_repository(
    name = "com_github_libp2p_go_cidranger",
    importpath = "github.com/libp2p/go-cidranger",
    sum = "h1:ewPN8EZ0dd1LSnrtuwd4709PXVcITVeuwbag38yPW7c=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_libp2p_go_conn_security",
    importpath = "github.com/libp2p/go-conn-security",
    sum = "h1:4kMMrqrt9EUNCNjX1xagSJC+bq16uqjMe9lk1KBMVNs=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_libp2p_go_conn_security_multistream",
    importpath = "github.com/libp2p/go-conn-security-multistream",
    sum = "h1:aqGmto+ttL/uJgX0JtQI0tD21CIEy5eYd1Hlp0juHY0=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_libp2p_go_doh_resolver",
    importpath = "github.com/libp2p/go-doh-resolver",
    sum = "h1:1wbVGkB4Tdj4WEvjAuYknOPyt4vSSDn9thnj13pKPaY=",
    version = "v0.3.1",
)

go_repository(
    name = "com_github_libp2p_go_eventbus",
    importpath = "github.com/libp2p/go-eventbus",
    sum = "h1:VanAdErQnpTioN2TowqNcOijf6YwhuODe4pPKSDpxGc=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_libp2p_go_flow_metrics",
    importpath = "github.com/libp2p/go-flow-metrics",
    sum = "h1:0gxuFd2GuK7IIP5pKljLwps6TvcuYgvG7Atqi3INF5s=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_libp2p_go_libp2p",
    importpath = "github.com/libp2p/go-libp2p",
    sum = "h1:8VXadcPNni74ODoZ+7326LMAppFYmz1fRQOUuT5iZvQ=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_asn_util",
    importpath = "github.com/libp2p/go-libp2p-asn-util",
    sum = "h1:BM7aaOF7RpmNn9+9g6uTjGJ0cTzWr5j9i9IKeun2M8U=",
    version = "v0.0.0-20200825225859-85005c6cf052",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_autonat",
    importpath = "github.com/libp2p/go-libp2p-autonat",
    sum = "h1:aCWAu43Ri4nU0ZPO7NyLzUvvfqd0nE3dX0R/ZGYVgOU=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_blankhost",
    importpath = "github.com/libp2p/go-libp2p-blankhost",
    sum = "h1:X919sCh+KLqJcNRApj43xCSiQRYqOSI88Fdf55ngf78=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_circuit",
    importpath = "github.com/libp2p/go-libp2p-circuit",
    sum = "h1:eniLL3Y9aq/sryfyV1IAHj5rlvuyj3b7iz8tSiZpdhY=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_connmgr",
    importpath = "github.com/libp2p/go-libp2p-connmgr",
    sum = "h1:TMS0vc0TCBomtQJyWr7fYxcVYYhx+q/2gF++G5Jkl/w=",
    version = "v0.2.4",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_core",
    build_directives = [
        "gazelle:proto disable",
        "gazelle:proto disable_global",
    ],
    importpath = "github.com/libp2p/go-libp2p-core",
    sum = "h1:86uOwW+O6Uc7NbaK4diuLZo2/Ikvqw2rgyV03VcSbLE=",
    version = "v0.0.2",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_crypto",
    importpath = "github.com/libp2p/go-libp2p-crypto",
    sum = "h1:k9MFy+o2zGDNGsaoZl0MA3iZ75qXxr9OOoAZF+sD5OQ=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_discovery",
    importpath = "github.com/libp2p/go-libp2p-discovery",
    sum = "h1:j+R6cokKcGbnZLf4kcNwpx6mDEUPF3N6SrqMymQhmvs=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_gostream",
    importpath = "github.com/libp2p/go-libp2p-gostream",
    sum = "h1:rnas//vRdHYCr7bjraZJISPwZV8OGMjeX5k5fN5Ax44=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_host",
    importpath = "github.com/libp2p/go-libp2p-host",
    sum = "h1:BB/1Z+4X0rjKP5lbQTmjEjLbDVbrcmLOlA6QDsN5/j4=",
    version = "v0.0.3",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_http",
    importpath = "github.com/libp2p/go-libp2p-http",
    sum = "h1:GYeVd+RZzkRa8XFLITqOpcrIQG6KbFLPJqII6HHBHzY=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_interface_connmgr",
    importpath = "github.com/libp2p/go-libp2p-interface-connmgr",
    sum = "h1:KG/KNYL2tYzXAfMvQN5K1aAGTYSYUMJ1prgYa2/JI1E=",
    version = "v0.0.5",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_interface_pnet",
    importpath = "github.com/libp2p/go-libp2p-interface-pnet",
    sum = "h1:7GnzRrBTJHEsofi1ahFdPN9Si6skwXQE9UqR2S+Pkh8=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_kad_dht",
    build_directives = [
        "gazelle:proto disable",
        "gazelle:proto disable_global",
    ],
    importpath = "github.com/libp2p/go-libp2p-kad-dht",
    sum = "h1:wQgzOpoc+dcPVDb3h0HNWUjon5JiYEqsA4iNBUtIA7A=",
    version = "v0.13.1",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_kbucket",
    importpath = "github.com/libp2p/go-libp2p-kbucket",
    sum = "h1:spZAcgxifvFZHBD8tErvppbnNiKA5uokDu3CV7axu70=",
    version = "v0.4.7",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_loggables",
    importpath = "github.com/libp2p/go-libp2p-loggables",
    sum = "h1:h3w8QFfCt2UJl/0/NW4K829HX/0S4KD31PQ7m8UXXO8=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_metrics",
    importpath = "github.com/libp2p/go-libp2p-metrics",
    sum = "h1:yumdPC/P2VzINdmcKZd0pciSUCpou+s0lwYCjBbzQZU=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_mplex",
    importpath = "github.com/libp2p/go-libp2p-mplex",
    sum = "h1:E1xaJBQnbSiTHGI1gaBKmKhu1TUKkErKJnE8iGvirYI=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_nat",
    importpath = "github.com/libp2p/go-libp2p-nat",
    sum = "h1:+KXK324yaY701On8a0aGjTnw8467kW3ExKcqW2wwmyw=",
    version = "v0.0.4",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_net",
    importpath = "github.com/libp2p/go-libp2p-net",
    sum = "h1:qP06u4TYXfl7uW/hzqPhlVVTSA2nw1B/bHBJaUnbh6M=",
    version = "v0.0.2",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_netutil",
    importpath = "github.com/libp2p/go-libp2p-netutil",
    sum = "h1:zscYDNVEcGxyUpMd0JReUZTrpMfia8PmLKcKF72EAMQ=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_noise",
    importpath = "github.com/libp2p/go-libp2p-noise",
    sum = "h1:MRt5XGfYziDXIUy2udtMWfPmzZqUDYoC1FZoKnqPzwk=",
    version = "v0.2.2",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_peer",
    importpath = "github.com/libp2p/go-libp2p-peer",
    sum = "h1:EQ8kMjaCUwt/Y5uLgjT8iY2qg0mGUT0N1zUjer50DsY=",
    version = "v0.2.0",
)

# v0.2.8 is released, but causes issues due to its reliance on a now
# deprecated package. Skip this version; try a later version once
# released.
#
# NB: don't look at GitHub releases, look at the tags. Not every tag is
# added to the repo "releases" page.
go_repository(
    name = "com_github_libp2p_go_libp2p_peerstore",
    importpath = "github.com/libp2p/go-libp2p-peerstore",
    sum = "h1:MKh7pRNPHSh1fLPj8u/M/s/napdmeNpoi9BRy9lPN0E=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_pnet",
    importpath = "github.com/libp2p/go-libp2p-pnet",
    sum = "h1:J6htxttBipJujEjz1y0a5+eYoiPcFHhSYHH6na5f0/k=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_protocol",
    importpath = "github.com/libp2p/go-libp2p-protocol",
    sum = "h1:HdqhEyhg0ToCaxgMhnOmUO8snQtt/kQlcjVk3UoJU3c=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_pubsub",
    build_directives = [
        "gazelle:proto disable",
        "gazelle:proto disable_global",
    ],
    importpath = "github.com/libp2p/go-libp2p-pubsub",
    sum = "h1:rHl9/Xok4zX3zgi0pg0XnUj9Xj2OeXO8oTu85q2+YA8=",
    version = "v0.5.4",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_pubsub_router",
    importpath = "github.com/libp2p/go-libp2p-pubsub-router",
    sum = "h1:KjzTLIOBCt0+/4wH6epTxD/Qu4Up/IyeKHlj9MhWRJI=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_quic_transport",
    importpath = "github.com/libp2p/go-libp2p-quic-transport",
    sum = "h1:p1YQDZRHH4Cv2LPtHubqlQ9ggz4CKng/REZuXZbZMhM=",
    version = "v0.11.2",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_record",
    build_directives = [
        "gazelle:proto disable",
        "gazelle:proto disable_global",
    ],
    importpath = "github.com/libp2p/go-libp2p-record",
    sum = "h1:wHwBGbFzymoIl69BpgwIu0O6ta3TXGcMPvHUAcodzRc=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_routing",
    importpath = "github.com/libp2p/go-libp2p-routing",
    sum = "h1:hPMAWktf9rYi3ME4MG48qE7dq1ofJxiQbfdvpNntjhc=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_routing_helpers",
    importpath = "github.com/libp2p/go-libp2p-routing-helpers",
    sum = "h1:xY61alxJ6PurSi+MXbywZpelvuU4U4p/gPTxjqCqTzY=",
    version = "v0.2.3",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_secio",
    importpath = "github.com/libp2p/go-libp2p-secio",
    sum = "h1:NNP5KLxuP97sE5Bu3iuwOWyT/dKEGMN5zSLMWdB7GTQ=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_swarm",
    importpath = "github.com/libp2p/go-libp2p-swarm",
    sum = "h1:HrFk2p0awrGEgch9JXK/qp/hfjqQfgNxpLWnCiWPg5s=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_testing",
    importpath = "github.com/libp2p/go-libp2p-testing",
    sum = "h1:bdij4bKaaND7tCsaXVjRfYkMpvoOeKj9AVQGJllA6jM=",
    version = "v0.0.3",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_tls",
    importpath = "github.com/libp2p/go-libp2p-tls",
    sum = "h1:twKMhMu44jQO+HgQK9X8NHO5HkeJu2QbhLzLJpa8oNM=",
    version = "v0.1.3",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_transport",
    importpath = "github.com/libp2p/go-libp2p-transport",
    sum = "h1:pV6+UlRxyDpASSGD+60vMvdifSCby6JkJDfi+yUMHac=",
    version = "v0.0.5",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_transport_upgrader",
    importpath = "github.com/libp2p/go-libp2p-transport-upgrader",
    sum = "h1:PZMS9lhjK9VytzMCW3tWHAXtKXmlURSc3ZdvwEcKCzw=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_xor",
    importpath = "github.com/libp2p/go-libp2p-xor",
    sum = "h1:GcTNu27BMpOTtMnQqun03+kbtHA1qTxJ/J8cZRRYu2k=",
    version = "v0.0.0-20200501025846-71e284145d58",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_yamux",
    importpath = "github.com/libp2p/go-libp2p-yamux",
    sum = "h1:TSPZ5cMMz/wdoYsye/wU1TE4G3LDGMoeEN0xgnCKU/I=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_libp2p_go_maddr_filter",
    importpath = "github.com/libp2p/go-maddr-filter",
    sum = "h1:hx8HIuuwk34KePddrp2mM5ivgPkZ09JH4AvsALRbFUs=",
    version = "v0.0.4",
)

go_repository(
    name = "com_github_libp2p_go_mplex",
    importpath = "github.com/libp2p/go-mplex",
    sum = "h1:/nBTy5+1yRyY82YaO6HXQRnO5IAGsXTjEJaR3LdTPc0=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_libp2p_go_msgio",
    importpath = "github.com/libp2p/go-msgio",
    sum = "h1:ivPvEKHxmVkTClHzg6RXTYHqaJQ0V9cDbq+6lKb3UV0=",
    version = "v0.0.2",
)

go_repository(
    name = "com_github_libp2p_go_nat",
    importpath = "github.com/libp2p/go-nat",
    sum = "h1:l6fKV+p0Xa354EqQOQP+d8CivdLM4kl5GxC1hSc/UeI=",
    version = "v0.0.3",
)

go_repository(
    name = "com_github_libp2p_go_netroute",
    importpath = "github.com/libp2p/go-netroute",
    sum = "h1:ruPJStbYyXVYGQ81uzEDzuvbYRLKRrLvTYd33yomC38=",
    version = "v0.1.6",
)

go_repository(
    name = "com_github_libp2p_go_openssl",
    importpath = "github.com/libp2p/go-openssl",
    sum = "h1:eCAzdLejcNVBzP/iZM9vqHnQm+XyCEbSSIheIPRGNsw=",
    version = "v0.0.7",
)

go_repository(
    name = "com_github_libp2p_go_reuseport",
    importpath = "github.com/libp2p/go-reuseport",
    sum = "h1:7PhkfH73VXfPJYKQ6JwS5I/eVcoyYi9IMNGc6FWpFLw=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_libp2p_go_reuseport_transport",
    importpath = "github.com/libp2p/go-reuseport-transport",
    sum = "h1:WglMwyXyBu61CMkjCCtnmqNqnjib0GIEjMiHTwR/KN4=",
    version = "v0.0.2",
)

go_repository(
    name = "com_github_libp2p_go_sockaddr",
    importpath = "github.com/libp2p/go-sockaddr",
    sum = "h1:yD80l2ZOdGksnOyHrhxDdTDFrf7Oy+v3FMVArIRgZxQ=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_libp2p_go_socket_activation",
    importpath = "github.com/libp2p/go-socket-activation",
    sum = "h1:VLU3IbrUUqu4DMhxA9857Q63qUpEAbCz5RqSnLCx5jE=",
    version = "v0.0.2",
)

go_repository(
    name = "com_github_libp2p_go_stream_muxer",
    importpath = "github.com/libp2p/go-stream-muxer",
    sum = "h1:Ce6e2Pyu+b5MC1k3eeFtAax0pW4gc6MosYSLV05UeLw=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_libp2p_go_stream_muxer_multistream",
    importpath = "github.com/libp2p/go-stream-muxer-multistream",
    sum = "h1:714bRJ4Zy9mdhyTLJ+ZKiROmAFwUHpeRidG+q7LTQOg=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_libp2p_go_tcp_transport",
    importpath = "github.com/libp2p/go-tcp-transport",
    sum = "h1:IGhowvEqyMFknOar4FWCKSWE0zL36UFKQtiRQD60/8o=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_libp2p_go_testutil",
    importpath = "github.com/libp2p/go-testutil",
    sum = "h1:4QhjaWGO89udplblLVpgGDOQjzFlRavZOjuEnz2rLMc=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_libp2p_go_ws_transport",
    importpath = "github.com/libp2p/go-ws-transport",
    sum = "h1:F+0OvvdmPTDsVc4AjPHjV7L7Pk1B7D5QwtDcKE2oag4=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_libp2p_go_yamux",
    importpath = "github.com/libp2p/go-yamux",
    sum = "h1:s6J6o7+ajoQMjHe7BEnq+EynOj5D2EoG8CuQgL3F2vg=",
    version = "v1.2.2",
)

go_repository(
    name = "com_github_libp2p_go_yamux_v2",
    importpath = "github.com/libp2p/go-yamux/v2",
    sum = "h1:RwtpYZ2/wVviZ5+3pjC8qdQ4TKnrak0/E01N1UWoAFU=",
    version = "v2.2.0",
)

go_repository(
    name = "com_github_lightstep_lightstep_tracer_common_golang_gogo",
    importpath = "github.com/lightstep/lightstep-tracer-common/golang/gogo",
    sum = "h1:143Bb8f8DuGWck/xpNUOckBVYfFbBTnLevfRZ1aVVqo=",
    version = "v0.0.0-20190605223551-bc2310a04743",
)

go_repository(
    name = "com_github_lightstep_lightstep_tracer_go",
    importpath = "github.com/lightstep/lightstep-tracer-go",
    sum = "h1:vi1F1IQ8N7hNWytK9DpJsUfQhGuNSc19z330K6vl4zk=",
    version = "v0.18.1",
)

go_repository(
    name = "com_github_lucas_clemente_quic_go",
    importpath = "github.com/lucas-clemente/quic-go",
    sum = "h1:8LqqL7nBQFDUINadW0fHV/xSaCQJgmJC0Gv+qUnjd78=",
    version = "v0.21.2",
)

go_repository(
    name = "com_github_lyft_protoc_gen_validate",
    importpath = "github.com/lyft/protoc-gen-validate",
    sum = "h1:KNt/RhmQTOLr7Aj8PsJ7mTronaFyx80mRTT9qF261dA=",
    version = "v0.0.13",
)

go_repository(
    name = "com_github_mailru_easyjson",
    importpath = "github.com/mailru/easyjson",
    sum = "h1:hB2xlXdHp/pmPZq0y3QnmWAArdw9PqbmotexnWx/FU8=",
    version = "v0.0.0-20190626092158-b2ccc519800e",
)

go_repository(
    name = "com_github_marten_seemann_qpack",
    importpath = "github.com/marten-seemann/qpack",
    sum = "h1:jvTsT/HpCn2UZJdP+UUB53FfUUgeOyG5K1ns0OJOGVs=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_marten_seemann_qtls",
    importpath = "github.com/marten-seemann/qtls",
    sum = "h1:ECsuYUKalRL240rRD4Ri33ISb7kAQ3qGDlrrl55b2pc=",
    version = "v0.10.0",
)

go_repository(
    name = "com_github_marten_seemann_qtls_go1_15",
    importpath = "github.com/marten-seemann/qtls-go1-15",
    sum = "h1:Ci4EIUN6Rlb+D6GmLdej/bCQ4nPYNtVXQB+xjiXE1nk=",
    version = "v0.1.5",
)

go_repository(
    name = "com_github_marten_seemann_qtls_go1_16",
    importpath = "github.com/marten-seemann/qtls-go1-16",
    sum = "h1:xbHbOGGhrenVtII6Co8akhLEdrawwB2iHl5yhJRpnco=",
    version = "v0.1.4",
)

go_repository(
    name = "com_github_marten_seemann_qtls_go1_17",
    importpath = "github.com/marten-seemann/qtls-go1-17",
    sum = "h1:/rpmWuGvceLwwWuaKPdjpR4JJEUH0tq64/I3hvzaNLM=",
    version = "v0.1.0-rc.1",
)

go_repository(
    name = "com_github_marten_seemann_tcp",
    importpath = "github.com/marten-seemann/tcp",
    sum = "h1:br0buuQ854V8u83wA0rVZ8ttrq5CpaPZdvrK0LP2lOk=",
    version = "v0.0.0-20210406111302-dfbc87cc63fd",
)

go_repository(
    name = "com_github_mgutz_ansi",
    importpath = "github.com/mgutz/ansi",
    sum = "h1:j7+1HpAFS1zy5+Q4qx1fWh90gTKwiN4QCGoY9TWyyO4=",
    version = "v0.0.0-20170206155736-9520e82c474b",
)

go_repository(
    name = "com_github_microcosm_cc_bluemonday",
    importpath = "github.com/microcosm-cc/bluemonday",
    sum = "h1:SIYunPjnlXcW+gVfvm0IlSeR5U3WZUOLfVmqg85Go44=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_mikioh_tcp",
    importpath = "github.com/mikioh/tcp",
    sum = "h1:bzE/A84HN25pxAuk9Eej1Kz9OUelF97nAc82bDquQI8=",
    version = "v0.0.0-20190314235350-803a9b46060c",
)

go_repository(
    name = "com_github_mikioh_tcpinfo",
    importpath = "github.com/mikioh/tcpinfo",
    sum = "h1:z78hV3sbSMAUoyUMM0I83AUIT6Hu17AWfgjzIbtrYFc=",
    version = "v0.0.0-20190314235526-30a79bb1804b",
)

go_repository(
    name = "com_github_mikioh_tcpopt",
    importpath = "github.com/mikioh/tcpopt",
    sum = "h1:PTfri+PuQmWDqERdnNMiD9ZejrlswWrCpBEZgWOiTrc=",
    version = "v0.0.0-20190314235656-172688c1accc",
)

go_repository(
    name = "com_github_minio_blake2b_simd",
    importpath = "github.com/minio/blake2b-simd",
    sum = "h1:lYpkrQH5ajf0OXOcUbGjvZxxijuBwbbmlSxLiuofa+g=",
    version = "v0.0.0-20160723061019-3f5f724cb5b1",
)

go_repository(
    name = "com_github_minio_sha256_simd",
    importpath = "github.com/minio/sha256-simd",
    sum = "h1:v1ta+49hkWZyvaKwrQB8elexRqm6Y0aMLjCNsrYxo6g=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_mr_tron_base58",
    importpath = "github.com/mr-tron/base58",
    sum = "h1:T/HDJBh4ZCPbU39/+c3rRvE0uKBQlU27+QI8LJ4t64o=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_multiformats_go_base32",
    importpath = "github.com/multiformats/go-base32",
    sum = "h1:tw5+NhuwaOjJCC5Pp82QuXbrmLzWg7uxlMFp8Nq/kkI=",
    version = "v0.0.3",
)

go_repository(
    name = "com_github_multiformats_go_base36",
    importpath = "github.com/multiformats/go-base36",
    sum = "h1:JR6TyF7JjGd3m6FbLU2cOxhC0Li8z8dLNGQ89tUg4F4=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_multiformats_go_multiaddr",
    importpath = "github.com/multiformats/go-multiaddr",
    sum = "h1:WgMSI84/eRLdbptXMkMWDXPjPq7SPLIgGUVm2eroyU4=",
    version = "v0.0.4",
)

go_repository(
    name = "com_github_multiformats_go_multiaddr_dns",
    importpath = "github.com/multiformats/go-multiaddr-dns",
    sum = "h1:/Bbsgsy3R6e3jf2qBahzNHzww6usYaZ0NhNH3sqdFS8=",
    version = "v0.0.2",
)

go_repository(
    name = "com_github_multiformats_go_multiaddr_fmt",
    importpath = "github.com/multiformats/go-multiaddr-fmt",
    sum = "h1:5YjeOIzbX8OTKVaN72aOzGIYW7PnrZrnkDyOfAWRSMA=",
    version = "v0.0.1",
)

# This is deprecated; replaced by "net" subpackage of:
#   github.com/multiformats/go-multiaddr
# We are waiting for its users to replace it, after which we can scrub
# this dependency.
go_repository(
    name = "com_github_multiformats_go_multiaddr_net",
    importpath = "github.com/multiformats/go-multiaddr-net",
    sum = "h1:76O59E3FavvHqNg7jvzWzsPSW5JSi/ek0E4eiDVbg9g=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_multiformats_go_multibase",
    importpath = "github.com/multiformats/go-multibase",
    sum = "h1:l/B6bJDQjvQ5G52jw4QGSYeOTZoAwIO77RblWplfIqk=",
    version = "v0.0.3",
)

go_repository(
    name = "com_github_multiformats_go_multicodec",
    importpath = "github.com/multiformats/go-multicodec",
    sum = "h1:tstDwfIjiHbnIjeM5Lp+pMrSeN+LCMsEwOrkPmWm03A=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_multiformats_go_multihash",
    importpath = "github.com/multiformats/go-multihash",
    sum = "h1:hWOPdrNqDjwHDx82vsYGSDZNyktOJJ2dzZJzFkOV1jM=",
    version = "v0.0.15",
)

go_repository(
    name = "com_github_multiformats_go_multistream",
    importpath = "github.com/multiformats/go-multistream",
    sum = "h1:UpO6jrsjqs46mqAK3n6wKRYFhugss9ArzbyUzU+4wkQ=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_multiformats_go_varint",
    importpath = "github.com/multiformats/go-varint",
    sum = "h1:gk85QWKxh3TazbLxED/NlDVv8+q+ReFJk7Y2W/KhfNY=",
    version = "v0.0.6",
)

go_repository(
    name = "com_github_nats_io_jwt",
    importpath = "github.com/nats-io/jwt",
    sum = "h1:+RB5hMpXUUA2dfxuhBTEkMOrYmM+gKIZYS1KjSostMI=",
    version = "v0.3.2",
)

go_repository(
    name = "com_github_nats_io_nats_go",
    importpath = "github.com/nats-io/nats.go",
    sum = "h1:ik3HbLhZ0YABLto7iX80pZLPw/6dx3T+++MZJwLnMrQ=",
    version = "v1.9.1",
)

go_repository(
    name = "com_github_nats_io_nats_server_v2",
    importpath = "github.com/nats-io/nats-server/v2",
    sum = "h1:i2Ly0B+1+rzNZHHWtD4ZwKi+OU5l+uQo1iDHZ2PmiIc=",
    version = "v2.1.2",
)

go_repository(
    name = "com_github_nats_io_nkeys",
    importpath = "github.com/nats-io/nkeys",
    sum = "h1:6JrEfig+HzTH85yxzhSVbjHRJv9cn0p6n3IngIcM5/k=",
    version = "v0.1.3",
)

go_repository(
    name = "com_github_nats_io_nuid",
    importpath = "github.com/nats-io/nuid",
    sum = "h1:5iA8DT8V7q8WK2EScv2padNa/rTESc1KdnPw4TC2paw=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_neelance_astrewrite",
    importpath = "github.com/neelance/astrewrite",
    sum = "h1:D6paGObi5Wud7xg83MaEFyjxQB1W5bz5d0IFppr+ymk=",
    version = "v0.0.0-20160511093645-99348263ae86",
)

go_repository(
    name = "com_github_neelance_sourcemap",
    importpath = "github.com/neelance/sourcemap",
    sum = "h1:eFXv9Nu1lGbrNbj619aWwZfVF5HBrm9Plte8aNptuTI=",
    version = "v0.0.0-20151028013722-8c68805598ab",
)

go_repository(
    name = "com_github_oklog_oklog",
    importpath = "github.com/oklog/oklog",
    sum = "h1:wVfs8F+in6nTBMkA7CbRw+zZMIB7nNM825cM1wuzoTk=",
    version = "v0.3.2",
)

go_repository(
    name = "com_github_oklog_run",
    importpath = "github.com/oklog/run",
    sum = "h1:Ru7dDtJNOyC66gQ5dQmaCa0qIsAUFY3sFpK1Xk8igrw=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_op_go_logging",
    importpath = "github.com/op/go-logging",
    sum = "h1:lDH9UUVJtmYCjyT0CI4q8xvlXPxeZ0gYCVvWbmPlp88=",
    version = "v0.0.0-20160315200505-970db520ece7",
)

go_repository(
    name = "com_github_opentracing_basictracer_go",
    importpath = "github.com/opentracing/basictracer-go",
    sum = "h1:YyUAhaEfjoWXclZVJ9sGoNct7j4TVk7lZWlQw5UXuoo=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_opentracing_contrib_go_observer",
    importpath = "github.com/opentracing-contrib/go-observer",
    sum = "h1:lM6RxxfUMrYL/f8bWEUqdXrANWtrL7Nndbm9iFN0DlU=",
    version = "v0.0.0-20170622124052-a52f23424492",
)

go_repository(
    name = "com_github_openzipkin_contrib_zipkin_go_opentracing",
    importpath = "github.com/openzipkin-contrib/zipkin-go-opentracing",
    sum = "h1:ZCnq+JUrvXcDVhX/xRolRBZifmabN1HcS1wrPSvxhrU=",
    version = "v0.4.5",
)

go_repository(
    name = "com_github_openzipkin_zipkin_go",
    importpath = "github.com/openzipkin/zipkin-go",
    sum = "h1:nY8Hti+WKaP0cRsSeQ026wU03QsM762XBeCXBb9NAWI=",
    version = "v0.2.2",
)

go_repository(
    name = "com_github_pact_foundation_pact_go",
    importpath = "github.com/pact-foundation/pact-go",
    sum = "h1:OYkFijGHoZAYbOIb1LWXrwKQbMMRUv1oQ89blD2Mh2Q=",
    version = "v1.0.4",
)

go_repository(
    name = "com_github_performancecopilot_speed",
    importpath = "github.com/performancecopilot/speed",
    sum = "h1:2WnRzIquHa5QxaJKShDkLM+sc0JPuwhXzK8OYOyt3Vg=",
    version = "v3.0.0+incompatible",
)

go_repository(
    name = "com_github_pierrec_lz4",
    importpath = "github.com/pierrec/lz4",
    sum = "h1:2xWsjqPFWcplujydGg4WmhC/6fZqK42wMM8aXeqhl0I=",
    version = "v2.0.5+incompatible",
)

go_repository(
    name = "com_github_pkg_profile",
    importpath = "github.com/pkg/profile",
    sum = "h1:F++O52m40owAmADcojzM+9gyjmMOY/T4oYJkgFDH8RE=",
    version = "v1.2.1",
)

go_repository(
    name = "com_github_polydawn_refmt",
    importpath = "github.com/polydawn/refmt",
    sum = "h1:ZOcivgkkFRnjfoTcGsDq3UQYiBmekwLA+qg0OjyB/ls=",
    version = "v0.0.0-20201211092308-30ac6d18308e",
)

go_repository(
    name = "com_github_prometheus_statsd_exporter",
    importpath = "github.com/prometheus/statsd_exporter",
    sum = "h1:M0hQphnq2WyWKS5CefQL8PqWwBOBPhiAkyLo5l4ZYvE=",
    version = "v0.20.0",
)

go_repository(
    name = "com_github_rcrowley_go_metrics",
    importpath = "github.com/rcrowley/go-metrics",
    sum = "h1:9ZKAASQSHhDYGoxY8uLVpewe1GDZ2vu2Tr/vTdVAkFQ=",
    version = "v0.0.0-20181016184325-3113b8401b8a",
)

go_repository(
    name = "com_github_rwcarlsen_goexif",
    importpath = "github.com/rwcarlsen/goexif",
    sum = "h1:CmH9+J6ZSsIjUK3dcGsnCnO41eRBOnY12zwkn5qVwgc=",
    version = "v0.0.0-20190401172101-9e8deecbddbd",
)

go_repository(
    name = "com_github_samuel_go_zookeeper",
    importpath = "github.com/samuel/go-zookeeper",
    sum = "h1:p3Vo3i64TCLY7gIfzeQaUJ+kppEO5WQG3cL8iE8tGHU=",
    version = "v0.0.0-20190923202752-2cc03de413da",
)

go_repository(
    name = "com_github_shopify_sarama",
    importpath = "github.com/Shopify/sarama",
    sum = "h1:9oksLxC6uxVPHPVYUmq6xhr1BOF/hHobWH2UzO67z1s=",
    version = "v1.19.0",
)

go_repository(
    name = "com_github_shopify_toxiproxy",
    importpath = "github.com/Shopify/toxiproxy",
    sum = "h1:TKdv8HiTLgE5wdJuEML90aBgNWsokNbMijUGhmcoBJc=",
    version = "v2.1.4+incompatible",
)

go_repository(
    name = "com_github_shurcool_component",
    importpath = "github.com/shurcooL/component",
    sum = "h1:Fth6mevc5rX7glNLpbAMJnqKlfIkcTjZCSHEeqvKbcI=",
    version = "v0.0.0-20170202220835-f88ec8f54cc4",
)

go_repository(
    name = "com_github_shurcool_events",
    importpath = "github.com/shurcooL/events",
    sum = "h1:vabduItPAIz9px5iryD5peyx7O3Ya8TBThapgXim98o=",
    version = "v0.0.0-20181021180414-410e4ca65f48",
)

go_repository(
    name = "com_github_shurcool_github_flavored_markdown",
    importpath = "github.com/shurcooL/github_flavored_markdown",
    sum = "h1:qb9IthCFBmROJ6YBS31BEMeSYjOscSiG+EO+JVNTz64=",
    version = "v0.0.0-20181002035957-2122de532470",
)

go_repository(
    name = "com_github_shurcool_gofontwoff",
    importpath = "github.com/shurcooL/gofontwoff",
    sum = "h1:Yoy/IzG4lULT6qZg62sVC+qyBL8DQkmD2zv6i7OImrc=",
    version = "v0.0.0-20180329035133-29b52fc0a18d",
)

go_repository(
    name = "com_github_shurcool_gopherjslib",
    importpath = "github.com/shurcooL/gopherjslib",
    sum = "h1:UOk+nlt1BJtTcH15CT7iNO7YVWTfTv/DNwEAQHLIaDQ=",
    version = "v0.0.0-20160914041154-feb6d3990c2c",
)

go_repository(
    name = "com_github_shurcool_highlight_diff",
    importpath = "github.com/shurcooL/highlight_diff",
    sum = "h1:vYEG87HxbU6dXj5npkeulCS96Dtz5xg3jcfCgpcvbIw=",
    version = "v0.0.0-20170515013008-09bb4053de1b",
)

go_repository(
    name = "com_github_shurcool_highlight_go",
    importpath = "github.com/shurcooL/highlight_go",
    sum = "h1:7pDq9pAMCQgRohFmd25X8hIH8VxmT3TaDm+r9LHxgBk=",
    version = "v0.0.0-20181028180052-98c3abbbae20",
)

go_repository(
    name = "com_github_shurcool_home",
    importpath = "github.com/shurcooL/home",
    sum = "h1:MPblCbqA5+z6XARjScMfz1TqtJC7TuTRj0U9VqIBs6k=",
    version = "v0.0.0-20181020052607-80b7ffcb30f9",
)

go_repository(
    name = "com_github_shurcool_htmlg",
    importpath = "github.com/shurcooL/htmlg",
    sum = "h1:crYRwvwjdVh1biHzzciFHe8DrZcYrVcZFlJtykhRctg=",
    version = "v0.0.0-20170918183704-d01228ac9e50",
)

go_repository(
    name = "com_github_shurcool_httperror",
    importpath = "github.com/shurcooL/httperror",
    sum = "h1:eHRtZoIi6n9Wo1uR+RU44C247msLWwyA89hVKwRLkMk=",
    version = "v0.0.0-20170206035902-86b7830d14cc",
)

go_repository(
    name = "com_github_shurcool_httpfs",
    importpath = "github.com/shurcooL/httpfs",
    sum = "h1:SWV2fHctRpRrp49VXJ6UZja7gU9QLHwRpIPBN89SKEo=",
    version = "v0.0.0-20171119174359-809beceb2371",
)

go_repository(
    name = "com_github_shurcool_httpgzip",
    importpath = "github.com/shurcooL/httpgzip",
    sum = "h1:fxoFD0in0/CBzXoyNhMTjvBZYW6ilSnTw7N7y/8vkmM=",
    version = "v0.0.0-20180522190206-b1c53ac65af9",
)

go_repository(
    name = "com_github_shurcool_issues",
    importpath = "github.com/shurcooL/issues",
    sum = "h1:T4wuULTrzCKMFlg3HmKHgXAF8oStFb/+lOIupLV2v+o=",
    version = "v0.0.0-20181008053335-6292fdc1e191",
)

go_repository(
    name = "com_github_shurcool_issuesapp",
    importpath = "github.com/shurcooL/issuesapp",
    sum = "h1:Y+TeIabU8sJD10Qwd/zMty2/LEaT9GNDaA6nyZf+jgo=",
    version = "v0.0.0-20180602232740-048589ce2241",
)

go_repository(
    name = "com_github_shurcool_notifications",
    importpath = "github.com/shurcooL/notifications",
    sum = "h1:TQVQrsyNaimGwF7bIhzoVC9QkKm4KsWd8cECGzFx8gI=",
    version = "v0.0.0-20181007000457-627ab5aea122",
)

go_repository(
    name = "com_github_shurcool_octicon",
    importpath = "github.com/shurcooL/octicon",
    sum = "h1:bu666BQci+y4S0tVRVjsHUeRon6vUXmsGBwdowgMrg4=",
    version = "v0.0.0-20181028054416-fa4f57f9efb2",
)

go_repository(
    name = "com_github_shurcool_reactions",
    importpath = "github.com/shurcooL/reactions",
    sum = "h1:LneqU9PHDsg/AkPDU3AkqMxnMYL+imaqkpflHu73us8=",
    version = "v0.0.0-20181006231557-f2e0b4ca5b82",
)

go_repository(
    name = "com_github_shurcool_users",
    importpath = "github.com/shurcooL/users",
    sum = "h1:YGaxtkYjb8mnTvtufv2LKLwCQu2/C7qFB7UtrOlTWOY=",
    version = "v0.0.0-20180125191416-49c67e49c537",
)

go_repository(
    name = "com_github_shurcool_webdavfs",
    importpath = "github.com/shurcooL/webdavfs",
    sum = "h1:JtcyT0rk/9PKOdnKQzuDR+FSjh7SGtJwpgVpfZBRKlQ=",
    version = "v0.0.0-20170829043945-18c3829fa133",
)

go_repository(
    name = "com_github_smola_gocompat",
    importpath = "github.com/smola/gocompat",
    sum = "h1:6b1oIMlUXIpz//VKEDzPVBK8KG7beVwmHIUEBIs/Pns=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_sony_gobreaker",
    importpath = "github.com/sony/gobreaker",
    sum = "h1:oMnRNZXX5j85zso6xCPRNPtmAycat+WcoKbklScLDgQ=",
    version = "v0.4.1",
)

go_repository(
    name = "com_github_sourcegraph_annotate",
    importpath = "github.com/sourcegraph/annotate",
    sum = "h1:yKm7XZV6j9Ev6lojP2XaIshpT4ymkqhMeSghO5Ps00E=",
    version = "v0.0.0-20160123013949-f4cad6c6324d",
)

go_repository(
    name = "com_github_sourcegraph_syntaxhighlight",
    importpath = "github.com/sourcegraph/syntaxhighlight",
    sum = "h1:qpG93cPwA5f7s/ZPBJnGOYQNK/vKsaDaseuKT5Asee8=",
    version = "v0.0.0-20170531221838-bd320f5d308e",
)

go_repository(
    name = "com_github_spacemonkeygo_openssl",
    importpath = "github.com/spacemonkeygo/openssl",
    sum = "h1:/eS3yfGjQKG+9kayBkj0ip1BGhq6zJ3eaVksphxAaek=",
    version = "v0.0.0-20181017203307-c2dcc5cca94a",
)

go_repository(
    name = "com_github_spacemonkeygo_spacelog",
    importpath = "github.com/spacemonkeygo/spacelog",
    sum = "h1:RC6RW7j+1+HkWaX/Yh71Ee5ZHaHYt7ZP4sQgUrm6cDU=",
    version = "v0.0.0-20180420211403-2296661a0572",
)

go_repository(
    name = "com_github_src_d_envconfig",
    importpath = "github.com/src-d/envconfig",
    sum = "h1:/AJi6DtjFhZKNx3OB2qMsq7y4yT5//AeSZIe7rk+PX8=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_stebalien_go_bitfield",
    importpath = "github.com/Stebalien/go-bitfield",
    sum = "h1:X3kbSSPUaJK60wV2hjOPZwmpljr6VGCqdq4cBLhbQBo=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_streadway_amqp",
    importpath = "github.com/streadway/amqp",
    sum = "h1:WhxRHzgeVGETMlmVfqhRn8RIeeNoPr2Czh33I4Zdccw=",
    version = "v0.0.0-20190827072141-edfb9018d271",
)

go_repository(
    name = "com_github_streadway_handy",
    importpath = "github.com/streadway/handy",
    sum = "h1:AhmOdSHeswKHBjhsLs/7+1voOxT+LLrSk/Nxvk35fug=",
    version = "v0.0.0-20190108123426-d5acb3125c2a",
)

go_repository(
    name = "com_github_syndtr_goleveldb",
    importpath = "github.com/syndtr/goleveldb",
    sum = "h1:fBdIW9lB4Iz0n9khmH8w27SJ3QEJ7+IgjPEwGSZiFdE=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_tarm_serial",
    importpath = "github.com/tarm/serial",
    sum = "h1:UyzmZLoiDWMRywV4DUYb9Fbt8uiOSooupjTq10vpvnU=",
    version = "v0.0.0-20180830185346-98f6abe2eb07",
)

go_repository(
    name = "com_github_texttheater_golang_levenshtein",
    importpath = "github.com/texttheater/golang-levenshtein",
    sum = "h1:T5PdfK/M1xyrHwynxMIVMWLS7f/qHwfslZphxtGnw7s=",
    version = "v0.0.0-20180516184445-d188e65d659e",
)

go_repository(
    name = "com_github_tv42_httpunix",
    importpath = "github.com/tv42/httpunix",
    sum = "h1:u6SKchux2yDvFQnDHS3lPnIRmfVJ5Sxy3ao2SIdysLQ=",
    version = "v0.0.0-20191220191345-2ba4b9c3382c",
)

go_repository(
    name = "com_github_urfave_cli_v2",
    importpath = "github.com/urfave/cli/v2",
    sum = "h1:+HU9SCbu8GnEUFtIBfuUNXN39ofWViIEJIp6SURMpCg=",
    version = "v2.0.0",
)

go_repository(
    name = "com_github_viant_assertly",
    importpath = "github.com/viant/assertly",
    sum = "h1:5x1GzBaRteIwTr5RAGFVG14uNeRFxVNbXPWrK2qAgpc=",
    version = "v0.4.8",
)

go_repository(
    name = "com_github_viant_toolbox",
    importpath = "github.com/viant/toolbox",
    sum = "h1:6TteTDQ68CjgcCe8wH3D3ZhUQQOJXMTbj/D9rkk2a1k=",
    version = "v0.24.0",
)

go_repository(
    name = "com_github_vividcortex_gohistogram",
    importpath = "github.com/VividCortex/gohistogram",
    sum = "h1:6+hBz+qvs0JOrrNhhmR7lFxo5sINxBCGXrdtl/UvroE=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_wangjia184_sortedset",
    importpath = "github.com/wangjia184/sortedset",
    sum = "h1:kZiWylALnUy4kzoKJemjH8eqwCl3RjW1r1ITCjjW7G8=",
    version = "v0.0.0-20160527075905-f5d03557ba30",
)

go_repository(
    name = "com_github_warpfork_go_wish",
    importpath = "github.com/warpfork/go-wish",
    sum = "h1:G++j5e0OC488te356JvdhaM8YS6nMsjLAYF7JxCv07w=",
    version = "v0.0.0-20200122115046-b9ea61034e4a",
)

go_repository(
    name = "com_github_whyrusleeping_base32",
    importpath = "github.com/whyrusleeping/base32",
    sum = "h1:BCPnHtcboadS0DvysUuJXZ4lWVv5Bh5i7+tbIyi+ck4=",
    version = "v0.0.0-20170828182744-c30ac30633cc",
)

go_repository(
    name = "com_github_whyrusleeping_cbor_gen",
    importpath = "github.com/whyrusleeping/cbor-gen",
    sum = "h1:WXhVOwj2USAXB5oMDwRl3piOux2XMV9TANaYxXHdkoE=",
    version = "v0.0.0-20200123233031-1cdf64d27158",
)

go_repository(
    name = "com_github_whyrusleeping_chunker",
    importpath = "github.com/whyrusleeping/chunker",
    sum = "h1:jQa4QT2UP9WYv2nzyawpKMOCl+Z/jW7djv2/J50lj9E=",
    version = "v0.0.0-20181014151217-fe64bd25879f",
)

go_repository(
    name = "com_github_whyrusleeping_go_keyspace",
    importpath = "github.com/whyrusleeping/go-keyspace",
    sum = "h1:EKhdznlJHPMoKr0XTrX+IlJs1LH3lyx2nfr1dOlZ79k=",
    version = "v0.0.0-20160322163242-5b898ac5add1",
)

go_repository(
    name = "com_github_whyrusleeping_go_logging",
    importpath = "github.com/whyrusleeping/go-logging",
    sum = "h1:9lDbC6Rz4bwmou+oE6Dt4Cb2BGMur5eR/GYptkKUVHo=",
    version = "v0.0.0-20170515211332-0457bb6b88fc",
)

go_repository(
    name = "com_github_whyrusleeping_go_notifier",
    importpath = "github.com/whyrusleeping/go-notifier",
    sum = "h1:M/lL30eFZTKnomXY6huvM6G0+gVquFNf6mxghaWlFUg=",
    version = "v0.0.0-20170827234753-097c5d47330f",
)

go_repository(
    name = "com_github_whyrusleeping_go_sysinfo",
    importpath = "github.com/whyrusleeping/go-sysinfo",
    sum = "h1:ctS9Anw/KozviCCtK6VWMz5kPL9nbQzbQY4yfqlIV4M=",
    version = "v0.0.0-20190219211824-4a357d4b90b1",
)

go_repository(
    name = "com_github_whyrusleeping_mafmt",
    importpath = "github.com/whyrusleeping/mafmt",
    sum = "h1:TCghSl5kkwEE0j+sU/gudyhVMRlpBin8fMBBHg59EbA=",
    version = "v1.2.8",
)

go_repository(
    name = "com_github_whyrusleeping_mdns",
    importpath = "github.com/whyrusleeping/mdns",
    sum = "h1:nMCC9Pwz1pxfC1Y6mYncdk+kq8d5aLx0Q+/gyZGE44M=",
    version = "v0.0.0-20180901202407-ef14215e6b30",
)

go_repository(
    name = "com_github_whyrusleeping_multiaddr_filter",
    importpath = "github.com/whyrusleeping/multiaddr-filter",
    sum = "h1:E9S12nwJwEOXe2d6gT6qxdvqMnNq+VnSsKPgm2ZZNds=",
    version = "v0.0.0-20160516205228-e903e4adabd7",
)

go_repository(
    name = "com_github_whyrusleeping_timecache",
    importpath = "github.com/whyrusleeping/timecache",
    sum = "h1:lYbXeSvJi5zk5GLKVuid9TVjS9a0OmLIDKTfoZBL6Ow=",
    version = "v0.0.0-20160911033111-cfcb2f1abfee",
)

go_repository(
    name = "com_github_x_cray_logrus_prefixed_formatter",
    importpath = "github.com/x-cray/logrus-prefixed-formatter",
    sum = "h1:00txxvfBM9muc0jiLIEAkAcIMJzfthRT6usrui8uGmg=",
    version = "v0.5.2",
)

go_repository(
    name = "com_shuralyov_dmitri_app_changes",
    importpath = "dmitri.shuralyov.com/app/changes",
    sum = "h1:hJiie5Bf3QucGRa4ymsAUOxyhYwGEz1xrsVk0P8erlw=",
    version = "v0.0.0-20180602232624-0a106ad413e3",
)

go_repository(
    name = "com_shuralyov_dmitri_html_belt",
    importpath = "dmitri.shuralyov.com/html/belt",
    sum = "h1:SPOUaucgtVls75mg+X7CXigS71EnsfVUK/2CgVrwqgw=",
    version = "v0.0.0-20180602232347-f7d459c86be0",
)

go_repository(
    name = "com_shuralyov_dmitri_service_change",
    importpath = "dmitri.shuralyov.com/service/change",
    sum = "h1:GvWw74lx5noHocd+f6HBMXK6DuggBB1dhVkuGZbv7qM=",
    version = "v0.0.0-20181023043359-a85b471d5412",
)

go_repository(
    name = "com_shuralyov_dmitri_state",
    importpath = "dmitri.shuralyov.com/state",
    sum = "h1:ivON6cwHK1OH26MZyWDCnbTRZZf0IhNsENoNAKFS1g4=",
    version = "v0.0.0-20180228185332-28bcc343414c",
)

go_repository(
    name = "com_sourcegraph_sourcegraph_appdash",
    importpath = "sourcegraph.com/sourcegraph/appdash",
    sum = "h1:ucqkfpjg9WzSUubAO62csmucvxl4/JeW3F4I4909XkM=",
    version = "v0.0.0-20190731080439-ebfcffb1b5c0",
)

go_repository(
    name = "com_sourcegraph_sourcegraph_go_diff",
    importpath = "sourcegraph.com/sourcegraph/go-diff",
    sum = "h1:eTiIR0CoWjGzJcnQ3OkhIl/b9GJovq4lSAVRt0ZFEG8=",
    version = "v0.5.0",
)

go_repository(
    name = "in_gopkg_inf_v0",
    importpath = "gopkg.in/inf.v0",
    sum = "h1:73M5CoZyi3ZLMOyDlQh031Cx6N9NDJ2Vvfl76EDAgDc=",
    version = "v0.9.1",
)

go_repository(
    name = "in_gopkg_src_d_go_cli_v0",
    importpath = "gopkg.in/src-d/go-cli.v0",
    sum = "h1:mXa4inJUuWOoA4uEROxtJ3VMELMlVkIxIfcR0HBekAM=",
    version = "v0.0.0-20181105080154-d492247bbc0d",
)

go_repository(
    name = "in_gopkg_src_d_go_log_v1",
    importpath = "gopkg.in/src-d/go-log.v1",
    sum = "h1:heWvX7J6qbGWbeFS/aRmiy1eYaT+QMV6wNvHDyMjQV4=",
    version = "v1.0.1",
)

go_repository(
    name = "io_opencensus_go_contrib_exporter_prometheus",
    importpath = "contrib.go.opencensus.io/exporter/prometheus",
    sum = "h1:08FMdJYpItzsknogU6PiiNo7XQZg/25GjH236+YCwD0=",
    version = "v0.3.0",
)

go_repository(
    name = "org_apache_git_thrift_git",
    importpath = "git.apache.org/thrift.git",
    sum = "h1:OR8VhtwhcAI3U48/rzBsVOuHi0zDPzYI1xASVcdSgR8=",
    version = "v0.0.0-20180902110319-2566ecd5d999",
)

go_repository(
    name = "org_bazil_fuse",
    importpath = "bazil.org/fuse",
    sum = "h1:utDghgcjE8u+EBjHOgYT+dJPcnDF05KqWMBcjuJy510=",
    version = "v0.0.0-20200117225306-7b5117fecadc",
)

go_repository(
    name = "org_go4",
    importpath = "go4.org",
    sum = "h1:BNJlw5kRTzdmyfh5U8F93HA2OwkP7ZGwA51eJ/0wKOU=",
    version = "v0.0.0-20200411211856-f5505b9728dd",
)

go_repository(
    name = "org_go4_grpc",
    importpath = "grpc.go4.org",
    sum = "h1:tmXTu+dfa+d9Evp8NpJdgOy6+rt8/x4yG7qPBrtNfLY=",
    version = "v0.0.0-20170609214715-11d0a25b4919",
)

go_repository(
    name = "org_golang_x_build",
    importpath = "golang.org/x/build",
    sum = "h1:E2M5QgjZ/Jg+ObCQAudsXxuTsLj7Nl5RV/lZcQZmKSo=",
    version = "v0.0.0-20190111050920-041ab4dc3f9d",
)

go_repository(
    name = "org_golang_x_perf",
    importpath = "golang.org/x/perf",
    sum = "h1:xYq6+9AtI+xP3M4r0N1hCkHrInHDBohhquRgx9Kk6gI=",
    version = "v0.0.0-20180704124530-6e6d33e29852",
)

go_repository(
    name = "org_uber_go_dig",
    importpath = "go.uber.org/dig",
    sum = "h1:yLmDDj9/zuDjv3gz8GQGviXMs9TfysIUMUilCpgzUJY=",
    version = "v1.10.0",
)

go_repository(
    name = "org_uber_go_fx",
    importpath = "go.uber.org/fx",
    sum = "h1:CFNTr1oin5OJ0VCZ8EycL3wzF29Jz2g0xe55RFsf2a4=",
    version = "v1.13.1",
)

go_repository(
    name = "org_uber_go_goleak",
    importpath = "go.uber.org/goleak",
    sum = "h1:z+mqJhf6ss6BSfSM671tgKyZBFPTTJM+HLxnhPC3wu0=",
    version = "v1.1.10",
)

go_repository(
    name = "com_github_gin_contrib_sse",
    importpath = "github.com/gin-contrib/sse",
    sum = "h1:Y/yl/+YNO8GZSjAhjMsSuLt29uWRFHdHYUb5lYOV9qE=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_gin_gonic_gin",
    importpath = "github.com/gin-gonic/gin",
    sum = "h1:aMBzLJ/GMEYmv1UWs2FFTcPISLrQH2mRgL9Glz8xows=",
    version = "v1.7.3",
)

go_repository(
    name = "com_github_go_playground_assert_v2",
    importpath = "github.com/go-playground/assert/v2",
    sum = "h1:MsBgLAaY856+nPRTKrp3/OZK38U/wa0CcBYNjji3q3A=",
    version = "v2.0.1",
)

go_repository(
    name = "com_github_go_playground_locales",
    importpath = "github.com/go-playground/locales",
    sum = "h1:HyWk6mgj5qFqCT5fjGBuRArbVDfE4hi8+e8ceBS/t7Q=",
    version = "v0.13.0",
)

go_repository(
    name = "com_github_go_playground_universal_translator",
    importpath = "github.com/go-playground/universal-translator",
    sum = "h1:icxd5fm+REJzpZx7ZfpaD876Lmtgy7VtROAbHHXk8no=",
    version = "v0.17.0",
)

go_repository(
    name = "com_github_go_playground_validator_v10",
    importpath = "github.com/go-playground/validator/v10",
    sum = "h1:pH2c5ADXtd66mxoE0Zm9SUhxE20r7aM3F26W0hOn+GE=",
    version = "v10.4.1",
)

go_repository(
    name = "com_github_leodido_go_urn",
    importpath = "github.com/leodido/go-urn",
    sum = "h1:hpXL4XnriNwQ/ABnpepYM/1vCLWNDfUNts8dX3xTG6Y=",
    version = "v1.2.0",
)

# ERROR: /home/adrian/.cache/bazel/_bazel_adrian/1645628646ea614ae25fb34fd432471c/external/com_github_ipfs_go_merkledag/pb/BUILD.bazel:12:17: no such package '@com_github_ipfs_go_merkledag//github.com/gogo/protobuf/gogoproto': BUILD file not found in directory 'github.com/gogo/protobuf/gogoproto' of external repository @com_github_ipfs_go_merkledag. Add a BUILD file to a directory to mark it as a package. and referenced by '@com_github_ipfs_go_merkledag//pb:merkledag_pb_go_proto'
# build_directives fixes this
# https://stackoverflow.com/questions/68776492/bazel-not-adding-build-file-to-external-dependency/68822701#68822701
go_repository(
    name = "com_github_ipfs_go_merkledag",
    build_directives = [
        "gazelle:proto disable",
        "gazelle:proto disable_global",
    ],
    importpath = "github.com/ipfs/go-merkledag",
    sum = "h1:MRqj40QkrWkvPswXs4EfSslhZ4RVPRbxwX11js0t1xY=",
    version = "v0.3.2",
)

go_repository(
    name = "com_github_r3labs_diff_v2",
    importpath = "github.com/r3labs/diff/v2",
    sum = "h1:OwxATYI673fTkqzQFcf154Lck9DZNgJsSdYKEg+3tC0=",
    version = "v2.13.6",
)

go_repository(
    name = "com_github_vmihailenco_msgpack",
    importpath = "github.com/vmihailenco/msgpack",
    sum = "h1:dSLoQfGFAo3F6OoNhwUmLwVgaUXK79GlxNBwueZn0xI=",
    version = "v4.0.4+incompatible",
)
