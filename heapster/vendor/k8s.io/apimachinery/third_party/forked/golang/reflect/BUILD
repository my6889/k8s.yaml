package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "go_default_test",
    srcs = ["deep_equal_test.go"],
    importpath = "k8s.io/apimachinery/third_party/forked/golang/reflect",
    library = ":go_default_library",
)

go_library(
    name = "go_default_library",
    srcs = ["deep_equal.go"],
    importpath = "k8s.io/apimachinery/third_party/forked/golang/reflect",
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
)
