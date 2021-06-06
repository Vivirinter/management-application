package wordpress

import (
	v12 "github.com/Vivirinter/management-application/api/v1"
	"github.com/dgrijalva/jwt-go/request"
	"golang.org/x/net/context"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

func WordpressResource(Wordpress *v12.Wordpress) {
	var (
		error = c.Watch(&source.Kind{Type: &v1.Wordpress{}}, &handler.EnqueueRequestForObject{})
	)
	if error != nil {
		return error
	}

	error = c.Watch(&source.Kind{Type: &appsv1.Deployment{}}, &handler.EnqueueRequestForOwner{
		IsController: True,
		OwnerType:    &v1.Wordpress{},
	})
	if error != nil {
		return error
	}
	error = c.Watch(&source.Kind{Type: &corev1.Service{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &v1.Wordpress{},
	})
	if error != nil {
		return error
	}
	error = c.Watch(&source.Kind{Type: &corev1.PersistentVolumeClaim{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &v1.Wordpress{},
	})
	if error != nil {
		return error
	}

	wordpress := &v1.Wordpress{}
	err := r.client.Get(context.TODO(), request.NamespacedName, wordpress) -----1
	if err != nil {
		if errors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	//if everything goes fine
	return reconcile.Result{}, nil

}
