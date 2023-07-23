package all

import (
	_ "github.com/solodba/devcloud/mpaas/apps/configmap/api"
	_ "github.com/solodba/devcloud/mpaas/apps/namespace/api"
	_ "github.com/solodba/devcloud/mpaas/apps/node/api"
	_ "github.com/solodba/devcloud/mpaas/apps/pod/api"
	_ "github.com/solodba/devcloud/mpaas/apps/pv/api"
	_ "github.com/solodba/devcloud/mpaas/apps/pvc/api"
	_ "github.com/solodba/devcloud/mpaas/apps/secret/api"
	_ "github.com/solodba/devcloud/mpaas/apps/svc/api"
	_ "github.com/solodba/devcloud/mpaas/apps/test/api"
)
