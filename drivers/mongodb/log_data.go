package mongodb

type LogData struct {
	Request string `bson:"request"`
	Method  string `bson:"method"`
}
