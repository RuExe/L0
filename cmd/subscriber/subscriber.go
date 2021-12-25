package subscriber

import (
	"L0/internal/config"
	"L0/internal/domain"
	"L0/internal/repository"
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

type Subscriber struct {
	config  *config.Config
	storage *repository.OrderStorage
}

func NewSubscriber(config *config.Config, storage *repository.OrderStorage) *Subscriber {
	return &Subscriber{
		config:  config,
		storage: storage,
	}
}

func (s *Subscriber) Subscribe() {
	opts := []nats.Option{nats.Name("NATS Sample Subscriber")}
	opts = setupConnOptions(opts)

	nc, err := nats.Connect(s.config.SubscriberConfig.Url, opts...)
	if err != nil {
		log.Println(err)
		return
	}

	subj, i := "foo", 0

	nc.Subscribe(subj, func(msg *nats.Msg) {
		i += 1
		s.action(msg, i)
	})
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on [%s]", subj)
}

func (s *Subscriber) action(m *nats.Msg, i int) {
	order := domain.Order{}
	err := json.Unmarshal(m.Data, &order)
	if err != nil {
		log.Println(err)
		return
	}

	(*s.storage).Add(order)
	log.Printf("[#%d] Received on [%s]: '%s'", i, m.Subject, string(m.Data))
}

func setupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		log.Printf("Disconnected due to:%s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Fatalf("Exiting: %v", nc.LastError())
	}))
	return opts
}
