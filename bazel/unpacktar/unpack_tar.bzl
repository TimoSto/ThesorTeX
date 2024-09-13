def _unpack_tar_impl(ctx):
    # tell bazel that this rule creates a directory with the given name
    out_dir = ctx.actions.declare_directory(ctx.attr.out_dir_path)

    # combine the args
    args = [out_dir.path] + [ctx.file.tar.path] + [ctx.attr.strip_components]

    ctx.actions.run(
        inputs = [ctx.file.tar],  # all files to which the action needs access
        outputs = [out_dir],
        arguments = args,
        executable = ctx.executable.tool,
    )

    # This is necessary, otherwise the output will not be available to other targets and won't be in the bazel/bin dir
    return [DefaultInfo(files = depset([out_dir]))]

unpack_tar = rule(
    implementation = _unpack_tar_impl,
    attrs = {
        "tar": attr.label(allow_single_file = True),
        "out_dir_path": attr.string(mandatory = True),  # this cannot be named "out", because gazelle would then add an incorrect embedsrc
        "tool": attr.label(
            executable = True,
            cfg = "exec",
            allow_files = True,
            default = Label("//bazel/unpacktar:unpack_tar"),
        ),
        "strip_components": attr.string(default = "0"),
    },
)
