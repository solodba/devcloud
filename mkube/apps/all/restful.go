package all

import (
	_ "github.com/solodba/devcloud/mkube/apps/configmap/api"
	_ "github.com/solodba/devcloud/mkube/apps/cronjob/api"
	_ "github.com/solodba/devcloud/mkube/apps/daemonset/api"
	_ "github.com/solodba/devcloud/mkube/apps/deployment/api"
	_ "github.com/solodba/devcloud/mkube/apps/ingress/api"
	_ "github.com/solodba/devcloud/mkube/apps/job/api"
	_ "github.com/solodba/devcloud/mkube/apps/namespace/api"
	_ "github.com/solodba/devcloud/mkube/apps/node/api"
	_ "github.com/solodba/devcloud/mkube/apps/pod/api"
	_ "github.com/solodba/devcloud/mkube/apps/pv/api"
	_ "github.com/solodba/devcloud/mkube/apps/pvc/api"
	_ "github.com/solodba/devcloud/mkube/apps/rbac/api"
	_ "github.com/solodba/devcloud/mkube/apps/sc/api"
	_ "github.com/solodba/devcloud/mkube/apps/secret/api"
	_ "github.com/solodba/devcloud/mkube/apps/statefulset/api"
	_ "github.com/solodba/devcloud/mkube/apps/svc/api"
	_ "github.com/solodba/devcloud/mkube/apps/test/api"
)
