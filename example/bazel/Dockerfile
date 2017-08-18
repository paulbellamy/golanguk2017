FROM insready/bazel

# Extract and install bazel so it's ready to go
RUN bazel version

# Don't make me stutter
ENTRYPOINT ["bazel", "--output_base=.cache"]
