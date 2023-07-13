package conf

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *MongoDB) GetClient() (*mongo.Client, error) {
	opts := &options.ClientOptions{
		Hosts: m.Endpoints,
	}
	cred := options.Credential{}
	if m.Username != "" && m.Password != "" {
		cred.PasswordSet = true
		cred.Username = m.Username
		cred.Password = m.Password
		cred.AuthSource = m.AuthSource
		opts.SetAuth(cred)
	}
	opts.SetTimeout(time.Second * time.Duration(5))
	client, err := mongo.NewClient(opts)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(5))
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (m *MongoDB) Client() (*mongo.Client, error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if m.client == nil {
		client, err := m.GetClient()
		if err != nil {
			return nil, err
		}
		m.client = client
	}
	return m.client, nil
}

func (m *MongoDB) GetDB() (*mongo.Database, error) {
	client, err := m.Client()
	if err != nil {
		return nil, err
	}
	return client.Database(m.Database), nil
}
