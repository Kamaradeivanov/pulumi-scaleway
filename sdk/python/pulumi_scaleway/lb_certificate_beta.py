# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables


class LbCertificateBeta(pulumi.CustomResource):
    common_name: pulumi.Output[str]
    """
    The main domain name of the certificate
    """
    custom_certificate: pulumi.Output[dict]
    """
    The custom type certificate type configuration

      * `certificateChain` (`str`)
    """
    fingerprint: pulumi.Output[str]
    """
    The identifier (SHA-1) of the certificate
    """
    lb_id: pulumi.Output[str]
    """
    The load-balancer ID
    """
    letsencrypt: pulumi.Output[dict]
    """
    The Let's Encrypt type certificate configuration

      * `common_name` (`str`)
      * `subjectAlternativeNames` (`list`)
    """
    name: pulumi.Output[str]
    """
    The name of the load-balancer certificate
    """
    not_valid_after: pulumi.Output[str]
    """
    The not valid after validity bound timestamp
    """
    not_valid_before: pulumi.Output[str]
    """
    The not valid before validity bound timestamp
    """
    status: pulumi.Output[str]
    """
    The status of certificate
    """
    subject_alternative_name: pulumi.Output[str]
    """
    The alternative domain names of the certificate
    """
    def __init__(__self__, resource_name, opts=None, custom_certificate=None, lb_id=None, letsencrypt=None, name=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a LbCertificateBeta resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] custom_certificate: The custom type certificate type configuration
        :param pulumi.Input[str] lb_id: The load-balancer ID
        :param pulumi.Input[dict] letsencrypt: The Let's Encrypt type certificate configuration
        :param pulumi.Input[str] name: The name of the load-balancer certificate

        The **custom_certificate** object supports the following:

          * `certificateChain` (`pulumi.Input[str]`)

        The **letsencrypt** object supports the following:

          * `common_name` (`pulumi.Input[str]`)
          * `subjectAlternativeNames` (`pulumi.Input[list]`)
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

            __props__['custom_certificate'] = custom_certificate
            if lb_id is None:
                raise TypeError("Missing required property 'lb_id'")
            __props__['lb_id'] = lb_id
            __props__['letsencrypt'] = letsencrypt
            __props__['name'] = name
            __props__['common_name'] = None
            __props__['fingerprint'] = None
            __props__['not_valid_after'] = None
            __props__['not_valid_before'] = None
            __props__['status'] = None
            __props__['subject_alternative_name'] = None
        super(LbCertificateBeta, __self__).__init__(
            'scaleway:index/lbCertificateBeta:LbCertificateBeta',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, common_name=None, custom_certificate=None, fingerprint=None, lb_id=None, letsencrypt=None, name=None, not_valid_after=None, not_valid_before=None, status=None, subject_alternative_name=None):
        """
        Get an existing LbCertificateBeta resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] common_name: The main domain name of the certificate
        :param pulumi.Input[dict] custom_certificate: The custom type certificate type configuration
        :param pulumi.Input[str] fingerprint: The identifier (SHA-1) of the certificate
        :param pulumi.Input[str] lb_id: The load-balancer ID
        :param pulumi.Input[dict] letsencrypt: The Let's Encrypt type certificate configuration
        :param pulumi.Input[str] name: The name of the load-balancer certificate
        :param pulumi.Input[str] not_valid_after: The not valid after validity bound timestamp
        :param pulumi.Input[str] not_valid_before: The not valid before validity bound timestamp
        :param pulumi.Input[str] status: The status of certificate
        :param pulumi.Input[str] subject_alternative_name: The alternative domain names of the certificate

        The **custom_certificate** object supports the following:

          * `certificateChain` (`pulumi.Input[str]`)

        The **letsencrypt** object supports the following:

          * `common_name` (`pulumi.Input[str]`)
          * `subjectAlternativeNames` (`pulumi.Input[list]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["common_name"] = common_name
        __props__["custom_certificate"] = custom_certificate
        __props__["fingerprint"] = fingerprint
        __props__["lb_id"] = lb_id
        __props__["letsencrypt"] = letsencrypt
        __props__["name"] = name
        __props__["not_valid_after"] = not_valid_after
        __props__["not_valid_before"] = not_valid_before
        __props__["status"] = status
        __props__["subject_alternative_name"] = subject_alternative_name
        return LbCertificateBeta(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
