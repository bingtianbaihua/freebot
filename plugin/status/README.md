## status

给 issue 或 PR 加上 `status/` 前缀的标签，用于指定状态。

状态是唯一的，加上某个状态，会删除之前的状态。每一种状态都可以设置前置条件，在满足某些条件下才能被更改为指定状态。

### cmd

```
/status {status}
```

### extra

参考配置:

```
{
    "extra": {
        "label_precondition": {
            "wip": [],
            "wait-review": [],
            "request-changes": [],
            "approved": [{
                "is_owner": true
            }],
            "testing": [{
                "required_labels": ["status/approved"]
            }],
            "merge-ready": [
                {
                    "is_owner": true
                },
                {
                    "is_qa": true,
                    "required_labels": ["status/testing"]
                }
            ]
        }
    }
}
```

上面的配置表示 wip, wait-review, request-changes 这三种状态任何人可以添加。

approved 状态只能由 owner 修改。

testing 状态需要处于 approved 状态才能修改。

merge-ready 状态有两种情况都可以，一种是 owner 可以直接修改，另外一种是 QA 可以修改且需要处于 testing 状态。

### Support Events

issue_comment, pull_request, pull_request_review_comment

### Support Actions

created
