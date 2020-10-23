package service

import "go-gin-pj/model"

// WebTestService - test service
type WebTestService struct{}

// GetList - get test
func (WebTestService) GetList() model.WebTest {

	// test code
	test := model.WebTest{ID: 1, Content: "test"}
	// TODO : DbEngine
	// if err != nil {
	// 	panic(err)
	// }
	return test
}
