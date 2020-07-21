package barion

import (
	"context"
	"log"
	"net/http"
)

type callbackRequest struct {
	PaymentId string
}

func (c *client) CallbackHandler(paymentStateHandler func(state *PaymentState)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		keys, ok := r.URL.Query()["paymentId"]

		if !ok || len(keys[0]) < 1 {
			log.Println("Url Param 'PaymentId' is missing")
			http.Error(w, "Url Param 'PaymentId' is missing", http.StatusBadRequest)
			return
		}

		key := keys[0]

		state, err := c.GetPaymentState(context.TODO(), key)
		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		go func() {
			paymentStateHandler(state)
		}()
		_, _ = w.Write([]byte("OK"))
		return
	})
}
