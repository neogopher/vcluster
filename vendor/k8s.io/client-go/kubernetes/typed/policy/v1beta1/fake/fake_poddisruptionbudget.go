/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "k8s.io/api/policy/v1beta1"
	policyv1beta1 "k8s.io/client-go/applyconfigurations/policy/v1beta1"
	gentype "k8s.io/client-go/gentype"
	typedpolicyv1beta1 "k8s.io/client-go/kubernetes/typed/policy/v1beta1"
)

// fakePodDisruptionBudgets implements PodDisruptionBudgetInterface
type fakePodDisruptionBudgets struct {
	*gentype.FakeClientWithListAndApply[*v1beta1.PodDisruptionBudget, *v1beta1.PodDisruptionBudgetList, *policyv1beta1.PodDisruptionBudgetApplyConfiguration]
	Fake *FakePolicyV1beta1
}

func newFakePodDisruptionBudgets(fake *FakePolicyV1beta1, namespace string) typedpolicyv1beta1.PodDisruptionBudgetInterface {
	return &fakePodDisruptionBudgets{
		gentype.NewFakeClientWithListAndApply[*v1beta1.PodDisruptionBudget, *v1beta1.PodDisruptionBudgetList, *policyv1beta1.PodDisruptionBudgetApplyConfiguration](
			fake.Fake,
			namespace,
			v1beta1.SchemeGroupVersion.WithResource("poddisruptionbudgets"),
			v1beta1.SchemeGroupVersion.WithKind("PodDisruptionBudget"),
			func() *v1beta1.PodDisruptionBudget { return &v1beta1.PodDisruptionBudget{} },
			func() *v1beta1.PodDisruptionBudgetList { return &v1beta1.PodDisruptionBudgetList{} },
			func(dst, src *v1beta1.PodDisruptionBudgetList) { dst.ListMeta = src.ListMeta },
			func(list *v1beta1.PodDisruptionBudgetList) []*v1beta1.PodDisruptionBudget {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1beta1.PodDisruptionBudgetList, items []*v1beta1.PodDisruptionBudget) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
