load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    visibility = ["//visibility:private"],
    deps = ["//vendor/github.com/sirupsen/logrus:go_default_library"],
)

go_binary(
    name = "api",
    library = ":go_default_library",
    gc_linkopts = [
      "-linkmode",
      "external",
      "-extldflags",
      "-static",
    ],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["main_test.go"],
    library = ":go_default_library",
)
