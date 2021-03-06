# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables


class SecurityGroup(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    The description of the security group
    """
    enable_default_security: pulumi.Output[bool]
    """
    Add default security group rules
    """
    inbound_default_policy: pulumi.Output[str]
    """
    Default inbound traffic policy for this security group
    """
    name: pulumi.Output[str]
    """
    The name of the security group
    """
    outbound_default_policy: pulumi.Output[str]
    """
    Default outbound traffic policy for this security group
    """
    stateful: pulumi.Output[bool]
    """
    Mark security group as stateful
    """
    def __init__(__self__, resource_name, opts=None, description=None, enable_default_security=None, inbound_default_policy=None, name=None, outbound_default_policy=None, stateful=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a SecurityGroup resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the security group
        :param pulumi.Input[bool] enable_default_security: Add default security group rules
        :param pulumi.Input[str] inbound_default_policy: Default inbound traffic policy for this security group
        :param pulumi.Input[str] name: The name of the security group
        :param pulumi.Input[str] outbound_default_policy: Default outbound traffic policy for this security group
        :param pulumi.Input[bool] stateful: Mark security group as stateful
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

            if description is None:
                raise TypeError("Missing required property 'description'")
            __props__['description'] = description
            __props__['enable_default_security'] = enable_default_security
            __props__['inbound_default_policy'] = inbound_default_policy
            __props__['name'] = name
            __props__['outbound_default_policy'] = outbound_default_policy
            __props__['stateful'] = stateful
        super(SecurityGroup, __self__).__init__(
            'scaleway:index/securityGroup:SecurityGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, enable_default_security=None, inbound_default_policy=None, name=None, outbound_default_policy=None, stateful=None):
        """
        Get an existing SecurityGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the security group
        :param pulumi.Input[bool] enable_default_security: Add default security group rules
        :param pulumi.Input[str] inbound_default_policy: Default inbound traffic policy for this security group
        :param pulumi.Input[str] name: The name of the security group
        :param pulumi.Input[str] outbound_default_policy: Default outbound traffic policy for this security group
        :param pulumi.Input[bool] stateful: Mark security group as stateful
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["enable_default_security"] = enable_default_security
        __props__["inbound_default_policy"] = inbound_default_policy
        __props__["name"] = name
        __props__["outbound_default_policy"] = outbound_default_policy
        __props__["stateful"] = stateful
        return SecurityGroup(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
