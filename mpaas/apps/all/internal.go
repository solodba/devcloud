package all

import (
	_ "github.com/solodba/devcloud/mpaas/apps/configmap/impl"
	_ "github.com/solodba/devcloud/mpaas/apps/daemonset/impl"
	_ "github.com/solodba/devcloud/mpaas/apps/deployment/impl"
	_ "github.com/solodba/devcloud/mpaas/apps/ingress/impl"
	_ "github.com/solodba/devcloud/mpaas/apps/job/impl"
	_ "github.com/solodba/devcloud/mpaas/apps/namespace/impl"
	_ "github.com/solodba/devcloud/mpaas/apps/node/impl"
	_ "github.com/solodba/devcloud/mpaas/apps/pod/impl"
	_ "github.com/solodba/devcloud/mpaas/apps/pv/impl"
	_ "github.com/solodba/devcloud/mpaas/apps/pvc/impl"
	_ "github.com/solodba/devcloud/mpaas/apps/sc/impl"
	_ "github.com/solodba/devcloud/mpaas/apps/secret/impl"
	_ "github.com/solodba/devcloud/mpaas/apps/statefulset/impl"
	_ "github.com/solodba/devcloud/mpaas/apps/svc/impl"
)
