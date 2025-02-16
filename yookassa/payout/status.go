package yoopayout

type Status string

const (
	Pending   Status = "pending"
	Succeeded Status = "succeeded"
	Canceled  Status = "canceled"
)
