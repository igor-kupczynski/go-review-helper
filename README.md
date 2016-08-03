Review Helper for Github
========================

A simple utility to help you review a large PR.

It returns a list of files sorted by number of changes.

## Install

1. TBD
2. Create a file `~/.review-helper.json.in`
3. Put a following content there
  ```
{"token": "your-github-access-token-with-repo-permissions"}
  ```

## Usage

```
$ review-helper <org> <repo> <pr-number>
```

For example to help you review the task cancelation mechanism in elasticsearch -
https://github.com/elastic/elasticsearch/pull/16320 - use the following command.

```
$ review-helper elastic elasticsearch 16320
ore/src/main/java/org/elasticsearch/tasks/TaskManager.java -> 297
core/src/main/java/org/elasticsearch/action/admin/cluster/node/tasks/cancel/TransportCancelTasksAction.java -> 285
core/src/main/java/org/elasticsearch/action/admin/cluster/node/tasks/cancel/CancelTasksRequest.java -> 73
core/src/main/java/org/elasticsearch/rest/action/admin/cluster/node/tasks/RestCancelTasksAction.java -> 62
core/src/main/java/org/elasticsearch/tasks/CancellableTask.java -> 59
core/src/main/java/org/elasticsearch/action/admin/cluster/node/tasks/cancel/CancelTasksAction.java -> 46
core/src/main/java/org/elasticsearch/action/support/tasks/TransportTasksAction.java -> 44
core/src/main/java/org/elasticsearch/action/admin/cluster/node/tasks/cancel/CancelTasksResponse.java -> 42
core/src/main/java/org/elasticsearch/action/admin/cluster/node/tasks/cancel/CancelTasksRequestBuilder.java -> 34
core/src/main/java/org/elasticsearch/action/support/tasks/BaseTasksRequest.java -> 28
core/src/main/java/org/elasticsearch/client/ClusterAdminClient.java -> 26
core/src/main/java/org/elasticsearch/action/support/ChildTaskActionRequest.java -> 22
core/src/main/java/org/elasticsearch/client/support/AbstractClient.java -> 20
core/src/main/java/org/elasticsearch/client/Requests.java -> 14
core/src/main/java/org/elasticsearch/action/support/nodes/TransportNodesAction.java -> 13
core/src/main/java/org/elasticsearch/action/support/broadcast/TransportBroadcastAction.java -> 12
core/src/main/java/org/elasticsearch/action/support/replication/ReplicationRequest.java -> 9
core/src/main/java/org/elasticsearch/tasks/Task.java -> 9
core/src/main/java/org/elasticsearch/action/support/ChildTaskRequest.java -> 8
core/src/main/java/org/elasticsearch/action/admin/cluster/node/tasks/list/ListTasksResponse.java -> 6
core/src/main/java/org/elasticsearch/rest/action/admin/cluster/node/tasks/RestListTasksAction.java -> 4
core/src/main/java/org/elasticsearch/action/admin/cluster/node/tasks/list/TransportListTasksAction.java -> 4
core/src/main/java/org/elasticsearch/action/ActionModule.java -> 3
core/src/main/java/org/elasticsearch/action/support/replication/ReplicationTask.java -> 2
core/src/main/java/org/elasticsearch/common/network/NetworkModule.java -> 2
core/src/main/java/org/elasticsearch/action/support/master/TransportMasterNodeAction.java -> 2
core/src/main/java/org/elasticsearch/action/support/replication/TransportReplicationAction.java -> 1
core/src/main/java/org/elasticsearch/action/support/replication/TransportBroadcastReplicationAction.java -> 1
core/src/main/java/org/elasticsearch/action/support/broadcast/node/TransportBroadcastByNodeAction.java -> 1
core/src/main/java/org/elasticsearch/cluster/service/InternalClusterService.java -> 1
```

Looking at the output it seems you should start with `TaskManager.java` and
`TransportCancelTasksAction.java`!
