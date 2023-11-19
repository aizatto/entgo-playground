package main

import (
	"context"
	"log"
	"os"
	"queuetable/ent"
	"queuetable/ent/user"
	"queuetable/internal/clients"

	"github.com/google/uuid"
)

func main() {
	err := run()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()

	entClient, err := clients.EntClient(ctx)
	if err != nil {
		return err
	}
	defer entClient.Close()

	tx, err := entClient.Tx(ctx)
	if err != nil {
		return err
	}

	err = createuser(ctx, tx)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return err
		}
	}

	return nil
}

func createuser(ctx context.Context, tx *ent.Tx) error {
	queueid := uuid.New()

	_, err := tx.
		User.
		Create().
		SetQueueObjID(queueid).
		Save(ctx)
	if err != nil {
		return err
	}

	_, err = tx.
		Queue.
		Create().
		SetInstruction("create").
		SetObjTable(user.Table).
		SetObjID(queueid).
		Save(ctx)

	return err
}
