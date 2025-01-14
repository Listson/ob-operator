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

package converter

import (
	"github.com/pkg/errors"

	cloudv1 "github.com/oceanbase/ob-operator/apis/cloud/v1"
	myconfig "github.com/oceanbase/ob-operator/pkg/config"
	observerconst "github.com/oceanbase/ob-operator/pkg/controllers/observer/const"
	"github.com/oceanbase/ob-operator/pkg/controllers/observer/model"
	"github.com/oceanbase/ob-operator/pkg/controllers/observer/sql"
	statefulappCore "github.com/oceanbase/ob-operator/pkg/controllers/statefulapp/const"
)

func IsAllOBServerActive(obServerList []model.AllServer, obClusters []cloudv1.Cluster) bool {
	obServerCurrentReplicas := make(map[string]bool)
	for _, obServer := range obServerList {
		if obServer.Status == observerconst.OBServerActive {
			obServerCurrentReplicas[obServer.Zone] = true
		}
	}
	for _, obCluster := range obClusters {
		if obCluster.Cluster == myconfig.ClusterName {
			for _, zone := range obCluster.Zone {
				tmp := obServerCurrentReplicas[zone.Name]
				if !tmp {
					return false
				}
			}
			return true
		}
	}
	return false
}

func IsOBServerDeleted(clusterIP, podIP string) bool {
	obServerList := sql.GetOBServer(clusterIP)
	for _, obServer := range obServerList {
		if obServer.SvrIP == podIP {
			return false
		}
	}
	return true
}

func IsPodNotInOBServerList(zoneName, ip string, nodeMap map[string][]cloudv1.OBNode) bool {
	zoneIPList := nodeMap[zoneName]
	if len(zoneIPList) > 0 {
		for _, tmpIP := range zoneIPList {
			if tmpIP.ServerIP == ip {
				return false
			}
		}
		return true
	}
	return false
}

func IsOBServerInactiveOrDeletingAndNotInPodList(server cloudv1.OBNode, podRunningList []string) bool {
	if server.Status == observerconst.OBServerInactive || server.Status == observerconst.OBServerDeleting {
		for _, podIP := range podRunningList {
			if podIP == server.ServerIP {
				return false
			}
		}
		return true
	}
	return false
}

func GetInfoForAddServerByZone(clusterIP string, statefulApp cloudv1.StatefulApp) (error, string, string) {
	obServerList := sql.GetOBServer(clusterIP)
	if len(obServerList) == 0 {
		return errors.New(observerconst.DataBaseError), "", ""
	}

	nodeMap := GenerateNodeMapByOBServerList(obServerList)

	// judge witch ip need add
	for _, subset := range statefulApp.Status.Subsets {
		for _, pod := range subset.Pods {
			if pod.PodPhase == statefulappCore.PodStatusRunning {
				status := IsPodNotInOBServerList(subset.Name, pod.PodIP, nodeMap)
				// Pod IP not in OBServerList, need to add server
				// do one thing at a time
				if status {
					return nil, subset.Name, pod.PodIP
				}
			}
		}
	}

	return errors.New("none ip need add"), "", ""
}

func GetInfoForDelServerByZone(clusterIP string, clusterSpec cloudv1.Cluster, statefulApp cloudv1.StatefulApp) (error, string, string) {
	obServerList := sql.GetOBServer(clusterIP)
	if len(obServerList) == 0 {
		return errors.New(observerconst.DataBaseError), "", ""
	}

	nodeMap := GenerateNodeMapByOBServerList(obServerList)

	// judge witch ip need del
	for _, subset := range statefulApp.Status.Subsets {
		runningPodList := getRunningPodListFromSubsetStatus(subset)
		zoneSpec := GetZoneSpecFromClusterSpec(subset.Name, clusterSpec)
		// number of observer in db > replica
		if len(nodeMap[subset.Name]) > int(zoneSpec.Replicas) {
			for _, pod := range nodeMap[subset.Name] {
				status := IsOBServerInactiveOrDeletingAndNotInPodList(pod, runningPodList)
				if status {
					// OBServer IP not in PodList, need to delete server
					// do one thing at a time
					return nil, subset.Name, pod.ServerIP
				}
			}
		}
	}

	return errors.New("none ip need del"), "", ""
}

func getRunningPodListFromSubsetStatus(subset cloudv1.SubsetStatus) []string {
	runningPodList := make([]string, 0)
	for _, pod := range subset.Pods {
		if pod.PodPhase == statefulappCore.PodStatusRunning {
			runningPodList = append(runningPodList, pod.PodIP)
		}
	}
	return runningPodList
}
