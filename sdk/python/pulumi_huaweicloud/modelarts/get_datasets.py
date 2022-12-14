# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetDatasetsResult',
    'AwaitableGetDatasetsResult',
    'get_datasets',
    'get_datasets_output',
]

@pulumi.output_type
class GetDatasetsResult:
    """
    A collection of values returned by getDatasets.
    """
    def __init__(__self__, datasets=None, id=None, name=None, region=None, type=None):
        if datasets and not isinstance(datasets, list):
            raise TypeError("Expected argument 'datasets' to be a list")
        pulumi.set(__self__, "datasets", datasets)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if type and not isinstance(type, int):
            raise TypeError("Expected argument 'type' to be a int")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def datasets(self) -> Sequence['outputs.GetDatasetsDatasetResult']:
        """
        Indicates a list of all datasets found. Structure is documented below.
        """
        return pulumi.get(self, "datasets")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of label.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def type(self) -> Optional[int]:
        """
        The field type. Valid values include: `String`, `Short`, `Int`, `Long`, `Double`, `Float`, `Byte`,
        `Date`, `Timestamp`, `Bool`.
        """
        return pulumi.get(self, "type")


class AwaitableGetDatasetsResult(GetDatasetsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatasetsResult(
            datasets=self.datasets,
            id=self.id,
            name=self.name,
            region=self.region,
            type=self.type)


def get_datasets(name: Optional[str] = None,
                 region: Optional[str] = None,
                 type: Optional[int] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatasetsResult:
    """
    Use this data source to get a list of ModelArts datasets.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    test = huaweicloud.ModelArts.get_datasets(name="your_dataset_name",
        type=1)
    ```


    :param str name: Specifies the name of datasets.
    :param str region: Specifies the region in which to obtain datasets. If omitted, the provider-level region
           will be used.
    :param int type: Specifies the type of datasets. The options are:
           + **0**: Image classification, supported formats: `.jpg`, `.png`, `.jpeg`, `.bmp`.
           + **1**: Object detection, supported formats: `.jpg`, `.png`, `.jpeg`, `.bmp`.
           + **3**: Image segmentation, supported formats: `.jpg`, `.png`, `.jpeg`, `.bmp`.
           + **100**: Text classification, supported formats: `.txt`, `.csv`.
           + **200**: Sound classification, Supported formats: `.wav`.
           + **400**: Table type, supported formats: Carbon type.
           + **600**: Video, supported formats: `.mp4`
           + **900**: Free format.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['region'] = region
    __args__['type'] = type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:ModelArts/getDatasets:getDatasets', __args__, opts=opts, typ=GetDatasetsResult).value

    return AwaitableGetDatasetsResult(
        datasets=__ret__.datasets,
        id=__ret__.id,
        name=__ret__.name,
        region=__ret__.region,
        type=__ret__.type)


@_utilities.lift_output_func(get_datasets)
def get_datasets_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                        region: Optional[pulumi.Input[Optional[str]]] = None,
                        type: Optional[pulumi.Input[Optional[int]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatasetsResult]:
    """
    Use this data source to get a list of ModelArts datasets.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    test = huaweicloud.ModelArts.get_datasets(name="your_dataset_name",
        type=1)
    ```


    :param str name: Specifies the name of datasets.
    :param str region: Specifies the region in which to obtain datasets. If omitted, the provider-level region
           will be used.
    :param int type: Specifies the type of datasets. The options are:
           + **0**: Image classification, supported formats: `.jpg`, `.png`, `.jpeg`, `.bmp`.
           + **1**: Object detection, supported formats: `.jpg`, `.png`, `.jpeg`, `.bmp`.
           + **3**: Image segmentation, supported formats: `.jpg`, `.png`, `.jpeg`, `.bmp`.
           + **100**: Text classification, supported formats: `.txt`, `.csv`.
           + **200**: Sound classification, Supported formats: `.wav`.
           + **400**: Table type, supported formats: Carbon type.
           + **600**: Video, supported formats: `.mp4`
           + **900**: Free format.
    """
    ...
