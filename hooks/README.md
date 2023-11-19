# Hooks

To play around with the [Hooks](https://entgo.io/docs/hooks) feature.

Experiment with:
1. What happens if a hook fails:
   1. before `next.Mutate(ctx, m)`
   2. after `next.Mutate(ctx, m)`

Goal:

To figure out where to place the correct hooks, and what to do if a hook fails.

For example:

1. Where should we place a call to elasticsearch to create a record, and how should we handle the error?
  1. Should we rollback everything if just elasticsearch fails?
  2. Do we rollback just the elasticsearch delete if it failed?
  3. How do we isolate that it was elasticsearch that failed?

### 
If Elasticsearch fails:

Option 1: Still create the record
Cons:
1. Data inconsistency

Option 2: Put the create in a queue, and let the queue consume and create the record
Pro:
1. More resiliency

Cons:
1. Another service to read the queue
2. Data can look inconsitent to user until queue is processed

Concerns:
1. What if the queue service is down

Option 3: Put instructions into the same datastore as the ent, and have a service read the instructions
Pro:
1. Rollback on the same datastore is easily achievable

Cons:
1. Sqlite does not allow nested transactions

How can this be achieved without know the global id? A secondary ID then.

Sample scenario:

Scenario 1: System is down, insert, then an update comes in.
We can cancel the update, as it should be the same as insert.


## References

1. [Transaction Hooks](https://entgo.io/docs/transactions/#hooks)

## Gotchas

1. After adding a hook to `Hooks()`, generate your ent schema again.
2. Import the runtime closer to where you instantiate *ent.Client:  `<project>/ent/runtime`

## How to use

### OnCreate

#### Before

Expectation and Actual:
1. User is not created
1. Row is not created

```sh
go run main.go before
```

#### After

Expectation:
1. User is not created
1. Row is not created

Actual:
1. *ent.User is not created
1. Row is created

```sh
go run main.go after
```

#### Never

Expectation and Actual:
1. User is created
1. Row is created

```sh
go run main.go never
```

### Transaction

Questions: what if something fails on `tx.Commit() after`

#### After

Actual:
1. `tx.Commit()` does not return an error
1. Transaction is not committed, you have to rollback

#### Never

Actual:
1. Transaction is committed

## Generate

```sh
go generate ./ent
```