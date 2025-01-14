/*
Copyright (c) 2021 OceanBase
ob-operator is licensed under Mulan PSL v2.
You can use this software according to the terms and conditions of the Mulan PSL v2.
You may obtain a copy of Mulan PSL v2 at:
         http://license.coscl.org.cn/MulanPSL2
THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
See the Mulan PSL v2 for more details.
*/

package cable

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"k8s.io/klog/v2"

	cloudv1 "github.com/oceanbase/ob-operator/apis/cloud/v1"
	myconfig "github.com/oceanbase/ob-operator/pkg/config"
	observerconst "github.com/oceanbase/ob-operator/pkg/controllers/observer/const"
	"github.com/oceanbase/ob-operator/pkg/infrastructure/ob"
)

func GenerateRSListFromSubset(subsets []cloudv1.SubsetStatus) string {
	var rsList string
	for _, subset := range subsets {
		podList := subset.Pods
		if len(podList) > 0 {
			podIP := podList[0].PodIP
			if podIP != "" {
				if rsList == "" {
					rsList = fmt.Sprintf("%s:%s:%s", podIP, ob.OBSERVER_RPC_PORT, ob.OBSERVER_MYSQL_PORT)
				} else {
					rsList = fmt.Sprintf("%s,%s:%s:%s", rsList, podIP, ob.OBSERVER_RPC_PORT, ob.OBSERVER_MYSQL_PORT)
				}
			} else {
				klog.Errorln("pod ip is empty", subsets)
			}
		}
	}
	return rsList
}

func GenerateRSListFromRootServiceStatus(topology []cloudv1.ClusterRootServiceStatus) string {
	var rsList string
	for _, cluster := range topology {
		if cluster.Cluster == myconfig.ClusterName {
			for _, zone := range cluster.Zone {
				if zone.ServerIP != "" && zone.Status == observerconst.OBServerActive {
					if rsList == "" {
						rsList = fmt.Sprintf("%s:%s:%s", zone.ServerIP, ob.OBSERVER_RPC_PORT, ob.OBSERVER_MYSQL_PORT)
					} else {
						rsList = fmt.Sprintf("%s,%s:%s:%s", rsList, zone.ServerIP, ob.OBSERVER_RPC_PORT, ob.OBSERVER_MYSQL_PORT)
					}
				}
			}
		}
		return rsList
	}
	return rsList
}

func GenerateOBServerStartArgs(obCluster cloudv1.OBCluster, zoneName, rsList string) map[string]interface{} {
	obServerStartArgs := make(map[string]interface{})
	obServerStartArgs["clusterName"] = obCluster.Name
	obServerStartArgs["clusterId"] = obCluster.Spec.ClusterID
	obServerStartArgs["zoneName"] = zoneName
	obServerStartArgs["rsList"] = rsList
	cpu, _ := obCluster.Spec.Resources.CPU.AsInt64()
	obServerStartArgs["cpuLimit"] = cpu
	memory, _ := obCluster.Spec.Resources.Memory.AsInt64()
	obServerStartArgs["memoryLimit"] = memory / 1024 / 1024 / 1024
	return obServerStartArgs
}

func GenerateOBClusterBootstrapArgs(subsets []cloudv1.SubsetStatus) (string, error) {
	rsList := make([]ob.RSInfo, 0)
	for _, subset := range subsets {
		podList := subset.Pods
		if len(podList) > 0 {
			podIP := podList[0].PodIP
			port, err := strconv.Atoi(ob.OBSERVER_RPC_PORT)
			if err != nil {
				klog.Errorln("the zone don't have first server", subset)
				return "", errors.New("the zone don't have first server")
			}
			var rsInfo ob.RSInfo
			rsInfo.Region = subset.Region
			rsInfo.Zone = subset.Name
			rsInfo.Server.Ip = podIP
			rsInfo.Server.Port = port
			rsList = append(rsList, rsInfo)
		} else {
			klog.Errorln("the zone don't have first server", subset)
			return "", errors.New("the zone don't have first server")
		}
	}
	var bootstrapParam ob.BootstrapParam
	bootstrapParam.ClusterType = ob.PRIMARY
	bootstrapParam.RSInfoList = rsList
	obclusterBootstrapArgs := ob.GenerateBootstrapSQL(bootstrapParam)
	klog.Infoln("OBCluster bootstrap args", obclusterBootstrapArgs)
	return obclusterBootstrapArgs, nil
}
