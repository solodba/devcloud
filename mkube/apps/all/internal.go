package all

import (
	_ "github.com/solodba/devcloud/mkube/apps/configmap/impl"
	_ "github.com/solodba/devcloud/mkube/apps/cronjob/impl"
	_ "github.com/solodba/devcloud/mkube/apps/daemonset/impl"
	_ "github.com/solodba/devcloud/mkube/apps/deployment/impl"
	_ "github.com/solodba/devcloud/mkube/apps/harbor/impl"
	_ "github.com/solodba/devcloud/mkube/apps/ingress/impl"
	_ "github.com/solodba/devcloud/mkube/apps/job/impl"
	_ "github.com/solodba/devcloud/mkube/apps/namespace/impl"
	_ "github.com/solodba/devcloud/mkube/apps/node/impl"
	_ "github.com/solodba/devcloud/mkube/apps/pod/impl"
	_ "github.com/solodba/devcloud/mkube/apps/pv/impl"
	_ "github.com/solodba/devcloud/mkube/apps/pvc/impl"
	_ "github.com/solodba/devcloud/mkube/apps/rbac/impl"
	_ "github.com/solodba/devcloud/mkube/apps/sc/impl"
	_ "github.com/solodba/devcloud/mkube/apps/secret/impl"
	_ "github.com/solodba/devcloud/mkube/apps/statefulset/impl"
	_ "github.com/solodba/devcloud/mkube/apps/svc/impl"
)
