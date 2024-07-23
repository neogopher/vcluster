package generic

import (
	"github.com/loft-sh/vcluster/pkg/syncer/synccontext"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func WithRecorder(mapper synccontext.Mapper) synccontext.Mapper {
	return &recorder{
		Mapper: mapper,
	}
}

type recorder struct {
	synccontext.Mapper
}

func (n *recorder) VirtualToHost(ctx *synccontext.SyncContext, req types.NamespacedName, vObj client.Object) (retName types.NamespacedName) {
	defer func() {
		RecordMapping(ctx, retName, req, n.GroupVersionKind())
	}()

	// check store first
	pName, ok := VirtualToHostFromStore(ctx, req, n.GroupVersionKind())
	if ok {
		return pName
	}

	return n.Mapper.VirtualToHost(ctx, req, vObj)
}

func (n *recorder) HostToVirtual(ctx *synccontext.SyncContext, req types.NamespacedName, pObj client.Object) (retName types.NamespacedName) {
	defer func() {
		RecordMapping(ctx, req, retName, n.GroupVersionKind())
	}()

	// check store first
	vName, ok := HostToVirtualFromStore(ctx, req, n.GroupVersionKind())
	if ok {
		return vName
	}

	return n.Mapper.HostToVirtual(ctx, req, pObj)
}

func (n *recorder) IsManaged(ctx *synccontext.SyncContext, pObj client.Object) (bool, error) {
	if ctx != nil && ctx.Mappings != nil && ctx.Mappings.Store() != nil {
		_, ok := ctx.Mappings.Store().HostToVirtualName(ctx, synccontext.Object{
			GroupVersionKind: n.GroupVersionKind(),
			NamespacedName: types.NamespacedName{
				Name:      pObj.GetName(),
				Namespace: pObj.GetNamespace(),
			},
		})
		if ok {
			return true, nil
		}
	}

	return n.Mapper.IsManaged(ctx, pObj)
}

func RecordMapping(ctx *synccontext.SyncContext, pName, vName types.NamespacedName, gvk schema.GroupVersionKind) {
	if pName.Name == "" || vName.Name == "" {
		return
	}

	if ctx != nil && ctx.Mappings != nil && ctx.Mappings.Store() != nil {
		// check if we have the owning object in the context
		belongsTo, ok := synccontext.MappingFrom(ctx)
		if !ok {
			return
		}

		// record the reference
		err := ctx.Mappings.Store().RecordReference(ctx, synccontext.NameMapping{
			GroupVersionKind: gvk,

			HostName:    pName,
			VirtualName: vName,
		}, belongsTo)
		if err != nil {
			klog.FromContext(ctx).Error(err, "record name mapping", "host", pName, "virtual", vName)
		}
	}
}

func HostToVirtualFromStore(ctx *synccontext.SyncContext, req types.NamespacedName, gvk schema.GroupVersionKind) (types.NamespacedName, bool) {
	if ctx != nil && ctx.Mappings != nil && ctx.Mappings.Store() != nil {
		return ctx.Mappings.Store().HostToVirtualName(ctx, synccontext.Object{
			GroupVersionKind: gvk,
			NamespacedName:   req,
		})
	}

	return types.NamespacedName{}, false
}

func VirtualToHostFromStore(ctx *synccontext.SyncContext, req types.NamespacedName, gvk schema.GroupVersionKind) (types.NamespacedName, bool) {
	if ctx != nil && ctx.Mappings != nil && ctx.Mappings.Store() != nil {
		return ctx.Mappings.Store().VirtualToHostName(ctx, synccontext.Object{
			GroupVersionKind: gvk,
			NamespacedName:   req,
		})
	}

	return types.NamespacedName{}, false
}
