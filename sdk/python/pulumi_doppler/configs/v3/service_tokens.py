# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ._enums import *

__all__ = ['ServiceTokensArgs', 'ServiceTokens']

@pulumi.input_type
class ServiceTokensArgs:
    def __init__(__self__, *,
                 config: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 access: Optional[pulumi.Input['ServiceTokensAccess']] = None,
                 expire_at: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServiceTokens resource.
        :param pulumi.Input[str] config: Name of the config object.
        :param pulumi.Input[str] project: Unique identifier for the project object.
        :param pulumi.Input['ServiceTokensAccess'] access: Token's capabilities.
        :param pulumi.Input[str] expire_at: Unix timestamp of when token should expire.
        :param pulumi.Input[str] name: Name of the service token.
        """
        if config is None:
            config = 'CONFIG_NAME'
        pulumi.set(__self__, "config", config)
        if project is None:
            project = 'PROJECT_NAME'
        pulumi.set(__self__, "project", project)
        if access is None:
            access = 'read'
        if access is not None:
            pulumi.set(__self__, "access", access)
        if expire_at is not None:
            pulumi.set(__self__, "expire_at", expire_at)
        if name is None:
            name = 'TOKEN_NAME'
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def config(self) -> pulumi.Input[str]:
        """
        Name of the config object.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: pulumi.Input[str]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        Unique identifier for the project object.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def access(self) -> Optional[pulumi.Input['ServiceTokensAccess']]:
        """
        Token's capabilities.
        """
        return pulumi.get(self, "access")

    @access.setter
    def access(self, value: Optional[pulumi.Input['ServiceTokensAccess']]):
        pulumi.set(self, "access", value)

    @property
    @pulumi.getter(name="expireAt")
    def expire_at(self) -> Optional[pulumi.Input[str]]:
        """
        Unix timestamp of when token should expire.
        """
        return pulumi.get(self, "expire_at")

    @expire_at.setter
    def expire_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expire_at", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the service token.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class ServiceTokens(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access: Optional[pulumi.Input['ServiceTokensAccess']] = None,
                 config: Optional[pulumi.Input[str]] = None,
                 expire_at: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a ServiceTokens resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input['ServiceTokensAccess'] access: Token's capabilities.
        :param pulumi.Input[str] config: Name of the config object.
        :param pulumi.Input[str] expire_at: Unix timestamp of when token should expire.
        :param pulumi.Input[str] name: Name of the service token.
        :param pulumi.Input[str] project: Unique identifier for the project object.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceTokensArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ServiceTokens resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ServiceTokensArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceTokensArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access: Optional[pulumi.Input['ServiceTokensAccess']] = None,
                 config: Optional[pulumi.Input[str]] = None,
                 expire_at: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceTokensArgs.__new__(ServiceTokensArgs)

            if access is None:
                access = 'read'
            __props__.__dict__["access"] = access
            if config is None:
                config = 'CONFIG_NAME'
            if config is None and not opts.urn:
                raise TypeError("Missing required property 'config'")
            __props__.__dict__["config"] = config
            __props__.__dict__["expire_at"] = expire_at
            if name is None:
                name = 'TOKEN_NAME'
            __props__.__dict__["name"] = name
            if project is None:
                project = 'PROJECT_NAME'
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["token"] = None
        super(ServiceTokens, __self__).__init__(
            'doppler-native:configs/v3:ServiceTokens',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ServiceTokens':
        """
        Get an existing ServiceTokens resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ServiceTokensArgs.__new__(ServiceTokensArgs)

        __props__.__dict__["access"] = None
        __props__.__dict__["config"] = None
        __props__.__dict__["expire_at"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["token"] = None
        return ServiceTokens(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def access(self) -> pulumi.Output[Optional['ServiceTokensAccess']]:
        """
        Token's capabilities.
        """
        return pulumi.get(self, "access")

    @property
    @pulumi.getter
    def config(self) -> pulumi.Output[str]:
        """
        Name of the config object.
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter(name="expireAt")
    def expire_at(self) -> pulumi.Output[Optional[str]]:
        """
        Unix timestamp of when token should expire.
        """
        return pulumi.get(self, "expire_at")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the service token.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        Unique identifier for the project object.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[Optional['outputs.TokenProperties']]:
        return pulumi.get(self, "token")

