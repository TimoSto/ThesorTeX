package feedback

type Feedback struct {
	ID      string `dynamodbav:"id"`
	Date    string `dynamodbav:"date"`
	Message string `dynamodbav:"message"`
}
