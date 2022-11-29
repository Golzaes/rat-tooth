package aliyun

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/aliyun/aliyun-log-go-sdk/producer"
	log "github.com/golzaes/rat-tooth/observation/logs"
	"google.golang.org/protobuf/proto"
)

// Logger see more detail https://github.com/aliyun/aliyun-log-go-sdk
type Logger interface {
	log.Logger
	Producer() *producer.Producer
	Close() error
}

type aliyunLog struct {
	producer *producer.Producer
	opts     *Sls
}

type SlsOption func(*Sls)

type Sls struct {
	accessKey    string
	accessSecret string
	endpoint     string
	projectName  string
	logStorage   string
}

func (a *aliyunLog) Producer() *producer.Producer {
	return a.producer
}

func (a *aliyunLog) Close() error {
	return a.producer.Close(5000)
}

func (a *aliyunLog) Log(level log.Level, kv ...interface{}) error {
	contents := make([]*sls.LogContent, 0, len(kv)/2+1)
	contents = append(contents, &sls.LogContent{
		Key:   newString(level.Key()),
		Value: newString(level.String()),
	})
	for i := 0; i < len(kv); i += 2 {
		contents = append(contents, &sls.LogContent{
			Key:   newString(toString(kv[i])),
			Value: newString(toString(kv[i+1])),
		})
	}
	logInst := &sls.Log{
		Time:     proto.Uint32(uint32(time.Now().Unix())),
		Contents: contents,
	}
	return a.producer.SendLog(a.opts.projectName, a.opts.logStorage, "", "", logInst)
}

// NewAliyunLog creates a new aliyun logger instance with default configuration
func NewAliyunLog(options ...SlsOption) Logger {
	s := &Sls{}
	s.Init()
	for _, f := range options {
		f(s)
	}
	producerConfig := producer.GetDefaultProducerConfig()
	producerConfig.Endpoint = s.endpoint
	producerConfig.AccessKeyID = s.accessKey
	producerConfig.AccessKeySecret = s.accessSecret
	producerInst := producer.InitProducer(producerConfig)
	return &aliyunLog{
		opts:     s,
		producer: producerInst,
	}
}

// Init initializes the Sls private variables and sets default
// configuration for the Sls
func (s *Sls) Init() {
	s.projectName = `project_name`
	s.logStorage = `app`
}

//WithEndpoint allows you to set a custom endpoint
func WithEndpoint(endpoint string) SlsOption {
	return func(alc *Sls) {
		alc.endpoint = endpoint
	}
}

//WithProject allows you to set a custom Project name
func WithProject(projectName string) SlsOption {
	return func(alc *Sls) {
		alc.projectName = projectName
	}
}

//WithLogStore allows you to set a custom log store
func WithLogStore(logStore string) SlsOption {
	return func(alc *Sls) {
		alc.logStorage = logStore
	}
}

// WithAccessKey allows you to set a custom AccessKey
func WithAccessKey(ak string) SlsOption {
	return func(alc *Sls) {
		alc.accessKey = ak
	}
}

// WithAccessSecret allows you to set a custom AccessSecret
func WithAccessSecret(as string) SlsOption {
	return func(alc *Sls) {
		alc.accessSecret = as
	}
}

// newString string convert to *string
func newString(s string) *string {
	return &s
}

// toString convert any type to string
func toString(v interface{}) string {
	var key string
	if v == nil {
		return key
	}
	switch v := v.(type) {
	case float64:
		key = strconv.FormatFloat(v, 'f', -1, 64)
	case float32:
		key = strconv.FormatFloat(float64(v), 'f', -1, 32)
	case int:
		key = strconv.Itoa(v)
	case uint:
		key = strconv.FormatUint(uint64(v), 10)
	case int8:
		key = strconv.Itoa(int(v))
	case uint8:
		key = strconv.FormatUint(uint64(v), 10)
	case int16:
		key = strconv.Itoa(int(v))
	case uint16:
		key = strconv.FormatUint(uint64(v), 10)
	case int32:
		key = strconv.Itoa(int(v))
	case uint32:
		key = strconv.FormatUint(uint64(v), 10)
	case int64:
		key = strconv.FormatInt(v, 10)
	case uint64:
		key = strconv.FormatUint(v, 10)
	case string:
		key = v
	case bool:
		key = strconv.FormatBool(v)
	case []byte:
		key = string(v)
	case fmt.Stringer:
		key = v.String()
	default:
		newValue, _ := json.Marshal(v)
		key = string(newValue)
	}
	return key
}
