# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetAccountSSHKeyResult:
    """
    A collection of values returned by getAccountSSHKey.
    """
    def __init__(__self__, id=None, name=None, organization_id=None, public_key=None, ssh_key_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        __self__.organization_id = organization_id
        if public_key and not isinstance(public_key, str):
            raise TypeError("Expected argument 'public_key' to be a str")
        __self__.public_key = public_key
        if ssh_key_id and not isinstance(ssh_key_id, str):
            raise TypeError("Expected argument 'ssh_key_id' to be a str")
        __self__.ssh_key_id = ssh_key_id
class AwaitableGetAccountSSHKeyResult(GetAccountSSHKeyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccountSSHKeyResult(
            id=self.id,
            name=self.name,
            organization_id=self.organization_id,
            public_key=self.public_key,
            ssh_key_id=self.ssh_key_id)

def get_account_ssh_key(name=None,organization_id=None,ssh_key_id=None,opts=None):
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()


    __args__['name'] = name
    __args__['organizationId'] = organization_id
    __args__['sshKeyId'] = ssh_key_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('scaleway:index/getAccountSSHKey:getAccountSSHKey', __args__, opts=opts).value

    return AwaitableGetAccountSSHKeyResult(
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        organization_id=__ret__.get('organizationId'),
        public_key=__ret__.get('publicKey'),
        ssh_key_id=__ret__.get('sshKeyId'))
