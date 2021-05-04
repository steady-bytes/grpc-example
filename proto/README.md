# supported types

1. Service
  An rpc service interface that runs on port 55000. The current supported methods are Query, that takes a key/val map as a parameter.
  
  __Required Archtectural Services__ Services for all event driven systems.
  A. CommandHandler - Process a command into an event
      Apply()
  B. EventStore - Stores the event, and forwards based on the aggregate type to the correct data stream
      Store()
      

2. Command
  Business processes that needs to be achived. Something like `PlaceOrder` where a few asyncrounus tasks need to be carried out.

  A message that a client of the system executes to invoke some action on the remaining system
  ```

3. Events
  Completed task with data the system has processed.

4. Workers
  A process that is listening do a data stream and processing messages from that stream.

5.
  
