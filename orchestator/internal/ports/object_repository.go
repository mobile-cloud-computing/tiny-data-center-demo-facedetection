package ports

type ObjectRepository interface {
	CreateBucketIfNotExist(bucket string) error
}
