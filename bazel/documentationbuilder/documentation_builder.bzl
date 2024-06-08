
def _build_documentation_impl(ctx):
    out = ctx.actions.declare_file(ctx.label.name)

    args = ["-out-dir=" + out.path, "-src=" + ctx.file.src.path]

    ctx.actions.run(
        outputs = [out],
        arguments = args,
        progress_message = "Merging into",
        executable = ctx.executable.tool,
    )

    return [DefaultInfo(files = depset([out]))]


build_documentations = rule(
    implementation = _build_documentation_impl,
    attrs = {
        "src": attr.label(allow_single_file = True),
        "tool": attr.label(
              executable = True,
              cfg = "exec",
              default = Label("//bazel/documentationbuilder:documentationbuilder"),
          )
    }
)
