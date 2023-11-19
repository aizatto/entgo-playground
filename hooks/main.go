package main

import (
	"context"
	"fmt"
	"hooks/ent"
	"log"
	"os"

	_ "hooks/ent/runtime"
	"hooks/ent/user"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
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

	entClient, err := getEntClient(ctx)
	if err != nil {
		return err
	}
	defer entClient.Close()

	if len(os.Args) < 3 {
		return errors.New("expected arguments: <oncreate|onupdate|ondelete|commit|rollback> <before|after|never>")
	}

	state := os.Args[2]

	switch os.Args[1] {
	case "oncreate":
		return onCreate(ctx, entClient, state)
	case "commit":
		return commit(ctx, entClient, state)
	default:
		return fmt.Errorf("unknown argument: %s", os.Args[1])
	}
}

func onCreate(ctx context.Context, entClient *ent.Client, state string) error {
	var entuser *ent.User
	var err error

	switch state {
	case "before":
		entuser, err = entClient.
			User.
			Create().
			SetFail(user.FailBefore).
			Save(ctx)
	case "after":
		entuser, err = entClient.
			User.
			Create().
			SetFail(user.FailAfter).
			Save(ctx)
	case "never":
		entuser, err = entClient.
			User.
			Create().
			SetFail(user.FailNever).
			Save(ctx)
	default:
		err = fmt.Errorf("unknown argument: %s", state)
	}

	if entuser != nil {
		log.Println("*ent.User is created")
	}

	count, _ := entClient.User.Query().Count(ctx)
	if count != 0 {
		log.Println("user row is created")
	}

	return err
}

func commit(ctx context.Context, entClient *ent.Client, state1 string) error {
	txstate := "never"
	if len(os.Args) > 3 {
		txstate = os.Args[3]
	}

	tx, err := entClient.Tx(ctx)
	if err != nil {
		return err
	}

	tx.OnCommit(func(next ent.Committer) ent.Committer {
		return ent.CommitFunc(func(ctx context.Context, tx *ent.Tx) error {
			log.Println("tx: before commit")
			if txstate == "beforecommit" {
				return errors.New("tx: failed before commit")
			}

			err := next.Commit(ctx, tx)

			log.Println("tx: after commit")
			if txstate == "aftercommit" {
				return errors.New("tx: failed after commit")
			}

			return err
		})
	})

	tx.OnRollback(func(next ent.Rollbacker) ent.Rollbacker {
		return ent.RollbackFunc(func(ctx context.Context, tx *ent.Tx) error {
			log.Println("tx: before rollback")
			if txstate == "afterrollback" {
				return errors.New("tx: failed before rollbac")
			}

			err := next.Rollback(ctx, tx)

			log.Println("tx: after rollback")
			if txstate == "beforerollback" {
				return errors.New("tx: failed after rollback")
			}

			return err
		})
	})

	err = onCreate(ctx, tx.Client(), state1)
	if err != nil {
		// todo: decide if we should swallow this error if its known
		// return errors.Wrap(err, "tx failed onCreate")
		err = tx.Rollback()
		if err != nil {
			return errors.Wrap(err, "tx: failed to rollback on create")
		}
		return nil
	}

	err = tx.Commit()
	if err != nil {
		// todo: decide if we should swallow this error if its known
		// return errors.Wrap(err, "tx: failed to commit")
		err = tx.Rollback()
		if err != nil {
			return errors.Wrap(err, "tx: failed to rollback on commit")
		}
	}
	return nil
}

func getEntClient(ctx context.Context) (*ent.Client, error) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		return nil, errors.Wrap(err, "failed opening connection to database")
	}

	if err := client.Schema.Create(ctx); err != nil {
		return nil, errors.Wrap(err, "failed creating schema resources")
	}

	return client.Debug(), err
}
