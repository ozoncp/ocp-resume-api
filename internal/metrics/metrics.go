package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	actions  *prometheus.CounterVec
	requests *prometheus.CounterVec
)

const (
	label  string = "action"
	create string = "create"
	update string = "update"
	remove string = "remove"
	descr  string = "describe"
)

// Register создает и регистрирует необходимые счетчики
func Register() {
	requests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "resume_requests",
			Help: "Number of requests for resume CRUD actions.",
		},
		[]string{label},
	)
	actions = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "resume_successful_actions",
			Help: "Number of resume CRUD actions successfully performed.",
		},
		[]string{label},
	)

	prometheus.MustRegister(requests, actions)
}

func incrementRequests(req string, times int) {
	if requests == nil {
		return
	}
	requests.With(prometheus.Labels{label: req}).Add(float64(times))
}

func incrementActions(act string, times int) {
	if actions == nil {
		return
	}
	actions.With(prometheus.Labels{label: act}).Add(float64(times))
}

// IncrementCreateRequests увеливает счетчик запросов на создание Resume
func IncrementCreateRequests(times int) {
	incrementRequests(create, times)
}

// IncrementUpdateRequests увеливает счетчик запросов на обновление Resume
func IncrementUpdateRequests(times int) {
	incrementRequests(update, times)
}

// IncrementRemoveRequests увеливает счетчик запросов на удаление Resume
func IncrementRemoveRequests(times int) {
	incrementRequests(remove, times)
}

// IncrementDescribeRequests увеливает счетчик запросов на получение Resume
func IncrementDescribeRequests(times int) {
	incrementRequests(descr, times)
}

// IncrementDescribeRequests увеливает счетчик запросов на получение Resume
func IncrementDescribe(times int) {
	incrementActions(descr, times)
}

// IncrementCreates увеливает счетчик успешных созданий Resume
func IncrementCreates(times int) {
	incrementActions(create, times)
}

// IncrementUpdates увеливает счетчик успешных обновлений Resume
func IncrementUpdates(times int) {
	incrementActions(update, times)
}

// IncrementRemoves увеливает счетчик успешных удалений Resume
func IncrementRemoves(times int) {
	incrementActions(remove, times)
}
