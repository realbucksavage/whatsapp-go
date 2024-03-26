package whatsapp

type MessageResponse struct {
	MessagingProduct string            `json:"messaging_product,omitempty"`
	Contacts         []ResponseContact `json:"contacts,omitempty"`
	Error            *APIError         `json:"error,omitempty"`
	Messages         []struct {
		ID string `json:"id" json:"id,omitempty"`
	} `json:"messages,omitempty"`
}

type ResponseContact struct {
	Input string `json:"input,omitempty"`
	WAID  string `json:"wa_id,omitempty"`
}

type APIError struct {
	Code         int
	ErrorSubcode int
	FBTraceID    string
	Message      string
	Type         string
}

func (err *APIError) Error() string {
	return err.Message
}

type MessageOption func(m *Message)

func WithOpaqueData(data string) MessageOption {
	return func(m *Message) {
		m.BizOpaqueCallbackData = data
	}
}

func apply(m *Message, opts ...MessageOption) {
	for _, o := range opts {
		o(m)
	}
}

func NewText(text string) *Text {
	return &Text{Body: text}
}

func NewTextMessage(to string, text *Text, opts ...MessageOption) *Message {
	msg := newMessage(to)
	msg.Type = TypeText
	msg.Text = text

	if opts != nil {
		apply(msg, opts...)
	}

	return msg
}

func newMessage(to string) *Message {
	return &Message{
		MessagingProduct: "whatsapp",
		To:               to,
	}
}
