
def _build_documentation_impl(ctx):
    out = ctx.actions.declare_file(ctx.label.name)

    ctx.actions.run(
        outputs = [out],
        arguments = [out.path],
        progress_message = "Merging into",
        executable = ctx.executable.tool,
    )

    return [DefaultInfo(files = depset([out]))]


build_documentations = rule(
    implementation = _build_documentation_impl,
    attrs = {
        "tool": attr.label(
              executable = True,
              cfg = "exec",
              default = Label("//bazel/documentationbuilder:documentationbuilder"),
          )
    }
)
