package consumer

import "encoding/json"

type consumerService struct {
	consumerStorage consumerStorage
}

func NewConsumerService(consumerStorage consumerStorage) *consumerService {
	return &consumerService{
		consumerStorage: consumerStorage,
	}
}

func (s *consumerService) KafKaUserTopic(msg []byte) error {

	var user User
	if err := json.Unmarshal(msg, &user); err != nil {
		return err
	}
	if err := s.consumerStorage.Insert(user); err != nil {
		return err
	}
	return nil
}
