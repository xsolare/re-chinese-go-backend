package response

type Response struct {
	RequestId string `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Code      int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Msg       string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	Status    string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

type response struct {
	Response
	Data interface{} `json:"data"`
}

func (e *response) SetData(data interface{}) {
	e.Data = data
}

func (e *response) SetTraceID(id string) {
	e.RequestId = id
}

func (e *response) SetMsg(s string) {
	e.Msg = s
}

func (e *response) SetCode(code int32) {
	e.Code = code
}

func (e *response) SetSuccess(success bool) {
	if !success {
		e.Status = "error"
	}
}

func (e response) Clone() Responses {
	return &e
}
