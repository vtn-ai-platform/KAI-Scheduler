/*
Copyright 2025 NVIDIA CORPORATION
SPDX-License-Identifier: Apache-2.0
*/
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	schedulingv2alpha2 "github.com/NVIDIA/KAI-scheduler/pkg/apis/client/clientset/versioned/typed/scheduling/v2alpha2"
	v2alpha2 "github.com/NVIDIA/KAI-scheduler/pkg/apis/scheduling/v2alpha2"
	gentype "k8s.io/client-go/gentype"
)

// fakePodGroups implements PodGroupInterface
type fakePodGroups struct {
	*gentype.FakeClientWithList[*v2alpha2.PodGroup, *v2alpha2.PodGroupList]
	Fake *FakeSchedulingV2alpha2
}

func newFakePodGroups(fake *FakeSchedulingV2alpha2, namespace string) schedulingv2alpha2.PodGroupInterface {
	return &fakePodGroups{
		gentype.NewFakeClientWithList[*v2alpha2.PodGroup, *v2alpha2.PodGroupList](
			fake.Fake,
			namespace,
			v2alpha2.SchemeGroupVersion.WithResource("podgroups"),
			v2alpha2.SchemeGroupVersion.WithKind("PodGroup"),
			func() *v2alpha2.PodGroup { return &v2alpha2.PodGroup{} },
			func() *v2alpha2.PodGroupList { return &v2alpha2.PodGroupList{} },
			func(dst, src *v2alpha2.PodGroupList) { dst.ListMeta = src.ListMeta },
			func(list *v2alpha2.PodGroupList) []*v2alpha2.PodGroup { return gentype.ToPointerSlice(list.Items) },
			func(list *v2alpha2.PodGroupList, items []*v2alpha2.PodGroup) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
