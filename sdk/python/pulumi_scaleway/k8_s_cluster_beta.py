# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables


class K8SClusterBeta(pulumi.CustomResource):
    admission_plugins: pulumi.Output[list]
    """
    The list of admission plugins to enable on the cluster
    """
    apiserver_url: pulumi.Output[str]
    """
    Kubernetes API server URL
    """
    auto_upgrade: pulumi.Output[dict]
    """
    The auto upgrade configuration for the cluster

      * `enable` (`bool`)
      * `maintenanceWindowDay` (`str`)
      * `maintenanceWindowStartHour` (`float`)
    """
    autoscaler_config: pulumi.Output[dict]
    """
    The autoscaler configuration for the cluster

      * `balanceSimilarNodeGroups` (`bool`)
      * `disableScaleDown` (`bool`)
      * `estimator` (`str`)
      * `expander` (`str`)
      * `expendablePodsPriorityCutoff` (`float`)
      * `ignoreDaemonsetsUtilization` (`bool`)
      * `scaleDownDelayAfterAdd` (`str`)
      * `scaleDownUnneededTime` (`str`)
    """
    cni: pulumi.Output[str]
    """
    The CNI plugin of the cluster
    """
    created_at: pulumi.Output[str]
    """
    The date and time of the creation of the Kubernetes cluster
    """
    default_pool: pulumi.Output[dict]
    """
    Default pool created for the cluster on creation

      * `autohealing` (`bool`)
      * `autoscaling` (`bool`)
      * `container_runtime` (`str`)
      * `created_at` (`str`)
      * `max_size` (`float`)
      * `min_size` (`float`)
      * `node_type` (`str`)
      * `nodes` (`list`)
        * `name` (`str`)
        * `public_ip` (`str`)
        * `publicIpV6` (`str`)
        * `status` (`str`)

      * `placement_group_id` (`str`)
      * `poolId` (`str`)
      * `size` (`float`)
      * `status` (`str`)
      * `tags` (`list`)
      * `updated_at` (`str`)
      * `wait_for_pool_ready` (`bool`)
    """
    description: pulumi.Output[str]
    """
    The description of the cluster
    """
    enable_dashboard: pulumi.Output[bool]
    """
    Enable the dashboard on the cluster
    """
    feature_gates: pulumi.Output[list]
    """
    The list of feature gates to enable on the cluster
    """
    ingress: pulumi.Output[str]
    """
    The ingress to be deployed on the cluster
    """
    kubeconfig: pulumi.Output[dict]
    """
    The kubeconfig configuration file of the Kubernetes cluster

      * `clusterCaCertificate` (`str`)
      * `configFile` (`str`)
      * `host` (`str`)
      * `token` (`str`)
    """
    name: pulumi.Output[str]
    """
    The name of the cluster
    """
    organization_id: pulumi.Output[str]
    """
    The organization_id you want to attach the resource to
    """
    region: pulumi.Output[str]
    """
    The region you want to attach the resource to
    """
    status: pulumi.Output[str]
    """
    The status of the cluster
    """
    tags: pulumi.Output[list]
    """
    The tags associated with the cluster
    """
    updated_at: pulumi.Output[str]
    """
    The date and time of the last update of the Kubernetes cluster
    """
    upgrade_available: pulumi.Output[bool]
    """
    True if an upgrade is available
    """
    version: pulumi.Output[str]
    """
    The version of the cluster
    """
    wildcard_dns: pulumi.Output[str]
    """
    Wildcard DNS pointing to all the ready nodes
    """
    def __init__(__self__, resource_name, opts=None, admission_plugins=None, auto_upgrade=None, autoscaler_config=None, cni=None, default_pool=None, description=None, enable_dashboard=None, feature_gates=None, ingress=None, name=None, organization_id=None, region=None, tags=None, version=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a K8SClusterBeta resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] admission_plugins: The list of admission plugins to enable on the cluster
        :param pulumi.Input[dict] auto_upgrade: The auto upgrade configuration for the cluster
        :param pulumi.Input[dict] autoscaler_config: The autoscaler configuration for the cluster
        :param pulumi.Input[str] cni: The CNI plugin of the cluster
        :param pulumi.Input[dict] default_pool: Default pool created for the cluster on creation
        :param pulumi.Input[str] description: The description of the cluster
        :param pulumi.Input[bool] enable_dashboard: Enable the dashboard on the cluster
        :param pulumi.Input[list] feature_gates: The list of feature gates to enable on the cluster
        :param pulumi.Input[str] ingress: The ingress to be deployed on the cluster
        :param pulumi.Input[str] name: The name of the cluster
        :param pulumi.Input[str] organization_id: The organization_id you want to attach the resource to
        :param pulumi.Input[str] region: The region you want to attach the resource to
        :param pulumi.Input[list] tags: The tags associated with the cluster
        :param pulumi.Input[str] version: The version of the cluster

        The **auto_upgrade** object supports the following:

          * `enable` (`pulumi.Input[bool]`)
          * `maintenanceWindowDay` (`pulumi.Input[str]`)
          * `maintenanceWindowStartHour` (`pulumi.Input[float]`)

        The **autoscaler_config** object supports the following:

          * `balanceSimilarNodeGroups` (`pulumi.Input[bool]`)
          * `disableScaleDown` (`pulumi.Input[bool]`)
          * `estimator` (`pulumi.Input[str]`)
          * `expander` (`pulumi.Input[str]`)
          * `expendablePodsPriorityCutoff` (`pulumi.Input[float]`)
          * `ignoreDaemonsetsUtilization` (`pulumi.Input[bool]`)
          * `scaleDownDelayAfterAdd` (`pulumi.Input[str]`)
          * `scaleDownUnneededTime` (`pulumi.Input[str]`)

        The **default_pool** object supports the following:

          * `autohealing` (`pulumi.Input[bool]`)
          * `autoscaling` (`pulumi.Input[bool]`)
          * `container_runtime` (`pulumi.Input[str]`)
          * `created_at` (`pulumi.Input[str]`)
          * `max_size` (`pulumi.Input[float]`)
          * `min_size` (`pulumi.Input[float]`)
          * `node_type` (`pulumi.Input[str]`)
          * `nodes` (`pulumi.Input[list]`)
            * `name` (`pulumi.Input[str]`)
            * `public_ip` (`pulumi.Input[str]`)
            * `publicIpV6` (`pulumi.Input[str]`)
            * `status` (`pulumi.Input[str]`)

          * `placement_group_id` (`pulumi.Input[str]`)
          * `poolId` (`pulumi.Input[str]`)
          * `size` (`pulumi.Input[float]`)
          * `status` (`pulumi.Input[str]`)
          * `tags` (`pulumi.Input[list]`)
          * `updated_at` (`pulumi.Input[str]`)
          * `wait_for_pool_ready` (`pulumi.Input[bool]`)
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['admission_plugins'] = admission_plugins
            __props__['auto_upgrade'] = auto_upgrade
            __props__['autoscaler_config'] = autoscaler_config
            if cni is None:
                raise TypeError("Missing required property 'cni'")
            __props__['cni'] = cni
            if default_pool is not None:
                warnings.warn("This fields is deprecated and will be removed in the next major version, please use scaleway_k8s_pool_beta instead.", DeprecationWarning)
                pulumi.log.warn("default_pool is deprecated: This fields is deprecated and will be removed in the next major version, please use scaleway_k8s_pool_beta instead.")
            __props__['default_pool'] = default_pool
            __props__['description'] = description
            __props__['enable_dashboard'] = enable_dashboard
            __props__['feature_gates'] = feature_gates
            __props__['ingress'] = ingress
            __props__['name'] = name
            __props__['organization_id'] = organization_id
            __props__['region'] = region
            __props__['tags'] = tags
            if version is None:
                raise TypeError("Missing required property 'version'")
            __props__['version'] = version
            __props__['apiserver_url'] = None
            __props__['created_at'] = None
            __props__['kubeconfig'] = None
            __props__['status'] = None
            __props__['updated_at'] = None
            __props__['upgrade_available'] = None
            __props__['wildcard_dns'] = None
        super(K8SClusterBeta, __self__).__init__(
            'scaleway:index/k8SClusterBeta:K8SClusterBeta',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, admission_plugins=None, apiserver_url=None, auto_upgrade=None, autoscaler_config=None, cni=None, created_at=None, default_pool=None, description=None, enable_dashboard=None, feature_gates=None, ingress=None, kubeconfig=None, name=None, organization_id=None, region=None, status=None, tags=None, updated_at=None, upgrade_available=None, version=None, wildcard_dns=None):
        """
        Get an existing K8SClusterBeta resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] admission_plugins: The list of admission plugins to enable on the cluster
        :param pulumi.Input[str] apiserver_url: Kubernetes API server URL
        :param pulumi.Input[dict] auto_upgrade: The auto upgrade configuration for the cluster
        :param pulumi.Input[dict] autoscaler_config: The autoscaler configuration for the cluster
        :param pulumi.Input[str] cni: The CNI plugin of the cluster
        :param pulumi.Input[str] created_at: The date and time of the creation of the Kubernetes cluster
        :param pulumi.Input[dict] default_pool: Default pool created for the cluster on creation
        :param pulumi.Input[str] description: The description of the cluster
        :param pulumi.Input[bool] enable_dashboard: Enable the dashboard on the cluster
        :param pulumi.Input[list] feature_gates: The list of feature gates to enable on the cluster
        :param pulumi.Input[str] ingress: The ingress to be deployed on the cluster
        :param pulumi.Input[dict] kubeconfig: The kubeconfig configuration file of the Kubernetes cluster
        :param pulumi.Input[str] name: The name of the cluster
        :param pulumi.Input[str] organization_id: The organization_id you want to attach the resource to
        :param pulumi.Input[str] region: The region you want to attach the resource to
        :param pulumi.Input[str] status: The status of the cluster
        :param pulumi.Input[list] tags: The tags associated with the cluster
        :param pulumi.Input[str] updated_at: The date and time of the last update of the Kubernetes cluster
        :param pulumi.Input[bool] upgrade_available: True if an upgrade is available
        :param pulumi.Input[str] version: The version of the cluster
        :param pulumi.Input[str] wildcard_dns: Wildcard DNS pointing to all the ready nodes

        The **auto_upgrade** object supports the following:

          * `enable` (`pulumi.Input[bool]`)
          * `maintenanceWindowDay` (`pulumi.Input[str]`)
          * `maintenanceWindowStartHour` (`pulumi.Input[float]`)

        The **autoscaler_config** object supports the following:

          * `balanceSimilarNodeGroups` (`pulumi.Input[bool]`)
          * `disableScaleDown` (`pulumi.Input[bool]`)
          * `estimator` (`pulumi.Input[str]`)
          * `expander` (`pulumi.Input[str]`)
          * `expendablePodsPriorityCutoff` (`pulumi.Input[float]`)
          * `ignoreDaemonsetsUtilization` (`pulumi.Input[bool]`)
          * `scaleDownDelayAfterAdd` (`pulumi.Input[str]`)
          * `scaleDownUnneededTime` (`pulumi.Input[str]`)

        The **default_pool** object supports the following:

          * `autohealing` (`pulumi.Input[bool]`)
          * `autoscaling` (`pulumi.Input[bool]`)
          * `container_runtime` (`pulumi.Input[str]`)
          * `created_at` (`pulumi.Input[str]`)
          * `max_size` (`pulumi.Input[float]`)
          * `min_size` (`pulumi.Input[float]`)
          * `node_type` (`pulumi.Input[str]`)
          * `nodes` (`pulumi.Input[list]`)
            * `name` (`pulumi.Input[str]`)
            * `public_ip` (`pulumi.Input[str]`)
            * `publicIpV6` (`pulumi.Input[str]`)
            * `status` (`pulumi.Input[str]`)

          * `placement_group_id` (`pulumi.Input[str]`)
          * `poolId` (`pulumi.Input[str]`)
          * `size` (`pulumi.Input[float]`)
          * `status` (`pulumi.Input[str]`)
          * `tags` (`pulumi.Input[list]`)
          * `updated_at` (`pulumi.Input[str]`)
          * `wait_for_pool_ready` (`pulumi.Input[bool]`)

        The **kubeconfig** object supports the following:

          * `clusterCaCertificate` (`pulumi.Input[str]`)
          * `configFile` (`pulumi.Input[str]`)
          * `host` (`pulumi.Input[str]`)
          * `token` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["admission_plugins"] = admission_plugins
        __props__["apiserver_url"] = apiserver_url
        __props__["auto_upgrade"] = auto_upgrade
        __props__["autoscaler_config"] = autoscaler_config
        __props__["cni"] = cni
        __props__["created_at"] = created_at
        __props__["default_pool"] = default_pool
        __props__["description"] = description
        __props__["enable_dashboard"] = enable_dashboard
        __props__["feature_gates"] = feature_gates
        __props__["ingress"] = ingress
        __props__["kubeconfig"] = kubeconfig
        __props__["name"] = name
        __props__["organization_id"] = organization_id
        __props__["region"] = region
        __props__["status"] = status
        __props__["tags"] = tags
        __props__["updated_at"] = updated_at
        __props__["upgrade_available"] = upgrade_available
        __props__["version"] = version
        __props__["wildcard_dns"] = wildcard_dns
        return K8SClusterBeta(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
