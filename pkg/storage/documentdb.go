package storage

import "context"


type IDocumentDB interface {
	Close(ctx context.Context)  error
}
