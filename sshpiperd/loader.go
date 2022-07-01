package main

import (
	_ "github.com/nickolastone/sshpiper/sshpiperd/upstream/database"
	_ "github.com/nickolastone/sshpiper/sshpiperd/upstream/grpcupstream"
	_ "github.com/nickolastone/sshpiper/sshpiperd/upstream/kubernetes"
	_ "github.com/nickolastone/sshpiper/sshpiperd/upstream/workingdir"
	_ "github.com/nickolastone/sshpiper/sshpiperd/upstream/yaml"

	_ "github.com/nickolastone/sshpiper/sshpiperd/challenger/authy"
	_ "github.com/nickolastone/sshpiper/sshpiperd/challenger/azdevicecode"
	_ "github.com/nickolastone/sshpiper/sshpiperd/challenger/pome"

	_ "github.com/nickolastone/sshpiper/sshpiperd/auditor/typescriptlogger"
)
