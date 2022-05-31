package tag

func (r *repository) DeleteTag(tagName string) error {
	return r.DeleteTagRepo(tagName)
}
