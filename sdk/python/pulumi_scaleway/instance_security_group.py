# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables


class InstanceSecurityGroup(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    The description of the security group
    """
    external_rules: pulumi.Output[bool]
    inbound_default_policy: pulumi.Output[str]
    """
    Default inbound traffic policy for this security group
    """
    inbound_rules: pulumi.Output[list]
    """
    Inbound rules for this security group

      * `action` (`str`)
      * `ip` (`str`)
      * `ip_range` (`str`)
      * `port` (`float`)
      * `portRange` (`str`)
      * `protocol` (`str`)
    """
    name: pulumi.Output[str]
    """
    The name of the security group
    """
    organization_id: pulumi.Output[str]
    """
    The organization_id you want to attach the resource to
    """
    outbound_default_policy: pulumi.Output[str]
    """
    Default outbound traffic policy for this security group
    """
    outbound_rules: pulumi.Output[list]
    """
    Outbound rules for this security group

      * `action` (`str`)
      * `ip` (`str`)
      * `ip_range` (`str`)
      * `port` (`float`)
      * `portRange` (`str`)
      * `protocol` (`str`)
    """
    stateful: pulumi.Output[bool]
    """
    The stateful value of the security group
    """
    zone: pulumi.Output[str]
    """
    The zone you want to attach the resource to
    """
    def __init__(__self__, resource_name, opts=None, description=None, external_rules=None, inbound_default_policy=None, inbound_rules=None, name=None, organization_id=None, outbound_default_policy=None, outbound_rules=None, stateful=None, zone=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a InstanceSecurityGroup resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the security group
        :param pulumi.Input[str] inbound_default_policy: Default inbound traffic policy for this security group
        :param pulumi.Input[list] inbound_rules: Inbound rules for this security group
        :param pulumi.Input[str] name: The name of the security group
        :param pulumi.Input[str] organization_id: The organization_id you want to attach the resource to
        :param pulumi.Input[str] outbound_default_policy: Default outbound traffic policy for this security group
        :param pulumi.Input[list] outbound_rules: Outbound rules for this security group
        :param pulumi.Input[bool] stateful: The stateful value of the security group
        :param pulumi.Input[str] zone: The zone you want to attach the resource to

        The **inbound_rules** object supports the following:

          * `action` (`pulumi.Input[str]`)
          * `ip` (`pulumi.Input[str]`)
          * `ip_range` (`pulumi.Input[str]`)
          * `port` (`pulumi.Input[float]`)
          * `portRange` (`pulumi.Input[str]`)
          * `protocol` (`pulumi.Input[str]`)

        The **outbound_rules** object supports the following:

          * `action` (`pulumi.Input[str]`)
          * `ip` (`pulumi.Input[str]`)
          * `ip_range` (`pulumi.Input[str]`)
          * `port` (`pulumi.Input[float]`)
          * `portRange` (`pulumi.Input[str]`)
          * `protocol` (`pulumi.Input[str]`)
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

            __props__['description'] = description
            __props__['external_rules'] = external_rules
            __props__['inbound_default_policy'] = inbound_default_policy
            __props__['inbound_rules'] = inbound_rules
            __props__['name'] = name
            __props__['organization_id'] = organization_id
            __props__['outbound_default_policy'] = outbound_default_policy
            __props__['outbound_rules'] = outbound_rules
            __props__['stateful'] = stateful
            __props__['zone'] = zone
        super(InstanceSecurityGroup, __self__).__init__(
            'scaleway:index/instanceSecurityGroup:InstanceSecurityGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, external_rules=None, inbound_default_policy=None, inbound_rules=None, name=None, organization_id=None, outbound_default_policy=None, outbound_rules=None, stateful=None, zone=None):
        """
        Get an existing InstanceSecurityGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the security group
        :param pulumi.Input[str] inbound_default_policy: Default inbound traffic policy for this security group
        :param pulumi.Input[list] inbound_rules: Inbound rules for this security group
        :param pulumi.Input[str] name: The name of the security group
        :param pulumi.Input[str] organization_id: The organization_id you want to attach the resource to
        :param pulumi.Input[str] outbound_default_policy: Default outbound traffic policy for this security group
        :param pulumi.Input[list] outbound_rules: Outbound rules for this security group
        :param pulumi.Input[bool] stateful: The stateful value of the security group
        :param pulumi.Input[str] zone: The zone you want to attach the resource to

        The **inbound_rules** object supports the following:

          * `action` (`pulumi.Input[str]`)
          * `ip` (`pulumi.Input[str]`)
          * `ip_range` (`pulumi.Input[str]`)
          * `port` (`pulumi.Input[float]`)
          * `portRange` (`pulumi.Input[str]`)
          * `protocol` (`pulumi.Input[str]`)

        The **outbound_rules** object supports the following:

          * `action` (`pulumi.Input[str]`)
          * `ip` (`pulumi.Input[str]`)
          * `ip_range` (`pulumi.Input[str]`)
          * `port` (`pulumi.Input[float]`)
          * `portRange` (`pulumi.Input[str]`)
          * `protocol` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["external_rules"] = external_rules
        __props__["inbound_default_policy"] = inbound_default_policy
        __props__["inbound_rules"] = inbound_rules
        __props__["name"] = name
        __props__["organization_id"] = organization_id
        __props__["outbound_default_policy"] = outbound_default_policy
        __props__["outbound_rules"] = outbound_rules
        __props__["stateful"] = stateful
        __props__["zone"] = zone
        return InstanceSecurityGroup(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
