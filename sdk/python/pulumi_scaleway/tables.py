# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

_SNAKE_TO_CAMEL_CASE_TABLE = {
    "access_key": "accessKey",
    "additional_volume_ids": "additionalVolumeIds",
    "admission_plugins": "admissionPlugins",
    "apiserver_url": "apiserverUrl",
    "auto_upgrade": "autoUpgrade",
    "autoscaler_config": "autoscalerConfig",
    "backend_id": "backendId",
    "boot_type": "bootType",
    "certificate_id": "certificateId",
    "cloud_init": "cloudInit",
    "cluster_id": "clusterId",
    "common_name": "commonName",
    "container_runtime": "containerRuntime",
    "created_at": "createdAt",
    "creation_ip": "creationIp",
    "current_size": "currentSize",
    "custom_certificate": "customCertificate",
    "default_pool": "defaultPool",
    "disable_backup": "disableBackup",
    "dynamic_ip_required": "dynamicIpRequired",
    "enable_dashboard": "enableDashboard",
    "enable_default_security": "enableDefaultSecurity",
    "enable_dynamic_ip": "enableDynamicIp",
    "enable_ipv6": "enableIpv6",
    "endpoint_ip": "endpointIp",
    "endpoint_port": "endpointPort",
    "expiration_date": "expirationDate",
    "external_rules": "externalRules",
    "feature_gates": "featureGates",
    "forward_port": "forwardPort",
    "forward_port_algorithm": "forwardPortAlgorithm",
    "forward_protocol": "forwardProtocol",
    "from_snapshot_id": "fromSnapshotId",
    "from_volume_id": "fromVolumeId",
    "health_check_delay": "healthCheckDelay",
    "health_check_http": "healthCheckHttp",
    "health_check_https": "healthCheckHttps",
    "health_check_max_retries": "healthCheckMaxRetries",
    "health_check_port": "healthCheckPort",
    "health_check_tcp": "healthCheckTcp",
    "health_check_timeout": "healthCheckTimeout",
    "inbound_default_policy": "inboundDefaultPolicy",
    "inbound_port": "inboundPort",
    "inbound_rules": "inboundRules",
    "ip_address": "ipAddress",
    "ip_id": "ipId",
    "ip_range": "ipRange",
    "ipv6_address": "ipv6Address",
    "ipv6_gateway": "ipv6Gateway",
    "ipv6_prefix_length": "ipv6PrefixLength",
    "is_ha_cluster": "isHaCluster",
    "is_public": "isPublic",
    "lb_id": "lbId",
    "max_size": "maxSize",
    "min_size": "minSize",
    "node_type": "nodeType",
    "not_valid_after": "notValidAfter",
    "not_valid_before": "notValidBefore",
    "offer_id": "offerId",
    "on_marked_down_action": "onMarkedDownAction",
    "organization_id": "organizationId",
    "os_id": "osId",
    "outbound_default_policy": "outboundDefaultPolicy",
    "outbound_rules": "outboundRules",
    "placement_group_id": "placementGroupId",
    "placement_group_policy_respected": "placementGroupPolicyRespected",
    "policy_mode": "policyMode",
    "policy_respected": "policyRespected",
    "policy_type": "policyType",
    "private_ip": "privateIp",
    "public_ip": "publicIp",
    "public_ipv6": "publicIpv6",
    "public_key": "publicKey",
    "read_replicas": "readReplicas",
    "root_volume": "rootVolume",
    "secret_key": "secretKey",
    "security_group": "securityGroup",
    "security_group_id": "securityGroupId",
    "send_proxy_v2": "sendProxyV2",
    "server_id": "serverId",
    "server_ips": "serverIps",
    "size_in_gb": "sizeInGb",
    "ssh_key_ids": "sshKeyIds",
    "state_detail": "stateDetail",
    "sticky_sessions": "stickySessions",
    "sticky_sessions_cookie_name": "stickySessionsCookieName",
    "subject_alternative_name": "subjectAlternativeName",
    "timeout_client": "timeoutClient",
    "timeout_connect": "timeoutConnect",
    "timeout_server": "timeoutServer",
    "timeout_tunnel": "timeoutTunnel",
    "updated_at": "updatedAt",
    "upgrade_available": "upgradeAvailable",
    "user_datas": "userDatas",
    "user_id": "userId",
    "user_name": "userName",
    "wait_for_pool_ready": "waitForPoolReady",
    "wildcard_dns": "wildcardDns",
}

_CAMEL_TO_SNAKE_CASE_TABLE = {
    "accessKey": "access_key",
    "additionalVolumeIds": "additional_volume_ids",
    "admissionPlugins": "admission_plugins",
    "apiserverUrl": "apiserver_url",
    "autoUpgrade": "auto_upgrade",
    "autoscalerConfig": "autoscaler_config",
    "backendId": "backend_id",
    "bootType": "boot_type",
    "certificateId": "certificate_id",
    "cloudInit": "cloud_init",
    "clusterId": "cluster_id",
    "commonName": "common_name",
    "containerRuntime": "container_runtime",
    "createdAt": "created_at",
    "creationIp": "creation_ip",
    "currentSize": "current_size",
    "customCertificate": "custom_certificate",
    "defaultPool": "default_pool",
    "disableBackup": "disable_backup",
    "dynamicIpRequired": "dynamic_ip_required",
    "enableDashboard": "enable_dashboard",
    "enableDefaultSecurity": "enable_default_security",
    "enableDynamicIp": "enable_dynamic_ip",
    "enableIpv6": "enable_ipv6",
    "endpointIp": "endpoint_ip",
    "endpointPort": "endpoint_port",
    "expirationDate": "expiration_date",
    "externalRules": "external_rules",
    "featureGates": "feature_gates",
    "forwardPort": "forward_port",
    "forwardPortAlgorithm": "forward_port_algorithm",
    "forwardProtocol": "forward_protocol",
    "fromSnapshotId": "from_snapshot_id",
    "fromVolumeId": "from_volume_id",
    "healthCheckDelay": "health_check_delay",
    "healthCheckHttp": "health_check_http",
    "healthCheckHttps": "health_check_https",
    "healthCheckMaxRetries": "health_check_max_retries",
    "healthCheckPort": "health_check_port",
    "healthCheckTcp": "health_check_tcp",
    "healthCheckTimeout": "health_check_timeout",
    "inboundDefaultPolicy": "inbound_default_policy",
    "inboundPort": "inbound_port",
    "inboundRules": "inbound_rules",
    "ipAddress": "ip_address",
    "ipId": "ip_id",
    "ipRange": "ip_range",
    "ipv6Address": "ipv6_address",
    "ipv6Gateway": "ipv6_gateway",
    "ipv6PrefixLength": "ipv6_prefix_length",
    "isHaCluster": "is_ha_cluster",
    "isPublic": "is_public",
    "lbId": "lb_id",
    "maxSize": "max_size",
    "minSize": "min_size",
    "nodeType": "node_type",
    "notValidAfter": "not_valid_after",
    "notValidBefore": "not_valid_before",
    "offerId": "offer_id",
    "onMarkedDownAction": "on_marked_down_action",
    "organizationId": "organization_id",
    "osId": "os_id",
    "outboundDefaultPolicy": "outbound_default_policy",
    "outboundRules": "outbound_rules",
    "placementGroupId": "placement_group_id",
    "placementGroupPolicyRespected": "placement_group_policy_respected",
    "policyMode": "policy_mode",
    "policyRespected": "policy_respected",
    "policyType": "policy_type",
    "privateIp": "private_ip",
    "publicIp": "public_ip",
    "publicIpv6": "public_ipv6",
    "publicKey": "public_key",
    "readReplicas": "read_replicas",
    "rootVolume": "root_volume",
    "secretKey": "secret_key",
    "securityGroup": "security_group",
    "securityGroupId": "security_group_id",
    "sendProxyV2": "send_proxy_v2",
    "serverId": "server_id",
    "serverIps": "server_ips",
    "sizeInGb": "size_in_gb",
    "sshKeyIds": "ssh_key_ids",
    "stateDetail": "state_detail",
    "stickySessions": "sticky_sessions",
    "stickySessionsCookieName": "sticky_sessions_cookie_name",
    "subjectAlternativeName": "subject_alternative_name",
    "timeoutClient": "timeout_client",
    "timeoutConnect": "timeout_connect",
    "timeoutServer": "timeout_server",
    "timeoutTunnel": "timeout_tunnel",
    "updatedAt": "updated_at",
    "upgradeAvailable": "upgrade_available",
    "userDatas": "user_datas",
    "userId": "user_id",
    "userName": "user_name",
    "waitForPoolReady": "wait_for_pool_ready",
    "wildcardDns": "wildcard_dns",
}
