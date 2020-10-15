# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables


class Token(pulumi.CustomResource):
    access_key: pulumi.Output[str]
    """
    The access_key.
    """
    creation_ip: pulumi.Output[str]
    """
    The ip used to create the key.
    """
    description: pulumi.Output[str]
    """
    The token description.
    """
    email: pulumi.Output[str]
    """
    The account email. Defaults to registered user.
    """
    expiration_date: pulumi.Output[str]
    """
    The tokens expiration date
    """
    expires: pulumi.Output[bool]
    """
    Defines if the token is set to expire
    """
    password: pulumi.Output[str]
    """
    User password, in case a login is require
    """
    secret_key: pulumi.Output[str]
    """
    The secret_key.
    """
    user_id: pulumi.Output[str]
    """
    The userid of the associated user.
    """
    def __init__(__self__, resource_name, opts=None, description=None, email=None, expires=None, password=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a Token resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The token description.
        :param pulumi.Input[str] email: The account email. Defaults to registered user.
        :param pulumi.Input[bool] expires: Defines if the token is set to expire
        :param pulumi.Input[str] password: User password, in case a login is require
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
            __props__['email'] = email
            __props__['expires'] = expires
            __props__['password'] = password
            __props__['access_key'] = None
            __props__['creation_ip'] = None
            __props__['expiration_date'] = None
            __props__['secret_key'] = None
            __props__['user_id'] = None
        super(Token, __self__).__init__(
            'scaleway:index/token:Token',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, access_key=None, creation_ip=None, description=None, email=None, expiration_date=None, expires=None, password=None, secret_key=None, user_id=None):
        """
        Get an existing Token resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key: The access_key.
        :param pulumi.Input[str] creation_ip: The ip used to create the key.
        :param pulumi.Input[str] description: The token description.
        :param pulumi.Input[str] email: The account email. Defaults to registered user.
        :param pulumi.Input[str] expiration_date: The tokens expiration date
        :param pulumi.Input[bool] expires: Defines if the token is set to expire
        :param pulumi.Input[str] password: User password, in case a login is require
        :param pulumi.Input[str] secret_key: The secret_key.
        :param pulumi.Input[str] user_id: The userid of the associated user.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["access_key"] = access_key
        __props__["creation_ip"] = creation_ip
        __props__["description"] = description
        __props__["email"] = email
        __props__["expiration_date"] = expiration_date
        __props__["expires"] = expires
        __props__["password"] = password
        __props__["secret_key"] = secret_key
        __props__["user_id"] = user_id
        return Token(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
