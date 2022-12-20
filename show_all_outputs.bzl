def format(target):
    """
    This is used by `bazel output-query` command together with cquery.
    """
    outputs = target.files.to_list()
    return outputs[-1].path if len(outputs) > 0 else "(missing)"
