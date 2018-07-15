package testService

type TestService struct {
	Key    string
	Secret string
}

type TestServiceData struct {
	Value string
}

type TestServiceResp struct {
	RespVal string
}

func (t *TestService) Perform(val TestServiceData) (*TestServiceResp, error) {
	return &TestServiceResp{RespVal: val.Value}, nil
}

func (t *TestService) HandleResponse(resp *TestServiceResp) error {
	return nil
}
