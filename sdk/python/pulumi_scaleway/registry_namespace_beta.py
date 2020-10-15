# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables


class RegistryNamespaceBeta(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    The description of the container registry namespace
    """
    endpoint: pulumi.Output[str]
    """
    The endpoint reachable by docker
    """
    is_public: pulumi.Output[bool]
    """
    Define the default visibity policy
    """
    name: pulumi.Output[str]
    """
    The name of the container registry namespace
    """
    organization_id: pulumi.Output[str]
    """
    The organization_id you want to attach the resource to
    """
    region: pulumi.Output[str]
    """
    The region you want to attach the resource to
    """
    def __init__(__self__, resource_name, opts=None, description=None, is_public=None, name=None, organization_id=None, region=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a RegistryNamespaceBeta resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the container registry namespace
        :param pulumi.Input[bool] is_public: Define the default visibity policy
        :param pulumi.Input[str] name: The name of the container registry namespace
        :param pulumi.Input[str] organization_id: The organization_id you want to attach the resource to
        :param pulumi.Input[str] region: The region you want to attach the resource to
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
            __props__['is_public'] = is_public
            __props__['name'] = name
            __props__['organization_id'] = organization_id
            __props__['region'] = region
            __props__['endpoint'] = None
        super(RegistryNamespaceBeta, __self__).__init__(
            'scaleway:index/registryNamespaceBeta:RegistryNamespaceBeta',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, endpoint=None, is_public=None, name=None, organization_id=None, region=None):
        """
        Get an existing RegistryNamespaceBeta resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the container registry namespace
        :param pulumi.Input[str] endpoint: The endpoint reachable by docker
        :param pulumi.Input[bool] is_public: Define the default visibity policy
        :param pulumi.Input[str] name: The name of the container registry namespace
        :param pulumi.Input[str] organization_id: The organization_id you want to attach the resource to
        :param pulumi.Input[str] region: The region you want to attach the resource to
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["endpoint"] = endpoint
        __props__["is_public"] = is_public
        __props__["name"] = name
        __props__["organization_id"] = organization_id
        __props__["region"] = region
        return RegistryNamespaceBeta(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop