# Queue Table

Goal:

1. Reuse a queue.
1. Figure out a data model that allows me to not worry about relying on another service for queues.

Requirements:

1. `queue_obj_id` is a unique id but is not `queues.id`
1. Polymorphic relationship
1. queue not reading the same row

Alternatives:

1. Instead of a `queue_obj_id`, can use a `global id` similar to relay's global id

Concerns:

1. Multiple queues reading from the `queues` table
1. What if the process reading the queue dies

## Conclusion

We don't have to use the same datastore, we can use another service. As long as there is only a single check afterwards we are ok.

Just remember to check the error code of putting it into the queue, and if it fails, rollback.

Concerns / Limitations:
1. Multiple creates, deletes in the same transaction
1. Cascading create calls to queue

You can put the desired action in the ent hook, but as long as it is just one.

## Gotchas

1. Main table uuid is required

## Generate

```sh
go generate ./ent
```

## References

1. [Database Locking Techniques with Ent](https://entgo.io/blog/2021/07/22/database-locking-techniques-with-ent/)
  1. In pessimisitic locking, SQLite does not support `SELECT ... FOR UPDATE`
  1. [MySQL InnoDB Locking Reads](https://dev.mysql.com/doc/refman/8.0/en/innodb-locking-reads.html)
1. [Ent Go Documentation: Row level locks](https://entgo.io/docs/feature-flags/#row-level-locks)