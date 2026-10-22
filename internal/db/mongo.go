package db

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)


const defaultMongoConnUrl = "mongodb://localhost:27017"

func mongoConnUrl() string {
	if url := os.Getenv("MONGO_URI"); url != "" {
		return url
	}
	return defaultMongoConnUrl
}

var instance *MongoDBConnection
var once sync.Once

type MongoDBConnection struct {
	client *mongo.Client
	dbName string
	mu     sync.RWMutex
}

func (conn *MongoDBConnection) Connect(ctx context.Context, dbName string) error {
	conn.mu.Lock()
	defer conn.mu.Unlock()
	
	options := options.Client().ApplyURI(mongoConnUrl()).SetMaxPoolSize(10).SetMinPoolSize(2).SetConnectTimeout(10 * time.Second).SetMaxConnIdleTime(5 * time.Minute).SetCompressors([]string{"zlib", "snappy"})

	client, err := mongo.Connect(options)
	if err != nil {
		return fmt.Errorf("mongo.Connect: %w", err)
	}

	pingCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := client.Ping(pingCtx, readpref.Primary()); err != nil {
		client.Disconnect(ctx)
		return fmt.Errorf("ping failed: %v", err)
	}

	conn.dbName = dbName
	conn.client = client
	return nil
}

func (conn *MongoDBConnection) Disconnect(ctx context.Context) error {
	conn.mu.RLock()
	defer conn.mu.RUnlock()
	if conn.client == nil {
		return nil
	}
	return conn.client.Disconnect(ctx)
}

func (conn *MongoDBConnection) Ping(ctx context.Context) error {
	conn.mu.RLock()
	defer conn.mu.RUnlock()
	pingctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := conn.client.Ping(pingctx, readpref.Primary()); err != nil {
		return fmt.Errorf("ping failed: %v", err)
	}
	return nil
}

func (conn *MongoDBConnection) GetCollection(ctx context.Context, name string) *mongo.Collection {
	return conn.client.Database(conn.dbName).Collection(name)
}

func GetDB() *MongoDBConnection {
	once.Do(func() {
		instance = &MongoDBConnection{}
	})
	return instance
}
