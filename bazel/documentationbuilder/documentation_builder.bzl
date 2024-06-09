# This rule call the documentationbuilder tool and gives the output as a directory
# We need to do this in a rule, because newer version of pkg_tar require declare_directory, if we input directories

def _build_documentation_impl(ctx):

    # declare the out directory from the attribute
    # this gives us an object with the full path to the desired output in the bin directory

    out = ctx.actions.declare_directory(ctx.attr.out)

    # cli args for the go tool

    args = ["-out-dir=" + out.path, "-src=" + ctx.file.src.path]

    ctx.actions.run(
        inputs = [ctx.file.src],
        outputs = [out],
        arguments = args,
        progress_message = "Building documentation...",
        executable = ctx.executable.tool,
    )

    return [DefaultInfo(files = depset([out]))]


build_documentations = rule(
    implementation = _build_documentation_impl,
    attrs = {
        "src": attr.label(allow_single_file = True),
        "out": attr.string(default = "docs"),
        "tool": attr.label(
              executable = True,
              cfg = "exec",
              default = Label("//tools/documentationbuilder:documentationbuilder"),
          )
    }
)
