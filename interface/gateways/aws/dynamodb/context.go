package dynamodb

import "context"

type txWriteContextKey struct{}

func GetTxWriteByContext(ctx context.Context) *txWrite {
	txWrite, ok := ctx.Value(txWriteContextKey{}).(*txWrite)
	if !ok {
		return nil
	}
	return txWrite
}

func SetTxWriteByContext(ctx context.Context, txWrite *txWrite) context.Context {
	return context.WithValue(ctx, txWriteContextKey{}, txWrite)
}
