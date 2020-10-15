# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables


class IP(pulumi.CustomResource):
    ip: pulumi.Output[str]
    """
    The ipv4 address of the ip
    """
    reverse: pulumi.Output[str]
    """
    The ipv4 reverse dns
    """
    server: pulumi.Output[str]
    """
    The server associated with the ip
    """
    def __init__(__self__, resource_name, opts=None, reverse=None, server=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a IP resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] reverse: The ipv4 reverse dns
        :param pulumi.Input[str] server: The server associated with the ip
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

            if reverse is not None:
                warnings.warn("use scaleway_ip_reverse_dns resource instead", DeprecationWarning)
                pulumi.log.warn("reverse is deprecated: use scaleway_ip_reverse_dns resource instead")
            __props__['reverse'] = reverse
            __props__['server'] = server
            __props__['ip'] = None
        super(IP, __self__).__init__(
            'scaleway:index/iP:IP',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, ip=None, reverse=None, server=None):
        """
        Get an existing IP resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ip: The ipv4 address of the ip
        :param pulumi.Input[str] reverse: The ipv4 reverse dns
        :param pulumi.Input[str] server: The server associated with the ip
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["ip"] = ip
        __props__["reverse"] = reverse
        __props__["server"] = server
        return IP(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop