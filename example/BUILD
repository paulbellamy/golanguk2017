# go

load("@io_bazel_rules_go//go:def.bzl", "gazelle", "go_prefix")

go_prefix("github.com/paulbellamy/golanguk2017/example")

gazelle(
  name = "gazelle",
  external = "vendored",
)



# docker

load("@io_bazel_rules_docker//docker:docker.bzl", "docker_build")

docker_build(
  name = "lasersteak",
  base = "@alpine//image",
  cmd = ["/api"],
  files = ["//cmd/api"],
  tars = ["//js:build"],
  ports = ["8000"],
)
