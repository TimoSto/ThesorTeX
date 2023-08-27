package sendfeedback

import (
	"encoding/json"
	"fmt"
	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/services/contact/internal/feedback"
	"io"
	"net/http"
)

type message struct {
	Subject string
	Message string
}

func GetFeedbackHandler(store feedback.Store) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var data message
		b, err := io.ReadAll(r.Body)
		fmt.Println(b, err)
		err = json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			log.Error("could not unmarshal feedback message: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = store.SaveFeedback(data.Subject, data.Message)
		if err != nil {
			fmt.Println(err)
			log.Error("could not save the feedback: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	return http.HandlerFunc(fn)
}
