package api

import "hal9000/internal/app/api/controller"


type ApiApplication struct {
	testCtl *controller.TestController
}


func NewApiApplication(testctl *controller.TestController) *ApiApplication{
	return &ApiApplication{testCtl:testctl}
}


