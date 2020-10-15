# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetRegistryNamespaceBetaResult:
    """
    A collection of values returned by getRegistryNamespaceBeta.
    """
    def __init__(__self__, description=None, endpoint=None, id=None, is_public=None, name=None, namespace_id=None, organization_id=None, region=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        if endpoint and not isinstance(endpoint, str):
            raise TypeError("Expected argument 'endpoint' to be a str")
        __self__.endpoint = endpoint
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if is_public and not isinstance(is_public, bool):
            raise TypeError("Expected argument 'is_public' to be a bool")
        __self__.is_public = is_public
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if namespace_id and not isinstance(namespace_id, str):
            raise TypeError("Expected argument 'namespace_id' to be a str")
        __self__.namespace_id = namespace_id
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        __self__.organization_id = organization_id
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
class AwaitableGetRegistryNamespaceBetaResult(GetRegistryNamespaceBetaResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegistryNamespaceBetaResult(
            description=self.description,
            endpoint=self.endpoint,
            id=self.id,
            is_public=self.is_public,
            name=self.name,
            namespace_id=self.namespace_id,
            organization_id=self.organization_id,
            region=self.region)

def get_registry_namespace_beta(name=None,namespace_id=None,region=None,opts=None):
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()


    __args__['name'] = name
    __args__['namespaceId'] = namespace_id
    __args__['region'] = region
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('scaleway:index/getRegistryNamespaceBeta:getRegistryNamespaceBeta', __args__, opts=opts).value

    return AwaitableGetRegistryNamespaceBetaResult(
        description=__ret__.get('description'),
        endpoint=__ret__.get('endpoint'),
        id=__ret__.get('id'),
        is_public=__ret__.get('isPublic'),
        name=__ret__.get('name'),
        namespace_id=__ret__.get('namespaceId'),
        organization_id=__ret__.get('organizationId'),
        region=__ret__.get('region'))
