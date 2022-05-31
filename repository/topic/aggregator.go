package topic

import "bareksa-take-home-test-michael-koh/core/entity"

func (r *repository) GetTopicsBySerials(serials []string) ([]entity.Topic, error) {
	topics, err := r.GetTopicsBySerialsRepo(serials)
	if err != nil {
		return nil, err
	}

	res := make([]entity.Topic, 0)
	for _, i := range topics {
		res = append(res, entity.Topic{
			Serial: i.Serial,
			Title:  i.Title,
		})
	}

	return res, nil
}
