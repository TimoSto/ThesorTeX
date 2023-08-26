package feedback

type Feedback struct {
	ID      string `dynamodbav:"id"`
	Date    string `dynamodbav:"date"`
	Subject string `dynamodbav:"subject"`
	Message string `dynamodbav:"message"`
}
