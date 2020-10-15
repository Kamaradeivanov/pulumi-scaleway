// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface BaremetalServerIp {
    address: string;
    id: string;
    reverse: string;
    version: string;
}

export interface GetBaremetalOfferCpus {
    coreCount: number;
    frequency: number;
    name: string;
    threadCount: number;
}

export interface GetBaremetalOfferDisk {
    capacity: number;
    type: string;
}

export interface GetBaremetalOfferMemory {
    capacity: number;
    frequency: number;
    isEcc: boolean;
    type: string;
}

export interface GetInstanceSecurityGroupInboundRule {
    action: string;
    ip: string;
    ipRange: string;
    port: number;
    portRange: string;
    protocol: string;
}

export interface GetInstanceSecurityGroupOutboundRule {
    action: string;
    ip: string;
    ipRange: string;
    port: number;
    portRange: string;
    protocol: string;
}

export interface GetInstanceServerRootVolume {
    deleteOnTermination: boolean;
    sizeInGb: number;
    volumeId: string;
}

export interface GetInstanceServerUserData {
    key: string;
    value: string;
}

export interface InstanceSecurityGroupInboundRule {
    action: string;
    ip?: string;
    ipRange?: string;
    port?: number;
    portRange?: string;
    protocol?: string;
}

export interface InstanceSecurityGroupOutboundRule {
    action: string;
    ip?: string;
    ipRange?: string;
    port?: number;
    portRange?: string;
    protocol?: string;
}

export interface InstanceSecurityGroupRulesInboundRule {
    action: string;
    ip?: string;
    ipRange?: string;
    port?: number;
    portRange?: string;
    protocol?: string;
}

export interface InstanceSecurityGroupRulesOutboundRule {
    action: string;
    ip?: string;
    ipRange?: string;
    port?: number;
    portRange?: string;
    protocol?: string;
}

export interface InstanceServerRootVolume {
    deleteOnTermination?: boolean;
    sizeInGb: number;
    volumeId: string;
}

export interface InstanceServerUserData {
    key: string;
    value: string;
}

export interface K8SClusterBetaAutoUpgrade {
    enable: boolean;
    maintenanceWindowDay: string;
    maintenanceWindowStartHour: number;
}

export interface K8SClusterBetaAutoscalerConfig {
    balanceSimilarNodeGroups?: boolean;
    disableScaleDown?: boolean;
    estimator?: string;
    expander?: string;
    expendablePodsPriorityCutoff?: number;
    ignoreDaemonsetsUtilization?: boolean;
    scaleDownDelayAfterAdd?: string;
    scaleDownUnneededTime?: string;
}

export interface K8SClusterBetaDefaultPool {
    /**
     * @deprecated This fields is deprecated and will be removed in the next major version, please use scaleway_k8s_pool_beta instead.
     */
    autohealing?: boolean;
    /**
     * @deprecated This fields is deprecated and will be removed in the next major version, please use scaleway_k8s_pool_beta instead.
     */
    autoscaling?: boolean;
    /**
     * @deprecated This fields is deprecated and will be removed in the next major version, please use scaleway_k8s_pool_beta instead.
     */
    containerRuntime?: string;
    createdAt: string;
    /**
     * @deprecated This fields is deprecated and will be removed in the next major version, please use scaleway_k8s_pool_beta instead.
     */
    maxSize: number;
    /**
     * @deprecated This fields is deprecated and will be removed in the next major version, please use scaleway_k8s_pool_beta instead.
     */
    minSize: number;
    /**
     * @deprecated This fields is deprecated and will be removed in the next major version, please use scaleway_k8s_pool_beta instead.
     */
    nodeType: string;
    nodes: outputs.K8SClusterBetaDefaultPoolNode[];
    /**
     * @deprecated This fields is deprecated and will be removed in the next major version, please use scaleway_k8s_pool_beta instead.
     */
    placementGroupId?: string;
    poolId: string;
    /**
     * @deprecated This fields is deprecated and will be removed in the next major version, please use scaleway_k8s_pool_beta instead.
     */
    size: number;
    status: string;
    /**
     * @deprecated This fields is deprecated and will be removed in the next major version, please use scaleway_k8s_pool_beta instead.
     */
    tags?: string[];
    updatedAt: string;
    /**
     * @deprecated This fields is deprecated and will be removed in the next major version, please use scaleway_k8s_pool_beta instead.
     */
    waitForPoolReady?: boolean;
}

export interface K8SClusterBetaDefaultPoolNode {
    name: string;
    publicIp: string;
    publicIpV6: string;
    status: string;
}

export interface K8SClusterBetaKubeconfig {
    clusterCaCertificate: string;
    configFile: string;
    host: string;
    token: string;
}

export interface K8SPoolBetaNode {
    name: string;
    publicIp: string;
    publicIpV6: string;
    status: string;
}

export interface LbBackendBetaHealthCheckHttp {
    code?: number;
    method?: string;
    uri: string;
}

export interface LbBackendBetaHealthCheckHttps {
    code?: number;
    method?: string;
    uri: string;
}

export interface LbBackendBetaHealthCheckTcp {
}

export interface LbCertificateBetaCustomCertificate {
    certificateChain: string;
}

export interface LbCertificateBetaLetsencrypt {
    commonName: string;
    subjectAlternativeNames?: string[];
}

export interface LbFrontendBetaAcl {
    action: outputs.LbFrontendBetaAclAction;
    match: outputs.LbFrontendBetaAclMatch;
    name: string;
    organizationId: string;
    region: string;
}

export interface LbFrontendBetaAclAction {
    type: string;
}

export interface LbFrontendBetaAclMatch {
    httpFilter?: string;
    httpFilterValues?: string[];
    invert?: boolean;
    ipSubnets?: string[];
}

export interface RdbInstanceBetaReadReplica {
    ip: string;
    name: string;
    port: number;
}

export interface ServerVolume {
    sizeInGb: number;
    type: string;
    volumeId: string;
}