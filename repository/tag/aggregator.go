package tag

import "bareksa-take-home-test-michael-koh/core/entity"

func (r *repository) DeleteTag(tagName string) error {
	return r.DeleteTagRepo(tagName)
}

func (r *repository) GetTags(page, size int) ([]entity.Tag, error) {
	tags, err := r.GetTagsRepo(page, size)
	if err != nil {
		return nil, err
	}

	resTag := make([]entity.Tag, 0)
	for _, i := range tags {
		resTag = append(resTag, entity.Tag{
			Name: i.Name,
		})
	}

	return resTag, nil
}

func (r *repository) CreateTag(tagName string) error {
	return r.CreateTagRepo(tagName)
}

func (r *repository) UpdateTag(tagName, newTagName string) error {
	return r.UpdateTagRepo(tagName, newTagName)
}
