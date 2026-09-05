package invalid
func Cast[To,From any](x From)To{return To(x)}
